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

// List configurations for the given product family, product line and product for the given subscription.
func ListProductsAndConfigurations(ctx *pulumi.Context, args *ListProductsAndConfigurationsArgs, opts ...pulumi.InvokeOption) (*ListProductsAndConfigurationsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListProductsAndConfigurationsResult
	err := ctx.Invoke("azure-native:edgeorder/v20220501preview:listProductsAndConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProductsAndConfigurationsArgs struct {
	// Holds details about product hierarchy information and filterable property.
	ConfigurationFilter *ConfigurationFilter `pulumi:"configurationFilter"`
	// Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details.
	CustomerSubscriptionDetails *CustomerSubscriptionDetails `pulumi:"customerSubscriptionDetails"`
	// $skipToken is supported on list of configurations, which provides the next page in the list of configurations.
	SkipToken *string `pulumi:"skipToken"`
}

// The list of configurations.
type ListProductsAndConfigurationsResult struct {
	// Link for the next set of configurations.
	NextLink *string `pulumi:"nextLink"`
	// List of configurations.
	Value []ConfigurationResponse `pulumi:"value"`
}

func ListProductsAndConfigurationsOutput(ctx *pulumi.Context, args ListProductsAndConfigurationsOutputArgs, opts ...pulumi.InvokeOption) ListProductsAndConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListProductsAndConfigurationsResult, error) {
			args := v.(ListProductsAndConfigurationsArgs)
			r, err := ListProductsAndConfigurations(ctx, &args, opts...)
			var s ListProductsAndConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListProductsAndConfigurationsResultOutput)
}

type ListProductsAndConfigurationsOutputArgs struct {
	// Holds details about product hierarchy information and filterable property.
	ConfigurationFilter ConfigurationFilterPtrInput `pulumi:"configurationFilter"`
	// Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details.
	CustomerSubscriptionDetails CustomerSubscriptionDetailsPtrInput `pulumi:"customerSubscriptionDetails"`
	// $skipToken is supported on list of configurations, which provides the next page in the list of configurations.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListProductsAndConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductsAndConfigurationsArgs)(nil)).Elem()
}

// The list of configurations.
type ListProductsAndConfigurationsResultOutput struct{ *pulumi.OutputState }

func (ListProductsAndConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductsAndConfigurationsResult)(nil)).Elem()
}

func (o ListProductsAndConfigurationsResultOutput) ToListProductsAndConfigurationsResultOutput() ListProductsAndConfigurationsResultOutput {
	return o
}

func (o ListProductsAndConfigurationsResultOutput) ToListProductsAndConfigurationsResultOutputWithContext(ctx context.Context) ListProductsAndConfigurationsResultOutput {
	return o
}

func (o ListProductsAndConfigurationsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListProductsAndConfigurationsResult] {
	return pulumix.Output[ListProductsAndConfigurationsResult]{
		OutputState: o.OutputState,
	}
}

// Link for the next set of configurations.
func (o ListProductsAndConfigurationsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListProductsAndConfigurationsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of configurations.
func (o ListProductsAndConfigurationsResultOutput) Value() ConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ListProductsAndConfigurationsResult) []ConfigurationResponse { return v.Value }).(ConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListProductsAndConfigurationsResultOutput{})
}
