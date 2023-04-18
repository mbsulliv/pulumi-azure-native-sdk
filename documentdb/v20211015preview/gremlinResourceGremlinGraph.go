// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB Gremlin graph.
type GremlinResourceGremlinGraph struct {
	pulumi.CustomResourceState

	// Identity for the resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                `pulumi:"name"`
	Options  GremlinGraphGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource GremlinGraphGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGremlinResourceGremlinGraph registers a new resource with the given unique name, arguments, and options.
func NewGremlinResourceGremlinGraph(ctx *pulumi.Context,
	name string, args *GremlinResourceGremlinGraphArgs, opts ...pulumi.ResourceOption) (*GremlinResourceGremlinGraph, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Resource = args.Resource.ToGremlinGraphResourceOutput().ApplyT(func(v GremlinGraphResource) GremlinGraphResource { return *v.Defaults() }).(GremlinGraphResourceOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:GremlinResourceGremlinGraph"),
		},
	})
	opts = append(opts, aliases)
	var resource GremlinResourceGremlinGraph
	err := ctx.RegisterResource("azure-native:documentdb/v20211015preview:GremlinResourceGremlinGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGremlinResourceGremlinGraph gets an existing GremlinResourceGremlinGraph resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGremlinResourceGremlinGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GremlinResourceGremlinGraphState, opts ...pulumi.ResourceOption) (*GremlinResourceGremlinGraph, error) {
	var resource GremlinResourceGremlinGraph
	err := ctx.ReadResource("azure-native:documentdb/v20211015preview:GremlinResourceGremlinGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GremlinResourceGremlinGraph resources.
type gremlinResourceGremlinGraphState struct {
}

type GremlinResourceGremlinGraphState struct {
}

func (GremlinResourceGremlinGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinResourceGremlinGraphState)(nil)).Elem()
}

type gremlinResourceGremlinGraphArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// Cosmos DB graph name.
	GraphName *string `pulumi:"graphName"`
	// Identity for the resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options *CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a Gremlin graph
	Resource GremlinGraphResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GremlinResourceGremlinGraph resource.
type GremlinResourceGremlinGraphArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// Cosmos DB graph name.
	GraphName pulumi.StringPtrInput
	// Identity for the resource.
	Identity ManagedServiceIdentityPtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsPtrInput
	// The standard JSON format of a Gremlin graph
	Resource GremlinGraphResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (GremlinResourceGremlinGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinResourceGremlinGraphArgs)(nil)).Elem()
}

type GremlinResourceGremlinGraphInput interface {
	pulumi.Input

	ToGremlinResourceGremlinGraphOutput() GremlinResourceGremlinGraphOutput
	ToGremlinResourceGremlinGraphOutputWithContext(ctx context.Context) GremlinResourceGremlinGraphOutput
}

func (*GremlinResourceGremlinGraph) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinResourceGremlinGraph)(nil)).Elem()
}

func (i *GremlinResourceGremlinGraph) ToGremlinResourceGremlinGraphOutput() GremlinResourceGremlinGraphOutput {
	return i.ToGremlinResourceGremlinGraphOutputWithContext(context.Background())
}

func (i *GremlinResourceGremlinGraph) ToGremlinResourceGremlinGraphOutputWithContext(ctx context.Context) GremlinResourceGremlinGraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinResourceGremlinGraphOutput)
}

type GremlinResourceGremlinGraphOutput struct{ *pulumi.OutputState }

func (GremlinResourceGremlinGraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinResourceGremlinGraph)(nil)).Elem()
}

func (o GremlinResourceGremlinGraphOutput) ToGremlinResourceGremlinGraphOutput() GremlinResourceGremlinGraphOutput {
	return o
}

func (o GremlinResourceGremlinGraphOutput) ToGremlinResourceGremlinGraphOutputWithContext(ctx context.Context) GremlinResourceGremlinGraphOutput {
	return o
}

// Identity for the resource.
func (o GremlinResourceGremlinGraphOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *GremlinResourceGremlinGraph) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The location of the resource group to which the resource belongs.
func (o GremlinResourceGremlinGraphOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinResourceGremlinGraph) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o GremlinResourceGremlinGraphOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GremlinResourceGremlinGraph) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GremlinResourceGremlinGraphOutput) Options() GremlinGraphGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v *GremlinResourceGremlinGraph) GremlinGraphGetPropertiesResponseOptionsPtrOutput {
		return v.Options
	}).(GremlinGraphGetPropertiesResponseOptionsPtrOutput)
}

func (o GremlinResourceGremlinGraphOutput) Resource() GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *GremlinResourceGremlinGraph) GremlinGraphGetPropertiesResponseResourcePtrOutput {
		return v.Resource
	}).(GremlinGraphGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o GremlinResourceGremlinGraphOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GremlinResourceGremlinGraph) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o GremlinResourceGremlinGraphOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GremlinResourceGremlinGraph) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GremlinResourceGremlinGraphOutput{})
}
