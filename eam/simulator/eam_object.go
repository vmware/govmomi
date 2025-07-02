// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"time"

	"github.com/vmware/govmomi/eam/methods"
	"github.com/vmware/govmomi/eam/mo"
	"github.com/vmware/govmomi/eam/types"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

// EamObject contains the fields and functions common to all objects.
type EamObject mo.EamObject

func (m *EamObject) Reference() vim.ManagedObjectReference {
	return m.Self
}

func (m *EamObject) AddIssue(
	ctx *simulator.Context,
	req *types.AddIssue) soap.HasFault {

	// Get the typed issue to ensure the correct type of issue is stored and
	// returned to the caller.
	issue := issueType(req.Issue)

	// Get the base issue in order to assign an issue key and timestamp.
	baseIssue := issue.GetIssue()
	baseIssue.Key = nextAvailableIssueKey()
	baseIssue.Time = time.Now().UTC()

	// Store and return the typed issue.
	m.Issue = append(m.Issue, issue)

	return &methods.AddIssueBody{
		Res: &types.AddIssueResponse{
			Returnval: issue,
		},
	}
}

func (m *EamObject) QueryIssue(
	ctx *simulator.Context,
	req *types.QueryIssue) soap.HasFault {

	var issues []types.BaseIssue

	if len(req.IssueKey) == 0 {
		// If no keys were specified then return all issues.
		issues = m.Issue
	} else {
		// Get only the issues for the specified keys.
		for _, issueKey := range req.IssueKey {
			for _, issue := range m.Issue {
				if issue.GetIssue().Key == issueKey {
					issues = append(issues, issue)
				}
			}
		}
	}

	return &methods.QueryIssueBody{
		Res: &types.QueryIssueResponse{
			Returnval: issues,
		},
	}
}

func (m *EamObject) Resolve(
	ctx *simulator.Context,
	req *types.Resolve) soap.HasFault {

	// notFoundKeys is a list of issue keys that were sent but
	// not found for the given object.
	notFoundKeys := []int32{}

	// issueExists is a helper function that returns true
	issueExists := func(issueKey int32) bool {
		for _, k := range req.IssueKey {
			if k == issueKey {
				return true
			}
		}
		return false
	}

	// Iterate over the object's issues, and if a key matches, then remove
	// the issue from the list of the object's issues. If a key does not match
	// then record the key as notFound.
	for i := 0; i < len(m.Issue); i++ {
		issueKey := m.Issue[i].GetIssue().Key

		if ok := issueExists(issueKey); ok {
			// Update the object's issue list so that it no longer includes
			// the current issue.
			m.Issue = append(m.Issue[:i], m.Issue[i+1:]...)
			i--

			// Ensure the key is removed from the global key space.
			freeIssueKey(issueKey)
		} else {
			notFoundKeys = append(notFoundKeys, issueKey)
		}
	}

	return &methods.ResolveBody{
		Res: &types.ResolveResponse{
			Returnval: notFoundKeys,
		},
	}
}

func (m *EamObject) ResolveAll(
	ctx *simulator.Context,
	req *types.ResolveAll) soap.HasFault {

	// Iterate over the issues and ensure each one of their keys are removed
	// from the global key space.
	for _, issue := range m.Issue {
		freeIssueKey(issue.GetIssue().Key)
	}

	// Reset the object's issues.
	m.Issue = m.Issue[:0]

	return &methods.ResolveAllBody{Res: &types.ResolveAllResponse{}}
}
