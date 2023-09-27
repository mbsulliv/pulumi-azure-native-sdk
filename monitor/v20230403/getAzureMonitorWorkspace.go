// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230403

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns the specified Azure Monitor Workspace
func LookupAzureMonitorWorkspace(ctx *pulumi.Context, args *LookupAzureMonitorWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupAzureMonitorWorkspaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAzureMonitorWorkspaceResult
	err := ctx.Invoke("azure-native:monitor/v20230403:getAzureMonitorWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAzureMonitorWorkspaceArgs struct {
	// The name of the Azure Monitor Workspace. The name is case insensitive
	AzureMonitorWorkspaceName string `pulumi:"azureMonitorWorkspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Monitor Workspace definition
type LookupAzureMonitorWorkspaceResult struct {
	// The immutable Id of the Azure Monitor Workspace. This property is read-only.
	AccountId string `pulumi:"accountId"`
	// The Data Collection Rule and Endpoint used for ingestion by default.
	DefaultIngestionSettings AzureMonitorWorkspaceResponseDefaultIngestionSettings `pulumi:"defaultIngestionSettings"`
	// Resource entity tag (ETag)
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Properties related to the metrics container in the Azure Monitor Workspace
	Metrics AzureMonitorWorkspaceResponseMetrics `pulumi:"metrics"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of private endpoint connections
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The provisioning state of the Azure Monitor Workspace. Set to Succeeded if everything is healthy.
	ProvisioningState string `pulumi:"provisioningState"`
	// Gets or sets allow or disallow public network access to Azure Monitor Workspace
	PublicNetworkAccess string `pulumi:"publicNetworkAccess"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAzureMonitorWorkspaceOutput(ctx *pulumi.Context, args LookupAzureMonitorWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupAzureMonitorWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzureMonitorWorkspaceResult, error) {
			args := v.(LookupAzureMonitorWorkspaceArgs)
			r, err := LookupAzureMonitorWorkspace(ctx, &args, opts...)
			var s LookupAzureMonitorWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzureMonitorWorkspaceResultOutput)
}

type LookupAzureMonitorWorkspaceOutputArgs struct {
	// The name of the Azure Monitor Workspace. The name is case insensitive
	AzureMonitorWorkspaceName pulumi.StringInput `pulumi:"azureMonitorWorkspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAzureMonitorWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureMonitorWorkspaceArgs)(nil)).Elem()
}

// An Azure Monitor Workspace definition
type LookupAzureMonitorWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupAzureMonitorWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureMonitorWorkspaceResult)(nil)).Elem()
}

func (o LookupAzureMonitorWorkspaceResultOutput) ToLookupAzureMonitorWorkspaceResultOutput() LookupAzureMonitorWorkspaceResultOutput {
	return o
}

func (o LookupAzureMonitorWorkspaceResultOutput) ToLookupAzureMonitorWorkspaceResultOutputWithContext(ctx context.Context) LookupAzureMonitorWorkspaceResultOutput {
	return o
}

func (o LookupAzureMonitorWorkspaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAzureMonitorWorkspaceResult] {
	return pulumix.Output[LookupAzureMonitorWorkspaceResult]{
		OutputState: o.OutputState,
	}
}

// The immutable Id of the Azure Monitor Workspace. This property is read-only.
func (o LookupAzureMonitorWorkspaceResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.AccountId }).(pulumi.StringOutput)
}

// The Data Collection Rule and Endpoint used for ingestion by default.
func (o LookupAzureMonitorWorkspaceResultOutput) DefaultIngestionSettings() AzureMonitorWorkspaceResponseDefaultIngestionSettingsOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) AzureMonitorWorkspaceResponseDefaultIngestionSettings {
		return v.DefaultIngestionSettings
	}).(AzureMonitorWorkspaceResponseDefaultIngestionSettingsOutput)
}

// Resource entity tag (ETag)
func (o LookupAzureMonitorWorkspaceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAzureMonitorWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupAzureMonitorWorkspaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Properties related to the metrics container in the Azure Monitor Workspace
func (o LookupAzureMonitorWorkspaceResultOutput) Metrics() AzureMonitorWorkspaceResponseMetricsOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) AzureMonitorWorkspaceResponseMetrics { return v.Metrics }).(AzureMonitorWorkspaceResponseMetricsOutput)
}

// The name of the resource
func (o LookupAzureMonitorWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connections
func (o LookupAzureMonitorWorkspaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioning state of the Azure Monitor Workspace. Set to Succeeded if everything is healthy.
func (o LookupAzureMonitorWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets allow or disallow public network access to Azure Monitor Workspace
func (o LookupAzureMonitorWorkspaceResultOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAzureMonitorWorkspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAzureMonitorWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAzureMonitorWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureMonitorWorkspaceResultOutput{})
}
