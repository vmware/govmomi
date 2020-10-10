/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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

package main

import (
	"crypto/tls"
	"expvar"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"github.com/google/uuid"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/types"

	// Register vcsim optional endpoints
	_ "github.com/vmware/govmomi/cns/simulator"
	_ "github.com/vmware/govmomi/lookup/simulator"
	_ "github.com/vmware/govmomi/pbm/simulator"
	_ "github.com/vmware/govmomi/sts/simulator"
	_ "github.com/vmware/govmomi/vapi/cluster/simulator"
	_ "github.com/vmware/govmomi/vapi/namespace/simulator"
	_ "github.com/vmware/govmomi/vapi/simulator"
)

func main() {
	model := simulator.VPX()

	flag.IntVar(&model.Datacenter, "dc", model.Datacenter, "Number of datacenters")
	flag.IntVar(&model.Cluster, "cluster", model.Cluster, "Number of clusters")
	flag.IntVar(&model.ClusterHost, "host", model.ClusterHost, "Number of hosts per cluster")
	flag.IntVar(&model.Host, "standalone-host", model.Host, "Number of standalone hosts")
	flag.IntVar(&model.Datastore, "ds", model.Datastore, "Number of local datastores")
	flag.IntVar(&model.Machine, "vm", model.Machine, "Number of virtual machines per resource pool")
	flag.IntVar(&model.Pool, "pool", model.Pool, "Number of resource pools per compute resource")
	flag.IntVar(&model.App, "app", model.App, "Number of virtual apps per compute resource")
	flag.IntVar(&model.Pod, "pod", model.Pod, "Number of storage pods per datacenter")
	flag.IntVar(&model.Portgroup, "pg", model.Portgroup, "Number of port groups")
	flag.IntVar(&model.PortgroupNSX, "pg-nsx", model.PortgroupNSX, "Number of NSX backed port groups")
	flag.IntVar(&model.OpaqueNetwork, "nsx", model.OpaqueNetwork, "Number of NSX backed opaque networks")
	flag.IntVar(&model.Folder, "folder", model.Folder, "Number of folders")
	flag.BoolVar(&model.Autostart, "autostart", model.Autostart, "Autostart model created VMs")
	v := &model.ServiceContent.About.ApiVersion
	flag.StringVar(v, "api-version", *v, "API version")

	isESX := flag.Bool("esx", false, "Simulate standalone ESX")
	isTLS := flag.Bool("tls", true, "Enable TLS")
	cert := flag.String("tlscert", "", "Path to TLS certificate file")
	key := flag.String("tlskey", "", "Path to TLS key file")
	env := flag.String("E", "-", "Output vcsim variables to the given fifo or stdout")
	listen := flag.String("l", "127.0.0.1:8989", "Listen address for vcsim")
	user := flag.String("username", "", "Login username for vcsim (any username allowed by default)")
	pass := flag.String("password", "", "Login password for vcsim (any password allowed by default)")
	tunnel := flag.Int("tunnel", -1, "SDK tunnel port")
	flag.BoolVar(&simulator.Trace, "trace", simulator.Trace, "Trace SOAP to -trace-file")
	trace := flag.String("trace-file", "", "Trace output file (defaults to stderr)")
	stdinExit := flag.Bool("stdinexit", false, "Press any key to exit")
	dir := flag.String("load", "", "Load model from directory")

	flag.IntVar(&model.DelayConfig.Delay, "delay", model.DelayConfig.Delay, "Method response delay across all methods")
	methodDelayP := flag.String("method-delay", "", "Delay per method on the form 'method1:delay1,method2:delay2...'")
	flag.Float64Var(&model.DelayConfig.DelayJitter, "delay-jitter", model.DelayConfig.DelayJitter, "Delay jitter coefficient of variation (tip: 0.5 is a good starting value)")

	flag.Parse()

	if *trace != "" {
		var err error
		simulator.TraceFile, err = os.Create(*trace)
		if err != nil {
			log.Fatal(err)
		}
		simulator.Trace = true
	}

	methodDelay := *methodDelayP
	u := &url.URL{Host: *listen}
	if *user != "" {
		u.User = url.UserPassword(secret(user), secret(pass))
	}

	switch flag.Arg(0) {
	case "uuidgen": // util-linux not installed on Travis CI
		fmt.Println(uuid.New().String())
		return
	}

	if methodDelay != "" {
		m := make(map[string]int)
		for _, s := range strings.Split(methodDelay, ",") {
			s = strings.TrimSpace(s)
			tuples := strings.Split(s, ":")
			if len(tuples) == 2 {
				key := tuples[0]
				value, err := strconv.Atoi(tuples[1])
				if err != nil {
					log.Fatalf("Incorrect format of method-delay argument: %s", err)
				}
				m[key] = value
			} else {
				log.Fatal("Incorrect method delay format.")
			}
		}
		model.DelayConfig.MethodDelay = m
	}

	var err error
	out := os.Stdout

	if *env != "-" {
		out, err = os.OpenFile(*env, os.O_WRONLY, 0)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = updateHostTemplate(u.Host); err != nil {
		log.Fatal(err)
	}

	if *isESX {
		opts := model
		model = simulator.ESX()
		// Preserve options that also apply to ESX
		model.Datastore = opts.Datastore
		model.Machine = opts.Machine
		model.Autostart = opts.Autostart
		model.DelayConfig.Delay = opts.DelayConfig.Delay
		model.DelayConfig.MethodDelay = opts.DelayConfig.MethodDelay
		model.DelayConfig.DelayJitter = opts.DelayConfig.DelayJitter
	}

	tag := " (govmomi simulator)"
	model.ServiceContent.About.Name += tag
	model.ServiceContent.About.OsType = runtime.GOOS + "-" + runtime.GOARCH

	esx.HostSystem.Summary.Hardware.Vendor += tag

	if *dir == "" {
		err = model.Create()
	} else {
		err = model.Load(*dir)
	}
	if err != nil {
		log.Fatal(err)
	}

	model.Service.RegisterEndpoints = true
	model.Service.Listen = u
	if *isTLS {
		model.Service.TLS = new(tls.Config)
		if *cert != "" {
			c, err := tls.LoadX509KeyPair(*cert, *key)
			if err != nil {
				log.Fatal(err)
			}

			model.Service.TLS.Certificates = []tls.Certificate{c}
		}
	}

	expvar.Publish("vcsim", expvar.Func(func() interface{} {
		count := model.Count()

		return struct {
			Registry *simulator.Registry
			Model    *simulator.Model
		}{
			simulator.Map,
			&count,
		}
	}))

	model.Service.ServeMux = http.DefaultServeMux // expvar.init registers "/debug/vars" with the DefaultServeMux

	s := model.Service.NewServer()

	if *tunnel >= 0 {
		s.Tunnel = *tunnel
		if err := s.StartTunnel(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Fprintf(out, "export GOVC_URL=%s GOVC_SIM_PID=%d\n", s.URL, os.Getpid())
	if out != os.Stdout {
		err = out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	if *stdinExit {
		fmt.Fprintf(out, "Press any key to exit")
		go func() {
			os.Stdin.Read(make([]byte, 1))
			sig <- syscall.SIGTERM
		}()
	}

	<-sig

	model.Remove()

	if *trace != "" {
		_ = simulator.TraceFile.Close()
	}
}

func updateHostTemplate(ip string) error {
	addr, port, err := net.SplitHostPort(ip)
	if err != nil {
		return err
	}
	if port != "0" { // server starts after the model is created, skipping auto-selected ports for now
		n, err := strconv.Atoi(port)
		if err != nil {
			return err
		}
		esx.HostSystem.Summary.Config.Port = int32(n)
	}

	nics := [][]types.HostVirtualNic{
		esx.HostConfigInfo.Network.Vnic,
		esx.HostConfigInfo.Vmotion.NetConfig.CandidateVnic,
	}

	for _, nic := range esx.HostConfigInfo.VirtualNicManagerInfo.NetConfig {
		nics = append(nics, nic.CandidateVnic)
	}

	for _, nic := range nics {
		for i := range nic {
			nic[i].Spec.Ip.IpAddress = addr // replace "127.0.0.1" with $addr
		}
	}

	return nil
}

func secret(s *string) string {
	val, err := session.Secret(*s)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
