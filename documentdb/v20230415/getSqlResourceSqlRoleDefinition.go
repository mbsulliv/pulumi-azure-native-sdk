// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230415

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the properties of an existing Azure Cosmos DB SQL Role Definition with the given Id.
func LookupSqlResourceSqlRoleDefinition(ctx *pulumi.Context, args *LookupSqlResourceSqlRoleDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlRoleDefinitionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlResourceSqlRoleDefinitionResult
	err := ctx.Invoke("azure-native:documentdb/v20230415:getSqlResourceSqlRoleDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlRoleDefinitionArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The GUID for the Role Definition.
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}

// An Azure Cosmos DB SQL Role Definition.
type LookupSqlResourceSqlRoleDefinitionResult struct {
	// A set of fully qualified Scopes at or below which Role Assignments may be created using this Role Definition. This will allow application of this Role Definition on the entire database account or any underlying Database / Collection. Must have at least one element. Scopes higher than Database account are not enforceable as assignable Scopes. Note that resources referenced in assignable Scopes need not exist.
	AssignableScopes []string `pulumi:"assignableScopes"`
	// The unique resource identifier of the database account.
	Id string `pulumi:"id"`
	// The name of the database account.
	Name string `pulumi:"name"`
	// The set of operations allowed through this Role Definition.
	Permissions []PermissionResponse `pulumi:"permissions"`
	// A user-friendly name for the Role Definition. Must be unique for the database account.
	RoleName *string `pulumi:"roleName"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupSqlResourceSqlRoleDefinitionOutput(ctx *pulumi.Context, args LookupSqlResourceSqlRoleDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupSqlResourceSqlRoleDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlResourceSqlRoleDefinitionResult, error) {
			args := v.(LookupSqlResourceSqlRoleDefinitionArgs)
			r, err := LookupSqlResourceSqlRoleDefinition(ctx, &args, opts...)
			var s LookupSqlResourceSqlRoleDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlResourceSqlRoleDefinitionResultOutput)
}

type LookupSqlResourceSqlRoleDefinitionOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The GUID for the Role Definition.
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (LookupSqlResourceSqlRoleDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlRoleDefinitionArgs)(nil)).Elem()
}

// An Azure Cosmos DB SQL Role Definition.
type LookupSqlResourceSqlRoleDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupSqlResourceSqlRoleDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlRoleDefinitionResult)(nil)).Elem()
}

func (o LookupSqlResourceSqlRoleDefinitionResultOutput) ToLookupSqlResourceSqlRoleDefinitionResultOutput() LookupSqlResourceSqlRoleDefinitionResultOutput {
	return o
}

func (o LookupSqlResourceSqlRoleDefinitionResultOutput) ToLookupSqlResourceSqlRoleDefinitionResultOutputWithContext(ctx context.Context) LookupSqlResourceSqlRoleDefinitionResultOutput {
	return o
}

// A set of fully qualified Scopes at or below which Role Assignments may be created using this Role Definition. This will allow application of this Role Definition on the entire database account or any underlying Database / Collection. Must have at least one element. Scopes higher than Database account are not enforceable as assignable Scopes. Note that resources referenced in assignable Scopes need not exist.
func (o LookupSqlResourceSqlRoleDefinitionResultOutput) AssignableScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleDefinitionResult) []string { return v.AssignableScopes }).(pulumi.StringArrayOutput)
}

// The unique resource identifier of the database account.
func (o LookupSqlResourceSqlRoleDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the database account.
func (o LookupSqlResourceSqlRoleDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The set of operations allowed through this Role Definition.
func (o LookupSqlResourceSqlRoleDefinitionResultOutput) Permissions() PermissionResponseArrayOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleDefinitionResult) []PermissionResponse { return v.Permissions }).(PermissionResponseArrayOutput)
}

// A user-friendly name for the Role Definition. Must be unique for the database account.
func (o LookupSqlResourceSqlRoleDefinitionResultOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleDefinitionResult) *string { return v.RoleName }).(pulumi.StringPtrOutput)
}

// The type of Azure resource.
func (o LookupSqlResourceSqlRoleDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlResourceSqlRoleDefinitionResultOutput{})
}
