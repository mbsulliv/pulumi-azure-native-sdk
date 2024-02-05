// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220510preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies the assessment mode.
type AssessmentModeTypes string

const (
	AssessmentModeTypesImageDefault        = AssessmentModeTypes("ImageDefault")
	AssessmentModeTypesAutomaticByPlatform = AssessmentModeTypes("AutomaticByPlatform")
)

func (AssessmentModeTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentModeTypes)(nil)).Elem()
}

func (e AssessmentModeTypes) ToAssessmentModeTypesOutput() AssessmentModeTypesOutput {
	return pulumi.ToOutput(e).(AssessmentModeTypesOutput)
}

func (e AssessmentModeTypes) ToAssessmentModeTypesOutputWithContext(ctx context.Context) AssessmentModeTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssessmentModeTypesOutput)
}

func (e AssessmentModeTypes) ToAssessmentModeTypesPtrOutput() AssessmentModeTypesPtrOutput {
	return e.ToAssessmentModeTypesPtrOutputWithContext(context.Background())
}

func (e AssessmentModeTypes) ToAssessmentModeTypesPtrOutputWithContext(ctx context.Context) AssessmentModeTypesPtrOutput {
	return AssessmentModeTypes(e).ToAssessmentModeTypesOutputWithContext(ctx).ToAssessmentModeTypesPtrOutputWithContext(ctx)
}

func (e AssessmentModeTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentModeTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentModeTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssessmentModeTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssessmentModeTypesOutput struct{ *pulumi.OutputState }

func (AssessmentModeTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentModeTypes)(nil)).Elem()
}

func (o AssessmentModeTypesOutput) ToAssessmentModeTypesOutput() AssessmentModeTypesOutput {
	return o
}

func (o AssessmentModeTypesOutput) ToAssessmentModeTypesOutputWithContext(ctx context.Context) AssessmentModeTypesOutput {
	return o
}

func (o AssessmentModeTypesOutput) ToAssessmentModeTypesPtrOutput() AssessmentModeTypesPtrOutput {
	return o.ToAssessmentModeTypesPtrOutputWithContext(context.Background())
}

func (o AssessmentModeTypesOutput) ToAssessmentModeTypesPtrOutputWithContext(ctx context.Context) AssessmentModeTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentModeTypes) *AssessmentModeTypes {
		return &v
	}).(AssessmentModeTypesPtrOutput)
}

func (o AssessmentModeTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssessmentModeTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentModeTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssessmentModeTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentModeTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentModeTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssessmentModeTypesPtrOutput struct{ *pulumi.OutputState }

func (AssessmentModeTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentModeTypes)(nil)).Elem()
}

func (o AssessmentModeTypesPtrOutput) ToAssessmentModeTypesPtrOutput() AssessmentModeTypesPtrOutput {
	return o
}

func (o AssessmentModeTypesPtrOutput) ToAssessmentModeTypesPtrOutputWithContext(ctx context.Context) AssessmentModeTypesPtrOutput {
	return o
}

func (o AssessmentModeTypesPtrOutput) Elem() AssessmentModeTypesOutput {
	return o.ApplyT(func(v *AssessmentModeTypes) AssessmentModeTypes {
		if v != nil {
			return *v
		}
		var ret AssessmentModeTypes
		return ret
	}).(AssessmentModeTypesOutput)
}

func (o AssessmentModeTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentModeTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssessmentModeTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AssessmentModeTypesInput is an input type that accepts values of the AssessmentModeTypes enum
// A concrete instance of `AssessmentModeTypesInput` can be one of the following:
//
//	AssessmentModeTypesImageDefault
//	AssessmentModeTypesAutomaticByPlatform
type AssessmentModeTypesInput interface {
	pulumi.Input

	ToAssessmentModeTypesOutput() AssessmentModeTypesOutput
	ToAssessmentModeTypesOutputWithContext(context.Context) AssessmentModeTypesOutput
}

var assessmentModeTypesPtrType = reflect.TypeOf((**AssessmentModeTypes)(nil)).Elem()

type AssessmentModeTypesPtrInput interface {
	pulumi.Input

	ToAssessmentModeTypesPtrOutput() AssessmentModeTypesPtrOutput
	ToAssessmentModeTypesPtrOutputWithContext(context.Context) AssessmentModeTypesPtrOutput
}

type assessmentModeTypesPtr string

func AssessmentModeTypesPtr(v string) AssessmentModeTypesPtrInput {
	return (*assessmentModeTypesPtr)(&v)
}

func (*assessmentModeTypesPtr) ElementType() reflect.Type {
	return assessmentModeTypesPtrType
}

func (in *assessmentModeTypesPtr) ToAssessmentModeTypesPtrOutput() AssessmentModeTypesPtrOutput {
	return pulumi.ToOutput(in).(AssessmentModeTypesPtrOutput)
}

func (in *assessmentModeTypesPtr) ToAssessmentModeTypesPtrOutputWithContext(ctx context.Context) AssessmentModeTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssessmentModeTypesPtrOutput)
}

// Specifies the patch mode.
type PatchModeTypes string

const (
	PatchModeTypesImageDefault        = PatchModeTypes("ImageDefault")
	PatchModeTypesAutomaticByPlatform = PatchModeTypes("AutomaticByPlatform")
	PatchModeTypesAutomaticByOS       = PatchModeTypes("AutomaticByOS")
	PatchModeTypesManual              = PatchModeTypes("Manual")
)

func (PatchModeTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*PatchModeTypes)(nil)).Elem()
}

func (e PatchModeTypes) ToPatchModeTypesOutput() PatchModeTypesOutput {
	return pulumi.ToOutput(e).(PatchModeTypesOutput)
}

func (e PatchModeTypes) ToPatchModeTypesOutputWithContext(ctx context.Context) PatchModeTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PatchModeTypesOutput)
}

func (e PatchModeTypes) ToPatchModeTypesPtrOutput() PatchModeTypesPtrOutput {
	return e.ToPatchModeTypesPtrOutputWithContext(context.Background())
}

func (e PatchModeTypes) ToPatchModeTypesPtrOutputWithContext(ctx context.Context) PatchModeTypesPtrOutput {
	return PatchModeTypes(e).ToPatchModeTypesOutputWithContext(ctx).ToPatchModeTypesPtrOutputWithContext(ctx)
}

func (e PatchModeTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PatchModeTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PatchModeTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PatchModeTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PatchModeTypesOutput struct{ *pulumi.OutputState }

func (PatchModeTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PatchModeTypes)(nil)).Elem()
}

func (o PatchModeTypesOutput) ToPatchModeTypesOutput() PatchModeTypesOutput {
	return o
}

func (o PatchModeTypesOutput) ToPatchModeTypesOutputWithContext(ctx context.Context) PatchModeTypesOutput {
	return o
}

func (o PatchModeTypesOutput) ToPatchModeTypesPtrOutput() PatchModeTypesPtrOutput {
	return o.ToPatchModeTypesPtrOutputWithContext(context.Background())
}

func (o PatchModeTypesOutput) ToPatchModeTypesPtrOutputWithContext(ctx context.Context) PatchModeTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PatchModeTypes) *PatchModeTypes {
		return &v
	}).(PatchModeTypesPtrOutput)
}

func (o PatchModeTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PatchModeTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PatchModeTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PatchModeTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PatchModeTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PatchModeTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PatchModeTypesPtrOutput struct{ *pulumi.OutputState }

func (PatchModeTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchModeTypes)(nil)).Elem()
}

func (o PatchModeTypesPtrOutput) ToPatchModeTypesPtrOutput() PatchModeTypesPtrOutput {
	return o
}

func (o PatchModeTypesPtrOutput) ToPatchModeTypesPtrOutputWithContext(ctx context.Context) PatchModeTypesPtrOutput {
	return o
}

func (o PatchModeTypesPtrOutput) Elem() PatchModeTypesOutput {
	return o.ApplyT(func(v *PatchModeTypes) PatchModeTypes {
		if v != nil {
			return *v
		}
		var ret PatchModeTypes
		return ret
	}).(PatchModeTypesOutput)
}

func (o PatchModeTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PatchModeTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PatchModeTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PatchModeTypesInput is an input type that accepts values of the PatchModeTypes enum
// A concrete instance of `PatchModeTypesInput` can be one of the following:
//
//	PatchModeTypesImageDefault
//	PatchModeTypesAutomaticByPlatform
//	PatchModeTypesAutomaticByOS
//	PatchModeTypesManual
type PatchModeTypesInput interface {
	pulumi.Input

	ToPatchModeTypesOutput() PatchModeTypesOutput
	ToPatchModeTypesOutputWithContext(context.Context) PatchModeTypesOutput
}

var patchModeTypesPtrType = reflect.TypeOf((**PatchModeTypes)(nil)).Elem()

type PatchModeTypesPtrInput interface {
	pulumi.Input

	ToPatchModeTypesPtrOutput() PatchModeTypesPtrOutput
	ToPatchModeTypesPtrOutputWithContext(context.Context) PatchModeTypesPtrOutput
}

type patchModeTypesPtr string

func PatchModeTypesPtr(v string) PatchModeTypesPtrInput {
	return (*patchModeTypesPtr)(&v)
}

func (*patchModeTypesPtr) ElementType() reflect.Type {
	return patchModeTypesPtrType
}

func (in *patchModeTypesPtr) ToPatchModeTypesPtrOutput() PatchModeTypesPtrOutput {
	return pulumi.ToOutput(in).(PatchModeTypesPtrOutput)
}

func (in *patchModeTypesPtr) ToPatchModeTypesPtrOutputWithContext(ctx context.Context) PatchModeTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PatchModeTypesPtrOutput)
}

// The identity type.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
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

// The level code.
type StatusLevelTypes string

const (
	StatusLevelTypesInfo    = StatusLevelTypes("Info")
	StatusLevelTypesWarning = StatusLevelTypes("Warning")
	StatusLevelTypesError   = StatusLevelTypes("Error")
)

func (StatusLevelTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusLevelTypes)(nil)).Elem()
}

func (e StatusLevelTypes) ToStatusLevelTypesOutput() StatusLevelTypesOutput {
	return pulumi.ToOutput(e).(StatusLevelTypesOutput)
}

func (e StatusLevelTypes) ToStatusLevelTypesOutputWithContext(ctx context.Context) StatusLevelTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StatusLevelTypesOutput)
}

func (e StatusLevelTypes) ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput {
	return e.ToStatusLevelTypesPtrOutputWithContext(context.Background())
}

func (e StatusLevelTypes) ToStatusLevelTypesPtrOutputWithContext(ctx context.Context) StatusLevelTypesPtrOutput {
	return StatusLevelTypes(e).ToStatusLevelTypesOutputWithContext(ctx).ToStatusLevelTypesPtrOutputWithContext(ctx)
}

func (e StatusLevelTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusLevelTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusLevelTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StatusLevelTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StatusLevelTypesOutput struct{ *pulumi.OutputState }

func (StatusLevelTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusLevelTypes)(nil)).Elem()
}

func (o StatusLevelTypesOutput) ToStatusLevelTypesOutput() StatusLevelTypesOutput {
	return o
}

func (o StatusLevelTypesOutput) ToStatusLevelTypesOutputWithContext(ctx context.Context) StatusLevelTypesOutput {
	return o
}

func (o StatusLevelTypesOutput) ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput {
	return o.ToStatusLevelTypesPtrOutputWithContext(context.Background())
}

func (o StatusLevelTypesOutput) ToStatusLevelTypesPtrOutputWithContext(ctx context.Context) StatusLevelTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StatusLevelTypes) *StatusLevelTypes {
		return &v
	}).(StatusLevelTypesPtrOutput)
}

func (o StatusLevelTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StatusLevelTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusLevelTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StatusLevelTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusLevelTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusLevelTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatusLevelTypesPtrOutput struct{ *pulumi.OutputState }

func (StatusLevelTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StatusLevelTypes)(nil)).Elem()
}

func (o StatusLevelTypesPtrOutput) ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput {
	return o
}

func (o StatusLevelTypesPtrOutput) ToStatusLevelTypesPtrOutputWithContext(ctx context.Context) StatusLevelTypesPtrOutput {
	return o
}

func (o StatusLevelTypesPtrOutput) Elem() StatusLevelTypesOutput {
	return o.ApplyT(func(v *StatusLevelTypes) StatusLevelTypes {
		if v != nil {
			return *v
		}
		var ret StatusLevelTypes
		return ret
	}).(StatusLevelTypesOutput)
}

func (o StatusLevelTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusLevelTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StatusLevelTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// StatusLevelTypesInput is an input type that accepts values of the StatusLevelTypes enum
// A concrete instance of `StatusLevelTypesInput` can be one of the following:
//
//	StatusLevelTypesInfo
//	StatusLevelTypesWarning
//	StatusLevelTypesError
type StatusLevelTypesInput interface {
	pulumi.Input

	ToStatusLevelTypesOutput() StatusLevelTypesOutput
	ToStatusLevelTypesOutputWithContext(context.Context) StatusLevelTypesOutput
}

var statusLevelTypesPtrType = reflect.TypeOf((**StatusLevelTypes)(nil)).Elem()

type StatusLevelTypesPtrInput interface {
	pulumi.Input

	ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput
	ToStatusLevelTypesPtrOutputWithContext(context.Context) StatusLevelTypesPtrOutput
}

type statusLevelTypesPtr string

func StatusLevelTypesPtr(v string) StatusLevelTypesPtrInput {
	return (*statusLevelTypesPtr)(&v)
}

func (*statusLevelTypesPtr) ElementType() reflect.Type {
	return statusLevelTypesPtrType
}

func (in *statusLevelTypesPtr) ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput {
	return pulumi.ToOutput(in).(StatusLevelTypesPtrOutput)
}

func (in *statusLevelTypesPtr) ToStatusLevelTypesPtrOutputWithContext(ctx context.Context) StatusLevelTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatusLevelTypesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentModeTypesOutput{})
	pulumi.RegisterOutputType(AssessmentModeTypesPtrOutput{})
	pulumi.RegisterOutputType(PatchModeTypesOutput{})
	pulumi.RegisterOutputType(PatchModeTypesPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(StatusLevelTypesOutput{})
	pulumi.RegisterOutputType(StatusLevelTypesPtrOutput{})
}
