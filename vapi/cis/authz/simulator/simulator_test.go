// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/cis/authz"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"

	_ "github.com/vmware/govmomi/vapi/cis/authz/simulator"
	_ "github.com/vmware/govmomi/vapi/simulator"
)

type response struct {
	code int
	body []byte
}

// wire sends requests shaped like those of a plain JSON-over-HTTP vAPI client:
// the session id in the vmware-api-session-id header and the id in the URL path.
type wire struct {
	rc      *rest.Client
	session string
}

func (c *wire) do(ctx context.Context, t *testing.T, method, path string, body any) response {
	t.Helper()

	u := *c.rc.URL()
	u.Path, u.RawQuery, _ = strings.Cut(path, "?")
	var rdr io.Reader = http.NoBody
	if body != nil {
		b, err := json.Marshal(body)
		require.NoError(t, err)
		rdr = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), rdr)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if c.session != "" {
		req.Header.Set("vmware-api-session-id", c.session)
	}

	var res response
	err = c.rc.Client.Do(ctx, req, func(r *http.Response) error {
		res.code = r.StatusCode
		res.body, err = io.ReadAll(r.Body)
		return err
	})
	require.NoError(t, err)
	return res
}

func createSpec(user string, roleIDs ...string) map[string]authz.CreateSpec {
	return map[string]authz.CreateSpec{"create_spec": {
		ResourceID: authz.GlobalPermission,
		Principal:  authz.Principal{Type: authz.PrincipalUser, UserName: user},
		RoleID:     roleIDs,
		Propagate:  true,
	}}
}

func TestPermissionWire(t *testing.T) {
	const (
		permissionPath = "/rest/com/vmware/cis/authz/permission"
		user           = `vsphere.local\wcp-cluster`
		globalID       = "username=" + user + ";doc=urn:acl:global:permissions"
	)

	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		rc := rest.NewClient(vc)
		require.NoError(t, rc.Login(ctx, simulator.DefaultLogin))
		c := &wire{rc: rc, session: rc.SessionID()}

		// A role created through AuthorizationManager is accepted by id.
		am := object.NewAuthorizationManager(vc)
		_, err := am.AddRole(ctx, "authz-test", []string{"System.View"})
		require.NoError(t, err)
		roles, err := am.RoleList(ctx)
		require.NoError(t, err)
		role := roles.ByName("authz-test")
		require.NotNil(t, role)
		roleID := strconv.Itoa(int(role.RoleId))

		// Every path requires a session.
		anon := &wire{rc: rc}
		assert.Equal(t, http.StatusUnauthorized, anon.do(ctx, t, http.MethodPost, permissionPath, createSpec(user, roleID)).code)
		assert.Equal(t, http.StatusUnauthorized, anon.do(ctx, t, http.MethodGet, permissionPath, nil).code)
		assert.Equal(t, http.StatusUnauthorized, anon.do(ctx, t, http.MethodDelete, permissionPath+"/id:"+globalID, nil).code)

		// Create returns the permission id, "value" wrapped.
		res := c.do(ctx, t, http.MethodPost, permissionPath, createSpec(user, roleID))
		require.Equal(t, http.StatusOK, res.code, string(res.body))
		var created struct {
			Value string `json:"value"`
		}
		require.NoError(t, json.Unmarshal(res.body, &created))
		assert.Equal(t, globalID, created.Value)

		// A second Global Permission for the principal fails with the already-exists message id.
		res = c.do(ctx, t, http.MethodPost, permissionPath, createSpec(user, "-1"))
		assert.Equal(t, http.StatusBadRequest, res.code)
		assert.Contains(t, string(res.body), authz.AlreadyExistsGlobal)
		var vapiErr struct {
			Type  string `json:"type"`
			Value struct {
				Messages []rest.LocalizableMessage `json:"messages"`
			} `json:"value"`
		}
		require.NoError(t, json.Unmarshal(res.body, &vapiErr))
		assert.Equal(t, "com.vmware.vapi.std.errors.already_exists", vapiErr.Type)
		require.Len(t, vapiErr.Value.Messages, 1)
		assert.Equal(t, authz.AlreadyExistsGlobal, vapiErr.Value.Messages[0].ID)
		assert.Equal(t, []string{user}, vapiErr.Value.Messages[0].Args)

		// Read path: get by id and list-detail. The backslash in the id is sent path-escaped (%5C), as url.URL encodes it.
		res = c.do(ctx, t, http.MethodGet, permissionPath+"/id:"+globalID, nil)
		require.Equal(t, http.StatusOK, res.code, string(res.body))
		var got struct {
			Value authz.Info `json:"value"`
		}
		require.NoError(t, json.Unmarshal(res.body, &got))
		assert.Equal(t, authz.Info{
			ID:         globalID,
			ResourceID: authz.GlobalPermission,
			Principal:  authz.Principal{Type: authz.PrincipalUser, UserName: user},
			RoleID:     []string{roleID},
			Propagate:  true,
		}, got.Value)

		res = c.do(ctx, t, http.MethodPost, permissionPath+"?~action=list-detail", nil)
		require.Equal(t, http.StatusOK, res.code, string(res.body))
		var list struct {
			Value []authz.Info `json:"value"`
		}
		require.NoError(t, json.Unmarshal(res.body, &list))
		assert.Equal(t, []authz.Info{got.Value}, list.Value)

		// Delete by id, then 404 once absent.
		res = c.do(ctx, t, http.MethodDelete, permissionPath+"/id:"+globalID, nil)
		assert.Equal(t, http.StatusOK, res.code, string(res.body))
		res = c.do(ctx, t, http.MethodDelete, permissionPath+"/id:"+globalID, nil)
		assert.Equal(t, http.StatusNotFound, res.code)
		assert.Contains(t, string(res.body), "com.vmware.vapi.std.errors.not_found")
		res = c.do(ctx, t, http.MethodGet, permissionPath+"/id:"+globalID, nil)
		assert.Equal(t, http.StatusNotFound, res.code)

		// The principal can be granted again once its permission is removed.
		res = c.do(ctx, t, http.MethodPost, permissionPath, createSpec(user, roleID))
		assert.Equal(t, http.StatusOK, res.code, string(res.body))
	})
}

func TestPermissionInvalid(t *testing.T) {
	const permissionPath = "/rest/com/vmware/cis/authz/permission"

	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		rc := rest.NewClient(vc)
		require.NoError(t, rc.Login(ctx, simulator.DefaultLogin))
		c := &wire{rc: rc, session: rc.SessionID()}

		tests := []struct {
			name string
			spec map[string]authz.CreateSpec
			kind string
		}{
			{"unknown role", createSpec(`vsphere.local\u`, "4242"), "invalid_argument"},
			{"non-numeric role", createSpec(`vsphere.local\u`, "Admin"), "invalid_argument"},
			{"no role", createSpec(`vsphere.local\u`), "invalid_argument"},
			{"two roles", createSpec(`vsphere.local\u`, "-1", "-2"), "invalid_argument"},
			{"no principal", createSpec("", "-1"), "invalid_argument"},
		}

		entity := createSpec(`vsphere.local\u`, "-1")
		spec := entity["create_spec"]
		spec.ResourceID = authz.DynamicID{Type: "Folder", ID: "group-d1"}
		entity["create_spec"] = spec
		tests = append(tests, struct {
			name string
			spec map[string]authz.CreateSpec
			kind string
		}{"entity resource", entity, "unsupported"})

		for _, test := range tests {
			res := c.do(ctx, t, http.MethodPost, permissionPath, test.spec)
			assert.Equal(t, http.StatusBadRequest, res.code, test.name)
			assert.Contains(t, string(res.body), "com.vmware.vapi.std.errors."+test.kind, test.name)
		}

		res := c.do(ctx, t, http.MethodGet, permissionPath+"/not-an-id", nil)
		assert.Equal(t, http.StatusNotFound, res.code)

		res = c.do(ctx, t, http.MethodGet, permissionPath, nil)
		require.Equal(t, http.StatusOK, res.code)
		assert.JSONEq(t, `{"value":[]}`, string(res.body))
	})
}
