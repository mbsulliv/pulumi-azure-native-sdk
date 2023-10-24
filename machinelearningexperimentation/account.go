// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningexperimentation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An object that represents a machine learning team account.
// Azure REST API version: 2017-05-01-preview. Prior API version in Azure Native 1.x: 2017-05-01-preview.
type Account struct {
	pulumi.CustomResourceState

	// The immutable id associated with this team account.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The creation date of the machine learning team account in ISO8601 format.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The description of this workspace.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The uri for this machine learning team account.
	DiscoveryUri pulumi.StringOutput `pulumi:"discoveryUri"`
	// The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// The fully qualified arm id of the user key vault.
	KeyVaultId pulumi.StringOutput `pulumi:"keyVaultId"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment state of team account resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The no of users/seats who can access this team account. This property defines the charge on the team account.
	Seats pulumi.StringPtrOutput `pulumi:"seats"`
	// The properties of the storage account for the machine learning team account.
	StorageAccount StorageAccountPropertiesResponseOutput `pulumi:"storageAccount"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The fully qualified arm id of the vso account to be used for this team account.
	VsoAccountId pulumi.StringOutput `pulumi:"vsoAccountId"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyVaultId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccount == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccount'")
	}
	if args.VsoAccountId == nil {
		return nil, errors.New("invalid value for required argument 'VsoAccountId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningexperimentation/v20170501preview:Account"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Account
	err := ctx.RegisterResource("azure-native:machinelearningexperimentation:Account", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:machinelearningexperimentation:Account", name, id, state, &resource, opts...)
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
	// The name of the machine learning team account.
	AccountName *string `pulumi:"accountName"`
	// The description of this workspace.
	Description *string `pulumi:"description"`
	// The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
	FriendlyName *string `pulumi:"friendlyName"`
	// The fully qualified arm id of the user key vault.
	KeyVaultId string `pulumi:"keyVaultId"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// The name of the resource group to which the machine learning team account belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The no of users/seats who can access this team account. This property defines the charge on the team account.
	Seats *string `pulumi:"seats"`
	// The properties of the storage account for the machine learning team account.
	StorageAccount StorageAccountProperties `pulumi:"storageAccount"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The fully qualified arm id of the vso account to be used for this team account.
	VsoAccountId string `pulumi:"vsoAccountId"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the machine learning team account.
	AccountName pulumi.StringPtrInput
	// The description of this workspace.
	Description pulumi.StringPtrInput
	// The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
	FriendlyName pulumi.StringPtrInput
	// The fully qualified arm id of the user key vault.
	KeyVaultId pulumi.StringInput
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringPtrInput
	// The name of the resource group to which the machine learning team account belongs.
	ResourceGroupName pulumi.StringInput
	// The no of users/seats who can access this team account. This property defines the charge on the team account.
	Seats pulumi.StringPtrInput
	// The properties of the storage account for the machine learning team account.
	StorageAccount StorageAccountPropertiesInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The fully qualified arm id of the vso account to be used for this team account.
	VsoAccountId pulumi.StringInput
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

// The immutable id associated with this team account.
func (o AccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The creation date of the machine learning team account in ISO8601 format.
func (o AccountOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The description of this workspace.
func (o AccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The uri for this machine learning team account.
func (o AccountOutput) DiscoveryUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.DiscoveryUri }).(pulumi.StringOutput)
}

// The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
func (o AccountOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// The fully qualified arm id of the user key vault.
func (o AccountOutput) KeyVaultId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.KeyVaultId }).(pulumi.StringOutput)
}

// The location of the resource. This cannot be changed after the resource is created.
func (o AccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource.
func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current deployment state of team account resource. The provisioningState is to indicate states for resource provisioning.
func (o AccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The no of users/seats who can access this team account. This property defines the charge on the team account.
func (o AccountOutput) Seats() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Seats }).(pulumi.StringPtrOutput)
}

// The properties of the storage account for the machine learning team account.
func (o AccountOutput) StorageAccount() StorageAccountPropertiesResponseOutput {
	return o.ApplyT(func(v *Account) StorageAccountPropertiesResponseOutput { return v.StorageAccount }).(StorageAccountPropertiesResponseOutput)
}

// The tags of the resource.
func (o AccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Account) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The fully qualified arm id of the vso account to be used for this team account.
func (o AccountOutput) VsoAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.VsoAccountId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
