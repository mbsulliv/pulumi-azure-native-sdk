// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package netapp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get details of the specified subvolume
// Azure REST API version: 2022-11-01.
func GetSubvolumeMetadata(ctx *pulumi.Context, args *GetSubvolumeMetadataArgs, opts ...pulumi.InvokeOption) (*GetSubvolumeMetadataResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetSubvolumeMetadataResult
	err := ctx.Invoke("azure-native:netapp:getSubvolumeMetadata", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSubvolumeMetadataArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the subvolume.
	SubvolumeName string `pulumi:"subvolumeName"`
	// The name of the volume
	VolumeName string `pulumi:"volumeName"`
}

// Result of the post subvolume and action is to get metadata of the subvolume.
type GetSubvolumeMetadataResult struct {
	// Most recent access time and date
	AccessedTimeStamp *string `pulumi:"accessedTimeStamp"`
	// Bytes used
	BytesUsed *float64 `pulumi:"bytesUsed"`
	// Most recent change time and date
	ChangedTimeStamp *string `pulumi:"changedTimeStamp"`
	// Creation time and date
	CreationTimeStamp *string `pulumi:"creationTimeStamp"`
	// Resource Id
	Id string `pulumi:"id"`
	// Most recent modification time and date
	ModifiedTimeStamp *string `pulumi:"modifiedTimeStamp"`
	// Resource name
	Name string `pulumi:"name"`
	// Path to the parent subvolume
	ParentPath *string `pulumi:"parentPath"`
	// Path to the subvolume
	Path *string `pulumi:"path"`
	// Permissions of the subvolume
	Permissions *string `pulumi:"permissions"`
	// Azure lifecycle management
	ProvisioningState *string `pulumi:"provisioningState"`
	// Size of subvolume
	Size *float64 `pulumi:"size"`
	// Resource type
	Type string `pulumi:"type"`
}

func GetSubvolumeMetadataOutput(ctx *pulumi.Context, args GetSubvolumeMetadataOutputArgs, opts ...pulumi.InvokeOption) GetSubvolumeMetadataResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSubvolumeMetadataResult, error) {
			args := v.(GetSubvolumeMetadataArgs)
			r, err := GetSubvolumeMetadata(ctx, &args, opts...)
			var s GetSubvolumeMetadataResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSubvolumeMetadataResultOutput)
}

type GetSubvolumeMetadataOutputArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the subvolume.
	SubvolumeName pulumi.StringInput `pulumi:"subvolumeName"`
	// The name of the volume
	VolumeName pulumi.StringInput `pulumi:"volumeName"`
}

func (GetSubvolumeMetadataOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubvolumeMetadataArgs)(nil)).Elem()
}

// Result of the post subvolume and action is to get metadata of the subvolume.
type GetSubvolumeMetadataResultOutput struct{ *pulumi.OutputState }

func (GetSubvolumeMetadataResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubvolumeMetadataResult)(nil)).Elem()
}

func (o GetSubvolumeMetadataResultOutput) ToGetSubvolumeMetadataResultOutput() GetSubvolumeMetadataResultOutput {
	return o
}

func (o GetSubvolumeMetadataResultOutput) ToGetSubvolumeMetadataResultOutputWithContext(ctx context.Context) GetSubvolumeMetadataResultOutput {
	return o
}

func (o GetSubvolumeMetadataResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetSubvolumeMetadataResult] {
	return pulumix.Output[GetSubvolumeMetadataResult]{
		OutputState: o.OutputState,
	}
}

// Most recent access time and date
func (o GetSubvolumeMetadataResultOutput) AccessedTimeStamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.AccessedTimeStamp }).(pulumi.StringPtrOutput)
}

// Bytes used
func (o GetSubvolumeMetadataResultOutput) BytesUsed() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *float64 { return v.BytesUsed }).(pulumi.Float64PtrOutput)
}

// Most recent change time and date
func (o GetSubvolumeMetadataResultOutput) ChangedTimeStamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.ChangedTimeStamp }).(pulumi.StringPtrOutput)
}

// Creation time and date
func (o GetSubvolumeMetadataResultOutput) CreationTimeStamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.CreationTimeStamp }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o GetSubvolumeMetadataResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) string { return v.Id }).(pulumi.StringOutput)
}

// Most recent modification time and date
func (o GetSubvolumeMetadataResultOutput) ModifiedTimeStamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.ModifiedTimeStamp }).(pulumi.StringPtrOutput)
}

// Resource name
func (o GetSubvolumeMetadataResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) string { return v.Name }).(pulumi.StringOutput)
}

// Path to the parent subvolume
func (o GetSubvolumeMetadataResultOutput) ParentPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.ParentPath }).(pulumi.StringPtrOutput)
}

// Path to the subvolume
func (o GetSubvolumeMetadataResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// Permissions of the subvolume
func (o GetSubvolumeMetadataResultOutput) Permissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.Permissions }).(pulumi.StringPtrOutput)
}

// Azure lifecycle management
func (o GetSubvolumeMetadataResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Size of subvolume
func (o GetSubvolumeMetadataResultOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *float64 { return v.Size }).(pulumi.Float64PtrOutput)
}

// Resource type
func (o GetSubvolumeMetadataResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSubvolumeMetadataResultOutput{})
}
