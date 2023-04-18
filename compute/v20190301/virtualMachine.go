// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a Virtual Machine.
//
// Deprecated: Version 2019-03-01 will be removed in v2 of the provider.
type VirtualMachine struct {
	pulumi.CustomResourceState

	// Specifies additional capabilities enabled or disabled on the virtual machine.
	AdditionalCapabilities AdditionalCapabilitiesResponsePtrOutput `pulumi:"additionalCapabilities"`
	// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
	AvailabilitySet SubResourceResponsePtrOutput `pulumi:"availabilitySet"`
	// Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
	BillingProfile BillingProfileResponsePtrOutput `pulumi:"billingProfile"`
	// Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
	DiagnosticsProfile DiagnosticsProfileResponsePtrOutput `pulumi:"diagnosticsProfile"`
	// Specifies the eviction policy for the Azure Spot virtual machine. Only supported value is 'Deallocate'. <br><br>Minimum api-version: 2019-03-01
	EvictionPolicy pulumi.StringPtrOutput `pulumi:"evictionPolicy"`
	// Specifies the hardware settings for the virtual machine.
	HardwareProfile HardwareProfileResponsePtrOutput `pulumi:"hardwareProfile"`
	// Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
	Host SubResourceResponsePtrOutput `pulumi:"host"`
	// The identity of the virtual machine, if configured.
	Identity VirtualMachineIdentityResponsePtrOutput `pulumi:"identity"`
	// The virtual machine instance view.
	InstanceView VirtualMachineInstanceViewResponseOutput `pulumi:"instanceView"`
	// Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. <br><br> Possible values are: <br><br> Windows_Client <br><br> Windows_Server <br><br> If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Minimum api-version: 2015-06-15
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the network interfaces of the virtual machine.
	NetworkProfile NetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// Specifies the operating system settings for the virtual machine.
	OsProfile OSProfileResponsePtrOutput `pulumi:"osProfile"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanResponsePtrOutput `pulumi:"plan"`
	// Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
	Priority pulumi.StringPtrOutput `pulumi:"priority"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup SubResourceResponsePtrOutput `pulumi:"proximityPlacementGroup"`
	// The virtual machine child extension resources.
	Resources VirtualMachineExtensionResponseArrayOutput `pulumi:"resources"`
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile StorageProfileResponsePtrOutput `pulumi:"storageProfile"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
	VirtualMachineScaleSet SubResourceResponsePtrOutput `pulumi:"virtualMachineScaleSet"`
	// Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
	VmId pulumi.StringOutput `pulumi:"vmId"`
	// The virtual machine zones.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20150615:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160330:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230301:VirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachine
	err := ctx.RegisterResource("azure-native:compute/v20190301:VirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachine gets an existing VirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	var resource VirtualMachine
	err := ctx.ReadResource("azure-native:compute/v20190301:VirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachine resources.
type virtualMachineState struct {
}

type VirtualMachineState struct {
}

func (VirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineState)(nil)).Elem()
}

type virtualMachineArgs struct {
	// Specifies additional capabilities enabled or disabled on the virtual machine.
	AdditionalCapabilities *AdditionalCapabilities `pulumi:"additionalCapabilities"`
	// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
	AvailabilitySet *SubResource `pulumi:"availabilitySet"`
	// Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
	BillingProfile *BillingProfile `pulumi:"billingProfile"`
	// Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
	DiagnosticsProfile *DiagnosticsProfile `pulumi:"diagnosticsProfile"`
	// Specifies the eviction policy for the Azure Spot virtual machine. Only supported value is 'Deallocate'. <br><br>Minimum api-version: 2019-03-01
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// Specifies the hardware settings for the virtual machine.
	HardwareProfile *HardwareProfile `pulumi:"hardwareProfile"`
	// Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
	Host *SubResource `pulumi:"host"`
	// The identity of the virtual machine, if configured.
	Identity *VirtualMachineIdentity `pulumi:"identity"`
	// Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. <br><br> Possible values are: <br><br> Windows_Client <br><br> Windows_Server <br><br> If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Minimum api-version: 2015-06-15
	LicenseType *string `pulumi:"licenseType"`
	// Resource location
	Location *string `pulumi:"location"`
	// Specifies the network interfaces of the virtual machine.
	NetworkProfile *NetworkProfile `pulumi:"networkProfile"`
	// Specifies the operating system settings for the virtual machine.
	OsProfile *OSProfile `pulumi:"osProfile"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan *Plan `pulumi:"plan"`
	// Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
	Priority *string `pulumi:"priority"`
	// Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup *SubResource `pulumi:"proximityPlacementGroup"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile *StorageProfile `pulumi:"storageProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
	VirtualMachineScaleSet *SubResource `pulumi:"virtualMachineScaleSet"`
	// The name of the virtual machine.
	VmName *string `pulumi:"vmName"`
	// The virtual machine zones.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a VirtualMachine resource.
type VirtualMachineArgs struct {
	// Specifies additional capabilities enabled or disabled on the virtual machine.
	AdditionalCapabilities AdditionalCapabilitiesPtrInput
	// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
	AvailabilitySet SubResourcePtrInput
	// Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
	BillingProfile BillingProfilePtrInput
	// Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
	DiagnosticsProfile DiagnosticsProfilePtrInput
	// Specifies the eviction policy for the Azure Spot virtual machine. Only supported value is 'Deallocate'. <br><br>Minimum api-version: 2019-03-01
	EvictionPolicy pulumi.StringPtrInput
	// Specifies the hardware settings for the virtual machine.
	HardwareProfile HardwareProfilePtrInput
	// Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
	Host SubResourcePtrInput
	// The identity of the virtual machine, if configured.
	Identity VirtualMachineIdentityPtrInput
	// Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. <br><br> Possible values are: <br><br> Windows_Client <br><br> Windows_Server <br><br> If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Minimum api-version: 2015-06-15
	LicenseType pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Specifies the network interfaces of the virtual machine.
	NetworkProfile NetworkProfilePtrInput
	// Specifies the operating system settings for the virtual machine.
	OsProfile OSProfilePtrInput
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanPtrInput
	// Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
	Priority pulumi.StringPtrInput
	// Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup SubResourcePtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile StorageProfilePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
	VirtualMachineScaleSet SubResourcePtrInput
	// The name of the virtual machine.
	VmName pulumi.StringPtrInput
	// The virtual machine zones.
	Zones pulumi.StringArrayInput
}

func (VirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineArgs)(nil)).Elem()
}

type VirtualMachineInput interface {
	pulumi.Input

	ToVirtualMachineOutput() VirtualMachineOutput
	ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput
}

func (*VirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachine)(nil)).Elem()
}

func (i *VirtualMachine) ToVirtualMachineOutput() VirtualMachineOutput {
	return i.ToVirtualMachineOutputWithContext(context.Background())
}

func (i *VirtualMachine) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineOutput)
}

type VirtualMachineOutput struct{ *pulumi.OutputState }

func (VirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachine)(nil)).Elem()
}

func (o VirtualMachineOutput) ToVirtualMachineOutput() VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return o
}

// Specifies additional capabilities enabled or disabled on the virtual machine.
func (o VirtualMachineOutput) AdditionalCapabilities() AdditionalCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) AdditionalCapabilitiesResponsePtrOutput { return v.AdditionalCapabilities }).(AdditionalCapabilitiesResponsePtrOutput)
}

// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
func (o VirtualMachineOutput) AvailabilitySet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) SubResourceResponsePtrOutput { return v.AvailabilitySet }).(SubResourceResponsePtrOutput)
}

// Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
func (o VirtualMachineOutput) BillingProfile() BillingProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) BillingProfileResponsePtrOutput { return v.BillingProfile }).(BillingProfileResponsePtrOutput)
}

// Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
func (o VirtualMachineOutput) DiagnosticsProfile() DiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) DiagnosticsProfileResponsePtrOutput { return v.DiagnosticsProfile }).(DiagnosticsProfileResponsePtrOutput)
}

// Specifies the eviction policy for the Azure Spot virtual machine. Only supported value is 'Deallocate'. <br><br>Minimum api-version: 2019-03-01
func (o VirtualMachineOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.EvictionPolicy }).(pulumi.StringPtrOutput)
}

// Specifies the hardware settings for the virtual machine.
func (o VirtualMachineOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) HardwareProfileResponsePtrOutput { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

// Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
func (o VirtualMachineOutput) Host() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) SubResourceResponsePtrOutput { return v.Host }).(SubResourceResponsePtrOutput)
}

// The identity of the virtual machine, if configured.
func (o VirtualMachineOutput) Identity() VirtualMachineIdentityResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachineIdentityResponsePtrOutput { return v.Identity }).(VirtualMachineIdentityResponsePtrOutput)
}

// The virtual machine instance view.
func (o VirtualMachineOutput) InstanceView() VirtualMachineInstanceViewResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachineInstanceViewResponseOutput { return v.InstanceView }).(VirtualMachineInstanceViewResponseOutput)
}

// Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. <br><br> Possible values are: <br><br> Windows_Client <br><br> Windows_Server <br><br> If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Minimum api-version: 2015-06-15
func (o VirtualMachineOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// Resource location
func (o VirtualMachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o VirtualMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the network interfaces of the virtual machine.
func (o VirtualMachineOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

// Specifies the operating system settings for the virtual machine.
func (o VirtualMachineOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) OSProfileResponsePtrOutput { return v.OsProfile }).(OSProfileResponsePtrOutput)
}

// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
func (o VirtualMachineOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) PlanResponsePtrOutput { return v.Plan }).(PlanResponsePtrOutput)
}

// Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
func (o VirtualMachineOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Priority }).(pulumi.StringPtrOutput)
}

// The provisioning state, which only appears in the response.
func (o VirtualMachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
func (o VirtualMachineOutput) ProximityPlacementGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) SubResourceResponsePtrOutput { return v.ProximityPlacementGroup }).(SubResourceResponsePtrOutput)
}

// The virtual machine child extension resources.
func (o VirtualMachineOutput) Resources() VirtualMachineExtensionResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachineExtensionResponseArrayOutput { return v.Resources }).(VirtualMachineExtensionResponseArrayOutput)
}

// Specifies the storage settings for the virtual machine disks.
func (o VirtualMachineOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) StorageProfileResponsePtrOutput { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

// Resource tags
func (o VirtualMachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o VirtualMachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
func (o VirtualMachineOutput) VirtualMachineScaleSet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) SubResourceResponsePtrOutput { return v.VirtualMachineScaleSet }).(SubResourceResponsePtrOutput)
}

// Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
func (o VirtualMachineOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.VmId }).(pulumi.StringOutput)
}

// The virtual machine zones.
func (o VirtualMachineOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineOutput{})
}
