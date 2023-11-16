// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230415

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the SQL container under an existing Azure Cosmos DB database account.
func LookupSqlResourceSqlContainer(ctx *pulumi.Context, args *LookupSqlResourceSqlContainerArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlContainerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlResourceSqlContainerResult
	err := ctx.Invoke("azure-native:documentdb/v20230415:getSqlResourceSqlContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSqlResourceSqlContainerArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB container name.
	ContainerName string `pulumi:"containerName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Cosmos DB container.
type LookupSqlResourceSqlContainerResult struct {
	// The unique resource identifier of the ARM resource.
	Id string `pulumi:"id"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     string                                     `pulumi:"name"`
	Options  *SqlContainerGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *SqlContainerGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupSqlResourceSqlContainerResult
func (val *LookupSqlResourceSqlContainerResult) Defaults() *LookupSqlResourceSqlContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Resource = tmp.Resource.Defaults()

	return &tmp
}

func LookupSqlResourceSqlContainerOutput(ctx *pulumi.Context, args LookupSqlResourceSqlContainerOutputArgs, opts ...pulumi.InvokeOption) LookupSqlResourceSqlContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlResourceSqlContainerResult, error) {
			args := v.(LookupSqlResourceSqlContainerArgs)
			r, err := LookupSqlResourceSqlContainer(ctx, &args, opts...)
			var s LookupSqlResourceSqlContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlResourceSqlContainerResultOutput)
}

type LookupSqlResourceSqlContainerOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Cosmos DB container name.
	ContainerName pulumi.StringInput `pulumi:"containerName"`
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSqlResourceSqlContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlContainerArgs)(nil)).Elem()
}

// An Azure Cosmos DB container.
type LookupSqlResourceSqlContainerResultOutput struct{ *pulumi.OutputState }

func (LookupSqlResourceSqlContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlContainerResult)(nil)).Elem()
}

func (o LookupSqlResourceSqlContainerResultOutput) ToLookupSqlResourceSqlContainerResultOutput() LookupSqlResourceSqlContainerResultOutput {
	return o
}

func (o LookupSqlResourceSqlContainerResultOutput) ToLookupSqlResourceSqlContainerResultOutputWithContext(ctx context.Context) LookupSqlResourceSqlContainerResultOutput {
	return o
}

// The unique resource identifier of the ARM resource.
func (o LookupSqlResourceSqlContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource group to which the resource belongs.
func (o LookupSqlResourceSqlContainerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o LookupSqlResourceSqlContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlContainerResultOutput) Options() SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) *SqlContainerGetPropertiesResponseOptions {
		return v.Options
	}).(SqlContainerGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupSqlResourceSqlContainerResultOutput) Resource() SqlContainerGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) *SqlContainerGetPropertiesResponseResource {
		return v.Resource
	}).(SqlContainerGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o LookupSqlResourceSqlContainerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o LookupSqlResourceSqlContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlResourceSqlContainerResultOutput{})
}
