// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuresphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists device insights for catalog.
// API Version: 2022-09-01-preview.
func ListCatalogDeviceInsights(ctx *pulumi.Context, args *ListCatalogDeviceInsightsArgs, opts ...pulumi.InvokeOption) (*ListCatalogDeviceInsightsResult, error) {
	var rv ListCatalogDeviceInsightsResult
	err := ctx.Invoke("azure-native:azuresphere:listCatalogDeviceInsights", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListCatalogDeviceInsightsArgs struct {
	// Name of catalog
	CatalogName string `pulumi:"catalogName"`
	// Filter the result list using the given expression
	Filter *string `pulumi:"filter"`
	// The maximum number of result items per page.
	Maxpagesize *int `pulumi:"maxpagesize"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The number of result items to skip.
	Skip *int `pulumi:"skip"`
	// The number of result items to return.
	Top *int `pulumi:"top"`
}

// Paged collection of DeviceInsight items
type ListCatalogDeviceInsightsResult struct {
	// The link to the next page of items
	NextLink *string `pulumi:"nextLink"`
	// The DeviceInsight items on this page
	Value []DeviceInsightResponse `pulumi:"value"`
}

func ListCatalogDeviceInsightsOutput(ctx *pulumi.Context, args ListCatalogDeviceInsightsOutputArgs, opts ...pulumi.InvokeOption) ListCatalogDeviceInsightsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListCatalogDeviceInsightsResult, error) {
			args := v.(ListCatalogDeviceInsightsArgs)
			r, err := ListCatalogDeviceInsights(ctx, &args, opts...)
			var s ListCatalogDeviceInsightsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListCatalogDeviceInsightsResultOutput)
}

type ListCatalogDeviceInsightsOutputArgs struct {
	// Name of catalog
	CatalogName pulumi.StringInput `pulumi:"catalogName"`
	// Filter the result list using the given expression
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// The maximum number of result items per page.
	Maxpagesize pulumi.IntPtrInput `pulumi:"maxpagesize"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The number of result items to skip.
	Skip pulumi.IntPtrInput `pulumi:"skip"`
	// The number of result items to return.
	Top pulumi.IntPtrInput `pulumi:"top"`
}

func (ListCatalogDeviceInsightsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCatalogDeviceInsightsArgs)(nil)).Elem()
}

// Paged collection of DeviceInsight items
type ListCatalogDeviceInsightsResultOutput struct{ *pulumi.OutputState }

func (ListCatalogDeviceInsightsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCatalogDeviceInsightsResult)(nil)).Elem()
}

func (o ListCatalogDeviceInsightsResultOutput) ToListCatalogDeviceInsightsResultOutput() ListCatalogDeviceInsightsResultOutput {
	return o
}

func (o ListCatalogDeviceInsightsResultOutput) ToListCatalogDeviceInsightsResultOutputWithContext(ctx context.Context) ListCatalogDeviceInsightsResultOutput {
	return o
}

// The link to the next page of items
func (o ListCatalogDeviceInsightsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListCatalogDeviceInsightsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// The DeviceInsight items on this page
func (o ListCatalogDeviceInsightsResultOutput) Value() DeviceInsightResponseArrayOutput {
	return o.ApplyT(func(v ListCatalogDeviceInsightsResult) []DeviceInsightResponse { return v.Value }).(DeviceInsightResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListCatalogDeviceInsightsResultOutput{})
}