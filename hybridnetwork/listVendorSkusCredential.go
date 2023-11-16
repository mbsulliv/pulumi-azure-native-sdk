// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridnetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Generate credentials for publishing SKU images.
// Azure REST API version: 2022-01-01-preview.
func ListVendorSkusCredential(ctx *pulumi.Context, args *ListVendorSkusCredentialArgs, opts ...pulumi.InvokeOption) (*ListVendorSkusCredentialResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListVendorSkusCredentialResult
	err := ctx.Invoke("azure-native:hybridnetwork:listVendorSkusCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVendorSkusCredentialArgs struct {
	// The name of the sku.
	SkuName string `pulumi:"skuName"`
	// The name of the vendor.
	VendorName string `pulumi:"vendorName"`
}

// The Sku credential definition.
type ListVendorSkusCredentialResult struct {
	// The Acr server url
	AcrServerUrl *string `pulumi:"acrServerUrl"`
	// The credential value.
	AcrToken *string `pulumi:"acrToken"`
	// The UTC time when credential will expire.
	Expiry *string `pulumi:"expiry"`
	// The repositories that could be accessed using the current credential.
	Repositories []string `pulumi:"repositories"`
	// The username of the sku credential.
	Username *string `pulumi:"username"`
}

func ListVendorSkusCredentialOutput(ctx *pulumi.Context, args ListVendorSkusCredentialOutputArgs, opts ...pulumi.InvokeOption) ListVendorSkusCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListVendorSkusCredentialResult, error) {
			args := v.(ListVendorSkusCredentialArgs)
			r, err := ListVendorSkusCredential(ctx, &args, opts...)
			var s ListVendorSkusCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListVendorSkusCredentialResultOutput)
}

type ListVendorSkusCredentialOutputArgs struct {
	// The name of the sku.
	SkuName pulumi.StringInput `pulumi:"skuName"`
	// The name of the vendor.
	VendorName pulumi.StringInput `pulumi:"vendorName"`
}

func (ListVendorSkusCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVendorSkusCredentialArgs)(nil)).Elem()
}

// The Sku credential definition.
type ListVendorSkusCredentialResultOutput struct{ *pulumi.OutputState }

func (ListVendorSkusCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVendorSkusCredentialResult)(nil)).Elem()
}

func (o ListVendorSkusCredentialResultOutput) ToListVendorSkusCredentialResultOutput() ListVendorSkusCredentialResultOutput {
	return o
}

func (o ListVendorSkusCredentialResultOutput) ToListVendorSkusCredentialResultOutputWithContext(ctx context.Context) ListVendorSkusCredentialResultOutput {
	return o
}

// The Acr server url
func (o ListVendorSkusCredentialResultOutput) AcrServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListVendorSkusCredentialResult) *string { return v.AcrServerUrl }).(pulumi.StringPtrOutput)
}

// The credential value.
func (o ListVendorSkusCredentialResultOutput) AcrToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListVendorSkusCredentialResult) *string { return v.AcrToken }).(pulumi.StringPtrOutput)
}

// The UTC time when credential will expire.
func (o ListVendorSkusCredentialResultOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListVendorSkusCredentialResult) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

// The repositories that could be accessed using the current credential.
func (o ListVendorSkusCredentialResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListVendorSkusCredentialResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

// The username of the sku credential.
func (o ListVendorSkusCredentialResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListVendorSkusCredentialResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVendorSkusCredentialResultOutput{})
}
