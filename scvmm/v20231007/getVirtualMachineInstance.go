// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231007

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about a virtual machine instance.
func LookupVirtualMachineInstance(ctx *pulumi.Context, args *LookupVirtualMachineInstanceArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineInstanceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualMachineInstanceResult
	err := ctx.Invoke("azure-native:scvmm/v20231007:getVirtualMachineInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineInstanceArgs struct {
	// The fully qualified Azure Resource manager identifier of the Hybrid Compute machine resource to be extended.
	ResourceUri string `pulumi:"resourceUri"`
}

// Define the virtualMachineInstance.
type LookupVirtualMachineInstanceResult struct {
	// Availability Sets in vm.
	AvailabilitySets []VirtualMachineInstancePropertiesResponseAvailabilitySets `pulumi:"availabilitySets"`
	// Gets or sets the extended location.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Hardware properties.
	HardwareProfile *HardwareProfileResponse `pulumi:"hardwareProfile"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Gets the infrastructure profile.
	InfrastructureProfile *InfrastructureProfileResponse `pulumi:"infrastructureProfile"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Network properties.
	NetworkProfile *NetworkProfileResponse `pulumi:"networkProfile"`
	// OS properties.
	OsProfile *OsProfileForVMInstanceResponse `pulumi:"osProfile"`
	// Gets the power state of the virtual machine.
	PowerState string `pulumi:"powerState"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Storage properties.
	StorageProfile *StorageProfileResponse `pulumi:"storageProfile"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupVirtualMachineInstanceOutput(ctx *pulumi.Context, args LookupVirtualMachineInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineInstanceResult, error) {
			args := v.(LookupVirtualMachineInstanceArgs)
			r, err := LookupVirtualMachineInstance(ctx, &args, opts...)
			var s LookupVirtualMachineInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineInstanceResultOutput)
}

type LookupVirtualMachineInstanceOutputArgs struct {
	// The fully qualified Azure Resource manager identifier of the Hybrid Compute machine resource to be extended.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupVirtualMachineInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineInstanceArgs)(nil)).Elem()
}

// Define the virtualMachineInstance.
type LookupVirtualMachineInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineInstanceResult)(nil)).Elem()
}

func (o LookupVirtualMachineInstanceResultOutput) ToLookupVirtualMachineInstanceResultOutput() LookupVirtualMachineInstanceResultOutput {
	return o
}

func (o LookupVirtualMachineInstanceResultOutput) ToLookupVirtualMachineInstanceResultOutputWithContext(ctx context.Context) LookupVirtualMachineInstanceResultOutput {
	return o
}

// Availability Sets in vm.
func (o LookupVirtualMachineInstanceResultOutput) AvailabilitySets() VirtualMachineInstancePropertiesResponseAvailabilitySetsArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineInstanceResult) []VirtualMachineInstancePropertiesResponseAvailabilitySets {
		return v.AvailabilitySets
	}).(VirtualMachineInstancePropertiesResponseAvailabilitySetsArrayOutput)
}

// Gets or sets the extended location.
func (o LookupVirtualMachineInstanceResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineInstanceResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Hardware properties.
func (o LookupVirtualMachineInstanceResultOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineInstanceResult) *HardwareProfileResponse { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupVirtualMachineInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets the infrastructure profile.
func (o LookupVirtualMachineInstanceResultOutput) InfrastructureProfile() InfrastructureProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineInstanceResult) *InfrastructureProfileResponse {
		return v.InfrastructureProfile
	}).(InfrastructureProfileResponsePtrOutput)
}

// The name of the resource
func (o LookupVirtualMachineInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network properties.
func (o LookupVirtualMachineInstanceResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineInstanceResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

// OS properties.
func (o LookupVirtualMachineInstanceResultOutput) OsProfile() OsProfileForVMInstanceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineInstanceResult) *OsProfileForVMInstanceResponse { return v.OsProfile }).(OsProfileForVMInstanceResponsePtrOutput)
}

// Gets the power state of the virtual machine.
func (o LookupVirtualMachineInstanceResultOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineInstanceResult) string { return v.PowerState }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupVirtualMachineInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Storage properties.
func (o LookupVirtualMachineInstanceResultOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineInstanceResult) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVirtualMachineInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVirtualMachineInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineInstanceResultOutput{})
}
