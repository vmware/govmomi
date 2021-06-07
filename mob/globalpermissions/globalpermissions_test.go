package globalpermissions_test

import (
	"context"
	"fmt"
	"github.com/vmware/govmomi/mob/globalpermissions"
	"github.com/vmware/govmomi/mob/mobclient"
	"github.com/vmware/govmomi/vim25/soap"
	"net/url"
	"testing"
)

func TestManager_GlobalPermissionsWorkflow(t *testing.T) {

	userinfo := url.UserPassword("user", "pass")
	url := url.URL{Host: "vc_ip", Scheme: "https"}

	mob := mobclient.NewClient(soap.NewClient(&url, true))
	err := mob.Login(context.Background(), userinfo)
	if err != nil {
		t.Fatal("error while logging in ", err)
	}

	manager := globalpermissions.NewManager(mob)
	permissions, err := manager.ListGlobalPermission()
	if err != nil {
		t.Error(err)
	}
	for _, val := range permissions {
		fmt.Println(val)
	}

	err = manager.CreateGlobalPermission("user/group", true, false, 100)
	if err != nil {
		t.Error(err)
	}
}
