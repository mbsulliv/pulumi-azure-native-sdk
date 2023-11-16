// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The managed private endpoint resource type.
func LookupManagedPrivateEndpoint(ctx *pulumi.Context, args *LookupManagedPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupManagedPrivateEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedPrivateEndpointResult
	err := ctx.Invoke("azure-native:dashboard/v20230901:getManagedPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedPrivateEndpointArgs struct {
	// The managed private endpoint name of Azure Managed Grafana.
	ManagedPrivateEndpointName string `pulumi:"managedPrivateEndpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workspace name of Azure Managed Grafana.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The managed private endpoint resource type.
type LookupManagedPrivateEndpointResult struct {
	// The state of managed private endpoint connection.
	ConnectionState ManagedPrivateEndpointConnectionStateResponse `pulumi:"connectionState"`
	// The group Ids of the managed private endpoint.
	GroupIds []string `pulumi:"groupIds"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The ARM resource ID of the resource for which the managed private endpoint is pointing to.
	PrivateLinkResourceId *string `pulumi:"privateLinkResourceId"`
	// The region of the resource to which the managed private endpoint is pointing to.
	PrivateLinkResourceRegion *string `pulumi:"privateLinkResourceRegion"`
	// The private IP of private endpoint after approval. This property is empty before connection is approved.
	PrivateLinkServicePrivateIP string `pulumi:"privateLinkServicePrivateIP"`
	// The URL of the data store behind the private link service. It would be the URL in the Grafana data source configuration page without the protocol and port.
	PrivateLinkServiceUrl *string `pulumi:"privateLinkServiceUrl"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// User input request message of the managed private endpoint.
	RequestMessage *string `pulumi:"requestMessage"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupManagedPrivateEndpointOutput(ctx *pulumi.Context, args LookupManagedPrivateEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupManagedPrivateEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedPrivateEndpointResult, error) {
			args := v.(LookupManagedPrivateEndpointArgs)
			r, err := LookupManagedPrivateEndpoint(ctx, &args, opts...)
			var s LookupManagedPrivateEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedPrivateEndpointResultOutput)
}

type LookupManagedPrivateEndpointOutputArgs struct {
	// The managed private endpoint name of Azure Managed Grafana.
	ManagedPrivateEndpointName pulumi.StringInput `pulumi:"managedPrivateEndpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The workspace name of Azure Managed Grafana.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupManagedPrivateEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedPrivateEndpointArgs)(nil)).Elem()
}

// The managed private endpoint resource type.
type LookupManagedPrivateEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupManagedPrivateEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedPrivateEndpointResult)(nil)).Elem()
}

func (o LookupManagedPrivateEndpointResultOutput) ToLookupManagedPrivateEndpointResultOutput() LookupManagedPrivateEndpointResultOutput {
	return o
}

func (o LookupManagedPrivateEndpointResultOutput) ToLookupManagedPrivateEndpointResultOutputWithContext(ctx context.Context) LookupManagedPrivateEndpointResultOutput {
	return o
}

// The state of managed private endpoint connection.
func (o LookupManagedPrivateEndpointResultOutput) ConnectionState() ManagedPrivateEndpointConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) ManagedPrivateEndpointConnectionStateResponse {
		return v.ConnectionState
	}).(ManagedPrivateEndpointConnectionStateResponseOutput)
}

// The group Ids of the managed private endpoint.
func (o LookupManagedPrivateEndpointResultOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupManagedPrivateEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupManagedPrivateEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupManagedPrivateEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ARM resource ID of the resource for which the managed private endpoint is pointing to.
func (o LookupManagedPrivateEndpointResultOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

// The region of the resource to which the managed private endpoint is pointing to.
func (o LookupManagedPrivateEndpointResultOutput) PrivateLinkResourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) *string { return v.PrivateLinkResourceRegion }).(pulumi.StringPtrOutput)
}

// The private IP of private endpoint after approval. This property is empty before connection is approved.
func (o LookupManagedPrivateEndpointResultOutput) PrivateLinkServicePrivateIP() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.PrivateLinkServicePrivateIP }).(pulumi.StringOutput)
}

// The URL of the data store behind the private link service. It would be the URL in the Grafana data source configuration page without the protocol and port.
func (o LookupManagedPrivateEndpointResultOutput) PrivateLinkServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) *string { return v.PrivateLinkServiceUrl }).(pulumi.StringPtrOutput)
}

// Provisioning state of the resource.
func (o LookupManagedPrivateEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// User input request message of the managed private endpoint.
func (o LookupManagedPrivateEndpointResultOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupManagedPrivateEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupManagedPrivateEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupManagedPrivateEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedPrivateEndpointResultOutput{})
}
