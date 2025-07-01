// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ssoadmin

import (
	"context"
	"fmt"
	"math"
	"path"
	"reflect"
	"strings"

	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/lookup"
	ltypes "github.com/vmware/govmomi/lookup/types"
	"github.com/vmware/govmomi/ssoadmin/methods"
	"github.com/vmware/govmomi/ssoadmin/types"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

const (
	Namespace  = "sso"
	Version    = "version2"
	basePath   = "/sso-adminserver"
	Path       = basePath + vim25.Path
	SystemPath = basePath + "/system-sdk"
)

var (
	ServiceInstance = vim.ManagedObjectReference{
		Type:  "SsoAdminServiceInstance",
		Value: "SsoAdminServiceInstance",
	}
)

type Client struct {
	*soap.Client

	RoundTripper   soap.RoundTripper
	ServiceContent types.AdminServiceContent
	GroupCheck     types.GroupcheckServiceContent
	Domain         string
	Limit          int32
}

func init() {
	// Fault types are not in the ssoadmin.wsdl
	vim.Add("SsoFaultNotAuthenticated", reflect.TypeOf((*vim.NotAuthenticated)(nil)).Elem())
	vim.Add("SsoFaultNoPermission", reflect.TypeOf((*vim.NoPermission)(nil)).Elem())
	vim.Add("SsoFaultInvalidCredentials", reflect.TypeOf((*vim.InvalidLogin)(nil)).Elem())
	vim.Add("SsoAdminFaultDuplicateSolutionCertificateFaultFault", reflect.TypeOf((*vim.InvalidArgument)(nil)).Elem())
}

func getEndpointURL(ctx context.Context, c *vim25.Client) string {
	// Services running on vCenter can bypass lookup service using the
	// system-sdk path. This avoids the need to lookup the system domain.
	if useSidecar := internal.UsingEnvoySidecar(c); useSidecar {
		return fmt.Sprintf("http://%s%s", c.URL().Host, SystemPath)
	}
	return getEndpointURLFromLookupService(ctx, c)
}

func getEndpointURLFromLookupService(ctx context.Context, c *vim25.Client) string {
	filter := &ltypes.LookupServiceRegistrationFilter{
		ServiceType: &ltypes.LookupServiceRegistrationServiceType{
			Product: "com.vmware.cis",
			Type:    "cs.identity",
		},
		EndpointType: &ltypes.LookupServiceRegistrationEndpointType{
			Protocol: "vmomi",
			Type:     "com.vmware.cis.cs.identity.admin",
		},
	}

	return lookup.EndpointURL(ctx, c, Path, filter)
}

func NewClient(ctx context.Context, c *vim25.Client) (*Client, error) {
	url := getEndpointURL(ctx, c)
	sc := c.NewServiceClient(url, Namespace)
	sc.Version = Version

	admin := &Client{
		Client:       sc,
		RoundTripper: sc,
		Domain:       "vsphere.local", // Default
		Limit:        math.MaxInt32,
	}
	if url != Path && !internal.UsingEnvoySidecar(c) {
		admin.Domain = path.Base(url)
	}

	{
		req := types.SsoAdminServiceInstance{
			This: ServiceInstance,
		}

		res, err := methods.SsoAdminServiceInstance(ctx, admin, &req)
		if err != nil {
			return nil, err
		}

		admin.ServiceContent = res.Returnval
	}

	{
		req := types.SsoGroupcheckServiceInstance{
			This: vim.ManagedObjectReference{
				Type: "SsoGroupcheckServiceInstance", Value: "ServiceInstance",
			},
		}

		res, err := methods.SsoGroupcheckServiceInstance(ctx, admin, &req)
		if err != nil {
			return nil, err
		}

		admin.GroupCheck = res.Returnval
	}

	return admin, nil
}

// RoundTrip dispatches to the RoundTripper field.
func (c *Client) RoundTrip(ctx context.Context, req, res soap.HasFault) error {
	// Drop any operationID header, not used by ssoadmin
	ctx = context.WithValue(ctx, vim.ID{}, "")
	return c.RoundTripper.RoundTrip(ctx, req, res)
}

func (c *Client) parseID(name string) types.PrincipalId {
	p := strings.SplitN(name, "@", 2)
	id := types.PrincipalId{Name: p[0]}
	if len(p) == 2 {
		id.Domain = p[1]
	} else {
		id.Domain = c.Domain
	}
	return id
}

func (c *Client) CreateSolutionUser(ctx context.Context, name string, details types.AdminSolutionDetails) error {
	req := types.CreateLocalSolutionUser{
		This:        c.ServiceContent.PrincipalManagementService,
		UserName:    name,
		UserDetails: details,
	}

	_, err := methods.CreateLocalSolutionUser(ctx, c, &req)
	return err
}

func (c *Client) UpdateLocalPasswordPolicy(ctx context.Context, policy types.AdminPasswordPolicy) error {
	req := types.UpdateLocalPasswordPolicy{
		This:   c.ServiceContent.PasswordPolicyService,
		Policy: policy,
	}

	_, err := methods.UpdateLocalPasswordPolicy(ctx, c, &req)
	return err
}

func (c *Client) UpdateSolutionUser(ctx context.Context, name string, details types.AdminSolutionDetails) error {
	req := types.UpdateLocalSolutionUserDetails{
		This:        c.ServiceContent.PrincipalManagementService,
		UserName:    name,
		UserDetails: details,
	}

	_, err := methods.UpdateLocalSolutionUserDetails(ctx, c, &req)
	return err
}

func (c *Client) DeletePrincipal(ctx context.Context, name string) error {
	req := types.DeleteLocalPrincipal{
		This:          c.ServiceContent.PrincipalManagementService,
		PrincipalName: name,
	}

	_, err := methods.DeleteLocalPrincipal(ctx, c, &req)
	return err
}

func (c *Client) AddUsersToGroup(ctx context.Context, groupName string, userIDs ...types.PrincipalId) error {
	req := types.AddUsersToLocalGroup{
		This:      c.ServiceContent.PrincipalManagementService,
		GroupName: groupName,
		UserIds:   userIDs,
	}

	_, err := methods.AddUsersToLocalGroup(ctx, c, &req)
	return err
}

func (c *Client) RemoveUsersFromGroup(ctx context.Context, groupName string, userIDs ...types.PrincipalId) error {
	req := types.RemovePrincipalsFromLocalGroup{
		This:          c.ServiceContent.PrincipalManagementService,
		GroupName:     groupName,
		PrincipalsIds: userIDs,
	}

	_, err := methods.RemovePrincipalsFromLocalGroup(ctx, c, &req)
	return err
}

func (c *Client) AddGroupsToGroup(ctx context.Context, groupName string, groupIDs ...types.PrincipalId) error {
	req := types.AddGroupsToLocalGroup{
		This:      c.ServiceContent.PrincipalManagementService,
		GroupName: groupName,
		GroupIds:  groupIDs,
	}

	_, err := methods.AddGroupsToLocalGroup(ctx, c, &req)
	return err
}

func (c *Client) CreateGroup(ctx context.Context, name string, details types.AdminGroupDetails) error {
	req := types.CreateLocalGroup{
		This:         c.ServiceContent.PrincipalManagementService,
		GroupName:    name,
		GroupDetails: details,
	}

	_, err := methods.CreateLocalGroup(ctx, c, &req)
	return err
}

func (c *Client) UpdateGroup(ctx context.Context, name string, details types.AdminGroupDetails) error {
	req := types.UpdateLocalGroupDetails{
		This:         c.ServiceContent.PrincipalManagementService,
		GroupName:    name,
		GroupDetails: details,
	}

	_, err := methods.UpdateLocalGroupDetails(ctx, c, &req)
	return err
}

func (c *Client) CreatePersonUser(ctx context.Context, name string, details types.AdminPersonDetails, password string) error {
	req := types.CreateLocalPersonUser{
		This:        c.ServiceContent.PrincipalManagementService,
		UserName:    name,
		UserDetails: details,
		Password:    password,
	}

	_, err := methods.CreateLocalPersonUser(ctx, c, &req)
	return err
}

func (c *Client) UpdatePersonUser(ctx context.Context, name string, details types.AdminPersonDetails) error {
	req := types.UpdateLocalPersonUserDetails{
		This:        c.ServiceContent.PrincipalManagementService,
		UserName:    name,
		UserDetails: details,
	}

	_, err := methods.UpdateLocalPersonUserDetails(ctx, c, &req)
	return err
}

func (c *Client) ResetPersonPassword(ctx context.Context, name string, password string) error {
	req := types.ResetLocalPersonUserPassword{
		This:        c.ServiceContent.PrincipalManagementService,
		UserName:    name,
		NewPassword: password,
	}

	_, err := methods.ResetLocalPersonUserPassword(ctx, c, &req)
	return err
}

func (c *Client) FindSolutionUser(ctx context.Context, name string) (*types.AdminSolutionUser, error) {
	req := types.FindSolutionUser{
		This:     c.ServiceContent.PrincipalDiscoveryService,
		UserName: name,
	}

	res, err := methods.FindSolutionUser(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) FindPersonUser(ctx context.Context, name string) (*types.AdminPersonUser, error) {
	req := types.FindPersonUser{
		This:   c.ServiceContent.PrincipalDiscoveryService,
		UserId: c.parseID(name),
	}

	res, err := methods.FindPersonUser(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) FindUser(ctx context.Context, name string) (*types.AdminUser, error) {
	req := types.FindUser{
		This:   c.ServiceContent.PrincipalDiscoveryService,
		UserId: c.parseID(name),
	}

	res, err := methods.FindUser(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) FindSolutionUsers(ctx context.Context, search string) ([]types.AdminSolutionUser, error) {
	req := types.FindSolutionUsers{
		This:         c.ServiceContent.PrincipalDiscoveryService,
		SearchString: search,
		Limit:        c.Limit,
	}

	res, err := methods.FindSolutionUsers(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) FindPersonUsers(ctx context.Context, search string) ([]types.AdminPersonUser, error) {
	req := types.FindPersonUsers{
		This: c.ServiceContent.PrincipalDiscoveryService,
		Criteria: types.AdminPrincipalDiscoveryServiceSearchCriteria{
			Domain:       c.Domain,
			SearchString: search,
		},
		Limit: c.Limit,
	}

	res, err := methods.FindPersonUsers(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) FindGroup(ctx context.Context, name string) (*types.AdminGroup, error) {
	req := types.FindGroup{
		This:    c.ServiceContent.PrincipalDiscoveryService,
		GroupId: c.parseID(name),
	}

	res, err := methods.FindGroup(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) FindGroups(ctx context.Context, search string) ([]types.AdminGroup, error) {
	req := types.FindGroups{
		This: c.ServiceContent.PrincipalDiscoveryService,
		Criteria: types.AdminPrincipalDiscoveryServiceSearchCriteria{
			Domain:       c.Domain,
			SearchString: search,
		},
		Limit: c.Limit,
	}

	res, err := methods.FindGroups(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) FindUsersInGroup(ctx context.Context, name string, search string) ([]types.AdminUser, error) {
	req := types.FindUsersInGroup{
		This:         c.ServiceContent.PrincipalDiscoveryService,
		GroupId:      c.parseID(name),
		SearchString: search,
		Limit:        c.Limit,
	}

	res, err := methods.FindUsersInGroup(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) FindGroupsInGroup(ctx context.Context, name string, search string) ([]types.AdminGroup, error) {
	req := types.FindGroupsInGroup{
		This:         c.ServiceContent.PrincipalDiscoveryService,
		GroupId:      c.parseID(name),
		SearchString: search,
		Limit:        c.Limit,
	}

	res, err := methods.FindGroupsInGroup(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) FindParentGroups(ctx context.Context, id types.PrincipalId, groups ...types.PrincipalId) ([]types.PrincipalId, error) {
	if len(groups) == 0 {
		req := types.FindAllParentGroups{
			This:   c.GroupCheck.GroupCheckService,
			UserId: id,
		}
		res, err := methods.FindAllParentGroups(ctx, c, &req)
		if err != nil {
			return nil, err
		}
		return res.Returnval, nil
	}

	return nil, nil
}

func (c *Client) GetLocalPasswordPolicy(ctx context.Context) (*types.AdminPasswordPolicy, error) {
	req := types.GetLocalPasswordPolicy{
		This: c.ServiceContent.PasswordPolicyService,
	}

	res, err := methods.GetLocalPasswordPolicy(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (c *Client) Login(ctx context.Context) error {
	req := types.Login{
		This: c.ServiceContent.SessionManager,
	}

	_, err := methods.Login(ctx, c, &req)
	return err
}

func (c *Client) Logout(ctx context.Context) error {
	req := types.Logout{
		This: c.ServiceContent.SessionManager,
	}

	_, err := methods.Logout(ctx, c, &req)
	return err
}

func (c *Client) SetRole(ctx context.Context, id types.PrincipalId, role string) (bool, error) {
	req := types.SetRole{
		This:   c.ServiceContent.RoleManagementService,
		UserId: id,
		Role:   role,
	}

	res, err := methods.SetRole(ctx, c, &req)
	if err != nil {
		return false, err
	}

	return res.Returnval, nil
}

func (c *Client) GrantWSTrustRole(ctx context.Context, id types.PrincipalId, role string) (bool, error) {
	req := types.GrantWSTrustRole{
		This:   c.ServiceContent.RoleManagementService,
		UserId: id,
		Role:   role,
	}

	res, err := methods.GrantWSTrustRole(ctx, c, &req)
	if err != nil {
		return false, err
	}

	return res.Returnval, nil
}

func (c *Client) RevokeWSTrustRole(ctx context.Context, id types.PrincipalId, role string) (bool, error) {
	req := types.RevokeWSTrustRole{
		This:   c.ServiceContent.RoleManagementService,
		UserId: id,
		Role:   role,
	}

	res, err := methods.RevokeWSTrustRole(ctx, c, &req)
	if err != nil {
		return false, err
	}

	return res.Returnval, nil
}

func (c *Client) IdentitySources(ctx context.Context) (*types.IdentitySources, error) {
	req := types.Get{
		This: c.ServiceContent.IdentitySourceManagementService,
	}

	res, err := methods.Get(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (c *Client) GetDefaultDomains(ctx context.Context) ([]string, error) {
	req := types.GetDefaultDomains{
		This: c.ServiceContent.IdentitySourceManagementService,
	}

	res, err := methods.GetDefaultDomains(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) SetDefaultDomains(ctx context.Context, domain string) error {
	req := types.SetDefaultDomains{
		This:        c.ServiceContent.IdentitySourceManagementService,
		DomainNames: domain,
	}

	_, err := methods.SetDefaultDomains(ctx, c, &req)
	return err
}

func (c *Client) RegisterLdap(ctx context.Context, stype string, name string, alias string, details types.LdapIdentitySourceDetails, auth types.SsoAdminIdentitySourceManagementServiceAuthenticationCredentails) error {
	req := types.RegisterLdap{
		This:               c.ServiceContent.IdentitySourceManagementService,
		ServerType:         stype,
		DomainName:         name,
		DomainAlias:        alias,
		Details:            details,
		AuthenticationType: "password",
		AuthnCredentials:   &auth,
	}

	_, err := methods.RegisterLdap(ctx, c, &req)
	return err
}

func (c *Client) UpdateLdap(ctx context.Context, name string, details types.LdapIdentitySourceDetails) error {
	req := types.UpdateLdap{
		This:       c.ServiceContent.IdentitySourceManagementService,
		DomainName: name,
		Details:    details,
	}

	_, err := methods.UpdateLdap(ctx, c, &req)
	return err
}

func (c *Client) UpdateLdapAuthnType(ctx context.Context, name string, auth types.SsoAdminIdentitySourceManagementServiceAuthenticationCredentails) error {
	req := types.UpdateLdapAuthnType{
		This:               c.ServiceContent.IdentitySourceManagementService,
		DomainName:         name,
		AuthenticationType: "password",
		AuthnCredentials:   &auth,
	}

	_, err := methods.UpdateLdapAuthnType(ctx, c, &req)
	return err
}

func (c *Client) GetTrustedCertificates(ctx context.Context) ([]string, error) {
	req := types.GetTrustedCertificates{
		This: c.ServiceContent.ConfigurationManagementService,
	}

	res, err := methods.GetTrustedCertificates(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}
