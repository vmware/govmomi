// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"archive/tar"
	"bytes"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/nfc"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/ovf"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi"
	"github.com/vmware/govmomi/vapi/internal"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vapi/tags"
	"github.com/vmware/govmomi/vapi/vcenter"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	vim "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
	"github.com/vmware/govmomi/vmdk"
)

type item struct {
	*library.Item
	File     []library.File
	Template *types.ManagedObjectReference
}

type content struct {
	*library.Library
	Item map[string]*item
	Subs map[string]*library.Subscriber
	VMTX map[string]*types.ManagedObjectReference
}

type update struct {
	*sync.WaitGroup
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
	Map         *simulator.Registry
	ServeMux    *http.ServeMux
	URL         url.URL
	Category    map[string]*tags.Category
	Tag         map[string]*tags.Tag
	Association map[string]map[internal.AssociatedObject]bool
	Session     map[string]*rest.Session
	Library     map[string]*content
	Update      map[string]update
	Download    map[string]download
	Policies    []library.ContentSecurityPoliciesInfo
	Trust       map[string]library.TrustedCertificate
}

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		if r.IsVPX() {
			patterns, h := New(s.Listen, r)
			for _, p := range patterns {
				s.Handle(p, h)
			}
		}
	})
}

// New creates a vAPI simulator.
func New(u *url.URL, r *simulator.Registry) ([]string, http.Handler) {
	s := &handler{
		Map:         r,
		ServeMux:    http.NewServeMux(),
		URL:         *u,
		Category:    make(map[string]*tags.Category),
		Tag:         make(map[string]*tags.Tag),
		Association: make(map[string]map[internal.AssociatedObject]bool),
		Session:     make(map[string]*rest.Session),
		Library:     make(map[string]*content),
		Update:      make(map[string]update),
		Download:    make(map[string]download),
		Policies:    defaultSecurityPolicies(),
		Trust:       make(map[string]library.TrustedCertificate),
	}

	handlers := []struct {
		p string
		m http.HandlerFunc
	}{
		// /rest/ patterns.
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
		{internal.SubscribedLibraryPath + "/", s.libraryID},
		{internal.Subscriptions, s.subscriptions},
		{internal.Subscriptions + "/", s.subscriptionsID},
		{internal.LibraryItemPath, s.libraryItem},
		{internal.LibraryItemPath + "/", s.libraryItemID},
		{internal.LibraryItemStoragePath, s.libraryItemStorage},
		{internal.LibraryItemStoragePath + "/", s.libraryItemStorageID},
		{internal.SubscribedLibraryItem + "/", s.libraryItemID},
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
		{internal.VCenterOVFLibraryItem, s.libraryItemOVF},
		{internal.VCenterOVFLibraryItem + "/", s.libraryItemOVFID},
		{internal.VCenterVMTXLibraryItem, s.libraryItemCreateTemplate},
		{internal.VCenterVMTXLibraryItem + "/", s.libraryItemTemplateID},
		{internal.DebugEcho, s.debugEcho},
		// /api/ patterns.
		{internal.SecurityPoliciesPath, s.librarySecurityPolicies},
		{internal.TrustedCertificatesPath, s.libraryTrustedCertificates},
		{internal.TrustedCertificatesPath + "/", s.libraryTrustedCertificatesID},
	}

	for i := range handlers {
		h := handlers[i]
		s.HandleFunc(h.p, h.m)
	}

	return []string{rest.Path + "/", vapi.Path + "/"}, s
}

func (s *handler) withClient(f func(context.Context, *vim25.Client) error) error {
	return WithClient(s.URL, f)
}

// WithClient creates invokes f with an authenticated vim25.Client.
func WithClient(u url.URL, f func(context.Context, *vim25.Client) error) error {
	ctx := context.Background()
	c, err := govmomi.NewClient(ctx, &u, true)
	if err != nil {
		return err
	}
	defer func() {
		_ = c.Logout(ctx)
	}()
	return f(ctx, c.Client)
}

// RunTask creates a Task with the given spec and sets the task state based on error returned by f.
func RunTask(u url.URL, spec types.CreateTask, f func(context.Context, *vim25.Client) error) string {
	var id string

	err := WithClient(u, func(ctx context.Context, c *vim25.Client) error {
		spec.This = *c.ServiceContent.TaskManager
		if spec.TaskTypeId == "" {
			spec.TaskTypeId = "com.vmware.govmomi.simulator.test"
		}
		res, err := methods.CreateTask(ctx, c, &spec)
		if err != nil {
			return err
		}

		ref := res.Returnval.Task
		task := object.NewTask(c, ref)
		id = ref.Value + ":" + uuid.NewString()

		if err = task.SetState(ctx, types.TaskInfoStateRunning, nil, nil); err != nil {
			return err
		}

		var fault *types.LocalizedMethodFault
		state := types.TaskInfoStateSuccess
		if f != nil {
			err = f(ctx, c)
		}

		if err != nil {
			fault = &types.LocalizedMethodFault{
				Fault:            &types.SystemError{Reason: err.Error()},
				LocalizedMessage: err.Error(),
			}
			state = types.TaskInfoStateError
		}

		return task.SetState(ctx, state, nil, fault)
	})

	if err != nil {
		panic(err) // should not happen
	}

	return id
}

// HandleFunc wraps the given handler with authorization checks and passes to http.ServeMux.HandleFunc
func (s *handler) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	// Rest paths have been moved from /rest/* to /api/*. Account for both the legacy and new cases here.
	if !strings.HasPrefix(pattern, rest.Path) && !strings.HasPrefix(pattern, vapi.Path) {
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
		return u, s.Map.SessionManager().Authenticate(s.URL, &vim.Login{UserName: u, Password: p})
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
		ids = append(
			ids,
			vim.ManagedObjectReference{
				Type:  id.Type,
				Value: id.Value,
			})
	}
	return ids, nil
}

// AttachedTags is meant for internal use via simulator.Registry.tagManager
func (s *handler) AttachedTags(ref vim.ManagedObjectReference) ([]vim.VslmTagEntry, vim.BaseMethodFault) {
	oid := internal.AssociatedObject{
		Type:  ref.Type,
		Value: ref.Value,
	}
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
	s.Association[t.ID][internal.AssociatedObject{
		Type:  ref.Type,
		Value: ref.Value,
	}] = true
	return nil
}

// DetachTag is meant for internal use via simulator.Registry.tagManager
func (s *handler) DetachTag(id vim.ManagedObjectReference, tag vim.VslmTagEntry) vim.BaseMethodFault {
	t := s.findTag(tag)
	if t == nil {
		return new(vim.NotFound)
	}
	delete(s.Association[t.ID], internal.AssociatedObject{
		Type:  id.Type,
		Value: id.Value,
	})
	return nil
}

// StatusOK responds with http.StatusOK and encodes val, if specified, to JSON
// For use with "/api" endpoints.
func StatusOK(w http.ResponseWriter, val ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if len(val) == 0 {
		return
	}

	err := json.NewEncoder(w).Encode(val[0])

	if err != nil {
		log.Panic(err)
	}
}

// OK responds with http.StatusOK and encodes val, if specified, to JSON
// For use with "/rest" endpoints where the response is a "value" wrapped structure.
func OK(w http.ResponseWriter, val ...interface{}) {
	if len(val) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}

	s := struct {
		Value interface{} `json:"value,omitempty"`
	}{
		val[0],
	}

	StatusOK(w, s)
}

// BadRequest responds with http.StatusBadRequest and json encoded vAPI error of type kind.
// For use with "/rest" endpoints where the response is a "value" wrapped structure.
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

// ApiErrorAlreadyExists responds with a REST error of type "ALREADY_EXISTS".
// For use with "/api" endpoints.
func ApiErrorAlreadyExists(w http.ResponseWriter) {
	apiError(w, http.StatusBadRequest, "ALREADY_EXISTS")
}

// ApiErrorGeneral responds with a REST error of type "ERROR".
// For use with "/api" endpoints.
func ApiErrorGeneral(w http.ResponseWriter) {
	apiError(w, http.StatusInternalServerError, "ERROR")
}

// ApiErrorInvalidArgument responds with a REST error of type "INVALID_ARGUMENT".
// For use with "/api" endpoints.
func ApiErrorInvalidArgument(w http.ResponseWriter) {
	apiError(w, http.StatusBadRequest, "INVALID_ARGUMENT")
}

// ApiErrorNotAllowedInCurrentState responds with a REST error of type "NOT_ALLOWED_IN_CURRENT_STATE".
// For use with "/api" endpoints.
func ApiErrorNotAllowedInCurrentState(w http.ResponseWriter) {
	apiError(w, http.StatusBadRequest, "NOT_ALLOWED_IN_CURRENT_STATE")
}

// ApiErrorNotFound responds with a REST error of type "NOT_FOUND".
// For use with "/api" endpoints.
func ApiErrorNotFound(w http.ResponseWriter) {
	apiError(w, http.StatusNotFound, "NOT_FOUND")
}

// ApiErrorResourceInUse responds with a REST error of type "RESOURCE_IN_USE".
// For use with "/api" endpoints.
func ApiErrorResourceInUse(w http.ResponseWriter) {
	apiError(w, http.StatusBadRequest, "RESOURCE_IN_USE")
}

// ApiErrorUnauthorized responds with a REST error of type "UNAUTHORIZED".
// For use with "/api" endpoints.
func ApiErrorUnauthorized(w http.ResponseWriter) {
	apiError(w, http.StatusBadRequest, "UNAUTHORIZED")
}

// ApiErrorUnsupported responds with a REST error of type "UNSUPPORTED".
// For use with "/api" endpoints.
func ApiErrorUnsupported(w http.ResponseWriter) {
	apiError(w, http.StatusBadRequest, "UNSUPPORTED")
}

func apiError(w http.ResponseWriter, statusCode int, errorType string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(`{"error_type":"%s", "messages":[]}`, errorType)))
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

	// Use ServeHTTP directly and not via handler otherwise the path values like "{id}" are not set
	s.ServeMux.ServeHTTP(w, r)
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

func (s *handler) expiredSession(id string, now time.Time) bool {
	expired := true
	s.Lock()
	session, ok := s.Session[id]
	if ok {
		expired = now.Sub(session.LastAccessed) > simulator.SessionIdleTimeout
		if expired {
			delete(s.Session, id)
		}
	}
	s.Unlock()
	return expired
}

func (s *handler) session(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get(internal.SessionCookieName)
	useHeaderAuthn := strings.ToLower(r.Header.Get(internal.UseHeaderAuthn))

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
		simulator.SessionIdleWatch(context.Background(), id, s.expiredSession)
		if useHeaderAuthn != "true" {
			http.SetCookie(w, &http.Cookie{
				Name:  internal.SessionCookieName,
				Value: id,
				Path:  rest.Path,
			})
		}
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
				if tag.Name == spec.Tag.Name && tag.CategoryID == spec.Tag.CategoryID {
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

// TODO: support cardinality checks
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

	case "attach-multiple-tags-to-object":
		// TODO: add check if target (moref) exist or return 403 as per API behavior

		res := struct {
			Success bool             `json:"success"`
			Errors  tags.BatchErrors `json:"error_messages,omitempty"`
		}{}

		for _, id := range spec.TagIDs {
			if _, exists := s.Association[id]; !exists {
				log.Printf("association tag not found: %s", id)
				res.Errors = append(res.Errors, tags.BatchError{
					Type:    "cis.tagging.objectNotFound.error",
					Message: fmt.Sprintf("Tagging object %s not found", id),
				})
			} else {
				s.Association[id][*spec.ObjectID] = true
			}
		}

		if len(res.Errors) == 0 {
			res.Success = true
		}
		OK(w, res)

	case "detach-multiple-tags-from-object":
		// TODO: add check if target (moref) exist or return 403 as per API behavior

		res := struct {
			Success bool             `json:"success"`
			Errors  tags.BatchErrors `json:"error_messages,omitempty"`
		}{}

		for _, id := range spec.TagIDs {
			if _, exists := s.Association[id]; !exists {
				log.Printf("association tag not found: %s", id)
				res.Errors = append(res.Errors, tags.BatchError{
					Type:    "cis.tagging.objectNotFound.error",
					Message: fmt.Sprintf("Tagging object %s not found", id),
				})
			} else {
				s.Association[id][*spec.ObjectID] = false
			}
		}

		if len(res.Errors) == 0 {
			res.Success = true
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
	var specs struct {
		ObjectIDs []internal.AssociatedObject `json:"object_ids"`
	}
	switch s.action(r) {
	case "attach", "detach", "list-attached-objects":
		if !s.decode(r, w, &spec) {
			return
		}
	case "attach-tag-to-multiple-objects":
		if !s.decode(r, w, &specs) {
			return
		}
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
	case "attach-tag-to-multiple-objects":
		for _, obj := range specs.ObjectIDs {
			s.Association[id][obj] = true
		}
		OK(w)
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
			if !s.isValidSecurityPolicy(spec.Library.SecurityPolicyID) {
				http.NotFound(w, r)
				return
			}

			id := uuid.New().String()
			spec.Library.ID = id
			spec.Library.ServerGUID = uuid.New().String()
			spec.Library.CreationTime = types.NewTime(time.Now())
			spec.Library.LastModifiedTime = types.NewTime(time.Now())
			spec.Library.UnsetSecurityPolicyID = spec.Library.SecurityPolicyID == ""
			dir := s.libraryPath(&spec.Library, "")
			if err := os.Mkdir(dir, 0750); err != nil {
				s.error(w, err)
				return
			}
			s.Library[id] = &content{
				Library: &spec.Library,
				Item:    make(map[string]*item),
				Subs:    make(map[string]*library.Subscriber),
				VMTX:    make(map[string]*types.ManagedObjectReference),
			}

			pub := spec.Library.Publication
			if pub != nil && pub.Published != nil && *pub.Published {
				// Generate PublishURL as real vCenter does
				pub.PublishURL = (&url.URL{
					Scheme: s.URL.Scheme,
					Host:   s.URL.Host,
					Path:   "/cls/vcsp/lib/" + id,
				}).String()
			}

			s.syncSubLib(s.Library[id])

			spec.Library.StateInfo = &library.StateInfo{State: "ACTIVE"}

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

func (s *handler) syncSubLib(dstLib *content) error {

	sub := dstLib.Subscription
	if sub == nil {
		return nil
	}

	lastSyncTime := time.Now().UTC()
	dstLib.LastSyncTime = &lastSyncTime

	var syncAll bool
	if sub.OnDemand != nil && !*sub.OnDemand {
		syncAll = true
	}

	srcLib, ok := s.Library[path.Base(sub.SubscriptionURL)]
	if !ok {
		return nil
	}

	if dstLib.Item == nil {
		dstLib.Item = map[string]*item{}
	}

	// handledSrcItems tracks which items from the source library have been
	// seen when iterating over the existing, subscribed library. This enables
	// the addition of *new* items from the source library that do not yet exist
	// in the subscribed, destination library.
	handledSrcItems := map[string]struct{}{}

	// Update any items that already exist in the subscribed library.
	for _, dstItem := range dstLib.Item {

		// Indicate this source item has been seen.
		handledSrcItems[dstItem.SourceID] = struct{}{}

		// Synchronize the item.
		if err := s.syncItem(
			dstItem,
			dstLib,
			srcLib,
			syncAll,
			srcLib.LastSyncTime); err != nil {

			return err
		}
	}

	// Add any new items from the published library.
	for _, srcItem := range srcLib.Item {

		// Skip any source items that were handled above.
		if _, ok := handledSrcItems[srcItem.ID]; ok {
			continue
		}

		now := time.Now().UTC()

		// Create the destination item.
		dstItem := &item{
			Item: &library.Item{
				// Give the copy a unique ID.
				ID: uuid.NewString(),

				// Track the source item's ID.
				SourceID: srcItem.ID,

				// Track the library to which the new item belongs.
				LibraryID: dstLib.ID,

				// Ensure the creation/modified times are set.
				CreationTime:     &now,
				LastModifiedTime: &now,
			},
		}

		// Add the new item to the subscribed library.
		dstLib.Item[dstItem.ID] = dstItem

		// Synchronize the item.
		if err := s.syncItem(
			dstItem,
			dstLib,
			srcLib,
			syncAll,
			dstLib.LastSyncTime); err != nil {

			return err
		}
	}

	return nil
}

func (s *handler) evictLibrary(lib *content) {
	for i := range lib.Item {
		s.evictItem(lib.Item[i])
	}
}

func (s *handler) evictItem(item *item) {
	item.Cached = false
	for i := range item.File {
		item.File[i].Cached = &item.Cached
	}
}

var ovfOrManifestRx = regexp.MustCompile(`(?i)^.+\.(ovf|mf)$`)

func (s *handler) syncItem(
	dstItem *item,
	dstLib,
	srcLib *content,
	syncAll bool,
	lastSyncTime *time.Time) error {

	// dstLib is nil when this function is called by the workflow for deploying
	// a subscribed library item.
	if dstLib == nil {
		var ok bool
		if dstLib, ok = s.Library[dstItem.LibraryID]; !ok {
			return fmt.Errorf("cannot find sub library id %q", dstItem.LibraryID)
		}
	}

	// srcLib is nil when this function is used to synchronize an individual
	// item versus synchronizing the entire library.
	if srcLib == nil {
		sub := dstLib.Subscription
		if sub == nil {
			return nil
		}
		var ok bool
		srcLibID := path.Base(sub.SubscriptionURL)
		if srcLib, ok = s.Library[srcLibID]; !ok {
			return fmt.Errorf("cannot find pub library id %q", srcLibID)
		}
	}

	// Get the path to the destination library item on the local filesystem.
	dstItemPath := s.libraryPath(dstLib.Library, dstItem.ID)

	// Get the source item.
	srcItem, ok := srcLib.Item[dstItem.SourceID]
	if !ok {
		// The source item is no more, so delete the destination item.
		delete(dstLib.Item, dstItem.ID)

		// Clean up the destination item's files as well.
		os.RemoveAll(dstItemPath)

		return nil
	}

	// lastSyncTime is nil when this function is used to synchronize an
	// individual item versus synchronizing the entire library.
	if lastSyncTime == nil {
		now := time.Now().UTC()
		lastSyncTime = &now
	}
	dstItem.LastSyncTime = lastSyncTime

	// There is nothing to sync if the metadata and content versions have not
	// changed, the item is already cached, and syncAll is false.
	if dstItem.MetadataVersion == srcItem.MetadataVersion &&
		dstItem.ContentVersion == srcItem.ContentVersion &&
		dstItem.Cached && !syncAll {

		return nil
	}

	// Since there was a modification, update the last mod time.
	dstItem.LastModifiedTime = lastSyncTime

	// Copy information from the srcItem to dstItem.
	dstItem.Name = srcItem.Name
	dstItem.ContentVersion = srcItem.ContentVersion
	dstItem.MetadataVersion = srcItem.MetadataVersion
	dstItem.Type = srcItem.Type
	dstItem.Description = srcItem.Description
	dstItem.Version = srcItem.Version

	// Update the destination item's files from the source.
	dstItem.File = make([]library.File, len(srcItem.File))
	copy(dstItem.File, srcItem.File)

	// If the destination item was previously cached or syncAll was used, then
	// mark the destination item as cached.
	dstItem.Cached = dstItem.Cached || syncAll
	fileIsCached := true
	fileIsNotCached := false
	fileZeroSize := int64(0)

	// Ensure a directory exists on the local filesystem for the destination
	// item.
	if err := os.MkdirAll(dstItemPath, 0750); err != nil {
		return fmt.Errorf(
			"failed to make directory for library %q item %q: %w",
			dstLib.ID,
			dstItem.ID,
			err)
	}

	// Update the the destination item's files.
	srcItemPath := s.libraryPath(srcLib.Library, srcItem.ID)
	for i := range dstItem.File {
		var (
			dstFile = &dstItem.File[i]
			srcFile = srcItem.File[i]
		)

		if !isValidFileName(dstFile.Name) || !isValidFileName(srcFile.Name) {
			return errors.New("invalid file name")
		}

		var (
			dstFilePath = path.Join(dstItemPath, dstFile.Name)
			srcFilePath = path.Join(srcItemPath, srcFile.Name)
		)

		// .ovf and .mf files are always cached.
		if ovfOrManifestRx.MatchString(dstFile.Name) {
			dstFile.Cached = &fileIsCached
			if err := copyFile(dstFilePath, srcFilePath); err != nil {
				return err
			}
			continue
		}

		// For other file types, the behavior depends on syncAll:
		//
		// - false -- Create the destination file as a placeholder but do not
		//            mark it as cached.
		// - true  -- Copy the source file to the destination and mark it as
		//            cached.
		if !syncAll {
			if err := createFile(dstFilePath); err != nil {
				return err
			}

			// Ensure the empty file does not indicate it is cached and does not
			// report a size.
			dstFile.Cached = &fileIsNotCached
			dstFile.Size = &fileZeroSize
		} else {
			if err := copyFile(dstFilePath, srcFilePath); err != nil {
				return err
			}

			// Ensure the file reports that it is cached.
			dstFile.Cached = &fileIsCached
		}
	}

	return nil
}

const (
	createOrCopyFlags = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	createOrCopyMode  = os.FileMode(0664)
)

func createFile(dstPath string) error {
	f, err := os.OpenFile(dstPath, createOrCopyFlags, createOrCopyMode)
	if err != nil {
		return fmt.Errorf("failed to create %q: %w", dstPath, err)
	}
	return f.Close()
}

// TODO: considering using object.DatastoreFileManager.Copy here instead
func openFile(dstPath string, flag int, perm os.FileMode) (*os.File, error) {
	backing := simulator.VirtualDiskBackingFileName(dstPath)
	if backing == dstPath {
		// dstPath is not a .vmdk file
		return os.OpenFile(dstPath, flag, perm)
	}

	// Generate the descriptor file using dstPath
	extent := vmdk.Extent{Info: filepath.Base(backing)}
	desc := vmdk.NewDescriptor(extent)

	f, err := os.OpenFile(dstPath, flag, perm)
	if err != nil {
		return nil, err
	}

	if err = desc.Write(f); err != nil {
		_ = f.Close()
		return nil, err
	}

	if err = f.Close(); err != nil {
		return nil, err
	}

	// Create ${name}-flat.vmdk to store contents
	return os.OpenFile(backing, flag, perm)
}

func sourceFile(srcPath string) (*os.File, error) {
	// Open ${name}-flat.vmdk if src is a .vmdk
	srcPath = simulator.VirtualDiskBackingFileName(srcPath)
	return os.Open(srcPath)
}

func copyFile(dstPath, srcPath string) error {
	srcStat, err := os.Stat(srcPath)
	if err != nil {
		return fmt.Errorf("failed to stat %q: %w", srcPath, err)
	}

	if !srcStat.Mode().IsRegular() {
		return fmt.Errorf("%q is not a regular file", srcPath)
	}

	src, err := sourceFile(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open %q: %w", srcPath, err)
	}
	defer src.Close()

	dst, err := openFile(dstPath, createOrCopyFlags, createOrCopyMode)
	if err != nil {
		return fmt.Errorf("failed to create %q: %w", dstPath, err)
	}
	defer dst.Close()

	// Copy the file using a 1MiB buffer.
	if _, err = copyReaderToWriter(dst, dstPath, src, srcPath); err != nil {
		return err
	}

	return nil
}

// copyReaderToWriter copies the contents of src to dst using a 1MiB buffer.
func copyReaderToWriter(
	dst io.Writer, dstName string,
	src io.Reader, srcName string) (int64, error) {

	buf := make([]byte, 1 /* byte */ *1024 /* kibibyte */ *1024 /* mebibyte */)
	n, err := io.CopyBuffer(dst, src, buf)
	if err != nil {
		return 0, fmt.Errorf("failed to copy %q to %q: %w", srcName, dstName, err)
	}

	return n, nil
}

func (s *handler) publish(w http.ResponseWriter, r *http.Request, sids []internal.SubscriptionDestination, l *content, vmtx *item) bool {
	var ids []string
	if len(sids) == 0 {
		for sid := range l.Subs {
			ids = append(ids, sid)
		}
	} else {
		for _, dst := range sids {
			ids = append(ids, dst.ID)
		}
	}

	for _, sid := range ids {
		sub, ok := l.Subs[sid]
		if !ok {
			log.Printf("library subscription not found: %s", sid)
			http.NotFound(w, r)
			return false
		}

		slib := s.Library[sub.LibraryID]
		if slib.VMTX[vmtx.ID] != nil {
			return true // already cloned
		}

		ds := &vcenter.DiskStorage{Datastore: l.Library.Storage[0].DatastoreID}
		ref, err := s.cloneVM(vmtx.Template.Value, vmtx.Name, sub.Placement, ds)
		if err != nil {
			s.error(w, err)
			return false
		}

		slib.VMTX[vmtx.ID] = ref
	}

	return true
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
		p := s.libraryPath(l.Library, "")
		if err := os.RemoveAll(p); err != nil {
			s.error(w, err)
			return
		}
		for _, item := range l.Item {
			s.deleteVM(item.Template)
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
	case http.MethodPost:
		switch s.action(r) {
		case "publish":
			var spec internal.SubscriptionDestinationSpec
			if !s.decode(r, w, &spec) {
				return
			}
			for _, item := range l.Item {
				if item.Type != library.ItemTypeVMTX {
					continue
				}
				if !s.publish(w, r, spec.Subscriptions, l, item) {
					return
				}
			}
			OK(w)
		case "sync":
			if l.Type == "SUBSCRIBED" {
				l.LastSyncTime = types.NewTime(time.Now())
				if err := s.syncSubLib(l); err != nil {
					BadRequest(w, err.Error())
				} else {
					OK(w)
				}
			} else {
				http.NotFound(w, r)
			}
		case "evict":
			s.evictLibrary(l)
			OK(w)
		}
	case http.MethodGet:
		OK(w, l)
	}
}

func (s *handler) subscriptions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("library")
	l, ok := s.Library[id]
	if !ok {
		log.Printf("library not found: %s", id)
		http.NotFound(w, r)
		return
	}

	var res []library.SubscriberSummary
	for sid, slib := range l.Subs {
		res = append(res, library.SubscriberSummary{
			LibraryID:              slib.LibraryID,
			LibraryName:            slib.LibraryName,
			SubscriptionID:         sid,
			LibraryVcenterHostname: "",
		})
	}
	OK(w, res)
}

func (s *handler) subscriptionsID(w http.ResponseWriter, r *http.Request) {
	id := s.id(r)
	l, ok := s.Library[id]
	if !ok {
		log.Printf("library not found: %s", id)
		http.NotFound(w, r)
		return
	}

	switch s.action(r) {
	case "get":
		var dst internal.SubscriptionDestination
		if !s.decode(r, w, &dst) {
			return
		}

		sub, ok := l.Subs[dst.ID]
		if !ok {
			log.Printf("library subscription not found: %s", dst.ID)
			http.NotFound(w, r)
			return
		}

		OK(w, sub)
	case "delete":
		var dst internal.SubscriptionDestination
		if !s.decode(r, w, &dst) {
			return
		}

		delete(l.Subs, dst.ID)

		OK(w)
	case "create", "":
		var spec struct {
			Sub struct {
				SubscriberLibrary library.SubscriberLibrary `json:"subscribed_library"`
			} `json:"spec"`
		}
		if !s.decode(r, w, &spec) {
			return
		}

		sub := spec.Sub.SubscriberLibrary
		slib, ok := s.Library[sub.LibraryID]
		if !ok {
			log.Printf("library not found: %s", sub.LibraryID)
			http.NotFound(w, r)
			return
		}

		id := uuid.New().String()
		l.Subs[id] = &library.Subscriber{
			LibraryID:       slib.ID,
			LibraryName:     slib.Name,
			LibraryLocation: sub.Target,
			Placement:       sub.Placement,
			Vcenter:         sub.Vcenter,
		}

		OK(w, id)
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
			if l.Type == "SUBSCRIBED" {
				BadRequest(w, "com.vmware.vapi.std.errors.invalid_element_type")
				return
			}
			for _, item := range l.Item {
				if item.Name == spec.Item.Name {
					BadRequest(w, "com.vmware.vapi.std.errors.already_exists")
					return
				}
			}

			if !isValidFileName(spec.Item.Name) {
				ApiErrorInvalidArgument(w)
				return
			}

			id = uuid.New().String()
			spec.Item.ID = id
			spec.Item.CreationTime = types.NewTime(time.Now())
			spec.Item.LastModifiedTime = types.NewTime(time.Now())

			// Local items are always marked Cached=true
			spec.Item.Cached = true

			// Local items start with a ContentVersion="1"
			spec.Item.ContentVersion = getVersionString("")
			spec.Item.MetadataVersion = getVersionString("")

			if l.SecurityPolicyID != "" {
				// TODO: verify signed items
				spec.Item.SecurityCompliance = types.NewBool(false)
				spec.Item.CertificateVerification = &library.ItemCertificateVerification{
					Status: "NOT_AVAILABLE",
				}
			}
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
		log.Printf("libraryItemID: library item not found: %q", id)
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		p := s.libraryPath(l.Library, id)
		if err := os.RemoveAll(p); err != nil {
			s.error(w, err)
			return
		}
		s.deleteVM(l.Item[item.ID].Template)
		delete(l.Item, item.ID)
		OK(w)
	case http.MethodPatch:
		var spec struct {
			library.Item `json:"update_spec"`
		}
		if s.decode(r, w, &spec) {
			item.Patch(&spec.Item)
			OK(w)
		}
	case http.MethodPost:
		switch s.action(r) {
		case "copy":
			var spec struct {
				library.Item `json:"destination_create_spec"`
			}
			if !s.decode(r, w, &spec) {
				return
			}

			l, ok = s.Library[spec.LibraryID]
			if !ok {
				log.Printf("library not found: %q", spec.LibraryID)
				http.NotFound(w, r)
				return
			}
			if spec.Name == "" {
				BadRequest(w, "com.vmware.vapi.std.errors.invalid_argument")
			}

			id := uuid.New().String()
			nitem := item.cp()
			nitem.ID = id
			nitem.LibraryID = spec.LibraryID
			l.Item[id] = nitem

			OK(w, id)
		case "sync":
			if l.Type == "SUBSCRIBED" || l.Publication != nil {
				var spec internal.SubscriptionItemDestinationSpec
				if s.decode(r, w, &spec) {
					if l.Publication != nil {
						if s.publish(w, r, spec.Subscriptions, l, item) {
							OK(w)
						}
					}
					if l.Type == "SUBSCRIBED" {
						if err := s.syncItem(item, l, nil, spec.Force, nil); err != nil {
							BadRequest(w, err.Error())
						} else {
							OK(w)
						}
					}
				}
			} else {
				http.NotFound(w, r)
			}
		case "publish":
			var spec internal.SubscriptionDestinationSpec
			if s.decode(r, w, &spec) {
				if s.publish(w, r, spec.Subscriptions, l, item) {
					OK(w)
				}
			}
		case "evict":
			s.evictItem(item)
			OK(w, id)
		}
	case http.MethodGet:
		OK(w, item)
	}
}

func (s *handler) libraryItemByID(id string) (*content, *item) {
	for _, l := range s.Library {
		if item, ok := l.Item[id]; ok {
			return l, item
		}
	}

	log.Printf("library for item %q not found", id)

	return nil, nil
}

func (s *handler) libraryItemStorageByID(id string) ([]library.Storage, bool) {
	lib, item := s.libraryItemByID(id)
	if item == nil {
		return nil, false
	}

	storage := make([]library.Storage, len(item.File))

	for i, file := range item.File {
		storage[i] = library.Storage{
			StorageBacking: lib.Storage[0],
			StorageURIs: []string{
				path.Join(s.libraryPath(lib.Library, id), file.Name),
			},
			Name:    file.Name,
			Version: file.Version,
		}
		if file.Checksum != nil {
			storage[i].Checksum = *file.Checksum
		}
		if file.Size != nil {
			storage[i].Size = *file.Size
		}
		if file.Cached != nil {
			storage[i].Cached = *file.Cached
		}
	}

	return storage, true
}

func (s *handler) libraryItemStorage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("library_item_id")
	storage, ok := s.libraryItemStorageByID(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	OK(w, storage)
}

func (s *handler) libraryItemStorageID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := s.id(r)
	storage, ok := s.libraryItemStorageByID(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	var spec struct {
		Name string `json:"file_name"`
	}

	if s.decode(r, w, &spec) {
		for _, file := range storage {
			if file.Name == spec.Name {
				OK(w, []library.Storage{file})
				return
			}
		}
		http.NotFound(w, r)
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
			lib, item := s.libraryItemByID(spec.Session.LibraryItemID)
			if lib == nil {
				log.Printf("library for item %q not found", item.ID)
				http.NotFound(w, r)
				return
			}
			session := &library.Session{
				ID:                        uuid.New().String(),
				LibraryItemID:             item.ID,
				LibraryItemContentVersion: item.ContentVersion,
				ClientProgress:            0,
				State:                     "ACTIVE",
				ExpirationTime:            types.NewTime(time.Now().Add(time.Hour)),
			}
			s.Update[session.ID] = update{
				WaitGroup: new(sync.WaitGroup),
				Session:   session,
				Library:   lib.Library,
				File:      make(map[string]*library.UpdateFile),
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
		if up.State != "ERROR" {
			up.State = state
		}
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
			go func() {
				up.Wait() // wait for any PULL sources to complete
				done("DONE")
			}()
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

func (s *handler) libraryItemProbe(endpoint library.TransferEndpoint) *library.ProbeResult {
	p := &library.ProbeResult{
		Status: "SUCCESS",
	}

	result := func() *library.ProbeResult {
		for i, m := range p.ErrorMessages {
			p.ErrorMessages[i].DefaultMessage = fmt.Sprintf(m.DefaultMessage, m.Args[0])
		}
		return p
	}

	u, err := url.Parse(endpoint.URI)
	if err != nil {
		p.Status = "INVALID_URL"
		p.ErrorMessages = []rest.LocalizableMessage{{
			Args:           []string{endpoint.URI},
			ID:             "com.vmware.vdcs.cls-main.invalid_url_format",
			DefaultMessage: "Invalid URL format for %s",
		}}
		return result()
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		p.Status = "INVALID_URL"
		p.ErrorMessages = []rest.LocalizableMessage{{
			Args:           []string{endpoint.URI},
			ID:             "com.vmware.vdcs.cls-main.file_probe_unsupported_uri_scheme",
			DefaultMessage: "The specified URI %s is not supported",
		}}
		return result()
	}

	res, err := http.Head(endpoint.URI)
	if err != nil {
		id := "com.vmware.vdcs.cls-main.http_request_error"
		p.Status = "INVALID_URL"

		if soap.IsCertificateUntrusted(err) {
			var info object.HostCertificateInfo
			_ = info.FromURL(u, nil)

			id = "com.vmware.vdcs.cls-main.http_request_error_peer_not_authenticated"
			p.Status = "CERTIFICATE_ERROR"
			p.SSLThumbprint = info.ThumbprintSHA1
		}

		p.ErrorMessages = []rest.LocalizableMessage{{
			Args:           []string{err.Error()},
			ID:             id,
			DefaultMessage: "HTTP request error: %s",
		}}

		return result()
	}
	_ = res.Body.Close()

	if res.TLS != nil {
		p.SSLThumbprint = soap.ThumbprintSHA1(res.TLS.PeerCertificates[0])
	}

	return result()
}

func (s *handler) libraryItemUpdateSessionFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		switch s.action(r) {
		case "probe":
			var spec struct {
				SourceEndpoint library.TransferEndpoint `json:"source_endpoint"`
			}
			if s.decode(r, w, &spec) {
				res := s.libraryItemProbe(spec.SourceEndpoint)
				OK(w, res)
			}
		default:
			http.NotFound(w, r)
		}
		return
	case http.MethodGet:
	default:
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
	defer up.Done()
	done := func(err error) {
		s.Lock()
		info.Status = "READY"
		if err != nil {
			log.Printf("PULL %s: %s", info.SourceEndpoint.URI, err)
			info.Status = "ERROR"
			info.ErrorMessage = &rest.LocalizableMessage{DefaultMessage: err.Error()}
			up.State = info.Status
			up.ErrorMessage = info.ErrorMessage
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

	err = s.libraryItemFileCreate(&up, info.Name, res.Body, info.Checksum)
	done(err)
}

func hasChecksum(c *library.Checksum) bool {
	return c != nil && c.Checksum != ""
}

var checksum = map[string]func() hash.Hash{
	"MD5":    md5.New,
	"SHA1":   sha1.New,
	"SHA256": sha256.New,
	"SHA512": sha512.New,
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
				Checksum:         spec.File.Checksum,
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
				if hasChecksum(info.Checksum) && checksum[info.Checksum.Algorithm] == nil {
					BadRequest(w, "com.vmware.vapi.std.errors.invalid_argument")
					return
				}
				info.SourceEndpoint = spec.File.SourceEndpoint
				info.Status = "TRANSFERRING"
				up.Add(1)
				go s.pullSource(up, info)
			}
			up.File[id] = info
			OK(w, info)
		}
	case "get":
		var spec struct {
			File string `json:"file_name"`
		}
		if s.decode(r, w, &spec) {
			for _, f := range up.File {
				if f.Name == spec.File {
					OK(w, f)
					return
				}
			}
		}
	case "remove":
		if up.State != "ACTIVE" {
			s.error(w, fmt.Errorf("removeFile not allowed in state %s", up.State))
			return
		}
		delete(s.Update, id)
		OK(w)
	case "validate":
		if up.State != "ACTIVE" {
			BadRequest(w, "com.vmware.vapi.std.errors.not_allowed_in_current_state")
			return
		}
		var res library.UpdateFileValidation
		// TODO check missing_files, validate .ovf
		OK(w, res)
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
			lib, item := s.libraryItemByID(spec.Session.LibraryItemID)
			if item == nil {
				http.NotFound(w, r)
				return
			}

			session := &library.Session{
				ID:                        uuid.New().String(),
				LibraryItemID:             spec.Session.LibraryItemID,
				LibraryItemContentVersion: item.ContentVersion,
				ClientProgress:            0,
				State:                     "ACTIVE",
				ExpirationTime:            types.NewTime(time.Now().Add(time.Hour)),
			}
			s.Download[session.ID] = download{
				Session: session,
				Library: lib.Library,
				File:    make(map[string]*library.DownloadFile),
			}
			for _, file := range item.File {
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
func (s *handler) libraryPath(l *library.Library, id string) string {
	dsref := types.ManagedObjectReference{
		Type:  "Datastore",
		Value: l.Storage[0].DatastoreID,
	}
	ds := s.Map.Get(dsref).(*simulator.Datastore)

	if !isValidFileName(l.ID) || !isValidFileName(id) {
		panic("invalid file name")
	}

	return path.Join(append([]string{ds.Info.GetDatastoreInfo().Url, "contentlib-" + l.ID}, id)...)
}

func (s *handler) libraryItemFileCreate(
	up *update,
	dstFileName string,
	body io.ReadCloser,
	cs *library.Checksum) error {

	defer body.Close()

	if !isValidFileName(dstFileName) {
		return errors.New("invalid file name")
	}

	dstItemPath := s.libraryPath(up.Library, up.Session.LibraryItemID)
	if err := os.MkdirAll(dstItemPath, 0750); err != nil {
		return err
	}

	// handleFile is used to process non-OVA files or files inside of an OVA.
	handleFile := func(
		fileName string,
		src io.Reader,
		doChecksum bool) (library.File, error) {

		dstFilePath := path.Join(dstItemPath, fileName)

		dst, err := openFile(dstFilePath, createOrCopyFlags, createOrCopyMode)
		if err != nil {
			return library.File{}, err
		}
		defer dst.Close()

		var h hash.Hash

		if doChecksum {
			if hasChecksum(cs) {
				h = checksum[cs.Algorithm]()
				src = io.TeeReader(src, h)
			}
		}

		n, err := copyReaderToWriter(dst, dstFilePath, src, fileName)
		if err != nil {
			return library.File{}, err
		}

		if h != nil {
			if sum := fmt.Sprintf("%x", h.Sum(nil)); sum != cs.Checksum {
				return library.File{}, fmt.Errorf(
					"checksum mismatch: file=%s, alg=%s, actual=%s, expected=%s",
					fileName, cs.Algorithm, sum, cs.Checksum)
			}
		}

		return library.File{
			Cached:  types.NewBool(true),
			Name:    fileName,
			Size:    &n,
			Version: "1",
		}, nil
	}

	// If the file being uploaded is not an OVA then it can be received
	// directly.
	if !strings.EqualFold(path.Ext(dstFileName), ".ova") {

		// Handle the non-OVA file.
		f, err := handleFile(dstFileName, body, true)
		if err != nil {
			return err
		}

		// Update the library item with the uploaded file.
		i := s.Library[up.Library.ID].Item[up.Session.LibraryItemID]
		i.File = append(i.File, f)
		return nil
	}

	// If this is an OVA then the entire OVA is hashed.
	var (
		h   hash.Hash
		src io.Reader = body
	)

	// See if the provided checksum is using a supported algorithm.
	if hasChecksum(cs) {
		h = checksum[cs.Algorithm]()
		src = io.TeeReader(src, h)
	}

	// Otherwise the contents of the OVA should be uploaded.
	r := tar.NewReader(src)

	// Collect the files from the OVA.
	var files []library.File
	for {
		h, err := r.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to unwind ova: %w", err)
		}
		if isValidFileName(h.Name) {

			// Tell the handleFile method *not* to do a checksum on the file
			// from the OVA. The checksum will occur on the entire OVA once its
			// contents have been read.
			f, err := handleFile(h.Name, io.LimitReader(r, h.Size), false)
			if err != nil {
				return err
			}

			files = append(files, f)
		}
	}

	// If there was a checksum provided then verify the entire OVA matches the
	// provided checksum.
	if h != nil {
		if sum := fmt.Sprintf("%x", h.Sum(nil)); sum != cs.Checksum {
			return fmt.Errorf(
				"checksum mismatch: file=%s, alg=%s, actual=%s, expected=%s",
				dstFileName, cs.Algorithm, sum, cs.Checksum)
		}
	}

	// Update the library item with the uploaded files.
	i := s.Library[up.Library.ID].Item[up.Session.LibraryItemID]
	i.File = files

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
		p := path.Join(s.libraryPath(dl.Library, dl.Session.LibraryItemID), name)
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

	err := s.libraryItemFileCreate(up, name, r.Body, nil)
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

func (i *item) cp() *item {
	nitem := *i.Item

	nfile := make([]library.File, len(i.File))
	copy(nfile, i.File)

	var nref *types.ManagedObjectReference
	if i.Template != nil {
		iref := *i.Template
		nref = &iref
	}

	return &item{
		Item:     &nitem,
		File:     nfile,
		Template: nref,
	}
}

func (i *item) ovf() string {
	for _, f := range i.File {
		if strings.HasSuffix(f.Name, ".ovf") {
			return f.Name
		}
	}
	return ""
}

func vmConfigSpec(ctx context.Context, c *vim25.Client, deploy vcenter.Deploy) (*types.VirtualMachineConfigSpec, error) {
	if deploy.VmConfigSpec == nil {
		return nil, nil
	}

	b, err := base64.StdEncoding.DecodeString(deploy.VmConfigSpec.XML)
	if err != nil {
		return nil, err
	}

	var spec *types.VirtualMachineConfigSpec

	dec := xml.NewDecoder(bytes.NewReader(b))
	dec.TypeFunc = c.Types
	err = dec.Decode(&spec)
	if err != nil {
		return nil, err
	}

	return spec, nil
}

func (s *handler) libraryDeploy(ctx context.Context, c *vim25.Client, lib *library.Library, item *item, deploy vcenter.Deploy) (*nfc.LeaseInfo, error) {
	config, err := vmConfigSpec(ctx, c, deploy)
	if err != nil {
		return nil, err
	}

	name := item.ovf()
	desc, err := os.ReadFile(filepath.Join(s.libraryPath(lib, item.ID), name))
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

	v, err := view.NewManager(c).CreateContainerView(ctx, c.ServiceContent.RootFolder, nil, true)
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

	m := ovf.NewManager(c)
	spec, err := m.CreateImportSpec(ctx, string(desc), pool, ds, &cisp)
	if err != nil {
		return nil, err
	}
	if spec.Error != nil {
		return nil, errors.New(spec.Error[0].LocalizedMessage)
	}

	if config != nil {
		if vmImportSpec, ok := spec.ImportSpec.(*types.VirtualMachineImportSpec); ok {
			var configSpecs []types.BaseVirtualDeviceConfigSpec

			// Remove devices that we don't want to carry over from the import spec. Otherwise, since we
			// just reconfigure the VM with the provided ConfigSpec later these devices won't be removed.
			for _, d := range vmImportSpec.ConfigSpec.DeviceChange {
				switch d.GetVirtualDeviceConfigSpec().Device.(type) {
				case types.BaseVirtualEthernetCard:
				default:
					configSpecs = append(configSpecs, d)
				}
			}
			vmImportSpec.ConfigSpec.DeviceChange = configSpecs
		}
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

	lease := nfc.NewLease(c, res.Returnval)
	info, err := lease.Wait(ctx, spec.FileItem)
	if err != nil {
		return nil, err
	}

	if err = lease.Complete(ctx); err != nil {
		return nil, err
	}

	if config != nil {
		if err = s.reconfigVM(info.Entity, *config); err != nil {
			return nil, err
		}
	}

	return info, nil
}

func (s *handler) libraryItemOVF(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req vcenter.OVF
	if !s.decode(r, w, &req) {
		return
	}

	switch {
	case req.Target.LibraryItemID != "":
	case req.Target.LibraryID != "":
		l, ok := s.Library[req.Target.LibraryID]
		if !ok {
			http.NotFound(w, r)
		}

		id := uuid.New().String()
		l.Item[id] = &item{
			Item: &library.Item{
				ID:               id,
				LibraryID:        l.Library.ID,
				Name:             req.Spec.Name,
				Description:      &req.Spec.Description,
				Type:             library.ItemTypeOVF,
				CreationTime:     types.NewTime(time.Now()),
				LastModifiedTime: types.NewTime(time.Now()),
			},
		}

		res := vcenter.CreateResult{
			Succeeded: true,
			ID:        id,
		}
		OK(w, res)
	default:
		BadRequest(w, "com.vmware.vapi.std.errors.invalid_argument")
		return
	}
}

func (s *handler) libraryItemOVFID(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("libraryItemOVFID: library item not found: %q", id)
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
		err := s.withClient(func(ctx context.Context, c *vim25.Client) error {
			info, err := s.libraryDeploy(ctx, c, lib, item, spec.Deploy)
			if err != nil {
				return err
			}
			id := vcenter.ResourceID{
				Type:  info.Entity.Type,
				Value: info.Entity.Value,
			}
			d.Succeeded = true
			d.ResourceID = &id
			return nil
		})
		if err != nil {
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

func (s *handler) deleteVM(ref *types.ManagedObjectReference) {
	if ref == nil {
		return
	}
	_ = s.withClient(func(ctx context.Context, c *vim25.Client) error {
		_, _ = object.NewVirtualMachine(c, *ref).Destroy(ctx)
		return nil
	})
}

func (s *handler) reconfigVM(ref types.ManagedObjectReference, config types.VirtualMachineConfigSpec) error {
	return s.withClient(func(ctx context.Context, c *vim25.Client) error {
		vm := object.NewVirtualMachine(c, ref)
		task, err := vm.Reconfigure(ctx, config)
		if err != nil {
			return err
		}
		return task.Wait(ctx)
	})
}

func (s *handler) cloneVM(source string, name string, p *library.Placement, storage *vcenter.DiskStorage) (*types.ManagedObjectReference, error) {
	var folder, pool, host, ds *types.ManagedObjectReference
	if p.Folder != "" {
		folder = &types.ManagedObjectReference{Type: "Folder", Value: p.Folder}
	}
	if p.ResourcePool != "" {
		pool = &types.ManagedObjectReference{Type: "ResourcePool", Value: p.ResourcePool}
	}
	if p.Host != "" {
		host = &types.ManagedObjectReference{Type: "HostSystem", Value: p.Host}
	}
	if storage != nil {
		if storage.Datastore != "" {
			ds = &types.ManagedObjectReference{Type: "Datastore", Value: storage.Datastore}
		}
	}

	spec := types.VirtualMachineCloneSpec{
		Template: true,
		Location: types.VirtualMachineRelocateSpec{
			Folder:    folder,
			Pool:      pool,
			Host:      host,
			Datastore: ds,
		},
	}

	var ref *types.ManagedObjectReference

	return ref, s.withClient(func(ctx context.Context, c *vim25.Client) error {
		vm := object.NewVirtualMachine(c, types.ManagedObjectReference{Type: "VirtualMachine", Value: source})

		task, err := vm.Clone(ctx, object.NewFolder(c, *folder), name, spec)
		if err != nil {
			return err
		}
		res, err := task.WaitForResult(ctx, nil)
		if err != nil {
			return err
		}
		ref = types.NewReference(res.Result.(types.ManagedObjectReference))
		return nil
	})
}

func (s *handler) libraryItemCreateTemplate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var spec struct {
		vcenter.Template `json:"spec"`
	}
	if !s.decode(r, w, &spec) {
		return
	}

	l, ok := s.Library[spec.Library]
	if !ok {
		http.NotFound(w, r)
		return
	}

	ds := &vcenter.DiskStorage{Datastore: l.Library.Storage[0].DatastoreID}
	ref, err := s.cloneVM(spec.SourceVM, spec.Name, spec.Placement, ds)
	if err != nil {
		BadRequest(w, err.Error())
		return
	}

	id := uuid.New().String()
	l.Item[id] = &item{
		Item: &library.Item{
			ID:               id,
			LibraryID:        l.Library.ID,
			Name:             spec.Name,
			Type:             library.ItemTypeVMTX,
			CreationTime:     types.NewTime(time.Now()),
			LastModifiedTime: types.NewTime(time.Now()),
		},
		Template: ref,
	}

	OK(w, id)
}

func (s *handler) libraryItemTemplateID(w http.ResponseWriter, r *http.Request) {
	// Go's ServeMux doesn't support wildcard matching, hacking around that for now to support
	// CheckOuts, e.g. "/vcenter/vm-template/library-items/{item}/check-outs/{vm}?action=check-in"
	p := strings.TrimPrefix(r.URL.Path, rest.Path+internal.VCenterVMTXLibraryItem+"/")
	route := strings.Split(p, "/")
	if len(route) == 0 {
		http.NotFound(w, r)
		return
	}

	id := route[0]
	ok := false

	var item *item
	for _, l := range s.Library {
		item, ok = l.Item[id]
		if ok {
			break
		}
	}
	if !ok {
		log.Printf("libraryItemTemplateID: library item not found: %q", id)
		http.NotFound(w, r)
		return
	}

	if item.Type != library.ItemTypeVMTX {
		BadRequest(w, "com.vmware.vapi.std.errors.invalid_argument")
		return
	}

	if len(route) > 1 {
		switch route[1] {
		case "check-outs":
			s.libraryItemCheckOuts(item, w, r)
			return
		default:
			http.NotFound(w, r)
			return
		}
	}

	if r.Method == http.MethodGet {
		// TODO: add mock data
		t := &vcenter.TemplateInfo{}
		OK(w, t)
		return
	}

	var spec struct {
		vcenter.DeployTemplate `json:"spec"`
	}
	if !s.decode(r, w, &spec) {
		return
	}

	switch r.URL.Query().Get("action") {
	case "deploy":
		p := spec.Placement
		if p == nil {
			BadRequest(w, "com.vmware.vapi.std.errors.invalid_argument")
			return
		}
		if p.Cluster == "" && p.Host == "" && p.ResourcePool == "" {
			BadRequest(w, "com.vmware.vapi.std.errors.invalid_argument")
			return
		}

		s.syncItem(item, nil, nil, true, nil)
		ref, err := s.cloneVM(item.Template.Value, spec.Name, p, spec.DiskStorage)
		if err != nil {
			BadRequest(w, err.Error())
			return
		}
		OK(w, ref.Value)
	default:
		http.NotFound(w, r)
	}
}

func (s *handler) libraryItemCheckOuts(item *item, w http.ResponseWriter, r *http.Request) {
	switch r.URL.Query().Get("action") {
	case "check-out":
		var spec struct {
			*vcenter.CheckOut `json:"spec"`
		}
		if !s.decode(r, w, &spec) {
			return
		}

		ref, err := s.cloneVM(item.Template.Value, spec.Name, spec.Placement, nil)
		if err != nil {
			BadRequest(w, err.Error())
			return
		}
		OK(w, ref.Value)
	case "check-in":
		// TODO: increment ContentVersion
		OK(w, "0")
	default:
		http.NotFound(w, r)
	}
}

// defaultSecurityPolicies generates the initial set of security policies always present on vCenter.
func defaultSecurityPolicies() []library.ContentSecurityPoliciesInfo {
	policyID, _ := uuid.NewUUID()
	return []library.ContentSecurityPoliciesInfo{
		{
			ItemTypeRules: map[string]string{
				"ovf": "OVF_STRICT_VERIFICATION",
			},
			Name:   "OVF default policy",
			Policy: policyID.String(),
		},
	}
}

func (s *handler) librarySecurityPolicies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		StatusOK(w, s.Policies)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *handler) isValidSecurityPolicy(policy string) bool {
	if policy == "" {
		return true
	}

	for _, p := range s.Policies {
		if p.Policy == policy {
			return true
		}
	}
	return false
}

func (s *handler) libraryTrustedCertificates(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var res struct {
			Certificates []library.TrustedCertificateSummary `json:"certificates"`
		}
		for id, cert := range s.Trust {
			res.Certificates = append(res.Certificates, library.TrustedCertificateSummary{
				TrustedCertificate: cert,
				ID:                 id,
			})
		}

		StatusOK(w, &res)
	case http.MethodPost:
		var info library.TrustedCertificate
		if s.decode(r, w, &info) {
			block, _ := pem.Decode([]byte(info.Text))
			if block == nil {
				s.error(w, errors.New("invalid certificate"))
				return
			}
			_, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				s.error(w, err)
				return
			}

			id := uuid.New().String()
			for x, cert := range s.Trust {
				if info.Text == cert.Text {
					id = x // existing certificate
					break
				}
			}
			s.Trust[id] = info

			w.WriteHeader(http.StatusCreated)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *handler) libraryTrustedCertificatesID(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	cert, ok := s.Trust[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		StatusOK(w, &cert)
	case http.MethodDelete:
		delete(s.Trust, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *handler) debugEcho(w http.ResponseWriter, r *http.Request) {
	r.Write(w)
}

func isValidFileName(s string) bool {
	return !strings.Contains(s, "/") &&
		!strings.Contains(s, "\\") &&
		!strings.Contains(s, "..")
}

func getVersionString(current string) string {
	if current == "" {
		return "1"
	}
	i, err := strconv.Atoi(current)
	if err != nil {
		panic(err)
	}
	i += 1
	return strconv.Itoa(i)
}
