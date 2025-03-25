// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cns

import (
	"bytes"
	"context"
	"crypto/sha1"
	"crypto/tls"
	"errors"
	"fmt"

	"github.com/vmware/govmomi"
	cnstypes "github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/object"
	vim25types "github.com/vmware/govmomi/vim25/types"
)

// DefaultVCenterPort is the default port used to access vCenter.
const DefaultVCenterPort string = "443"

// GetTaskInfo gets the task info given a task
func GetTaskInfo(ctx context.Context, task *object.Task) (*vim25types.TaskInfo, error) {
	taskInfo, err := task.WaitForResult(ctx, nil)
	if err != nil {
		return nil, err
	}
	return taskInfo, nil
}

// GetQuerySnapshotsTaskResult gets the task result of QuerySnapshots given a task info
func GetQuerySnapshotsTaskResult(ctx context.Context, taskInfo *vim25types.TaskInfo) (*cnstypes.CnsSnapshotQueryResult, error) {
	if taskInfo == nil {
		return nil, errors.New("TaskInfo is empty")
	}
	if taskInfo.Result != nil {
		snapshotQueryResult := taskInfo.Result.(cnstypes.CnsSnapshotQueryResult)
		return &snapshotQueryResult, nil
	}
	return nil, errors.New("TaskInfo result is empty")
}

// GetTaskResult gets the task result given a task info
func GetTaskResult(ctx context.Context, taskInfo *vim25types.TaskInfo) (cnstypes.BaseCnsVolumeOperationResult, error) {
	if taskInfo == nil {
		return nil, errors.New("TaskInfo is empty")
	}
	if taskInfo.Result != nil {
		volumeOperationBatchResult := taskInfo.Result.(cnstypes.CnsVolumeOperationBatchResult)
		if len(volumeOperationBatchResult.VolumeResults) == 0 {
			return nil, errors.New("Cannot get VolumeOperationResult")
		}
		return volumeOperationBatchResult.VolumeResults[0], nil
	}
	return nil, errors.New("TaskInfo result is empty")
}

// GetTaskResultArray gets the task result array for a specified task info
func GetTaskResultArray(ctx context.Context, taskInfo *vim25types.TaskInfo) ([]cnstypes.BaseCnsVolumeOperationResult, error) {
	if taskInfo == nil {
		return nil, errors.New("TaskInfo is empty")
	}
	if taskInfo.Result != nil {
		volumeOperationBatchResult := taskInfo.Result.(cnstypes.CnsVolumeOperationBatchResult)
		if len(volumeOperationBatchResult.VolumeResults) == 0 {
			return nil, errors.New("Cannot get VolumeOperationResult")
		}
		return volumeOperationBatchResult.VolumeResults, nil
	}
	return nil, errors.New("TaskInfo result is empty")
}

// dropUnknownCreateSpecElements helps drop newly added elements in the CnsVolumeCreateSpec, which are not known to the prior vSphere releases
func dropUnknownCreateSpecElements(c *Client, createSpecList []cnstypes.CnsVolumeCreateSpec) []cnstypes.CnsVolumeCreateSpec {
	updatedcreateSpecList := make([]cnstypes.CnsVolumeCreateSpec, 0, len(createSpecList))
	switch c.Version {
	case ReleaseVSAN67u3:
		// Dropping optional fields not known to vSAN 6.7U3
		for _, createSpec := range createSpecList {
			createSpec.Metadata.ContainerCluster.ClusterFlavor = ""
			createSpec.Metadata.ContainerCluster.ClusterDistribution = ""
			createSpec.Metadata.ContainerClusterArray = nil
			var updatedEntityMetadata []cnstypes.BaseCnsEntityMetadata
			for _, entityMetadata := range createSpec.Metadata.EntityMetadata {
				k8sEntityMetadata := any(entityMetadata).(*cnstypes.CnsKubernetesEntityMetadata)
				k8sEntityMetadata.ClusterID = ""
				k8sEntityMetadata.ReferredEntity = nil
				updatedEntityMetadata = append(updatedEntityMetadata, cnstypes.BaseCnsEntityMetadata(k8sEntityMetadata))
			}
			createSpec.Metadata.EntityMetadata = updatedEntityMetadata
			_, ok := createSpec.BackingObjectDetails.(*cnstypes.CnsBlockBackingDetails)
			if ok {
				createSpec.BackingObjectDetails.(*cnstypes.CnsBlockBackingDetails).BackingDiskUrlPath = ""
			}
			updatedcreateSpecList = append(updatedcreateSpecList, createSpec)
		}
		createSpecList = updatedcreateSpecList
	case ReleaseVSAN70:
		// Dropping optional fields not known to vSAN 7.0
		for _, createSpec := range createSpecList {
			createSpec.Metadata.ContainerCluster.ClusterDistribution = ""
			var updatedContainerClusterArray []cnstypes.CnsContainerCluster
			for _, containerCluster := range createSpec.Metadata.ContainerClusterArray {
				containerCluster.ClusterDistribution = ""
				updatedContainerClusterArray = append(updatedContainerClusterArray, containerCluster)
			}
			createSpec.Metadata.ContainerClusterArray = updatedContainerClusterArray
			_, ok := createSpec.BackingObjectDetails.(*cnstypes.CnsBlockBackingDetails)
			if ok {
				createSpec.BackingObjectDetails.(*cnstypes.CnsBlockBackingDetails).BackingDiskUrlPath = ""
			}
			updatedcreateSpecList = append(updatedcreateSpecList, createSpec)
		}
		createSpecList = updatedcreateSpecList
	case ReleaseVSAN70u1:
		// Dropping optional fields not known to vSAN 7.0U1
		for _, createSpec := range createSpecList {
			createSpec.Metadata.ContainerCluster.ClusterDistribution = ""
			var updatedContainerClusterArray []cnstypes.CnsContainerCluster
			for _, containerCluster := range createSpec.Metadata.ContainerClusterArray {
				containerCluster.ClusterDistribution = ""
				updatedContainerClusterArray = append(updatedContainerClusterArray, containerCluster)
			}
			createSpec.Metadata.ContainerClusterArray = updatedContainerClusterArray
			updatedcreateSpecList = append(updatedcreateSpecList, createSpec)
		}
		createSpecList = updatedcreateSpecList
	}
	return createSpecList
}

// dropUnknownVolumeMetadataUpdateSpecElements helps drop newly added elements in the CnsVolumeMetadataUpdateSpec, which are not known to the prior vSphere releases
func dropUnknownVolumeMetadataUpdateSpecElements(c *Client, updateSpecList []cnstypes.CnsVolumeMetadataUpdateSpec) []cnstypes.CnsVolumeMetadataUpdateSpec {
	// Dropping optional fields not known to vSAN 6.7U3
	if c.Version == ReleaseVSAN67u3 {
		updatedUpdateSpecList := make([]cnstypes.CnsVolumeMetadataUpdateSpec, 0, len(updateSpecList))
		for _, updateSpec := range updateSpecList {
			updateSpec.Metadata.ContainerCluster.ClusterFlavor = ""
			updateSpec.Metadata.ContainerCluster.ClusterDistribution = ""
			var updatedEntityMetadata []cnstypes.BaseCnsEntityMetadata
			for _, entityMetadata := range updateSpec.Metadata.EntityMetadata {
				k8sEntityMetadata := any(entityMetadata).(*cnstypes.CnsKubernetesEntityMetadata)
				k8sEntityMetadata.ClusterID = ""
				k8sEntityMetadata.ReferredEntity = nil
				updatedEntityMetadata = append(updatedEntityMetadata, cnstypes.BaseCnsEntityMetadata(k8sEntityMetadata))
			}
			updateSpec.Metadata.ContainerClusterArray = nil
			updateSpec.Metadata.EntityMetadata = updatedEntityMetadata
			updatedUpdateSpecList = append(updatedUpdateSpecList, updateSpec)
		}
		updateSpecList = updatedUpdateSpecList
	} else if c.Version == ReleaseVSAN70 || c.Version == ReleaseVSAN70u1 {
		updatedUpdateSpecList := make([]cnstypes.CnsVolumeMetadataUpdateSpec, 0, len(updateSpecList))
		for _, updateSpec := range updateSpecList {
			updateSpec.Metadata.ContainerCluster.ClusterDistribution = ""
			var updatedContainerClusterArray []cnstypes.CnsContainerCluster
			for _, containerCluster := range updateSpec.Metadata.ContainerClusterArray {
				containerCluster.ClusterDistribution = ""
				updatedContainerClusterArray = append(updatedContainerClusterArray, containerCluster)
			}
			updateSpec.Metadata.ContainerClusterArray = updatedContainerClusterArray
			updatedUpdateSpecList = append(updatedUpdateSpecList, updateSpec)
		}
		updateSpecList = updatedUpdateSpecList
	}
	return updateSpecList
}

// GetServiceLocatorInstance takes as input VC userName, VC password, VC client
// and returns a service locator instance for the VC.
func GetServiceLocatorInstance(ctx context.Context, userName string, password string, vcClient *govmomi.Client) (*vim25types.ServiceLocator, error) {
	hostPortStr := fmt.Sprintf("%s:%s", vcClient.URL().Hostname(), DefaultVCenterPort)
	url := fmt.Sprintf("https://%s/sdk", hostPortStr)

	thumbprint, err := getSslThumbprint(ctx, hostPortStr, vcClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get ssl thumbprint. Error: %+v", err)
	}

	serviceLocatorInstance := &vim25types.ServiceLocator{
		InstanceUuid: vcClient.ServiceContent.About.InstanceUuid,
		Url:          url,
		Credential: &vim25types.ServiceLocatorNamePassword{
			Username: userName,
			Password: password,
		},
		SslThumbprint: thumbprint,
	}

	return serviceLocatorInstance, nil
}

// getSslThumbprint connects to the given network address, initiates a TLS handshake
// and retrieves the SSL thumprint from the resulting TLS connection.
func getSslThumbprint(ctx context.Context, addr string, vcClient *govmomi.Client) (string, error) {
	conn, err := vcClient.Client.DefaultTransport().DialTLSContext(ctx, "tcp", addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	tlsConn, ok := conn.(*tls.Conn)
	if !ok {
		return "", fmt.Errorf("cannot convert net connection to tls connection")
	}

	cert := tlsConn.ConnectionState().PeerCertificates[0]
	thumbPrint := sha1.Sum(cert.Raw)

	// Get hex representation for each byte of the thumbprint separated by colon.
	// e.g. B9:12:79:B9:36:1B:B5:C1:2F:20:4A:DD:BD:0C:3D:31:82:99:CB:5C
	var buf bytes.Buffer
	for i, f := range thumbPrint {
		if i > 0 {
			fmt.Fprintf(&buf, ":")
		}
		fmt.Fprintf(&buf, "%02X", f)
	}

	return buf.String(), nil
}
