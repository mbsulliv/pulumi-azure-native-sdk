// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

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
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Contents of the Policy as defined by the format.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewProductPolicy registers a new resource with the given unique name, arguments, and options.
func NewProductPolicy(ctx *pulumi.Context,
	name string, args *ProductPolicyArgs, opts ...pulumi.ResourceOption) (*ProductPolicy, error) {
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
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if args.Format == nil {
		args.Format = pulumi.StringPtr("xml")
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
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ProductPolicy"),
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
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:ProductPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProductPolicy
	err := ctx.RegisterResource("azure-native:apimanagement/v20220801:ProductPolicy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:apimanagement/v20220801:ProductPolicy", name, id, state, &resource, opts...)
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
	Format *string `pulumi:"format"`
	// The identifier of the Policy.
	PolicyId *string `pulumi:"policyId"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Contents of the Policy as defined by the format.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a ProductPolicy resource.
type ProductPolicyArgs struct {
	// Format of the policyContent.
	Format pulumi.StringPtrInput
	// The identifier of the Policy.
	PolicyId pulumi.StringPtrInput
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Contents of the Policy as defined by the format.
	Value pulumi.StringInput
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
func (o ProductPolicyOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductPolicy) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ProductPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ProductPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Contents of the Policy as defined by the format.
func (o ProductPolicyOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductPolicy) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProductPolicyOutput{})
}
