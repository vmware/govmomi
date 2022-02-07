/*
   Copyright (c) 2022 VMware, Inc. All Rights Reserved.

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

package simulator

import (
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/ssoadmin"
	"github.com/vmware/govmomi/ssoadmin/methods"
	"github.com/vmware/govmomi/ssoadmin/types"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		if r.IsVPX() {
			s.RegisterSDK(New())
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

type SessionManager struct {
	vim.ManagedObjectReference
}

type IdentitySourceManagementService struct {
	vim.ManagedObjectReference
}

func New() *simulator.Registry {
	r := simulator.NewRegistry()
	r.Namespace = ssoadmin.Namespace
	r.Path = ssoadmin.Path

	r.Put(&ServiceInstance{
		ManagedObjectReference: ssoadmin.ServiceInstance,
	})

	r.Put(&GroupcheckServiceInstance{
		ManagedObjectReference: groupcheckServiceInstance,
	})

	r.Put(&SessionManager{
		ManagedObjectReference: content.SessionManager,
	})

	r.Put(&IdentitySourceManagementService{
		ManagedObjectReference: content.IdentitySourceManagementService,
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
			Returnval: types.GroupcheckServiceContent{},
		},
	}
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
