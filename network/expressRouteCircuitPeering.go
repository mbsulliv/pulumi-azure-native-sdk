// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Peering in an ExpressRouteCircuit resource.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01.
//
// Other available API versions: 2017-09-01, 2019-02-01, 2019-06-01, 2019-08-01, 2023-04-01, 2023-05-01.
type ExpressRouteCircuitPeering struct {
	pulumi.CustomResourceState

	// The Azure ASN.
	AzureASN pulumi.IntPtrOutput `pulumi:"azureASN"`
	// The list of circuit connections associated with Azure Private Peering for this circuit.
	Connections ExpressRouteCircuitConnectionResponseArrayOutput `pulumi:"connections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The ExpressRoute connection.
	ExpressRouteConnection ExpressRouteConnectionIdResponsePtrOutput `pulumi:"expressRouteConnection"`
	// The GatewayManager Etag.
	GatewayManagerEtag pulumi.StringPtrOutput `pulumi:"gatewayManagerEtag"`
	// The IPv6 peering configuration.
	Ipv6PeeringConfig Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput `pulumi:"ipv6PeeringConfig"`
	// Who was the last to modify the peering.
	LastModifiedBy pulumi.StringOutput `pulumi:"lastModifiedBy"`
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig ExpressRouteCircuitPeeringConfigResponsePtrOutput `pulumi:"microsoftPeeringConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The peer ASN.
	PeerASN pulumi.Float64PtrOutput `pulumi:"peerASN"`
	// The list of peered circuit connections associated with Azure Private Peering for this circuit.
	PeeredConnections PeerExpressRouteCircuitConnectionResponseArrayOutput `pulumi:"peeredConnections"`
	// The peering type.
	PeeringType pulumi.StringPtrOutput `pulumi:"peeringType"`
	// The primary port.
	PrimaryAzurePort pulumi.StringPtrOutput `pulumi:"primaryAzurePort"`
	// The primary address prefix.
	PrimaryPeerAddressPrefix pulumi.StringPtrOutput `pulumi:"primaryPeerAddressPrefix"`
	// The provisioning state of the express route circuit peering resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The reference to the RouteFilter resource.
	RouteFilter SubResourceResponsePtrOutput `pulumi:"routeFilter"`
	// The secondary port.
	SecondaryAzurePort pulumi.StringPtrOutput `pulumi:"secondaryAzurePort"`
	// The secondary address prefix.
	SecondaryPeerAddressPrefix pulumi.StringPtrOutput `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key.
	SharedKey pulumi.StringPtrOutput `pulumi:"sharedKey"`
	// The peering state.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The peering stats of express route circuit.
	Stats ExpressRouteCircuitStatsResponsePtrOutput `pulumi:"stats"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The VLAN ID.
	VlanId pulumi.IntPtrOutput `pulumi:"vlanId"`
}

// NewExpressRouteCircuitPeering registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuitPeering(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitPeeringArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitPeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CircuitName == nil {
		return nil, errors.New("invalid value for required argument 'CircuitName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20150501preview:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:ExpressRouteCircuitPeering"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ExpressRouteCircuitPeering
	err := ctx.RegisterResource("azure-native:network:ExpressRouteCircuitPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuitPeering gets an existing ExpressRouteCircuitPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuitPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitPeeringState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitPeering, error) {
	var resource ExpressRouteCircuitPeering
	err := ctx.ReadResource("azure-native:network:ExpressRouteCircuitPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuitPeering resources.
type expressRouteCircuitPeeringState struct {
}

type ExpressRouteCircuitPeeringState struct {
}

func (ExpressRouteCircuitPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitPeeringState)(nil)).Elem()
}

type expressRouteCircuitPeeringArgs struct {
	// The Azure ASN.
	AzureASN *int `pulumi:"azureASN"`
	// The name of the express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// The list of circuit connections associated with Azure Private Peering for this circuit.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Connections []ExpressRouteCircuitConnectionType `pulumi:"connections"`
	// The GatewayManager Etag.
	GatewayManagerEtag *string `pulumi:"gatewayManagerEtag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The IPv6 peering configuration.
	Ipv6PeeringConfig *Ipv6ExpressRouteCircuitPeeringConfig `pulumi:"ipv6PeeringConfig"`
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig *ExpressRouteCircuitPeeringConfig `pulumi:"microsoftPeeringConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The peer ASN.
	PeerASN *float64 `pulumi:"peerASN"`
	// The name of the peering.
	PeeringName *string `pulumi:"peeringName"`
	// The peering type.
	PeeringType *string `pulumi:"peeringType"`
	// The primary port.
	PrimaryAzurePort *string `pulumi:"primaryAzurePort"`
	// The primary address prefix.
	PrimaryPeerAddressPrefix *string `pulumi:"primaryPeerAddressPrefix"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The reference to the RouteFilter resource.
	RouteFilter *SubResource `pulumi:"routeFilter"`
	// The secondary port.
	SecondaryAzurePort *string `pulumi:"secondaryAzurePort"`
	// The secondary address prefix.
	SecondaryPeerAddressPrefix *string `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key.
	SharedKey *string `pulumi:"sharedKey"`
	// The peering state.
	State *string `pulumi:"state"`
	// The peering stats of express route circuit.
	Stats *ExpressRouteCircuitStats `pulumi:"stats"`
	// The VLAN ID.
	VlanId *int `pulumi:"vlanId"`
}

// The set of arguments for constructing a ExpressRouteCircuitPeering resource.
type ExpressRouteCircuitPeeringArgs struct {
	// The Azure ASN.
	AzureASN pulumi.IntPtrInput
	// The name of the express route circuit.
	CircuitName pulumi.StringInput
	// The list of circuit connections associated with Azure Private Peering for this circuit.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Connections ExpressRouteCircuitConnectionTypeArrayInput
	// The GatewayManager Etag.
	GatewayManagerEtag pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The IPv6 peering configuration.
	Ipv6PeeringConfig Ipv6ExpressRouteCircuitPeeringConfigPtrInput
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig ExpressRouteCircuitPeeringConfigPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The peer ASN.
	PeerASN pulumi.Float64PtrInput
	// The name of the peering.
	PeeringName pulumi.StringPtrInput
	// The peering type.
	PeeringType pulumi.StringPtrInput
	// The primary port.
	PrimaryAzurePort pulumi.StringPtrInput
	// The primary address prefix.
	PrimaryPeerAddressPrefix pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The reference to the RouteFilter resource.
	RouteFilter SubResourcePtrInput
	// The secondary port.
	SecondaryAzurePort pulumi.StringPtrInput
	// The secondary address prefix.
	SecondaryPeerAddressPrefix pulumi.StringPtrInput
	// The shared key.
	SharedKey pulumi.StringPtrInput
	// The peering state.
	State pulumi.StringPtrInput
	// The peering stats of express route circuit.
	Stats ExpressRouteCircuitStatsPtrInput
	// The VLAN ID.
	VlanId pulumi.IntPtrInput
}

func (ExpressRouteCircuitPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitPeeringArgs)(nil)).Elem()
}

type ExpressRouteCircuitPeeringInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringOutput() ExpressRouteCircuitPeeringOutput
	ToExpressRouteCircuitPeeringOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringOutput
}

func (*ExpressRouteCircuitPeering) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeering)(nil)).Elem()
}

func (i *ExpressRouteCircuitPeering) ToExpressRouteCircuitPeeringOutput() ExpressRouteCircuitPeeringOutput {
	return i.ToExpressRouteCircuitPeeringOutputWithContext(context.Background())
}

func (i *ExpressRouteCircuitPeering) ToExpressRouteCircuitPeeringOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringOutput)
}

func (i *ExpressRouteCircuitPeering) ToOutput(ctx context.Context) pulumix.Output[*ExpressRouteCircuitPeering] {
	return pulumix.Output[*ExpressRouteCircuitPeering]{
		OutputState: i.ToExpressRouteCircuitPeeringOutputWithContext(ctx).OutputState,
	}
}

type ExpressRouteCircuitPeeringOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeering)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringOutput) ToExpressRouteCircuitPeeringOutput() ExpressRouteCircuitPeeringOutput {
	return o
}

func (o ExpressRouteCircuitPeeringOutput) ToExpressRouteCircuitPeeringOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringOutput {
	return o
}

func (o ExpressRouteCircuitPeeringOutput) ToOutput(ctx context.Context) pulumix.Output[*ExpressRouteCircuitPeering] {
	return pulumix.Output[*ExpressRouteCircuitPeering]{
		OutputState: o.OutputState,
	}
}

// The Azure ASN.
func (o ExpressRouteCircuitPeeringOutput) AzureASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.IntPtrOutput { return v.AzureASN }).(pulumi.IntPtrOutput)
}

// The list of circuit connections associated with Azure Private Peering for this circuit.
func (o ExpressRouteCircuitPeeringOutput) Connections() ExpressRouteCircuitConnectionResponseArrayOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) ExpressRouteCircuitConnectionResponseArrayOutput {
		return v.Connections
	}).(ExpressRouteCircuitConnectionResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o ExpressRouteCircuitPeeringOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The ExpressRoute connection.
func (o ExpressRouteCircuitPeeringOutput) ExpressRouteConnection() ExpressRouteConnectionIdResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) ExpressRouteConnectionIdResponsePtrOutput {
		return v.ExpressRouteConnection
	}).(ExpressRouteConnectionIdResponsePtrOutput)
}

// The GatewayManager Etag.
func (o ExpressRouteCircuitPeeringOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.StringPtrOutput { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

// The IPv6 peering configuration.
func (o ExpressRouteCircuitPeeringOutput) Ipv6PeeringConfig() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
		return v.Ipv6PeeringConfig
	}).(Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

// Who was the last to modify the peering.
func (o ExpressRouteCircuitPeeringOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.StringOutput { return v.LastModifiedBy }).(pulumi.StringOutput)
}

// The Microsoft peering configuration.
func (o ExpressRouteCircuitPeeringOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) ExpressRouteCircuitPeeringConfigResponsePtrOutput {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o ExpressRouteCircuitPeeringOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The peer ASN.
func (o ExpressRouteCircuitPeeringOutput) PeerASN() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.Float64PtrOutput { return v.PeerASN }).(pulumi.Float64PtrOutput)
}

// The list of peered circuit connections associated with Azure Private Peering for this circuit.
func (o ExpressRouteCircuitPeeringOutput) PeeredConnections() PeerExpressRouteCircuitConnectionResponseArrayOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) PeerExpressRouteCircuitConnectionResponseArrayOutput {
		return v.PeeredConnections
	}).(PeerExpressRouteCircuitConnectionResponseArrayOutput)
}

// The peering type.
func (o ExpressRouteCircuitPeeringOutput) PeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.StringPtrOutput { return v.PeeringType }).(pulumi.StringPtrOutput)
}

// The primary port.
func (o ExpressRouteCircuitPeeringOutput) PrimaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.StringPtrOutput { return v.PrimaryAzurePort }).(pulumi.StringPtrOutput)
}

// The primary address prefix.
func (o ExpressRouteCircuitPeeringOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.StringPtrOutput { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

// The provisioning state of the express route circuit peering resource.
func (o ExpressRouteCircuitPeeringOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The reference to the RouteFilter resource.
func (o ExpressRouteCircuitPeeringOutput) RouteFilter() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) SubResourceResponsePtrOutput { return v.RouteFilter }).(SubResourceResponsePtrOutput)
}

// The secondary port.
func (o ExpressRouteCircuitPeeringOutput) SecondaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.StringPtrOutput { return v.SecondaryAzurePort }).(pulumi.StringPtrOutput)
}

// The secondary address prefix.
func (o ExpressRouteCircuitPeeringOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.StringPtrOutput { return v.SecondaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

// The shared key.
func (o ExpressRouteCircuitPeeringOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.StringPtrOutput { return v.SharedKey }).(pulumi.StringPtrOutput)
}

// The peering state.
func (o ExpressRouteCircuitPeeringOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// The peering stats of express route circuit.
func (o ExpressRouteCircuitPeeringOutput) Stats() ExpressRouteCircuitStatsResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) ExpressRouteCircuitStatsResponsePtrOutput { return v.Stats }).(ExpressRouteCircuitStatsResponsePtrOutput)
}

// Type of the resource.
func (o ExpressRouteCircuitPeeringOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The VLAN ID.
func (o ExpressRouteCircuitPeeringOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) pulumi.IntPtrOutput { return v.VlanId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringOutput{})
}
