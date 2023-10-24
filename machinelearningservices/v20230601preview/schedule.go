// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
type Schedule struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	ScheduleProperties ScheduleResponseOutput `pulumi:"scheduleProperties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSchedule registers a new resource with the given unique name, arguments, and options.
func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScheduleProperties == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleProperties'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.ScheduleProperties = args.ScheduleProperties.ToScheduleTypeOutput().ApplyT(func(v ScheduleType) ScheduleType { return *v.Defaults() }).(ScheduleTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:Schedule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Schedule
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230601preview:Schedule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230601preview:Schedule", name, id, state, &resource, opts...)
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
	// Schedule name.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// [Required] Additional attributes of the entity.
	ScheduleProperties ScheduleType `pulumi:"scheduleProperties"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Schedule resource.
type ScheduleArgs struct {
	// Schedule name.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// [Required] Additional attributes of the entity.
	ScheduleProperties ScheduleTypeInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
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

func (i *Schedule) ToOutput(ctx context.Context) pulumix.Output[*Schedule] {
	return pulumix.Output[*Schedule]{
		OutputState: i.ToScheduleOutputWithContext(ctx).OutputState,
	}
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

func (o ScheduleOutput) ToOutput(ctx context.Context) pulumix.Output[*Schedule] {
	return pulumix.Output[*Schedule]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o ScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o ScheduleOutput) ScheduleProperties() ScheduleResponseOutput {
	return o.ApplyT(func(v *Schedule) ScheduleResponseOutput { return v.ScheduleProperties }).(ScheduleResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ScheduleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Schedule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduleOutput{})
}
