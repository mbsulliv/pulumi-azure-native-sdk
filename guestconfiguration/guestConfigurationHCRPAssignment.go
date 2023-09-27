// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guestconfiguration

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Guest configuration assignment is an association between a machine and guest configuration.
// Azure REST API version: 2022-01-25. Prior API version in Azure Native 1.x: 2020-06-25
type GuestConfigurationHCRPAssignment struct {
	pulumi.CustomResourceState

	// Region where the VM is located.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the guest configuration assignment.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of the Guest configuration assignment.
	Properties GuestConfigurationAssignmentPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGuestConfigurationHCRPAssignment registers a new resource with the given unique name, arguments, and options.
func NewGuestConfigurationHCRPAssignment(ctx *pulumi.Context,
	name string, args *GuestConfigurationHCRPAssignmentArgs, opts ...pulumi.ResourceOption) (*GuestConfigurationHCRPAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MachineName == nil {
		return nil, errors.New("invalid value for required argument 'MachineName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20181120:GuestConfigurationHCRPAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20200625:GuestConfigurationHCRPAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20210125:GuestConfigurationHCRPAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20220125:GuestConfigurationHCRPAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GuestConfigurationHCRPAssignment
	err := ctx.RegisterResource("azure-native:guestconfiguration:GuestConfigurationHCRPAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGuestConfigurationHCRPAssignment gets an existing GuestConfigurationHCRPAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGuestConfigurationHCRPAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestConfigurationHCRPAssignmentState, opts ...pulumi.ResourceOption) (*GuestConfigurationHCRPAssignment, error) {
	var resource GuestConfigurationHCRPAssignment
	err := ctx.ReadResource("azure-native:guestconfiguration:GuestConfigurationHCRPAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GuestConfigurationHCRPAssignment resources.
type guestConfigurationHCRPAssignmentState struct {
}

type GuestConfigurationHCRPAssignmentState struct {
}

func (GuestConfigurationHCRPAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationHCRPAssignmentState)(nil)).Elem()
}

type guestConfigurationHCRPAssignmentArgs struct {
	// Name of the guest configuration assignment.
	GuestConfigurationAssignmentName *string `pulumi:"guestConfigurationAssignmentName"`
	// Region where the VM is located.
	Location *string `pulumi:"location"`
	// The name of the ARC machine.
	MachineName string `pulumi:"machineName"`
	// Name of the guest configuration assignment.
	Name *string `pulumi:"name"`
	// Properties of the Guest configuration assignment.
	Properties *GuestConfigurationAssignmentProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a GuestConfigurationHCRPAssignment resource.
type GuestConfigurationHCRPAssignmentArgs struct {
	// Name of the guest configuration assignment.
	GuestConfigurationAssignmentName pulumi.StringPtrInput
	// Region where the VM is located.
	Location pulumi.StringPtrInput
	// The name of the ARC machine.
	MachineName pulumi.StringInput
	// Name of the guest configuration assignment.
	Name pulumi.StringPtrInput
	// Properties of the Guest configuration assignment.
	Properties GuestConfigurationAssignmentPropertiesPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
}

func (GuestConfigurationHCRPAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationHCRPAssignmentArgs)(nil)).Elem()
}

type GuestConfigurationHCRPAssignmentInput interface {
	pulumi.Input

	ToGuestConfigurationHCRPAssignmentOutput() GuestConfigurationHCRPAssignmentOutput
	ToGuestConfigurationHCRPAssignmentOutputWithContext(ctx context.Context) GuestConfigurationHCRPAssignmentOutput
}

func (*GuestConfigurationHCRPAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationHCRPAssignment)(nil)).Elem()
}

func (i *GuestConfigurationHCRPAssignment) ToGuestConfigurationHCRPAssignmentOutput() GuestConfigurationHCRPAssignmentOutput {
	return i.ToGuestConfigurationHCRPAssignmentOutputWithContext(context.Background())
}

func (i *GuestConfigurationHCRPAssignment) ToGuestConfigurationHCRPAssignmentOutputWithContext(ctx context.Context) GuestConfigurationHCRPAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationHCRPAssignmentOutput)
}

func (i *GuestConfigurationHCRPAssignment) ToOutput(ctx context.Context) pulumix.Output[*GuestConfigurationHCRPAssignment] {
	return pulumix.Output[*GuestConfigurationHCRPAssignment]{
		OutputState: i.ToGuestConfigurationHCRPAssignmentOutputWithContext(ctx).OutputState,
	}
}

type GuestConfigurationHCRPAssignmentOutput struct{ *pulumi.OutputState }

func (GuestConfigurationHCRPAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationHCRPAssignment)(nil)).Elem()
}

func (o GuestConfigurationHCRPAssignmentOutput) ToGuestConfigurationHCRPAssignmentOutput() GuestConfigurationHCRPAssignmentOutput {
	return o
}

func (o GuestConfigurationHCRPAssignmentOutput) ToGuestConfigurationHCRPAssignmentOutputWithContext(ctx context.Context) GuestConfigurationHCRPAssignmentOutput {
	return o
}

func (o GuestConfigurationHCRPAssignmentOutput) ToOutput(ctx context.Context) pulumix.Output[*GuestConfigurationHCRPAssignment] {
	return pulumix.Output[*GuestConfigurationHCRPAssignment]{
		OutputState: o.OutputState,
	}
}

// Region where the VM is located.
func (o GuestConfigurationHCRPAssignmentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationHCRPAssignment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the guest configuration assignment.
func (o GuestConfigurationHCRPAssignmentOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationHCRPAssignment) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Properties of the Guest configuration assignment.
func (o GuestConfigurationHCRPAssignmentOutput) Properties() GuestConfigurationAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *GuestConfigurationHCRPAssignment) GuestConfigurationAssignmentPropertiesResponseOutput {
		return v.Properties
	}).(GuestConfigurationAssignmentPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o GuestConfigurationHCRPAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GuestConfigurationHCRPAssignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o GuestConfigurationHCRPAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestConfigurationHCRPAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GuestConfigurationHCRPAssignmentOutput{})
}
