// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a specific private endpoint connection under a topic, domain, or partner namespace.
// Azure REST API version: 2022-06-15.
//
// Other available API versions: 2023-06-01-preview.
func LookupPrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:eventgrid:getPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionArgs struct {
	// The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name).
	ParentName string `pulumi:"parentName"`
	// The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\'.
	ParentType string `pulumi:"parentType"`
	// The name of the private endpoint connection connection.
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupPrivateEndpointConnectionResult struct {
	// GroupIds from the private link service resource.
	GroupIds []string `pulumi:"groupIds"`
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// The Private Endpoint resource for this Connection.
	PrivateEndpoint *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// Details about the state of the connection.
	PrivateLinkServiceConnectionState *ConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the Private Endpoint Connection.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

func LookupPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionResult, error) {
			args := v.(LookupPrivateEndpointConnectionArgs)
			r, err := LookupPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionResultOutput)
}

type LookupPrivateEndpointConnectionOutputArgs struct {
	// The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name).
	ParentName pulumi.StringInput `pulumi:"parentName"`
	// The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\'.
	ParentType pulumi.StringInput `pulumi:"parentType"`
	// The name of the private endpoint connection connection.
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionArgs)(nil)).Elem()
}

type LookupPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionResultOutput) ToLookupPrivateEndpointConnectionResultOutput() LookupPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionResultOutput) ToLookupPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPrivateEndpointConnectionResult] {
	return pulumix.Output[LookupPrivateEndpointConnectionResult]{
		OutputState: o.OutputState,
	}
}

// GroupIds from the private link service resource.
func (o LookupPrivateEndpointConnectionResultOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// Fully qualified identifier of the resource.
func (o LookupPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the resource.
func (o LookupPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Private Endpoint resource for this Connection.
func (o LookupPrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

// Details about the state of the connection.
func (o LookupPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() ConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) *ConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionStateResponsePtrOutput)
}

// Provisioning state of the Private Endpoint Connection.
func (o LookupPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Type of the resource.
func (o LookupPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionResultOutput{})
}
