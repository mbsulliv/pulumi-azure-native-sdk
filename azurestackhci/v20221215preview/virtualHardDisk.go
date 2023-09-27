// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221215preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The virtual hard disk resource definition.
type VirtualHardDisk struct {
	pulumi.CustomResourceState

	BlockSizeBytes pulumi.IntPtrOutput `pulumi:"blockSizeBytes"`
	// Storage ContainerID of the storage container to be used for VHD
	ContainerId pulumi.StringPtrOutput `pulumi:"containerId"`
	// The format of the actual VHD file [vhd, vhdx]
	DiskFileFormat pulumi.StringPtrOutput `pulumi:"diskFileFormat"`
	// Size of the disk in GB
	DiskSizeGB pulumi.Float64PtrOutput `pulumi:"diskSizeGB"`
	// Boolean for enabling dynamic sizing on the virtual hard disk
	Dynamic pulumi.BoolPtrOutput `pulumi:"dynamic"`
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine [V1, V2]
	HyperVGeneration pulumi.StringPtrOutput `pulumi:"hyperVGeneration"`
	// The geo-location where the resource lives
	Location           pulumi.StringOutput `pulumi:"location"`
	LogicalSectorBytes pulumi.IntPtrOutput `pulumi:"logicalSectorBytes"`
	// The name of the resource
	Name                pulumi.StringOutput `pulumi:"name"`
	PhysicalSectorBytes pulumi.IntPtrOutput `pulumi:"physicalSectorBytes"`
	// Provisioning state of the virtual hard disk.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The observed state of virtual hard disks
	Status VirtualHardDiskStatusResponseOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualHardDisk registers a new resource with the given unique name, arguments, and options.
func NewVirtualHardDisk(ctx *pulumi.Context,
	name string, args *VirtualHardDiskArgs, opts ...pulumi.ResourceOption) (*VirtualHardDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci:VirtualHardDisk"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210701preview:VirtualHardDisk"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:VirtualHardDisk"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230701preview:VirtualHardDisk"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualHardDisk
	err := ctx.RegisterResource("azure-native:azurestackhci/v20221215preview:VirtualHardDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHardDisk gets an existing VirtualHardDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHardDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHardDiskState, opts ...pulumi.ResourceOption) (*VirtualHardDisk, error) {
	var resource VirtualHardDisk
	err := ctx.ReadResource("azure-native:azurestackhci/v20221215preview:VirtualHardDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHardDisk resources.
type virtualHardDiskState struct {
}

type VirtualHardDiskState struct {
}

func (VirtualHardDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHardDiskState)(nil)).Elem()
}

type virtualHardDiskArgs struct {
	BlockSizeBytes *int `pulumi:"blockSizeBytes"`
	// Storage ContainerID of the storage container to be used for VHD
	ContainerId *string `pulumi:"containerId"`
	// The format of the actual VHD file [vhd, vhdx]
	DiskFileFormat *string `pulumi:"diskFileFormat"`
	// Size of the disk in GB
	DiskSizeGB *float64 `pulumi:"diskSizeGB"`
	// Boolean for enabling dynamic sizing on the virtual hard disk
	Dynamic *bool `pulumi:"dynamic"`
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine [V1, V2]
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// The geo-location where the resource lives
	Location            *string `pulumi:"location"`
	LogicalSectorBytes  *int    `pulumi:"logicalSectorBytes"`
	PhysicalSectorBytes *int    `pulumi:"physicalSectorBytes"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Name of the virtual hard disk
	VirtualHardDiskName *string `pulumi:"virtualHardDiskName"`
}

// The set of arguments for constructing a VirtualHardDisk resource.
type VirtualHardDiskArgs struct {
	BlockSizeBytes pulumi.IntPtrInput
	// Storage ContainerID of the storage container to be used for VHD
	ContainerId pulumi.StringPtrInput
	// The format of the actual VHD file [vhd, vhdx]
	DiskFileFormat pulumi.StringPtrInput
	// Size of the disk in GB
	DiskSizeGB pulumi.Float64PtrInput
	// Boolean for enabling dynamic sizing on the virtual hard disk
	Dynamic pulumi.BoolPtrInput
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationPtrInput
	// The hypervisor generation of the Virtual Machine [V1, V2]
	HyperVGeneration pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location            pulumi.StringPtrInput
	LogicalSectorBytes  pulumi.IntPtrInput
	PhysicalSectorBytes pulumi.IntPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Name of the virtual hard disk
	VirtualHardDiskName pulumi.StringPtrInput
}

func (VirtualHardDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHardDiskArgs)(nil)).Elem()
}

type VirtualHardDiskInput interface {
	pulumi.Input

	ToVirtualHardDiskOutput() VirtualHardDiskOutput
	ToVirtualHardDiskOutputWithContext(ctx context.Context) VirtualHardDiskOutput
}

func (*VirtualHardDisk) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDisk)(nil)).Elem()
}

func (i *VirtualHardDisk) ToVirtualHardDiskOutput() VirtualHardDiskOutput {
	return i.ToVirtualHardDiskOutputWithContext(context.Background())
}

func (i *VirtualHardDisk) ToVirtualHardDiskOutputWithContext(ctx context.Context) VirtualHardDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHardDiskOutput)
}

func (i *VirtualHardDisk) ToOutput(ctx context.Context) pulumix.Output[*VirtualHardDisk] {
	return pulumix.Output[*VirtualHardDisk]{
		OutputState: i.ToVirtualHardDiskOutputWithContext(ctx).OutputState,
	}
}

type VirtualHardDiskOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDisk)(nil)).Elem()
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskOutput() VirtualHardDiskOutput {
	return o
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskOutputWithContext(ctx context.Context) VirtualHardDiskOutput {
	return o
}

func (o VirtualHardDiskOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualHardDisk] {
	return pulumix.Output[*VirtualHardDisk]{
		OutputState: o.OutputState,
	}
}

func (o VirtualHardDiskOutput) BlockSizeBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualHardDisk) pulumi.IntPtrOutput { return v.BlockSizeBytes }).(pulumi.IntPtrOutput)
}

// Storage ContainerID of the storage container to be used for VHD
func (o VirtualHardDiskOutput) ContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHardDisk) pulumi.StringPtrOutput { return v.ContainerId }).(pulumi.StringPtrOutput)
}

// The format of the actual VHD file [vhd, vhdx]
func (o VirtualHardDiskOutput) DiskFileFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHardDisk) pulumi.StringPtrOutput { return v.DiskFileFormat }).(pulumi.StringPtrOutput)
}

// Size of the disk in GB
func (o VirtualHardDiskOutput) DiskSizeGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualHardDisk) pulumi.Float64PtrOutput { return v.DiskSizeGB }).(pulumi.Float64PtrOutput)
}

// Boolean for enabling dynamic sizing on the virtual hard disk
func (o VirtualHardDiskOutput) Dynamic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualHardDisk) pulumi.BoolPtrOutput { return v.Dynamic }).(pulumi.BoolPtrOutput)
}

// The extendedLocation of the resource.
func (o VirtualHardDiskOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHardDisk) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The hypervisor generation of the Virtual Machine [V1, V2]
func (o VirtualHardDiskOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHardDisk) pulumi.StringPtrOutput { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o VirtualHardDiskOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHardDisk) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualHardDiskOutput) LogicalSectorBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualHardDisk) pulumi.IntPtrOutput { return v.LogicalSectorBytes }).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o VirtualHardDiskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHardDisk) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualHardDiskOutput) PhysicalSectorBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualHardDisk) pulumi.IntPtrOutput { return v.PhysicalSectorBytes }).(pulumi.IntPtrOutput)
}

// Provisioning state of the virtual hard disk.
func (o VirtualHardDiskOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHardDisk) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The observed state of virtual hard disks
func (o VirtualHardDiskOutput) Status() VirtualHardDiskStatusResponseOutput {
	return o.ApplyT(func(v *VirtualHardDisk) VirtualHardDiskStatusResponseOutput { return v.Status }).(VirtualHardDiskStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o VirtualHardDiskOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualHardDisk) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o VirtualHardDiskOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualHardDisk) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o VirtualHardDiskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHardDisk) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualHardDiskOutput{})
}
