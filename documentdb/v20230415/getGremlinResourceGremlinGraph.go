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

// Gets the Gremlin graph under an existing Azure Cosmos DB database account.
func LookupGremlinResourceGremlinGraph(ctx *pulumi.Context, args *LookupGremlinResourceGremlinGraphArgs, opts ...pulumi.InvokeOption) (*LookupGremlinResourceGremlinGraphResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGremlinResourceGremlinGraphResult
	err := ctx.Invoke("azure-native:documentdb/v20230415:getGremlinResourceGremlinGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupGremlinResourceGremlinGraphArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// Cosmos DB graph name.
	GraphName string `pulumi:"graphName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Cosmos DB Gremlin graph.
type LookupGremlinResourceGremlinGraphResult struct {
	// The unique resource identifier of the ARM resource.
	Id string `pulumi:"id"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     string                                     `pulumi:"name"`
	Options  *GremlinGraphGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *GremlinGraphGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupGremlinResourceGremlinGraphResult
func (val *LookupGremlinResourceGremlinGraphResult) Defaults() *LookupGremlinResourceGremlinGraphResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Resource = tmp.Resource.Defaults()

	return &tmp
}

func LookupGremlinResourceGremlinGraphOutput(ctx *pulumi.Context, args LookupGremlinResourceGremlinGraphOutputArgs, opts ...pulumi.InvokeOption) LookupGremlinResourceGremlinGraphResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGremlinResourceGremlinGraphResult, error) {
			args := v.(LookupGremlinResourceGremlinGraphArgs)
			r, err := LookupGremlinResourceGremlinGraph(ctx, &args, opts...)
			var s LookupGremlinResourceGremlinGraphResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGremlinResourceGremlinGraphResultOutput)
}

type LookupGremlinResourceGremlinGraphOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// Cosmos DB graph name.
	GraphName pulumi.StringInput `pulumi:"graphName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGremlinResourceGremlinGraphOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGremlinResourceGremlinGraphArgs)(nil)).Elem()
}

// An Azure Cosmos DB Gremlin graph.
type LookupGremlinResourceGremlinGraphResultOutput struct{ *pulumi.OutputState }

func (LookupGremlinResourceGremlinGraphResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGremlinResourceGremlinGraphResult)(nil)).Elem()
}

func (o LookupGremlinResourceGremlinGraphResultOutput) ToLookupGremlinResourceGremlinGraphResultOutput() LookupGremlinResourceGremlinGraphResultOutput {
	return o
}

func (o LookupGremlinResourceGremlinGraphResultOutput) ToLookupGremlinResourceGremlinGraphResultOutputWithContext(ctx context.Context) LookupGremlinResourceGremlinGraphResultOutput {
	return o
}

func (o LookupGremlinResourceGremlinGraphResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGremlinResourceGremlinGraphResult] {
	return pulumix.Output[LookupGremlinResourceGremlinGraphResult]{
		OutputState: o.OutputState,
	}
}

// The unique resource identifier of the ARM resource.
func (o LookupGremlinResourceGremlinGraphResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinGraphResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource group to which the resource belongs.
func (o LookupGremlinResourceGremlinGraphResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinGraphResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o LookupGremlinResourceGremlinGraphResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinGraphResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGremlinResourceGremlinGraphResultOutput) Options() GremlinGraphGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinGraphResult) *GremlinGraphGetPropertiesResponseOptions {
		return v.Options
	}).(GremlinGraphGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupGremlinResourceGremlinGraphResultOutput) Resource() GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinGraphResult) *GremlinGraphGetPropertiesResponseResource {
		return v.Resource
	}).(GremlinGraphGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o LookupGremlinResourceGremlinGraphResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinGraphResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o LookupGremlinResourceGremlinGraphResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinGraphResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGremlinResourceGremlinGraphResultOutput{})
}
