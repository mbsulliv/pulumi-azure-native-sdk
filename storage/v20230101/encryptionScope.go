// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Encryption Scope resource.
type EncryptionScope struct {
	pulumi.CustomResourceState

	// Gets the creation date and time of the encryption scope in UTC.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The key vault properties for the encryption scope. This is a required field if encryption scope 'source' attribute is set to 'Microsoft.KeyVault'.
	KeyVaultProperties EncryptionScopeKeyVaultPropertiesResponsePtrOutput `pulumi:"keyVaultProperties"`
	// Gets the last modification date and time of the encryption scope in UTC.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// A boolean indicating whether or not the service applies a secondary layer of encryption with platform managed keys for data at rest.
	RequireInfrastructureEncryption pulumi.BoolPtrOutput `pulumi:"requireInfrastructureEncryption"`
	// The provider for the encryption scope. Possible values (case-insensitive):  Microsoft.Storage, Microsoft.KeyVault.
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// The state of the encryption scope. Possible values (case-insensitive):  Enabled, Disabled.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEncryptionScope registers a new resource with the given unique name, arguments, and options.
func NewEncryptionScope(ctx *pulumi.Context,
	name string, args *EncryptionScopeArgs, opts ...pulumi.ResourceOption) (*EncryptionScope, error) {
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
			Type: pulumi.String("azure-native:storage:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220901:EncryptionScope"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20230401:EncryptionScope"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EncryptionScope
	err := ctx.RegisterResource("azure-native:storage/v20230101:EncryptionScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEncryptionScope gets an existing EncryptionScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEncryptionScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EncryptionScopeState, opts ...pulumi.ResourceOption) (*EncryptionScope, error) {
	var resource EncryptionScope
	err := ctx.ReadResource("azure-native:storage/v20230101:EncryptionScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EncryptionScope resources.
type encryptionScopeState struct {
}

type EncryptionScopeState struct {
}

func (EncryptionScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptionScopeState)(nil)).Elem()
}

type encryptionScopeArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the encryption scope within the specified storage account. Encryption scope names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	EncryptionScopeName *string `pulumi:"encryptionScopeName"`
	// The key vault properties for the encryption scope. This is a required field if encryption scope 'source' attribute is set to 'Microsoft.KeyVault'.
	KeyVaultProperties *EncryptionScopeKeyVaultProperties `pulumi:"keyVaultProperties"`
	// A boolean indicating whether or not the service applies a secondary layer of encryption with platform managed keys for data at rest.
	RequireInfrastructureEncryption *bool `pulumi:"requireInfrastructureEncryption"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The provider for the encryption scope. Possible values (case-insensitive):  Microsoft.Storage, Microsoft.KeyVault.
	Source *string `pulumi:"source"`
	// The state of the encryption scope. Possible values (case-insensitive):  Enabled, Disabled.
	State *string `pulumi:"state"`
}

// The set of arguments for constructing a EncryptionScope resource.
type EncryptionScopeArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// The name of the encryption scope within the specified storage account. Encryption scope names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	EncryptionScopeName pulumi.StringPtrInput
	// The key vault properties for the encryption scope. This is a required field if encryption scope 'source' attribute is set to 'Microsoft.KeyVault'.
	KeyVaultProperties EncryptionScopeKeyVaultPropertiesPtrInput
	// A boolean indicating whether or not the service applies a secondary layer of encryption with platform managed keys for data at rest.
	RequireInfrastructureEncryption pulumi.BoolPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The provider for the encryption scope. Possible values (case-insensitive):  Microsoft.Storage, Microsoft.KeyVault.
	Source pulumi.StringPtrInput
	// The state of the encryption scope. Possible values (case-insensitive):  Enabled, Disabled.
	State pulumi.StringPtrInput
}

func (EncryptionScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptionScopeArgs)(nil)).Elem()
}

type EncryptionScopeInput interface {
	pulumi.Input

	ToEncryptionScopeOutput() EncryptionScopeOutput
	ToEncryptionScopeOutputWithContext(ctx context.Context) EncryptionScopeOutput
}

func (*EncryptionScope) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionScope)(nil)).Elem()
}

func (i *EncryptionScope) ToEncryptionScopeOutput() EncryptionScopeOutput {
	return i.ToEncryptionScopeOutputWithContext(context.Background())
}

func (i *EncryptionScope) ToEncryptionScopeOutputWithContext(ctx context.Context) EncryptionScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionScopeOutput)
}

type EncryptionScopeOutput struct{ *pulumi.OutputState }

func (EncryptionScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionScope)(nil)).Elem()
}

func (o EncryptionScopeOutput) ToEncryptionScopeOutput() EncryptionScopeOutput {
	return o
}

func (o EncryptionScopeOutput) ToEncryptionScopeOutputWithContext(ctx context.Context) EncryptionScopeOutput {
	return o
}

// Gets the creation date and time of the encryption scope in UTC.
func (o EncryptionScopeOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The key vault properties for the encryption scope. This is a required field if encryption scope 'source' attribute is set to 'Microsoft.KeyVault'.
func (o EncryptionScopeOutput) KeyVaultProperties() EncryptionScopeKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionScope) EncryptionScopeKeyVaultPropertiesResponsePtrOutput {
		return v.KeyVaultProperties
	}).(EncryptionScopeKeyVaultPropertiesResponsePtrOutput)
}

// Gets the last modification date and time of the encryption scope in UTC.
func (o EncryptionScopeOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o EncryptionScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A boolean indicating whether or not the service applies a secondary layer of encryption with platform managed keys for data at rest.
func (o EncryptionScopeOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.BoolPtrOutput { return v.RequireInfrastructureEncryption }).(pulumi.BoolPtrOutput)
}

// The provider for the encryption scope. Possible values (case-insensitive):  Microsoft.Storage, Microsoft.KeyVault.
func (o EncryptionScopeOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

// The state of the encryption scope. Possible values (case-insensitive):  Enabled, Disabled.
func (o EncryptionScopeOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EncryptionScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EncryptionScopeOutput{})
}
