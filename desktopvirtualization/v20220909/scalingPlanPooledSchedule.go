// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220909

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a ScalingPlanPooledSchedule definition.
type ScalingPlanPooledSchedule struct {
	pulumi.CustomResourceState

	// Set of days of the week on which this schedule is active.
	DaysOfWeek pulumi.StringArrayOutput `pulumi:"daysOfWeek"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Load balancing algorithm for off-peak period.
	OffPeakLoadBalancingAlgorithm pulumi.StringPtrOutput `pulumi:"offPeakLoadBalancingAlgorithm"`
	// Starting time for off-peak period.
	OffPeakStartTime TimeResponsePtrOutput `pulumi:"offPeakStartTime"`
	// Load balancing algorithm for peak period.
	PeakLoadBalancingAlgorithm pulumi.StringPtrOutput `pulumi:"peakLoadBalancingAlgorithm"`
	// Starting time for peak period.
	PeakStartTime TimeResponsePtrOutput `pulumi:"peakStartTime"`
	// Capacity threshold for ramp down period.
	RampDownCapacityThresholdPct pulumi.IntPtrOutput `pulumi:"rampDownCapacityThresholdPct"`
	// Should users be logged off forcefully from hosts.
	RampDownForceLogoffUsers pulumi.BoolPtrOutput `pulumi:"rampDownForceLogoffUsers"`
	// Load balancing algorithm for ramp down period.
	RampDownLoadBalancingAlgorithm pulumi.StringPtrOutput `pulumi:"rampDownLoadBalancingAlgorithm"`
	// Minimum host percentage for ramp down period.
	RampDownMinimumHostsPct pulumi.IntPtrOutput `pulumi:"rampDownMinimumHostsPct"`
	// Notification message for users during ramp down period.
	RampDownNotificationMessage pulumi.StringPtrOutput `pulumi:"rampDownNotificationMessage"`
	// Starting time for ramp down period.
	RampDownStartTime TimeResponsePtrOutput `pulumi:"rampDownStartTime"`
	// Specifies when to stop hosts during ramp down period.
	RampDownStopHostsWhen pulumi.StringPtrOutput `pulumi:"rampDownStopHostsWhen"`
	// Number of minutes to wait to stop hosts during ramp down period.
	RampDownWaitTimeMinutes pulumi.IntPtrOutput `pulumi:"rampDownWaitTimeMinutes"`
	// Capacity threshold for ramp up period.
	RampUpCapacityThresholdPct pulumi.IntPtrOutput `pulumi:"rampUpCapacityThresholdPct"`
	// Load balancing algorithm for ramp up period.
	RampUpLoadBalancingAlgorithm pulumi.StringPtrOutput `pulumi:"rampUpLoadBalancingAlgorithm"`
	// Minimum host percentage for ramp up period.
	RampUpMinimumHostsPct pulumi.IntPtrOutput `pulumi:"rampUpMinimumHostsPct"`
	// Starting time for ramp up period.
	RampUpStartTime TimeResponsePtrOutput `pulumi:"rampUpStartTime"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewScalingPlanPooledSchedule registers a new resource with the given unique name, arguments, and options.
func NewScalingPlanPooledSchedule(ctx *pulumi.Context,
	name string, args *ScalingPlanPooledScheduleArgs, opts ...pulumi.ResourceOption) (*ScalingPlanPooledSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScalingPlanName == nil {
		return nil, errors.New("invalid value for required argument 'ScalingPlanName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:ScalingPlanPooledSchedule"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:ScalingPlanPooledSchedule"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20221014preview:ScalingPlanPooledSchedule"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20230707preview:ScalingPlanPooledSchedule"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20230905:ScalingPlanPooledSchedule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ScalingPlanPooledSchedule
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20220909:ScalingPlanPooledSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingPlanPooledSchedule gets an existing ScalingPlanPooledSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingPlanPooledSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingPlanPooledScheduleState, opts ...pulumi.ResourceOption) (*ScalingPlanPooledSchedule, error) {
	var resource ScalingPlanPooledSchedule
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20220909:ScalingPlanPooledSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingPlanPooledSchedule resources.
type scalingPlanPooledScheduleState struct {
}

type ScalingPlanPooledScheduleState struct {
}

func (ScalingPlanPooledScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanPooledScheduleState)(nil)).Elem()
}

type scalingPlanPooledScheduleArgs struct {
	// Set of days of the week on which this schedule is active.
	DaysOfWeek []string `pulumi:"daysOfWeek"`
	// Load balancing algorithm for off-peak period.
	OffPeakLoadBalancingAlgorithm *string `pulumi:"offPeakLoadBalancingAlgorithm"`
	// Starting time for off-peak period.
	OffPeakStartTime *Time `pulumi:"offPeakStartTime"`
	// Load balancing algorithm for peak period.
	PeakLoadBalancingAlgorithm *string `pulumi:"peakLoadBalancingAlgorithm"`
	// Starting time for peak period.
	PeakStartTime *Time `pulumi:"peakStartTime"`
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
	RampDownStartTime *Time `pulumi:"rampDownStartTime"`
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
	RampUpStartTime *Time `pulumi:"rampUpStartTime"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the scaling plan.
	ScalingPlanName string `pulumi:"scalingPlanName"`
	// The name of the ScalingPlanSchedule
	ScalingPlanScheduleName *string `pulumi:"scalingPlanScheduleName"`
}

// The set of arguments for constructing a ScalingPlanPooledSchedule resource.
type ScalingPlanPooledScheduleArgs struct {
	// Set of days of the week on which this schedule is active.
	DaysOfWeek pulumi.StringArrayInput
	// Load balancing algorithm for off-peak period.
	OffPeakLoadBalancingAlgorithm pulumi.StringPtrInput
	// Starting time for off-peak period.
	OffPeakStartTime TimePtrInput
	// Load balancing algorithm for peak period.
	PeakLoadBalancingAlgorithm pulumi.StringPtrInput
	// Starting time for peak period.
	PeakStartTime TimePtrInput
	// Capacity threshold for ramp down period.
	RampDownCapacityThresholdPct pulumi.IntPtrInput
	// Should users be logged off forcefully from hosts.
	RampDownForceLogoffUsers pulumi.BoolPtrInput
	// Load balancing algorithm for ramp down period.
	RampDownLoadBalancingAlgorithm pulumi.StringPtrInput
	// Minimum host percentage for ramp down period.
	RampDownMinimumHostsPct pulumi.IntPtrInput
	// Notification message for users during ramp down period.
	RampDownNotificationMessage pulumi.StringPtrInput
	// Starting time for ramp down period.
	RampDownStartTime TimePtrInput
	// Specifies when to stop hosts during ramp down period.
	RampDownStopHostsWhen pulumi.StringPtrInput
	// Number of minutes to wait to stop hosts during ramp down period.
	RampDownWaitTimeMinutes pulumi.IntPtrInput
	// Capacity threshold for ramp up period.
	RampUpCapacityThresholdPct pulumi.IntPtrInput
	// Load balancing algorithm for ramp up period.
	RampUpLoadBalancingAlgorithm pulumi.StringPtrInput
	// Minimum host percentage for ramp up period.
	RampUpMinimumHostsPct pulumi.IntPtrInput
	// Starting time for ramp up period.
	RampUpStartTime TimePtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the scaling plan.
	ScalingPlanName pulumi.StringInput
	// The name of the ScalingPlanSchedule
	ScalingPlanScheduleName pulumi.StringPtrInput
}

func (ScalingPlanPooledScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanPooledScheduleArgs)(nil)).Elem()
}

type ScalingPlanPooledScheduleInput interface {
	pulumi.Input

	ToScalingPlanPooledScheduleOutput() ScalingPlanPooledScheduleOutput
	ToScalingPlanPooledScheduleOutputWithContext(ctx context.Context) ScalingPlanPooledScheduleOutput
}

func (*ScalingPlanPooledSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPlanPooledSchedule)(nil)).Elem()
}

func (i *ScalingPlanPooledSchedule) ToScalingPlanPooledScheduleOutput() ScalingPlanPooledScheduleOutput {
	return i.ToScalingPlanPooledScheduleOutputWithContext(context.Background())
}

func (i *ScalingPlanPooledSchedule) ToScalingPlanPooledScheduleOutputWithContext(ctx context.Context) ScalingPlanPooledScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPlanPooledScheduleOutput)
}

func (i *ScalingPlanPooledSchedule) ToOutput(ctx context.Context) pulumix.Output[*ScalingPlanPooledSchedule] {
	return pulumix.Output[*ScalingPlanPooledSchedule]{
		OutputState: i.ToScalingPlanPooledScheduleOutputWithContext(ctx).OutputState,
	}
}

type ScalingPlanPooledScheduleOutput struct{ *pulumi.OutputState }

func (ScalingPlanPooledScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPlanPooledSchedule)(nil)).Elem()
}

func (o ScalingPlanPooledScheduleOutput) ToScalingPlanPooledScheduleOutput() ScalingPlanPooledScheduleOutput {
	return o
}

func (o ScalingPlanPooledScheduleOutput) ToScalingPlanPooledScheduleOutputWithContext(ctx context.Context) ScalingPlanPooledScheduleOutput {
	return o
}

func (o ScalingPlanPooledScheduleOutput) ToOutput(ctx context.Context) pulumix.Output[*ScalingPlanPooledSchedule] {
	return pulumix.Output[*ScalingPlanPooledSchedule]{
		OutputState: o.OutputState,
	}
}

// Set of days of the week on which this schedule is active.
func (o ScalingPlanPooledScheduleOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringArrayOutput { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o ScalingPlanPooledScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Load balancing algorithm for off-peak period.
func (o ScalingPlanPooledScheduleOutput) OffPeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringPtrOutput { return v.OffPeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

// Starting time for off-peak period.
func (o ScalingPlanPooledScheduleOutput) OffPeakStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) TimeResponsePtrOutput { return v.OffPeakStartTime }).(TimeResponsePtrOutput)
}

// Load balancing algorithm for peak period.
func (o ScalingPlanPooledScheduleOutput) PeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringPtrOutput { return v.PeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

// Starting time for peak period.
func (o ScalingPlanPooledScheduleOutput) PeakStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) TimeResponsePtrOutput { return v.PeakStartTime }).(TimeResponsePtrOutput)
}

// Capacity threshold for ramp down period.
func (o ScalingPlanPooledScheduleOutput) RampDownCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.IntPtrOutput { return v.RampDownCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

// Should users be logged off forcefully from hosts.
func (o ScalingPlanPooledScheduleOutput) RampDownForceLogoffUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.BoolPtrOutput { return v.RampDownForceLogoffUsers }).(pulumi.BoolPtrOutput)
}

// Load balancing algorithm for ramp down period.
func (o ScalingPlanPooledScheduleOutput) RampDownLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringPtrOutput { return v.RampDownLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

// Minimum host percentage for ramp down period.
func (o ScalingPlanPooledScheduleOutput) RampDownMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.IntPtrOutput { return v.RampDownMinimumHostsPct }).(pulumi.IntPtrOutput)
}

// Notification message for users during ramp down period.
func (o ScalingPlanPooledScheduleOutput) RampDownNotificationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringPtrOutput { return v.RampDownNotificationMessage }).(pulumi.StringPtrOutput)
}

// Starting time for ramp down period.
func (o ScalingPlanPooledScheduleOutput) RampDownStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) TimeResponsePtrOutput { return v.RampDownStartTime }).(TimeResponsePtrOutput)
}

// Specifies when to stop hosts during ramp down period.
func (o ScalingPlanPooledScheduleOutput) RampDownStopHostsWhen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringPtrOutput { return v.RampDownStopHostsWhen }).(pulumi.StringPtrOutput)
}

// Number of minutes to wait to stop hosts during ramp down period.
func (o ScalingPlanPooledScheduleOutput) RampDownWaitTimeMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.IntPtrOutput { return v.RampDownWaitTimeMinutes }).(pulumi.IntPtrOutput)
}

// Capacity threshold for ramp up period.
func (o ScalingPlanPooledScheduleOutput) RampUpCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.IntPtrOutput { return v.RampUpCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

// Load balancing algorithm for ramp up period.
func (o ScalingPlanPooledScheduleOutput) RampUpLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringPtrOutput { return v.RampUpLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

// Minimum host percentage for ramp up period.
func (o ScalingPlanPooledScheduleOutput) RampUpMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.IntPtrOutput { return v.RampUpMinimumHostsPct }).(pulumi.IntPtrOutput)
}

// Starting time for ramp up period.
func (o ScalingPlanPooledScheduleOutput) RampUpStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) TimeResponsePtrOutput { return v.RampUpStartTime }).(TimeResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ScalingPlanPooledScheduleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ScalingPlanPooledScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalingPlanPooledScheduleOutput{})
}
