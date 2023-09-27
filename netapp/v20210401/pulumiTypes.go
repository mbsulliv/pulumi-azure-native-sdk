// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToOutput(ctx context.Context) pulumix.Output[SystemDataResponse] {
	return pulumix.Output[SystemDataResponse]{
		OutputState: o.OutputState,
	}
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

// Volume details using the backup policy
type VolumeBackups struct {
	// Total count of backups for volume
	BackupsCount *int `pulumi:"backupsCount"`
	// Policy enabled
	PolicyEnabled *bool `pulumi:"policyEnabled"`
	// Volume name
	VolumeName *string `pulumi:"volumeName"`
}

// VolumeBackupsInput is an input type that accepts VolumeBackupsArgs and VolumeBackupsOutput values.
// You can construct a concrete instance of `VolumeBackupsInput` via:
//
//	VolumeBackupsArgs{...}
type VolumeBackupsInput interface {
	pulumi.Input

	ToVolumeBackupsOutput() VolumeBackupsOutput
	ToVolumeBackupsOutputWithContext(context.Context) VolumeBackupsOutput
}

// Volume details using the backup policy
type VolumeBackupsArgs struct {
	// Total count of backups for volume
	BackupsCount pulumi.IntPtrInput `pulumi:"backupsCount"`
	// Policy enabled
	PolicyEnabled pulumi.BoolPtrInput `pulumi:"policyEnabled"`
	// Volume name
	VolumeName pulumi.StringPtrInput `pulumi:"volumeName"`
}

func (VolumeBackupsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeBackups)(nil)).Elem()
}

func (i VolumeBackupsArgs) ToVolumeBackupsOutput() VolumeBackupsOutput {
	return i.ToVolumeBackupsOutputWithContext(context.Background())
}

func (i VolumeBackupsArgs) ToVolumeBackupsOutputWithContext(ctx context.Context) VolumeBackupsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeBackupsOutput)
}

func (i VolumeBackupsArgs) ToOutput(ctx context.Context) pulumix.Output[VolumeBackups] {
	return pulumix.Output[VolumeBackups]{
		OutputState: i.ToVolumeBackupsOutputWithContext(ctx).OutputState,
	}
}

// VolumeBackupsArrayInput is an input type that accepts VolumeBackupsArray and VolumeBackupsArrayOutput values.
// You can construct a concrete instance of `VolumeBackupsArrayInput` via:
//
//	VolumeBackupsArray{ VolumeBackupsArgs{...} }
type VolumeBackupsArrayInput interface {
	pulumi.Input

	ToVolumeBackupsArrayOutput() VolumeBackupsArrayOutput
	ToVolumeBackupsArrayOutputWithContext(context.Context) VolumeBackupsArrayOutput
}

type VolumeBackupsArray []VolumeBackupsInput

func (VolumeBackupsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeBackups)(nil)).Elem()
}

func (i VolumeBackupsArray) ToVolumeBackupsArrayOutput() VolumeBackupsArrayOutput {
	return i.ToVolumeBackupsArrayOutputWithContext(context.Background())
}

func (i VolumeBackupsArray) ToVolumeBackupsArrayOutputWithContext(ctx context.Context) VolumeBackupsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeBackupsArrayOutput)
}

func (i VolumeBackupsArray) ToOutput(ctx context.Context) pulumix.Output[[]VolumeBackups] {
	return pulumix.Output[[]VolumeBackups]{
		OutputState: i.ToVolumeBackupsArrayOutputWithContext(ctx).OutputState,
	}
}

// Volume details using the backup policy
type VolumeBackupsOutput struct{ *pulumi.OutputState }

func (VolumeBackupsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeBackups)(nil)).Elem()
}

func (o VolumeBackupsOutput) ToVolumeBackupsOutput() VolumeBackupsOutput {
	return o
}

func (o VolumeBackupsOutput) ToVolumeBackupsOutputWithContext(ctx context.Context) VolumeBackupsOutput {
	return o
}

func (o VolumeBackupsOutput) ToOutput(ctx context.Context) pulumix.Output[VolumeBackups] {
	return pulumix.Output[VolumeBackups]{
		OutputState: o.OutputState,
	}
}

// Total count of backups for volume
func (o VolumeBackupsOutput) BackupsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VolumeBackups) *int { return v.BackupsCount }).(pulumi.IntPtrOutput)
}

// Policy enabled
func (o VolumeBackupsOutput) PolicyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeBackups) *bool { return v.PolicyEnabled }).(pulumi.BoolPtrOutput)
}

// Volume name
func (o VolumeBackupsOutput) VolumeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeBackups) *string { return v.VolumeName }).(pulumi.StringPtrOutput)
}

type VolumeBackupsArrayOutput struct{ *pulumi.OutputState }

func (VolumeBackupsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeBackups)(nil)).Elem()
}

func (o VolumeBackupsArrayOutput) ToVolumeBackupsArrayOutput() VolumeBackupsArrayOutput {
	return o
}

func (o VolumeBackupsArrayOutput) ToVolumeBackupsArrayOutputWithContext(ctx context.Context) VolumeBackupsArrayOutput {
	return o
}

func (o VolumeBackupsArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]VolumeBackups] {
	return pulumix.Output[[]VolumeBackups]{
		OutputState: o.OutputState,
	}
}

func (o VolumeBackupsArrayOutput) Index(i pulumi.IntInput) VolumeBackupsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeBackups {
		return vs[0].([]VolumeBackups)[vs[1].(int)]
	}).(VolumeBackupsOutput)
}

// Volume details using the backup policy
type VolumeBackupsResponse struct {
	// Total count of backups for volume
	BackupsCount *int `pulumi:"backupsCount"`
	// Policy enabled
	PolicyEnabled *bool `pulumi:"policyEnabled"`
	// Volume name
	VolumeName *string `pulumi:"volumeName"`
}

// Volume details using the backup policy
type VolumeBackupsResponseOutput struct{ *pulumi.OutputState }

func (VolumeBackupsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeBackupsResponse)(nil)).Elem()
}

func (o VolumeBackupsResponseOutput) ToVolumeBackupsResponseOutput() VolumeBackupsResponseOutput {
	return o
}

func (o VolumeBackupsResponseOutput) ToVolumeBackupsResponseOutputWithContext(ctx context.Context) VolumeBackupsResponseOutput {
	return o
}

func (o VolumeBackupsResponseOutput) ToOutput(ctx context.Context) pulumix.Output[VolumeBackupsResponse] {
	return pulumix.Output[VolumeBackupsResponse]{
		OutputState: o.OutputState,
	}
}

// Total count of backups for volume
func (o VolumeBackupsResponseOutput) BackupsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VolumeBackupsResponse) *int { return v.BackupsCount }).(pulumi.IntPtrOutput)
}

// Policy enabled
func (o VolumeBackupsResponseOutput) PolicyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeBackupsResponse) *bool { return v.PolicyEnabled }).(pulumi.BoolPtrOutput)
}

// Volume name
func (o VolumeBackupsResponseOutput) VolumeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeBackupsResponse) *string { return v.VolumeName }).(pulumi.StringPtrOutput)
}

type VolumeBackupsResponseArrayOutput struct{ *pulumi.OutputState }

func (VolumeBackupsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeBackupsResponse)(nil)).Elem()
}

func (o VolumeBackupsResponseArrayOutput) ToVolumeBackupsResponseArrayOutput() VolumeBackupsResponseArrayOutput {
	return o
}

func (o VolumeBackupsResponseArrayOutput) ToVolumeBackupsResponseArrayOutputWithContext(ctx context.Context) VolumeBackupsResponseArrayOutput {
	return o
}

func (o VolumeBackupsResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]VolumeBackupsResponse] {
	return pulumix.Output[[]VolumeBackupsResponse]{
		OutputState: o.OutputState,
	}
}

func (o VolumeBackupsResponseArrayOutput) Index(i pulumi.IntInput) VolumeBackupsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeBackupsResponse {
		return vs[0].([]VolumeBackupsResponse)[vs[1].(int)]
	}).(VolumeBackupsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VolumeBackupsOutput{})
	pulumi.RegisterOutputType(VolumeBackupsArrayOutput{})
	pulumi.RegisterOutputType(VolumeBackupsResponseOutput{})
	pulumi.RegisterOutputType(VolumeBackupsResponseArrayOutput{})
}
