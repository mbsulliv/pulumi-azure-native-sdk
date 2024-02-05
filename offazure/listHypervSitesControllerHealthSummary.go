// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package offazure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Method to get site health summary.
// Azure REST API version: 2023-06-06.
//
// Other available API versions: 2023-10-01-preview.
func ListHypervSitesControllerHealthSummary(ctx *pulumi.Context, args *ListHypervSitesControllerHealthSummaryArgs, opts ...pulumi.InvokeOption) (*ListHypervSitesControllerHealthSummaryResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListHypervSitesControllerHealthSummaryResult
	err := ctx.Invoke("azure-native:offazure:listHypervSitesControllerHealthSummary", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListHypervSitesControllerHealthSummaryArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name
	SiteName string `pulumi:"siteName"`
}

// Collection of SiteHealthSummary.
type ListHypervSitesControllerHealthSummaryResult struct {
	// Gets the value of next link.
	NextLink string `pulumi:"nextLink"`
	// Gets the list of SiteHealthSummary.
	Value []SiteHealthSummaryResponse `pulumi:"value"`
}

func ListHypervSitesControllerHealthSummaryOutput(ctx *pulumi.Context, args ListHypervSitesControllerHealthSummaryOutputArgs, opts ...pulumi.InvokeOption) ListHypervSitesControllerHealthSummaryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListHypervSitesControllerHealthSummaryResult, error) {
			args := v.(ListHypervSitesControllerHealthSummaryArgs)
			r, err := ListHypervSitesControllerHealthSummary(ctx, &args, opts...)
			var s ListHypervSitesControllerHealthSummaryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListHypervSitesControllerHealthSummaryResultOutput)
}

type ListHypervSitesControllerHealthSummaryOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site name
	SiteName pulumi.StringInput `pulumi:"siteName"`
}

func (ListHypervSitesControllerHealthSummaryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListHypervSitesControllerHealthSummaryArgs)(nil)).Elem()
}

// Collection of SiteHealthSummary.
type ListHypervSitesControllerHealthSummaryResultOutput struct{ *pulumi.OutputState }

func (ListHypervSitesControllerHealthSummaryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListHypervSitesControllerHealthSummaryResult)(nil)).Elem()
}

func (o ListHypervSitesControllerHealthSummaryResultOutput) ToListHypervSitesControllerHealthSummaryResultOutput() ListHypervSitesControllerHealthSummaryResultOutput {
	return o
}

func (o ListHypervSitesControllerHealthSummaryResultOutput) ToListHypervSitesControllerHealthSummaryResultOutputWithContext(ctx context.Context) ListHypervSitesControllerHealthSummaryResultOutput {
	return o
}

// Gets the value of next link.
func (o ListHypervSitesControllerHealthSummaryResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListHypervSitesControllerHealthSummaryResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Gets the list of SiteHealthSummary.
func (o ListHypervSitesControllerHealthSummaryResultOutput) Value() SiteHealthSummaryResponseArrayOutput {
	return o.ApplyT(func(v ListHypervSitesControllerHealthSummaryResult) []SiteHealthSummaryResponse { return v.Value }).(SiteHealthSummaryResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListHypervSitesControllerHealthSummaryResultOutput{})
}
