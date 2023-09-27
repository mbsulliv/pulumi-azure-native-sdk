// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// Represents a Template Spec artifact containing an embedded Azure Resource Manager template for use as a linked template.
type LinkedTemplateArtifact struct {
	// A filesystem safe relative path of the artifact.
	Path string `pulumi:"path"`
	// The Azure Resource Manager template.
	Template interface{} `pulumi:"template"`
}

// LinkedTemplateArtifactInput is an input type that accepts LinkedTemplateArtifactArgs and LinkedTemplateArtifactOutput values.
// You can construct a concrete instance of `LinkedTemplateArtifactInput` via:
//
//	LinkedTemplateArtifactArgs{...}
type LinkedTemplateArtifactInput interface {
	pulumi.Input

	ToLinkedTemplateArtifactOutput() LinkedTemplateArtifactOutput
	ToLinkedTemplateArtifactOutputWithContext(context.Context) LinkedTemplateArtifactOutput
}

// Represents a Template Spec artifact containing an embedded Azure Resource Manager template for use as a linked template.
type LinkedTemplateArtifactArgs struct {
	// A filesystem safe relative path of the artifact.
	Path pulumi.StringInput `pulumi:"path"`
	// The Azure Resource Manager template.
	Template pulumi.Input `pulumi:"template"`
}

func (LinkedTemplateArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedTemplateArtifact)(nil)).Elem()
}

func (i LinkedTemplateArtifactArgs) ToLinkedTemplateArtifactOutput() LinkedTemplateArtifactOutput {
	return i.ToLinkedTemplateArtifactOutputWithContext(context.Background())
}

func (i LinkedTemplateArtifactArgs) ToLinkedTemplateArtifactOutputWithContext(ctx context.Context) LinkedTemplateArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedTemplateArtifactOutput)
}

func (i LinkedTemplateArtifactArgs) ToOutput(ctx context.Context) pulumix.Output[LinkedTemplateArtifact] {
	return pulumix.Output[LinkedTemplateArtifact]{
		OutputState: i.ToLinkedTemplateArtifactOutputWithContext(ctx).OutputState,
	}
}

// LinkedTemplateArtifactArrayInput is an input type that accepts LinkedTemplateArtifactArray and LinkedTemplateArtifactArrayOutput values.
// You can construct a concrete instance of `LinkedTemplateArtifactArrayInput` via:
//
//	LinkedTemplateArtifactArray{ LinkedTemplateArtifactArgs{...} }
type LinkedTemplateArtifactArrayInput interface {
	pulumi.Input

	ToLinkedTemplateArtifactArrayOutput() LinkedTemplateArtifactArrayOutput
	ToLinkedTemplateArtifactArrayOutputWithContext(context.Context) LinkedTemplateArtifactArrayOutput
}

type LinkedTemplateArtifactArray []LinkedTemplateArtifactInput

func (LinkedTemplateArtifactArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedTemplateArtifact)(nil)).Elem()
}

func (i LinkedTemplateArtifactArray) ToLinkedTemplateArtifactArrayOutput() LinkedTemplateArtifactArrayOutput {
	return i.ToLinkedTemplateArtifactArrayOutputWithContext(context.Background())
}

func (i LinkedTemplateArtifactArray) ToLinkedTemplateArtifactArrayOutputWithContext(ctx context.Context) LinkedTemplateArtifactArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedTemplateArtifactArrayOutput)
}

func (i LinkedTemplateArtifactArray) ToOutput(ctx context.Context) pulumix.Output[[]LinkedTemplateArtifact] {
	return pulumix.Output[[]LinkedTemplateArtifact]{
		OutputState: i.ToLinkedTemplateArtifactArrayOutputWithContext(ctx).OutputState,
	}
}

// Represents a Template Spec artifact containing an embedded Azure Resource Manager template for use as a linked template.
type LinkedTemplateArtifactOutput struct{ *pulumi.OutputState }

func (LinkedTemplateArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedTemplateArtifact)(nil)).Elem()
}

func (o LinkedTemplateArtifactOutput) ToLinkedTemplateArtifactOutput() LinkedTemplateArtifactOutput {
	return o
}

func (o LinkedTemplateArtifactOutput) ToLinkedTemplateArtifactOutputWithContext(ctx context.Context) LinkedTemplateArtifactOutput {
	return o
}

func (o LinkedTemplateArtifactOutput) ToOutput(ctx context.Context) pulumix.Output[LinkedTemplateArtifact] {
	return pulumix.Output[LinkedTemplateArtifact]{
		OutputState: o.OutputState,
	}
}

// A filesystem safe relative path of the artifact.
func (o LinkedTemplateArtifactOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LinkedTemplateArtifact) string { return v.Path }).(pulumi.StringOutput)
}

// The Azure Resource Manager template.
func (o LinkedTemplateArtifactOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v LinkedTemplateArtifact) interface{} { return v.Template }).(pulumi.AnyOutput)
}

type LinkedTemplateArtifactArrayOutput struct{ *pulumi.OutputState }

func (LinkedTemplateArtifactArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedTemplateArtifact)(nil)).Elem()
}

func (o LinkedTemplateArtifactArrayOutput) ToLinkedTemplateArtifactArrayOutput() LinkedTemplateArtifactArrayOutput {
	return o
}

func (o LinkedTemplateArtifactArrayOutput) ToLinkedTemplateArtifactArrayOutputWithContext(ctx context.Context) LinkedTemplateArtifactArrayOutput {
	return o
}

func (o LinkedTemplateArtifactArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]LinkedTemplateArtifact] {
	return pulumix.Output[[]LinkedTemplateArtifact]{
		OutputState: o.OutputState,
	}
}

func (o LinkedTemplateArtifactArrayOutput) Index(i pulumi.IntInput) LinkedTemplateArtifactOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedTemplateArtifact {
		return vs[0].([]LinkedTemplateArtifact)[vs[1].(int)]
	}).(LinkedTemplateArtifactOutput)
}

// Represents a Template Spec artifact containing an embedded Azure Resource Manager template for use as a linked template.
type LinkedTemplateArtifactResponse struct {
	// A filesystem safe relative path of the artifact.
	Path string `pulumi:"path"`
	// The Azure Resource Manager template.
	Template interface{} `pulumi:"template"`
}

// Represents a Template Spec artifact containing an embedded Azure Resource Manager template for use as a linked template.
type LinkedTemplateArtifactResponseOutput struct{ *pulumi.OutputState }

func (LinkedTemplateArtifactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedTemplateArtifactResponse)(nil)).Elem()
}

func (o LinkedTemplateArtifactResponseOutput) ToLinkedTemplateArtifactResponseOutput() LinkedTemplateArtifactResponseOutput {
	return o
}

func (o LinkedTemplateArtifactResponseOutput) ToLinkedTemplateArtifactResponseOutputWithContext(ctx context.Context) LinkedTemplateArtifactResponseOutput {
	return o
}

func (o LinkedTemplateArtifactResponseOutput) ToOutput(ctx context.Context) pulumix.Output[LinkedTemplateArtifactResponse] {
	return pulumix.Output[LinkedTemplateArtifactResponse]{
		OutputState: o.OutputState,
	}
}

// A filesystem safe relative path of the artifact.
func (o LinkedTemplateArtifactResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LinkedTemplateArtifactResponse) string { return v.Path }).(pulumi.StringOutput)
}

// The Azure Resource Manager template.
func (o LinkedTemplateArtifactResponseOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v LinkedTemplateArtifactResponse) interface{} { return v.Template }).(pulumi.AnyOutput)
}

type LinkedTemplateArtifactResponseArrayOutput struct{ *pulumi.OutputState }

func (LinkedTemplateArtifactResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedTemplateArtifactResponse)(nil)).Elem()
}

func (o LinkedTemplateArtifactResponseArrayOutput) ToLinkedTemplateArtifactResponseArrayOutput() LinkedTemplateArtifactResponseArrayOutput {
	return o
}

func (o LinkedTemplateArtifactResponseArrayOutput) ToLinkedTemplateArtifactResponseArrayOutputWithContext(ctx context.Context) LinkedTemplateArtifactResponseArrayOutput {
	return o
}

func (o LinkedTemplateArtifactResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]LinkedTemplateArtifactResponse] {
	return pulumix.Output[[]LinkedTemplateArtifactResponse]{
		OutputState: o.OutputState,
	}
}

func (o LinkedTemplateArtifactResponseArrayOutput) Index(i pulumi.IntInput) LinkedTemplateArtifactResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedTemplateArtifactResponse {
		return vs[0].([]LinkedTemplateArtifactResponse)[vs[1].(int)]
	}).(LinkedTemplateArtifactResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToOutput(ctx context.Context) pulumix.Output[SystemDataResponse] {
	return pulumix.Output[SystemDataResponse]{
		OutputState: o.OutputState,
	}
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

// High-level information about a Template Spec version.
type TemplateSpecVersionInfoResponse struct {
	// Template Spec version description.
	Description string `pulumi:"description"`
	// The timestamp of when the version was created.
	TimeCreated string `pulumi:"timeCreated"`
	// The timestamp of when the version was last modified.
	TimeModified string `pulumi:"timeModified"`
}

// High-level information about a Template Spec version.
type TemplateSpecVersionInfoResponseOutput struct{ *pulumi.OutputState }

func (TemplateSpecVersionInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecVersionInfoResponse)(nil)).Elem()
}

func (o TemplateSpecVersionInfoResponseOutput) ToTemplateSpecVersionInfoResponseOutput() TemplateSpecVersionInfoResponseOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseOutput) ToTemplateSpecVersionInfoResponseOutputWithContext(ctx context.Context) TemplateSpecVersionInfoResponseOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseOutput) ToOutput(ctx context.Context) pulumix.Output[TemplateSpecVersionInfoResponse] {
	return pulumix.Output[TemplateSpecVersionInfoResponse]{
		OutputState: o.OutputState,
	}
}

// Template Spec version description.
func (o TemplateSpecVersionInfoResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecVersionInfoResponse) string { return v.Description }).(pulumi.StringOutput)
}

// The timestamp of when the version was created.
func (o TemplateSpecVersionInfoResponseOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecVersionInfoResponse) string { return v.TimeCreated }).(pulumi.StringOutput)
}

// The timestamp of when the version was last modified.
func (o TemplateSpecVersionInfoResponseOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecVersionInfoResponse) string { return v.TimeModified }).(pulumi.StringOutput)
}

type TemplateSpecVersionInfoResponseMapOutput struct{ *pulumi.OutputState }

func (TemplateSpecVersionInfoResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TemplateSpecVersionInfoResponse)(nil)).Elem()
}

func (o TemplateSpecVersionInfoResponseMapOutput) ToTemplateSpecVersionInfoResponseMapOutput() TemplateSpecVersionInfoResponseMapOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseMapOutput) ToTemplateSpecVersionInfoResponseMapOutputWithContext(ctx context.Context) TemplateSpecVersionInfoResponseMapOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]TemplateSpecVersionInfoResponse] {
	return pulumix.Output[map[string]TemplateSpecVersionInfoResponse]{
		OutputState: o.OutputState,
	}
}

func (o TemplateSpecVersionInfoResponseMapOutput) MapIndex(k pulumi.StringInput) TemplateSpecVersionInfoResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TemplateSpecVersionInfoResponse {
		return vs[0].(map[string]TemplateSpecVersionInfoResponse)[vs[1].(string)]
	}).(TemplateSpecVersionInfoResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedTemplateArtifactOutput{})
	pulumi.RegisterOutputType(LinkedTemplateArtifactArrayOutput{})
	pulumi.RegisterOutputType(LinkedTemplateArtifactResponseOutput{})
	pulumi.RegisterOutputType(LinkedTemplateArtifactResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TemplateSpecVersionInfoResponseOutput{})
	pulumi.RegisterOutputType(TemplateSpecVersionInfoResponseMapOutput{})
}
