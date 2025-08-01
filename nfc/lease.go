// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package nfc

import (
	"context"
	"fmt"
	"io"
	"path"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type Lease struct {
	types.ManagedObjectReference

	c *vim25.Client
}

func NewLease(c *vim25.Client, ref types.ManagedObjectReference) *Lease {
	return &Lease{ref, c}
}

// Abort wraps methods.Abort
func (l *Lease) Abort(ctx context.Context, fault *types.LocalizedMethodFault) error {
	req := types.HttpNfcLeaseAbort{
		This:  l.Reference(),
		Fault: fault,
	}

	_, err := methods.HttpNfcLeaseAbort(ctx, l.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// Complete wraps methods.Complete
func (l *Lease) Complete(ctx context.Context) error {
	req := types.HttpNfcLeaseComplete{
		This: l.Reference(),
	}

	_, err := methods.HttpNfcLeaseComplete(ctx, l.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// GetManifest wraps methods.GetManifest
func (l *Lease) GetManifest(ctx context.Context) ([]types.HttpNfcLeaseManifestEntry, error) {
	req := types.HttpNfcLeaseGetManifest{
		This: l.Reference(),
	}

	res, err := methods.HttpNfcLeaseGetManifest(ctx, l.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

// Progress wraps methods.Progress
func (l *Lease) Progress(ctx context.Context, percent int32) error {
	req := types.HttpNfcLeaseProgress{
		This:    l.Reference(),
		Percent: percent,
	}

	_, err := methods.HttpNfcLeaseProgress(ctx, l.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// Properties returns a mo.HttpNfcLease with the specified properties.
// If no properties are requested, all properties are returned.
func (l *Lease) Properties(
	ctx context.Context,
	props ...string) (mo.HttpNfcLease, error) {

	if len(props) == 0 {
		props = []string{
			"initializeProgress",
			"transferProgress",
			"mode",
			"capabilities",
			"info",
			"state",
			"error",
		}
	}

	var o mo.HttpNfcLease
	pc := property.DefaultCollector(l.c)
	return o, pc.RetrieveOne(ctx, l.Reference(), props, &o)
}

type LeaseInfo struct {
	types.HttpNfcLeaseInfo

	Items []FileItem
}

func (l *Lease) newLeaseInfo(li *types.HttpNfcLeaseInfo, items []types.OvfFileItem) (*LeaseInfo, error) {
	info := &LeaseInfo{
		HttpNfcLeaseInfo: *li,
	}

	for _, device := range li.DeviceUrl {
		u, err := l.c.ParseURL(device.Url)
		if err != nil {
			return nil, err
		}

		if device.SslThumbprint != "" {
			// TODO: prefer host management IP
			l.c.SetThumbprint(u.Host, device.SslThumbprint)
		}

		if len(items) == 0 {
			// this is an export
			item := types.OvfFileItem{
				DeviceId: device.Key,
				Path:     device.TargetId,
				Size:     device.FileSize,
			}

			if item.Size == 0 {
				item.Size = li.TotalDiskCapacityInKB * 1024
			}

			if item.Path == "" {
				item.Path = path.Base(device.Url)
			}

			info.Items = append(info.Items, NewFileItem(u, item))

			continue
		}

		// this is an import
		for _, item := range items {
			if device.ImportKey == item.DeviceId {
				fi := NewFileItem(u, item)
				fi.Thumbprint = device.SslThumbprint

				info.Items = append(info.Items, fi)

				break
			}
		}
	}

	return info, nil
}

func (l *Lease) Wait(ctx context.Context, items []types.OvfFileItem) (*LeaseInfo, error) {
	var lease mo.HttpNfcLease

	pc := property.DefaultCollector(l.c)
	err := property.Wait(ctx, pc, l.Reference(), []string{"state", "info", "error"}, func(pc []types.PropertyChange) bool {
		done := false

		for _, c := range pc {
			if c.Val == nil {
				continue
			}

			switch c.Name {
			case "error":
				val := c.Val.(types.LocalizedMethodFault)
				lease.Error = &val
				done = true
			case "info":
				val := c.Val.(types.HttpNfcLeaseInfo)
				lease.Info = &val
			case "state":
				lease.State = c.Val.(types.HttpNfcLeaseState)
				if lease.State != types.HttpNfcLeaseStateInitializing {
					done = true
				}
			}
		}

		return done
	})

	if err != nil {
		return nil, err
	}

	if lease.State == types.HttpNfcLeaseStateReady {
		return l.newLeaseInfo(lease.Info, items)
	}

	if lease.Error != nil {
		return nil, &task.Error{LocalizedMethodFault: lease.Error}
	}

	return nil, fmt.Errorf("unexpected nfc lease state: %s", lease.State)
}

func (l *Lease) StartUpdater(ctx context.Context, info *LeaseInfo) *LeaseUpdater {
	return newLeaseUpdater(ctx, l, info)
}

func (l *Lease) Upload(ctx context.Context, item FileItem, f io.Reader, opts soap.Upload) error {
	if opts.Progress == nil {
		opts.Progress = item
	}

	// Non-disk files (such as .iso) use the PUT method.
	// Overwrite: t header is also required in this case (ovftool does the same)
	if item.Create {
		opts.Method = "PUT"
		opts.Headers = map[string]string{
			"Overwrite": "t",
		}
	} else {
		opts.Method = "POST"
		opts.Type = "application/x-vnd.vmware-streamVmdk"
	}

	return l.c.Upload(ctx, f, item.URL, &opts)
}

func (l *Lease) DownloadFile(ctx context.Context, file string, item FileItem, opts soap.Download) error {
	if opts.Progress == nil {
		opts.Progress = item
	}

	return l.c.DownloadFile(ctx, file, item.URL, &opts)
}
