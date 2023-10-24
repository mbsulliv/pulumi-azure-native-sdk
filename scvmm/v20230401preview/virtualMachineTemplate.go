// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The VirtualMachineTemplates resource definition.
type VirtualMachineTemplate struct {
	pulumi.CustomResourceState

	// Gets or sets computer name.
	ComputerName pulumi.StringOutput `pulumi:"computerName"`
	// Gets or sets the desired number of vCPUs for the vm.
	CpuCount pulumi.IntOutput `pulumi:"cpuCount"`
	// Gets or sets the disks of the template.
	Disks VirtualDiskResponseArrayOutput `pulumi:"disks"`
	// Gets or sets a value indicating whether to enable dynamic memory or not.
	DynamicMemoryEnabled pulumi.StringOutput `pulumi:"dynamicMemoryEnabled"`
	// Gets or sets the max dynamic memory for the vm.
	DynamicMemoryMaxMB pulumi.IntOutput `pulumi:"dynamicMemoryMaxMB"`
	// Gets or sets the min dynamic memory for the vm.
	DynamicMemoryMinMB pulumi.IntOutput `pulumi:"dynamicMemoryMinMB"`
	// The extended location.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// Gets or sets the generation for the vm.
	Generation pulumi.IntOutput `pulumi:"generation"`
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId pulumi.StringPtrOutput `pulumi:"inventoryItemId"`
	// Gets or sets a value indicating whether the vm template is customizable or not.
	IsCustomizable pulumi.StringOutput `pulumi:"isCustomizable"`
	// Gets highly available property.
	IsHighlyAvailable pulumi.StringOutput `pulumi:"isHighlyAvailable"`
	// Gets or sets a value indicating whether to enable processor compatibility mode for live migration of VMs.
	LimitCpuForMigration pulumi.StringOutput `pulumi:"limitCpuForMigration"`
	// Gets or sets the location.
	Location pulumi.StringOutput `pulumi:"location"`
	// MemoryMB is the desired size of a virtual machine's memory, in MB.
	MemoryMB pulumi.IntOutput `pulumi:"memoryMB"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the network interfaces of the template.
	NetworkInterfaces NetworkInterfacesResponseArrayOutput `pulumi:"networkInterfaces"`
	// Gets or sets os name.
	OsName pulumi.StringOutput `pulumi:"osName"`
	// Gets or sets the type of the os.
	OsType pulumi.StringOutput `pulumi:"osType"`
	// Gets or sets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system data.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique ID of the virtual machine template.
	Uuid pulumi.StringPtrOutput `pulumi:"uuid"`
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId pulumi.StringPtrOutput `pulumi:"vmmServerId"`
}

// NewVirtualMachineTemplate registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineTemplate(ctx *pulumi.Context,
	name string, args *VirtualMachineTemplateArgs, opts ...pulumi.ResourceOption) (*VirtualMachineTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scvmm:VirtualMachineTemplate"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20200605preview:VirtualMachineTemplate"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20220521preview:VirtualMachineTemplate"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20231007:VirtualMachineTemplate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualMachineTemplate
	err := ctx.RegisterResource("azure-native:scvmm/v20230401preview:VirtualMachineTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineTemplate gets an existing VirtualMachineTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineTemplateState, opts ...pulumi.ResourceOption) (*VirtualMachineTemplate, error) {
	var resource VirtualMachineTemplate
	err := ctx.ReadResource("azure-native:scvmm/v20230401preview:VirtualMachineTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineTemplate resources.
type virtualMachineTemplateState struct {
}

type VirtualMachineTemplateState struct {
}

func (VirtualMachineTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineTemplateState)(nil)).Elem()
}

type virtualMachineTemplateArgs struct {
	// The extended location.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Unique ID of the virtual machine template.
	Uuid *string `pulumi:"uuid"`
	// Name of the VirtualMachineTemplate.
	VirtualMachineTemplateName *string `pulumi:"virtualMachineTemplateName"`
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId *string `pulumi:"vmmServerId"`
}

// The set of arguments for constructing a VirtualMachineTemplate resource.
type VirtualMachineTemplateArgs struct {
	// The extended location.
	ExtendedLocation ExtendedLocationInput
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId pulumi.StringPtrInput
	// Gets or sets the location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Unique ID of the virtual machine template.
	Uuid pulumi.StringPtrInput
	// Name of the VirtualMachineTemplate.
	VirtualMachineTemplateName pulumi.StringPtrInput
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId pulumi.StringPtrInput
}

func (VirtualMachineTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineTemplateArgs)(nil)).Elem()
}

type VirtualMachineTemplateInput interface {
	pulumi.Input

	ToVirtualMachineTemplateOutput() VirtualMachineTemplateOutput
	ToVirtualMachineTemplateOutputWithContext(ctx context.Context) VirtualMachineTemplateOutput
}

func (*VirtualMachineTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineTemplate)(nil)).Elem()
}

func (i *VirtualMachineTemplate) ToVirtualMachineTemplateOutput() VirtualMachineTemplateOutput {
	return i.ToVirtualMachineTemplateOutputWithContext(context.Background())
}

func (i *VirtualMachineTemplate) ToVirtualMachineTemplateOutputWithContext(ctx context.Context) VirtualMachineTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineTemplateOutput)
}

func (i *VirtualMachineTemplate) ToOutput(ctx context.Context) pulumix.Output[*VirtualMachineTemplate] {
	return pulumix.Output[*VirtualMachineTemplate]{
		OutputState: i.ToVirtualMachineTemplateOutputWithContext(ctx).OutputState,
	}
}

type VirtualMachineTemplateOutput struct{ *pulumi.OutputState }

func (VirtualMachineTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineTemplate)(nil)).Elem()
}

func (o VirtualMachineTemplateOutput) ToVirtualMachineTemplateOutput() VirtualMachineTemplateOutput {
	return o
}

func (o VirtualMachineTemplateOutput) ToVirtualMachineTemplateOutputWithContext(ctx context.Context) VirtualMachineTemplateOutput {
	return o
}

func (o VirtualMachineTemplateOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualMachineTemplate] {
	return pulumix.Output[*VirtualMachineTemplate]{
		OutputState: o.OutputState,
	}
}

// Gets or sets computer name.
func (o VirtualMachineTemplateOutput) ComputerName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.ComputerName }).(pulumi.StringOutput)
}

// Gets or sets the desired number of vCPUs for the vm.
func (o VirtualMachineTemplateOutput) CpuCount() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.IntOutput { return v.CpuCount }).(pulumi.IntOutput)
}

// Gets or sets the disks of the template.
func (o VirtualMachineTemplateOutput) Disks() VirtualDiskResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) VirtualDiskResponseArrayOutput { return v.Disks }).(VirtualDiskResponseArrayOutput)
}

// Gets or sets a value indicating whether to enable dynamic memory or not.
func (o VirtualMachineTemplateOutput) DynamicMemoryEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.DynamicMemoryEnabled }).(pulumi.StringOutput)
}

// Gets or sets the max dynamic memory for the vm.
func (o VirtualMachineTemplateOutput) DynamicMemoryMaxMB() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.IntOutput { return v.DynamicMemoryMaxMB }).(pulumi.IntOutput)
}

// Gets or sets the min dynamic memory for the vm.
func (o VirtualMachineTemplateOutput) DynamicMemoryMinMB() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.IntOutput { return v.DynamicMemoryMinMB }).(pulumi.IntOutput)
}

// The extended location.
func (o VirtualMachineTemplateOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Gets or sets the generation for the vm.
func (o VirtualMachineTemplateOutput) Generation() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.IntOutput { return v.Generation }).(pulumi.IntOutput)
}

// Gets or sets the inventory Item ID for the resource.
func (o VirtualMachineTemplateOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Gets or sets a value indicating whether the vm template is customizable or not.
func (o VirtualMachineTemplateOutput) IsCustomizable() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.IsCustomizable }).(pulumi.StringOutput)
}

// Gets highly available property.
func (o VirtualMachineTemplateOutput) IsHighlyAvailable() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.IsHighlyAvailable }).(pulumi.StringOutput)
}

// Gets or sets a value indicating whether to enable processor compatibility mode for live migration of VMs.
func (o VirtualMachineTemplateOutput) LimitCpuForMigration() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.LimitCpuForMigration }).(pulumi.StringOutput)
}

// Gets or sets the location.
func (o VirtualMachineTemplateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// MemoryMB is the desired size of a virtual machine's memory, in MB.
func (o VirtualMachineTemplateOutput) MemoryMB() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.IntOutput { return v.MemoryMB }).(pulumi.IntOutput)
}

// Resource Name
func (o VirtualMachineTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the network interfaces of the template.
func (o VirtualMachineTemplateOutput) NetworkInterfaces() NetworkInterfacesResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) NetworkInterfacesResponseArrayOutput { return v.NetworkInterfaces }).(NetworkInterfacesResponseArrayOutput)
}

// Gets or sets os name.
func (o VirtualMachineTemplateOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.OsName }).(pulumi.StringOutput)
}

// Gets or sets the type of the os.
func (o VirtualMachineTemplateOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.OsType }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state.
func (o VirtualMachineTemplateOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system data.
func (o VirtualMachineTemplateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o VirtualMachineTemplateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource Type
func (o VirtualMachineTemplateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Unique ID of the virtual machine template.
func (o VirtualMachineTemplateOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringPtrOutput { return v.Uuid }).(pulumi.StringPtrOutput)
}

// ARM Id of the vmmServer resource in which this resource resides.
func (o VirtualMachineTemplateOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringPtrOutput { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineTemplateOutput{})
}
