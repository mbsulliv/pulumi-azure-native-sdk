// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets virtual machines by resource name
func LookupVirtualmachineRetrieve(ctx *pulumi.Context, args *LookupVirtualmachineRetrieveArgs, opts ...pulumi.InvokeOption) (*LookupVirtualmachineRetrieveResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualmachineRetrieveResult
	err := ctx.Invoke("azure-native:azurestackhci/v20210901preview:getVirtualmachineRetrieve", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualmachineRetrieveArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	VirtualmachinesName string `pulumi:"virtualmachinesName"`
}

// The virtual machine resource definition.
type LookupVirtualmachineRetrieveResult struct {
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Guest agent status properties.
	GuestAgentProfile *GuestAgentProfileResponse `pulumi:"guestAgentProfile"`
	// HardwareProfile - Specifies the hardware settings for the virtual machine.
	HardwareProfile *VirtualmachinesPropertiesResponseHardwareProfile `pulumi:"hardwareProfile"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// NetworkProfile - describes the network configuration the virtual machine
	NetworkProfile *VirtualmachinesPropertiesResponseNetworkProfile `pulumi:"networkProfile"`
	// OsProfile - describes the configuration of the operating system and sets login data
	OsProfile         *VirtualmachinesPropertiesResponseOsProfile `pulumi:"osProfile"`
	ProvisioningState string                                      `pulumi:"provisioningState"`
	// name of the object to be used in moc
	ResourceName *string `pulumi:"resourceName"`
	// SecurityProfile - Specifies the security settings for the virtual machine.
	SecurityProfile *VirtualmachinesPropertiesResponseSecurityProfile `pulumi:"securityProfile"`
	// VirtualMachineStatus defines the observed state of virtualmachines
	Status VirtualMachineStatusResponse `pulumi:"status"`
	// StorageProfile - contains information about the disks and storage information for the virtual machine
	StorageProfile *VirtualmachinesPropertiesResponseStorageProfile `pulumi:"storageProfile"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Unique identifier for the vm resource.
	VmId string `pulumi:"vmId"`
}

// Defaults sets the appropriate defaults for LookupVirtualmachineRetrieveResult
func (val *LookupVirtualmachineRetrieveResult) Defaults() *LookupVirtualmachineRetrieveResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.HardwareProfile = tmp.HardwareProfile.Defaults()

	tmp.SecurityProfile = tmp.SecurityProfile.Defaults()

	return &tmp
}

func LookupVirtualmachineRetrieveOutput(ctx *pulumi.Context, args LookupVirtualmachineRetrieveOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualmachineRetrieveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualmachineRetrieveResult, error) {
			args := v.(LookupVirtualmachineRetrieveArgs)
			r, err := LookupVirtualmachineRetrieve(ctx, &args, opts...)
			var s LookupVirtualmachineRetrieveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualmachineRetrieveResultOutput)
}

type LookupVirtualmachineRetrieveOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualmachinesName pulumi.StringInput `pulumi:"virtualmachinesName"`
}

func (LookupVirtualmachineRetrieveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualmachineRetrieveArgs)(nil)).Elem()
}

// The virtual machine resource definition.
type LookupVirtualmachineRetrieveResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualmachineRetrieveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualmachineRetrieveResult)(nil)).Elem()
}

func (o LookupVirtualmachineRetrieveResultOutput) ToLookupVirtualmachineRetrieveResultOutput() LookupVirtualmachineRetrieveResultOutput {
	return o
}

func (o LookupVirtualmachineRetrieveResultOutput) ToLookupVirtualmachineRetrieveResultOutputWithContext(ctx context.Context) LookupVirtualmachineRetrieveResultOutput {
	return o
}

func (o LookupVirtualmachineRetrieveResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualmachineRetrieveResult] {
	return pulumix.Output[LookupVirtualmachineRetrieveResult]{
		OutputState: o.OutputState,
	}
}

// The extendedLocation of the resource.
func (o LookupVirtualmachineRetrieveResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Guest agent status properties.
func (o LookupVirtualmachineRetrieveResultOutput) GuestAgentProfile() GuestAgentProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) *GuestAgentProfileResponse { return v.GuestAgentProfile }).(GuestAgentProfileResponsePtrOutput)
}

// HardwareProfile - Specifies the hardware settings for the virtual machine.
func (o LookupVirtualmachineRetrieveResultOutput) HardwareProfile() VirtualmachinesPropertiesResponseHardwareProfilePtrOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) *VirtualmachinesPropertiesResponseHardwareProfile {
		return v.HardwareProfile
	}).(VirtualmachinesPropertiesResponseHardwareProfilePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupVirtualmachineRetrieveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupVirtualmachineRetrieveResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupVirtualmachineRetrieveResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupVirtualmachineRetrieveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) string { return v.Name }).(pulumi.StringOutput)
}

// NetworkProfile - describes the network configuration the virtual machine
func (o LookupVirtualmachineRetrieveResultOutput) NetworkProfile() VirtualmachinesPropertiesResponseNetworkProfilePtrOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) *VirtualmachinesPropertiesResponseNetworkProfile {
		return v.NetworkProfile
	}).(VirtualmachinesPropertiesResponseNetworkProfilePtrOutput)
}

// OsProfile - describes the configuration of the operating system and sets login data
func (o LookupVirtualmachineRetrieveResultOutput) OsProfile() VirtualmachinesPropertiesResponseOsProfilePtrOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) *VirtualmachinesPropertiesResponseOsProfile {
		return v.OsProfile
	}).(VirtualmachinesPropertiesResponseOsProfilePtrOutput)
}

func (o LookupVirtualmachineRetrieveResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// name of the object to be used in moc
func (o LookupVirtualmachineRetrieveResultOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

// SecurityProfile - Specifies the security settings for the virtual machine.
func (o LookupVirtualmachineRetrieveResultOutput) SecurityProfile() VirtualmachinesPropertiesResponseSecurityProfilePtrOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) *VirtualmachinesPropertiesResponseSecurityProfile {
		return v.SecurityProfile
	}).(VirtualmachinesPropertiesResponseSecurityProfilePtrOutput)
}

// VirtualMachineStatus defines the observed state of virtualmachines
func (o LookupVirtualmachineRetrieveResultOutput) Status() VirtualMachineStatusResponseOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) VirtualMachineStatusResponse { return v.Status }).(VirtualMachineStatusResponseOutput)
}

// StorageProfile - contains information about the disks and storage information for the virtual machine
func (o LookupVirtualmachineRetrieveResultOutput) StorageProfile() VirtualmachinesPropertiesResponseStorageProfilePtrOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) *VirtualmachinesPropertiesResponseStorageProfile {
		return v.StorageProfile
	}).(VirtualmachinesPropertiesResponseStorageProfilePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupVirtualmachineRetrieveResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupVirtualmachineRetrieveResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVirtualmachineRetrieveResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) string { return v.Type }).(pulumi.StringOutput)
}

// Unique identifier for the vm resource.
func (o LookupVirtualmachineRetrieveResultOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualmachineRetrieveResult) string { return v.VmId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualmachineRetrieveResultOutput{})
}
