// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// List product families for the given subscription.
func ListProductsAndConfigurationProductFamilies(ctx *pulumi.Context, args *ListProductsAndConfigurationProductFamiliesArgs, opts ...pulumi.InvokeOption) (*ListProductsAndConfigurationProductFamiliesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListProductsAndConfigurationProductFamiliesResult
	err := ctx.Invoke("azure-native:edgeorder/v20220501preview:listProductsAndConfigurationProductFamilies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProductsAndConfigurationProductFamiliesArgs struct {
	// Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details.
	CustomerSubscriptionDetails *CustomerSubscriptionDetails `pulumi:"customerSubscriptionDetails"`
	// $expand is supported on configurations parameter for product, which provides details on the configurations for the product.
	Expand *string `pulumi:"expand"`
	// Dictionary of filterable properties on product family.
	FilterableProperties map[string][]FilterableProperty `pulumi:"filterableProperties"`
	// $skipToken is supported on list of product families, which provides the next page in the list of product families.
	SkipToken *string `pulumi:"skipToken"`
}

// The list of product families.
type ListProductsAndConfigurationProductFamiliesResult struct {
	// Link for the next set of product families.
	NextLink *string `pulumi:"nextLink"`
	// List of product families.
	Value []ProductFamilyResponse `pulumi:"value"`
}

func ListProductsAndConfigurationProductFamiliesOutput(ctx *pulumi.Context, args ListProductsAndConfigurationProductFamiliesOutputArgs, opts ...pulumi.InvokeOption) ListProductsAndConfigurationProductFamiliesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListProductsAndConfigurationProductFamiliesResult, error) {
			args := v.(ListProductsAndConfigurationProductFamiliesArgs)
			r, err := ListProductsAndConfigurationProductFamilies(ctx, &args, opts...)
			var s ListProductsAndConfigurationProductFamiliesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListProductsAndConfigurationProductFamiliesResultOutput)
}

type ListProductsAndConfigurationProductFamiliesOutputArgs struct {
	// Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details.
	CustomerSubscriptionDetails CustomerSubscriptionDetailsPtrInput `pulumi:"customerSubscriptionDetails"`
	// $expand is supported on configurations parameter for product, which provides details on the configurations for the product.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// Dictionary of filterable properties on product family.
	FilterableProperties FilterablePropertyArrayMapInput `pulumi:"filterableProperties"`
	// $skipToken is supported on list of product families, which provides the next page in the list of product families.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListProductsAndConfigurationProductFamiliesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductsAndConfigurationProductFamiliesArgs)(nil)).Elem()
}

// The list of product families.
type ListProductsAndConfigurationProductFamiliesResultOutput struct{ *pulumi.OutputState }

func (ListProductsAndConfigurationProductFamiliesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductsAndConfigurationProductFamiliesResult)(nil)).Elem()
}

func (o ListProductsAndConfigurationProductFamiliesResultOutput) ToListProductsAndConfigurationProductFamiliesResultOutput() ListProductsAndConfigurationProductFamiliesResultOutput {
	return o
}

func (o ListProductsAndConfigurationProductFamiliesResultOutput) ToListProductsAndConfigurationProductFamiliesResultOutputWithContext(ctx context.Context) ListProductsAndConfigurationProductFamiliesResultOutput {
	return o
}

func (o ListProductsAndConfigurationProductFamiliesResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListProductsAndConfigurationProductFamiliesResult] {
	return pulumix.Output[ListProductsAndConfigurationProductFamiliesResult]{
		OutputState: o.OutputState,
	}
}

// Link for the next set of product families.
func (o ListProductsAndConfigurationProductFamiliesResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListProductsAndConfigurationProductFamiliesResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of product families.
func (o ListProductsAndConfigurationProductFamiliesResultOutput) Value() ProductFamilyResponseArrayOutput {
	return o.ApplyT(func(v ListProductsAndConfigurationProductFamiliesResult) []ProductFamilyResponse { return v.Value }).(ProductFamilyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListProductsAndConfigurationProductFamiliesResultOutput{})
}
