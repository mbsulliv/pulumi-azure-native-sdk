// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221012preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Schedule to execute a task.
type Schedule struct {
	pulumi.CustomResourceState

	// The frequency of this scheduled task.
	Frequency pulumi.StringOutput `pulumi:"frequency"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Indicates whether or not this scheduled task is enabled.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The target time to trigger the action. The format is HH:MM.
	Time pulumi.StringOutput `pulumi:"time"`
	// The IANA timezone id at which the schedule should execute.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSchedule registers a new resource with the given unique name, arguments, and options.
func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Time == nil {
		return nil, errors.New("invalid value for required argument 'Time'")
	}
	if args.TimeZone == nil {
		return nil, errors.New("invalid value for required argument 'TimeZone'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221111preview:Schedule"),
		},
	})
	opts = append(opts, aliases)
	var resource Schedule
	err := ctx.RegisterResource("azure-native:devcenter/v20221012preview:Schedule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:devcenter/v20221012preview:Schedule", name, id, state, &resource, opts...)
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
	// The frequency of this scheduled task.
	Frequency string `pulumi:"frequency"`
	// Name of the pool.
	PoolName string `pulumi:"poolName"`
	// The name of the project.
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the schedule that uniquely identifies it.
	ScheduleName *string `pulumi:"scheduleName"`
	// Indicates whether or not this scheduled task is enabled.
	State *string `pulumi:"state"`
	// The target time to trigger the action. The format is HH:MM.
	Time string `pulumi:"time"`
	// The IANA timezone id at which the schedule should execute.
	TimeZone string `pulumi:"timeZone"`
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int `pulumi:"top"`
	// Supported type this scheduled task represents.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Schedule resource.
type ScheduleArgs struct {
	// The frequency of this scheduled task.
	Frequency pulumi.StringInput
	// Name of the pool.
	PoolName pulumi.StringInput
	// The name of the project.
	ProjectName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the schedule that uniquely identifies it.
	ScheduleName pulumi.StringPtrInput
	// Indicates whether or not this scheduled task is enabled.
	State pulumi.StringPtrInput
	// The target time to trigger the action. The format is HH:MM.
	Time pulumi.StringInput
	// The IANA timezone id at which the schedule should execute.
	TimeZone pulumi.StringInput
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top pulumi.IntPtrInput
	// Supported type this scheduled task represents.
	Type pulumi.StringInput
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

// The frequency of this scheduled task.
func (o ScheduleOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Frequency }).(pulumi.StringOutput)
}

// The name of the resource
func (o ScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o ScheduleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Indicates whether or not this scheduled task is enabled.
func (o ScheduleOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ScheduleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Schedule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The target time to trigger the action. The format is HH:MM.
func (o ScheduleOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Time }).(pulumi.StringOutput)
}

// The IANA timezone id at which the schedule should execute.
func (o ScheduleOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduleOutput{})
}
