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
// Copyright 2016-2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vmware/govmomi/toolbox"
	gv "github.com/vmware/govmomi/version"
)

// This example can be run on a VM hosted by ESX, Fusion or Workstation
func main() {
	ver := flag.Bool("version", false, "Show vcsim version number")
	flag.Parse()

	// if -version is requested, print it and exit cleanly
	if *ver {
		v := gv.GitVersion
		if v == "" {
			v = gv.Version
		}
		fmt.Printf("toolbox %s\n", v)
		os.Exit(0)
	}

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
