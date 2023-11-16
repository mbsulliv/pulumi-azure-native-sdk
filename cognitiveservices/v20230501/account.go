// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Cognitive Services account is an Azure resource representing the provisioned account, it's type, location and SKU.
type Account struct {
	pulumi.CustomResourceState

	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Identity for the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The Kind of the resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of Cognitive Services account.
	Properties AccountPropertiesResponseOutput `pulumi:"properties"`
	// The resource model definition representing SKU
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
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
	if args.Properties != nil {
		args.Properties = args.Properties.ToAccountPropertiesPtrOutput().ApplyT(func(v *AccountProperties) *AccountProperties { return v.Defaults() }).(AccountPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cognitiveservices:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20160201preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20170418:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20210430:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20211001:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20220301:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20221001:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20221201:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20231001preview:Account"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Account
	err := ctx.RegisterResource("azure-native:cognitiveservices/v20230501:Account", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:cognitiveservices/v20230501:Account", name, id, state, &resource, opts...)
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
	// The name of Cognitive Services account.
	AccountName *string `pulumi:"accountName"`
	// Identity for the resource.
	Identity *Identity `pulumi:"identity"`
	// The Kind of the resource.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Properties of Cognitive Services account.
	Properties *AccountProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource model definition representing SKU
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of Cognitive Services account.
	AccountName pulumi.StringPtrInput
	// Identity for the resource.
	Identity IdentityPtrInput
	// The Kind of the resource.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Properties of Cognitive Services account.
	Properties AccountPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The resource model definition representing SKU
	Sku SkuPtrInput
	// Resource tags.
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

// Resource Etag.
func (o AccountOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o AccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Account) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The Kind of the resource.
func (o AccountOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o AccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of Cognitive Services account.
func (o AccountOutput) Properties() AccountPropertiesResponseOutput {
	return o.ApplyT(func(v *Account) AccountPropertiesResponseOutput { return v.Properties }).(AccountPropertiesResponseOutput)
}

// The resource model definition representing SKU
func (o AccountOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Account) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o AccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Account) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Account) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
