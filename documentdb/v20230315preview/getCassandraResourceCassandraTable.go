// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230315preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the Cassandra table under an existing Azure Cosmos DB database account.
func LookupCassandraResourceCassandraTable(ctx *pulumi.Context, args *LookupCassandraResourceCassandraTableArgs, opts ...pulumi.InvokeOption) (*LookupCassandraResourceCassandraTableResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCassandraResourceCassandraTableResult
	err := ctx.Invoke("azure-native:documentdb/v20230315preview:getCassandraResourceCassandraTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCassandraResourceCassandraTableArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB keyspace name.
	KeyspaceName string `pulumi:"keyspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cosmos DB table name.
	TableName string `pulumi:"tableName"`
}

// An Azure Cosmos DB Cassandra table.
type LookupCassandraResourceCassandraTableResult struct {
	// The unique resource identifier of the ARM resource.
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     string                                       `pulumi:"name"`
	Options  *CassandraTableGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *CassandraTableGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupCassandraResourceCassandraTableOutput(ctx *pulumi.Context, args LookupCassandraResourceCassandraTableOutputArgs, opts ...pulumi.InvokeOption) LookupCassandraResourceCassandraTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCassandraResourceCassandraTableResult, error) {
			args := v.(LookupCassandraResourceCassandraTableArgs)
			r, err := LookupCassandraResourceCassandraTable(ctx, &args, opts...)
			var s LookupCassandraResourceCassandraTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCassandraResourceCassandraTableResultOutput)
}

type LookupCassandraResourceCassandraTableOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Cosmos DB keyspace name.
	KeyspaceName pulumi.StringInput `pulumi:"keyspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Cosmos DB table name.
	TableName pulumi.StringInput `pulumi:"tableName"`
}

func (LookupCassandraResourceCassandraTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraResourceCassandraTableArgs)(nil)).Elem()
}

// An Azure Cosmos DB Cassandra table.
type LookupCassandraResourceCassandraTableResultOutput struct{ *pulumi.OutputState }

func (LookupCassandraResourceCassandraTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraResourceCassandraTableResult)(nil)).Elem()
}

func (o LookupCassandraResourceCassandraTableResultOutput) ToLookupCassandraResourceCassandraTableResultOutput() LookupCassandraResourceCassandraTableResultOutput {
	return o
}

func (o LookupCassandraResourceCassandraTableResultOutput) ToLookupCassandraResourceCassandraTableResultOutputWithContext(ctx context.Context) LookupCassandraResourceCassandraTableResultOutput {
	return o
}

func (o LookupCassandraResourceCassandraTableResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCassandraResourceCassandraTableResult] {
	return pulumix.Output[LookupCassandraResourceCassandraTableResult]{
		OutputState: o.OutputState,
	}
}

// The unique resource identifier of the ARM resource.
func (o LookupCassandraResourceCassandraTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupCassandraResourceCassandraTableResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The location of the resource group to which the resource belongs.
func (o LookupCassandraResourceCassandraTableResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o LookupCassandraResourceCassandraTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCassandraResourceCassandraTableResultOutput) Options() CassandraTableGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) *CassandraTableGetPropertiesResponseOptions {
		return v.Options
	}).(CassandraTableGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupCassandraResourceCassandraTableResultOutput) Resource() CassandraTableGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) *CassandraTableGetPropertiesResponseResource {
		return v.Resource
	}).(CassandraTableGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o LookupCassandraResourceCassandraTableResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o LookupCassandraResourceCassandraTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraResourceCassandraTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCassandraResourceCassandraTableResultOutput{})
}
