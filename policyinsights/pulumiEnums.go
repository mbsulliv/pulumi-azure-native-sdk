// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package policyinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The compliance state that should be set on the resource.
type ComplianceState string

const (
	// The resource is in compliance with the policy.
	ComplianceStateCompliant = ComplianceState("Compliant")
	// The resource is not in compliance with the policy.
	ComplianceStateNonCompliant = ComplianceState("NonCompliant")
	// The compliance state of the resource is not known.
	ComplianceStateUnknown = ComplianceState("Unknown")
)

func (ComplianceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ComplianceState)(nil)).Elem()
}

func (e ComplianceState) ToComplianceStateOutput() ComplianceStateOutput {
	return pulumi.ToOutput(e).(ComplianceStateOutput)
}

func (e ComplianceState) ToComplianceStateOutputWithContext(ctx context.Context) ComplianceStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ComplianceStateOutput)
}

func (e ComplianceState) ToComplianceStatePtrOutput() ComplianceStatePtrOutput {
	return e.ToComplianceStatePtrOutputWithContext(context.Background())
}

func (e ComplianceState) ToComplianceStatePtrOutputWithContext(ctx context.Context) ComplianceStatePtrOutput {
	return ComplianceState(e).ToComplianceStateOutputWithContext(ctx).ToComplianceStatePtrOutputWithContext(ctx)
}

func (e ComplianceState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComplianceState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComplianceState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComplianceState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComplianceStateOutput struct{ *pulumi.OutputState }

func (ComplianceStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComplianceState)(nil)).Elem()
}

func (o ComplianceStateOutput) ToComplianceStateOutput() ComplianceStateOutput {
	return o
}

func (o ComplianceStateOutput) ToComplianceStateOutputWithContext(ctx context.Context) ComplianceStateOutput {
	return o
}

func (o ComplianceStateOutput) ToComplianceStatePtrOutput() ComplianceStatePtrOutput {
	return o.ToComplianceStatePtrOutputWithContext(context.Background())
}

func (o ComplianceStateOutput) ToComplianceStatePtrOutputWithContext(ctx context.Context) ComplianceStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComplianceState) *ComplianceState {
		return &v
	}).(ComplianceStatePtrOutput)
}

func (o ComplianceStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ComplianceStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComplianceState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ComplianceStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComplianceStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComplianceState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ComplianceStatePtrOutput struct{ *pulumi.OutputState }

func (ComplianceStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComplianceState)(nil)).Elem()
}

func (o ComplianceStatePtrOutput) ToComplianceStatePtrOutput() ComplianceStatePtrOutput {
	return o
}

func (o ComplianceStatePtrOutput) ToComplianceStatePtrOutputWithContext(ctx context.Context) ComplianceStatePtrOutput {
	return o
}

func (o ComplianceStatePtrOutput) Elem() ComplianceStateOutput {
	return o.ApplyT(func(v *ComplianceState) ComplianceState {
		if v != nil {
			return *v
		}
		var ret ComplianceState
		return ret
	}).(ComplianceStateOutput)
}

func (o ComplianceStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComplianceStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ComplianceState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ComplianceStateInput is an input type that accepts values of the ComplianceState enum
// A concrete instance of `ComplianceStateInput` can be one of the following:
//
//	ComplianceStateCompliant
//	ComplianceStateNonCompliant
//	ComplianceStateUnknown
type ComplianceStateInput interface {
	pulumi.Input

	ToComplianceStateOutput() ComplianceStateOutput
	ToComplianceStateOutputWithContext(context.Context) ComplianceStateOutput
}

var complianceStatePtrType = reflect.TypeOf((**ComplianceState)(nil)).Elem()

type ComplianceStatePtrInput interface {
	pulumi.Input

	ToComplianceStatePtrOutput() ComplianceStatePtrOutput
	ToComplianceStatePtrOutputWithContext(context.Context) ComplianceStatePtrOutput
}

type complianceStatePtr string

func ComplianceStatePtr(v string) ComplianceStatePtrInput {
	return (*complianceStatePtr)(&v)
}

func (*complianceStatePtr) ElementType() reflect.Type {
	return complianceStatePtrType
}

func (in *complianceStatePtr) ToComplianceStatePtrOutput() ComplianceStatePtrOutput {
	return pulumi.ToOutput(in).(ComplianceStatePtrOutput)
}

func (in *complianceStatePtr) ToComplianceStatePtrOutputWithContext(ctx context.Context) ComplianceStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ComplianceStatePtrOutput)
}

// The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
type ResourceDiscoveryMode string

const (
	// Remediate resources that are already known to be non-compliant.
	ResourceDiscoveryModeExistingNonCompliant = ResourceDiscoveryMode("ExistingNonCompliant")
	// Re-evaluate the compliance state of resources and then remediate the resources found to be non-compliant.
	ResourceDiscoveryModeReEvaluateCompliance = ResourceDiscoveryMode("ReEvaluateCompliance")
)

func (ResourceDiscoveryMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceDiscoveryMode)(nil)).Elem()
}

func (e ResourceDiscoveryMode) ToResourceDiscoveryModeOutput() ResourceDiscoveryModeOutput {
	return pulumi.ToOutput(e).(ResourceDiscoveryModeOutput)
}

func (e ResourceDiscoveryMode) ToResourceDiscoveryModeOutputWithContext(ctx context.Context) ResourceDiscoveryModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceDiscoveryModeOutput)
}

func (e ResourceDiscoveryMode) ToResourceDiscoveryModePtrOutput() ResourceDiscoveryModePtrOutput {
	return e.ToResourceDiscoveryModePtrOutputWithContext(context.Background())
}

func (e ResourceDiscoveryMode) ToResourceDiscoveryModePtrOutputWithContext(ctx context.Context) ResourceDiscoveryModePtrOutput {
	return ResourceDiscoveryMode(e).ToResourceDiscoveryModeOutputWithContext(ctx).ToResourceDiscoveryModePtrOutputWithContext(ctx)
}

func (e ResourceDiscoveryMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceDiscoveryMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceDiscoveryMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceDiscoveryMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceDiscoveryModeOutput struct{ *pulumi.OutputState }

func (ResourceDiscoveryModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceDiscoveryMode)(nil)).Elem()
}

func (o ResourceDiscoveryModeOutput) ToResourceDiscoveryModeOutput() ResourceDiscoveryModeOutput {
	return o
}

func (o ResourceDiscoveryModeOutput) ToResourceDiscoveryModeOutputWithContext(ctx context.Context) ResourceDiscoveryModeOutput {
	return o
}

func (o ResourceDiscoveryModeOutput) ToResourceDiscoveryModePtrOutput() ResourceDiscoveryModePtrOutput {
	return o.ToResourceDiscoveryModePtrOutputWithContext(context.Background())
}

func (o ResourceDiscoveryModeOutput) ToResourceDiscoveryModePtrOutputWithContext(ctx context.Context) ResourceDiscoveryModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceDiscoveryMode) *ResourceDiscoveryMode {
		return &v
	}).(ResourceDiscoveryModePtrOutput)
}

func (o ResourceDiscoveryModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceDiscoveryModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceDiscoveryMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceDiscoveryModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceDiscoveryModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceDiscoveryMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceDiscoveryModePtrOutput struct{ *pulumi.OutputState }

func (ResourceDiscoveryModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceDiscoveryMode)(nil)).Elem()
}

func (o ResourceDiscoveryModePtrOutput) ToResourceDiscoveryModePtrOutput() ResourceDiscoveryModePtrOutput {
	return o
}

func (o ResourceDiscoveryModePtrOutput) ToResourceDiscoveryModePtrOutputWithContext(ctx context.Context) ResourceDiscoveryModePtrOutput {
	return o
}

func (o ResourceDiscoveryModePtrOutput) Elem() ResourceDiscoveryModeOutput {
	return o.ApplyT(func(v *ResourceDiscoveryMode) ResourceDiscoveryMode {
		if v != nil {
			return *v
		}
		var ret ResourceDiscoveryMode
		return ret
	}).(ResourceDiscoveryModeOutput)
}

func (o ResourceDiscoveryModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceDiscoveryModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceDiscoveryMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ResourceDiscoveryModeInput is an input type that accepts values of the ResourceDiscoveryMode enum
// A concrete instance of `ResourceDiscoveryModeInput` can be one of the following:
//
//	ResourceDiscoveryModeExistingNonCompliant
//	ResourceDiscoveryModeReEvaluateCompliance
type ResourceDiscoveryModeInput interface {
	pulumi.Input

	ToResourceDiscoveryModeOutput() ResourceDiscoveryModeOutput
	ToResourceDiscoveryModeOutputWithContext(context.Context) ResourceDiscoveryModeOutput
}

var resourceDiscoveryModePtrType = reflect.TypeOf((**ResourceDiscoveryMode)(nil)).Elem()

type ResourceDiscoveryModePtrInput interface {
	pulumi.Input

	ToResourceDiscoveryModePtrOutput() ResourceDiscoveryModePtrOutput
	ToResourceDiscoveryModePtrOutputWithContext(context.Context) ResourceDiscoveryModePtrOutput
}

type resourceDiscoveryModePtr string

func ResourceDiscoveryModePtr(v string) ResourceDiscoveryModePtrInput {
	return (*resourceDiscoveryModePtr)(&v)
}

func (*resourceDiscoveryModePtr) ElementType() reflect.Type {
	return resourceDiscoveryModePtrType
}

func (in *resourceDiscoveryModePtr) ToResourceDiscoveryModePtrOutput() ResourceDiscoveryModePtrOutput {
	return pulumi.ToOutput(in).(ResourceDiscoveryModePtrOutput)
}

func (in *resourceDiscoveryModePtr) ToResourceDiscoveryModePtrOutputWithContext(ctx context.Context) ResourceDiscoveryModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceDiscoveryModePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ComplianceStateOutput{})
	pulumi.RegisterOutputType(ComplianceStatePtrOutput{})
	pulumi.RegisterOutputType(ResourceDiscoveryModeOutput{})
	pulumi.RegisterOutputType(ResourceDiscoveryModePtrOutput{})
}
