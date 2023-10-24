// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A managed instance key.
// Azure REST API version: 2021-11-01. Prior API version in Azure Native 1.x: 2020-11-01-preview.
//
// Other available API versions: 2022-11-01-preview, 2023-02-01-preview.
type ManagedInstanceKey struct {
	pulumi.CustomResourceState

	// Key auto rotation opt-in flag. Either true or false.
	AutoRotationEnabled pulumi.BoolOutput `pulumi:"autoRotationEnabled"`
	// The key creation date.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Kind of encryption protector. This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Thumbprint of the key.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedInstanceKey registers a new resource with the given unique name, arguments, and options.
func NewManagedInstanceKey(ctx *pulumi.Context,
	name string, args *ManagedInstanceKeyArgs, opts ...pulumi.ResourceOption) (*ManagedInstanceKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerKeyType == nil {
		return nil, errors.New("invalid value for required argument 'ServerKeyType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql/v20171001preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ManagedInstanceKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:ManagedInstanceKey"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedInstanceKey
	err := ctx.RegisterResource("azure-native:sql:ManagedInstanceKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedInstanceKey gets an existing ManagedInstanceKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedInstanceKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstanceKeyState, opts ...pulumi.ResourceOption) (*ManagedInstanceKey, error) {
	var resource ManagedInstanceKey
	err := ctx.ReadResource("azure-native:sql:ManagedInstanceKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedInstanceKey resources.
type managedInstanceKeyState struct {
}

type ManagedInstanceKeyState struct {
}

func (ManagedInstanceKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceKeyState)(nil)).Elem()
}

type managedInstanceKeyArgs struct {
	// The name of the managed instance key to be operated on (updated or created).
	KeyName *string `pulumi:"keyName"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The key type like 'ServiceManaged', 'AzureKeyVault'.
	ServerKeyType string `pulumi:"serverKeyType"`
	// The URI of the key. If the ServerKeyType is AzureKeyVault, then the URI is required.
	Uri *string `pulumi:"uri"`
}

// The set of arguments for constructing a ManagedInstanceKey resource.
type ManagedInstanceKeyArgs struct {
	// The name of the managed instance key to be operated on (updated or created).
	KeyName pulumi.StringPtrInput
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The key type like 'ServiceManaged', 'AzureKeyVault'.
	ServerKeyType pulumi.StringInput
	// The URI of the key. If the ServerKeyType is AzureKeyVault, then the URI is required.
	Uri pulumi.StringPtrInput
}

func (ManagedInstanceKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceKeyArgs)(nil)).Elem()
}

type ManagedInstanceKeyInput interface {
	pulumi.Input

	ToManagedInstanceKeyOutput() ManagedInstanceKeyOutput
	ToManagedInstanceKeyOutputWithContext(ctx context.Context) ManagedInstanceKeyOutput
}

func (*ManagedInstanceKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceKey)(nil)).Elem()
}

func (i *ManagedInstanceKey) ToManagedInstanceKeyOutput() ManagedInstanceKeyOutput {
	return i.ToManagedInstanceKeyOutputWithContext(context.Background())
}

func (i *ManagedInstanceKey) ToManagedInstanceKeyOutputWithContext(ctx context.Context) ManagedInstanceKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceKeyOutput)
}

func (i *ManagedInstanceKey) ToOutput(ctx context.Context) pulumix.Output[*ManagedInstanceKey] {
	return pulumix.Output[*ManagedInstanceKey]{
		OutputState: i.ToManagedInstanceKeyOutputWithContext(ctx).OutputState,
	}
}

type ManagedInstanceKeyOutput struct{ *pulumi.OutputState }

func (ManagedInstanceKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceKey)(nil)).Elem()
}

func (o ManagedInstanceKeyOutput) ToManagedInstanceKeyOutput() ManagedInstanceKeyOutput {
	return o
}

func (o ManagedInstanceKeyOutput) ToManagedInstanceKeyOutputWithContext(ctx context.Context) ManagedInstanceKeyOutput {
	return o
}

func (o ManagedInstanceKeyOutput) ToOutput(ctx context.Context) pulumix.Output[*ManagedInstanceKey] {
	return pulumix.Output[*ManagedInstanceKey]{
		OutputState: o.OutputState,
	}
}

// Key auto rotation opt-in flag. Either true or false.
func (o ManagedInstanceKeyOutput) AutoRotationEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ManagedInstanceKey) pulumi.BoolOutput { return v.AutoRotationEnabled }).(pulumi.BoolOutput)
}

// The key creation date.
func (o ManagedInstanceKeyOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceKey) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// Kind of encryption protector. This is metadata used for the Azure portal experience.
func (o ManagedInstanceKeyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceKey) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource name.
func (o ManagedInstanceKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Thumbprint of the key.
func (o ManagedInstanceKeyOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceKey) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

// Resource type.
func (o ManagedInstanceKeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceKey) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedInstanceKeyOutput{})
}
