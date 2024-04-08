// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The mode that is used to deploy resources. This value can be either Incremental or Complete. In Incremental mode, resources are deployed without deleting existing resources that are not included in the template. In Complete mode, resources are deployed and existing resources in the resource group that are not included in the template are deleted. Be careful when using Complete mode as you may unintentionally delete resources.
type DeploymentMode string

const (
	DeploymentModeIncremental = DeploymentMode("Incremental")
	DeploymentModeComplete    = DeploymentMode("Complete")
)

func (DeploymentMode) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentMode)(nil)).Elem()
}

func (e DeploymentMode) ToDeploymentModeOutput() DeploymentModeOutput {
	return pulumi.ToOutput(e).(DeploymentModeOutput)
}

func (e DeploymentMode) ToDeploymentModeOutputWithContext(ctx context.Context) DeploymentModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeploymentModeOutput)
}

func (e DeploymentMode) ToDeploymentModePtrOutput() DeploymentModePtrOutput {
	return e.ToDeploymentModePtrOutputWithContext(context.Background())
}

func (e DeploymentMode) ToDeploymentModePtrOutputWithContext(ctx context.Context) DeploymentModePtrOutput {
	return DeploymentMode(e).ToDeploymentModeOutputWithContext(ctx).ToDeploymentModePtrOutputWithContext(ctx)
}

func (e DeploymentMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeploymentMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeploymentMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeploymentMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeploymentModeOutput struct{ *pulumi.OutputState }

func (DeploymentModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentMode)(nil)).Elem()
}

func (o DeploymentModeOutput) ToDeploymentModeOutput() DeploymentModeOutput {
	return o
}

func (o DeploymentModeOutput) ToDeploymentModeOutputWithContext(ctx context.Context) DeploymentModeOutput {
	return o
}

func (o DeploymentModeOutput) ToDeploymentModePtrOutput() DeploymentModePtrOutput {
	return o.ToDeploymentModePtrOutputWithContext(context.Background())
}

func (o DeploymentModeOutput) ToDeploymentModePtrOutputWithContext(ctx context.Context) DeploymentModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentMode) *DeploymentMode {
		return &v
	}).(DeploymentModePtrOutput)
}

func (o DeploymentModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeploymentModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeploymentMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeploymentModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeploymentModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeploymentMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeploymentModePtrOutput struct{ *pulumi.OutputState }

func (DeploymentModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentMode)(nil)).Elem()
}

func (o DeploymentModePtrOutput) ToDeploymentModePtrOutput() DeploymentModePtrOutput {
	return o
}

func (o DeploymentModePtrOutput) ToDeploymentModePtrOutputWithContext(ctx context.Context) DeploymentModePtrOutput {
	return o
}

func (o DeploymentModePtrOutput) Elem() DeploymentModeOutput {
	return o.ApplyT(func(v *DeploymentMode) DeploymentMode {
		if v != nil {
			return *v
		}
		var ret DeploymentMode
		return ret
	}).(DeploymentModeOutput)
}

func (o DeploymentModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeploymentModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeploymentMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DeploymentModeInput is an input type that accepts values of the DeploymentMode enum
// A concrete instance of `DeploymentModeInput` can be one of the following:
//
//	DeploymentModeIncremental
//	DeploymentModeComplete
type DeploymentModeInput interface {
	pulumi.Input

	ToDeploymentModeOutput() DeploymentModeOutput
	ToDeploymentModeOutputWithContext(context.Context) DeploymentModeOutput
}

var deploymentModePtrType = reflect.TypeOf((**DeploymentMode)(nil)).Elem()

type DeploymentModePtrInput interface {
	pulumi.Input

	ToDeploymentModePtrOutput() DeploymentModePtrOutput
	ToDeploymentModePtrOutputWithContext(context.Context) DeploymentModePtrOutput
}

type deploymentModePtr string

func DeploymentModePtr(v string) DeploymentModePtrInput {
	return (*deploymentModePtr)(&v)
}

func (*deploymentModePtr) ElementType() reflect.Type {
	return deploymentModePtrType
}

func (in *deploymentModePtr) ToDeploymentModePtrOutput() DeploymentModePtrOutput {
	return pulumi.ToOutput(in).(DeploymentModePtrOutput)
}

func (in *deploymentModePtr) ToDeploymentModePtrOutputWithContext(ctx context.Context) DeploymentModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeploymentModePtrOutput)
}

// The scope to be used for evaluation of parameters, variables and functions in a nested template.
type ExpressionEvaluationOptionsScopeType string

const (
	ExpressionEvaluationOptionsScopeTypeNotSpecified = ExpressionEvaluationOptionsScopeType("NotSpecified")
	ExpressionEvaluationOptionsScopeTypeOuter        = ExpressionEvaluationOptionsScopeType("Outer")
	ExpressionEvaluationOptionsScopeTypeInner        = ExpressionEvaluationOptionsScopeType("Inner")
)

func (ExpressionEvaluationOptionsScopeType) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressionEvaluationOptionsScopeType)(nil)).Elem()
}

func (e ExpressionEvaluationOptionsScopeType) ToExpressionEvaluationOptionsScopeTypeOutput() ExpressionEvaluationOptionsScopeTypeOutput {
	return pulumi.ToOutput(e).(ExpressionEvaluationOptionsScopeTypeOutput)
}

func (e ExpressionEvaluationOptionsScopeType) ToExpressionEvaluationOptionsScopeTypeOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsScopeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpressionEvaluationOptionsScopeTypeOutput)
}

func (e ExpressionEvaluationOptionsScopeType) ToExpressionEvaluationOptionsScopeTypePtrOutput() ExpressionEvaluationOptionsScopeTypePtrOutput {
	return e.ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(context.Background())
}

func (e ExpressionEvaluationOptionsScopeType) ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsScopeTypePtrOutput {
	return ExpressionEvaluationOptionsScopeType(e).ToExpressionEvaluationOptionsScopeTypeOutputWithContext(ctx).ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(ctx)
}

func (e ExpressionEvaluationOptionsScopeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressionEvaluationOptionsScopeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressionEvaluationOptionsScopeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpressionEvaluationOptionsScopeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpressionEvaluationOptionsScopeTypeOutput struct{ *pulumi.OutputState }

func (ExpressionEvaluationOptionsScopeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressionEvaluationOptionsScopeType)(nil)).Elem()
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToExpressionEvaluationOptionsScopeTypeOutput() ExpressionEvaluationOptionsScopeTypeOutput {
	return o
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToExpressionEvaluationOptionsScopeTypeOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsScopeTypeOutput {
	return o
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToExpressionEvaluationOptionsScopeTypePtrOutput() ExpressionEvaluationOptionsScopeTypePtrOutput {
	return o.ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(context.Background())
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsScopeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressionEvaluationOptionsScopeType) *ExpressionEvaluationOptionsScopeType {
		return &v
	}).(ExpressionEvaluationOptionsScopeTypePtrOutput)
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressionEvaluationOptionsScopeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressionEvaluationOptionsScopeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpressionEvaluationOptionsScopeTypePtrOutput struct{ *pulumi.OutputState }

func (ExpressionEvaluationOptionsScopeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressionEvaluationOptionsScopeType)(nil)).Elem()
}

func (o ExpressionEvaluationOptionsScopeTypePtrOutput) ToExpressionEvaluationOptionsScopeTypePtrOutput() ExpressionEvaluationOptionsScopeTypePtrOutput {
	return o
}

func (o ExpressionEvaluationOptionsScopeTypePtrOutput) ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsScopeTypePtrOutput {
	return o
}

func (o ExpressionEvaluationOptionsScopeTypePtrOutput) Elem() ExpressionEvaluationOptionsScopeTypeOutput {
	return o.ApplyT(func(v *ExpressionEvaluationOptionsScopeType) ExpressionEvaluationOptionsScopeType {
		if v != nil {
			return *v
		}
		var ret ExpressionEvaluationOptionsScopeType
		return ret
	}).(ExpressionEvaluationOptionsScopeTypeOutput)
}

func (o ExpressionEvaluationOptionsScopeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressionEvaluationOptionsScopeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpressionEvaluationOptionsScopeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ExpressionEvaluationOptionsScopeTypeInput is an input type that accepts values of the ExpressionEvaluationOptionsScopeType enum
// A concrete instance of `ExpressionEvaluationOptionsScopeTypeInput` can be one of the following:
//
//	ExpressionEvaluationOptionsScopeTypeNotSpecified
//	ExpressionEvaluationOptionsScopeTypeOuter
//	ExpressionEvaluationOptionsScopeTypeInner
type ExpressionEvaluationOptionsScopeTypeInput interface {
	pulumi.Input

	ToExpressionEvaluationOptionsScopeTypeOutput() ExpressionEvaluationOptionsScopeTypeOutput
	ToExpressionEvaluationOptionsScopeTypeOutputWithContext(context.Context) ExpressionEvaluationOptionsScopeTypeOutput
}

var expressionEvaluationOptionsScopeTypePtrType = reflect.TypeOf((**ExpressionEvaluationOptionsScopeType)(nil)).Elem()

type ExpressionEvaluationOptionsScopeTypePtrInput interface {
	pulumi.Input

	ToExpressionEvaluationOptionsScopeTypePtrOutput() ExpressionEvaluationOptionsScopeTypePtrOutput
	ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(context.Context) ExpressionEvaluationOptionsScopeTypePtrOutput
}

type expressionEvaluationOptionsScopeTypePtr string

func ExpressionEvaluationOptionsScopeTypePtr(v string) ExpressionEvaluationOptionsScopeTypePtrInput {
	return (*expressionEvaluationOptionsScopeTypePtr)(&v)
}

func (*expressionEvaluationOptionsScopeTypePtr) ElementType() reflect.Type {
	return expressionEvaluationOptionsScopeTypePtrType
}

func (in *expressionEvaluationOptionsScopeTypePtr) ToExpressionEvaluationOptionsScopeTypePtrOutput() ExpressionEvaluationOptionsScopeTypePtrOutput {
	return pulumi.ToOutput(in).(ExpressionEvaluationOptionsScopeTypePtrOutput)
}

func (in *expressionEvaluationOptionsScopeTypePtr) ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsScopeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpressionEvaluationOptionsScopeTypePtrOutput)
}

// The extended location type.
type ExtendedLocationType string

const (
	ExtendedLocationTypeEdgeZone = ExtendedLocationType("EdgeZone")
)

func (ExtendedLocationType) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationType)(nil)).Elem()
}

func (e ExtendedLocationType) ToExtendedLocationTypeOutput() ExtendedLocationTypeOutput {
	return pulumi.ToOutput(e).(ExtendedLocationTypeOutput)
}

func (e ExtendedLocationType) ToExtendedLocationTypeOutputWithContext(ctx context.Context) ExtendedLocationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExtendedLocationTypeOutput)
}

func (e ExtendedLocationType) ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput {
	return e.ToExtendedLocationTypePtrOutputWithContext(context.Background())
}

func (e ExtendedLocationType) ToExtendedLocationTypePtrOutputWithContext(ctx context.Context) ExtendedLocationTypePtrOutput {
	return ExtendedLocationType(e).ToExtendedLocationTypeOutputWithContext(ctx).ToExtendedLocationTypePtrOutputWithContext(ctx)
}

func (e ExtendedLocationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExtendedLocationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExtendedLocationTypeOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationType)(nil)).Elem()
}

func (o ExtendedLocationTypeOutput) ToExtendedLocationTypeOutput() ExtendedLocationTypeOutput {
	return o
}

func (o ExtendedLocationTypeOutput) ToExtendedLocationTypeOutputWithContext(ctx context.Context) ExtendedLocationTypeOutput {
	return o
}

func (o ExtendedLocationTypeOutput) ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput {
	return o.ToExtendedLocationTypePtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypeOutput) ToExtendedLocationTypePtrOutputWithContext(ctx context.Context) ExtendedLocationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocationType) *ExtendedLocationType {
		return &v
	}).(ExtendedLocationTypePtrOutput)
}

func (o ExtendedLocationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExtendedLocationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExtendedLocationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationTypePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationType)(nil)).Elem()
}

func (o ExtendedLocationTypePtrOutput) ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput {
	return o
}

func (o ExtendedLocationTypePtrOutput) ToExtendedLocationTypePtrOutputWithContext(ctx context.Context) ExtendedLocationTypePtrOutput {
	return o
}

func (o ExtendedLocationTypePtrOutput) Elem() ExtendedLocationTypeOutput {
	return o.ApplyT(func(v *ExtendedLocationType) ExtendedLocationType {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationType
		return ret
	}).(ExtendedLocationTypeOutput)
}

func (o ExtendedLocationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExtendedLocationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ExtendedLocationTypeInput is an input type that accepts values of the ExtendedLocationType enum
// A concrete instance of `ExtendedLocationTypeInput` can be one of the following:
//
//	ExtendedLocationTypeEdgeZone
type ExtendedLocationTypeInput interface {
	pulumi.Input

	ToExtendedLocationTypeOutput() ExtendedLocationTypeOutput
	ToExtendedLocationTypeOutputWithContext(context.Context) ExtendedLocationTypeOutput
}

var extendedLocationTypePtrType = reflect.TypeOf((**ExtendedLocationType)(nil)).Elem()

type ExtendedLocationTypePtrInput interface {
	pulumi.Input

	ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput
	ToExtendedLocationTypePtrOutputWithContext(context.Context) ExtendedLocationTypePtrOutput
}

type extendedLocationTypePtr string

func ExtendedLocationTypePtr(v string) ExtendedLocationTypePtrInput {
	return (*extendedLocationTypePtr)(&v)
}

func (*extendedLocationTypePtr) ElementType() reflect.Type {
	return extendedLocationTypePtrType
}

func (in *extendedLocationTypePtr) ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput {
	return pulumi.ToOutput(in).(ExtendedLocationTypePtrOutput)
}

func (in *extendedLocationTypePtr) ToExtendedLocationTypePtrOutputWithContext(ctx context.Context) ExtendedLocationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExtendedLocationTypePtrOutput)
}

// The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.
type OnErrorDeploymentType string

const (
	OnErrorDeploymentTypeLastSuccessful     = OnErrorDeploymentType("LastSuccessful")
	OnErrorDeploymentTypeSpecificDeployment = OnErrorDeploymentType("SpecificDeployment")
)

func (OnErrorDeploymentType) ElementType() reflect.Type {
	return reflect.TypeOf((*OnErrorDeploymentType)(nil)).Elem()
}

func (e OnErrorDeploymentType) ToOnErrorDeploymentTypeOutput() OnErrorDeploymentTypeOutput {
	return pulumi.ToOutput(e).(OnErrorDeploymentTypeOutput)
}

func (e OnErrorDeploymentType) ToOnErrorDeploymentTypeOutputWithContext(ctx context.Context) OnErrorDeploymentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OnErrorDeploymentTypeOutput)
}

func (e OnErrorDeploymentType) ToOnErrorDeploymentTypePtrOutput() OnErrorDeploymentTypePtrOutput {
	return e.ToOnErrorDeploymentTypePtrOutputWithContext(context.Background())
}

func (e OnErrorDeploymentType) ToOnErrorDeploymentTypePtrOutputWithContext(ctx context.Context) OnErrorDeploymentTypePtrOutput {
	return OnErrorDeploymentType(e).ToOnErrorDeploymentTypeOutputWithContext(ctx).ToOnErrorDeploymentTypePtrOutputWithContext(ctx)
}

func (e OnErrorDeploymentType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OnErrorDeploymentType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OnErrorDeploymentType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OnErrorDeploymentType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OnErrorDeploymentTypeOutput struct{ *pulumi.OutputState }

func (OnErrorDeploymentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnErrorDeploymentType)(nil)).Elem()
}

func (o OnErrorDeploymentTypeOutput) ToOnErrorDeploymentTypeOutput() OnErrorDeploymentTypeOutput {
	return o
}

func (o OnErrorDeploymentTypeOutput) ToOnErrorDeploymentTypeOutputWithContext(ctx context.Context) OnErrorDeploymentTypeOutput {
	return o
}

func (o OnErrorDeploymentTypeOutput) ToOnErrorDeploymentTypePtrOutput() OnErrorDeploymentTypePtrOutput {
	return o.ToOnErrorDeploymentTypePtrOutputWithContext(context.Background())
}

func (o OnErrorDeploymentTypeOutput) ToOnErrorDeploymentTypePtrOutputWithContext(ctx context.Context) OnErrorDeploymentTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OnErrorDeploymentType) *OnErrorDeploymentType {
		return &v
	}).(OnErrorDeploymentTypePtrOutput)
}

func (o OnErrorDeploymentTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OnErrorDeploymentTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OnErrorDeploymentType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OnErrorDeploymentTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OnErrorDeploymentTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OnErrorDeploymentType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OnErrorDeploymentTypePtrOutput struct{ *pulumi.OutputState }

func (OnErrorDeploymentTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnErrorDeploymentType)(nil)).Elem()
}

func (o OnErrorDeploymentTypePtrOutput) ToOnErrorDeploymentTypePtrOutput() OnErrorDeploymentTypePtrOutput {
	return o
}

func (o OnErrorDeploymentTypePtrOutput) ToOnErrorDeploymentTypePtrOutputWithContext(ctx context.Context) OnErrorDeploymentTypePtrOutput {
	return o
}

func (o OnErrorDeploymentTypePtrOutput) Elem() OnErrorDeploymentTypeOutput {
	return o.ApplyT(func(v *OnErrorDeploymentType) OnErrorDeploymentType {
		if v != nil {
			return *v
		}
		var ret OnErrorDeploymentType
		return ret
	}).(OnErrorDeploymentTypeOutput)
}

func (o OnErrorDeploymentTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OnErrorDeploymentTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OnErrorDeploymentType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// OnErrorDeploymentTypeInput is an input type that accepts values of the OnErrorDeploymentType enum
// A concrete instance of `OnErrorDeploymentTypeInput` can be one of the following:
//
//	OnErrorDeploymentTypeLastSuccessful
//	OnErrorDeploymentTypeSpecificDeployment
type OnErrorDeploymentTypeInput interface {
	pulumi.Input

	ToOnErrorDeploymentTypeOutput() OnErrorDeploymentTypeOutput
	ToOnErrorDeploymentTypeOutputWithContext(context.Context) OnErrorDeploymentTypeOutput
}

var onErrorDeploymentTypePtrType = reflect.TypeOf((**OnErrorDeploymentType)(nil)).Elem()

type OnErrorDeploymentTypePtrInput interface {
	pulumi.Input

	ToOnErrorDeploymentTypePtrOutput() OnErrorDeploymentTypePtrOutput
	ToOnErrorDeploymentTypePtrOutputWithContext(context.Context) OnErrorDeploymentTypePtrOutput
}

type onErrorDeploymentTypePtr string

func OnErrorDeploymentTypePtr(v string) OnErrorDeploymentTypePtrInput {
	return (*onErrorDeploymentTypePtr)(&v)
}

func (*onErrorDeploymentTypePtr) ElementType() reflect.Type {
	return onErrorDeploymentTypePtrType
}

func (in *onErrorDeploymentTypePtr) ToOnErrorDeploymentTypePtrOutput() OnErrorDeploymentTypePtrOutput {
	return pulumi.ToOutput(in).(OnErrorDeploymentTypePtrOutput)
}

func (in *onErrorDeploymentTypePtr) ToOnErrorDeploymentTypePtrOutputWithContext(ctx context.Context) OnErrorDeploymentTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OnErrorDeploymentTypePtrOutput)
}

// The identity type.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
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
//	ResourceIdentityType_SystemAssigned_UserAssigned
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

func init() {
	pulumi.RegisterOutputType(DeploymentModeOutput{})
	pulumi.RegisterOutputType(DeploymentModePtrOutput{})
	pulumi.RegisterOutputType(ExpressionEvaluationOptionsScopeTypeOutput{})
	pulumi.RegisterOutputType(ExpressionEvaluationOptionsScopeTypePtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypeOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypePtrOutput{})
	pulumi.RegisterOutputType(OnErrorDeploymentTypeOutput{})
	pulumi.RegisterOutputType(OnErrorDeploymentTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
