// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/vmware/govmomi/eam/object"
	"github.com/vmware/govmomi/eam/types"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

func TestAgency(t *testing.T) {

	// Create a finder that sets the default datacenter.
	finder := find.NewFinder(client.vim, true)

	// Get the datacenter to use when creating the agency.
	datacenter, err := finder.DefaultDatacenter(client.ctx)
	if err != nil {
		t.Fatal(err)
	}
	finder.SetDatacenter(datacenter)

	// Get the "vm" folder.
	folder, err := finder.DefaultFolder(client.ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Get the cluster to use when creating the agency.
	computeResource, err := finder.ClusterComputeResourceOrDefault(client.ctx, "")
	if err != nil {
		t.Fatal(err)
	}

	// Get the resource pool to use when creating the agency.
	pool, err := computeResource.ResourcePool(client.ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Get the datastore to use when creating the agency.
	datastore, err := finder.DatastoreOrDefault(client.ctx, "")
	if err != nil {
		t.Fatal(err)
	}

	// Get the network to use when creating the agency.
	network, err := finder.NetworkOrDefault(client.ctx, "DVS0")
	if err != nil {
		t.Fatal(err)
	}

	const (
		initialGoalState = string(types.EamObjectRuntimeInfoGoalStateEnabled)
	)
	var (
		agency       object.Agency
		agencyConfig = &types.AgencyConfigInfo{
			AgencyName: t.Name(),
			AgentName:  t.Name(),
			AgentConfig: []types.AgentConfigInfo{
				{
					HostVersion: "1",
				},
			},
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
		}
	)

	agencyExists := func(agency object.Agency) (bool, error) {
		agencies, err := client.eam.Agencies(client.ctx)
		if err != nil {
			return false, err
		}
		for _, a := range agencies {
			if a.Reference() == agency.Reference() {
				return true, nil
			}
		}
		return false, nil
	}

	testCreate := func(t *testing.T) {
		var err error
		if agency, err = client.eam.CreateAgency(
			client.ctx, agencyConfig, initialGoalState); err != nil {
			t.Fatal(err)
		}
		if ok, err := agencyExists(agency); !ok {
			if err != nil {
				t.Fatal(err)
			}
			t.Fatal("agency not returned by Agencies")
		}
	}

	testConfig := func(t *testing.T) {
		t.Parallel()

		baseConfig, err := agency.Config(client.ctx)
		if err != nil {
			t.Fatal(err)
		}
		config := baseConfig.GetAgencyConfigInfo()
		if config.AgencyName != agencyConfig.AgencyName {
			t.Fatalf(
				"unexpected agency name: exp=%v, act=%v",
				agencyConfig.AgencyName,
				config.AgencyName)
		}
		if config.AgentName != agencyConfig.AgentName {
			t.Fatalf(
				"unexpected agency agent name: exp=%v, act=%v",
				agencyConfig.AgentName,
				config.AgentName)
		}
	}

	// This test waits up to 10 seconds for the agency.runtime.issue
	// list to be non-empty. Because this test is run in parallel with
	// the TestAgency.Created.Runtime.Issues.Add test, there should be an
	// issue returned by the tested call before 10 seconds has elapsed.
	testRuntimeIssues := func(t *testing.T) {
		t.Parallel()
		const (
			waitTotalSecs    = 10
			waitIntervalSecs = time.Duration(1) * time.Second
		)
		hasIssues := false
		for i := 0; i < waitTotalSecs; i++ {
			runtime, err := agency.Runtime(client.ctx)
			if err != nil {
				t.Fatal(err)
			}
			if len(runtime.Issue) > 0 {
				hasIssues = true
				break
			}
			time.Sleep(waitIntervalSecs)
		}
		if !hasIssues {
			t.Fatalf(
				"agency.runtime had no issues after %d seconds",
				waitTotalSecs)
		}
	}

	testRuntimeGoalState := func(t *testing.T) {
		validateExpectedGoalState := func(expGoalState any) error {
			runtime, err := agency.Runtime(client.ctx)
			if err != nil {
				return err
			}
			if runtime.GoalState != fmt.Sprintf("%s", expGoalState) {
				return fmt.Errorf(
					"unexpected agency goal state: exp=%v, act=%v",
					expGoalState,
					runtime.GoalState)
			}
			return nil
		}

		t.Run("Initial", func(t *testing.T) {
			if err := validateExpectedGoalState(
				initialGoalState); err != nil {
				t.Fatal(err)
			}
		})
		t.Run("Disabled", func(t *testing.T) {
			if err := agency.Disable(client.ctx); err != nil {
				t.Fatal(err)
			}
			if err := validateExpectedGoalState(
				types.EamObjectRuntimeInfoGoalStateDisabled); err != nil {
				t.Fatal(err)
			}
		})
		t.Run("Uninstalled", func(t *testing.T) {
			if err := agency.Uninstall(client.ctx); err != nil {
				t.Fatal(err)
			}
			if err := validateExpectedGoalState(
				types.EamObjectRuntimeInfoGoalStateUninstalled); err != nil {
				t.Fatal(err)
			}
		})
		t.Run("Enabled", func(t *testing.T) {
			if err := agency.Enable(client.ctx); err != nil {
				t.Fatal(err)
			}
			if err := validateExpectedGoalState(
				types.EamObjectRuntimeInfoGoalStateEnabled); err != nil {
				t.Fatal(err)
			}
		})
	}

	testRuntime := func(t *testing.T) {
		t.Parallel()

		runtime, err := agency.Runtime(client.ctx)
		if err != nil {
			t.Fatal(err)
		}
		if runtime.GoalState != initialGoalState {
			t.Fatalf(
				"unexpected agency goal state: exp=%v, act=%v",
				initialGoalState,
				runtime.GoalState)
		}

		t.Run("Issues", testRuntimeIssues)
		t.Run("GoalState", testRuntimeGoalState)
	}

	testIssues := func(t *testing.T) {
		t.Parallel()

		var issueKey int32

		t.Run("Add", func(t *testing.T) {
			baseIssue, err := agency.AddIssue(client.ctx, &types.OrphanedAgency{
				AgencyIssue: types.AgencyIssue{
					Agency:     agency.Reference(),
					AgencyName: agencyConfig.AgencyName,
				},
			})
			if err != nil {
				t.Fatal(err)
			}
			baseAgencyIssue, ok := baseIssue.(types.BaseAgencyIssue)
			if !ok {
				t.Fatalf(
					"unexpected issue type: exp=%v, act=%T",
					"types.BaseAgencyIssue",
					baseIssue)
			}
			issue := baseAgencyIssue.GetAgencyIssue()
			if issue == nil {
				t.Fatal("returned issue is nil")
			}
			if issue.Key == 0 {
				t.Fatal("issue.Key == 0")
			}
			if issue.Time.IsZero() {
				t.Fatal("issue.Time is not set")
			}
			if issue.Agency != agency.Reference() {
				t.Fatalf(
					"unexpected agency moRef: exp=%v, act=%v",
					agency.Reference(),
					issue.Agency)
			}
			if issue.AgencyName != agencyConfig.AgencyName {
				t.Fatalf(
					"unexpected agency name: exp=%v, act=%v",
					agencyConfig.AgencyName,
					issue.AgencyName)
			}
			issueKey = issue.Key
			t.Logf("added new issue to agency: agency=%v, issueKey=%d",
				agency.Reference(), issueKey)
		})

		t.Run("Query", func(t *testing.T) {
			validateIssueIsInList := func(
				issues []types.BaseIssue,
				inErr error) error {

				if inErr != nil {
					return inErr
				}
				if len(issues) == 0 {
					return errors.New("no issues returned")
				}
				foundIssueKey := false
				for _, baseIssue := range issues {
					issue := baseIssue.GetIssue()
					if issue == nil {
						return errors.New("returned issue is nil")
					}
					if issue.Key == issueKey {
						foundIssueKey = true
						break
					}
				}
				if !foundIssueKey {
					return fmt.Errorf(
						"did not find expected issue: key=%d", issueKey)
				}
				return nil
			}

			t.Run("All", func(t *testing.T) {
				t.Parallel()
				if err := validateIssueIsInList(
					agency.Issues(client.ctx)); err != nil {
					t.Fatal(err)
				}
			})
			t.Run("ByKey", func(t *testing.T) {
				t.Parallel()
				if err := validateIssueIsInList(
					agency.Issues(client.ctx, issueKey)); err != nil {
					t.Fatal(err)
				}
			})
		})
	}

	testAgents := func(t *testing.T) {
		t.Parallel()

		var agent *object.Agent

		t.Run("Query", func(t *testing.T) {
			agents, err := agency.Agents(client.ctx)
			if err != nil {
				t.Fatal(err)
			}
			if len(agents) != 1 {
				t.Fatal("expected one agent")
			}
			agent = &agents[0]
		})

		t.Run("Queried", func(t *testing.T) {
			t.Run("Config", func(t *testing.T) {
				t.Parallel()
				config, err := agent.Config(client.ctx)
				if err != nil {
					t.Fatal(err)
				}
				if config.HostVersion != "1" {
					t.Fatalf(
						"unexpected agent config host version: exp=%v, act=%v",
						"1",
						config.HostVersion)
				}
			})
			t.Run("Runtime", func(t *testing.T) {
				t.Parallel()
				var runtime *types.AgentRuntimeInfo

				t.Run("Query", func(t *testing.T) {
					var err error
					if runtime, err = agent.Runtime(client.ctx); err != nil {
						t.Fatal(err)
					} else if runtime.Agency == nil {
						t.Fatal("agent runtime.Agency is nil")
					}
				})
				t.Run("Properties", func(t *testing.T) {
					t.Run("Agency", func(t *testing.T) {
						t.Parallel()
						if *runtime.Agency != agency.Reference() {
							t.Fatalf(
								"unexpected agent runtime.Agency: exp=%v, act=%v",
								agency.Reference(),
								*runtime.Agency)
						}
					})
					t.Run("VirtualMachine", func(t *testing.T) {
						t.Parallel()
						if runtime.Vm == nil {
							t.Fatal("runtime.Vm is nil")
						}
					})
				})
			})
		})
	}

	testDestroy := func(t *testing.T) {
		t.Run("HappyPath", func(t *testing.T) {
			if err := agency.Destroy(client.ctx); err != nil {
				t.Fatal(err)
			}
		})

		// Attempt to destroy the agency a second time, asserting that a
		// ManagedObjectNotFound error will occur.
		t.Run("NotFound", func(t *testing.T) {
			if err := agency.Destroy(client.ctx); err == nil {
				t.Fatal("error did not occur")
			} else if fault := soap.ToSoapFault(err); fault == nil {
				t.Fatalf("soap fault did not occur: %+v", err)
			} else if fault, ok := fault.VimFault().(vim.ManagedObjectNotFound); !ok {
				t.Fatalf("expected soap fault did not occur: %+v", fault)
			} else if fault.Obj != agency.Reference() {
				t.Fatalf("unexpected error details: exp=%v, act=%v",
					agency.Reference(), fault.Obj)
			}
		})

		t.Run("Verify", func(t *testing.T) {
			if ok, err := agencyExists(agency); err != nil {
				t.Fatal(err)
			} else if ok {
				t.Fatal("agency still returned after being destroyed")
			}
		})
	}

	t.Run("Create", testCreate)

	t.Run("Created", func(t *testing.T) {
		t.Run("Config", testConfig)
		t.Run("Issues", testIssues)
		t.Run("Runtime", testRuntime)
		t.Run("Agents", testAgents)
	})

	t.Run("Destroy", testDestroy)
}
