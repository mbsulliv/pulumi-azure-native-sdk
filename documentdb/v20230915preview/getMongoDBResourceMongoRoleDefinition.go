// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230915preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the properties of an existing Azure Cosmos DB Mongo Role Definition with the given Id.
func LookupMongoDBResourceMongoRoleDefinition(ctx *pulumi.Context, args *LookupMongoDBResourceMongoRoleDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupMongoDBResourceMongoRoleDefinitionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMongoDBResourceMongoRoleDefinitionResult
	err := ctx.Invoke("azure-native:documentdb/v20230915preview:getMongoDBResourceMongoRoleDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoDBResourceMongoRoleDefinitionArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The ID for the Role Definition {dbName.roleName}.
	MongoRoleDefinitionId string `pulumi:"mongoRoleDefinitionId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Cosmos DB Mongo Role Definition.
type LookupMongoDBResourceMongoRoleDefinitionResult struct {
	// The database name for which access is being granted for this Role Definition.
	DatabaseName *string `pulumi:"databaseName"`
	// The unique resource identifier of the database account.
	Id string `pulumi:"id"`
	// The name of the database account.
	Name string `pulumi:"name"`
	// A set of privileges contained by the Role Definition. This will allow application of this Role Definition on the entire database account or any underlying Database / Collection. Scopes higher than Database are not enforceable as privilege.
	Privileges []PrivilegeResponse `pulumi:"privileges"`
	// A user-friendly name for the Role Definition. Must be unique for the database account.
	RoleName *string `pulumi:"roleName"`
	// The set of roles inherited by this Role Definition.
	Roles []RoleResponse `pulumi:"roles"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupMongoDBResourceMongoRoleDefinitionOutput(ctx *pulumi.Context, args LookupMongoDBResourceMongoRoleDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupMongoDBResourceMongoRoleDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMongoDBResourceMongoRoleDefinitionResult, error) {
			args := v.(LookupMongoDBResourceMongoRoleDefinitionArgs)
			r, err := LookupMongoDBResourceMongoRoleDefinition(ctx, &args, opts...)
			var s LookupMongoDBResourceMongoRoleDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMongoDBResourceMongoRoleDefinitionResultOutput)
}

type LookupMongoDBResourceMongoRoleDefinitionOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The ID for the Role Definition {dbName.roleName}.
	MongoRoleDefinitionId pulumi.StringInput `pulumi:"mongoRoleDefinitionId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMongoDBResourceMongoRoleDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoRoleDefinitionArgs)(nil)).Elem()
}

// An Azure Cosmos DB Mongo Role Definition.
type LookupMongoDBResourceMongoRoleDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupMongoDBResourceMongoRoleDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoRoleDefinitionResult)(nil)).Elem()
}

func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) ToLookupMongoDBResourceMongoRoleDefinitionResultOutput() LookupMongoDBResourceMongoRoleDefinitionResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) ToLookupMongoDBResourceMongoRoleDefinitionResultOutputWithContext(ctx context.Context) LookupMongoDBResourceMongoRoleDefinitionResultOutput {
	return o
}

// The database name for which access is being granted for this Role Definition.
func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// The unique resource identifier of the database account.
func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the database account.
func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

// A set of privileges contained by the Role Definition. This will allow application of this Role Definition on the entire database account or any underlying Database / Collection. Scopes higher than Database are not enforceable as privilege.
func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) Privileges() PrivilegeResponseArrayOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) []PrivilegeResponse { return v.Privileges }).(PrivilegeResponseArrayOutput)
}

// A user-friendly name for the Role Definition. Must be unique for the database account.
func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) *string { return v.RoleName }).(pulumi.StringPtrOutput)
}

// The set of roles inherited by this Role Definition.
func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) Roles() RoleResponseArrayOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) []RoleResponse { return v.Roles }).(RoleResponseArrayOutput)
}

// The type of Azure resource.
func (o LookupMongoDBResourceMongoRoleDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoRoleDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMongoDBResourceMongoRoleDefinitionResultOutput{})
}
