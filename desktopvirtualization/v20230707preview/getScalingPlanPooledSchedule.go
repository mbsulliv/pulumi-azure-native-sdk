// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230707preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a ScalingPlanPooledSchedule.
func LookupScalingPlanPooledSchedule(ctx *pulumi.Context, args *LookupScalingPlanPooledScheduleArgs, opts ...pulumi.InvokeOption) (*LookupScalingPlanPooledScheduleResult, error) {
	var rv LookupScalingPlanPooledScheduleResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20230707preview:getScalingPlanPooledSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScalingPlanPooledScheduleArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the scaling plan.
	ScalingPlanName string `pulumi:"scalingPlanName"`
	// The name of the ScalingPlanSchedule
	ScalingPlanScheduleName string `pulumi:"scalingPlanScheduleName"`
}

// Represents a ScalingPlanPooledSchedule definition.
type LookupScalingPlanPooledScheduleResult struct {
	// Set of days of the week on which this schedule is active.
	DaysOfWeek []string `pulumi:"daysOfWeek"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Load balancing algorithm for off-peak period.
	OffPeakLoadBalancingAlgorithm *string `pulumi:"offPeakLoadBalancingAlgorithm"`
	// Starting time for off-peak period.
	OffPeakStartTime *TimeResponse `pulumi:"offPeakStartTime"`
	// Load balancing algorithm for peak period.
	PeakLoadBalancingAlgorithm *string `pulumi:"peakLoadBalancingAlgorithm"`
	// Starting time for peak period.
	PeakStartTime *TimeResponse `pulumi:"peakStartTime"`
	// Capacity threshold for ramp down period.
	RampDownCapacityThresholdPct *int `pulumi:"rampDownCapacityThresholdPct"`
	// Should users be logged off forcefully from hosts.
	RampDownForceLogoffUsers *bool `pulumi:"rampDownForceLogoffUsers"`
	// Load balancing algorithm for ramp down period.
	RampDownLoadBalancingAlgorithm *string `pulumi:"rampDownLoadBalancingAlgorithm"`
	// Minimum host percentage for ramp down period.
	RampDownMinimumHostsPct *int `pulumi:"rampDownMinimumHostsPct"`
	// Notification message for users during ramp down period.
	RampDownNotificationMessage *string `pulumi:"rampDownNotificationMessage"`
	// Starting time for ramp down period.
	RampDownStartTime *TimeResponse `pulumi:"rampDownStartTime"`
	// Specifies when to stop hosts during ramp down period.
	RampDownStopHostsWhen *string `pulumi:"rampDownStopHostsWhen"`
	// Number of minutes to wait to stop hosts during ramp down period.
	RampDownWaitTimeMinutes *int `pulumi:"rampDownWaitTimeMinutes"`
	// Capacity threshold for ramp up period.
	RampUpCapacityThresholdPct *int `pulumi:"rampUpCapacityThresholdPct"`
	// Load balancing algorithm for ramp up period.
	RampUpLoadBalancingAlgorithm *string `pulumi:"rampUpLoadBalancingAlgorithm"`
	// Minimum host percentage for ramp up period.
	RampUpMinimumHostsPct *int `pulumi:"rampUpMinimumHostsPct"`
	// Starting time for ramp up period.
	RampUpStartTime *TimeResponse `pulumi:"rampUpStartTime"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupScalingPlanPooledScheduleOutput(ctx *pulumi.Context, args LookupScalingPlanPooledScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupScalingPlanPooledScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScalingPlanPooledScheduleResult, error) {
			args := v.(LookupScalingPlanPooledScheduleArgs)
			r, err := LookupScalingPlanPooledSchedule(ctx, &args, opts...)
			var s LookupScalingPlanPooledScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScalingPlanPooledScheduleResultOutput)
}

type LookupScalingPlanPooledScheduleOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the scaling plan.
	ScalingPlanName pulumi.StringInput `pulumi:"scalingPlanName"`
	// The name of the ScalingPlanSchedule
	ScalingPlanScheduleName pulumi.StringInput `pulumi:"scalingPlanScheduleName"`
}

func (LookupScalingPlanPooledScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalingPlanPooledScheduleArgs)(nil)).Elem()
}

// Represents a ScalingPlanPooledSchedule definition.
type LookupScalingPlanPooledScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupScalingPlanPooledScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalingPlanPooledScheduleResult)(nil)).Elem()
}

func (o LookupScalingPlanPooledScheduleResultOutput) ToLookupScalingPlanPooledScheduleResultOutput() LookupScalingPlanPooledScheduleResultOutput {
	return o
}

func (o LookupScalingPlanPooledScheduleResultOutput) ToLookupScalingPlanPooledScheduleResultOutputWithContext(ctx context.Context) LookupScalingPlanPooledScheduleResultOutput {
	return o
}

// Set of days of the week on which this schedule is active.
func (o LookupScalingPlanPooledScheduleResultOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupScalingPlanPooledScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupScalingPlanPooledScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Load balancing algorithm for off-peak period.
func (o LookupScalingPlanPooledScheduleResultOutput) OffPeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *string { return v.OffPeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

// Starting time for off-peak period.
func (o LookupScalingPlanPooledScheduleResultOutput) OffPeakStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *TimeResponse { return v.OffPeakStartTime }).(TimeResponsePtrOutput)
}

// Load balancing algorithm for peak period.
func (o LookupScalingPlanPooledScheduleResultOutput) PeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *string { return v.PeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

// Starting time for peak period.
func (o LookupScalingPlanPooledScheduleResultOutput) PeakStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *TimeResponse { return v.PeakStartTime }).(TimeResponsePtrOutput)
}

// Capacity threshold for ramp down period.
func (o LookupScalingPlanPooledScheduleResultOutput) RampDownCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *int { return v.RampDownCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

// Should users be logged off forcefully from hosts.
func (o LookupScalingPlanPooledScheduleResultOutput) RampDownForceLogoffUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *bool { return v.RampDownForceLogoffUsers }).(pulumi.BoolPtrOutput)
}

// Load balancing algorithm for ramp down period.
func (o LookupScalingPlanPooledScheduleResultOutput) RampDownLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *string { return v.RampDownLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

// Minimum host percentage for ramp down period.
func (o LookupScalingPlanPooledScheduleResultOutput) RampDownMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *int { return v.RampDownMinimumHostsPct }).(pulumi.IntPtrOutput)
}

// Notification message for users during ramp down period.
func (o LookupScalingPlanPooledScheduleResultOutput) RampDownNotificationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *string { return v.RampDownNotificationMessage }).(pulumi.StringPtrOutput)
}

// Starting time for ramp down period.
func (o LookupScalingPlanPooledScheduleResultOutput) RampDownStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *TimeResponse { return v.RampDownStartTime }).(TimeResponsePtrOutput)
}

// Specifies when to stop hosts during ramp down period.
func (o LookupScalingPlanPooledScheduleResultOutput) RampDownStopHostsWhen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *string { return v.RampDownStopHostsWhen }).(pulumi.StringPtrOutput)
}

// Number of minutes to wait to stop hosts during ramp down period.
func (o LookupScalingPlanPooledScheduleResultOutput) RampDownWaitTimeMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *int { return v.RampDownWaitTimeMinutes }).(pulumi.IntPtrOutput)
}

// Capacity threshold for ramp up period.
func (o LookupScalingPlanPooledScheduleResultOutput) RampUpCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *int { return v.RampUpCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

// Load balancing algorithm for ramp up period.
func (o LookupScalingPlanPooledScheduleResultOutput) RampUpLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *string { return v.RampUpLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

// Minimum host percentage for ramp up period.
func (o LookupScalingPlanPooledScheduleResultOutput) RampUpMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *int { return v.RampUpMinimumHostsPct }).(pulumi.IntPtrOutput)
}

// Starting time for ramp up period.
func (o LookupScalingPlanPooledScheduleResultOutput) RampUpStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *TimeResponse { return v.RampUpStartTime }).(TimeResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupScalingPlanPooledScheduleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupScalingPlanPooledScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScalingPlanPooledScheduleResultOutput{})
}
