// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies the kind of blueprint artifact.
type ArtifactKind string

const (
	ArtifactKindTemplate         = ArtifactKind("template")
	ArtifactKindRoleAssignment   = ArtifactKind("roleAssignment")
	ArtifactKindPolicyAssignment = ArtifactKind("policyAssignment")
)

// Lock mode.
type AssignmentLockMode string

const (
	AssignmentLockModeNone                    = AssignmentLockMode("None")
	AssignmentLockModeAllResourcesReadOnly    = AssignmentLockMode("AllResourcesReadOnly")
	AssignmentLockModeAllResourcesDoNotDelete = AssignmentLockMode("AllResourcesDoNotDelete")
)

func (AssignmentLockMode) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentLockMode)(nil)).Elem()
}

func (e AssignmentLockMode) ToAssignmentLockModeOutput() AssignmentLockModeOutput {
	return pulumi.ToOutput(e).(AssignmentLockModeOutput)
}

func (e AssignmentLockMode) ToAssignmentLockModeOutputWithContext(ctx context.Context) AssignmentLockModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssignmentLockModeOutput)
}

func (e AssignmentLockMode) ToAssignmentLockModePtrOutput() AssignmentLockModePtrOutput {
	return e.ToAssignmentLockModePtrOutputWithContext(context.Background())
}

func (e AssignmentLockMode) ToAssignmentLockModePtrOutputWithContext(ctx context.Context) AssignmentLockModePtrOutput {
	return AssignmentLockMode(e).ToAssignmentLockModeOutputWithContext(ctx).ToAssignmentLockModePtrOutputWithContext(ctx)
}

func (e AssignmentLockMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssignmentLockMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssignmentLockMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssignmentLockMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssignmentLockModeOutput struct{ *pulumi.OutputState }

func (AssignmentLockModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentLockMode)(nil)).Elem()
}

func (o AssignmentLockModeOutput) ToAssignmentLockModeOutput() AssignmentLockModeOutput {
	return o
}

func (o AssignmentLockModeOutput) ToAssignmentLockModeOutputWithContext(ctx context.Context) AssignmentLockModeOutput {
	return o
}

func (o AssignmentLockModeOutput) ToAssignmentLockModePtrOutput() AssignmentLockModePtrOutput {
	return o.ToAssignmentLockModePtrOutputWithContext(context.Background())
}

func (o AssignmentLockModeOutput) ToAssignmentLockModePtrOutputWithContext(ctx context.Context) AssignmentLockModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignmentLockMode) *AssignmentLockMode {
		return &v
	}).(AssignmentLockModePtrOutput)
}

func (o AssignmentLockModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssignmentLockModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssignmentLockMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssignmentLockModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssignmentLockModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssignmentLockMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssignmentLockModePtrOutput struct{ *pulumi.OutputState }

func (AssignmentLockModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentLockMode)(nil)).Elem()
}

func (o AssignmentLockModePtrOutput) ToAssignmentLockModePtrOutput() AssignmentLockModePtrOutput {
	return o
}

func (o AssignmentLockModePtrOutput) ToAssignmentLockModePtrOutputWithContext(ctx context.Context) AssignmentLockModePtrOutput {
	return o
}

func (o AssignmentLockModePtrOutput) Elem() AssignmentLockModeOutput {
	return o.ApplyT(func(v *AssignmentLockMode) AssignmentLockMode {
		if v != nil {
			return *v
		}
		var ret AssignmentLockMode
		return ret
	}).(AssignmentLockModeOutput)
}

func (o AssignmentLockModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssignmentLockModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssignmentLockMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AssignmentLockModeInput is an input type that accepts values of the AssignmentLockMode enum
// A concrete instance of `AssignmentLockModeInput` can be one of the following:
//
//	AssignmentLockModeNone
//	AssignmentLockModeAllResourcesReadOnly
//	AssignmentLockModeAllResourcesDoNotDelete
type AssignmentLockModeInput interface {
	pulumi.Input

	ToAssignmentLockModeOutput() AssignmentLockModeOutput
	ToAssignmentLockModeOutputWithContext(context.Context) AssignmentLockModeOutput
}

var assignmentLockModePtrType = reflect.TypeOf((**AssignmentLockMode)(nil)).Elem()

type AssignmentLockModePtrInput interface {
	pulumi.Input

	ToAssignmentLockModePtrOutput() AssignmentLockModePtrOutput
	ToAssignmentLockModePtrOutputWithContext(context.Context) AssignmentLockModePtrOutput
}

type assignmentLockModePtr string

func AssignmentLockModePtr(v string) AssignmentLockModePtrInput {
	return (*assignmentLockModePtr)(&v)
}

func (*assignmentLockModePtr) ElementType() reflect.Type {
	return assignmentLockModePtrType
}

func (in *assignmentLockModePtr) ToAssignmentLockModePtrOutput() AssignmentLockModePtrOutput {
	return pulumi.ToOutput(in).(AssignmentLockModePtrOutput)
}

func (in *assignmentLockModePtr) ToAssignmentLockModePtrOutputWithContext(ctx context.Context) AssignmentLockModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssignmentLockModePtrOutput)
}

// The scope where this blueprint definition can be assigned.
type BlueprintTargetScope string

const (
	// The blueprint targets a subscription during blueprint assignment.
	BlueprintTargetScopeSubscription = BlueprintTargetScope("subscription")
	// The blueprint targets a management group during blueprint assignment. This is reserved for future use.
	BlueprintTargetScopeManagementGroup = BlueprintTargetScope("managementGroup")
)

func (BlueprintTargetScope) ElementType() reflect.Type {
	return reflect.TypeOf((*BlueprintTargetScope)(nil)).Elem()
}

func (e BlueprintTargetScope) ToBlueprintTargetScopeOutput() BlueprintTargetScopeOutput {
	return pulumi.ToOutput(e).(BlueprintTargetScopeOutput)
}

func (e BlueprintTargetScope) ToBlueprintTargetScopeOutputWithContext(ctx context.Context) BlueprintTargetScopeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BlueprintTargetScopeOutput)
}

func (e BlueprintTargetScope) ToBlueprintTargetScopePtrOutput() BlueprintTargetScopePtrOutput {
	return e.ToBlueprintTargetScopePtrOutputWithContext(context.Background())
}

func (e BlueprintTargetScope) ToBlueprintTargetScopePtrOutputWithContext(ctx context.Context) BlueprintTargetScopePtrOutput {
	return BlueprintTargetScope(e).ToBlueprintTargetScopeOutputWithContext(ctx).ToBlueprintTargetScopePtrOutputWithContext(ctx)
}

func (e BlueprintTargetScope) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlueprintTargetScope) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlueprintTargetScope) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BlueprintTargetScope) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BlueprintTargetScopeOutput struct{ *pulumi.OutputState }

func (BlueprintTargetScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlueprintTargetScope)(nil)).Elem()
}

func (o BlueprintTargetScopeOutput) ToBlueprintTargetScopeOutput() BlueprintTargetScopeOutput {
	return o
}

func (o BlueprintTargetScopeOutput) ToBlueprintTargetScopeOutputWithContext(ctx context.Context) BlueprintTargetScopeOutput {
	return o
}

func (o BlueprintTargetScopeOutput) ToBlueprintTargetScopePtrOutput() BlueprintTargetScopePtrOutput {
	return o.ToBlueprintTargetScopePtrOutputWithContext(context.Background())
}

func (o BlueprintTargetScopeOutput) ToBlueprintTargetScopePtrOutputWithContext(ctx context.Context) BlueprintTargetScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlueprintTargetScope) *BlueprintTargetScope {
		return &v
	}).(BlueprintTargetScopePtrOutput)
}

func (o BlueprintTargetScopeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BlueprintTargetScopeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlueprintTargetScope) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BlueprintTargetScopeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlueprintTargetScopeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlueprintTargetScope) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BlueprintTargetScopePtrOutput struct{ *pulumi.OutputState }

func (BlueprintTargetScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlueprintTargetScope)(nil)).Elem()
}

func (o BlueprintTargetScopePtrOutput) ToBlueprintTargetScopePtrOutput() BlueprintTargetScopePtrOutput {
	return o
}

func (o BlueprintTargetScopePtrOutput) ToBlueprintTargetScopePtrOutputWithContext(ctx context.Context) BlueprintTargetScopePtrOutput {
	return o
}

func (o BlueprintTargetScopePtrOutput) Elem() BlueprintTargetScopeOutput {
	return o.ApplyT(func(v *BlueprintTargetScope) BlueprintTargetScope {
		if v != nil {
			return *v
		}
		var ret BlueprintTargetScope
		return ret
	}).(BlueprintTargetScopeOutput)
}

func (o BlueprintTargetScopePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlueprintTargetScopePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BlueprintTargetScope) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// BlueprintTargetScopeInput is an input type that accepts values of the BlueprintTargetScope enum
// A concrete instance of `BlueprintTargetScopeInput` can be one of the following:
//
//	BlueprintTargetScopeSubscription
//	BlueprintTargetScopeManagementGroup
type BlueprintTargetScopeInput interface {
	pulumi.Input

	ToBlueprintTargetScopeOutput() BlueprintTargetScopeOutput
	ToBlueprintTargetScopeOutputWithContext(context.Context) BlueprintTargetScopeOutput
}

var blueprintTargetScopePtrType = reflect.TypeOf((**BlueprintTargetScope)(nil)).Elem()

type BlueprintTargetScopePtrInput interface {
	pulumi.Input

	ToBlueprintTargetScopePtrOutput() BlueprintTargetScopePtrOutput
	ToBlueprintTargetScopePtrOutputWithContext(context.Context) BlueprintTargetScopePtrOutput
}

type blueprintTargetScopePtr string

func BlueprintTargetScopePtr(v string) BlueprintTargetScopePtrInput {
	return (*blueprintTargetScopePtr)(&v)
}

func (*blueprintTargetScopePtr) ElementType() reflect.Type {
	return blueprintTargetScopePtrType
}

func (in *blueprintTargetScopePtr) ToBlueprintTargetScopePtrOutput() BlueprintTargetScopePtrOutput {
	return pulumi.ToOutput(in).(BlueprintTargetScopePtrOutput)
}

func (in *blueprintTargetScopePtr) ToBlueprintTargetScopePtrOutputWithContext(ctx context.Context) BlueprintTargetScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BlueprintTargetScopePtrOutput)
}

// Type of the managed identity.
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone           = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned   = ManagedServiceIdentityType("UserAssigned")
)

func (ManagedServiceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return e.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return ManagedServiceIdentityType(e).ToManagedServiceIdentityTypeOutputWithContext(ctx).ToManagedServiceIdentityTypePtrOutputWithContext(ctx)
}

func (e ManagedServiceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedServiceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentityType) *ManagedServiceIdentityType {
		return &v
	}).(ManagedServiceIdentityTypePtrOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedServiceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) Elem() ManagedServiceIdentityTypeOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityType) ManagedServiceIdentityType {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityType
		return ret
	}).(ManagedServiceIdentityTypeOutput)
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedServiceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ManagedServiceIdentityTypeInput is an input type that accepts values of the ManagedServiceIdentityType enum
// A concrete instance of `ManagedServiceIdentityTypeInput` can be one of the following:
//
//	ManagedServiceIdentityTypeNone
//	ManagedServiceIdentityTypeSystemAssigned
//	ManagedServiceIdentityTypeUserAssigned
type ManagedServiceIdentityTypeInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput
	ToManagedServiceIdentityTypeOutputWithContext(context.Context) ManagedServiceIdentityTypeOutput
}

var managedServiceIdentityTypePtrType = reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()

type ManagedServiceIdentityTypePtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput
	ToManagedServiceIdentityTypePtrOutputWithContext(context.Context) ManagedServiceIdentityTypePtrOutput
}

type managedServiceIdentityTypePtr string

func ManagedServiceIdentityTypePtr(v string) ManagedServiceIdentityTypePtrInput {
	return (*managedServiceIdentityTypePtr)(&v)
}

func (*managedServiceIdentityTypePtr) ElementType() reflect.Type {
	return managedServiceIdentityTypePtrType
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ManagedServiceIdentityTypePtrOutput)
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedServiceIdentityTypePtrOutput)
}

// Allowed data types for Resource Manager template parameters.
type TemplateParameterType string

const (
	TemplateParameterTypeString       = TemplateParameterType("string")
	TemplateParameterTypeArray        = TemplateParameterType("array")
	TemplateParameterTypeBool         = TemplateParameterType("bool")
	TemplateParameterTypeInt          = TemplateParameterType("int")
	TemplateParameterTypeObject       = TemplateParameterType("object")
	TemplateParameterTypeSecureObject = TemplateParameterType("secureObject")
	TemplateParameterTypeSecureString = TemplateParameterType("secureString")
)

func (TemplateParameterType) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateParameterType)(nil)).Elem()
}

func (e TemplateParameterType) ToTemplateParameterTypeOutput() TemplateParameterTypeOutput {
	return pulumi.ToOutput(e).(TemplateParameterTypeOutput)
}

func (e TemplateParameterType) ToTemplateParameterTypeOutputWithContext(ctx context.Context) TemplateParameterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TemplateParameterTypeOutput)
}

func (e TemplateParameterType) ToTemplateParameterTypePtrOutput() TemplateParameterTypePtrOutput {
	return e.ToTemplateParameterTypePtrOutputWithContext(context.Background())
}

func (e TemplateParameterType) ToTemplateParameterTypePtrOutputWithContext(ctx context.Context) TemplateParameterTypePtrOutput {
	return TemplateParameterType(e).ToTemplateParameterTypeOutputWithContext(ctx).ToTemplateParameterTypePtrOutputWithContext(ctx)
}

func (e TemplateParameterType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TemplateParameterType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TemplateParameterType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TemplateParameterType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TemplateParameterTypeOutput struct{ *pulumi.OutputState }

func (TemplateParameterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateParameterType)(nil)).Elem()
}

func (o TemplateParameterTypeOutput) ToTemplateParameterTypeOutput() TemplateParameterTypeOutput {
	return o
}

func (o TemplateParameterTypeOutput) ToTemplateParameterTypeOutputWithContext(ctx context.Context) TemplateParameterTypeOutput {
	return o
}

func (o TemplateParameterTypeOutput) ToTemplateParameterTypePtrOutput() TemplateParameterTypePtrOutput {
	return o.ToTemplateParameterTypePtrOutputWithContext(context.Background())
}

func (o TemplateParameterTypeOutput) ToTemplateParameterTypePtrOutputWithContext(ctx context.Context) TemplateParameterTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TemplateParameterType) *TemplateParameterType {
		return &v
	}).(TemplateParameterTypePtrOutput)
}

func (o TemplateParameterTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TemplateParameterTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TemplateParameterType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TemplateParameterTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TemplateParameterTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TemplateParameterType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TemplateParameterTypePtrOutput struct{ *pulumi.OutputState }

func (TemplateParameterTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateParameterType)(nil)).Elem()
}

func (o TemplateParameterTypePtrOutput) ToTemplateParameterTypePtrOutput() TemplateParameterTypePtrOutput {
	return o
}

func (o TemplateParameterTypePtrOutput) ToTemplateParameterTypePtrOutputWithContext(ctx context.Context) TemplateParameterTypePtrOutput {
	return o
}

func (o TemplateParameterTypePtrOutput) Elem() TemplateParameterTypeOutput {
	return o.ApplyT(func(v *TemplateParameterType) TemplateParameterType {
		if v != nil {
			return *v
		}
		var ret TemplateParameterType
		return ret
	}).(TemplateParameterTypeOutput)
}

func (o TemplateParameterTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TemplateParameterTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TemplateParameterType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TemplateParameterTypeInput is an input type that accepts values of the TemplateParameterType enum
// A concrete instance of `TemplateParameterTypeInput` can be one of the following:
//
//	TemplateParameterTypeString
//	TemplateParameterTypeArray
//	TemplateParameterTypeBool
//	TemplateParameterTypeInt
//	TemplateParameterTypeObject
//	TemplateParameterTypeSecureObject
//	TemplateParameterTypeSecureString
type TemplateParameterTypeInput interface {
	pulumi.Input

	ToTemplateParameterTypeOutput() TemplateParameterTypeOutput
	ToTemplateParameterTypeOutputWithContext(context.Context) TemplateParameterTypeOutput
}

var templateParameterTypePtrType = reflect.TypeOf((**TemplateParameterType)(nil)).Elem()

type TemplateParameterTypePtrInput interface {
	pulumi.Input

	ToTemplateParameterTypePtrOutput() TemplateParameterTypePtrOutput
	ToTemplateParameterTypePtrOutputWithContext(context.Context) TemplateParameterTypePtrOutput
}

type templateParameterTypePtr string

func TemplateParameterTypePtr(v string) TemplateParameterTypePtrInput {
	return (*templateParameterTypePtr)(&v)
}

func (*templateParameterTypePtr) ElementType() reflect.Type {
	return templateParameterTypePtrType
}

func (in *templateParameterTypePtr) ToTemplateParameterTypePtrOutput() TemplateParameterTypePtrOutput {
	return pulumi.ToOutput(in).(TemplateParameterTypePtrOutput)
}

func (in *templateParameterTypePtr) ToTemplateParameterTypePtrOutputWithContext(ctx context.Context) TemplateParameterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TemplateParameterTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssignmentLockModeOutput{})
	pulumi.RegisterOutputType(AssignmentLockModePtrOutput{})
	pulumi.RegisterOutputType(BlueprintTargetScopeOutput{})
	pulumi.RegisterOutputType(BlueprintTargetScopePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(TemplateParameterTypeOutput{})
	pulumi.RegisterOutputType(TemplateParameterTypePtrOutput{})
}
