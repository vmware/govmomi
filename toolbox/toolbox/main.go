// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vmware/govmomi/toolbox"
)

// This example can be run on a VM hosted by ESX, Fusion or Workstation
func main() {
	flag.Parse()

	in := toolbox.NewBackdoorChannelIn()
	out := toolbox.NewBackdoorChannelOut()

	service := toolbox.NewService(in, out)

	if os.Getuid() == 0 {
		service.Power.Halt.Handler = toolbox.Halt
		service.Power.Reboot.Handler = toolbox.Reboot
	}

	err := service.Start()
	if err != nil {
		log.Fatal(err)
	}

	// handle the signals and gracefully shutdown the service
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("signal %s received", <-sig)
		service.Stop()
	}()

	service.Wait()
}
