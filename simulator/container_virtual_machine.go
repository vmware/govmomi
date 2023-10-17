/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package simulator

import (
	"archive/tar"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

const ContainerBackingOptionKey = "RUN.container"

var (
	toolsRunning = []types.PropertyChange{
		{Name: "guest.toolsStatus", Val: types.VirtualMachineToolsStatusToolsOk},
		{Name: "guest.toolsRunningStatus", Val: string(types.VirtualMachineToolsRunningStatusGuestToolsRunning)},
	}

	toolsNotRunning = []types.PropertyChange{
		{Name: "guest.toolsStatus", Val: types.VirtualMachineToolsStatusToolsNotRunning},
		{Name: "guest.toolsRunningStatus", Val: string(types.VirtualMachineToolsRunningStatusGuestToolsNotRunning)},
	}
)

type simVM struct {
	vm *VirtualMachine
	c  *container
}

// createSimulationVM inspects the provided VirtualMachine and creates a simVM binding for it if
// the vm.Config.ExtraConfig set contains a key "RUN.container".
// If the ExtraConfig set does not contain that key, this returns nil.
// Methods on the simVM type are written to check for nil object so the return from this call can be blindly
// assigned and invoked without the caller caring about whether a binding for a backing container was warranted.
func createSimulationVM(vm *VirtualMachine) *simVM {
	svm := &simVM{
		vm: vm,
	}

	for _, opt := range vm.Config.ExtraConfig {
		val := opt.GetOptionValue()
		if val.Key == ContainerBackingOptionKey {
			return svm
		}
	}

	return nil
}

// applies container network settings to vm.Guest properties.
func (svm *simVM) syncNetworkConfigToVMGuestProperties() error {
	if svm == nil {
		return nil
	}

	out, detail, err := svm.c.inspect()
	if err != nil {
		return err
	}

	svm.vm.Config.Annotation = "inspect"
	svm.vm.logPrintf("%s: %s", svm.vm.Config.Annotation, string(out))

	netS := detail.NetworkSettings.networkSettings

	// ? Why is this valid - we're taking the first entry while iterating over a MAP
	for _, n := range detail.NetworkSettings.Networks {
		netS = n
		break
	}

	if detail.State.Paused {
		svm.vm.Runtime.PowerState = types.VirtualMachinePowerStateSuspended
	} else if detail.State.Running {
		svm.vm.Runtime.PowerState = types.VirtualMachinePowerStatePoweredOn
	} else {
		svm.vm.Runtime.PowerState = types.VirtualMachinePowerStatePoweredOff
	}

	svm.vm.Guest.IpAddress = netS.IPAddress
	svm.vm.Summary.Guest.IpAddress = netS.IPAddress

	if len(svm.vm.Guest.Net) != 0 {
		net := &svm.vm.Guest.Net[0]
		net.IpAddress = []string{netS.IPAddress}
		net.MacAddress = netS.MacAddress
		net.IpConfig = &types.NetIpConfigInfo{
			IpAddress: []types.NetIpConfigInfoIpAddress{{
				IpAddress:    netS.IPAddress,
				PrefixLength: int32(netS.IPPrefixLen),
				State:        string(types.NetIpConfigInfoIpAddressStatusPreferred),
			}},
		}
	}

	for _, d := range svm.vm.Config.Hardware.Device {
		if eth, ok := d.(types.BaseVirtualEthernetCard); ok {
			eth.GetVirtualEthernetCard().MacAddress = netS.MacAddress
			break
		}
	}

	return nil
}

func (svm *simVM) prepareGuestOperation(auth types.BaseGuestAuthentication) types.BaseMethodFault {
	if svm == nil || svm.c == nil || svm.c.id == "" {
		return new(types.GuestOperationsUnavailable)
	}

	if svm.vm.Runtime.PowerState != types.VirtualMachinePowerStatePoweredOn {
		return &types.InvalidPowerState{
			RequestedState: types.VirtualMachinePowerStatePoweredOn,
			ExistingState:  svm.vm.Runtime.PowerState,
		}
	}

	switch creds := auth.(type) {
	case *types.NamePasswordAuthentication:
		if creds.Username == "" || creds.Password == "" {
			return new(types.InvalidGuestLogin)
		}
	default:
		return new(types.InvalidGuestLogin)
	}

	return nil
}

// populateDMI writes BIOS UUID DMI files to a container volume
func (svm *simVM) populateDMI() error {
	if svm.c == nil {
		return nil
	}

	files := []tarEntry{
		{
			&tar.Header{
				Name: "product_uuid",
				Mode: 0444,
			},
			[]byte(productUUID(svm.vm.uid)),
		},
		{
			&tar.Header{
				Name: "product_serial",
				Mode: 0444,
			},
			[]byte(productSerial(svm.vm.uid)),
		},
	}

	_, err := svm.c.createVolume("dmi", []string{deleteWithContainer}, files)
	return err
}

// start runs the container if specified by the RUN.container extraConfig property.
// lazily creates a container backing if specified by an ExtraConfig property with key "RUN.container"
func (svm *simVM) start(ctx *Context) error {
	if svm == nil {
		return nil
	}

	if svm.c != nil && svm.c.id != "" {
		err := svm.c.start(ctx)
		if err != nil {
			log.Printf("%s %s: %s", svm.vm.Name, "start", err)
		} else {
			ctx.Map.Update(svm.vm, toolsRunning)
		}

		return err
	}

	var args []string
	var env []string
	var ports []string
	mountDMI := true

	for _, opt := range svm.vm.Config.ExtraConfig {
		val := opt.GetOptionValue()
		if val.Key == ContainerBackingOptionKey {
			run := val.Value.(string)
			err := json.Unmarshal([]byte(run), &args)
			if err != nil {
				args = []string{run}
			}

			continue
		}

		if val.Key == "RUN.mountdmi" {
			var mount bool
			err := json.Unmarshal([]byte(val.Value.(string)), &mount)
			if err == nil {
				mountDMI = mount
			}

			continue
		}

		if strings.HasPrefix(val.Key, "RUN.port.") {
			// ? would this not make more sense as a set of tuples in the value?
			// or inlined into the RUN.container freeform string as is the case with the nginx volume in the examples?
			sKey := strings.Split(val.Key, ".")
			containerPort := sKey[len(sKey)-1]
			ports = append(ports, fmt.Sprintf("%s:%s", val.Value.(string), containerPort))

			continue
		}

		if strings.HasPrefix(val.Key, "RUN.env.") {
			sKey := strings.Split(val.Key, ".")
			envKey := sKey[len(sKey)-1]
			env = append(env, fmt.Sprintf("%s=%s", envKey, val.Value.(string)))
		}

		if strings.HasPrefix(val.Key, "guestinfo.") {
			key := strings.Replace(strings.ToUpper(val.Key), ".", "_", -1)
			env = append(env, fmt.Sprintf("VMX_%s=%s", key, val.Value.(string)))

			continue
		}
	}

	if len(args) == 0 {
		// not an error - it's simply a simVM that shouldn't be backed by a container
		return nil
	}

	if len(env) != 0 {
		// Configure env as the data access method for cloud-init-vmware-guestinfo
		env = append(env, "VMX_GUESTINFO=true")
	}

	volumes := []string{}
	if mountDMI {
		volumes = append(volumes, constructVolumeName(svm.vm.Name, svm.vm.uid.String(), "dmi")+":/sys/class/dmi/id")
	}

	var err error
	svm.c, err = create(ctx, svm.vm.Name, svm.vm.uid.String(), nil, volumes, ports, env, args[0], args[1:])
	if err != nil {
		return err
	}

	if mountDMI {
		// not combined with the test assembling volumes because we want to have the container name first.
		// cannot add a label to a volume after creation, so if we want to associate with the container ID the
		// container must come first
		err = svm.populateDMI()
		if err != nil {
			return err
		}
	}

	err = svm.c.start(ctx)
	if err != nil {
		log.Printf("%s %s: %s %s", svm.vm.Name, "start", args, err)
		return err
	}

	ctx.Map.Update(svm.vm, toolsRunning)

	svm.vm.logPrintf("%s: %s", args, svm.c.id)

	if err = svm.syncNetworkConfigToVMGuestProperties(); err != nil {
		log.Printf("%s inspect %s: %s", svm.vm.Name, svm.c.id, err)
	}

	callback := func(details *containerDetails, c *container) error {
		spoofctx := SpoofContext()

		if c.id == "" && svm.vm != nil {
			// If the container cannot be found then destroy this VM unless the VM is no longer configured for container backing (svm.vm == nil)
			taskRef := svm.vm.DestroyTask(spoofctx, &types.Destroy_Task{This: svm.vm.Self}).(*methods.Destroy_TaskBody).Res.Returnval
			task, ok := spoofctx.Map.Get(taskRef).(*Task)
			if !ok {
				panic(fmt.Sprintf("couldn't retrieve task for moref %+q while deleting VM %s", taskRef, svm.vm.Name))
			}

			// Wait for the task to complete and see if there is an error.
			task.Wait()
			if task.Info.Error != nil {
				msg := fmt.Sprintf("failed to destroy vm: err=%v", *task.Info.Error)
				svm.vm.logPrintf(msg)

				return errors.New(msg)
			}
		}

		return svm.syncNetworkConfigToVMGuestProperties()
	}

	// Start watching the container resource.
	err = svm.c.watchContainer(context.Background(), callback)
	if _, ok := err.(uninitializedContainer); ok {
		// the container has been deleted before we could watch, despite successful launch so clean up.
		callback(nil, svm.c)

		// successful launch so nil the error
		return nil
	}

	return err
}

// stop the container (if any) for the given vm.
func (svm *simVM) stop(ctx *Context) error {
	if svm == nil || svm.c == nil {
		return nil
	}

	err := svm.c.stop(ctx)
	if err != nil {
		log.Printf("%s %s: %s", svm.vm.Name, "stop", err)

		return err
	}

	ctx.Map.Update(svm.vm, toolsNotRunning)

	return nil
}

// pause the container (if any) for the given vm.
func (svm *simVM) pause(ctx *Context) error {
	if svm == nil || svm.c == nil {
		return nil
	}

	err := svm.c.pause(ctx)
	if err != nil {
		log.Printf("%s %s: %s", svm.vm.Name, "pause", err)

		return err
	}

	ctx.Map.Update(svm.vm, toolsNotRunning)

	return nil
}

// restart the container (if any) for the given vm.
func (svm *simVM) restart(ctx *Context) error {
	if svm == nil || svm.c == nil {
		return nil
	}

	err := svm.c.restart(ctx)
	if err != nil {
		log.Printf("%s %s: %s", svm.vm.Name, "restart", err)

		return err
	}

	ctx.Map.Update(svm.vm, toolsRunning)

	return nil
}

// remove the container (if any) for the given vm.
func (svm *simVM) remove(ctx *Context) error {
	if svm == nil || svm.c == nil {
		return nil
	}

	err := svm.c.remove(ctx)
	if err != nil {
		log.Printf("%s %s: %s", svm.vm.Name, "remove", err)

		return err
	}

	return nil
}

func (svm *simVM) exec(ctx *Context, auth types.BaseGuestAuthentication, args []string) (string, types.BaseMethodFault) {
	if svm == nil || svm.c == nil {
		return "", nil
	}

	fault := svm.prepareGuestOperation(auth)
	if fault != nil {
		return "", fault
	}

	out, err := svm.c.exec(ctx, args)
	if err != nil {
		log.Printf("%s: %s (%s)", svm.vm.Name, args, string(out))
		return "", new(types.GuestOperationsFault)
	}

	return strings.TrimSpace(string(out)), nil
}

func guestUpload(id string, file string, r *http.Request) error {
	// TODO: decide behaviour for no container
	err := copyToGuest(id, file, r.ContentLength, r.Body)
	_ = r.Body.Close()
	return err
}

func guestDownload(id string, file string, w http.ResponseWriter) error {
	// TODO: decide behaviour for no container
	sink := func(len int64, r io.Reader) error {
		w.Header().Set("Content-Length", strconv.FormatInt(len, 10))
		_, err := io.Copy(w, r)
		return err
	}

	err := copyFromGuest(id, file, sink)
	return err
}

const guestPrefix = "/guestFile/"

// ServeGuest handles container guest file upload/download
func ServeGuest(w http.ResponseWriter, r *http.Request) {
	// Real vCenter form: /guestFile?id=139&token=...
	// vcsim form:        /guestFile/tmp/foo/bar?id=ebc8837b8cb6&token=...

	id := r.URL.Query().Get("id")
	file := strings.TrimPrefix(r.URL.Path, guestPrefix[:len(guestPrefix)-1])
	var err error

	switch r.Method {
	case http.MethodPut:
		err = guestUpload(id, file, r)
	case http.MethodGet:
		err = guestDownload(id, file, w)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err != nil {
		log.Printf("%s %s: %s", r.Method, r.URL, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// productSerial returns the uuid in /sys/class/dmi/id/product_serial format
func productSerial(id uuid.UUID) string {
	var dst [len(id)*2 + len(id) - 1]byte

	j := 0
	for i := 0; i < len(id); i++ {
		hex.Encode(dst[j:j+2], id[i:i+1])
		j += 3
		if j < len(dst) {
			s := j - 1
			if s == len(dst)/2 {
				dst[s] = '-'
			} else {
				dst[s] = ' '
			}
		}
	}

	return fmt.Sprintf("VMware-%s", string(dst[:]))
}

// productUUID returns the uuid in /sys/class/dmi/id/product_uuid format
func productUUID(id uuid.UUID) string {
	var dst [36]byte

	hex.Encode(dst[0:2], id[3:4])
	hex.Encode(dst[2:4], id[2:3])
	hex.Encode(dst[4:6], id[1:2])
	hex.Encode(dst[6:8], id[0:1])
	dst[8] = '-'
	hex.Encode(dst[9:11], id[5:6])
	hex.Encode(dst[11:13], id[4:5])
	dst[13] = '-'
	hex.Encode(dst[14:16], id[7:8])
	hex.Encode(dst[16:18], id[6:7])
	dst[18] = '-'
	hex.Encode(dst[19:23], id[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], id[10:])

	return strings.ToUpper(string(dst[:]))
}
