// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Content type contract details.
// Azure REST API version: 2022-08-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2022-09-01-preview, 2023-03-01-preview.
type ContentItem struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the content item.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewContentItem registers a new resource with the given unique name, arguments, and options.
func NewContentItem(ctx *pulumi.Context,
	name string, args *ContentItemArgs, opts ...pulumi.ResourceOption) (*ContentItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContentTypeId == nil {
		return nil, errors.New("invalid value for required argument 'ContentTypeId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:ContentItem"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ContentItem
	err := ctx.RegisterResource("azure-native:apimanagement:ContentItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContentItem gets an existing ContentItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContentItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContentItemState, opts ...pulumi.ResourceOption) (*ContentItem, error) {
	var resource ContentItem
	err := ctx.ReadResource("azure-native:apimanagement:ContentItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContentItem resources.
type contentItemState struct {
}

type ContentItemState struct {
}

func (ContentItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*contentItemState)(nil)).Elem()
}

type contentItemArgs struct {
	// Content item identifier.
	ContentItemId *string `pulumi:"contentItemId"`
	// Content type identifier.
	ContentTypeId string `pulumi:"contentTypeId"`
	// Properties of the content item.
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ContentItem resource.
type ContentItemArgs struct {
	// Content item identifier.
	ContentItemId pulumi.StringPtrInput
	// Content type identifier.
	ContentTypeId pulumi.StringInput
	// Properties of the content item.
	Properties pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (ContentItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contentItemArgs)(nil)).Elem()
}

type ContentItemInput interface {
	pulumi.Input

	ToContentItemOutput() ContentItemOutput
	ToContentItemOutputWithContext(ctx context.Context) ContentItemOutput
}

func (*ContentItem) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentItem)(nil)).Elem()
}

func (i *ContentItem) ToContentItemOutput() ContentItemOutput {
	return i.ToContentItemOutputWithContext(context.Background())
}

func (i *ContentItem) ToContentItemOutputWithContext(ctx context.Context) ContentItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentItemOutput)
}

type ContentItemOutput struct{ *pulumi.OutputState }

func (ContentItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentItem)(nil)).Elem()
}

func (o ContentItemOutput) ToContentItemOutput() ContentItemOutput {
	return o
}

func (o ContentItemOutput) ToContentItemOutputWithContext(ctx context.Context) ContentItemOutput {
	return o
}

// The name of the resource
func (o ContentItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the content item.
func (o ContentItemOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ContentItem) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ContentItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentItem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContentItemOutput{})
}
