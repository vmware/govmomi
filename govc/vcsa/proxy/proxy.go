/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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

package logging

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	vnetworking "github.com/vmware/govmomi/vapi/appliance/networking"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vcsa.networking.proxy.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *info) Description() string {
	return `Retrieve the VC networking proxy configuration

Examples:
  govc vcsa.networking.proxy.info`
}

type proxyResult struct {
    proxy vnetworking.ProxyConfig
    noProxy []string
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	fwd := vnetworking.NewManager(c)

	proxyRes, err := fwd.Proxy(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	noProxyRes, err := fwd.NoProxy(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var res proxyResult
	res.proxy = proxyRes
	res.noProxy = noProxyRes
	return cmd.WriteResult(res)
}

func (res proxyResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	
	printProxyConfig("HTTP", res.proxy.Http, tw)
	printProxyConfig("HTTPS", res.proxy.Https, tw)
	printProxyConfig("FTP", res.proxy.Ftp, tw)
	fmt.Fprintf(tw, "No Proxy addresses:\t%s\n", strings.Join(res.noProxy, ", "))

	return tw.Flush()
}

func printProxyConfig(proxyName string, proxyProtocolConfig vnetworking.ProtocolProxyConfig, w io.Writer) {
	if (!proxyProtocolConfig.Enabled) {
		fmt.Fprintf(w, "%s proxy:\tDisabled\n", proxyName)
		return
	}

	fmt.Fprintf(w, "%s proxy:\tEnabled\n", proxyName)
	fmt.Fprintf(w, "\tServer:\t%s\n", proxyProtocolConfig.Server)
	fmt.Fprintf(w, "\tPort:\t%d\n", proxyProtocolConfig.Port)
	if (proxyProtocolConfig.Username != "") {
		fmt.Fprintf(w, "\tUsername:\t%s\n", proxyProtocolConfig.Username)
	}
}
