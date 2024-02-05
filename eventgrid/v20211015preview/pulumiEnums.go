// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
type AdvancedFilterOperatorType string

const (
	AdvancedFilterOperatorTypeNumberIn                  = AdvancedFilterOperatorType("NumberIn")
	AdvancedFilterOperatorTypeNumberNotIn               = AdvancedFilterOperatorType("NumberNotIn")
	AdvancedFilterOperatorTypeNumberLessThan            = AdvancedFilterOperatorType("NumberLessThan")
	AdvancedFilterOperatorTypeNumberGreaterThan         = AdvancedFilterOperatorType("NumberGreaterThan")
	AdvancedFilterOperatorTypeNumberLessThanOrEquals    = AdvancedFilterOperatorType("NumberLessThanOrEquals")
	AdvancedFilterOperatorTypeNumberGreaterThanOrEquals = AdvancedFilterOperatorType("NumberGreaterThanOrEquals")
	AdvancedFilterOperatorTypeBoolEquals                = AdvancedFilterOperatorType("BoolEquals")
	AdvancedFilterOperatorTypeStringIn                  = AdvancedFilterOperatorType("StringIn")
	AdvancedFilterOperatorTypeStringNotIn               = AdvancedFilterOperatorType("StringNotIn")
	AdvancedFilterOperatorTypeStringBeginsWith          = AdvancedFilterOperatorType("StringBeginsWith")
	AdvancedFilterOperatorTypeStringEndsWith            = AdvancedFilterOperatorType("StringEndsWith")
	AdvancedFilterOperatorTypeStringContains            = AdvancedFilterOperatorType("StringContains")
	AdvancedFilterOperatorTypeNumberInRange             = AdvancedFilterOperatorType("NumberInRange")
	AdvancedFilterOperatorTypeNumberNotInRange          = AdvancedFilterOperatorType("NumberNotInRange")
	AdvancedFilterOperatorTypeStringNotBeginsWith       = AdvancedFilterOperatorType("StringNotBeginsWith")
	AdvancedFilterOperatorTypeStringNotEndsWith         = AdvancedFilterOperatorType("StringNotEndsWith")
	AdvancedFilterOperatorTypeStringNotContains         = AdvancedFilterOperatorType("StringNotContains")
	AdvancedFilterOperatorTypeIsNullOrUndefined         = AdvancedFilterOperatorType("IsNullOrUndefined")
	AdvancedFilterOperatorTypeIsNotNull                 = AdvancedFilterOperatorType("IsNotNull")
)

// Activation state of the partner destination.
type PartnerDestinationActivationState string

const (
	PartnerDestinationActivationStateNeverActivated = PartnerDestinationActivationState("NeverActivated")
	PartnerDestinationActivationStateActivated      = PartnerDestinationActivationState("Activated")
)

func (PartnerDestinationActivationState) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerDestinationActivationState)(nil)).Elem()
}

func (e PartnerDestinationActivationState) ToPartnerDestinationActivationStateOutput() PartnerDestinationActivationStateOutput {
	return pulumi.ToOutput(e).(PartnerDestinationActivationStateOutput)
}

func (e PartnerDestinationActivationState) ToPartnerDestinationActivationStateOutputWithContext(ctx context.Context) PartnerDestinationActivationStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PartnerDestinationActivationStateOutput)
}

func (e PartnerDestinationActivationState) ToPartnerDestinationActivationStatePtrOutput() PartnerDestinationActivationStatePtrOutput {
	return e.ToPartnerDestinationActivationStatePtrOutputWithContext(context.Background())
}

func (e PartnerDestinationActivationState) ToPartnerDestinationActivationStatePtrOutputWithContext(ctx context.Context) PartnerDestinationActivationStatePtrOutput {
	return PartnerDestinationActivationState(e).ToPartnerDestinationActivationStateOutputWithContext(ctx).ToPartnerDestinationActivationStatePtrOutputWithContext(ctx)
}

func (e PartnerDestinationActivationState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartnerDestinationActivationState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartnerDestinationActivationState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PartnerDestinationActivationState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PartnerDestinationActivationStateOutput struct{ *pulumi.OutputState }

func (PartnerDestinationActivationStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerDestinationActivationState)(nil)).Elem()
}

func (o PartnerDestinationActivationStateOutput) ToPartnerDestinationActivationStateOutput() PartnerDestinationActivationStateOutput {
	return o
}

func (o PartnerDestinationActivationStateOutput) ToPartnerDestinationActivationStateOutputWithContext(ctx context.Context) PartnerDestinationActivationStateOutput {
	return o
}

func (o PartnerDestinationActivationStateOutput) ToPartnerDestinationActivationStatePtrOutput() PartnerDestinationActivationStatePtrOutput {
	return o.ToPartnerDestinationActivationStatePtrOutputWithContext(context.Background())
}

func (o PartnerDestinationActivationStateOutput) ToPartnerDestinationActivationStatePtrOutputWithContext(ctx context.Context) PartnerDestinationActivationStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PartnerDestinationActivationState) *PartnerDestinationActivationState {
		return &v
	}).(PartnerDestinationActivationStatePtrOutput)
}

func (o PartnerDestinationActivationStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PartnerDestinationActivationStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartnerDestinationActivationState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PartnerDestinationActivationStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartnerDestinationActivationStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartnerDestinationActivationState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PartnerDestinationActivationStatePtrOutput struct{ *pulumi.OutputState }

func (PartnerDestinationActivationStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerDestinationActivationState)(nil)).Elem()
}

func (o PartnerDestinationActivationStatePtrOutput) ToPartnerDestinationActivationStatePtrOutput() PartnerDestinationActivationStatePtrOutput {
	return o
}

func (o PartnerDestinationActivationStatePtrOutput) ToPartnerDestinationActivationStatePtrOutputWithContext(ctx context.Context) PartnerDestinationActivationStatePtrOutput {
	return o
}

func (o PartnerDestinationActivationStatePtrOutput) Elem() PartnerDestinationActivationStateOutput {
	return o.ApplyT(func(v *PartnerDestinationActivationState) PartnerDestinationActivationState {
		if v != nil {
			return *v
		}
		var ret PartnerDestinationActivationState
		return ret
	}).(PartnerDestinationActivationStateOutput)
}

func (o PartnerDestinationActivationStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartnerDestinationActivationStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PartnerDestinationActivationState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PartnerDestinationActivationStateInput is an input type that accepts values of the PartnerDestinationActivationState enum
// A concrete instance of `PartnerDestinationActivationStateInput` can be one of the following:
//
//	PartnerDestinationActivationStateNeverActivated
//	PartnerDestinationActivationStateActivated
type PartnerDestinationActivationStateInput interface {
	pulumi.Input

	ToPartnerDestinationActivationStateOutput() PartnerDestinationActivationStateOutput
	ToPartnerDestinationActivationStateOutputWithContext(context.Context) PartnerDestinationActivationStateOutput
}

var partnerDestinationActivationStatePtrType = reflect.TypeOf((**PartnerDestinationActivationState)(nil)).Elem()

type PartnerDestinationActivationStatePtrInput interface {
	pulumi.Input

	ToPartnerDestinationActivationStatePtrOutput() PartnerDestinationActivationStatePtrOutput
	ToPartnerDestinationActivationStatePtrOutputWithContext(context.Context) PartnerDestinationActivationStatePtrOutput
}

type partnerDestinationActivationStatePtr string

func PartnerDestinationActivationStatePtr(v string) PartnerDestinationActivationStatePtrInput {
	return (*partnerDestinationActivationStatePtr)(&v)
}

func (*partnerDestinationActivationStatePtr) ElementType() reflect.Type {
	return partnerDestinationActivationStatePtrType
}

func (in *partnerDestinationActivationStatePtr) ToPartnerDestinationActivationStatePtrOutput() PartnerDestinationActivationStatePtrOutput {
	return pulumi.ToOutput(in).(PartnerDestinationActivationStatePtrOutput)
}

func (in *partnerDestinationActivationStatePtr) ToPartnerDestinationActivationStatePtrOutputWithContext(ctx context.Context) PartnerDestinationActivationStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PartnerDestinationActivationStatePtrOutput)
}

// Provisioning state of the partner destination.
type PartnerDestinationProvisioningState string

const (
	PartnerDestinationProvisioningStateCreating  = PartnerDestinationProvisioningState("Creating")
	PartnerDestinationProvisioningStateUpdating  = PartnerDestinationProvisioningState("Updating")
	PartnerDestinationProvisioningStateDeleting  = PartnerDestinationProvisioningState("Deleting")
	PartnerDestinationProvisioningStateSucceeded = PartnerDestinationProvisioningState("Succeeded")
	PartnerDestinationProvisioningStateCanceled  = PartnerDestinationProvisioningState("Canceled")
	PartnerDestinationProvisioningStateFailed    = PartnerDestinationProvisioningState("Failed")
)

func (PartnerDestinationProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerDestinationProvisioningState)(nil)).Elem()
}

func (e PartnerDestinationProvisioningState) ToPartnerDestinationProvisioningStateOutput() PartnerDestinationProvisioningStateOutput {
	return pulumi.ToOutput(e).(PartnerDestinationProvisioningStateOutput)
}

func (e PartnerDestinationProvisioningState) ToPartnerDestinationProvisioningStateOutputWithContext(ctx context.Context) PartnerDestinationProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PartnerDestinationProvisioningStateOutput)
}

func (e PartnerDestinationProvisioningState) ToPartnerDestinationProvisioningStatePtrOutput() PartnerDestinationProvisioningStatePtrOutput {
	return e.ToPartnerDestinationProvisioningStatePtrOutputWithContext(context.Background())
}

func (e PartnerDestinationProvisioningState) ToPartnerDestinationProvisioningStatePtrOutputWithContext(ctx context.Context) PartnerDestinationProvisioningStatePtrOutput {
	return PartnerDestinationProvisioningState(e).ToPartnerDestinationProvisioningStateOutputWithContext(ctx).ToPartnerDestinationProvisioningStatePtrOutputWithContext(ctx)
}

func (e PartnerDestinationProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartnerDestinationProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartnerDestinationProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PartnerDestinationProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PartnerDestinationProvisioningStateOutput struct{ *pulumi.OutputState }

func (PartnerDestinationProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerDestinationProvisioningState)(nil)).Elem()
}

func (o PartnerDestinationProvisioningStateOutput) ToPartnerDestinationProvisioningStateOutput() PartnerDestinationProvisioningStateOutput {
	return o
}

func (o PartnerDestinationProvisioningStateOutput) ToPartnerDestinationProvisioningStateOutputWithContext(ctx context.Context) PartnerDestinationProvisioningStateOutput {
	return o
}

func (o PartnerDestinationProvisioningStateOutput) ToPartnerDestinationProvisioningStatePtrOutput() PartnerDestinationProvisioningStatePtrOutput {
	return o.ToPartnerDestinationProvisioningStatePtrOutputWithContext(context.Background())
}

func (o PartnerDestinationProvisioningStateOutput) ToPartnerDestinationProvisioningStatePtrOutputWithContext(ctx context.Context) PartnerDestinationProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PartnerDestinationProvisioningState) *PartnerDestinationProvisioningState {
		return &v
	}).(PartnerDestinationProvisioningStatePtrOutput)
}

func (o PartnerDestinationProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PartnerDestinationProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartnerDestinationProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PartnerDestinationProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartnerDestinationProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartnerDestinationProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PartnerDestinationProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (PartnerDestinationProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerDestinationProvisioningState)(nil)).Elem()
}

func (o PartnerDestinationProvisioningStatePtrOutput) ToPartnerDestinationProvisioningStatePtrOutput() PartnerDestinationProvisioningStatePtrOutput {
	return o
}

func (o PartnerDestinationProvisioningStatePtrOutput) ToPartnerDestinationProvisioningStatePtrOutputWithContext(ctx context.Context) PartnerDestinationProvisioningStatePtrOutput {
	return o
}

func (o PartnerDestinationProvisioningStatePtrOutput) Elem() PartnerDestinationProvisioningStateOutput {
	return o.ApplyT(func(v *PartnerDestinationProvisioningState) PartnerDestinationProvisioningState {
		if v != nil {
			return *v
		}
		var ret PartnerDestinationProvisioningState
		return ret
	}).(PartnerDestinationProvisioningStateOutput)
}

func (o PartnerDestinationProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartnerDestinationProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PartnerDestinationProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PartnerDestinationProvisioningStateInput is an input type that accepts values of the PartnerDestinationProvisioningState enum
// A concrete instance of `PartnerDestinationProvisioningStateInput` can be one of the following:
//
//	PartnerDestinationProvisioningStateCreating
//	PartnerDestinationProvisioningStateUpdating
//	PartnerDestinationProvisioningStateDeleting
//	PartnerDestinationProvisioningStateSucceeded
//	PartnerDestinationProvisioningStateCanceled
//	PartnerDestinationProvisioningStateFailed
type PartnerDestinationProvisioningStateInput interface {
	pulumi.Input

	ToPartnerDestinationProvisioningStateOutput() PartnerDestinationProvisioningStateOutput
	ToPartnerDestinationProvisioningStateOutputWithContext(context.Context) PartnerDestinationProvisioningStateOutput
}

var partnerDestinationProvisioningStatePtrType = reflect.TypeOf((**PartnerDestinationProvisioningState)(nil)).Elem()

type PartnerDestinationProvisioningStatePtrInput interface {
	pulumi.Input

	ToPartnerDestinationProvisioningStatePtrOutput() PartnerDestinationProvisioningStatePtrOutput
	ToPartnerDestinationProvisioningStatePtrOutputWithContext(context.Context) PartnerDestinationProvisioningStatePtrOutput
}

type partnerDestinationProvisioningStatePtr string

func PartnerDestinationProvisioningStatePtr(v string) PartnerDestinationProvisioningStatePtrInput {
	return (*partnerDestinationProvisioningStatePtr)(&v)
}

func (*partnerDestinationProvisioningStatePtr) ElementType() reflect.Type {
	return partnerDestinationProvisioningStatePtrType
}

func (in *partnerDestinationProvisioningStatePtr) ToPartnerDestinationProvisioningStatePtrOutput() PartnerDestinationProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(PartnerDestinationProvisioningStatePtrOutput)
}

func (in *partnerDestinationProvisioningStatePtr) ToPartnerDestinationProvisioningStatePtrOutputWithContext(ctx context.Context) PartnerDestinationProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PartnerDestinationProvisioningStatePtrOutput)
}

// Visibility state of the partner registration.
type PartnerRegistrationVisibilityState string

const (
	PartnerRegistrationVisibilityStateHidden             = PartnerRegistrationVisibilityState("Hidden")
	PartnerRegistrationVisibilityStatePublicPreview      = PartnerRegistrationVisibilityState("PublicPreview")
	PartnerRegistrationVisibilityStateGenerallyAvailable = PartnerRegistrationVisibilityState("GenerallyAvailable")
)

func (PartnerRegistrationVisibilityState) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerRegistrationVisibilityState)(nil)).Elem()
}

func (e PartnerRegistrationVisibilityState) ToPartnerRegistrationVisibilityStateOutput() PartnerRegistrationVisibilityStateOutput {
	return pulumi.ToOutput(e).(PartnerRegistrationVisibilityStateOutput)
}

func (e PartnerRegistrationVisibilityState) ToPartnerRegistrationVisibilityStateOutputWithContext(ctx context.Context) PartnerRegistrationVisibilityStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PartnerRegistrationVisibilityStateOutput)
}

func (e PartnerRegistrationVisibilityState) ToPartnerRegistrationVisibilityStatePtrOutput() PartnerRegistrationVisibilityStatePtrOutput {
	return e.ToPartnerRegistrationVisibilityStatePtrOutputWithContext(context.Background())
}

func (e PartnerRegistrationVisibilityState) ToPartnerRegistrationVisibilityStatePtrOutputWithContext(ctx context.Context) PartnerRegistrationVisibilityStatePtrOutput {
	return PartnerRegistrationVisibilityState(e).ToPartnerRegistrationVisibilityStateOutputWithContext(ctx).ToPartnerRegistrationVisibilityStatePtrOutputWithContext(ctx)
}

func (e PartnerRegistrationVisibilityState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartnerRegistrationVisibilityState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartnerRegistrationVisibilityState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PartnerRegistrationVisibilityState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PartnerRegistrationVisibilityStateOutput struct{ *pulumi.OutputState }

func (PartnerRegistrationVisibilityStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerRegistrationVisibilityState)(nil)).Elem()
}

func (o PartnerRegistrationVisibilityStateOutput) ToPartnerRegistrationVisibilityStateOutput() PartnerRegistrationVisibilityStateOutput {
	return o
}

func (o PartnerRegistrationVisibilityStateOutput) ToPartnerRegistrationVisibilityStateOutputWithContext(ctx context.Context) PartnerRegistrationVisibilityStateOutput {
	return o
}

func (o PartnerRegistrationVisibilityStateOutput) ToPartnerRegistrationVisibilityStatePtrOutput() PartnerRegistrationVisibilityStatePtrOutput {
	return o.ToPartnerRegistrationVisibilityStatePtrOutputWithContext(context.Background())
}

func (o PartnerRegistrationVisibilityStateOutput) ToPartnerRegistrationVisibilityStatePtrOutputWithContext(ctx context.Context) PartnerRegistrationVisibilityStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PartnerRegistrationVisibilityState) *PartnerRegistrationVisibilityState {
		return &v
	}).(PartnerRegistrationVisibilityStatePtrOutput)
}

func (o PartnerRegistrationVisibilityStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PartnerRegistrationVisibilityStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartnerRegistrationVisibilityState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PartnerRegistrationVisibilityStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartnerRegistrationVisibilityStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartnerRegistrationVisibilityState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PartnerRegistrationVisibilityStatePtrOutput struct{ *pulumi.OutputState }

func (PartnerRegistrationVisibilityStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerRegistrationVisibilityState)(nil)).Elem()
}

func (o PartnerRegistrationVisibilityStatePtrOutput) ToPartnerRegistrationVisibilityStatePtrOutput() PartnerRegistrationVisibilityStatePtrOutput {
	return o
}

func (o PartnerRegistrationVisibilityStatePtrOutput) ToPartnerRegistrationVisibilityStatePtrOutputWithContext(ctx context.Context) PartnerRegistrationVisibilityStatePtrOutput {
	return o
}

func (o PartnerRegistrationVisibilityStatePtrOutput) Elem() PartnerRegistrationVisibilityStateOutput {
	return o.ApplyT(func(v *PartnerRegistrationVisibilityState) PartnerRegistrationVisibilityState {
		if v != nil {
			return *v
		}
		var ret PartnerRegistrationVisibilityState
		return ret
	}).(PartnerRegistrationVisibilityStateOutput)
}

func (o PartnerRegistrationVisibilityStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartnerRegistrationVisibilityStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PartnerRegistrationVisibilityState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PartnerRegistrationVisibilityStateInput is an input type that accepts values of the PartnerRegistrationVisibilityState enum
// A concrete instance of `PartnerRegistrationVisibilityStateInput` can be one of the following:
//
//	PartnerRegistrationVisibilityStateHidden
//	PartnerRegistrationVisibilityStatePublicPreview
//	PartnerRegistrationVisibilityStateGenerallyAvailable
type PartnerRegistrationVisibilityStateInput interface {
	pulumi.Input

	ToPartnerRegistrationVisibilityStateOutput() PartnerRegistrationVisibilityStateOutput
	ToPartnerRegistrationVisibilityStateOutputWithContext(context.Context) PartnerRegistrationVisibilityStateOutput
}

var partnerRegistrationVisibilityStatePtrType = reflect.TypeOf((**PartnerRegistrationVisibilityState)(nil)).Elem()

type PartnerRegistrationVisibilityStatePtrInput interface {
	pulumi.Input

	ToPartnerRegistrationVisibilityStatePtrOutput() PartnerRegistrationVisibilityStatePtrOutput
	ToPartnerRegistrationVisibilityStatePtrOutputWithContext(context.Context) PartnerRegistrationVisibilityStatePtrOutput
}

type partnerRegistrationVisibilityStatePtr string

func PartnerRegistrationVisibilityStatePtr(v string) PartnerRegistrationVisibilityStatePtrInput {
	return (*partnerRegistrationVisibilityStatePtr)(&v)
}

func (*partnerRegistrationVisibilityStatePtr) ElementType() reflect.Type {
	return partnerRegistrationVisibilityStatePtrType
}

func (in *partnerRegistrationVisibilityStatePtr) ToPartnerRegistrationVisibilityStatePtrOutput() PartnerRegistrationVisibilityStatePtrOutput {
	return pulumi.ToOutput(in).(PartnerRegistrationVisibilityStatePtrOutput)
}

func (in *partnerRegistrationVisibilityStatePtr) ToPartnerRegistrationVisibilityStatePtrOutputWithContext(ctx context.Context) PartnerRegistrationVisibilityStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PartnerRegistrationVisibilityStatePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PartnerDestinationActivationStateOutput{})
	pulumi.RegisterOutputType(PartnerDestinationActivationStatePtrOutput{})
	pulumi.RegisterOutputType(PartnerDestinationProvisioningStateOutput{})
	pulumi.RegisterOutputType(PartnerDestinationProvisioningStatePtrOutput{})
	pulumi.RegisterOutputType(PartnerRegistrationVisibilityStateOutput{})
	pulumi.RegisterOutputType(PartnerRegistrationVisibilityStatePtrOutput{})
}
