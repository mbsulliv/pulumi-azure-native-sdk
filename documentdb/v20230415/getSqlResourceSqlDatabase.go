// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230415

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the SQL database under an existing Azure Cosmos DB database account with the provided name.
func LookupSqlResourceSqlDatabase(ctx *pulumi.Context, args *LookupSqlResourceSqlDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlDatabaseResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlResourceSqlDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20230415:getSqlResourceSqlDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlDatabaseArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Cosmos DB SQL database.
type LookupSqlResourceSqlDatabaseResult struct {
	// The unique resource identifier of the ARM resource.
	Id string `pulumi:"id"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     string                                    `pulumi:"name"`
	Options  *SqlDatabaseGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *SqlDatabaseGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupSqlResourceSqlDatabaseOutput(ctx *pulumi.Context, args LookupSqlResourceSqlDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupSqlResourceSqlDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlResourceSqlDatabaseResult, error) {
			args := v.(LookupSqlResourceSqlDatabaseArgs)
			r, err := LookupSqlResourceSqlDatabase(ctx, &args, opts...)
			var s LookupSqlResourceSqlDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlResourceSqlDatabaseResultOutput)
}

type LookupSqlResourceSqlDatabaseOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSqlResourceSqlDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlDatabaseArgs)(nil)).Elem()
}

// An Azure Cosmos DB SQL database.
type LookupSqlResourceSqlDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupSqlResourceSqlDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlDatabaseResult)(nil)).Elem()
}

func (o LookupSqlResourceSqlDatabaseResultOutput) ToLookupSqlResourceSqlDatabaseResultOutput() LookupSqlResourceSqlDatabaseResultOutput {
	return o
}

func (o LookupSqlResourceSqlDatabaseResultOutput) ToLookupSqlResourceSqlDatabaseResultOutputWithContext(ctx context.Context) LookupSqlResourceSqlDatabaseResultOutput {
	return o
}

// The unique resource identifier of the ARM resource.
func (o LookupSqlResourceSqlDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource group to which the resource belongs.
func (o LookupSqlResourceSqlDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o LookupSqlResourceSqlDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlDatabaseResultOutput) Options() SqlDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) *SqlDatabaseGetPropertiesResponseOptions { return v.Options }).(SqlDatabaseGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupSqlResourceSqlDatabaseResultOutput) Resource() SqlDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) *SqlDatabaseGetPropertiesResponseResource {
		return v.Resource
	}).(SqlDatabaseGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o LookupSqlResourceSqlDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o LookupSqlResourceSqlDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlResourceSqlDatabaseResultOutput{})
}