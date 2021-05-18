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

package object

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/eam"
	"github.com/vmware/govmomi/eam/methods"
	"github.com/vmware/govmomi/eam/types"
	vim "github.com/vmware/govmomi/vim25/types"
)

// EamObject contains the fields and functions common to all objects.
type EamObject struct {
	c *eam.Client
	r vim.ManagedObjectReference
}

func (m EamObject) String() string {
	return fmt.Sprintf("%v", m.Reference())
}

func (m EamObject) Reference() vim.ManagedObjectReference {
	return m.r
}

func (m EamObject) Client() *eam.Client {
	return m.c
}

func (m EamObject) AddIssue(
	ctx context.Context,
	issue types.BaseIssue) (types.BaseIssue, error) {

	resp, err := methods.AddIssue(ctx, m.c, &types.AddIssue{
		This:  m.r,
		Issue: issue,
	})
	if err != nil {
		return nil, err
	}
	return resp.Returnval, nil
}

func (m EamObject) Issues(
	ctx context.Context,
	issueKeys ...int32) ([]types.BaseIssue, error) {

	resp, err := methods.QueryIssue(ctx, m.c, &types.QueryIssue{
		This:     m.r,
		IssueKey: issueKeys,
	})
	if err != nil {
		return nil, err
	}
	return resp.Returnval, nil
}

func (m EamObject) Resolve(
	ctx context.Context,
	issueKeys []int32) ([]int32, error) {

	resp, err := methods.Resolve(ctx, m.c, &types.Resolve{
		This:     m.r,
		IssueKey: issueKeys,
	})
	if err != nil {
		return nil, err
	}
	return resp.Returnval, nil
}

func (m EamObject) ResolveAll(ctx context.Context) error {

	_, err := methods.ResolveAll(ctx, m.c, &types.ResolveAll{
		This: m.r,
	})
	if err != nil {
		return err
	}
	return nil
}
