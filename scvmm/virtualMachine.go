// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scvmm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The VirtualMachines resource definition.
// Azure REST API version: 2022-05-21-preview. Prior API version in Azure Native 1.x: 2020-06-05-preview.
//
// Other available API versions: 2023-04-01-preview.
type VirtualMachine struct {
	pulumi.CustomResourceState

	// Availability Sets in vm.
	AvailabilitySets VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput `pulumi:"availabilitySets"`
	// Type of checkpoint supported for the vm.
	CheckpointType pulumi.StringPtrOutput `pulumi:"checkpointType"`
	// Checkpoints in the vm.
	Checkpoints CheckpointResponseArrayOutput `pulumi:"checkpoints"`
	// ARM Id of the cloud resource to use for deploying the vm.
	CloudId pulumi.StringPtrOutput `pulumi:"cloudId"`
	// The extended location.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// Gets or sets the generation for the vm.
	Generation pulumi.IntPtrOutput `pulumi:"generation"`
	// Guest agent status properties.
	GuestAgentProfile GuestAgentProfileResponsePtrOutput `pulumi:"guestAgentProfile"`
	// Hardware properties.
	HardwareProfile HardwareProfileResponsePtrOutput `pulumi:"hardwareProfile"`
	// The identity of the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId pulumi.StringPtrOutput `pulumi:"inventoryItemId"`
	// Last restored checkpoint in the vm.
	LastRestoredVMCheckpoint CheckpointResponseOutput `pulumi:"lastRestoredVMCheckpoint"`
	// Gets or sets the location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Network properties.
	NetworkProfile NetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// OS properties.
	OsProfile OsProfileResponsePtrOutput `pulumi:"osProfile"`
	// Gets the power state of the virtual machine.
	PowerState pulumi.StringOutput `pulumi:"powerState"`
	// Gets or sets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Storage properties.
	StorageProfile StorageProfileResponsePtrOutput `pulumi:"storageProfile"`
	// The system data.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// ARM Id of the template resource to use for deploying the vm.
	TemplateId pulumi.StringPtrOutput `pulumi:"templateId"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique ID of the virtual machine.
	Uuid pulumi.StringPtrOutput `pulumi:"uuid"`
	// VMName is the name of VM on the SCVMM server.
	VmName pulumi.StringPtrOutput `pulumi:"vmName"`
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId pulumi.StringPtrOutput `pulumi:"vmmServerId"`
}

// NewVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
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
			Type: pulumi.String("azure-native:scvmm/v20200605preview:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20220521preview:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20230401preview:VirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualMachine
	err := ctx.RegisterResource("azure-native:scvmm:VirtualMachine", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:scvmm:VirtualMachine", name, id, state, &resource, opts...)
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
	// Availability Sets in vm.
	AvailabilitySets []VirtualMachinePropertiesAvailabilitySets `pulumi:"availabilitySets"`
	// Type of checkpoint supported for the vm.
	CheckpointType *string `pulumi:"checkpointType"`
	// Checkpoints in the vm.
	Checkpoints []Checkpoint `pulumi:"checkpoints"`
	// ARM Id of the cloud resource to use for deploying the vm.
	CloudId *string `pulumi:"cloudId"`
	// The extended location.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// Gets or sets the generation for the vm.
	Generation *int `pulumi:"generation"`
	// Guest agent status properties.
	GuestAgentProfile *GuestAgentProfile `pulumi:"guestAgentProfile"`
	// Hardware properties.
	HardwareProfile *HardwareProfile `pulumi:"hardwareProfile"`
	// The identity of the resource.
	Identity *Identity `pulumi:"identity"`
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// Network properties.
	NetworkProfile *NetworkProfile `pulumi:"networkProfile"`
	// OS properties.
	OsProfile *OsProfile `pulumi:"osProfile"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Storage properties.
	StorageProfile *StorageProfile `pulumi:"storageProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// ARM Id of the template resource to use for deploying the vm.
	TemplateId *string `pulumi:"templateId"`
	// Unique ID of the virtual machine.
	Uuid *string `pulumi:"uuid"`
	// Name of the VirtualMachine.
	VirtualMachineName *string `pulumi:"virtualMachineName"`
	// VMName is the name of VM on the SCVMM server.
	VmName *string `pulumi:"vmName"`
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId *string `pulumi:"vmmServerId"`
}

// The set of arguments for constructing a VirtualMachine resource.
type VirtualMachineArgs struct {
	// Availability Sets in vm.
	AvailabilitySets VirtualMachinePropertiesAvailabilitySetsArrayInput
	// Type of checkpoint supported for the vm.
	CheckpointType pulumi.StringPtrInput
	// Checkpoints in the vm.
	Checkpoints CheckpointArrayInput
	// ARM Id of the cloud resource to use for deploying the vm.
	CloudId pulumi.StringPtrInput
	// The extended location.
	ExtendedLocation ExtendedLocationInput
	// Gets or sets the generation for the vm.
	Generation pulumi.IntPtrInput
	// Guest agent status properties.
	GuestAgentProfile GuestAgentProfilePtrInput
	// Hardware properties.
	HardwareProfile HardwareProfilePtrInput
	// The identity of the resource.
	Identity IdentityPtrInput
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId pulumi.StringPtrInput
	// Gets or sets the location.
	Location pulumi.StringPtrInput
	// Network properties.
	NetworkProfile NetworkProfilePtrInput
	// OS properties.
	OsProfile OsProfilePtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Storage properties.
	StorageProfile StorageProfilePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// ARM Id of the template resource to use for deploying the vm.
	TemplateId pulumi.StringPtrInput
	// Unique ID of the virtual machine.
	Uuid pulumi.StringPtrInput
	// Name of the VirtualMachine.
	VirtualMachineName pulumi.StringPtrInput
	// VMName is the name of VM on the SCVMM server.
	VmName pulumi.StringPtrInput
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId pulumi.StringPtrInput
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

func (i *VirtualMachine) ToOutput(ctx context.Context) pulumix.Output[*VirtualMachine] {
	return pulumix.Output[*VirtualMachine]{
		OutputState: i.ToVirtualMachineOutputWithContext(ctx).OutputState,
	}
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

func (o VirtualMachineOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualMachine] {
	return pulumix.Output[*VirtualMachine]{
		OutputState: o.OutputState,
	}
}

// Availability Sets in vm.
func (o VirtualMachineOutput) AvailabilitySets() VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput {
		return v.AvailabilitySets
	}).(VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput)
}

// Type of checkpoint supported for the vm.
func (o VirtualMachineOutput) CheckpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.CheckpointType }).(pulumi.StringPtrOutput)
}

// Checkpoints in the vm.
func (o VirtualMachineOutput) Checkpoints() CheckpointResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) CheckpointResponseArrayOutput { return v.Checkpoints }).(CheckpointResponseArrayOutput)
}

// ARM Id of the cloud resource to use for deploying the vm.
func (o VirtualMachineOutput) CloudId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.CloudId }).(pulumi.StringPtrOutput)
}

// The extended location.
func (o VirtualMachineOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Gets or sets the generation for the vm.
func (o VirtualMachineOutput) Generation() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntPtrOutput { return v.Generation }).(pulumi.IntPtrOutput)
}

// Guest agent status properties.
func (o VirtualMachineOutput) GuestAgentProfile() GuestAgentProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) GuestAgentProfileResponsePtrOutput { return v.GuestAgentProfile }).(GuestAgentProfileResponsePtrOutput)
}

// Hardware properties.
func (o VirtualMachineOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) HardwareProfileResponsePtrOutput { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

// The identity of the resource.
func (o VirtualMachineOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// Gets or sets the inventory Item ID for the resource.
func (o VirtualMachineOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Last restored checkpoint in the vm.
func (o VirtualMachineOutput) LastRestoredVMCheckpoint() CheckpointResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) CheckpointResponseOutput { return v.LastRestoredVMCheckpoint }).(CheckpointResponseOutput)
}

// Gets or sets the location.
func (o VirtualMachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource Name
func (o VirtualMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network properties.
func (o VirtualMachineOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

// OS properties.
func (o VirtualMachineOutput) OsProfile() OsProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) OsProfileResponsePtrOutput { return v.OsProfile }).(OsProfileResponsePtrOutput)
}

// Gets the power state of the virtual machine.
func (o VirtualMachineOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.PowerState }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state.
func (o VirtualMachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Storage properties.
func (o VirtualMachineOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) StorageProfileResponsePtrOutput { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

// The system data.
func (o VirtualMachineOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o VirtualMachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// ARM Id of the template resource to use for deploying the vm.
func (o VirtualMachineOutput) TemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.TemplateId }).(pulumi.StringPtrOutput)
}

// Resource Type
func (o VirtualMachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Unique ID of the virtual machine.
func (o VirtualMachineOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Uuid }).(pulumi.StringPtrOutput)
}

// VMName is the name of VM on the SCVMM server.
func (o VirtualMachineOutput) VmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.VmName }).(pulumi.StringPtrOutput)
}

// ARM Id of the vmmServer resource in which this resource resides.
func (o VirtualMachineOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineOutput{})
}
