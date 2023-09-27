// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a named add-on of an app.
func LookupWebAppPremierAddOnSlot(ctx *pulumi.Context, args *LookupWebAppPremierAddOnSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPremierAddOnSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppPremierAddOnSlotResult
	err := ctx.Invoke("azure-native:web/v20160801:getWebAppPremierAddOnSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPremierAddOnSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Add-on name.
	PremierAddOnName string `pulumi:"premierAddOnName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the named add-on for the production slot.
	Slot string `pulumi:"slot"`
}

// Premier add-on.
type LookupWebAppPremierAddOnSlotResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Premier add on Marketplace offer.
	MarketplaceOffer *string `pulumi:"marketplaceOffer"`
	// Premier add on Marketplace publisher.
	MarketplacePublisher *string `pulumi:"marketplacePublisher"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Premier add on Name.
	PremierAddOnName *string `pulumi:"premierAddOnName"`
	// Premier add on Product.
	Product *string `pulumi:"product"`
	// Premier add on SKU.
	Sku *string `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// Premier add on Vendor.
	Vendor *string `pulumi:"vendor"`
}

func LookupWebAppPremierAddOnSlotOutput(ctx *pulumi.Context, args LookupWebAppPremierAddOnSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppPremierAddOnSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppPremierAddOnSlotResult, error) {
			args := v.(LookupWebAppPremierAddOnSlotArgs)
			r, err := LookupWebAppPremierAddOnSlot(ctx, &args, opts...)
			var s LookupWebAppPremierAddOnSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppPremierAddOnSlotResultOutput)
}

type LookupWebAppPremierAddOnSlotOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Add-on name.
	PremierAddOnName pulumi.StringInput `pulumi:"premierAddOnName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the named add-on for the production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppPremierAddOnSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPremierAddOnSlotArgs)(nil)).Elem()
}

// Premier add-on.
type LookupWebAppPremierAddOnSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppPremierAddOnSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPremierAddOnSlotResult)(nil)).Elem()
}

func (o LookupWebAppPremierAddOnSlotResultOutput) ToLookupWebAppPremierAddOnSlotResultOutput() LookupWebAppPremierAddOnSlotResultOutput {
	return o
}

func (o LookupWebAppPremierAddOnSlotResultOutput) ToLookupWebAppPremierAddOnSlotResultOutputWithContext(ctx context.Context) LookupWebAppPremierAddOnSlotResultOutput {
	return o
}

func (o LookupWebAppPremierAddOnSlotResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWebAppPremierAddOnSlotResult] {
	return pulumix.Output[LookupWebAppPremierAddOnSlotResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id.
func (o LookupWebAppPremierAddOnSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupWebAppPremierAddOnSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Location.
func (o LookupWebAppPremierAddOnSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

// Premier add on Marketplace offer.
func (o LookupWebAppPremierAddOnSlotResultOutput) MarketplaceOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) *string { return v.MarketplaceOffer }).(pulumi.StringPtrOutput)
}

// Premier add on Marketplace publisher.
func (o LookupWebAppPremierAddOnSlotResultOutput) MarketplacePublisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) *string { return v.MarketplacePublisher }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppPremierAddOnSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Premier add on Name.
func (o LookupWebAppPremierAddOnSlotResultOutput) PremierAddOnName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) *string { return v.PremierAddOnName }).(pulumi.StringPtrOutput)
}

// Premier add on Product.
func (o LookupWebAppPremierAddOnSlotResultOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) *string { return v.Product }).(pulumi.StringPtrOutput)
}

// Premier add on SKU.
func (o LookupWebAppPremierAddOnSlotResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupWebAppPremierAddOnSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupWebAppPremierAddOnSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

// Premier add on Vendor.
func (o LookupWebAppPremierAddOnSlotResultOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) *string { return v.Vendor }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppPremierAddOnSlotResultOutput{})
}
