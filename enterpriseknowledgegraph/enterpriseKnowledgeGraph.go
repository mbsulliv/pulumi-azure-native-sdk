// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package enterpriseknowledgegraph

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// EnterpriseKnowledgeGraph resource definition
// Azure REST API version: 2018-12-03. Prior API version in Azure Native 1.x: 2018-12-03.
type EnterpriseKnowledgeGraph struct {
	pulumi.CustomResourceState

	// Specifies the location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The set of properties specific to EnterpriseKnowledgeGraph resource
	Properties EnterpriseKnowledgeGraphPropertiesResponseOutput `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEnterpriseKnowledgeGraph registers a new resource with the given unique name, arguments, and options.
func NewEnterpriseKnowledgeGraph(ctx *pulumi.Context,
	name string, args *EnterpriseKnowledgeGraphArgs, opts ...pulumi.ResourceOption) (*EnterpriseKnowledgeGraph, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:enterpriseknowledgegraph/v20181203:EnterpriseKnowledgeGraph"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EnterpriseKnowledgeGraph
	err := ctx.RegisterResource("azure-native:enterpriseknowledgegraph:EnterpriseKnowledgeGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterpriseKnowledgeGraph gets an existing EnterpriseKnowledgeGraph resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterpriseKnowledgeGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseKnowledgeGraphState, opts ...pulumi.ResourceOption) (*EnterpriseKnowledgeGraph, error) {
	var resource EnterpriseKnowledgeGraph
	err := ctx.ReadResource("azure-native:enterpriseknowledgegraph:EnterpriseKnowledgeGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterpriseKnowledgeGraph resources.
type enterpriseKnowledgeGraphState struct {
}

type EnterpriseKnowledgeGraphState struct {
}

func (EnterpriseKnowledgeGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseKnowledgeGraphState)(nil)).Elem()
}

type enterpriseKnowledgeGraphArgs struct {
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// The set of properties specific to EnterpriseKnowledgeGraph resource
	Properties *EnterpriseKnowledgeGraphProperties `pulumi:"properties"`
	// The name of the EnterpriseKnowledgeGraph resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the EnterpriseKnowledgeGraph resource.
	ResourceName *string `pulumi:"resourceName"`
	// Gets or sets the SKU of the resource.
	Sku *Sku `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EnterpriseKnowledgeGraph resource.
type EnterpriseKnowledgeGraphArgs struct {
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// The set of properties specific to EnterpriseKnowledgeGraph resource
	Properties EnterpriseKnowledgeGraphPropertiesPtrInput
	// The name of the EnterpriseKnowledgeGraph resource group in the user subscription.
	ResourceGroupName pulumi.StringInput
	// The name of the EnterpriseKnowledgeGraph resource.
	ResourceName pulumi.StringPtrInput
	// Gets or sets the SKU of the resource.
	Sku SkuPtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
}

func (EnterpriseKnowledgeGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseKnowledgeGraphArgs)(nil)).Elem()
}

type EnterpriseKnowledgeGraphInput interface {
	pulumi.Input

	ToEnterpriseKnowledgeGraphOutput() EnterpriseKnowledgeGraphOutput
	ToEnterpriseKnowledgeGraphOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphOutput
}

func (*EnterpriseKnowledgeGraph) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseKnowledgeGraph)(nil)).Elem()
}

func (i *EnterpriseKnowledgeGraph) ToEnterpriseKnowledgeGraphOutput() EnterpriseKnowledgeGraphOutput {
	return i.ToEnterpriseKnowledgeGraphOutputWithContext(context.Background())
}

func (i *EnterpriseKnowledgeGraph) ToEnterpriseKnowledgeGraphOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseKnowledgeGraphOutput)
}

type EnterpriseKnowledgeGraphOutput struct{ *pulumi.OutputState }

func (EnterpriseKnowledgeGraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseKnowledgeGraph)(nil)).Elem()
}

func (o EnterpriseKnowledgeGraphOutput) ToEnterpriseKnowledgeGraphOutput() EnterpriseKnowledgeGraphOutput {
	return o
}

func (o EnterpriseKnowledgeGraphOutput) ToEnterpriseKnowledgeGraphOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphOutput {
	return o
}

// Specifies the location of the resource.
func (o EnterpriseKnowledgeGraphOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraph) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Specifies the name of the resource.
func (o EnterpriseKnowledgeGraphOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraph) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The set of properties specific to EnterpriseKnowledgeGraph resource
func (o EnterpriseKnowledgeGraphOutput) Properties() EnterpriseKnowledgeGraphPropertiesResponseOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraph) EnterpriseKnowledgeGraphPropertiesResponseOutput {
		return v.Properties
	}).(EnterpriseKnowledgeGraphPropertiesResponseOutput)
}

// Gets or sets the SKU of the resource.
func (o EnterpriseKnowledgeGraphOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraph) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Contains resource tags defined as key/value pairs.
func (o EnterpriseKnowledgeGraphOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraph) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the type of the resource.
func (o EnterpriseKnowledgeGraphOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraph) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnterpriseKnowledgeGraphOutput{})
}
