// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Response to put/get patch schedules for Redis cache.
type PatchSchedule struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of patch schedules for a Redis cache.
	ScheduleEntries ScheduleEntryResponseArrayOutput `pulumi:"scheduleEntries"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPatchSchedule registers a new resource with the given unique name, arguments, and options.
func NewPatchSchedule(ctx *pulumi.Context,
	name string, args *PatchScheduleArgs, opts ...pulumi.ResourceOption) (*PatchSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScheduleEntries == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleEntries'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cache:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20171001:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20180301:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20190701:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20200601:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20201201:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210601:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220501:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220601:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20230401:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20230801:PatchSchedule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PatchSchedule
	err := ctx.RegisterResource("azure-native:cache/v20230501preview:PatchSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPatchSchedule gets an existing PatchSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPatchSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PatchScheduleState, opts ...pulumi.ResourceOption) (*PatchSchedule, error) {
	var resource PatchSchedule
	err := ctx.ReadResource("azure-native:cache/v20230501preview:PatchSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PatchSchedule resources.
type patchScheduleState struct {
}

type PatchScheduleState struct {
}

func (PatchScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*patchScheduleState)(nil)).Elem()
}

type patchScheduleArgs struct {
	// Default string modeled as parameter for auto generation to work correctly.
	Default *string `pulumi:"default"`
	// The name of the Redis cache.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntry `pulumi:"scheduleEntries"`
}

// The set of arguments for constructing a PatchSchedule resource.
type PatchScheduleArgs struct {
	// Default string modeled as parameter for auto generation to work correctly.
	Default pulumi.StringPtrInput
	// The name of the Redis cache.
	Name pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// List of patch schedules for a Redis cache.
	ScheduleEntries ScheduleEntryArrayInput
}

func (PatchScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*patchScheduleArgs)(nil)).Elem()
}

type PatchScheduleInput interface {
	pulumi.Input

	ToPatchScheduleOutput() PatchScheduleOutput
	ToPatchScheduleOutputWithContext(ctx context.Context) PatchScheduleOutput
}

func (*PatchSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchSchedule)(nil)).Elem()
}

func (i *PatchSchedule) ToPatchScheduleOutput() PatchScheduleOutput {
	return i.ToPatchScheduleOutputWithContext(context.Background())
}

func (i *PatchSchedule) ToPatchScheduleOutputWithContext(ctx context.Context) PatchScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchScheduleOutput)
}

type PatchScheduleOutput struct{ *pulumi.OutputState }

func (PatchScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchSchedule)(nil)).Elem()
}

func (o PatchScheduleOutput) ToPatchScheduleOutput() PatchScheduleOutput {
	return o
}

func (o PatchScheduleOutput) ToPatchScheduleOutputWithContext(ctx context.Context) PatchScheduleOutput {
	return o
}

// The geo-location where the resource lives
func (o PatchScheduleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchSchedule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o PatchScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of patch schedules for a Redis cache.
func (o PatchScheduleOutput) ScheduleEntries() ScheduleEntryResponseArrayOutput {
	return o.ApplyT(func(v *PatchSchedule) ScheduleEntryResponseArrayOutput { return v.ScheduleEntries }).(ScheduleEntryResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PatchScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchSchedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PatchScheduleOutput{})
}
