// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The managed application definition artifact type.
type ApplicationArtifactType string

const (
	ApplicationArtifactTypeNotSpecified = ApplicationArtifactType("NotSpecified")
	ApplicationArtifactTypeTemplate     = ApplicationArtifactType("Template")
	ApplicationArtifactTypeCustom       = ApplicationArtifactType("Custom")
)

func (ApplicationArtifactType) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationArtifactType)(nil)).Elem()
}

func (e ApplicationArtifactType) ToApplicationArtifactTypeOutput() ApplicationArtifactTypeOutput {
	return pulumi.ToOutput(e).(ApplicationArtifactTypeOutput)
}

func (e ApplicationArtifactType) ToApplicationArtifactTypeOutputWithContext(ctx context.Context) ApplicationArtifactTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationArtifactTypeOutput)
}

func (e ApplicationArtifactType) ToApplicationArtifactTypePtrOutput() ApplicationArtifactTypePtrOutput {
	return e.ToApplicationArtifactTypePtrOutputWithContext(context.Background())
}

func (e ApplicationArtifactType) ToApplicationArtifactTypePtrOutputWithContext(ctx context.Context) ApplicationArtifactTypePtrOutput {
	return ApplicationArtifactType(e).ToApplicationArtifactTypeOutputWithContext(ctx).ToApplicationArtifactTypePtrOutputWithContext(ctx)
}

func (e ApplicationArtifactType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationArtifactType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationArtifactType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationArtifactType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationArtifactTypeOutput struct{ *pulumi.OutputState }

func (ApplicationArtifactTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationArtifactType)(nil)).Elem()
}

func (o ApplicationArtifactTypeOutput) ToApplicationArtifactTypeOutput() ApplicationArtifactTypeOutput {
	return o
}

func (o ApplicationArtifactTypeOutput) ToApplicationArtifactTypeOutputWithContext(ctx context.Context) ApplicationArtifactTypeOutput {
	return o
}

func (o ApplicationArtifactTypeOutput) ToApplicationArtifactTypePtrOutput() ApplicationArtifactTypePtrOutput {
	return o.ToApplicationArtifactTypePtrOutputWithContext(context.Background())
}

func (o ApplicationArtifactTypeOutput) ToApplicationArtifactTypePtrOutputWithContext(ctx context.Context) ApplicationArtifactTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationArtifactType) *ApplicationArtifactType {
		return &v
	}).(ApplicationArtifactTypePtrOutput)
}

func (o ApplicationArtifactTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationArtifactTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationArtifactType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationArtifactTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationArtifactTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationArtifactType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationArtifactTypePtrOutput struct{ *pulumi.OutputState }

func (ApplicationArtifactTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationArtifactType)(nil)).Elem()
}

func (o ApplicationArtifactTypePtrOutput) ToApplicationArtifactTypePtrOutput() ApplicationArtifactTypePtrOutput {
	return o
}

func (o ApplicationArtifactTypePtrOutput) ToApplicationArtifactTypePtrOutputWithContext(ctx context.Context) ApplicationArtifactTypePtrOutput {
	return o
}

func (o ApplicationArtifactTypePtrOutput) Elem() ApplicationArtifactTypeOutput {
	return o.ApplyT(func(v *ApplicationArtifactType) ApplicationArtifactType {
		if v != nil {
			return *v
		}
		var ret ApplicationArtifactType
		return ret
	}).(ApplicationArtifactTypeOutput)
}

func (o ApplicationArtifactTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationArtifactTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationArtifactType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ApplicationArtifactTypeInput is an input type that accepts ApplicationArtifactTypeArgs and ApplicationArtifactTypeOutput values.
// You can construct a concrete instance of `ApplicationArtifactTypeInput` via:
//
//	ApplicationArtifactTypeArgs{...}
type ApplicationArtifactTypeInput interface {
	pulumi.Input

	ToApplicationArtifactTypeOutput() ApplicationArtifactTypeOutput
	ToApplicationArtifactTypeOutputWithContext(context.Context) ApplicationArtifactTypeOutput
}

var applicationArtifactTypePtrType = reflect.TypeOf((**ApplicationArtifactType)(nil)).Elem()

type ApplicationArtifactTypePtrInput interface {
	pulumi.Input

	ToApplicationArtifactTypePtrOutput() ApplicationArtifactTypePtrOutput
	ToApplicationArtifactTypePtrOutputWithContext(context.Context) ApplicationArtifactTypePtrOutput
}

type applicationArtifactTypePtr string

func ApplicationArtifactTypePtr(v string) ApplicationArtifactTypePtrInput {
	return (*applicationArtifactTypePtr)(&v)
}

func (*applicationArtifactTypePtr) ElementType() reflect.Type {
	return applicationArtifactTypePtrType
}

func (in *applicationArtifactTypePtr) ToApplicationArtifactTypePtrOutput() ApplicationArtifactTypePtrOutput {
	return pulumi.ToOutput(in).(ApplicationArtifactTypePtrOutput)
}

func (in *applicationArtifactTypePtr) ToApplicationArtifactTypePtrOutputWithContext(ctx context.Context) ApplicationArtifactTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationArtifactTypePtrOutput)
}

func (in *applicationArtifactTypePtr) ToOutput(ctx context.Context) pulumix.Output[*ApplicationArtifactType] {
	return pulumix.Output[*ApplicationArtifactType]{
		OutputState: in.ToApplicationArtifactTypePtrOutputWithContext(ctx).OutputState,
	}
}

// The managed application definition artifact name.
type ApplicationDefinitionArtifactName string

const (
	ApplicationDefinitionArtifactNameNotSpecified                = ApplicationDefinitionArtifactName("NotSpecified")
	ApplicationDefinitionArtifactNameApplicationResourceTemplate = ApplicationDefinitionArtifactName("ApplicationResourceTemplate")
	ApplicationDefinitionArtifactNameCreateUiDefinition          = ApplicationDefinitionArtifactName("CreateUiDefinition")
	ApplicationDefinitionArtifactNameMainTemplateParameters      = ApplicationDefinitionArtifactName("MainTemplateParameters")
)

// The managed application lock level.
type ApplicationLockLevel string

const (
	ApplicationLockLevelCanNotDelete = ApplicationLockLevel("CanNotDelete")
	ApplicationLockLevelReadOnly     = ApplicationLockLevel("ReadOnly")
	ApplicationLockLevelNone         = ApplicationLockLevel("None")
)

func (ApplicationLockLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLockLevel)(nil)).Elem()
}

func (e ApplicationLockLevel) ToApplicationLockLevelOutput() ApplicationLockLevelOutput {
	return pulumi.ToOutput(e).(ApplicationLockLevelOutput)
}

func (e ApplicationLockLevel) ToApplicationLockLevelOutputWithContext(ctx context.Context) ApplicationLockLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationLockLevelOutput)
}

func (e ApplicationLockLevel) ToApplicationLockLevelPtrOutput() ApplicationLockLevelPtrOutput {
	return e.ToApplicationLockLevelPtrOutputWithContext(context.Background())
}

func (e ApplicationLockLevel) ToApplicationLockLevelPtrOutputWithContext(ctx context.Context) ApplicationLockLevelPtrOutput {
	return ApplicationLockLevel(e).ToApplicationLockLevelOutputWithContext(ctx).ToApplicationLockLevelPtrOutputWithContext(ctx)
}

func (e ApplicationLockLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationLockLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationLockLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationLockLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationLockLevelOutput struct{ *pulumi.OutputState }

func (ApplicationLockLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLockLevel)(nil)).Elem()
}

func (o ApplicationLockLevelOutput) ToApplicationLockLevelOutput() ApplicationLockLevelOutput {
	return o
}

func (o ApplicationLockLevelOutput) ToApplicationLockLevelOutputWithContext(ctx context.Context) ApplicationLockLevelOutput {
	return o
}

func (o ApplicationLockLevelOutput) ToApplicationLockLevelPtrOutput() ApplicationLockLevelPtrOutput {
	return o.ToApplicationLockLevelPtrOutputWithContext(context.Background())
}

func (o ApplicationLockLevelOutput) ToApplicationLockLevelPtrOutputWithContext(ctx context.Context) ApplicationLockLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationLockLevel) *ApplicationLockLevel {
		return &v
	}).(ApplicationLockLevelPtrOutput)
}

func (o ApplicationLockLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationLockLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationLockLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationLockLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationLockLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationLockLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationLockLevelPtrOutput struct{ *pulumi.OutputState }

func (ApplicationLockLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLockLevel)(nil)).Elem()
}

func (o ApplicationLockLevelPtrOutput) ToApplicationLockLevelPtrOutput() ApplicationLockLevelPtrOutput {
	return o
}

func (o ApplicationLockLevelPtrOutput) ToApplicationLockLevelPtrOutputWithContext(ctx context.Context) ApplicationLockLevelPtrOutput {
	return o
}

func (o ApplicationLockLevelPtrOutput) Elem() ApplicationLockLevelOutput {
	return o.ApplyT(func(v *ApplicationLockLevel) ApplicationLockLevel {
		if v != nil {
			return *v
		}
		var ret ApplicationLockLevel
		return ret
	}).(ApplicationLockLevelOutput)
}

func (o ApplicationLockLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationLockLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationLockLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ApplicationLockLevelInput is an input type that accepts ApplicationLockLevelArgs and ApplicationLockLevelOutput values.
// You can construct a concrete instance of `ApplicationLockLevelInput` via:
//
//	ApplicationLockLevelArgs{...}
type ApplicationLockLevelInput interface {
	pulumi.Input

	ToApplicationLockLevelOutput() ApplicationLockLevelOutput
	ToApplicationLockLevelOutputWithContext(context.Context) ApplicationLockLevelOutput
}

var applicationLockLevelPtrType = reflect.TypeOf((**ApplicationLockLevel)(nil)).Elem()

type ApplicationLockLevelPtrInput interface {
	pulumi.Input

	ToApplicationLockLevelPtrOutput() ApplicationLockLevelPtrOutput
	ToApplicationLockLevelPtrOutputWithContext(context.Context) ApplicationLockLevelPtrOutput
}

type applicationLockLevelPtr string

func ApplicationLockLevelPtr(v string) ApplicationLockLevelPtrInput {
	return (*applicationLockLevelPtr)(&v)
}

func (*applicationLockLevelPtr) ElementType() reflect.Type {
	return applicationLockLevelPtrType
}

func (in *applicationLockLevelPtr) ToApplicationLockLevelPtrOutput() ApplicationLockLevelPtrOutput {
	return pulumi.ToOutput(in).(ApplicationLockLevelPtrOutput)
}

func (in *applicationLockLevelPtr) ToApplicationLockLevelPtrOutputWithContext(ctx context.Context) ApplicationLockLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationLockLevelPtrOutput)
}

func (in *applicationLockLevelPtr) ToOutput(ctx context.Context) pulumix.Output[*ApplicationLockLevel] {
	return pulumix.Output[*ApplicationLockLevel]{
		OutputState: in.ToApplicationLockLevelPtrOutputWithContext(ctx).OutputState,
	}
}

// The managed application management mode.
type ApplicationManagementMode string

const (
	ApplicationManagementModeNotSpecified = ApplicationManagementMode("NotSpecified")
	ApplicationManagementModeUnmanaged    = ApplicationManagementMode("Unmanaged")
	ApplicationManagementModeManaged      = ApplicationManagementMode("Managed")
)

// The managed application deployment mode.
type DeploymentMode string

const (
	DeploymentModeNotSpecified = DeploymentMode("NotSpecified")
	DeploymentModeIncremental  = DeploymentMode("Incremental")
	DeploymentModeComplete     = DeploymentMode("Complete")
)

// JIT approval mode.
type JitApprovalMode string

const (
	JitApprovalModeNotSpecified  = JitApprovalMode("NotSpecified")
	JitApprovalModeAutoApprove   = JitApprovalMode("AutoApprove")
	JitApprovalModeManualApprove = JitApprovalMode("ManualApprove")
)

// The approver type.
type JitApproverType string

const (
	JitApproverTypeUser  = JitApproverType("user")
	JitApproverTypeGroup = JitApproverType("group")
)

// The type of JIT schedule.
type JitSchedulingType string

const (
	JitSchedulingTypeNotSpecified = JitSchedulingType("NotSpecified")
	JitSchedulingTypeOnce         = JitSchedulingType("Once")
	JitSchedulingTypeRecurring    = JitSchedulingType("Recurring")
)

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

// ResourceIdentityTypeInput is an input type that accepts ResourceIdentityTypeArgs and ResourceIdentityTypeOutput values.
// You can construct a concrete instance of `ResourceIdentityTypeInput` via:
//
//	ResourceIdentityTypeArgs{...}
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

func (in *resourceIdentityTypePtr) ToOutput(ctx context.Context) pulumix.Output[*ResourceIdentityType] {
	return pulumix.Output[*ResourceIdentityType]{
		OutputState: in.ToResourceIdentityTypePtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterOutputType(ApplicationArtifactTypeOutput{})
	pulumi.RegisterOutputType(ApplicationArtifactTypePtrOutput{})
	pulumi.RegisterOutputType(ApplicationLockLevelOutput{})
	pulumi.RegisterOutputType(ApplicationLockLevelPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
