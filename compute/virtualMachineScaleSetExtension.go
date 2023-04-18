// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a Virtual Machine Scale Set Extension.
// API Version: 2021-03-01.
type VirtualMachineScaleSetExtension struct {
	pulumi.CustomResourceState

	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrOutput `pulumi:"autoUpgradeMinorVersion"`
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade pulumi.BoolPtrOutput `pulumi:"enableAutomaticUpgrade"`
	// If a value is provided and is different from the previous value, the extension handler will be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrOutput `pulumi:"forceUpdateTag"`
	// The name of the extension.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.AnyOutput `pulumi:"protectedSettings"`
	// Collection of extension names after which this extension needs to be provisioned.
	ProvisionAfterExtensions pulumi.StringArrayOutput `pulumi:"provisionAfterExtensions"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrOutput `pulumi:"publisher"`
	// Json formatted public settings for the extension.
	Settings pulumi.AnyOutput `pulumi:"settings"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrOutput `pulumi:"typeHandlerVersion"`
}

// NewVirtualMachineScaleSetExtension registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineScaleSetExtension(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetExtensionArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetExtension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmScaleSetName == nil {
		return nil, errors.New("invalid value for required argument 'VmScaleSetName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20170330:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230301:VirtualMachineScaleSetExtension"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineScaleSetExtension
	err := ctx.RegisterResource("azure-native:compute:VirtualMachineScaleSetExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineScaleSetExtension gets an existing VirtualMachineScaleSetExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineScaleSetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetExtensionState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetExtension, error) {
	var resource VirtualMachineScaleSetExtension
	err := ctx.ReadResource("azure-native:compute:VirtualMachineScaleSetExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineScaleSetExtension resources.
type virtualMachineScaleSetExtensionState struct {
}

type VirtualMachineScaleSetExtensionState struct {
}

func (VirtualMachineScaleSetExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetExtensionState)(nil)).Elem()
}

type virtualMachineScaleSetExtensionArgs struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade *bool `pulumi:"enableAutomaticUpgrade"`
	// If a value is provided and is different from the previous value, the extension handler will be forced to update even if the extension configuration has not changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The name of the extension.
	Name *string `pulumi:"name"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings interface{} `pulumi:"protectedSettings"`
	// Collection of extension names after which this extension needs to be provisioned.
	ProvisionAfterExtensions []string `pulumi:"provisionAfterExtensions"`
	// The name of the extension handler publisher.
	Publisher *string `pulumi:"publisher"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Json formatted public settings for the extension.
	Settings interface{} `pulumi:"settings"`
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type *string `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
	// The name of the VM scale set where the extension should be create or updated.
	VmScaleSetName string `pulumi:"vmScaleSetName"`
	// The name of the VM scale set extension.
	VmssExtensionName *string `pulumi:"vmssExtensionName"`
}

// The set of arguments for constructing a VirtualMachineScaleSetExtension resource.
type VirtualMachineScaleSetExtensionArgs struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade pulumi.BoolPtrInput
	// If a value is provided and is different from the previous value, the extension handler will be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrInput
	// The name of the extension.
	Name pulumi.StringPtrInput
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.Input
	// Collection of extension names after which this extension needs to be provisioned.
	ProvisionAfterExtensions pulumi.StringArrayInput
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Json formatted public settings for the extension.
	Settings pulumi.Input
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type pulumi.StringPtrInput
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrInput
	// The name of the VM scale set where the extension should be create or updated.
	VmScaleSetName pulumi.StringInput
	// The name of the VM scale set extension.
	VmssExtensionName pulumi.StringPtrInput
}

func (VirtualMachineScaleSetExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetExtensionArgs)(nil)).Elem()
}

type VirtualMachineScaleSetExtensionInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionOutput() VirtualMachineScaleSetExtensionOutput
	ToVirtualMachineScaleSetExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionOutput
}

func (*VirtualMachineScaleSetExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtension)(nil)).Elem()
}

func (i *VirtualMachineScaleSetExtension) ToVirtualMachineScaleSetExtensionOutput() VirtualMachineScaleSetExtensionOutput {
	return i.ToVirtualMachineScaleSetExtensionOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSetExtension) ToVirtualMachineScaleSetExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionOutput)
}

type VirtualMachineScaleSetExtensionOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtension)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionOutput) ToVirtualMachineScaleSetExtensionOutput() VirtualMachineScaleSetExtensionOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionOutput) ToVirtualMachineScaleSetExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionOutput {
	return o
}

// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
func (o VirtualMachineScaleSetExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.BoolPtrOutput { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
func (o VirtualMachineScaleSetExtensionOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.BoolPtrOutput { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

// If a value is provided and is different from the previous value, the extension handler will be forced to update even if the extension configuration has not changed.
func (o VirtualMachineScaleSetExtensionOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

// The name of the extension.
func (o VirtualMachineScaleSetExtensionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
func (o VirtualMachineScaleSetExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.AnyOutput { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

// Collection of extension names after which this extension needs to be provisioned.
func (o VirtualMachineScaleSetExtensionOutput) ProvisionAfterExtensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringArrayOutput { return v.ProvisionAfterExtensions }).(pulumi.StringArrayOutput)
}

// The provisioning state, which only appears in the response.
func (o VirtualMachineScaleSetExtensionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The name of the extension handler publisher.
func (o VirtualMachineScaleSetExtensionOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringPtrOutput { return v.Publisher }).(pulumi.StringPtrOutput)
}

// Json formatted public settings for the extension.
func (o VirtualMachineScaleSetExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.AnyOutput { return v.Settings }).(pulumi.AnyOutput)
}

// Resource type
func (o VirtualMachineScaleSetExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the version of the script handler.
func (o VirtualMachineScaleSetExtensionOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringPtrOutput { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionOutput{})
}
