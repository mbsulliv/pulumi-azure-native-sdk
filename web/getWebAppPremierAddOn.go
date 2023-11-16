// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Gets a named add-on of an app.
// Azure REST API version: 2022-09-01.
//
// Other available API versions: 2016-08-01, 2020-10-01, 2023-01-01.
func LookupWebAppPremierAddOn(ctx *pulumi.Context, args *LookupWebAppPremierAddOnArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPremierAddOnResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppPremierAddOnResult
	err := ctx.Invoke("azure-native:web:getWebAppPremierAddOn", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPremierAddOnArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Add-on name.
	PremierAddOnName string `pulumi:"premierAddOnName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Premier add-on.
type LookupWebAppPremierAddOnResult struct {
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

func LookupWebAppPremierAddOnOutput(ctx *pulumi.Context, args LookupWebAppPremierAddOnOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppPremierAddOnResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppPremierAddOnResult, error) {
			args := v.(LookupWebAppPremierAddOnArgs)
			r, err := LookupWebAppPremierAddOn(ctx, &args, opts...)
			var s LookupWebAppPremierAddOnResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppPremierAddOnResultOutput)
}

type LookupWebAppPremierAddOnOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Add-on name.
	PremierAddOnName pulumi.StringInput `pulumi:"premierAddOnName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppPremierAddOnOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPremierAddOnArgs)(nil)).Elem()
}

// Premier add-on.
type LookupWebAppPremierAddOnResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppPremierAddOnResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPremierAddOnResult)(nil)).Elem()
}

func (o LookupWebAppPremierAddOnResultOutput) ToLookupWebAppPremierAddOnResultOutput() LookupWebAppPremierAddOnResultOutput {
	return o
}

func (o LookupWebAppPremierAddOnResultOutput) ToLookupWebAppPremierAddOnResultOutputWithContext(ctx context.Context) LookupWebAppPremierAddOnResultOutput {
	return o
}

// Resource Id.
func (o LookupWebAppPremierAddOnResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupWebAppPremierAddOnResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Location.
func (o LookupWebAppPremierAddOnResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) string { return v.Location }).(pulumi.StringOutput)
}

// Premier add on Marketplace offer.
func (o LookupWebAppPremierAddOnResultOutput) MarketplaceOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) *string { return v.MarketplaceOffer }).(pulumi.StringPtrOutput)
}

// Premier add on Marketplace publisher.
func (o LookupWebAppPremierAddOnResultOutput) MarketplacePublisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) *string { return v.MarketplacePublisher }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppPremierAddOnResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) string { return v.Name }).(pulumi.StringOutput)
}

// Premier add on Product.
func (o LookupWebAppPremierAddOnResultOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) *string { return v.Product }).(pulumi.StringPtrOutput)
}

// Premier add on SKU.
func (o LookupWebAppPremierAddOnResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupWebAppPremierAddOnResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupWebAppPremierAddOnResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) string { return v.Type }).(pulumi.StringOutput)
}

// Premier add on Vendor.
func (o LookupWebAppPremierAddOnResultOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) *string { return v.Vendor }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppPremierAddOnResultOutput{})
}
