/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package govmomi

import (
	"net/http"
	"net/url"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type SessionManager struct {
	types.ManagedObjectReference
	c           *Client
	userSession *types.UserSession
}

func NewSessionManager(c *Client, ref types.ManagedObjectReference) SessionManager {
	return SessionManager{
		ManagedObjectReference: ref,
		c: c,
	}
}

func (sm SessionManager) Reference() types.ManagedObjectReference {
	return sm.ManagedObjectReference
}

func (sm *SessionManager) Login(u url.Userinfo) error {
	req := types.Login{
		This: sm.Reference(),
	}

	req.UserName = u.Username()
	if pw, ok := u.Password(); ok {
		req.Password = pw
	}

	login, err := methods.Login(sm.c, &req)
	sm.userSession = &login.Returnval
	return err
}

func (sm *SessionManager) Logout() error {
	req := types.Logout{
		This: sm.Reference(),
	}

	_, err := methods.Logout(sm.c, &req)
	// we've logged out - lets close any idle connections
	t := sm.c.Client.Transport.(*http.Transport)
	t.CloseIdleConnections()

	return err
}

func (sm *SessionManager) UserSession() (*types.UserSession, error) {

	if sm.userSession == nil {
		var mgr mo.SessionManager
		err := mo.RetrieveProperties(sm.c, sm.c.ServiceContent.PropertyCollector, sm.Reference(), &mgr)
		if err != nil {
			return nil, err
		}
		sm.userSession = mgr.CurrentSession
	}

	return sm.userSession, nil
}

func (sm *SessionManager) SessionIsActive() (bool, error) {
	user, err := sm.UserSession()
	if err != nil {
		return false, err
	}

	req := types.SessionIsActive{
		This:      sm.Reference(),
		SessionID: user.Key,
		UserName:  user.UserName,
	}

	active, err := methods.SessionIsActive(sm.c, &req)
	if err != nil {
		return false, err
	}

	return active.Returnval, err
}
