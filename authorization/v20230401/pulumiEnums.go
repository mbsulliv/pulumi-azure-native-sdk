// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
type EnforcementMode string

const (
	// The policy effect is enforced during resource creation or update.
	EnforcementModeDefault = EnforcementMode("Default")
	// The policy effect is not enforced during resource creation or update.
	EnforcementModeDoNotEnforce = EnforcementMode("DoNotEnforce")
)

func (EnforcementMode) ElementType() reflect.Type {
	return reflect.TypeOf((*EnforcementMode)(nil)).Elem()
}

func (e EnforcementMode) ToEnforcementModeOutput() EnforcementModeOutput {
	return pulumi.ToOutput(e).(EnforcementModeOutput)
}

func (e EnforcementMode) ToEnforcementModeOutputWithContext(ctx context.Context) EnforcementModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnforcementModeOutput)
}

func (e EnforcementMode) ToEnforcementModePtrOutput() EnforcementModePtrOutput {
	return e.ToEnforcementModePtrOutputWithContext(context.Background())
}

func (e EnforcementMode) ToEnforcementModePtrOutputWithContext(ctx context.Context) EnforcementModePtrOutput {
	return EnforcementMode(e).ToEnforcementModeOutputWithContext(ctx).ToEnforcementModePtrOutputWithContext(ctx)
}

func (e EnforcementMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnforcementMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnforcementMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnforcementMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnforcementModeOutput struct{ *pulumi.OutputState }

func (EnforcementModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnforcementMode)(nil)).Elem()
}

func (o EnforcementModeOutput) ToEnforcementModeOutput() EnforcementModeOutput {
	return o
}

func (o EnforcementModeOutput) ToEnforcementModeOutputWithContext(ctx context.Context) EnforcementModeOutput {
	return o
}

func (o EnforcementModeOutput) ToEnforcementModePtrOutput() EnforcementModePtrOutput {
	return o.ToEnforcementModePtrOutputWithContext(context.Background())
}

func (o EnforcementModeOutput) ToEnforcementModePtrOutputWithContext(ctx context.Context) EnforcementModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnforcementMode) *EnforcementMode {
		return &v
	}).(EnforcementModePtrOutput)
}

func (o EnforcementModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnforcementModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnforcementMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnforcementModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnforcementModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnforcementMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnforcementModePtrOutput struct{ *pulumi.OutputState }

func (EnforcementModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnforcementMode)(nil)).Elem()
}

func (o EnforcementModePtrOutput) ToEnforcementModePtrOutput() EnforcementModePtrOutput {
	return o
}

func (o EnforcementModePtrOutput) ToEnforcementModePtrOutputWithContext(ctx context.Context) EnforcementModePtrOutput {
	return o
}

func (o EnforcementModePtrOutput) Elem() EnforcementModeOutput {
	return o.ApplyT(func(v *EnforcementMode) EnforcementMode {
		if v != nil {
			return *v
		}
		var ret EnforcementMode
		return ret
	}).(EnforcementModeOutput)
}

func (o EnforcementModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnforcementModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnforcementMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EnforcementModeInput is an input type that accepts values of the EnforcementMode enum
// A concrete instance of `EnforcementModeInput` can be one of the following:
//
//	EnforcementModeDefault
//	EnforcementModeDoNotEnforce
type EnforcementModeInput interface {
	pulumi.Input

	ToEnforcementModeOutput() EnforcementModeOutput
	ToEnforcementModeOutputWithContext(context.Context) EnforcementModeOutput
}

var enforcementModePtrType = reflect.TypeOf((**EnforcementMode)(nil)).Elem()

type EnforcementModePtrInput interface {
	pulumi.Input

	ToEnforcementModePtrOutput() EnforcementModePtrOutput
	ToEnforcementModePtrOutputWithContext(context.Context) EnforcementModePtrOutput
}

type enforcementModePtr string

func EnforcementModePtr(v string) EnforcementModePtrInput {
	return (*enforcementModePtr)(&v)
}

func (*enforcementModePtr) ElementType() reflect.Type {
	return enforcementModePtrType
}

func (in *enforcementModePtr) ToEnforcementModePtrOutput() EnforcementModePtrOutput {
	return pulumi.ToOutput(in).(EnforcementModePtrOutput)
}

func (in *enforcementModePtr) ToEnforcementModePtrOutputWithContext(ctx context.Context) EnforcementModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnforcementModePtrOutput)
}

// The override kind.
type OverrideKind string

const (
	// It will override the policy effect type.
	OverrideKindPolicyEffect = OverrideKind("policyEffect")
)

func (OverrideKind) ElementType() reflect.Type {
	return reflect.TypeOf((*OverrideKind)(nil)).Elem()
}

func (e OverrideKind) ToOverrideKindOutput() OverrideKindOutput {
	return pulumi.ToOutput(e).(OverrideKindOutput)
}

func (e OverrideKind) ToOverrideKindOutputWithContext(ctx context.Context) OverrideKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OverrideKindOutput)
}

func (e OverrideKind) ToOverrideKindPtrOutput() OverrideKindPtrOutput {
	return e.ToOverrideKindPtrOutputWithContext(context.Background())
}

func (e OverrideKind) ToOverrideKindPtrOutputWithContext(ctx context.Context) OverrideKindPtrOutput {
	return OverrideKind(e).ToOverrideKindOutputWithContext(ctx).ToOverrideKindPtrOutputWithContext(ctx)
}

func (e OverrideKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OverrideKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OverrideKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OverrideKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OverrideKindOutput struct{ *pulumi.OutputState }

func (OverrideKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OverrideKind)(nil)).Elem()
}

func (o OverrideKindOutput) ToOverrideKindOutput() OverrideKindOutput {
	return o
}

func (o OverrideKindOutput) ToOverrideKindOutputWithContext(ctx context.Context) OverrideKindOutput {
	return o
}

func (o OverrideKindOutput) ToOverrideKindPtrOutput() OverrideKindPtrOutput {
	return o.ToOverrideKindPtrOutputWithContext(context.Background())
}

func (o OverrideKindOutput) ToOverrideKindPtrOutputWithContext(ctx context.Context) OverrideKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OverrideKind) *OverrideKind {
		return &v
	}).(OverrideKindPtrOutput)
}

func (o OverrideKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OverrideKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OverrideKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OverrideKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OverrideKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OverrideKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OverrideKindPtrOutput struct{ *pulumi.OutputState }

func (OverrideKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OverrideKind)(nil)).Elem()
}

func (o OverrideKindPtrOutput) ToOverrideKindPtrOutput() OverrideKindPtrOutput {
	return o
}

func (o OverrideKindPtrOutput) ToOverrideKindPtrOutputWithContext(ctx context.Context) OverrideKindPtrOutput {
	return o
}

func (o OverrideKindPtrOutput) Elem() OverrideKindOutput {
	return o.ApplyT(func(v *OverrideKind) OverrideKind {
		if v != nil {
			return *v
		}
		var ret OverrideKind
		return ret
	}).(OverrideKindOutput)
}

func (o OverrideKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OverrideKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OverrideKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// OverrideKindInput is an input type that accepts values of the OverrideKind enum
// A concrete instance of `OverrideKindInput` can be one of the following:
//
//	OverrideKindPolicyEffect
type OverrideKindInput interface {
	pulumi.Input

	ToOverrideKindOutput() OverrideKindOutput
	ToOverrideKindOutputWithContext(context.Context) OverrideKindOutput
}

var overrideKindPtrType = reflect.TypeOf((**OverrideKind)(nil)).Elem()

type OverrideKindPtrInput interface {
	pulumi.Input

	ToOverrideKindPtrOutput() OverrideKindPtrOutput
	ToOverrideKindPtrOutputWithContext(context.Context) OverrideKindPtrOutput
}

type overrideKindPtr string

func OverrideKindPtr(v string) OverrideKindPtrInput {
	return (*overrideKindPtr)(&v)
}

func (*overrideKindPtr) ElementType() reflect.Type {
	return overrideKindPtrType
}

func (in *overrideKindPtr) ToOverrideKindPtrOutput() OverrideKindPtrOutput {
	return pulumi.ToOutput(in).(OverrideKindPtrOutput)
}

func (in *overrideKindPtr) ToOverrideKindPtrOutputWithContext(ctx context.Context) OverrideKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OverrideKindPtrOutput)
}

// The data type of the parameter.
type ParameterType string

const (
	ParameterTypeString   = ParameterType("String")
	ParameterTypeArray    = ParameterType("Array")
	ParameterTypeObject   = ParameterType("Object")
	ParameterTypeBoolean  = ParameterType("Boolean")
	ParameterTypeInteger  = ParameterType("Integer")
	ParameterTypeFloat    = ParameterType("Float")
	ParameterTypeDateTime = ParameterType("DateTime")
)

func (ParameterType) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterType)(nil)).Elem()
}

func (e ParameterType) ToParameterTypeOutput() ParameterTypeOutput {
	return pulumi.ToOutput(e).(ParameterTypeOutput)
}

func (e ParameterType) ToParameterTypeOutputWithContext(ctx context.Context) ParameterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ParameterTypeOutput)
}

func (e ParameterType) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return e.ToParameterTypePtrOutputWithContext(context.Background())
}

func (e ParameterType) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return ParameterType(e).ToParameterTypeOutputWithContext(ctx).ToParameterTypePtrOutputWithContext(ctx)
}

func (e ParameterType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ParameterType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ParameterType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ParameterType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ParameterTypeOutput struct{ *pulumi.OutputState }

func (ParameterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterType)(nil)).Elem()
}

func (o ParameterTypeOutput) ToParameterTypeOutput() ParameterTypeOutput {
	return o
}

func (o ParameterTypeOutput) ToParameterTypeOutputWithContext(ctx context.Context) ParameterTypeOutput {
	return o
}

func (o ParameterTypeOutput) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return o.ToParameterTypePtrOutputWithContext(context.Background())
}

func (o ParameterTypeOutput) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ParameterType) *ParameterType {
		return &v
	}).(ParameterTypePtrOutput)
}

func (o ParameterTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ParameterTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ParameterType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ParameterTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ParameterTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ParameterType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ParameterTypePtrOutput struct{ *pulumi.OutputState }

func (ParameterTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterType)(nil)).Elem()
}

func (o ParameterTypePtrOutput) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return o
}

func (o ParameterTypePtrOutput) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return o
}

func (o ParameterTypePtrOutput) Elem() ParameterTypeOutput {
	return o.ApplyT(func(v *ParameterType) ParameterType {
		if v != nil {
			return *v
		}
		var ret ParameterType
		return ret
	}).(ParameterTypeOutput)
}

func (o ParameterTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ParameterTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ParameterType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ParameterTypeInput is an input type that accepts values of the ParameterType enum
// A concrete instance of `ParameterTypeInput` can be one of the following:
//
//	ParameterTypeString
//	ParameterTypeArray
//	ParameterTypeObject
//	ParameterTypeBoolean
//	ParameterTypeInteger
//	ParameterTypeFloat
//	ParameterTypeDateTime
type ParameterTypeInput interface {
	pulumi.Input

	ToParameterTypeOutput() ParameterTypeOutput
	ToParameterTypeOutputWithContext(context.Context) ParameterTypeOutput
}

var parameterTypePtrType = reflect.TypeOf((**ParameterType)(nil)).Elem()

type ParameterTypePtrInput interface {
	pulumi.Input

	ToParameterTypePtrOutput() ParameterTypePtrOutput
	ToParameterTypePtrOutputWithContext(context.Context) ParameterTypePtrOutput
}

type parameterTypePtr string

func ParameterTypePtr(v string) ParameterTypePtrInput {
	return (*parameterTypePtr)(&v)
}

func (*parameterTypePtr) ElementType() reflect.Type {
	return parameterTypePtrType
}

func (in *parameterTypePtr) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return pulumi.ToOutput(in).(ParameterTypePtrOutput)
}

func (in *parameterTypePtr) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ParameterTypePtrOutput)
}

// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
type PolicyType string

const (
	PolicyTypeNotSpecified = PolicyType("NotSpecified")
	PolicyTypeBuiltIn      = PolicyType("BuiltIn")
	PolicyTypeCustom       = PolicyType("Custom")
	PolicyTypeStatic       = PolicyType("Static")
)

func (PolicyType) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyType)(nil)).Elem()
}

func (e PolicyType) ToPolicyTypeOutput() PolicyTypeOutput {
	return pulumi.ToOutput(e).(PolicyTypeOutput)
}

func (e PolicyType) ToPolicyTypeOutputWithContext(ctx context.Context) PolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PolicyTypeOutput)
}

func (e PolicyType) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return e.ToPolicyTypePtrOutputWithContext(context.Background())
}

func (e PolicyType) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return PolicyType(e).ToPolicyTypeOutputWithContext(ctx).ToPolicyTypePtrOutputWithContext(ctx)
}

func (e PolicyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PolicyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PolicyTypeOutput struct{ *pulumi.OutputState }

func (PolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyType)(nil)).Elem()
}

func (o PolicyTypeOutput) ToPolicyTypeOutput() PolicyTypeOutput {
	return o
}

func (o PolicyTypeOutput) ToPolicyTypeOutputWithContext(ctx context.Context) PolicyTypeOutput {
	return o
}

func (o PolicyTypeOutput) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return o.ToPolicyTypePtrOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyType) *PolicyType {
		return &v
	}).(PolicyTypePtrOutput)
}

func (o PolicyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PolicyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolicyTypePtrOutput struct{ *pulumi.OutputState }

func (PolicyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyType)(nil)).Elem()
}

func (o PolicyTypePtrOutput) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return o
}

func (o PolicyTypePtrOutput) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return o
}

func (o PolicyTypePtrOutput) Elem() PolicyTypeOutput {
	return o.ApplyT(func(v *PolicyType) PolicyType {
		if v != nil {
			return *v
		}
		var ret PolicyType
		return ret
	}).(PolicyTypeOutput)
}

func (o PolicyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PolicyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PolicyTypeInput is an input type that accepts values of the PolicyType enum
// A concrete instance of `PolicyTypeInput` can be one of the following:
//
//	PolicyTypeNotSpecified
//	PolicyTypeBuiltIn
//	PolicyTypeCustom
//	PolicyTypeStatic
type PolicyTypeInput interface {
	pulumi.Input

	ToPolicyTypeOutput() PolicyTypeOutput
	ToPolicyTypeOutputWithContext(context.Context) PolicyTypeOutput
}

var policyTypePtrType = reflect.TypeOf((**PolicyType)(nil)).Elem()

type PolicyTypePtrInput interface {
	pulumi.Input

	ToPolicyTypePtrOutput() PolicyTypePtrOutput
	ToPolicyTypePtrOutputWithContext(context.Context) PolicyTypePtrOutput
}

type policyTypePtr string

func PolicyTypePtr(v string) PolicyTypePtrInput {
	return (*policyTypePtr)(&v)
}

func (*policyTypePtr) ElementType() reflect.Type {
	return policyTypePtrType
}

func (in *policyTypePtr) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return pulumi.ToOutput(in).(PolicyTypePtrOutput)
}

func (in *policyTypePtr) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PolicyTypePtrOutput)
}

// The identity type. This is the only required field when adding a system or user assigned identity to a resource.
type ResourceIdentityType string

const (
	// Indicates that a system assigned identity is associated with the resource.
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	// Indicates that a system assigned identity is associated with the resource.
	ResourceIdentityTypeUserAssigned = ResourceIdentityType("UserAssigned")
	// Indicates that no identity is associated with the resource or that the existing identity should be removed.
	ResourceIdentityTypeNone = ResourceIdentityType("None")
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

// ResourceIdentityTypeInput is an input type that accepts values of the ResourceIdentityType enum
// A concrete instance of `ResourceIdentityTypeInput` can be one of the following:
//
//	ResourceIdentityTypeSystemAssigned
//	ResourceIdentityTypeUserAssigned
//	ResourceIdentityTypeNone
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

// The selector kind.
type SelectorKind string

const (
	// The selector kind to filter policies by the resource location.
	SelectorKindResourceLocation = SelectorKind("resourceLocation")
	// The selector kind to filter policies by the resource type.
	SelectorKindResourceType = SelectorKind("resourceType")
	// The selector kind to filter policies by the resource without location.
	SelectorKindResourceWithoutLocation = SelectorKind("resourceWithoutLocation")
	// The selector kind to filter policies by the policy definition reference ID.
	SelectorKindPolicyDefinitionReferenceId = SelectorKind("policyDefinitionReferenceId")
)

func (SelectorKind) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectorKind)(nil)).Elem()
}

func (e SelectorKind) ToSelectorKindOutput() SelectorKindOutput {
	return pulumi.ToOutput(e).(SelectorKindOutput)
}

func (e SelectorKind) ToSelectorKindOutputWithContext(ctx context.Context) SelectorKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SelectorKindOutput)
}

func (e SelectorKind) ToSelectorKindPtrOutput() SelectorKindPtrOutput {
	return e.ToSelectorKindPtrOutputWithContext(context.Background())
}

func (e SelectorKind) ToSelectorKindPtrOutputWithContext(ctx context.Context) SelectorKindPtrOutput {
	return SelectorKind(e).ToSelectorKindOutputWithContext(ctx).ToSelectorKindPtrOutputWithContext(ctx)
}

func (e SelectorKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SelectorKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SelectorKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SelectorKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SelectorKindOutput struct{ *pulumi.OutputState }

func (SelectorKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectorKind)(nil)).Elem()
}

func (o SelectorKindOutput) ToSelectorKindOutput() SelectorKindOutput {
	return o
}

func (o SelectorKindOutput) ToSelectorKindOutputWithContext(ctx context.Context) SelectorKindOutput {
	return o
}

func (o SelectorKindOutput) ToSelectorKindPtrOutput() SelectorKindPtrOutput {
	return o.ToSelectorKindPtrOutputWithContext(context.Background())
}

func (o SelectorKindOutput) ToSelectorKindPtrOutputWithContext(ctx context.Context) SelectorKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SelectorKind) *SelectorKind {
		return &v
	}).(SelectorKindPtrOutput)
}

func (o SelectorKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SelectorKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SelectorKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SelectorKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SelectorKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SelectorKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SelectorKindPtrOutput struct{ *pulumi.OutputState }

func (SelectorKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SelectorKind)(nil)).Elem()
}

func (o SelectorKindPtrOutput) ToSelectorKindPtrOutput() SelectorKindPtrOutput {
	return o
}

func (o SelectorKindPtrOutput) ToSelectorKindPtrOutputWithContext(ctx context.Context) SelectorKindPtrOutput {
	return o
}

func (o SelectorKindPtrOutput) Elem() SelectorKindOutput {
	return o.ApplyT(func(v *SelectorKind) SelectorKind {
		if v != nil {
			return *v
		}
		var ret SelectorKind
		return ret
	}).(SelectorKindOutput)
}

func (o SelectorKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SelectorKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SelectorKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SelectorKindInput is an input type that accepts values of the SelectorKind enum
// A concrete instance of `SelectorKindInput` can be one of the following:
//
//	SelectorKindResourceLocation
//	SelectorKindResourceType
//	SelectorKindResourceWithoutLocation
//	SelectorKindPolicyDefinitionReferenceId
type SelectorKindInput interface {
	pulumi.Input

	ToSelectorKindOutput() SelectorKindOutput
	ToSelectorKindOutputWithContext(context.Context) SelectorKindOutput
}

var selectorKindPtrType = reflect.TypeOf((**SelectorKind)(nil)).Elem()

type SelectorKindPtrInput interface {
	pulumi.Input

	ToSelectorKindPtrOutput() SelectorKindPtrOutput
	ToSelectorKindPtrOutputWithContext(context.Context) SelectorKindPtrOutput
}

type selectorKindPtr string

func SelectorKindPtr(v string) SelectorKindPtrInput {
	return (*selectorKindPtr)(&v)
}

func (*selectorKindPtr) ElementType() reflect.Type {
	return selectorKindPtrType
}

func (in *selectorKindPtr) ToSelectorKindPtrOutput() SelectorKindPtrOutput {
	return pulumi.ToOutput(in).(SelectorKindPtrOutput)
}

func (in *selectorKindPtr) ToSelectorKindPtrOutputWithContext(ctx context.Context) SelectorKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SelectorKindPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EnforcementModeOutput{})
	pulumi.RegisterOutputType(EnforcementModePtrOutput{})
	pulumi.RegisterOutputType(OverrideKindOutput{})
	pulumi.RegisterOutputType(OverrideKindPtrOutput{})
	pulumi.RegisterOutputType(ParameterTypeOutput{})
	pulumi.RegisterOutputType(ParameterTypePtrOutput{})
	pulumi.RegisterOutputType(PolicyTypeOutput{})
	pulumi.RegisterOutputType(PolicyTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(SelectorKindOutput{})
	pulumi.RegisterOutputType(SelectorKindPtrOutput{})
}
