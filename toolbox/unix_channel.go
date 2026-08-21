// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"fmt"
	"net"
)

// UnixChannel implements the Channel interface over a Unix stream socket.
//
// Wire format: DataMap packets — the same framing real vmtoolsd uses over
// AF_VSOCK, and the same framing VsockChannel speaks. One framing across every
// GuestRPC transport means the server never has to guess which it is reading,
// and it keeps the simulator exercising the parse path a real guest exercises.
//
// fastClose is NOT set: the connection is opened once by Start and reused for
// the lifetime of the channel, so the server must keep serving it. (VsockChannel
// sets fastClose because it opens a fresh connection per exchange.)
//
// This is compatible with the ChannelOut.Request pattern used throughout the
// toolbox package: Send("info-set guestinfo.xxx value") then Receive() returns
// "1 " on success or "0 <error>" on failure.
//
// Usage (guest side – inside a container or test binary):
//
//	ch := toolbox.NewUnixChannelOut("/run/vmware/rpc.sock")
//	out := &toolbox.ChannelOut{Channel: ch}
//	if err := ch.Start(); err != nil { ... }
//	reply, err := out.Request([]byte("info-set guestinfo.foo bar"))
//
// The host side (govmomi/simulator.GuestRPCServer) speaks the same framing.
type UnixChannel struct {
	path string
	conn net.Conn
}

// NewUnixChannelOut returns a Channel that connects to the unix socket at path.
// Call Start() before use.
func NewUnixChannelOut(path string) Channel {
	return &UnixChannel{path: path}
}

// Start connects to the unix socket.
func (c *UnixChannel) Start() error {
	conn, err := net.Dial("unix", c.path)
	if err != nil {
		return fmt.Errorf("UnixChannel.Start %s: %w", c.path, err)
	}
	c.conn = conn
	return nil
}

// Stop closes the connection.
func (c *UnixChannel) Stop() error {
	if c.conn == nil {
		return nil
	}
	err := c.conn.Close()
	c.conn = nil
	return err
}

// Send transmits buf as a DataMap packet to the server.
func (c *UnixChannel) Send(buf []byte) error {
	return WriteDataMapPacket(c.conn, buf, false)
}

// Receive reads one DataMap packet from the server.
func (c *UnixChannel) Receive() ([]byte, error) {
	payload, _, err := ReadDataMapPacket(c.conn)
	return payload, err
}
