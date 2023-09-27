// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A common class for general resource information.
type LocalNetworkGateway struct {
	pulumi.CustomResourceState

	// Local network gateway's BGP speaker settings.
	BgpSettings BgpSettingsResponsePtrOutput `pulumi:"bgpSettings"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// FQDN of local network gateway.
	Fqdn pulumi.StringPtrOutput `pulumi:"fqdn"`
	// IP address of local network gateway.
	GatewayIpAddress pulumi.StringPtrOutput `pulumi:"gatewayIpAddress"`
	// Local network site address space.
	LocalNetworkAddressSpace AddressSpaceResponsePtrOutput `pulumi:"localNetworkAddressSpace"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the local network gateway resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource GUID property of the local network gateway resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLocalNetworkGateway registers a new resource with the given unique name, arguments, and options.
func NewLocalNetworkGateway(ctx *pulumi.Context,
	name string, args *LocalNetworkGatewayArgs, opts ...pulumi.ResourceOption) (*LocalNetworkGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:LocalNetworkGateway"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LocalNetworkGateway
	err := ctx.RegisterResource("azure-native:network/v20230401:LocalNetworkGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalNetworkGateway gets an existing LocalNetworkGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalNetworkGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalNetworkGatewayState, opts ...pulumi.ResourceOption) (*LocalNetworkGateway, error) {
	var resource LocalNetworkGateway
	err := ctx.ReadResource("azure-native:network/v20230401:LocalNetworkGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalNetworkGateway resources.
type localNetworkGatewayState struct {
}

type LocalNetworkGatewayState struct {
}

func (LocalNetworkGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*localNetworkGatewayState)(nil)).Elem()
}

type localNetworkGatewayArgs struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings *BgpSettings `pulumi:"bgpSettings"`
	// FQDN of local network gateway.
	Fqdn *string `pulumi:"fqdn"`
	// IP address of local network gateway.
	GatewayIpAddress *string `pulumi:"gatewayIpAddress"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Local network site address space.
	LocalNetworkAddressSpace *AddressSpace `pulumi:"localNetworkAddressSpace"`
	// The name of the local network gateway.
	LocalNetworkGatewayName *string `pulumi:"localNetworkGatewayName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LocalNetworkGateway resource.
type LocalNetworkGatewayArgs struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings BgpSettingsPtrInput
	// FQDN of local network gateway.
	Fqdn pulumi.StringPtrInput
	// IP address of local network gateway.
	GatewayIpAddress pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Local network site address space.
	LocalNetworkAddressSpace AddressSpacePtrInput
	// The name of the local network gateway.
	LocalNetworkGatewayName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (LocalNetworkGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localNetworkGatewayArgs)(nil)).Elem()
}

type LocalNetworkGatewayInput interface {
	pulumi.Input

	ToLocalNetworkGatewayOutput() LocalNetworkGatewayOutput
	ToLocalNetworkGatewayOutputWithContext(ctx context.Context) LocalNetworkGatewayOutput
}

func (*LocalNetworkGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNetworkGateway)(nil)).Elem()
}

func (i *LocalNetworkGateway) ToLocalNetworkGatewayOutput() LocalNetworkGatewayOutput {
	return i.ToLocalNetworkGatewayOutputWithContext(context.Background())
}

func (i *LocalNetworkGateway) ToLocalNetworkGatewayOutputWithContext(ctx context.Context) LocalNetworkGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNetworkGatewayOutput)
}

func (i *LocalNetworkGateway) ToOutput(ctx context.Context) pulumix.Output[*LocalNetworkGateway] {
	return pulumix.Output[*LocalNetworkGateway]{
		OutputState: i.ToLocalNetworkGatewayOutputWithContext(ctx).OutputState,
	}
}

type LocalNetworkGatewayOutput struct{ *pulumi.OutputState }

func (LocalNetworkGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNetworkGateway)(nil)).Elem()
}

func (o LocalNetworkGatewayOutput) ToLocalNetworkGatewayOutput() LocalNetworkGatewayOutput {
	return o
}

func (o LocalNetworkGatewayOutput) ToLocalNetworkGatewayOutputWithContext(ctx context.Context) LocalNetworkGatewayOutput {
	return o
}

func (o LocalNetworkGatewayOutput) ToOutput(ctx context.Context) pulumix.Output[*LocalNetworkGateway] {
	return pulumix.Output[*LocalNetworkGateway]{
		OutputState: o.OutputState,
	}
}

// Local network gateway's BGP speaker settings.
func (o LocalNetworkGatewayOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) BgpSettingsResponsePtrOutput { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LocalNetworkGatewayOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// FQDN of local network gateway.
func (o LocalNetworkGatewayOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringPtrOutput { return v.Fqdn }).(pulumi.StringPtrOutput)
}

// IP address of local network gateway.
func (o LocalNetworkGatewayOutput) GatewayIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringPtrOutput { return v.GatewayIpAddress }).(pulumi.StringPtrOutput)
}

// Local network site address space.
func (o LocalNetworkGatewayOutput) LocalNetworkAddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) AddressSpaceResponsePtrOutput { return v.LocalNetworkAddressSpace }).(AddressSpaceResponsePtrOutput)
}

// Resource location.
func (o LocalNetworkGatewayOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LocalNetworkGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the local network gateway resource.
func (o LocalNetworkGatewayOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the local network gateway resource.
func (o LocalNetworkGatewayOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Resource tags.
func (o LocalNetworkGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LocalNetworkGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LocalNetworkGatewayOutput{})
}
