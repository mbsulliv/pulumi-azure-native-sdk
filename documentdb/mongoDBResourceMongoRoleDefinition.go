// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package documentdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB Mongo Role Definition.
// Azure REST API version: 2023-04-15. Prior API version in Azure Native 1.x: 2021-10-15-preview.
//
// Other available API versions: 2023-03-01-preview, 2023-09-15, 2023-09-15-preview.
type MongoDBResourceMongoRoleDefinition struct {
	pulumi.CustomResourceState

	// The database name for which access is being granted for this Role Definition.
	DatabaseName pulumi.StringPtrOutput `pulumi:"databaseName"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// A set of privileges contained by the Role Definition. This will allow application of this Role Definition on the entire database account or any underlying Database / Collection. Scopes higher than Database are not enforceable as privilege.
	Privileges PrivilegeResponseArrayOutput `pulumi:"privileges"`
	// A user-friendly name for the Role Definition. Must be unique for the database account.
	RoleName pulumi.StringPtrOutput `pulumi:"roleName"`
	// The set of roles inherited by this Role Definition.
	Roles RoleResponseArrayOutput `pulumi:"roles"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMongoDBResourceMongoRoleDefinition registers a new resource with the given unique name, arguments, and options.
func NewMongoDBResourceMongoRoleDefinition(ctx *pulumi.Context,
	name string, args *MongoDBResourceMongoRoleDefinitionArgs, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoRoleDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115preview:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230301preview:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230415:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:MongoDBResourceMongoRoleDefinition"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MongoDBResourceMongoRoleDefinition
	err := ctx.RegisterResource("azure-native:documentdb:MongoDBResourceMongoRoleDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMongoDBResourceMongoRoleDefinition gets an existing MongoDBResourceMongoRoleDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMongoDBResourceMongoRoleDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoDBResourceMongoRoleDefinitionState, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoRoleDefinition, error) {
	var resource MongoDBResourceMongoRoleDefinition
	err := ctx.ReadResource("azure-native:documentdb:MongoDBResourceMongoRoleDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MongoDBResourceMongoRoleDefinition resources.
type mongoDBResourceMongoRoleDefinitionState struct {
}

type MongoDBResourceMongoRoleDefinitionState struct {
}

func (MongoDBResourceMongoRoleDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoRoleDefinitionState)(nil)).Elem()
}

type mongoDBResourceMongoRoleDefinitionArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The database name for which access is being granted for this Role Definition.
	DatabaseName *string `pulumi:"databaseName"`
	// The ID for the Role Definition {dbName.roleName}.
	MongoRoleDefinitionId *string `pulumi:"mongoRoleDefinitionId"`
	// A set of privileges contained by the Role Definition. This will allow application of this Role Definition on the entire database account or any underlying Database / Collection. Scopes higher than Database are not enforceable as privilege.
	Privileges []Privilege `pulumi:"privileges"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A user-friendly name for the Role Definition. Must be unique for the database account.
	RoleName *string `pulumi:"roleName"`
	// The set of roles inherited by this Role Definition.
	Roles []Role `pulumi:"roles"`
	// Indicates whether the Role Definition was built-in or user created.
	Type *MongoRoleDefinitionType `pulumi:"type"`
}

// The set of arguments for constructing a MongoDBResourceMongoRoleDefinition resource.
type MongoDBResourceMongoRoleDefinitionArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// The database name for which access is being granted for this Role Definition.
	DatabaseName pulumi.StringPtrInput
	// The ID for the Role Definition {dbName.roleName}.
	MongoRoleDefinitionId pulumi.StringPtrInput
	// A set of privileges contained by the Role Definition. This will allow application of this Role Definition on the entire database account or any underlying Database / Collection. Scopes higher than Database are not enforceable as privilege.
	Privileges PrivilegeArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// A user-friendly name for the Role Definition. Must be unique for the database account.
	RoleName pulumi.StringPtrInput
	// The set of roles inherited by this Role Definition.
	Roles RoleArrayInput
	// Indicates whether the Role Definition was built-in or user created.
	Type MongoRoleDefinitionTypePtrInput
}

func (MongoDBResourceMongoRoleDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoRoleDefinitionArgs)(nil)).Elem()
}

type MongoDBResourceMongoRoleDefinitionInput interface {
	pulumi.Input

	ToMongoDBResourceMongoRoleDefinitionOutput() MongoDBResourceMongoRoleDefinitionOutput
	ToMongoDBResourceMongoRoleDefinitionOutputWithContext(ctx context.Context) MongoDBResourceMongoRoleDefinitionOutput
}

func (*MongoDBResourceMongoRoleDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBResourceMongoRoleDefinition)(nil)).Elem()
}

func (i *MongoDBResourceMongoRoleDefinition) ToMongoDBResourceMongoRoleDefinitionOutput() MongoDBResourceMongoRoleDefinitionOutput {
	return i.ToMongoDBResourceMongoRoleDefinitionOutputWithContext(context.Background())
}

func (i *MongoDBResourceMongoRoleDefinition) ToMongoDBResourceMongoRoleDefinitionOutputWithContext(ctx context.Context) MongoDBResourceMongoRoleDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBResourceMongoRoleDefinitionOutput)
}

type MongoDBResourceMongoRoleDefinitionOutput struct{ *pulumi.OutputState }

func (MongoDBResourceMongoRoleDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBResourceMongoRoleDefinition)(nil)).Elem()
}

func (o MongoDBResourceMongoRoleDefinitionOutput) ToMongoDBResourceMongoRoleDefinitionOutput() MongoDBResourceMongoRoleDefinitionOutput {
	return o
}

func (o MongoDBResourceMongoRoleDefinitionOutput) ToMongoDBResourceMongoRoleDefinitionOutputWithContext(ctx context.Context) MongoDBResourceMongoRoleDefinitionOutput {
	return o
}

// The database name for which access is being granted for this Role Definition.
func (o MongoDBResourceMongoRoleDefinitionOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoRoleDefinition) pulumi.StringPtrOutput { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// The name of the database account.
func (o MongoDBResourceMongoRoleDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoRoleDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A set of privileges contained by the Role Definition. This will allow application of this Role Definition on the entire database account or any underlying Database / Collection. Scopes higher than Database are not enforceable as privilege.
func (o MongoDBResourceMongoRoleDefinitionOutput) Privileges() PrivilegeResponseArrayOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoRoleDefinition) PrivilegeResponseArrayOutput { return v.Privileges }).(PrivilegeResponseArrayOutput)
}

// A user-friendly name for the Role Definition. Must be unique for the database account.
func (o MongoDBResourceMongoRoleDefinitionOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoRoleDefinition) pulumi.StringPtrOutput { return v.RoleName }).(pulumi.StringPtrOutput)
}

// The set of roles inherited by this Role Definition.
func (o MongoDBResourceMongoRoleDefinitionOutput) Roles() RoleResponseArrayOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoRoleDefinition) RoleResponseArrayOutput { return v.Roles }).(RoleResponseArrayOutput)
}

// The type of Azure resource.
func (o MongoDBResourceMongoRoleDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoRoleDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MongoDBResourceMongoRoleDefinitionOutput{})
}
