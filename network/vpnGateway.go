// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VpnGateway Resource.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01.
//
// Other available API versions: 2018-07-01, 2023-04-01, 2023-05-01, 2023-06-01.
type VpnGateway struct {
	pulumi.CustomResourceState

	// Local network gateway's BGP speaker settings.
	BgpSettings BgpSettingsResponsePtrOutput `pulumi:"bgpSettings"`
	// List of all vpn connections to the gateway.
	Connections VpnConnectionResponseArrayOutput `pulumi:"connections"`
	// Enable BGP routes translation for NAT on this VpnGateway.
	EnableBgpRouteTranslationForNat pulumi.BoolPtrOutput `pulumi:"enableBgpRouteTranslationForNat"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// List of all IPs configured on the gateway.
	IpConfigurations VpnGatewayIpConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// Enable Routing Preference property for the Public IP Interface of the VpnGateway.
	IsRoutingPreferenceInternet pulumi.BoolPtrOutput `pulumi:"isRoutingPreferenceInternet"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of all the nat Rules associated with the gateway.
	NatRules VpnGatewayNatRuleResponseArrayOutput `pulumi:"natRules"`
	// The provisioning state of the VPN gateway resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The VirtualHub to which the gateway belongs.
	VirtualHub SubResourceResponsePtrOutput `pulumi:"virtualHub"`
	// The scale unit for this vpn gateway.
	VpnGatewayScaleUnit pulumi.IntPtrOutput `pulumi:"vpnGatewayScaleUnit"`
}

// NewVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewVpnGateway(ctx *pulumi.Context,
	name string, args *VpnGatewayArgs, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20180401:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:VpnGateway"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VpnGateway
	err := ctx.RegisterResource("azure-native:network:VpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnGateway gets an existing VpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnGatewayState, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
	var resource VpnGateway
	err := ctx.ReadResource("azure-native:network:VpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnGateway resources.
type vpnGatewayState struct {
}

type VpnGatewayState struct {
}

func (VpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayState)(nil)).Elem()
}

type vpnGatewayArgs struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings *BgpSettings `pulumi:"bgpSettings"`
	// List of all vpn connections to the gateway.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Connections []VpnConnectionType `pulumi:"connections"`
	// Enable BGP routes translation for NAT on this VpnGateway.
	EnableBgpRouteTranslationForNat *bool `pulumi:"enableBgpRouteTranslationForNat"`
	// The name of the gateway.
	GatewayName *string `pulumi:"gatewayName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Enable Routing Preference property for the Public IP Interface of the VpnGateway.
	IsRoutingPreferenceInternet *bool `pulumi:"isRoutingPreferenceInternet"`
	// Resource location.
	Location *string `pulumi:"location"`
	// List of all the nat Rules associated with the gateway.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	NatRules []VpnGatewayNatRule `pulumi:"natRules"`
	// The resource group name of the VpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The VirtualHub to which the gateway belongs.
	VirtualHub *SubResource `pulumi:"virtualHub"`
	// The scale unit for this vpn gateway.
	VpnGatewayScaleUnit *int `pulumi:"vpnGatewayScaleUnit"`
}

// The set of arguments for constructing a VpnGateway resource.
type VpnGatewayArgs struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings BgpSettingsPtrInput
	// List of all vpn connections to the gateway.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Connections VpnConnectionTypeArrayInput
	// Enable BGP routes translation for NAT on this VpnGateway.
	EnableBgpRouteTranslationForNat pulumi.BoolPtrInput
	// The name of the gateway.
	GatewayName pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Enable Routing Preference property for the Public IP Interface of the VpnGateway.
	IsRoutingPreferenceInternet pulumi.BoolPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// List of all the nat Rules associated with the gateway.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	NatRules VpnGatewayNatRuleArrayInput
	// The resource group name of the VpnGateway.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The VirtualHub to which the gateway belongs.
	VirtualHub SubResourcePtrInput
	// The scale unit for this vpn gateway.
	VpnGatewayScaleUnit pulumi.IntPtrInput
}

func (VpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayArgs)(nil)).Elem()
}

type VpnGatewayInput interface {
	pulumi.Input

	ToVpnGatewayOutput() VpnGatewayOutput
	ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput
}

func (*VpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnGateway)(nil)).Elem()
}

func (i *VpnGateway) ToVpnGatewayOutput() VpnGatewayOutput {
	return i.ToVpnGatewayOutputWithContext(context.Background())
}

func (i *VpnGateway) ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayOutput)
}

type VpnGatewayOutput struct{ *pulumi.OutputState }

func (VpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnGateway)(nil)).Elem()
}

func (o VpnGatewayOutput) ToVpnGatewayOutput() VpnGatewayOutput {
	return o
}

func (o VpnGatewayOutput) ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput {
	return o
}

// Local network gateway's BGP speaker settings.
func (o VpnGatewayOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *VpnGateway) BgpSettingsResponsePtrOutput { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

// List of all vpn connections to the gateway.
func (o VpnGatewayOutput) Connections() VpnConnectionResponseArrayOutput {
	return o.ApplyT(func(v *VpnGateway) VpnConnectionResponseArrayOutput { return v.Connections }).(VpnConnectionResponseArrayOutput)
}

// Enable BGP routes translation for NAT on this VpnGateway.
func (o VpnGatewayOutput) EnableBgpRouteTranslationForNat() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpnGateway) pulumi.BoolPtrOutput { return v.EnableBgpRouteTranslationForNat }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o VpnGatewayOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnGateway) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// List of all IPs configured on the gateway.
func (o VpnGatewayOutput) IpConfigurations() VpnGatewayIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VpnGateway) VpnGatewayIpConfigurationResponseArrayOutput { return v.IpConfigurations }).(VpnGatewayIpConfigurationResponseArrayOutput)
}

// Enable Routing Preference property for the Public IP Interface of the VpnGateway.
func (o VpnGatewayOutput) IsRoutingPreferenceInternet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpnGateway) pulumi.BoolPtrOutput { return v.IsRoutingPreferenceInternet }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o VpnGatewayOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnGateway) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o VpnGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of all the nat Rules associated with the gateway.
func (o VpnGatewayOutput) NatRules() VpnGatewayNatRuleResponseArrayOutput {
	return o.ApplyT(func(v *VpnGateway) VpnGatewayNatRuleResponseArrayOutput { return v.NatRules }).(VpnGatewayNatRuleResponseArrayOutput)
}

// The provisioning state of the VPN gateway resource.
func (o VpnGatewayOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnGateway) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o VpnGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpnGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o VpnGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The VirtualHub to which the gateway belongs.
func (o VpnGatewayOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VpnGateway) SubResourceResponsePtrOutput { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

// The scale unit for this vpn gateway.
func (o VpnGatewayOutput) VpnGatewayScaleUnit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpnGateway) pulumi.IntPtrOutput { return v.VpnGatewayScaleUnit }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnGatewayOutput{})
}
