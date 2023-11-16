// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified peering for the ExpressRouteCrossConnection.
func LookupExpressRouteCrossConnectionPeering(ctx *pulumi.Context, args *LookupExpressRouteCrossConnectionPeeringArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCrossConnectionPeeringResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExpressRouteCrossConnectionPeeringResult
	err := ctx.Invoke("azure-native:network/v20230501:getExpressRouteCrossConnectionPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCrossConnectionPeeringArgs struct {
	// The name of the ExpressRouteCrossConnection.
	CrossConnectionName string `pulumi:"crossConnectionName"`
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Peering in an ExpressRoute Cross Connection resource.
type LookupExpressRouteCrossConnectionPeeringResult struct {
	// The Azure ASN.
	AzureASN int `pulumi:"azureASN"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The GatewayManager Etag.
	GatewayManagerEtag *string `pulumi:"gatewayManagerEtag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The IPv6 peering configuration.
	Ipv6PeeringConfig *Ipv6ExpressRouteCircuitPeeringConfigResponse `pulumi:"ipv6PeeringConfig"`
	// Who was the last to modify the peering.
	LastModifiedBy string `pulumi:"lastModifiedBy"`
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig *ExpressRouteCircuitPeeringConfigResponse `pulumi:"microsoftPeeringConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The peer ASN.
	PeerASN *float64 `pulumi:"peerASN"`
	// The peering type.
	PeeringType *string `pulumi:"peeringType"`
	// The primary port.
	PrimaryAzurePort string `pulumi:"primaryAzurePort"`
	// The primary address prefix.
	PrimaryPeerAddressPrefix *string `pulumi:"primaryPeerAddressPrefix"`
	// The provisioning state of the express route cross connection peering resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The secondary port.
	SecondaryAzurePort string `pulumi:"secondaryAzurePort"`
	// The secondary address prefix.
	SecondaryPeerAddressPrefix *string `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key.
	SharedKey *string `pulumi:"sharedKey"`
	// The peering state.
	State *string `pulumi:"state"`
	// The VLAN ID.
	VlanId *int `pulumi:"vlanId"`
}

func LookupExpressRouteCrossConnectionPeeringOutput(ctx *pulumi.Context, args LookupExpressRouteCrossConnectionPeeringOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRouteCrossConnectionPeeringResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRouteCrossConnectionPeeringResult, error) {
			args := v.(LookupExpressRouteCrossConnectionPeeringArgs)
			r, err := LookupExpressRouteCrossConnectionPeering(ctx, &args, opts...)
			var s LookupExpressRouteCrossConnectionPeeringResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRouteCrossConnectionPeeringResultOutput)
}

type LookupExpressRouteCrossConnectionPeeringOutputArgs struct {
	// The name of the ExpressRouteCrossConnection.
	CrossConnectionName pulumi.StringInput `pulumi:"crossConnectionName"`
	// The name of the peering.
	PeeringName pulumi.StringInput `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRouteCrossConnectionPeeringOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCrossConnectionPeeringArgs)(nil)).Elem()
}

// Peering in an ExpressRoute Cross Connection resource.
type LookupExpressRouteCrossConnectionPeeringResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRouteCrossConnectionPeeringResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCrossConnectionPeeringResult)(nil)).Elem()
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) ToLookupExpressRouteCrossConnectionPeeringResultOutput() LookupExpressRouteCrossConnectionPeeringResultOutput {
	return o
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) ToLookupExpressRouteCrossConnectionPeeringResultOutputWithContext(ctx context.Context) LookupExpressRouteCrossConnectionPeeringResultOutput {
	return o
}

// The Azure ASN.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) AzureASN() pulumi.IntOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) int { return v.AzureASN }).(pulumi.IntOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The GatewayManager Etag.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The IPv6 peering configuration.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) Ipv6PeeringConfig() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *Ipv6ExpressRouteCircuitPeeringConfigResponse {
		return v.Ipv6PeeringConfig
	}).(Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

// Who was the last to modify the peering.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

// The Microsoft peering configuration.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *ExpressRouteCircuitPeeringConfigResponse {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The peer ASN.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) PeerASN() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *float64 { return v.PeerASN }).(pulumi.Float64PtrOutput)
}

// The peering type.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) PeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.PeeringType }).(pulumi.StringPtrOutput)
}

// The primary port.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) PrimaryAzurePort() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) string { return v.PrimaryAzurePort }).(pulumi.StringOutput)
}

// The primary address prefix.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

// The provisioning state of the express route cross connection peering resource.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The secondary port.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) SecondaryAzurePort() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) string { return v.SecondaryAzurePort }).(pulumi.StringOutput)
}

// The secondary address prefix.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.SecondaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

// The shared key.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

// The peering state.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// The VLAN ID.
func (o LookupExpressRouteCrossConnectionPeeringResultOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *int { return v.VlanId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRouteCrossConnectionPeeringResultOutput{})
}
