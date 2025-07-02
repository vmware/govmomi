// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library_test

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"

	_ "github.com/vmware/govmomi/vapi/simulator"
)

func ExampleManager_CreateLibrary() {
	simulator.Run(func(ctx context.Context, vc *vim25.Client) error {
		c := rest.NewClient(vc)

		err := c.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			return err
		}

		ds, err := find.NewFinder(vc).DefaultDatastore(ctx)
		if err != nil {
			return err
		}

		m := library.NewManager(c)

		id, err := m.CreateLibrary(ctx, library.Library{
			Name: "example",
			Type: "LOCAL",
			Storage: []library.StorageBacking{{
				DatastoreID: ds.Reference().Value,
				Type:        "DATASTORE",
			}},
		})
		if err != nil {
			return err
		}

		l, err := m.GetLibraryByID(ctx, id)
		if err != nil {
			return err
		}

		fmt.Println("created library", l.Name)
		return nil
	})
	// Output: created library example
}

func ExampleManager_CreateLibrary_subscribed() {
	simulator.Run(func(ctx context.Context, vc *vim25.Client) error {
		c := rest.NewClient(vc)

		err := c.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			return err
		}

		ds, err := find.NewFinder(vc).DefaultDatastore(ctx)
		if err != nil {
			return err
		}

		m := library.NewManager(c)

		pubLibID, err := m.CreateLibrary(ctx, library.Library{
			Name: "my-pub-lib",
			Type: "LOCAL",
			Storage: []library.StorageBacking{
				{
					DatastoreID: ds.Reference().Value,
					Type:        "DATASTORE",
				},
			},
			Publication: &library.Publication{
				Published: types.New(true),
			},
		})
		if err != nil {
			return err
		}

		pubLib, err := m.GetLibraryByID(ctx, pubLibID)
		if err != nil {
			return err
		}

		fmt.Println("created library", pubLib.Name)

		// Upload an OVA.
		pubItemID, err := m.CreateLibraryItem(ctx, library.Item{
			Name:      "my-image",
			Type:      "OVF",
			LibraryID: pubLib.ID,
		})
		if err != nil {
			return err
		}

		pubItem, err := m.GetLibraryItem(ctx, pubItemID)
		if err != nil {
			return err
		}

		fmt.Println("  created library item", pubItem.Name)

		uploadSessionID, err := m.CreateLibraryItemUpdateSession(
			ctx,
			library.Session{
				LibraryItemID: pubItemID,
			})

		uploadFn := func(path string) error {
			f, err := os.Open(filepath.Clean(path))
			if err != nil {
				return err
			}
			defer f.Close()

			fi, err := f.Stat()
			if err != nil {
				return err
			}

			info := library.UpdateFile{
				Name:       filepath.Base(path),
				SourceType: "PUSH",
				Size:       fi.Size(),
			}

			update, err := m.AddLibraryItemFile(ctx, uploadSessionID, info)
			if err != nil {
				return err
			}

			u, err := url.Parse(update.UploadEndpoint.URI)
			if err != nil {
				return err
			}

			p := soap.DefaultUpload
			p.ContentLength = info.Size

			return m.Client.Upload(ctx, f, u, &p)
		}

		if err := uploadFn("./testdata/ttylinux-pc_i486-16.1.ova"); err != nil {
			return err
		}

		if err := m.CompleteLibraryItemUpdateSession(
			ctx, uploadSessionID); err != nil {

			return err
		}

		pubItemStor, err := m.ListLibraryItemStorage(ctx, pubItemID)
		if err != nil {
			return err
		}

		for i := range pubItemStor {
			is := pubItemStor[i]
			fmt.Printf(
				"    uploaded library item file %s, cached=%v, size=%d\n",
				is.Name, is.Cached, is.Size)
		}

		// Create a subscribed library that points to the one above.
		subLibID, err := m.CreateLibrary(ctx, library.Library{
			Name: "my-sub-lib",
			Type: "SUBSCRIBED",
			Storage: []library.StorageBacking{
				{
					DatastoreID: ds.Reference().Value,
					Type:        "DATASTORE",
				},
			},
			Subscription: &library.Subscription{
				SubscriptionURL: pubLib.Publication.PublishURL,
				OnDemand:        types.New(false),
			},
		})
		if err != nil {
			return err
		}

		subLib, err := m.GetLibraryByID(ctx, subLibID)
		if err != nil {
			return err
		}

		fmt.Println("created library", subLib.Name)

		subItemIDs, err := m.ListLibraryItems(ctx, subLibID)
		if err != nil {
			return err
		}

		for i := range subItemIDs {
			subItemID := subItemIDs[i]

			subItem, err := m.GetLibraryItem(ctx, subItemID)
			if err != nil {
				return err
			}

			fmt.Println("  got subscribed library item", subItem.Name)

			subItemStor, err := m.ListLibraryItemStorage(ctx, subItemID)
			if err != nil {
				return err
			}
			for i := range subItemStor {
				is := subItemStor[i]
				fmt.Printf(
					"    library item file %s, cached=%v, size=%d\n",
					is.Name, is.Cached, is.Size)
			}
		}

		return nil
	})

	// Output:
	// created library my-pub-lib
	//   created library item my-image
	//     uploaded library item file ttylinux-pc_i486-16.1.ovf, cached=true, size=5005
	//     uploaded library item file ttylinux-pc_i486-16.1.mf, cached=true, size=159
	//     uploaded library item file ttylinux-pc_i486-16.1-disk1.vmdk, cached=true, size=1047552
	// created library my-sub-lib
	//   got subscribed library item my-image
	//     library item file ttylinux-pc_i486-16.1.ovf, cached=true, size=5005
	//     library item file ttylinux-pc_i486-16.1.mf, cached=true, size=159
	//     library item file ttylinux-pc_i486-16.1-disk1.vmdk, cached=true, size=1047552
}

func ExampleManager_CreateLibrary_subscribed_ondemand() {
	simulator.Run(func(ctx context.Context, vc *vim25.Client) error {
		c := rest.NewClient(vc)

		err := c.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			return err
		}

		ds, err := find.NewFinder(vc).DefaultDatastore(ctx)
		if err != nil {
			return err
		}

		m := library.NewManager(c)

		pubLibID, err := m.CreateLibrary(ctx, library.Library{
			Name: "my-pub-lib",
			Type: "LOCAL",
			Storage: []library.StorageBacking{
				{
					DatastoreID: ds.Reference().Value,
					Type:        "DATASTORE",
				},
			},
			Publication: &library.Publication{
				Published: types.New(true),
			},
		})
		if err != nil {
			return err
		}

		pubLib, err := m.GetLibraryByID(ctx, pubLibID)
		if err != nil {
			return err
		}

		fmt.Println("created library", pubLib.Name)

		// Upload an OVA.
		pubItemID, err := m.CreateLibraryItem(ctx, library.Item{
			Name:      "my-image",
			Type:      "OVF",
			LibraryID: pubLib.ID,
		})
		if err != nil {
			return err
		}

		pubItem, err := m.GetLibraryItem(ctx, pubItemID)
		if err != nil {
			return err
		}

		fmt.Println("  created library item", pubItem.Name)

		uploadSessionID, err := m.CreateLibraryItemUpdateSession(
			ctx,
			library.Session{
				LibraryItemID: pubItemID,
			})

		uploadFn := func(path string) error {
			f, err := os.Open(filepath.Clean(path))
			if err != nil {
				return err
			}
			defer f.Close()

			fi, err := f.Stat()
			if err != nil {
				return err
			}

			info := library.UpdateFile{
				Name:       filepath.Base(path),
				SourceType: "PUSH",
				Size:       fi.Size(),
			}

			update, err := m.AddLibraryItemFile(ctx, uploadSessionID, info)
			if err != nil {
				return err
			}

			u, err := url.Parse(update.UploadEndpoint.URI)
			if err != nil {
				return err
			}

			p := soap.DefaultUpload
			p.ContentLength = info.Size

			return m.Client.Upload(ctx, f, u, &p)
		}

		if err := uploadFn("./testdata/ttylinux-pc_i486-16.1.ova"); err != nil {
			return err
		}

		if err := m.CompleteLibraryItemUpdateSession(
			ctx, uploadSessionID); err != nil {

			return err
		}

		pubItemStor, err := m.ListLibraryItemStorage(ctx, pubItemID)
		if err != nil {
			return err
		}

		for i := range pubItemStor {
			is := pubItemStor[i]
			fmt.Printf(
				"    uploaded library item file %s, cached=%v, size=%d\n",
				is.Name, is.Cached, is.Size)
		}

		// Create a subscribed library that points to the one above.
		subLibID, err := m.CreateLibrary(ctx, library.Library{
			Name: "my-sub-lib",
			Type: "SUBSCRIBED",
			Storage: []library.StorageBacking{
				{
					DatastoreID: ds.Reference().Value,
					Type:        "DATASTORE",
				},
			},
			Subscription: &library.Subscription{
				SubscriptionURL: pubLib.Publication.PublishURL,
				OnDemand:        types.New(true),
			},
		})
		if err != nil {
			return err
		}

		subLib, err := m.GetLibraryByID(ctx, subLibID)
		if err != nil {
			return err
		}

		fmt.Println("created library", subLib.Name)

		subItemIDs, err := m.ListLibraryItems(ctx, subLibID)
		if err != nil {
			return err
		}

		for i := range subItemIDs {
			subItemID := subItemIDs[i]

			subItem, err := m.GetLibraryItem(ctx, subItemID)
			if err != nil {
				return err
			}

			fmt.Println("  got subscribed library item", subItem.Name)

			// List the item's storage prior to being synced.
			subItemStor, err := m.ListLibraryItemStorage(ctx, subItemID)
			if err != nil {
				return err
			}
			for i := range subItemStor {
				is := subItemStor[i]
				fmt.Printf(
					"    library item file %s, cached=%v, size=%d\n",
					is.Name, is.Cached, is.Size)
			}

			// Synchronize the item.
			if err := m.SyncLibraryItem(ctx, subItem, true); err != nil {
				return err
			}

			fmt.Println("  sync'd (force=true) subscribed library item", subItem.Name)

			// List the item's storage after being synced.
			subItemStor, err = m.ListLibraryItemStorage(ctx, subItemID)
			if err != nil {
				return err
			}
			for i := range subItemStor {
				is := subItemStor[i]
				fmt.Printf(
					"    library item file %s, cached=%v, size=%d\n",
					is.Name, is.Cached, is.Size)
			}
		}

		return nil
	})

	// Output:
	// created library my-pub-lib
	//   created library item my-image
	//     uploaded library item file ttylinux-pc_i486-16.1.ovf, cached=true, size=5005
	//     uploaded library item file ttylinux-pc_i486-16.1.mf, cached=true, size=159
	//     uploaded library item file ttylinux-pc_i486-16.1-disk1.vmdk, cached=true, size=1047552
	// created library my-sub-lib
	//   got subscribed library item my-image
	//     library item file ttylinux-pc_i486-16.1.ovf, cached=true, size=5005
	//     library item file ttylinux-pc_i486-16.1.mf, cached=true, size=159
	//     library item file ttylinux-pc_i486-16.1-disk1.vmdk, cached=false, size=0
	//   sync'd (force=true) subscribed library item my-image
	//     library item file ttylinux-pc_i486-16.1.ovf, cached=true, size=5005
	//     library item file ttylinux-pc_i486-16.1.mf, cached=true, size=159
	//     library item file ttylinux-pc_i486-16.1-disk1.vmdk, cached=true, size=1047552
}
