// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurestackhci

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The operation to get the extension.
// Azure REST API version: 2022-12-15-preview.
func LookupMachineExtension(ctx *pulumi.Context, args *LookupMachineExtensionArgs, opts ...pulumi.InvokeOption) (*LookupMachineExtensionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMachineExtensionResult
	err := ctx.Invoke("azure-native:azurestackhci:getMachineExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineExtensionArgs struct {
	// The name of the machine extension.
	ExtensionName string `pulumi:"extensionName"`
	// The name of the machine containing the extension.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes a Machine Extension.
type LookupMachineExtensionResult struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The Id.
	Id string `pulumi:"id"`
	// The machine extension instance view.
	InstanceView *MachineExtensionPropertiesResponseInstanceView `pulumi:"instanceView"`
	// The location.
	Location *string `pulumi:"location"`
	// The name.
	Name string `pulumi:"name"`
	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
	ProtectedSettings interface{} `pulumi:"protectedSettings"`
	// The provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// The name of the extension handler publisher.
	Publisher *string `pulumi:"publisher"`
	// Json formatted public settings for the extension.
	Settings interface{} `pulumi:"settings"`
	// The system data.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
}

func LookupMachineExtensionOutput(ctx *pulumi.Context, args LookupMachineExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupMachineExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMachineExtensionResult, error) {
			args := v.(LookupMachineExtensionArgs)
			r, err := LookupMachineExtension(ctx, &args, opts...)
			var s LookupMachineExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMachineExtensionResultOutput)
}

type LookupMachineExtensionOutputArgs struct {
	// The name of the machine extension.
	ExtensionName pulumi.StringInput `pulumi:"extensionName"`
	// The name of the machine containing the extension.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMachineExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineExtensionArgs)(nil)).Elem()
}

// Describes a Machine Extension.
type LookupMachineExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupMachineExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineExtensionResult)(nil)).Elem()
}

func (o LookupMachineExtensionResultOutput) ToLookupMachineExtensionResultOutput() LookupMachineExtensionResultOutput {
	return o
}

func (o LookupMachineExtensionResultOutput) ToLookupMachineExtensionResultOutputWithContext(ctx context.Context) LookupMachineExtensionResultOutput {
	return o
}

// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
func (o LookupMachineExtensionResultOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

// How the extension handler should be forced to update even if the extension configuration has not changed.
func (o LookupMachineExtensionResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

// The Id.
func (o LookupMachineExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The machine extension instance view.
func (o LookupMachineExtensionResultOutput) InstanceView() MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) *MachineExtensionPropertiesResponseInstanceView {
		return v.InstanceView
	}).(MachineExtensionPropertiesResponseInstanceViewPtrOutput)
}

// The location.
func (o LookupMachineExtensionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name.
func (o LookupMachineExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
func (o LookupMachineExtensionResultOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

// The provisioning state, which only appears in the response.
func (o LookupMachineExtensionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The name of the extension handler publisher.
func (o LookupMachineExtensionResultOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

// Json formatted public settings for the extension.
func (o LookupMachineExtensionResultOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

// The system data.
func (o LookupMachineExtensionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The Resource tags.
func (o LookupMachineExtensionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupMachineExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

// Specifies the version of the script handler.
func (o LookupMachineExtensionResultOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineExtensionResultOutput{})
}
