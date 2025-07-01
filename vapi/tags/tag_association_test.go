// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package tags_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vapi/tags"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

func TestManager_AttachMultipleTagsToObject(t *testing.T) {
	// tags to create in the system before executing tests
	// returned IDs will be appended to urns in args struct
	type fields struct {
		namedTags []string
	}

	type args struct {
		// if fields.namedTags are specified this field will be populated with the related URNs
		urns []string
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantTags int // number of expected tags attached to ref
		wantErr  error
	}{
		{
			name:   "tag not in URN-format",
			fields: fields{}, // no pre-existing tags required for this test case
			args: args{
				urns: []string{"not-urn-tag"},
			},
			wantTags: 0,
			wantErr:  errors.New("specified tag is not a URN: \"not-urn-tag\""),
		},
		{
			name: "two valid tags on one category",
			fields: fields{
				namedTags: []string{"valid-tag-1", "valid-tag-2"},
			},
			args:     args{}, // will be auto-populated
			wantTags: 2,
			wantErr:  nil,
		},
		{
			name: "one valid and one invalid (not-URN) tag on one category",
			fields: fields{
				namedTags: []string{"valid-tag"}, // create one valid tag before running test
			},
			args: args{
				urns: []string{"not-urn-tag"}, // force one additional non-URN tag
			},
			wantTags: 0,
			wantErr:  errors.New("specified tag is not a URN: \"not-urn-tag\""),
		},
		{
			name: "one valid and one not existing tag on one category",
			fields: fields{
				namedTags: []string{"valid-tag"}, // create one valid tag before running test
			},
			args: args{
				urns: []string{"urn:vmomi:InventoryServiceTag:31e55277-ca60-482a-899b-232184be224c:GLOBAL"}, // does not exist
			},
			wantTags: 1,
			wantErr: tags.BatchErrors{
				{
					Type:    "cis.tagging.objectNotFound.error",
					Message: "Tagging object urn:vmomi:InventoryServiceTag:31e55277-ca60-482a-899b-232184be224c:GLOBAL not found",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			simulator.Run(func(ctx context.Context, vc *vim25.Client) error {
				vm, err := getRef(t, ctx, vc)
				if err != nil {
					t.Errorf("get virtual machine: %v", err)
				}

				c := rest.NewClient(vc)
				_ = c.Login(ctx, simulator.DefaultLogin)

				m := tags.NewManager(c)

				// seed simulator with URNs if required for this test
				if len(tt.fields.namedTags) > 0 {
					idMap, err := createTags(t, ctx, m, tt.fields.namedTags)
					if err != nil {
						t.Errorf("set up tags: %v", err)
					}

					for _, id := range idMap {
						tt.args.urns = append(tt.args.urns, id)
					}
				}

				err = m.AttachMultipleTagsToObject(ctx, tt.args.urns, vm)
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("AttachMultipleTagsToObject() error = %v, wantErr %v", err, tt.wantErr)
				}

				attached, err := getTags(t, ctx, m, vm)
				if len(attached) != tt.wantTags {
					t.Errorf("AttachMultipleTagsToObject() attachedTags = %d, wantTags %d", len(attached), tt.wantTags)
				}

				return nil
			})
		})
	}
}

func TestManager_DetachMultipleTagsFromObject(t *testing.T) {
	// tags which will be optionally created and attached to the target ref
	// before executing tests - returned IDs for created tags will be appended to
	// urns in args struct
	type field struct {
		name   string
		create bool
		attach bool
	}

	type args struct {
		// if fields.namedTags && .create are specified this field will be populated
		// with the related URNs
		urns []string
	}

	tests := []struct {
		name     string
		fields   []field
		args     args
		wantTags int // number of expected tags attached to ref
		wantErr  error
	}{
		{
			name:   "tag not in URN-format",
			fields: []field{}, // no pre-existing tags required for this test case
			args: args{
				urns: []string{"not-urn-tag"},
			},
			wantTags: 0,
			wantErr:  errors.New("specified tag is not a URN: \"not-urn-tag\""),
		},
		{
			name: "no-op",
			fields: []field{
				{
					name:   "valid-tag",
					create: true,
					attach: false, // exists but not attached
				},
			},
			args:     args{}, // will be auto-populated
			wantTags: 0,
			wantErr:  nil,
		},
		{
			name: "two valid tags on one category",
			fields: []field{
				{
					name:   "valid-tag-1",
					create: true,
					attach: true,
				},
				{
					name:   "valid-tag-2",
					create: true,
					attach: true,
				},
			},
			args:     args{}, // will be auto-populated
			wantTags: 0,
			wantErr:  nil,
		},
		{
			name: "one valid and one invalid (not-URN) tag on one category",
			fields: []field{
				{
					name:   "valid-tag-1",
					create: true,
					attach: true,
				},
			},
			args: args{
				urns: []string{"not-urn-tag"}, // force one additional non-URN tag
			},
			wantTags: 1, // won't be detached
			wantErr:  errors.New("specified tag is not a URN: \"not-urn-tag\""),
		},
		{
			name: "one valid and one not existing tag on one category",
			fields: []field{
				{
					name:   "valid-tag-1",
					create: true,
					attach: true,
				},
			},
			args: args{
				urns: []string{"urn:vmomi:InventoryServiceTag:31e55277-ca60-482a-899b-232184be224c:GLOBAL"}, // does not exist
			},
			wantTags: 0, // existing will be detached
			wantErr: tags.BatchErrors{
				{
					Type:    "cis.tagging.objectNotFound.error",
					Message: "Tagging object urn:vmomi:InventoryServiceTag:31e55277-ca60-482a-899b-232184be224c:GLOBAL not found",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			simulator.Run(func(ctx context.Context, vc *vim25.Client) error {
				vm, err := getRef(t, ctx, vc)
				if err != nil {
					t.Errorf("get virtual machine: %v", err)
				}

				c := rest.NewClient(vc)
				_ = c.Login(ctx, simulator.DefaultLogin)

				m := tags.NewManager(c)

				// seed simulator with URNs and attach to ref if required for this test
				if len(tt.fields) > 0 {
					var create []string
					for _, t := range tt.fields {
						if t.create {
							create = append(create, t.name)
						}
					}

					idMap, err := createTags(t, ctx, m, create)
					if err != nil {
						t.Errorf("create tags: %v", err)
					}

					// transform name tags to created tag URNs before attaching
					var attach []string
					for _, t := range tt.fields {
						if t.attach {
							attach = append(attach, idMap[t.name])
						}
					}

					err = attachTags(t, ctx, m, vm, attach)
					if err != nil {
						t.Errorf("attach tags: %v", err)
					}

					for _, id := range idMap {
						tt.args.urns = append(tt.args.urns, id)
					}
				}

				err = m.DetachMultipleTagsFromObject(ctx, tt.args.urns, vm)
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("DetachMultipleTagsFromObject() error = %v, wantErr %v", err, tt.wantErr)
				}

				attached, err := getTags(t, ctx, m, vm)
				if len(attached) != tt.wantTags {
					t.Errorf("DetachMultipleTagsFromObject() attachedTags = %d, wantTags %d", len(attached), tt.wantTags)
				}

				return nil
			})
		})
	}
}

// createTags creates the given tag to category mappings and returns a map of
// names to IDs (URNs) for all created tags
func createTags(t *testing.T, ctx context.Context, mgr *tags.Manager, tagNames []string) (map[string]string, error) {
	t.Helper()

	cat := tags.Category{
		Name:        "test-category-1",
		Description: "category used for testing against simulator",
		Cardinality: "MULTIPLE", // simulator currently does not support cardinality validation
	}

	catID, err := mgr.CreateCategory(ctx, &cat)
	if err != nil {
		return nil, err
	}

	mapping := map[string]string{}
	for _, name := range tagNames {
		id, err := mgr.CreateTag(ctx, &tags.Tag{Name: name, CategoryID: catID})
		if err != nil {
			return nil, err
		}
		mapping[name] = id
	}

	return mapping, nil
}

// attachTags attaches the given tags on the given ref
func attachTags(t *testing.T, ctx context.Context, mgr *tags.Manager, ref mo.Reference, tagIDs []string) error {
	t.Helper()

	for _, id := range tagIDs {
		err := mgr.AttachTag(ctx, id, ref)
		if err != nil {
			return err
		}
	}

	return nil
}

func getTags(t *testing.T, ctx context.Context, mgr *tags.Manager, ref mo.Reference) ([]tags.Tag, error) {
	t.Helper()

	attached, err := mgr.GetAttachedTags(ctx, ref)
	if err != nil {
		return nil, err
	}

	return attached, nil
}

// getRef returns the first virtual machine found in the inventory
func getRef(t *testing.T, ctx context.Context, client *vim25.Client) (mo.Reference, error) {
	t.Helper()

	// Create view of VirtualMachine objects
	m := view.NewManager(client)

	v, err := m.CreateContainerView(ctx, client.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		return nil, err
	}

	defer v.Destroy(ctx)

	// Retrieve summary property for all machines
	// Reference: https://developer.broadcom.com/xapis/vsphere-web-services-api/latest/vim.VirtualMachine.html
	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		return nil, err
	}

	if len(vms) < 1 {
		return nil, errors.New("no existing virtual machine found")
	}
	return vms[0].Self, nil
}
