// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified private endpoint connection associated with the key vault.
func LookupPrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:keyvault/v20230201:getPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionArgs struct {
	// Name of the private endpoint connection associated with the key vault.
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// Name of the resource group that contains the key vault.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the key vault.
	VaultName string `pulumi:"vaultName"`
}

// Private endpoint connection resource.
type LookupPrivateEndpointConnectionResult struct {
	// Modified whenever there is a change in the state of private endpoint connection.
	Etag *string `pulumi:"etag"`
	// Fully qualified identifier of the key vault resource.
	Id string `pulumi:"id"`
	// Azure location of the key vault resource.
	Location string `pulumi:"location"`
	// Name of the key vault resource.
	Name string `pulumi:"name"`
	// Properties of the private endpoint object.
	PrivateEndpoint *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// Approval state of the private link connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the private endpoint connection.
	ProvisioningState string `pulumi:"provisioningState"`
	// Tags assigned to the key vault resource.
	Tags map[string]string `pulumi:"tags"`
	// Resource type of the key vault resource.
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
	// Name of the private endpoint connection associated with the key vault.
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// Name of the resource group that contains the key vault.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the key vault.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionArgs)(nil)).Elem()
}

// Private endpoint connection resource.
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

// Modified whenever there is a change in the state of private endpoint connection.
func (o LookupPrivateEndpointConnectionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified identifier of the key vault resource.
func (o LookupPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Azure location of the key vault resource.
func (o LookupPrivateEndpointConnectionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of the key vault resource.
func (o LookupPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the private endpoint object.
func (o LookupPrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

// Approval state of the private link connection.
func (o LookupPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

// Provisioning state of the private endpoint connection.
func (o LookupPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Tags assigned to the key vault resource.
func (o LookupPrivateEndpointConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type of the key vault resource.
func (o LookupPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionResultOutput{})
}
