# ESX Agent Manager (EAM) Simulator

This simulator package works with the existing vC Sim. Please see [`simulator_test.go`](simulator_test.go) for an example of how to use the EAM simulator with vC Sim.

## Use Docker to Simulate Agent VMs

It is possible to run the simulator test whereby the creation of agent VMs results in the creation of containers in Docker to simulate the lifecycle of the VMs. Docker must be installed and running, but other than that, simply set the value of the `AgentConfigInfo.OvfPackageUrl` field to a:

* non-empty string,
* that does not start with `./`, `/`, or `https?:`.

If those conditions are met, then the EAM simulator treats this value as the name of a Docker image, and will create a simualted VM with `RUN.container` key set to the value, causing the core vC Simulator to simulate the VM with a container based on the specified image.


### Example of Simulated Agent VMs

There is a test that provides an example of this simulated agent VM behavior. Again, Docker must be installed and running, but otherwise just execute the following:

```shell
go test -v ./simulator/ -powerOnVMs
```

The flag `-powerOnVMs` causes the the simulator test code to create the agencies with agent configuration information that uses the `nginx` container to simulate two VMs. The presence of the flag then causes the VMs to be powered on, and the test waits for the agent runtime information to report the VMs' power state and IP addresses.
