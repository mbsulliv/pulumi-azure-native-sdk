// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230915preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the Gremlin databases under an existing Azure Cosmos DB database account with the provided name.
func LookupGremlinResourceGremlinDatabase(ctx *pulumi.Context, args *LookupGremlinResourceGremlinDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupGremlinResourceGremlinDatabaseResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGremlinResourceGremlinDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20230915preview:getGremlinResourceGremlinDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGremlinResourceGremlinDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Cosmos DB Gremlin database.
type LookupGremlinResourceGremlinDatabaseResult struct {
	// The unique resource identifier of the ARM resource.
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     string                                        `pulumi:"name"`
	Options  *GremlinDatabaseGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *GremlinDatabaseGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupGremlinResourceGremlinDatabaseOutput(ctx *pulumi.Context, args LookupGremlinResourceGremlinDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupGremlinResourceGremlinDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGremlinResourceGremlinDatabaseResult, error) {
			args := v.(LookupGremlinResourceGremlinDatabaseArgs)
			r, err := LookupGremlinResourceGremlinDatabase(ctx, &args, opts...)
			var s LookupGremlinResourceGremlinDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGremlinResourceGremlinDatabaseResultOutput)
}

type LookupGremlinResourceGremlinDatabaseOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGremlinResourceGremlinDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGremlinResourceGremlinDatabaseArgs)(nil)).Elem()
}

// An Azure Cosmos DB Gremlin database.
type LookupGremlinResourceGremlinDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupGremlinResourceGremlinDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGremlinResourceGremlinDatabaseResult)(nil)).Elem()
}

func (o LookupGremlinResourceGremlinDatabaseResultOutput) ToLookupGremlinResourceGremlinDatabaseResultOutput() LookupGremlinResourceGremlinDatabaseResultOutput {
	return o
}

func (o LookupGremlinResourceGremlinDatabaseResultOutput) ToLookupGremlinResourceGremlinDatabaseResultOutputWithContext(ctx context.Context) LookupGremlinResourceGremlinDatabaseResultOutput {
	return o
}

// The unique resource identifier of the ARM resource.
func (o LookupGremlinResourceGremlinDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupGremlinResourceGremlinDatabaseResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The location of the resource group to which the resource belongs.
func (o LookupGremlinResourceGremlinDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o LookupGremlinResourceGremlinDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGremlinResourceGremlinDatabaseResultOutput) Options() GremlinDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) *GremlinDatabaseGetPropertiesResponseOptions {
		return v.Options
	}).(GremlinDatabaseGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupGremlinResourceGremlinDatabaseResultOutput) Resource() GremlinDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) *GremlinDatabaseGetPropertiesResponseResource {
		return v.Resource
	}).(GremlinDatabaseGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o LookupGremlinResourceGremlinDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o LookupGremlinResourceGremlinDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGremlinResourceGremlinDatabaseResultOutput{})
}
