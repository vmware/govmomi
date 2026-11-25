// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"net/http"
	"slices"
	"sort"
	"strconv"

	"github.com/vmware/govmomi/vapi/tagging"
)

// page applies the "names" and "marker" and "page_size" query parameters of the
// "/api/vcenter/tagging" list operations to the given ids (already name filtered).
// The marker is the last id of the previous page.
// Returns the ids of the requested page and the marker for the next page, if any.
// On failure, the error response has been sent and ok is false.
func (s *handler) page(w http.ResponseWriter, r *http.Request, ids []string) (page []string, marker string, ok bool) {
	q := r.URL.Query()

	size := tagging.DefaultListPageSize
	if v := q.Get("page_size"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 1 {
			apiError(w, http.StatusBadRequest, "INVALID_ARGUMENT")
			return nil, "", false
		}
		size = n
	}

	sort.Strings(ids)

	if m := q.Get("marker"); m != "" {
		i, found := slices.BinarySearch(ids, m)
		if !found {
			// The marker may refer to an item that has since been deleted,
			// but an unknown marker is the documented NOT_FOUND case.
			apiError(w, http.StatusNotFound, "NOT_FOUND")
			return nil, "", false
		}
		ids = ids[i+1:]
	}

	if len(ids) > size {
		ids = ids[:size]
		marker = ids[size-1]
	}

	return ids, marker, true
}

func (s *handler) apiCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	names := r.URL.Query()["names"]

	var ids []string
	for id, c := range s.Category {
		if len(names) == 0 || slices.Contains(names, c.Name) {
			ids = append(ids, id)
		}
	}

	ids, marker, ok := s.page(w, r, ids)
	if !ok {
		return
	}

	res := tagging.CategoryListResult{Marker: marker, Items: []tagging.CategoryListItem{}}
	for _, id := range ids {
		c := s.Category[id]
		res.Items = append(res.Items, tagging.CategoryListItem{
			CategoryID: id,
			Info: tagging.CategoryInfo{
				Name:            c.Name,
				Description:     c.Description,
				Cardinality:     c.Cardinality,
				AssociableTypes: c.AssociableTypes,
				UsedBy:          c.UsedBy,
			},
		})
	}

	StatusOK(w, res)
}

func (s *handler) apiTags(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	names := r.URL.Query()["names"]

	var ids []string
	for id, t := range s.Tag {
		if len(names) == 0 || slices.Contains(names, t.Name) {
			ids = append(ids, id)
		}
	}

	ids, marker, ok := s.page(w, r, ids)
	if !ok {
		return
	}

	res := tagging.TagListResult{Marker: marker, Items: []tagging.TagListItem{}}
	for _, id := range ids {
		t := s.Tag[id]
		res.Items = append(res.Items, tagging.TagListItem{
			Tag: id,
			Info: tagging.TagInfo{
				Name:        t.Name,
				Category:    t.CategoryID,
				Description: t.Description,
				UsedBy:      t.UsedBy,
			},
		})
	}

	StatusOK(w, res)
}
