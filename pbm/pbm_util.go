// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package pbm

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/vmware/govmomi/pbm/types"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	vim "github.com/vmware/govmomi/vim25/types"
)

// A struct to capture pbm create spec details.
type CapabilityProfileCreateSpec struct {
	Name             string
	SubProfileName   string
	Description      string
	Category         string
	CapabilityList   []Capability
	K8sCompliantName string
}

// A struct to capture pbm capability instance details.
type Capability struct {
	ID           string
	Namespace    string
	PropertyList []Property
}

// A struct to capture pbm property instance details.
type Property struct {
	ID       string
	Operator string
	Value    string
	DataType string
}

func CreateCapabilityProfileSpec(pbmCreateSpec CapabilityProfileCreateSpec) (*types.PbmCapabilityProfileCreateSpec, error) {
	capabilities, err := createCapabilityInstances(pbmCreateSpec.CapabilityList)
	if err != nil {
		return nil, err
	}

	pbmCapabilityProfileSpec := types.PbmCapabilityProfileCreateSpec{
		Name:        pbmCreateSpec.Name,
		Description: pbmCreateSpec.Description,
		Category:    pbmCreateSpec.Category,
		ResourceType: types.PbmProfileResourceType{
			ResourceType: string(types.PbmProfileResourceTypeEnumSTORAGE),
		},
		Constraints: &types.PbmCapabilitySubProfileConstraints{
			SubProfiles: []types.PbmCapabilitySubProfile{
				types.PbmCapabilitySubProfile{
					Capability: capabilities,
					Name:       pbmCreateSpec.SubProfileName,
				},
			},
		},
		K8sCompliantName: pbmCreateSpec.K8sCompliantName,
	}
	return &pbmCapabilityProfileSpec, nil
}

func createCapabilityInstances(rules []Capability) ([]types.PbmCapabilityInstance, error) {
	var capabilityInstances []types.PbmCapabilityInstance
	for _, capabilityRule := range rules {
		capability := types.PbmCapabilityInstance{
			Id: types.PbmCapabilityMetadataUniqueId{
				Namespace: capabilityRule.Namespace,
				Id:        capabilityRule.ID,
			},
		}

		var propertyInstances []types.PbmCapabilityPropertyInstance
		for _, propertyRule := range capabilityRule.PropertyList {
			property := types.PbmCapabilityPropertyInstance{
				Id: propertyRule.ID,
			}
			if propertyRule.Operator != "" {
				property.Operator = propertyRule.Operator
			}
			var err error
			switch strings.ToLower(propertyRule.DataType) {
			case "int":
				// Go int32 is marshalled to xsi:int whereas Go int is marshalled to xsi:long when sending down the wire.
				var val int32
				val, err = verifyPropertyValueIsInt(propertyRule.Value, propertyRule.DataType)
				property.Value = val
			case "bool":
				var val bool
				val, err = verifyPropertyValueIsBoolean(propertyRule.Value, propertyRule.DataType)
				property.Value = val
			case "string":
				property.Value = propertyRule.Value
			case "set":
				set := types.PbmCapabilityDiscreteSet{}
				for _, val := range strings.Split(propertyRule.Value, ",") {
					set.Values = append(set.Values, val)
				}
				property.Value = set
			default:
				return nil, fmt.Errorf("invalid value: %q with datatype: %q", propertyRule.Value, propertyRule.Value)
			}
			if err != nil {
				return nil, fmt.Errorf("invalid value: %q with datatype: %q", propertyRule.Value, propertyRule.Value)
			}
			propertyInstances = append(propertyInstances, property)
		}
		constraintInstances := []types.PbmCapabilityConstraintInstance{
			types.PbmCapabilityConstraintInstance{
				PropertyInstance: propertyInstances,
			},
		}
		capability.Constraint = constraintInstances
		capabilityInstances = append(capabilityInstances, capability)
	}
	return capabilityInstances, nil
}

// Verify if the capability value is of type integer.
func verifyPropertyValueIsInt(propertyValue string, dataType string) (int32, error) {
	val, err := strconv.ParseInt(propertyValue, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(val), nil
}

// Verify if the capability value is of type integer.
func verifyPropertyValueIsBoolean(propertyValue string, dataType string) (bool, error) {
	val, err := strconv.ParseBool(propertyValue)
	if err != nil {
		return false, err
	}
	return val, nil
}

// ProfileMap contains a map of storage profiles by name.
type ProfileMap struct {
	Name    map[string]types.BasePbmProfile
	Profile []types.BasePbmProfile
}

// ProfileMap builds a map of storage profiles by name.
func (c *Client) ProfileMap(ctx context.Context, uid ...string) (*ProfileMap, error) {
	m := &ProfileMap{Name: make(map[string]types.BasePbmProfile)}

	rtype := types.PbmProfileResourceType{
		ResourceType: string(types.PbmProfileResourceTypeEnumSTORAGE),
	}

	category := types.PbmProfileCategoryEnumREQUIREMENT

	var ids []types.PbmProfileId
	if len(uid) == 0 {
		var err error
		ids, err = c.QueryProfile(ctx, rtype, string(category))
		if err != nil {
			return nil, err
		}
	} else {
		ids = make([]types.PbmProfileId, len(uid))
		for i, id := range uid {
			ids[i].UniqueId = id
		}
	}

	profiles, err := c.RetrieveContent(ctx, ids)
	if err != nil {
		return nil, err
	}
	m.Profile = profiles

	for _, p := range profiles {
		base := p.GetPbmProfile()
		m.Name[base.Name] = p
		m.Name[base.ProfileId.UniqueId] = p
	}

	return m, nil
}

// DatastoreMap contains a map of Datastore by name.
type DatastoreMap struct {
	Name         map[string]string
	PlacementHub []types.PbmPlacementHub
}

// DatastoreMap returns a map of Datastore by name.
// The root reference can be a ClusterComputeResource or Folder.
func (c *Client) DatastoreMap(ctx context.Context, vc *vim25.Client, root vim.ManagedObjectReference) (*DatastoreMap, error) {
	m := &DatastoreMap{Name: make(map[string]string)}

	prop := []string{"name"}
	var content []vim.ObjectContent

	if root.Type == "ClusterComputeResource" {
		pc := property.DefaultCollector(vc)
		var cluster mo.ClusterComputeResource

		if err := pc.RetrieveOne(ctx, root, []string{"datastore"}, &cluster); err != nil {
			return nil, err
		}

		if err := pc.Retrieve(ctx, cluster.Datastore, prop, &content); err != nil {
			return nil, err
		}
	} else {
		kind := []string{"Datastore"}
		m := view.NewManager(vc)

		v, err := m.CreateContainerView(ctx, root, kind, true)
		if err != nil {
			return nil, err
		}

		err = v.Retrieve(ctx, kind, prop, &content)
		_ = v.Destroy(ctx)
		if err != nil {
			return nil, err
		}
	}

	for _, item := range content {
		m.PlacementHub = append(m.PlacementHub, types.PbmPlacementHub{
			HubType: item.Obj.Type,
			HubId:   item.Obj.Value,
		})
		m.Name[item.Obj.Value] = item.PropSet[0].Val.(string)
	}

	return m, nil
}
