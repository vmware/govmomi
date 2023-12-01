package simulator

import (
	"fmt"
	"log"
	"testing"

	"github.com/vmware/govmomi/simulator/vpx"
)

func TestKmipCryptoManager(t *testing.T) {

	model := VPX()

	_ = New(NewServiceInstance(SpoofContext(), model.ServiceContent, model.RootFolder)) // 2nd pass panics w/o copying RoleList

	kmip := Map.Get(*vpx.ServiceContent.CryptoManager).(*CryptoManagerKmip)
	ans := kmip.IsKmsClusterActive("kmipcluster")
	fmt.Println(ans)
	log.Println(ans)

	cl, err := initClient()
	err = cl.kclient.Connect()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Client connected\n")

	err = cl.kclient.Connect()
	if err != nil {
		panic(err)
	}
	resp, err := cl.CreateKey()
	resp, err = cl.CreateKey()
	resp, err = cl.CreateKey()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Key created and error is \n")
	log.Println(resp)
	log.Println(err)

}
