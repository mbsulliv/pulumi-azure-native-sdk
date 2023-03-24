// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221212preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachine struct {
	pulumi.CustomResourceState

	// The name of the administrator to which the ssh public keys will be added into the authorized keys.
	AdminUsername pulumi.StringOutput `pulumi:"adminUsername"`
	// The resource ID of the bare metal machine the virtual machine has landed to.
	BareMetalMachineId pulumi.StringOutput `pulumi:"bareMetalMachineId"`
	// Selects the boot method for the virtual machine.
	BootMethod                     pulumi.StringPtrOutput          `pulumi:"bootMethod"`
	CloudServicesNetworkAttachment NetworkAttachmentResponseOutput `pulumi:"cloudServicesNetworkAttachment"`
	// The resource ID of the cluster the virtual machine is created for.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The number of CPU cores in the virtual machine.
	CpuCores pulumi.Float64Output `pulumi:"cpuCores"`
	// The more detailed status of the virtual machine.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage pulumi.StringOutput `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// Field Deprecated, the value will be ignored if provided. The indicator of whether one of the specified CPU cores is isolated to run the emulator thread for this virtual machine.
	IsolateEmulatorThread pulumi.StringPtrOutput `pulumi:"isolateEmulatorThread"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The memory size of the virtual machine in GB.
	MemorySizeGB pulumi.Float64Output `pulumi:"memorySizeGB"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of network attachments to the virtual machine.
	NetworkAttachments NetworkAttachmentResponseArrayOutput `pulumi:"networkAttachments"`
	// The Base64 encoded cloud-init network data.
	NetworkData pulumi.StringPtrOutput `pulumi:"networkData"`
	// The scheduling hints for the virtual machine.
	PlacementHints VirtualMachinePlacementHintResponseArrayOutput `pulumi:"placementHints"`
	// The power state of the virtual machine.
	PowerState pulumi.StringOutput `pulumi:"powerState"`
	// The provisioning state of the virtual machine.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The list of ssh public keys. Each key will be added to the virtual machine using the cloud-init ssh_authorized_keys mechanism for the adminUsername.
	SshPublicKeys  SshPublicKeyResponseArrayOutput `pulumi:"sshPublicKeys"`
	StorageProfile StorageProfileResponseOutput    `pulumi:"storageProfile"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The Base64 encoded cloud-init user data.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// Field Deprecated, use virtualizationModel instead. The type of the virtio interface.
	VirtioInterface pulumi.StringPtrOutput `pulumi:"virtioInterface"`
	// The type of the device model to use.
	VmDeviceModel pulumi.StringPtrOutput `pulumi:"vmDeviceModel"`
	// The virtual machine image that is currently provisioned to the OS disk, using the full url and tag notation used to pull the image.
	VmImage                      pulumi.StringOutput                         `pulumi:"vmImage"`
	VmImageRepositoryCredentials ImageRepositoryCredentialsResponsePtrOutput `pulumi:"vmImageRepositoryCredentials"`
	// The resource IDs of volumes that are attached to the virtual machine.
	Volumes pulumi.StringArrayOutput `pulumi:"volumes"`
}

// NewVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminUsername == nil {
		return nil, errors.New("invalid value for required argument 'AdminUsername'")
	}
	if args.CloudServicesNetworkAttachment == nil {
		return nil, errors.New("invalid value for required argument 'CloudServicesNetworkAttachment'")
	}
	if args.CpuCores == nil {
		return nil, errors.New("invalid value for required argument 'CpuCores'")
	}
	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.MemorySizeGB == nil {
		return nil, errors.New("invalid value for required argument 'MemorySizeGB'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageProfile == nil {
		return nil, errors.New("invalid value for required argument 'StorageProfile'")
	}
	if args.VmImage == nil {
		return nil, errors.New("invalid value for required argument 'VmImage'")
	}
	if isZero(args.BootMethod) {
		args.BootMethod = pulumi.StringPtr("UEFI")
	}
	if isZero(args.IsolateEmulatorThread) {
		args.IsolateEmulatorThread = pulumi.StringPtr("True")
	}
	args.StorageProfile = args.StorageProfile.ToStorageProfileOutput().ApplyT(func(v StorageProfile) StorageProfile { return *v.Defaults() }).(StorageProfileOutput)
	if isZero(args.VirtioInterface) {
		args.VirtioInterface = pulumi.StringPtr("Modern")
	}
	if isZero(args.VmDeviceModel) {
		args.VmDeviceModel = pulumi.StringPtr("T2")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkcloud:VirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachine
	err := ctx.RegisterResource("azure-native:networkcloud/v20221212preview:VirtualMachine", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:networkcloud/v20221212preview:VirtualMachine", name, id, state, &resource, opts...)
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
	// The name of the administrator to which the ssh public keys will be added into the authorized keys.
	AdminUsername string `pulumi:"adminUsername"`
	// Selects the boot method for the virtual machine.
	BootMethod                     *string           `pulumi:"bootMethod"`
	CloudServicesNetworkAttachment NetworkAttachment `pulumi:"cloudServicesNetworkAttachment"`
	// The number of CPU cores in the virtual machine.
	CpuCores float64 `pulumi:"cpuCores"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// Field Deprecated, the value will be ignored if provided. The indicator of whether one of the specified CPU cores is isolated to run the emulator thread for this virtual machine.
	IsolateEmulatorThread *string `pulumi:"isolateEmulatorThread"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The memory size of the virtual machine in GB.
	MemorySizeGB float64 `pulumi:"memorySizeGB"`
	// The list of network attachments to the virtual machine.
	NetworkAttachments []NetworkAttachment `pulumi:"networkAttachments"`
	// The Base64 encoded cloud-init network data.
	NetworkData *string `pulumi:"networkData"`
	// The scheduling hints for the virtual machine.
	PlacementHints []VirtualMachinePlacementHint `pulumi:"placementHints"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of ssh public keys. Each key will be added to the virtual machine using the cloud-init ssh_authorized_keys mechanism for the adminUsername.
	SshPublicKeys  []SshPublicKey `pulumi:"sshPublicKeys"`
	StorageProfile StorageProfile `pulumi:"storageProfile"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The Base64 encoded cloud-init user data.
	UserData *string `pulumi:"userData"`
	// Field Deprecated, use virtualizationModel instead. The type of the virtio interface.
	VirtioInterface *string `pulumi:"virtioInterface"`
	// The name of the virtual machine.
	VirtualMachineName *string `pulumi:"virtualMachineName"`
	// The type of the device model to use.
	VmDeviceModel *string `pulumi:"vmDeviceModel"`
	// The virtual machine image that is currently provisioned to the OS disk, using the full url and tag notation used to pull the image.
	VmImage                      string                      `pulumi:"vmImage"`
	VmImageRepositoryCredentials *ImageRepositoryCredentials `pulumi:"vmImageRepositoryCredentials"`
}

// The set of arguments for constructing a VirtualMachine resource.
type VirtualMachineArgs struct {
	// The name of the administrator to which the ssh public keys will be added into the authorized keys.
	AdminUsername pulumi.StringInput
	// Selects the boot method for the virtual machine.
	BootMethod                     pulumi.StringPtrInput
	CloudServicesNetworkAttachment NetworkAttachmentInput
	// The number of CPU cores in the virtual machine.
	CpuCores pulumi.Float64Input
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationInput
	// Field Deprecated, the value will be ignored if provided. The indicator of whether one of the specified CPU cores is isolated to run the emulator thread for this virtual machine.
	IsolateEmulatorThread pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The memory size of the virtual machine in GB.
	MemorySizeGB pulumi.Float64Input
	// The list of network attachments to the virtual machine.
	NetworkAttachments NetworkAttachmentArrayInput
	// The Base64 encoded cloud-init network data.
	NetworkData pulumi.StringPtrInput
	// The scheduling hints for the virtual machine.
	PlacementHints VirtualMachinePlacementHintArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The list of ssh public keys. Each key will be added to the virtual machine using the cloud-init ssh_authorized_keys mechanism for the adminUsername.
	SshPublicKeys  SshPublicKeyArrayInput
	StorageProfile StorageProfileInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The Base64 encoded cloud-init user data.
	UserData pulumi.StringPtrInput
	// Field Deprecated, use virtualizationModel instead. The type of the virtio interface.
	VirtioInterface pulumi.StringPtrInput
	// The name of the virtual machine.
	VirtualMachineName pulumi.StringPtrInput
	// The type of the device model to use.
	VmDeviceModel pulumi.StringPtrInput
	// The virtual machine image that is currently provisioned to the OS disk, using the full url and tag notation used to pull the image.
	VmImage                      pulumi.StringInput
	VmImageRepositoryCredentials ImageRepositoryCredentialsPtrInput
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

// The name of the administrator to which the ssh public keys will be added into the authorized keys.
func (o VirtualMachineOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.AdminUsername }).(pulumi.StringOutput)
}

// The resource ID of the bare metal machine the virtual machine has landed to.
func (o VirtualMachineOutput) BareMetalMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.BareMetalMachineId }).(pulumi.StringOutput)
}

// Selects the boot method for the virtual machine.
func (o VirtualMachineOutput) BootMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.BootMethod }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) CloudServicesNetworkAttachment() NetworkAttachmentResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) NetworkAttachmentResponseOutput { return v.CloudServicesNetworkAttachment }).(NetworkAttachmentResponseOutput)
}

// The resource ID of the cluster the virtual machine is created for.
func (o VirtualMachineOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The number of CPU cores in the virtual machine.
func (o VirtualMachineOutput) CpuCores() pulumi.Float64Output {
	return o.ApplyT(func(v *VirtualMachine) pulumi.Float64Output { return v.CpuCores }).(pulumi.Float64Output)
}

// The more detailed status of the virtual machine.
func (o VirtualMachineOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o VirtualMachineOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o VirtualMachineOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Field Deprecated, the value will be ignored if provided. The indicator of whether one of the specified CPU cores is isolated to run the emulator thread for this virtual machine.
func (o VirtualMachineOutput) IsolateEmulatorThread() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.IsolateEmulatorThread }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o VirtualMachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The memory size of the virtual machine in GB.
func (o VirtualMachineOutput) MemorySizeGB() pulumi.Float64Output {
	return o.ApplyT(func(v *VirtualMachine) pulumi.Float64Output { return v.MemorySizeGB }).(pulumi.Float64Output)
}

// The name of the resource
func (o VirtualMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of network attachments to the virtual machine.
func (o VirtualMachineOutput) NetworkAttachments() NetworkAttachmentResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) NetworkAttachmentResponseArrayOutput { return v.NetworkAttachments }).(NetworkAttachmentResponseArrayOutput)
}

// The Base64 encoded cloud-init network data.
func (o VirtualMachineOutput) NetworkData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.NetworkData }).(pulumi.StringPtrOutput)
}

// The scheduling hints for the virtual machine.
func (o VirtualMachineOutput) PlacementHints() VirtualMachinePlacementHintResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachinePlacementHintResponseArrayOutput { return v.PlacementHints }).(VirtualMachinePlacementHintResponseArrayOutput)
}

// The power state of the virtual machine.
func (o VirtualMachineOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.PowerState }).(pulumi.StringOutput)
}

// The provisioning state of the virtual machine.
func (o VirtualMachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The list of ssh public keys. Each key will be added to the virtual machine using the cloud-init ssh_authorized_keys mechanism for the adminUsername.
func (o VirtualMachineOutput) SshPublicKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) SshPublicKeyResponseArrayOutput { return v.SshPublicKeys }).(SshPublicKeyResponseArrayOutput)
}

func (o VirtualMachineOutput) StorageProfile() StorageProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) StorageProfileResponseOutput { return v.StorageProfile }).(StorageProfileResponseOutput)
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

// The Base64 encoded cloud-init user data.
func (o VirtualMachineOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.UserData }).(pulumi.StringPtrOutput)
}

// Field Deprecated, use virtualizationModel instead. The type of the virtio interface.
func (o VirtualMachineOutput) VirtioInterface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.VirtioInterface }).(pulumi.StringPtrOutput)
}

// The type of the device model to use.
func (o VirtualMachineOutput) VmDeviceModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.VmDeviceModel }).(pulumi.StringPtrOutput)
}

// The virtual machine image that is currently provisioned to the OS disk, using the full url and tag notation used to pull the image.
func (o VirtualMachineOutput) VmImage() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.VmImage }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) VmImageRepositoryCredentials() ImageRepositoryCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) ImageRepositoryCredentialsResponsePtrOutput {
		return v.VmImageRepositoryCredentials
	}).(ImageRepositoryCredentialsResponsePtrOutput)
}

// The resource IDs of volumes that are attached to the virtual machine.
func (o VirtualMachineOutput) Volumes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringArrayOutput { return v.Volumes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineOutput{})
}