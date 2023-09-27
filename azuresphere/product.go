// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuresphere

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An product resource belonging to a catalog resource.
// Azure REST API version: 2022-09-01-preview. Prior API version in Azure Native 1.x: 2022-09-01-preview
type Product struct {
	pulumi.CustomResourceState

	// Description of the product
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProduct registers a new resource with the given unique name, arguments, and options.
func NewProduct(ctx *pulumi.Context,
	name string, args *ProductArgs, opts ...pulumi.ResourceOption) (*Product, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogName == nil {
		return nil, errors.New("invalid value for required argument 'CatalogName'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azuresphere/v20220901preview:Product"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Product
	err := ctx.RegisterResource("azure-native:azuresphere:Product", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProduct gets an existing Product resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductState, opts ...pulumi.ResourceOption) (*Product, error) {
	var resource Product
	err := ctx.ReadResource("azure-native:azuresphere:Product", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Product resources.
type productState struct {
}

type ProductState struct {
}

func (ProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*productState)(nil)).Elem()
}

type productArgs struct {
	// Name of catalog
	CatalogName string `pulumi:"catalogName"`
	// Description of the product
	Description string `pulumi:"description"`
	// Name of product.
	ProductName *string `pulumi:"productName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Product resource.
type ProductArgs struct {
	// Name of catalog
	CatalogName pulumi.StringInput
	// Description of the product
	Description pulumi.StringInput
	// Name of product.
	ProductName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productArgs)(nil)).Elem()
}

type ProductInput interface {
	pulumi.Input

	ToProductOutput() ProductOutput
	ToProductOutputWithContext(ctx context.Context) ProductOutput
}

func (*Product) ElementType() reflect.Type {
	return reflect.TypeOf((**Product)(nil)).Elem()
}

func (i *Product) ToProductOutput() ProductOutput {
	return i.ToProductOutputWithContext(context.Background())
}

func (i *Product) ToProductOutputWithContext(ctx context.Context) ProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductOutput)
}

func (i *Product) ToOutput(ctx context.Context) pulumix.Output[*Product] {
	return pulumix.Output[*Product]{
		OutputState: i.ToProductOutputWithContext(ctx).OutputState,
	}
}

type ProductOutput struct{ *pulumi.OutputState }

func (ProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Product)(nil)).Elem()
}

func (o ProductOutput) ToProductOutput() ProductOutput {
	return o
}

func (o ProductOutput) ToProductOutputWithContext(ctx context.Context) ProductOutput {
	return o
}

func (o ProductOutput) ToOutput(ctx context.Context) pulumix.Output[*Product] {
	return pulumix.Output[*Product]{
		OutputState: o.OutputState,
	}
}

// Description of the product
func (o ProductOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Product) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The name of the resource
func (o ProductOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Product) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o ProductOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Product) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ProductOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Product) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ProductOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Product) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProductOutput{})
}
