/*
Copyright (c) 2014-2023 VMware, Inc. All Rights Reserved.

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

package importx

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"path"
	"strings"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/nfc"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/ovf"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type ovfx struct {
	*flags.DatastoreFlag
	*flags.HostSystemFlag
	*flags.OutputFlag
	*flags.ResourcePoolFlag
	*flags.FolderFlag

	*ArchiveFlag
	*OptionsFlag

	Name           string
	VerifyManifest bool
	Hidden         bool

	Client       *vim25.Client
	Datacenter   *object.Datacenter
	Datastore    *object.Datastore
	ResourcePool *object.ResourcePool
}

func init() {
	cli.Register("import.ovf", &ovfx{})
}

func (cmd *ovfx) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)
	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)

	cmd.ArchiveFlag, ctx = newArchiveFlag(ctx)
	cmd.ArchiveFlag.Register(ctx, f)
	cmd.OptionsFlag, ctx = newOptionsFlag(ctx)
	cmd.OptionsFlag.Register(ctx, f)

	f.StringVar(&cmd.Name, "name", "", "Name to use for new entity")
	f.BoolVar(&cmd.VerifyManifest, "m", false, "Verify checksum of uploaded files against manifest (.mf)")
	f.BoolVar(&cmd.Hidden, "hidden", false, "Enable hidden properties")
}

func (cmd *ovfx) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ArchiveFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OptionsFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.FolderFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ovfx) Usage() string {
	return "PATH_TO_OVF"
}

func (cmd *ovfx) Run(ctx context.Context, f *flag.FlagSet) error {
	fpath, err := cmd.Prepare(f)
	if err != nil {
		return err
	}

	archive := &FileArchive{Path: fpath}
	archive.Client = cmd.Client

	cmd.Archive = archive

	moref, err := cmd.Import(fpath)
	if err != nil {
		return err
	}

	vm := object.NewVirtualMachine(cmd.Client, *moref)
	return cmd.Deploy(vm, cmd.OutputFlag)
}

func (cmd *ovfx) Prepare(f *flag.FlagSet) (string, error) {
	var err error

	args := f.Args()
	if len(args) != 1 {
		return "", errors.New("no file specified")
	}

	cmd.Client, err = cmd.DatastoreFlag.Client()
	if err != nil {
		return "", err
	}

	cmd.Datacenter, err = cmd.DatastoreFlag.Datacenter()
	if err != nil {
		return "", err
	}

	cmd.Datastore, err = cmd.DatastoreFlag.Datastore()
	if err != nil {
		return "", err
	}

	cmd.ResourcePool, err = cmd.ResourcePoolFlag.ResourcePoolIfSpecified()
	if err != nil {
		return "", err
	}

	return f.Arg(0), nil
}

func (cmd *ovfx) Map(op []Property) (p []types.KeyValue) {
	for _, v := range op {
		p = append(p, types.KeyValue{
			Key:   v.Key,
			Value: v.Value,
		})
	}

	return
}

func (cmd *ovfx) validateNetwork(e *ovf.Envelope, net Network) {
	var names []string

	if e.Network != nil {
		for _, n := range e.Network.Networks {
			if n.Name == net.Name {
				return
			}
			names = append(names, n.Name)
		}
	}

	_, _ = cmd.Log(fmt.Sprintf("Warning: invalid NetworkMapping.Name=%q, valid names=%s\n", net.Name, names))
}

func (cmd *ovfx) NetworkMap(e *ovf.Envelope) ([]types.OvfNetworkMapping, error) {
	ctx := context.TODO()
	finder, err := cmd.DatastoreFlag.Finder()
	if err != nil {
		return nil, err
	}

	var nmap []types.OvfNetworkMapping
	for _, m := range cmd.Options.NetworkMapping {
		if m.Network == "" {
			continue // Not set, let vSphere choose the default network
		}
		cmd.validateNetwork(e, m)

		var ref types.ManagedObjectReference

		net, err := finder.Network(ctx, m.Network)
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

	return nmap, err
}

func (cmd *ovfx) Import(fpath string) (*types.ManagedObjectReference, error) {
	ctx := context.TODO()

	o, err := cmd.ReadOvf(fpath)
	if err != nil {
		return nil, err
	}

	e, err := cmd.ReadEnvelope(o)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ovf: %s", err)
	}

	name := "Govc Virtual Appliance"
	if e.VirtualSystem != nil {
		name = e.VirtualSystem.ID
		if e.VirtualSystem.Name != nil {
			name = *e.VirtualSystem.Name
		}

		if cmd.Hidden {
			// TODO: userConfigurable is optional and defaults to false, so we should *add* userConfigurable=true
			// if not set for a Property. But, there'd be a bunch more work involved to preserve other data in doing
			// a complete xml.Marshal of the .ovf
			o = bytes.ReplaceAll(o, []byte(`userConfigurable="false"`), []byte(`userConfigurable="true"`))
		}
	}

	// Override name from options if specified
	if cmd.Options.Name != nil {
		name = *cmd.Options.Name
	}

	// Override name from arguments if specified
	if cmd.Name != "" {
		name = cmd.Name
	}

	nmap, err := cmd.NetworkMap(e)
	if err != nil {
		return nil, err
	}

	cisp := types.OvfCreateImportSpecParams{
		DiskProvisioning:   cmd.Options.DiskProvisioning,
		EntityName:         name,
		IpAllocationPolicy: cmd.Options.IPAllocationPolicy,
		IpProtocol:         cmd.Options.IPProtocol,
		OvfManagerCommonParams: types.OvfManagerCommonParams{
			DeploymentOption: cmd.Options.Deployment,
			Locale:           "US"},
		PropertyMapping: cmd.Map(cmd.Options.PropertyMapping),
		NetworkMapping:  nmap,
	}

	host, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return nil, err
	}

	if cmd.ResourcePool == nil {
		if host == nil {
			cmd.ResourcePool, err = cmd.ResourcePoolFlag.ResourcePool()
		} else {
			cmd.ResourcePool, err = host.ResourcePool(ctx)
		}
		if err != nil {
			return nil, err
		}
	}

	m := ovf.NewManager(cmd.Client)
	spec, err := m.CreateImportSpec(ctx, string(o), cmd.ResourcePool, cmd.Datastore, cisp)
	if err != nil {
		return nil, err
	}
	if spec.Error != nil {
		return nil, errors.New(spec.Error[0].LocalizedMessage)
	}
	if spec.Warning != nil {
		for _, w := range spec.Warning {
			_, _ = cmd.Log(fmt.Sprintf("Warning: %s\n", w.LocalizedMessage))
		}
	}

	if cmd.Options.Annotation != "" {
		switch s := spec.ImportSpec.(type) {
		case *types.VirtualMachineImportSpec:
			s.ConfigSpec.Annotation = cmd.Options.Annotation
		case *types.VirtualAppImportSpec:
			s.VAppConfigSpec.Annotation = cmd.Options.Annotation
		}
	}

	var folder *object.Folder
	// The folder argument must not be set on a VM in a vApp, otherwise causes
	// InvalidArgument fault: A specified parameter was not correct: pool
	if cmd.ResourcePool.Reference().Type != "VirtualApp" {
		folder, err = cmd.FolderOrDefault("vm")
		if err != nil {
			return nil, err
		}
	}

	if cmd.VerifyManifest {
		err = cmd.readManifest(fpath)
		if err != nil {
			return nil, err
		}
	}

	lease, err := cmd.ResourcePool.ImportVApp(ctx, spec.ImportSpec, folder, host)
	if err != nil {
		return nil, err
	}

	info, err := lease.Wait(ctx, spec.FileItem)
	if err != nil {
		return nil, err
	}

	u := lease.StartUpdater(ctx, info)
	defer u.Done()

	for _, i := range info.Items {
		err = cmd.Upload(ctx, lease, i)
		if err != nil {
			return nil, err
		}
	}

	return &info.Entity, lease.Complete(ctx)
}

func (cmd *ovfx) Upload(ctx context.Context, lease *nfc.Lease, item nfc.FileItem) error {
	file := item.Path

	f, size, err := cmd.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	logger := cmd.ProgressLogger(fmt.Sprintf("Uploading %s... ", path.Base(file)))
	defer logger.Wait()

	opts := soap.Upload{
		ContentLength: size,
		Progress:      logger,
	}

	err = lease.Upload(ctx, item, f, opts)
	if err != nil {
		return err
	}

	if cmd.VerifyManifest {
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
		return cmd.validateChecksum(ctx, lease, file, mapImportKeyToKey(leaseInfo.DeviceUrl, item.DeviceId))
	}
	return nil
}

func (cmd *ovfx) validateChecksum(ctx context.Context, lease *nfc.Lease, file string, key string) error {
	sum, found := cmd.manifest[file]
	if !found {
		msg := fmt.Sprintf("missing checksum for %v in manifest file", file)
		return errors.New(msg)
	}
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
