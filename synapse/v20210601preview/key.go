// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A workspace key
type Key struct {
	pulumi.CustomResourceState

	// Used to activate the workspace after a customer managed key is provided.
	IsActiveCMK pulumi.BoolPtrOutput `pulumi:"isActiveCMK"`
	// The Key Vault Url of the workspace key.
	KeyVaultUrl pulumi.StringPtrOutput `pulumi:"keyVaultUrl"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:Key"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:Key"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:Key"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:Key"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:Key"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:Key"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:Key"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Key
	err := ctx.RegisterResource("azure-native:synapse/v20210601preview:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("azure-native:synapse/v20210601preview:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
}

type KeyState struct {
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// Used to activate the workspace after a customer managed key is provided.
	IsActiveCMK *bool `pulumi:"isActiveCMK"`
	// The name of the workspace key
	KeyName *string `pulumi:"keyName"`
	// The Key Vault Url of the workspace key.
	KeyVaultUrl *string `pulumi:"keyVaultUrl"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// Used to activate the workspace after a customer managed key is provided.
	IsActiveCMK pulumi.BoolPtrInput
	// The name of the workspace key
	KeyName pulumi.StringPtrInput
	// The Key Vault Url of the workspace key.
	KeyVaultUrl pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

func (i *Key) ToOutput(ctx context.Context) pulumix.Output[*Key] {
	return pulumix.Output[*Key]{
		OutputState: i.ToKeyOutputWithContext(ctx).OutputState,
	}
}

type KeyOutput struct{ *pulumi.OutputState }

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

func (o KeyOutput) ToOutput(ctx context.Context) pulumix.Output[*Key] {
	return pulumix.Output[*Key]{
		OutputState: o.OutputState,
	}
}

// Used to activate the workspace after a customer managed key is provided.
func (o KeyOutput) IsActiveCMK() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.BoolPtrOutput { return v.IsActiveCMK }).(pulumi.BoolPtrOutput)
}

// The Key Vault Url of the workspace key.
func (o KeyOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.KeyVaultUrl }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o KeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o KeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyOutput{})
}
