// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package crypto_test

import (
	"context"
	"math"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi/crypto"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

// CryptoManagerKmip should implement the Reference interface.
var _ object.Reference = crypto.ManagerKmip{}

func TestCryptoManagerKmip(t *testing.T) {

	t.Run("RegisterKmipCluster", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			providerID := uuid.NewString()

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID,
				types.KmipClusterInfoKmsManagementTypeUnknown))

			isValid, err := m.IsValidProvider(ctx, providerID)
			assert.NoError(t, err)
			assert.True(t, isValid)

			err = m.RegisterKmsCluster(
				ctx,
				providerID,
				types.KmipClusterInfoKmsManagementTypeUnknown)
			assert.EqualError(t, err, "ServerFaultCode: Already registered")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))
		})
	})

	t.Run("GetClusterStatus", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			providerID := uuid.NewString()

			status, err := m.GetClusterStatus(ctx, providerID)
			assert.EqualError(t, err, "invalid cluster ID")
			assert.Nil(t, status)

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID,
				types.KmipClusterInfoKmsManagementTypeUnknown))

			status, err = m.GetClusterStatus(ctx, providerID)
			assert.NoError(t, err)
			assert.NotNil(t, status)
			assert.Equal(t, providerID, status.ClusterId.Id)
		})
	})

	t.Run("UnregisterKmsCluster", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			providerID := uuid.NewString()

			err = m.UnregisterKmsCluster(ctx, providerID)
			assert.EqualError(t, err, "ServerFaultCode: Invalid cluster ID")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID,
				types.KmipClusterInfoKmsManagementTypeUnknown))

			isValid, err := m.IsValidProvider(ctx, providerID)
			assert.NoError(t, err)
			assert.True(t, isValid)

			assert.NoError(t, m.UnregisterKmsCluster(ctx, providerID))

			isValid, err = m.IsValidProvider(ctx, providerID)
			assert.NoError(t, err)
			assert.False(t, isValid)
		})
	})

	t.Run("IsValidProvider", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			providerID := uuid.NewString()

			ok, err := m.IsValidProvider(ctx, providerID)
			assert.NoError(t, err)
			assert.False(t, ok)

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID,
				types.KmipClusterInfoKmsManagementTypeUnknown))

			ok, err = m.IsValidProvider(ctx, providerID)
			assert.NoError(t, err)
			assert.True(t, ok)
		})
	})

	t.Run("GetDefaultKmsClusterID", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			provider1ID := uuid.NewString()
			provider2ID := uuid.NewString()
			provider3ID := uuid.NewString()

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider1ID,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider2ID,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider3ID,
				types.KmipClusterInfoKmsManagementTypeNativeProvider))

			defaultProviderID, err := m.GetDefaultKmsClusterID(ctx, nil, true)
			assert.EqualError(t, err, "ServerFaultCode: No default provider")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))
			assert.Empty(t, defaultProviderID)

			assert.NoError(t, m.MarkDefault(ctx, provider3ID))
			defaultProviderID, err = m.GetDefaultKmsClusterID(ctx, nil, true)
			assert.NoError(t, err)
			assert.Equal(t, provider3ID, defaultProviderID)

			// Assert setting the default a second time does not return an
			// error.
			assert.NoError(t, m.MarkDefault(ctx, provider3ID))
			defaultProviderID, err = m.GetDefaultKmsClusterID(ctx, nil, true)
			assert.NoError(t, err)
			assert.Equal(t, provider3ID, defaultProviderID)

			fakeMoRef := types.ManagedObjectReference{
				Type:  "fake",
				Value: "fake",
			}

			defaultProviderID, err = m.GetDefaultKmsClusterID(ctx, &fakeMoRef, true)
			assert.EqualError(t, err, "ServerFaultCode: No default provider")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))
			assert.Empty(t, defaultProviderID)

			assert.NoError(t, m.SetDefaultKmsClusterId(
				ctx, provider2ID, &fakeMoRef))
			defaultProviderID, err = m.GetDefaultKmsClusterID(ctx, &fakeMoRef, true)
			assert.NoError(t, err)
			assert.Equal(t, provider2ID, defaultProviderID)

			// Assert setting the default for an entity a second time does not
			// return an error.
			assert.NoError(t, m.SetDefaultKmsClusterId(
				ctx, provider2ID, &fakeMoRef))
			defaultProviderID, err = m.GetDefaultKmsClusterID(ctx, &fakeMoRef, true)
			assert.NoError(t, err)
			assert.Equal(t, provider2ID, defaultProviderID)

			// Remove the default for the entity.
			assert.NoError(t, m.SetDefaultKmsClusterId(ctx, "", &fakeMoRef))
			defaultProviderID, err = m.GetDefaultKmsClusterID(ctx, &fakeMoRef, true)
			assert.EqualError(t, err, "ServerFaultCode: No default provider")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))
			assert.Empty(t, defaultProviderID)

			// Remove the default.
			assert.NoError(t, m.SetDefaultKmsClusterId(ctx, "", nil))
			assert.EqualError(t, err, "ServerFaultCode: No default provider")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))
			assert.Empty(t, defaultProviderID)
		})
	})

	t.Run("RegisterKmipServer", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			providerID := uuid.NewString()
			serverName := uuid.NewString()

			serverSpec := types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: providerID,
				},
				Info: types.KmipServerInfo{
					Name: serverName,
				},
			}

			err = m.RegisterKmsCluster(
				ctx,
				providerID,
				types.KmipClusterInfoKmsManagementTypeVCenter)
			assert.True(t, fault.Is(err, &types.InvalidArgument{}))

			assert.NoError(t, m.RegisterKmipServer(ctx, serverSpec))

			ok, err := m.IsValidServer(ctx, providerID, serverName)
			assert.NoError(t, err)
			assert.True(t, ok)

			err = m.RegisterKmipServer(ctx, serverSpec)
			assert.EqualError(t, err, "ServerFaultCode: Already registered")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))
		})
	})

	t.Run("GetServerStatus", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			providerID := uuid.NewString()
			serverName := uuid.NewString()

			status, err := m.GetServerStatus(ctx, providerID, serverName)
			assert.EqualError(t, err, "invalid cluster ID")
			assert.Nil(t, status)

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID,
				types.KmipClusterInfoKmsManagementTypeUnknown))

			status, err = m.GetServerStatus(ctx, providerID, serverName)
			assert.EqualError(t, err, "invalid server name")
			assert.Nil(t, status)

			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: providerID,
				},
				Info: types.KmipServerInfo{
					Name: serverName,
				},
			}))

			status, err = m.GetServerStatus(ctx, providerID, serverName)
			assert.NoError(t, err)
			assert.NotNil(t, status)
			assert.Equal(t, serverName, status.Name)
		})
	})

	t.Run("ListKmipServers", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			clusters, err := m.ListKmipServers(ctx, nil)
			assert.NoError(t, err)
			assert.Len(t, clusters, 0)

			provider1ID := uuid.NewString()
			provider2ID := uuid.NewString()
			provider3ID := uuid.NewString()

			provider1serverName1 := uuid.NewString()
			provider1serverName2 := uuid.NewString()
			provider2serverName1 := uuid.NewString()
			provider2serverName2 := uuid.NewString()
			provider2serverName3 := uuid.NewString()
			provider3serverName1 := uuid.NewString()

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider1ID,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider2ID,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider3ID,
				types.KmipClusterInfoKmsManagementTypeNativeProvider))

			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: provider1ID,
				},
				Info: types.KmipServerInfo{
					Name: provider1serverName1,
				},
			}))
			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: provider1ID,
				},
				Info: types.KmipServerInfo{
					Name: provider1serverName2,
				},
			}))
			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: provider2ID,
				},
				Info: types.KmipServerInfo{
					Name: provider2serverName1,
				},
			}))
			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: provider2ID,
				},
				Info: types.KmipServerInfo{
					Name: provider2serverName2,
				},
			}))
			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: provider2ID,
				},
				Info: types.KmipServerInfo{
					Name: provider2serverName3,
				},
			}))
			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: provider3ID,
				},
				Info: types.KmipServerInfo{
					Name: provider3serverName1,
				},
			}))

			clusters, err = m.ListKmipServers(ctx, nil)
			assert.NoError(t, err)
			assert.Len(t, clusters, 3)

			assert.Equal(t, clusters[0].ClusterId.Id, provider1ID)
			assert.Equal(t, clusters[0].ManagementType, string(types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.Len(t, clusters[0].Servers, 2)
			assert.Equal(t, clusters[0].Servers[0].Name, provider1serverName1)
			assert.Equal(t, clusters[0].Servers[1].Name, provider1serverName2)

			assert.Equal(t, clusters[1].ClusterId.Id, provider2ID)
			assert.Equal(t, clusters[1].ManagementType, string(types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.Len(t, clusters[1].Servers, 3)
			assert.Equal(t, clusters[1].Servers[0].Name, provider2serverName1)
			assert.Equal(t, clusters[1].Servers[1].Name, provider2serverName2)
			assert.Equal(t, clusters[1].Servers[2].Name, provider2serverName3)

			assert.Equal(t, clusters[2].ClusterId.Id, provider3ID)
			assert.Equal(t, clusters[2].ManagementType, string(types.KmipClusterInfoKmsManagementTypeNativeProvider))
			assert.Len(t, clusters[2].Servers, 1)
			assert.Equal(t, clusters[2].Servers[0].Name, provider3serverName1)

			// List all with a limit.
			clusters, err = m.ListKmipServers(ctx, types.NewInt32(math.MaxInt32))
			assert.NoError(t, err)
			assert.Len(t, clusters, 3)

			assert.Equal(t, clusters[0].ClusterId.Id, provider1ID)
			assert.Equal(t, clusters[0].ManagementType, string(types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.Len(t, clusters[0].Servers, 2)
			assert.Equal(t, clusters[0].Servers[0].Name, provider1serverName1)
			assert.Equal(t, clusters[0].Servers[1].Name, provider1serverName2)

			assert.Equal(t, clusters[1].ClusterId.Id, provider2ID)
			assert.Equal(t, clusters[1].ManagementType, string(types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.Len(t, clusters[1].Servers, 3)
			assert.Equal(t, clusters[1].Servers[0].Name, provider2serverName1)
			assert.Equal(t, clusters[1].Servers[1].Name, provider2serverName2)
			assert.Equal(t, clusters[1].Servers[2].Name, provider2serverName3)

			assert.Equal(t, clusters[2].ClusterId.Id, provider3ID)
			assert.Equal(t, clusters[2].ManagementType, string(types.KmipClusterInfoKmsManagementTypeNativeProvider))
			assert.Len(t, clusters[2].Servers, 1)
			assert.Equal(t, clusters[2].Servers[0].Name, provider3serverName1)

			// List the first cluster.
			clusters, err = m.ListKmipServers(ctx, types.NewInt32(1))
			assert.NoError(t, err)
			assert.Len(t, clusters, 1)

			assert.Equal(t, clusters[0].ClusterId.Id, provider1ID)
			assert.Equal(t, clusters[0].ManagementType, string(types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.Len(t, clusters[0].Servers, 2)
			assert.Equal(t, clusters[0].Servers[0].Name, provider1serverName1)
			assert.Equal(t, clusters[0].Servers[1].Name, provider1serverName2)

			// List the first and second cluster.
			clusters, err = m.ListKmipServers(ctx, types.NewInt32(2))
			assert.NoError(t, err)
			assert.Len(t, clusters, 2)

			assert.Equal(t, clusters[0].ClusterId.Id, provider1ID)
			assert.Equal(t, clusters[0].ManagementType, string(types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.Len(t, clusters[0].Servers, 2)
			assert.Equal(t, clusters[0].Servers[0].Name, provider1serverName1)
			assert.Equal(t, clusters[0].Servers[1].Name, provider1serverName2)

			assert.Equal(t, clusters[1].ClusterId.Id, provider2ID)
			assert.Equal(t, clusters[1].ManagementType, string(types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.Len(t, clusters[1].Servers, 3)
			assert.Equal(t, clusters[1].Servers[0].Name, provider2serverName1)
			assert.Equal(t, clusters[1].Servers[1].Name, provider2serverName2)
			assert.Equal(t, clusters[1].Servers[2].Name, provider2serverName3)
		})
	})

	t.Run("UpdateKmipServer", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			providerID := uuid.NewString()
			serverName := uuid.NewString()

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID,
				types.KmipClusterInfoKmsManagementTypeUnknown))

			spec := types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: providerID,
				},
				Info: types.KmipServerInfo{
					Name: serverName,
				},
			}

			assert.NoError(t, m.RegisterKmipServer(ctx, spec))

			ok, err := m.IsValidServer(ctx, providerID, serverName)
			assert.NoError(t, err)
			assert.True(t, ok)

			spec.ClusterId.Id = "invalid"
			spec.Info.Name = "invalid"
			spec.Info.Port = 123

			err = m.UpdateKmipServer(ctx, spec)
			assert.EqualError(t, err, "ServerFaultCode: Invalid cluster ID")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))

			clusters, err := m.ListKmipServers(ctx, nil)
			assert.NoError(t, err)
			assert.Len(t, clusters, 1)
			assert.Len(t, clusters[0].Servers, 1)
			assert.Equal(t, int32(0), clusters[0].Servers[0].Port)

			spec.ClusterId.Id = providerID

			err = m.UpdateKmipServer(ctx, spec)
			assert.EqualError(t, err, "ServerFaultCode: Invalid server name")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))

			clusters, err = m.ListKmipServers(ctx, nil)
			assert.NoError(t, err)
			assert.Len(t, clusters, 1)
			assert.Len(t, clusters[0].Servers, 1)
			assert.Equal(t, int32(0), clusters[0].Servers[0].Port)

			spec.Info.Name = serverName

			assert.NoError(t, m.UpdateKmipServer(ctx, spec))

			clusters, err = m.ListKmipServers(ctx, nil)
			assert.NoError(t, err)
			assert.Len(t, clusters, 1)
			assert.Len(t, clusters[0].Servers, 1)
			assert.Equal(t, int32(123), clusters[0].Servers[0].Port)
		})
	})

	t.Run("RemoveKmipServer", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			providerID := uuid.NewString()
			serverName := uuid.NewString()

			err = m.RemoveKmipServer(ctx, providerID, serverName)
			assert.EqualError(t, err, "ServerFaultCode: Invalid cluster ID")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID,
				types.KmipClusterInfoKmsManagementTypeUnknown))

			err = m.RemoveKmipServer(ctx, providerID, serverName)
			assert.EqualError(t, err, "ServerFaultCode: Invalid server name")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))

			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: providerID,
				},
				Info: types.KmipServerInfo{
					Name: serverName,
				},
			}))

			ok, err := m.IsValidServer(ctx, providerID, serverName)
			assert.NoError(t, err)
			assert.True(t, ok)

			assert.NoError(t, m.RemoveKmipServer(ctx, providerID, serverName))

			ok, err = m.IsValidServer(ctx, providerID, serverName)
			assert.NoError(t, err)
			assert.False(t, ok)
		})
	})

	t.Run("IsValidServer", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			providerID := uuid.NewString()
			serverName := uuid.NewString()

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID,
				types.KmipClusterInfoKmsManagementTypeUnknown))

			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: providerID,
				},
				Info: types.KmipServerInfo{
					Name: serverName,
				},
			}))

			ok, err := m.IsValidServer(ctx, providerID, serverName)
			assert.NoError(t, err)
			assert.True(t, ok)

			assert.NoError(t, m.RemoveKmipServer(ctx, providerID, serverName))

			ok, err = m.IsValidServer(ctx, providerID, serverName)
			assert.NoError(t, err)
			assert.False(t, ok)
		})
	})

	t.Run("GetStatus", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			status, err := m.GetStatus(ctx)
			assert.NoError(t, err)
			assert.Nil(t, status)

			provider1ID := uuid.NewString()
			provider2ID := uuid.NewString()
			provider3ID := uuid.NewString()

			provider1serverName1 := uuid.NewString()
			provider1serverName2 := uuid.NewString()
			provider2serverName1 := uuid.NewString()
			provider2serverName2 := uuid.NewString()
			provider2serverName3 := uuid.NewString()
			provider3serverName1 := uuid.NewString()

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider1ID,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider2ID,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider3ID,
				types.KmipClusterInfoKmsManagementTypeNativeProvider))

			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: provider1ID,
				},
				Info: types.KmipServerInfo{
					Name: provider1serverName1,
				},
			}))
			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: provider1ID,
				},
				Info: types.KmipServerInfo{
					Name: provider1serverName2,
				},
			}))
			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: provider2ID,
				},
				Info: types.KmipServerInfo{
					Name: provider2serverName1,
				},
			}))
			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: provider2ID,
				},
				Info: types.KmipServerInfo{
					Name: provider2serverName2,
				},
			}))
			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: provider2ID,
				},
				Info: types.KmipServerInfo{
					Name: provider2serverName3,
				},
			}))
			assert.NoError(t, m.RegisterKmipServer(ctx, types.KmipServerSpec{
				ClusterId: types.KeyProviderId{
					Id: provider3ID,
				},
				Info: types.KmipServerInfo{
					Name: provider3serverName1,
				},
			}))

			status, err = m.GetStatus(ctx)
			assert.NoError(t, err)
			assert.NotNil(t, status)
			assert.Len(t, status, 3)

			assert.Equal(t, status[0].ClusterId.Id, provider1ID)
			assert.Equal(t, status[0].ManagementType, string(types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.Len(t, status[0].Servers, 2)
			assert.Equal(t, status[0].Servers[0].Name, provider1serverName1)
			assert.Equal(t, status[0].Servers[1].Name, provider1serverName2)

			assert.Equal(t, status[1].ClusterId.Id, provider2ID)
			assert.Equal(t, status[1].ManagementType, string(types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.Len(t, status[1].Servers, 3)
			assert.Equal(t, status[1].Servers[0].Name, provider2serverName1)
			assert.Equal(t, status[1].Servers[1].Name, provider2serverName2)
			assert.Equal(t, status[1].Servers[2].Name, provider2serverName3)

			assert.Equal(t, status[2].ClusterId.Id, provider3ID)
			assert.Equal(t, status[2].ManagementType, string(types.KmipClusterInfoKmsManagementTypeNativeProvider))
			assert.Len(t, status[2].Servers, 1)
			assert.Equal(t, status[2].Servers[0].Name, provider3serverName1)

			status, err = m.GetStatus(ctx, types.KmipClusterInfo{
				ClusterId: types.KeyProviderId{
					Id: provider2ID,
				},
				Servers: []types.KmipServerInfo{
					{
						Name: provider2serverName2,
					},
				},
			})
			assert.NoError(t, err)
			assert.NotNil(t, status)
			assert.Len(t, status, 1)

			assert.Equal(t, status[0].ClusterId.Id, provider2ID)
			assert.Equal(t, status[0].ManagementType, string(types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.Len(t, status[0].Servers, 1)
			assert.Equal(t, status[0].Servers[0].Name, provider2serverName2)
		})
	})

	t.Run("IsDefaultProviderNative", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			provider1ID := uuid.NewString()
			provider2ID := uuid.NewString()
			provider3ID := uuid.NewString()

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider1ID,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider2ID,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider3ID,
				types.KmipClusterInfoKmsManagementTypeNativeProvider))

			ok, err := m.IsDefaultProviderNative(ctx, nil, false)
			assert.EqualError(t, err, "ServerFaultCode: No default provider")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))
			assert.False(t, ok)

			assert.NoError(t, m.MarkDefault(ctx, provider3ID))

			ok, err = m.IsDefaultProviderNative(ctx, nil, false)
			assert.NoError(t, err)
			assert.True(t, ok)
		})
	})

	t.Run("MarkDefault", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			provider1ID := uuid.NewString()
			provider2ID := uuid.NewString()
			provider3ID := uuid.NewString()

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider1ID,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider2ID,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider3ID,
				types.KmipClusterInfoKmsManagementTypeNativeProvider))

			assert.NoError(t, m.MarkDefault(ctx, provider2ID))
			defaultProviderID, err := m.GetDefaultKmsClusterID(ctx, nil, true)
			assert.NoError(t, err)
			assert.Equal(t, provider2ID, defaultProviderID)

			assert.NoError(t, m.MarkDefault(ctx, provider1ID))
			defaultProviderID, err = m.GetDefaultKmsClusterID(ctx, nil, true)
			assert.NoError(t, err)
			assert.Equal(t, provider1ID, defaultProviderID)

			assert.NoError(t, m.MarkDefault(ctx, provider3ID))
			defaultProviderID, err = m.GetDefaultKmsClusterID(ctx, nil, true)
			assert.NoError(t, err)
			assert.Equal(t, provider3ID, defaultProviderID)

			assert.NoError(t, m.MarkDefault(ctx, ""))
			defaultProviderID, err = m.GetDefaultKmsClusterID(ctx, nil, true)
			assert.EqualError(t, err, "ServerFaultCode: No default provider")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))
			assert.Empty(t, defaultProviderID)
		})
	})

	t.Run("SetDefaultKmsClusterId", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			provider1ID := uuid.NewString()
			provider2ID := uuid.NewString()
			provider3ID := uuid.NewString()

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider1ID,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider2ID,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				provider3ID,
				types.KmipClusterInfoKmsManagementTypeNativeProvider))

			assert.NoError(t, m.SetDefaultKmsClusterId(ctx, provider2ID, nil))
			defaultProviderID, err := m.GetDefaultKmsClusterID(ctx, nil, true)
			assert.NoError(t, err)
			assert.Equal(t, provider2ID, defaultProviderID)

			assert.NoError(t, m.SetDefaultKmsClusterId(ctx, provider1ID, nil))
			defaultProviderID, err = m.GetDefaultKmsClusterID(ctx, nil, true)
			assert.NoError(t, err)
			assert.Equal(t, provider1ID, defaultProviderID)

			assert.NoError(t, m.SetDefaultKmsClusterId(ctx, provider3ID, nil))
			defaultProviderID, err = m.GetDefaultKmsClusterID(ctx, nil, true)
			assert.NoError(t, err)
			assert.Equal(t, provider3ID, defaultProviderID)

			err = m.SetDefaultKmsClusterId(ctx, "invalid", nil)
			assert.EqualError(t, err, "ServerFaultCode: Invalid cluster ID")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))

			assert.NoError(t, m.SetDefaultKmsClusterId(ctx, "", nil))
			defaultProviderID, err = m.GetDefaultKmsClusterID(ctx, nil, true)
			assert.EqualError(t, err, "ServerFaultCode: No default provider")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))
			assert.Empty(t, defaultProviderID)
		})
	})

	t.Run("GenerateKey", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			providerID1 := uuid.NewString()
			providerID2 := uuid.NewString()

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID1,
				types.KmipClusterInfoKmsManagementTypeNativeProvider))

			keyID, err := m.GenerateKey(ctx, "")
			assert.EqualError(t, err, "ServerFaultCode: No default provider")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))
			assert.Empty(t, keyID)

			assert.NoError(t, m.MarkDefault(ctx, providerID1))

			keyID, err = m.GenerateKey(ctx, providerID1)
			assert.EqualError(t, err,
				"ServerFaultCode: Cannot generate keys with native key provider")
			assert.True(t, fault.Is(err, &types.RuntimeFault{}))
			assert.Empty(t, keyID)

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID2,
				types.KmipClusterInfoKmsManagementTypeUnknown))

			keyID, err = m.GenerateKey(ctx, providerID2)
			assert.NoError(t, err)
			assert.NotEmpty(t, keyID)

			assert.NoError(t, m.MarkDefault(ctx, providerID2))

			keyID, err = m.GenerateKey(ctx, "")
			assert.NoError(t, err)
			assert.NotEmpty(t, keyID)
		})
	})

	t.Run("ListKeys", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			providerID1 := uuid.NewString()
			providerID2 := uuid.NewString()
			providerID3 := uuid.NewString()

			keys, err := m.ListKeys(ctx, nil)
			assert.NoError(t, err)
			assert.Len(t, keys, 0)

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID1,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID2,
				types.KmipClusterInfoKmsManagementTypeUnknown))
			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID3,
				types.KmipClusterInfoKmsManagementTypeNativeProvider))

			keyID1, err := m.GenerateKey(ctx, providerID2)
			assert.NoError(t, err)
			assert.NotEmpty(t, keyID1)

			assert.NoError(t, m.MarkDefault(ctx, providerID2))
			keyID2, err := m.GenerateKey(ctx, "")
			assert.NoError(t, err)
			assert.NotEmpty(t, keyID2)

			assert.NoError(t, m.MarkDefault(ctx, providerID1))
			keyID3, err := m.GenerateKey(ctx, "")
			assert.NoError(t, err)
			assert.NotEmpty(t, keyID3)

			keys, err = m.ListKeys(ctx, nil)
			assert.NoError(t, err)
			assert.Len(t, keys, 3)
			assert.ElementsMatch(t, keys, []types.CryptoKeyId{
				{
					KeyId:      keyID1,
					ProviderId: &types.KeyProviderId{Id: providerID2},
				},
				{
					KeyId:      keyID2,
					ProviderId: &types.KeyProviderId{Id: providerID2},
				},
				{
					KeyId:      keyID3,
					ProviderId: &types.KeyProviderId{Id: providerID1},
				},
			})

			keys, err = m.ListKeys(ctx, types.NewInt32(math.MaxInt32))
			assert.NoError(t, err)
			assert.Len(t, keys, 3)
			assert.ElementsMatch(t, keys, []types.CryptoKeyId{
				{
					KeyId:      keyID1,
					ProviderId: &types.KeyProviderId{Id: providerID2},
				},
				{
					KeyId:      keyID2,
					ProviderId: &types.KeyProviderId{Id: providerID2},
				},
				{
					KeyId:      keyID3,
					ProviderId: &types.KeyProviderId{Id: providerID1},
				},
			})

			keys, err = m.ListKeys(ctx, types.NewInt32(1))
			assert.NoError(t, err)
			assert.Len(t, keys, 1)

			keys, err = m.ListKeys(ctx, types.NewInt32(2))
			assert.NoError(t, err)
			assert.Len(t, keys, 2)
		})
	})

	t.Run("IsValidKey", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			m, err := crypto.GetManagerKmip(c)
			assert.NoError(t, err)

			providerID := uuid.NewString()

			assert.NoError(t, m.RegisterKmsCluster(
				ctx,
				providerID,
				types.KmipClusterInfoKmsManagementTypeUnknown))

			assert.NoError(t, m.MarkDefault(ctx, providerID))

			keyID, err := m.GenerateKey(ctx, "")
			assert.NoError(t, err)
			assert.NotEmpty(t, keyID)

			ok, err := m.IsValidKey(ctx, providerID, keyID)
			assert.NoError(t, err)
			assert.True(t, ok)
		})
	})
}
