// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Policy Contract details.
type ProductPolicy struct {
	pulumi.CustomResourceState

	// Format of the policyContent.
	ContentFormat pulumi.StringPtrOutput `pulumi:"contentFormat"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent pulumi.StringOutput `pulumi:"policyContent"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProductPolicy registers a new resource with the given unique name, arguments, and options.
func NewProductPolicy(ctx *pulumi.Context,
	name string, args *ProductPolicyArgs, opts ...pulumi.ResourceOption) (*ProductPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyContent == nil {
		return nil, errors.New("invalid value for required argument 'PolicyContent'")
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
	if args.ContentFormat == nil {
		args.ContentFormat = pulumi.StringPtr("xml")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:ProductPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProductPolicy
	err := ctx.RegisterResource("azure-native:apimanagement/v20180601preview:ProductPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProductPolicy gets an existing ProductPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProductPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductPolicyState, opts ...pulumi.ResourceOption) (*ProductPolicy, error) {
	var resource ProductPolicy
	err := ctx.ReadResource("azure-native:apimanagement/v20180601preview:ProductPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProductPolicy resources.
type productPolicyState struct {
}

type ProductPolicyState struct {
}

func (ProductPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*productPolicyState)(nil)).Elem()
}

type productPolicyArgs struct {
	// Format of the policyContent.
	ContentFormat *string `pulumi:"contentFormat"`
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent string `pulumi:"policyContent"`
	// The identifier of the Policy.
	PolicyId *string `pulumi:"policyId"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ProductPolicy resource.
type ProductPolicyArgs struct {
	// Format of the policyContent.
	ContentFormat pulumi.StringPtrInput
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent pulumi.StringInput
	// The identifier of the Policy.
	PolicyId pulumi.StringPtrInput
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (ProductPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productPolicyArgs)(nil)).Elem()
}

type ProductPolicyInput interface {
	pulumi.Input

	ToProductPolicyOutput() ProductPolicyOutput
	ToProductPolicyOutputWithContext(ctx context.Context) ProductPolicyOutput
}

func (*ProductPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductPolicy)(nil)).Elem()
}

func (i *ProductPolicy) ToProductPolicyOutput() ProductPolicyOutput {
	return i.ToProductPolicyOutputWithContext(context.Background())
}

func (i *ProductPolicy) ToProductPolicyOutputWithContext(ctx context.Context) ProductPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductPolicyOutput)
}

func (i *ProductPolicy) ToOutput(ctx context.Context) pulumix.Output[*ProductPolicy] {
	return pulumix.Output[*ProductPolicy]{
		OutputState: i.ToProductPolicyOutputWithContext(ctx).OutputState,
	}
}

type ProductPolicyOutput struct{ *pulumi.OutputState }

func (ProductPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductPolicy)(nil)).Elem()
}

func (o ProductPolicyOutput) ToProductPolicyOutput() ProductPolicyOutput {
	return o
}

func (o ProductPolicyOutput) ToProductPolicyOutputWithContext(ctx context.Context) ProductPolicyOutput {
	return o
}

func (o ProductPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*ProductPolicy] {
	return pulumix.Output[*ProductPolicy]{
		OutputState: o.OutputState,
	}
}

// Format of the policyContent.
func (o ProductPolicyOutput) ContentFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductPolicy) pulumi.StringPtrOutput { return v.ContentFormat }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ProductPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Json escaped Xml Encoded contents of the Policy.
func (o ProductPolicyOutput) PolicyContent() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductPolicy) pulumi.StringOutput { return v.PolicyContent }).(pulumi.StringOutput)
}

// Resource type for API Management resource.
func (o ProductPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProductPolicyOutput{})
}
