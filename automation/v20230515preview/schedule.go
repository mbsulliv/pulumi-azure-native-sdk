// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230515preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the schedule.
type Schedule struct {
	pulumi.CustomResourceState

	// Gets or sets the advanced schedule.
	AdvancedSchedule AdvancedScheduleResponsePtrOutput `pulumi:"advancedSchedule"`
	// Gets or sets the creation time.
	CreationTime pulumi.StringPtrOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets or sets the end time of the schedule.
	ExpiryTime pulumi.StringPtrOutput `pulumi:"expiryTime"`
	// Gets or sets the expiry time's offset in minutes.
	ExpiryTimeOffsetMinutes pulumi.Float64PtrOutput `pulumi:"expiryTimeOffsetMinutes"`
	// Gets or sets the frequency of the schedule.
	Frequency pulumi.StringPtrOutput `pulumi:"frequency"`
	// Gets or sets the interval of the schedule.
	Interval pulumi.AnyOutput `pulumi:"interval"`
	// Gets or sets a value indicating whether this schedule is enabled.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// Gets or sets the last modified time.
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the next run time of the schedule.
	NextRun pulumi.StringPtrOutput `pulumi:"nextRun"`
	// Gets or sets the next run time's offset in minutes.
	NextRunOffsetMinutes pulumi.Float64PtrOutput `pulumi:"nextRunOffsetMinutes"`
	// Gets or sets the start time of the schedule.
	StartTime pulumi.StringPtrOutput `pulumi:"startTime"`
	// Gets the start time's offset in minutes.
	StartTimeOffsetMinutes pulumi.Float64Output `pulumi:"startTimeOffsetMinutes"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Gets or sets the time zone of the schedule.
	TimeZone pulumi.StringPtrOutput `pulumi:"timeZone"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSchedule registers a new resource with the given unique name, arguments, and options.
func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StartTime == nil {
		return nil, errors.New("invalid value for required argument 'StartTime'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20220808:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20231101:Schedule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Schedule
	err := ctx.RegisterResource("azure-native:automation/v20230515preview:Schedule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:automation/v20230515preview:Schedule", name, id, state, &resource, opts...)
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
	// Gets or sets the AdvancedSchedule.
	AdvancedSchedule *AdvancedSchedule `pulumi:"advancedSchedule"`
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Gets or sets the description of the schedule.
	Description *string `pulumi:"description"`
	// Gets or sets the end time of the schedule.
	ExpiryTime *string `pulumi:"expiryTime"`
	// Gets or sets the frequency of the schedule.
	Frequency string `pulumi:"frequency"`
	// Gets or sets the interval of the schedule.
	Interval interface{} `pulumi:"interval"`
	// Gets or sets the name of the Schedule.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The schedule name.
	ScheduleName *string `pulumi:"scheduleName"`
	// Gets or sets the start time of the schedule.
	StartTime string `pulumi:"startTime"`
	// Gets or sets the time zone of the schedule.
	TimeZone *string `pulumi:"timeZone"`
}

// The set of arguments for constructing a Schedule resource.
type ScheduleArgs struct {
	// Gets or sets the AdvancedSchedule.
	AdvancedSchedule AdvancedSchedulePtrInput
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Gets or sets the description of the schedule.
	Description pulumi.StringPtrInput
	// Gets or sets the end time of the schedule.
	ExpiryTime pulumi.StringPtrInput
	// Gets or sets the frequency of the schedule.
	Frequency pulumi.StringInput
	// Gets or sets the interval of the schedule.
	Interval pulumi.Input
	// Gets or sets the name of the Schedule.
	Name pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// The schedule name.
	ScheduleName pulumi.StringPtrInput
	// Gets or sets the start time of the schedule.
	StartTime pulumi.StringInput
	// Gets or sets the time zone of the schedule.
	TimeZone pulumi.StringPtrInput
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

// Gets or sets the advanced schedule.
func (o ScheduleOutput) AdvancedSchedule() AdvancedScheduleResponsePtrOutput {
	return o.ApplyT(func(v *Schedule) AdvancedScheduleResponsePtrOutput { return v.AdvancedSchedule }).(AdvancedScheduleResponsePtrOutput)
}

// Gets or sets the creation time.
func (o ScheduleOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the description.
func (o ScheduleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets or sets the end time of the schedule.
func (o ScheduleOutput) ExpiryTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.ExpiryTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the expiry time's offset in minutes.
func (o ScheduleOutput) ExpiryTimeOffsetMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.Float64PtrOutput { return v.ExpiryTimeOffsetMinutes }).(pulumi.Float64PtrOutput)
}

// Gets or sets the frequency of the schedule.
func (o ScheduleOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.Frequency }).(pulumi.StringPtrOutput)
}

// Gets or sets the interval of the schedule.
func (o ScheduleOutput) Interval() pulumi.AnyOutput {
	return o.ApplyT(func(v *Schedule) pulumi.AnyOutput { return v.Interval }).(pulumi.AnyOutput)
}

// Gets or sets a value indicating whether this schedule is enabled.
func (o ScheduleOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time.
func (o ScheduleOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the next run time of the schedule.
func (o ScheduleOutput) NextRun() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.NextRun }).(pulumi.StringPtrOutput)
}

// Gets or sets the next run time's offset in minutes.
func (o ScheduleOutput) NextRunOffsetMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.Float64PtrOutput { return v.NextRunOffsetMinutes }).(pulumi.Float64PtrOutput)
}

// Gets or sets the start time of the schedule.
func (o ScheduleOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.StartTime }).(pulumi.StringPtrOutput)
}

// Gets the start time's offset in minutes.
func (o ScheduleOutput) StartTimeOffsetMinutes() pulumi.Float64Output {
	return o.ApplyT(func(v *Schedule) pulumi.Float64Output { return v.StartTimeOffsetMinutes }).(pulumi.Float64Output)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ScheduleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Schedule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets the time zone of the schedule.
func (o ScheduleOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.TimeZone }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduleOutput{})
}
