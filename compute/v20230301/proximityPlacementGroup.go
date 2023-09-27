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

// Specifies information about the proximity placement group.
type ProximityPlacementGroup struct {
	pulumi.CustomResourceState

	// A list of references to all availability sets in the proximity placement group.
	AvailabilitySets SubResourceWithColocationStatusResponseArrayOutput `pulumi:"availabilitySets"`
	// Describes colocation status of the Proximity Placement Group.
	ColocationStatus InstanceViewStatusResponsePtrOutput `pulumi:"colocationStatus"`
	// Specifies the user intent of the proximity placement group.
	Intent ProximityPlacementGroupPropertiesResponseIntentPtrOutput `pulumi:"intent"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the type of the proximity placement group. Possible values are: **Standard** : Co-locate resources within an Azure region or Availability Zone. **Ultra** : For future use.
	ProximityPlacementGroupType pulumi.StringPtrOutput `pulumi:"proximityPlacementGroupType"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of references to all virtual machine scale sets in the proximity placement group.
	VirtualMachineScaleSets SubResourceWithColocationStatusResponseArrayOutput `pulumi:"virtualMachineScaleSets"`
	// A list of references to all virtual machines in the proximity placement group.
	VirtualMachines SubResourceWithColocationStatusResponseArrayOutput `pulumi:"virtualMachines"`
	// Specifies the Availability Zone where virtual machine, virtual machine scale set or availability set associated with the  proximity placement group can be created.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewProximityPlacementGroup registers a new resource with the given unique name, arguments, and options.
func NewProximityPlacementGroup(ctx *pulumi.Context,
	name string, args *ProximityPlacementGroupArgs, opts ...pulumi.ResourceOption) (*ProximityPlacementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230701:ProximityPlacementGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProximityPlacementGroup
	err := ctx.RegisterResource("azure-native:compute/v20230301:ProximityPlacementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProximityPlacementGroup gets an existing ProximityPlacementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProximityPlacementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProximityPlacementGroupState, opts ...pulumi.ResourceOption) (*ProximityPlacementGroup, error) {
	var resource ProximityPlacementGroup
	err := ctx.ReadResource("azure-native:compute/v20230301:ProximityPlacementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProximityPlacementGroup resources.
type proximityPlacementGroupState struct {
}

type ProximityPlacementGroupState struct {
}

func (ProximityPlacementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*proximityPlacementGroupState)(nil)).Elem()
}

type proximityPlacementGroupArgs struct {
	// Describes colocation status of the Proximity Placement Group.
	ColocationStatus *InstanceViewStatus `pulumi:"colocationStatus"`
	// Specifies the user intent of the proximity placement group.
	Intent *ProximityPlacementGroupPropertiesIntent `pulumi:"intent"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the proximity placement group.
	ProximityPlacementGroupName *string `pulumi:"proximityPlacementGroupName"`
	// Specifies the type of the proximity placement group. Possible values are: **Standard** : Co-locate resources within an Azure region or Availability Zone. **Ultra** : For future use.
	ProximityPlacementGroupType *string `pulumi:"proximityPlacementGroupType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Availability Zone where virtual machine, virtual machine scale set or availability set associated with the  proximity placement group can be created.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a ProximityPlacementGroup resource.
type ProximityPlacementGroupArgs struct {
	// Describes colocation status of the Proximity Placement Group.
	ColocationStatus InstanceViewStatusPtrInput
	// Specifies the user intent of the proximity placement group.
	Intent ProximityPlacementGroupPropertiesIntentPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the proximity placement group.
	ProximityPlacementGroupName pulumi.StringPtrInput
	// Specifies the type of the proximity placement group. Possible values are: **Standard** : Co-locate resources within an Azure region or Availability Zone. **Ultra** : For future use.
	ProximityPlacementGroupType pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Specifies the Availability Zone where virtual machine, virtual machine scale set or availability set associated with the  proximity placement group can be created.
	Zones pulumi.StringArrayInput
}

func (ProximityPlacementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*proximityPlacementGroupArgs)(nil)).Elem()
}

type ProximityPlacementGroupInput interface {
	pulumi.Input

	ToProximityPlacementGroupOutput() ProximityPlacementGroupOutput
	ToProximityPlacementGroupOutputWithContext(ctx context.Context) ProximityPlacementGroupOutput
}

func (*ProximityPlacementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ProximityPlacementGroup)(nil)).Elem()
}

func (i *ProximityPlacementGroup) ToProximityPlacementGroupOutput() ProximityPlacementGroupOutput {
	return i.ToProximityPlacementGroupOutputWithContext(context.Background())
}

func (i *ProximityPlacementGroup) ToProximityPlacementGroupOutputWithContext(ctx context.Context) ProximityPlacementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProximityPlacementGroupOutput)
}

func (i *ProximityPlacementGroup) ToOutput(ctx context.Context) pulumix.Output[*ProximityPlacementGroup] {
	return pulumix.Output[*ProximityPlacementGroup]{
		OutputState: i.ToProximityPlacementGroupOutputWithContext(ctx).OutputState,
	}
}

type ProximityPlacementGroupOutput struct{ *pulumi.OutputState }

func (ProximityPlacementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProximityPlacementGroup)(nil)).Elem()
}

func (o ProximityPlacementGroupOutput) ToProximityPlacementGroupOutput() ProximityPlacementGroupOutput {
	return o
}

func (o ProximityPlacementGroupOutput) ToProximityPlacementGroupOutputWithContext(ctx context.Context) ProximityPlacementGroupOutput {
	return o
}

func (o ProximityPlacementGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*ProximityPlacementGroup] {
	return pulumix.Output[*ProximityPlacementGroup]{
		OutputState: o.OutputState,
	}
}

// A list of references to all availability sets in the proximity placement group.
func (o ProximityPlacementGroupOutput) AvailabilitySets() SubResourceWithColocationStatusResponseArrayOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) SubResourceWithColocationStatusResponseArrayOutput {
		return v.AvailabilitySets
	}).(SubResourceWithColocationStatusResponseArrayOutput)
}

// Describes colocation status of the Proximity Placement Group.
func (o ProximityPlacementGroupOutput) ColocationStatus() InstanceViewStatusResponsePtrOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) InstanceViewStatusResponsePtrOutput { return v.ColocationStatus }).(InstanceViewStatusResponsePtrOutput)
}

// Specifies the user intent of the proximity placement group.
func (o ProximityPlacementGroupOutput) Intent() ProximityPlacementGroupPropertiesResponseIntentPtrOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) ProximityPlacementGroupPropertiesResponseIntentPtrOutput {
		return v.Intent
	}).(ProximityPlacementGroupPropertiesResponseIntentPtrOutput)
}

// Resource location
func (o ProximityPlacementGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o ProximityPlacementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the type of the proximity placement group. Possible values are: **Standard** : Co-locate resources within an Azure region or Availability Zone. **Ultra** : For future use.
func (o ProximityPlacementGroupOutput) ProximityPlacementGroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) pulumi.StringPtrOutput { return v.ProximityPlacementGroupType }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o ProximityPlacementGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o ProximityPlacementGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A list of references to all virtual machine scale sets in the proximity placement group.
func (o ProximityPlacementGroupOutput) VirtualMachineScaleSets() SubResourceWithColocationStatusResponseArrayOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) SubResourceWithColocationStatusResponseArrayOutput {
		return v.VirtualMachineScaleSets
	}).(SubResourceWithColocationStatusResponseArrayOutput)
}

// A list of references to all virtual machines in the proximity placement group.
func (o ProximityPlacementGroupOutput) VirtualMachines() SubResourceWithColocationStatusResponseArrayOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) SubResourceWithColocationStatusResponseArrayOutput {
		return v.VirtualMachines
	}).(SubResourceWithColocationStatusResponseArrayOutput)
}

// Specifies the Availability Zone where virtual machine, virtual machine scale set or availability set associated with the  proximity placement group can be created.
func (o ProximityPlacementGroupOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ProximityPlacementGroupOutput{})
}
