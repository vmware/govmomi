/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package importer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/nfc"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/ovf"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/progress"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type Importer struct {
	Log progress.LogFunc

	Name           string
	VerifyManifest bool
	Hidden         bool

	Client *vim25.Client
	Finder *find.Finder
	Sinker progress.Sinker

	Datacenter   *object.Datacenter
	Datastore    *object.Datastore
	ResourcePool *object.ResourcePool
	Host         *object.HostSystem
	Folder       *object.Folder

	Archive  Archive
	Manifest map[string]*library.Checksum
}

func (imp *Importer) manifestPath(fpath string) string {
	base := filepath.Base(fpath)
	ext := filepath.Ext(base)
	return filepath.Join(filepath.Dir(fpath), strings.Replace(base, ext, ".mf", 1))
}

func (imp *Importer) ReadManifest(fpath string) error {
	mf, _, err := imp.Archive.Open(imp.manifestPath(fpath))
	if err != nil {
		msg := fmt.Sprintf("failed to read manifest %q: %s", mf, err)
		return errors.New(msg)
	}
	imp.Manifest, err = library.ReadManifest(mf)
	_ = mf.Close()
	return err
}

func (imp *Importer) Import(ctx context.Context, fpath string, opts Options) (*types.ManagedObjectReference, error) {

	o, err := ReadOvf(fpath, imp.Archive)
	if err != nil {
		return nil, err
	}

	e, err := ReadEnvelope(o)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ovf: %s", err)
	}

	if e.VirtualSystem != nil {
		if e.VirtualSystem != nil {
			if opts.Name == nil {
				opts.Name = &e.VirtualSystem.ID
				if e.VirtualSystem.Name != nil {
					opts.Name = e.VirtualSystem.Name
				}
			}
		}
		if imp.Hidden {
			// TODO: userConfigurable is optional and defaults to false, so we should *add* userConfigurable=true
			// if not set for a Property. But, there'd be a bunch more work involved to preserve other data in doing
			// a complete xml.Marshal of the .ovf
			o = bytes.ReplaceAll(o, []byte(`userConfigurable="false"`), []byte(`userConfigurable="true"`))
		}
	}

	name := "Govc Virtual Appliance"
	if opts.Name != nil {
		name = *opts.Name
	}

	nmap, err := imp.NetworkMap(ctx, e, opts.NetworkMapping)
	if err != nil {
		return nil, err
	}

	cisp := types.OvfCreateImportSpecParams{
		DiskProvisioning:   opts.DiskProvisioning,
		EntityName:         name,
		IpAllocationPolicy: opts.IPAllocationPolicy,
		IpProtocol:         opts.IPProtocol,
		OvfManagerCommonParams: types.OvfManagerCommonParams{
			DeploymentOption: opts.Deployment,
			Locale:           "US"},
		PropertyMapping: OVFMap(opts.PropertyMapping),
		NetworkMapping:  nmap,
	}

	m := ovf.NewManager(imp.Client)
	spec, err := m.CreateImportSpec(ctx, string(o), imp.ResourcePool, imp.Datastore, cisp)
	if err != nil {
		return nil, err
	}
	if spec.Error != nil {
		return nil, errors.New(spec.Error[0].LocalizedMessage)
	}
	if spec.Warning != nil {
		for _, w := range spec.Warning {
			_, _ = imp.Log(fmt.Sprintf("Warning: %s\n", w.LocalizedMessage))
		}
	}

	if opts.Annotation != "" {
		switch s := spec.ImportSpec.(type) {
		case *types.VirtualMachineImportSpec:
			s.ConfigSpec.Annotation = opts.Annotation
		case *types.VirtualAppImportSpec:
			s.VAppConfigSpec.Annotation = opts.Annotation
		}
	}

	if imp.VerifyManifest {
		if err := imp.ReadManifest(fpath); err != nil {
			return nil, err
		}
	}

	lease, err := imp.ResourcePool.ImportVApp(ctx, spec.ImportSpec, imp.Folder, imp.Host)
	if err != nil {
		return nil, err
	}

	info, err := lease.Wait(ctx, spec.FileItem)
	if err != nil {
		_ = lease.Abort(ctx, nil)
		return nil, err
	}

	u := lease.StartUpdater(ctx, info)
	defer u.Done()

	for _, i := range info.Items {
		if err := imp.Upload(ctx, lease, i); err != nil {
			_ = lease.Abort(ctx, &types.LocalizedMethodFault{
				Fault: &types.FileFault{
					File: i.Path,
				},
			})
			return nil, err
		}
	}

	return &info.Entity, lease.Complete(ctx)
}

func (imp *Importer) NetworkMap(ctx context.Context, e *ovf.Envelope, networks []Network) ([]types.OvfNetworkMapping, error) {
	var nmap []types.OvfNetworkMapping
	for _, m := range networks {
		if m.Network == "" {
			continue // Not set, let vSphere choose the default network
		}
		if err := ValidateNetwork(e, m); err != nil && imp.Log != nil {
			_, _ = imp.Log(err.Error() + "\n")
		}

		var ref types.ManagedObjectReference

		net, err := imp.Finder.Network(ctx, m.Network)
		if err != nil {
			switch err.(type) {
			case *find.NotFoundError:
				if !ref.FromString(m.Network) {
					return nil, err
				} // else this is a raw MO ref
			default:
				return nil, err
			}
		} else {
			ref = net.Reference()
		}

		nmap = append(nmap, types.OvfNetworkMapping{
			Name:    m.Name,
			Network: ref,
		})
	}

	return nmap, nil
}

func OVFMap(op []Property) (p []types.KeyValue) {
	for _, v := range op {
		p = append(p, types.KeyValue{
			Key:   v.Key,
			Value: v.Value,
		})
	}

	return
}

func ValidateNetwork(e *ovf.Envelope, net Network) error {
	var names []string

	if e.Network != nil {
		for _, n := range e.Network.Networks {
			if n.Name == net.Name {
				return nil
			}
			names = append(names, n.Name)
		}
	}

	return fmt.Errorf("warning: invalid NetworkMapping.Name=%q, valid names=%s", net.Name, names)
}

func ValidateChecksum(ctx context.Context, lease *nfc.Lease, sum *library.Checksum, file string, key string) error {
	// Perform the checksum match eagerly, after each file upload, instead
	// of after uploading all the files, to provide fail-fast behavior.
	// (Trade-off here is multiple GetManifest() API calls to the server.)
	manifests, err := lease.GetManifest(ctx)
	if err != nil {
		return err
	}
	for _, m := range manifests {
		if m.Key == key {
			// Compare server-side computed checksum of uploaded file
			// against the client's manifest entry (assuming client's
			// manifest has correct checksums - client doesn't compute
			// checksum of the file before uploading).

			// Try matching sha1 first (newer versions have moved to sha256).
			if strings.ToUpper(sum.Algorithm) == "SHA1" {
				if sum.Checksum != m.Sha1 {
					msg := fmt.Sprintf("manifest checksum %v mismatch with uploaded checksum %v for file %v",
						sum.Checksum, m.Sha1, file)
					return errors.New(msg)
				}
				// Uploaded file checksum computed by server matches with local manifest entry.
				return nil
			}
			// If not sha1, check for other types (in a separate field).
			if !strings.EqualFold(sum.Algorithm, m.ChecksumType) {
				msg := fmt.Sprintf("manifest checksum type %v mismatch with uploaded checksum type %v for file %v",
					sum.Algorithm, m.ChecksumType, file)
				return errors.New(msg)
			}
			if !strings.EqualFold(sum.Checksum, m.Checksum) {
				msg := fmt.Sprintf("manifest checksum %v mismatch with uploaded checksum %v for file %v",
					sum.Checksum, m.Checksum, file)
				return errors.New(msg)
			}
			// Uploaded file checksum computed by server matches with local manifest entry.
			return nil
		}
	}
	msg := fmt.Sprintf("missing manifest entry on server for uploaded file %v (key %v), manifests=%#v", file, key, manifests)
	return errors.New(msg)
}

func (imp *Importer) Upload(ctx context.Context, lease *nfc.Lease, item nfc.FileItem) error {
	file := item.Path

	f, size, err := imp.Archive.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	logger := progress.NewProgressLogger(imp.Log, fmt.Sprintf("Uploading %s... ", path.Base(file)))
	defer logger.Wait()

	opts := soap.Upload{
		ContentLength: size,
		Progress:      logger,
	}

	err = lease.Upload(ctx, item, f, opts)
	if err != nil {
		return err
	}

	if imp.VerifyManifest {
		mapImportKeyToKey := func(urls []types.HttpNfcLeaseDeviceUrl, importKey string) string {
			for _, url := range urls {
				if url.ImportKey == importKey {
					return url.Key
				}
			}
			return ""
		}
		leaseInfo, err := lease.Wait(ctx, nil)
		if err != nil {
			return err
		}
		sum, ok := imp.Manifest[file]
		if !ok {
			return fmt.Errorf("missing checksum for %v in manifest file", file)
		}
		return ValidateChecksum(ctx, lease, sum, file, mapImportKeyToKey(leaseInfo.DeviceUrl, item.DeviceId))
	}
	return nil
}
