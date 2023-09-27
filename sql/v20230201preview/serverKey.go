// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A server key.
type ServerKey struct {
	pulumi.CustomResourceState

	// Key auto rotation opt-in flag. Either true or false.
	AutoRotationEnabled pulumi.BoolOutput `pulumi:"autoRotationEnabled"`
	// The server key creation date.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Kind of encryption protector. This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Subregion of the server key.
	Subregion pulumi.StringOutput `pulumi:"subregion"`
	// Thumbprint of the server key.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServerKey registers a new resource with the given unique name, arguments, and options.
func NewServerKey(ctx *pulumi.Context,
	name string, args *ServerKeyArgs, opts ...pulumi.ResourceOption) (*ServerKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerKeyType == nil {
		return nil, errors.New("invalid value for required argument 'ServerKeyType'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ServerKey"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServerKey
	err := ctx.RegisterResource("azure-native:sql/v20230201preview:ServerKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerKey gets an existing ServerKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerKeyState, opts ...pulumi.ResourceOption) (*ServerKey, error) {
	var resource ServerKey
	err := ctx.ReadResource("azure-native:sql/v20230201preview:ServerKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerKey resources.
type serverKeyState struct {
}

type ServerKeyState struct {
}

func (ServerKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverKeyState)(nil)).Elem()
}

type serverKeyArgs struct {
	// The name of the server key to be operated on (updated or created). The key name is required to be in the format of 'vault_key_version'. For example, if the keyId is https://YourVaultName.vault.azure.net/keys/YourKeyName/YourKeyVersion, then the server key name should be formatted as: YourVaultName_YourKeyName_YourKeyVersion
	KeyName *string `pulumi:"keyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The server key type like 'ServiceManaged', 'AzureKeyVault'.
	ServerKeyType string `pulumi:"serverKeyType"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The URI of the server key. If the ServerKeyType is AzureKeyVault, then the URI is required. The AKV URI is required to be in this format: 'https://YourVaultName.vault.azure.net/keys/YourKeyName/YourKeyVersion'
	Uri *string `pulumi:"uri"`
}

// The set of arguments for constructing a ServerKey resource.
type ServerKeyArgs struct {
	// The name of the server key to be operated on (updated or created). The key name is required to be in the format of 'vault_key_version'. For example, if the keyId is https://YourVaultName.vault.azure.net/keys/YourKeyName/YourKeyVersion, then the server key name should be formatted as: YourVaultName_YourKeyName_YourKeyVersion
	KeyName pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The server key type like 'ServiceManaged', 'AzureKeyVault'.
	ServerKeyType pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The URI of the server key. If the ServerKeyType is AzureKeyVault, then the URI is required. The AKV URI is required to be in this format: 'https://YourVaultName.vault.azure.net/keys/YourKeyName/YourKeyVersion'
	Uri pulumi.StringPtrInput
}

func (ServerKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverKeyArgs)(nil)).Elem()
}

type ServerKeyInput interface {
	pulumi.Input

	ToServerKeyOutput() ServerKeyOutput
	ToServerKeyOutputWithContext(ctx context.Context) ServerKeyOutput
}

func (*ServerKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerKey)(nil)).Elem()
}

func (i *ServerKey) ToServerKeyOutput() ServerKeyOutput {
	return i.ToServerKeyOutputWithContext(context.Background())
}

func (i *ServerKey) ToServerKeyOutputWithContext(ctx context.Context) ServerKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerKeyOutput)
}

func (i *ServerKey) ToOutput(ctx context.Context) pulumix.Output[*ServerKey] {
	return pulumix.Output[*ServerKey]{
		OutputState: i.ToServerKeyOutputWithContext(ctx).OutputState,
	}
}

type ServerKeyOutput struct{ *pulumi.OutputState }

func (ServerKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerKey)(nil)).Elem()
}

func (o ServerKeyOutput) ToServerKeyOutput() ServerKeyOutput {
	return o
}

func (o ServerKeyOutput) ToServerKeyOutputWithContext(ctx context.Context) ServerKeyOutput {
	return o
}

func (o ServerKeyOutput) ToOutput(ctx context.Context) pulumix.Output[*ServerKey] {
	return pulumix.Output[*ServerKey]{
		OutputState: o.OutputState,
	}
}

// Key auto rotation opt-in flag. Either true or false.
func (o ServerKeyOutput) AutoRotationEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.BoolOutput { return v.AutoRotationEnabled }).(pulumi.BoolOutput)
}

// The server key creation date.
func (o ServerKeyOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// Kind of encryption protector. This is metadata used for the Azure portal experience.
func (o ServerKeyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o ServerKeyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o ServerKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Subregion of the server key.
func (o ServerKeyOutput) Subregion() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.Subregion }).(pulumi.StringOutput)
}

// Thumbprint of the server key.
func (o ServerKeyOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

// Resource type.
func (o ServerKeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerKeyOutput{})
}
