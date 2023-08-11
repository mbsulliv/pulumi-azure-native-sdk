// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a Virtual Machine Extension.
type VirtualMachineExtension struct {
	pulumi.CustomResourceState

	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrOutput `pulumi:"autoUpgradeMinorVersion"`
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade pulumi.BoolPtrOutput `pulumi:"enableAutomaticUpgrade"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrOutput `pulumi:"forceUpdateTag"`
	// The virtual machine extension instance view.
	InstanceView VirtualMachineExtensionInstanceViewResponsePtrOutput `pulumi:"instanceView"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.AnyOutput `pulumi:"protectedSettings"`
	// The extensions protected settings that are passed by reference, and consumed from key vault
	ProtectedSettingsFromKeyVault pulumi.AnyOutput `pulumi:"protectedSettingsFromKeyVault"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrOutput `pulumi:"publisher"`
	// Json formatted public settings for the extension.
	Settings pulumi.AnyOutput `pulumi:"settings"`
	// Indicates whether failures stemming from the extension will be suppressed (Operational failures such as not connecting to the VM will not be suppressed regardless of this value). The default is false.
	SuppressFailures pulumi.BoolPtrOutput `pulumi:"suppressFailures"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrOutput `pulumi:"typeHandlerVersion"`
}

// NewVirtualMachineExtension registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineExtension(ctx *pulumi.Context,
	name string, args *VirtualMachineExtensionArgs, opts ...pulumi.ResourceOption) (*VirtualMachineExtension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmName == nil {
		return nil, errors.New("invalid value for required argument 'VmName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20150615:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160330:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230301:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230701:VirtualMachineExtension"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineExtension
	err := ctx.RegisterResource("azure-native:compute/v20211101:VirtualMachineExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineExtension gets an existing VirtualMachineExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineExtensionState, opts ...pulumi.ResourceOption) (*VirtualMachineExtension, error) {
	var resource VirtualMachineExtension
	err := ctx.ReadResource("azure-native:compute/v20211101:VirtualMachineExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineExtension resources.
type virtualMachineExtensionState struct {
}

type VirtualMachineExtensionState struct {
}

func (VirtualMachineExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineExtensionState)(nil)).Elem()
}

type virtualMachineExtensionArgs struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade *bool `pulumi:"enableAutomaticUpgrade"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The virtual machine extension instance view.
	InstanceView *VirtualMachineExtensionInstanceView `pulumi:"instanceView"`
	// Resource location
	Location *string `pulumi:"location"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings interface{} `pulumi:"protectedSettings"`
	// The extensions protected settings that are passed by reference, and consumed from key vault
	ProtectedSettingsFromKeyVault interface{} `pulumi:"protectedSettingsFromKeyVault"`
	// The name of the extension handler publisher.
	Publisher *string `pulumi:"publisher"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Json formatted public settings for the extension.
	Settings interface{} `pulumi:"settings"`
	// Indicates whether failures stemming from the extension will be suppressed (Operational failures such as not connecting to the VM will not be suppressed regardless of this value). The default is false.
	SuppressFailures *bool `pulumi:"suppressFailures"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type *string `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
	// The name of the virtual machine extension.
	VmExtensionName *string `pulumi:"vmExtensionName"`
	// The name of the virtual machine where the extension should be created or updated.
	VmName string `pulumi:"vmName"`
}

// The set of arguments for constructing a VirtualMachineExtension resource.
type VirtualMachineExtensionArgs struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
	EnableAutomaticUpgrade pulumi.BoolPtrInput
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrInput
	// The virtual machine extension instance view.
	InstanceView VirtualMachineExtensionInstanceViewPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.Input
	// The extensions protected settings that are passed by reference, and consumed from key vault
	ProtectedSettingsFromKeyVault pulumi.Input
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Json formatted public settings for the extension.
	Settings pulumi.Input
	// Indicates whether failures stemming from the extension will be suppressed (Operational failures such as not connecting to the VM will not be suppressed regardless of this value). The default is false.
	SuppressFailures pulumi.BoolPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type pulumi.StringPtrInput
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrInput
	// The name of the virtual machine extension.
	VmExtensionName pulumi.StringPtrInput
	// The name of the virtual machine where the extension should be created or updated.
	VmName pulumi.StringInput
}

func (VirtualMachineExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineExtensionArgs)(nil)).Elem()
}

type VirtualMachineExtensionInput interface {
	pulumi.Input

	ToVirtualMachineExtensionOutput() VirtualMachineExtensionOutput
	ToVirtualMachineExtensionOutputWithContext(ctx context.Context) VirtualMachineExtensionOutput
}

func (*VirtualMachineExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineExtension)(nil)).Elem()
}

func (i *VirtualMachineExtension) ToVirtualMachineExtensionOutput() VirtualMachineExtensionOutput {
	return i.ToVirtualMachineExtensionOutputWithContext(context.Background())
}

func (i *VirtualMachineExtension) ToVirtualMachineExtensionOutputWithContext(ctx context.Context) VirtualMachineExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineExtensionOutput)
}

type VirtualMachineExtensionOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineExtension)(nil)).Elem()
}

func (o VirtualMachineExtensionOutput) ToVirtualMachineExtensionOutput() VirtualMachineExtensionOutput {
	return o
}

func (o VirtualMachineExtensionOutput) ToVirtualMachineExtensionOutputWithContext(ctx context.Context) VirtualMachineExtensionOutput {
	return o
}

// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
func (o VirtualMachineExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.BoolPtrOutput { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
func (o VirtualMachineExtensionOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.BoolPtrOutput { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

// How the extension handler should be forced to update even if the extension configuration has not changed.
func (o VirtualMachineExtensionOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

// The virtual machine extension instance view.
func (o VirtualMachineExtensionOutput) InstanceView() VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) VirtualMachineExtensionInstanceViewResponsePtrOutput {
		return v.InstanceView
	}).(VirtualMachineExtensionInstanceViewResponsePtrOutput)
}

// Resource location
func (o VirtualMachineExtensionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o VirtualMachineExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
func (o VirtualMachineExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.AnyOutput { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

// The extensions protected settings that are passed by reference, and consumed from key vault
func (o VirtualMachineExtensionOutput) ProtectedSettingsFromKeyVault() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.AnyOutput { return v.ProtectedSettingsFromKeyVault }).(pulumi.AnyOutput)
}

// The provisioning state, which only appears in the response.
func (o VirtualMachineExtensionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The name of the extension handler publisher.
func (o VirtualMachineExtensionOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringPtrOutput { return v.Publisher }).(pulumi.StringPtrOutput)
}

// Json formatted public settings for the extension.
func (o VirtualMachineExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.AnyOutput { return v.Settings }).(pulumi.AnyOutput)
}

// Indicates whether failures stemming from the extension will be suppressed (Operational failures such as not connecting to the VM will not be suppressed regardless of this value). The default is false.
func (o VirtualMachineExtensionOutput) SuppressFailures() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.BoolPtrOutput { return v.SuppressFailures }).(pulumi.BoolPtrOutput)
}

// Resource tags
func (o VirtualMachineExtensionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o VirtualMachineExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the version of the script handler.
func (o VirtualMachineExtensionOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringPtrOutput { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineExtensionOutput{})
}
