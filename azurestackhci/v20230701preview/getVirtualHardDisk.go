// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a virtual hard disk
func LookupVirtualHardDisk(ctx *pulumi.Context, args *LookupVirtualHardDiskArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHardDiskResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualHardDiskResult
	err := ctx.Invoke("azure-native:azurestackhci/v20230701preview:getVirtualHardDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHardDiskArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the virtual hard disk
	VirtualHardDiskName string `pulumi:"virtualHardDiskName"`
}

// The virtual hard disk resource definition.
type LookupVirtualHardDiskResult struct {
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
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine [V1, V2]
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location           string `pulumi:"location"`
	LogicalSectorBytes *int   `pulumi:"logicalSectorBytes"`
	// The name of the resource
	Name                string `pulumi:"name"`
	PhysicalSectorBytes *int   `pulumi:"physicalSectorBytes"`
	// Provisioning state of the virtual hard disk.
	ProvisioningState string `pulumi:"provisioningState"`
	// The observed state of virtual hard disks
	Status VirtualHardDiskStatusResponse `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupVirtualHardDiskOutput(ctx *pulumi.Context, args LookupVirtualHardDiskOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualHardDiskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualHardDiskResult, error) {
			args := v.(LookupVirtualHardDiskArgs)
			r, err := LookupVirtualHardDisk(ctx, &args, opts...)
			var s LookupVirtualHardDiskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualHardDiskResultOutput)
}

type LookupVirtualHardDiskOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the virtual hard disk
	VirtualHardDiskName pulumi.StringInput `pulumi:"virtualHardDiskName"`
}

func (LookupVirtualHardDiskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHardDiskArgs)(nil)).Elem()
}

// The virtual hard disk resource definition.
type LookupVirtualHardDiskResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualHardDiskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHardDiskResult)(nil)).Elem()
}

func (o LookupVirtualHardDiskResultOutput) ToLookupVirtualHardDiskResultOutput() LookupVirtualHardDiskResultOutput {
	return o
}

func (o LookupVirtualHardDiskResultOutput) ToLookupVirtualHardDiskResultOutputWithContext(ctx context.Context) LookupVirtualHardDiskResultOutput {
	return o
}

func (o LookupVirtualHardDiskResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualHardDiskResult] {
	return pulumix.Output[LookupVirtualHardDiskResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupVirtualHardDiskResultOutput) BlockSizeBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) *int { return v.BlockSizeBytes }).(pulumi.IntPtrOutput)
}

// Storage ContainerID of the storage container to be used for VHD
func (o LookupVirtualHardDiskResultOutput) ContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) *string { return v.ContainerId }).(pulumi.StringPtrOutput)
}

// The format of the actual VHD file [vhd, vhdx]
func (o LookupVirtualHardDiskResultOutput) DiskFileFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) *string { return v.DiskFileFormat }).(pulumi.StringPtrOutput)
}

// Size of the disk in GB
func (o LookupVirtualHardDiskResultOutput) DiskSizeGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) *float64 { return v.DiskSizeGB }).(pulumi.Float64PtrOutput)
}

// Boolean for enabling dynamic sizing on the virtual hard disk
func (o LookupVirtualHardDiskResultOutput) Dynamic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) *bool { return v.Dynamic }).(pulumi.BoolPtrOutput)
}

// The extendedLocation of the resource.
func (o LookupVirtualHardDiskResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The hypervisor generation of the Virtual Machine [V1, V2]
func (o LookupVirtualHardDiskResultOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) *string { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupVirtualHardDiskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupVirtualHardDiskResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualHardDiskResultOutput) LogicalSectorBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) *int { return v.LogicalSectorBytes }).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o LookupVirtualHardDiskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualHardDiskResultOutput) PhysicalSectorBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) *int { return v.PhysicalSectorBytes }).(pulumi.IntPtrOutput)
}

// Provisioning state of the virtual hard disk.
func (o LookupVirtualHardDiskResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The observed state of virtual hard disks
func (o LookupVirtualHardDiskResultOutput) Status() VirtualHardDiskStatusResponseOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) VirtualHardDiskStatusResponse { return v.Status }).(VirtualHardDiskStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVirtualHardDiskResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupVirtualHardDiskResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVirtualHardDiskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHardDiskResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualHardDiskResultOutput{})
}
