// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/tls"
	"encoding/pem"
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
	_ "github.com/vmware/govmomi/eam/simulator"
	_ "github.com/vmware/govmomi/lookup/simulator"
	_ "github.com/vmware/govmomi/pbm/simulator"
	_ "github.com/vmware/govmomi/ssoadmin/simulator"
	_ "github.com/vmware/govmomi/sts/simulator"
	_ "github.com/vmware/govmomi/vapi/appliance/simulator"
	_ "github.com/vmware/govmomi/vapi/cis/tasks/simulator"
	_ "github.com/vmware/govmomi/vapi/cluster/simulator"
	_ "github.com/vmware/govmomi/vapi/crypto/simulator"
	_ "github.com/vmware/govmomi/vapi/esx/settings/simulator"
	_ "github.com/vmware/govmomi/vapi/namespace/simulator"
	_ "github.com/vmware/govmomi/vapi/simulator"
	_ "github.com/vmware/govmomi/vapi/vcenter/consumptiondomains/simulator"
	_ "github.com/vmware/govmomi/vapi/vm/simulator"
	_ "github.com/vmware/govmomi/vsan/simulator"
	_ "github.com/vmware/govmomi/vslm/simulator"
)

var (
	buildVersion string
	buildCommit  string
	buildDate    string
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
	hostIPBase := flag.String("host-ip-base", "", "Base IPv4 address to assign HostSystems from, incrementing by one per host (e.g. 10.20.8.0)")
	flag.BoolVar(&model.Autostart, "autostart", model.Autostart, "Autostart model created VMs")
	apiVersion := flag.String("api-version", "", "Override the API version reported by the server (default is inferred from -version)")
	version := flag.String("version", model.ServiceContent.About.Version, "Product version reported by the server and hosts")
	build := flag.String("build", "", "Build number reported by the server (default is inferred for known versions)")

	isESX := flag.Bool("esx", false, "Simulate standalone ESX")
	isTLS := flag.Bool("tls", true, "Enable TLS")
	cert := flag.String("tlscert", "", "Path to TLS certificate file")
	key := flag.String("tlskey", "", "Path to TLS key file")
	certOut := flag.String("tlscertout", "/tmp/vcsim.crt", "Write the TLS certificate to this file (empty to disable)")
	hostBuild := flag.String("host-build", "", "ESXi build number reported by hosts (default is inferred for known versions)")
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

	versionSet := false
	apiVersionSet := false
	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "version":
			versionSet = true
		case "api-version":
			apiVersionSet = true
		}
	})

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
	case "uuidgen":
		fmt.Println(uuid.New().String())
		os.Exit(0)
	case "version":
		fmt.Printf("Build Version: %s\n", buildVersion)
		fmt.Printf("Build Commit: %s\n", buildCommit)
		fmt.Printf("Build Date: %s\n", buildDate)
		os.Exit(0)
	}

	if methodDelay != "" {
		m := make(map[string]int)
		for s := range strings.SplitSeq(methodDelay, ",") {
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
		simulator.TaskDelay.MethodDelay = m
	}

	var err error

	if err = updateHostTemplate(u.Host); err != nil {
		log.Fatal(err)
	}

	if *hostIPBase != "" {
		ip := net.ParseIP(*hostIPBase).To4()
		if ip == nil {
			log.Fatalf("-host-ip-base: invalid IPv4 address: %s", *hostIPBase)
		}
		model.HostIPBase = ip
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

	// A supplied product version is shared by the server and its ESXi hosts.
	// Infer each API version and build number from that common version unless
	// explicitly overridden.
	if versionSet {
		model.ServiceContent.About.Version = *version
		esx.HostConfigInfo.Product.Version = *version
		esx.HostConfigInfo.Product.LicenseProductVersion = *version
		if serverAPI, hostAPI, ok := releaseAPIVersions(*version); ok {
			if *isESX {
				model.ServiceContent.About.ApiVersion = hostAPI
			} else {
				model.ServiceContent.About.ApiVersion = serverAPI
			}
			esx.HostConfigInfo.Product.ApiVersion = hostAPI
		} else {
			log.Printf("cannot infer API versions from %q; use -api-version to set the server API version", *version)
		}
	}
	if apiVersionSet {
		model.ServiceContent.About.ApiVersion = *apiVersion
	}
	if *build != "" {
		model.ServiceContent.About.Build = *build
	} else if versionSet {
		if build, ok := releaseBuild(*version, *isESX); ok {
			model.ServiceContent.About.Build = build
		} else {
			log.Printf("no known build number for version %q; use -build to set it", *version)
		}
	}
	if *hostBuild != "" {
		esx.HostConfigInfo.Product.Build = *hostBuild
	} else if versionSet {
		if build, ok := releaseBuild(*version, true); ok {
			esx.HostConfigInfo.Product.Build = build
		} else {
			log.Printf("no known ESXi build number for version %q; use -host-build to set it", *version)
		}
	}
	esx.HostConfigInfo.Product.FullName = fmt.Sprintf("%s %s build-%s", esx.HostConfigInfo.Product.Name, esx.HostConfigInfo.Product.Version, esx.HostConfigInfo.Product.Build)

	tag := " (govmomi simulator)"
	model.ServiceContent.About.FullName = fmt.Sprintf("%s %s build-%s%s", model.ServiceContent.About.Name, model.ServiceContent.About.Version, model.ServiceContent.About.Build, tag)
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

	expvar.Publish("vcsim", expvar.Func(func() any {
		count := model.Count()

		return struct {
			Registry *simulator.Registry `json:"registry"`
			Model    *simulator.Model    `json:"model"`
		}{
			model.Service.Context.Map,
			&count,
		}
	}))

	model.Service.ServeMux = http.DefaultServeMux // expvar.init registers "/debug/vars" with the DefaultServeMux

	s := model.Service.NewServer()
	if *isTLS && *certOut != "" {
		var certPEM []byte
		for _, der := range s.TLS.Certificates[0].Certificate {
			certPEM = append(certPEM, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})...)
		}
		if err := os.WriteFile(*certOut, certPEM, 0644); err != nil {
			log.Fatal(err)
		}
	}

	if *tunnel >= 0 {
		s.Tunnel = *tunnel
		if err := s.StartTunnel(); err != nil {
			log.Fatal(err)
		}
	}

	out := os.Stdout

	if *env != "-" {
		out, err = os.OpenFile(*env, os.O_WRONLY, 0)
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = fmt.Fprintf(out, "export GOVC_URL=%s GOVC_SIM_PID=%d\n", s.URL, os.Getpid())
	if err != nil {
		log.Fatal(err)
	}
	if out != os.Stdout {
		err = out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	if *stdinExit {
		fmt.Println("Press any key to exit")
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
	esx.HostSystem.Summary.ManagementServerIp = addr
	if port != "0" { // server starts after the model is created, skipping auto-selected ports for now
		n, err := strconv.ParseInt(port, 10, 32)
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

// releaseBuild returns the GA or commonly released build number for a product
// version. Patch builds can share a product version, so callers can override
// these values with -build or -host-build. The mappings are based on Broadcom
// KB 326316 (vCenter) and KB 316595 (ESXi).
func releaseBuild(version string, esx bool) (string, bool) {
	var builds map[string]string
	if esx {
		builds = map[string]string{
			"8.0.0":      "20513097",
			"8.0.1":      "21495797",
			"8.0.2":      "22380479",
			"8.0.3":      "24022510",
			"9.0.0":      "24755229",
			"9.0.0.0":    "24755229",
			"9.0.0.0100": "24813472",
			"9.0.1":      "24957456",
			"9.0.1.0":    "24957456",
			"9.0.2":      "25148076",
			"9.0.2.0":    "25148076",
			"9.0.2.0100": "25595025",
			"9.1":        "25370933",
			"9.1.0":      "25370933",
			"9.1.1.0":    "25714478",
		}
	} else {
		builds = map[string]string{
			"8.0.0":      "20519528",
			"8.0.1":      "21560480",
			"8.0.2":      "22385739",
			"8.0.3":      "24022515",
			"9.0.0":      "24755230",
			"9.0.0.0":    "24755230",
			"9.0.1":      "24957454",
			"9.0.1.0":    "24957454",
			"9.0.2":      "25148086",
			"9.0.2.0":    "25148086",
			"9.0.2.0100": "25629525",
			"9.1":        "25370922",
			"9.1.0":      "25370922",
			"9.1.1.0":    "25712839",
		}
	}
	build, ok := builds[version]
	return build, ok
}

func releaseAPIVersions(version string) (server, host string, ok bool) {
	parts := strings.Split(version, ".")
	if len(parts) < 2 {
		return "", "", false
	}

	for _, part := range parts[:min(len(parts), 3)] {
		if _, err := strconv.Atoi(part); err != nil {
			return "", "", false
		}
	}

	server = strings.Join(parts[:2], ".")
	hostParts := parts[:min(len(parts), 3)]
	if len(hostParts) < 3 {
		hostParts = append(hostParts, "0")
	}
	host = strings.Join(hostParts, ".") + ".0"
	return server, host, true
}
