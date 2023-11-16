// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurestackhci

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The virtual machine resource definition.
// Azure REST API version: 2022-12-15-preview.
type VirtualMachine struct {
	pulumi.CustomResourceState

	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// Guest agent status properties.
	GuestAgentProfile GuestAgentProfileResponsePtrOutput `pulumi:"guestAgentProfile"`
	// HardwareProfile - Specifies the hardware settings for the virtual machine.
	HardwareProfile VirtualMachinePropertiesResponseHardwareProfilePtrOutput `pulumi:"hardwareProfile"`
	// Identity for the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// NetworkProfile - describes the network configuration the virtual machine
	NetworkProfile VirtualMachinePropertiesResponseNetworkProfilePtrOutput `pulumi:"networkProfile"`
	// OsProfile - describes the configuration of the operating system and sets login data
	OsProfile VirtualMachinePropertiesResponseOsProfilePtrOutput `pulumi:"osProfile"`
	// Provisioning state of the virtual machine.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// SecurityProfile - Specifies the security settings for the virtual machine.
	SecurityProfile VirtualMachinePropertiesResponseSecurityProfilePtrOutput `pulumi:"securityProfile"`
	// The observed state of virtual machines
	Status VirtualMachineStatusResponseOutput `pulumi:"status"`
	// StorageProfile - contains information about the disks and storage information for the virtual machine
	StorageProfile VirtualMachinePropertiesResponseStorageProfilePtrOutput `pulumi:"storageProfile"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique identifier for the vm resource.
	VmId pulumi.StringOutput `pulumi:"vmId"`
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
	if args.HardwareProfile != nil {
		args.HardwareProfile = args.HardwareProfile.ToVirtualMachinePropertiesHardwareProfilePtrOutput().ApplyT(func(v *VirtualMachinePropertiesHardwareProfile) *VirtualMachinePropertiesHardwareProfile {
			return v.Defaults()
		}).(VirtualMachinePropertiesHardwareProfilePtrOutput)
	}
	if args.SecurityProfile != nil {
		args.SecurityProfile = args.SecurityProfile.ToVirtualMachinePropertiesSecurityProfilePtrOutput().ApplyT(func(v *VirtualMachinePropertiesSecurityProfile) *VirtualMachinePropertiesSecurityProfile {
			return v.Defaults()
		}).(VirtualMachinePropertiesSecurityProfilePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210701preview:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221215preview:VirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualMachine
	err := ctx.RegisterResource("azure-native:azurestackhci:VirtualMachine", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:azurestackhci:VirtualMachine", name, id, state, &resource, opts...)
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
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// HardwareProfile - Specifies the hardware settings for the virtual machine.
	HardwareProfile *VirtualMachinePropertiesHardwareProfile `pulumi:"hardwareProfile"`
	// Identity for the resource.
	Identity *Identity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// NetworkProfile - describes the network configuration the virtual machine
	NetworkProfile *VirtualMachinePropertiesNetworkProfile `pulumi:"networkProfile"`
	// OsProfile - describes the configuration of the operating system and sets login data
	OsProfile *VirtualMachinePropertiesOsProfile `pulumi:"osProfile"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SecurityProfile - Specifies the security settings for the virtual machine.
	SecurityProfile *VirtualMachinePropertiesSecurityProfile `pulumi:"securityProfile"`
	// StorageProfile - contains information about the disks and storage information for the virtual machine
	StorageProfile *VirtualMachinePropertiesStorageProfile `pulumi:"storageProfile"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Name of the virtual machine
	VirtualMachineName *string `pulumi:"virtualMachineName"`
}

// The set of arguments for constructing a VirtualMachine resource.
type VirtualMachineArgs struct {
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationPtrInput
	// HardwareProfile - Specifies the hardware settings for the virtual machine.
	HardwareProfile VirtualMachinePropertiesHardwareProfilePtrInput
	// Identity for the resource.
	Identity IdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// NetworkProfile - describes the network configuration the virtual machine
	NetworkProfile VirtualMachinePropertiesNetworkProfilePtrInput
	// OsProfile - describes the configuration of the operating system and sets login data
	OsProfile VirtualMachinePropertiesOsProfilePtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// SecurityProfile - Specifies the security settings for the virtual machine.
	SecurityProfile VirtualMachinePropertiesSecurityProfilePtrInput
	// StorageProfile - contains information about the disks and storage information for the virtual machine
	StorageProfile VirtualMachinePropertiesStorageProfilePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Name of the virtual machine
	VirtualMachineName pulumi.StringPtrInput
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

// The extendedLocation of the resource.
func (o VirtualMachineOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Guest agent status properties.
func (o VirtualMachineOutput) GuestAgentProfile() GuestAgentProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) GuestAgentProfileResponsePtrOutput { return v.GuestAgentProfile }).(GuestAgentProfileResponsePtrOutput)
}

// HardwareProfile - Specifies the hardware settings for the virtual machine.
func (o VirtualMachineOutput) HardwareProfile() VirtualMachinePropertiesResponseHardwareProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachinePropertiesResponseHardwareProfilePtrOutput {
		return v.HardwareProfile
	}).(VirtualMachinePropertiesResponseHardwareProfilePtrOutput)
}

// Identity for the resource.
func (o VirtualMachineOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o VirtualMachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o VirtualMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// NetworkProfile - describes the network configuration the virtual machine
func (o VirtualMachineOutput) NetworkProfile() VirtualMachinePropertiesResponseNetworkProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachinePropertiesResponseNetworkProfilePtrOutput {
		return v.NetworkProfile
	}).(VirtualMachinePropertiesResponseNetworkProfilePtrOutput)
}

// OsProfile - describes the configuration of the operating system and sets login data
func (o VirtualMachineOutput) OsProfile() VirtualMachinePropertiesResponseOsProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachinePropertiesResponseOsProfilePtrOutput { return v.OsProfile }).(VirtualMachinePropertiesResponseOsProfilePtrOutput)
}

// Provisioning state of the virtual machine.
func (o VirtualMachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// SecurityProfile - Specifies the security settings for the virtual machine.
func (o VirtualMachineOutput) SecurityProfile() VirtualMachinePropertiesResponseSecurityProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachinePropertiesResponseSecurityProfilePtrOutput {
		return v.SecurityProfile
	}).(VirtualMachinePropertiesResponseSecurityProfilePtrOutput)
}

// The observed state of virtual machines
func (o VirtualMachineOutput) Status() VirtualMachineStatusResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachineStatusResponseOutput { return v.Status }).(VirtualMachineStatusResponseOutput)
}

// StorageProfile - contains information about the disks and storage information for the virtual machine
func (o VirtualMachineOutput) StorageProfile() VirtualMachinePropertiesResponseStorageProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachinePropertiesResponseStorageProfilePtrOutput {
		return v.StorageProfile
	}).(VirtualMachinePropertiesResponseStorageProfilePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o VirtualMachineOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o VirtualMachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o VirtualMachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Unique identifier for the vm resource.
func (o VirtualMachineOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.VmId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineOutput{})
}
