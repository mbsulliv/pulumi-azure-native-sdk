// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a role assignment for the caller on a enrollment Account. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
// Azure REST API version: 2019-10-01-preview.
func LookupBillingRoleAssignmentByEnrollmentAccount(ctx *pulumi.Context, args *LookupBillingRoleAssignmentByEnrollmentAccountArgs, opts ...pulumi.InvokeOption) (*LookupBillingRoleAssignmentByEnrollmentAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBillingRoleAssignmentByEnrollmentAccountResult
	err := ctx.Invoke("azure-native:billing:getBillingRoleAssignmentByEnrollmentAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBillingRoleAssignmentByEnrollmentAccountArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName string `pulumi:"billingAccountName"`
	// The ID that uniquely identifies a role assignment.
	BillingRoleAssignmentName string `pulumi:"billingRoleAssignmentName"`
	// The ID that uniquely identifies an enrollment account.
	EnrollmentAccountName string `pulumi:"enrollmentAccountName"`
}

// The role assignment
type LookupBillingRoleAssignmentByEnrollmentAccountResult struct {
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

func LookupBillingRoleAssignmentByEnrollmentAccountOutput(ctx *pulumi.Context, args LookupBillingRoleAssignmentByEnrollmentAccountOutputArgs, opts ...pulumi.InvokeOption) LookupBillingRoleAssignmentByEnrollmentAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBillingRoleAssignmentByEnrollmentAccountResult, error) {
			args := v.(LookupBillingRoleAssignmentByEnrollmentAccountArgs)
			r, err := LookupBillingRoleAssignmentByEnrollmentAccount(ctx, &args, opts...)
			var s LookupBillingRoleAssignmentByEnrollmentAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBillingRoleAssignmentByEnrollmentAccountResultOutput)
}

type LookupBillingRoleAssignmentByEnrollmentAccountOutputArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName pulumi.StringInput `pulumi:"billingAccountName"`
	// The ID that uniquely identifies a role assignment.
	BillingRoleAssignmentName pulumi.StringInput `pulumi:"billingRoleAssignmentName"`
	// The ID that uniquely identifies an enrollment account.
	EnrollmentAccountName pulumi.StringInput `pulumi:"enrollmentAccountName"`
}

func (LookupBillingRoleAssignmentByEnrollmentAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingRoleAssignmentByEnrollmentAccountArgs)(nil)).Elem()
}

// The role assignment
type LookupBillingRoleAssignmentByEnrollmentAccountResultOutput struct{ *pulumi.OutputState }

func (LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingRoleAssignmentByEnrollmentAccountResult)(nil)).Elem()
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) ToLookupBillingRoleAssignmentByEnrollmentAccountResultOutput() LookupBillingRoleAssignmentByEnrollmentAccountResultOutput {
	return o
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) ToLookupBillingRoleAssignmentByEnrollmentAccountResultOutputWithContext(ctx context.Context) LookupBillingRoleAssignmentByEnrollmentAccountResultOutput {
	return o
}

// The principal Id of the user who created the role assignment.
func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) CreatedByPrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string { return v.CreatedByPrincipalId }).(pulumi.StringOutput)
}

// The tenant Id of the user who created the role assignment.
func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) CreatedByPrincipalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string {
		return v.CreatedByPrincipalTenantId
	}).(pulumi.StringOutput)
}

// The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) CreatedByUserEmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string {
		return v.CreatedByUserEmailAddress
	}).(pulumi.StringOutput)
}

// The date the role assignment was created.
func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

// Resource Id.
func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// The principal id of the user to whom the role was assigned.
func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

// The principal tenant id of the user to whom the role was assigned.
func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) PrincipalTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) *string { return v.PrincipalTenantId }).(pulumi.StringPtrOutput)
}

// The ID of the role definition.
func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) *string { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

// The scope at which the role was assigned.
func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string { return v.Scope }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) UserAuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) *string { return v.UserAuthenticationType }).(pulumi.StringPtrOutput)
}

// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) UserEmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) *string { return v.UserEmailAddress }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBillingRoleAssignmentByEnrollmentAccountResultOutput{})
}
