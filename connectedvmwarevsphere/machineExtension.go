// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Describes a Machine Extension.
// Azure REST API version: 2022-07-15-preview. Prior API version in Azure Native 1.x: 2020-10-01-preview.
//
// Other available API versions: 2022-01-10-preview, 2023-03-01-preview.
type MachineExtension struct {
	pulumi.CustomResourceState

	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrOutput `pulumi:"autoUpgradeMinorVersion"`
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version available.
	EnableAutomaticUpgrade pulumi.BoolPtrOutput `pulumi:"enableAutomaticUpgrade"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrOutput `pulumi:"forceUpdateTag"`
	// The machine extension instance view.
	InstanceView MachineExtensionPropertiesResponseInstanceViewPtrOutput `pulumi:"instanceView"`
	// Gets or sets the location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets or sets the name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.AnyOutput `pulumi:"protectedSettings"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrOutput `pulumi:"publisher"`
	// Json formatted public settings for the extension.
	Settings pulumi.AnyOutput `pulumi:"settings"`
	// The system data.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Gets or sets the Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets or sets the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrOutput `pulumi:"typeHandlerVersion"`
}

// NewMachineExtension registers a new resource with the given unique name, arguments, and options.
func NewMachineExtension(ctx *pulumi.Context,
	name string, args *MachineExtensionArgs, opts ...pulumi.ResourceOption) (*MachineExtension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualMachineName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20230301preview:MachineExtension"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MachineExtension
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere:MachineExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineExtension gets an existing MachineExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineExtensionState, opts ...pulumi.ResourceOption) (*MachineExtension, error) {
	var resource MachineExtension
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere:MachineExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineExtension resources.
type machineExtensionState struct {
}

type MachineExtensionState struct {
}

func (MachineExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineExtensionState)(nil)).Elem()
}

type machineExtensionArgs struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version available.
	EnableAutomaticUpgrade *bool `pulumi:"enableAutomaticUpgrade"`
	// The name of the machine extension.
	ExtensionName *string `pulumi:"extensionName"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings interface{} `pulumi:"protectedSettings"`
	// The name of the extension handler publisher.
	Publisher *string `pulumi:"publisher"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Json formatted public settings for the extension.
	Settings interface{} `pulumi:"settings"`
	// Gets or sets the Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type *string `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
	// The name of the machine where the extension should be created or updated.
	VirtualMachineName string `pulumi:"virtualMachineName"`
}

// The set of arguments for constructing a MachineExtension resource.
type MachineExtensionArgs struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version available.
	EnableAutomaticUpgrade pulumi.BoolPtrInput
	// The name of the machine extension.
	ExtensionName pulumi.StringPtrInput
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag pulumi.StringPtrInput
	// Gets or sets the location.
	Location pulumi.StringPtrInput
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings pulumi.Input
	// The name of the extension handler publisher.
	Publisher pulumi.StringPtrInput
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput
	// Json formatted public settings for the extension.
	Settings pulumi.Input
	// Gets or sets the Resource tags.
	Tags pulumi.StringMapInput
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type pulumi.StringPtrInput
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrInput
	// The name of the machine where the extension should be created or updated.
	VirtualMachineName pulumi.StringInput
}

func (MachineExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineExtensionArgs)(nil)).Elem()
}

type MachineExtensionInput interface {
	pulumi.Input

	ToMachineExtensionOutput() MachineExtensionOutput
	ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput
}

func (*MachineExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtension)(nil)).Elem()
}

func (i *MachineExtension) ToMachineExtensionOutput() MachineExtensionOutput {
	return i.ToMachineExtensionOutputWithContext(context.Background())
}

func (i *MachineExtension) ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionOutput)
}

func (i *MachineExtension) ToOutput(ctx context.Context) pulumix.Output[*MachineExtension] {
	return pulumix.Output[*MachineExtension]{
		OutputState: i.ToMachineExtensionOutputWithContext(ctx).OutputState,
	}
}

type MachineExtensionOutput struct{ *pulumi.OutputState }

func (MachineExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtension)(nil)).Elem()
}

func (o MachineExtensionOutput) ToMachineExtensionOutput() MachineExtensionOutput {
	return o
}

func (o MachineExtensionOutput) ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput {
	return o
}

func (o MachineExtensionOutput) ToOutput(ctx context.Context) pulumix.Output[*MachineExtension] {
	return pulumix.Output[*MachineExtension]{
		OutputState: o.OutputState,
	}
}

// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
func (o MachineExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.BoolPtrOutput { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version available.
func (o MachineExtensionOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.BoolPtrOutput { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

// How the extension handler should be forced to update even if the extension configuration has not changed.
func (o MachineExtensionOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

// The machine extension instance view.
func (o MachineExtensionOutput) InstanceView() MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o.ApplyT(func(v *MachineExtension) MachineExtensionPropertiesResponseInstanceViewPtrOutput {
		return v.InstanceView
	}).(MachineExtensionPropertiesResponseInstanceViewPtrOutput)
}

// Gets or sets the location.
func (o MachineExtensionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets or sets the name.
func (o MachineExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
func (o MachineExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.AnyOutput { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

// The provisioning state, which only appears in the response.
func (o MachineExtensionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The name of the extension handler publisher.
func (o MachineExtensionOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringPtrOutput { return v.Publisher }).(pulumi.StringPtrOutput)
}

// Json formatted public settings for the extension.
func (o MachineExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.AnyOutput { return v.Settings }).(pulumi.AnyOutput)
}

// The system data.
func (o MachineExtensionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MachineExtension) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets the Resource tags.
func (o MachineExtensionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o MachineExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the version of the script handler.
func (o MachineExtensionOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringPtrOutput { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineExtensionOutput{})
}
