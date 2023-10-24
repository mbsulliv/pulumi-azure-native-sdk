// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified Virtual Router Peering.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01.
func LookupVirtualRouterPeering(ctx *pulumi.Context, args *LookupVirtualRouterPeeringArgs, opts ...pulumi.InvokeOption) (*LookupVirtualRouterPeeringResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualRouterPeeringResult
	err := ctx.Invoke("azure-native:network:getVirtualRouterPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualRouterPeeringArgs struct {
	// The name of the Virtual Router Peering.
	PeeringName string `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Virtual Router.
	VirtualRouterName string `pulumi:"virtualRouterName"`
}

// Virtual Router Peering resource.
type LookupVirtualRouterPeeringResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Name of the virtual router peering that is unique within a virtual router.
	Name *string `pulumi:"name"`
	// Peer ASN.
	PeerAsn *float64 `pulumi:"peerAsn"`
	// Peer IP.
	PeerIp *string `pulumi:"peerIp"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Peering type.
	Type string `pulumi:"type"`
}

func LookupVirtualRouterPeeringOutput(ctx *pulumi.Context, args LookupVirtualRouterPeeringOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualRouterPeeringResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualRouterPeeringResult, error) {
			args := v.(LookupVirtualRouterPeeringArgs)
			r, err := LookupVirtualRouterPeering(ctx, &args, opts...)
			var s LookupVirtualRouterPeeringResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualRouterPeeringResultOutput)
}

type LookupVirtualRouterPeeringOutputArgs struct {
	// The name of the Virtual Router Peering.
	PeeringName pulumi.StringInput `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Virtual Router.
	VirtualRouterName pulumi.StringInput `pulumi:"virtualRouterName"`
}

func (LookupVirtualRouterPeeringOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualRouterPeeringArgs)(nil)).Elem()
}

// Virtual Router Peering resource.
type LookupVirtualRouterPeeringResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualRouterPeeringResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualRouterPeeringResult)(nil)).Elem()
}

func (o LookupVirtualRouterPeeringResultOutput) ToLookupVirtualRouterPeeringResultOutput() LookupVirtualRouterPeeringResultOutput {
	return o
}

func (o LookupVirtualRouterPeeringResultOutput) ToLookupVirtualRouterPeeringResultOutputWithContext(ctx context.Context) LookupVirtualRouterPeeringResultOutput {
	return o
}

func (o LookupVirtualRouterPeeringResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualRouterPeeringResult] {
	return pulumix.Output[LookupVirtualRouterPeeringResult]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVirtualRouterPeeringResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupVirtualRouterPeeringResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Name of the virtual router peering that is unique within a virtual router.
func (o LookupVirtualRouterPeeringResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Peer ASN.
func (o LookupVirtualRouterPeeringResultOutput) PeerAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) *float64 { return v.PeerAsn }).(pulumi.Float64PtrOutput)
}

// Peer IP.
func (o LookupVirtualRouterPeeringResultOutput) PeerIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) *string { return v.PeerIp }).(pulumi.StringPtrOutput)
}

// The provisioning state of the resource.
func (o LookupVirtualRouterPeeringResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Peering type.
func (o LookupVirtualRouterPeeringResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualRouterPeeringResultOutput{})
}
