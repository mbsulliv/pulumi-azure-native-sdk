// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define the virtualMachineInstance.
type VirtualMachineInstance struct {
	pulumi.CustomResourceState

	// Gets or sets the extended location.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// Hardware properties.
	HardwareProfile HardwareProfileResponsePtrOutput `pulumi:"hardwareProfile"`
	// Gets the infrastructure profile.
	InfrastructureProfile InfrastructureProfileResponsePtrOutput `pulumi:"infrastructureProfile"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Network properties.
	NetworkProfile NetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// OS properties.
	OsProfile OsProfileForVMInstanceResponsePtrOutput `pulumi:"osProfile"`
	// Placement properties.
	PlacementProfile PlacementProfileResponsePtrOutput `pulumi:"placementProfile"`
	// Gets the power state of the virtual machine.
	PowerState pulumi.StringOutput `pulumi:"powerState"`
	// Gets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets or sets a unique identifier for the vm resource.
	ResourceUid pulumi.StringOutput `pulumi:"resourceUid"`
	// Gets the security profile.
	SecurityProfile SecurityProfileResponsePtrOutput `pulumi:"securityProfile"`
	// The resource status information.
	Statuses ResourceStatusResponseArrayOutput `pulumi:"statuses"`
	// Storage properties.
	StorageProfile StorageProfileResponsePtrOutput `pulumi:"storageProfile"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualMachineInstance registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineInstance(ctx *pulumi.Context,
	name string, args *VirtualMachineInstanceArgs, opts ...pulumi.ResourceOption) (*VirtualMachineInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere:VirtualMachineInstance"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20231001:VirtualMachineInstance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualMachineInstance
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere/v20230301preview:VirtualMachineInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineInstance gets an existing VirtualMachineInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineInstanceState, opts ...pulumi.ResourceOption) (*VirtualMachineInstance, error) {
	var resource VirtualMachineInstance
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere/v20230301preview:VirtualMachineInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineInstance resources.
type virtualMachineInstanceState struct {
}

type VirtualMachineInstanceState struct {
}

func (VirtualMachineInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineInstanceState)(nil)).Elem()
}

type virtualMachineInstanceArgs struct {
	// Gets or sets the extended location.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Hardware properties.
	HardwareProfile *HardwareProfile `pulumi:"hardwareProfile"`
	// Gets the infrastructure profile.
	InfrastructureProfile *InfrastructureProfile `pulumi:"infrastructureProfile"`
	// Network properties.
	NetworkProfile *NetworkProfile `pulumi:"networkProfile"`
	// OS properties.
	OsProfile *OsProfileForVMInstance `pulumi:"osProfile"`
	// Placement properties.
	PlacementProfile *PlacementProfile `pulumi:"placementProfile"`
	// The fully qualified Azure Resource manager identifier of the Hybrid Compute machine resource to be extended.
	ResourceUri string `pulumi:"resourceUri"`
	// Gets the security profile.
	SecurityProfile *SecurityProfile `pulumi:"securityProfile"`
	// Storage properties.
	StorageProfile *StorageProfile `pulumi:"storageProfile"`
}

// The set of arguments for constructing a VirtualMachineInstance resource.
type VirtualMachineInstanceArgs struct {
	// Gets or sets the extended location.
	ExtendedLocation ExtendedLocationPtrInput
	// Hardware properties.
	HardwareProfile HardwareProfilePtrInput
	// Gets the infrastructure profile.
	InfrastructureProfile InfrastructureProfilePtrInput
	// Network properties.
	NetworkProfile NetworkProfilePtrInput
	// OS properties.
	OsProfile OsProfileForVMInstancePtrInput
	// Placement properties.
	PlacementProfile PlacementProfilePtrInput
	// The fully qualified Azure Resource manager identifier of the Hybrid Compute machine resource to be extended.
	ResourceUri pulumi.StringInput
	// Gets the security profile.
	SecurityProfile SecurityProfilePtrInput
	// Storage properties.
	StorageProfile StorageProfilePtrInput
}

func (VirtualMachineInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineInstanceArgs)(nil)).Elem()
}

type VirtualMachineInstanceInput interface {
	pulumi.Input

	ToVirtualMachineInstanceOutput() VirtualMachineInstanceOutput
	ToVirtualMachineInstanceOutputWithContext(ctx context.Context) VirtualMachineInstanceOutput
}

func (*VirtualMachineInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineInstance)(nil)).Elem()
}

func (i *VirtualMachineInstance) ToVirtualMachineInstanceOutput() VirtualMachineInstanceOutput {
	return i.ToVirtualMachineInstanceOutputWithContext(context.Background())
}

func (i *VirtualMachineInstance) ToVirtualMachineInstanceOutputWithContext(ctx context.Context) VirtualMachineInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineInstanceOutput)
}

type VirtualMachineInstanceOutput struct{ *pulumi.OutputState }

func (VirtualMachineInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineInstance)(nil)).Elem()
}

func (o VirtualMachineInstanceOutput) ToVirtualMachineInstanceOutput() VirtualMachineInstanceOutput {
	return o
}

func (o VirtualMachineInstanceOutput) ToVirtualMachineInstanceOutputWithContext(ctx context.Context) VirtualMachineInstanceOutput {
	return o
}

// Gets or sets the extended location.
func (o VirtualMachineInstanceOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Hardware properties.
func (o VirtualMachineInstanceOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) HardwareProfileResponsePtrOutput { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

// Gets the infrastructure profile.
func (o VirtualMachineInstanceOutput) InfrastructureProfile() InfrastructureProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) InfrastructureProfileResponsePtrOutput { return v.InfrastructureProfile }).(InfrastructureProfileResponsePtrOutput)
}

// The name of the resource
func (o VirtualMachineInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network properties.
func (o VirtualMachineInstanceOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

// OS properties.
func (o VirtualMachineInstanceOutput) OsProfile() OsProfileForVMInstanceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) OsProfileForVMInstanceResponsePtrOutput { return v.OsProfile }).(OsProfileForVMInstanceResponsePtrOutput)
}

// Placement properties.
func (o VirtualMachineInstanceOutput) PlacementProfile() PlacementProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) PlacementProfileResponsePtrOutput { return v.PlacementProfile }).(PlacementProfileResponsePtrOutput)
}

// Gets the power state of the virtual machine.
func (o VirtualMachineInstanceOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) pulumi.StringOutput { return v.PowerState }).(pulumi.StringOutput)
}

// Gets the provisioning state.
func (o VirtualMachineInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for the vm resource.
func (o VirtualMachineInstanceOutput) ResourceUid() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) pulumi.StringOutput { return v.ResourceUid }).(pulumi.StringOutput)
}

// Gets the security profile.
func (o VirtualMachineInstanceOutput) SecurityProfile() SecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) SecurityProfileResponsePtrOutput { return v.SecurityProfile }).(SecurityProfileResponsePtrOutput)
}

// The resource status information.
func (o VirtualMachineInstanceOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// Storage properties.
func (o VirtualMachineInstanceOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) StorageProfileResponsePtrOutput { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o VirtualMachineInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o VirtualMachineInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineInstanceOutput{})
}
