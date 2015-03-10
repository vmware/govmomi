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

package session

import (
	"net/url"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type Manager struct {
	roundTripper   soap.RoundTripper
	serviceContent types.ServiceContent
	userSession    *types.UserSession
}

func NewManager(rt soap.RoundTripper, sc types.ServiceContent) *Manager {
	m := Manager{
		roundTripper:   rt,
		serviceContent: sc,
	}

	return &m
}

func (sm Manager) Reference() types.ManagedObjectReference {
	return *sm.serviceContent.SessionManager
}

func (sm *Manager) Login(ctx context.Context, u *url.Userinfo) error {
	req := types.Login{
		This: sm.Reference(),
	}

	if u != nil {
		req.UserName = u.Username()
		if pw, ok := u.Password(); ok {
			req.Password = pw
		}
	}

	login, err := methods.Login(ctx, sm.roundTripper, &req)
	if err != nil {
		return err
	}

	sm.userSession = &login.Returnval
	return err
}

func (sm *Manager) Logout(ctx context.Context) error {
	req := types.Logout{
		This: sm.Reference(),
	}

	_, err := methods.Logout(ctx, sm.roundTripper, &req)
	return err
}

func (sm *Manager) UserSession(ctx context.Context) (*types.UserSession, error) {
	if sm.userSession == nil {
		var mgr mo.SessionManager
		err := mo.RetrieveProperties(ctx, sm.roundTripper, sm.serviceContent.PropertyCollector, sm.Reference(), &mgr)
		if err != nil {
			return nil, err
		}
		sm.userSession = mgr.CurrentSession
	}

	return sm.userSession, nil
}

func (sm *Manager) SessionIsActive(ctx context.Context) (bool, error) {
	user, err := sm.UserSession(ctx)
	if err != nil {
		return false, err
	}

	req := types.SessionIsActive{
		This:      sm.Reference(),
		SessionID: user.Key,
		UserName:  user.UserName,
	}

	active, err := methods.SessionIsActive(ctx, sm.roundTripper, &req)
	if err != nil {
		return false, err
	}

	return active.Returnval, err
}
