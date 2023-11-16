// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package edgeorder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This method provides the list of product families for the given subscription.
// Azure REST API version: 2021-12-01.
//
// Other available API versions: 2020-12-01-preview.
func ListProductFamilies(ctx *pulumi.Context, args *ListProductFamiliesArgs, opts ...pulumi.InvokeOption) (*ListProductFamiliesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListProductFamiliesResult
	err := ctx.Invoke("azure-native:edgeorder:listProductFamilies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProductFamiliesArgs struct {
	// Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details
	CustomerSubscriptionDetails *CustomerSubscriptionDetails `pulumi:"customerSubscriptionDetails"`
	// $expand is supported on configurations parameter for product, which provides details on the configurations for the product.
	Expand *string `pulumi:"expand"`
	// Dictionary of filterable properties on product family.
	FilterableProperties map[string][]FilterableProperty `pulumi:"filterableProperties"`
	// $skipToken is supported on list of product families, which provides the next page in the list of product families.
	SkipToken *string `pulumi:"skipToken"`
}

// The list of product families.
type ListProductFamiliesResult struct {
	// Link for the next set of product families.
	NextLink *string `pulumi:"nextLink"`
	// List of product families.
	Value []ProductFamilyResponse `pulumi:"value"`
}

func ListProductFamiliesOutput(ctx *pulumi.Context, args ListProductFamiliesOutputArgs, opts ...pulumi.InvokeOption) ListProductFamiliesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListProductFamiliesResult, error) {
			args := v.(ListProductFamiliesArgs)
			r, err := ListProductFamilies(ctx, &args, opts...)
			var s ListProductFamiliesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListProductFamiliesResultOutput)
}

type ListProductFamiliesOutputArgs struct {
	// Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details
	CustomerSubscriptionDetails CustomerSubscriptionDetailsPtrInput `pulumi:"customerSubscriptionDetails"`
	// $expand is supported on configurations parameter for product, which provides details on the configurations for the product.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// Dictionary of filterable properties on product family.
	FilterableProperties FilterablePropertyArrayMapInput `pulumi:"filterableProperties"`
	// $skipToken is supported on list of product families, which provides the next page in the list of product families.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListProductFamiliesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductFamiliesArgs)(nil)).Elem()
}

// The list of product families.
type ListProductFamiliesResultOutput struct{ *pulumi.OutputState }

func (ListProductFamiliesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductFamiliesResult)(nil)).Elem()
}

func (o ListProductFamiliesResultOutput) ToListProductFamiliesResultOutput() ListProductFamiliesResultOutput {
	return o
}

func (o ListProductFamiliesResultOutput) ToListProductFamiliesResultOutputWithContext(ctx context.Context) ListProductFamiliesResultOutput {
	return o
}

// Link for the next set of product families.
func (o ListProductFamiliesResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListProductFamiliesResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of product families.
func (o ListProductFamiliesResultOutput) Value() ProductFamilyResponseArrayOutput {
	return o.ApplyT(func(v ListProductFamiliesResult) []ProductFamilyResponse { return v.Value }).(ProductFamilyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListProductFamiliesResultOutput{})
}
