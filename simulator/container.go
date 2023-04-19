/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package simulator

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"time"
)

var (
	shell = "/bin/sh"
)

func init() {
	if sh, err := exec.LookPath("bash"); err != nil {
		shell = sh
	}
}

// container provides methods to manage a container within a simulator VM lifecycle.
type container struct {
	id   string
	name string
}

type networkSettings struct {
	Gateway     string
	IPAddress   string
	IPPrefixLen int
	MacAddress  string
}

type containerDetails struct {
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

func extractNameAndUid(containerName string) (name string, uid string, err error) {
	parts := strings.Split(strings.TrimPrefix(containerName, "vcsim-"), "-")
	if len(parts) != 2 {
		err = fmt.Errorf("container name does not match expected vcsim-name-uid format: %s", containerName)
		return
	}

	return parts[0], parts[1], nil
}

type tarEntry struct {
	header  *tar.Header
	content []byte
}

// From https://docs.docker.com/engine/reference/commandline/cp/ :
// > It is not possible to copy certain system files such as resources under /proc, /sys, /dev, tmpfs, and mounts created by the user in the container.
// > However, you can still copy such files by manually running tar in docker exec.
// TODO: look at whether this can useful combine with populateVolume for the tar portion or whether the duplication is low enough to make sense
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

	errin := tw.Close()
	errout := stdin.Close()

	errwait := cmd.Wait()

	return errors.Join(err, errout, errin, errwait)
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
	errwait := cmd.Wait()

	return errors.Join(err, errwait)
}

// populateVolume creates a volume tightly associated with the specified container, populated with the provided files
// If the header.Size is omitted or set to zero, then len(content+1) is used.
func populateVolume(containerName string, volumeName string, files []tarEntry) error {
	image := os.Getenv("VCSIM_BUSYBOX")
	if image == "" {
		image = "busybox"
	}

	// TODO: do we need to cap name lengths so as not to overflow?
	name := sanitizeName(containerName) + "--" + sanitizeName(volumeName)
	cmd := exec.Command("docker", "run", "--rm", "-i", "-v", name+":"+"/"+name, image, "tar", "-C", "/"+name, "-xf", "-")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
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

	err1 := tw.Close()
	err2 := stdin.Close()
	err = errors.Join(err1, err2)

	if err3 := cmd.Wait(); err3 != nil {
		stderr := ""
		if xerr, ok := err.(*exec.ExitError); ok {
			stderr = string(xerr.Stderr)
		}
		log.Printf("%s %s: %s %s", name, cmd.Args, err, stderr)

		return errors.Join(err, err3)
	}

	return err
}

// create
//   - name - pretty name, eg. vm name
//   - id - uuid or similar - this is merged into container name rather than dictating containerID
//   - networks - set of bridges to connect the container to
//   - volumes - colon separated tuple of volume name to mount path. Passed directly to docker via -v so mount options can be postfixed.
//   - env - array of environment vairables in name=value form
//   - image - the name of the container image to use, including tag
//   - args - the command+args to pass to the container
func create(ctx *Context, name string, id string, networks []string, volumes []string, ports []string, env []string, image string, args []string) (*container, error) {
	if len(image) == 0 {
		return nil, errors.New("cannot create container backing without an image")
	}

	var c container
	c.name = constructContainerName(name, id)

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

// populateVolume takes the specified files and writes them into a volume named for the container.
func (c *container) populateVolume(name string, files []tarEntry) error {
	return populateVolume(c.name, name, files)
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
	if c.id == "" {
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
	return
}

// start
//   - if the container already exists, start it or unpause it.
func (c *container) start(ctx *Context) error {
	if c.id == "" {
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
	if c.id == "" {
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
	if c.id == "" {
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
	if c.id == "" {
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
	if c.id == "" {
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
// returns:
//
//	err - joined err from deletion of container and matching volume name
func (c *container) remove(ctx *Context) error {
	if c.id == "" {
		// consider absence success
		return nil
	}

	cmd := exec.Command("docker", "rm", "-v", "-f", c.id)
	err := cmd.Run()
	if err != nil {
		log.Printf("%s %s: %s", c.name, cmd.Args, err)
	}

	// TODO: modify this to list all volumes with c.name prefix and delete them - necessary because populateVolume was generalized
	cmd = exec.Command("docker", "volume", "rm", "-f", c.name)
	err2 := cmd.Run()
	if err2 != nil {
		log.Printf("%s %s: %s", c.name, cmd.Args, err2)
	}

	combinedErr := errors.Join(err, err2)

	if combinedErr == nil {
		c.id = ""
	}

	return combinedErr
}

// watchContainer monitors the underlying container and updates
// properties based on the container status. This occurs until either
// the container or the VM is removed.
// returns:
//
//	err - uninitializedContainer error - if c.id is empty
func (c *container) watchContainer(ctx *Context, updateFn func(*Context, *containerDetails, *container) error) error {
	if c.id == "" {
		return uninitializedContainer(errors.New("Attempt to watch uninitialized container"))
	}

	// Update the VM from the container at regular intervals until the done
	// channel is closed.
	go func() {
		inspectInterval := time.Duration(5 * time.Second)
		if d, err := time.ParseDuration(os.Getenv("VCSIM_INSPECT_INTERVAL")); err == nil {
			inspectInterval = d
		}
		ticker := time.NewTicker(inspectInterval)

		for {
			select {
			case <-ticker.C:
				_, details, err := c.inspect()
				var rmErr error
				var removing bool
				if _, ok := err.(uninitializedContainer); ok {
					removing = true
					rmErr = c.remove(ctx)
				}

				updateErr := updateFn(ctx, &details, c)
				err = errors.Join(rmErr, updateErr)
				if removing && err == nil {
					// if we don't succeed we want to re-try
					ticker.Stop()
					return
				}
				// TODO: log err?
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
