// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox

// noopChannel implements Channel with no-op operations.
// It is used as the TCLO "in" channel when host-to-guest delivery of TCLO
// commands is not implemented. The Service's event loop calls in.Send(nil) and
// in.Receive() on each tick; the no-op implementation causes the service to
// back off to maxDelay (100 ms) and keep running without doing any work,
// keeping vmtoolsd alive for systemd while awaiting guest-initiated RPCI calls.
type noopChannel struct{}

// NewNoopChannelIn returns a Channel whose Send/Receive are permanent no-ops.
// Use this as the rpcIn argument to NewService when the host will not send
// TCLO commands (e.g., govmomi/simulator container-backed VMs where
// host-to-guest DialVM is not yet implemented).
func NewNoopChannelIn() Channel { return &noopChannel{} }

func (c *noopChannel) Start() error             { return nil }
func (c *noopChannel) Stop() error              { return nil }
func (c *noopChannel) Send(_ []byte) error      { return nil }
func (c *noopChannel) Receive() ([]byte, error) { return nil, nil }
