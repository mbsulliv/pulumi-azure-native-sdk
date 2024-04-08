// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NetApp account resource
type Account struct {
	pulumi.CustomResourceState

	// Active Directories
	ActiveDirectories ActiveDirectoryResponseArrayOutput `pulumi:"activeDirectories"`
	// Shows the status of disableShowmount for all volumes under the subscription, null equals false
	DisableShowmount pulumi.BoolOutput `pulumi:"disableShowmount"`
	// Encryption settings
	Encryption AccountEncryptionResponsePtrOutput `pulumi:"encryption"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The identity used for the resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
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
	if args.Encryption != nil {
		args.Encryption = args.Encryption.ToAccountEncryptionPtrOutput().ApplyT(func(v *AccountEncryption) *AccountEncryption { return v.Defaults() }).(AccountEncryptionPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20170815:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190501:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190601:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190701:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190801:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191001:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191101:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200201:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200301:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20211001:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220501:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220901:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20230501:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20230501preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20230701:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20230701preview:Account"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Account
	err := ctx.RegisterResource("azure-native:netapp/v20221101preview:Account", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:netapp/v20221101preview:Account", name, id, state, &resource, opts...)
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
	// The name of the NetApp account
	AccountName *string `pulumi:"accountName"`
	// Active Directories
	ActiveDirectories []ActiveDirectory `pulumi:"activeDirectories"`
	// Encryption settings
	Encryption *AccountEncryption `pulumi:"encryption"`
	// The identity used for the resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringPtrInput
	// Active Directories
	ActiveDirectories ActiveDirectoryArrayInput
	// Encryption settings
	Encryption AccountEncryptionPtrInput
	// The identity used for the resource.
	Identity ManagedServiceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
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

// Active Directories
func (o AccountOutput) ActiveDirectories() ActiveDirectoryResponseArrayOutput {
	return o.ApplyT(func(v *Account) ActiveDirectoryResponseArrayOutput { return v.ActiveDirectories }).(ActiveDirectoryResponseArrayOutput)
}

// Shows the status of disableShowmount for all volumes under the subscription, null equals false
func (o AccountOutput) DisableShowmount() pulumi.BoolOutput {
	return o.ApplyT(func(v *Account) pulumi.BoolOutput { return v.DisableShowmount }).(pulumi.BoolOutput)
}

// Encryption settings
func (o AccountOutput) Encryption() AccountEncryptionResponsePtrOutput {
	return o.ApplyT(func(v *Account) AccountEncryptionResponsePtrOutput { return v.Encryption }).(AccountEncryptionResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o AccountOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The identity used for the resource.
func (o AccountOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Account) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o AccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o AccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
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
