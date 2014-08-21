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

package guest

import (
	"flag"

	"net"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/govc/flags"
)

type GuestFlag struct {
	*flags.ClientFlag
	*flags.SearchFlag

	*AuthFlag
}

func NewGuestFlag() *GuestFlag {
	return &GuestFlag{
		SearchFlag: flags.NewSearchFlag(flags.SearchVirtualMachines),
	}
}

func (flag *GuestFlag) Register(f *flag.FlagSet) {
}

func (flag *GuestFlag) Process() error {
	return nil
}

func (flag *GuestFlag) FileManager() (*govmomi.GuestFileManager, error) {
	c, err := flag.Client()
	if err != nil {
		return nil, err
	}
	return c.GuestOperationsManager().FileManager()
}

func (flag *GuestFlag) ProcessManager() (*govmomi.GuestProcessManager, error) {
	c, err := flag.Client()
	if err != nil {
		return nil, err
	}
	return c.GuestOperationsManager().ProcessManager()
}

func (flag *GuestFlag) ParseURL(urlStr string) (*url.URL, error) {
	c, err := flag.Client()
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	host, port, _ := net.SplitHostPort(u.Host)

	if host == "*" {
		host, _, _ := net.SplitHostPort(c.Client.URL().Host)
		u.Host = net.JoinHostPort(host, port)
	}

	return u, nil
}
