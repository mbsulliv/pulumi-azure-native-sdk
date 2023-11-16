// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Execute Insights for an entity.
func GetEntityInsights(ctx *pulumi.Context, args *GetEntityInsightsArgs, opts ...pulumi.InvokeOption) (*GetEntityInsightsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetEntityInsightsResult
	err := ctx.Invoke("azure-native:securityinsights/v20230201preview:getEntityInsights", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEntityInsightsArgs struct {
	// Indicates if query time range should be extended with default time range of the query. Default value is false
	AddDefaultExtendedTimeRange *bool `pulumi:"addDefaultExtendedTimeRange"`
	// The end timeline date, so the results returned are before this date.
	EndTime string `pulumi:"endTime"`
	// entity ID
	EntityId string `pulumi:"entityId"`
	// List of Insights Query Id. If empty, default value is all insights of this entity
	InsightQueryIds []string `pulumi:"insightQueryIds"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The start timeline date, so the results returned are after this date.
	StartTime string `pulumi:"startTime"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The Get Insights result operation response.
type GetEntityInsightsResult struct {
	// The metadata from the get insights operation results.
	MetaData *GetInsightsResultsMetadataResponse `pulumi:"metaData"`
	// The insights result values.
	Value []EntityInsightItemResponse `pulumi:"value"`
}

func GetEntityInsightsOutput(ctx *pulumi.Context, args GetEntityInsightsOutputArgs, opts ...pulumi.InvokeOption) GetEntityInsightsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEntityInsightsResult, error) {
			args := v.(GetEntityInsightsArgs)
			r, err := GetEntityInsights(ctx, &args, opts...)
			var s GetEntityInsightsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEntityInsightsResultOutput)
}

type GetEntityInsightsOutputArgs struct {
	// Indicates if query time range should be extended with default time range of the query. Default value is false
	AddDefaultExtendedTimeRange pulumi.BoolPtrInput `pulumi:"addDefaultExtendedTimeRange"`
	// The end timeline date, so the results returned are before this date.
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// entity ID
	EntityId pulumi.StringInput `pulumi:"entityId"`
	// List of Insights Query Id. If empty, default value is all insights of this entity
	InsightQueryIds pulumi.StringArrayInput `pulumi:"insightQueryIds"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The start timeline date, so the results returned are after this date.
	StartTime pulumi.StringInput `pulumi:"startTime"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (GetEntityInsightsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntityInsightsArgs)(nil)).Elem()
}

// The Get Insights result operation response.
type GetEntityInsightsResultOutput struct{ *pulumi.OutputState }

func (GetEntityInsightsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntityInsightsResult)(nil)).Elem()
}

func (o GetEntityInsightsResultOutput) ToGetEntityInsightsResultOutput() GetEntityInsightsResultOutput {
	return o
}

func (o GetEntityInsightsResultOutput) ToGetEntityInsightsResultOutputWithContext(ctx context.Context) GetEntityInsightsResultOutput {
	return o
}

// The metadata from the get insights operation results.
func (o GetEntityInsightsResultOutput) MetaData() GetInsightsResultsMetadataResponsePtrOutput {
	return o.ApplyT(func(v GetEntityInsightsResult) *GetInsightsResultsMetadataResponse { return v.MetaData }).(GetInsightsResultsMetadataResponsePtrOutput)
}

// The insights result values.
func (o GetEntityInsightsResultOutput) Value() EntityInsightItemResponseArrayOutput {
	return o.ApplyT(func(v GetEntityInsightsResult) []EntityInsightItemResponse { return v.Value }).(EntityInsightItemResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEntityInsightsResultOutput{})
}
