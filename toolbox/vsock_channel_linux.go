// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package toolbox

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"unsafe"
)

const (
	afVsock       = 40  // AF_VSOCK
	vmaddrCIDHost = 2   // VMADDR_CID_HOST — the ESX hypervisor CID
	vsockRPCIPort = 976 // VMware RPCI / GuestRPC vsock port
)

// sockaddrVM mirrors struct sockaddr_vm from <linux/vm_sockets.h> (amd64 layout).
type sockaddrVM struct {
	family    uint16
	reserved1 uint16
	port      uint32
	cid       uint32
	flags     uint8
	pad       [3]uint8
}

// VsockChannel implements the Channel interface over AF_VSOCK using DataMap
// framing. This is the same wire protocol used by the real vmtoolsd binary over
// its vsock connection to ESX.
//
// The channel is lazy: Start() succeeds immediately without opening a socket.
// Each Send call opens a fresh AF_VSOCK connection, writes a DataMap packet
// with FAST_CLOSE=true (one-shot model matching real vmtoolsd --cmd behavior),
// and keeps the connection open for the paired Receive call. Receive reads the
// response and closes the connection.
//
// Thread safety: a single VsockChannel must not be used concurrently from
// multiple goroutines. The Service (service.go) calls Send+Receive only from
// its single event goroutine.
type VsockChannel struct {
	cid  uint32
	port uint32
	conn net.Conn // open connection after Send, closed after Receive
}

// NewVsockChannelOut returns a Channel that speaks RPCI via AF_VSOCK to the
// hypervisor CID (2) on the GuestRPC port (976) using DataMap framing.
func NewVsockChannelOut() Channel {
	return &VsockChannel{cid: vmaddrCIDHost, port: vsockRPCIPort}
}

// Start is a no-op; the connection is opened lazily on the first Send call.
// This allows Service.Start() to succeed even when running outside a VM.
func (c *VsockChannel) Start() error { return nil }

// Stop closes any open connection.
func (c *VsockChannel) Stop() error {
	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		return err
	}
	return nil
}

// Send opens a new AF_VSOCK connection to the hypervisor and writes a DataMap
// packet containing payload. The connection is left open for the paired
// Receive call. If a previous connection was not closed, it is closed first.
//
// socket() is called via syscall.Socket and connect() via syscall.Syscall;
// both result in the kernel socket(2)/connect(2) system calls, making them
// visible to the vcsim seccomp filter which replaces the AF_VSOCK FD with a
// Unix socketpair.
func (c *VsockChannel) Send(buf []byte) error {
	if c.conn != nil {
		_ = c.conn.Close()
		c.conn = nil
	}

	conn, err := dialVsockLinux(c.cid, c.port)
	if err != nil {
		return err
	}
	c.conn = conn

	return WriteDataMapPacket(c.conn, buf, true /* fastClose */)
}

// Receive reads a DataMap response from the open connection and closes it.
func (c *VsockChannel) Receive() ([]byte, error) {
	if c.conn == nil {
		return nil, fmt.Errorf("VsockChannel: Receive called without a prior Send")
	}
	defer func() {
		_ = c.conn.Close()
		c.conn = nil
	}()

	payload, _, err := ReadDataMapPacket(c.conn)
	return payload, err
}

// IsVsockAvailable returns true if the kernel supports creating AF_VSOCK
// sockets on this system. On govmomi/simulator container-backed VMs the seccomp
// intercept handles the socket() call, so this returns true inside containers
// with RUN.vmci=true. On bare containers or non-VM hosts without the vmci kernel
// module it returns false, enabling the backdoor fallback.
func IsVsockAvailable() bool {
	fd, err := syscall.Socket(afVsock, syscall.SOCK_STREAM, 0)
	if err != nil {
		return false
	}
	_ = syscall.Close(fd)
	return true
}

// dialVsockLinux opens an AF_VSOCK SOCK_STREAM socket and connects it to
// cid:port using raw syscalls (required for seccomp intercept compatibility).
func dialVsockLinux(cid, port uint32) (net.Conn, error) {
	fd, err := syscall.Socket(afVsock, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, fmt.Errorf("VsockChannel: socket(AF_VSOCK): %w", err)
	}

	addr := sockaddrVM{family: afVsock, port: port, cid: cid}
	_, _, errno := syscall.Syscall(
		syscall.SYS_CONNECT,
		uintptr(fd),
		uintptr(unsafe.Pointer(&addr)),
		unsafe.Sizeof(addr),
	)
	if errno != 0 {
		_ = syscall.Close(fd)
		return nil, fmt.Errorf("VsockChannel: connect(cid=%d, port=%d): %w", cid, port, errno)
	}

	f := os.NewFile(uintptr(fd), fmt.Sprintf("vsock-cid%d-port%d", cid, port))
	conn, err := net.FileConn(f)
	_ = f.Close() // net.FileConn duplicates the fd; the original can be closed
	if err != nil {
		return nil, fmt.Errorf("VsockChannel: FileConn: %w", err)
	}
	return conn, nil
}
