// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
type SecurityAlertPolicyState string

const (
	SecurityAlertPolicyStateNew      = SecurityAlertPolicyState("New")
	SecurityAlertPolicyStateEnabled  = SecurityAlertPolicyState("Enabled")
	SecurityAlertPolicyStateDisabled = SecurityAlertPolicyState("Disabled")
)

func (SecurityAlertPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAlertPolicyState)(nil)).Elem()
}

func (e SecurityAlertPolicyState) ToSecurityAlertPolicyStateOutput() SecurityAlertPolicyStateOutput {
	return pulumi.ToOutput(e).(SecurityAlertPolicyStateOutput)
}

func (e SecurityAlertPolicyState) ToSecurityAlertPolicyStateOutputWithContext(ctx context.Context) SecurityAlertPolicyStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityAlertPolicyStateOutput)
}

func (e SecurityAlertPolicyState) ToSecurityAlertPolicyStatePtrOutput() SecurityAlertPolicyStatePtrOutput {
	return e.ToSecurityAlertPolicyStatePtrOutputWithContext(context.Background())
}

func (e SecurityAlertPolicyState) ToSecurityAlertPolicyStatePtrOutputWithContext(ctx context.Context) SecurityAlertPolicyStatePtrOutput {
	return SecurityAlertPolicyState(e).ToSecurityAlertPolicyStateOutputWithContext(ctx).ToSecurityAlertPolicyStatePtrOutputWithContext(ctx)
}

func (e SecurityAlertPolicyState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityAlertPolicyState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityAlertPolicyState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityAlertPolicyState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityAlertPolicyStateOutput struct{ *pulumi.OutputState }

func (SecurityAlertPolicyStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAlertPolicyState)(nil)).Elem()
}

func (o SecurityAlertPolicyStateOutput) ToSecurityAlertPolicyStateOutput() SecurityAlertPolicyStateOutput {
	return o
}

func (o SecurityAlertPolicyStateOutput) ToSecurityAlertPolicyStateOutputWithContext(ctx context.Context) SecurityAlertPolicyStateOutput {
	return o
}

func (o SecurityAlertPolicyStateOutput) ToSecurityAlertPolicyStatePtrOutput() SecurityAlertPolicyStatePtrOutput {
	return o.ToSecurityAlertPolicyStatePtrOutputWithContext(context.Background())
}

func (o SecurityAlertPolicyStateOutput) ToSecurityAlertPolicyStatePtrOutputWithContext(ctx context.Context) SecurityAlertPolicyStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAlertPolicyState) *SecurityAlertPolicyState {
		return &v
	}).(SecurityAlertPolicyStatePtrOutput)
}

func (o SecurityAlertPolicyStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityAlertPolicyStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityAlertPolicyState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityAlertPolicyStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityAlertPolicyStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityAlertPolicyState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityAlertPolicyStatePtrOutput struct{ *pulumi.OutputState }

func (SecurityAlertPolicyStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAlertPolicyState)(nil)).Elem()
}

func (o SecurityAlertPolicyStatePtrOutput) ToSecurityAlertPolicyStatePtrOutput() SecurityAlertPolicyStatePtrOutput {
	return o
}

func (o SecurityAlertPolicyStatePtrOutput) ToSecurityAlertPolicyStatePtrOutputWithContext(ctx context.Context) SecurityAlertPolicyStatePtrOutput {
	return o
}

func (o SecurityAlertPolicyStatePtrOutput) Elem() SecurityAlertPolicyStateOutput {
	return o.ApplyT(func(v *SecurityAlertPolicyState) SecurityAlertPolicyState {
		if v != nil {
			return *v
		}
		var ret SecurityAlertPolicyState
		return ret
	}).(SecurityAlertPolicyStateOutput)
}

func (o SecurityAlertPolicyStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityAlertPolicyStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityAlertPolicyState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SecurityAlertPolicyStateInput is an input type that accepts SecurityAlertPolicyStateArgs and SecurityAlertPolicyStateOutput values.
// You can construct a concrete instance of `SecurityAlertPolicyStateInput` via:
//
//	SecurityAlertPolicyStateArgs{...}
type SecurityAlertPolicyStateInput interface {
	pulumi.Input

	ToSecurityAlertPolicyStateOutput() SecurityAlertPolicyStateOutput
	ToSecurityAlertPolicyStateOutputWithContext(context.Context) SecurityAlertPolicyStateOutput
}

var securityAlertPolicyStatePtrType = reflect.TypeOf((**SecurityAlertPolicyState)(nil)).Elem()

type SecurityAlertPolicyStatePtrInput interface {
	pulumi.Input

	ToSecurityAlertPolicyStatePtrOutput() SecurityAlertPolicyStatePtrOutput
	ToSecurityAlertPolicyStatePtrOutputWithContext(context.Context) SecurityAlertPolicyStatePtrOutput
}

type securityAlertPolicyStatePtr string

func SecurityAlertPolicyStatePtr(v string) SecurityAlertPolicyStatePtrInput {
	return (*securityAlertPolicyStatePtr)(&v)
}

func (*securityAlertPolicyStatePtr) ElementType() reflect.Type {
	return securityAlertPolicyStatePtrType
}

func (in *securityAlertPolicyStatePtr) ToSecurityAlertPolicyStatePtrOutput() SecurityAlertPolicyStatePtrOutput {
	return pulumi.ToOutput(in).(SecurityAlertPolicyStatePtrOutput)
}

func (in *securityAlertPolicyStatePtr) ToSecurityAlertPolicyStatePtrOutputWithContext(ctx context.Context) SecurityAlertPolicyStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityAlertPolicyStatePtrOutput)
}

func (in *securityAlertPolicyStatePtr) ToOutput(ctx context.Context) pulumix.Output[*SecurityAlertPolicyState] {
	return pulumix.Output[*SecurityAlertPolicyState]{
		OutputState: in.ToSecurityAlertPolicyStatePtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterOutputType(SecurityAlertPolicyStateOutput{})
	pulumi.RegisterOutputType(SecurityAlertPolicyStatePtrOutput{})
}
