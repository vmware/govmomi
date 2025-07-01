// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"time"

	"github.com/vmware/govmomi/pbm/types"
	vim "github.com/vmware/govmomi/vim25/types"
)

const DefaultEncryptionProfileID = "4d5f673c-536f-11e6-beb8-9e71128cae77"

var defaultEncryptionProfile = &types.PbmCapabilityProfile{
	PbmProfile: types.PbmProfile{
		ProfileId: types.PbmProfileId{
			UniqueId: DefaultEncryptionProfileID,
		},
		Name:            "VM Encryption Policy",
		Description:     "Sample storage policy for VMware's VM and virtual disk encryption",
		CreationTime:    time.Now(),
		CreatedBy:       "Temporary user handle",
		LastUpdatedTime: time.Now(),
		LastUpdatedBy:   "Temporary user handle",
	},
	ProfileCategory: "REQUIREMENT",
	ResourceType: types.PbmProfileResourceType{
		ResourceType: "STORAGE",
	},
	Constraints: &types.PbmCapabilitySubProfileConstraints{
		PbmCapabilityConstraints: types.PbmCapabilityConstraints{},
		SubProfiles: []types.PbmCapabilitySubProfile{
			{
				Name: "sp-1",
				Capability: []types.PbmCapabilityInstance{
					{
						Id: types.PbmCapabilityMetadataUniqueId{
							Namespace: "com.vmware.storageprofile.dataservice",
							Id:        "ad5a249d-cbc2-43af-9366-694d7664fa52",
						},
						Constraint: []types.PbmCapabilityConstraintInstance{
							{
								PropertyInstance: []types.PbmCapabilityPropertyInstance{
									{
										Id:       "ad5a249d-cbc2-43af-9366-694d7664fa52",
										Operator: "",
										Value:    "ad5a249d-cbc2-43af-9366-694d7664fa52",
									},
								},
							},
						},
					},
				},
				ForceProvision: vim.NewBool(false),
			},
		},
	},
	GenerationId:             0,
	IsDefault:                false,
	SystemCreatedProfileType: "",
	LineOfService:            "",
}

// vcenter67DefaultProfiles is a captured from vCenter 6.7's default set of PBM
// profiles.
var vcenter67DefaultProfiles = []types.BasePbmProfile{
	&types.PbmCapabilityProfile{
		PbmProfile: types.PbmProfile{
			ProfileId: types.PbmProfileId{
				UniqueId: "aa6d5a82-1c88-45da-85d3-3d74b91a5bad",
			},
			Name:            "vSAN Default Storage Policy",
			Description:     "Storage policy used as default for vSAN datastores",
			CreationTime:    time.Now(),
			CreatedBy:       "Temporary user handle",
			LastUpdatedTime: time.Now(),
			LastUpdatedBy:   "Temporary user handle",
		},
		ProfileCategory: "REQUIREMENT",
		ResourceType: types.PbmProfileResourceType{
			ResourceType: "STORAGE",
		},
		Constraints: &types.PbmCapabilitySubProfileConstraints{
			PbmCapabilityConstraints: types.PbmCapabilityConstraints{},
			SubProfiles: []types.PbmCapabilitySubProfile{
				{
					Name: "VSAN sub-profile",
					Capability: []types.PbmCapabilityInstance{
						{
							Id: types.PbmCapabilityMetadataUniqueId{
								Namespace: "VSAN",
								Id:        "hostFailuresToTolerate",
							},
							Constraint: []types.PbmCapabilityConstraintInstance{
								{
									PropertyInstance: []types.PbmCapabilityPropertyInstance{
										{
											Id:       "hostFailuresToTolerate",
											Operator: "",
											Value:    int32(1),
										},
									},
								},
							},
						},
						{
							Id: types.PbmCapabilityMetadataUniqueId{
								Namespace: "VSAN",
								Id:        "stripeWidth",
							},
							Constraint: []types.PbmCapabilityConstraintInstance{
								{
									PropertyInstance: []types.PbmCapabilityPropertyInstance{
										{
											Id:       "stripeWidth",
											Operator: "",
											Value:    int32(1),
										},
									},
								},
							},
						},
						{
							Id: types.PbmCapabilityMetadataUniqueId{
								Namespace: "VSAN",
								Id:        "forceProvisioning",
							},
							Constraint: []types.PbmCapabilityConstraintInstance{
								{
									PropertyInstance: []types.PbmCapabilityPropertyInstance{
										{
											Id:       "forceProvisioning",
											Operator: "",
											Value:    bool(false),
										},
									},
								},
							},
						},
						{
							Id: types.PbmCapabilityMetadataUniqueId{
								Namespace: "VSAN",
								Id:        "proportionalCapacity",
							},
							Constraint: []types.PbmCapabilityConstraintInstance{
								{
									PropertyInstance: []types.PbmCapabilityPropertyInstance{
										{
											Id:       "proportionalCapacity",
											Operator: "",
											Value:    int32(0),
										},
									},
								},
							},
						},
						{
							Id: types.PbmCapabilityMetadataUniqueId{
								Namespace: "VSAN",
								Id:        "cacheReservation",
							},
							Constraint: []types.PbmCapabilityConstraintInstance{
								{
									PropertyInstance: []types.PbmCapabilityPropertyInstance{
										{
											Id:       "cacheReservation",
											Operator: "",
											Value:    int32(0),
										},
									},
								},
							},
						},
					},
					ForceProvision: (*bool)(nil),
				},
			},
		},
		GenerationId:             0,
		IsDefault:                false,
		SystemCreatedProfileType: "VsanDefaultProfile",
		LineOfService:            "",
	},
	&types.PbmCapabilityProfile{
		PbmProfile: types.PbmProfile{
			ProfileId: types.PbmProfileId{
				UniqueId: "f4e5bade-15a2-4805-bf8e-52318c4ce443",
			},
			Name:            "VVol No Requirements Policy",
			Description:     "Allow the datastore to determine the best placement strategy for storage objects",
			CreationTime:    time.Now(),
			CreatedBy:       "Temporary user handle",
			LastUpdatedTime: time.Now(),
			LastUpdatedBy:   "Temporary user handle",
		},
		ProfileCategory: "REQUIREMENT",
		ResourceType: types.PbmProfileResourceType{
			ResourceType: "STORAGE",
		},
		Constraints:              &types.PbmCapabilityConstraints{},
		GenerationId:             0,
		IsDefault:                false,
		SystemCreatedProfileType: "VVolDefaultProfile",
		LineOfService:            "",
	},
	defaultEncryptionProfile,
	&types.PbmCapabilityProfile{
		PbmProfile: types.PbmProfile{
			ProfileId: types.PbmProfileId{
				UniqueId: "c268da1b-b343-49f7-a468-b1deeb7078e0",
			},
			Name:            "Host-local PMem Default Storage Policy",
			Description:     "Storage policy used as default for Host-local PMem datastores",
			CreationTime:    time.Now(),
			CreatedBy:       "Temporary user handle",
			LastUpdatedTime: time.Now(),
			LastUpdatedBy:   "Temporary user handle",
		},
		ProfileCategory: "REQUIREMENT",
		ResourceType: types.PbmProfileResourceType{
			ResourceType: "STORAGE",
		},
		Constraints: &types.PbmCapabilitySubProfileConstraints{
			PbmCapabilityConstraints: types.PbmCapabilityConstraints{},
			SubProfiles: []types.PbmCapabilitySubProfile{
				{
					Name: "PMem sub-profile",
					Capability: []types.PbmCapabilityInstance{
						{
							Id: types.PbmCapabilityMetadataUniqueId{
								Namespace: "PMem",
								Id:        "PMemType",
							},
							Constraint: []types.PbmCapabilityConstraintInstance{
								{
									PropertyInstance: []types.PbmCapabilityPropertyInstance{
										{
											Id:       "PMemType",
											Operator: "",
											Value:    "LocalPMem",
										},
									},
								},
							},
						},
					},
					ForceProvision: (*bool)(nil),
				},
			},
		},
		GenerationId:             0,
		IsDefault:                false,
		SystemCreatedProfileType: "PmemDefaultProfile",
		LineOfService:            "",
	},
}
