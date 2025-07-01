// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"encoding/base64"
	"encoding/pem"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/ssoadmin"
	"github.com/vmware/govmomi/ssoadmin/methods"
	"github.com/vmware/govmomi/ssoadmin/types"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		if r.IsVPX() {
			s.RegisterSDK(New(r, s.Listen), ssoadmin.SystemPath)
		}
	})
}

var content = types.AdminServiceContent{
	SessionManager:                  vim.ManagedObjectReference{Type: "SsoSessionManager", Value: "ssoSessionManager"},
	ConfigurationManagementService:  vim.ManagedObjectReference{Type: "SsoAdminConfigurationManagementService", Value: "configurationManagementService"},
	SmtpManagementService:           vim.ManagedObjectReference{Type: "SsoAdminSmtpManagementService", Value: "smtpManagementService"},
	PrincipalDiscoveryService:       vim.ManagedObjectReference{Type: "SsoAdminPrincipalDiscoveryService", Value: "principalDiscoveryService"},
	PrincipalManagementService:      vim.ManagedObjectReference{Type: "SsoAdminPrincipalManagementService", Value: "principalManagementService"},
	RoleManagementService:           vim.ManagedObjectReference{Type: "SsoAdminRoleManagementService", Value: "roleManagementService"},
	PasswordPolicyService:           vim.ManagedObjectReference{Type: "SsoAdminPasswordPolicyService", Value: "passwordPolicyService"},
	LockoutPolicyService:            vim.ManagedObjectReference{Type: "SsoAdminLockoutPolicyService", Value: "lockoutPolicyService"},
	DomainManagementService:         vim.ManagedObjectReference{Type: "SsoAdminDomainManagementService", Value: "domainManagementService"},
	IdentitySourceManagementService: vim.ManagedObjectReference{Type: "SsoAdminIdentitySourceManagementService", Value: "identitySourceManagementService"},
	SystemManagementService:         vim.ManagedObjectReference{Type: "SsoAdminSystemManagementService", Value: "systemManagementService"},
	DeploymentInformationService:    vim.ManagedObjectReference{Type: "SsoAdminDeploymentInformationService", Value: "deploymentInformationService"},
	ReplicationService:              vim.ManagedObjectReference{Type: "SsoAdminReplicationService", Value: "replicationService"},
}

var groupcheckContent = types.GroupcheckServiceContent{
	SessionManager:    vim.ManagedObjectReference{Type: "SsoSessionManager", Value: "ssoSessionManager"},
	GroupCheckService: vim.ManagedObjectReference{Type: "SsoGroupcheckGroupCheckService", Value: "groupCheckService"},
}

var groupcheckServiceInstance = vim.ManagedObjectReference{
	Type: "SsoGroupcheckServiceInstance", Value: "ServiceInstance",
}

var IdentitySources = types.IdentitySources{
	System: types.IdentitySource{
		Name:    "vsphere.local",
		Domains: []types.Domain{{Name: "vsphere.local", Alias: "vmwarem"}},
	},
	LocalOS: &types.IdentitySource{
		Name:    "localos",
		Domains: []types.Domain{{Name: "localos", Alias: ""}},
	},
	NativeAD: nil,
	LDAPS: []types.LdapIdentitySource{
		{
			IdentitySource: types.IdentitySource{
				Name:    "example.com",
				Domains: []types.Domain{{Name: "example.com", Alias: "examplex"}},
			},
			Type: "ActiveDirectory",
			Details: types.LdapIdentitySourceDetails{
				FriendlyName: "foo",
				UserBaseDn:   "ou=People,dc=example,dc=org",
				GroupBaseDn:  "ou=People,dc=example,dc=org",
				PrimaryURL:   "ldap://10.168.194.120:389",
				FailoverURL:  "",
			},
			AuthenticationDetails: types.AuthenticationDetails{
				AuthenticationType: "PASSWORD",
				Username:           "cn=admin,dc=example,dc=org",
			},
		},
	},
}

type ServiceInstance struct {
	vim.ManagedObjectReference
}

type GroupcheckServiceInstance struct {
	vim.ManagedObjectReference
}

type GroupcheckService struct {
	vim.ManagedObjectReference

	m *PrincipalManagementService
}

type SessionManager struct {
	vim.ManagedObjectReference
}

type ConfigurationManagementService struct {
	vim.ManagedObjectReference
}

type IdentitySourceManagementService struct {
	vim.ManagedObjectReference
}

type PrincipalManagementService struct {
	vim.ManagedObjectReference

	dir map[types.PrincipalId]principal
}

type PrincipalDiscoveryService struct {
	vim.ManagedObjectReference

	m *PrincipalManagementService
}

type principal struct {
	group    *types.AdminGroup
	members  map[types.PrincipalId]principal
	solution *types.AdminSolutionUser
	person   *types.AdminPersonUser
	password string
}

func New(vc *simulator.Registry, u *url.URL) *simulator.Registry {
	r := simulator.NewRegistry()
	r.Namespace = ssoadmin.Namespace
	r.Path = ssoadmin.Path

	settings := vc.OptionManager().Setting
	for i := range settings {
		setting := settings[i].GetOptionValue()
		if setting.Key == "config.vpxd.sso.admin.uri" {
			endpoint, _ := url.Parse(setting.Value.(string))
			r.Path = endpoint.Path
		}
	}

	r.Put(&ServiceInstance{
		ManagedObjectReference: ssoadmin.ServiceInstance,
	})

	r.Put(&GroupcheckServiceInstance{
		ManagedObjectReference: groupcheckServiceInstance,
	})

	r.Put(&SessionManager{
		ManagedObjectReference: content.SessionManager,
	})

	r.Put(&ConfigurationManagementService{
		ManagedObjectReference: content.ConfigurationManagementService,
	})

	r.Put(&IdentitySourceManagementService{
		ManagedObjectReference: content.IdentitySourceManagementService,
	})

	m := &PrincipalManagementService{
		ManagedObjectReference: content.PrincipalManagementService,
		dir:                    make(map[types.PrincipalId]principal),
	}
	r.Put(m)
	m.createDefaultUser(u.User)
	vc.SessionManager().ValidLogin = m.validLogin

	r.Put(&PrincipalDiscoveryService{
		ManagedObjectReference: content.PrincipalDiscoveryService,
		m:                      m,
	})

	r.Put(&GroupcheckService{
		ManagedObjectReference: groupcheckContent.GroupCheckService,
		m:                      m,
	})

	return r
}

func (s *ServiceInstance) SsoAdminServiceInstance(ctx *simulator.Context, _ *types.SsoAdminServiceInstance) soap.HasFault {
	return &methods.SsoAdminServiceInstanceBody{
		Res: &types.SsoAdminServiceInstanceResponse{
			Returnval: content,
		},
	}
}

func (s *GroupcheckServiceInstance) SsoGroupcheckServiceInstance(ctx *simulator.Context, _ *types.SsoGroupcheckServiceInstance) soap.HasFault {
	return &methods.SsoGroupcheckServiceInstanceBody{
		Res: &types.SsoGroupcheckServiceInstanceResponse{
			Returnval: groupcheckContent,
		},
	}
}

func (s *GroupcheckService) FindAllParentGroups(ctx *simulator.Context, req *types.FindAllParentGroups) soap.HasFault {
	body := new(methods.FindAllParentGroupsBody)

	p, ok := s.m.dir[req.UserId]
	if !ok || p.person == nil {
		body.Fault_ = simulator.Fault("", new(vim.NotFound))
		return body
	}

	body.Res = new(types.FindAllParentGroupsResponse)

	for gid, p := range s.m.dir {
		if p.group == nil {
			continue
		}

		for id, m := range p.members {
			if m.person == nil {
				continue
			}
			if id == req.UserId {
				body.Res.Returnval = append(body.Res.Returnval, gid)
			}
		}
	}

	return body
}

func (s *SessionManager) Login(ctx *simulator.Context, req *types.Login) soap.HasFault {
	body := new(methods.LoginBody)

	// TODO: check for Assertion>Subject>NameID as simulator.SessionManager.LoginByToken does
	body.Res = new(types.LoginResponse)

	return body
}

func (s *SessionManager) Logout(ctx *simulator.Context, req *types.Logout) soap.HasFault {
	return &methods.LogoutBody{
		Res: new(types.LogoutResponse),
	}
}

func (*ConfigurationManagementService) GetTrustedCertificates(ctx *simulator.Context, _ *types.GetTrustedCertificates) soap.HasFault {
	m := ctx.For(vim25.Path).Map.SessionManager()

	var res []string

	// TODO: consider adding a vcsim -tlscacerts flag
	cacerts := os.Getenv("VCSIM_CACERTS")
	if cacerts != "" {
		pemCerts, err := os.ReadFile(cacerts)
		if err != nil {
			log.Fatal(err)
		}
		for len(pemCerts) > 0 {
			var block *pem.Block
			block, pemCerts = pem.Decode(pemCerts)
			if block == nil {
				break
			}
			if block.Type != "CERTIFICATE" || len(block.Headers) != 0 {
				continue
			}
			res = append(res, base64.StdEncoding.EncodeToString(block.Bytes))
		}
	} else if m.TLS != nil {
		res = append(res, m.TLSCert())
	}

	return &methods.GetTrustedCertificatesBody{
		Res: &types.GetTrustedCertificatesResponse{
			Returnval: res,
		},
	}
}

func (s *IdentitySourceManagementService) Get(ctx *simulator.Context, _ *types.Get) soap.HasFault {
	sources := IdentitySources
	sources.All = nil
	sources.All = append(sources.All, sources.System)

	if sources.LocalOS != nil {
		sources.All = append(sources.All, *sources.LocalOS)
	}

	if sources.NativeAD != nil {
		sources.All = append(sources.All, *sources.NativeAD)
	}

	for _, ldap := range sources.LDAPS {
		sources.All = append(sources.All, ldap.IdentitySource)
	}

	return &methods.GetBody{
		Res: &types.GetResponse{
			Returnval: sources,
		},
	}
}

func (s *PrincipalDiscoveryService) FindUser(ctx *simulator.Context, req *types.FindUser) soap.HasFault {
	body := &methods.FindUserBody{
		Res: new(types.FindUserResponse),
	}

	p, ok := s.m.dir[req.UserId]
	if !ok {
		return body
	}

	var user *types.AdminUser

	switch {
	case p.person != nil:
		user = &types.AdminUser{
			Kind:        "person",
			Id:          p.person.Id,
			Description: p.person.Details.Description,
		}
	case p.solution != nil:
		user = &types.AdminUser{
			Kind:        "solution",
			Id:          p.solution.Id,
			Description: p.solution.Details.Description,
		}

	}

	body.Res.Returnval = user

	return body
}

func (s *PrincipalDiscoveryService) FindPersonUser(ctx *simulator.Context, req *types.FindPersonUser) soap.HasFault {
	body := &methods.FindPersonUserBody{
		Res: new(types.FindPersonUserResponse),
	}

	p, ok := s.m.dir[req.UserId]
	if ok {
		body.Res.Returnval = p.person
	}

	return body
}

func (s *PrincipalDiscoveryService) FindPersonUsers(ctx *simulator.Context, req *types.FindPersonUsers) soap.HasFault {
	body := &methods.FindPersonUsersBody{
		Res: &types.FindPersonUsersResponse{},
	}

	for _, p := range s.m.dir {
		if p.person == nil {
			continue
		}

		if domain := req.Criteria.Domain; domain != "" {
			if p.person.Id.Domain != domain {
				continue
			}
		}

		if search := req.Criteria.SearchString; search != "" {
			if !strings.Contains(p.person.Id.Name, search) {
				continue
			}
		}

		body.Res.Returnval = append(body.Res.Returnval, *p.person)
	}

	return body
}

func (s *PrincipalManagementService) createDefaultUser(u *url.Userinfo) {
	user := &types.AdminPersonUser{
		Id: types.PrincipalId{
			Name:   u.Username(),
			Domain: IdentitySources.System.Name,
		},
	}

	pass, _ := u.Password()
	s.dir[user.Id] = principal{
		person:   user,
		password: pass,
	}
}

func (s *PrincipalManagementService) validLogin(login *vim.Login) bool {
	id := parseID(login.UserName)

	if p, ok := s.dir[id]; ok && p.person != nil {
		return p.password == login.Password
	}

	return false
}

func (s *PrincipalDiscoveryService) FindSolutionUser(ctx *simulator.Context, req *types.FindSolutionUser) soap.HasFault {
	body := &methods.FindSolutionUserBody{
		Res: new(types.FindSolutionUserResponse),
	}

	id := parseID(req.UserName)
	p, ok := s.m.dir[id]
	if ok {
		body.Res.Returnval = p.solution
	}

	return body
}

func (s *PrincipalDiscoveryService) FindSolutionUsers(ctx *simulator.Context, req *types.FindSolutionUsers) soap.HasFault {
	body := &methods.FindSolutionUsersBody{
		Res: new(types.FindSolutionUsersResponse),
	}

	for _, p := range s.m.dir {
		if p.solution == nil {
			continue
		}

		if search := req.SearchString; search != "" {
			if !strings.Contains(p.solution.Id.Name, search) {
				continue
			}
		}

		body.Res.Returnval = append(body.Res.Returnval, *p.solution)
	}

	return body
}

func (s *PrincipalManagementService) CreateLocalPersonUser(ctx *simulator.Context, req *types.CreateLocalPersonUser) soap.HasFault {
	body := new(methods.CreateLocalPersonUserBody)

	user := &types.AdminPersonUser{
		Id: types.PrincipalId{
			Name:   req.UserName,
			Domain: IdentitySources.System.Name,
		},
		Details: req.UserDetails,
	}

	if _, ok := s.dir[user.Id]; ok {
		body.Fault_ = simulator.Fault("", new(vim.DuplicateName))
		return body
	}

	s.dir[user.Id] = principal{
		person:   user,
		password: req.Password,
	}

	body.Res = new(types.CreateLocalPersonUserResponse)

	return body
}

func (s *PrincipalManagementService) UpdateLocalPersonUserDetails(ctx *simulator.Context, req *types.UpdateLocalPersonUserDetails) soap.HasFault {
	body := &methods.UpdateLocalPersonUserDetailsBody{
		Res: new(types.UpdateLocalPersonUserDetailsResponse),
	}

	id := parseID(req.UserName)
	p, ok := s.dir[id]
	if !ok || p.person == nil {
		body.Fault_ = simulator.Fault("", new(vim.NotFound))
		return body
	}

	p.person.Details = req.UserDetails

	return body
}

func (s *PrincipalManagementService) ResetLocalPersonUserPassword(ctx *simulator.Context, req *types.ResetLocalPersonUserPassword) soap.HasFault {
	body := &methods.ResetLocalPersonUserPasswordBody{
		Res: new(types.ResetLocalPersonUserPasswordResponse),
	}

	id := parseID(req.UserName)
	p, ok := s.dir[id]
	if !ok || p.person == nil {
		body.Fault_ = simulator.Fault("", new(vim.NotFound))
		return body
	}

	p.password = req.NewPassword

	return body
}

func (s *PrincipalManagementService) CreateLocalSolutionUser(ctx *simulator.Context, req *types.CreateLocalSolutionUser) soap.HasFault {
	body := new(methods.CreateLocalSolutionUserBody)

	user := &types.AdminSolutionUser{
		Id: types.PrincipalId{
			Name:   req.UserName,
			Domain: IdentitySources.System.Name,
		},
		Details: req.UserDetails,
	}

	if _, ok := s.dir[user.Id]; ok {
		body.Fault_ = simulator.Fault("", new(vim.DuplicateName))
		return body
	}

	s.dir[user.Id] = principal{
		solution: user,
	}

	body.Res = new(types.CreateLocalSolutionUserResponse)

	return body
}

func (s *PrincipalManagementService) UpdateLocalSolutionUserDetails(ctx *simulator.Context, req *types.UpdateLocalSolutionUserDetails) soap.HasFault {
	body := &methods.UpdateLocalSolutionUserDetailsBody{
		Res: new(types.UpdateLocalSolutionUserDetailsResponse),
	}

	id := parseID(req.UserName)
	p, ok := s.dir[id]
	if !ok || p.solution == nil {
		body.Fault_ = simulator.Fault("", new(vim.NotFound))
		return body
	}

	p.solution.Details = req.UserDetails

	return body
}

func (s *PrincipalManagementService) DeleteLocalPrincipal(ctx *simulator.Context, req *types.DeleteLocalPrincipal) soap.HasFault {
	body := new(methods.DeleteLocalPrincipalBody)

	id := parseID(req.PrincipalName)

	if _, ok := s.dir[id]; ok {
		delete(s.dir, id)
		for _, p := range s.dir {
			if p.group != nil {
				delete(p.members, id)
			}
		}
		body.Res = new(types.DeleteLocalPrincipalResponse)
	} else {
		body.Fault_ = simulator.Fault("", new(vim.NotFound))
	}

	return body
}

func parseID(name string) types.PrincipalId {
	p := strings.SplitN(name, "@", 2)
	id := types.PrincipalId{Name: p[0]}
	if len(p) == 2 {
		id.Domain = p[1]
	} else {
		id.Domain = IdentitySources.System.Name
	}
	return id
}

func (s *PrincipalDiscoveryService) FindGroup(ctx *simulator.Context, req *types.FindGroup) soap.HasFault {
	body := &methods.FindGroupBody{
		Res: new(types.FindGroupResponse),
	}

	p, ok := s.m.dir[req.GroupId]
	if !ok {
		return body
	}

	body.Res.Returnval = p.group

	return body
}

func (s *PrincipalDiscoveryService) FindGroups(ctx *simulator.Context, req *types.FindGroups) soap.HasFault {
	body := &methods.FindGroupsBody{
		Res: &types.FindGroupsResponse{},
	}

	for _, p := range s.m.dir {
		if p.group == nil {
			continue
		}

		if domain := req.Criteria.Domain; domain != "" {
			if p.group.Id.Domain != domain {
				continue
			}
		}

		if search := req.Criteria.SearchString; search != "" {
			if !strings.Contains(p.group.Id.Name, search) {
				continue
			}
		}

		body.Res.Returnval = append(body.Res.Returnval, *p.group)
	}

	return body
}

func (s *PrincipalDiscoveryService) FindGroupsInGroup(ctx *simulator.Context, req *types.FindGroupsInGroup) soap.HasFault {
	body := &methods.FindGroupsInGroupBody{
		Res: &types.FindGroupsInGroupResponse{},
	}

	for _, p := range s.m.dir {
		if p.group == nil {
			continue
		}

		for id := range p.members {
			if search := req.SearchString; search != "" {
				if !strings.Contains(id.Name, search) {
					continue
				}
			}

			body.Res.Returnval = append(body.Res.Returnval, *p.group)
		}
	}

	return body
}

func (s *PrincipalManagementService) CreateLocalGroup(ctx *simulator.Context, req *types.CreateLocalGroup) soap.HasFault {
	body := new(methods.CreateLocalGroupBody)

	group := &types.AdminGroup{
		Id: types.PrincipalId{
			Name:   req.GroupName,
			Domain: IdentitySources.System.Name,
		},
		Details: req.GroupDetails,
	}

	if _, ok := s.dir[group.Id]; ok {
		body.Fault_ = simulator.Fault("", new(vim.DuplicateName))
		return body
	}

	s.dir[group.Id] = principal{
		group:   group,
		members: make(map[types.PrincipalId]principal),
	}

	body.Res = new(types.CreateLocalGroupResponse)

	return body
}

func (s *PrincipalManagementService) UpdateLocalGroupDetails(ctx *simulator.Context, req *types.UpdateLocalGroupDetails) soap.HasFault {
	body := &methods.UpdateLocalGroupDetailsBody{
		Res: new(types.UpdateLocalGroupDetailsResponse),
	}

	id := parseID(req.GroupName)
	p, ok := s.dir[id]
	if !ok || p.group == nil {
		body.Fault_ = simulator.Fault("", new(vim.NotFound))
		return body
	}

	p.group.Details = req.GroupDetails

	return body
}

func (s *PrincipalManagementService) AddUsersToLocalGroup(ctx *simulator.Context, req *types.AddUsersToLocalGroup) soap.HasFault {
	body := &methods.AddUsersToLocalGroupBody{
		Res: new(types.AddUsersToLocalGroupResponse),
	}

	id := parseID(req.GroupName)
	g, ok := s.dir[id]
	if !ok || g.group == nil {
		body.Fault_ = simulator.Fault("", new(vim.NotFound))
		return body
	}

	res := new(types.AddUsersToLocalGroupResponse)

	for _, id := range req.UserIds {
		p, ok := s.dir[id]
		if !ok || p.person == nil {
			body.Fault_ = simulator.Fault("", new(vim.NotFound))
			return body
		}

		added := false
		if _, ok := g.members[id]; !ok {
			g.members[id] = p
		} else {
			added = true
		}

		res.Returnval = append(res.Returnval, added)
	}

	body.Res = res

	return body
}

func (s *PrincipalManagementService) AddGroupsToLocalGroup(ctx *simulator.Context, req *types.AddGroupsToLocalGroup) soap.HasFault {
	body := &methods.AddGroupsToLocalGroupBody{
		Res: new(types.AddGroupsToLocalGroupResponse),
	}

	id := parseID(req.GroupName)
	g, ok := s.dir[id]
	if !ok || g.group == nil {
		body.Fault_ = simulator.Fault("", new(vim.NotFound))
		return body
	}

	res := new(types.AddGroupsToLocalGroupResponse)

	for _, id := range req.GroupIds {
		p, ok := s.dir[id]
		if !ok || p.group == nil {
			body.Fault_ = simulator.Fault("", new(vim.NotFound))
			return body
		}

		added := false
		if _, ok := g.members[id]; !ok {
			g.members[id] = p
		} else {
			added = true
		}

		res.Returnval = append(res.Returnval, added)
	}

	body.Res = res

	return body
}

func (s *PrincipalManagementService) RemovePrincipalsFromLocalGroup(ctx *simulator.Context, req *types.RemovePrincipalsFromLocalGroup) soap.HasFault {
	body := &methods.RemovePrincipalsFromLocalGroupBody{
		Res: new(types.RemovePrincipalsFromLocalGroupResponse),
	}

	id := parseID(req.GroupName)
	g, ok := s.dir[id]
	if !ok || g.group == nil {
		body.Fault_ = simulator.Fault("", new(vim.NotFound))
		return body
	}

	res := new(types.RemovePrincipalsFromLocalGroupResponse)

	for _, id := range req.PrincipalsIds {
		_, ok := s.dir[id]
		if !ok {
			body.Fault_ = simulator.Fault("", new(vim.NotFound))
			return body
		}

		removed := false
		if _, ok := g.members[id]; ok {
			delete(g.members, id)
			removed = true
		}

		res.Returnval = append(res.Returnval, removed)
	}

	body.Res = res

	return body
}
