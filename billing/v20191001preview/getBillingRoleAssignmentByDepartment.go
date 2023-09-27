// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a role assignment for the caller on a department. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
func LookupBillingRoleAssignmentByDepartment(ctx *pulumi.Context, args *LookupBillingRoleAssignmentByDepartmentArgs, opts ...pulumi.InvokeOption) (*LookupBillingRoleAssignmentByDepartmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBillingRoleAssignmentByDepartmentResult
	err := ctx.Invoke("azure-native:billing/v20191001preview:getBillingRoleAssignmentByDepartment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBillingRoleAssignmentByDepartmentArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName string `pulumi:"billingAccountName"`
	// The ID that uniquely identifies a role assignment.
	BillingRoleAssignmentName string `pulumi:"billingRoleAssignmentName"`
	// The ID that uniquely identifies a department.
	DepartmentName string `pulumi:"departmentName"`
}

// The role assignment
type LookupBillingRoleAssignmentByDepartmentResult struct {
	// The principal Id of the user who created the role assignment.
	CreatedByPrincipalId string `pulumi:"createdByPrincipalId"`
	// The tenant Id of the user who created the role assignment.
	CreatedByPrincipalTenantId string `pulumi:"createdByPrincipalTenantId"`
	// The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
	CreatedByUserEmailAddress string `pulumi:"createdByUserEmailAddress"`
	// The date the role assignment was created.
	CreatedOn string `pulumi:"createdOn"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The principal id of the user to whom the role was assigned.
	PrincipalId *string `pulumi:"principalId"`
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId *string `pulumi:"principalTenantId"`
	// The ID of the role definition.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The scope at which the role was assigned.
	Scope string `pulumi:"scope"`
	// Resource type.
	Type string `pulumi:"type"`
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType *string `pulumi:"userAuthenticationType"`
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress *string `pulumi:"userEmailAddress"`
}

func LookupBillingRoleAssignmentByDepartmentOutput(ctx *pulumi.Context, args LookupBillingRoleAssignmentByDepartmentOutputArgs, opts ...pulumi.InvokeOption) LookupBillingRoleAssignmentByDepartmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBillingRoleAssignmentByDepartmentResult, error) {
			args := v.(LookupBillingRoleAssignmentByDepartmentArgs)
			r, err := LookupBillingRoleAssignmentByDepartment(ctx, &args, opts...)
			var s LookupBillingRoleAssignmentByDepartmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBillingRoleAssignmentByDepartmentResultOutput)
}

type LookupBillingRoleAssignmentByDepartmentOutputArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName pulumi.StringInput `pulumi:"billingAccountName"`
	// The ID that uniquely identifies a role assignment.
	BillingRoleAssignmentName pulumi.StringInput `pulumi:"billingRoleAssignmentName"`
	// The ID that uniquely identifies a department.
	DepartmentName pulumi.StringInput `pulumi:"departmentName"`
}

func (LookupBillingRoleAssignmentByDepartmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingRoleAssignmentByDepartmentArgs)(nil)).Elem()
}

// The role assignment
type LookupBillingRoleAssignmentByDepartmentResultOutput struct{ *pulumi.OutputState }

func (LookupBillingRoleAssignmentByDepartmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingRoleAssignmentByDepartmentResult)(nil)).Elem()
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) ToLookupBillingRoleAssignmentByDepartmentResultOutput() LookupBillingRoleAssignmentByDepartmentResultOutput {
	return o
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) ToLookupBillingRoleAssignmentByDepartmentResultOutputWithContext(ctx context.Context) LookupBillingRoleAssignmentByDepartmentResultOutput {
	return o
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBillingRoleAssignmentByDepartmentResult] {
	return pulumix.Output[LookupBillingRoleAssignmentByDepartmentResult]{
		OutputState: o.OutputState,
	}
}

// The principal Id of the user who created the role assignment.
func (o LookupBillingRoleAssignmentByDepartmentResultOutput) CreatedByPrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.CreatedByPrincipalId }).(pulumi.StringOutput)
}

// The tenant Id of the user who created the role assignment.
func (o LookupBillingRoleAssignmentByDepartmentResultOutput) CreatedByPrincipalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.CreatedByPrincipalTenantId }).(pulumi.StringOutput)
}

// The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
func (o LookupBillingRoleAssignmentByDepartmentResultOutput) CreatedByUserEmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.CreatedByUserEmailAddress }).(pulumi.StringOutput)
}

// The date the role assignment was created.
func (o LookupBillingRoleAssignmentByDepartmentResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

// Resource Id.
func (o LookupBillingRoleAssignmentByDepartmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupBillingRoleAssignmentByDepartmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The principal id of the user to whom the role was assigned.
func (o LookupBillingRoleAssignmentByDepartmentResultOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

// The principal tenant id of the user to whom the role was assigned.
func (o LookupBillingRoleAssignmentByDepartmentResultOutput) PrincipalTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) *string { return v.PrincipalTenantId }).(pulumi.StringPtrOutput)
}

// The ID of the role definition.
func (o LookupBillingRoleAssignmentByDepartmentResultOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) *string { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

// The scope at which the role was assigned.
func (o LookupBillingRoleAssignmentByDepartmentResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.Scope }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupBillingRoleAssignmentByDepartmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.Type }).(pulumi.StringOutput)
}

// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
func (o LookupBillingRoleAssignmentByDepartmentResultOutput) UserAuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) *string { return v.UserAuthenticationType }).(pulumi.StringPtrOutput)
}

// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
func (o LookupBillingRoleAssignmentByDepartmentResultOutput) UserEmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) *string { return v.UserEmailAddress }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBillingRoleAssignmentByDepartmentResultOutput{})
}
