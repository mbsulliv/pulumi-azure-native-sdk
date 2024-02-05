// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scom

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies that the image or disk that is being used was licensed on-premises. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json)
type HybridLicenseType string

const (
	HybridLicenseTypeNone               = HybridLicenseType("None")
	HybridLicenseTypeAzureHybridBenefit = HybridLicenseType("AzureHybridBenefit")
)

func (HybridLicenseType) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridLicenseType)(nil)).Elem()
}

func (e HybridLicenseType) ToHybridLicenseTypeOutput() HybridLicenseTypeOutput {
	return pulumi.ToOutput(e).(HybridLicenseTypeOutput)
}

func (e HybridLicenseType) ToHybridLicenseTypeOutputWithContext(ctx context.Context) HybridLicenseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HybridLicenseTypeOutput)
}

func (e HybridLicenseType) ToHybridLicenseTypePtrOutput() HybridLicenseTypePtrOutput {
	return e.ToHybridLicenseTypePtrOutputWithContext(context.Background())
}

func (e HybridLicenseType) ToHybridLicenseTypePtrOutputWithContext(ctx context.Context) HybridLicenseTypePtrOutput {
	return HybridLicenseType(e).ToHybridLicenseTypeOutputWithContext(ctx).ToHybridLicenseTypePtrOutputWithContext(ctx)
}

func (e HybridLicenseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HybridLicenseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HybridLicenseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HybridLicenseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HybridLicenseTypeOutput struct{ *pulumi.OutputState }

func (HybridLicenseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridLicenseType)(nil)).Elem()
}

func (o HybridLicenseTypeOutput) ToHybridLicenseTypeOutput() HybridLicenseTypeOutput {
	return o
}

func (o HybridLicenseTypeOutput) ToHybridLicenseTypeOutputWithContext(ctx context.Context) HybridLicenseTypeOutput {
	return o
}

func (o HybridLicenseTypeOutput) ToHybridLicenseTypePtrOutput() HybridLicenseTypePtrOutput {
	return o.ToHybridLicenseTypePtrOutputWithContext(context.Background())
}

func (o HybridLicenseTypeOutput) ToHybridLicenseTypePtrOutputWithContext(ctx context.Context) HybridLicenseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HybridLicenseType) *HybridLicenseType {
		return &v
	}).(HybridLicenseTypePtrOutput)
}

func (o HybridLicenseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HybridLicenseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HybridLicenseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HybridLicenseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HybridLicenseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HybridLicenseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HybridLicenseTypePtrOutput struct{ *pulumi.OutputState }

func (HybridLicenseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridLicenseType)(nil)).Elem()
}

func (o HybridLicenseTypePtrOutput) ToHybridLicenseTypePtrOutput() HybridLicenseTypePtrOutput {
	return o
}

func (o HybridLicenseTypePtrOutput) ToHybridLicenseTypePtrOutputWithContext(ctx context.Context) HybridLicenseTypePtrOutput {
	return o
}

func (o HybridLicenseTypePtrOutput) Elem() HybridLicenseTypeOutput {
	return o.ApplyT(func(v *HybridLicenseType) HybridLicenseType {
		if v != nil {
			return *v
		}
		var ret HybridLicenseType
		return ret
	}).(HybridLicenseTypeOutput)
}

func (o HybridLicenseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HybridLicenseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HybridLicenseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// HybridLicenseTypeInput is an input type that accepts values of the HybridLicenseType enum
// A concrete instance of `HybridLicenseTypeInput` can be one of the following:
//
//	HybridLicenseTypeNone
//	HybridLicenseTypeAzureHybridBenefit
type HybridLicenseTypeInput interface {
	pulumi.Input

	ToHybridLicenseTypeOutput() HybridLicenseTypeOutput
	ToHybridLicenseTypeOutputWithContext(context.Context) HybridLicenseTypeOutput
}

var hybridLicenseTypePtrType = reflect.TypeOf((**HybridLicenseType)(nil)).Elem()

type HybridLicenseTypePtrInput interface {
	pulumi.Input

	ToHybridLicenseTypePtrOutput() HybridLicenseTypePtrOutput
	ToHybridLicenseTypePtrOutputWithContext(context.Context) HybridLicenseTypePtrOutput
}

type hybridLicenseTypePtr string

func HybridLicenseTypePtr(v string) HybridLicenseTypePtrInput {
	return (*hybridLicenseTypePtr)(&v)
}

func (*hybridLicenseTypePtr) ElementType() reflect.Type {
	return hybridLicenseTypePtrType
}

func (in *hybridLicenseTypePtr) ToHybridLicenseTypePtrOutput() HybridLicenseTypePtrOutput {
	return pulumi.ToOutput(in).(HybridLicenseTypePtrOutput)
}

func (in *hybridLicenseTypePtr) ToHybridLicenseTypePtrOutputWithContext(ctx context.Context) HybridLicenseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HybridLicenseTypePtrOutput)
}

// The identity type
type ManagedIdentityType string

const (
	ManagedIdentityTypeNone                         = ManagedIdentityType("None")
	ManagedIdentityTypeUserAssigned                 = ManagedIdentityType("UserAssigned")
	ManagedIdentityTypeSystemAssigned               = ManagedIdentityType("SystemAssigned")
	ManagedIdentityType_SystemAssigned_UserAssigned = ManagedIdentityType("SystemAssigned,UserAssigned")
)

func (ManagedIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityType)(nil)).Elem()
}

func (e ManagedIdentityType) ToManagedIdentityTypeOutput() ManagedIdentityTypeOutput {
	return pulumi.ToOutput(e).(ManagedIdentityTypeOutput)
}

func (e ManagedIdentityType) ToManagedIdentityTypeOutputWithContext(ctx context.Context) ManagedIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedIdentityTypeOutput)
}

func (e ManagedIdentityType) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return e.ToManagedIdentityTypePtrOutputWithContext(context.Background())
}

func (e ManagedIdentityType) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return ManagedIdentityType(e).ToManagedIdentityTypeOutputWithContext(ctx).ToManagedIdentityTypePtrOutputWithContext(ctx)
}

func (e ManagedIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedIdentityTypeOutput struct{ *pulumi.OutputState }

func (ManagedIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityType)(nil)).Elem()
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypeOutput() ManagedIdentityTypeOutput {
	return o
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypeOutputWithContext(ctx context.Context) ManagedIdentityTypeOutput {
	return o
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return o.ToManagedIdentityTypePtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityType) *ManagedIdentityType {
		return &v
	}).(ManagedIdentityTypePtrOutput)
}

func (o ManagedIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityType)(nil)).Elem()
}

func (o ManagedIdentityTypePtrOutput) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return o
}

func (o ManagedIdentityTypePtrOutput) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return o
}

func (o ManagedIdentityTypePtrOutput) Elem() ManagedIdentityTypeOutput {
	return o.ApplyT(func(v *ManagedIdentityType) ManagedIdentityType {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityType
		return ret
	}).(ManagedIdentityTypeOutput)
}

func (o ManagedIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ManagedIdentityTypeInput is an input type that accepts values of the ManagedIdentityType enum
// A concrete instance of `ManagedIdentityTypeInput` can be one of the following:
//
//	ManagedIdentityTypeNone
//	ManagedIdentityTypeUserAssigned
//	ManagedIdentityTypeSystemAssigned
//	ManagedIdentityType_SystemAssigned_UserAssigned
type ManagedIdentityTypeInput interface {
	pulumi.Input

	ToManagedIdentityTypeOutput() ManagedIdentityTypeOutput
	ToManagedIdentityTypeOutputWithContext(context.Context) ManagedIdentityTypeOutput
}

var managedIdentityTypePtrType = reflect.TypeOf((**ManagedIdentityType)(nil)).Elem()

type ManagedIdentityTypePtrInput interface {
	pulumi.Input

	ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput
	ToManagedIdentityTypePtrOutputWithContext(context.Context) ManagedIdentityTypePtrOutput
}

type managedIdentityTypePtr string

func ManagedIdentityTypePtr(v string) ManagedIdentityTypePtrInput {
	return (*managedIdentityTypePtr)(&v)
}

func (*managedIdentityTypePtr) ElementType() reflect.Type {
	return managedIdentityTypePtrType
}

func (in *managedIdentityTypePtr) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ManagedIdentityTypePtrOutput)
}

func (in *managedIdentityTypePtr) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedIdentityTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HybridLicenseTypeOutput{})
	pulumi.RegisterOutputType(HybridLicenseTypePtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTypeOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTypePtrOutput{})
}
