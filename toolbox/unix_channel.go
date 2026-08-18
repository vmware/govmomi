// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// UnixChannel implements the Channel interface over a Unix stream socket.
//
// Wire format: each message is framed as a 4-byte little-endian uint32 length
// prefix followed by that many payload bytes. The same framing is used in both
// directions (Send and Receive). Zero-length messages are valid (empty payload).
//
// This framing is compatible with the ChannelOut.Request pattern used throughout
// the toolbox package: Send("info-set guestinfo.xxx value") then Receive() returns
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

// Send transmits buf as a length-prefixed frame to the server.
func (c *UnixChannel) Send(buf []byte) error {
	return WriteUnixFrame(c.conn, buf)
}

// Receive reads one length-prefixed frame from the server.
func (c *UnixChannel) Receive() ([]byte, error) {
	return ReadUnixFrame(c.conn)
}

// WriteUnixFrame writes a 4-byte LE length prefix followed by buf.
// Exported so that govmomi/simulator can share the same framing.
func WriteUnixFrame(w io.Writer, buf []byte) error {
	var hdr [4]byte
	binary.LittleEndian.PutUint32(hdr[:], uint32(len(buf)))
	if _, err := w.Write(hdr[:]); err != nil {
		return err
	}
	if len(buf) == 0 {
		return nil
	}
	_, err := w.Write(buf)
	return err
}

// ReadUnixFrame reads a 4-byte LE length prefix and then that many bytes.
// Exported so that govmomi/simulator can share the same framing.
func ReadUnixFrame(r io.Reader) ([]byte, error) {
	var hdr [4]byte
	if _, err := io.ReadFull(r, hdr[:]); err != nil {
		return nil, err
	}
	n := binary.LittleEndian.Uint32(hdr[:])
	if n == 0 {
		return nil, nil
	}
	buf := make([]byte, n)
	_, err := io.ReadFull(r, buf)
	return buf, err
}
