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

// Contract details.
// Azure REST API version: 2022-08-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2022-09-01-preview, 2023-03-01-preview.
type ProductGroup struct {
	pulumi.CustomResourceState

	// true if the group is one of the three system groups (Administrators, Developers, or Guests); otherwise false.
	BuiltIn pulumi.BoolOutput `pulumi:"builtIn"`
	// Group description. Can contain HTML formatting tags.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Group name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// For external groups, this property contains the id of the group from the external identity provider, e.g. for Azure Active Directory `aad://<tenant>.onmicrosoft.com/groups/<group object id>`; otherwise the value is null.
	ExternalId pulumi.StringPtrOutput `pulumi:"externalId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProductGroup registers a new resource with the given unique name, arguments, and options.
func NewProductGroup(ctx *pulumi.Context,
	name string, args *ProductGroupArgs, opts ...pulumi.ResourceOption) (*ProductGroup, error) {
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
			Type: pulumi.String("azure-native:apimanagement/v20170301:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:ProductGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProductGroup
	err := ctx.RegisterResource("azure-native:apimanagement:ProductGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProductGroup gets an existing ProductGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProductGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductGroupState, opts ...pulumi.ResourceOption) (*ProductGroup, error) {
	var resource ProductGroup
	err := ctx.ReadResource("azure-native:apimanagement:ProductGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProductGroup resources.
type productGroupState struct {
}

type ProductGroupState struct {
}

func (ProductGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*productGroupState)(nil)).Elem()
}

type productGroupArgs struct {
	// Group identifier. Must be unique in the current API Management service instance.
	GroupId *string `pulumi:"groupId"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ProductGroup resource.
type ProductGroupArgs struct {
	// Group identifier. Must be unique in the current API Management service instance.
	GroupId pulumi.StringPtrInput
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (ProductGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productGroupArgs)(nil)).Elem()
}

type ProductGroupInput interface {
	pulumi.Input

	ToProductGroupOutput() ProductGroupOutput
	ToProductGroupOutputWithContext(ctx context.Context) ProductGroupOutput
}

func (*ProductGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductGroup)(nil)).Elem()
}

func (i *ProductGroup) ToProductGroupOutput() ProductGroupOutput {
	return i.ToProductGroupOutputWithContext(context.Background())
}

func (i *ProductGroup) ToProductGroupOutputWithContext(ctx context.Context) ProductGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductGroupOutput)
}

type ProductGroupOutput struct{ *pulumi.OutputState }

func (ProductGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductGroup)(nil)).Elem()
}

func (o ProductGroupOutput) ToProductGroupOutput() ProductGroupOutput {
	return o
}

func (o ProductGroupOutput) ToProductGroupOutputWithContext(ctx context.Context) ProductGroupOutput {
	return o
}

// true if the group is one of the three system groups (Administrators, Developers, or Guests); otherwise false.
func (o ProductGroupOutput) BuiltIn() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProductGroup) pulumi.BoolOutput { return v.BuiltIn }).(pulumi.BoolOutput)
}

// Group description. Can contain HTML formatting tags.
func (o ProductGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Group name.
func (o ProductGroupOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductGroup) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// For external groups, this property contains the id of the group from the external identity provider, e.g. for Azure Active Directory `aad://<tenant>.onmicrosoft.com/groups/<group object id>`; otherwise the value is null.
func (o ProductGroupOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductGroup) pulumi.StringPtrOutput { return v.ExternalId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ProductGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ProductGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProductGroupOutput{})
}
