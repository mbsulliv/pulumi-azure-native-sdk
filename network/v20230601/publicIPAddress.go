// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Public IP address resource.
type PublicIPAddress struct {
	pulumi.CustomResourceState

	// The DDoS protection custom policy associated with the public IP address.
	DdosSettings DdosSettingsResponsePtrOutput `pulumi:"ddosSettings"`
	// Specify what happens to the public IP address when the VM using it is deleted
	DeleteOption pulumi.StringPtrOutput `pulumi:"deleteOption"`
	// The FQDN of the DNS record associated with the public IP address.
	DnsSettings PublicIPAddressDnsSettingsResponsePtrOutput `pulumi:"dnsSettings"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The extended location of the public ip address.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The idle timeout of the public IP address.
	IdleTimeoutInMinutes pulumi.IntPtrOutput `pulumi:"idleTimeoutInMinutes"`
	// The IP address associated with the public IP address resource.
	IpAddress pulumi.StringPtrOutput `pulumi:"ipAddress"`
	// The IP configuration associated with the public IP address.
	IpConfiguration IPConfigurationResponseOutput `pulumi:"ipConfiguration"`
	// The list of tags associated with the public IP address.
	IpTags IpTagResponseArrayOutput `pulumi:"ipTags"`
	// The linked public IP address of the public IP address resource.
	LinkedPublicIPAddress PublicIPAddressResponsePtrOutput `pulumi:"linkedPublicIPAddress"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Migration phase of Public IP Address.
	MigrationPhase pulumi.StringPtrOutput `pulumi:"migrationPhase"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The NatGateway for the Public IP address.
	NatGateway NatGatewayResponsePtrOutput `pulumi:"natGateway"`
	// The provisioning state of the public IP address resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The public IP address version.
	PublicIPAddressVersion pulumi.StringPtrOutput `pulumi:"publicIPAddressVersion"`
	// The public IP address allocation method.
	PublicIPAllocationMethod pulumi.StringPtrOutput `pulumi:"publicIPAllocationMethod"`
	// The Public IP Prefix this Public IP Address should be allocated from.
	PublicIPPrefix SubResourceResponsePtrOutput `pulumi:"publicIPPrefix"`
	// The resource GUID property of the public IP address resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// The service public IP address of the public IP address resource.
	ServicePublicIPAddress PublicIPAddressResponsePtrOutput `pulumi:"servicePublicIPAddress"`
	// The public IP address SKU.
	Sku PublicIPAddressSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewPublicIPAddress registers a new resource with the given unique name, arguments, and options.
func NewPublicIPAddress(ctx *pulumi.Context,
	name string, args *PublicIPAddressArgs, opts ...pulumi.ResourceOption) (*PublicIPAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:PublicIPAddress"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PublicIPAddress
	err := ctx.RegisterResource("azure-native:network/v20230601:PublicIPAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicIPAddress gets an existing PublicIPAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicIPAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIPAddressState, opts ...pulumi.ResourceOption) (*PublicIPAddress, error) {
	var resource PublicIPAddress
	err := ctx.ReadResource("azure-native:network/v20230601:PublicIPAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicIPAddress resources.
type publicIPAddressState struct {
}

type PublicIPAddressState struct {
}

func (PublicIPAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPAddressState)(nil)).Elem()
}

type publicIPAddressArgs struct {
	// The DDoS protection custom policy associated with the public IP address.
	DdosSettings *DdosSettings `pulumi:"ddosSettings"`
	// Specify what happens to the public IP address when the VM using it is deleted
	DeleteOption *string `pulumi:"deleteOption"`
	// The FQDN of the DNS record associated with the public IP address.
	DnsSettings *PublicIPAddressDnsSettings `pulumi:"dnsSettings"`
	// The extended location of the public ip address.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The idle timeout of the public IP address.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The IP address associated with the public IP address resource.
	IpAddress *string `pulumi:"ipAddress"`
	// The list of tags associated with the public IP address.
	IpTags []IpTag `pulumi:"ipTags"`
	// The linked public IP address of the public IP address resource.
	LinkedPublicIPAddress *PublicIPAddressType `pulumi:"linkedPublicIPAddress"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Migration phase of Public IP Address.
	MigrationPhase *string `pulumi:"migrationPhase"`
	// The NatGateway for the Public IP address.
	NatGateway *NatGatewayType `pulumi:"natGateway"`
	// The public IP address version.
	PublicIPAddressVersion *string `pulumi:"publicIPAddressVersion"`
	// The public IP address allocation method.
	PublicIPAllocationMethod *string `pulumi:"publicIPAllocationMethod"`
	// The Public IP Prefix this Public IP Address should be allocated from.
	PublicIPPrefix *SubResource `pulumi:"publicIPPrefix"`
	// The name of the public IP address.
	PublicIpAddressName *string `pulumi:"publicIpAddressName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The service public IP address of the public IP address resource.
	ServicePublicIPAddress *PublicIPAddressType `pulumi:"servicePublicIPAddress"`
	// The public IP address SKU.
	Sku *PublicIPAddressSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a PublicIPAddress resource.
type PublicIPAddressArgs struct {
	// The DDoS protection custom policy associated with the public IP address.
	DdosSettings DdosSettingsPtrInput
	// Specify what happens to the public IP address when the VM using it is deleted
	DeleteOption pulumi.StringPtrInput
	// The FQDN of the DNS record associated with the public IP address.
	DnsSettings PublicIPAddressDnsSettingsPtrInput
	// The extended location of the public ip address.
	ExtendedLocation ExtendedLocationPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The idle timeout of the public IP address.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// The IP address associated with the public IP address resource.
	IpAddress pulumi.StringPtrInput
	// The list of tags associated with the public IP address.
	IpTags IpTagArrayInput
	// The linked public IP address of the public IP address resource.
	LinkedPublicIPAddress PublicIPAddressTypePtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Migration phase of Public IP Address.
	MigrationPhase pulumi.StringPtrInput
	// The NatGateway for the Public IP address.
	NatGateway NatGatewayTypePtrInput
	// The public IP address version.
	PublicIPAddressVersion pulumi.StringPtrInput
	// The public IP address allocation method.
	PublicIPAllocationMethod pulumi.StringPtrInput
	// The Public IP Prefix this Public IP Address should be allocated from.
	PublicIPPrefix SubResourcePtrInput
	// The name of the public IP address.
	PublicIpAddressName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The service public IP address of the public IP address resource.
	ServicePublicIPAddress PublicIPAddressTypePtrInput
	// The public IP address SKU.
	Sku PublicIPAddressSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones pulumi.StringArrayInput
}

func (PublicIPAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPAddressArgs)(nil)).Elem()
}

type PublicIPAddressInput interface {
	pulumi.Input

	ToPublicIPAddressOutput() PublicIPAddressOutput
	ToPublicIPAddressOutputWithContext(ctx context.Context) PublicIPAddressOutput
}

func (*PublicIPAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddress)(nil)).Elem()
}

func (i *PublicIPAddress) ToPublicIPAddressOutput() PublicIPAddressOutput {
	return i.ToPublicIPAddressOutputWithContext(context.Background())
}

func (i *PublicIPAddress) ToPublicIPAddressOutputWithContext(ctx context.Context) PublicIPAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressOutput)
}

type PublicIPAddressOutput struct{ *pulumi.OutputState }

func (PublicIPAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddress)(nil)).Elem()
}

func (o PublicIPAddressOutput) ToPublicIPAddressOutput() PublicIPAddressOutput {
	return o
}

func (o PublicIPAddressOutput) ToPublicIPAddressOutputWithContext(ctx context.Context) PublicIPAddressOutput {
	return o
}

// The DDoS protection custom policy associated with the public IP address.
func (o PublicIPAddressOutput) DdosSettings() DdosSettingsResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) DdosSettingsResponsePtrOutput { return v.DdosSettings }).(DdosSettingsResponsePtrOutput)
}

// Specify what happens to the public IP address when the VM using it is deleted
func (o PublicIPAddressOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringPtrOutput { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

// The FQDN of the DNS record associated with the public IP address.
func (o PublicIPAddressOutput) DnsSettings() PublicIPAddressDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) PublicIPAddressDnsSettingsResponsePtrOutput { return v.DnsSettings }).(PublicIPAddressDnsSettingsResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o PublicIPAddressOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the public ip address.
func (o PublicIPAddressOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The idle timeout of the public IP address.
func (o PublicIPAddressOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.IntPtrOutput { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

// The IP address associated with the public IP address resource.
func (o PublicIPAddressOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringPtrOutput { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// The IP configuration associated with the public IP address.
func (o PublicIPAddressOutput) IpConfiguration() IPConfigurationResponseOutput {
	return o.ApplyT(func(v *PublicIPAddress) IPConfigurationResponseOutput { return v.IpConfiguration }).(IPConfigurationResponseOutput)
}

// The list of tags associated with the public IP address.
func (o PublicIPAddressOutput) IpTags() IpTagResponseArrayOutput {
	return o.ApplyT(func(v *PublicIPAddress) IpTagResponseArrayOutput { return v.IpTags }).(IpTagResponseArrayOutput)
}

// The linked public IP address of the public IP address resource.
func (o PublicIPAddressOutput) LinkedPublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) PublicIPAddressResponsePtrOutput { return v.LinkedPublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

// Resource location.
func (o PublicIPAddressOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Migration phase of Public IP Address.
func (o PublicIPAddressOutput) MigrationPhase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringPtrOutput { return v.MigrationPhase }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o PublicIPAddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The NatGateway for the Public IP address.
func (o PublicIPAddressOutput) NatGateway() NatGatewayResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) NatGatewayResponsePtrOutput { return v.NatGateway }).(NatGatewayResponsePtrOutput)
}

// The provisioning state of the public IP address resource.
func (o PublicIPAddressOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The public IP address version.
func (o PublicIPAddressOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringPtrOutput { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

// The public IP address allocation method.
func (o PublicIPAddressOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringPtrOutput { return v.PublicIPAllocationMethod }).(pulumi.StringPtrOutput)
}

// The Public IP Prefix this Public IP Address should be allocated from.
func (o PublicIPAddressOutput) PublicIPPrefix() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) SubResourceResponsePtrOutput { return v.PublicIPPrefix }).(SubResourceResponsePtrOutput)
}

// The resource GUID property of the public IP address resource.
func (o PublicIPAddressOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The service public IP address of the public IP address resource.
func (o PublicIPAddressOutput) ServicePublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) PublicIPAddressResponsePtrOutput { return v.ServicePublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

// The public IP address SKU.
func (o PublicIPAddressOutput) Sku() PublicIPAddressSkuResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) PublicIPAddressSkuResponsePtrOutput { return v.Sku }).(PublicIPAddressSkuResponsePtrOutput)
}

// Resource tags.
func (o PublicIPAddressOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o PublicIPAddressOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A list of availability zones denoting the IP allocated for the resource needs to come from.
func (o PublicIPAddressOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(PublicIPAddressOutput{})
}
