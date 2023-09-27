// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230415

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves the properties of an existing Azure Cosmos DB Mongo User Definition with the given Id.
func LookupMongoDBResourceMongoUserDefinition(ctx *pulumi.Context, args *LookupMongoDBResourceMongoUserDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupMongoDBResourceMongoUserDefinitionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMongoDBResourceMongoUserDefinitionResult
	err := ctx.Invoke("azure-native:documentdb/v20230415:getMongoDBResourceMongoUserDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoDBResourceMongoUserDefinitionArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The ID for the User Definition {dbName.userName}.
	MongoUserDefinitionId string `pulumi:"mongoUserDefinitionId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Cosmos DB User Definition
type LookupMongoDBResourceMongoUserDefinitionResult struct {
	// A custom definition for the USer Definition.
	CustomData *string `pulumi:"customData"`
	// The database name for which access is being granted for this User Definition.
	DatabaseName *string `pulumi:"databaseName"`
	// The unique resource identifier of the database account.
	Id string `pulumi:"id"`
	// The Mongo Auth mechanism. For now, we only support auth mechanism SCRAM-SHA-256.
	Mechanisms *string `pulumi:"mechanisms"`
	// The name of the database account.
	Name string `pulumi:"name"`
	// The password for User Definition. Response does not contain user password.
	Password *string `pulumi:"password"`
	// The set of roles inherited by the User Definition.
	Roles []RoleResponse `pulumi:"roles"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
	// The user name for User Definition.
	UserName *string `pulumi:"userName"`
}

func LookupMongoDBResourceMongoUserDefinitionOutput(ctx *pulumi.Context, args LookupMongoDBResourceMongoUserDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupMongoDBResourceMongoUserDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMongoDBResourceMongoUserDefinitionResult, error) {
			args := v.(LookupMongoDBResourceMongoUserDefinitionArgs)
			r, err := LookupMongoDBResourceMongoUserDefinition(ctx, &args, opts...)
			var s LookupMongoDBResourceMongoUserDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMongoDBResourceMongoUserDefinitionResultOutput)
}

type LookupMongoDBResourceMongoUserDefinitionOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The ID for the User Definition {dbName.userName}.
	MongoUserDefinitionId pulumi.StringInput `pulumi:"mongoUserDefinitionId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMongoDBResourceMongoUserDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoUserDefinitionArgs)(nil)).Elem()
}

// An Azure Cosmos DB User Definition
type LookupMongoDBResourceMongoUserDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupMongoDBResourceMongoUserDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoUserDefinitionResult)(nil)).Elem()
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) ToLookupMongoDBResourceMongoUserDefinitionResultOutput() LookupMongoDBResourceMongoUserDefinitionResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) ToLookupMongoDBResourceMongoUserDefinitionResultOutputWithContext(ctx context.Context) LookupMongoDBResourceMongoUserDefinitionResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMongoDBResourceMongoUserDefinitionResult] {
	return pulumix.Output[LookupMongoDBResourceMongoUserDefinitionResult]{
		OutputState: o.OutputState,
	}
}

// A custom definition for the USer Definition.
func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

// The database name for which access is being granted for this User Definition.
func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// The unique resource identifier of the database account.
func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Mongo Auth mechanism. For now, we only support auth mechanism SCRAM-SHA-256.
func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) Mechanisms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) *string { return v.Mechanisms }).(pulumi.StringPtrOutput)
}

// The name of the database account.
func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The password for User Definition. Response does not contain user password.
func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

// The set of roles inherited by the User Definition.
func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) Roles() RoleResponseArrayOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) []RoleResponse { return v.Roles }).(RoleResponseArrayOutput)
}

// The type of Azure resource.
func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

// The user name for User Definition.
func (o LookupMongoDBResourceMongoUserDefinitionResultOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoUserDefinitionResult) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMongoDBResourceMongoUserDefinitionResultOutput{})
}
