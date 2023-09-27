// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcegraph

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a single graph query by its resourceName.
// Azure REST API version: 2020-04-01-preview.
func LookupGraphQuery(ctx *pulumi.Context, args *LookupGraphQueryArgs, opts ...pulumi.InvokeOption) (*LookupGraphQueryResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGraphQueryResult
	err := ctx.Invoke("azure-native:resourcegraph:getGraphQuery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGraphQueryArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Graph Query resource.
	ResourceName string `pulumi:"resourceName"`
}

// Graph Query entity definition.
type LookupGraphQueryResult struct {
	// The description of a graph query.
	Description *string `pulumi:"description"`
	// This will be used to handle Optimistic Concurrency.
	Etag *string `pulumi:"etag"`
	// Azure resource Id
	Id string `pulumi:"id"`
	// The location of the resource
	Location string `pulumi:"location"`
	// Azure resource name. This is GUID value. The display name should be assigned within properties field.
	Name string `pulumi:"name"`
	// KQL query that will be graph.
	Query string `pulumi:"query"`
	// Enum indicating a type of graph query.
	ResultKind string `pulumi:"resultKind"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this graph query definition.
	TimeModified string `pulumi:"timeModified"`
	// Azure resource type
	Type string `pulumi:"type"`
}

func LookupGraphQueryOutput(ctx *pulumi.Context, args LookupGraphQueryOutputArgs, opts ...pulumi.InvokeOption) LookupGraphQueryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGraphQueryResult, error) {
			args := v.(LookupGraphQueryArgs)
			r, err := LookupGraphQuery(ctx, &args, opts...)
			var s LookupGraphQueryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGraphQueryResultOutput)
}

type LookupGraphQueryOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Graph Query resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupGraphQueryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphQueryArgs)(nil)).Elem()
}

// Graph Query entity definition.
type LookupGraphQueryResultOutput struct{ *pulumi.OutputState }

func (LookupGraphQueryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphQueryResult)(nil)).Elem()
}

func (o LookupGraphQueryResultOutput) ToLookupGraphQueryResultOutput() LookupGraphQueryResultOutput {
	return o
}

func (o LookupGraphQueryResultOutput) ToLookupGraphQueryResultOutputWithContext(ctx context.Context) LookupGraphQueryResultOutput {
	return o
}

func (o LookupGraphQueryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGraphQueryResult] {
	return pulumix.Output[LookupGraphQueryResult]{
		OutputState: o.OutputState,
	}
}

// The description of a graph query.
func (o LookupGraphQueryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// This will be used to handle Optimistic Concurrency.
func (o LookupGraphQueryResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Azure resource Id
func (o LookupGraphQueryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource
func (o LookupGraphQueryResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name. This is GUID value. The display name should be assigned within properties field.
func (o LookupGraphQueryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.Name }).(pulumi.StringOutput)
}

// KQL query that will be graph.
func (o LookupGraphQueryResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.Query }).(pulumi.StringOutput)
}

// Enum indicating a type of graph query.
func (o LookupGraphQueryResultOutput) ResultKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.ResultKind }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupGraphQueryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupGraphQueryResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Date and time in UTC of the last modification that was made to this graph query definition.
func (o LookupGraphQueryResultOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.TimeModified }).(pulumi.StringOutput)
}

// Azure resource type
func (o LookupGraphQueryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQueryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGraphQueryResultOutput{})
}
