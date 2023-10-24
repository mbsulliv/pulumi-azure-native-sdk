// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storagesync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a ServerEndpoint.
// Azure REST API version: 2022-06-01.
//
// Other available API versions: 2017-06-05-preview, 2018-04-02, 2018-07-01, 2018-10-01, 2019-10-01, 2022-09-01.
func LookupServerEndpoint(ctx *pulumi.Context, args *LookupServerEndpointArgs, opts ...pulumi.InvokeOption) (*LookupServerEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServerEndpointResult
	err := ctx.Invoke("azure-native:storagesync:getServerEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerEndpointArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Server Endpoint object.
	ServerEndpointName string `pulumi:"serverEndpointName"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	// Name of Sync Group resource.
	SyncGroupName string `pulumi:"syncGroupName"`
}

// Server Endpoint object.
type LookupServerEndpointResult struct {
	// Cloud Tiering.
	CloudTiering *string `pulumi:"cloudTiering"`
	// Cloud tiering status. Only populated if cloud tiering is enabled.
	CloudTieringStatus ServerEndpointCloudTieringStatusResponse `pulumi:"cloudTieringStatus"`
	// Friendly Name
	FriendlyName *string `pulumi:"friendlyName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Policy for how namespace and files are recalled during FastDr.
	InitialDownloadPolicy *string `pulumi:"initialDownloadPolicy"`
	// Policy for how the initial upload sync session is performed.
	InitialUploadPolicy *string `pulumi:"initialUploadPolicy"`
	// Resource Last Operation Name
	LastOperationName string `pulumi:"lastOperationName"`
	// ServerEndpoint lastWorkflowId
	LastWorkflowId string `pulumi:"lastWorkflowId"`
	// Policy for enabling follow-the-sun business models: link local cache to cloud behavior to pre-populate before local access.
	LocalCacheMode *string `pulumi:"localCacheMode"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Offline data transfer
	OfflineDataTransfer *string `pulumi:"offlineDataTransfer"`
	// Offline data transfer share name
	OfflineDataTransferShareName *string `pulumi:"offlineDataTransferShareName"`
	// Offline data transfer storage account resource ID
	OfflineDataTransferStorageAccountResourceId string `pulumi:"offlineDataTransferStorageAccountResourceId"`
	// Offline data transfer storage account tenant ID
	OfflineDataTransferStorageAccountTenantId string `pulumi:"offlineDataTransferStorageAccountTenantId"`
	// ServerEndpoint Provisioning State
	ProvisioningState string `pulumi:"provisioningState"`
	// Recall status. Only populated if cloud tiering is enabled.
	RecallStatus ServerEndpointRecallStatusResponse `pulumi:"recallStatus"`
	// Server Local path.
	ServerLocalPath *string `pulumi:"serverLocalPath"`
	// Server name
	ServerName string `pulumi:"serverName"`
	// Server Resource Id.
	ServerResourceId *string `pulumi:"serverResourceId"`
	// Server Endpoint sync status
	SyncStatus ServerEndpointSyncStatusResponse `pulumi:"syncStatus"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Tier files older than days.
	TierFilesOlderThanDays *int `pulumi:"tierFilesOlderThanDays"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Level of free space to be maintained by Cloud Tiering if it is enabled.
	VolumeFreeSpacePercent *int `pulumi:"volumeFreeSpacePercent"`
}

func LookupServerEndpointOutput(ctx *pulumi.Context, args LookupServerEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupServerEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerEndpointResult, error) {
			args := v.(LookupServerEndpointArgs)
			r, err := LookupServerEndpoint(ctx, &args, opts...)
			var s LookupServerEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerEndpointResultOutput)
}

type LookupServerEndpointOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Server Endpoint object.
	ServerEndpointName pulumi.StringInput `pulumi:"serverEndpointName"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName pulumi.StringInput `pulumi:"storageSyncServiceName"`
	// Name of Sync Group resource.
	SyncGroupName pulumi.StringInput `pulumi:"syncGroupName"`
}

func (LookupServerEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerEndpointArgs)(nil)).Elem()
}

// Server Endpoint object.
type LookupServerEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupServerEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerEndpointResult)(nil)).Elem()
}

func (o LookupServerEndpointResultOutput) ToLookupServerEndpointResultOutput() LookupServerEndpointResultOutput {
	return o
}

func (o LookupServerEndpointResultOutput) ToLookupServerEndpointResultOutputWithContext(ctx context.Context) LookupServerEndpointResultOutput {
	return o
}

func (o LookupServerEndpointResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupServerEndpointResult] {
	return pulumix.Output[LookupServerEndpointResult]{
		OutputState: o.OutputState,
	}
}

// Cloud Tiering.
func (o LookupServerEndpointResultOutput) CloudTiering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.CloudTiering }).(pulumi.StringPtrOutput)
}

// Cloud tiering status. Only populated if cloud tiering is enabled.
func (o LookupServerEndpointResultOutput) CloudTieringStatus() ServerEndpointCloudTieringStatusResponseOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) ServerEndpointCloudTieringStatusResponse {
		return v.CloudTieringStatus
	}).(ServerEndpointCloudTieringStatusResponseOutput)
}

// Friendly Name
func (o LookupServerEndpointResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupServerEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// Policy for how namespace and files are recalled during FastDr.
func (o LookupServerEndpointResultOutput) InitialDownloadPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.InitialDownloadPolicy }).(pulumi.StringPtrOutput)
}

// Policy for how the initial upload sync session is performed.
func (o LookupServerEndpointResultOutput) InitialUploadPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.InitialUploadPolicy }).(pulumi.StringPtrOutput)
}

// Resource Last Operation Name
func (o LookupServerEndpointResultOutput) LastOperationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.LastOperationName }).(pulumi.StringOutput)
}

// ServerEndpoint lastWorkflowId
func (o LookupServerEndpointResultOutput) LastWorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.LastWorkflowId }).(pulumi.StringOutput)
}

// Policy for enabling follow-the-sun business models: link local cache to cloud behavior to pre-populate before local access.
func (o LookupServerEndpointResultOutput) LocalCacheMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.LocalCacheMode }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupServerEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// Offline data transfer
func (o LookupServerEndpointResultOutput) OfflineDataTransfer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.OfflineDataTransfer }).(pulumi.StringPtrOutput)
}

// Offline data transfer share name
func (o LookupServerEndpointResultOutput) OfflineDataTransferShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.OfflineDataTransferShareName }).(pulumi.StringPtrOutput)
}

// Offline data transfer storage account resource ID
func (o LookupServerEndpointResultOutput) OfflineDataTransferStorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.OfflineDataTransferStorageAccountResourceId }).(pulumi.StringOutput)
}

// Offline data transfer storage account tenant ID
func (o LookupServerEndpointResultOutput) OfflineDataTransferStorageAccountTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.OfflineDataTransferStorageAccountTenantId }).(pulumi.StringOutput)
}

// ServerEndpoint Provisioning State
func (o LookupServerEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Recall status. Only populated if cloud tiering is enabled.
func (o LookupServerEndpointResultOutput) RecallStatus() ServerEndpointRecallStatusResponseOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) ServerEndpointRecallStatusResponse { return v.RecallStatus }).(ServerEndpointRecallStatusResponseOutput)
}

// Server Local path.
func (o LookupServerEndpointResultOutput) ServerLocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.ServerLocalPath }).(pulumi.StringPtrOutput)
}

// Server name
func (o LookupServerEndpointResultOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.ServerName }).(pulumi.StringOutput)
}

// Server Resource Id.
func (o LookupServerEndpointResultOutput) ServerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.ServerResourceId }).(pulumi.StringPtrOutput)
}

// Server Endpoint sync status
func (o LookupServerEndpointResultOutput) SyncStatus() ServerEndpointSyncStatusResponseOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) ServerEndpointSyncStatusResponse { return v.SyncStatus }).(ServerEndpointSyncStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupServerEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Tier files older than days.
func (o LookupServerEndpointResultOutput) TierFilesOlderThanDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *int { return v.TierFilesOlderThanDays }).(pulumi.IntPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupServerEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

// Level of free space to be maintained by Cloud Tiering if it is enabled.
func (o LookupServerEndpointResultOutput) VolumeFreeSpacePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *int { return v.VolumeFreeSpacePercent }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerEndpointResultOutput{})
}
