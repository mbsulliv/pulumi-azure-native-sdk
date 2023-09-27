// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An account data transfer object.
// Azure REST API version: 2021-08-01. Prior API version in Azure Native 1.x: 2020-09-01
type Account struct {
	pulumi.CustomResourceState

	// Time at which the account was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Identity Info on the Account
	Identity IdentityResponseOutput `pulumi:"identity"`
	// Location of the azure resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the Account
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// System Data of the Azure resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tags on the azure resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
	// Email of the user who created the resource
	UserEmail pulumi.StringOutput `pulumi:"userEmail"`
	// Name of the user who created the resource
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:Account"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:Account"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:Account"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Account
	err := ctx.RegisterResource("azure-native:datashare:Account", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:datashare:Account", name, id, state, &resource, opts...)
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
	// The name of the share account.
	AccountName *string `pulumi:"accountName"`
	// Identity Info on the Account
	Identity Identity `pulumi:"identity"`
	// Location of the azure resource.
	Location *string `pulumi:"location"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags on the azure resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the share account.
	AccountName pulumi.StringPtrInput
	// Identity Info on the Account
	Identity IdentityInput
	// Location of the azure resource.
	Location pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Tags on the azure resource.
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

func (i *Account) ToOutput(ctx context.Context) pulumix.Output[*Account] {
	return pulumix.Output[*Account]{
		OutputState: i.ToAccountOutputWithContext(ctx).OutputState,
	}
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

func (o AccountOutput) ToOutput(ctx context.Context) pulumix.Output[*Account] {
	return pulumix.Output[*Account]{
		OutputState: o.OutputState,
	}
}

// Time at which the account was created.
func (o AccountOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Identity Info on the Account
func (o AccountOutput) Identity() IdentityResponseOutput {
	return o.ApplyT(func(v *Account) IdentityResponseOutput { return v.Identity }).(IdentityResponseOutput)
}

// Location of the azure resource.
func (o AccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the azure resource
func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the Account
func (o AccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o AccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Account) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags on the azure resource.
func (o AccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Account) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the azure resource
func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Email of the user who created the resource
func (o AccountOutput) UserEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.UserEmail }).(pulumi.StringOutput)
}

// Name of the user who created the resource
func (o AccountOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
