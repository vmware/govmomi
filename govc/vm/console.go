/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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

package vm

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"time"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type console struct {
	*flags.VirtualMachineFlag

	h5      bool
	wss     bool
	capture string

	sleep   int
	multi	int
}

func init() {
	cli.Register("vm.console", &console{})
}

func (cmd *console) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.BoolVar(&cmd.h5, "h5", false, "Generate HTML5 UI console link")
	f.BoolVar(&cmd.wss, "wss", false, "Generate WebSocket console link")

	f.StringVar(&cmd.capture, "capture", "", "Capture console screen shot to file")
	f.IntVar(&cmd.sleep, "sleep", 0, "Sleep before capture")
	f.IntVar(&cmd.multi, "multi", 1, "Capture multiple intervals")
}

func (cmd *console) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *console) Usage() string {
	return "VM"
}

func (cmd *console) Description() string {
	return `Generate console URL or screen capture for VM.

One of VMRC, VMware Player, VMware Fusion or VMware Workstation must be installed to
open VMRC console URLs.

Examples:
  govc vm.console my-vm
  govc vm.console -capture screen.png my-vm				# screen capture
  govc vm.console -capture - my-vm | display 				# screen capture to stdout
  govc vm.console -capture "{vm}-{r}.png" -multi 5 -sleep 2 my-vm 	# multiple, timed screenshots
  open $(govc vm.console my-vm)              				# MacOSX VMRC
  open $(govc vm.console -h5 my-vm)          				# MacOSX H5
  xdg-open $(govc vm.console my-vm)          				# Linux VMRC
  xdg-open $(govc vm.console -h5 my-vm)      				# Linux H5`
}

func (cmd *console) Run(ctx context.Context, f *flag.FlagSet) error {
	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	if len(vms) != 1 {
		return flag.ErrHelp
	}

	vm := vms[0]

	state, err := vm.PowerState(ctx)
	if err != nil {
		return err
	}

	if (cmd.capture != "" || cmd.wss) && state != types.VirtualMachinePowerStatePoweredOn {
		return fmt.Errorf("vm is not powered on (%s)", state)
	}

	c := vm.Client()

	u := c.URL()

	if cmd.capture != "" {
		multi := cmd.multi
		startTime := time.Now()

		if multi < 1 || cmd.capture == "-" {
			multi = 1
		}

		err := error(nil)
		ms := cmd.sleep*1000

		//accumulate runtime instead of multiplying in the loop
		//this is not the real runtime but what would be expected running through the loop in 0ms processing time
		expectedrunms := 0
		//waiting a little bit less because we are expecting 50ms runtime due to https download and storing
		expectedworktime := 50

		for n := 0; n < multi && err == nil; n++ {
			now := time.Now()

			if cmd.sleep > 0 {
				//calculate against startTime so that delays do not accumulate
				diff := int(now.UnixMilli()-startTime.UnixMilli())-expectedrunms
				relsleep := (ms-diff-expectedworktime)
				
				if relsleep > 0 {
					if relsleep > ms {
						relsleep = ms
					}
					time.Sleep(time.Duration(relsleep) * time.Millisecond) 
				}
				expectedrunms += ms
			}

			u.Path = "/screen"
			query := url.Values{"id": []string{vm.Reference().Value}}
			u.RawQuery = query.Encode()

			param := soap.DefaultDownload

			if cmd.capture == "-" {
				w, _, derr := c.Download(ctx, u, &param)
				if derr != nil {
					return derr
				}

				_, err = io.Copy(os.Stdout, w)
				if err != nil {
					return err
				}

				return w.Close()
			} else {
				// -capture "vm-{r}.png" for a relative name
				// relative is the unix time relative to the start time
				// this can be handy if you want capture the time between screenshots but unix time is unreadable
				
				fileName := cmd.capture
				fileName = strings.Replace(fileName,"{vm}",vm.Name(),-1)
				fileName = strings.Replace(fileName,"{d}",fmt.Sprintf("%06d",n),-1)
				fileName = strings.Replace(fileName,"{n}",fmt.Sprintf("%d",n),-1)
				fileName = strings.Replace(fileName,"{u}",fmt.Sprintf("%d",now.Unix()),-1)
				fileName = strings.Replace(fileName,"{r}",fmt.Sprintf("%06d",(now.Unix()-startTime.Unix())),-1)
				err = c.DownloadFile(ctx, fileName, u, &param)
			}
		}
		return err
	}

	if cmd.wss {
		ticket, err := vm.AcquireTicket(ctx, string(types.VirtualMachineTicketTypeWebmks))
		if err != nil {
			return err
		}

		link := fmt.Sprintf("wss://%s:%d/ticket/%s", ticket.Host, ticket.Port, ticket.Ticket)
		fmt.Fprintln(cmd.Out, link)
		return nil
	}

	m := session.NewManager(c)
	ticket, err := m.AcquireCloneTicket(ctx)
	if err != nil {
		return err
	}

	var link string

	if cmd.h5 {
		m := object.NewOptionManager(c, *c.ServiceContent.Setting)

		opt, err := m.Query(ctx, "VirtualCenter.FQDN")
		if err != nil {
			return err
		}

		fqdn := opt[0].GetOptionValue().Value.(string)

		var info object.HostCertificateInfo
		err = info.FromURL(u, nil)
		if err != nil {
			return err
		}

		u.Path = "/ui/webconsole.html"

		u.RawQuery = url.Values{
			"vmId":          []string{vm.Reference().Value},
			"vmName":        []string{vm.Name()},
			"serverGuid":    []string{c.ServiceContent.About.InstanceUuid},
			"host":          []string{fqdn},
			"sessionTicket": []string{ticket},
			"thumbprint":    []string{info.ThumbprintSHA1},
		}.Encode()

		link = u.String()
	} else {
		link = fmt.Sprintf("vmrc://clone:%s@%s/?moid=%s", ticket, u.Hostname(), vm.Reference().Value)
	}

	fmt.Fprintln(cmd.Out, link)

	return nil
}
