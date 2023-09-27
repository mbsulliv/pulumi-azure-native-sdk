// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package documentdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the MongoDB collection under an existing Azure Cosmos DB database account.
// Azure REST API version: 2023-04-15.
func LookupMongoDBResourceMongoDBCollection(ctx *pulumi.Context, args *LookupMongoDBResourceMongoDBCollectionArgs, opts ...pulumi.InvokeOption) (*LookupMongoDBResourceMongoDBCollectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMongoDBResourceMongoDBCollectionResult
	err := ctx.Invoke("azure-native:documentdb:getMongoDBResourceMongoDBCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoDBResourceMongoDBCollectionArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB collection name.
	CollectionName string `pulumi:"collectionName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Cosmos DB MongoDB collection.
type LookupMongoDBResourceMongoDBCollectionResult struct {
	// The unique resource identifier of the ARM resource.
	Id string `pulumi:"id"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     string                                          `pulumi:"name"`
	Options  *MongoDBCollectionGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *MongoDBCollectionGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupMongoDBResourceMongoDBCollectionOutput(ctx *pulumi.Context, args LookupMongoDBResourceMongoDBCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupMongoDBResourceMongoDBCollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMongoDBResourceMongoDBCollectionResult, error) {
			args := v.(LookupMongoDBResourceMongoDBCollectionArgs)
			r, err := LookupMongoDBResourceMongoDBCollection(ctx, &args, opts...)
			var s LookupMongoDBResourceMongoDBCollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMongoDBResourceMongoDBCollectionResultOutput)
}

type LookupMongoDBResourceMongoDBCollectionOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Cosmos DB collection name.
	CollectionName pulumi.StringInput `pulumi:"collectionName"`
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMongoDBResourceMongoDBCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoDBCollectionArgs)(nil)).Elem()
}

// An Azure Cosmos DB MongoDB collection.
type LookupMongoDBResourceMongoDBCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupMongoDBResourceMongoDBCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoDBCollectionResult)(nil)).Elem()
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) ToLookupMongoDBResourceMongoDBCollectionResultOutput() LookupMongoDBResourceMongoDBCollectionResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) ToLookupMongoDBResourceMongoDBCollectionResultOutputWithContext(ctx context.Context) LookupMongoDBResourceMongoDBCollectionResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMongoDBResourceMongoDBCollectionResult] {
	return pulumix.Output[LookupMongoDBResourceMongoDBCollectionResult]{
		OutputState: o.OutputState,
	}
}

// The unique resource identifier of the ARM resource.
func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource group to which the resource belongs.
func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Options() MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) *MongoDBCollectionGetPropertiesResponseOptions {
		return v.Options
	}).(MongoDBCollectionGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Resource() MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) *MongoDBCollectionGetPropertiesResponseResource {
		return v.Resource
	}).(MongoDBCollectionGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o LookupMongoDBResourceMongoDBCollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBCollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMongoDBResourceMongoDBCollectionResultOutput{})
}
