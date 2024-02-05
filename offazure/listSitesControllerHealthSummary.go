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
func ListSitesControllerHealthSummary(ctx *pulumi.Context, args *ListSitesControllerHealthSummaryArgs, opts ...pulumi.InvokeOption) (*ListSitesControllerHealthSummaryResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListSitesControllerHealthSummaryResult
	err := ctx.Invoke("azure-native:offazure:listSitesControllerHealthSummary", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSitesControllerHealthSummaryArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name
	SiteName string `pulumi:"siteName"`
}

// Collection of SiteHealthSummary.
type ListSitesControllerHealthSummaryResult struct {
	// Gets the value of next link.
	NextLink string `pulumi:"nextLink"`
	// Gets the list of SiteHealthSummary.
	Value []SiteHealthSummaryResponse `pulumi:"value"`
}

func ListSitesControllerHealthSummaryOutput(ctx *pulumi.Context, args ListSitesControllerHealthSummaryOutputArgs, opts ...pulumi.InvokeOption) ListSitesControllerHealthSummaryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSitesControllerHealthSummaryResult, error) {
			args := v.(ListSitesControllerHealthSummaryArgs)
			r, err := ListSitesControllerHealthSummary(ctx, &args, opts...)
			var s ListSitesControllerHealthSummaryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSitesControllerHealthSummaryResultOutput)
}

type ListSitesControllerHealthSummaryOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site name
	SiteName pulumi.StringInput `pulumi:"siteName"`
}

func (ListSitesControllerHealthSummaryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSitesControllerHealthSummaryArgs)(nil)).Elem()
}

// Collection of SiteHealthSummary.
type ListSitesControllerHealthSummaryResultOutput struct{ *pulumi.OutputState }

func (ListSitesControllerHealthSummaryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSitesControllerHealthSummaryResult)(nil)).Elem()
}

func (o ListSitesControllerHealthSummaryResultOutput) ToListSitesControllerHealthSummaryResultOutput() ListSitesControllerHealthSummaryResultOutput {
	return o
}

func (o ListSitesControllerHealthSummaryResultOutput) ToListSitesControllerHealthSummaryResultOutputWithContext(ctx context.Context) ListSitesControllerHealthSummaryResultOutput {
	return o
}

// Gets the value of next link.
func (o ListSitesControllerHealthSummaryResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListSitesControllerHealthSummaryResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Gets the list of SiteHealthSummary.
func (o ListSitesControllerHealthSummaryResultOutput) Value() SiteHealthSummaryResponseArrayOutput {
	return o.ApplyT(func(v ListSitesControllerHealthSummaryResult) []SiteHealthSummaryResponse { return v.Value }).(SiteHealthSummaryResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSitesControllerHealthSummaryResultOutput{})
}
