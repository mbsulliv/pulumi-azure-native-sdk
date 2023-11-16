// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220125

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Guest configuration assignment is an association between a machine and guest configuration.
type GuestConfigurationConnectedVMwarevSphereAssignment struct {
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

// NewGuestConfigurationConnectedVMwarevSphereAssignment registers a new resource with the given unique name, arguments, and options.
func NewGuestConfigurationConnectedVMwarevSphereAssignment(ctx *pulumi.Context,
	name string, args *GuestConfigurationConnectedVMwarevSphereAssignmentArgs, opts ...pulumi.ResourceOption) (*GuestConfigurationConnectedVMwarevSphereAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmName == nil {
		return nil, errors.New("invalid value for required argument 'VmName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:guestconfiguration:GuestConfigurationConnectedVMwarevSphereAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20200625:GuestConfigurationConnectedVMwarevSphereAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GuestConfigurationConnectedVMwarevSphereAssignment
	err := ctx.RegisterResource("azure-native:guestconfiguration/v20220125:GuestConfigurationConnectedVMwarevSphereAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGuestConfigurationConnectedVMwarevSphereAssignment gets an existing GuestConfigurationConnectedVMwarevSphereAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGuestConfigurationConnectedVMwarevSphereAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestConfigurationConnectedVMwarevSphereAssignmentState, opts ...pulumi.ResourceOption) (*GuestConfigurationConnectedVMwarevSphereAssignment, error) {
	var resource GuestConfigurationConnectedVMwarevSphereAssignment
	err := ctx.ReadResource("azure-native:guestconfiguration/v20220125:GuestConfigurationConnectedVMwarevSphereAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GuestConfigurationConnectedVMwarevSphereAssignment resources.
type guestConfigurationConnectedVMwarevSphereAssignmentState struct {
}

type GuestConfigurationConnectedVMwarevSphereAssignmentState struct {
}

func (GuestConfigurationConnectedVMwarevSphereAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationConnectedVMwarevSphereAssignmentState)(nil)).Elem()
}

type guestConfigurationConnectedVMwarevSphereAssignmentArgs struct {
	// Name of the guest configuration assignment.
	GuestConfigurationAssignmentName *string `pulumi:"guestConfigurationAssignmentName"`
	// Region where the VM is located.
	Location *string `pulumi:"location"`
	// Name of the guest configuration assignment.
	Name *string `pulumi:"name"`
	// Properties of the Guest configuration assignment.
	Properties *GuestConfigurationAssignmentProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual machine.
	VmName string `pulumi:"vmName"`
}

// The set of arguments for constructing a GuestConfigurationConnectedVMwarevSphereAssignment resource.
type GuestConfigurationConnectedVMwarevSphereAssignmentArgs struct {
	// Name of the guest configuration assignment.
	GuestConfigurationAssignmentName pulumi.StringPtrInput
	// Region where the VM is located.
	Location pulumi.StringPtrInput
	// Name of the guest configuration assignment.
	Name pulumi.StringPtrInput
	// Properties of the Guest configuration assignment.
	Properties GuestConfigurationAssignmentPropertiesPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the virtual machine.
	VmName pulumi.StringInput
}

func (GuestConfigurationConnectedVMwarevSphereAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationConnectedVMwarevSphereAssignmentArgs)(nil)).Elem()
}

type GuestConfigurationConnectedVMwarevSphereAssignmentInput interface {
	pulumi.Input

	ToGuestConfigurationConnectedVMwarevSphereAssignmentOutput() GuestConfigurationConnectedVMwarevSphereAssignmentOutput
	ToGuestConfigurationConnectedVMwarevSphereAssignmentOutputWithContext(ctx context.Context) GuestConfigurationConnectedVMwarevSphereAssignmentOutput
}

func (*GuestConfigurationConnectedVMwarevSphereAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationConnectedVMwarevSphereAssignment)(nil)).Elem()
}

func (i *GuestConfigurationConnectedVMwarevSphereAssignment) ToGuestConfigurationConnectedVMwarevSphereAssignmentOutput() GuestConfigurationConnectedVMwarevSphereAssignmentOutput {
	return i.ToGuestConfigurationConnectedVMwarevSphereAssignmentOutputWithContext(context.Background())
}

func (i *GuestConfigurationConnectedVMwarevSphereAssignment) ToGuestConfigurationConnectedVMwarevSphereAssignmentOutputWithContext(ctx context.Context) GuestConfigurationConnectedVMwarevSphereAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationConnectedVMwarevSphereAssignmentOutput)
}

type GuestConfigurationConnectedVMwarevSphereAssignmentOutput struct{ *pulumi.OutputState }

func (GuestConfigurationConnectedVMwarevSphereAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationConnectedVMwarevSphereAssignment)(nil)).Elem()
}

func (o GuestConfigurationConnectedVMwarevSphereAssignmentOutput) ToGuestConfigurationConnectedVMwarevSphereAssignmentOutput() GuestConfigurationConnectedVMwarevSphereAssignmentOutput {
	return o
}

func (o GuestConfigurationConnectedVMwarevSphereAssignmentOutput) ToGuestConfigurationConnectedVMwarevSphereAssignmentOutputWithContext(ctx context.Context) GuestConfigurationConnectedVMwarevSphereAssignmentOutput {
	return o
}

// Region where the VM is located.
func (o GuestConfigurationConnectedVMwarevSphereAssignmentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationConnectedVMwarevSphereAssignment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the guest configuration assignment.
func (o GuestConfigurationConnectedVMwarevSphereAssignmentOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationConnectedVMwarevSphereAssignment) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Properties of the Guest configuration assignment.
func (o GuestConfigurationConnectedVMwarevSphereAssignmentOutput) Properties() GuestConfigurationAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *GuestConfigurationConnectedVMwarevSphereAssignment) GuestConfigurationAssignmentPropertiesResponseOutput {
		return v.Properties
	}).(GuestConfigurationAssignmentPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o GuestConfigurationConnectedVMwarevSphereAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GuestConfigurationConnectedVMwarevSphereAssignment) SystemDataResponseOutput {
		return v.SystemData
	}).(SystemDataResponseOutput)
}

// The type of the resource.
func (o GuestConfigurationConnectedVMwarevSphereAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestConfigurationConnectedVMwarevSphereAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GuestConfigurationConnectedVMwarevSphereAssignmentOutput{})
}
