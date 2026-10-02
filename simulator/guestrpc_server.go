// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/vmware/govmomi/toolbox"
	"github.com/vmware/govmomi/vim25/types"
)

const (
	// GuestRPCSocketName is the well-known path inside every container-backed VM
	// with RUN.vmci=true. The host-side unix socket is bind-mounted to this path.
	GuestRPCSocketName = "/run/vmware/rpc.sock"

	// GuestRPCSocketDir is the well-known DIRECTORY path inside every container-backed VM
	// with RUN.vmci=true. vcsim bind-mounts this directory (not just the socket file) so
	// the mount remains visible even when RUN.nestedContainers=true adds --tmpfs /run.
	// A file bind-mount target under a tmpfs must already exist as a file; a directory
	// bind-mount creates a new submount that overrides the tmpfs contents at that path.
	GuestRPCSocketDir = "/run/vmware"

	// guestInfoPrefix is the only ExtraConfig namespace a guest may read or
	// write via GuestRPC. Real VMX restricts guest access to guestinfo.*, and
	// vcsim reads other namespaces (RUN.*) as host-side container directives —
	// so accepting them here would let guest code inspect or reconfigure its
	// own backing on next start.
	guestInfoPrefix = "guestinfo."
)

// isGuestReadable reports whether key is a namespace a guest may read via
// info-get. Real VMX restricts guest reads to guestinfo.*.
func isGuestReadable(key string) bool {
	return strings.HasPrefix(key, guestInfoPrefix)
}

// isGuestWritable reports whether key is a namespace a guest may write via
// info-set. Real VMX additionally makes any key containing "/" read-only
// from the guest (used for host/tools-managed sub-namespaces under
// guestinfo.*), so a writable key must be readable AND slash-free.
func isGuestWritable(key string) bool {
	return isGuestReadable(key) && !strings.Contains(key, "/")
}

// GuestRPCServer is a per-VM host-side server that implements the toolbox RPCI
// protocol over a Unix stream socket.
//
// Wire format: DataMap packets, the same framing real vmtoolsd uses over
// AF_VSOCK (toolbox.ReadDataMapPacket / WriteDataMapPacket). Using one framing
// across every GuestRPC transport means this server never has to guess what it
// is reading. Response status prefix: "1 " = success, "0 " = error (same as
// toolbox.rpciOK / rpciERR).
//
// Each VM with RUN.vmci=true has its own GuestRPCServer at a unique socket path.
// Multiple concurrent VMs do not share any state.
//
// Lifecycle:
//
//	srv = newGuestRPCServer(vm, path)
//	srv.Start(ctx)    — before container create; removes stale socket from prior crash
//	(container runs; guest code connects and issues RPCI commands)
//	srv.Stop()        — from simVM.stop() / simVM.remove()
//
// Recovery: Start removes any stale socket file before binding, so a crashed test
// does not block the next run. The socket file is also removed by Stop().
type GuestRPCServer struct {
	vm         *VirtualMachine
	ctx        *Context // captures the PowerOn Context for AutoUpdate
	socketPath string   // host-side absolute path

	ln   net.Listener
	done chan struct{}
	once sync.Once
	wg   sync.WaitGroup

	// rpcMu serializes concurrent infoGet / infoSet calls from multiple
	// serveConn goroutines.  Without this, concurrent goroutines sharing the
	// same s.ctx bypass ObjectLock's mutual exclusion (ObjectLock.Acquire
	// grants re-entry to the same *Context pointer, so all serveConn goroutines
	// that share s.ctx can enter WithLock/AutoUpdate simultaneously, causing a
	// data race on ExtraConfig).
	//
	// Lock order: rpcMu → ObjectLock (via WithLock/AutoUpdate).  This order is
	// consistent because no code holds ObjectLock and then acquires rpcMu.
	rpcMu sync.RWMutex
}

// newGuestRPCServer creates (but does not start) a GuestRPCServer.
func newGuestRPCServer(vm *VirtualMachine, socketPath string) *GuestRPCServer {
	return &GuestRPCServer{
		vm:         vm,
		socketPath: socketPath,
		done:       make(chan struct{}),
	}
}

// GuestRPCSocketDirPath returns the canonical per-VM host-side socket directory.
// The directory is bind-mounted at GuestRPCSocketDir inside the container.
// Using a directory mount (not a file mount) means the submount is visible even
// when RUN.nestedContainers=true adds --tmpfs /run: the directory mount is applied
// after the tmpfs in the OCI spec and creates a new submount at /run/vmware/,
// whereas a file bind-mount requires the target file to already exist in the tmpfs.
func GuestRPCSocketDirPath(vmUID string) string {
	return filepath.Join(os.TempDir(), fmt.Sprintf("vcsim-rpc-%s", vmUID))
}

// GuestRPCSocketPath returns the canonical per-VM host-side unix socket path.
// The socket resides inside GuestRPCSocketDirPath so the directory bind-mount exposes it.
func GuestRPCSocketPath(vmUID string) string {
	return filepath.Join(GuestRPCSocketDirPath(vmUID), "rpc.sock")
}

// Start begins listening on the unix socket and launches the accept loop.
// ctx is the PowerOn/start Context; it is stored and used for AutoUpdate calls.
// Creates the socket directory and removes any stale socket from a prior crash.
func (s *GuestRPCServer) Start(ctx *Context) error {
	s.ctx = ctx

	// Create the socket directory.  The directory is bind-mounted into the
	// container (not the socket file), so it must exist before docker create runs.
	socketDir := filepath.Dir(s.socketPath)
	if err := os.MkdirAll(socketDir, 0o700); err != nil {
		return fmt.Errorf("GuestRPCServer %s: mkdir socket dir: %w", s.vm.Name, err)
	}

	// Remove stale socket from a prior crash.  If a live process still holds
	// the socket, net.Listen will fail — surfacing the conflict rather than
	// silently breaking it.
	_ = os.Remove(s.socketPath)

	ln, err := net.Listen("unix", s.socketPath)
	if err != nil {
		return fmt.Errorf("GuestRPCServer %s: %w", s.vm.Name, err)
	}
	s.ln = ln

	s.wg.Add(1)
	go s.acceptLoop()

	return nil
}

// Stop closes the listener, signals the accept loop to exit, waits for all
// handler goroutines to finish, and removes the socket file.
// Safe to call multiple times (idempotent via sync.Once).
func (s *GuestRPCServer) Stop() {
	s.once.Do(func() {
		close(s.done)
		if s.ln != nil {
			_ = s.ln.Close()
		}
	})
	s.wg.Wait()
	_ = os.Remove(s.socketPath)
	// Best-effort cleanup of the socket directory created in Start.
	// os.Remove fails silently if the directory is not empty (e.g. left-over
	// files from a test crash), which is acceptable.
	_ = os.Remove(filepath.Dir(s.socketPath))
}

func (s *GuestRPCServer) acceptLoop() {
	defer s.wg.Done()

	for {
		conn, err := s.ln.Accept()
		if err != nil {
			select {
			case <-s.done:
				return // normal shutdown; listener was closed by Stop()
			default:
				log.Printf("GuestRPCServer %s: accept: %v", s.vm.Name, err)
				return
			}
		}
		if Trace {
			log.Printf("GuestRPCServer %s: new connection from %s", s.vm.Name, conn.RemoteAddr())
		}
		s.wg.Add(1)
		go s.serveConn(conn)
	}
}

func (s *GuestRPCServer) serveConn(conn net.Conn) {
	defer s.wg.Done()
	defer conn.Close()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("GuestRPCServer %s: panic in serveConn: %v\n%s",
				s.vm.Name, r, debug.Stack())
		}
	}()

	// One framing on this socket: DataMap, the format real vmtoolsd speaks.
	// There is nothing to detect, which is deliberate — two self-delimiting
	// framings sharing one socket cannot be told apart reliably by any length
	// heuristic, and the guest chooses the lengths by padding its payload.
	s.serveConnDataMap(conn, conn)
}

// serveConnDataMap handles connections using the DataMap framing protocol,
// which is the wire format used by real vmtoolsd binaries over vsock.
//
// DataMap framing is documented in toolbox.ReadDataMapPacket / WriteDataMapPacket.
// The RPC command is the GUESTRPCPKT_FIELD_PAYLOAD string; the response is a
// DataMap packet with FIELD_TYPE=TYPE_DATA and FIELD_PAYLOAD="1 result"|"0 err".
func (s *GuestRPCServer) serveConnDataMap(r io.Reader, w io.Writer) {
	for {
		select {
		case <-s.done:
			return
		default:
		}

		payload, fastClose, err := toolbox.ReadDataMapPacket(r)
		if err != nil {
			return // EOF or peer closed
		}

		rpcCmd := strings.TrimRight(string(payload), "\x00")
		if Trace {
			log.Printf("GuestRPCServer %s: [vsock/DataMap] cmd=%q fastClose=%v",
				s.vm.Name, rpcCmd, fastClose)
		}

		resp := s.dispatch(rpcCmd)
		if werr := toolbox.WriteDataMapPacket(w, []byte(resp), false); werr != nil {
			return
		}

		if fastClose {
			// Client set FAST_CLOSE: it will close the connection after reading
			// our response; no need to loop.
			return
		}
	}
}

// dispatch parses and executes an RPCI command.
func (s *GuestRPCServer) dispatch(cmd string) string {
	switch {
	case strings.HasPrefix(cmd, "info-get "):
		return s.infoGet(strings.TrimPrefix(cmd, "info-get "))

	case strings.HasPrefix(cmd, "info-set "):
		rest := strings.TrimPrefix(cmd, "info-set ")
		// "info-set guestinfo.key value with spaces"
		parts := strings.SplitN(rest, " ", 2)
		if len(parts) == 2 {
			return s.infoSet(parts[0], parts[1])
		}
		return "0 Bad syntax for info-set"

	case strings.HasPrefix(cmd, "tools.capability."),
		strings.HasPrefix(cmd, "tools.set-state"),
		strings.HasPrefix(cmd, "tools.set.version"),
		strings.HasPrefix(cmd, "Capabilities_Register"):
		// Acknowledge without semantic effect.
		return "1 "

	default:
		return "0 Unknown command"
	}
}

// infoGet reads key from the VM's ExtraConfig under the object lock.
func (s *GuestRPCServer) infoGet(key string) string {
	if !isGuestReadable(key) {
		return "0 Permission denied"
	}

	s.rpcMu.RLock()
	defer s.rpcMu.RUnlock()

	var result string
	found := false

	s.ctx.WithLock(s.vm, func() {
		for _, opt := range s.vm.Config.ExtraConfig {
			v := opt.GetOptionValue()
			if v.Key == key {
				if v.Value != nil {
					result = fmt.Sprintf("%v", v.Value)
				}
				found = true
				break
			}
		}
	})

	if found {
		return "1 " + result
	}
	return "0 No value found"
}

// infoSet writes key=value to the VM's ExtraConfig and emits a PropertyChange.
func (s *GuestRPCServer) infoSet(key, value string) string {
	if !isGuestWritable(key) {
		return "0 Permission denied"
	}

	s.rpcMu.Lock()
	defer s.rpcMu.Unlock()

	// The value is guest-supplied and routinely carries cloud-init userdata,
	// so it is never logged; the key alone identifies the operation.
	if Trace {
		log.Printf("GuestRPCServer %s: info-set %s", s.vm.Name, key)
	}
	s.ctx.AutoUpdate(s.vm, func() {
		for _, opt := range s.vm.Config.ExtraConfig {
			v := opt.GetOptionValue()
			if v.Key == key {
				v.Value = value
				return
			}
		}
		s.vm.Config.ExtraConfig = append(s.vm.Config.ExtraConfig,
			&types.OptionValue{Key: key, Value: value})
	})
	return "1 "
}

// vmciEnabled reports whether the GuestRPC simulation layer should be
// activated for this VM: a per-VM unix socket server speaking RPCI, plus
// injection of the toolbox binary as /usr/bin/vmware-rpctool.
//
// A seccomp AF_VSOCK interception layer is future work and is NOT part of this
// implementation; nothing here provides AF_VSOCK inside a container.
// Returns true when RUN.vmci is set to "true"
// (case-insensitive, whitespace-trimmed) in extraConfig.
func vmciEnabled(extraConfig []types.BaseOptionValue) bool {
	for _, opt := range extraConfig {
		v := opt.GetOptionValue()
		if v.Key != "RUN.vmci" {
			continue
		}
		s, ok := v.Value.(string)
		return ok && strings.EqualFold(strings.TrimSpace(s), "true")
	}
	return false
}

// guestRPCVolumeMount returns the -v argument for bind-mounting the per-VM
// GuestRPC socket directory into the container.
//
// We bind-mount the DIRECTORY (not the socket file) so the mount remains visible
// when RUN.nestedContainers=true adds --tmpfs /run. A file bind-mount requires the
// target to already exist as a file inside a fresh tmpfs (it doesn't); a directory
// bind-mount creates a new submount at /run/vmware/ that overrides the empty tmpfs
// content at that path, making rpc.sock accessible at its well-known container path.
func guestRPCVolumeMount(socketDirPath string) string {
	return fmt.Sprintf("%s:%s", socketDirPath, GuestRPCSocketDir)
}

// hasVolumeDest returns true if any RUN.volume.* ExtraConfig entry has dest as
// its container-side mount destination.  The volume value format is
// "hostPath:containerPath[:options]".  Used to avoid injecting a duplicate
// auto-managed volume when the caller has already provided one.
func hasVolumeDest(extraConfig []types.BaseOptionValue, dest string) bool {
	for _, opt := range extraConfig {
		v := opt.GetOptionValue()
		if !strings.HasPrefix(v.Key, "RUN.volume.") {
			continue
		}
		s, ok := v.Value.(string)
		if !ok {
			continue
		}
		// Value format: "hostPath:containerPath" or "hostPath:containerPath:opts"
		parts := strings.SplitN(s, ":", 3)
		if len(parts) >= 2 && parts[1] == dest {
			return true
		}
	}
	return false
}
