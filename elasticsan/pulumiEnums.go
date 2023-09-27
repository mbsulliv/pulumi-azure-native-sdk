// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsan

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The action of virtual network rule.
type Action string

const (
	ActionAllow = Action("Allow")
)

func (Action) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (e Action) ToActionOutput() ActionOutput {
	return pulumi.ToOutput(e).(ActionOutput)
}

func (e Action) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ActionOutput)
}

func (e Action) ToActionPtrOutput() ActionPtrOutput {
	return e.ToActionPtrOutputWithContext(context.Background())
}

func (e Action) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return Action(e).ToActionOutputWithContext(ctx).ToActionPtrOutputWithContext(ctx)
}

func (e Action) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Action) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Action) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Action) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ActionOutput struct{ *pulumi.OutputState }

func (ActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (o ActionOutput) ToActionOutput() ActionOutput {
	return o
}

func (o ActionOutput) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return o
}

func (o ActionOutput) ToActionPtrOutput() ActionPtrOutput {
	return o.ToActionPtrOutputWithContext(context.Background())
}

func (o ActionOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Action) *Action {
		return &v
	}).(ActionPtrOutput)
}

func (o ActionOutput) ToOutput(ctx context.Context) pulumix.Output[Action] {
	return pulumix.Output[Action]{
		OutputState: o.OutputState,
	}
}

func (o ActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Action) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Action) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ActionPtrOutput struct{ *pulumi.OutputState }

func (ActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (o ActionPtrOutput) ToActionPtrOutput() ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*Action] {
	return pulumix.Output[*Action]{
		OutputState: o.OutputState,
	}
}

func (o ActionPtrOutput) Elem() ActionOutput {
	return o.ApplyT(func(v *Action) Action {
		if v != nil {
			return *v
		}
		var ret Action
		return ret
	}).(ActionOutput)
}

func (o ActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Action) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ActionInput is an input type that accepts ActionArgs and ActionOutput values.
// You can construct a concrete instance of `ActionInput` via:
//
//	ActionArgs{...}
type ActionInput interface {
	pulumi.Input

	ToActionOutput() ActionOutput
	ToActionOutputWithContext(context.Context) ActionOutput
}

var actionPtrType = reflect.TypeOf((**Action)(nil)).Elem()

type ActionPtrInput interface {
	pulumi.Input

	ToActionPtrOutput() ActionPtrOutput
	ToActionPtrOutputWithContext(context.Context) ActionPtrOutput
}

type actionPtr string

func ActionPtr(v string) ActionPtrInput {
	return (*actionPtr)(&v)
}

func (*actionPtr) ElementType() reflect.Type {
	return actionPtrType
}

func (in *actionPtr) ToActionPtrOutput() ActionPtrOutput {
	return pulumi.ToOutput(in).(ActionPtrOutput)
}

func (in *actionPtr) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ActionPtrOutput)
}

func (in *actionPtr) ToOutput(ctx context.Context) pulumix.Output[*Action] {
	return pulumix.Output[*Action]{
		OutputState: in.ToActionPtrOutputWithContext(ctx).OutputState,
	}
}

// Type of encryption
type EncryptionType string

const (
	// Volume is encrypted at rest with Platform managed key. It is the default encryption type.
	EncryptionTypeEncryptionAtRestWithPlatformKey = EncryptionType("EncryptionAtRestWithPlatformKey")
)

// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusFailed   = PrivateEndpointServiceConnectionStatus("Failed")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

// The sku name.
type SkuName string

const (
	// Premium locally redundant storage
	SkuName_Premium_LRS = SkuName("Premium_LRS")
	// Premium zone redundant storage
	SkuName_Premium_ZRS = SkuName("Premium_ZRS")
)

// The sku tier.
type SkuTier string

const (
	// Premium Tier
	SkuTierPremium = SkuTier("Premium")
)

// Type of storage target
type StorageTargetType string

const (
	StorageTargetTypeIscsi = StorageTargetType("Iscsi")
	StorageTargetTypeNone  = StorageTargetType("None")
)

// This enumerates the possible sources of a volume creation.
type VolumeCreateOption string

const (
	VolumeCreateOptionNone = VolumeCreateOption("None")
)

func (VolumeCreateOption) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeCreateOption)(nil)).Elem()
}

func (e VolumeCreateOption) ToVolumeCreateOptionOutput() VolumeCreateOptionOutput {
	return pulumi.ToOutput(e).(VolumeCreateOptionOutput)
}

func (e VolumeCreateOption) ToVolumeCreateOptionOutputWithContext(ctx context.Context) VolumeCreateOptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VolumeCreateOptionOutput)
}

func (e VolumeCreateOption) ToVolumeCreateOptionPtrOutput() VolumeCreateOptionPtrOutput {
	return e.ToVolumeCreateOptionPtrOutputWithContext(context.Background())
}

func (e VolumeCreateOption) ToVolumeCreateOptionPtrOutputWithContext(ctx context.Context) VolumeCreateOptionPtrOutput {
	return VolumeCreateOption(e).ToVolumeCreateOptionOutputWithContext(ctx).ToVolumeCreateOptionPtrOutputWithContext(ctx)
}

func (e VolumeCreateOption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VolumeCreateOption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VolumeCreateOption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VolumeCreateOption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VolumeCreateOptionOutput struct{ *pulumi.OutputState }

func (VolumeCreateOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeCreateOption)(nil)).Elem()
}

func (o VolumeCreateOptionOutput) ToVolumeCreateOptionOutput() VolumeCreateOptionOutput {
	return o
}

func (o VolumeCreateOptionOutput) ToVolumeCreateOptionOutputWithContext(ctx context.Context) VolumeCreateOptionOutput {
	return o
}

func (o VolumeCreateOptionOutput) ToVolumeCreateOptionPtrOutput() VolumeCreateOptionPtrOutput {
	return o.ToVolumeCreateOptionPtrOutputWithContext(context.Background())
}

func (o VolumeCreateOptionOutput) ToVolumeCreateOptionPtrOutputWithContext(ctx context.Context) VolumeCreateOptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeCreateOption) *VolumeCreateOption {
		return &v
	}).(VolumeCreateOptionPtrOutput)
}

func (o VolumeCreateOptionOutput) ToOutput(ctx context.Context) pulumix.Output[VolumeCreateOption] {
	return pulumix.Output[VolumeCreateOption]{
		OutputState: o.OutputState,
	}
}

func (o VolumeCreateOptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VolumeCreateOptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VolumeCreateOption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VolumeCreateOptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VolumeCreateOptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VolumeCreateOption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VolumeCreateOptionPtrOutput struct{ *pulumi.OutputState }

func (VolumeCreateOptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeCreateOption)(nil)).Elem()
}

func (o VolumeCreateOptionPtrOutput) ToVolumeCreateOptionPtrOutput() VolumeCreateOptionPtrOutput {
	return o
}

func (o VolumeCreateOptionPtrOutput) ToVolumeCreateOptionPtrOutputWithContext(ctx context.Context) VolumeCreateOptionPtrOutput {
	return o
}

func (o VolumeCreateOptionPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*VolumeCreateOption] {
	return pulumix.Output[*VolumeCreateOption]{
		OutputState: o.OutputState,
	}
}

func (o VolumeCreateOptionPtrOutput) Elem() VolumeCreateOptionOutput {
	return o.ApplyT(func(v *VolumeCreateOption) VolumeCreateOption {
		if v != nil {
			return *v
		}
		var ret VolumeCreateOption
		return ret
	}).(VolumeCreateOptionOutput)
}

func (o VolumeCreateOptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VolumeCreateOptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VolumeCreateOption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// VolumeCreateOptionInput is an input type that accepts VolumeCreateOptionArgs and VolumeCreateOptionOutput values.
// You can construct a concrete instance of `VolumeCreateOptionInput` via:
//
//	VolumeCreateOptionArgs{...}
type VolumeCreateOptionInput interface {
	pulumi.Input

	ToVolumeCreateOptionOutput() VolumeCreateOptionOutput
	ToVolumeCreateOptionOutputWithContext(context.Context) VolumeCreateOptionOutput
}

var volumeCreateOptionPtrType = reflect.TypeOf((**VolumeCreateOption)(nil)).Elem()

type VolumeCreateOptionPtrInput interface {
	pulumi.Input

	ToVolumeCreateOptionPtrOutput() VolumeCreateOptionPtrOutput
	ToVolumeCreateOptionPtrOutputWithContext(context.Context) VolumeCreateOptionPtrOutput
}

type volumeCreateOptionPtr string

func VolumeCreateOptionPtr(v string) VolumeCreateOptionPtrInput {
	return (*volumeCreateOptionPtr)(&v)
}

func (*volumeCreateOptionPtr) ElementType() reflect.Type {
	return volumeCreateOptionPtrType
}

func (in *volumeCreateOptionPtr) ToVolumeCreateOptionPtrOutput() VolumeCreateOptionPtrOutput {
	return pulumi.ToOutput(in).(VolumeCreateOptionPtrOutput)
}

func (in *volumeCreateOptionPtr) ToVolumeCreateOptionPtrOutputWithContext(ctx context.Context) VolumeCreateOptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VolumeCreateOptionPtrOutput)
}

func (in *volumeCreateOptionPtr) ToOutput(ctx context.Context) pulumix.Output[*VolumeCreateOption] {
	return pulumix.Output[*VolumeCreateOption]{
		OutputState: in.ToVolumeCreateOptionPtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionPtrOutput{})
	pulumi.RegisterOutputType(VolumeCreateOptionOutput{})
	pulumi.RegisterOutputType(VolumeCreateOptionPtrOutput{})
}
