package simulator

import (
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type CryptoManagerKmip struct {
	mo.CryptoManagerKmip
}

func (m *CryptoManagerKmip) init(r *Registry) {
	m.Enabled = true
	m.KmipServers = []types.KmipClusterInfo{
		{
			ClusterId:    types.KeyProviderId{Id: "kmipcluster"},
			Servers:      []types.KmipServerInfo{{Name: "kmipserver", Address: "localhost", Port: 5696}},
			UseAsDefault: true,
		},
	}

	root := r.content().CryptoManager // take the cryptomanager details from servicefolder
	m.CryptoManagerKmip.Self = *root
	//	m.CryptoManagerKmip =

}

func (m *CryptoManagerKmip) IsKmsClusterActive(clusterId string) bool {
	i := 0
	for i = 0; i < len(m.CryptoManagerKmip.KmipServers); i++ {
		if m.CryptoManagerKmip.KmipServers[i].ClusterId.Id == clusterId {
			return true
		}
	}

	return false
}

// Add create/get key methods here and expose them as API so that we can call these APIs from BVTs
func CreateKey() {

}

func GetKey() {

}

// create key and get key funcs here
