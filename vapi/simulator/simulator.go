/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

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
	"archive/tar"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/nfc"
	"github.com/vmware/govmomi/ovf"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/internal"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vapi/tags"
	"github.com/vmware/govmomi/vapi/vcenter"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
	vim "github.com/vmware/govmomi/vim25/types"
)

type item struct {
	*library.Item
	File []library.File
}

type content struct {
	*library.Library
	Item map[string]*item
}

type update struct {
	*library.Session
	Library *library.Library
	File    map[string]*library.UpdateFile
}

type download struct {
	*library.Session
	Library *library.Library
	File    map[string]*library.DownloadFile
}

type handler struct {
	sync.Mutex
	ServeMux    *http.ServeMux
	URL         url.URL
	Category    map[string]*tags.Category
	Tag         map[string]*tags.Tag
	Association map[string]map[internal.AssociatedObject]bool
	Session     map[string]*rest.Session
	Library     map[string]content
	Update      map[string]update
	Download    map[string]download
}

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		if r.IsVPX() {
			path, handler := New(s.Listen, r.OptionManager().Setting)
			s.Handle(path, handler)
		}
	})
}

// New creates a vAPI simulator.
func New(u *url.URL, settings []vim.BaseOptionValue) (string, http.Handler) {
	s := &handler{
		ServeMux:    http.NewServeMux(),
		URL:         *u,
		Category:    make(map[string]*tags.Category),
		Tag:         make(map[string]*tags.Tag),
		Association: make(map[string]map[internal.AssociatedObject]bool),
		Session:     make(map[string]*rest.Session),
		Library:     make(map[string]content),
		Update:      make(map[string]update),
		Download:    make(map[string]download),
	}

	handlers := []struct {
		p string
		m http.HandlerFunc
	}{
		{internal.SessionPath, s.session},
		{internal.CategoryPath, s.category},
		{internal.CategoryPath + "/", s.categoryID},
		{internal.TagPath, s.tag},
		{internal.TagPath + "/", s.tagID},
		{internal.AssociationPath, s.association},
		{internal.AssociationPath + "/", s.associationID},
		{internal.LibraryPath, s.library},
		{internal.LocalLibraryPath, s.library},
		{internal.SubscribedLibraryPath, s.library},
		{internal.LibraryPath + "/", s.libraryID},
		{internal.LocalLibraryPath + "/", s.libraryID},
		{internal.LibraryItemPath, s.libraryItem},
		{internal.LibraryItemPath + "/", s.libraryItemID},
		{internal.LibraryItemUpdateSession, s.libraryItemUpdateSession},
		{internal.LibraryItemUpdateSession + "/", s.libraryItemUpdateSessionID},
		{internal.LibraryItemUpdateSessionFile, s.libraryItemUpdateSessionFile},
		{internal.LibraryItemUpdateSessionFile + "/", s.libraryItemUpdateSessionFileID},
		{internal.LibraryItemDownloadSession, s.libraryItemDownloadSession},
		{internal.LibraryItemDownloadSession + "/", s.libraryItemDownloadSessionID},
		{internal.LibraryItemDownloadSessionFile, s.libraryItemDownloadSessionFile},
		{internal.LibraryItemDownloadSessionFile + "/", s.libraryItemDownloadSessionFileID},
		{internal.LibraryItemFileData + "/", s.libraryItemFileData},
		{internal.LibraryItemFilePath, s.libraryItemFile},
		{internal.LibraryItemFilePath + "/", s.libraryItemFileID},
		{internal.VCenterOVFLibraryItem + "/", s.libraryItemDeployID},
	}

	for i := range handlers {
		h := handlers[i]
		s.HandleFunc(h.p, h.m)
	}

	return rest.Path + "/", s
}

// HandleFunc wraps the given handler with authorization checks and passes to http.ServeMux.HandleFunc
func (s *handler) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	if !strings.HasPrefix(pattern, rest.Path) {
		pattern = rest.Path + pattern
	}

	s.ServeMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		s.Lock()
		defer s.Unlock()

		if !s.isAuthorized(r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		handler(w, r)
	})
}

func (s *handler) isAuthorized(r *http.Request) bool {
	if r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, internal.SessionPath) && s.action(r) == "" {
		return true
	}
	id := r.Header.Get(internal.SessionCookieName)
	if id == "" {
		if cookie, err := r.Cookie(internal.SessionCookieName); err == nil {
			id = cookie.Value
			r.Header.Set(internal.SessionCookieName, id)
		}
	}
	info, ok := s.Session[id]
	if ok {
		info.LastAccessed = time.Now()
	} else {
		_, ok = s.Update[id]
	}
	return ok
}

func (s *handler) hasAuthorization(r *http.Request) (string, bool) {
	u, p, ok := r.BasicAuth()
	if ok { // user+pass auth
		if u == "" || p == "" {
			return u, false
		}
		return u, true
	}
	auth := r.Header.Get("Authorization")
	return "TODO", strings.HasPrefix(auth, "SIGN ") // token auth
}

func (s *handler) findTag(e vim.VslmTagEntry) *tags.Tag {
	for _, c := range s.Category {
		if c.Name == e.ParentCategoryName {
			for _, t := range s.Tag {
				if t.Name == e.TagName && t.CategoryID == c.ID {
					return t
				}
			}
		}
	}
	return nil
}

// AttachedObjects is meant for internal use via simulator.Registry.tagManager
func (s *handler) AttachedObjects(tag vim.VslmTagEntry) ([]vim.ManagedObjectReference, vim.BaseMethodFault) {
	t := s.findTag(tag)
	if t == nil {
		return nil, new(vim.NotFound)
	}
	var ids []vim.ManagedObjectReference
	for id := range s.Association[t.ID] {
		ids = append(ids, vim.ManagedObjectReference(id))
	}
	return ids, nil
}

// AttachedTags is meant for internal use via simulator.Registry.tagManager
func (s *handler) AttachedTags(ref vim.ManagedObjectReference) ([]vim.VslmTagEntry, vim.BaseMethodFault) {
	oid := internal.AssociatedObject(ref)
	var tags []vim.VslmTagEntry
	for id, objs := range s.Association {
		if objs[oid] {
			tag := s.Tag[id]
			cat := s.Category[tag.CategoryID]
			tags = append(tags, vim.VslmTagEntry{
				TagName:            tag.Name,
				ParentCategoryName: cat.Name,
			})
		}
	}
	return tags, nil
}

// AttachTag is meant for internal use via simulator.Registry.tagManager
func (s *handler) AttachTag(ref vim.ManagedObjectReference, tag vim.VslmTagEntry) vim.BaseMethodFault {
	t := s.findTag(tag)
	if t == nil {
		return new(vim.NotFound)
	}
	s.Association[t.ID][internal.AssociatedObject(ref)] = true
	return nil
}

// DetachTag is meant for internal use via simulator.Registry.tagManager
func (s *handler) DetachTag(id vim.ManagedObjectReference, tag vim.VslmTagEntry) vim.BaseMethodFault {
	t := s.findTag(tag)
	if t == nil {
		return new(vim.NotFound)
	}
	delete(s.Association[t.ID], internal.AssociatedObject(id))
	return nil
}

// OK responds with http.StatusOK and json encoded val if given.
func OK(w http.ResponseWriter, val ...interface{}) {
	w.WriteHeader(http.StatusOK)

	if len(val) == 0 {
		return
	}

	err := json.NewEncoder(w).Encode(struct {
		Value interface{} `json:"value,omitempty"`
	}{
		val[0],
	})

	if err != nil {
		log.Panic(err)
	}
}

// BadRequest responds with http.StatusBadRequest and json encoded vAPI error of type kind.
func BadRequest(w http.ResponseWriter, kind string) {
	w.WriteHeader(http.StatusBadRequest)

	err := json.NewEncoder(w).Encode(struct {
		Type  string `json:"type"`
		Value struct {
			Messages []string `json:"messages,omitempty"`
		} `json:"value,omitempty"`
	}{
		Type: kind,
	})

	if err != nil {
		log.Panic(err)
	}
}

func (*handler) error(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	log.Print(err)
}

// ServeHTTP handles vAPI requests.
func (s *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost, http.MethodDelete, http.MethodGet, http.MethodPatch, http.MethodPut:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	h, _ := s.ServeMux.Handler(r)
	h.ServeHTTP(w, r)
}

func (s *handler) decode(r *http.Request, w http.ResponseWriter, val interface{}) bool {
	return Decode(r, w, val)
}

// Decode the request Body into val.
// Returns true on success, otherwise false and sends the http.StatusBadRequest response.
func Decode(r *http.Request, w http.ResponseWriter, val interface{}) bool {
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(val)
	if err != nil {
		log.Printf("%s %s: %s", r.Method, r.RequestURI, err)
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	return true
}

func (s *handler) session(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get(internal.SessionCookieName)

	switch r.Method {
	case http.MethodPost:
		if s.action(r) != "" {
			if session, ok := s.Session[id]; ok {
				OK(w, session)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
			return
		}
		user, ok := s.hasAuthorization(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		id = uuid.New().String()
		now := time.Now()
		s.Session[id] = &rest.Session{User: user, Created: now, LastAccessed: now}
		http.SetCookie(w, &http.Cookie{
			Name:  internal.SessionCookieName,
			Value: id,
			Path:  rest.Path,
		})
		OK(w, id)
	case http.MethodDelete:
		delete(s.Session, id)
		OK(w)
	case http.MethodGet:
		OK(w, s.Session[id])
	}
}

func (s *handler) action(r *http.Request) string {
	return r.URL.Query().Get("~action")
}

func (s *handler) id(r *http.Request) string {
	base := path.Base(r.URL.Path)
	id := strings.TrimPrefix(base, "id:")
	if id == base {
		return "" // trigger 404 Not Found w/o id: prefix
	}
	return id
}

func newID(kind string) string {
	return fmt.Sprintf("urn:vmomi:InventoryService%s:%s:GLOBAL", kind, uuid.New().String())
}

func (s *handler) category(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var spec struct {
			Category tags.Category `json:"create_spec"`
		}
		if s.decode(r, w, &spec) {
			for _, category := range s.Category {
				if category.Name == spec.Category.Name {
					BadRequest(w, "com.vmware.vapi.std.errors.already_exists")
					return
				}
			}
			id := newID("Category")
			spec.Category.ID = id
			s.Category[id] = &spec.Category
			OK(w, id)
		}
	case http.MethodGet:
		var ids []string
		for id := range s.Category {
			ids = append(ids, id)
		}

		OK(w, ids)
	}
}

func (s *handler) categoryID(w http.ResponseWriter, r *http.Request) {
	id := s.id(r)

	o, ok := s.Category[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		delete(s.Category, id)
		for ix, tag := range s.Tag {
			if tag.CategoryID == id {
				delete(s.Tag, ix)
				delete(s.Association, ix)
			}
		}
		OK(w)
	case http.MethodPatch:
		var spec struct {
			Category tags.Category `json:"update_spec"`
		}
		if s.decode(r, w, &spec) {
			ntypes := len(spec.Category.AssociableTypes)
			if ntypes != 0 {
				// Validate that AssociableTypes is only appended to.
				etypes := len(o.AssociableTypes)
				fail := ntypes < etypes
				if !fail {
					fail = !reflect.DeepEqual(o.AssociableTypes, spec.Category.AssociableTypes[:etypes])
				}
				if fail {
					BadRequest(w, "com.vmware.vapi.std.errors.invalid_argument")
					return
				}
			}
			o.Patch(&spec.Category)
			OK(w)
		}
	case http.MethodGet:
		OK(w, o)
	}
}

func (s *handler) tag(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var spec struct {
			Tag tags.Tag `json:"create_spec"`
		}
		if s.decode(r, w, &spec) {
			for _, tag := range s.Tag {
				if tag.Name == spec.Tag.Name {
					BadRequest(w, "com.vmware.vapi.std.errors.already_exists")
					return
				}
			}
			id := newID("Tag")
			spec.Tag.ID = id
			s.Tag[id] = &spec.Tag
			s.Association[id] = make(map[internal.AssociatedObject]bool)
			OK(w, id)
		}
	case http.MethodGet:
		var ids []string
		for id := range s.Tag {
			ids = append(ids, id)
		}
		OK(w, ids)
	}
}

func (s *handler) tagID(w http.ResponseWriter, r *http.Request) {
	id := s.id(r)

	switch s.action(r) {
	case "list-tags-for-category":
		var ids []string
		for _, tag := range s.Tag {
			if tag.CategoryID == id {
				ids = append(ids, tag.ID)
			}
		}
		OK(w, ids)
		return
	}

	o, ok := s.Tag[id]
	if !ok {
		log.Printf("tag not found: %s", id)
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		delete(s.Tag, id)
		delete(s.Association, id)
		OK(w)
	case http.MethodPatch:
		var spec struct {
			Tag tags.Tag `json:"update_spec"`
		}
		if s.decode(r, w, &spec) {
			o.Patch(&spec.Tag)
			OK(w)
		}
	case http.MethodGet:
		OK(w, o)
	}
}

func (s *handler) association(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var spec struct {
		internal.Association
		TagIDs    []string                    `json:"tag_ids,omitempty"`
		ObjectIDs []internal.AssociatedObject `json:"object_ids,omitempty"`
	}
	if !s.decode(r, w, &spec) {
		return
	}

	switch s.action(r) {
	case "list-attached-tags":
		var ids []string
		for id, objs := range s.Association {
			if objs[*spec.ObjectID] {
				ids = append(ids, id)
			}
		}
		OK(w, ids)
	case "list-attached-objects-on-tags":
		var res []tags.AttachedObjects
		for _, id := range spec.TagIDs {
			o := tags.AttachedObjects{TagID: id}
			for i := range s.Association[id] {
				o.ObjectIDs = append(o.ObjectIDs, i)
			}
			res = append(res, o)
		}
		OK(w, res)
	case "list-attached-tags-on-objects":
		var res []tags.AttachedTags
		for _, ref := range spec.ObjectIDs {
			o := tags.AttachedTags{ObjectID: ref}
			for id, objs := range s.Association {
				if objs[ref] {
					o.TagIDs = append(o.TagIDs, id)
				}
			}
			res = append(res, o)
		}
		OK(w, res)
	}
}

func (s *handler) associationID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := s.id(r)
	if _, exists := s.Association[id]; !exists {
		log.Printf("association tag not found: %s", id)
		http.NotFound(w, r)
		return
	}

	var spec internal.Association
	if !s.decode(r, w, &spec) {
		return
	}

	switch s.action(r) {
	case "attach":
		s.Association[id][*spec.ObjectID] = true
		OK(w)
	case "detach":
		delete(s.Association[id], *spec.ObjectID)
		OK(w)
	case "list-attached-objects":
		var ids []internal.AssociatedObject
		for id := range s.Association[id] {
			ids = append(ids, id)
		}
		OK(w, ids)
	}
}

func (s *handler) library(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var spec struct {
			Library library.Library `json:"create_spec"`
			Find    library.Find    `json:"spec"`
		}
		if !s.decode(r, w, &spec) {
			return
		}

		switch s.action(r) {
		case "find":
			var ids []string
			for _, l := range s.Library {
				if spec.Find.Type != "" {
					if spec.Find.Type != l.Library.Type {
						continue
					}
				}
				if spec.Find.Name != "" {
					if !strings.EqualFold(l.Library.Name, spec.Find.Name) {
						continue
					}
				}
				ids = append(ids, l.ID)
			}
			OK(w, ids)
		case "":
			id := uuid.New().String()
			spec.Library.ID = id
			dir := libraryPath(&spec.Library, "")
			if err := os.Mkdir(dir, 0750); err != nil {
				s.error(w, err)
				return
			}
			s.Library[id] = content{
				Library: &spec.Library,
				Item:    make(map[string]*item),
			}
			OK(w, id)
		}
	case http.MethodGet:
		var ids []string
		for id := range s.Library {
			ids = append(ids, id)
		}
		OK(w, ids)
	}
}

func (s *handler) libraryID(w http.ResponseWriter, r *http.Request) {
	id := s.id(r)
	l, ok := s.Library[id]
	if !ok {
		log.Printf("library not found: %s", id)
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		p := libraryPath(l.Library, "")
		if err := os.RemoveAll(p); err != nil {
			s.error(w, err)
			return
		}
		delete(s.Library, id)
		OK(w)
	case http.MethodPatch:
		var spec struct {
			Library library.Library `json:"update_spec"`
		}
		if s.decode(r, w, &spec) {
			l.Patch(&spec.Library)
			OK(w)
		}
	case http.MethodGet:
		OK(w, l)
	}
}

func (s *handler) libraryItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var spec struct {
			Item library.Item     `json:"create_spec"`
			Find library.FindItem `json:"spec"`
		}
		if !s.decode(r, w, &spec) {
			return
		}

		switch s.action(r) {
		case "find":
			var ids []string
			for _, l := range s.Library {
				if spec.Find.LibraryID != "" {
					if spec.Find.LibraryID != l.ID {
						continue
					}
				}
				for _, i := range l.Item {
					if spec.Find.Name != "" {
						if spec.Find.Name != i.Name {
							continue
						}
					}
					if spec.Find.Type != "" {
						if spec.Find.Type != i.Type {
							continue
						}
					}
					ids = append(ids, i.ID)
				}
			}
			OK(w, ids)
		case "create", "":
			id := spec.Item.LibraryID
			l, ok := s.Library[id]
			if !ok {
				log.Printf("library not found: %s", id)
				http.NotFound(w, r)
				return
			}
			for _, item := range l.Item {
				if item.Name == spec.Item.Name {
					BadRequest(w, "com.vmware.vapi.std.errors.already_exists")
					return
				}
			}
			id = uuid.New().String()
			spec.Item.ID = id
			l.Item[id] = &item{Item: &spec.Item}
			OK(w, id)
		}
	case http.MethodGet:
		id := r.URL.Query().Get("library_id")
		l, ok := s.Library[id]
		if !ok {
			log.Printf("library not found: %s", id)
			http.NotFound(w, r)
			return
		}

		var ids []string
		for id := range l.Item {
			ids = append(ids, id)
		}
		OK(w, ids)
	}
}

func (s *handler) libraryItemID(w http.ResponseWriter, r *http.Request) {
	id := s.id(r)
	lid := r.URL.Query().Get("library_id")
	if lid == "" {
		if l := s.itemLibrary(id); l != nil {
			lid = l.ID
		}
	}
	l, ok := s.Library[lid]
	if !ok {
		log.Printf("library not found: %q", lid)
		http.NotFound(w, r)
		return
	}
	item, ok := l.Item[id]
	if !ok {
		log.Printf("library item not found: %q", id)
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		p := libraryPath(l.Library, id)
		if err := os.RemoveAll(p); err != nil {
			s.error(w, err)
			return
		}
		delete(l.Item, item.ID)
		OK(w)
	case http.MethodPatch:
		var spec struct {
			Item library.Item `json:"update_spec"`
		}
		if s.decode(r, w, &spec) {
			item.Patch(&spec.Item)
			OK(w)
		}
	case http.MethodGet:
		OK(w, item)
	}
}

func (s *handler) libraryItemUpdateSession(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var ids []string
		for id := range s.Update {
			ids = append(ids, id)
		}
		OK(w, ids)
	case http.MethodPost:
		var spec struct {
			Session library.Session `json:"create_spec"`
		}
		if !s.decode(r, w, &spec) {
			return
		}

		switch s.action(r) {
		case "create", "":
			lib := s.itemLibrary(spec.Session.LibraryItemID)
			if lib == nil {
				log.Printf("library for item %q not found", spec.Session.LibraryItemID)
				http.NotFound(w, r)
				return
			}
			session := &library.Session{
				ID:                        uuid.New().String(),
				LibraryItemID:             spec.Session.LibraryItemID,
				LibraryItemContentVersion: "1",
				ClientProgress:            0,
				State:                     "ACTIVE",
				ExpirationTime:            types.NewTime(time.Now().Add(time.Hour)),
			}
			s.Update[session.ID] = update{
				Session: session,
				Library: lib,
				File:    make(map[string]*library.UpdateFile),
			}
			OK(w, session.ID)
		}
	}
}

func (s *handler) libraryItemUpdateSessionID(w http.ResponseWriter, r *http.Request) {
	id := s.id(r)
	up, ok := s.Update[id]
	if !ok {
		log.Printf("update session not found: %s", id)
		http.NotFound(w, r)
		return
	}

	session := up.Session
	done := func(state string) {
		up.State = state
		go time.AfterFunc(session.ExpirationTime.Sub(time.Now()), func() {
			s.Lock()
			delete(s.Update, id)
			s.Unlock()
		})
	}

	switch r.Method {
	case http.MethodGet:
		OK(w, session)
	case http.MethodPost:
		switch s.action(r) {
		case "cancel":
			done("CANCELED")
		case "complete":
			done("DONE")
		case "fail":
			done("ERROR")
		case "keep-alive":
			session.ExpirationTime = types.NewTime(time.Now().Add(time.Hour))
		}
		OK(w)
	case http.MethodDelete:
		delete(s.Update, id)
		OK(w)
	}
}

func (s *handler) libraryItemUpdateSessionFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("update_session_id")
	up, ok := s.Update[id]
	if !ok {
		log.Printf("update session not found: %s", id)
		http.NotFound(w, r)
		return
	}

	var files []*library.UpdateFile
	for _, f := range up.File {
		files = append(files, f)
	}
	OK(w, files)
}

func (s *handler) pullSource(up update, info *library.UpdateFile) {
	done := func(err error) {
		s.Lock()
		info.Status = "READY"
		if err != nil {
			log.Printf("PULL %s: %s", info.SourceEndpoint.URI, err)
			info.Status = "ERROR"
			up.State = "ERROR"
			up.ErrorMessage = &rest.LocalizableMessage{DefaultMessage: err.Error()}
		}
		s.Unlock()
	}

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	res, err := c.Get(info.SourceEndpoint.URI)
	if err != nil {
		done(err)
		return
	}

	err = s.libraryItemFileCreate(&up, info.Name, res.Body)
	done(err)
}

func (s *handler) libraryItemUpdateSessionFileID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := s.id(r)
	up, ok := s.Update[id]
	if !ok {
		log.Printf("update session not found: %s", id)
		http.NotFound(w, r)
		return
	}

	switch s.action(r) {
	case "add":
		var spec struct {
			File library.UpdateFile `json:"file_spec"`
		}
		if s.decode(r, w, &spec) {
			id = uuid.New().String()
			info := &library.UpdateFile{
				Name:             spec.File.Name,
				SourceType:       spec.File.SourceType,
				Status:           "WAITING_FOR_TRANSFER",
				BytesTransferred: 0,
			}
			switch info.SourceType {
			case "PUSH":
				u := url.URL{
					Scheme: s.URL.Scheme,
					Host:   s.URL.Host,
					Path:   path.Join(rest.Path, internal.LibraryItemFileData, id, info.Name),
				}
				info.UploadEndpoint = &library.TransferEndpoint{URI: u.String()}
			case "PULL":
				info.SourceEndpoint = spec.File.SourceEndpoint
				go s.pullSource(up, info)
			}
			up.File[id] = info
			OK(w, info)
		}
	case "get":
		OK(w, up.Session)
	case "list":
		var ids []string
		for id := range up.File {
			ids = append(ids, id)
		}
		OK(w, ids)
	case "remove":
		delete(s.Update, id)
		OK(w)
	case "validate":
		// TODO
	}
}

func (s *handler) libraryItemDownloadSession(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var ids []string
		for id := range s.Download {
			ids = append(ids, id)
		}
		OK(w, ids)
	case http.MethodPost:
		var spec struct {
			Session library.Session `json:"create_spec"`
		}
		if !s.decode(r, w, &spec) {
			return
		}

		switch s.action(r) {
		case "create", "":
			var lib *library.Library
			var files []library.File
			for _, l := range s.Library {
				if item, ok := l.Item[spec.Session.LibraryItemID]; ok {
					lib = l.Library
					files = item.File
					break
				}
			}
			if lib == nil {
				log.Printf("library for item %q not found", spec.Session.LibraryItemID)
				http.NotFound(w, r)
				return
			}
			session := &library.Session{
				ID:                        uuid.New().String(),
				LibraryItemID:             spec.Session.LibraryItemID,
				LibraryItemContentVersion: "1",
				ClientProgress:            0,
				State:                     "ACTIVE",
				ExpirationTime:            types.NewTime(time.Now().Add(time.Hour)),
			}
			s.Download[session.ID] = download{
				Session: session,
				Library: lib,
				File:    make(map[string]*library.DownloadFile),
			}
			for _, file := range files {
				s.Download[session.ID].File[file.Name] = &library.DownloadFile{
					Name:   file.Name,
					Status: "UNPREPARED",
				}
			}
			OK(w, session.ID)
		}
	}
}

func (s *handler) libraryItemDownloadSessionID(w http.ResponseWriter, r *http.Request) {
	id := s.id(r)
	up, ok := s.Download[id]
	if !ok {
		log.Printf("download session not found: %s", id)
		http.NotFound(w, r)
		return
	}

	session := up.Session
	switch r.Method {
	case http.MethodGet:
		OK(w, session)
	case http.MethodPost:
		switch s.action(r) {
		case "cancel", "complete", "fail":
			delete(s.Download, id) // TODO: fully mock VC's behavior
		case "keep-alive":
			session.ExpirationTime = types.NewTime(time.Now().Add(time.Hour))
		}
		OK(w)
	case http.MethodDelete:
		delete(s.Download, id)
		OK(w)
	}
}

func (s *handler) libraryItemDownloadSessionFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("download_session_id")
	dl, ok := s.Download[id]
	if !ok {
		log.Printf("download session not found: %s", id)
		http.NotFound(w, r)
		return
	}

	var files []*library.DownloadFile
	for _, f := range dl.File {
		files = append(files, f)
	}
	OK(w, files)
}

func (s *handler) libraryItemDownloadSessionFileID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := s.id(r)
	dl, ok := s.Download[id]
	if !ok {
		log.Printf("download session not found: %s", id)
		http.NotFound(w, r)
		return
	}

	var spec struct {
		File string `json:"file_name"`
	}

	switch s.action(r) {
	case "prepare":
		if s.decode(r, w, &spec) {
			u := url.URL{
				Scheme: s.URL.Scheme,
				Host:   s.URL.Host,
				Path:   path.Join(rest.Path, internal.LibraryItemFileData, id, spec.File),
			}
			info := &library.DownloadFile{
				Name:             spec.File,
				Status:           "PREPARED",
				BytesTransferred: 0,
				DownloadEndpoint: &library.TransferEndpoint{
					URI: u.String(),
				},
			}
			dl.File[spec.File] = info
			OK(w, info)
		}
	case "get":
		if s.decode(r, w, &spec) {
			OK(w, dl.File[spec.File])
		}
	}
}

func (s *handler) itemLibrary(id string) *library.Library {
	for _, l := range s.Library {
		if _, ok := l.Item[id]; ok {
			return l.Library
		}
	}
	return nil
}

func (s *handler) updateFileInfo(id string) *update {
	for _, up := range s.Update {
		for i := range up.File {
			if i == id {
				return &up
			}
		}
	}
	return nil
}

// libraryPath returns the local Datastore fs path for a Library or Item if id is specified.
func libraryPath(l *library.Library, id string) string {
	// DatastoreID (moref) format is "$local-path@$ds-folder-id",
	// see simulator.HostDatastoreSystem.CreateLocalDatastore
	ds := strings.SplitN(l.Storage[0].DatastoreID, "@", 2)[0]
	return path.Join(append([]string{ds, "contentlib-" + l.ID}, id)...)
}

func (s *handler) libraryItemFileCreate(up *update, name string, body io.ReadCloser) error {
	var in io.Reader = body
	dir := libraryPath(up.Library, up.Session.LibraryItemID)
	if err := os.MkdirAll(dir, 0750); err != nil {
		return err
	}

	if path.Ext(name) == ".ova" {
		// All we need is the .ovf, vcsim has no use for .vmdk or .mf
		r := tar.NewReader(body)
		for {
			h, err := r.Next()
			if err != nil {
				return err
			}

			if path.Ext(h.Name) == ".ovf" {
				name = h.Name
				in = io.LimitReader(body, h.Size)
				break
			}
		}
	}

	file, err := os.Create(path.Join(dir, name))
	if err != nil {
		return err
	}

	n, err := io.Copy(file, in)
	_ = body.Close()
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	i := s.Library[up.Library.ID].Item[up.Session.LibraryItemID]
	i.File = append(i.File, library.File{
		Cached:  types.NewBool(true),
		Name:    name,
		Size:    types.NewInt64(n),
		Version: "1",
	})

	return nil
}

func (s *handler) libraryItemFileData(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	id, name := p[len(p)-2], p[len(p)-1]

	if r.Method == http.MethodGet {
		dl, ok := s.Download[id]
		if !ok {
			log.Printf("library download not found: %s", id)
			http.NotFound(w, r)
			return
		}
		p := path.Join(libraryPath(dl.Library, dl.Session.LibraryItemID), name)
		f, err := os.Open(p)
		if err != nil {
			s.error(w, err)
			return
		}
		_, err = io.Copy(w, f)
		if err != nil {
			log.Printf("copy %s: %s", p, err)
		}
		_ = f.Close()
		return
	}

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	up := s.updateFileInfo(id)
	if up == nil {
		log.Printf("library update not found: %s", id)
		http.NotFound(w, r)
		return
	}

	err := s.libraryItemFileCreate(up, name, r.Body)
	if err != nil {
		s.error(w, err)
	}
}

func (s *handler) libraryItemFile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("library_item_id")
	for _, l := range s.Library {
		if i, ok := l.Item[id]; ok {
			OK(w, i.File)
			return
		}
	}
	http.NotFound(w, r)
}

func (s *handler) libraryItemFileID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := s.id(r)
	var spec struct {
		Name string `json:"name"`
	}
	if !s.decode(r, w, &spec) {
		return
	}
	for _, l := range s.Library {
		if i, ok := l.Item[id]; ok {
			for _, f := range i.File {
				if f.Name == spec.Name {
					OK(w, f)
					return
				}
			}
		}
	}
	http.NotFound(w, r)
}

func (i *item) ovf() string {
	for _, f := range i.File {
		if strings.HasSuffix(f.Name, ".ovf") {
			return f.Name
		}
	}
	return ""
}

func (s *handler) libraryDeploy(lib *library.Library, item *item, deploy vcenter.Deploy) (*nfc.LeaseInfo, error) {
	name := item.ovf()
	desc, err := ioutil.ReadFile(filepath.Join(libraryPath(lib, item.ID), name))
	if err != nil {
		return nil, err
	}
	ds := types.ManagedObjectReference{Type: "Datastore", Value: deploy.DeploymentSpec.DefaultDatastoreID}
	pool := types.ManagedObjectReference{Type: "ResourcePool", Value: deploy.Target.ResourcePoolID}
	var folder, host *types.ManagedObjectReference
	if deploy.Target.FolderID != "" {
		folder = &types.ManagedObjectReference{Type: "Folder", Value: deploy.Target.FolderID}
	}
	if deploy.Target.HostID != "" {
		host = &types.ManagedObjectReference{Type: "HostSystem", Value: deploy.Target.HostID}
	}

	ctx := context.Background()
	c, err := govmomi.NewClient(ctx, &s.URL, true)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = c.Logout(ctx)
	}()

	v, err := view.NewManager(c.Client).CreateContainerView(ctx, c.ServiceContent.RootFolder, nil, true)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = v.Destroy(ctx)
	}()
	refs, err := v.Find(ctx, []string{"Network"}, nil)
	if err != nil {
		return nil, err
	}

	var network []types.OvfNetworkMapping
	for _, net := range deploy.NetworkMappings {
		for i := range refs {
			if refs[i].Value == net.Value {
				network = append(network, types.OvfNetworkMapping{Name: net.Key, Network: refs[i]})
				break
			}
		}
	}

	if ds.Value == "" {
		// Datastore is optional in the deploy spec, but not in OvfManager.CreateImportSpec
		refs, err = v.Find(ctx, []string{"Datastore"}, nil)
		if err != nil {
			return nil, err
		}
		// TODO: consider StorageProfileID
		ds = refs[0]
	}

	cisp := types.OvfCreateImportSpecParams{
		DiskProvisioning: deploy.DeploymentSpec.StorageProvisioning,
		EntityName:       deploy.DeploymentSpec.Name,
		NetworkMapping:   network,
	}

	for _, p := range deploy.AdditionalParams {
		switch p.Type {
		case vcenter.TypePropertyParams:
			for _, prop := range p.Properties {
				cisp.PropertyMapping = append(cisp.PropertyMapping, types.KeyValue{
					Key:   prop.ID,
					Value: prop.Value,
				})
			}
		case vcenter.TypeDeploymentOptionParams:
			cisp.OvfManagerCommonParams.DeploymentOption = p.SelectedKey
		}
	}

	m := ovf.NewManager(c.Client)
	spec, err := m.CreateImportSpec(ctx, string(desc), pool, ds, cisp)
	if err != nil {
		return nil, err
	}
	if spec.Error != nil {
		return nil, errors.New(spec.Error[0].LocalizedMessage)
	}

	req := types.ImportVApp{
		This:   pool,
		Spec:   spec.ImportSpec,
		Folder: folder,
		Host:   host,
	}
	res, err := methods.ImportVApp(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	lease := nfc.NewLease(c.Client, res.Returnval)
	info, err := lease.Wait(ctx, spec.FileItem)
	if err != nil {
		return nil, err
	}

	return info, lease.Complete(ctx)
}

func (s *handler) libraryItemDeployID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := s.id(r)
	ok := false
	var lib *library.Library
	var item *item
	for _, l := range s.Library {
		item, ok = l.Item[id]
		if ok {
			lib = l.Library
			break
		}
	}
	if !ok {
		log.Printf("library item not found: %q", id)
		http.NotFound(w, r)
		return
	}

	var spec struct {
		vcenter.Deploy
	}
	if !s.decode(r, w, &spec) {
		return
	}

	switch s.action(r) {
	case "deploy":
		var d vcenter.Deployment
		info, err := s.libraryDeploy(lib, item, spec.Deploy)
		if err == nil {
			id := vcenter.ResourceID(info.Entity)
			d.Succeeded = true
			d.ResourceID = &id
		} else {
			d.Error = &vcenter.DeploymentError{
				Errors: []vcenter.OVFError{{
					Category: "SERVER",
					Error: &vcenter.Error{
						Class: "com.vmware.vapi.std.errors.error",
						Messages: []rest.LocalizableMessage{
							{
								DefaultMessage: err.Error(),
							},
						},
					},
				}},
			}
		}
		OK(w, d)
	case "filter":
		res := vcenter.FilterResponse{
			Name: item.Name,
		}
		OK(w, res)
	default:
		http.NotFound(w, r)
	}
}
