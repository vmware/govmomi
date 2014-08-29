/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package importx

import (
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"sync"
	"time"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type ovf struct {
	*flags.DatastoreFlag
	*flags.ResourcePoolFlag
	*flags.SearchFlag
	*flags.OutputFlag

	Client       *govmomi.Client
	Datacenter   *govmomi.Datacenter
	Datastore    *govmomi.Datastore
	ResourcePool *govmomi.ResourcePool

	Archive
}

func init() {
	cli.Register("import.ovf", newOvf())
}

func newOvf() *ovf {
	return &ovf{
		SearchFlag: flags.NewSearchFlag(flags.SearchHosts),
	}
}

func (cmd *ovf) Register(f *flag.FlagSet) {
}

func (cmd *ovf) Process() error { return nil }

func (cmd *ovf) Run(f *flag.FlagSet) error {
	file, err := cmd.Prepare(f)

	if err != nil {
		return err
	}

	cmd.Archive = &FileArchive{file}

	return cmd.Import(file)
}

func (cmd *ovf) Prepare(f *flag.FlagSet) (importable, error) {
	var err error

	args := f.Args()
	if len(args) != 1 {
		return "", errors.New("no file to import")
	}

	file := importable(f.Arg(0))

	cmd.Client, err = cmd.DatastoreFlag.Client()
	if err != nil {
		return "", err
	}

	cmd.Datacenter, err = cmd.DatastoreFlag.Datacenter()
	if err != nil {
		return "", err
	}

	cmd.Datastore, err = cmd.DatastoreFlag.Datastore()
	if err != nil {
		return "", err
	}

	cmd.ResourcePool, err = cmd.ResourcePoolFlag.ResourcePool()
	if err != nil {
		return "", err
	}

	return file, nil
}

func (cmd *ovf) ReadAll(name string) ([]byte, error) {
	f, _, err := cmd.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

func (cmd *ovf) Import(i importable) error {
	c := cmd.Client

	desc, err := cmd.ReadAll(string(i))
	if err != nil {
		return err
	}

	// extract name from .ovf for use as VM name
	ovf := struct {
		VirtualSystem struct {
			Name string
		}
	}{}

	if err := xml.Unmarshal(desc, &ovf); err != nil {
		return fmt.Errorf("failed to parse ovf: %s", err.Error())
	}

	cisp := types.OvfCreateImportSpecParams{
		EntityName: ovf.VirtualSystem.Name,
		OvfManagerCommonParams: types.OvfManagerCommonParams{
			Locale: "US",
		},
	}

	spec, err := c.OvfManager().CreateImportSpec(string(desc), cmd.ResourcePool, cmd.Datastore, cisp)
	if err != nil {
		return err
	}

	if spec.Error != nil {
		return errors.New(spec.Error[0].LocalizedMessage)
	}

	if spec.Warning != nil {
		for _, w := range spec.Warning {
			_, _ = cmd.Log(fmt.Sprintf("Warning: %s\n", w.LocalizedMessage))
		}
	}

	// TODO: ImportSpec may have unitNumber==0, but this field is optional in the wsdl
	// and hence omitempty in the struct tag; but unitNumber is required for certain devices.
	s := &spec.ImportSpec.(*types.VirtualMachineImportSpec).ConfigSpec
	for _, d := range s.DeviceChange {
		n := &d.GetVirtualDeviceConfigSpec().Device.GetVirtualDevice().UnitNumber
		if *n == 0 {
			*n = -1
		}
	}

	var host *govmomi.HostSystem
	if cmd.SearchFlag.IsSet() {
		if host, err = cmd.HostSystem(); err != nil {
			return err
		}
	}

	// TODO: need a folder option
	folders, err := cmd.Datacenter.Folders(c)
	if err != nil {
		return err
	}
	folder := &folders.VmFolder

	lease, err := cmd.ResourcePool.ImportVApp(c, spec.ImportSpec, folder, host)
	if err != nil {
		return err
	}

	info, err := lease.Wait(c)
	if err != nil {
		return err
	}

	for _, device := range info.DeviceUrl {
		for _, item := range spec.FileItem {
			if device.ImportKey != item.DeviceId {
				continue
			}

			u, err := c.Client.ParseURL(device.Url)
			if err != nil {
				return err
			}

			err = cmd.Upload(lease, &item, u)
			if err != nil {
				return err
			}
		}
	}

	return lease.HttpNfcLeaseComplete(c)
}

// LeaseUpdater consumes an Upload.Progress channel (in) used to update HttpNfcLeaseProgress.
// Progress is forwarded to another channel (out), which can in turn be consumed by the ProgressLogger.
func (cmd *ovf) LeaseUpdater(lease *govmomi.HttpNfcLease, in <-chan vim25.Progress, out chan<- vim25.Progress) *sync.WaitGroup {
	var wg sync.WaitGroup

	go func() {
		var p vim25.Progress
		var ok bool
		var err error
		var percent int

		tick := time.NewTicker(2 * time.Second)
		defer tick.Stop()
		defer wg.Done()

		for ok = true; ok && err == nil; {
			select {
			case p, ok = <-in:
				if !ok {
					break
				}
				percent = int(p.Percentage())
				err = p.Error()
				out <- p // Forward to the ProgressLogger
			case <-tick.C:
				// From the vim api HttpNfcLeaseProgress(percent) doc, percent ==
				// "Completion status represented as an integer in the 0-100 range."
				// Always report the current value of percent,
				// as it will renew the lease even if the value hasn't changed or is 0
				err = lease.HttpNfcLeaseProgress(cmd.Client, percent)
			}
		}
	}()

	wg.Add(1)

	return &wg
}

func (cmd *ovf) Upload(lease *govmomi.HttpNfcLease, item *types.OvfFileItem, u *url.URL) error {
	file := item.Path

	f, size, err := cmd.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	in := make(chan vim25.Progress)

	out := make(chan vim25.Progress)

	wg := cmd.LeaseUpdater(lease, in, out)

	pwg := cmd.ProgressLogger(fmt.Sprintf("Uploading %s... ", importable(file).Base()), out)

	// defer queue is LIFO..
	defer pwg.Wait() // .... 3) wait for ProgressLogger to return
	defer close(out) // .... 2) propagate close to chained channel
	defer wg.Wait()  // .... 1) wait for Progress channel to close

	opts := soap.Upload{
		Type:          "application/x-vnd.vmware-streamVmdk",
		Method:        "POST",
		ProgressCh:    in,
		ContentLength: size,
	}

	if item.Create {
		opts.Method = "PUT"
	}

	return cmd.Client.Client.Upload(f, u, &opts)
}
