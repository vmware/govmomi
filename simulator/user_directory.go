/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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
	"strings"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type UserDirectoryBackend interface {
	AddEntity(string, bool)
	RemoveEntity(string, bool)
	SearchEntities(bool, bool, func(string) bool) []types.BaseUserSearchResult
}

type userSearchResultList struct {
	results []*types.UserSearchResult
}

func (l *userSearchResultList) AddEntity(id string, group bool) {
	l.results = append(l.results, &types.UserSearchResult{
		FullName:  id,
		Group:     group,
		Principal: id,
	})
}

func (l *userSearchResultList) RemoveEntity(id string, group bool) {
	for i, ug := range l.results {
		if ug.Group == group && ug.Principal == id {
			l.results = append(l.results[:i], l.results[i+1:]...)
			return
		}
	}
}

func (l *userSearchResultList) SearchEntities(users, groups bool, principalFilter func(string) bool) (res []types.BaseUserSearchResult) {
	for _, r := range l.results {
		if users && !r.Group || groups && r.Group {
			if principalFilter(r.Principal) {
				res = append(res, r)
			}
		}
	}
	return
}

var DefaultUserGroup = &userSearchResultList{
	results: []*types.UserSearchResult{
		{FullName: "root", Group: true, Principal: "root"},
		{FullName: "root", Group: false, Principal: "root"},
		{FullName: "administrator", Group: false, Principal: "admin"},
	},
}

type UserDirectory struct {
	mo.UserDirectory

	backend *userSearchResultList
}

func (u *UserDirectory) Backend() UserDirectoryBackend {
	return u.backend
}

func (m *UserDirectory) init(*Registry) {
	m.backend = DefaultUserGroup
}

func (u *UserDirectory) RetrieveUserGroups(req *types.RetrieveUserGroups) soap.HasFault {
	compare := compareFunc(req.SearchStr, req.ExactMatch)

	res := u.backend.SearchEntities(req.FindUsers, req.FindGroups, compare)

	body := &methods.RetrieveUserGroupsBody{
		Res: &types.RetrieveUserGroupsResponse{
			Returnval: res,
		},
	}

	return body
}

func compareFunc(compared string, exactly bool) func(string) bool {
	return func(s string) bool {
		if exactly {
			return s == compared
		}
		return strings.Contains(strings.ToLower(s), strings.ToLower(compared))
	}
}
