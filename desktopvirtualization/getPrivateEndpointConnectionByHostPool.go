// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package desktopvirtualization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a private endpoint connection.
// Azure REST API version: 2022-10-14-preview.
func LookupPrivateEndpointConnectionByHostPool(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionByHostPoolArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionByHostPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateEndpointConnectionByHostPoolResult
	err := ctx.Invoke("azure-native:desktopvirtualization:getPrivateEndpointConnectionByHostPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionByHostPoolArgs struct {
	// The name of the host pool within the specified resource group
	HostPoolName string `pulumi:"hostPoolName"`
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Private Endpoint Connection resource.
type LookupPrivateEndpointConnectionByHostPoolResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupPrivateEndpointConnectionByHostPoolOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionByHostPoolOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionByHostPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionByHostPoolResult, error) {
			args := v.(LookupPrivateEndpointConnectionByHostPoolArgs)
			r, err := LookupPrivateEndpointConnectionByHostPool(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionByHostPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionByHostPoolResultOutput)
}

type LookupPrivateEndpointConnectionByHostPoolOutputArgs struct {
	// The name of the host pool within the specified resource group
	HostPoolName pulumi.StringInput `pulumi:"hostPoolName"`
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateEndpointConnectionByHostPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionByHostPoolArgs)(nil)).Elem()
}

// The Private Endpoint Connection resource.
type LookupPrivateEndpointConnectionByHostPoolResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionByHostPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionByHostPoolResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) ToLookupPrivateEndpointConnectionByHostPoolResultOutput() LookupPrivateEndpointConnectionByHostPoolResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) ToLookupPrivateEndpointConnectionByHostPoolResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionByHostPoolResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPrivateEndpointConnectionByHostPoolResult] {
	return pulumix.Output[LookupPrivateEndpointConnectionByHostPoolResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPrivateEndpointConnectionByHostPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByHostPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionByHostPoolResultOutput{})
}
