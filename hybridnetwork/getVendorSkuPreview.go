// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridnetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the preview information of a vendor sku.
// Azure REST API version: 2022-01-01-preview.
func LookupVendorSkuPreview(ctx *pulumi.Context, args *LookupVendorSkuPreviewArgs, opts ...pulumi.InvokeOption) (*LookupVendorSkuPreviewResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVendorSkuPreviewResult
	err := ctx.Invoke("azure-native:hybridnetwork:getVendorSkuPreview", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVendorSkuPreviewArgs struct {
	// Preview subscription ID.
	PreviewSubscription string `pulumi:"previewSubscription"`
	// The name of the vendor sku.
	SkuName string `pulumi:"skuName"`
	// The name of the vendor.
	VendorName string `pulumi:"vendorName"`
}

// Customer subscription which can use a sku.
type LookupVendorSkuPreviewResult struct {
	// The ARM ID of the resource.
	Id string `pulumi:"id"`
	// The preview subscription ID.
	Name string `pulumi:"name"`
	// The provisioning state of the PreviewSubscription resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupVendorSkuPreviewOutput(ctx *pulumi.Context, args LookupVendorSkuPreviewOutputArgs, opts ...pulumi.InvokeOption) LookupVendorSkuPreviewResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVendorSkuPreviewResult, error) {
			args := v.(LookupVendorSkuPreviewArgs)
			r, err := LookupVendorSkuPreview(ctx, &args, opts...)
			var s LookupVendorSkuPreviewResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVendorSkuPreviewResultOutput)
}

type LookupVendorSkuPreviewOutputArgs struct {
	// Preview subscription ID.
	PreviewSubscription pulumi.StringInput `pulumi:"previewSubscription"`
	// The name of the vendor sku.
	SkuName pulumi.StringInput `pulumi:"skuName"`
	// The name of the vendor.
	VendorName pulumi.StringInput `pulumi:"vendorName"`
}

func (LookupVendorSkuPreviewOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVendorSkuPreviewArgs)(nil)).Elem()
}

// Customer subscription which can use a sku.
type LookupVendorSkuPreviewResultOutput struct{ *pulumi.OutputState }

func (LookupVendorSkuPreviewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVendorSkuPreviewResult)(nil)).Elem()
}

func (o LookupVendorSkuPreviewResultOutput) ToLookupVendorSkuPreviewResultOutput() LookupVendorSkuPreviewResultOutput {
	return o
}

func (o LookupVendorSkuPreviewResultOutput) ToLookupVendorSkuPreviewResultOutputWithContext(ctx context.Context) LookupVendorSkuPreviewResultOutput {
	return o
}

// The ARM ID of the resource.
func (o LookupVendorSkuPreviewResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkuPreviewResult) string { return v.Id }).(pulumi.StringOutput)
}

// The preview subscription ID.
func (o LookupVendorSkuPreviewResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkuPreviewResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the PreviewSubscription resource.
func (o LookupVendorSkuPreviewResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkuPreviewResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system meta data relating to this resource.
func (o LookupVendorSkuPreviewResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVendorSkuPreviewResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupVendorSkuPreviewResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkuPreviewResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVendorSkuPreviewResultOutput{})
}
