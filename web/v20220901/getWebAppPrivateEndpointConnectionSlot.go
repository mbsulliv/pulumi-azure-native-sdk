// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Gets a private endpoint connection
func LookupWebAppPrivateEndpointConnectionSlot(ctx *pulumi.Context, args *LookupWebAppPrivateEndpointConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPrivateEndpointConnectionSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppPrivateEndpointConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20220901:getWebAppPrivateEndpointConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPrivateEndpointConnectionSlotArgs struct {
	// Name of the site.
	Name string `pulumi:"name"`
	// Name of the private endpoint connection.
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the site deployment slot.
	Slot string `pulumi:"slot"`
}

// Remote Private Endpoint Connection ARM resource.
type LookupWebAppPrivateEndpointConnectionSlotResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Private IPAddresses mapped to the remote private endpoint
	IpAddresses []string `pulumi:"ipAddresses"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// PrivateEndpoint of a remote private endpoint connection
	PrivateEndpoint *ArmIdWrapperResponse `pulumi:"privateEndpoint"`
	// The state of a private link connection
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWebAppPrivateEndpointConnectionSlotOutput(ctx *pulumi.Context, args LookupWebAppPrivateEndpointConnectionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppPrivateEndpointConnectionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppPrivateEndpointConnectionSlotResult, error) {
			args := v.(LookupWebAppPrivateEndpointConnectionSlotArgs)
			r, err := LookupWebAppPrivateEndpointConnectionSlot(ctx, &args, opts...)
			var s LookupWebAppPrivateEndpointConnectionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppPrivateEndpointConnectionSlotResultOutput)
}

type LookupWebAppPrivateEndpointConnectionSlotOutputArgs struct {
	// Name of the site.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the private endpoint connection.
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the site deployment slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppPrivateEndpointConnectionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPrivateEndpointConnectionSlotArgs)(nil)).Elem()
}

// Remote Private Endpoint Connection ARM resource.
type LookupWebAppPrivateEndpointConnectionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppPrivateEndpointConnectionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPrivateEndpointConnectionSlotResult)(nil)).Elem()
}

func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) ToLookupWebAppPrivateEndpointConnectionSlotResultOutput() LookupWebAppPrivateEndpointConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) ToLookupWebAppPrivateEndpointConnectionSlotResultOutputWithContext(ctx context.Context) LookupWebAppPrivateEndpointConnectionSlotResultOutput {
	return o
}

// Resource Id.
func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Private IPAddresses mapped to the remote private endpoint
func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// Kind of resource.
func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// PrivateEndpoint of a remote private endpoint connection
func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) *ArmIdWrapperResponse {
		return v.PrivateEndpoint
	}).(ArmIdWrapperResponsePtrOutput)
}

// The state of a private link connection
func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) *PrivateLinkConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppPrivateEndpointConnectionSlotResultOutput{})
}
