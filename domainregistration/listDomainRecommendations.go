// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domainregistration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Get domain name recommendations based on keywords.
// Azure REST API version: 2022-09-01.
//
// Other available API versions: 2015-04-01, 2018-02-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2023-01-01.
func ListDomainRecommendations(ctx *pulumi.Context, args *ListDomainRecommendationsArgs, opts ...pulumi.InvokeOption) (*ListDomainRecommendationsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDomainRecommendationsResult
	err := ctx.Invoke("azure-native:domainregistration:listDomainRecommendations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDomainRecommendationsArgs struct {
	// Keywords to be used for generating domain recommendations.
	Keywords *string `pulumi:"keywords"`
	// Maximum number of recommendations.
	MaxDomainRecommendations *int `pulumi:"maxDomainRecommendations"`
}

// Collection of domain name identifiers.
type ListDomainRecommendationsResult struct {
	// Link to next page of resources.
	NextLink string `pulumi:"nextLink"`
	// Collection of resources.
	Value []NameIdentifierResponse `pulumi:"value"`
}

func ListDomainRecommendationsOutput(ctx *pulumi.Context, args ListDomainRecommendationsOutputArgs, opts ...pulumi.InvokeOption) ListDomainRecommendationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDomainRecommendationsResult, error) {
			args := v.(ListDomainRecommendationsArgs)
			r, err := ListDomainRecommendations(ctx, &args, opts...)
			var s ListDomainRecommendationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDomainRecommendationsResultOutput)
}

type ListDomainRecommendationsOutputArgs struct {
	// Keywords to be used for generating domain recommendations.
	Keywords pulumi.StringPtrInput `pulumi:"keywords"`
	// Maximum number of recommendations.
	MaxDomainRecommendations pulumi.IntPtrInput `pulumi:"maxDomainRecommendations"`
}

func (ListDomainRecommendationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomainRecommendationsArgs)(nil)).Elem()
}

// Collection of domain name identifiers.
type ListDomainRecommendationsResultOutput struct{ *pulumi.OutputState }

func (ListDomainRecommendationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomainRecommendationsResult)(nil)).Elem()
}

func (o ListDomainRecommendationsResultOutput) ToListDomainRecommendationsResultOutput() ListDomainRecommendationsResultOutput {
	return o
}

func (o ListDomainRecommendationsResultOutput) ToListDomainRecommendationsResultOutputWithContext(ctx context.Context) ListDomainRecommendationsResultOutput {
	return o
}

// Link to next page of resources.
func (o ListDomainRecommendationsResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListDomainRecommendationsResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Collection of resources.
func (o ListDomainRecommendationsResultOutput) Value() NameIdentifierResponseArrayOutput {
	return o.ApplyT(func(v ListDomainRecommendationsResult) []NameIdentifierResponse { return v.Value }).(NameIdentifierResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDomainRecommendationsResultOutput{})
}
