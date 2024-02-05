// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180301beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Payment channel for the SaasSubscription.
type PaymentChannelType string

const (
	PaymentChannelTypeSubscriptionDelegated = PaymentChannelType("SubscriptionDelegated")
	PaymentChannelTypeCustomerDelegated     = PaymentChannelType("CustomerDelegated")
)

func (PaymentChannelType) ElementType() reflect.Type {
	return reflect.TypeOf((*PaymentChannelType)(nil)).Elem()
}

func (e PaymentChannelType) ToPaymentChannelTypeOutput() PaymentChannelTypeOutput {
	return pulumi.ToOutput(e).(PaymentChannelTypeOutput)
}

func (e PaymentChannelType) ToPaymentChannelTypeOutputWithContext(ctx context.Context) PaymentChannelTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PaymentChannelTypeOutput)
}

func (e PaymentChannelType) ToPaymentChannelTypePtrOutput() PaymentChannelTypePtrOutput {
	return e.ToPaymentChannelTypePtrOutputWithContext(context.Background())
}

func (e PaymentChannelType) ToPaymentChannelTypePtrOutputWithContext(ctx context.Context) PaymentChannelTypePtrOutput {
	return PaymentChannelType(e).ToPaymentChannelTypeOutputWithContext(ctx).ToPaymentChannelTypePtrOutputWithContext(ctx)
}

func (e PaymentChannelType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PaymentChannelType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PaymentChannelType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PaymentChannelType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PaymentChannelTypeOutput struct{ *pulumi.OutputState }

func (PaymentChannelTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PaymentChannelType)(nil)).Elem()
}

func (o PaymentChannelTypeOutput) ToPaymentChannelTypeOutput() PaymentChannelTypeOutput {
	return o
}

func (o PaymentChannelTypeOutput) ToPaymentChannelTypeOutputWithContext(ctx context.Context) PaymentChannelTypeOutput {
	return o
}

func (o PaymentChannelTypeOutput) ToPaymentChannelTypePtrOutput() PaymentChannelTypePtrOutput {
	return o.ToPaymentChannelTypePtrOutputWithContext(context.Background())
}

func (o PaymentChannelTypeOutput) ToPaymentChannelTypePtrOutputWithContext(ctx context.Context) PaymentChannelTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PaymentChannelType) *PaymentChannelType {
		return &v
	}).(PaymentChannelTypePtrOutput)
}

func (o PaymentChannelTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PaymentChannelTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PaymentChannelType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PaymentChannelTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PaymentChannelTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PaymentChannelType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PaymentChannelTypePtrOutput struct{ *pulumi.OutputState }

func (PaymentChannelTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PaymentChannelType)(nil)).Elem()
}

func (o PaymentChannelTypePtrOutput) ToPaymentChannelTypePtrOutput() PaymentChannelTypePtrOutput {
	return o
}

func (o PaymentChannelTypePtrOutput) ToPaymentChannelTypePtrOutputWithContext(ctx context.Context) PaymentChannelTypePtrOutput {
	return o
}

func (o PaymentChannelTypePtrOutput) Elem() PaymentChannelTypeOutput {
	return o.ApplyT(func(v *PaymentChannelType) PaymentChannelType {
		if v != nil {
			return *v
		}
		var ret PaymentChannelType
		return ret
	}).(PaymentChannelTypeOutput)
}

func (o PaymentChannelTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PaymentChannelTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PaymentChannelType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PaymentChannelTypeInput is an input type that accepts values of the PaymentChannelType enum
// A concrete instance of `PaymentChannelTypeInput` can be one of the following:
//
//	PaymentChannelTypeSubscriptionDelegated
//	PaymentChannelTypeCustomerDelegated
type PaymentChannelTypeInput interface {
	pulumi.Input

	ToPaymentChannelTypeOutput() PaymentChannelTypeOutput
	ToPaymentChannelTypeOutputWithContext(context.Context) PaymentChannelTypeOutput
}

var paymentChannelTypePtrType = reflect.TypeOf((**PaymentChannelType)(nil)).Elem()

type PaymentChannelTypePtrInput interface {
	pulumi.Input

	ToPaymentChannelTypePtrOutput() PaymentChannelTypePtrOutput
	ToPaymentChannelTypePtrOutputWithContext(context.Context) PaymentChannelTypePtrOutput
}

type paymentChannelTypePtr string

func PaymentChannelTypePtr(v string) PaymentChannelTypePtrInput {
	return (*paymentChannelTypePtr)(&v)
}

func (*paymentChannelTypePtr) ElementType() reflect.Type {
	return paymentChannelTypePtrType
}

func (in *paymentChannelTypePtr) ToPaymentChannelTypePtrOutput() PaymentChannelTypePtrOutput {
	return pulumi.ToOutput(in).(PaymentChannelTypePtrOutput)
}

func (in *paymentChannelTypePtr) ToPaymentChannelTypePtrOutputWithContext(ctx context.Context) PaymentChannelTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PaymentChannelTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PaymentChannelTypeOutput{})
	pulumi.RegisterOutputType(PaymentChannelTypePtrOutput{})
}
