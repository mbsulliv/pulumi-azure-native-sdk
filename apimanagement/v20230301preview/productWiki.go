// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Wiki properties
type ProductWiki struct {
	pulumi.CustomResourceState

	// Collection wiki documents included into this wiki.
	Documents WikiDocumentationContractResponseArrayOutput `pulumi:"documents"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProductWiki registers a new resource with the given unique name, arguments, and options.
func NewProductWiki(ctx *pulumi.Context,
	name string, args *ProductWikiArgs, opts ...pulumi.ResourceOption) (*ProductWiki, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ProductWiki"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:ProductWiki"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:ProductWiki"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProductWiki
	err := ctx.RegisterResource("azure-native:apimanagement/v20230301preview:ProductWiki", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProductWiki gets an existing ProductWiki resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProductWiki(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductWikiState, opts ...pulumi.ResourceOption) (*ProductWiki, error) {
	var resource ProductWiki
	err := ctx.ReadResource("azure-native:apimanagement/v20230301preview:ProductWiki", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProductWiki resources.
type productWikiState struct {
}

type ProductWikiState struct {
}

func (ProductWikiState) ElementType() reflect.Type {
	return reflect.TypeOf((*productWikiState)(nil)).Elem()
}

type productWikiArgs struct {
	// Collection wiki documents included into this wiki.
	Documents []WikiDocumentationContract `pulumi:"documents"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ProductWiki resource.
type ProductWikiArgs struct {
	// Collection wiki documents included into this wiki.
	Documents WikiDocumentationContractArrayInput
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (ProductWikiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productWikiArgs)(nil)).Elem()
}

type ProductWikiInput interface {
	pulumi.Input

	ToProductWikiOutput() ProductWikiOutput
	ToProductWikiOutputWithContext(ctx context.Context) ProductWikiOutput
}

func (*ProductWiki) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductWiki)(nil)).Elem()
}

func (i *ProductWiki) ToProductWikiOutput() ProductWikiOutput {
	return i.ToProductWikiOutputWithContext(context.Background())
}

func (i *ProductWiki) ToProductWikiOutputWithContext(ctx context.Context) ProductWikiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductWikiOutput)
}

type ProductWikiOutput struct{ *pulumi.OutputState }

func (ProductWikiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductWiki)(nil)).Elem()
}

func (o ProductWikiOutput) ToProductWikiOutput() ProductWikiOutput {
	return o
}

func (o ProductWikiOutput) ToProductWikiOutputWithContext(ctx context.Context) ProductWikiOutput {
	return o
}

// Collection wiki documents included into this wiki.
func (o ProductWikiOutput) Documents() WikiDocumentationContractResponseArrayOutput {
	return o.ApplyT(func(v *ProductWiki) WikiDocumentationContractResponseArrayOutput { return v.Documents }).(WikiDocumentationContractResponseArrayOutput)
}

// The name of the resource
func (o ProductWikiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductWiki) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ProductWikiOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductWiki) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProductWikiOutput{})
}
