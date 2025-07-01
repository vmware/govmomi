// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"testing"
	"time"

	"github.com/vmware/govmomi/eam"
	"github.com/vmware/govmomi/eam/object"
	"github.com/vmware/govmomi/eam/types"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/find"
	vimobject "github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/session"
	vcsim "github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"

	// making sure the SDK endpoint is registered
	_ "github.com/vmware/govmomi/eam/simulator"
)

const waitLoopMessage = `

################################################################################
# When executed with the flags -powerOnVMs and -waitForExit, this test will    #
# pause here and update the screen with the status of the agent VMs until      #
# SIGINT is sent to this process.                                              #
################################################################################
`

var (
	flagPowerOnVMs = flag.Bool("powerOnVMs", false, "Powers on the VMs in the test with Docker")
	flagWaitToExit = flag.Bool("waitToExit", false, "Waits for user input to exit the test")
)

func TestSimulator(t *testing.T) {
	vcsim.Test(func(ctx context.Context, vimClient *vim25.Client) {
		// Create a finder that sets the default datacenter.
		finder := find.NewFinder(vimClient, true)

		// Get the datacenter to use when creating the agency.
		datacenter, err := finder.DefaultDatacenter(ctx)
		if err != nil {
			t.Fatal(err)
		}
		finder.SetDatacenter(datacenter)

		// Get the "vm" folder.
		folder, err := finder.DefaultFolder(ctx)
		if err != nil {
			t.Fatal(err)
		}

		// Get the cluster to use when creating the agency.
		computeResource, err := finder.ClusterComputeResourceOrDefault(ctx, "")
		if err != nil {
			t.Fatal(err)
		}

		// Get the resource pool to use when creating the agency.
		pool, err := computeResource.ResourcePool(ctx)
		if err != nil {
			t.Fatal(err)
		}

		// Get the datastore to use when creating the agency.
		datastore, err := finder.DatastoreOrDefault(ctx, "")
		if err != nil {
			t.Fatal(err)
		}

		// Get the network to use when creating the agency.
		network, err := finder.NetworkOrDefault(ctx, "DVS0")
		if err != nil {
			t.Fatal(err)
		}

		// Get an EAM client.
		eamClient := eam.NewClient(vimClient)

		// Get the EAM root object.
		mgr := object.NewEsxAgentManager(eamClient, eam.EsxAgentManager)

		// Define a function that will list and print the agency MoRefs.
		listAgencies := func() int {
			t.Log("listing agencies")
			agencies, err := mgr.Agencies(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if len(agencies) == 0 {
				t.Log("no agencies")
				return 0
			}
			for _, obj := range agencies {
				t.Logf("agency: %v", obj.Reference())
				config, err := obj.Config(ctx)
				if err != nil {
					t.Fatal(err)
				}
				t.Logf("agency config: %+v", config)
				runtime, err := obj.Runtime(ctx)
				if err != nil {
					t.Fatal(err)
				}
				t.Logf("agency runtime: %+v", runtime)

				agents, err := obj.Agents(ctx)
				if err != nil {
					t.Fatal(err)
				}
				if len(agents) == 0 {
					t.Log("no agents")
				} else {
					for _, a := range agents {
						t.Logf("agent: %v", a.Reference())
						config, err := a.Config(ctx)
						if err != nil {
							t.Fatal(err)
						}
						t.Logf("agent config: %+v", config)
						runtime, err := a.Runtime(ctx)
						if err != nil {
							t.Fatal(err)
						}
						t.Logf("agent runtime: %+v", runtime)
					}
				}
			}
			return len(agencies)
		}

		// List and print the agency MoRefs. There are none.
		if listAgencies() > 0 {
			t.Fatal("no agencies expected")
		}

		// Create a new agency.
		t.Log("creating a new agency")
		agency, err := mgr.CreateAgency(
			ctx,
			&types.AgencyConfigInfo{
				AgencyName: "nginx",
				AgentVmDatastore: []vim.ManagedObjectReference{
					datastore.Reference(),
				},
				Folders: []types.AgencyVMFolder{
					{
						FolderId:     folder.Reference(),
						DatacenterId: datacenter.Reference(),
					},
				},
				ResourcePools: []types.AgencyVMResourcePool{
					{
						ResourcePoolId:    pool.Reference(),
						ComputeResourceId: computeResource.Reference(),
					},
				},
				AgentVmNetwork: []vim.ManagedObjectReference{
					network.Reference(),
				},
				AgentConfig: []types.AgentConfigInfo{
					{
						OvfPackageUrl: "nginx",
					},
					{
						OvfPackageUrl: "nginx",
					},
				},
			},
			string(types.EamObjectRuntimeInfoGoalStateEnabled),
		)
		if err != nil {
			if soap.IsSoapFault(err) {
				fault := soap.ToSoapFault(err).VimFault()
				t.Fatalf("%[1]T %[1]v", fault)
			} else {
				t.Fatalf("%[1]T %[1]v", err)
			}
		}
		t.Logf("created agency: %v", agency.Reference())

		// List the agencies again, and this time the newly created agency will be
		// printed to the console.
		if listAgencies() != 1 {
			t.Fatal("one agency expected")
		}

		// Check whether or not we want to power on the VMs with Docker.
		if *flagPowerOnVMs {
			agencies, err := mgr.Agencies(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if len(agencies) == 0 {
				t.Fatal("no agencies")
			}
			agency := agencies[0]

			// Wait for the agent VMs to have IP addresses
			{
				agents, err := agency.Agents(ctx)
				if err != nil {
					t.Fatal(err)
				}
				if len(agents) == 0 {
					t.Fatal("no agents")
				}

				wait := make(chan struct{})
				done := make(chan struct{})
				errs := make(chan error)
				msgs := make(chan string)
				var wg sync.WaitGroup
				wg.Add(len(agents))

				for _, agent := range agents {
					agent := agent
					go func() {
						var (
							vmPowerState string
							vmIp         string
							once         sync.Once
							vmon         = map[vim.ManagedObjectReference]struct{}{}
						)
						for {
							runtime, err := agent.Runtime(ctx)
							if err != nil {
								errs <- err
								return
							}
							if runtime.Vm == nil {
								errs <- fmt.Errorf("vm is nil for agent %s", agent.Reference())
								return
							}
							if _, ok := vmon[*runtime.Vm]; !ok {
								vm := vimobject.NewVirtualMachine(vimClient, *runtime.Vm)
								if _, err := vm.PowerOn(ctx); err != nil {
									errs <- err
									return
								}
								vmon[*runtime.Vm] = struct{}{}
							}

							if vmIp != runtime.VmIp || vmPowerState != string(runtime.VmPowerState) {
								vmIp = runtime.VmIp
								vmPowerState = string(runtime.VmPowerState)
								msgs <- fmt.Sprintf(
									"%v: name=%s, powerState=%s, ipAddr=%s",
									*runtime.Vm,
									runtime.VmName,
									vmPowerState,
									vmIp)
							}
							if vmIp != "" {
								once.Do(func() {
									wg.Done()
								})
							}
							select {
							case <-time.After(1 * time.Second):
							case <-done:
								return
							}
						}
					}()
				}

				go func() {
					wg.Wait()
					if *flagWaitToExit {
						<-wait
					}
					close(done)
				}()

				go func() {
					defer close(wait)
					if !*flagWaitToExit {
						return
					}
					t.Log(waitLoopMessage)
					c := make(chan os.Signal, 1)
					signal.Notify(c, os.Interrupt)
					for {
						select {
						case <-time.After(1 * time.Second):
						case <-c:
							return
						}
					}
				}()

				t.Log("waiting for the agent VMs to power on and get IP addresses")
				timeout := time.After(10 * time.Minute)
				func() {
					for {
						select {
						case msg := <-msgs:
							t.Log(msg)
						case err := <-errs:
							t.Fatal(err)
						case <-done:
							return
						case <-timeout:
							t.Fatal("timed out waiting for agent VMs to power on and get IP addresses")
							return
						}
					}
				}()
			}
		}

		// Destroy the agency.
		t.Log("destroying agency")
		if err := agency.Destroy(ctx); err != nil {
			t.Fatal(err)
		}

		if listAgencies() != 0 {
			t.Fatal("no agencies expected")
		}
	})
}

func TestNotAuthenticated(t *testing.T) {
	vcsim.Test(func(ctx context.Context, vimClient *vim25.Client) {

		t.Run("TerminateSession", func(t *testing.T) {
			// Terminate the session.
			sessionManager := session.NewManager(vimClient)
			if err := sessionManager.Logout(ctx); err != nil {
				t.Fatalf("logout failed: %v", err)
			}
		})

		t.Run("ValidateFaults", func(t *testing.T) {

			t.Run("vim.NotAuthenticated", func(t *testing.T) {
				t.Parallel()

				// Create a finder to get the default datacenter.
				finder := find.NewFinder(vimClient, true)

				// Try to get the default datacenter, but receive a NotAuthenticated
				// error.
				_, err := finder.DefaultDatacenter(ctx)

				if !fault.Is(err, &vim.NotAuthenticated{}) {
					t.Fatal(err)
				}
			})

			t.Run("eam.EamInvalidLogin", func(t *testing.T) {
				t.Parallel()

				// Get an EAM client.
				eamClient := eam.NewClient(vimClient)

				// Get the EAM root object.
				mgr := object.NewEsxAgentManager(eamClient, eam.EsxAgentManager)

				// Try to list the agencies, but receive an EamInvalidLogin error.
				_, err := mgr.Agencies(ctx)

				if !fault.Is(err, &types.EamInvalidLogin{}) {
					t.Fatalf("err=%[1]T %+[1]v", err)
				}
			})
		})

	})
}
