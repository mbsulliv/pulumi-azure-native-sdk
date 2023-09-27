// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Specifies information about the capacity reservation group that the capacity reservations should be assigned to. Currently, a capacity reservation can only be added to a capacity reservation group at creation time. An existing capacity reservation cannot be added or moved to another capacity reservation group.
type CapacityReservationGroup struct {
	pulumi.CustomResourceState

	// A list of all capacity reservation resource ids that belong to capacity reservation group.
	CapacityReservations SubResourceReadOnlyResponseArrayOutput `pulumi:"capacityReservations"`
	// The capacity reservation group instance view which has the list of instance views for all the capacity reservations that belong to the capacity reservation group.
	InstanceView CapacityReservationGroupInstanceViewResponseOutput `pulumi:"instanceView"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of references to all virtual machines associated to the capacity reservation group.
	VirtualMachinesAssociated SubResourceReadOnlyResponseArrayOutput `pulumi:"virtualMachinesAssociated"`
	// Availability Zones to use for this capacity reservation group. The zones can be assigned only during creation. If not provided, the group supports only regional resources in the region. If provided, enforces each capacity reservation in the group to be in one of the zones.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewCapacityReservationGroup registers a new resource with the given unique name, arguments, and options.
func NewCapacityReservationGroup(ctx *pulumi.Context,
	name string, args *CapacityReservationGroupArgs, opts ...pulumi.ResourceOption) (*CapacityReservationGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:CapacityReservationGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:CapacityReservationGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:CapacityReservationGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:CapacityReservationGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:CapacityReservationGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:CapacityReservationGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:CapacityReservationGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230701:CapacityReservationGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CapacityReservationGroup
	err := ctx.RegisterResource("azure-native:compute/v20230301:CapacityReservationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapacityReservationGroup gets an existing CapacityReservationGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapacityReservationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityReservationGroupState, opts ...pulumi.ResourceOption) (*CapacityReservationGroup, error) {
	var resource CapacityReservationGroup
	err := ctx.ReadResource("azure-native:compute/v20230301:CapacityReservationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CapacityReservationGroup resources.
type capacityReservationGroupState struct {
}

type CapacityReservationGroupState struct {
}

func (CapacityReservationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationGroupState)(nil)).Elem()
}

type capacityReservationGroupArgs struct {
	// The name of the capacity reservation group.
	CapacityReservationGroupName *string `pulumi:"capacityReservationGroupName"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Availability Zones to use for this capacity reservation group. The zones can be assigned only during creation. If not provided, the group supports only regional resources in the region. If provided, enforces each capacity reservation in the group to be in one of the zones.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a CapacityReservationGroup resource.
type CapacityReservationGroupArgs struct {
	// The name of the capacity reservation group.
	CapacityReservationGroupName pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Availability Zones to use for this capacity reservation group. The zones can be assigned only during creation. If not provided, the group supports only regional resources in the region. If provided, enforces each capacity reservation in the group to be in one of the zones.
	Zones pulumi.StringArrayInput
}

func (CapacityReservationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationGroupArgs)(nil)).Elem()
}

type CapacityReservationGroupInput interface {
	pulumi.Input

	ToCapacityReservationGroupOutput() CapacityReservationGroupOutput
	ToCapacityReservationGroupOutputWithContext(ctx context.Context) CapacityReservationGroupOutput
}

func (*CapacityReservationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservationGroup)(nil)).Elem()
}

func (i *CapacityReservationGroup) ToCapacityReservationGroupOutput() CapacityReservationGroupOutput {
	return i.ToCapacityReservationGroupOutputWithContext(context.Background())
}

func (i *CapacityReservationGroup) ToCapacityReservationGroupOutputWithContext(ctx context.Context) CapacityReservationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationGroupOutput)
}

func (i *CapacityReservationGroup) ToOutput(ctx context.Context) pulumix.Output[*CapacityReservationGroup] {
	return pulumix.Output[*CapacityReservationGroup]{
		OutputState: i.ToCapacityReservationGroupOutputWithContext(ctx).OutputState,
	}
}

type CapacityReservationGroupOutput struct{ *pulumi.OutputState }

func (CapacityReservationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservationGroup)(nil)).Elem()
}

func (o CapacityReservationGroupOutput) ToCapacityReservationGroupOutput() CapacityReservationGroupOutput {
	return o
}

func (o CapacityReservationGroupOutput) ToCapacityReservationGroupOutputWithContext(ctx context.Context) CapacityReservationGroupOutput {
	return o
}

func (o CapacityReservationGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*CapacityReservationGroup] {
	return pulumix.Output[*CapacityReservationGroup]{
		OutputState: o.OutputState,
	}
}

// A list of all capacity reservation resource ids that belong to capacity reservation group.
func (o CapacityReservationGroupOutput) CapacityReservations() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) SubResourceReadOnlyResponseArrayOutput {
		return v.CapacityReservations
	}).(SubResourceReadOnlyResponseArrayOutput)
}

// The capacity reservation group instance view which has the list of instance views for all the capacity reservations that belong to the capacity reservation group.
func (o CapacityReservationGroupOutput) InstanceView() CapacityReservationGroupInstanceViewResponseOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) CapacityReservationGroupInstanceViewResponseOutput {
		return v.InstanceView
	}).(CapacityReservationGroupInstanceViewResponseOutput)
}

// Resource location
func (o CapacityReservationGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o CapacityReservationGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource tags
func (o CapacityReservationGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o CapacityReservationGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A list of references to all virtual machines associated to the capacity reservation group.
func (o CapacityReservationGroupOutput) VirtualMachinesAssociated() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) SubResourceReadOnlyResponseArrayOutput {
		return v.VirtualMachinesAssociated
	}).(SubResourceReadOnlyResponseArrayOutput)
}

// Availability Zones to use for this capacity reservation group. The zones can be assigned only during creation. If not provided, the group supports only regional resources in the region. If provided, enforces each capacity reservation in the group to be in one of the zones.
func (o CapacityReservationGroupOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(CapacityReservationGroupOutput{})
}
