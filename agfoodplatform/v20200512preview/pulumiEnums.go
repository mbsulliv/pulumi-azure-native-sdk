// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200512preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
type SkuTier string

const (
	SkuTierFree     = SkuTier("Free")
	SkuTierBasic    = SkuTier("Basic")
	SkuTierStandard = SkuTier("Standard")
	SkuTierPremium  = SkuTier("Premium")
)

func (SkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (e SkuTier) ToSkuTierOutput() SkuTierOutput {
	return pulumi.ToOutput(e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return e.ToSkuTierPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return SkuTier(e).ToSkuTierOutputWithContext(ctx).ToSkuTierPtrOutputWithContext(ctx)
}

func (e SkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuTierOutput struct{ *pulumi.OutputState }

func (SkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (o SkuTierOutput) ToSkuTierOutput() SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o.ToSkuTierPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuTier) *SkuTier {
		return &v
	}).(SkuTierPtrOutput)
}

func (o SkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuTierPtrOutput struct{ *pulumi.OutputState }

func (SkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuTier)(nil)).Elem()
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) Elem() SkuTierOutput {
	return o.ApplyT(func(v *SkuTier) SkuTier {
		if v != nil {
			return *v
		}
		var ret SkuTier
		return ret
	}).(SkuTierOutput)
}

func (o SkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SkuTierInput is an input type that accepts values of the SkuTier enum
// A concrete instance of `SkuTierInput` can be one of the following:
//
//	SkuTierFree
//	SkuTierBasic
//	SkuTierStandard
//	SkuTierPremium
type SkuTierInput interface {
	pulumi.Input

	ToSkuTierOutput() SkuTierOutput
	ToSkuTierOutputWithContext(context.Context) SkuTierOutput
}

var skuTierPtrType = reflect.TypeOf((**SkuTier)(nil)).Elem()

type SkuTierPtrInput interface {
	pulumi.Input

	ToSkuTierPtrOutput() SkuTierPtrOutput
	ToSkuTierPtrOutputWithContext(context.Context) SkuTierPtrOutput
}

type skuTierPtr string

func SkuTierPtr(v string) SkuTierPtrInput {
	return (*skuTierPtr)(&v)
}

func (*skuTierPtr) ElementType() reflect.Type {
	return skuTierPtrType
}

func (in *skuTierPtr) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return pulumi.ToOutput(in).(SkuTierPtrOutput)
}

func (in *skuTierPtr) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuTierPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SkuTierOutput{})
	pulumi.RegisterOutputType(SkuTierPtrOutput{})
}
