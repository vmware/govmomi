// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"bytes"
	"fmt"
)

// Channel abstracts the guest<->vmx RPC transport
type Channel interface {
	Start() error
	Stop() error
	Send([]byte) error
	Receive() ([]byte, error)
}

var (
	rpciOK  = []byte{'1', ' '}
	rpciERR = []byte{'0', ' '}
)

// ChannelOut extends Channel to provide RPCI protocol helpers
type ChannelOut struct {
	Channel
}

// Request sends an RPC command to the vmx and checks the return code for success or error
func (c *ChannelOut) Request(request []byte) ([]byte, error) {
	if err := c.Send(request); err != nil {
		return nil, err
	}

	reply, err := c.Receive()
	if err != nil {
		return nil, err
	}

	if bytes.HasPrefix(reply, rpciOK) {
		return reply[2:], nil
	}

	return nil, fmt.Errorf("request %q: %q", request, reply)
}
