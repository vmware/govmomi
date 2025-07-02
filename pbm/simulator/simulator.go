// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"regexp"
	"slices"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/pbm"
	"github.com/vmware/govmomi/pbm/methods"
	"github.com/vmware/govmomi/pbm/types"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

var content = types.PbmServiceInstanceContent{
	AboutInfo: types.PbmAboutInfo{
		Name:         "PBM",
		Version:      "2.0",
		InstanceUuid: "df09f335-be97-4f33-8c27-315faaaad6fc",
	},
	SessionManager:            vim.ManagedObjectReference{Type: "PbmSessionManager", Value: "SessionManager"},
	CapabilityMetadataManager: vim.ManagedObjectReference{Type: "PbmCapabilityMetadataManager", Value: "CapabilityMetadataManager"},
	ProfileManager:            vim.ManagedObjectReference{Type: "PbmProfileProfileManager", Value: "ProfileManager"},
	ComplianceManager:         vim.ManagedObjectReference{Type: "PbmComplianceManager", Value: "complianceManager"},
	PlacementSolver:           vim.ManagedObjectReference{Type: "PbmPlacementSolver", Value: "placementSolver"},
	ReplicationManager:        &vim.ManagedObjectReference{Type: "PbmReplicationManager", Value: "ReplicationManager"},
}

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		if r.IsVPX() {
			s.RegisterSDK(New())
		}
	})
}

func New() *simulator.Registry {
	r := simulator.NewRegistry()
	r.Namespace = pbm.Namespace
	r.Path = pbm.Path
	r.Cookie = simulator.SOAPCookie

	r.Put(&ServiceInstance{
		ManagedObjectReference: pbm.ServiceInstance,
		Content:                content,
	})

	profileManager := &ProfileManager{
		ManagedObjectReference: content.ProfileManager,
	}
	profileManager.init(r)
	r.Put(profileManager)

	r.Put(&PlacementSolver{
		ManagedObjectReference: content.PlacementSolver,
	})

	return r
}

type ServiceInstance struct {
	vim.ManagedObjectReference

	Content types.PbmServiceInstanceContent
}

func (s *ServiceInstance) PbmRetrieveServiceContent(_ *types.PbmRetrieveServiceContent) soap.HasFault {
	return &methods.PbmRetrieveServiceContentBody{
		Res: &types.PbmRetrieveServiceContentResponse{
			Returnval: s.Content,
		},
	}
}

type ProfileManager struct {
	vim.ManagedObjectReference

	profiles       []types.BasePbmProfile
	profileDetails map[string]types.PbmProfileDetails
}

// This function  is to deal with non-ascii characters like some non-english languages
// to convert to more similar and meaningful characters for further consideration.
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func getK8sCompliantNameForPolicy(policyName string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	policyName, _, _ = transform.String(t, policyName)
	str := strings.ToLower(policyName)
	reg := regexp.MustCompile("[^a-z0-9.]+")

	convertedString := reg.ReplaceAllString(str, "-")

	// K8sCompliantName cannot begin with '-' or '.', simply add '0' to resolve it.
	if convertedString[0] == '-' || convertedString[0] == '.' {
		convertedString = "0" + convertedString
	}

	// K8sCompliantName cannot end with '-' or '.' as well, add another '0' if so.
	if convertedString[len(convertedString)-1] == '-' || convertedString[len(convertedString)-1] == '.' {
		convertedString += "0"
	}

	return convertedString
}

func (m *ProfileManager) init(_ *simulator.Registry) {
	m.profiles = slices.Clone(vcenter67DefaultProfiles)
	// Set K8s compliant names for default profiles
	for i := range m.profiles {
		b, ok := m.profiles[i].(types.BasePbmCapabilityProfile)
		if !ok {
			continue
		}
		p := b.GetPbmCapabilityProfile()
		p.K8sCompliantName = getK8sCompliantNameForPolicy(p.Name)
	}

	// Ensure the default encryption profile has the encryption IOFILTER as this
	// is required when detecting whether a policy supports encryption.
	m.profileDetails = map[string]types.PbmProfileDetails{
		defaultEncryptionProfile.ProfileId.UniqueId: {
			Profile: defaultEncryptionProfile,
			IofInfos: []types.PbmIofilterInfo{
				{
					FilterType: string(types.PbmIofilterInfoFilterTypeENCRYPTION),
				},
			},
		},
	}
}

func (m *ProfileManager) PbmQueryProfile(req *types.PbmQueryProfile) soap.HasFault {
	body := new(methods.PbmQueryProfileBody)
	body.Res = new(types.PbmQueryProfileResponse)

	for i := range m.profiles {
		b, ok := m.profiles[i].(types.BasePbmCapabilityProfile)
		if !ok {
			continue
		}
		p := b.GetPbmCapabilityProfile()

		if p.ResourceType != req.ResourceType {
			continue
		}

		if req.ProfileCategory != "" {
			if p.ProfileCategory != req.ProfileCategory {
				continue
			}
		}

		body.Res.Returnval = append(body.Res.Returnval, types.PbmProfileId{
			UniqueId: p.ProfileId.UniqueId,
		})
	}

	return body
}

func (m *ProfileManager) PbmQueryProfileDetails(req *types.PbmQueryProfileDetails) soap.HasFault {
	body := new(methods.PbmQueryProfileDetailsBody)
	body.Res = new(types.PbmQueryProfileDetailsResponse)

	for i := range m.profiles {
		b, ok := m.profiles[i].(types.BasePbmCapabilityProfile)
		if !ok {
			continue
		}
		p := b.GetPbmCapabilityProfile()

		if req.ProfileCategory != "" {
			if p.ProfileCategory != req.ProfileCategory {
				continue
			}
		}

		body.Res.Returnval = append(body.Res.Returnval, types.PbmProfileDetails{
			Profile: p,
		})
	}

	return body
}

func (m *ProfileManager) PbmQueryAssociatedProfile(req *types.PbmQueryAssociatedProfile) soap.HasFault {
	body := new(methods.PbmQueryAssociatedProfileBody)
	body.Res = new(types.PbmQueryAssociatedProfileResponse)

	return body
}

func (m *ProfileManager) PbmQueryAssociatedProfiles(req *types.PbmQueryAssociatedProfiles) soap.HasFault {
	body := new(methods.PbmQueryAssociatedProfilesBody)
	body.Res = new(types.PbmQueryAssociatedProfilesResponse)

	return body
}

func (m *ProfileManager) PbmRetrieveContent(req *types.PbmRetrieveContent) soap.HasFault {
	body := new(methods.PbmRetrieveContentBody)
	if len(req.ProfileIds) == 0 {
		body.Fault_ = simulator.Fault("", new(vim.InvalidRequest))
		return body
	}

	var res []types.BasePbmProfile

	match := func(id string) bool {
		for _, p := range m.profiles {
			if id == p.GetPbmProfile().ProfileId.UniqueId {
				res = append(res, p)
				return true
			}
		}
		return false
	}

	for _, p := range req.ProfileIds {
		if match(p.UniqueId) {
			continue
		}

		body.Fault_ = simulator.Fault("", &vim.InvalidArgument{InvalidProperty: "profileId"})
		return body
	}

	body.Res = &types.PbmRetrieveContentResponse{Returnval: res}

	return body
}

func (m *ProfileManager) PbmCreate(ctx *simulator.Context, req *types.PbmCreate) soap.HasFault {
	body := new(methods.PbmCreateBody)
	body.Res = new(types.PbmCreateResponse)

	profile := &types.PbmCapabilityProfile{
		PbmProfile: types.PbmProfile{
			ProfileId: types.PbmProfileId{
				UniqueId: uuid.New().String(),
			},
			Name:            req.CreateSpec.Name,
			Description:     req.CreateSpec.Description,
			CreationTime:    time.Now(),
			CreatedBy:       ctx.Session.UserName,
			LastUpdatedTime: time.Now(),
			LastUpdatedBy:   ctx.Session.UserName,
		},
		ProfileCategory:          req.CreateSpec.Category,
		ResourceType:             req.CreateSpec.ResourceType,
		Constraints:              req.CreateSpec.Constraints,
		GenerationId:             0,
		IsDefault:                false,
		SystemCreatedProfileType: "",
		LineOfService:            "",
		K8sCompliantName:         req.CreateSpec.K8sCompliantName,
	}

	m.profiles = append(m.profiles, profile)
	body.Res.Returnval.UniqueId = profile.PbmProfile.ProfileId.UniqueId

	return body
}

func (m *ProfileManager) PbmDelete(req *types.PbmDelete) soap.HasFault {
	body := new(methods.PbmDeleteBody)

	for _, id := range req.ProfileId {
		for i, p := range m.profiles {
			pid := p.GetPbmProfile().ProfileId

			if id == pid {
				m.profiles = append(m.profiles[:i], m.profiles[i+1:]...)
				break
			}
		}
	}

	body.Res = new(types.PbmDeleteResponse)

	return body
}

func (m *ProfileManager) PbmQueryIOFiltersFromProfileId(req *types.PbmQueryIOFiltersFromProfileId) soap.HasFault {
	body := methods.PbmQueryIOFiltersFromProfileIdBody{
		Res: &types.PbmQueryIOFiltersFromProfileIdResponse{},
	}

	for i := range req.ProfileIds {
		profileID := req.ProfileIds[i]
		if profileDetails, ok := m.profileDetails[profileID.UniqueId]; ok {
			body.Res.Returnval = append(
				body.Res.Returnval,
				types.PbmProfileToIofilterMap{
					Key:       profileID,
					Iofilters: profileDetails.IofInfos,
				})
		} else {
			body.Fault_ = simulator.Fault("Invalid profile ID", &vim.RuntimeFault{})
			break
		}
	}

	if body.Fault_ != nil {
		body.Res = nil
	}

	return &body
}

func (m *ProfileManager) PbmResolveK8sCompliantNames(req *types.PbmResolveK8sCompliantNames) soap.HasFault {
	body := new(methods.PbmResolveK8sCompliantNamesBody)
	body.Res = new(types.PbmResolveK8sCompliantNamesResponse)

	for i := range m.profiles {
		b, ok := m.profiles[i].(types.BasePbmCapabilityProfile)
		if !ok {
			continue
		}
		p := b.GetPbmCapabilityProfile()
		p.OtherK8sCompliantNames = append(p.OtherK8sCompliantNames, p.K8sCompliantName+"-latebinding")
	}
	return body
}

func (m *ProfileManager) PbmUpdateK8sCompliantNames(req *types.PbmUpdateK8sCompliantNames) soap.HasFault {
	body := new(methods.PbmUpdateK8sCompliantNamesBody)
	body.Res = new(types.PbmUpdateK8sCompliantNamesResponse)

	for i := range m.profiles {
		b, ok := m.profiles[i].(types.BasePbmCapabilityProfile)
		if !ok {
			continue
		}
		p := b.GetPbmCapabilityProfile()
		if p.ProfileId.UniqueId == req.K8sCompliantNameSpec.ProfileId {
			if p.K8sCompliantName == "" {
				p.K8sCompliantName = req.K8sCompliantNameSpec.K8sCompliantName
			} else if p.K8sCompliantName != req.K8sCompliantNameSpec.K8sCompliantName {
				body.Fault_ = simulator.Fault("Invalid Argument", &vim.RuntimeFault{})
				break
			}
			if slices.Contains(req.K8sCompliantNameSpec.OtherK8sCompliantNames, p.K8sCompliantName) {
				body.Fault_ = simulator.Fault("Duplicate Name", &vim.RuntimeFault{})
				break
			}
			p.OtherK8sCompliantNames = req.K8sCompliantNameSpec.OtherK8sCompliantNames
			break
		}
	}
	return body
}

type PlacementSolver struct {
	vim.ManagedObjectReference
}

func (m *PlacementSolver) PbmCheckRequirements(ctx *simulator.Context, req *types.PbmCheckRequirements) soap.HasFault {
	body := new(methods.PbmCheckRequirementsBody)
	body.Res = new(types.PbmCheckRequirementsResponse)

	for _, ds := range ctx.For(vim25.Path).Map.All("Datastore") {
		// TODO: filter
		ref := ds.Reference()
		body.Res.Returnval = append(body.Res.Returnval, types.PbmPlacementCompatibilityResult{
			Hub: types.PbmPlacementHub{
				HubType: ref.Type,
				HubId:   ref.Value,
			},
			MatchingResources: nil,
			HowMany:           0,
			Utilization:       nil,
			Warning:           nil,
			Error:             nil,
		})
	}

	return body
}

func (m *PlacementSolver) PbmCheckCompatibility(ctx *simulator.Context, _ *types.PbmCheckCompatibility) soap.HasFault {
	body := new(methods.PbmCheckCompatibilityBody)
	body.Res = new(types.PbmCheckCompatibilityResponse)

	for _, ds := range ctx.For(vim25.Path).Map.All("Datastore") {
		// TODO: filter
		ref := ds.Reference()
		body.Res.Returnval = append(body.Res.Returnval, types.PbmPlacementCompatibilityResult{
			Hub: types.PbmPlacementHub{
				HubType: ref.Type,
				HubId:   ref.Value,
			},
			MatchingResources: nil,
			HowMany:           0,
			Utilization:       nil,
			Warning:           nil,
			Error:             nil,
		})
	}

	return body
}
