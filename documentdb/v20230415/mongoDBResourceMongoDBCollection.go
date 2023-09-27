// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230415

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An Azure Cosmos DB MongoDB collection.
type MongoDBResourceMongoDBCollection struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                     `pulumi:"name"`
	Options  MongoDBCollectionGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource MongoDBCollectionGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMongoDBResourceMongoDBCollection registers a new resource with the given unique name, arguments, and options.
func NewMongoDBResourceMongoDBCollection(ctx *pulumi.Context,
	name string, args *MongoDBResourceMongoDBCollectionArgs, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoDBCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230301preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:MongoDBResourceMongoDBCollection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MongoDBResourceMongoDBCollection
	err := ctx.RegisterResource("azure-native:documentdb/v20230415:MongoDBResourceMongoDBCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMongoDBResourceMongoDBCollection gets an existing MongoDBResourceMongoDBCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMongoDBResourceMongoDBCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoDBResourceMongoDBCollectionState, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoDBCollection, error) {
	var resource MongoDBResourceMongoDBCollection
	err := ctx.ReadResource("azure-native:documentdb/v20230415:MongoDBResourceMongoDBCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MongoDBResourceMongoDBCollection resources.
type mongoDBResourceMongoDBCollectionState struct {
}

type MongoDBResourceMongoDBCollectionState struct {
}

func (MongoDBResourceMongoDBCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoDBCollectionState)(nil)).Elem()
}

type mongoDBResourceMongoDBCollectionArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB collection name.
	CollectionName *string `pulumi:"collectionName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options *CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a MongoDB collection
	Resource MongoDBCollectionResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MongoDBResourceMongoDBCollection resource.
type MongoDBResourceMongoDBCollectionArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB collection name.
	CollectionName pulumi.StringPtrInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsPtrInput
	// The standard JSON format of a MongoDB collection
	Resource MongoDBCollectionResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (MongoDBResourceMongoDBCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoDBCollectionArgs)(nil)).Elem()
}

type MongoDBResourceMongoDBCollectionInput interface {
	pulumi.Input

	ToMongoDBResourceMongoDBCollectionOutput() MongoDBResourceMongoDBCollectionOutput
	ToMongoDBResourceMongoDBCollectionOutputWithContext(ctx context.Context) MongoDBResourceMongoDBCollectionOutput
}

func (*MongoDBResourceMongoDBCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBResourceMongoDBCollection)(nil)).Elem()
}

func (i *MongoDBResourceMongoDBCollection) ToMongoDBResourceMongoDBCollectionOutput() MongoDBResourceMongoDBCollectionOutput {
	return i.ToMongoDBResourceMongoDBCollectionOutputWithContext(context.Background())
}

func (i *MongoDBResourceMongoDBCollection) ToMongoDBResourceMongoDBCollectionOutputWithContext(ctx context.Context) MongoDBResourceMongoDBCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBResourceMongoDBCollectionOutput)
}

func (i *MongoDBResourceMongoDBCollection) ToOutput(ctx context.Context) pulumix.Output[*MongoDBResourceMongoDBCollection] {
	return pulumix.Output[*MongoDBResourceMongoDBCollection]{
		OutputState: i.ToMongoDBResourceMongoDBCollectionOutputWithContext(ctx).OutputState,
	}
}

type MongoDBResourceMongoDBCollectionOutput struct{ *pulumi.OutputState }

func (MongoDBResourceMongoDBCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBResourceMongoDBCollection)(nil)).Elem()
}

func (o MongoDBResourceMongoDBCollectionOutput) ToMongoDBResourceMongoDBCollectionOutput() MongoDBResourceMongoDBCollectionOutput {
	return o
}

func (o MongoDBResourceMongoDBCollectionOutput) ToMongoDBResourceMongoDBCollectionOutputWithContext(ctx context.Context) MongoDBResourceMongoDBCollectionOutput {
	return o
}

func (o MongoDBResourceMongoDBCollectionOutput) ToOutput(ctx context.Context) pulumix.Output[*MongoDBResourceMongoDBCollection] {
	return pulumix.Output[*MongoDBResourceMongoDBCollection]{
		OutputState: o.OutputState,
	}
}

// The location of the resource group to which the resource belongs.
func (o MongoDBResourceMongoDBCollectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoDBCollection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o MongoDBResourceMongoDBCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoDBCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MongoDBResourceMongoDBCollectionOutput) Options() MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoDBCollection) MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
		return v.Options
	}).(MongoDBCollectionGetPropertiesResponseOptionsPtrOutput)
}

func (o MongoDBResourceMongoDBCollectionOutput) Resource() MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoDBCollection) MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
		return v.Resource
	}).(MongoDBCollectionGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o MongoDBResourceMongoDBCollectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoDBCollection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o MongoDBResourceMongoDBCollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoDBCollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MongoDBResourceMongoDBCollectionOutput{})
}
