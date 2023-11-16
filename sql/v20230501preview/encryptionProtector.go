// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The server encryption protector.
type EncryptionProtector struct {
	pulumi.CustomResourceState

	// Key auto rotation opt-in flag. Either true or false.
	AutoRotationEnabled pulumi.BoolPtrOutput `pulumi:"autoRotationEnabled"`
	// Kind of encryption protector. This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the server key.
	ServerKeyName pulumi.StringPtrOutput `pulumi:"serverKeyName"`
	// The encryption protector type like 'ServiceManaged', 'AzureKeyVault'.
	ServerKeyType pulumi.StringOutput `pulumi:"serverKeyType"`
	// Subregion of the encryption protector.
	Subregion pulumi.StringOutput `pulumi:"subregion"`
	// Thumbprint of the server key.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The URI of the server key.
	Uri pulumi.StringOutput `pulumi:"uri"`
}

// NewEncryptionProtector registers a new resource with the given unique name, arguments, and options.
func NewEncryptionProtector(ctx *pulumi.Context,
	name string, args *EncryptionProtectorArgs, opts ...pulumi.ResourceOption) (*EncryptionProtector, error) {
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
			Type: pulumi.String("azure-native:sql:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:EncryptionProtector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EncryptionProtector
	err := ctx.RegisterResource("azure-native:sql/v20230501preview:EncryptionProtector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEncryptionProtector gets an existing EncryptionProtector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEncryptionProtector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EncryptionProtectorState, opts ...pulumi.ResourceOption) (*EncryptionProtector, error) {
	var resource EncryptionProtector
	err := ctx.ReadResource("azure-native:sql/v20230501preview:EncryptionProtector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EncryptionProtector resources.
type encryptionProtectorState struct {
}

type EncryptionProtectorState struct {
}

func (EncryptionProtectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptionProtectorState)(nil)).Elem()
}

type encryptionProtectorArgs struct {
	// Key auto rotation opt-in flag. Either true or false.
	AutoRotationEnabled *bool `pulumi:"autoRotationEnabled"`
	// The name of the encryption protector to be updated.
	EncryptionProtectorName *string `pulumi:"encryptionProtectorName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server key.
	ServerKeyName *string `pulumi:"serverKeyName"`
	// The encryption protector type like 'ServiceManaged', 'AzureKeyVault'.
	ServerKeyType string `pulumi:"serverKeyType"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// The set of arguments for constructing a EncryptionProtector resource.
type EncryptionProtectorArgs struct {
	// Key auto rotation opt-in flag. Either true or false.
	AutoRotationEnabled pulumi.BoolPtrInput
	// The name of the encryption protector to be updated.
	EncryptionProtectorName pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server key.
	ServerKeyName pulumi.StringPtrInput
	// The encryption protector type like 'ServiceManaged', 'AzureKeyVault'.
	ServerKeyType pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
}

func (EncryptionProtectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptionProtectorArgs)(nil)).Elem()
}

type EncryptionProtectorInput interface {
	pulumi.Input

	ToEncryptionProtectorOutput() EncryptionProtectorOutput
	ToEncryptionProtectorOutputWithContext(ctx context.Context) EncryptionProtectorOutput
}

func (*EncryptionProtector) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProtector)(nil)).Elem()
}

func (i *EncryptionProtector) ToEncryptionProtectorOutput() EncryptionProtectorOutput {
	return i.ToEncryptionProtectorOutputWithContext(context.Background())
}

func (i *EncryptionProtector) ToEncryptionProtectorOutputWithContext(ctx context.Context) EncryptionProtectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionProtectorOutput)
}

type EncryptionProtectorOutput struct{ *pulumi.OutputState }

func (EncryptionProtectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProtector)(nil)).Elem()
}

func (o EncryptionProtectorOutput) ToEncryptionProtectorOutput() EncryptionProtectorOutput {
	return o
}

func (o EncryptionProtectorOutput) ToEncryptionProtectorOutputWithContext(ctx context.Context) EncryptionProtectorOutput {
	return o
}

// Key auto rotation opt-in flag. Either true or false.
func (o EncryptionProtectorOutput) AutoRotationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionProtector) pulumi.BoolPtrOutput { return v.AutoRotationEnabled }).(pulumi.BoolPtrOutput)
}

// Kind of encryption protector. This is metadata used for the Azure portal experience.
func (o EncryptionProtectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionProtector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o EncryptionProtectorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionProtector) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o EncryptionProtectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionProtector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the server key.
func (o EncryptionProtectorOutput) ServerKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionProtector) pulumi.StringPtrOutput { return v.ServerKeyName }).(pulumi.StringPtrOutput)
}

// The encryption protector type like 'ServiceManaged', 'AzureKeyVault'.
func (o EncryptionProtectorOutput) ServerKeyType() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionProtector) pulumi.StringOutput { return v.ServerKeyType }).(pulumi.StringOutput)
}

// Subregion of the encryption protector.
func (o EncryptionProtectorOutput) Subregion() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionProtector) pulumi.StringOutput { return v.Subregion }).(pulumi.StringOutput)
}

// Thumbprint of the server key.
func (o EncryptionProtectorOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionProtector) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

// Resource type.
func (o EncryptionProtectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionProtector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The URI of the server key.
func (o EncryptionProtectorOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptionProtector) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EncryptionProtectorOutput{})
}
