// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230915preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the Tables under an existing Azure Cosmos DB database account with the provided name.
func LookupTableResourceTable(ctx *pulumi.Context, args *LookupTableResourceTableArgs, opts ...pulumi.InvokeOption) (*LookupTableResourceTableResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTableResourceTableResult
	err := ctx.Invoke("azure-native:documentdb/v20230915preview:getTableResourceTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTableResourceTableArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cosmos DB table name.
	TableName string `pulumi:"tableName"`
}

// An Azure Cosmos DB Table.
type LookupTableResourceTableResult struct {
	// The unique resource identifier of the ARM resource.
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     string                              `pulumi:"name"`
	Options  *TableGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *TableGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupTableResourceTableOutput(ctx *pulumi.Context, args LookupTableResourceTableOutputArgs, opts ...pulumi.InvokeOption) LookupTableResourceTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTableResourceTableResult, error) {
			args := v.(LookupTableResourceTableArgs)
			r, err := LookupTableResourceTable(ctx, &args, opts...)
			var s LookupTableResourceTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTableResourceTableResultOutput)
}

type LookupTableResourceTableOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Cosmos DB table name.
	TableName pulumi.StringInput `pulumi:"tableName"`
}

func (LookupTableResourceTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableResourceTableArgs)(nil)).Elem()
}

// An Azure Cosmos DB Table.
type LookupTableResourceTableResultOutput struct{ *pulumi.OutputState }

func (LookupTableResourceTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableResourceTableResult)(nil)).Elem()
}

func (o LookupTableResourceTableResultOutput) ToLookupTableResourceTableResultOutput() LookupTableResourceTableResultOutput {
	return o
}

func (o LookupTableResourceTableResultOutput) ToLookupTableResourceTableResultOutputWithContext(ctx context.Context) LookupTableResourceTableResultOutput {
	return o
}

// The unique resource identifier of the ARM resource.
func (o LookupTableResourceTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupTableResourceTableResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The location of the resource group to which the resource belongs.
func (o LookupTableResourceTableResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o LookupTableResourceTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTableResourceTableResultOutput) Options() TableGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) *TableGetPropertiesResponseOptions { return v.Options }).(TableGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupTableResourceTableResultOutput) Resource() TableGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) *TableGetPropertiesResponseResource { return v.Resource }).(TableGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o LookupTableResourceTableResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o LookupTableResourceTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTableResourceTableResultOutput{})
}
