// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220515preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB User Definition
type MongoDBResourceMongoUserDefinition struct {
	pulumi.CustomResourceState

	// A custom definition for the USer Definition.
	CustomData pulumi.StringPtrOutput `pulumi:"customData"`
	// The database name for which access is being granted for this User Definition.
	DatabaseName pulumi.StringPtrOutput `pulumi:"databaseName"`
	// The Mongo Auth mechanism. For now, we only support auth mechanism SCRAM-SHA-256.
	Mechanisms pulumi.StringPtrOutput `pulumi:"mechanisms"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password for User Definition. Response does not contain user password.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The set of roles inherited by the User Definition.
	Roles RoleResponseArrayOutput `pulumi:"roles"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The user name for User Definition.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
}

// NewMongoDBResourceMongoUserDefinition registers a new resource with the given unique name, arguments, and options.
func NewMongoDBResourceMongoUserDefinition(ctx *pulumi.Context,
	name string, args *MongoDBResourceMongoUserDefinitionArgs, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoUserDefinition, error) {
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
			Type: pulumi.String("azure-native:documentdb:MongoDBResourceMongoUserDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:MongoDBResourceMongoUserDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:MongoDBResourceMongoUserDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:MongoDBResourceMongoUserDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:MongoDBResourceMongoUserDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:MongoDBResourceMongoUserDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:MongoDBResourceMongoUserDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:MongoDBResourceMongoUserDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource MongoDBResourceMongoUserDefinition
	err := ctx.RegisterResource("azure-native:documentdb/v20220515preview:MongoDBResourceMongoUserDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMongoDBResourceMongoUserDefinition gets an existing MongoDBResourceMongoUserDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMongoDBResourceMongoUserDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoDBResourceMongoUserDefinitionState, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoUserDefinition, error) {
	var resource MongoDBResourceMongoUserDefinition
	err := ctx.ReadResource("azure-native:documentdb/v20220515preview:MongoDBResourceMongoUserDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MongoDBResourceMongoUserDefinition resources.
type mongoDBResourceMongoUserDefinitionState struct {
}

type MongoDBResourceMongoUserDefinitionState struct {
}

func (MongoDBResourceMongoUserDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoUserDefinitionState)(nil)).Elem()
}

type mongoDBResourceMongoUserDefinitionArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// A custom definition for the USer Definition.
	CustomData *string `pulumi:"customData"`
	// The database name for which access is being granted for this User Definition.
	DatabaseName *string `pulumi:"databaseName"`
	// The Mongo Auth mechanism. For now, we only support auth mechanism SCRAM-SHA-256.
	Mechanisms *string `pulumi:"mechanisms"`
	// The ID for the User Definition {dbName.userName}.
	MongoUserDefinitionId *string `pulumi:"mongoUserDefinitionId"`
	// The password for User Definition. Response does not contain user password.
	Password *string `pulumi:"password"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The set of roles inherited by the User Definition.
	Roles []Role `pulumi:"roles"`
	// The user name for User Definition.
	UserName *string `pulumi:"userName"`
}

// The set of arguments for constructing a MongoDBResourceMongoUserDefinition resource.
type MongoDBResourceMongoUserDefinitionArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// A custom definition for the USer Definition.
	CustomData pulumi.StringPtrInput
	// The database name for which access is being granted for this User Definition.
	DatabaseName pulumi.StringPtrInput
	// The Mongo Auth mechanism. For now, we only support auth mechanism SCRAM-SHA-256.
	Mechanisms pulumi.StringPtrInput
	// The ID for the User Definition {dbName.userName}.
	MongoUserDefinitionId pulumi.StringPtrInput
	// The password for User Definition. Response does not contain user password.
	Password pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The set of roles inherited by the User Definition.
	Roles RoleArrayInput
	// The user name for User Definition.
	UserName pulumi.StringPtrInput
}

func (MongoDBResourceMongoUserDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoUserDefinitionArgs)(nil)).Elem()
}

type MongoDBResourceMongoUserDefinitionInput interface {
	pulumi.Input

	ToMongoDBResourceMongoUserDefinitionOutput() MongoDBResourceMongoUserDefinitionOutput
	ToMongoDBResourceMongoUserDefinitionOutputWithContext(ctx context.Context) MongoDBResourceMongoUserDefinitionOutput
}

func (*MongoDBResourceMongoUserDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBResourceMongoUserDefinition)(nil)).Elem()
}

func (i *MongoDBResourceMongoUserDefinition) ToMongoDBResourceMongoUserDefinitionOutput() MongoDBResourceMongoUserDefinitionOutput {
	return i.ToMongoDBResourceMongoUserDefinitionOutputWithContext(context.Background())
}

func (i *MongoDBResourceMongoUserDefinition) ToMongoDBResourceMongoUserDefinitionOutputWithContext(ctx context.Context) MongoDBResourceMongoUserDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBResourceMongoUserDefinitionOutput)
}

type MongoDBResourceMongoUserDefinitionOutput struct{ *pulumi.OutputState }

func (MongoDBResourceMongoUserDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBResourceMongoUserDefinition)(nil)).Elem()
}

func (o MongoDBResourceMongoUserDefinitionOutput) ToMongoDBResourceMongoUserDefinitionOutput() MongoDBResourceMongoUserDefinitionOutput {
	return o
}

func (o MongoDBResourceMongoUserDefinitionOutput) ToMongoDBResourceMongoUserDefinitionOutputWithContext(ctx context.Context) MongoDBResourceMongoUserDefinitionOutput {
	return o
}

// A custom definition for the USer Definition.
func (o MongoDBResourceMongoUserDefinitionOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoUserDefinition) pulumi.StringPtrOutput { return v.CustomData }).(pulumi.StringPtrOutput)
}

// The database name for which access is being granted for this User Definition.
func (o MongoDBResourceMongoUserDefinitionOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoUserDefinition) pulumi.StringPtrOutput { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// The Mongo Auth mechanism. For now, we only support auth mechanism SCRAM-SHA-256.
func (o MongoDBResourceMongoUserDefinitionOutput) Mechanisms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoUserDefinition) pulumi.StringPtrOutput { return v.Mechanisms }).(pulumi.StringPtrOutput)
}

// The name of the database account.
func (o MongoDBResourceMongoUserDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoUserDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The password for User Definition. Response does not contain user password.
func (o MongoDBResourceMongoUserDefinitionOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoUserDefinition) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The set of roles inherited by the User Definition.
func (o MongoDBResourceMongoUserDefinitionOutput) Roles() RoleResponseArrayOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoUserDefinition) RoleResponseArrayOutput { return v.Roles }).(RoleResponseArrayOutput)
}

// The type of Azure resource.
func (o MongoDBResourceMongoUserDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoUserDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The user name for User Definition.
func (o MongoDBResourceMongoUserDefinitionOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoUserDefinition) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MongoDBResourceMongoUserDefinitionOutput{})
}
