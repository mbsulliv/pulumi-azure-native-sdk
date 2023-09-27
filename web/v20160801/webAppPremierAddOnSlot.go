// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Premier add-on.
type WebAppPremierAddOnSlot struct {
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
	// Premier add on Name.
	PremierAddOnName pulumi.StringPtrOutput `pulumi:"premierAddOnName"`
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

// NewWebAppPremierAddOnSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppPremierAddOnSlot(ctx *pulumi.Context,
	name string, args *WebAppPremierAddOnSlotArgs, opts ...pulumi.ResourceOption) (*WebAppPremierAddOnSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppPremierAddOnSlot"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppPremierAddOnSlot
	err := ctx.RegisterResource("azure-native:web/v20160801:WebAppPremierAddOnSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppPremierAddOnSlot gets an existing WebAppPremierAddOnSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppPremierAddOnSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppPremierAddOnSlotState, opts ...pulumi.ResourceOption) (*WebAppPremierAddOnSlot, error) {
	var resource WebAppPremierAddOnSlot
	err := ctx.ReadResource("azure-native:web/v20160801:WebAppPremierAddOnSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppPremierAddOnSlot resources.
type webAppPremierAddOnSlotState struct {
}

type WebAppPremierAddOnSlotState struct {
}

func (WebAppPremierAddOnSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPremierAddOnSlotState)(nil)).Elem()
}

type webAppPremierAddOnSlotArgs struct {
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
	// Premier add on Name.
	PremierAddOnName *string `pulumi:"premierAddOnName"`
	// Premier add on Product.
	Product *string `pulumi:"product"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Premier add on SKU.
	Sku *string `pulumi:"sku"`
	// Name of the deployment slot. If a slot is not specified, the API will update the named add-on for the production slot.
	Slot string `pulumi:"slot"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Premier add on Vendor.
	Vendor *string `pulumi:"vendor"`
}

// The set of arguments for constructing a WebAppPremierAddOnSlot resource.
type WebAppPremierAddOnSlotArgs struct {
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
	// Premier add on Name.
	PremierAddOnName pulumi.StringPtrInput
	// Premier add on Product.
	Product pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Premier add on SKU.
	Sku pulumi.StringPtrInput
	// Name of the deployment slot. If a slot is not specified, the API will update the named add-on for the production slot.
	Slot pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Premier add on Vendor.
	Vendor pulumi.StringPtrInput
}

func (WebAppPremierAddOnSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPremierAddOnSlotArgs)(nil)).Elem()
}

type WebAppPremierAddOnSlotInput interface {
	pulumi.Input

	ToWebAppPremierAddOnSlotOutput() WebAppPremierAddOnSlotOutput
	ToWebAppPremierAddOnSlotOutputWithContext(ctx context.Context) WebAppPremierAddOnSlotOutput
}

func (*WebAppPremierAddOnSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppPremierAddOnSlot)(nil)).Elem()
}

func (i *WebAppPremierAddOnSlot) ToWebAppPremierAddOnSlotOutput() WebAppPremierAddOnSlotOutput {
	return i.ToWebAppPremierAddOnSlotOutputWithContext(context.Background())
}

func (i *WebAppPremierAddOnSlot) ToWebAppPremierAddOnSlotOutputWithContext(ctx context.Context) WebAppPremierAddOnSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPremierAddOnSlotOutput)
}

func (i *WebAppPremierAddOnSlot) ToOutput(ctx context.Context) pulumix.Output[*WebAppPremierAddOnSlot] {
	return pulumix.Output[*WebAppPremierAddOnSlot]{
		OutputState: i.ToWebAppPremierAddOnSlotOutputWithContext(ctx).OutputState,
	}
}

type WebAppPremierAddOnSlotOutput struct{ *pulumi.OutputState }

func (WebAppPremierAddOnSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppPremierAddOnSlot)(nil)).Elem()
}

func (o WebAppPremierAddOnSlotOutput) ToWebAppPremierAddOnSlotOutput() WebAppPremierAddOnSlotOutput {
	return o
}

func (o WebAppPremierAddOnSlotOutput) ToWebAppPremierAddOnSlotOutputWithContext(ctx context.Context) WebAppPremierAddOnSlotOutput {
	return o
}

func (o WebAppPremierAddOnSlotOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppPremierAddOnSlot] {
	return pulumix.Output[*WebAppPremierAddOnSlot]{
		OutputState: o.OutputState,
	}
}

// Kind of resource.
func (o WebAppPremierAddOnSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPremierAddOnSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Location.
func (o WebAppPremierAddOnSlotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPremierAddOnSlot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Premier add on Marketplace offer.
func (o WebAppPremierAddOnSlotOutput) MarketplaceOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPremierAddOnSlot) pulumi.StringPtrOutput { return v.MarketplaceOffer }).(pulumi.StringPtrOutput)
}

// Premier add on Marketplace publisher.
func (o WebAppPremierAddOnSlotOutput) MarketplacePublisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPremierAddOnSlot) pulumi.StringPtrOutput { return v.MarketplacePublisher }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppPremierAddOnSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPremierAddOnSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Premier add on Name.
func (o WebAppPremierAddOnSlotOutput) PremierAddOnName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPremierAddOnSlot) pulumi.StringPtrOutput { return v.PremierAddOnName }).(pulumi.StringPtrOutput)
}

// Premier add on Product.
func (o WebAppPremierAddOnSlotOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPremierAddOnSlot) pulumi.StringPtrOutput { return v.Product }).(pulumi.StringPtrOutput)
}

// Premier add on SKU.
func (o WebAppPremierAddOnSlotOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPremierAddOnSlot) pulumi.StringPtrOutput { return v.Sku }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o WebAppPremierAddOnSlotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebAppPremierAddOnSlot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o WebAppPremierAddOnSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPremierAddOnSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Premier add on Vendor.
func (o WebAppPremierAddOnSlotOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPremierAddOnSlot) pulumi.StringPtrOutput { return v.Vendor }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppPremierAddOnSlotOutput{})
}
