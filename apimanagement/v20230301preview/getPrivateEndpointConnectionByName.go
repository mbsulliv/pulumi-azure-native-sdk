// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the Private Endpoint Connection specified by its identifier.
func LookupPrivateEndpointConnectionByName(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionByNameArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionByNameResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateEndpointConnectionByNameResult
	err := ctx.Invoke("azure-native:apimanagement/v20230301preview:getPrivateEndpointConnectionByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionByNameArgs struct {
	// Name of the private endpoint connection.
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The Private Endpoint Connection resource.
type LookupPrivateEndpointConnectionByNameResult struct {
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
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupPrivateEndpointConnectionByNameOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionByNameOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionByNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionByNameResult, error) {
			args := v.(LookupPrivateEndpointConnectionByNameArgs)
			r, err := LookupPrivateEndpointConnectionByName(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionByNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionByNameResultOutput)
}

type LookupPrivateEndpointConnectionByNameOutputArgs struct {
	// Name of the private endpoint connection.
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupPrivateEndpointConnectionByNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionByNameArgs)(nil)).Elem()
}

// The Private Endpoint Connection resource.
type LookupPrivateEndpointConnectionByNameResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionByNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionByNameResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionByNameResultOutput) ToLookupPrivateEndpointConnectionByNameResultOutput() LookupPrivateEndpointConnectionByNameResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionByNameResultOutput) ToLookupPrivateEndpointConnectionByNameResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionByNameResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPrivateEndpointConnectionByNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByNameResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPrivateEndpointConnectionByNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByNameResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o LookupPrivateEndpointConnectionByNameResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByNameResult) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o LookupPrivateEndpointConnectionByNameResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByNameResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o LookupPrivateEndpointConnectionByNameResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByNameResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPrivateEndpointConnectionByNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionByNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionByNameResultOutput{})
}
