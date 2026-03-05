// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// Container Backing for Virtual Machines
//
// This file implements container-backed VMs for vcsim. When a VM is created with
// the ExtraConfig key "RUN.container" set to a container image, vcsim will create
// a real container to back the simulated VM.
//
// # ExtraConfig Options
//
// The following ExtraConfig keys control container backing behavior:
//
//   - RUN.container: Container image to use (e.g., "nginx:latest", "alpine:3.20")
//     Can be a simple image name or JSON array with command args.
//
//   - RUN.mountdmi: Boolean (default: true). Mount /sys/class/dmi/id for DMI info.
//     Set to "false" for rootless containers that can't mount system paths.
//
//   - RUN.network: Container network to use (e.g., "podman", "bridge").
//     Important for rootless podman which needs explicit network for IP assignment.
//
//   - RUN.nestedContainers: Boolean (default: false). Enable nested container mode
//     for running Kubernetes or other container workloads inside the container.
//     When true, adds kind-style flags:
//     --cgroupns=private, --security-opt seccomp=unconfined,
//     --security-opt apparmor=unconfined, --tmpfs /tmp, --tmpfs /run,
//     --volume /var, --volume /lib/modules:/lib/modules:ro, --device /dev/fuse.
//     Reference: https://github.com/kubernetes-sigs/kind/blob/main/pkg/cluster/internal/providers/docker/provision.go
//
//   - RUN.port.<containerPort>: Map container port to host port.
//     Example: RUN.port.80 = "8080" maps container port 80 to host port 8080.
//
//   - RUN.env.<name>: Set environment variable in the container.
//     Example: RUN.env.DEBUG = "true" sets DEBUG=true in the container.
//
//   - guestinfo.*: Passed as VMX_GUESTINFO_* environment variables.
//     Used by cloud-init VMware datasource with EnvVar transport.
//
// # Example: Basic Container
//
//	spec := types.VirtualMachineConfigSpec{
//		Name: "web-server",
//		ExtraConfig: []types.BaseOptionValue{
//			&types.OptionValue{Key: "RUN.container", Value: "nginx:latest"},
//			&types.OptionValue{Key: "RUN.port.80", Value: "8080"},
//		},
//	}
//
// # Example: Kubernetes Node (Nested Containers)
//
//	spec := types.VirtualMachineConfigSpec{
//		Name: "k8s-node",
//		ExtraConfig: []types.BaseOptionValue{
//			&types.OptionValue{Key: "RUN.container", Value: "my-k8s-image:latest"},
//			&types.OptionValue{Key: "RUN.mountdmi", Value: "false"},
//			&types.OptionValue{Key: "RUN.nestedContainers", Value: "true"},
//			&types.OptionValue{Key: "guestinfo.metadata", Value: metadataEncoded},
//			&types.OptionValue{Key: "guestinfo.userdata", Value: userdataEncoded},
//		},
//	}

package simulator

import (
	"archive/tar"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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
// If ctx is provided, property change notifications will be triggered for all modified properties.
// Uses PropertyDiff to generate granular property changes.
func (svm *simVM) syncNetworkConfigToVMGuestProperties(ctx *Context) error {
	if svm == nil {
		return nil
	}

	out, detail, err := svm.c.inspect()
	if err != nil {
		return err
	}

	svm.vm.Config.Annotation = "inspect"
	svm.vm.logPrintf("%s: %s", svm.vm.Config.Annotation, string(out))

	// Create a checkpoint of the current VM state before modifications
	checkpoint := Checkpoint(&svm.vm.VirtualMachine)

	// Update power state based on container state
	if detail.State.Paused {
		svm.vm.Runtime.PowerState = types.VirtualMachinePowerStateSuspended
	} else if detail.State.Running {
		svm.vm.Runtime.PowerState = types.VirtualMachinePowerStatePoweredOn
	} else {
		svm.vm.Runtime.PowerState = types.VirtualMachinePowerStatePoweredOff
	}

	// Get the primary network settings (first network or default)
	primaryNet := detail.NetworkSettings.networkSettings
	for _, n := range detail.NetworkSettings.Networks {
		primaryNet = n
		break
	}

	// Update primary IP address
	svm.vm.Guest.IpAddress = primaryNet.IPAddress
	svm.vm.Summary.Guest.IpAddress = primaryNet.IPAddress

	// Update hostname
	if detail.Config.Hostname != "" {
		svm.vm.Guest.HostName = detail.Config.Hostname
		svm.vm.Summary.Guest.HostName = detail.Config.Hostname
	}

	// Build Guest.Net from all container networks
	var guestNics []types.GuestNicInfo
	nicIndex := int32(0)
	for networkName, netSettings := range detail.NetworkSettings.Networks {
		if netSettings.IPAddress == "" {
			continue
		}

		nic := types.GuestNicInfo{
			Network:      networkName,
			IpAddress:    []string{netSettings.IPAddress},
			MacAddress:   netSettings.MacAddress,
			Connected:    true,
			DeviceConfigId: nicIndex,
			IpConfig: &types.NetIpConfigInfo{
				IpAddress: []types.NetIpConfigInfoIpAddress{{
					IpAddress:    netSettings.IPAddress,
					PrefixLength: int32(netSettings.IPPrefixLen),
					State:        string(types.NetIpConfigInfoIpAddressStatusPreferred),
				}},
			},
		}
		guestNics = append(guestNics, nic)
		nicIndex++
	}

	// If no networks found in the Networks map, use the default network settings
	if len(guestNics) == 0 && primaryNet.IPAddress != "" {
		nic := types.GuestNicInfo{
			Network:      "default",
			IpAddress:    []string{primaryNet.IPAddress},
			MacAddress:   primaryNet.MacAddress,
			Connected:    true,
			DeviceConfigId: 0,
			IpConfig: &types.NetIpConfigInfo{
				IpAddress: []types.NetIpConfigInfoIpAddress{{
					IpAddress:    primaryNet.IPAddress,
					PrefixLength: int32(primaryNet.IPPrefixLen),
					State:        string(types.NetIpConfigInfoIpAddressStatusPreferred),
				}},
			},
		}
		guestNics = append(guestNics, nic)
	}

	svm.vm.Guest.Net = guestNics

	// Build IP stack info with DNS and routing
	if primaryNet.IPAddress != "" {
		gsi := types.GuestStackInfo{
			DnsConfig: &types.NetDnsConfigInfo{
				Dhcp:         false,
				HostName:     svm.vm.Guest.HostName,
				DomainName:   detail.Config.Domainname,
				IpAddress:    detail.Config.DNS,
				SearchDomain: nil,
			},
			IpRouteConfig: &types.NetIpRouteConfigInfo{
				IpRoute: []types.NetIpRouteConfigInfoIpRoute{{
					Network:      "0.0.0.0",
					PrefixLength: 0,
					Gateway: types.NetIpRouteConfigInfoGateway{
						IpAddress: primaryNet.Gateway,
						Device:    "0",
					},
				}},
			},
		}
		svm.vm.Guest.IpStack = []types.GuestStackInfo{gsi}
	}

	// Update MAC address on virtual ethernet card to match primary network
	for _, d := range svm.vm.Config.Hardware.Device {
		if eth, ok := d.(types.BaseVirtualEthernetCard); ok {
			eth.GetVirtualEthernetCard().MacAddress = primaryNet.MacAddress
			break
		}
	}

	// Use PropertyDiff to generate granular property changes
	if ctx != nil {
		changes := PropertyDiff(checkpoint, &svm.vm.VirtualMachine)
		if len(changes) > 0 {
			ctx.Update(svm.vm, changes)
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
			ctx.Update(svm.vm, toolsRunning)
		}

		return err
	}

	var args []string
	var env []string
	var ports []string
	var networks []string
	mountDMI := true
	nestedContainers := false

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

		if val.Key == "RUN.nestedContainers" {
			// Enable nested container mode for running Kubernetes or other container
			// workloads inside the container. This adds flags adapted from kind's
			// provision.go including --cgroupns=private, security-opt unconfined,
			// tmpfs mounts, and /dev/fuse device.
			var nested bool
			err := json.Unmarshal([]byte(val.Value.(string)), &nested)
			if err == nil {
				nestedContainers = nested
			}

			continue
		}

		if val.Key == "RUN.network" {
			// Specify container network (e.g., "podman" for rootless podman bridge network)
			// This is important for rootless podman which doesn't assign IPs without a network
			networks = append(networks, val.Value.(string))
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
	svm.c, err = create(ctx, svm.vm.Name, svm.vm.uid.String(), networks, volumes, ports, env, nestedContainers, args[0], args[1:])
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

	ctx.Update(svm.vm, toolsRunning)

	svm.vm.logPrintf("%s: %s", args, svm.c.id)

	// Sync network config, retrying a few times to allow the container to get an IP
	// Container runtimes may take a moment to assign an IP address after start
	for i := 0; i < 5; i++ {
		if err = svm.syncNetworkConfigToVMGuestProperties(ctx); err != nil {
			log.Printf("%s inspect %s: %s", svm.vm.Name, svm.c.id, err)
			break
		}
		if svm.vm.Guest.IpAddress != "" {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	callback := func(callbackCtx *Context, details *containerDetails, c *container) error {
		if c.id == "" && svm.vm != nil {
			// If the container cannot be found then destroy this VM unless the VM is no longer configured for container backing (svm.vm == nil)
			taskRef := svm.vm.DestroyTask(callbackCtx, &types.Destroy_Task{This: svm.vm.Self}).(*methods.Destroy_TaskBody).Res.Returnval
			task, ok := callbackCtx.Map.Get(taskRef).(*Task)
			if !ok {
				panic(fmt.Sprintf("couldn't retrieve task for moref %+q while deleting VM %s", taskRef, svm.vm.Name))
			}

			// Wait for the task to complete and see if there is an error.
			task.Wait()
			if task.Info.Error != nil {
				err := fmt.Errorf("failed to destroy vm: %v", task.Info.Error)
				svm.vm.logPrintf("%s", err.Error())
				return err
			}
		}

		return svm.syncNetworkConfigToVMGuestProperties(callbackCtx)
	}

	// Start watching the container resource.
	err = svm.c.watchContainer(ctx, callback)
	if _, ok := err.(uninitializedContainer); ok {
		// the container has been deleted before we could watch, despite successful launch so clean up.
		callback(ctx, nil, svm.c)

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

	ctx.Update(svm.vm, toolsNotRunning)

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

	ctx.Update(svm.vm, toolsNotRunning)

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

	ctx.Update(svm.vm, toolsRunning)

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
