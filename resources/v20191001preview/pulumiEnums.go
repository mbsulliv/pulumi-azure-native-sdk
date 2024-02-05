// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
type CleanupOptions string

const (
	CleanupOptionsAlways       = CleanupOptions("Always")
	CleanupOptionsOnSuccess    = CleanupOptions("OnSuccess")
	CleanupOptionsOnExpiration = CleanupOptions("OnExpiration")
)

func (CleanupOptions) ElementType() reflect.Type {
	return reflect.TypeOf((*CleanupOptions)(nil)).Elem()
}

func (e CleanupOptions) ToCleanupOptionsOutput() CleanupOptionsOutput {
	return pulumi.ToOutput(e).(CleanupOptionsOutput)
}

func (e CleanupOptions) ToCleanupOptionsOutputWithContext(ctx context.Context) CleanupOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CleanupOptionsOutput)
}

func (e CleanupOptions) ToCleanupOptionsPtrOutput() CleanupOptionsPtrOutput {
	return e.ToCleanupOptionsPtrOutputWithContext(context.Background())
}

func (e CleanupOptions) ToCleanupOptionsPtrOutputWithContext(ctx context.Context) CleanupOptionsPtrOutput {
	return CleanupOptions(e).ToCleanupOptionsOutputWithContext(ctx).ToCleanupOptionsPtrOutputWithContext(ctx)
}

func (e CleanupOptions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CleanupOptions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CleanupOptions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CleanupOptions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CleanupOptionsOutput struct{ *pulumi.OutputState }

func (CleanupOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CleanupOptions)(nil)).Elem()
}

func (o CleanupOptionsOutput) ToCleanupOptionsOutput() CleanupOptionsOutput {
	return o
}

func (o CleanupOptionsOutput) ToCleanupOptionsOutputWithContext(ctx context.Context) CleanupOptionsOutput {
	return o
}

func (o CleanupOptionsOutput) ToCleanupOptionsPtrOutput() CleanupOptionsPtrOutput {
	return o.ToCleanupOptionsPtrOutputWithContext(context.Background())
}

func (o CleanupOptionsOutput) ToCleanupOptionsPtrOutputWithContext(ctx context.Context) CleanupOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CleanupOptions) *CleanupOptions {
		return &v
	}).(CleanupOptionsPtrOutput)
}

func (o CleanupOptionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CleanupOptionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CleanupOptions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CleanupOptionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CleanupOptionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CleanupOptions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CleanupOptionsPtrOutput struct{ *pulumi.OutputState }

func (CleanupOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CleanupOptions)(nil)).Elem()
}

func (o CleanupOptionsPtrOutput) ToCleanupOptionsPtrOutput() CleanupOptionsPtrOutput {
	return o
}

func (o CleanupOptionsPtrOutput) ToCleanupOptionsPtrOutputWithContext(ctx context.Context) CleanupOptionsPtrOutput {
	return o
}

func (o CleanupOptionsPtrOutput) Elem() CleanupOptionsOutput {
	return o.ApplyT(func(v *CleanupOptions) CleanupOptions {
		if v != nil {
			return *v
		}
		var ret CleanupOptions
		return ret
	}).(CleanupOptionsOutput)
}

func (o CleanupOptionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CleanupOptionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CleanupOptions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CleanupOptionsInput is an input type that accepts values of the CleanupOptions enum
// A concrete instance of `CleanupOptionsInput` can be one of the following:
//
//	CleanupOptionsAlways
//	CleanupOptionsOnSuccess
//	CleanupOptionsOnExpiration
type CleanupOptionsInput interface {
	pulumi.Input

	ToCleanupOptionsOutput() CleanupOptionsOutput
	ToCleanupOptionsOutputWithContext(context.Context) CleanupOptionsOutput
}

var cleanupOptionsPtrType = reflect.TypeOf((**CleanupOptions)(nil)).Elem()

type CleanupOptionsPtrInput interface {
	pulumi.Input

	ToCleanupOptionsPtrOutput() CleanupOptionsPtrOutput
	ToCleanupOptionsPtrOutputWithContext(context.Context) CleanupOptionsPtrOutput
}

type cleanupOptionsPtr string

func CleanupOptionsPtr(v string) CleanupOptionsPtrInput {
	return (*cleanupOptionsPtr)(&v)
}

func (*cleanupOptionsPtr) ElementType() reflect.Type {
	return cleanupOptionsPtrType
}

func (in *cleanupOptionsPtr) ToCleanupOptionsPtrOutput() CleanupOptionsPtrOutput {
	return pulumi.ToOutput(in).(CleanupOptionsPtrOutput)
}

func (in *cleanupOptionsPtr) ToCleanupOptionsPtrOutputWithContext(ctx context.Context) CleanupOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CleanupOptionsPtrOutput)
}

// Type of the managed identity.
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeUserAssigned = ManagedServiceIdentityType("UserAssigned")
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

// Type of the script.
type ScriptType string

const (
	ScriptTypeAzurePowerShell = ScriptType("AzurePowerShell")
	ScriptTypeAzureCLI        = ScriptType("AzureCLI")
)

func init() {
	pulumi.RegisterOutputType(CleanupOptionsOutput{})
	pulumi.RegisterOutputType(CleanupOptionsPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypePtrOutput{})
}
