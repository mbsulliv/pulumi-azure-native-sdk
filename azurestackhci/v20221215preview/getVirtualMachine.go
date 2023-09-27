// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221215preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a virtual machine
func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:azurestackhci/v20221215preview:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualMachineArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the virtual machine
	VirtualMachineName string `pulumi:"virtualMachineName"`
}

// The virtual machine resource definition.
type LookupVirtualMachineResult struct {
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Guest agent status properties.
	GuestAgentProfile *GuestAgentProfileResponse `pulumi:"guestAgentProfile"`
	// HardwareProfile - Specifies the hardware settings for the virtual machine.
	HardwareProfile *VirtualMachinePropertiesResponseHardwareProfile `pulumi:"hardwareProfile"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// NetworkProfile - describes the network configuration the virtual machine
	NetworkProfile *VirtualMachinePropertiesResponseNetworkProfile `pulumi:"networkProfile"`
	// OsProfile - describes the configuration of the operating system and sets login data
	OsProfile *VirtualMachinePropertiesResponseOsProfile `pulumi:"osProfile"`
	// Provisioning state of the virtual machine.
	ProvisioningState string `pulumi:"provisioningState"`
	// SecurityProfile - Specifies the security settings for the virtual machine.
	SecurityProfile *VirtualMachinePropertiesResponseSecurityProfile `pulumi:"securityProfile"`
	// The observed state of virtual machines
	Status VirtualMachineStatusResponse `pulumi:"status"`
	// StorageProfile - contains information about the disks and storage information for the virtual machine
	StorageProfile *VirtualMachinePropertiesResponseStorageProfile `pulumi:"storageProfile"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Unique identifier for the vm resource.
	VmId string `pulumi:"vmId"`
}

// Defaults sets the appropriate defaults for LookupVirtualMachineResult
func (val *LookupVirtualMachineResult) Defaults() *LookupVirtualMachineResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.HardwareProfile = tmp.HardwareProfile.Defaults()

	tmp.SecurityProfile = tmp.SecurityProfile.Defaults()

	return &tmp
}

func LookupVirtualMachineOutput(ctx *pulumi.Context, args LookupVirtualMachineOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineResult, error) {
			args := v.(LookupVirtualMachineArgs)
			r, err := LookupVirtualMachine(ctx, &args, opts...)
			var s LookupVirtualMachineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineResultOutput)
}

type LookupVirtualMachineOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the virtual machine
	VirtualMachineName pulumi.StringInput `pulumi:"virtualMachineName"`
}

func (LookupVirtualMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineArgs)(nil)).Elem()
}

// The virtual machine resource definition.
type LookupVirtualMachineResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineResult)(nil)).Elem()
}

func (o LookupVirtualMachineResultOutput) ToLookupVirtualMachineResultOutput() LookupVirtualMachineResultOutput {
	return o
}

func (o LookupVirtualMachineResultOutput) ToLookupVirtualMachineResultOutputWithContext(ctx context.Context) LookupVirtualMachineResultOutput {
	return o
}

func (o LookupVirtualMachineResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualMachineResult] {
	return pulumix.Output[LookupVirtualMachineResult]{
		OutputState: o.OutputState,
	}
}

// The extendedLocation of the resource.
func (o LookupVirtualMachineResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Guest agent status properties.
func (o LookupVirtualMachineResultOutput) GuestAgentProfile() GuestAgentProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *GuestAgentProfileResponse { return v.GuestAgentProfile }).(GuestAgentProfileResponsePtrOutput)
}

// HardwareProfile - Specifies the hardware settings for the virtual machine.
func (o LookupVirtualMachineResultOutput) HardwareProfile() VirtualMachinePropertiesResponseHardwareProfilePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *VirtualMachinePropertiesResponseHardwareProfile {
		return v.HardwareProfile
	}).(VirtualMachinePropertiesResponseHardwareProfilePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupVirtualMachineResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupVirtualMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupVirtualMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

// NetworkProfile - describes the network configuration the virtual machine
func (o LookupVirtualMachineResultOutput) NetworkProfile() VirtualMachinePropertiesResponseNetworkProfilePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *VirtualMachinePropertiesResponseNetworkProfile {
		return v.NetworkProfile
	}).(VirtualMachinePropertiesResponseNetworkProfilePtrOutput)
}

// OsProfile - describes the configuration of the operating system and sets login data
func (o LookupVirtualMachineResultOutput) OsProfile() VirtualMachinePropertiesResponseOsProfilePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *VirtualMachinePropertiesResponseOsProfile { return v.OsProfile }).(VirtualMachinePropertiesResponseOsProfilePtrOutput)
}

// Provisioning state of the virtual machine.
func (o LookupVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// SecurityProfile - Specifies the security settings for the virtual machine.
func (o LookupVirtualMachineResultOutput) SecurityProfile() VirtualMachinePropertiesResponseSecurityProfilePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *VirtualMachinePropertiesResponseSecurityProfile {
		return v.SecurityProfile
	}).(VirtualMachinePropertiesResponseSecurityProfilePtrOutput)
}

// The observed state of virtual machines
func (o LookupVirtualMachineResultOutput) Status() VirtualMachineStatusResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) VirtualMachineStatusResponse { return v.Status }).(VirtualMachineStatusResponseOutput)
}

// StorageProfile - contains information about the disks and storage information for the virtual machine
func (o LookupVirtualMachineResultOutput) StorageProfile() VirtualMachinePropertiesResponseStorageProfilePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *VirtualMachinePropertiesResponseStorageProfile {
		return v.StorageProfile
	}).(VirtualMachinePropertiesResponseStorageProfilePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVirtualMachineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupVirtualMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVirtualMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

// Unique identifier for the vm resource.
func (o LookupVirtualMachineResultOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.VmId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineResultOutput{})
}
