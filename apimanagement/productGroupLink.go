// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Product-group link details.
// Azure REST API version: 2022-09-01-preview.
type ProductGroupLink struct {
	pulumi.CustomResourceState

	// Full resource Id of a group.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProductGroupLink registers a new resource with the given unique name, arguments, and options.
func NewProductGroupLink(ctx *pulumi.Context,
	name string, args *ProductGroupLinkArgs, opts ...pulumi.ResourceOption) (*ProductGroupLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
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
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:ProductGroupLink"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:ProductGroupLink"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProductGroupLink
	err := ctx.RegisterResource("azure-native:apimanagement:ProductGroupLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProductGroupLink gets an existing ProductGroupLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProductGroupLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductGroupLinkState, opts ...pulumi.ResourceOption) (*ProductGroupLink, error) {
	var resource ProductGroupLink
	err := ctx.ReadResource("azure-native:apimanagement:ProductGroupLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProductGroupLink resources.
type productGroupLinkState struct {
}

type ProductGroupLinkState struct {
}

func (ProductGroupLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*productGroupLinkState)(nil)).Elem()
}

type productGroupLinkArgs struct {
	// Full resource Id of a group.
	GroupId string `pulumi:"groupId"`
	// Product-Group link identifier. Must be unique in the current API Management service instance.
	GroupLinkId *string `pulumi:"groupLinkId"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ProductGroupLink resource.
type ProductGroupLinkArgs struct {
	// Full resource Id of a group.
	GroupId pulumi.StringInput
	// Product-Group link identifier. Must be unique in the current API Management service instance.
	GroupLinkId pulumi.StringPtrInput
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (ProductGroupLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productGroupLinkArgs)(nil)).Elem()
}

type ProductGroupLinkInput interface {
	pulumi.Input

	ToProductGroupLinkOutput() ProductGroupLinkOutput
	ToProductGroupLinkOutputWithContext(ctx context.Context) ProductGroupLinkOutput
}

func (*ProductGroupLink) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductGroupLink)(nil)).Elem()
}

func (i *ProductGroupLink) ToProductGroupLinkOutput() ProductGroupLinkOutput {
	return i.ToProductGroupLinkOutputWithContext(context.Background())
}

func (i *ProductGroupLink) ToProductGroupLinkOutputWithContext(ctx context.Context) ProductGroupLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductGroupLinkOutput)
}

func (i *ProductGroupLink) ToOutput(ctx context.Context) pulumix.Output[*ProductGroupLink] {
	return pulumix.Output[*ProductGroupLink]{
		OutputState: i.ToProductGroupLinkOutputWithContext(ctx).OutputState,
	}
}

type ProductGroupLinkOutput struct{ *pulumi.OutputState }

func (ProductGroupLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductGroupLink)(nil)).Elem()
}

func (o ProductGroupLinkOutput) ToProductGroupLinkOutput() ProductGroupLinkOutput {
	return o
}

func (o ProductGroupLinkOutput) ToProductGroupLinkOutputWithContext(ctx context.Context) ProductGroupLinkOutput {
	return o
}

func (o ProductGroupLinkOutput) ToOutput(ctx context.Context) pulumix.Output[*ProductGroupLink] {
	return pulumix.Output[*ProductGroupLink]{
		OutputState: o.OutputState,
	}
}

// Full resource Id of a group.
func (o ProductGroupLinkOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductGroupLink) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The name of the resource
func (o ProductGroupLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductGroupLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ProductGroupLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductGroupLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProductGroupLinkOutput{})
}
