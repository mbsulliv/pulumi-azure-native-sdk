// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A schedule.
// Azure REST API version: 2018-09-15. Prior API version in Azure Native 1.x: 2018-09-15.
//
// Other available API versions: 2016-05-15.
type Schedule struct {
	pulumi.CustomResourceState

	// The creation date of the schedule.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// If the schedule will occur once each day of the week, specify the daily recurrence.
	DailyRecurrence DayDetailsResponsePtrOutput `pulumi:"dailyRecurrence"`
	// If the schedule will occur multiple times a day, specify the hourly recurrence.
	HourlyRecurrence HourDetailsResponsePtrOutput `pulumi:"hourlyRecurrence"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Notification settings.
	NotificationSettings NotificationSettingsResponsePtrOutput `pulumi:"notificationSettings"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The status of the schedule (i.e. Enabled, Disabled)
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource ID to which the schedule belongs
	TargetResourceId pulumi.StringPtrOutput `pulumi:"targetResourceId"`
	// The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).
	TaskType pulumi.StringPtrOutput `pulumi:"taskType"`
	// The time zone ID (e.g. China Standard Time, Greenland Standard Time, Pacific Standard time, etc.). The possible values for this property can be found in `IReadOnlyCollection<string> TimeZoneConverter.TZConvert.KnownWindowsTimeZoneIds` (https://github.com/mattjohnsonpint/TimeZoneConverter/blob/main/README.md)
	TimeZoneId pulumi.StringPtrOutput `pulumi:"timeZoneId"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringOutput `pulumi:"uniqueIdentifier"`
	// If the schedule will occur only some days of the week, specify the weekly recurrence.
	WeeklyRecurrence WeekDetailsResponsePtrOutput `pulumi:"weeklyRecurrence"`
}

// NewSchedule registers a new resource with the given unique name, arguments, and options.
func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.NotificationSettings != nil {
		args.NotificationSettings = args.NotificationSettings.ToNotificationSettingsPtrOutput().ApplyT(func(v *NotificationSettings) *NotificationSettings { return v.Defaults() }).(NotificationSettingsPtrOutput)
	}
	if args.Status == nil {
		args.Status = pulumi.StringPtr("Disabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:Schedule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Schedule
	err := ctx.RegisterResource("azure-native:devtestlab:Schedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchedule gets an existing Schedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleState, opts ...pulumi.ResourceOption) (*Schedule, error) {
	var resource Schedule
	err := ctx.ReadResource("azure-native:devtestlab:Schedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schedule resources.
type scheduleState struct {
}

type ScheduleState struct {
}

func (ScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleState)(nil)).Elem()
}

type scheduleArgs struct {
	// If the schedule will occur once each day of the week, specify the daily recurrence.
	DailyRecurrence *DayDetails `pulumi:"dailyRecurrence"`
	// If the schedule will occur multiple times a day, specify the hourly recurrence.
	HourlyRecurrence *HourDetails `pulumi:"hourlyRecurrence"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the schedule.
	Name *string `pulumi:"name"`
	// Notification settings.
	NotificationSettings *NotificationSettings `pulumi:"notificationSettings"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The status of the schedule (i.e. Enabled, Disabled)
	Status *string `pulumi:"status"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The resource ID to which the schedule belongs
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).
	TaskType *string `pulumi:"taskType"`
	// The time zone ID (e.g. China Standard Time, Greenland Standard Time, Pacific Standard time, etc.). The possible values for this property can be found in `IReadOnlyCollection<string> TimeZoneConverter.TZConvert.KnownWindowsTimeZoneIds` (https://github.com/mattjohnsonpint/TimeZoneConverter/blob/main/README.md)
	TimeZoneId *string `pulumi:"timeZoneId"`
	// If the schedule will occur only some days of the week, specify the weekly recurrence.
	WeeklyRecurrence *WeekDetails `pulumi:"weeklyRecurrence"`
}

// The set of arguments for constructing a Schedule resource.
type ScheduleArgs struct {
	// If the schedule will occur once each day of the week, specify the daily recurrence.
	DailyRecurrence DayDetailsPtrInput
	// If the schedule will occur multiple times a day, specify the hourly recurrence.
	HourlyRecurrence HourDetailsPtrInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the schedule.
	Name pulumi.StringPtrInput
	// Notification settings.
	NotificationSettings NotificationSettingsPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The status of the schedule (i.e. Enabled, Disabled)
	Status pulumi.StringPtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The resource ID to which the schedule belongs
	TargetResourceId pulumi.StringPtrInput
	// The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).
	TaskType pulumi.StringPtrInput
	// The time zone ID (e.g. China Standard Time, Greenland Standard Time, Pacific Standard time, etc.). The possible values for this property can be found in `IReadOnlyCollection<string> TimeZoneConverter.TZConvert.KnownWindowsTimeZoneIds` (https://github.com/mattjohnsonpint/TimeZoneConverter/blob/main/README.md)
	TimeZoneId pulumi.StringPtrInput
	// If the schedule will occur only some days of the week, specify the weekly recurrence.
	WeeklyRecurrence WeekDetailsPtrInput
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleArgs)(nil)).Elem()
}

type ScheduleInput interface {
	pulumi.Input

	ToScheduleOutput() ScheduleOutput
	ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput
}

func (*Schedule) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (i *Schedule) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i *Schedule) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}

type ScheduleOutput struct{ *pulumi.OutputState }

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

// The creation date of the schedule.
func (o ScheduleOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// If the schedule will occur once each day of the week, specify the daily recurrence.
func (o ScheduleOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Schedule) DayDetailsResponsePtrOutput { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

// If the schedule will occur multiple times a day, specify the hourly recurrence.
func (o ScheduleOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Schedule) HourDetailsResponsePtrOutput { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

// The location of the resource.
func (o ScheduleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o ScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Notification settings.
func (o ScheduleOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Schedule) NotificationSettingsResponsePtrOutput { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

// The provisioning status of the resource.
func (o ScheduleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The status of the schedule (i.e. Enabled, Disabled)
func (o ScheduleOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// The tags of the resource.
func (o ScheduleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource ID to which the schedule belongs
func (o ScheduleOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

// The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).
func (o ScheduleOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.TaskType }).(pulumi.StringPtrOutput)
}

// The time zone ID (e.g. China Standard Time, Greenland Standard Time, Pacific Standard time, etc.). The possible values for this property can be found in `IReadOnlyCollection<string> TimeZoneConverter.TZConvert.KnownWindowsTimeZoneIds` (https://github.com/mattjohnsonpint/TimeZoneConverter/blob/main/README.md)
func (o ScheduleOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o ScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o ScheduleOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

// If the schedule will occur only some days of the week, specify the weekly recurrence.
func (o ScheduleOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Schedule) WeekDetailsResponsePtrOutput { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduleOutput{})
}
