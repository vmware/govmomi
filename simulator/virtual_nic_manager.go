package simulator


import (
        "github.com/vmware/govmomi/vim25/mo"
        "github.com/vmware/govmomi/vim25/types"
        "github.com/vmware/govmomi/vim25/soap"
        "github.com/vmware/govmomi/vim25/methods"
)

type HostVirtualNicManager struct {
        mo.HostVirtualNicManager

        Host *mo.HostSystem
}

func NewHostVirtualNicManager(host *mo.HostSystem) *HostVirtualNicManager{
        return &HostVirtualNicManager{
                Host: host,
                HostVirtualNicManager: mo.HostVirtualNicManager{
                        Info: types.HostVirtualNicManagerInfo{
                                NetConfig: []types.VirtualNicManagerNetConfig{
                                        {
                                                NicType: "management",
                                                CandidateVnic: []types.HostVirtualNic{
                                                        {
                                                                Spec:  types.HostVirtualNicSpec{
                                                                        Ip:  &types.HostIpConfig{
                                                                                IpAddress:  "192.168.0.1",
                                                                        },
                                                                },
                                                        },
                                                },
                                        },
                                },
                        },
                },
        }
}


func (s *HostVirtualNicManager) QueryNetConfig(c *types.QueryNetConfig) soap.HasFault {
        r := &methods.QueryNetConfigBody{}
        var Netconfig *types.VirtualNicManagerNetConfig



        for _, netconfig := range s.Info.NetConfig {
                if netconfig.NicType == c.NicType {
                        Netconfig = &netconfig
                        break
                }
        }
        r.Res = &types.QueryNetConfigResponse{
                Returnval: Netconfig,
        }
        return r
}
