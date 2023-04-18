// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about the model view or the instance view of a virtual machine.
func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:compute/v20230301:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineArgs struct {
	// The expand expression to apply on the operation. 'InstanceView' retrieves a snapshot of the runtime properties of the virtual machine that is managed by the platform and can change outside of control plane operations. 'UserData' retrieves the UserData property as part of the VM model view that was provided by the user during the VM Create/Update operation.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual machine.
	VmName string `pulumi:"vmName"`
}

// Describes a Virtual Machine.
type LookupVirtualMachineResult struct {
	// Specifies additional capabilities enabled or disabled on the virtual machine.
	AdditionalCapabilities *AdditionalCapabilitiesResponse `pulumi:"additionalCapabilities"`
	// Specifies the gallery applications that should be made available to the VM/VMSS.
	ApplicationProfile *ApplicationProfileResponse `pulumi:"applicationProfile"`
	// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Availability sets overview](https://docs.microsoft.com/azure/virtual-machines/availability-set-overview). For more information on Azure planned maintenance, see [Maintenance and updates for Virtual Machines in Azure](https://docs.microsoft.com/azure/virtual-machines/maintenance-and-updates). Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
	AvailabilitySet *SubResourceResponse `pulumi:"availabilitySet"`
	// Specifies the billing related details of a Azure Spot virtual machine. Minimum api-version: 2019-03-01.
	BillingProfile *BillingProfileResponse `pulumi:"billingProfile"`
	// Specifies information about the capacity reservation that is used to allocate virtual machine. Minimum api-version: 2021-04-01.
	CapacityReservation *CapacityReservationProfileResponse `pulumi:"capacityReservation"`
	// Specifies the boot diagnostic settings state. Minimum api-version: 2015-06-15.
	DiagnosticsProfile *DiagnosticsProfileResponse `pulumi:"diagnosticsProfile"`
	// Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// The extended location of the Virtual Machine.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). Minimum api-version: 2020-06-01.
	ExtensionsTimeBudget *string `pulumi:"extensionsTimeBudget"`
	// Specifies the hardware settings for the virtual machine.
	HardwareProfile *HardwareProfileResponse `pulumi:"hardwareProfile"`
	// Specifies information about the dedicated host that the virtual machine resides in. Minimum api-version: 2018-10-01.
	Host *SubResourceResponse `pulumi:"host"`
	// Specifies information about the dedicated host group that the virtual machine resides in. **Note:** User cannot specify both host and hostGroup properties. Minimum api-version: 2020-06-01.
	HostGroup *SubResourceResponse `pulumi:"hostGroup"`
	// Resource Id
	Id string `pulumi:"id"`
	// The identity of the virtual machine, if configured.
	Identity *VirtualMachineIdentityResponse `pulumi:"identity"`
	// The virtual machine instance view.
	InstanceView VirtualMachineInstanceViewResponse `pulumi:"instanceView"`
	// Specifies that the image or disk that is being used was licensed on-premises. <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15
	LicenseType *string `pulumi:"licenseType"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Specifies the network interfaces of the virtual machine.
	NetworkProfile *NetworkProfileResponse `pulumi:"networkProfile"`
	// Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
	OsProfile *OSProfileResponse `pulumi:"osProfile"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan *PlanResponse `pulumi:"plan"`
	// Specifies the scale set logical fault domain into which the Virtual Machine will be created. By default, the Virtual Machine will by automatically assigned to a fault domain that best maintains balance across available fault domains. This is applicable only if the 'virtualMachineScaleSet' property of this Virtual Machine is set. The Virtual Machine Scale Set that is referenced, must have 'platformFaultDomainCount' greater than 1. This property cannot be updated once the Virtual Machine is created. Fault domain assignment can be viewed in the Virtual Machine Instance View. Minimum api‐version: 2020‐12‐01.
	PlatformFaultDomain *int `pulumi:"platformFaultDomain"`
	// Specifies the priority for the virtual machine. Minimum api-version: 2019-03-01
	Priority *string `pulumi:"priority"`
	// The provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// Specifies information about the proximity placement group that the virtual machine should be assigned to. Minimum api-version: 2018-04-01.
	ProximityPlacementGroup *SubResourceResponse `pulumi:"proximityPlacementGroup"`
	// The virtual machine child extension resources.
	Resources []VirtualMachineExtensionResponse `pulumi:"resources"`
	// Specifies Scheduled Event related configurations.
	ScheduledEventsProfile *ScheduledEventsProfileResponse `pulumi:"scheduledEventsProfile"`
	// Specifies the Security related profile settings for the virtual machine.
	SecurityProfile *SecurityProfileResponse `pulumi:"securityProfile"`
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile *StorageProfileResponse `pulumi:"storageProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Specifies the time at which the Virtual Machine resource was created. Minimum api-version: 2021-11-01.
	TimeCreated string `pulumi:"timeCreated"`
	// Resource type
	Type string `pulumi:"type"`
	// UserData for the VM, which must be base-64 encoded. Customer should not pass any secrets in here. Minimum api-version: 2021-03-01.
	UserData *string `pulumi:"userData"`
	// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. This property cannot exist along with a non-null properties.availabilitySet reference. Minimum api‐version: 2019‐03‐01.
	VirtualMachineScaleSet *SubResourceResponse `pulumi:"virtualMachineScaleSet"`
	// Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
	VmId string `pulumi:"vmId"`
	// The virtual machine zones.
	Zones []string `pulumi:"zones"`
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
	// The expand expression to apply on the operation. 'InstanceView' retrieves a snapshot of the runtime properties of the virtual machine that is managed by the platform and can change outside of control plane operations. 'UserData' retrieves the UserData property as part of the VM model view that was provided by the user during the VM Create/Update operation.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the virtual machine.
	VmName pulumi.StringInput `pulumi:"vmName"`
}

func (LookupVirtualMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineArgs)(nil)).Elem()
}

// Describes a Virtual Machine.
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

// Specifies additional capabilities enabled or disabled on the virtual machine.
func (o LookupVirtualMachineResultOutput) AdditionalCapabilities() AdditionalCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *AdditionalCapabilitiesResponse { return v.AdditionalCapabilities }).(AdditionalCapabilitiesResponsePtrOutput)
}

// Specifies the gallery applications that should be made available to the VM/VMSS.
func (o LookupVirtualMachineResultOutput) ApplicationProfile() ApplicationProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *ApplicationProfileResponse { return v.ApplicationProfile }).(ApplicationProfileResponsePtrOutput)
}

// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Availability sets overview](https://docs.microsoft.com/azure/virtual-machines/availability-set-overview). For more information on Azure planned maintenance, see [Maintenance and updates for Virtual Machines in Azure](https://docs.microsoft.com/azure/virtual-machines/maintenance-and-updates). Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
func (o LookupVirtualMachineResultOutput) AvailabilitySet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SubResourceResponse { return v.AvailabilitySet }).(SubResourceResponsePtrOutput)
}

// Specifies the billing related details of a Azure Spot virtual machine. Minimum api-version: 2019-03-01.
func (o LookupVirtualMachineResultOutput) BillingProfile() BillingProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *BillingProfileResponse { return v.BillingProfile }).(BillingProfileResponsePtrOutput)
}

// Specifies information about the capacity reservation that is used to allocate virtual machine. Minimum api-version: 2021-04-01.
func (o LookupVirtualMachineResultOutput) CapacityReservation() CapacityReservationProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *CapacityReservationProfileResponse { return v.CapacityReservation }).(CapacityReservationProfileResponsePtrOutput)
}

// Specifies the boot diagnostic settings state. Minimum api-version: 2015-06-15.
func (o LookupVirtualMachineResultOutput) DiagnosticsProfile() DiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *DiagnosticsProfileResponse { return v.DiagnosticsProfile }).(DiagnosticsProfileResponsePtrOutput)
}

// Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
func (o LookupVirtualMachineResultOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.EvictionPolicy }).(pulumi.StringPtrOutput)
}

// The extended location of the Virtual Machine.
func (o LookupVirtualMachineResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). Minimum api-version: 2020-06-01.
func (o LookupVirtualMachineResultOutput) ExtensionsTimeBudget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.ExtensionsTimeBudget }).(pulumi.StringPtrOutput)
}

// Specifies the hardware settings for the virtual machine.
func (o LookupVirtualMachineResultOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *HardwareProfileResponse { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

// Specifies information about the dedicated host that the virtual machine resides in. Minimum api-version: 2018-10-01.
func (o LookupVirtualMachineResultOutput) Host() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SubResourceResponse { return v.Host }).(SubResourceResponsePtrOutput)
}

// Specifies information about the dedicated host group that the virtual machine resides in. **Note:** User cannot specify both host and hostGroup properties. Minimum api-version: 2020-06-01.
func (o LookupVirtualMachineResultOutput) HostGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SubResourceResponse { return v.HostGroup }).(SubResourceResponsePtrOutput)
}

// Resource Id
func (o LookupVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the virtual machine, if configured.
func (o LookupVirtualMachineResultOutput) Identity() VirtualMachineIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *VirtualMachineIdentityResponse { return v.Identity }).(VirtualMachineIdentityResponsePtrOutput)
}

// The virtual machine instance view.
func (o LookupVirtualMachineResultOutput) InstanceView() VirtualMachineInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) VirtualMachineInstanceViewResponse { return v.InstanceView }).(VirtualMachineInstanceViewResponseOutput)
}

// Specifies that the image or disk that is being used was licensed on-premises. <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15
func (o LookupVirtualMachineResultOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// Resource location
func (o LookupVirtualMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupVirtualMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the network interfaces of the virtual machine.
func (o LookupVirtualMachineResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

// Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
func (o LookupVirtualMachineResultOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *OSProfileResponse { return v.OsProfile }).(OSProfileResponsePtrOutput)
}

// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
func (o LookupVirtualMachineResultOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

// Specifies the scale set logical fault domain into which the Virtual Machine will be created. By default, the Virtual Machine will by automatically assigned to a fault domain that best maintains balance across available fault domains. This is applicable only if the 'virtualMachineScaleSet' property of this Virtual Machine is set. The Virtual Machine Scale Set that is referenced, must have 'platformFaultDomainCount' greater than 1. This property cannot be updated once the Virtual Machine is created. Fault domain assignment can be viewed in the Virtual Machine Instance View. Minimum api‐version: 2020‐12‐01.
func (o LookupVirtualMachineResultOutput) PlatformFaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.PlatformFaultDomain }).(pulumi.IntPtrOutput)
}

// Specifies the priority for the virtual machine. Minimum api-version: 2019-03-01
func (o LookupVirtualMachineResultOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

// The provisioning state, which only appears in the response.
func (o LookupVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specifies information about the proximity placement group that the virtual machine should be assigned to. Minimum api-version: 2018-04-01.
func (o LookupVirtualMachineResultOutput) ProximityPlacementGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SubResourceResponse { return v.ProximityPlacementGroup }).(SubResourceResponsePtrOutput)
}

// The virtual machine child extension resources.
func (o LookupVirtualMachineResultOutput) Resources() VirtualMachineExtensionResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []VirtualMachineExtensionResponse { return v.Resources }).(VirtualMachineExtensionResponseArrayOutput)
}

// Specifies Scheduled Event related configurations.
func (o LookupVirtualMachineResultOutput) ScheduledEventsProfile() ScheduledEventsProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *ScheduledEventsProfileResponse { return v.ScheduledEventsProfile }).(ScheduledEventsProfileResponsePtrOutput)
}

// Specifies the Security related profile settings for the virtual machine.
func (o LookupVirtualMachineResultOutput) SecurityProfile() SecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SecurityProfileResponse { return v.SecurityProfile }).(SecurityProfileResponsePtrOutput)
}

// Specifies the storage settings for the virtual machine disks.
func (o LookupVirtualMachineResultOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

// Resource tags
func (o LookupVirtualMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the time at which the Virtual Machine resource was created. Minimum api-version: 2021-11-01.
func (o LookupVirtualMachineResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

// Resource type
func (o LookupVirtualMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

// UserData for the VM, which must be base-64 encoded. Customer should not pass any secrets in here. Minimum api-version: 2021-03-01.
func (o LookupVirtualMachineResultOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.UserData }).(pulumi.StringPtrOutput)
}

// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. This property cannot exist along with a non-null properties.availabilitySet reference. Minimum api‐version: 2019‐03‐01.
func (o LookupVirtualMachineResultOutput) VirtualMachineScaleSet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SubResourceResponse { return v.VirtualMachineScaleSet }).(SubResourceResponsePtrOutput)
}

// Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
func (o LookupVirtualMachineResultOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.VmId }).(pulumi.StringOutput)
}

// The virtual machine zones.
func (o LookupVirtualMachineResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineResultOutput{})
}
