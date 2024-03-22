// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231004preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a ScalingPlanPersonalSchedule definition.
type ScalingPlanPersonalSchedule struct {
	pulumi.CustomResourceState

	// Set of days of the week on which this schedule is active.
	DaysOfWeek pulumi.StringArrayOutput `pulumi:"daysOfWeek"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Action to be taken after a user disconnect during the off-peak period.
	OffPeakActionOnDisconnect pulumi.StringPtrOutput `pulumi:"offPeakActionOnDisconnect"`
	// Action to be taken after a logoff during the off-peak period.
	OffPeakActionOnLogoff pulumi.StringPtrOutput `pulumi:"offPeakActionOnLogoff"`
	// The time in minutes to wait before performing the desired session handling action when a user disconnects during the off-peak period.
	OffPeakMinutesToWaitOnDisconnect pulumi.IntPtrOutput `pulumi:"offPeakMinutesToWaitOnDisconnect"`
	// The time in minutes to wait before performing the desired session handling action when a user logs off during the off-peak period.
	OffPeakMinutesToWaitOnLogoff pulumi.IntPtrOutput `pulumi:"offPeakMinutesToWaitOnLogoff"`
	// Starting time for off-peak period.
	OffPeakStartTime TimeResponsePtrOutput `pulumi:"offPeakStartTime"`
	// The desired configuration of Start VM On Connect for the hostpool during the off-peak phase.
	OffPeakStartVMOnConnect pulumi.StringPtrOutput `pulumi:"offPeakStartVMOnConnect"`
	// Action to be taken after a user disconnect during the peak period.
	PeakActionOnDisconnect pulumi.StringPtrOutput `pulumi:"peakActionOnDisconnect"`
	// Action to be taken after a logoff during the peak period.
	PeakActionOnLogoff pulumi.StringPtrOutput `pulumi:"peakActionOnLogoff"`
	// The time in minutes to wait before performing the desired session handling action when a user disconnects during the peak period.
	PeakMinutesToWaitOnDisconnect pulumi.IntPtrOutput `pulumi:"peakMinutesToWaitOnDisconnect"`
	// The time in minutes to wait before performing the desired session handling action when a user logs off during the peak period.
	PeakMinutesToWaitOnLogoff pulumi.IntPtrOutput `pulumi:"peakMinutesToWaitOnLogoff"`
	// Starting time for peak period.
	PeakStartTime TimeResponsePtrOutput `pulumi:"peakStartTime"`
	// The desired configuration of Start VM On Connect for the hostpool during the peak phase.
	PeakStartVMOnConnect pulumi.StringPtrOutput `pulumi:"peakStartVMOnConnect"`
	// Action to be taken after a user disconnect during the ramp down period.
	RampDownActionOnDisconnect pulumi.StringPtrOutput `pulumi:"rampDownActionOnDisconnect"`
	// Action to be taken after a logoff during the ramp down period.
	RampDownActionOnLogoff pulumi.StringPtrOutput `pulumi:"rampDownActionOnLogoff"`
	// The time in minutes to wait before performing the desired session handling action when a user disconnects during the ramp down period.
	RampDownMinutesToWaitOnDisconnect pulumi.IntPtrOutput `pulumi:"rampDownMinutesToWaitOnDisconnect"`
	// The time in minutes to wait before performing the desired session handling action when a user logs off during the ramp down period.
	RampDownMinutesToWaitOnLogoff pulumi.IntPtrOutput `pulumi:"rampDownMinutesToWaitOnLogoff"`
	// Starting time for ramp down period.
	RampDownStartTime TimeResponsePtrOutput `pulumi:"rampDownStartTime"`
	// The desired configuration of Start VM On Connect for the hostpool during the ramp down phase.
	RampDownStartVMOnConnect pulumi.StringPtrOutput `pulumi:"rampDownStartVMOnConnect"`
	// Action to be taken after a user disconnect during the ramp up period.
	RampUpActionOnDisconnect pulumi.StringPtrOutput `pulumi:"rampUpActionOnDisconnect"`
	// Action to be taken after a logoff during the ramp up period.
	RampUpActionOnLogoff pulumi.StringPtrOutput `pulumi:"rampUpActionOnLogoff"`
	// The desired startup behavior during the ramp up period for personal vms in the hostpool.
	RampUpAutoStartHosts pulumi.StringPtrOutput `pulumi:"rampUpAutoStartHosts"`
	// The time in minutes to wait before performing the desired session handling action when a user disconnects during the ramp up period.
	RampUpMinutesToWaitOnDisconnect pulumi.IntPtrOutput `pulumi:"rampUpMinutesToWaitOnDisconnect"`
	// The time in minutes to wait before performing the desired session handling action when a user logs off during the ramp up period.
	RampUpMinutesToWaitOnLogoff pulumi.IntPtrOutput `pulumi:"rampUpMinutesToWaitOnLogoff"`
	// Starting time for ramp up period.
	RampUpStartTime TimeResponsePtrOutput `pulumi:"rampUpStartTime"`
	// The desired configuration of Start VM On Connect for the hostpool during the ramp up phase. If this is disabled, session hosts must be turned on using rampUpAutoStartHosts or by turning them on manually.
	RampUpStartVMOnConnect pulumi.StringPtrOutput `pulumi:"rampUpStartVMOnConnect"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewScalingPlanPersonalSchedule registers a new resource with the given unique name, arguments, and options.
func NewScalingPlanPersonalSchedule(ctx *pulumi.Context,
	name string, args *ScalingPlanPersonalScheduleArgs, opts ...pulumi.ResourceOption) (*ScalingPlanPersonalSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScalingPlanName == nil {
		return nil, errors.New("invalid value for required argument 'ScalingPlanName'")
	}
	if args.OffPeakStartVMOnConnect == nil {
		args.OffPeakStartVMOnConnect = pulumi.StringPtr("Enable")
	}
	if args.PeakStartVMOnConnect == nil {
		args.PeakStartVMOnConnect = pulumi.StringPtr("Enable")
	}
	if args.RampDownStartVMOnConnect == nil {
		args.RampDownStartVMOnConnect = pulumi.StringPtr("Enable")
	}
	if args.RampUpStartVMOnConnect == nil {
		args.RampUpStartVMOnConnect = pulumi.StringPtr("Enable")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:ScalingPlanPersonalSchedule"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20230707preview:ScalingPlanPersonalSchedule"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20230905:ScalingPlanPersonalSchedule"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20231101preview:ScalingPlanPersonalSchedule"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20240116preview:ScalingPlanPersonalSchedule"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20240306preview:ScalingPlanPersonalSchedule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ScalingPlanPersonalSchedule
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20231004preview:ScalingPlanPersonalSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingPlanPersonalSchedule gets an existing ScalingPlanPersonalSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingPlanPersonalSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingPlanPersonalScheduleState, opts ...pulumi.ResourceOption) (*ScalingPlanPersonalSchedule, error) {
	var resource ScalingPlanPersonalSchedule
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20231004preview:ScalingPlanPersonalSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingPlanPersonalSchedule resources.
type scalingPlanPersonalScheduleState struct {
}

type ScalingPlanPersonalScheduleState struct {
}

func (ScalingPlanPersonalScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanPersonalScheduleState)(nil)).Elem()
}

type scalingPlanPersonalScheduleArgs struct {
	// Set of days of the week on which this schedule is active.
	DaysOfWeek []string `pulumi:"daysOfWeek"`
	// Action to be taken after a user disconnect during the off-peak period.
	OffPeakActionOnDisconnect *string `pulumi:"offPeakActionOnDisconnect"`
	// Action to be taken after a logoff during the off-peak period.
	OffPeakActionOnLogoff *string `pulumi:"offPeakActionOnLogoff"`
	// The time in minutes to wait before performing the desired session handling action when a user disconnects during the off-peak period.
	OffPeakMinutesToWaitOnDisconnect *int `pulumi:"offPeakMinutesToWaitOnDisconnect"`
	// The time in minutes to wait before performing the desired session handling action when a user logs off during the off-peak period.
	OffPeakMinutesToWaitOnLogoff *int `pulumi:"offPeakMinutesToWaitOnLogoff"`
	// Starting time for off-peak period.
	OffPeakStartTime *Time `pulumi:"offPeakStartTime"`
	// The desired configuration of Start VM On Connect for the hostpool during the off-peak phase.
	OffPeakStartVMOnConnect *string `pulumi:"offPeakStartVMOnConnect"`
	// Action to be taken after a user disconnect during the peak period.
	PeakActionOnDisconnect *string `pulumi:"peakActionOnDisconnect"`
	// Action to be taken after a logoff during the peak period.
	PeakActionOnLogoff *string `pulumi:"peakActionOnLogoff"`
	// The time in minutes to wait before performing the desired session handling action when a user disconnects during the peak period.
	PeakMinutesToWaitOnDisconnect *int `pulumi:"peakMinutesToWaitOnDisconnect"`
	// The time in minutes to wait before performing the desired session handling action when a user logs off during the peak period.
	PeakMinutesToWaitOnLogoff *int `pulumi:"peakMinutesToWaitOnLogoff"`
	// Starting time for peak period.
	PeakStartTime *Time `pulumi:"peakStartTime"`
	// The desired configuration of Start VM On Connect for the hostpool during the peak phase.
	PeakStartVMOnConnect *string `pulumi:"peakStartVMOnConnect"`
	// Action to be taken after a user disconnect during the ramp down period.
	RampDownActionOnDisconnect *string `pulumi:"rampDownActionOnDisconnect"`
	// Action to be taken after a logoff during the ramp down period.
	RampDownActionOnLogoff *string `pulumi:"rampDownActionOnLogoff"`
	// The time in minutes to wait before performing the desired session handling action when a user disconnects during the ramp down period.
	RampDownMinutesToWaitOnDisconnect *int `pulumi:"rampDownMinutesToWaitOnDisconnect"`
	// The time in minutes to wait before performing the desired session handling action when a user logs off during the ramp down period.
	RampDownMinutesToWaitOnLogoff *int `pulumi:"rampDownMinutesToWaitOnLogoff"`
	// Starting time for ramp down period.
	RampDownStartTime *Time `pulumi:"rampDownStartTime"`
	// The desired configuration of Start VM On Connect for the hostpool during the ramp down phase.
	RampDownStartVMOnConnect *string `pulumi:"rampDownStartVMOnConnect"`
	// Action to be taken after a user disconnect during the ramp up period.
	RampUpActionOnDisconnect *string `pulumi:"rampUpActionOnDisconnect"`
	// Action to be taken after a logoff during the ramp up period.
	RampUpActionOnLogoff *string `pulumi:"rampUpActionOnLogoff"`
	// The desired startup behavior during the ramp up period for personal vms in the hostpool.
	RampUpAutoStartHosts *string `pulumi:"rampUpAutoStartHosts"`
	// The time in minutes to wait before performing the desired session handling action when a user disconnects during the ramp up period.
	RampUpMinutesToWaitOnDisconnect *int `pulumi:"rampUpMinutesToWaitOnDisconnect"`
	// The time in minutes to wait before performing the desired session handling action when a user logs off during the ramp up period.
	RampUpMinutesToWaitOnLogoff *int `pulumi:"rampUpMinutesToWaitOnLogoff"`
	// Starting time for ramp up period.
	RampUpStartTime *Time `pulumi:"rampUpStartTime"`
	// The desired configuration of Start VM On Connect for the hostpool during the ramp up phase. If this is disabled, session hosts must be turned on using rampUpAutoStartHosts or by turning them on manually.
	RampUpStartVMOnConnect *string `pulumi:"rampUpStartVMOnConnect"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the scaling plan.
	ScalingPlanName string `pulumi:"scalingPlanName"`
	// The name of the ScalingPlanSchedule
	ScalingPlanScheduleName *string `pulumi:"scalingPlanScheduleName"`
}

// The set of arguments for constructing a ScalingPlanPersonalSchedule resource.
type ScalingPlanPersonalScheduleArgs struct {
	// Set of days of the week on which this schedule is active.
	DaysOfWeek pulumi.StringArrayInput
	// Action to be taken after a user disconnect during the off-peak period.
	OffPeakActionOnDisconnect pulumi.StringPtrInput
	// Action to be taken after a logoff during the off-peak period.
	OffPeakActionOnLogoff pulumi.StringPtrInput
	// The time in minutes to wait before performing the desired session handling action when a user disconnects during the off-peak period.
	OffPeakMinutesToWaitOnDisconnect pulumi.IntPtrInput
	// The time in minutes to wait before performing the desired session handling action when a user logs off during the off-peak period.
	OffPeakMinutesToWaitOnLogoff pulumi.IntPtrInput
	// Starting time for off-peak period.
	OffPeakStartTime TimePtrInput
	// The desired configuration of Start VM On Connect for the hostpool during the off-peak phase.
	OffPeakStartVMOnConnect pulumi.StringPtrInput
	// Action to be taken after a user disconnect during the peak period.
	PeakActionOnDisconnect pulumi.StringPtrInput
	// Action to be taken after a logoff during the peak period.
	PeakActionOnLogoff pulumi.StringPtrInput
	// The time in minutes to wait before performing the desired session handling action when a user disconnects during the peak period.
	PeakMinutesToWaitOnDisconnect pulumi.IntPtrInput
	// The time in minutes to wait before performing the desired session handling action when a user logs off during the peak period.
	PeakMinutesToWaitOnLogoff pulumi.IntPtrInput
	// Starting time for peak period.
	PeakStartTime TimePtrInput
	// The desired configuration of Start VM On Connect for the hostpool during the peak phase.
	PeakStartVMOnConnect pulumi.StringPtrInput
	// Action to be taken after a user disconnect during the ramp down period.
	RampDownActionOnDisconnect pulumi.StringPtrInput
	// Action to be taken after a logoff during the ramp down period.
	RampDownActionOnLogoff pulumi.StringPtrInput
	// The time in minutes to wait before performing the desired session handling action when a user disconnects during the ramp down period.
	RampDownMinutesToWaitOnDisconnect pulumi.IntPtrInput
	// The time in minutes to wait before performing the desired session handling action when a user logs off during the ramp down period.
	RampDownMinutesToWaitOnLogoff pulumi.IntPtrInput
	// Starting time for ramp down period.
	RampDownStartTime TimePtrInput
	// The desired configuration of Start VM On Connect for the hostpool during the ramp down phase.
	RampDownStartVMOnConnect pulumi.StringPtrInput
	// Action to be taken after a user disconnect during the ramp up period.
	RampUpActionOnDisconnect pulumi.StringPtrInput
	// Action to be taken after a logoff during the ramp up period.
	RampUpActionOnLogoff pulumi.StringPtrInput
	// The desired startup behavior during the ramp up period for personal vms in the hostpool.
	RampUpAutoStartHosts pulumi.StringPtrInput
	// The time in minutes to wait before performing the desired session handling action when a user disconnects during the ramp up period.
	RampUpMinutesToWaitOnDisconnect pulumi.IntPtrInput
	// The time in minutes to wait before performing the desired session handling action when a user logs off during the ramp up period.
	RampUpMinutesToWaitOnLogoff pulumi.IntPtrInput
	// Starting time for ramp up period.
	RampUpStartTime TimePtrInput
	// The desired configuration of Start VM On Connect for the hostpool during the ramp up phase. If this is disabled, session hosts must be turned on using rampUpAutoStartHosts or by turning them on manually.
	RampUpStartVMOnConnect pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the scaling plan.
	ScalingPlanName pulumi.StringInput
	// The name of the ScalingPlanSchedule
	ScalingPlanScheduleName pulumi.StringPtrInput
}

func (ScalingPlanPersonalScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanPersonalScheduleArgs)(nil)).Elem()
}

type ScalingPlanPersonalScheduleInput interface {
	pulumi.Input

	ToScalingPlanPersonalScheduleOutput() ScalingPlanPersonalScheduleOutput
	ToScalingPlanPersonalScheduleOutputWithContext(ctx context.Context) ScalingPlanPersonalScheduleOutput
}

func (*ScalingPlanPersonalSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPlanPersonalSchedule)(nil)).Elem()
}

func (i *ScalingPlanPersonalSchedule) ToScalingPlanPersonalScheduleOutput() ScalingPlanPersonalScheduleOutput {
	return i.ToScalingPlanPersonalScheduleOutputWithContext(context.Background())
}

func (i *ScalingPlanPersonalSchedule) ToScalingPlanPersonalScheduleOutputWithContext(ctx context.Context) ScalingPlanPersonalScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPlanPersonalScheduleOutput)
}

type ScalingPlanPersonalScheduleOutput struct{ *pulumi.OutputState }

func (ScalingPlanPersonalScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPlanPersonalSchedule)(nil)).Elem()
}

func (o ScalingPlanPersonalScheduleOutput) ToScalingPlanPersonalScheduleOutput() ScalingPlanPersonalScheduleOutput {
	return o
}

func (o ScalingPlanPersonalScheduleOutput) ToScalingPlanPersonalScheduleOutputWithContext(ctx context.Context) ScalingPlanPersonalScheduleOutput {
	return o
}

// Set of days of the week on which this schedule is active.
func (o ScalingPlanPersonalScheduleOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringArrayOutput { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o ScalingPlanPersonalScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Action to be taken after a user disconnect during the off-peak period.
func (o ScalingPlanPersonalScheduleOutput) OffPeakActionOnDisconnect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringPtrOutput { return v.OffPeakActionOnDisconnect }).(pulumi.StringPtrOutput)
}

// Action to be taken after a logoff during the off-peak period.
func (o ScalingPlanPersonalScheduleOutput) OffPeakActionOnLogoff() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringPtrOutput { return v.OffPeakActionOnLogoff }).(pulumi.StringPtrOutput)
}

// The time in minutes to wait before performing the desired session handling action when a user disconnects during the off-peak period.
func (o ScalingPlanPersonalScheduleOutput) OffPeakMinutesToWaitOnDisconnect() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.IntPtrOutput { return v.OffPeakMinutesToWaitOnDisconnect }).(pulumi.IntPtrOutput)
}

// The time in minutes to wait before performing the desired session handling action when a user logs off during the off-peak period.
func (o ScalingPlanPersonalScheduleOutput) OffPeakMinutesToWaitOnLogoff() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.IntPtrOutput { return v.OffPeakMinutesToWaitOnLogoff }).(pulumi.IntPtrOutput)
}

// Starting time for off-peak period.
func (o ScalingPlanPersonalScheduleOutput) OffPeakStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) TimeResponsePtrOutput { return v.OffPeakStartTime }).(TimeResponsePtrOutput)
}

// The desired configuration of Start VM On Connect for the hostpool during the off-peak phase.
func (o ScalingPlanPersonalScheduleOutput) OffPeakStartVMOnConnect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringPtrOutput { return v.OffPeakStartVMOnConnect }).(pulumi.StringPtrOutput)
}

// Action to be taken after a user disconnect during the peak period.
func (o ScalingPlanPersonalScheduleOutput) PeakActionOnDisconnect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringPtrOutput { return v.PeakActionOnDisconnect }).(pulumi.StringPtrOutput)
}

// Action to be taken after a logoff during the peak period.
func (o ScalingPlanPersonalScheduleOutput) PeakActionOnLogoff() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringPtrOutput { return v.PeakActionOnLogoff }).(pulumi.StringPtrOutput)
}

// The time in minutes to wait before performing the desired session handling action when a user disconnects during the peak period.
func (o ScalingPlanPersonalScheduleOutput) PeakMinutesToWaitOnDisconnect() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.IntPtrOutput { return v.PeakMinutesToWaitOnDisconnect }).(pulumi.IntPtrOutput)
}

// The time in minutes to wait before performing the desired session handling action when a user logs off during the peak period.
func (o ScalingPlanPersonalScheduleOutput) PeakMinutesToWaitOnLogoff() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.IntPtrOutput { return v.PeakMinutesToWaitOnLogoff }).(pulumi.IntPtrOutput)
}

// Starting time for peak period.
func (o ScalingPlanPersonalScheduleOutput) PeakStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) TimeResponsePtrOutput { return v.PeakStartTime }).(TimeResponsePtrOutput)
}

// The desired configuration of Start VM On Connect for the hostpool during the peak phase.
func (o ScalingPlanPersonalScheduleOutput) PeakStartVMOnConnect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringPtrOutput { return v.PeakStartVMOnConnect }).(pulumi.StringPtrOutput)
}

// Action to be taken after a user disconnect during the ramp down period.
func (o ScalingPlanPersonalScheduleOutput) RampDownActionOnDisconnect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringPtrOutput { return v.RampDownActionOnDisconnect }).(pulumi.StringPtrOutput)
}

// Action to be taken after a logoff during the ramp down period.
func (o ScalingPlanPersonalScheduleOutput) RampDownActionOnLogoff() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringPtrOutput { return v.RampDownActionOnLogoff }).(pulumi.StringPtrOutput)
}

// The time in minutes to wait before performing the desired session handling action when a user disconnects during the ramp down period.
func (o ScalingPlanPersonalScheduleOutput) RampDownMinutesToWaitOnDisconnect() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.IntPtrOutput { return v.RampDownMinutesToWaitOnDisconnect }).(pulumi.IntPtrOutput)
}

// The time in minutes to wait before performing the desired session handling action when a user logs off during the ramp down period.
func (o ScalingPlanPersonalScheduleOutput) RampDownMinutesToWaitOnLogoff() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.IntPtrOutput { return v.RampDownMinutesToWaitOnLogoff }).(pulumi.IntPtrOutput)
}

// Starting time for ramp down period.
func (o ScalingPlanPersonalScheduleOutput) RampDownStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) TimeResponsePtrOutput { return v.RampDownStartTime }).(TimeResponsePtrOutput)
}

// The desired configuration of Start VM On Connect for the hostpool during the ramp down phase.
func (o ScalingPlanPersonalScheduleOutput) RampDownStartVMOnConnect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringPtrOutput { return v.RampDownStartVMOnConnect }).(pulumi.StringPtrOutput)
}

// Action to be taken after a user disconnect during the ramp up period.
func (o ScalingPlanPersonalScheduleOutput) RampUpActionOnDisconnect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringPtrOutput { return v.RampUpActionOnDisconnect }).(pulumi.StringPtrOutput)
}

// Action to be taken after a logoff during the ramp up period.
func (o ScalingPlanPersonalScheduleOutput) RampUpActionOnLogoff() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringPtrOutput { return v.RampUpActionOnLogoff }).(pulumi.StringPtrOutput)
}

// The desired startup behavior during the ramp up period for personal vms in the hostpool.
func (o ScalingPlanPersonalScheduleOutput) RampUpAutoStartHosts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringPtrOutput { return v.RampUpAutoStartHosts }).(pulumi.StringPtrOutput)
}

// The time in minutes to wait before performing the desired session handling action when a user disconnects during the ramp up period.
func (o ScalingPlanPersonalScheduleOutput) RampUpMinutesToWaitOnDisconnect() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.IntPtrOutput { return v.RampUpMinutesToWaitOnDisconnect }).(pulumi.IntPtrOutput)
}

// The time in minutes to wait before performing the desired session handling action when a user logs off during the ramp up period.
func (o ScalingPlanPersonalScheduleOutput) RampUpMinutesToWaitOnLogoff() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.IntPtrOutput { return v.RampUpMinutesToWaitOnLogoff }).(pulumi.IntPtrOutput)
}

// Starting time for ramp up period.
func (o ScalingPlanPersonalScheduleOutput) RampUpStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) TimeResponsePtrOutput { return v.RampUpStartTime }).(TimeResponsePtrOutput)
}

// The desired configuration of Start VM On Connect for the hostpool during the ramp up phase. If this is disabled, session hosts must be turned on using rampUpAutoStartHosts or by turning them on manually.
func (o ScalingPlanPersonalScheduleOutput) RampUpStartVMOnConnect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringPtrOutput { return v.RampUpStartVMOnConnect }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ScalingPlanPersonalScheduleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ScalingPlanPersonalScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlanPersonalSchedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalingPlanPersonalScheduleOutput{})
}
