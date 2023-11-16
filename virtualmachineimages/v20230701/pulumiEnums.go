// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// If there is a validation error and this field is set to 'cleanup', the build VM and associated network resources will be cleaned up. This is the default behavior. If there is a validation error and this field is set to 'abort', the build VM will be preserved.
type OnBuildError string

const (
	OnBuildErrorCleanup = OnBuildError("cleanup")
	OnBuildErrorAbort   = OnBuildError("abort")
)

// The type of identity used for the image template. The type 'None' will remove any identities from the image template.
type ResourceIdentityType string

const (
	ResourceIdentityTypeUserAssigned = ResourceIdentityType("UserAssigned")
	ResourceIdentityTypeNone         = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ResourceIdentityTypeInput is an input type that accepts ResourceIdentityTypeArgs and ResourceIdentityTypeOutput values.
// You can construct a concrete instance of `ResourceIdentityTypeInput` via:
//
//	ResourceIdentityTypeArgs{...}
type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToOutput(ctx context.Context) pulumix.Output[*ResourceIdentityType] {
	return pulumix.Output[*ResourceIdentityType]{
		OutputState: in.ToResourceIdentityTypePtrOutputWithContext(ctx).OutputState,
	}
}

// Specifies the storage account type to be used to store the image in this region. Omit to use the default (Standard_LRS).
type SharedImageStorageAccountType string

const (
	SharedImageStorageAccountType_Standard_LRS = SharedImageStorageAccountType("Standard_LRS")
	SharedImageStorageAccountType_Standard_ZRS = SharedImageStorageAccountType("Standard_ZRS")
	SharedImageStorageAccountType_Premium_LRS  = SharedImageStorageAccountType("Premium_LRS")
)

// Enabling this field will improve VM boot time by optimizing the final customized image output.
type VMBootOptimizationState string

const (
	VMBootOptimizationStateEnabled  = VMBootOptimizationState("Enabled")
	VMBootOptimizationStateDisabled = VMBootOptimizationState("Disabled")
)

func (VMBootOptimizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*VMBootOptimizationState)(nil)).Elem()
}

func (e VMBootOptimizationState) ToVMBootOptimizationStateOutput() VMBootOptimizationStateOutput {
	return pulumi.ToOutput(e).(VMBootOptimizationStateOutput)
}

func (e VMBootOptimizationState) ToVMBootOptimizationStateOutputWithContext(ctx context.Context) VMBootOptimizationStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VMBootOptimizationStateOutput)
}

func (e VMBootOptimizationState) ToVMBootOptimizationStatePtrOutput() VMBootOptimizationStatePtrOutput {
	return e.ToVMBootOptimizationStatePtrOutputWithContext(context.Background())
}

func (e VMBootOptimizationState) ToVMBootOptimizationStatePtrOutputWithContext(ctx context.Context) VMBootOptimizationStatePtrOutput {
	return VMBootOptimizationState(e).ToVMBootOptimizationStateOutputWithContext(ctx).ToVMBootOptimizationStatePtrOutputWithContext(ctx)
}

func (e VMBootOptimizationState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VMBootOptimizationState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VMBootOptimizationState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VMBootOptimizationState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VMBootOptimizationStateOutput struct{ *pulumi.OutputState }

func (VMBootOptimizationStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMBootOptimizationState)(nil)).Elem()
}

func (o VMBootOptimizationStateOutput) ToVMBootOptimizationStateOutput() VMBootOptimizationStateOutput {
	return o
}

func (o VMBootOptimizationStateOutput) ToVMBootOptimizationStateOutputWithContext(ctx context.Context) VMBootOptimizationStateOutput {
	return o
}

func (o VMBootOptimizationStateOutput) ToVMBootOptimizationStatePtrOutput() VMBootOptimizationStatePtrOutput {
	return o.ToVMBootOptimizationStatePtrOutputWithContext(context.Background())
}

func (o VMBootOptimizationStateOutput) ToVMBootOptimizationStatePtrOutputWithContext(ctx context.Context) VMBootOptimizationStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VMBootOptimizationState) *VMBootOptimizationState {
		return &v
	}).(VMBootOptimizationStatePtrOutput)
}

func (o VMBootOptimizationStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VMBootOptimizationStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VMBootOptimizationState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VMBootOptimizationStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VMBootOptimizationStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VMBootOptimizationState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VMBootOptimizationStatePtrOutput struct{ *pulumi.OutputState }

func (VMBootOptimizationStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMBootOptimizationState)(nil)).Elem()
}

func (o VMBootOptimizationStatePtrOutput) ToVMBootOptimizationStatePtrOutput() VMBootOptimizationStatePtrOutput {
	return o
}

func (o VMBootOptimizationStatePtrOutput) ToVMBootOptimizationStatePtrOutputWithContext(ctx context.Context) VMBootOptimizationStatePtrOutput {
	return o
}

func (o VMBootOptimizationStatePtrOutput) Elem() VMBootOptimizationStateOutput {
	return o.ApplyT(func(v *VMBootOptimizationState) VMBootOptimizationState {
		if v != nil {
			return *v
		}
		var ret VMBootOptimizationState
		return ret
	}).(VMBootOptimizationStateOutput)
}

func (o VMBootOptimizationStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VMBootOptimizationStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VMBootOptimizationState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// VMBootOptimizationStateInput is an input type that accepts VMBootOptimizationStateArgs and VMBootOptimizationStateOutput values.
// You can construct a concrete instance of `VMBootOptimizationStateInput` via:
//
//	VMBootOptimizationStateArgs{...}
type VMBootOptimizationStateInput interface {
	pulumi.Input

	ToVMBootOptimizationStateOutput() VMBootOptimizationStateOutput
	ToVMBootOptimizationStateOutputWithContext(context.Context) VMBootOptimizationStateOutput
}

var vmbootOptimizationStatePtrType = reflect.TypeOf((**VMBootOptimizationState)(nil)).Elem()

type VMBootOptimizationStatePtrInput interface {
	pulumi.Input

	ToVMBootOptimizationStatePtrOutput() VMBootOptimizationStatePtrOutput
	ToVMBootOptimizationStatePtrOutputWithContext(context.Context) VMBootOptimizationStatePtrOutput
}

type vmbootOptimizationStatePtr string

func VMBootOptimizationStatePtr(v string) VMBootOptimizationStatePtrInput {
	return (*vmbootOptimizationStatePtr)(&v)
}

func (*vmbootOptimizationStatePtr) ElementType() reflect.Type {
	return vmbootOptimizationStatePtrType
}

func (in *vmbootOptimizationStatePtr) ToVMBootOptimizationStatePtrOutput() VMBootOptimizationStatePtrOutput {
	return pulumi.ToOutput(in).(VMBootOptimizationStatePtrOutput)
}

func (in *vmbootOptimizationStatePtr) ToVMBootOptimizationStatePtrOutputWithContext(ctx context.Context) VMBootOptimizationStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VMBootOptimizationStatePtrOutput)
}

func (in *vmbootOptimizationStatePtr) ToOutput(ctx context.Context) pulumix.Output[*VMBootOptimizationState] {
	return pulumix.Output[*VMBootOptimizationState]{
		OutputState: in.ToVMBootOptimizationStatePtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(VMBootOptimizationStateOutput{})
	pulumi.RegisterOutputType(VMBootOptimizationStatePtrOutput{})
}
