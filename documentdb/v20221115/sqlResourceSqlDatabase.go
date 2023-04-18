// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221115

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB SQL database.
type SqlResourceSqlDatabase struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                               `pulumi:"name"`
	Options  SqlDatabaseGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource SqlDatabaseGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlResourceSqlDatabase registers a new resource with the given unique name, arguments, and options.
func NewSqlResourceSqlDatabase(ctx *pulumi.Context,
	name string, args *SqlResourceSqlDatabaseArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:SqlResourceSqlDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlResourceSqlDatabase
	err := ctx.RegisterResource("azure-native:documentdb/v20221115:SqlResourceSqlDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlResourceSqlDatabase gets an existing SqlResourceSqlDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlResourceSqlDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlDatabaseState, opts ...pulumi.ResourceOption) (*SqlResourceSqlDatabase, error) {
	var resource SqlResourceSqlDatabase
	err := ctx.ReadResource("azure-native:documentdb/v20221115:SqlResourceSqlDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlResourceSqlDatabase resources.
type sqlResourceSqlDatabaseState struct {
}

type SqlResourceSqlDatabaseState struct {
}

func (SqlResourceSqlDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlDatabaseState)(nil)).Elem()
}

type sqlResourceSqlDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName *string `pulumi:"databaseName"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options *CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a SQL database
	Resource SqlDatabaseResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SqlResourceSqlDatabase resource.
type SqlResourceSqlDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringPtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsPtrInput
	// The standard JSON format of a SQL database
	Resource SqlDatabaseResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (SqlResourceSqlDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlDatabaseArgs)(nil)).Elem()
}

type SqlResourceSqlDatabaseInput interface {
	pulumi.Input

	ToSqlResourceSqlDatabaseOutput() SqlResourceSqlDatabaseOutput
	ToSqlResourceSqlDatabaseOutputWithContext(ctx context.Context) SqlResourceSqlDatabaseOutput
}

func (*SqlResourceSqlDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlDatabase)(nil)).Elem()
}

func (i *SqlResourceSqlDatabase) ToSqlResourceSqlDatabaseOutput() SqlResourceSqlDatabaseOutput {
	return i.ToSqlResourceSqlDatabaseOutputWithContext(context.Background())
}

func (i *SqlResourceSqlDatabase) ToSqlResourceSqlDatabaseOutputWithContext(ctx context.Context) SqlResourceSqlDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlDatabaseOutput)
}

type SqlResourceSqlDatabaseOutput struct{ *pulumi.OutputState }

func (SqlResourceSqlDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlDatabase)(nil)).Elem()
}

func (o SqlResourceSqlDatabaseOutput) ToSqlResourceSqlDatabaseOutput() SqlResourceSqlDatabaseOutput {
	return o
}

func (o SqlResourceSqlDatabaseOutput) ToSqlResourceSqlDatabaseOutputWithContext(ctx context.Context) SqlResourceSqlDatabaseOutput {
	return o
}

// The location of the resource group to which the resource belongs.
func (o SqlResourceSqlDatabaseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlDatabase) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o SqlResourceSqlDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlResourceSqlDatabaseOutput) Options() SqlDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlDatabase) SqlDatabaseGetPropertiesResponseOptionsPtrOutput { return v.Options }).(SqlDatabaseGetPropertiesResponseOptionsPtrOutput)
}

func (o SqlResourceSqlDatabaseOutput) Resource() SqlDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlDatabase) SqlDatabaseGetPropertiesResponseResourcePtrOutput { return v.Resource }).(SqlDatabaseGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o SqlResourceSqlDatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlResourceSqlDatabase) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o SqlResourceSqlDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlResourceSqlDatabaseOutput{})
}
