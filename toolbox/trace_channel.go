// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	Trace = false

	traceLog io.Writer = os.Stderr
)

func init() {
	flag.BoolVar(&Trace, "toolbox.trace", Trace, "Enable toolbox trace")
}

type TraceChannel struct {
	Channel
	log io.Writer
}

func NewTraceChannel(c Channel) Channel {
	if !Trace {
		return c
	}

	return &TraceChannel{
		Channel: c,
		log:     traceLog,
	}
}

func (d *TraceChannel) Start() error {
	err := d.Channel.Start()

	return err
}

func (d *TraceChannel) Stop() error {
	err := d.Channel.Stop()

	return err
}

func (d *TraceChannel) Send(buf []byte) error {
	if len(buf) > 0 {
		fmt.Fprintf(d.log, "SEND %d...\n%s\n", len(buf), hex.Dump(buf))
	}

	err := d.Channel.Send(buf)

	return err
}

func (d *TraceChannel) Receive() ([]byte, error) {
	buf, err := d.Channel.Receive()

	if err == nil && len(buf) > 0 {
		fmt.Fprintf(d.log, "RECV %d...\n%s\n", len(buf), hex.Dump(buf))
	}

	return buf, err
}
