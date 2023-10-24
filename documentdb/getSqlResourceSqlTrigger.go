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

// Gets the SQL trigger under an existing Azure Cosmos DB database account.
// Azure REST API version: 2023-04-15.
//
// Other available API versions: 2019-08-01, 2023-03-15-preview, 2023-09-15, 2023-09-15-preview.
func LookupSqlResourceSqlTrigger(ctx *pulumi.Context, args *LookupSqlResourceSqlTriggerArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlTriggerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlResourceSqlTriggerResult
	err := ctx.Invoke("azure-native:documentdb:getSqlResourceSqlTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlTriggerArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB container name.
	ContainerName string `pulumi:"containerName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cosmos DB trigger name.
	TriggerName string `pulumi:"triggerName"`
}

// An Azure Cosmos DB trigger.
type LookupSqlResourceSqlTriggerResult struct {
	// The unique resource identifier of the ARM resource.
	Id string `pulumi:"id"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     string                                   `pulumi:"name"`
	Resource *SqlTriggerGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupSqlResourceSqlTriggerOutput(ctx *pulumi.Context, args LookupSqlResourceSqlTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupSqlResourceSqlTriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlResourceSqlTriggerResult, error) {
			args := v.(LookupSqlResourceSqlTriggerArgs)
			r, err := LookupSqlResourceSqlTrigger(ctx, &args, opts...)
			var s LookupSqlResourceSqlTriggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlResourceSqlTriggerResultOutput)
}

type LookupSqlResourceSqlTriggerOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Cosmos DB container name.
	ContainerName pulumi.StringInput `pulumi:"containerName"`
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Cosmos DB trigger name.
	TriggerName pulumi.StringInput `pulumi:"triggerName"`
}

func (LookupSqlResourceSqlTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlTriggerArgs)(nil)).Elem()
}

// An Azure Cosmos DB trigger.
type LookupSqlResourceSqlTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupSqlResourceSqlTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlTriggerResult)(nil)).Elem()
}

func (o LookupSqlResourceSqlTriggerResultOutput) ToLookupSqlResourceSqlTriggerResultOutput() LookupSqlResourceSqlTriggerResultOutput {
	return o
}

func (o LookupSqlResourceSqlTriggerResultOutput) ToLookupSqlResourceSqlTriggerResultOutputWithContext(ctx context.Context) LookupSqlResourceSqlTriggerResultOutput {
	return o
}

func (o LookupSqlResourceSqlTriggerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSqlResourceSqlTriggerResult] {
	return pulumix.Output[LookupSqlResourceSqlTriggerResult]{
		OutputState: o.OutputState,
	}
}

// The unique resource identifier of the ARM resource.
func (o LookupSqlResourceSqlTriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlTriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource group to which the resource belongs.
func (o LookupSqlResourceSqlTriggerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlTriggerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o LookupSqlResourceSqlTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlTriggerResultOutput) Resource() SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlTriggerResult) *SqlTriggerGetPropertiesResponseResource { return v.Resource }).(SqlTriggerGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o LookupSqlResourceSqlTriggerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlTriggerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o LookupSqlResourceSqlTriggerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlTriggerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlResourceSqlTriggerResultOutput{})
}
