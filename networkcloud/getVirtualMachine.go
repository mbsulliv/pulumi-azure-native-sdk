// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkcloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get properties of the provided virtual machine.
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2023-07-01.
func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:networkcloud:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualMachineArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual machine.
	VirtualMachineName string `pulumi:"virtualMachineName"`
}

type LookupVirtualMachineResult struct {
	// The name of the administrator to which the ssh public keys will be added into the authorized keys.
	AdminUsername string `pulumi:"adminUsername"`
	// The cluster availability zone containing this virtual machine.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The resource ID of the bare metal machine the virtual machine has landed to.
	BareMetalMachineId string `pulumi:"bareMetalMachineId"`
	// Selects the boot method for the virtual machine.
	BootMethod *string `pulumi:"bootMethod"`
	// The cloud service network that provides platform-level services for the virtual machine.
	CloudServicesNetworkAttachment NetworkAttachmentResponse `pulumi:"cloudServicesNetworkAttachment"`
	// The resource ID of the cluster the virtual machine is created for.
	ClusterId string `pulumi:"clusterId"`
	// The number of CPU cores in the virtual machine.
	CpuCores float64 `pulumi:"cpuCores"`
	// The more detailed status of the virtual machine.
	DetailedStatus string `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage string `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Field Deprecated, the value will be ignored if provided. The indicator of whether one of the specified CPU cores is isolated to run the emulator thread for this virtual machine.
	IsolateEmulatorThread *string `pulumi:"isolateEmulatorThread"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The memory size of the virtual machine in GB.
	MemorySizeGB float64 `pulumi:"memorySizeGB"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The list of network attachments to the virtual machine.
	NetworkAttachments []NetworkAttachmentResponse `pulumi:"networkAttachments"`
	// The Base64 encoded cloud-init network data.
	NetworkData *string `pulumi:"networkData"`
	// The scheduling hints for the virtual machine.
	PlacementHints []VirtualMachinePlacementHintResponse `pulumi:"placementHints"`
	// The power state of the virtual machine.
	PowerState string `pulumi:"powerState"`
	// The provisioning state of the virtual machine.
	ProvisioningState string `pulumi:"provisioningState"`
	// The list of ssh public keys. Each key will be added to the virtual machine using the cloud-init ssh_authorized_keys mechanism for the adminUsername.
	SshPublicKeys []SshPublicKeyResponse `pulumi:"sshPublicKeys"`
	// The storage profile that specifies size and other parameters about the disks related to the virtual machine.
	StorageProfile StorageProfileResponse `pulumi:"storageProfile"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The Base64 encoded cloud-init user data.
	UserData *string `pulumi:"userData"`
	// Field Deprecated, use virtualizationModel instead. The type of the virtio interface.
	VirtioInterface *string `pulumi:"virtioInterface"`
	// The type of the device model to use.
	VmDeviceModel *string `pulumi:"vmDeviceModel"`
	// The virtual machine image that is currently provisioned to the OS disk, using the full url and tag notation used to pull the image.
	VmImage string `pulumi:"vmImage"`
	// The credentials used to login to the image repository that has access to the specified image.
	VmImageRepositoryCredentials *ImageRepositoryCredentialsResponse `pulumi:"vmImageRepositoryCredentials"`
	// The resource IDs of volumes that are attached to the virtual machine.
	Volumes []string `pulumi:"volumes"`
}

// Defaults sets the appropriate defaults for LookupVirtualMachineResult
func (val *LookupVirtualMachineResult) Defaults() *LookupVirtualMachineResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.BootMethod == nil {
		bootMethod_ := "UEFI"
		tmp.BootMethod = &bootMethod_
	}
	if tmp.IsolateEmulatorThread == nil {
		isolateEmulatorThread_ := "True"
		tmp.IsolateEmulatorThread = &isolateEmulatorThread_
	}
	tmp.StorageProfile = *tmp.StorageProfile.Defaults()

	if tmp.VirtioInterface == nil {
		virtioInterface_ := "Modern"
		tmp.VirtioInterface = &virtioInterface_
	}
	if tmp.VmDeviceModel == nil {
		vmDeviceModel_ := "T2"
		tmp.VmDeviceModel = &vmDeviceModel_
	}
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
	// The name of the virtual machine.
	VirtualMachineName pulumi.StringInput `pulumi:"virtualMachineName"`
}

func (LookupVirtualMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineArgs)(nil)).Elem()
}

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

// The name of the administrator to which the ssh public keys will be added into the authorized keys.
func (o LookupVirtualMachineResultOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.AdminUsername }).(pulumi.StringOutput)
}

// The cluster availability zone containing this virtual machine.
func (o LookupVirtualMachineResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The resource ID of the bare metal machine the virtual machine has landed to.
func (o LookupVirtualMachineResultOutput) BareMetalMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.BareMetalMachineId }).(pulumi.StringOutput)
}

// Selects the boot method for the virtual machine.
func (o LookupVirtualMachineResultOutput) BootMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.BootMethod }).(pulumi.StringPtrOutput)
}

// The cloud service network that provides platform-level services for the virtual machine.
func (o LookupVirtualMachineResultOutput) CloudServicesNetworkAttachment() NetworkAttachmentResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) NetworkAttachmentResponse { return v.CloudServicesNetworkAttachment }).(NetworkAttachmentResponseOutput)
}

// The resource ID of the cluster the virtual machine is created for.
func (o LookupVirtualMachineResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The number of CPU cores in the virtual machine.
func (o LookupVirtualMachineResultOutput) CpuCores() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVirtualMachineResult) float64 { return v.CpuCores }).(pulumi.Float64Output)
}

// The more detailed status of the virtual machine.
func (o LookupVirtualMachineResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o LookupVirtualMachineResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o LookupVirtualMachineResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

// Field Deprecated, the value will be ignored if provided. The indicator of whether one of the specified CPU cores is isolated to run the emulator thread for this virtual machine.
func (o LookupVirtualMachineResultOutput) IsolateEmulatorThread() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.IsolateEmulatorThread }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupVirtualMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

// The memory size of the virtual machine in GB.
func (o LookupVirtualMachineResultOutput) MemorySizeGB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVirtualMachineResult) float64 { return v.MemorySizeGB }).(pulumi.Float64Output)
}

// The name of the resource
func (o LookupVirtualMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of network attachments to the virtual machine.
func (o LookupVirtualMachineResultOutput) NetworkAttachments() NetworkAttachmentResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []NetworkAttachmentResponse { return v.NetworkAttachments }).(NetworkAttachmentResponseArrayOutput)
}

// The Base64 encoded cloud-init network data.
func (o LookupVirtualMachineResultOutput) NetworkData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.NetworkData }).(pulumi.StringPtrOutput)
}

// The scheduling hints for the virtual machine.
func (o LookupVirtualMachineResultOutput) PlacementHints() VirtualMachinePlacementHintResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []VirtualMachinePlacementHintResponse { return v.PlacementHints }).(VirtualMachinePlacementHintResponseArrayOutput)
}

// The power state of the virtual machine.
func (o LookupVirtualMachineResultOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.PowerState }).(pulumi.StringOutput)
}

// The provisioning state of the virtual machine.
func (o LookupVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The list of ssh public keys. Each key will be added to the virtual machine using the cloud-init ssh_authorized_keys mechanism for the adminUsername.
func (o LookupVirtualMachineResultOutput) SshPublicKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []SshPublicKeyResponse { return v.SshPublicKeys }).(SshPublicKeyResponseArrayOutput)
}

// The storage profile that specifies size and other parameters about the disks related to the virtual machine.
func (o LookupVirtualMachineResultOutput) StorageProfile() StorageProfileResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponseOutput)
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

// The Base64 encoded cloud-init user data.
func (o LookupVirtualMachineResultOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.UserData }).(pulumi.StringPtrOutput)
}

// Field Deprecated, use virtualizationModel instead. The type of the virtio interface.
func (o LookupVirtualMachineResultOutput) VirtioInterface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.VirtioInterface }).(pulumi.StringPtrOutput)
}

// The type of the device model to use.
func (o LookupVirtualMachineResultOutput) VmDeviceModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.VmDeviceModel }).(pulumi.StringPtrOutput)
}

// The virtual machine image that is currently provisioned to the OS disk, using the full url and tag notation used to pull the image.
func (o LookupVirtualMachineResultOutput) VmImage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.VmImage }).(pulumi.StringOutput)
}

// The credentials used to login to the image repository that has access to the specified image.
func (o LookupVirtualMachineResultOutput) VmImageRepositoryCredentials() ImageRepositoryCredentialsResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *ImageRepositoryCredentialsResponse {
		return v.VmImageRepositoryCredentials
	}).(ImageRepositoryCredentialsResponsePtrOutput)
}

// The resource IDs of volumes that are attached to the virtual machine.
func (o LookupVirtualMachineResultOutput) Volumes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []string { return v.Volumes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineResultOutput{})
}
