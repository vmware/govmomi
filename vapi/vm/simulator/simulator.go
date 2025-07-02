// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	vapi "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/vapi/vm/dataset"
	"github.com/vmware/govmomi/vapi/vm/internal"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

const (
	// Minimal VM hardware version which supports DataSets feature
	minVmHardwareVersionDataSets = "vmx-20"

	typeVM = "VirtualMachine"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		New(s.Listen).Register(s, r)
	})
}

type Handler struct {
	u        *url.URL
	registry *simulator.Registry
}

func New(u *url.URL) *Handler {
	h := &Handler{
		u: u,
	}
	return h
}

const (
	restPathPrefix = rest.Path + internal.LegacyVCenterVMPath + "/"
	apiPathPrefix  = internal.VCenterVMPath + "/"
)

func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	if r.IsVPX() {
		h.registry = r
		s.HandleFunc(restPathPrefix, h.handle)
		s.HandleFunc(apiPathPrefix, h.handle)
	}
}

// path starts with "/api/vcenter/vm"
func (h *Handler) handle(w http.ResponseWriter, r *http.Request) {
	// The standard http.ServeMux does not support placeholders, so traverse the path segment by segment.
	// 'tail' tracks the remaining path segments.
	p := r.URL.Path
	p = strings.TrimPrefix(p, restPathPrefix)
	p = strings.TrimPrefix(p, apiPathPrefix)
	tail := strings.Split(p, "/")
	if len(tail) == 0 {
		// "/api/vcenter/vm"
		http.NotFound(w, r)
	} else {
		// "/api/vcenter/vm/..."
		switch tail[0] {
		case "":
			http.NotFound(w, r)
			return
		default:
			vmId := tail[0]
			h.handleVm(w, r, tail[1:], vmId)
		}
	}
}

// path starts with "/api/vcenter/vm/{}"
func (h *Handler) handleVm(w http.ResponseWriter, r *http.Request, tail []string, vmId string) {
	vm := h.validateVmExists(w, r, vmId)
	if vm == nil {
		return
	}
	ctx := &simulator.Context{
		Context: context.Background(),
		Session: &simulator.Session{
			UserSession: types.UserSession{
				Key: uuid.New().String(),
			},
			Registry: h.registry,
		},
		Map: h.registry,
	}
	h.registry.WithLock(ctx, vm.Reference(), func() {
		if len(tail) == 0 {
			// "/api/vcenter/vm/{}"
			switch r.Method {
			case http.MethodDelete:
				h.deleteVM(w, r, ctx, vm)
			default:
				http.NotFound(w, r)
			}
		} else {
			// "/api/vcenter/vm/{}/..."
			switch tail[0] {
			case "data-sets":
				h.handleVmDataSets(w, r, tail[1:], vm)
			default:
				http.NotFound(w, r)
			}
		}
	})
}

// path starts with "/api/vcenter/vm/{}/data-sets"
func (h *Handler) handleVmDataSets(w http.ResponseWriter, r *http.Request, tail []string, vm *simulator.VirtualMachine) {
	if !h.validateVmHardwareVersionDataSets(w, r, vm) {
		return
	}
	if len(tail) == 0 {
		// "/api/vcenter/vm/{}/data-sets"
		switch r.Method {
		case http.MethodGet:
			h.listDataSets(w, r, vm)
		case http.MethodPost:
			h.createDataSet(w, r, vm)
		default:
			http.NotFound(w, r)
		}
	} else {
		// "/api/vcenter/vm/{}/data-sets/..."
		switch tail[0] {
		case "":
			http.NotFound(w, r)
			return
		default:
			dataSetId := tail[0]
			h.handleVmDataSet(w, r, tail[1:], vm, dataSetId)
		}
	}
}

// path starts with "/api/vcenter/vm/{}/data-sets/{}"
func (h *Handler) handleVmDataSet(w http.ResponseWriter, r *http.Request, tail []string, vm *simulator.VirtualMachine, dataSetId string) {
	dataSet := h.validateDataSetExists(w, r, vm, dataSetId)
	if dataSet == nil {
		return
	}
	if len(tail) == 0 {
		// "/api/vcenter/vm/{}/data-sets/{}"
		switch r.Method {
		case http.MethodGet:
			vapi.StatusOK(w, dataSet.Info)
		case http.MethodPatch:
			h.updateDataSet(w, r, vm, dataSet)
		case http.MethodDelete:
			h.deleteDataSet(w, r, vm, dataSet)
		default:
			http.NotFound(w, r)
		}
	} else {
		// "/api/vcenter/vm/{}/data-sets/{}/..."
		switch tail[0] {
		case "entries":
			h.handleVmDataSetEntries(w, r, tail[1:], vm, dataSet)
		default:
			http.NotFound(w, r)
		}

	}
}

// path starts with "/api/vcenter/vm/{}/data-sets/{}/entries"
func (h *Handler) handleVmDataSetEntries(w http.ResponseWriter, r *http.Request, tail []string, vm *simulator.VirtualMachine, dataSet *simulator.DataSet) {
	if len(tail) == 0 {
		// "/api/vcenter/vm/{}/data-sets/{}/entries"
		switch r.Method {
		case http.MethodGet:
			h.listDataSetEntries(w, r, vm, dataSet)
		default:
			http.NotFound(w, r)
		}
	} else {
		// "/api/vcenter/vm/{}/data-sets/{}/entries/..."
		switch tail[0] {
		case "":
			http.NotFound(w, r)
			return
		default:
			entryKey := tail[0]
			h.handleVmDataSetEntry(w, r, tail[1:], vm, dataSet, entryKey)
		}
	}
}

// path starts with "/api/vcenter/vm/{}/data-sets/{}/entries/{}"
func (h *Handler) handleVmDataSetEntry(w http.ResponseWriter, r *http.Request, tail []string, vm *simulator.VirtualMachine, dataSet *simulator.DataSet, entryKey string) {
	if len(tail) > 0 {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		h.getDataSetEntry(w, r, vm, dataSet, entryKey)
	case http.MethodPut:
		h.setDataSetEntry(w, r, vm, dataSet, entryKey)
	case http.MethodDelete:
		h.deleteDataSetEntry(w, r, vm, dataSet, entryKey)
	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) deleteVM(w http.ResponseWriter, r *http.Request, ctx *simulator.Context, vm *simulator.VirtualMachine) {
	taskRef := vm.DestroyTask(ctx, &types.Destroy_Task{This: vm.Self}).(*methods.Destroy_TaskBody).Res.Returnval
	task := ctx.Map.Get(taskRef).(*simulator.Task)
	task.Wait()
	if task.Info.Error != nil {
		log.Printf("%s %s: %v", r.Method, r.RequestURI, task.Info.Error)
		vapi.ApiErrorGeneral(w)
		return
	}
	vapi.StatusOK(w)
}

func (h *Handler) createDataSet(w http.ResponseWriter, r *http.Request, vm *simulator.VirtualMachine) {
	if !h.validateVmIsNotSuspended(w, r, vm) {
		return
	}
	var createSpec dataset.CreateSpec
	if !vapi.Decode(r, w, &createSpec) {
		return
	}
	if createSpec.Name == "" {
		vapi.ApiErrorInvalidArgument(w)
		return
	}
	if createSpec.Host == "" {
		vapi.ApiErrorInvalidArgument(w)
		return
	}
	if createSpec.Guest == "" {
		vapi.ApiErrorInvalidArgument(w)
		return
	}
	dataSetId := createSpec.Name
	_, ok := vm.DataSets[dataSetId]
	if ok {
		vapi.ApiErrorAlreadyExists(w)
		return
	}
	dataSet := &simulator.DataSet{
		Info: &dataset.Info{
			Name:                     createSpec.Name,
			Description:              createSpec.Description,
			Host:                     createSpec.Host,
			Guest:                    createSpec.Guest,
			Used:                     0,
			OmitFromSnapshotAndClone: getOrDefault(createSpec.OmitFromSnapshotAndClone, false),
		},
		ID:      dataSetId,
		Entries: make(map[string]string),
	}
	vm.DataSets[dataSetId] = dataSet
	vapi.StatusOK(w, dataSetId)
}

func (h *Handler) listDataSets(w http.ResponseWriter, r *http.Request, vm *simulator.VirtualMachine) {
	var result []dataset.Summary = make([]dataset.Summary, 0, len(vm.DataSets))
	for _, v := range vm.DataSets {
		summary := dataset.Summary{
			DataSet:     v.ID,
			Name:        v.Name,
			Description: v.Description,
		}
		result = append(result, summary)
	}
	vapi.StatusOK(w, result)
}

func (h *Handler) deleteDataSet(w http.ResponseWriter, r *http.Request, vm *simulator.VirtualMachine, dataSet *simulator.DataSet) {
	if !h.validateVmIsNotSuspended(w, r, vm) {
		return
	}
	force := strings.EqualFold(r.URL.Query().Get("force"), "true")
	if len(dataSet.Entries) > 0 && !force {
		// cannot delete non-empty data set without force
		vapi.ApiErrorResourceInUse(w)
		return
	}
	delete(vm.DataSets, dataSet.ID)
	vapi.StatusOK(w)
}

func (h *Handler) updateDataSet(w http.ResponseWriter, r *http.Request, vm *simulator.VirtualMachine, dataSet *simulator.DataSet) {
	if !h.validateVmIsNotSuspended(w, r, vm) {
		return
	}
	var updateSpec dataset.UpdateSpec
	if !vapi.Decode(r, w, &updateSpec) {
		return
	}
	if updateSpec.Description != nil {
		dataSet.Description = *updateSpec.Description
	}
	if updateSpec.Host != nil {
		dataSet.Host = *updateSpec.Host
	}
	if updateSpec.Guest != nil {
		dataSet.Guest = *updateSpec.Guest
	}
	if updateSpec.OmitFromSnapshotAndClone != nil {
		dataSet.OmitFromSnapshotAndClone = *updateSpec.OmitFromSnapshotAndClone
	}
	vapi.StatusOK(w)
}

func (h *Handler) listDataSetEntries(w http.ResponseWriter, r *http.Request, vm *simulator.VirtualMachine, dataSet *simulator.DataSet) {
	if dataSet.Host == dataset.AccessNone {
		vapi.ApiErrorUnauthorized(w)
		return
	}
	var result []string = make([]string, 0, len(dataSet.Entries))
	for k := range dataSet.Entries {
		result = append(result, k)
	}
	vapi.StatusOK(w, result)
}

func (h *Handler) getDataSetEntry(w http.ResponseWriter, r *http.Request, vm *simulator.VirtualMachine, dataSet *simulator.DataSet, entryKey string) {
	if dataSet.Host == dataset.AccessNone {
		vapi.ApiErrorUnauthorized(w)
		return
	}
	val, ok := dataSet.Entries[entryKey]
	if !ok {
		vapi.ApiErrorNotFound(w)
		return
	}
	vapi.StatusOK(w, val)
}

const (
	// A key can be at most 4096 bytes.
	entryKeyMaxLen = 4096
	// A value can be at most 1MB.
	entryValueMaxLen = 1024 * 1024
)

func (h *Handler) setDataSetEntry(w http.ResponseWriter, r *http.Request, vm *simulator.VirtualMachine, dataSet *simulator.DataSet, entryKey string) {
	var val string
	if !vapi.Decode(r, w, &val) {
		return
	}
	if !h.validateVmIsNotSuspended(w, r, vm) {
		return
	}
	if dataSet.Host != dataset.AccessReadWrite {
		vapi.ApiErrorUnauthorized(w)
		return
	}
	if len(entryKey) > entryKeyMaxLen {
		vapi.ApiErrorInvalidArgument(w)
		return
	}
	if len(val) > entryValueMaxLen {
		vapi.ApiErrorInvalidArgument(w)
		return
	}
	old, ok := dataSet.Entries[entryKey]
	if ok {
		dataSet.Used -= len(entryKey)
		dataSet.Used -= len(old)
	}
	dataSet.Entries[entryKey] = val
	dataSet.Used += len(entryKey)
	dataSet.Used += len(val)
	vapi.StatusOK(w)
}

func (h *Handler) deleteDataSetEntry(w http.ResponseWriter, r *http.Request, vm *simulator.VirtualMachine, dataSet *simulator.DataSet, entryKey string) {
	if !h.validateVmIsNotSuspended(w, r, vm) {
		return
	}
	if dataSet.Host != dataset.AccessReadWrite {
		vapi.ApiErrorUnauthorized(w)
		return
	}
	val, ok := dataSet.Entries[entryKey]
	if !ok {
		vapi.ApiErrorNotFound(w)
		return
	}
	dataSet.Used -= len(entryKey)
	dataSet.Used -= len(val)
	delete(dataSet.Entries, entryKey)
	vapi.StatusOK(w)
}

func getOrDefault(b *bool, defaultValue bool) bool {
	if b == nil {
		return defaultValue
	}
	return *b
}

func (h *Handler) validateVmExists(w http.ResponseWriter, r *http.Request, vmId string) *simulator.VirtualMachine {
	vm, ok := h.registry.Get(types.ManagedObjectReference{Type: typeVM, Value: vmId}).(*simulator.VirtualMachine)
	if !ok {
		vapi.ApiErrorNotFound(w)
		return nil
	}
	return vm
}

func (h *Handler) validateDataSetExists(w http.ResponseWriter, r *http.Request, vm *simulator.VirtualMachine, dataSetId string) *simulator.DataSet {
	dataSet, ok := vm.DataSets[dataSetId]
	if !ok {
		vapi.ApiErrorNotFound(w)
		return nil
	}
	return dataSet
}

func (h *Handler) validateVmHardwareVersionDataSets(w http.ResponseWriter, r *http.Request, vm *simulator.VirtualMachine) bool {
	version := vm.Config.Version
	if len(version) < len(minVmHardwareVersionDataSets) {
		vapi.ApiErrorUnsupported(w)
		return false
	}
	if len(version) == len(minVmHardwareVersionDataSets) && version < minVmHardwareVersionDataSets {
		vapi.ApiErrorUnsupported(w)
		return false
	}
	return true
}

func (h *Handler) validateVmIsNotSuspended(w http.ResponseWriter, r *http.Request, vm *simulator.VirtualMachine) bool {
	if vm.Summary.Runtime.PowerState == types.VirtualMachinePowerStateSuspended {
		vapi.ApiErrorNotAllowedInCurrentState(w)
		return false
	}
	return true
}
