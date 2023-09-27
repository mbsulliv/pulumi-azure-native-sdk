// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Graph Query entity definition.
type GraphQuery struct {
	pulumi.CustomResourceState

	// The description of a graph query.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// This will be used to handle Optimistic Concurrency. If not present, it will always overwrite the existing resource without checking conflict.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The location of the resource
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Azure resource name. This is GUID value. The display name should be assigned within properties field.
	Name pulumi.StringOutput `pulumi:"name"`
	// KQL query that will be graph.
	Query pulumi.StringOutput `pulumi:"query"`
	// Enum indicating a type of graph query.
	ResultKind pulumi.StringOutput `pulumi:"resultKind"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this graph query definition.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGraphQuery registers a new resource with the given unique name, arguments, and options.
func NewGraphQuery(ctx *pulumi.Context,
	name string, args *GraphQueryArgs, opts ...pulumi.ResourceOption) (*GraphQuery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resourcegraph:GraphQuery"),
		},
		{
			Type: pulumi.String("azure-native:resourcegraph/v20200401preview:GraphQuery"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GraphQuery
	err := ctx.RegisterResource("azure-native:resourcegraph/v20180901preview:GraphQuery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGraphQuery gets an existing GraphQuery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGraphQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphQueryState, opts ...pulumi.ResourceOption) (*GraphQuery, error) {
	var resource GraphQuery
	err := ctx.ReadResource("azure-native:resourcegraph/v20180901preview:GraphQuery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GraphQuery resources.
type graphQueryState struct {
}

type GraphQueryState struct {
}

func (GraphQueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQueryState)(nil)).Elem()
}

type graphQueryArgs struct {
	// The description of a graph query.
	Description *string `pulumi:"description"`
	// The location of the resource
	Location *string `pulumi:"location"`
	// KQL query that will be graph.
	Query string `pulumi:"query"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Graph Query resource.
	ResourceName *string `pulumi:"resourceName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GraphQuery resource.
type GraphQueryArgs struct {
	// The description of a graph query.
	Description pulumi.StringPtrInput
	// The location of the resource
	Location pulumi.StringPtrInput
	// KQL query that will be graph.
	Query pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the Graph Query resource.
	ResourceName pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (GraphQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQueryArgs)(nil)).Elem()
}

type GraphQueryInput interface {
	pulumi.Input

	ToGraphQueryOutput() GraphQueryOutput
	ToGraphQueryOutputWithContext(ctx context.Context) GraphQueryOutput
}

func (*GraphQuery) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQuery)(nil)).Elem()
}

func (i *GraphQuery) ToGraphQueryOutput() GraphQueryOutput {
	return i.ToGraphQueryOutputWithContext(context.Background())
}

func (i *GraphQuery) ToGraphQueryOutputWithContext(ctx context.Context) GraphQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQueryOutput)
}

func (i *GraphQuery) ToOutput(ctx context.Context) pulumix.Output[*GraphQuery] {
	return pulumix.Output[*GraphQuery]{
		OutputState: i.ToGraphQueryOutputWithContext(ctx).OutputState,
	}
}

type GraphQueryOutput struct{ *pulumi.OutputState }

func (GraphQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQuery)(nil)).Elem()
}

func (o GraphQueryOutput) ToGraphQueryOutput() GraphQueryOutput {
	return o
}

func (o GraphQueryOutput) ToGraphQueryOutputWithContext(ctx context.Context) GraphQueryOutput {
	return o
}

func (o GraphQueryOutput) ToOutput(ctx context.Context) pulumix.Output[*GraphQuery] {
	return pulumix.Output[*GraphQuery]{
		OutputState: o.OutputState,
	}
}

// The description of a graph query.
func (o GraphQueryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// This will be used to handle Optimistic Concurrency. If not present, it will always overwrite the existing resource without checking conflict.
func (o GraphQueryOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The location of the resource
func (o GraphQueryOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Azure resource name. This is GUID value. The display name should be assigned within properties field.
func (o GraphQueryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// KQL query that will be graph.
func (o GraphQueryOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

// Enum indicating a type of graph query.
func (o GraphQueryOutput) ResultKind() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringOutput { return v.ResultKind }).(pulumi.StringOutput)
}

// Resource tags
func (o GraphQueryOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Date and time in UTC of the last modification that was made to this graph query definition.
func (o GraphQueryOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringOutput { return v.TimeModified }).(pulumi.StringOutput)
}

// Azure resource type
func (o GraphQueryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GraphQueryOutput{})
}
