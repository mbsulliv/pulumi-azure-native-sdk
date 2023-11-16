// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a blueprint assignment.
type Assignment struct {
	pulumi.CustomResourceState

	// ID of the published version of a blueprint definition.
	BlueprintId pulumi.StringPtrOutput `pulumi:"blueprintId"`
	// Multi-line explain this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Managed identity for this blueprint assignment.
	Identity ManagedServiceIdentityResponseOutput `pulumi:"identity"`
	// The location of this blueprint assignment.
	Location pulumi.StringOutput `pulumi:"location"`
	// Defines how resources deployed by a blueprint assignment are locked.
	Locks AssignmentLockSettingsResponsePtrOutput `pulumi:"locks"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Blueprint assignment parameter values.
	Parameters ParameterValueResponseMapOutput `pulumi:"parameters"`
	// State of the blueprint assignment.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Names and locations of resource group placeholders.
	ResourceGroups ResourceGroupValueResponseMapOutput `pulumi:"resourceGroups"`
	// The target subscription scope of the blueprint assignment (format: '/subscriptions/{subscriptionId}'). For management group level assignments, the property is required.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Status of blueprint assignment. This field is readonly.
	Status AssignmentStatusResponseOutput `pulumi:"status"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAssignment registers a new resource with the given unique name, arguments, and options.
func NewAssignment(ctx *pulumi.Context,
	name string, args *AssignmentArgs, opts ...pulumi.ResourceOption) (*Assignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	if args.ResourceGroups == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroups'")
	}
	if args.ResourceScope == nil {
		return nil, errors.New("invalid value for required argument 'ResourceScope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:blueprint:Assignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Assignment
	err := ctx.RegisterResource("azure-native:blueprint/v20181101preview:Assignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssignment gets an existing Assignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssignmentState, opts ...pulumi.ResourceOption) (*Assignment, error) {
	var resource Assignment
	err := ctx.ReadResource("azure-native:blueprint/v20181101preview:Assignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Assignment resources.
type assignmentState struct {
}

type AssignmentState struct {
}

func (AssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentState)(nil)).Elem()
}

type assignmentArgs struct {
	// Name of the blueprint assignment.
	AssignmentName *string `pulumi:"assignmentName"`
	// ID of the published version of a blueprint definition.
	BlueprintId *string `pulumi:"blueprintId"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// Managed identity for this blueprint assignment.
	Identity ManagedServiceIdentity `pulumi:"identity"`
	// The location of this blueprint assignment.
	Location *string `pulumi:"location"`
	// Defines how resources deployed by a blueprint assignment are locked.
	Locks *AssignmentLockSettings `pulumi:"locks"`
	// Blueprint assignment parameter values.
	Parameters map[string]ParameterValue `pulumi:"parameters"`
	// Names and locations of resource group placeholders.
	ResourceGroups map[string]ResourceGroupValue `pulumi:"resourceGroups"`
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope string `pulumi:"resourceScope"`
	// The target subscription scope of the blueprint assignment (format: '/subscriptions/{subscriptionId}'). For management group level assignments, the property is required.
	Scope *string `pulumi:"scope"`
}

// The set of arguments for constructing a Assignment resource.
type AssignmentArgs struct {
	// Name of the blueprint assignment.
	AssignmentName pulumi.StringPtrInput
	// ID of the published version of a blueprint definition.
	BlueprintId pulumi.StringPtrInput
	// Multi-line explain this resource.
	Description pulumi.StringPtrInput
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrInput
	// Managed identity for this blueprint assignment.
	Identity ManagedServiceIdentityInput
	// The location of this blueprint assignment.
	Location pulumi.StringPtrInput
	// Defines how resources deployed by a blueprint assignment are locked.
	Locks AssignmentLockSettingsPtrInput
	// Blueprint assignment parameter values.
	Parameters ParameterValueMapInput
	// Names and locations of resource group placeholders.
	ResourceGroups ResourceGroupValueMapInput
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope pulumi.StringInput
	// The target subscription scope of the blueprint assignment (format: '/subscriptions/{subscriptionId}'). For management group level assignments, the property is required.
	Scope pulumi.StringPtrInput
}

func (AssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentArgs)(nil)).Elem()
}

type AssignmentInput interface {
	pulumi.Input

	ToAssignmentOutput() AssignmentOutput
	ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput
}

func (*Assignment) ElementType() reflect.Type {
	return reflect.TypeOf((**Assignment)(nil)).Elem()
}

func (i *Assignment) ToAssignmentOutput() AssignmentOutput {
	return i.ToAssignmentOutputWithContext(context.Background())
}

func (i *Assignment) ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentOutput)
}

type AssignmentOutput struct{ *pulumi.OutputState }

func (AssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Assignment)(nil)).Elem()
}

func (o AssignmentOutput) ToAssignmentOutput() AssignmentOutput {
	return o
}

func (o AssignmentOutput) ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput {
	return o
}

// ID of the published version of a blueprint definition.
func (o AssignmentOutput) BlueprintId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.BlueprintId }).(pulumi.StringPtrOutput)
}

// Multi-line explain this resource.
func (o AssignmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// One-liner string explain this resource.
func (o AssignmentOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Managed identity for this blueprint assignment.
func (o AssignmentOutput) Identity() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *Assignment) ManagedServiceIdentityResponseOutput { return v.Identity }).(ManagedServiceIdentityResponseOutput)
}

// The location of this blueprint assignment.
func (o AssignmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Defines how resources deployed by a blueprint assignment are locked.
func (o AssignmentOutput) Locks() AssignmentLockSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Assignment) AssignmentLockSettingsResponsePtrOutput { return v.Locks }).(AssignmentLockSettingsResponsePtrOutput)
}

// Name of this resource.
func (o AssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Blueprint assignment parameter values.
func (o AssignmentOutput) Parameters() ParameterValueResponseMapOutput {
	return o.ApplyT(func(v *Assignment) ParameterValueResponseMapOutput { return v.Parameters }).(ParameterValueResponseMapOutput)
}

// State of the blueprint assignment.
func (o AssignmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Names and locations of resource group placeholders.
func (o AssignmentOutput) ResourceGroups() ResourceGroupValueResponseMapOutput {
	return o.ApplyT(func(v *Assignment) ResourceGroupValueResponseMapOutput { return v.ResourceGroups }).(ResourceGroupValueResponseMapOutput)
}

// The target subscription scope of the blueprint assignment (format: '/subscriptions/{subscriptionId}'). For management group level assignments, the property is required.
func (o AssignmentOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// Status of blueprint assignment. This field is readonly.
func (o AssignmentOutput) Status() AssignmentStatusResponseOutput {
	return o.ApplyT(func(v *Assignment) AssignmentStatusResponseOutput { return v.Status }).(AssignmentStatusResponseOutput)
}

// Type of this resource.
func (o AssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssignmentOutput{})
}
