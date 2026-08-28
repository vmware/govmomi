// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// vsock-test is a minimal binary compiled for linux/amd64 and injected into
// container-backed vcsim VMs via RUN.volume.  It exercises the GuestRPC server
// over the unix socket that vcsim mounts at /run/vmware/rpc.sock.
//
// Usage:
//
//	vsock-test --socket /run/vmware/rpc.sock --rpc 'info-get guestinfo.test'
//	vsock-test --socket /run/vmware/rpc.sock --rpc 'info-set guestinfo.test hello'
//	vsock-test --socket /run/vmware/rpc.sock --roundtrip guestinfo.test hello
//
// Exit code 0 = success; non-zero = failure.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/vmware/govmomi/toolbox"
)

func main() {
	sockPath := flag.String("socket", "/run/vmware/rpc.sock", "path to GuestRPC unix socket")
	rpcCmd := flag.String("rpc", "", "single RPCI command to send (e.g. 'info-get guestinfo.foo')")
	roundtrip := flag.Bool("roundtrip", false, "if true, next two args are key and value for a set+get roundtrip")
	flag.Parse()

	ch := toolbox.NewUnixChannelOut(*sockPath)
	out := &toolbox.ChannelOut{Channel: ch}

	if err := ch.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "connect %s: %v\n", *sockPath, err)
		os.Exit(1)
	}
	defer ch.Stop()

	if *roundtrip {
		args := flag.Args()
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "--roundtrip requires two arguments: <key> <value>")
			os.Exit(1)
		}
		key, value := args[0], args[1]

		// Set the value.
		setCmd := fmt.Sprintf("info-set %s %s", key, value)
		if _, err := out.Request([]byte(setCmd)); err != nil {
			fmt.Fprintf(os.Stderr, "info-set: %v\n", err)
			os.Exit(1)
		}

		// Get it back.
		getCmd := fmt.Sprintf("info-get %s", key)
		reply, err := out.Request([]byte(getCmd))
		if err != nil {
			fmt.Fprintf(os.Stderr, "info-get: %v\n", err)
			os.Exit(1)
		}
		got := string(reply)
		if strings.TrimSpace(got) != strings.TrimSpace(value) {
			fmt.Fprintf(os.Stderr, "roundtrip mismatch: want %q got %q\n", value, got)
			os.Exit(1)
		}
		fmt.Printf("roundtrip OK: %s=%s\n", key, got)
		return
	}

	if *rpcCmd != "" {
		reply, err := out.Request([]byte(*rpcCmd))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", *rpcCmd, err)
			os.Exit(1)
		}
		fmt.Printf("1 %s\n", string(reply))
		return
	}

	fmt.Fprintln(os.Stderr, "specify --rpc or --roundtrip")
	os.Exit(1)
}
