// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Managed instance's Start/Stop schedule.
type StartStopManagedInstanceSchedule struct {
	pulumi.CustomResourceState

	// The description of the schedule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Timestamp when the next action will be executed in the corresponding schedule time zone.
	NextExecutionTime pulumi.StringOutput `pulumi:"nextExecutionTime"`
	// Next action to be executed (Start or Stop)
	NextRunAction pulumi.StringOutput `pulumi:"nextRunAction"`
	// Schedule list.
	ScheduleList ScheduleItemResponseArrayOutput `pulumi:"scheduleList"`
	// System data of the scheduled resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The time zone of the schedule.
	TimeZoneId pulumi.StringPtrOutput `pulumi:"timeZoneId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStartStopManagedInstanceSchedule registers a new resource with the given unique name, arguments, and options.
func NewStartStopManagedInstanceSchedule(ctx *pulumi.Context,
	name string, args *StartStopManagedInstanceScheduleArgs, opts ...pulumi.ResourceOption) (*StartStopManagedInstanceSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScheduleList == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleList'")
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("")
	}
	if args.TimeZoneId == nil {
		args.TimeZoneId = pulumi.StringPtr("UTC")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:StartStopManagedInstanceSchedule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:StartStopManagedInstanceSchedule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:StartStopManagedInstanceSchedule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:StartStopManagedInstanceSchedule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:StartStopManagedInstanceSchedule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StartStopManagedInstanceSchedule
	err := ctx.RegisterResource("azure-native:sql/v20230801preview:StartStopManagedInstanceSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStartStopManagedInstanceSchedule gets an existing StartStopManagedInstanceSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStartStopManagedInstanceSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StartStopManagedInstanceScheduleState, opts ...pulumi.ResourceOption) (*StartStopManagedInstanceSchedule, error) {
	var resource StartStopManagedInstanceSchedule
	err := ctx.ReadResource("azure-native:sql/v20230801preview:StartStopManagedInstanceSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StartStopManagedInstanceSchedule resources.
type startStopManagedInstanceScheduleState struct {
}

type StartStopManagedInstanceScheduleState struct {
}

func (StartStopManagedInstanceScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*startStopManagedInstanceScheduleState)(nil)).Elem()
}

type startStopManagedInstanceScheduleArgs struct {
	// The description of the schedule.
	Description *string `pulumi:"description"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Schedule list.
	ScheduleList []ScheduleItem `pulumi:"scheduleList"`
	// Name of the managed instance Start/Stop schedule.
	StartStopScheduleName *string `pulumi:"startStopScheduleName"`
	// The time zone of the schedule.
	TimeZoneId *string `pulumi:"timeZoneId"`
}

// The set of arguments for constructing a StartStopManagedInstanceSchedule resource.
type StartStopManagedInstanceScheduleArgs struct {
	// The description of the schedule.
	Description pulumi.StringPtrInput
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Schedule list.
	ScheduleList ScheduleItemArrayInput
	// Name of the managed instance Start/Stop schedule.
	StartStopScheduleName pulumi.StringPtrInput
	// The time zone of the schedule.
	TimeZoneId pulumi.StringPtrInput
}

func (StartStopManagedInstanceScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*startStopManagedInstanceScheduleArgs)(nil)).Elem()
}

type StartStopManagedInstanceScheduleInput interface {
	pulumi.Input

	ToStartStopManagedInstanceScheduleOutput() StartStopManagedInstanceScheduleOutput
	ToStartStopManagedInstanceScheduleOutputWithContext(ctx context.Context) StartStopManagedInstanceScheduleOutput
}

func (*StartStopManagedInstanceSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**StartStopManagedInstanceSchedule)(nil)).Elem()
}

func (i *StartStopManagedInstanceSchedule) ToStartStopManagedInstanceScheduleOutput() StartStopManagedInstanceScheduleOutput {
	return i.ToStartStopManagedInstanceScheduleOutputWithContext(context.Background())
}

func (i *StartStopManagedInstanceSchedule) ToStartStopManagedInstanceScheduleOutputWithContext(ctx context.Context) StartStopManagedInstanceScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartStopManagedInstanceScheduleOutput)
}

type StartStopManagedInstanceScheduleOutput struct{ *pulumi.OutputState }

func (StartStopManagedInstanceScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StartStopManagedInstanceSchedule)(nil)).Elem()
}

func (o StartStopManagedInstanceScheduleOutput) ToStartStopManagedInstanceScheduleOutput() StartStopManagedInstanceScheduleOutput {
	return o
}

func (o StartStopManagedInstanceScheduleOutput) ToStartStopManagedInstanceScheduleOutputWithContext(ctx context.Context) StartStopManagedInstanceScheduleOutput {
	return o
}

// The description of the schedule.
func (o StartStopManagedInstanceScheduleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StartStopManagedInstanceSchedule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o StartStopManagedInstanceScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StartStopManagedInstanceSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Timestamp when the next action will be executed in the corresponding schedule time zone.
func (o StartStopManagedInstanceScheduleOutput) NextExecutionTime() pulumi.StringOutput {
	return o.ApplyT(func(v *StartStopManagedInstanceSchedule) pulumi.StringOutput { return v.NextExecutionTime }).(pulumi.StringOutput)
}

// Next action to be executed (Start or Stop)
func (o StartStopManagedInstanceScheduleOutput) NextRunAction() pulumi.StringOutput {
	return o.ApplyT(func(v *StartStopManagedInstanceSchedule) pulumi.StringOutput { return v.NextRunAction }).(pulumi.StringOutput)
}

// Schedule list.
func (o StartStopManagedInstanceScheduleOutput) ScheduleList() ScheduleItemResponseArrayOutput {
	return o.ApplyT(func(v *StartStopManagedInstanceSchedule) ScheduleItemResponseArrayOutput { return v.ScheduleList }).(ScheduleItemResponseArrayOutput)
}

// System data of the scheduled resource.
func (o StartStopManagedInstanceScheduleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StartStopManagedInstanceSchedule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The time zone of the schedule.
func (o StartStopManagedInstanceScheduleOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StartStopManagedInstanceSchedule) pulumi.StringPtrOutput { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o StartStopManagedInstanceScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StartStopManagedInstanceSchedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StartStopManagedInstanceScheduleOutput{})
}
