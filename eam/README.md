# ESX Agent Manager (EAM)

ESX Agent Manager ([EAM](https://vdc-download.vmware.com/vmwb-repository/dcr-public/3d076a12-29a2-4d17-9269-cb8150b5a37f/8b5969e2-1a66-4425-af17-feff6d6f705d/SDK/eam/doc/index.html)) is a long-lived service on vCenter that acts as an intermediary for provisioning agent virtual machines and VIB modules on behalf of the user.

## HowTo

Please refer to [`simulator/simulator_test.go`](./simulator/simulator_test.go) for an example on how to use this package.

## Simulator

The [`simulator`](./simulator/) package provides an EAM simulator that can even simulate the lifecycle of agent VMs using containers.
