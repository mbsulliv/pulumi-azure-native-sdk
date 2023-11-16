// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VpnSite Resource.
type VpnSite struct {
	pulumi.CustomResourceState

	// The AddressSpace that contains an array of IP address ranges.
	AddressSpace AddressSpaceResponsePtrOutput `pulumi:"addressSpace"`
	// The set of bgp properties.
	BgpProperties BgpSettingsResponsePtrOutput `pulumi:"bgpProperties"`
	// The device properties.
	DeviceProperties DevicePropertiesResponsePtrOutput `pulumi:"deviceProperties"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The ip-address for the vpn-site.
	IpAddress pulumi.StringPtrOutput `pulumi:"ipAddress"`
	// IsSecuritySite flag.
	IsSecuritySite pulumi.BoolPtrOutput `pulumi:"isSecuritySite"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Office365 Policy.
	O365Policy O365PolicyPropertiesResponsePtrOutput `pulumi:"o365Policy"`
	// The provisioning state of the VPN site resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The key for vpn-site that can be used for connections.
	SiteKey pulumi.StringPtrOutput `pulumi:"siteKey"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The VirtualWAN to which the vpnSite belongs.
	VirtualWan SubResourceResponsePtrOutput `pulumi:"virtualWan"`
	// List of all vpn site links.
	VpnSiteLinks VpnSiteLinkResponseArrayOutput `pulumi:"vpnSiteLinks"`
}

// NewVpnSite registers a new resource with the given unique name, arguments, and options.
func NewVpnSite(ctx *pulumi.Context,
	name string, args *VpnSiteArgs, opts ...pulumi.ResourceOption) (*VpnSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:VpnSite"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VpnSite
	err := ctx.RegisterResource("azure-native:network/v20230201:VpnSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnSite gets an existing VpnSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnSiteState, opts ...pulumi.ResourceOption) (*VpnSite, error) {
	var resource VpnSite
	err := ctx.ReadResource("azure-native:network/v20230201:VpnSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnSite resources.
type vpnSiteState struct {
}

type VpnSiteState struct {
}

func (VpnSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSiteState)(nil)).Elem()
}

type vpnSiteArgs struct {
	// The AddressSpace that contains an array of IP address ranges.
	AddressSpace *AddressSpace `pulumi:"addressSpace"`
	// The set of bgp properties.
	BgpProperties *BgpSettings `pulumi:"bgpProperties"`
	// The device properties.
	DeviceProperties *DeviceProperties `pulumi:"deviceProperties"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The ip-address for the vpn-site.
	IpAddress *string `pulumi:"ipAddress"`
	// IsSecuritySite flag.
	IsSecuritySite *bool `pulumi:"isSecuritySite"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Office365 Policy.
	O365Policy *O365PolicyProperties `pulumi:"o365Policy"`
	// The resource group name of the VpnSite.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The key for vpn-site that can be used for connections.
	SiteKey *string `pulumi:"siteKey"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The VirtualWAN to which the vpnSite belongs.
	VirtualWan *SubResource `pulumi:"virtualWan"`
	// List of all vpn site links.
	VpnSiteLinks []VpnSiteLink `pulumi:"vpnSiteLinks"`
	// The name of the VpnSite being created or updated.
	VpnSiteName *string `pulumi:"vpnSiteName"`
}

// The set of arguments for constructing a VpnSite resource.
type VpnSiteArgs struct {
	// The AddressSpace that contains an array of IP address ranges.
	AddressSpace AddressSpacePtrInput
	// The set of bgp properties.
	BgpProperties BgpSettingsPtrInput
	// The device properties.
	DeviceProperties DevicePropertiesPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The ip-address for the vpn-site.
	IpAddress pulumi.StringPtrInput
	// IsSecuritySite flag.
	IsSecuritySite pulumi.BoolPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Office365 Policy.
	O365Policy O365PolicyPropertiesPtrInput
	// The resource group name of the VpnSite.
	ResourceGroupName pulumi.StringInput
	// The key for vpn-site that can be used for connections.
	SiteKey pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The VirtualWAN to which the vpnSite belongs.
	VirtualWan SubResourcePtrInput
	// List of all vpn site links.
	VpnSiteLinks VpnSiteLinkArrayInput
	// The name of the VpnSite being created or updated.
	VpnSiteName pulumi.StringPtrInput
}

func (VpnSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSiteArgs)(nil)).Elem()
}

type VpnSiteInput interface {
	pulumi.Input

	ToVpnSiteOutput() VpnSiteOutput
	ToVpnSiteOutputWithContext(ctx context.Context) VpnSiteOutput
}

func (*VpnSite) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnSite)(nil)).Elem()
}

func (i *VpnSite) ToVpnSiteOutput() VpnSiteOutput {
	return i.ToVpnSiteOutputWithContext(context.Background())
}

func (i *VpnSite) ToVpnSiteOutputWithContext(ctx context.Context) VpnSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteOutput)
}

type VpnSiteOutput struct{ *pulumi.OutputState }

func (VpnSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnSite)(nil)).Elem()
}

func (o VpnSiteOutput) ToVpnSiteOutput() VpnSiteOutput {
	return o
}

func (o VpnSiteOutput) ToVpnSiteOutputWithContext(ctx context.Context) VpnSiteOutput {
	return o
}

// The AddressSpace that contains an array of IP address ranges.
func (o VpnSiteOutput) AddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v *VpnSite) AddressSpaceResponsePtrOutput { return v.AddressSpace }).(AddressSpaceResponsePtrOutput)
}

// The set of bgp properties.
func (o VpnSiteOutput) BgpProperties() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *VpnSite) BgpSettingsResponsePtrOutput { return v.BgpProperties }).(BgpSettingsResponsePtrOutput)
}

// The device properties.
func (o VpnSiteOutput) DeviceProperties() DevicePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *VpnSite) DevicePropertiesResponsePtrOutput { return v.DeviceProperties }).(DevicePropertiesResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o VpnSiteOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnSite) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The ip-address for the vpn-site.
func (o VpnSiteOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnSite) pulumi.StringPtrOutput { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// IsSecuritySite flag.
func (o VpnSiteOutput) IsSecuritySite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpnSite) pulumi.BoolPtrOutput { return v.IsSecuritySite }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o VpnSiteOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnSite) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o VpnSiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnSite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Office365 Policy.
func (o VpnSiteOutput) O365Policy() O365PolicyPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *VpnSite) O365PolicyPropertiesResponsePtrOutput { return v.O365Policy }).(O365PolicyPropertiesResponsePtrOutput)
}

// The provisioning state of the VPN site resource.
func (o VpnSiteOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnSite) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The key for vpn-site that can be used for connections.
func (o VpnSiteOutput) SiteKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnSite) pulumi.StringPtrOutput { return v.SiteKey }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o VpnSiteOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpnSite) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o VpnSiteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnSite) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The VirtualWAN to which the vpnSite belongs.
func (o VpnSiteOutput) VirtualWan() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VpnSite) SubResourceResponsePtrOutput { return v.VirtualWan }).(SubResourceResponsePtrOutput)
}

// List of all vpn site links.
func (o VpnSiteOutput) VpnSiteLinks() VpnSiteLinkResponseArrayOutput {
	return o.ApplyT(func(v *VpnSite) VpnSiteLinkResponseArrayOutput { return v.VpnSiteLinks }).(VpnSiteLinkResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnSiteOutput{})
}
