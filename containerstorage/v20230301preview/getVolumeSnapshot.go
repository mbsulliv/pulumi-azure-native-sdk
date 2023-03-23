// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a VolumeSnapshot
func LookupVolumeSnapshot(ctx *pulumi.Context, args *LookupVolumeSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupVolumeSnapshotResult, error) {
	var rv LookupVolumeSnapshotResult
	err := ctx.Invoke("azure-native:containerstorage/v20230301preview:getVolumeSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeSnapshotArgs struct {
	// Pool Object
	PoolName string `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Volume Snapshot Resource
	VolumeSnapshotName string `pulumi:"volumeSnapshotName"`
}

// Concrete proxy resource types can be created by aliasing this type using a specific property type.
type LookupVolumeSnapshotResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// List of string mount options
	MountOptions []string `pulumi:"mountOptions"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// Reclaim Policy, Delete or Retain
	ReclaimPolicy string `pulumi:"reclaimPolicy"`
	// Reference to the source volume
	Source string `pulumi:"source"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Indicates how the volumes created from the snapshot should be attached
	VolumeMode string `pulumi:"volumeMode"`
}

func LookupVolumeSnapshotOutput(ctx *pulumi.Context, args LookupVolumeSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeSnapshotResult, error) {
			args := v.(LookupVolumeSnapshotArgs)
			r, err := LookupVolumeSnapshot(ctx, &args, opts...)
			var s LookupVolumeSnapshotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeSnapshotResultOutput)
}

type LookupVolumeSnapshotOutputArgs struct {
	// Pool Object
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Volume Snapshot Resource
	VolumeSnapshotName pulumi.StringInput `pulumi:"volumeSnapshotName"`
}

func (LookupVolumeSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeSnapshotArgs)(nil)).Elem()
}

// Concrete proxy resource types can be created by aliasing this type using a specific property type.
type LookupVolumeSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeSnapshotResult)(nil)).Elem()
}

func (o LookupVolumeSnapshotResultOutput) ToLookupVolumeSnapshotResultOutput() LookupVolumeSnapshotResultOutput {
	return o
}

func (o LookupVolumeSnapshotResultOutput) ToLookupVolumeSnapshotResultOutputWithContext(ctx context.Context) LookupVolumeSnapshotResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupVolumeSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of string mount options
func (o LookupVolumeSnapshotResultOutput) MountOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) []string { return v.MountOptions }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o LookupVolumeSnapshotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupVolumeSnapshotResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Reclaim Policy, Delete or Retain
func (o LookupVolumeSnapshotResultOutput) ReclaimPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) string { return v.ReclaimPolicy }).(pulumi.StringOutput)
}

// Reference to the source volume
func (o LookupVolumeSnapshotResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) string { return v.Source }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVolumeSnapshotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVolumeSnapshotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) string { return v.Type }).(pulumi.StringOutput)
}

// Indicates how the volumes created from the snapshot should be attached
func (o LookupVolumeSnapshotResultOutput) VolumeMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) string { return v.VolumeMode }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeSnapshotResultOutput{})
}
