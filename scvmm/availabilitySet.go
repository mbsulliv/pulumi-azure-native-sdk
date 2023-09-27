// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scvmm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The AvailabilitySets resource definition.
// Azure REST API version: 2022-05-21-preview. Prior API version in Azure Native 1.x: 2020-06-05-preview
type AvailabilitySet struct {
	pulumi.CustomResourceState

	// Name of the availability set.
	AvailabilitySetName pulumi.StringPtrOutput `pulumi:"availabilitySetName"`
	// The extended location.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// Gets or sets the location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system data.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId pulumi.StringPtrOutput `pulumi:"vmmServerId"`
}

// NewAvailabilitySet registers a new resource with the given unique name, arguments, and options.
func NewAvailabilitySet(ctx *pulumi.Context,
	name string, args *AvailabilitySetArgs, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scvmm/v20200605preview:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20220521preview:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20230401preview:AvailabilitySet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AvailabilitySet
	err := ctx.RegisterResource("azure-native:scvmm:AvailabilitySet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAvailabilitySet gets an existing AvailabilitySet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAvailabilitySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AvailabilitySetState, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	var resource AvailabilitySet
	err := ctx.ReadResource("azure-native:scvmm:AvailabilitySet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AvailabilitySet resources.
type availabilitySetState struct {
}

type AvailabilitySetState struct {
}

func (AvailabilitySetState) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilitySetState)(nil)).Elem()
}

type availabilitySetArgs struct {
	// Name of the availability set.
	AvailabilitySetName *string `pulumi:"availabilitySetName"`
	// The extended location.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId *string `pulumi:"vmmServerId"`
}

// The set of arguments for constructing a AvailabilitySet resource.
type AvailabilitySetArgs struct {
	// Name of the availability set.
	AvailabilitySetName pulumi.StringPtrInput
	// The extended location.
	ExtendedLocation ExtendedLocationPtrInput
	// Gets or sets the location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId pulumi.StringPtrInput
}

func (AvailabilitySetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilitySetArgs)(nil)).Elem()
}

type AvailabilitySetInput interface {
	pulumi.Input

	ToAvailabilitySetOutput() AvailabilitySetOutput
	ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput
}

func (*AvailabilitySet) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilitySet)(nil)).Elem()
}

func (i *AvailabilitySet) ToAvailabilitySetOutput() AvailabilitySetOutput {
	return i.ToAvailabilitySetOutputWithContext(context.Background())
}

func (i *AvailabilitySet) ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilitySetOutput)
}

func (i *AvailabilitySet) ToOutput(ctx context.Context) pulumix.Output[*AvailabilitySet] {
	return pulumix.Output[*AvailabilitySet]{
		OutputState: i.ToAvailabilitySetOutputWithContext(ctx).OutputState,
	}
}

type AvailabilitySetOutput struct{ *pulumi.OutputState }

func (AvailabilitySetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilitySet)(nil)).Elem()
}

func (o AvailabilitySetOutput) ToAvailabilitySetOutput() AvailabilitySetOutput {
	return o
}

func (o AvailabilitySetOutput) ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput {
	return o
}

func (o AvailabilitySetOutput) ToOutput(ctx context.Context) pulumix.Output[*AvailabilitySet] {
	return pulumix.Output[*AvailabilitySet]{
		OutputState: o.OutputState,
	}
}

// Name of the availability set.
func (o AvailabilitySetOutput) AvailabilitySetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringPtrOutput { return v.AvailabilitySetName }).(pulumi.StringPtrOutput)
}

// The extended location.
func (o AvailabilitySetOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Gets or sets the location.
func (o AvailabilitySetOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o AvailabilitySetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state.
func (o AvailabilitySetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system data.
func (o AvailabilitySetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AvailabilitySet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o AvailabilitySetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource Type
func (o AvailabilitySetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// ARM Id of the vmmServer resource in which this resource resides.
func (o AvailabilitySetOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringPtrOutput { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AvailabilitySetOutput{})
}
