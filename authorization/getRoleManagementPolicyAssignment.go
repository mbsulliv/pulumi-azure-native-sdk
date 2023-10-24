// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the specified role management policy assignment for a resource scope
// Azure REST API version: 2020-10-01.
//
// Other available API versions: 2020-10-01-preview.
func LookupRoleManagementPolicyAssignment(ctx *pulumi.Context, args *LookupRoleManagementPolicyAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleManagementPolicyAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRoleManagementPolicyAssignmentResult
	err := ctx.Invoke("azure-native:authorization:getRoleManagementPolicyAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleManagementPolicyAssignmentArgs struct {
	// The name of format {guid_guid} the role management policy assignment to get.
	RoleManagementPolicyAssignmentName string `pulumi:"roleManagementPolicyAssignmentName"`
	// The scope of the role management policy.
	Scope string `pulumi:"scope"`
}

// Role management policy
type LookupRoleManagementPolicyAssignmentResult struct {
	// The readonly computed rule applied to the policy.
	EffectiveRules []interface{} `pulumi:"effectiveRules"`
	// The role management policy Id.
	Id string `pulumi:"id"`
	// The role management policy name.
	Name string `pulumi:"name"`
	// Additional properties of scope, role definition and policy
	PolicyAssignmentProperties PolicyAssignmentPropertiesResponse `pulumi:"policyAssignmentProperties"`
	// The policy id role management policy assignment.
	PolicyId *string `pulumi:"policyId"`
	// The role definition of management policy assignment.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The role management policy scope.
	Scope *string `pulumi:"scope"`
	// The role management policy type.
	Type string `pulumi:"type"`
}

func LookupRoleManagementPolicyAssignmentOutput(ctx *pulumi.Context, args LookupRoleManagementPolicyAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupRoleManagementPolicyAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoleManagementPolicyAssignmentResult, error) {
			args := v.(LookupRoleManagementPolicyAssignmentArgs)
			r, err := LookupRoleManagementPolicyAssignment(ctx, &args, opts...)
			var s LookupRoleManagementPolicyAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoleManagementPolicyAssignmentResultOutput)
}

type LookupRoleManagementPolicyAssignmentOutputArgs struct {
	// The name of format {guid_guid} the role management policy assignment to get.
	RoleManagementPolicyAssignmentName pulumi.StringInput `pulumi:"roleManagementPolicyAssignmentName"`
	// The scope of the role management policy.
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupRoleManagementPolicyAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleManagementPolicyAssignmentArgs)(nil)).Elem()
}

// Role management policy
type LookupRoleManagementPolicyAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupRoleManagementPolicyAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleManagementPolicyAssignmentResult)(nil)).Elem()
}

func (o LookupRoleManagementPolicyAssignmentResultOutput) ToLookupRoleManagementPolicyAssignmentResultOutput() LookupRoleManagementPolicyAssignmentResultOutput {
	return o
}

func (o LookupRoleManagementPolicyAssignmentResultOutput) ToLookupRoleManagementPolicyAssignmentResultOutputWithContext(ctx context.Context) LookupRoleManagementPolicyAssignmentResultOutput {
	return o
}

func (o LookupRoleManagementPolicyAssignmentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRoleManagementPolicyAssignmentResult] {
	return pulumix.Output[LookupRoleManagementPolicyAssignmentResult]{
		OutputState: o.OutputState,
	}
}

// The readonly computed rule applied to the policy.
func (o LookupRoleManagementPolicyAssignmentResultOutput) EffectiveRules() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) []interface{} { return v.EffectiveRules }).(pulumi.ArrayOutput)
}

// The role management policy Id.
func (o LookupRoleManagementPolicyAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The role management policy name.
func (o LookupRoleManagementPolicyAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Additional properties of scope, role definition and policy
func (o LookupRoleManagementPolicyAssignmentResultOutput) PolicyAssignmentProperties() PolicyAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) PolicyAssignmentPropertiesResponse {
		return v.PolicyAssignmentProperties
	}).(PolicyAssignmentPropertiesResponseOutput)
}

// The policy id role management policy assignment.
func (o LookupRoleManagementPolicyAssignmentResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

// The role definition of management policy assignment.
func (o LookupRoleManagementPolicyAssignmentResultOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) *string { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

// The role management policy scope.
func (o LookupRoleManagementPolicyAssignmentResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// The role management policy type.
func (o LookupRoleManagementPolicyAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleManagementPolicyAssignmentResultOutput{})
}
