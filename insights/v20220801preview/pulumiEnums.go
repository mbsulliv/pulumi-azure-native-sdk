// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The criteria operator. Relevant and required only for rules of the kind LogAlert.
type ConditionOperator string

const (
	ConditionOperatorEquals             = ConditionOperator("Equals")
	ConditionOperatorGreaterThan        = ConditionOperator("GreaterThan")
	ConditionOperatorGreaterThanOrEqual = ConditionOperator("GreaterThanOrEqual")
	ConditionOperatorLessThan           = ConditionOperator("LessThan")
	ConditionOperatorLessThanOrEqual    = ConditionOperator("LessThanOrEqual")
)

// Operator for dimension values
type DimensionOperator string

const (
	DimensionOperatorInclude = DimensionOperator("Include")
	DimensionOperatorExclude = DimensionOperator("Exclude")
)

// Type of managed service identity.
type IdentityType string

const (
	IdentityTypeSystemAssigned = IdentityType("SystemAssigned")
	IdentityTypeUserAssigned   = IdentityType("UserAssigned")
	IdentityTypeNone           = IdentityType("None")
)

func (IdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (e IdentityType) ToIdentityTypeOutput() IdentityTypeOutput {
	return pulumi.ToOutput(e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return e.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (e IdentityType) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return IdentityType(e).ToIdentityTypeOutputWithContext(ctx).ToIdentityTypePtrOutputWithContext(ctx)
}

func (e IdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IdentityTypeOutput struct{ *pulumi.OutputState }

func (IdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (o IdentityTypeOutput) ToIdentityTypeOutput() IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityType) *IdentityType {
		return &v
	}).(IdentityTypePtrOutput)
}

func (o IdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentityTypePtrOutput struct{ *pulumi.OutputState }

func (IdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityType)(nil)).Elem()
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) Elem() IdentityTypeOutput {
	return o.ApplyT(func(v *IdentityType) IdentityType {
		if v != nil {
			return *v
		}
		var ret IdentityType
		return ret
	}).(IdentityTypeOutput)
}

func (o IdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IdentityTypeInput is an input type that accepts IdentityTypeArgs and IdentityTypeOutput values.
// You can construct a concrete instance of `IdentityTypeInput` via:
//
//	IdentityTypeArgs{...}
type IdentityTypeInput interface {
	pulumi.Input

	ToIdentityTypeOutput() IdentityTypeOutput
	ToIdentityTypeOutputWithContext(context.Context) IdentityTypeOutput
}

var identityTypePtrType = reflect.TypeOf((**IdentityType)(nil)).Elem()

type IdentityTypePtrInput interface {
	pulumi.Input

	ToIdentityTypePtrOutput() IdentityTypePtrOutput
	ToIdentityTypePtrOutputWithContext(context.Context) IdentityTypePtrOutput
}

type identityTypePtr string

func IdentityTypePtr(v string) IdentityTypePtrInput {
	return (*identityTypePtr)(&v)
}

func (*identityTypePtr) ElementType() reflect.Type {
	return identityTypePtrType
}

func (in *identityTypePtr) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return pulumi.ToOutput(in).(IdentityTypePtrOutput)
}

func (in *identityTypePtr) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IdentityTypePtrOutput)
}

func (in *identityTypePtr) ToOutput(ctx context.Context) pulumix.Output[*IdentityType] {
	return pulumix.Output[*IdentityType]{
		OutputState: in.ToIdentityTypePtrOutputWithContext(ctx).OutputState,
	}
}

// Indicates the type of scheduled query rule. The default is LogAlert.
type Kind string

const (
	KindLogAlert    = Kind("LogAlert")
	KindLogToMetric = Kind("LogToMetric")
)

// This determines if traffic is allowed over public network. By default it is enabled.
type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled            = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled           = PublicNetworkAccess("Disabled")
	PublicNetworkAccessSecuredByPerimeter = PublicNetworkAccess("SecuredByPerimeter")
)

// Aggregation type. Relevant and required only for rules of the kind LogAlert.
type TimeAggregation string

const (
	TimeAggregationCount   = TimeAggregation("Count")
	TimeAggregationAverage = TimeAggregation("Average")
	TimeAggregationMinimum = TimeAggregation("Minimum")
	TimeAggregationMaximum = TimeAggregation("Maximum")
	TimeAggregationTotal   = TimeAggregation("Total")
)

func init() {
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
}
