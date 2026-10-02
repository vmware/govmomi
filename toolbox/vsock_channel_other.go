// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

//go:build !linux

package toolbox

import "errors"

// VsockChannel is the stub for non-Linux platforms. AF_VSOCK is a Linux kernel
// socket family; this stub satisfies the Channel interface so that callers can
// use NewVsockChannelOut() in cross-platform code and handle the Start() error
// gracefully (e.g., by falling back to the backdoor channel).
type VsockChannel struct{}

// NewVsockChannelOut returns a VsockChannel stub that always errors on Start.
func NewVsockChannelOut() Channel { return &VsockChannel{} }

// IsVsockAvailable returns false on non-Linux platforms.
func IsVsockAvailable() bool { return false }

func (c *VsockChannel) Start() error {
	return errors.New("VsockChannel: AF_VSOCK is only supported on Linux")
}
func (c *VsockChannel) Stop() error         { return nil }
func (c *VsockChannel) Send(_ []byte) error { return errors.New("VsockChannel: not started") }
func (c *VsockChannel) Receive() ([]byte, error) {
	return nil, errors.New("VsockChannel: not started")
}
