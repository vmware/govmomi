// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package methods

import (
	"context"

	"github.com/vmware/govmomi/sms/types"
	"github.com/vmware/govmomi/vim25/soap"
)

type FailoverReplicationGroup_TaskBody struct {
	Req    *types.FailoverReplicationGroup_Task         `xml:"urn:sms FailoverReplicationGroup_Task,omitempty"`
	Res    *types.FailoverReplicationGroup_TaskResponse `xml:"urn:sms FailoverReplicationGroup_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FailoverReplicationGroup_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func FailoverReplicationGroup_Task(ctx context.Context, r soap.RoundTripper, req *types.FailoverReplicationGroup_Task) (*types.FailoverReplicationGroup_TaskResponse, error) {
	var reqBody, resBody FailoverReplicationGroup_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PrepareFailoverReplicationGroup_TaskBody struct {
	Req    *types.PrepareFailoverReplicationGroup_Task         `xml:"urn:sms PrepareFailoverReplicationGroup_Task,omitempty"`
	Res    *types.PrepareFailoverReplicationGroup_TaskResponse `xml:"urn:sms PrepareFailoverReplicationGroup_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PrepareFailoverReplicationGroup_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PrepareFailoverReplicationGroup_Task(ctx context.Context, r soap.RoundTripper, req *types.PrepareFailoverReplicationGroup_Task) (*types.PrepareFailoverReplicationGroup_TaskResponse, error) {
	var reqBody, resBody PrepareFailoverReplicationGroup_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PromoteReplicationGroup_TaskBody struct {
	Req    *types.PromoteReplicationGroup_Task         `xml:"urn:sms PromoteReplicationGroup_Task,omitempty"`
	Res    *types.PromoteReplicationGroup_TaskResponse `xml:"urn:sms PromoteReplicationGroup_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PromoteReplicationGroup_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PromoteReplicationGroup_Task(ctx context.Context, r soap.RoundTripper, req *types.PromoteReplicationGroup_Task) (*types.PromoteReplicationGroup_TaskResponse, error) {
	var reqBody, resBody PromoteReplicationGroup_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryAboutInfoBody struct {
	Req    *types.QueryAboutInfo         `xml:"urn:sms QueryAboutInfo,omitempty"`
	Res    *types.QueryAboutInfoResponse `xml:"urn:sms QueryAboutInfoResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAboutInfoBody) Fault() *soap.Fault { return b.Fault_ }

func QueryAboutInfo(ctx context.Context, r soap.RoundTripper, req *types.QueryAboutInfo) (*types.QueryAboutInfoResponse, error) {
	var reqBody, resBody QueryAboutInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryActiveAlarmBody struct {
	Req    *types.QueryActiveAlarm         `xml:"urn:sms QueryActiveAlarm,omitempty"`
	Res    *types.QueryActiveAlarmResponse `xml:"urn:sms QueryActiveAlarmResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryActiveAlarmBody) Fault() *soap.Fault { return b.Fault_ }

func QueryActiveAlarm(ctx context.Context, r soap.RoundTripper, req *types.QueryActiveAlarm) (*types.QueryActiveAlarmResponse, error) {
	var reqBody, resBody QueryActiveAlarmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryArrayBody struct {
	Req    *types.QueryArray         `xml:"urn:sms QueryArray,omitempty"`
	Res    *types.QueryArrayResponse `xml:"urn:sms QueryArrayResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryArrayBody) Fault() *soap.Fault { return b.Fault_ }

func QueryArray(ctx context.Context, r soap.RoundTripper, req *types.QueryArray) (*types.QueryArrayResponse, error) {
	var reqBody, resBody QueryArrayBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryArrayAssociatedWithLunBody struct {
	Req    *types.QueryArrayAssociatedWithLun         `xml:"urn:sms QueryArrayAssociatedWithLun,omitempty"`
	Res    *types.QueryArrayAssociatedWithLunResponse `xml:"urn:sms QueryArrayAssociatedWithLunResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryArrayAssociatedWithLunBody) Fault() *soap.Fault { return b.Fault_ }

func QueryArrayAssociatedWithLun(ctx context.Context, r soap.RoundTripper, req *types.QueryArrayAssociatedWithLun) (*types.QueryArrayAssociatedWithLunResponse, error) {
	var reqBody, resBody QueryArrayAssociatedWithLunBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryAssociatedBackingStoragePoolBody struct {
	Req    *types.QueryAssociatedBackingStoragePool         `xml:"urn:sms QueryAssociatedBackingStoragePool,omitempty"`
	Res    *types.QueryAssociatedBackingStoragePoolResponse `xml:"urn:sms QueryAssociatedBackingStoragePoolResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAssociatedBackingStoragePoolBody) Fault() *soap.Fault { return b.Fault_ }

func QueryAssociatedBackingStoragePool(ctx context.Context, r soap.RoundTripper, req *types.QueryAssociatedBackingStoragePool) (*types.QueryAssociatedBackingStoragePoolResponse, error) {
	var reqBody, resBody QueryAssociatedBackingStoragePoolBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDatastoreBackingPoolMappingBody struct {
	Req    *types.QueryDatastoreBackingPoolMapping         `xml:"urn:sms QueryDatastoreBackingPoolMapping,omitempty"`
	Res    *types.QueryDatastoreBackingPoolMappingResponse `xml:"urn:sms QueryDatastoreBackingPoolMappingResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDatastoreBackingPoolMappingBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDatastoreBackingPoolMapping(ctx context.Context, r soap.RoundTripper, req *types.QueryDatastoreBackingPoolMapping) (*types.QueryDatastoreBackingPoolMappingResponse, error) {
	var reqBody, resBody QueryDatastoreBackingPoolMappingBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDatastoreCapabilityBody struct {
	Req    *types.QueryDatastoreCapability         `xml:"urn:sms QueryDatastoreCapability,omitempty"`
	Res    *types.QueryDatastoreCapabilityResponse `xml:"urn:sms QueryDatastoreCapabilityResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDatastoreCapabilityBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDatastoreCapability(ctx context.Context, r soap.RoundTripper, req *types.QueryDatastoreCapability) (*types.QueryDatastoreCapabilityResponse, error) {
	var reqBody, resBody QueryDatastoreCapabilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDrsMigrationCapabilityForPerformanceBody struct {
	Req    *types.QueryDrsMigrationCapabilityForPerformance         `xml:"urn:sms QueryDrsMigrationCapabilityForPerformance,omitempty"`
	Res    *types.QueryDrsMigrationCapabilityForPerformanceResponse `xml:"urn:sms QueryDrsMigrationCapabilityForPerformanceResponse,omitempty"`
	Fault_ *soap.Fault                                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDrsMigrationCapabilityForPerformanceBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDrsMigrationCapabilityForPerformance(ctx context.Context, r soap.RoundTripper, req *types.QueryDrsMigrationCapabilityForPerformance) (*types.QueryDrsMigrationCapabilityForPerformanceResponse, error) {
	var reqBody, resBody QueryDrsMigrationCapabilityForPerformanceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDrsMigrationCapabilityForPerformanceExBody struct {
	Req    *types.QueryDrsMigrationCapabilityForPerformanceEx         `xml:"urn:sms QueryDrsMigrationCapabilityForPerformanceEx,omitempty"`
	Res    *types.QueryDrsMigrationCapabilityForPerformanceExResponse `xml:"urn:sms QueryDrsMigrationCapabilityForPerformanceExResponse,omitempty"`
	Fault_ *soap.Fault                                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDrsMigrationCapabilityForPerformanceExBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDrsMigrationCapabilityForPerformanceEx(ctx context.Context, r soap.RoundTripper, req *types.QueryDrsMigrationCapabilityForPerformanceEx) (*types.QueryDrsMigrationCapabilityForPerformanceExResponse, error) {
	var reqBody, resBody QueryDrsMigrationCapabilityForPerformanceExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryFaultDomainBody struct {
	Req    *types.QueryFaultDomain         `xml:"urn:sms QueryFaultDomain,omitempty"`
	Res    *types.QueryFaultDomainResponse `xml:"urn:sms QueryFaultDomainResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryFaultDomainBody) Fault() *soap.Fault { return b.Fault_ }

func QueryFaultDomain(ctx context.Context, r soap.RoundTripper, req *types.QueryFaultDomain) (*types.QueryFaultDomainResponse, error) {
	var reqBody, resBody QueryFaultDomainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryFileSystemAssociatedWithArrayBody struct {
	Req    *types.QueryFileSystemAssociatedWithArray         `xml:"urn:sms QueryFileSystemAssociatedWithArray,omitempty"`
	Res    *types.QueryFileSystemAssociatedWithArrayResponse `xml:"urn:sms QueryFileSystemAssociatedWithArrayResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryFileSystemAssociatedWithArrayBody) Fault() *soap.Fault { return b.Fault_ }

func QueryFileSystemAssociatedWithArray(ctx context.Context, r soap.RoundTripper, req *types.QueryFileSystemAssociatedWithArray) (*types.QueryFileSystemAssociatedWithArrayResponse, error) {
	var reqBody, resBody QueryFileSystemAssociatedWithArrayBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryHostAssociatedWithLunBody struct {
	Req    *types.QueryHostAssociatedWithLun         `xml:"urn:sms QueryHostAssociatedWithLun,omitempty"`
	Res    *types.QueryHostAssociatedWithLunResponse `xml:"urn:sms QueryHostAssociatedWithLunResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHostAssociatedWithLunBody) Fault() *soap.Fault { return b.Fault_ }

func QueryHostAssociatedWithLun(ctx context.Context, r soap.RoundTripper, req *types.QueryHostAssociatedWithLun) (*types.QueryHostAssociatedWithLunResponse, error) {
	var reqBody, resBody QueryHostAssociatedWithLunBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryLunAssociatedWithArrayBody struct {
	Req    *types.QueryLunAssociatedWithArray         `xml:"urn:sms QueryLunAssociatedWithArray,omitempty"`
	Res    *types.QueryLunAssociatedWithArrayResponse `xml:"urn:sms QueryLunAssociatedWithArrayResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryLunAssociatedWithArrayBody) Fault() *soap.Fault { return b.Fault_ }

func QueryLunAssociatedWithArray(ctx context.Context, r soap.RoundTripper, req *types.QueryLunAssociatedWithArray) (*types.QueryLunAssociatedWithArrayResponse, error) {
	var reqBody, resBody QueryLunAssociatedWithArrayBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryLunAssociatedWithPortBody struct {
	Req    *types.QueryLunAssociatedWithPort         `xml:"urn:sms QueryLunAssociatedWithPort,omitempty"`
	Res    *types.QueryLunAssociatedWithPortResponse `xml:"urn:sms QueryLunAssociatedWithPortResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryLunAssociatedWithPortBody) Fault() *soap.Fault { return b.Fault_ }

func QueryLunAssociatedWithPort(ctx context.Context, r soap.RoundTripper, req *types.QueryLunAssociatedWithPort) (*types.QueryLunAssociatedWithPortResponse, error) {
	var reqBody, resBody QueryLunAssociatedWithPortBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryNfsDatastoreAssociatedWithFileSystemBody struct {
	Req    *types.QueryNfsDatastoreAssociatedWithFileSystem         `xml:"urn:sms QueryNfsDatastoreAssociatedWithFileSystem,omitempty"`
	Res    *types.QueryNfsDatastoreAssociatedWithFileSystemResponse `xml:"urn:sms QueryNfsDatastoreAssociatedWithFileSystemResponse,omitempty"`
	Fault_ *soap.Fault                                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryNfsDatastoreAssociatedWithFileSystemBody) Fault() *soap.Fault { return b.Fault_ }

func QueryNfsDatastoreAssociatedWithFileSystem(ctx context.Context, r soap.RoundTripper, req *types.QueryNfsDatastoreAssociatedWithFileSystem) (*types.QueryNfsDatastoreAssociatedWithFileSystemResponse, error) {
	var reqBody, resBody QueryNfsDatastoreAssociatedWithFileSystemBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPointInTimeReplicaBody struct {
	Req    *types.QueryPointInTimeReplica         `xml:"urn:sms QueryPointInTimeReplica,omitempty"`
	Res    *types.QueryPointInTimeReplicaResponse `xml:"urn:sms QueryPointInTimeReplicaResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPointInTimeReplicaBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPointInTimeReplica(ctx context.Context, r soap.RoundTripper, req *types.QueryPointInTimeReplica) (*types.QueryPointInTimeReplicaResponse, error) {
	var reqBody, resBody QueryPointInTimeReplicaBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPortAssociatedWithArrayBody struct {
	Req    *types.QueryPortAssociatedWithArray         `xml:"urn:sms QueryPortAssociatedWithArray,omitempty"`
	Res    *types.QueryPortAssociatedWithArrayResponse `xml:"urn:sms QueryPortAssociatedWithArrayResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPortAssociatedWithArrayBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPortAssociatedWithArray(ctx context.Context, r soap.RoundTripper, req *types.QueryPortAssociatedWithArray) (*types.QueryPortAssociatedWithArrayResponse, error) {
	var reqBody, resBody QueryPortAssociatedWithArrayBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPortAssociatedWithLunBody struct {
	Req    *types.QueryPortAssociatedWithLun         `xml:"urn:sms QueryPortAssociatedWithLun,omitempty"`
	Res    *types.QueryPortAssociatedWithLunResponse `xml:"urn:sms QueryPortAssociatedWithLunResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPortAssociatedWithLunBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPortAssociatedWithLun(ctx context.Context, r soap.RoundTripper, req *types.QueryPortAssociatedWithLun) (*types.QueryPortAssociatedWithLunResponse, error) {
	var reqBody, resBody QueryPortAssociatedWithLunBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPortAssociatedWithProcessorBody struct {
	Req    *types.QueryPortAssociatedWithProcessor         `xml:"urn:sms QueryPortAssociatedWithProcessor,omitempty"`
	Res    *types.QueryPortAssociatedWithProcessorResponse `xml:"urn:sms QueryPortAssociatedWithProcessorResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPortAssociatedWithProcessorBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPortAssociatedWithProcessor(ctx context.Context, r soap.RoundTripper, req *types.QueryPortAssociatedWithProcessor) (*types.QueryPortAssociatedWithProcessorResponse, error) {
	var reqBody, resBody QueryPortAssociatedWithProcessorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryProcessorAssociatedWithArrayBody struct {
	Req    *types.QueryProcessorAssociatedWithArray         `xml:"urn:sms QueryProcessorAssociatedWithArray,omitempty"`
	Res    *types.QueryProcessorAssociatedWithArrayResponse `xml:"urn:sms QueryProcessorAssociatedWithArrayResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryProcessorAssociatedWithArrayBody) Fault() *soap.Fault { return b.Fault_ }

func QueryProcessorAssociatedWithArray(ctx context.Context, r soap.RoundTripper, req *types.QueryProcessorAssociatedWithArray) (*types.QueryProcessorAssociatedWithArrayResponse, error) {
	var reqBody, resBody QueryProcessorAssociatedWithArrayBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryProviderBody struct {
	Req    *types.QueryProvider         `xml:"urn:sms QueryProvider,omitempty"`
	Res    *types.QueryProviderResponse `xml:"urn:sms QueryProviderResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryProviderBody) Fault() *soap.Fault { return b.Fault_ }

func QueryProvider(ctx context.Context, r soap.RoundTripper, req *types.QueryProvider) (*types.QueryProviderResponse, error) {
	var reqBody, resBody QueryProviderBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryProviderInfoBody struct {
	Req    *types.QueryProviderInfo         `xml:"urn:sms QueryProviderInfo,omitempty"`
	Res    *types.QueryProviderInfoResponse `xml:"urn:sms QueryProviderInfoResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryProviderInfoBody) Fault() *soap.Fault { return b.Fault_ }

func QueryProviderInfo(ctx context.Context, r soap.RoundTripper, req *types.QueryProviderInfo) (*types.QueryProviderInfoResponse, error) {
	var reqBody, resBody QueryProviderInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryReplicationGroupBody struct {
	Req    *types.QueryReplicationGroup         `xml:"urn:sms QueryReplicationGroup,omitempty"`
	Res    *types.QueryReplicationGroupResponse `xml:"urn:sms QueryReplicationGroupResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryReplicationGroupBody) Fault() *soap.Fault { return b.Fault_ }

func QueryReplicationGroup(ctx context.Context, r soap.RoundTripper, req *types.QueryReplicationGroup) (*types.QueryReplicationGroupResponse, error) {
	var reqBody, resBody QueryReplicationGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryReplicationGroupInfoBody struct {
	Req    *types.QueryReplicationGroupInfo         `xml:"urn:sms QueryReplicationGroupInfo,omitempty"`
	Res    *types.QueryReplicationGroupInfoResponse `xml:"urn:sms QueryReplicationGroupInfoResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryReplicationGroupInfoBody) Fault() *soap.Fault { return b.Fault_ }

func QueryReplicationGroupInfo(ctx context.Context, r soap.RoundTripper, req *types.QueryReplicationGroupInfo) (*types.QueryReplicationGroupInfoResponse, error) {
	var reqBody, resBody QueryReplicationGroupInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryReplicationPeerBody struct {
	Req    *types.QueryReplicationPeer         `xml:"urn:sms QueryReplicationPeer,omitempty"`
	Res    *types.QueryReplicationPeerResponse `xml:"urn:sms QueryReplicationPeerResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryReplicationPeerBody) Fault() *soap.Fault { return b.Fault_ }

func QueryReplicationPeer(ctx context.Context, r soap.RoundTripper, req *types.QueryReplicationPeer) (*types.QueryReplicationPeerResponse, error) {
	var reqBody, resBody QueryReplicationPeerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QuerySessionManagerBody struct {
	Req    *types.QuerySessionManager         `xml:"urn:sms QuerySessionManager,omitempty"`
	Res    *types.QuerySessionManagerResponse `xml:"urn:sms QuerySessionManagerResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QuerySessionManagerBody) Fault() *soap.Fault { return b.Fault_ }

func QuerySessionManager(ctx context.Context, r soap.RoundTripper, req *types.QuerySessionManager) (*types.QuerySessionManagerResponse, error) {
	var reqBody, resBody QuerySessionManagerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QuerySmsTaskInfoBody struct {
	Req    *types.QuerySmsTaskInfo         `xml:"urn:sms QuerySmsTaskInfo,omitempty"`
	Res    *types.QuerySmsTaskInfoResponse `xml:"urn:sms QuerySmsTaskInfoResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QuerySmsTaskInfoBody) Fault() *soap.Fault { return b.Fault_ }

func QuerySmsTaskInfo(ctx context.Context, r soap.RoundTripper, req *types.QuerySmsTaskInfo) (*types.QuerySmsTaskInfoResponse, error) {
	var reqBody, resBody QuerySmsTaskInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QuerySmsTaskResultBody struct {
	Req    *types.QuerySmsTaskResult         `xml:"urn:sms QuerySmsTaskResult,omitempty"`
	Res    *types.QuerySmsTaskResultResponse `xml:"urn:sms QuerySmsTaskResultResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QuerySmsTaskResultBody) Fault() *soap.Fault { return b.Fault_ }

func QuerySmsTaskResult(ctx context.Context, r soap.RoundTripper, req *types.QuerySmsTaskResult) (*types.QuerySmsTaskResultResponse, error) {
	var reqBody, resBody QuerySmsTaskResultBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryStorageContainerBody struct {
	Req    *types.QueryStorageContainer         `xml:"urn:sms QueryStorageContainer,omitempty"`
	Res    *types.QueryStorageContainerResponse `xml:"urn:sms QueryStorageContainerResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryStorageContainerBody) Fault() *soap.Fault { return b.Fault_ }

func QueryStorageContainer(ctx context.Context, r soap.RoundTripper, req *types.QueryStorageContainer) (*types.QueryStorageContainerResponse, error) {
	var reqBody, resBody QueryStorageContainerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryStorageManagerBody struct {
	Req    *types.QueryStorageManager         `xml:"urn:sms QueryStorageManager,omitempty"`
	Res    *types.QueryStorageManagerResponse `xml:"urn:sms QueryStorageManagerResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryStorageManagerBody) Fault() *soap.Fault { return b.Fault_ }

func QueryStorageManager(ctx context.Context, r soap.RoundTripper, req *types.QueryStorageManager) (*types.QueryStorageManagerResponse, error) {
	var reqBody, resBody QueryStorageManagerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVmfsDatastoreAssociatedWithLunBody struct {
	Req    *types.QueryVmfsDatastoreAssociatedWithLun         `xml:"urn:sms QueryVmfsDatastoreAssociatedWithLun,omitempty"`
	Res    *types.QueryVmfsDatastoreAssociatedWithLunResponse `xml:"urn:sms QueryVmfsDatastoreAssociatedWithLunResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVmfsDatastoreAssociatedWithLunBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVmfsDatastoreAssociatedWithLun(ctx context.Context, r soap.RoundTripper, req *types.QueryVmfsDatastoreAssociatedWithLun) (*types.QueryVmfsDatastoreAssociatedWithLunResponse, error) {
	var reqBody, resBody QueryVmfsDatastoreAssociatedWithLunBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RegisterProvider_TaskBody struct {
	Req    *types.RegisterProvider_Task         `xml:"urn:sms RegisterProvider_Task,omitempty"`
	Res    *types.RegisterProvider_TaskResponse `xml:"urn:sms RegisterProvider_TaskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RegisterProvider_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RegisterProvider_Task(ctx context.Context, r soap.RoundTripper, req *types.RegisterProvider_Task) (*types.RegisterProvider_TaskResponse, error) {
	var reqBody, resBody RegisterProvider_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReverseReplicateGroup_TaskBody struct {
	Req    *types.ReverseReplicateGroup_Task         `xml:"urn:sms ReverseReplicateGroup_Task,omitempty"`
	Res    *types.ReverseReplicateGroup_TaskResponse `xml:"urn:sms ReverseReplicateGroup_TaskResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReverseReplicateGroup_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReverseReplicateGroup_Task(ctx context.Context, r soap.RoundTripper, req *types.ReverseReplicateGroup_Task) (*types.ReverseReplicateGroup_TaskResponse, error) {
	var reqBody, resBody ReverseReplicateGroup_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SmsRefreshCACertificatesAndCRLs_TaskBody struct {
	Req    *types.SmsRefreshCACertificatesAndCRLs_Task         `xml:"urn:sms SmsRefreshCACertificatesAndCRLs_Task,omitempty"`
	Res    *types.SmsRefreshCACertificatesAndCRLs_TaskResponse `xml:"urn:sms SmsRefreshCACertificatesAndCRLs_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SmsRefreshCACertificatesAndCRLs_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func SmsRefreshCACertificatesAndCRLs_Task(ctx context.Context, r soap.RoundTripper, req *types.SmsRefreshCACertificatesAndCRLs_Task) (*types.SmsRefreshCACertificatesAndCRLs_TaskResponse, error) {
	var reqBody, resBody SmsRefreshCACertificatesAndCRLs_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SyncReplicationGroup_TaskBody struct {
	Req    *types.SyncReplicationGroup_Task         `xml:"urn:sms SyncReplicationGroup_Task,omitempty"`
	Res    *types.SyncReplicationGroup_TaskResponse `xml:"urn:sms SyncReplicationGroup_TaskResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SyncReplicationGroup_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func SyncReplicationGroup_Task(ctx context.Context, r soap.RoundTripper, req *types.SyncReplicationGroup_Task) (*types.SyncReplicationGroup_TaskResponse, error) {
	var reqBody, resBody SyncReplicationGroup_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type TestFailoverReplicationGroupStart_TaskBody struct {
	Req    *types.TestFailoverReplicationGroupStart_Task         `xml:"urn:sms TestFailoverReplicationGroupStart_Task,omitempty"`
	Res    *types.TestFailoverReplicationGroupStart_TaskResponse `xml:"urn:sms TestFailoverReplicationGroupStart_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TestFailoverReplicationGroupStart_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func TestFailoverReplicationGroupStart_Task(ctx context.Context, r soap.RoundTripper, req *types.TestFailoverReplicationGroupStart_Task) (*types.TestFailoverReplicationGroupStart_TaskResponse, error) {
	var reqBody, resBody TestFailoverReplicationGroupStart_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type TestFailoverReplicationGroupStop_TaskBody struct {
	Req    *types.TestFailoverReplicationGroupStop_Task         `xml:"urn:sms TestFailoverReplicationGroupStop_Task,omitempty"`
	Res    *types.TestFailoverReplicationGroupStop_TaskResponse `xml:"urn:sms TestFailoverReplicationGroupStop_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TestFailoverReplicationGroupStop_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func TestFailoverReplicationGroupStop_Task(ctx context.Context, r soap.RoundTripper, req *types.TestFailoverReplicationGroupStop_Task) (*types.TestFailoverReplicationGroupStop_TaskResponse, error) {
	var reqBody, resBody TestFailoverReplicationGroupStop_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnregisterProvider_TaskBody struct {
	Req    *types.UnregisterProvider_Task         `xml:"urn:sms UnregisterProvider_Task,omitempty"`
	Res    *types.UnregisterProvider_TaskResponse `xml:"urn:sms UnregisterProvider_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnregisterProvider_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UnregisterProvider_Task(ctx context.Context, r soap.RoundTripper, req *types.UnregisterProvider_Task) (*types.UnregisterProvider_TaskResponse, error) {
	var reqBody, resBody UnregisterProvider_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpgradeVASAProvider_TaskBody struct {
	Req    *types.UpgradeVASAProvider_Task         `xml:"urn:sms UpgradeVASAProvider_Task,omitempty"`
	Res    *types.UpgradeVASAProvider_TaskResponse `xml:"urn:sms UpgradeVASAProvider_TaskResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpgradeVASAProvider_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpgradeVASAProvider_Task(ctx context.Context, r soap.RoundTripper, req *types.UpgradeVASAProvider_Task) (*types.UpgradeVASAProvider_TaskResponse, error) {
	var reqBody, resBody UpgradeVASAProvider_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VasaProviderReconnect_TaskBody struct {
	Req    *types.VasaProviderReconnect_Task         `xml:"urn:sms VasaProviderReconnect_Task,omitempty"`
	Res    *types.VasaProviderReconnect_TaskResponse `xml:"urn:sms VasaProviderReconnect_TaskResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VasaProviderReconnect_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VasaProviderReconnect_Task(ctx context.Context, r soap.RoundTripper, req *types.VasaProviderReconnect_Task) (*types.VasaProviderReconnect_TaskResponse, error) {
	var reqBody, resBody VasaProviderReconnect_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VasaProviderRefreshCertificate_TaskBody struct {
	Req    *types.VasaProviderRefreshCertificate_Task         `xml:"urn:sms VasaProviderRefreshCertificate_Task,omitempty"`
	Res    *types.VasaProviderRefreshCertificate_TaskResponse `xml:"urn:sms VasaProviderRefreshCertificate_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VasaProviderRefreshCertificate_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VasaProviderRefreshCertificate_Task(ctx context.Context, r soap.RoundTripper, req *types.VasaProviderRefreshCertificate_Task) (*types.VasaProviderRefreshCertificate_TaskResponse, error) {
	var reqBody, resBody VasaProviderRefreshCertificate_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VasaProviderRevokeCertificate_TaskBody struct {
	Req    *types.VasaProviderRevokeCertificate_Task         `xml:"urn:sms VasaProviderRevokeCertificate_Task,omitempty"`
	Res    *types.VasaProviderRevokeCertificate_TaskResponse `xml:"urn:sms VasaProviderRevokeCertificate_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VasaProviderRevokeCertificate_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VasaProviderRevokeCertificate_Task(ctx context.Context, r soap.RoundTripper, req *types.VasaProviderRevokeCertificate_Task) (*types.VasaProviderRevokeCertificate_TaskResponse, error) {
	var reqBody, resBody VasaProviderRevokeCertificate_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VasaProviderSync_TaskBody struct {
	Req    *types.VasaProviderSync_Task         `xml:"urn:sms VasaProviderSync_Task,omitempty"`
	Res    *types.VasaProviderSync_TaskResponse `xml:"urn:sms VasaProviderSync_TaskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VasaProviderSync_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VasaProviderSync_Task(ctx context.Context, r soap.RoundTripper, req *types.VasaProviderSync_Task) (*types.VasaProviderSync_TaskResponse, error) {
	var reqBody, resBody VasaProviderSync_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}
