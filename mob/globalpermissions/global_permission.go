/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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

package globalpermissions

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/vmware/govmomi/mob/internal"
	"github.com/vmware/govmomi/mob/mobclient"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const globalPermissionsServiceMoid = "authorizationService"
const voidMethodResultMessage = "Method Invocation Result: void"

// Manager extends rest.Client, adding tag related methods.
type Manager struct {
	client *mobclient.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *mobclient.Client) *Manager {
	return &Manager{
		client: client,
	}
}

type GlobalPermission struct {
	Principal string
	Group     bool
	Propagate bool
	RoleId    int64
}

func (m *Manager) CreateGlobalPermission(principalName string, isGroup bool, propagate bool, roleId int64) error {
	if principalName == "" {
		return fmt.Errorf("error while creating global permission, principal name can't be empty")
	}
	method := "AuthorizationService.AddGlobalAccessControlList"
	sessionCallRes := m.client.Resource(internal.SessionPath)
	sessionCallRes.WithParam(internal.Moid, globalPermissionsServiceMoid)
	sessionCallRes.WithParam(internal.Method, method)
	req := sessionCallRes.Request(http.MethodGet)
	out := ""
	xsrfToken, err := m.client.Do(context.Background(), req, &out)
	if err != nil {
		return err
	}
	if xsrfToken == "" {
		return fmt.Errorf("error while creating session for %s %s", method, err)
	}

	type principal struct {
		XMLName xml.Name `xml:"principal"`
		Name    string   `xml:"name"`
		Group   bool     `xml:"group"`
	}
	type permissions struct {
		XMLName   xml.Name  `xml:"permissions"`
		Principal principal `xml:"principal"`
		Propagate bool      `xml:"propagate"`
		Role      int64     `xml:"roles"`
		Version   int64     `xml:"version"`
	}

	createPermissionBody := permissions{Principal: principal{Name: principalName, Group: isGroup}, Propagate: propagate, Role: roleId}
	createPermissionBodyXml, _ := xml.MarshalIndent(createPermissionBody, "  ", "    ")

	form := url.Values{}
	form.Add(internal.XsrfToken, xsrfToken)
	form.Add("permissions", string(createPermissionBodyXml))
	postReq := sessionCallRes.Request(http.MethodPost, form.Encode())

	_, err = m.client.Do(context.Background(), postReq, &out)
	if err != nil {
		return err
	}
	if strings.Contains(out, voidMethodResultMessage) {
		return nil
	}
	errorIndex := strings.Index(out, "Method Invocation Result:")
	if errorIndex > 0 {
		out = out[errorIndex:]
	}
	return fmt.Errorf("error while creating global permission for user %s %s", principalName, out)
}

func (m *Manager) ListGlobalPermission() (map[string]GlobalPermission, error) {
	method := "AuthorizationService.GetPermissions"
	sessionCallRes := m.client.Resource(internal.SessionPath)
	sessionCallRes.WithParam(internal.Moid, globalPermissionsServiceMoid)
	sessionCallRes.WithParam(internal.Method, method)
	req := sessionCallRes.Request(http.MethodGet)
	out := ""
	xsrfToken, err := m.client.Do(context.Background(), req, &out)
	if err != nil {
		return nil, err
	}
	if xsrfToken == "" {
		return nil, fmt.Errorf("error while creating session for %s %s", method, err)
	}

	form := url.Values{}
	form.Add(internal.XsrfToken, xsrfToken)
	form.Add("docUri", "https://vmware.com")
	postReq := sessionCallRes.Request(http.MethodPost, form.Encode())

	_, err = m.client.Do(context.Background(), postReq, &out)
	if err != nil {
		return nil, err
	}
	if out == "" {
		return nil, fmt.Errorf("error while getting all global permissions, no response from server")
	}

	result := map[string]GlobalPermission{}
	permissionsObjs := strings.Split(out, "version")
	for i, permissionsObj := range permissionsObjs {
		if i == 0 {
			//trim for 1st value
			permissionsObj = permissionsObj[len(permissionsObj)-450:]
		}
		principalRegex, _ := regexp.Compile("<tr><td class=\"c2\">name</td><td class=\"c1\">string</td><td>(.+?)</td></tr>")
		matches := principalRegex.FindStringSubmatch(permissionsObj)
		if len(matches) < 2 {
			continue
		}
		principal := matches[1]

		groupRegex, _ := regexp.Compile("<tr><td class=\"c2\">group</td><td class=\"c1\">boolean</td><td>(.+?)</td></tr>")
		matches = groupRegex.FindStringSubmatch(permissionsObj)
		if len(matches) < 2 {
			continue
		}
		isGroup, err := strconv.ParseBool(matches[1])
		if err != nil {
			continue
		}

		propagateRegex, _ := regexp.Compile("<tr><td class=\"c2\">propagate</td><td class=\"c1\">boolean</td><td>(.+?)</td></tr>")
		matches = propagateRegex.FindStringSubmatch(permissionsObj)
		if len(matches) < 2 {
			continue
		}
		propagate, err := strconv.ParseBool(matches[1])
		if err != nil {
			continue
		}

		roleRegex, _ := regexp.Compile("<td class=\"c2\">roles</td><td class=\"c1\">ArrayOfLong</td><td><ul class=\"noindent\"><li>(.+?)</li></ul></td></tr><tr>")
		matches = roleRegex.FindStringSubmatch(permissionsObj)
		if len(matches) < 2 {
			continue
		}
		role, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			continue
		}

		result[principal] = GlobalPermission{principal, isGroup, propagate, role}
	}
	return result, nil
}

func (m *Manager) DeleteGlobalPermission(ctx context.Context, principalName string, isGroup bool) error {
	if principalName == "" {
		return fmt.Errorf("error while deleting global permission, principal name can't be empty")
	}
	method := "AuthorizationService.RemoveGlobalAccess"
	sessionCallRes := m.client.Resource(internal.SessionPath)
	sessionCallRes.WithParam(internal.Moid, globalPermissionsServiceMoid)
	sessionCallRes.WithParam(internal.Method, method)
	req := sessionCallRes.Request(http.MethodGet)
	out := ""
	xsrfToken, err := m.client.Do(context.Background(), req, &out)
	if err != nil {
		return err
	}
	if xsrfToken == "" {
		return fmt.Errorf("error while creating session for %s %s", method, err)
	}

	type principal struct {
		XMLName xml.Name `xml:"principals"`
		Name    string   `xml:"name"`
		Group   bool     `xml:"group"`
	}

	deletePermissionBody := principal{Name: principalName, Group: isGroup}
	deletePermissionBodyXml, _ := xml.MarshalIndent(deletePermissionBody, "  ", "    ")

	form := url.Values{}
	form.Add(internal.XsrfToken, xsrfToken)
	form.Add("principals", string(deletePermissionBodyXml))
	postReq := sessionCallRes.Request(http.MethodPost, form.Encode())

	_, err = m.client.Do(context.Background(), postReq, &out)
	if err != nil {
		return err
	}
	if strings.Contains(out, voidMethodResultMessage) {
		return nil
	}
	errorIndex := strings.Index(out, "Method Invocation Result:")
	if errorIndex > 0 {
		out = out[errorIndex:]
	}
	return fmt.Errorf("error while deleting global permission for user %s %s", principalName, out)
}
