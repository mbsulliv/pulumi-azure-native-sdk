// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180915

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A virtual machine.
type VirtualMachine struct {
	pulumi.CustomResourceState

	// Indicates whether another user can take ownership of the virtual machine
	AllowClaim pulumi.BoolPtrOutput `pulumi:"allowClaim"`
	// The applicable schedule for the virtual machine.
	ApplicableSchedule ApplicableScheduleResponseOutput `pulumi:"applicableSchedule"`
	// The artifact deployment status for the virtual machine.
	ArtifactDeploymentStatus ArtifactDeploymentStatusPropertiesResponseOutput `pulumi:"artifactDeploymentStatus"`
	// The artifacts to be installed on the virtual machine.
	Artifacts ArtifactInstallPropertiesResponseArrayOutput `pulumi:"artifacts"`
	// The resource identifier (Microsoft.Compute) of the virtual machine.
	ComputeId pulumi.StringOutput `pulumi:"computeId"`
	// The compute virtual machine properties.
	ComputeVm ComputeVmPropertiesResponseOutput `pulumi:"computeVm"`
	// The email address of creator of the virtual machine.
	CreatedByUser pulumi.StringOutput `pulumi:"createdByUser"`
	// The object identifier of the creator of the virtual machine.
	CreatedByUserId pulumi.StringOutput `pulumi:"createdByUserId"`
	// The creation date of the virtual machine.
	CreatedDate pulumi.StringPtrOutput `pulumi:"createdDate"`
	// The custom image identifier of the virtual machine.
	CustomImageId pulumi.StringPtrOutput `pulumi:"customImageId"`
	// New or existing data disks to attach to the virtual machine after creation
	DataDiskParameters DataDiskPropertiesResponseArrayOutput `pulumi:"dataDiskParameters"`
	// Indicates whether the virtual machine is to be created without a public IP address.
	DisallowPublicIpAddress pulumi.BoolPtrOutput `pulumi:"disallowPublicIpAddress"`
	// The resource ID of the environment that contains this virtual machine, if any.
	EnvironmentId pulumi.StringPtrOutput `pulumi:"environmentId"`
	// The expiration date for VM.
	ExpirationDate pulumi.StringPtrOutput `pulumi:"expirationDate"`
	// The fully-qualified domain name of the virtual machine.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The Microsoft Azure Marketplace image reference of the virtual machine.
	GalleryImageReference GalleryImageReferenceResponsePtrOutput `pulumi:"galleryImageReference"`
	// Indicates whether this virtual machine uses an SSH key for authentication.
	IsAuthenticationWithSshKey pulumi.BoolPtrOutput `pulumi:"isAuthenticationWithSshKey"`
	// The lab subnet name of the virtual machine.
	LabSubnetName pulumi.StringPtrOutput `pulumi:"labSubnetName"`
	// The lab virtual network identifier of the virtual machine.
	LabVirtualNetworkId pulumi.StringPtrOutput `pulumi:"labVirtualNetworkId"`
	// Last known compute power state captured in DTL
	LastKnownPowerState pulumi.StringOutput `pulumi:"lastKnownPowerState"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network interface properties.
	NetworkInterface NetworkInterfacePropertiesResponsePtrOutput `pulumi:"networkInterface"`
	// The notes of the virtual machine.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The OS type of the virtual machine.
	OsType pulumi.StringOutput `pulumi:"osType"`
	// The object identifier of the owner of the virtual machine.
	OwnerObjectId pulumi.StringPtrOutput `pulumi:"ownerObjectId"`
	// The user principal name of the virtual machine owner.
	OwnerUserPrincipalName pulumi.StringPtrOutput `pulumi:"ownerUserPrincipalName"`
	// The password of the virtual machine administrator.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The id of the plan associated with the virtual machine image
	PlanId pulumi.StringPtrOutput `pulumi:"planId"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Virtual Machine schedules to be created
	ScheduleParameters ScheduleCreationParameterResponseArrayOutput `pulumi:"scheduleParameters"`
	// The size of the virtual machine.
	Size pulumi.StringPtrOutput `pulumi:"size"`
	// The SSH key of the virtual machine administrator.
	SshKey pulumi.StringPtrOutput `pulumi:"sshKey"`
	// Storage type to use for virtual machine (i.e. Standard, Premium).
	StorageType pulumi.StringPtrOutput `pulumi:"storageType"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringOutput `pulumi:"uniqueIdentifier"`
	// The user name of the virtual machine.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
	// Tells source of creation of lab virtual machine. Output property only.
	VirtualMachineCreationSource pulumi.StringOutput `pulumi:"virtualMachineCreationSource"`
}

// NewVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.AllowClaim == nil {
		args.AllowClaim = pulumi.BoolPtr(false)
	}
	if args.DisallowPublicIpAddress == nil {
		args.DisallowPublicIpAddress = pulumi.BoolPtr(false)
	}
	if args.OwnerObjectId == nil {
		args.OwnerObjectId = pulumi.StringPtr("dynamicValue")
	}
	if args.StorageType == nil {
		args.StorageType = pulumi.StringPtr("labStorageType")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:VirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualMachine
	err := ctx.RegisterResource("azure-native:devtestlab/v20180915:VirtualMachine", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:devtestlab/v20180915:VirtualMachine", name, id, state, &resource, opts...)
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
	// Indicates whether another user can take ownership of the virtual machine
	AllowClaim *bool `pulumi:"allowClaim"`
	// The artifacts to be installed on the virtual machine.
	Artifacts []ArtifactInstallProperties `pulumi:"artifacts"`
	// The creation date of the virtual machine.
	CreatedDate *string `pulumi:"createdDate"`
	// The custom image identifier of the virtual machine.
	CustomImageId *string `pulumi:"customImageId"`
	// New or existing data disks to attach to the virtual machine after creation
	DataDiskParameters []DataDiskProperties `pulumi:"dataDiskParameters"`
	// Indicates whether the virtual machine is to be created without a public IP address.
	DisallowPublicIpAddress *bool `pulumi:"disallowPublicIpAddress"`
	// The resource ID of the environment that contains this virtual machine, if any.
	EnvironmentId *string `pulumi:"environmentId"`
	// The expiration date for VM.
	ExpirationDate *string `pulumi:"expirationDate"`
	// The Microsoft Azure Marketplace image reference of the virtual machine.
	GalleryImageReference *GalleryImageReference `pulumi:"galleryImageReference"`
	// Indicates whether this virtual machine uses an SSH key for authentication.
	IsAuthenticationWithSshKey *bool `pulumi:"isAuthenticationWithSshKey"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The lab subnet name of the virtual machine.
	LabSubnetName *string `pulumi:"labSubnetName"`
	// The lab virtual network identifier of the virtual machine.
	LabVirtualNetworkId *string `pulumi:"labVirtualNetworkId"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the virtual machine.
	Name *string `pulumi:"name"`
	// The network interface properties.
	NetworkInterface *NetworkInterfaceProperties `pulumi:"networkInterface"`
	// The notes of the virtual machine.
	Notes *string `pulumi:"notes"`
	// The object identifier of the owner of the virtual machine.
	OwnerObjectId *string `pulumi:"ownerObjectId"`
	// The user principal name of the virtual machine owner.
	OwnerUserPrincipalName *string `pulumi:"ownerUserPrincipalName"`
	// The password of the virtual machine administrator.
	Password *string `pulumi:"password"`
	// The id of the plan associated with the virtual machine image
	PlanId *string `pulumi:"planId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Virtual Machine schedules to be created
	ScheduleParameters []ScheduleCreationParameter `pulumi:"scheduleParameters"`
	// The size of the virtual machine.
	Size *string `pulumi:"size"`
	// The SSH key of the virtual machine administrator.
	SshKey *string `pulumi:"sshKey"`
	// Storage type to use for virtual machine (i.e. Standard, Premium).
	StorageType *string `pulumi:"storageType"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The user name of the virtual machine.
	UserName *string `pulumi:"userName"`
}

// The set of arguments for constructing a VirtualMachine resource.
type VirtualMachineArgs struct {
	// Indicates whether another user can take ownership of the virtual machine
	AllowClaim pulumi.BoolPtrInput
	// The artifacts to be installed on the virtual machine.
	Artifacts ArtifactInstallPropertiesArrayInput
	// The creation date of the virtual machine.
	CreatedDate pulumi.StringPtrInput
	// The custom image identifier of the virtual machine.
	CustomImageId pulumi.StringPtrInput
	// New or existing data disks to attach to the virtual machine after creation
	DataDiskParameters DataDiskPropertiesArrayInput
	// Indicates whether the virtual machine is to be created without a public IP address.
	DisallowPublicIpAddress pulumi.BoolPtrInput
	// The resource ID of the environment that contains this virtual machine, if any.
	EnvironmentId pulumi.StringPtrInput
	// The expiration date for VM.
	ExpirationDate pulumi.StringPtrInput
	// The Microsoft Azure Marketplace image reference of the virtual machine.
	GalleryImageReference GalleryImageReferencePtrInput
	// Indicates whether this virtual machine uses an SSH key for authentication.
	IsAuthenticationWithSshKey pulumi.BoolPtrInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The lab subnet name of the virtual machine.
	LabSubnetName pulumi.StringPtrInput
	// The lab virtual network identifier of the virtual machine.
	LabVirtualNetworkId pulumi.StringPtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the virtual machine.
	Name pulumi.StringPtrInput
	// The network interface properties.
	NetworkInterface NetworkInterfacePropertiesPtrInput
	// The notes of the virtual machine.
	Notes pulumi.StringPtrInput
	// The object identifier of the owner of the virtual machine.
	OwnerObjectId pulumi.StringPtrInput
	// The user principal name of the virtual machine owner.
	OwnerUserPrincipalName pulumi.StringPtrInput
	// The password of the virtual machine administrator.
	Password pulumi.StringPtrInput
	// The id of the plan associated with the virtual machine image
	PlanId pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Virtual Machine schedules to be created
	ScheduleParameters ScheduleCreationParameterArrayInput
	// The size of the virtual machine.
	Size pulumi.StringPtrInput
	// The SSH key of the virtual machine administrator.
	SshKey pulumi.StringPtrInput
	// Storage type to use for virtual machine (i.e. Standard, Premium).
	StorageType pulumi.StringPtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The user name of the virtual machine.
	UserName pulumi.StringPtrInput
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

// Indicates whether another user can take ownership of the virtual machine
func (o VirtualMachineOutput) AllowClaim() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.BoolPtrOutput { return v.AllowClaim }).(pulumi.BoolPtrOutput)
}

// The applicable schedule for the virtual machine.
func (o VirtualMachineOutput) ApplicableSchedule() ApplicableScheduleResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) ApplicableScheduleResponseOutput { return v.ApplicableSchedule }).(ApplicableScheduleResponseOutput)
}

// The artifact deployment status for the virtual machine.
func (o VirtualMachineOutput) ArtifactDeploymentStatus() ArtifactDeploymentStatusPropertiesResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) ArtifactDeploymentStatusPropertiesResponseOutput {
		return v.ArtifactDeploymentStatus
	}).(ArtifactDeploymentStatusPropertiesResponseOutput)
}

// The artifacts to be installed on the virtual machine.
func (o VirtualMachineOutput) Artifacts() ArtifactInstallPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) ArtifactInstallPropertiesResponseArrayOutput { return v.Artifacts }).(ArtifactInstallPropertiesResponseArrayOutput)
}

// The resource identifier (Microsoft.Compute) of the virtual machine.
func (o VirtualMachineOutput) ComputeId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.ComputeId }).(pulumi.StringOutput)
}

// The compute virtual machine properties.
func (o VirtualMachineOutput) ComputeVm() ComputeVmPropertiesResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) ComputeVmPropertiesResponseOutput { return v.ComputeVm }).(ComputeVmPropertiesResponseOutput)
}

// The email address of creator of the virtual machine.
func (o VirtualMachineOutput) CreatedByUser() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.CreatedByUser }).(pulumi.StringOutput)
}

// The object identifier of the creator of the virtual machine.
func (o VirtualMachineOutput) CreatedByUserId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.CreatedByUserId }).(pulumi.StringOutput)
}

// The creation date of the virtual machine.
func (o VirtualMachineOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

// The custom image identifier of the virtual machine.
func (o VirtualMachineOutput) CustomImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.CustomImageId }).(pulumi.StringPtrOutput)
}

// New or existing data disks to attach to the virtual machine after creation
func (o VirtualMachineOutput) DataDiskParameters() DataDiskPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) DataDiskPropertiesResponseArrayOutput { return v.DataDiskParameters }).(DataDiskPropertiesResponseArrayOutput)
}

// Indicates whether the virtual machine is to be created without a public IP address.
func (o VirtualMachineOutput) DisallowPublicIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.BoolPtrOutput { return v.DisallowPublicIpAddress }).(pulumi.BoolPtrOutput)
}

// The resource ID of the environment that contains this virtual machine, if any.
func (o VirtualMachineOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

// The expiration date for VM.
func (o VirtualMachineOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

// The fully-qualified domain name of the virtual machine.
func (o VirtualMachineOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// The Microsoft Azure Marketplace image reference of the virtual machine.
func (o VirtualMachineOutput) GalleryImageReference() GalleryImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) GalleryImageReferenceResponsePtrOutput { return v.GalleryImageReference }).(GalleryImageReferenceResponsePtrOutput)
}

// Indicates whether this virtual machine uses an SSH key for authentication.
func (o VirtualMachineOutput) IsAuthenticationWithSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.BoolPtrOutput { return v.IsAuthenticationWithSshKey }).(pulumi.BoolPtrOutput)
}

// The lab subnet name of the virtual machine.
func (o VirtualMachineOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

// The lab virtual network identifier of the virtual machine.
func (o VirtualMachineOutput) LabVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.LabVirtualNetworkId }).(pulumi.StringPtrOutput)
}

// Last known compute power state captured in DTL
func (o VirtualMachineOutput) LastKnownPowerState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.LastKnownPowerState }).(pulumi.StringOutput)
}

// The location of the resource.
func (o VirtualMachineOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o VirtualMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network interface properties.
func (o VirtualMachineOutput) NetworkInterface() NetworkInterfacePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) NetworkInterfacePropertiesResponsePtrOutput { return v.NetworkInterface }).(NetworkInterfacePropertiesResponsePtrOutput)
}

// The notes of the virtual machine.
func (o VirtualMachineOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// The OS type of the virtual machine.
func (o VirtualMachineOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.OsType }).(pulumi.StringOutput)
}

// The object identifier of the owner of the virtual machine.
func (o VirtualMachineOutput) OwnerObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.OwnerObjectId }).(pulumi.StringPtrOutput)
}

// The user principal name of the virtual machine owner.
func (o VirtualMachineOutput) OwnerUserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.OwnerUserPrincipalName }).(pulumi.StringPtrOutput)
}

// The password of the virtual machine administrator.
func (o VirtualMachineOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The id of the plan associated with the virtual machine image
func (o VirtualMachineOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.PlanId }).(pulumi.StringPtrOutput)
}

// The provisioning status of the resource.
func (o VirtualMachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Virtual Machine schedules to be created
func (o VirtualMachineOutput) ScheduleParameters() ScheduleCreationParameterResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) ScheduleCreationParameterResponseArrayOutput { return v.ScheduleParameters }).(ScheduleCreationParameterResponseArrayOutput)
}

// The size of the virtual machine.
func (o VirtualMachineOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Size }).(pulumi.StringPtrOutput)
}

// The SSH key of the virtual machine administrator.
func (o VirtualMachineOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.SshKey }).(pulumi.StringPtrOutput)
}

// Storage type to use for virtual machine (i.e. Standard, Premium).
func (o VirtualMachineOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.StorageType }).(pulumi.StringPtrOutput)
}

// The tags of the resource.
func (o VirtualMachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o VirtualMachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o VirtualMachineOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

// The user name of the virtual machine.
func (o VirtualMachineOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

// Tells source of creation of lab virtual machine. Output property only.
func (o VirtualMachineOutput) VirtualMachineCreationSource() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.VirtualMachineCreationSource }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineOutput{})
}
