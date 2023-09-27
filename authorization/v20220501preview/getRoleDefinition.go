// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get role definition by ID (GUID).
func LookupRoleDefinition(ctx *pulumi.Context, args *LookupRoleDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupRoleDefinitionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRoleDefinitionResult
	err := ctx.Invoke("azure-native:authorization/v20220501preview:getRoleDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleDefinitionArgs struct {
	// The ID of the role definition.
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
	// The scope of the operation or resource. Valid scopes are: subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}', or resource (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
	Scope string `pulumi:"scope"`
}

// Role definition.
type LookupRoleDefinitionResult struct {
	// Role definition assignable scopes.
	AssignableScopes []string `pulumi:"assignableScopes"`
	// Id of the user who created the assignment
	CreatedBy string `pulumi:"createdBy"`
	// Time it was created
	CreatedOn string `pulumi:"createdOn"`
	// The role definition description.
	Description *string `pulumi:"description"`
	// The role definition ID.
	Id string `pulumi:"id"`
	// The role definition name.
	Name string `pulumi:"name"`
	// Role definition permissions.
	Permissions []PermissionResponse `pulumi:"permissions"`
	// The role name.
	RoleName *string `pulumi:"roleName"`
	// The role type.
	RoleType *string `pulumi:"roleType"`
	// The role definition type.
	Type string `pulumi:"type"`
	// Id of the user who updated the assignment
	UpdatedBy string `pulumi:"updatedBy"`
	// Time it was updated
	UpdatedOn string `pulumi:"updatedOn"`
}

func LookupRoleDefinitionOutput(ctx *pulumi.Context, args LookupRoleDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupRoleDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoleDefinitionResult, error) {
			args := v.(LookupRoleDefinitionArgs)
			r, err := LookupRoleDefinition(ctx, &args, opts...)
			var s LookupRoleDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoleDefinitionResultOutput)
}

type LookupRoleDefinitionOutputArgs struct {
	// The ID of the role definition.
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
	// The scope of the operation or resource. Valid scopes are: subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}', or resource (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupRoleDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleDefinitionArgs)(nil)).Elem()
}

// Role definition.
type LookupRoleDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupRoleDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleDefinitionResult)(nil)).Elem()
}

func (o LookupRoleDefinitionResultOutput) ToLookupRoleDefinitionResultOutput() LookupRoleDefinitionResultOutput {
	return o
}

func (o LookupRoleDefinitionResultOutput) ToLookupRoleDefinitionResultOutputWithContext(ctx context.Context) LookupRoleDefinitionResultOutput {
	return o
}

func (o LookupRoleDefinitionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRoleDefinitionResult] {
	return pulumix.Output[LookupRoleDefinitionResult]{
		OutputState: o.OutputState,
	}
}

// Role definition assignable scopes.
func (o LookupRoleDefinitionResultOutput) AssignableScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) []string { return v.AssignableScopes }).(pulumi.StringArrayOutput)
}

// Id of the user who created the assignment
func (o LookupRoleDefinitionResultOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.CreatedBy }).(pulumi.StringOutput)
}

// Time it was created
func (o LookupRoleDefinitionResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

// The role definition description.
func (o LookupRoleDefinitionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The role definition ID.
func (o LookupRoleDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The role definition name.
func (o LookupRoleDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Role definition permissions.
func (o LookupRoleDefinitionResultOutput) Permissions() PermissionResponseArrayOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) []PermissionResponse { return v.Permissions }).(PermissionResponseArrayOutput)
}

// The role name.
func (o LookupRoleDefinitionResultOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) *string { return v.RoleName }).(pulumi.StringPtrOutput)
}

// The role type.
func (o LookupRoleDefinitionResultOutput) RoleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) *string { return v.RoleType }).(pulumi.StringPtrOutput)
}

// The role definition type.
func (o LookupRoleDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

// Id of the user who updated the assignment
func (o LookupRoleDefinitionResultOutput) UpdatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.UpdatedBy }).(pulumi.StringOutput)
}

// Time it was updated
func (o LookupRoleDefinitionResultOutput) UpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.UpdatedOn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleDefinitionResultOutput{})
}
