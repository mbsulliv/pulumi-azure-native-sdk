// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230315preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the MongoDB databases under an existing Azure Cosmos DB database account with the provided name.
func LookupMongoDBResourceMongoDBDatabase(ctx *pulumi.Context, args *LookupMongoDBResourceMongoDBDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupMongoDBResourceMongoDBDatabaseResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMongoDBResourceMongoDBDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20230315preview:getMongoDBResourceMongoDBDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoDBResourceMongoDBDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Cosmos DB MongoDB database.
type LookupMongoDBResourceMongoDBDatabaseResult struct {
	// The unique resource identifier of the ARM resource.
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     string                                        `pulumi:"name"`
	Options  *MongoDBDatabaseGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *MongoDBDatabaseGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupMongoDBResourceMongoDBDatabaseOutput(ctx *pulumi.Context, args LookupMongoDBResourceMongoDBDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupMongoDBResourceMongoDBDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMongoDBResourceMongoDBDatabaseResult, error) {
			args := v.(LookupMongoDBResourceMongoDBDatabaseArgs)
			r, err := LookupMongoDBResourceMongoDBDatabase(ctx, &args, opts...)
			var s LookupMongoDBResourceMongoDBDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMongoDBResourceMongoDBDatabaseResultOutput)
}

type LookupMongoDBResourceMongoDBDatabaseOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMongoDBResourceMongoDBDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoDBDatabaseArgs)(nil)).Elem()
}

// An Azure Cosmos DB MongoDB database.
type LookupMongoDBResourceMongoDBDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupMongoDBResourceMongoDBDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoDBResourceMongoDBDatabaseResult)(nil)).Elem()
}

func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) ToLookupMongoDBResourceMongoDBDatabaseResultOutput() LookupMongoDBResourceMongoDBDatabaseResultOutput {
	return o
}

func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) ToLookupMongoDBResourceMongoDBDatabaseResultOutputWithContext(ctx context.Context) LookupMongoDBResourceMongoDBDatabaseResultOutput {
	return o
}

// The unique resource identifier of the ARM resource.
func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The location of the resource group to which the resource belongs.
func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Options() MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) *MongoDBDatabaseGetPropertiesResponseOptions {
		return v.Options
	}).(MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Resource() MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) *MongoDBDatabaseGetPropertiesResponseResource {
		return v.Resource
	}).(MongoDBDatabaseGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o LookupMongoDBResourceMongoDBDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoDBResourceMongoDBDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMongoDBResourceMongoDBDatabaseResultOutput{})
}
