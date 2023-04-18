// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160319

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB MongoDB database.
//
// Deprecated: Version 2016-03-19 will be removed in v2 of the provider.
type DatabaseAccountMongoDBDatabase struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabaseAccountMongoDBDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccountMongoDBDatabase(ctx *pulumi.Context,
	name string, args *DatabaseAccountMongoDBDatabaseArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountMongoDBDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:DatabaseAccountMongoDBDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountMongoDBDatabase
	err := ctx.RegisterResource("azure-native:documentdb/v20160319:DatabaseAccountMongoDBDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAccountMongoDBDatabase gets an existing DatabaseAccountMongoDBDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAccountMongoDBDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountMongoDBDatabaseState, opts ...pulumi.ResourceOption) (*DatabaseAccountMongoDBDatabase, error) {
	var resource DatabaseAccountMongoDBDatabase
	err := ctx.ReadResource("azure-native:documentdb/v20160319:DatabaseAccountMongoDBDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccountMongoDBDatabase resources.
type databaseAccountMongoDBDatabaseState struct {
}

type DatabaseAccountMongoDBDatabaseState struct {
}

func (DatabaseAccountMongoDBDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBDatabaseState)(nil)).Elem()
}

type databaseAccountMongoDBDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName *string `pulumi:"databaseName"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options map[string]string `pulumi:"options"`
	// The standard JSON format of a MongoDB database
	Resource MongoDBDatabaseResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DatabaseAccountMongoDBDatabase resource.
type DatabaseAccountMongoDBDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options pulumi.StringMapInput
	// The standard JSON format of a MongoDB database
	Resource MongoDBDatabaseResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountMongoDBDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBDatabaseArgs)(nil)).Elem()
}

type DatabaseAccountMongoDBDatabaseInput interface {
	pulumi.Input

	ToDatabaseAccountMongoDBDatabaseOutput() DatabaseAccountMongoDBDatabaseOutput
	ToDatabaseAccountMongoDBDatabaseOutputWithContext(ctx context.Context) DatabaseAccountMongoDBDatabaseOutput
}

func (*DatabaseAccountMongoDBDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountMongoDBDatabase)(nil)).Elem()
}

func (i *DatabaseAccountMongoDBDatabase) ToDatabaseAccountMongoDBDatabaseOutput() DatabaseAccountMongoDBDatabaseOutput {
	return i.ToDatabaseAccountMongoDBDatabaseOutputWithContext(context.Background())
}

func (i *DatabaseAccountMongoDBDatabase) ToDatabaseAccountMongoDBDatabaseOutputWithContext(ctx context.Context) DatabaseAccountMongoDBDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountMongoDBDatabaseOutput)
}

type DatabaseAccountMongoDBDatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseAccountMongoDBDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountMongoDBDatabase)(nil)).Elem()
}

func (o DatabaseAccountMongoDBDatabaseOutput) ToDatabaseAccountMongoDBDatabaseOutput() DatabaseAccountMongoDBDatabaseOutput {
	return o
}

func (o DatabaseAccountMongoDBDatabaseOutput) ToDatabaseAccountMongoDBDatabaseOutputWithContext(ctx context.Context) DatabaseAccountMongoDBDatabaseOutput {
	return o
}

// The location of the resource group to which the resource belongs.
func (o DatabaseAccountMongoDBDatabaseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBDatabase) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the database account.
func (o DatabaseAccountMongoDBDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o DatabaseAccountMongoDBDatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBDatabase) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o DatabaseAccountMongoDBDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountMongoDBDatabaseOutput{})
}
