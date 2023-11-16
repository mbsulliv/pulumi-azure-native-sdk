// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storagesync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a given registered server.
// Azure REST API version: 2022-06-01.
//
// Other available API versions: 2017-06-05-preview, 2018-04-02, 2018-07-01, 2022-09-01.
func LookupRegisteredServer(ctx *pulumi.Context, args *LookupRegisteredServerArgs, opts ...pulumi.InvokeOption) (*LookupRegisteredServerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegisteredServerResult
	err := ctx.Invoke("azure-native:storagesync:getRegisteredServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegisteredServerArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// GUID identifying the on-premises server.
	ServerId string `pulumi:"serverId"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
}

// Registered Server resource.
type LookupRegisteredServerResult struct {
	// Registered Server Agent Version
	AgentVersion *string `pulumi:"agentVersion"`
	// Registered Server Agent Version Expiration Date
	AgentVersionExpirationDate string `pulumi:"agentVersionExpirationDate"`
	// Registered Server Agent Version Status
	AgentVersionStatus string `pulumi:"agentVersionStatus"`
	// Registered Server clusterId
	ClusterId *string `pulumi:"clusterId"`
	// Registered Server clusterName
	ClusterName *string `pulumi:"clusterName"`
	// Resource discoveryEndpointUri
	DiscoveryEndpointUri *string `pulumi:"discoveryEndpointUri"`
	// Friendly Name
	FriendlyName *string `pulumi:"friendlyName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Registered Server last heart beat
	LastHeartBeat *string `pulumi:"lastHeartBeat"`
	// Resource Last Operation Name
	LastOperationName *string `pulumi:"lastOperationName"`
	// Registered Server lastWorkflowId
	LastWorkflowId *string `pulumi:"lastWorkflowId"`
	// Management Endpoint Uri
	ManagementEndpointUri *string `pulumi:"managementEndpointUri"`
	// Monitoring Configuration
	MonitoringConfiguration *string `pulumi:"monitoringConfiguration"`
	// Telemetry Endpoint Uri
	MonitoringEndpointUri *string `pulumi:"monitoringEndpointUri"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Registered Server Provisioning State
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource Location
	ResourceLocation *string `pulumi:"resourceLocation"`
	// Registered Server Certificate
	ServerCertificate *string `pulumi:"serverCertificate"`
	// Registered Server serverId
	ServerId *string `pulumi:"serverId"`
	// Registered Server Management Error Code
	ServerManagementErrorCode *int `pulumi:"serverManagementErrorCode"`
	// Server name
	ServerName string `pulumi:"serverName"`
	// Registered Server OS Version
	ServerOSVersion *string `pulumi:"serverOSVersion"`
	// Registered Server serverRole
	ServerRole *string `pulumi:"serverRole"`
	// Service Location
	ServiceLocation *string `pulumi:"serviceLocation"`
	// Registered Server storageSyncServiceUid
	StorageSyncServiceUid *string `pulumi:"storageSyncServiceUid"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupRegisteredServerOutput(ctx *pulumi.Context, args LookupRegisteredServerOutputArgs, opts ...pulumi.InvokeOption) LookupRegisteredServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegisteredServerResult, error) {
			args := v.(LookupRegisteredServerArgs)
			r, err := LookupRegisteredServer(ctx, &args, opts...)
			var s LookupRegisteredServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegisteredServerResultOutput)
}

type LookupRegisteredServerOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// GUID identifying the on-premises server.
	ServerId pulumi.StringInput `pulumi:"serverId"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName pulumi.StringInput `pulumi:"storageSyncServiceName"`
}

func (LookupRegisteredServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegisteredServerArgs)(nil)).Elem()
}

// Registered Server resource.
type LookupRegisteredServerResultOutput struct{ *pulumi.OutputState }

func (LookupRegisteredServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegisteredServerResult)(nil)).Elem()
}

func (o LookupRegisteredServerResultOutput) ToLookupRegisteredServerResultOutput() LookupRegisteredServerResultOutput {
	return o
}

func (o LookupRegisteredServerResultOutput) ToLookupRegisteredServerResultOutputWithContext(ctx context.Context) LookupRegisteredServerResultOutput {
	return o
}

// Registered Server Agent Version
func (o LookupRegisteredServerResultOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

// Registered Server Agent Version Expiration Date
func (o LookupRegisteredServerResultOutput) AgentVersionExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) string { return v.AgentVersionExpirationDate }).(pulumi.StringOutput)
}

// Registered Server Agent Version Status
func (o LookupRegisteredServerResultOutput) AgentVersionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) string { return v.AgentVersionStatus }).(pulumi.StringOutput)
}

// Registered Server clusterId
func (o LookupRegisteredServerResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// Registered Server clusterName
func (o LookupRegisteredServerResultOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ClusterName }).(pulumi.StringPtrOutput)
}

// Resource discoveryEndpointUri
func (o LookupRegisteredServerResultOutput) DiscoveryEndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.DiscoveryEndpointUri }).(pulumi.StringPtrOutput)
}

// Friendly Name
func (o LookupRegisteredServerResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRegisteredServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// Registered Server last heart beat
func (o LookupRegisteredServerResultOutput) LastHeartBeat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.LastHeartBeat }).(pulumi.StringPtrOutput)
}

// Resource Last Operation Name
func (o LookupRegisteredServerResultOutput) LastOperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.LastOperationName }).(pulumi.StringPtrOutput)
}

// Registered Server lastWorkflowId
func (o LookupRegisteredServerResultOutput) LastWorkflowId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.LastWorkflowId }).(pulumi.StringPtrOutput)
}

// Management Endpoint Uri
func (o LookupRegisteredServerResultOutput) ManagementEndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ManagementEndpointUri }).(pulumi.StringPtrOutput)
}

// Monitoring Configuration
func (o LookupRegisteredServerResultOutput) MonitoringConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.MonitoringConfiguration }).(pulumi.StringPtrOutput)
}

// Telemetry Endpoint Uri
func (o LookupRegisteredServerResultOutput) MonitoringEndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.MonitoringEndpointUri }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupRegisteredServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Registered Server Provisioning State
func (o LookupRegisteredServerResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Resource Location
func (o LookupRegisteredServerResultOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

// Registered Server Certificate
func (o LookupRegisteredServerResultOutput) ServerCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ServerCertificate }).(pulumi.StringPtrOutput)
}

// Registered Server serverId
func (o LookupRegisteredServerResultOutput) ServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ServerId }).(pulumi.StringPtrOutput)
}

// Registered Server Management Error Code
func (o LookupRegisteredServerResultOutput) ServerManagementErrorCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *int { return v.ServerManagementErrorCode }).(pulumi.IntPtrOutput)
}

// Server name
func (o LookupRegisteredServerResultOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) string { return v.ServerName }).(pulumi.StringOutput)
}

// Registered Server OS Version
func (o LookupRegisteredServerResultOutput) ServerOSVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ServerOSVersion }).(pulumi.StringPtrOutput)
}

// Registered Server serverRole
func (o LookupRegisteredServerResultOutput) ServerRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ServerRole }).(pulumi.StringPtrOutput)
}

// Service Location
func (o LookupRegisteredServerResultOutput) ServiceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.ServiceLocation }).(pulumi.StringPtrOutput)
}

// Registered Server storageSyncServiceUid
func (o LookupRegisteredServerResultOutput) StorageSyncServiceUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) *string { return v.StorageSyncServiceUid }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRegisteredServerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRegisteredServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegisteredServerResultOutput{})
}
