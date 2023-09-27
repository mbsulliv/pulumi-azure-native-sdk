// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230415

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the SQL userDefinedFunction under an existing Azure Cosmos DB database account.
func LookupSqlResourceSqlUserDefinedFunction(ctx *pulumi.Context, args *LookupSqlResourceSqlUserDefinedFunctionArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlUserDefinedFunctionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlResourceSqlUserDefinedFunctionResult
	err := ctx.Invoke("azure-native:documentdb/v20230415:getSqlResourceSqlUserDefinedFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlUserDefinedFunctionArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB container name.
	ContainerName string `pulumi:"containerName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cosmos DB userDefinedFunction name.
	UserDefinedFunctionName string `pulumi:"userDefinedFunctionName"`
}

// An Azure Cosmos DB userDefinedFunction.
type LookupSqlResourceSqlUserDefinedFunctionResult struct {
	// The unique resource identifier of the ARM resource.
	Id string `pulumi:"id"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     string                                               `pulumi:"name"`
	Resource *SqlUserDefinedFunctionGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupSqlResourceSqlUserDefinedFunctionOutput(ctx *pulumi.Context, args LookupSqlResourceSqlUserDefinedFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupSqlResourceSqlUserDefinedFunctionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlResourceSqlUserDefinedFunctionResult, error) {
			args := v.(LookupSqlResourceSqlUserDefinedFunctionArgs)
			r, err := LookupSqlResourceSqlUserDefinedFunction(ctx, &args, opts...)
			var s LookupSqlResourceSqlUserDefinedFunctionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlResourceSqlUserDefinedFunctionResultOutput)
}

type LookupSqlResourceSqlUserDefinedFunctionOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Cosmos DB container name.
	ContainerName pulumi.StringInput `pulumi:"containerName"`
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Cosmos DB userDefinedFunction name.
	UserDefinedFunctionName pulumi.StringInput `pulumi:"userDefinedFunctionName"`
}

func (LookupSqlResourceSqlUserDefinedFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlUserDefinedFunctionArgs)(nil)).Elem()
}

// An Azure Cosmos DB userDefinedFunction.
type LookupSqlResourceSqlUserDefinedFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupSqlResourceSqlUserDefinedFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlUserDefinedFunctionResult)(nil)).Elem()
}

func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) ToLookupSqlResourceSqlUserDefinedFunctionResultOutput() LookupSqlResourceSqlUserDefinedFunctionResultOutput {
	return o
}

func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) ToLookupSqlResourceSqlUserDefinedFunctionResultOutputWithContext(ctx context.Context) LookupSqlResourceSqlUserDefinedFunctionResultOutput {
	return o
}

func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSqlResourceSqlUserDefinedFunctionResult] {
	return pulumix.Output[LookupSqlResourceSqlUserDefinedFunctionResult]{
		OutputState: o.OutputState,
	}
}

// The unique resource identifier of the ARM resource.
func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlUserDefinedFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource group to which the resource belongs.
func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlUserDefinedFunctionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlUserDefinedFunctionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) Resource() SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlUserDefinedFunctionResult) *SqlUserDefinedFunctionGetPropertiesResponseResource {
		return v.Resource
	}).(SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlUserDefinedFunctionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o LookupSqlResourceSqlUserDefinedFunctionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlUserDefinedFunctionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlResourceSqlUserDefinedFunctionResultOutput{})
}
