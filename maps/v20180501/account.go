// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure resource which represents access to a suite of Maps REST APIs.
type Account struct {
	pulumi.CustomResourceState

	// The location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Maps Account, which is unique within a Resource Group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The map account properties.
	Properties MapsAccountPropertiesResponseOutput `pulumi:"properties"`
	// The SKU of this account.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Gets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:maps:Account"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20170101preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20200201preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20210201:Account"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20210701preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20211201preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20230601:Account"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20230801preview:Account"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Account
	err := ctx.RegisterResource("azure-native:maps/v20180501:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:maps/v20180501:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
}

type AccountState struct {
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name of the Maps Account.
	AccountName *string `pulumi:"accountName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the Azure Resource Group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of this account.
	Sku Sku `pulumi:"sku"`
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the Maps Account.
	AccountName pulumi.StringPtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the Azure Resource Group.
	ResourceGroupName pulumi.StringInput
	// The SKU of this account.
	Sku SkuInput
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

// The location of the resource.
func (o AccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the Maps Account, which is unique within a Resource Group.
func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The map account properties.
func (o AccountOutput) Properties() MapsAccountPropertiesResponseOutput {
	return o.ApplyT(func(v *Account) MapsAccountPropertiesResponseOutput { return v.Properties }).(MapsAccountPropertiesResponseOutput)
}

// The SKU of this account.
func (o AccountOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Account) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// Gets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
func (o AccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Account) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
