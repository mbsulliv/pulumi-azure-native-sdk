// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Role management policy
type RoleManagementPolicyAssignment struct {
	pulumi.CustomResourceState

	// The readonly computed rule applied to the policy.
	EffectiveRules pulumi.ArrayOutput `pulumi:"effectiveRules"`
	// The role management policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Additional properties of scope, role definition and policy
	PolicyAssignmentProperties PolicyAssignmentPropertiesResponseOutput `pulumi:"policyAssignmentProperties"`
	// The policy id role management policy assignment.
	PolicyId pulumi.StringPtrOutput `pulumi:"policyId"`
	// The role definition of management policy assignment.
	RoleDefinitionId pulumi.StringPtrOutput `pulumi:"roleDefinitionId"`
	// The role management policy scope.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// The role management policy type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRoleManagementPolicyAssignment registers a new resource with the given unique name, arguments, and options.
func NewRoleManagementPolicyAssignment(ctx *pulumi.Context,
	name string, args *RoleManagementPolicyAssignmentArgs, opts ...pulumi.ResourceOption) (*RoleManagementPolicyAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:RoleManagementPolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20201001preview:RoleManagementPolicyAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RoleManagementPolicyAssignment
	err := ctx.RegisterResource("azure-native:authorization/v20201001:RoleManagementPolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleManagementPolicyAssignment gets an existing RoleManagementPolicyAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleManagementPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleManagementPolicyAssignmentState, opts ...pulumi.ResourceOption) (*RoleManagementPolicyAssignment, error) {
	var resource RoleManagementPolicyAssignment
	err := ctx.ReadResource("azure-native:authorization/v20201001:RoleManagementPolicyAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleManagementPolicyAssignment resources.
type roleManagementPolicyAssignmentState struct {
}

type RoleManagementPolicyAssignmentState struct {
}

func (RoleManagementPolicyAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleManagementPolicyAssignmentState)(nil)).Elem()
}

type roleManagementPolicyAssignmentArgs struct {
	// The policy id role management policy assignment.
	PolicyId *string `pulumi:"policyId"`
	// The role definition of management policy assignment.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The name of format {guid_guid} the role management policy assignment to upsert.
	RoleManagementPolicyAssignmentName *string `pulumi:"roleManagementPolicyAssignmentName"`
	// The role management policy scope.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a RoleManagementPolicyAssignment resource.
type RoleManagementPolicyAssignmentArgs struct {
	// The policy id role management policy assignment.
	PolicyId pulumi.StringPtrInput
	// The role definition of management policy assignment.
	RoleDefinitionId pulumi.StringPtrInput
	// The name of format {guid_guid} the role management policy assignment to upsert.
	RoleManagementPolicyAssignmentName pulumi.StringPtrInput
	// The role management policy scope.
	Scope pulumi.StringInput
}

func (RoleManagementPolicyAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleManagementPolicyAssignmentArgs)(nil)).Elem()
}

type RoleManagementPolicyAssignmentInput interface {
	pulumi.Input

	ToRoleManagementPolicyAssignmentOutput() RoleManagementPolicyAssignmentOutput
	ToRoleManagementPolicyAssignmentOutputWithContext(ctx context.Context) RoleManagementPolicyAssignmentOutput
}

func (*RoleManagementPolicyAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleManagementPolicyAssignment)(nil)).Elem()
}

func (i *RoleManagementPolicyAssignment) ToRoleManagementPolicyAssignmentOutput() RoleManagementPolicyAssignmentOutput {
	return i.ToRoleManagementPolicyAssignmentOutputWithContext(context.Background())
}

func (i *RoleManagementPolicyAssignment) ToRoleManagementPolicyAssignmentOutputWithContext(ctx context.Context) RoleManagementPolicyAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleManagementPolicyAssignmentOutput)
}

func (i *RoleManagementPolicyAssignment) ToOutput(ctx context.Context) pulumix.Output[*RoleManagementPolicyAssignment] {
	return pulumix.Output[*RoleManagementPolicyAssignment]{
		OutputState: i.ToRoleManagementPolicyAssignmentOutputWithContext(ctx).OutputState,
	}
}

type RoleManagementPolicyAssignmentOutput struct{ *pulumi.OutputState }

func (RoleManagementPolicyAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleManagementPolicyAssignment)(nil)).Elem()
}

func (o RoleManagementPolicyAssignmentOutput) ToRoleManagementPolicyAssignmentOutput() RoleManagementPolicyAssignmentOutput {
	return o
}

func (o RoleManagementPolicyAssignmentOutput) ToRoleManagementPolicyAssignmentOutputWithContext(ctx context.Context) RoleManagementPolicyAssignmentOutput {
	return o
}

func (o RoleManagementPolicyAssignmentOutput) ToOutput(ctx context.Context) pulumix.Output[*RoleManagementPolicyAssignment] {
	return pulumix.Output[*RoleManagementPolicyAssignment]{
		OutputState: o.OutputState,
	}
}

// The readonly computed rule applied to the policy.
func (o RoleManagementPolicyAssignmentOutput) EffectiveRules() pulumi.ArrayOutput {
	return o.ApplyT(func(v *RoleManagementPolicyAssignment) pulumi.ArrayOutput { return v.EffectiveRules }).(pulumi.ArrayOutput)
}

// The role management policy name.
func (o RoleManagementPolicyAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleManagementPolicyAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Additional properties of scope, role definition and policy
func (o RoleManagementPolicyAssignmentOutput) PolicyAssignmentProperties() PolicyAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *RoleManagementPolicyAssignment) PolicyAssignmentPropertiesResponseOutput {
		return v.PolicyAssignmentProperties
	}).(PolicyAssignmentPropertiesResponseOutput)
}

// The policy id role management policy assignment.
func (o RoleManagementPolicyAssignmentOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleManagementPolicyAssignment) pulumi.StringPtrOutput { return v.PolicyId }).(pulumi.StringPtrOutput)
}

// The role definition of management policy assignment.
func (o RoleManagementPolicyAssignmentOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleManagementPolicyAssignment) pulumi.StringPtrOutput { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

// The role management policy scope.
func (o RoleManagementPolicyAssignmentOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleManagementPolicyAssignment) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// The role management policy type.
func (o RoleManagementPolicyAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleManagementPolicyAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RoleManagementPolicyAssignmentOutput{})
}
