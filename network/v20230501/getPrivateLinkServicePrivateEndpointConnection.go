// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the specific private end point connection by specific private link service in the resource group.
func LookupPrivateLinkServicePrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateLinkServicePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkServicePrivateEndpointConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateLinkServicePrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:network/v20230501:getPrivateLinkServicePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPrivateLinkServicePrivateEndpointConnectionArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the private end point connection.
	PeConnectionName string `pulumi:"peConnectionName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the private link service.
	ServiceName string `pulumi:"serviceName"`
}

// PrivateEndpointConnection resource.
type LookupPrivateLinkServicePrivateEndpointConnectionResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The consumer link id.
	LinkIdentifier string `pulumi:"linkIdentifier"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// The location of the private endpoint.
	PrivateEndpointLocation string `pulumi:"privateEndpointLocation"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupPrivateLinkServicePrivateEndpointConnectionResult
func (val *LookupPrivateLinkServicePrivateEndpointConnectionResult) Defaults() *LookupPrivateLinkServicePrivateEndpointConnectionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PrivateEndpoint = *tmp.PrivateEndpoint.Defaults()

	return &tmp
}

func LookupPrivateLinkServicePrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupPrivateLinkServicePrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkServicePrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkServicePrivateEndpointConnectionResult, error) {
			args := v.(LookupPrivateLinkServicePrivateEndpointConnectionArgs)
			r, err := LookupPrivateLinkServicePrivateEndpointConnection(ctx, &args, opts...)
			var s LookupPrivateLinkServicePrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateLinkServicePrivateEndpointConnectionResultOutput)
}

type LookupPrivateLinkServicePrivateEndpointConnectionOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the private end point connection.
	PeConnectionName pulumi.StringInput `pulumi:"peConnectionName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the private link service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupPrivateLinkServicePrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkServicePrivateEndpointConnectionArgs)(nil)).Elem()
}

// PrivateEndpointConnection resource.
type LookupPrivateLinkServicePrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkServicePrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) ToLookupPrivateLinkServicePrivateEndpointConnectionResultOutput() LookupPrivateLinkServicePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) ToLookupPrivateLinkServicePrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupPrivateLinkServicePrivateEndpointConnectionResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The consumer link id.
func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) LinkIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) string { return v.LinkIdentifier }).(pulumi.StringOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The resource of private end point.
func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponseOutput)
}

// The location of the private endpoint.
func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) PrivateEndpointLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) string {
		return v.PrivateEndpointLocation
	}).(pulumi.StringOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource type.
func (o LookupPrivateLinkServicePrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicePrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkServicePrivateEndpointConnectionResultOutput{})
}
