// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/eam"
	"github.com/vmware/govmomi/eam/object"
	"github.com/vmware/govmomi/eam/simulator"
	vcsim "github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

var (
	client struct {
		*eam.Client
		ctx context.Context
		eam object.EsxAgentManager
		vim *vim25.Client
	}
)

func TestMain(m *testing.M) {
	client.ctx = context.Background()

	// Define a new model for vC Sim.
	model := vcsim.VPX()
	defer model.Remove()

	// Create the resources from the model.
	if err := model.Create(); err != nil {
		log.Fatal(err)
	}

	// Register the EAM endpoint.
	model.Service.RegisterSDK(simulator.New())

	// Start the simulator.
	server := model.Service.NewServer()
	defer server.Close()

	// Get a vCenter client to the simulator.
	govmomiClient, err := govmomi.NewClient(client.ctx, server.URL, true)
	if err != nil {
		log.Fatal(err)
	}
	client.vim = govmomiClient.Client

	// Get an EAM client.
	client.Client = eam.NewClient(client.vim)

	// Get the EAM root object.
	client.eam = object.NewEsxAgentManager(client.Client, eam.EsxAgentManager)

	// Run the tests.
	os.Exit(m.Run())
}
