// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Premier add-on.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2016-08-01, 2020-10-01.
type WebAppPremierAddOn struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Premier add on Marketplace offer.
	MarketplaceOffer pulumi.StringPtrOutput `pulumi:"marketplaceOffer"`
	// Premier add on Marketplace publisher.
	MarketplacePublisher pulumi.StringPtrOutput `pulumi:"marketplacePublisher"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Premier add on Product.
	Product pulumi.StringPtrOutput `pulumi:"product"`
	// Premier add on SKU.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Premier add on Vendor.
	Vendor pulumi.StringPtrOutput `pulumi:"vendor"`
}

// NewWebAppPremierAddOn registers a new resource with the given unique name, arguments, and options.
func NewWebAppPremierAddOn(ctx *pulumi.Context,
	name string, args *WebAppPremierAddOnArgs, opts ...pulumi.ResourceOption) (*WebAppPremierAddOn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppPremierAddOn"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppPremierAddOn
	err := ctx.RegisterResource("azure-native:web:WebAppPremierAddOn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppPremierAddOn gets an existing WebAppPremierAddOn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppPremierAddOn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppPremierAddOnState, opts ...pulumi.ResourceOption) (*WebAppPremierAddOn, error) {
	var resource WebAppPremierAddOn
	err := ctx.ReadResource("azure-native:web:WebAppPremierAddOn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppPremierAddOn resources.
type webAppPremierAddOnState struct {
}

type WebAppPremierAddOnState struct {
}

func (WebAppPremierAddOnState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPremierAddOnState)(nil)).Elem()
}

type webAppPremierAddOnArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Premier add on Marketplace offer.
	MarketplaceOffer *string `pulumi:"marketplaceOffer"`
	// Premier add on Marketplace publisher.
	MarketplacePublisher *string `pulumi:"marketplacePublisher"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Add-on name.
	PremierAddOnName *string `pulumi:"premierAddOnName"`
	// Premier add on Product.
	Product *string `pulumi:"product"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Premier add on SKU.
	Sku *string `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Premier add on Vendor.
	Vendor *string `pulumi:"vendor"`
}

// The set of arguments for constructing a WebAppPremierAddOn resource.
type WebAppPremierAddOnArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Premier add on Marketplace offer.
	MarketplaceOffer pulumi.StringPtrInput
	// Premier add on Marketplace publisher.
	MarketplacePublisher pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Add-on name.
	PremierAddOnName pulumi.StringPtrInput
	// Premier add on Product.
	Product pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Premier add on SKU.
	Sku pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Premier add on Vendor.
	Vendor pulumi.StringPtrInput
}

func (WebAppPremierAddOnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPremierAddOnArgs)(nil)).Elem()
}

type WebAppPremierAddOnInput interface {
	pulumi.Input

	ToWebAppPremierAddOnOutput() WebAppPremierAddOnOutput
	ToWebAppPremierAddOnOutputWithContext(ctx context.Context) WebAppPremierAddOnOutput
}

func (*WebAppPremierAddOn) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppPremierAddOn)(nil)).Elem()
}

func (i *WebAppPremierAddOn) ToWebAppPremierAddOnOutput() WebAppPremierAddOnOutput {
	return i.ToWebAppPremierAddOnOutputWithContext(context.Background())
}

func (i *WebAppPremierAddOn) ToWebAppPremierAddOnOutputWithContext(ctx context.Context) WebAppPremierAddOnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPremierAddOnOutput)
}

func (i *WebAppPremierAddOn) ToOutput(ctx context.Context) pulumix.Output[*WebAppPremierAddOn] {
	return pulumix.Output[*WebAppPremierAddOn]{
		OutputState: i.ToWebAppPremierAddOnOutputWithContext(ctx).OutputState,
	}
}

type WebAppPremierAddOnOutput struct{ *pulumi.OutputState }

func (WebAppPremierAddOnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppPremierAddOn)(nil)).Elem()
}

func (o WebAppPremierAddOnOutput) ToWebAppPremierAddOnOutput() WebAppPremierAddOnOutput {
	return o
}

func (o WebAppPremierAddOnOutput) ToWebAppPremierAddOnOutputWithContext(ctx context.Context) WebAppPremierAddOnOutput {
	return o
}

func (o WebAppPremierAddOnOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppPremierAddOn] {
	return pulumix.Output[*WebAppPremierAddOn]{
		OutputState: o.OutputState,
	}
}

// Kind of resource.
func (o WebAppPremierAddOnOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPremierAddOn) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Location.
func (o WebAppPremierAddOnOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPremierAddOn) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Premier add on Marketplace offer.
func (o WebAppPremierAddOnOutput) MarketplaceOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPremierAddOn) pulumi.StringPtrOutput { return v.MarketplaceOffer }).(pulumi.StringPtrOutput)
}

// Premier add on Marketplace publisher.
func (o WebAppPremierAddOnOutput) MarketplacePublisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPremierAddOn) pulumi.StringPtrOutput { return v.MarketplacePublisher }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppPremierAddOnOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPremierAddOn) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Premier add on Product.
func (o WebAppPremierAddOnOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPremierAddOn) pulumi.StringPtrOutput { return v.Product }).(pulumi.StringPtrOutput)
}

// Premier add on SKU.
func (o WebAppPremierAddOnOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPremierAddOn) pulumi.StringPtrOutput { return v.Sku }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o WebAppPremierAddOnOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebAppPremierAddOn) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o WebAppPremierAddOnOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPremierAddOn) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Premier add on Vendor.
func (o WebAppPremierAddOnOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPremierAddOn) pulumi.StringPtrOutput { return v.Vendor }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppPremierAddOnOutput{})
}
