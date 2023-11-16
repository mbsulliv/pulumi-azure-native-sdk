// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VpnGateway Resource.
type VpnGateway struct {
	pulumi.CustomResourceState

	// Local network gateway's BGP speaker settings.
	BgpSettings BgpSettingsResponsePtrOutput `pulumi:"bgpSettings"`
	// list of all vpn connections to the gateway.
	Connections VpnConnectionResponseArrayOutput `pulumi:"connections"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The policies applied to this vpn gateway.
	Policies PoliciesResponsePtrOutput `pulumi:"policies"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The VirtualHub to which the gateway belongs
	VirtualHub SubResourceResponsePtrOutput `pulumi:"virtualHub"`
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
			Type: pulumi.String("azure-native:network:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VpnGateway"),
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
	err := ctx.RegisterResource("azure-native:network/v20180701:VpnGateway", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:network/v20180701:VpnGateway", name, id, state, &resource, opts...)
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
	// list of all vpn connections to the gateway.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Connections []VpnConnectionType `pulumi:"connections"`
	// The name of the gateway.
	GatewayName *string `pulumi:"gatewayName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The policies applied to this vpn gateway.
	Policies *Policies `pulumi:"policies"`
	// The resource group name of the VpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The VirtualHub to which the gateway belongs
	VirtualHub *SubResource `pulumi:"virtualHub"`
}

// The set of arguments for constructing a VpnGateway resource.
type VpnGatewayArgs struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings BgpSettingsPtrInput
	// list of all vpn connections to the gateway.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Connections VpnConnectionTypeArrayInput
	// The name of the gateway.
	GatewayName pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The policies applied to this vpn gateway.
	Policies PoliciesPtrInput
	// The resource group name of the VpnGateway.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The VirtualHub to which the gateway belongs
	VirtualHub SubResourcePtrInput
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

// list of all vpn connections to the gateway.
func (o VpnGatewayOutput) Connections() VpnConnectionResponseArrayOutput {
	return o.ApplyT(func(v *VpnGateway) VpnConnectionResponseArrayOutput { return v.Connections }).(VpnConnectionResponseArrayOutput)
}

// Gets a unique read-only string that changes whenever the resource is updated.
func (o VpnGatewayOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnGateway) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource location.
func (o VpnGatewayOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnGateway) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o VpnGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The policies applied to this vpn gateway.
func (o VpnGatewayOutput) Policies() PoliciesResponsePtrOutput {
	return o.ApplyT(func(v *VpnGateway) PoliciesResponsePtrOutput { return v.Policies }).(PoliciesResponsePtrOutput)
}

// The provisioning state of the resource.
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

// The VirtualHub to which the gateway belongs
func (o VpnGatewayOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VpnGateway) SubResourceResponsePtrOutput { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnGatewayOutput{})
}
