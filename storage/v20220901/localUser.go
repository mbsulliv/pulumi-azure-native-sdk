// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The local user associated with the storage accounts.
type LocalUser struct {
	pulumi.CustomResourceState

	// Indicates whether shared key exists. Set it to false to remove existing shared key.
	HasSharedKey pulumi.BoolPtrOutput `pulumi:"hasSharedKey"`
	// Indicates whether ssh key exists. Set it to false to remove existing SSH key.
	HasSshKey pulumi.BoolPtrOutput `pulumi:"hasSshKey"`
	// Indicates whether ssh password exists. Set it to false to remove existing SSH password.
	HasSshPassword pulumi.BoolPtrOutput `pulumi:"hasSshPassword"`
	// Optional, local user home directory.
	HomeDirectory pulumi.StringPtrOutput `pulumi:"homeDirectory"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The permission scopes of the local user.
	PermissionScopes PermissionScopeResponseArrayOutput `pulumi:"permissionScopes"`
	// A unique Security Identifier that is generated by the server.
	Sid pulumi.StringOutput `pulumi:"sid"`
	// Optional, local user ssh authorized keys for SFTP.
	SshAuthorizedKeys SshPublicKeyResponseArrayOutput `pulumi:"sshAuthorizedKeys"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLocalUser registers a new resource with the given unique name, arguments, and options.
func NewLocalUser(ctx *pulumi.Context,
	name string, args *LocalUserArgs, opts ...pulumi.ResourceOption) (*LocalUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:LocalUser"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:LocalUser"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:LocalUser"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:LocalUser"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20230101:LocalUser"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LocalUser
	err := ctx.RegisterResource("azure-native:storage/v20220901:LocalUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalUser gets an existing LocalUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalUserState, opts ...pulumi.ResourceOption) (*LocalUser, error) {
	var resource LocalUser
	err := ctx.ReadResource("azure-native:storage/v20220901:LocalUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalUser resources.
type localUserState struct {
}

type LocalUserState struct {
}

func (LocalUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*localUserState)(nil)).Elem()
}

type localUserArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// Indicates whether shared key exists. Set it to false to remove existing shared key.
	HasSharedKey *bool `pulumi:"hasSharedKey"`
	// Indicates whether ssh key exists. Set it to false to remove existing SSH key.
	HasSshKey *bool `pulumi:"hasSshKey"`
	// Indicates whether ssh password exists. Set it to false to remove existing SSH password.
	HasSshPassword *bool `pulumi:"hasSshPassword"`
	// Optional, local user home directory.
	HomeDirectory *string `pulumi:"homeDirectory"`
	// The permission scopes of the local user.
	PermissionScopes []PermissionScope `pulumi:"permissionScopes"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Optional, local user ssh authorized keys for SFTP.
	SshAuthorizedKeys []SshPublicKey `pulumi:"sshAuthorizedKeys"`
	// The name of local user. The username must contain lowercase letters and numbers only. It must be unique only within the storage account.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a LocalUser resource.
type LocalUserArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// Indicates whether shared key exists. Set it to false to remove existing shared key.
	HasSharedKey pulumi.BoolPtrInput
	// Indicates whether ssh key exists. Set it to false to remove existing SSH key.
	HasSshKey pulumi.BoolPtrInput
	// Indicates whether ssh password exists. Set it to false to remove existing SSH password.
	HasSshPassword pulumi.BoolPtrInput
	// Optional, local user home directory.
	HomeDirectory pulumi.StringPtrInput
	// The permission scopes of the local user.
	PermissionScopes PermissionScopeArrayInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Optional, local user ssh authorized keys for SFTP.
	SshAuthorizedKeys SshPublicKeyArrayInput
	// The name of local user. The username must contain lowercase letters and numbers only. It must be unique only within the storage account.
	Username pulumi.StringPtrInput
}

func (LocalUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localUserArgs)(nil)).Elem()
}

type LocalUserInput interface {
	pulumi.Input

	ToLocalUserOutput() LocalUserOutput
	ToLocalUserOutputWithContext(ctx context.Context) LocalUserOutput
}

func (*LocalUser) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalUser)(nil)).Elem()
}

func (i *LocalUser) ToLocalUserOutput() LocalUserOutput {
	return i.ToLocalUserOutputWithContext(context.Background())
}

func (i *LocalUser) ToLocalUserOutputWithContext(ctx context.Context) LocalUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalUserOutput)
}

type LocalUserOutput struct{ *pulumi.OutputState }

func (LocalUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalUser)(nil)).Elem()
}

func (o LocalUserOutput) ToLocalUserOutput() LocalUserOutput {
	return o
}

func (o LocalUserOutput) ToLocalUserOutputWithContext(ctx context.Context) LocalUserOutput {
	return o
}

// Indicates whether shared key exists. Set it to false to remove existing shared key.
func (o LocalUserOutput) HasSharedKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.BoolPtrOutput { return v.HasSharedKey }).(pulumi.BoolPtrOutput)
}

// Indicates whether ssh key exists. Set it to false to remove existing SSH key.
func (o LocalUserOutput) HasSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.BoolPtrOutput { return v.HasSshKey }).(pulumi.BoolPtrOutput)
}

// Indicates whether ssh password exists. Set it to false to remove existing SSH password.
func (o LocalUserOutput) HasSshPassword() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.BoolPtrOutput { return v.HasSshPassword }).(pulumi.BoolPtrOutput)
}

// Optional, local user home directory.
func (o LocalUserOutput) HomeDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.StringPtrOutput { return v.HomeDirectory }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LocalUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The permission scopes of the local user.
func (o LocalUserOutput) PermissionScopes() PermissionScopeResponseArrayOutput {
	return o.ApplyT(func(v *LocalUser) PermissionScopeResponseArrayOutput { return v.PermissionScopes }).(PermissionScopeResponseArrayOutput)
}

// A unique Security Identifier that is generated by the server.
func (o LocalUserOutput) Sid() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.StringOutput { return v.Sid }).(pulumi.StringOutput)
}

// Optional, local user ssh authorized keys for SFTP.
func (o LocalUserOutput) SshAuthorizedKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v *LocalUser) SshPublicKeyResponseArrayOutput { return v.SshAuthorizedKeys }).(SshPublicKeyResponseArrayOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LocalUserOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LocalUser) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LocalUserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LocalUserOutput{})
}
