// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Timeline for an entity.
func GetEntitiesGetTimeline(ctx *pulumi.Context, args *GetEntitiesGetTimelineArgs, opts ...pulumi.InvokeOption) (*GetEntitiesGetTimelineResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetEntitiesGetTimelineResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getEntitiesGetTimeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEntitiesGetTimelineArgs struct {
	// The end timeline date, so the results returned are before this date.
	EndTime string `pulumi:"endTime"`
	// entity ID
	EntityId string `pulumi:"entityId"`
	// Array of timeline Item kinds.
	Kinds []string `pulumi:"kinds"`
	// The number of bucket for timeline queries aggregation.
	NumberOfBucket *int `pulumi:"numberOfBucket"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The start timeline date, so the results returned are after this date.
	StartTime string `pulumi:"startTime"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The entity timeline result operation response.
type GetEntitiesGetTimelineResult struct {
	// The metadata from the timeline operation results.
	MetaData *TimelineResultsMetadataResponse `pulumi:"metaData"`
	// The timeline result values.
	Value []interface{} `pulumi:"value"`
}

func GetEntitiesGetTimelineOutput(ctx *pulumi.Context, args GetEntitiesGetTimelineOutputArgs, opts ...pulumi.InvokeOption) GetEntitiesGetTimelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEntitiesGetTimelineResult, error) {
			args := v.(GetEntitiesGetTimelineArgs)
			r, err := GetEntitiesGetTimeline(ctx, &args, opts...)
			var s GetEntitiesGetTimelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEntitiesGetTimelineResultOutput)
}

type GetEntitiesGetTimelineOutputArgs struct {
	// The end timeline date, so the results returned are before this date.
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// entity ID
	EntityId pulumi.StringInput `pulumi:"entityId"`
	// Array of timeline Item kinds.
	Kinds pulumi.StringArrayInput `pulumi:"kinds"`
	// The number of bucket for timeline queries aggregation.
	NumberOfBucket pulumi.IntPtrInput `pulumi:"numberOfBucket"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The start timeline date, so the results returned are after this date.
	StartTime pulumi.StringInput `pulumi:"startTime"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (GetEntitiesGetTimelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntitiesGetTimelineArgs)(nil)).Elem()
}

// The entity timeline result operation response.
type GetEntitiesGetTimelineResultOutput struct{ *pulumi.OutputState }

func (GetEntitiesGetTimelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntitiesGetTimelineResult)(nil)).Elem()
}

func (o GetEntitiesGetTimelineResultOutput) ToGetEntitiesGetTimelineResultOutput() GetEntitiesGetTimelineResultOutput {
	return o
}

func (o GetEntitiesGetTimelineResultOutput) ToGetEntitiesGetTimelineResultOutputWithContext(ctx context.Context) GetEntitiesGetTimelineResultOutput {
	return o
}

func (o GetEntitiesGetTimelineResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetEntitiesGetTimelineResult] {
	return pulumix.Output[GetEntitiesGetTimelineResult]{
		OutputState: o.OutputState,
	}
}

// The metadata from the timeline operation results.
func (o GetEntitiesGetTimelineResultOutput) MetaData() TimelineResultsMetadataResponsePtrOutput {
	return o.ApplyT(func(v GetEntitiesGetTimelineResult) *TimelineResultsMetadataResponse { return v.MetaData }).(TimelineResultsMetadataResponsePtrOutput)
}

// The timeline result values.
func (o GetEntitiesGetTimelineResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetEntitiesGetTimelineResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEntitiesGetTimelineResultOutput{})
}
