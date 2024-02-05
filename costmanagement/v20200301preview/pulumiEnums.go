// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Method of cost allocation for the rule
type CostAllocationPolicyType string

const (
	CostAllocationPolicyTypeFixedProportion = CostAllocationPolicyType("FixedProportion")
)

func (CostAllocationPolicyType) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationPolicyType)(nil)).Elem()
}

func (e CostAllocationPolicyType) ToCostAllocationPolicyTypeOutput() CostAllocationPolicyTypeOutput {
	return pulumi.ToOutput(e).(CostAllocationPolicyTypeOutput)
}

func (e CostAllocationPolicyType) ToCostAllocationPolicyTypeOutputWithContext(ctx context.Context) CostAllocationPolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CostAllocationPolicyTypeOutput)
}

func (e CostAllocationPolicyType) ToCostAllocationPolicyTypePtrOutput() CostAllocationPolicyTypePtrOutput {
	return e.ToCostAllocationPolicyTypePtrOutputWithContext(context.Background())
}

func (e CostAllocationPolicyType) ToCostAllocationPolicyTypePtrOutputWithContext(ctx context.Context) CostAllocationPolicyTypePtrOutput {
	return CostAllocationPolicyType(e).ToCostAllocationPolicyTypeOutputWithContext(ctx).ToCostAllocationPolicyTypePtrOutputWithContext(ctx)
}

func (e CostAllocationPolicyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CostAllocationPolicyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CostAllocationPolicyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CostAllocationPolicyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CostAllocationPolicyTypeOutput struct{ *pulumi.OutputState }

func (CostAllocationPolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationPolicyType)(nil)).Elem()
}

func (o CostAllocationPolicyTypeOutput) ToCostAllocationPolicyTypeOutput() CostAllocationPolicyTypeOutput {
	return o
}

func (o CostAllocationPolicyTypeOutput) ToCostAllocationPolicyTypeOutputWithContext(ctx context.Context) CostAllocationPolicyTypeOutput {
	return o
}

func (o CostAllocationPolicyTypeOutput) ToCostAllocationPolicyTypePtrOutput() CostAllocationPolicyTypePtrOutput {
	return o.ToCostAllocationPolicyTypePtrOutputWithContext(context.Background())
}

func (o CostAllocationPolicyTypeOutput) ToCostAllocationPolicyTypePtrOutputWithContext(ctx context.Context) CostAllocationPolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CostAllocationPolicyType) *CostAllocationPolicyType {
		return &v
	}).(CostAllocationPolicyTypePtrOutput)
}

func (o CostAllocationPolicyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CostAllocationPolicyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CostAllocationPolicyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CostAllocationPolicyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CostAllocationPolicyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CostAllocationPolicyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CostAllocationPolicyTypePtrOutput struct{ *pulumi.OutputState }

func (CostAllocationPolicyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationPolicyType)(nil)).Elem()
}

func (o CostAllocationPolicyTypePtrOutput) ToCostAllocationPolicyTypePtrOutput() CostAllocationPolicyTypePtrOutput {
	return o
}

func (o CostAllocationPolicyTypePtrOutput) ToCostAllocationPolicyTypePtrOutputWithContext(ctx context.Context) CostAllocationPolicyTypePtrOutput {
	return o
}

func (o CostAllocationPolicyTypePtrOutput) Elem() CostAllocationPolicyTypeOutput {
	return o.ApplyT(func(v *CostAllocationPolicyType) CostAllocationPolicyType {
		if v != nil {
			return *v
		}
		var ret CostAllocationPolicyType
		return ret
	}).(CostAllocationPolicyTypeOutput)
}

func (o CostAllocationPolicyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CostAllocationPolicyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CostAllocationPolicyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CostAllocationPolicyTypeInput is an input type that accepts values of the CostAllocationPolicyType enum
// A concrete instance of `CostAllocationPolicyTypeInput` can be one of the following:
//
//	CostAllocationPolicyTypeFixedProportion
type CostAllocationPolicyTypeInput interface {
	pulumi.Input

	ToCostAllocationPolicyTypeOutput() CostAllocationPolicyTypeOutput
	ToCostAllocationPolicyTypeOutputWithContext(context.Context) CostAllocationPolicyTypeOutput
}

var costAllocationPolicyTypePtrType = reflect.TypeOf((**CostAllocationPolicyType)(nil)).Elem()

type CostAllocationPolicyTypePtrInput interface {
	pulumi.Input

	ToCostAllocationPolicyTypePtrOutput() CostAllocationPolicyTypePtrOutput
	ToCostAllocationPolicyTypePtrOutputWithContext(context.Context) CostAllocationPolicyTypePtrOutput
}

type costAllocationPolicyTypePtr string

func CostAllocationPolicyTypePtr(v string) CostAllocationPolicyTypePtrInput {
	return (*costAllocationPolicyTypePtr)(&v)
}

func (*costAllocationPolicyTypePtr) ElementType() reflect.Type {
	return costAllocationPolicyTypePtrType
}

func (in *costAllocationPolicyTypePtr) ToCostAllocationPolicyTypePtrOutput() CostAllocationPolicyTypePtrOutput {
	return pulumi.ToOutput(in).(CostAllocationPolicyTypePtrOutput)
}

func (in *costAllocationPolicyTypePtr) ToCostAllocationPolicyTypePtrOutputWithContext(ctx context.Context) CostAllocationPolicyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CostAllocationPolicyTypePtrOutput)
}

// Type of resources contained in this cost allocation rule
type CostAllocationResourceType string

const (
	// Indicates an Azure dimension such as a subscription id or resource group name is being used for allocation.
	CostAllocationResourceTypeDimension = CostAllocationResourceType("Dimension")
	// Allocates cost based on Azure Tag key value pairs.
	CostAllocationResourceTypeTag = CostAllocationResourceType("Tag")
)

func (CostAllocationResourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationResourceType)(nil)).Elem()
}

func (e CostAllocationResourceType) ToCostAllocationResourceTypeOutput() CostAllocationResourceTypeOutput {
	return pulumi.ToOutput(e).(CostAllocationResourceTypeOutput)
}

func (e CostAllocationResourceType) ToCostAllocationResourceTypeOutputWithContext(ctx context.Context) CostAllocationResourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CostAllocationResourceTypeOutput)
}

func (e CostAllocationResourceType) ToCostAllocationResourceTypePtrOutput() CostAllocationResourceTypePtrOutput {
	return e.ToCostAllocationResourceTypePtrOutputWithContext(context.Background())
}

func (e CostAllocationResourceType) ToCostAllocationResourceTypePtrOutputWithContext(ctx context.Context) CostAllocationResourceTypePtrOutput {
	return CostAllocationResourceType(e).ToCostAllocationResourceTypeOutputWithContext(ctx).ToCostAllocationResourceTypePtrOutputWithContext(ctx)
}

func (e CostAllocationResourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CostAllocationResourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CostAllocationResourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CostAllocationResourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CostAllocationResourceTypeOutput struct{ *pulumi.OutputState }

func (CostAllocationResourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationResourceType)(nil)).Elem()
}

func (o CostAllocationResourceTypeOutput) ToCostAllocationResourceTypeOutput() CostAllocationResourceTypeOutput {
	return o
}

func (o CostAllocationResourceTypeOutput) ToCostAllocationResourceTypeOutputWithContext(ctx context.Context) CostAllocationResourceTypeOutput {
	return o
}

func (o CostAllocationResourceTypeOutput) ToCostAllocationResourceTypePtrOutput() CostAllocationResourceTypePtrOutput {
	return o.ToCostAllocationResourceTypePtrOutputWithContext(context.Background())
}

func (o CostAllocationResourceTypeOutput) ToCostAllocationResourceTypePtrOutputWithContext(ctx context.Context) CostAllocationResourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CostAllocationResourceType) *CostAllocationResourceType {
		return &v
	}).(CostAllocationResourceTypePtrOutput)
}

func (o CostAllocationResourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CostAllocationResourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CostAllocationResourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CostAllocationResourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CostAllocationResourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CostAllocationResourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CostAllocationResourceTypePtrOutput struct{ *pulumi.OutputState }

func (CostAllocationResourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationResourceType)(nil)).Elem()
}

func (o CostAllocationResourceTypePtrOutput) ToCostAllocationResourceTypePtrOutput() CostAllocationResourceTypePtrOutput {
	return o
}

func (o CostAllocationResourceTypePtrOutput) ToCostAllocationResourceTypePtrOutputWithContext(ctx context.Context) CostAllocationResourceTypePtrOutput {
	return o
}

func (o CostAllocationResourceTypePtrOutput) Elem() CostAllocationResourceTypeOutput {
	return o.ApplyT(func(v *CostAllocationResourceType) CostAllocationResourceType {
		if v != nil {
			return *v
		}
		var ret CostAllocationResourceType
		return ret
	}).(CostAllocationResourceTypeOutput)
}

func (o CostAllocationResourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CostAllocationResourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CostAllocationResourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CostAllocationResourceTypeInput is an input type that accepts values of the CostAllocationResourceType enum
// A concrete instance of `CostAllocationResourceTypeInput` can be one of the following:
//
//	CostAllocationResourceTypeDimension
//	CostAllocationResourceTypeTag
type CostAllocationResourceTypeInput interface {
	pulumi.Input

	ToCostAllocationResourceTypeOutput() CostAllocationResourceTypeOutput
	ToCostAllocationResourceTypeOutputWithContext(context.Context) CostAllocationResourceTypeOutput
}

var costAllocationResourceTypePtrType = reflect.TypeOf((**CostAllocationResourceType)(nil)).Elem()

type CostAllocationResourceTypePtrInput interface {
	pulumi.Input

	ToCostAllocationResourceTypePtrOutput() CostAllocationResourceTypePtrOutput
	ToCostAllocationResourceTypePtrOutputWithContext(context.Context) CostAllocationResourceTypePtrOutput
}

type costAllocationResourceTypePtr string

func CostAllocationResourceTypePtr(v string) CostAllocationResourceTypePtrInput {
	return (*costAllocationResourceTypePtr)(&v)
}

func (*costAllocationResourceTypePtr) ElementType() reflect.Type {
	return costAllocationResourceTypePtrType
}

func (in *costAllocationResourceTypePtr) ToCostAllocationResourceTypePtrOutput() CostAllocationResourceTypePtrOutput {
	return pulumi.ToOutput(in).(CostAllocationResourceTypePtrOutput)
}

func (in *costAllocationResourceTypePtr) ToCostAllocationResourceTypePtrOutputWithContext(ctx context.Context) CostAllocationResourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CostAllocationResourceTypePtrOutput)
}

// Status of the rule
type RuleStatus string

const (
	// Rule is saved but not used to allocate costs.
	RuleStatusNotActive = RuleStatus("NotActive")
	// Rule is saved and impacting cost allocation.
	RuleStatusActive = RuleStatus("Active")
	// Rule is saved and cost allocation is being updated. Readonly value that cannot be submitted in a put request.
	RuleStatusProcessing = RuleStatus("Processing")
)

func (RuleStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleStatus)(nil)).Elem()
}

func (e RuleStatus) ToRuleStatusOutput() RuleStatusOutput {
	return pulumi.ToOutput(e).(RuleStatusOutput)
}

func (e RuleStatus) ToRuleStatusOutputWithContext(ctx context.Context) RuleStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RuleStatusOutput)
}

func (e RuleStatus) ToRuleStatusPtrOutput() RuleStatusPtrOutput {
	return e.ToRuleStatusPtrOutputWithContext(context.Background())
}

func (e RuleStatus) ToRuleStatusPtrOutputWithContext(ctx context.Context) RuleStatusPtrOutput {
	return RuleStatus(e).ToRuleStatusOutputWithContext(ctx).ToRuleStatusPtrOutputWithContext(ctx)
}

func (e RuleStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RuleStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RuleStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RuleStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RuleStatusOutput struct{ *pulumi.OutputState }

func (RuleStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleStatus)(nil)).Elem()
}

func (o RuleStatusOutput) ToRuleStatusOutput() RuleStatusOutput {
	return o
}

func (o RuleStatusOutput) ToRuleStatusOutputWithContext(ctx context.Context) RuleStatusOutput {
	return o
}

func (o RuleStatusOutput) ToRuleStatusPtrOutput() RuleStatusPtrOutput {
	return o.ToRuleStatusPtrOutputWithContext(context.Background())
}

func (o RuleStatusOutput) ToRuleStatusPtrOutputWithContext(ctx context.Context) RuleStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RuleStatus) *RuleStatus {
		return &v
	}).(RuleStatusPtrOutput)
}

func (o RuleStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RuleStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RuleStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RuleStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RuleStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RuleStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RuleStatusPtrOutput struct{ *pulumi.OutputState }

func (RuleStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleStatus)(nil)).Elem()
}

func (o RuleStatusPtrOutput) ToRuleStatusPtrOutput() RuleStatusPtrOutput {
	return o
}

func (o RuleStatusPtrOutput) ToRuleStatusPtrOutputWithContext(ctx context.Context) RuleStatusPtrOutput {
	return o
}

func (o RuleStatusPtrOutput) Elem() RuleStatusOutput {
	return o.ApplyT(func(v *RuleStatus) RuleStatus {
		if v != nil {
			return *v
		}
		var ret RuleStatus
		return ret
	}).(RuleStatusOutput)
}

func (o RuleStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RuleStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RuleStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// RuleStatusInput is an input type that accepts values of the RuleStatus enum
// A concrete instance of `RuleStatusInput` can be one of the following:
//
//	RuleStatusNotActive
//	RuleStatusActive
//	RuleStatusProcessing
type RuleStatusInput interface {
	pulumi.Input

	ToRuleStatusOutput() RuleStatusOutput
	ToRuleStatusOutputWithContext(context.Context) RuleStatusOutput
}

var ruleStatusPtrType = reflect.TypeOf((**RuleStatus)(nil)).Elem()

type RuleStatusPtrInput interface {
	pulumi.Input

	ToRuleStatusPtrOutput() RuleStatusPtrOutput
	ToRuleStatusPtrOutputWithContext(context.Context) RuleStatusPtrOutput
}

type ruleStatusPtr string

func RuleStatusPtr(v string) RuleStatusPtrInput {
	return (*ruleStatusPtr)(&v)
}

func (*ruleStatusPtr) ElementType() reflect.Type {
	return ruleStatusPtrType
}

func (in *ruleStatusPtr) ToRuleStatusPtrOutput() RuleStatusPtrOutput {
	return pulumi.ToOutput(in).(RuleStatusPtrOutput)
}

func (in *ruleStatusPtr) ToRuleStatusPtrOutputWithContext(ctx context.Context) RuleStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RuleStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CostAllocationPolicyTypeOutput{})
	pulumi.RegisterOutputType(CostAllocationPolicyTypePtrOutput{})
	pulumi.RegisterOutputType(CostAllocationResourceTypeOutput{})
	pulumi.RegisterOutputType(CostAllocationResourceTypePtrOutput{})
	pulumi.RegisterOutputType(RuleStatusOutput{})
	pulumi.RegisterOutputType(RuleStatusPtrOutput{})
}
