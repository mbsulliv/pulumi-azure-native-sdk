// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The recurrence type : weekly, monthly, etc.
type AccessReviewRecurrencePatternType string

const (
	AccessReviewRecurrencePatternTypeWeekly          = AccessReviewRecurrencePatternType("weekly")
	AccessReviewRecurrencePatternTypeAbsoluteMonthly = AccessReviewRecurrencePatternType("absoluteMonthly")
)

// The recurrence range type. The possible values are: endDate, noEnd, numbered.
type AccessReviewRecurrenceRangeType string

const (
	AccessReviewRecurrenceRangeTypeEndDate  = AccessReviewRecurrenceRangeType("endDate")
	AccessReviewRecurrenceRangeTypeNoEnd    = AccessReviewRecurrenceRangeType("noEnd")
	AccessReviewRecurrenceRangeTypeNumbered = AccessReviewRecurrenceRangeType("numbered")
)

func (AccessReviewRecurrenceRangeType) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewRecurrenceRangeType)(nil)).Elem()
}

func (e AccessReviewRecurrenceRangeType) ToAccessReviewRecurrenceRangeTypeOutput() AccessReviewRecurrenceRangeTypeOutput {
	return pulumi.ToOutput(e).(AccessReviewRecurrenceRangeTypeOutput)
}

func (e AccessReviewRecurrenceRangeType) ToAccessReviewRecurrenceRangeTypeOutputWithContext(ctx context.Context) AccessReviewRecurrenceRangeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessReviewRecurrenceRangeTypeOutput)
}

func (e AccessReviewRecurrenceRangeType) ToAccessReviewRecurrenceRangeTypePtrOutput() AccessReviewRecurrenceRangeTypePtrOutput {
	return e.ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(context.Background())
}

func (e AccessReviewRecurrenceRangeType) ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(ctx context.Context) AccessReviewRecurrenceRangeTypePtrOutput {
	return AccessReviewRecurrenceRangeType(e).ToAccessReviewRecurrenceRangeTypeOutputWithContext(ctx).ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(ctx)
}

func (e AccessReviewRecurrenceRangeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessReviewRecurrenceRangeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessReviewRecurrenceRangeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessReviewRecurrenceRangeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessReviewRecurrenceRangeTypeOutput struct{ *pulumi.OutputState }

func (AccessReviewRecurrenceRangeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewRecurrenceRangeType)(nil)).Elem()
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToAccessReviewRecurrenceRangeTypeOutput() AccessReviewRecurrenceRangeTypeOutput {
	return o
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToAccessReviewRecurrenceRangeTypeOutputWithContext(ctx context.Context) AccessReviewRecurrenceRangeTypeOutput {
	return o
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToAccessReviewRecurrenceRangeTypePtrOutput() AccessReviewRecurrenceRangeTypePtrOutput {
	return o.ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(context.Background())
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(ctx context.Context) AccessReviewRecurrenceRangeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessReviewRecurrenceRangeType) *AccessReviewRecurrenceRangeType {
		return &v
	}).(AccessReviewRecurrenceRangeTypePtrOutput)
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessReviewRecurrenceRangeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessReviewRecurrenceRangeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessReviewRecurrenceRangeTypePtrOutput struct{ *pulumi.OutputState }

func (AccessReviewRecurrenceRangeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessReviewRecurrenceRangeType)(nil)).Elem()
}

func (o AccessReviewRecurrenceRangeTypePtrOutput) ToAccessReviewRecurrenceRangeTypePtrOutput() AccessReviewRecurrenceRangeTypePtrOutput {
	return o
}

func (o AccessReviewRecurrenceRangeTypePtrOutput) ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(ctx context.Context) AccessReviewRecurrenceRangeTypePtrOutput {
	return o
}

func (o AccessReviewRecurrenceRangeTypePtrOutput) Elem() AccessReviewRecurrenceRangeTypeOutput {
	return o.ApplyT(func(v *AccessReviewRecurrenceRangeType) AccessReviewRecurrenceRangeType {
		if v != nil {
			return *v
		}
		var ret AccessReviewRecurrenceRangeType
		return ret
	}).(AccessReviewRecurrenceRangeTypeOutput)
}

func (o AccessReviewRecurrenceRangeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessReviewRecurrenceRangeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessReviewRecurrenceRangeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AccessReviewRecurrenceRangeTypeInput is an input type that accepts values of the AccessReviewRecurrenceRangeType enum
// A concrete instance of `AccessReviewRecurrenceRangeTypeInput` can be one of the following:
//
//	AccessReviewRecurrenceRangeTypeEndDate
//	AccessReviewRecurrenceRangeTypeNoEnd
//	AccessReviewRecurrenceRangeTypeNumbered
type AccessReviewRecurrenceRangeTypeInput interface {
	pulumi.Input

	ToAccessReviewRecurrenceRangeTypeOutput() AccessReviewRecurrenceRangeTypeOutput
	ToAccessReviewRecurrenceRangeTypeOutputWithContext(context.Context) AccessReviewRecurrenceRangeTypeOutput
}

var accessReviewRecurrenceRangeTypePtrType = reflect.TypeOf((**AccessReviewRecurrenceRangeType)(nil)).Elem()

type AccessReviewRecurrenceRangeTypePtrInput interface {
	pulumi.Input

	ToAccessReviewRecurrenceRangeTypePtrOutput() AccessReviewRecurrenceRangeTypePtrOutput
	ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(context.Context) AccessReviewRecurrenceRangeTypePtrOutput
}

type accessReviewRecurrenceRangeTypePtr string

func AccessReviewRecurrenceRangeTypePtr(v string) AccessReviewRecurrenceRangeTypePtrInput {
	return (*accessReviewRecurrenceRangeTypePtr)(&v)
}

func (*accessReviewRecurrenceRangeTypePtr) ElementType() reflect.Type {
	return accessReviewRecurrenceRangeTypePtrType
}

func (in *accessReviewRecurrenceRangeTypePtr) ToAccessReviewRecurrenceRangeTypePtrOutput() AccessReviewRecurrenceRangeTypePtrOutput {
	return pulumi.ToOutput(in).(AccessReviewRecurrenceRangeTypePtrOutput)
}

func (in *accessReviewRecurrenceRangeTypePtr) ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(ctx context.Context) AccessReviewRecurrenceRangeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessReviewRecurrenceRangeTypePtrOutput)
}

// Represents a reviewer's decision for a given review
type AccessReviewResult string

const (
	AccessReviewResultApprove     = AccessReviewResult("Approve")
	AccessReviewResultDeny        = AccessReviewResult("Deny")
	AccessReviewResultNotReviewed = AccessReviewResult("NotReviewed")
	AccessReviewResultDontKnow    = AccessReviewResult("DontKnow")
	AccessReviewResultNotNotified = AccessReviewResult("NotNotified")
)

func (AccessReviewResult) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewResult)(nil)).Elem()
}

func (e AccessReviewResult) ToAccessReviewResultOutput() AccessReviewResultOutput {
	return pulumi.ToOutput(e).(AccessReviewResultOutput)
}

func (e AccessReviewResult) ToAccessReviewResultOutputWithContext(ctx context.Context) AccessReviewResultOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessReviewResultOutput)
}

func (e AccessReviewResult) ToAccessReviewResultPtrOutput() AccessReviewResultPtrOutput {
	return e.ToAccessReviewResultPtrOutputWithContext(context.Background())
}

func (e AccessReviewResult) ToAccessReviewResultPtrOutputWithContext(ctx context.Context) AccessReviewResultPtrOutput {
	return AccessReviewResult(e).ToAccessReviewResultOutputWithContext(ctx).ToAccessReviewResultPtrOutputWithContext(ctx)
}

func (e AccessReviewResult) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessReviewResult) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessReviewResult) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessReviewResult) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessReviewResultOutput struct{ *pulumi.OutputState }

func (AccessReviewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewResult)(nil)).Elem()
}

func (o AccessReviewResultOutput) ToAccessReviewResultOutput() AccessReviewResultOutput {
	return o
}

func (o AccessReviewResultOutput) ToAccessReviewResultOutputWithContext(ctx context.Context) AccessReviewResultOutput {
	return o
}

func (o AccessReviewResultOutput) ToAccessReviewResultPtrOutput() AccessReviewResultPtrOutput {
	return o.ToAccessReviewResultPtrOutputWithContext(context.Background())
}

func (o AccessReviewResultOutput) ToAccessReviewResultPtrOutputWithContext(ctx context.Context) AccessReviewResultPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessReviewResult) *AccessReviewResult {
		return &v
	}).(AccessReviewResultPtrOutput)
}

func (o AccessReviewResultOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessReviewResultOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessReviewResult) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessReviewResultOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessReviewResultOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessReviewResult) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessReviewResultPtrOutput struct{ *pulumi.OutputState }

func (AccessReviewResultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessReviewResult)(nil)).Elem()
}

func (o AccessReviewResultPtrOutput) ToAccessReviewResultPtrOutput() AccessReviewResultPtrOutput {
	return o
}

func (o AccessReviewResultPtrOutput) ToAccessReviewResultPtrOutputWithContext(ctx context.Context) AccessReviewResultPtrOutput {
	return o
}

func (o AccessReviewResultPtrOutput) Elem() AccessReviewResultOutput {
	return o.ApplyT(func(v *AccessReviewResult) AccessReviewResult {
		if v != nil {
			return *v
		}
		var ret AccessReviewResult
		return ret
	}).(AccessReviewResultOutput)
}

func (o AccessReviewResultPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessReviewResultPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessReviewResult) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AccessReviewResultInput is an input type that accepts values of the AccessReviewResult enum
// A concrete instance of `AccessReviewResultInput` can be one of the following:
//
//	AccessReviewResultApprove
//	AccessReviewResultDeny
//	AccessReviewResultNotReviewed
//	AccessReviewResultDontKnow
//	AccessReviewResultNotNotified
type AccessReviewResultInput interface {
	pulumi.Input

	ToAccessReviewResultOutput() AccessReviewResultOutput
	ToAccessReviewResultOutputWithContext(context.Context) AccessReviewResultOutput
}

var accessReviewResultPtrType = reflect.TypeOf((**AccessReviewResult)(nil)).Elem()

type AccessReviewResultPtrInput interface {
	pulumi.Input

	ToAccessReviewResultPtrOutput() AccessReviewResultPtrOutput
	ToAccessReviewResultPtrOutputWithContext(context.Context) AccessReviewResultPtrOutput
}

type accessReviewResultPtr string

func AccessReviewResultPtr(v string) AccessReviewResultPtrInput {
	return (*accessReviewResultPtr)(&v)
}

func (*accessReviewResultPtr) ElementType() reflect.Type {
	return accessReviewResultPtrType
}

func (in *accessReviewResultPtr) ToAccessReviewResultPtrOutput() AccessReviewResultPtrOutput {
	return pulumi.ToOutput(in).(AccessReviewResultPtrOutput)
}

func (in *accessReviewResultPtr) ToAccessReviewResultPtrOutputWithContext(ctx context.Context) AccessReviewResultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessReviewResultPtrOutput)
}

// This specifies the behavior for the autoReview feature when an access review completes.
type DefaultDecisionType string

const (
	DefaultDecisionTypeApprove        = DefaultDecisionType("Approve")
	DefaultDecisionTypeDeny           = DefaultDecisionType("Deny")
	DefaultDecisionTypeRecommendation = DefaultDecisionType("Recommendation")
)

func (DefaultDecisionType) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultDecisionType)(nil)).Elem()
}

func (e DefaultDecisionType) ToDefaultDecisionTypeOutput() DefaultDecisionTypeOutput {
	return pulumi.ToOutput(e).(DefaultDecisionTypeOutput)
}

func (e DefaultDecisionType) ToDefaultDecisionTypeOutputWithContext(ctx context.Context) DefaultDecisionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DefaultDecisionTypeOutput)
}

func (e DefaultDecisionType) ToDefaultDecisionTypePtrOutput() DefaultDecisionTypePtrOutput {
	return e.ToDefaultDecisionTypePtrOutputWithContext(context.Background())
}

func (e DefaultDecisionType) ToDefaultDecisionTypePtrOutputWithContext(ctx context.Context) DefaultDecisionTypePtrOutput {
	return DefaultDecisionType(e).ToDefaultDecisionTypeOutputWithContext(ctx).ToDefaultDecisionTypePtrOutputWithContext(ctx)
}

func (e DefaultDecisionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultDecisionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultDecisionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DefaultDecisionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DefaultDecisionTypeOutput struct{ *pulumi.OutputState }

func (DefaultDecisionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultDecisionType)(nil)).Elem()
}

func (o DefaultDecisionTypeOutput) ToDefaultDecisionTypeOutput() DefaultDecisionTypeOutput {
	return o
}

func (o DefaultDecisionTypeOutput) ToDefaultDecisionTypeOutputWithContext(ctx context.Context) DefaultDecisionTypeOutput {
	return o
}

func (o DefaultDecisionTypeOutput) ToDefaultDecisionTypePtrOutput() DefaultDecisionTypePtrOutput {
	return o.ToDefaultDecisionTypePtrOutputWithContext(context.Background())
}

func (o DefaultDecisionTypeOutput) ToDefaultDecisionTypePtrOutputWithContext(ctx context.Context) DefaultDecisionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultDecisionType) *DefaultDecisionType {
		return &v
	}).(DefaultDecisionTypePtrOutput)
}

func (o DefaultDecisionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DefaultDecisionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultDecisionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DefaultDecisionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultDecisionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultDecisionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DefaultDecisionTypePtrOutput struct{ *pulumi.OutputState }

func (DefaultDecisionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultDecisionType)(nil)).Elem()
}

func (o DefaultDecisionTypePtrOutput) ToDefaultDecisionTypePtrOutput() DefaultDecisionTypePtrOutput {
	return o
}

func (o DefaultDecisionTypePtrOutput) ToDefaultDecisionTypePtrOutputWithContext(ctx context.Context) DefaultDecisionTypePtrOutput {
	return o
}

func (o DefaultDecisionTypePtrOutput) Elem() DefaultDecisionTypeOutput {
	return o.ApplyT(func(v *DefaultDecisionType) DefaultDecisionType {
		if v != nil {
			return *v
		}
		var ret DefaultDecisionType
		return ret
	}).(DefaultDecisionTypeOutput)
}

func (o DefaultDecisionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultDecisionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DefaultDecisionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DefaultDecisionTypeInput is an input type that accepts values of the DefaultDecisionType enum
// A concrete instance of `DefaultDecisionTypeInput` can be one of the following:
//
//	DefaultDecisionTypeApprove
//	DefaultDecisionTypeDeny
//	DefaultDecisionTypeRecommendation
type DefaultDecisionTypeInput interface {
	pulumi.Input

	ToDefaultDecisionTypeOutput() DefaultDecisionTypeOutput
	ToDefaultDecisionTypeOutputWithContext(context.Context) DefaultDecisionTypeOutput
}

var defaultDecisionTypePtrType = reflect.TypeOf((**DefaultDecisionType)(nil)).Elem()

type DefaultDecisionTypePtrInput interface {
	pulumi.Input

	ToDefaultDecisionTypePtrOutput() DefaultDecisionTypePtrOutput
	ToDefaultDecisionTypePtrOutputWithContext(context.Context) DefaultDecisionTypePtrOutput
}

type defaultDecisionTypePtr string

func DefaultDecisionTypePtr(v string) DefaultDecisionTypePtrInput {
	return (*defaultDecisionTypePtr)(&v)
}

func (*defaultDecisionTypePtr) ElementType() reflect.Type {
	return defaultDecisionTypePtrType
}

func (in *defaultDecisionTypePtr) ToDefaultDecisionTypePtrOutput() DefaultDecisionTypePtrOutput {
	return pulumi.ToOutput(in).(DefaultDecisionTypePtrOutput)
}

func (in *defaultDecisionTypePtr) ToDefaultDecisionTypePtrOutputWithContext(ctx context.Context) DefaultDecisionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DefaultDecisionTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessReviewRecurrenceRangeTypeOutput{})
	pulumi.RegisterOutputType(AccessReviewRecurrenceRangeTypePtrOutput{})
	pulumi.RegisterOutputType(AccessReviewResultOutput{})
	pulumi.RegisterOutputType(AccessReviewResultPtrOutput{})
	pulumi.RegisterOutputType(DefaultDecisionTypeOutput{})
	pulumi.RegisterOutputType(DefaultDecisionTypePtrOutput{})
}
