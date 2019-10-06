/*
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

// based on https://www.virtuallyghetto.com/2013/01/retrieving-esxtop-performance-data.html

package esxtop

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/template"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type esxtop struct {
	*flags.HostSystemFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("host.esxtop", &esxtop{})
}

func (cmd *esxtop) Usage() string {
	return "-host.dns=..."
}

func (cmd *esxtop) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *esxtop) Description() string {
	return `Fetch esxtop stats for a host.

This is experimental and should only be used to retrieve metrics not available through metrics.sample

Examples:
  govc host.esxtop -host.dns=... -json | jq '.Stats[] | select(.TypeName == "SchedGroup") | select( .Values.IsVM)'`
}

func (cmd *esxtop) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

type esxtopCounterValues struct {
	Stats []CounterValue
}

func (v *esxtopCounterValues) Write(w io.Writer) error {
	tmpl, err := template.New("test1").Parse(DefaultTemplate)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, v.Stats)
	return err
}

const esxtopServiceName = "Esxtop"

func queryEsxtopService(ctx context.Context, c *vim25.Client, hostName string) (*types.ManagedObjectReference, error) {
	req := types.QueryServiceList{
		This:        c.ServiceContent.ServiceManager.Reference(),
		ServiceName: esxtopServiceName, // not sure how this is used
		Location:    []string{"vmware.host." + hostName},
	}

	queryServiceList, err := methods.QueryServiceList(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	for _, svci := range queryServiceList.Returnval {
		if svci.ServiceName == esxtopServiceName {
			return &(svci.Service), nil
		}
	}

	return nil, fmt.Errorf("could not find service %s", esxtopServiceName)
}

func executeSimpleCommand(ctx context.Context, c *vim25.Client, svc *types.ManagedObjectReference, arguments []string) (string, error) {
	req := types.ExecuteSimpleCommand{
		This:      svc.Reference(),
		Arguments: arguments,
	}

	execSimpleCommand, err := methods.ExecuteSimpleCommand(ctx, c, &req)
	if err != nil {
		return "", err
	}
	return execSimpleCommand.Returnval, nil
}

func (cmd *esxtop) Run(ctx context.Context, f *flag.FlagSet) error {
	var c *vim25.Client
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	svc, err := queryEsxtopService(ctx, c, host.Name())
	if err != nil {
		return err
	}

	returnval, err := executeSimpleCommand(ctx, c, svc, []string{"CounterInfo"})
	if err != nil {
		return err
	}
	counterInfo := ParseCounterInfo(returnval)

	returnval, err = executeSimpleCommand(ctx, c, svc, []string{"FetchStats"})
	if err != nil {
		return err
	}
	stats := ParseCounterValues(returnval, counterInfo)

	returnval, err = executeSimpleCommand(ctx, c, svc, []string{"FreeStats"})
	if err != nil {
		return err
	}

	counterValues := esxtopCounterValues{stats}
	return cmd.WriteResult(&counterValues)
}
