// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Describes a VMSS VM Extension.
type VirtualMachineScaleSetVMExtension struct {
	pulumi.CustomResourceState

	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrOutput `pulumi:"autoUpgradeMinorVersion"`
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade pulumi.BoolPtrOutput `pulumi:"enableAutomaticUpgrade"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrOutput `pulumi:"forceUpdateTag"`
	// The virtual machine extension instance view.
	InstanceView VirtualMachineExtensionInstanceViewResponsePtrOutput `pulumi:"instanceView"`
	// The location of the extension.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the extension.
	Name pulumi.StringOutput `pulumi:"name"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.AnyOutput `pulumi:"protectedSettings"`
	// The extensions protected settings that are passed by reference, and consumed from key vault
	ProtectedSettingsFromKeyVault KeyVaultSecretReferenceResponsePtrOutput `pulumi:"protectedSettingsFromKeyVault"`
	// Collection of extension names after which this extension needs to be provisioned.
	ProvisionAfterExtensions pulumi.StringArrayOutput `pulumi:"provisionAfterExtensions"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrOutput `pulumi:"publisher"`
	// Json formatted public settings for the extension.
	Settings pulumi.AnyOutput `pulumi:"settings"`
	// Indicates whether failures stemming from the extension will be suppressed (Operational failures such as not connecting to the VM will not be suppressed regardless of this value). The default is false.
	SuppressFailures pulumi.BoolPtrOutput `pulumi:"suppressFailures"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrOutput `pulumi:"typeHandlerVersion"`
}

// NewVirtualMachineScaleSetVMExtension registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineScaleSetVMExtension(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetVMExtensionArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVMExtension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmScaleSetName == nil {
		return nil, errors.New("invalid value for required argument 'VmScaleSetName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230701:VirtualMachineScaleSetVMExtension"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualMachineScaleSetVMExtension
	err := ctx.RegisterResource("azure-native:compute/v20230301:VirtualMachineScaleSetVMExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineScaleSetVMExtension gets an existing VirtualMachineScaleSetVMExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineScaleSetVMExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetVMExtensionState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVMExtension, error) {
	var resource VirtualMachineScaleSetVMExtension
	err := ctx.ReadResource("azure-native:compute/v20230301:VirtualMachineScaleSetVMExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineScaleSetVMExtension resources.
type virtualMachineScaleSetVMExtensionState struct {
}

type VirtualMachineScaleSetVMExtensionState struct {
}

func (VirtualMachineScaleSetVMExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMExtensionState)(nil)).Elem()
}

type virtualMachineScaleSetVMExtensionArgs struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade *bool `pulumi:"enableAutomaticUpgrade"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The instance ID of the virtual machine.
	InstanceId string `pulumi:"instanceId"`
	// The virtual machine extension instance view.
	InstanceView *VirtualMachineExtensionInstanceView `pulumi:"instanceView"`
	// The location of the extension.
	Location *string `pulumi:"location"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings interface{} `pulumi:"protectedSettings"`
	// The extensions protected settings that are passed by reference, and consumed from key vault
	ProtectedSettingsFromKeyVault *KeyVaultSecretReference `pulumi:"protectedSettingsFromKeyVault"`
	// Collection of extension names after which this extension needs to be provisioned.
	ProvisionAfterExtensions []string `pulumi:"provisionAfterExtensions"`
	// The name of the extension handler publisher.
	Publisher *string `pulumi:"publisher"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Json formatted public settings for the extension.
	Settings interface{} `pulumi:"settings"`
	// Indicates whether failures stemming from the extension will be suppressed (Operational failures such as not connecting to the VM will not be suppressed regardless of this value). The default is false.
	SuppressFailures *bool `pulumi:"suppressFailures"`
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type *string `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
	// The name of the virtual machine extension.
	VmExtensionName *string `pulumi:"vmExtensionName"`
	// The name of the VM scale set.
	VmScaleSetName string `pulumi:"vmScaleSetName"`
}

// The set of arguments for constructing a VirtualMachineScaleSetVMExtension resource.
type VirtualMachineScaleSetVMExtensionArgs struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade pulumi.BoolPtrInput
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrInput
	// The instance ID of the virtual machine.
	InstanceId pulumi.StringInput
	// The virtual machine extension instance view.
	InstanceView VirtualMachineExtensionInstanceViewPtrInput
	// The location of the extension.
	Location pulumi.StringPtrInput
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.Input
	// The extensions protected settings that are passed by reference, and consumed from key vault
	ProtectedSettingsFromKeyVault KeyVaultSecretReferencePtrInput
	// Collection of extension names after which this extension needs to be provisioned.
	ProvisionAfterExtensions pulumi.StringArrayInput
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Json formatted public settings for the extension.
	Settings pulumi.Input
	// Indicates whether failures stemming from the extension will be suppressed (Operational failures such as not connecting to the VM will not be suppressed regardless of this value). The default is false.
	SuppressFailures pulumi.BoolPtrInput
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type pulumi.StringPtrInput
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrInput
	// The name of the virtual machine extension.
	VmExtensionName pulumi.StringPtrInput
	// The name of the VM scale set.
	VmScaleSetName pulumi.StringInput
}

func (VirtualMachineScaleSetVMExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMExtensionArgs)(nil)).Elem()
}

type VirtualMachineScaleSetVMExtensionInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMExtensionOutput() VirtualMachineScaleSetVMExtensionOutput
	ToVirtualMachineScaleSetVMExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMExtensionOutput
}

func (*VirtualMachineScaleSetVMExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMExtension)(nil)).Elem()
}

func (i *VirtualMachineScaleSetVMExtension) ToVirtualMachineScaleSetVMExtensionOutput() VirtualMachineScaleSetVMExtensionOutput {
	return i.ToVirtualMachineScaleSetVMExtensionOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSetVMExtension) ToVirtualMachineScaleSetVMExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMExtensionOutput)
}

func (i *VirtualMachineScaleSetVMExtension) ToOutput(ctx context.Context) pulumix.Output[*VirtualMachineScaleSetVMExtension] {
	return pulumix.Output[*VirtualMachineScaleSetVMExtension]{
		OutputState: i.ToVirtualMachineScaleSetVMExtensionOutputWithContext(ctx).OutputState,
	}
}

type VirtualMachineScaleSetVMExtensionOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMExtension)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMExtensionOutput) ToVirtualMachineScaleSetVMExtensionOutput() VirtualMachineScaleSetVMExtensionOutput {
	return o
}

func (o VirtualMachineScaleSetVMExtensionOutput) ToVirtualMachineScaleSetVMExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMExtensionOutput {
	return o
}

func (o VirtualMachineScaleSetVMExtensionOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualMachineScaleSetVMExtension] {
	return pulumix.Output[*VirtualMachineScaleSetVMExtension]{
		OutputState: o.OutputState,
	}
}

// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
func (o VirtualMachineScaleSetVMExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.BoolPtrOutput { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
func (o VirtualMachineScaleSetVMExtensionOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.BoolPtrOutput { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

// How the extension handler should be forced to update even if the extension configuration has not changed.
func (o VirtualMachineScaleSetVMExtensionOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

// The virtual machine extension instance view.
func (o VirtualMachineScaleSetVMExtensionOutput) InstanceView() VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) VirtualMachineExtensionInstanceViewResponsePtrOutput {
		return v.InstanceView
	}).(VirtualMachineExtensionInstanceViewResponsePtrOutput)
}

// The location of the extension.
func (o VirtualMachineScaleSetVMExtensionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the extension.
func (o VirtualMachineScaleSetVMExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
func (o VirtualMachineScaleSetVMExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.AnyOutput { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

// The extensions protected settings that are passed by reference, and consumed from key vault
func (o VirtualMachineScaleSetVMExtensionOutput) ProtectedSettingsFromKeyVault() KeyVaultSecretReferenceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) KeyVaultSecretReferenceResponsePtrOutput {
		return v.ProtectedSettingsFromKeyVault
	}).(KeyVaultSecretReferenceResponsePtrOutput)
}

// Collection of extension names after which this extension needs to be provisioned.
func (o VirtualMachineScaleSetVMExtensionOutput) ProvisionAfterExtensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringArrayOutput { return v.ProvisionAfterExtensions }).(pulumi.StringArrayOutput)
}

// The provisioning state, which only appears in the response.
func (o VirtualMachineScaleSetVMExtensionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The name of the extension handler publisher.
func (o VirtualMachineScaleSetVMExtensionOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringPtrOutput { return v.Publisher }).(pulumi.StringPtrOutput)
}

// Json formatted public settings for the extension.
func (o VirtualMachineScaleSetVMExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.AnyOutput { return v.Settings }).(pulumi.AnyOutput)
}

// Indicates whether failures stemming from the extension will be suppressed (Operational failures such as not connecting to the VM will not be suppressed regardless of this value). The default is false.
func (o VirtualMachineScaleSetVMExtensionOutput) SuppressFailures() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.BoolPtrOutput { return v.SuppressFailures }).(pulumi.BoolPtrOutput)
}

// Resource type
func (o VirtualMachineScaleSetVMExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the version of the script handler.
func (o VirtualMachineScaleSetVMExtensionOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringPtrOutput { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMExtensionOutput{})
}
