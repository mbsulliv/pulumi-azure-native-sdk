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

// Product-API link details.
// Azure REST API version: 2022-09-01-preview.
type ProductApiLink struct {
	pulumi.CustomResourceState

	// Full resource Id of an API.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProductApiLink registers a new resource with the given unique name, arguments, and options.
func NewProductApiLink(ctx *pulumi.Context,
	name string, args *ProductApiLinkArgs, opts ...pulumi.ResourceOption) (*ProductApiLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
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
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:ProductApiLink"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:ProductApiLink"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProductApiLink
	err := ctx.RegisterResource("azure-native:apimanagement:ProductApiLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProductApiLink gets an existing ProductApiLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProductApiLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductApiLinkState, opts ...pulumi.ResourceOption) (*ProductApiLink, error) {
	var resource ProductApiLink
	err := ctx.ReadResource("azure-native:apimanagement:ProductApiLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProductApiLink resources.
type productApiLinkState struct {
}

type ProductApiLinkState struct {
}

func (ProductApiLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*productApiLinkState)(nil)).Elem()
}

type productApiLinkArgs struct {
	// Full resource Id of an API.
	ApiId string `pulumi:"apiId"`
	// Product-API link identifier. Must be unique in the current API Management service instance.
	ApiLinkId *string `pulumi:"apiLinkId"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ProductApiLink resource.
type ProductApiLinkArgs struct {
	// Full resource Id of an API.
	ApiId pulumi.StringInput
	// Product-API link identifier. Must be unique in the current API Management service instance.
	ApiLinkId pulumi.StringPtrInput
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (ProductApiLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productApiLinkArgs)(nil)).Elem()
}

type ProductApiLinkInput interface {
	pulumi.Input

	ToProductApiLinkOutput() ProductApiLinkOutput
	ToProductApiLinkOutputWithContext(ctx context.Context) ProductApiLinkOutput
}

func (*ProductApiLink) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductApiLink)(nil)).Elem()
}

func (i *ProductApiLink) ToProductApiLinkOutput() ProductApiLinkOutput {
	return i.ToProductApiLinkOutputWithContext(context.Background())
}

func (i *ProductApiLink) ToProductApiLinkOutputWithContext(ctx context.Context) ProductApiLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductApiLinkOutput)
}

func (i *ProductApiLink) ToOutput(ctx context.Context) pulumix.Output[*ProductApiLink] {
	return pulumix.Output[*ProductApiLink]{
		OutputState: i.ToProductApiLinkOutputWithContext(ctx).OutputState,
	}
}

type ProductApiLinkOutput struct{ *pulumi.OutputState }

func (ProductApiLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductApiLink)(nil)).Elem()
}

func (o ProductApiLinkOutput) ToProductApiLinkOutput() ProductApiLinkOutput {
	return o
}

func (o ProductApiLinkOutput) ToProductApiLinkOutputWithContext(ctx context.Context) ProductApiLinkOutput {
	return o
}

func (o ProductApiLinkOutput) ToOutput(ctx context.Context) pulumix.Output[*ProductApiLink] {
	return pulumix.Output[*ProductApiLink]{
		OutputState: o.OutputState,
	}
}

// Full resource Id of an API.
func (o ProductApiLinkOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductApiLink) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// The name of the resource
func (o ProductApiLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductApiLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ProductApiLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductApiLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProductApiLinkOutput{})
}
