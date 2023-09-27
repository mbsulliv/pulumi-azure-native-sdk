// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Export logs that show total throttled Api requests for this subscription in the given time window.
func GetLogAnalyticExportThrottledRequests(ctx *pulumi.Context, args *GetLogAnalyticExportThrottledRequestsArgs, opts ...pulumi.InvokeOption) (*GetLogAnalyticExportThrottledRequestsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetLogAnalyticExportThrottledRequestsResult
	err := ctx.Invoke("azure-native:compute/v20230301:getLogAnalyticExportThrottledRequests", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLogAnalyticExportThrottledRequestsArgs struct {
	// SAS Uri of the logging blob container to which LogAnalytics Api writes output logs to.
	BlobContainerSasUri string `pulumi:"blobContainerSasUri"`
	// From time of the query
	FromTime string `pulumi:"fromTime"`
	// Group query result by Client Application ID.
	GroupByClientApplicationId *bool `pulumi:"groupByClientApplicationId"`
	// Group query result by Operation Name.
	GroupByOperationName *bool `pulumi:"groupByOperationName"`
	// Group query result by Resource Name.
	GroupByResourceName *bool `pulumi:"groupByResourceName"`
	// Group query result by Throttle Policy applied.
	GroupByThrottlePolicy *bool `pulumi:"groupByThrottlePolicy"`
	// Group query result by User Agent.
	GroupByUserAgent *bool `pulumi:"groupByUserAgent"`
	// The location upon which virtual-machine-sizes is queried.
	Location string `pulumi:"location"`
	// To time of the query
	ToTime string `pulumi:"toTime"`
}

// LogAnalytics operation status response
type GetLogAnalyticExportThrottledRequestsResult struct {
	// LogAnalyticsOutput
	Properties LogAnalyticsOutputResponse `pulumi:"properties"`
}

func GetLogAnalyticExportThrottledRequestsOutput(ctx *pulumi.Context, args GetLogAnalyticExportThrottledRequestsOutputArgs, opts ...pulumi.InvokeOption) GetLogAnalyticExportThrottledRequestsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLogAnalyticExportThrottledRequestsResult, error) {
			args := v.(GetLogAnalyticExportThrottledRequestsArgs)
			r, err := GetLogAnalyticExportThrottledRequests(ctx, &args, opts...)
			var s GetLogAnalyticExportThrottledRequestsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLogAnalyticExportThrottledRequestsResultOutput)
}

type GetLogAnalyticExportThrottledRequestsOutputArgs struct {
	// SAS Uri of the logging blob container to which LogAnalytics Api writes output logs to.
	BlobContainerSasUri pulumi.StringInput `pulumi:"blobContainerSasUri"`
	// From time of the query
	FromTime pulumi.StringInput `pulumi:"fromTime"`
	// Group query result by Client Application ID.
	GroupByClientApplicationId pulumi.BoolPtrInput `pulumi:"groupByClientApplicationId"`
	// Group query result by Operation Name.
	GroupByOperationName pulumi.BoolPtrInput `pulumi:"groupByOperationName"`
	// Group query result by Resource Name.
	GroupByResourceName pulumi.BoolPtrInput `pulumi:"groupByResourceName"`
	// Group query result by Throttle Policy applied.
	GroupByThrottlePolicy pulumi.BoolPtrInput `pulumi:"groupByThrottlePolicy"`
	// Group query result by User Agent.
	GroupByUserAgent pulumi.BoolPtrInput `pulumi:"groupByUserAgent"`
	// The location upon which virtual-machine-sizes is queried.
	Location pulumi.StringInput `pulumi:"location"`
	// To time of the query
	ToTime pulumi.StringInput `pulumi:"toTime"`
}

func (GetLogAnalyticExportThrottledRequestsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogAnalyticExportThrottledRequestsArgs)(nil)).Elem()
}

// LogAnalytics operation status response
type GetLogAnalyticExportThrottledRequestsResultOutput struct{ *pulumi.OutputState }

func (GetLogAnalyticExportThrottledRequestsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogAnalyticExportThrottledRequestsResult)(nil)).Elem()
}

func (o GetLogAnalyticExportThrottledRequestsResultOutput) ToGetLogAnalyticExportThrottledRequestsResultOutput() GetLogAnalyticExportThrottledRequestsResultOutput {
	return o
}

func (o GetLogAnalyticExportThrottledRequestsResultOutput) ToGetLogAnalyticExportThrottledRequestsResultOutputWithContext(ctx context.Context) GetLogAnalyticExportThrottledRequestsResultOutput {
	return o
}

func (o GetLogAnalyticExportThrottledRequestsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetLogAnalyticExportThrottledRequestsResult] {
	return pulumix.Output[GetLogAnalyticExportThrottledRequestsResult]{
		OutputState: o.OutputState,
	}
}

// LogAnalyticsOutput
func (o GetLogAnalyticExportThrottledRequestsResultOutput) Properties() LogAnalyticsOutputResponseOutput {
	return o.ApplyT(func(v GetLogAnalyticExportThrottledRequestsResult) LogAnalyticsOutputResponse { return v.Properties }).(LogAnalyticsOutputResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogAnalyticExportThrottledRequestsResultOutput{})
}
