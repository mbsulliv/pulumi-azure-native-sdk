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

// Virtual Appliance Site resource.
type VirtualApplianceSite struct {
	pulumi.CustomResourceState

	// Address Prefix.
	AddressPrefix pulumi.StringPtrOutput `pulumi:"addressPrefix"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Name of the virtual appliance site.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Office 365 Policy.
	O365Policy Office365PolicyPropertiesResponsePtrOutput `pulumi:"o365Policy"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Site type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualApplianceSite registers a new resource with the given unique name, arguments, and options.
func NewVirtualApplianceSite(ctx *pulumi.Context,
	name string, args *VirtualApplianceSiteArgs, opts ...pulumi.ResourceOption) (*VirtualApplianceSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkVirtualApplianceName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkVirtualApplianceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:VirtualApplianceSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:VirtualApplianceSite"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualApplianceSite
	err := ctx.RegisterResource("azure-native:network/v20230601:VirtualApplianceSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualApplianceSite gets an existing VirtualApplianceSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualApplianceSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualApplianceSiteState, opts ...pulumi.ResourceOption) (*VirtualApplianceSite, error) {
	var resource VirtualApplianceSite
	err := ctx.ReadResource("azure-native:network/v20230601:VirtualApplianceSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualApplianceSite resources.
type virtualApplianceSiteState struct {
}

type VirtualApplianceSiteState struct {
}

func (VirtualApplianceSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualApplianceSiteState)(nil)).Elem()
}

type virtualApplianceSiteArgs struct {
	// Address Prefix.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Name of the virtual appliance site.
	Name *string `pulumi:"name"`
	// The name of the Network Virtual Appliance.
	NetworkVirtualApplianceName string `pulumi:"networkVirtualApplianceName"`
	// Office 365 Policy.
	O365Policy *Office365PolicyProperties `pulumi:"o365Policy"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the site.
	SiteName *string `pulumi:"siteName"`
}

// The set of arguments for constructing a VirtualApplianceSite resource.
type VirtualApplianceSiteArgs struct {
	// Address Prefix.
	AddressPrefix pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Name of the virtual appliance site.
	Name pulumi.StringPtrInput
	// The name of the Network Virtual Appliance.
	NetworkVirtualApplianceName pulumi.StringInput
	// Office 365 Policy.
	O365Policy Office365PolicyPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the site.
	SiteName pulumi.StringPtrInput
}

func (VirtualApplianceSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualApplianceSiteArgs)(nil)).Elem()
}

type VirtualApplianceSiteInput interface {
	pulumi.Input

	ToVirtualApplianceSiteOutput() VirtualApplianceSiteOutput
	ToVirtualApplianceSiteOutputWithContext(ctx context.Context) VirtualApplianceSiteOutput
}

func (*VirtualApplianceSite) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualApplianceSite)(nil)).Elem()
}

func (i *VirtualApplianceSite) ToVirtualApplianceSiteOutput() VirtualApplianceSiteOutput {
	return i.ToVirtualApplianceSiteOutputWithContext(context.Background())
}

func (i *VirtualApplianceSite) ToVirtualApplianceSiteOutputWithContext(ctx context.Context) VirtualApplianceSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualApplianceSiteOutput)
}

type VirtualApplianceSiteOutput struct{ *pulumi.OutputState }

func (VirtualApplianceSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualApplianceSite)(nil)).Elem()
}

func (o VirtualApplianceSiteOutput) ToVirtualApplianceSiteOutput() VirtualApplianceSiteOutput {
	return o
}

func (o VirtualApplianceSiteOutput) ToVirtualApplianceSiteOutputWithContext(ctx context.Context) VirtualApplianceSiteOutput {
	return o
}

// Address Prefix.
func (o VirtualApplianceSiteOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualApplianceSite) pulumi.StringPtrOutput { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o VirtualApplianceSiteOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualApplianceSite) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Name of the virtual appliance site.
func (o VirtualApplianceSiteOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualApplianceSite) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Office 365 Policy.
func (o VirtualApplianceSiteOutput) O365Policy() Office365PolicyPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *VirtualApplianceSite) Office365PolicyPropertiesResponsePtrOutput { return v.O365Policy }).(Office365PolicyPropertiesResponsePtrOutput)
}

// The provisioning state of the resource.
func (o VirtualApplianceSiteOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualApplianceSite) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Site type.
func (o VirtualApplianceSiteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualApplianceSite) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualApplianceSiteOutput{})
}
