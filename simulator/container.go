// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	shell      = "/bin/sh"
	eventWatch eventWatcher
)

const (
	deleteWithContainer = "lifecycle=container"
	createdByVcsim      = "createdBy=vcsim"
)

func init() {
	if sh, err := exec.LookPath("bash"); err != nil {
		shell = sh
	}
}

type eventWatcher struct {
	sync.Mutex

	stdin   io.WriteCloser
	stdout  io.ReadCloser
	process *os.Process

	// watches is a map of container IDs to container objects
	watches map[string]*container
}

// container provides methods to manage a container within a simulator VM lifecycle.
type container struct {
	sync.Mutex

	id   string
	name string

	cancelWatch context.CancelFunc
	changes     chan struct{}
}

type networkSettings struct {
	Gateway     string
	IPAddress   string
	IPPrefixLen int
	MacAddress  string
}

type containerDetails struct {
	Config struct {
		Hostname   string
		Domainname string
		DNS        []string `json:"dns"`
	}
	State struct {
		Running bool
		Paused  bool
	}
	NetworkSettings struct {
		networkSettings
		Networks map[string]networkSettings
	}
}

type unknownContainer error
type uninitializedContainer error

var sanitizeNameRx = regexp.MustCompile(`[\(\)\s]`)

func sanitizeName(name string) string {
	return sanitizeNameRx.ReplaceAllString(name, "-")
}

func constructContainerName(name, uid string) string {
	return fmt.Sprintf("vcsim-%s-%s", sanitizeName(name), uid)
}

func constructVolumeName(containerName, uid, volumeName string) string {
	return constructContainerName(containerName, uid) + "--" + sanitizeName(volumeName)
}

func prefixToMask(prefix int) string {
	mask := net.CIDRMask(prefix, 32)
	return fmt.Sprintf("%d.%d.%d.%d", mask[0], mask[1], mask[2], mask[3])
}

type tarEntry struct {
	header  *tar.Header
	content []byte
}

// From https://docs.docker.com/engine/reference/commandline/cp/ :
// > It is not possible to copy certain system files such as resources under /proc, /sys, /dev, tmpfs, and mounts created by the user in the container.
// > However, you can still copy such files by manually running tar in docker exec.
func copyToGuest(id string, dest string, length int64, reader io.Reader) error {
	cmd := exec.Command("docker", "exec", "-i", id, "tar", "Cxf", path.Dir(dest), "-")
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	tw := tar.NewWriter(stdin)
	_ = tw.WriteHeader(&tar.Header{
		Name:    path.Base(dest),
		Size:    length,
		Mode:    0444,
		ModTime: time.Now(),
	})

	_, err = io.Copy(tw, reader)

	twErr := tw.Close()
	stdinErr := stdin.Close()

	waitErr := cmd.Wait()

	if err != nil || twErr != nil || stdinErr != nil || waitErr != nil {
		return fmt.Errorf("copy: {%s}, tw: {%s}, stdin: {%s}, wait: {%s}", err, twErr, stdinErr, waitErr)
	}

	return nil
}

func copyFromGuest(id string, src string, sink func(int64, io.Reader) error) error {
	cmd := exec.Command("docker", "exec", id, "tar", "Ccf", path.Dir(src), "-", path.Base(src))
	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}

	tr := tar.NewReader(stdout)
	header, err := tr.Next()
	if err != nil {
		return err
	}

	err = sink(header.Size, tr)
	waitErr := cmd.Wait()

	if err != nil || waitErr != nil {
		return fmt.Errorf("err: {%s}, wait: {%s}", err, waitErr)
	}

	return nil
}

// createVolume creates a volume populated with the provided files
// If the header.Size is omitted or set to zero, then len(content+1) is used.
// Docker appears to treat this volume create command as idempotent so long as it's identical
// to an existing volume, so we can use this both for creating volumes inline in container create (for labelling) and
// for population after.
// returns:
//
//	uid - string
//	err - error or nil
func createVolume(volumeName string, labels []string, files []tarEntry) (string, error) {
	image := os.Getenv("VCSIM_BUSYBOX")
	if image == "" {
		image = "busybox"
	}

	name := sanitizeName(volumeName)
	uid := ""

	// label the volume if specified - this requires the volume be created before use
	if len(labels) > 0 {
		run := []string{"volume", "create"}
		for i := range labels {
			run = append(run, "--label", labels[i])
		}
		run = append(run, name)
		cmd := exec.Command("docker", run...)
		out, err := cmd.Output()
		if err != nil {
			return "", err
		}
		uid = strings.TrimSpace(string(out))

		if name == "" {
			name = uid
		}
	}

	run := []string{"run", "--rm", "-i"}
	run = append(run, "-v", name+":/"+name)
	run = append(run, image, "tar", "-C", "/"+name, "-xf", "-")
	cmd := exec.Command("docker", run...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return uid, err
	}

	err = cmd.Start()
	if err != nil {
		return uid, err
	}

	tw := tar.NewWriter(stdin)

	for _, file := range files {
		header := file.header

		if header.Size == 0 && len(file.content) > 0 {
			header.Size = int64(len(file.content))
		}

		if header.ModTime.IsZero() {
			header.ModTime = time.Now()
		}

		if header.Mode == 0 {
			header.Mode = 0444
		}

		tarErr := tw.WriteHeader(header)
		if tarErr == nil {
			_, tarErr = tw.Write(file.content)
		}
	}

	err = nil
	twErr := tw.Close()
	stdinErr := stdin.Close()
	if twErr != nil || stdinErr != nil {
		err = fmt.Errorf("tw: {%s}, stdin: {%s}", twErr, stdinErr)
	}

	if waitErr := cmd.Wait(); waitErr != nil {
		stderr := ""
		if xerr, ok := waitErr.(*exec.ExitError); ok {
			stderr = string(xerr.Stderr)
		}
		log.Printf("%s %s: %s %s", name, cmd.Args, waitErr, stderr)

		err = fmt.Errorf("%s, wait: {%s}", err, waitErr)
		return uid, err
	}

	return uid, err
}

func getBridge(bridgeName string) (string, error) {
	// {"CreatedAt":"2023-07-11 19:22:25.45027052 +0000 UTC","Driver":"bridge","ID":"fe52c7502c5d","IPv6":"false","Internal":"false","Labels":"goodbye=,hello=","Name":"testnet","Scope":"local"}
	// podman has distinctly different fields at v4.4.1 so commented out fields that don't match. We only actually care about ID
	type bridgeNet struct {
		// CreatedAt string
		Driver string
		ID     string
		// IPv6      string
		// Internal  string
		// Labels    string
		Name string
		// Scope     string
	}

	// if the underlay bridge already exists, return that
	// we don't check for a specific label or similar so that it's possible to use a bridge created by other frameworks for composite testing
	var bridge bridgeNet
	cmd := exec.Command("docker", "network", "ls", "--format={{json .}}", "-f", fmt.Sprintf("name=%s$", bridgeName))
	out, err := cmd.Output()
	if err != nil {
		log.Printf("vcsim %s: %s, %s", cmd.Args, err, out)
		return "", err
	}

	// unfortunately docker returns an empty string not an empty json doc and podman returns '[]'
	// podman also returns an array of matches even when there's only one, so we normalize.
	str := strings.TrimSpace(string(out))
	str = strings.TrimPrefix(str, "[")
	str = strings.TrimSuffix(str, "]")
	if len(str) == 0 {
		return "", nil
	}

	err = json.Unmarshal([]byte(str), &bridge)
	if err != nil {
		log.Printf("vcsim %s: %s, %s", cmd.Args, err, str)
		return "", err
	}

	return bridge.ID, nil
}

// createBridge creates a bridge network if one does not already exist
// returns:
//
//	uid - string
//	err - error or nil
func createBridge(bridgeName string, labels ...string) (string, error) {

	id, err := getBridge(bridgeName)
	if err != nil {
		return "", err
	}

	if id != "" {
		return id, nil
	}

	run := []string{"network", "create", "--label", createdByVcsim}
	for i := range labels {
		run = append(run, "--label", labels[i])
	}
	run = append(run, bridgeName)

	cmd := exec.Command("docker", run...)
	out, err := cmd.Output()
	if err != nil {
		log.Printf("vcsim %s: %s: %s", cmd.Args, out, err)
		return "", err
	}

	// docker returns the ID regardless of whether you supply a name when creating the network, however
	// podman returns the pretty name, so we have to normalize
	id, err = getBridge(bridgeName)
	if err != nil {
		return "", err
	}

	return id, nil
}

// create
//   - name - pretty name, eg. vm name
//   - id - uuid or similar - this is merged into container name rather than dictating containerID
//   - networks - set of bridges to connect the container to
//   - volumes - colon separated tuple of volume name to mount path. Passed directly to docker via -v so mount options can be postfixed.
//   - env - array of environment vairables in name=value form
//   - optsAndImage - pass-though options and must include at least the container image to use, including tag if necessary
//   - args - the command+args to pass to the container
func create(ctx *Context, name string, id string, networks []string, volumes []string, ports []string, env []string, image string, args []string) (*container, error) {
	if len(image) == 0 {
		return nil, errors.New("cannot create container backing without an image")
	}

	var c container
	c.name = constructContainerName(name, id)
	c.changes = make(chan struct{})

	for i := range volumes {
		// we'll pre-create anonymous volumes, simply for labelling consistency
		volName := strings.Split(volumes[i], ":")
		createVolume(volName[0], []string{deleteWithContainer, "container=" + c.name}, nil)
	}

	// assemble env
	var dockerNet []string
	var dockerVol []string
	var dockerPort []string
	var dockerEnv []string

	for i := range env {
		dockerEnv = append(dockerEnv, "--env", env[i])
	}

	for i := range volumes {
		dockerVol = append(dockerVol, "-v", volumes[i])
	}

	for i := range ports {
		dockerPort = append(dockerPort, "-p", ports[i])
	}

	for i := range networks {
		dockerNet = append(dockerNet, "--network", networks[i])
	}

	run := []string{"docker", "create", "--name", c.name}
	run = append(run, dockerNet...)
	run = append(run, dockerVol...)
	run = append(run, dockerPort...)
	run = append(run, dockerEnv...)
	run = append(run, image)
	run = append(run, args...)

	// this combines all the run options into a single string that's passed to /bin/bash -c as the single argument to force bash parsing.
	// TODO: make this configurable behaviour so users also have the option of not escaping everything for bash
	cmd := exec.Command(shell, "-c", strings.Join(run, " "))
	out, err := cmd.Output()
	if err != nil {
		stderr := ""
		if xerr, ok := err.(*exec.ExitError); ok {
			stderr = string(xerr.Stderr)
		}
		log.Printf("%s %s: %s %s", name, cmd.Args, err, stderr)

		return nil, err
	}

	c.id = strings.TrimSpace(string(out))

	return &c, nil
}

// createVolume takes the specified files and writes them into a volume named for the container.
func (c *container) createVolume(name string, labels []string, files []tarEntry) (string, error) {
	return createVolume(c.name+"--"+name, append(labels, "container="+c.name), files)
}

// inspect retrieves and parses container properties into directly usable struct
// returns:
//
//	out - the stdout of the command
//	detail - basic struct populated with container details
//	err:
//		* if c.id is empty, or docker returns "No such object", will return an uninitializedContainer error
//		* err from either execution or parsing of json output
func (c *container) inspect() (out []byte, detail containerDetails, err error) {
	c.Lock()
	id := c.id
	c.Unlock()

	if id == "" {
		err = uninitializedContainer(errors.New("inspect of uninitialized container"))
		return
	}

	var details []containerDetails

	cmd := exec.Command("docker", "inspect", c.id)
	out, err = cmd.Output()
	if eErr, ok := err.(*exec.ExitError); ok {
		if strings.Contains(string(eErr.Stderr), "No such object") {
			err = uninitializedContainer(errors.New("inspect of uninitialized container"))
		}
	}

	if err != nil {
		return
	}

	if err = json.NewDecoder(bytes.NewReader(out)).Decode(&details); err != nil {
		return
	}

	if len(details) != 1 {
		err = fmt.Errorf("multiple containers (%d) match ID: %s", len(details), c.id)
		return
	}

	detail = details[0]

	// DNS setting
	f, oerr := os.Open("/etc/docker/daemon.json")
	if oerr != nil {
		return
	}
	err = json.NewDecoder(f).Decode(&detail.Config)
	_ = f.Close()

	return
}

// start
//   - if the container already exists, start it or unpause it.
func (c *container) start(ctx *Context) error {
	c.Lock()
	id := c.id
	c.Unlock()

	if id == "" {
		return uninitializedContainer(errors.New("start of uninitialized container"))
	}

	start := "start"
	_, detail, err := c.inspect()
	if err != nil {
		return err
	}

	if detail.State.Paused {
		start = "unpause"
	}

	cmd := exec.Command("docker", start, c.id)
	err = cmd.Run()
	if err != nil {
		log.Printf("%s %s: %s", c.name, cmd.Args, err)
	}

	return err
}

// pause the container (if any) for the given vm.
func (c *container) pause(ctx *Context) error {
	c.Lock()
	id := c.id
	c.Unlock()

	if id == "" {
		return uninitializedContainer(errors.New("pause of uninitialized container"))
	}

	cmd := exec.Command("docker", "pause", c.id)
	err := cmd.Run()
	if err != nil {
		log.Printf("%s %s: %s", c.name, cmd.Args, err)
	}

	return err
}

// restart the container (if any) for the given vm.
func (c *container) restart(ctx *Context) error {
	c.Lock()
	id := c.id
	c.Unlock()

	if id == "" {
		return uninitializedContainer(errors.New("restart of uninitialized container"))
	}

	cmd := exec.Command("docker", "restart", c.id)
	err := cmd.Run()
	if err != nil {
		log.Printf("%s %s: %s", c.name, cmd.Args, err)
	}

	return err
}

// stop the container (if any) for the given vm.
func (c *container) stop(ctx *Context) error {
	c.Lock()
	id := c.id
	c.Unlock()

	if id == "" {
		return uninitializedContainer(errors.New("stop of uninitialized container"))
	}

	cmd := exec.Command("docker", "stop", c.id)
	err := cmd.Run()
	if err != nil {
		log.Printf("%s %s: %s", c.name, cmd.Args, err)
	}

	return err
}

// exec invokes the specified command, with executable being the first of the args, in the specified container
// returns
//
//	 string - combined stdout and stderr from command
//	 err
//			* uninitializedContainer error - if c.id is empty
//		   	* err from cmd execution
func (c *container) exec(ctx *Context, args []string) (string, error) {
	c.Lock()
	id := c.id
	c.Unlock()

	if id == "" {
		return "", uninitializedContainer(errors.New("exec into uninitialized container"))
	}

	args = append([]string{"exec", c.id}, args...)
	cmd := exec.Command("docker", args...)
	res, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("%s: %s (%s)", c.name, cmd.Args, string(res))
		return "", err
	}

	return strings.TrimSpace(string(res)), nil
}

// remove the container (if any) for the given vm. Considers removal of an uninitialized container success.
// Also removes volumes and networks that indicate they are lifecycle coupled with this container.
// returns:
//
//	err - joined err from deletion of container and any volumes or networks that have coupled lifecycle
func (c *container) remove(ctx *Context) error {
	c.Lock()
	defer c.Unlock()

	if c.id == "" {
		// consider absence success
		return nil
	}

	cmd := exec.Command("docker", "rm", "-v", "-f", c.id)
	err := cmd.Run()
	if err != nil {
		log.Printf("%s %s: %s", c.name, cmd.Args, err)
		return err
	}

	cmd = exec.Command("docker", "volume", "ls", "-q", "--filter", "label=container="+c.name, "--filter", "label="+deleteWithContainer)
	volumesToReap, lsverr := cmd.Output()
	if lsverr != nil {
		log.Printf("%s %s: %s", c.name, cmd.Args, lsverr)
	}
	log.Printf("%s volumes: %s", c.name, volumesToReap)

	var rmverr error
	if len(volumesToReap) > 0 {
		run := []string{"volume", "rm", "-f"}
		run = append(run, strings.Split(string(volumesToReap), "\n")...)
		cmd = exec.Command("docker", run...)
		out, rmverr := cmd.Output()
		if rmverr != nil {
			log.Printf("%s %s: %s, %s", c.name, cmd.Args, rmverr, out)
		}
	}

	cmd = exec.Command("docker", "network", "ls", "-q", "--filter", "label=container="+c.name, "--filter", "label="+deleteWithContainer)
	networksToReap, lsnerr := cmd.Output()
	if lsnerr != nil {
		log.Printf("%s %s: %s", c.name, cmd.Args, lsnerr)
	}

	var rmnerr error
	if len(networksToReap) > 0 {
		run := []string{"network", "rm", "-f"}
		run = append(run, strings.Split(string(volumesToReap), "\n")...)
		cmd = exec.Command("docker", run...)
		rmnerr = cmd.Run()
		if rmnerr != nil {
			log.Printf("%s %s: %s", c.name, cmd.Args, rmnerr)
		}
	}

	if err != nil || lsverr != nil || rmverr != nil || lsnerr != nil || rmnerr != nil {
		return fmt.Errorf("err: {%s}, lsverr: {%s}, rmverr: {%s}, lsnerr:{%s}, rmerr: {%s}", err, lsverr, rmverr, lsnerr, rmnerr)
	}

	if c.cancelWatch != nil {
		c.cancelWatch()
		eventWatch.ignore(c)
	}
	c.id = ""
	return nil
}

// updated is a simple trigger allowing a caller to indicate that something has likely changed about the container
// and interested parties should re-inspect as needed.
func (c *container) updated() {
	consolidationWindow := 250 * time.Millisecond
	if d, err := time.ParseDuration(os.Getenv("VCSIM_EVENT_CONSOLIDATION_WINDOW")); err == nil {
		consolidationWindow = d
	}

	select {
	case c.changes <- struct{}{}:
		time.Sleep(consolidationWindow)
		// as this is only a hint to avoid waiting for the full inspect interval, we don't care about accumulating
		// multiple triggers. We do pause to allow large numbers of sequential updates to consolidate
	default:
	}
}

// watchContainer monitors the underlying container and updates
// properties based on the container status. This occurs until either
// the container or the VM is removed.
// returns:
//
//	err - uninitializedContainer error - if c.id is empty
func (c *container) watchContainer(ctx *Context, updateFn func(*containerDetails, *container) error) error {
	c.Lock()
	defer c.Unlock()

	if c.id == "" {
		return uninitializedContainer(errors.New("Attempt to watch uninitialized container"))
	}

	eventWatch.watch(c)

	cancelCtx, cancelFunc := context.WithCancel(ctx)
	c.cancelWatch = cancelFunc

	// Update the VM from the container at regular intervals until the done
	// channel is closed.
	go func() {
		inspectInterval := 10 * time.Second
		if d, err := time.ParseDuration(os.Getenv("VCSIM_INSPECT_INTERVAL")); err == nil {
			inspectInterval = d
		}
		ticker := time.NewTicker(inspectInterval)

		update := func() {
			_, details, err := c.inspect()
			var rmErr error
			var removing bool
			if _, ok := err.(uninitializedContainer); ok {
				removing = true
				rmErr = c.remove(ctx)
			}

			updateErr := updateFn(&details, c)
			// if we don't succeed we want to re-try
			if removing && rmErr == nil && updateErr == nil {
				ticker.Stop()
				return
			}
			if updateErr != nil {
				log.Printf("vcsim container watch: %s %s", c.id, updateErr)
			}
		}

		for {
			select {
			case <-c.changes:
				update()
			case <-ticker.C:
				update()
			case <-cancelCtx.Done():
				return
			}
		}
	}()

	return nil
}

func (w *eventWatcher) watch(c *container) {
	w.Lock()
	defer w.Unlock()

	if w.watches == nil {
		w.watches = make(map[string]*container)
	}

	w.watches[c.id] = c

	if w.stdin == nil {
		cmd := exec.Command("docker", "events", "--format", "'{{.ID}}'", "--filter", "Type=container")
		w.stdout, _ = cmd.StdoutPipe()
		w.stdin, _ = cmd.StdinPipe()
		err := cmd.Start()
		if err != nil {
			log.Printf("docker event watcher: %s %s", cmd.Args, err)
			w.stdin = nil
			w.stdout = nil
			w.process = nil

			return
		}

		w.process = cmd.Process

		go w.monitor()
	}
}

func (w *eventWatcher) ignore(c *container) {
	w.Lock()

	delete(w.watches, c.id)

	if len(w.watches) == 0 && w.stdin != nil {
		w.stop()
	}

	w.Unlock()
}

func (w *eventWatcher) monitor() {
	w.Lock()
	watches := len(w.watches)
	w.Unlock()

	if watches == 0 {
		return
	}

	scanner := bufio.NewScanner(w.stdout)
	for scanner.Scan() {
		id := strings.TrimSpace(scanner.Text())

		w.Lock()
		container := w.watches[id]
		w.Unlock()

		if container != nil {
			// this is called in a routine to allow an event consolidation window
			go container.updated()
		}
	}
}

func (w *eventWatcher) stop() {
	if w.stdin != nil {
		w.stdin.Close()
		w.stdin = nil
	}
	if w.stdout != nil {
		w.stdout.Close()
		w.stdout = nil
	}
	w.process.Kill()
}
