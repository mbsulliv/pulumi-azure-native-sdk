// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Method to get site health summary.
func ListServerSitesControllerHealthSummary(ctx *pulumi.Context, args *ListServerSitesControllerHealthSummaryArgs, opts ...pulumi.InvokeOption) (*ListServerSitesControllerHealthSummaryResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListServerSitesControllerHealthSummaryResult
	err := ctx.Invoke("azure-native:offazure/v20231001preview:listServerSitesControllerHealthSummary", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListServerSitesControllerHealthSummaryArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name
	SiteName string `pulumi:"siteName"`
}

// Collection of SiteHealthSummary.
type ListServerSitesControllerHealthSummaryResult struct {
	// Gets the value of next link.
	NextLink string `pulumi:"nextLink"`
	// Gets the list of SiteHealthSummary.
	Value []SiteHealthSummaryResponse `pulumi:"value"`
}

func ListServerSitesControllerHealthSummaryOutput(ctx *pulumi.Context, args ListServerSitesControllerHealthSummaryOutputArgs, opts ...pulumi.InvokeOption) ListServerSitesControllerHealthSummaryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListServerSitesControllerHealthSummaryResult, error) {
			args := v.(ListServerSitesControllerHealthSummaryArgs)
			r, err := ListServerSitesControllerHealthSummary(ctx, &args, opts...)
			var s ListServerSitesControllerHealthSummaryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListServerSitesControllerHealthSummaryResultOutput)
}

type ListServerSitesControllerHealthSummaryOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site name
	SiteName pulumi.StringInput `pulumi:"siteName"`
}

func (ListServerSitesControllerHealthSummaryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServerSitesControllerHealthSummaryArgs)(nil)).Elem()
}

// Collection of SiteHealthSummary.
type ListServerSitesControllerHealthSummaryResultOutput struct{ *pulumi.OutputState }

func (ListServerSitesControllerHealthSummaryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServerSitesControllerHealthSummaryResult)(nil)).Elem()
}

func (o ListServerSitesControllerHealthSummaryResultOutput) ToListServerSitesControllerHealthSummaryResultOutput() ListServerSitesControllerHealthSummaryResultOutput {
	return o
}

func (o ListServerSitesControllerHealthSummaryResultOutput) ToListServerSitesControllerHealthSummaryResultOutputWithContext(ctx context.Context) ListServerSitesControllerHealthSummaryResultOutput {
	return o
}

// Gets the value of next link.
func (o ListServerSitesControllerHealthSummaryResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListServerSitesControllerHealthSummaryResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Gets the list of SiteHealthSummary.
func (o ListServerSitesControllerHealthSummaryResultOutput) Value() SiteHealthSummaryResponseArrayOutput {
	return o.ApplyT(func(v ListServerSitesControllerHealthSummaryResult) []SiteHealthSummaryResponse { return v.Value }).(SiteHealthSummaryResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListServerSitesControllerHealthSummaryResultOutput{})
}
