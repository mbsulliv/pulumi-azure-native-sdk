// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
type LogSettings struct {
	// Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
	Category *string `pulumi:"category"`
	// A value indicating whether this log is enabled.
	Enabled bool `pulumi:"enabled"`
	// The retention policy for this log.
	RetentionPolicy *RetentionPolicy `pulumi:"retentionPolicy"`
}

// LogSettingsInput is an input type that accepts LogSettingsArgs and LogSettingsOutput values.
// You can construct a concrete instance of `LogSettingsInput` via:
//
//	LogSettingsArgs{...}
type LogSettingsInput interface {
	pulumi.Input

	ToLogSettingsOutput() LogSettingsOutput
	ToLogSettingsOutputWithContext(context.Context) LogSettingsOutput
}

// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
type LogSettingsArgs struct {
	// Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
	Category pulumi.StringPtrInput `pulumi:"category"`
	// A value indicating whether this log is enabled.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	// The retention policy for this log.
	RetentionPolicy RetentionPolicyPtrInput `pulumi:"retentionPolicy"`
}

func (LogSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSettings)(nil)).Elem()
}

func (i LogSettingsArgs) ToLogSettingsOutput() LogSettingsOutput {
	return i.ToLogSettingsOutputWithContext(context.Background())
}

func (i LogSettingsArgs) ToLogSettingsOutputWithContext(ctx context.Context) LogSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSettingsOutput)
}

func (i LogSettingsArgs) ToOutput(ctx context.Context) pulumix.Output[LogSettings] {
	return pulumix.Output[LogSettings]{
		OutputState: i.ToLogSettingsOutputWithContext(ctx).OutputState,
	}
}

// LogSettingsArrayInput is an input type that accepts LogSettingsArray and LogSettingsArrayOutput values.
// You can construct a concrete instance of `LogSettingsArrayInput` via:
//
//	LogSettingsArray{ LogSettingsArgs{...} }
type LogSettingsArrayInput interface {
	pulumi.Input

	ToLogSettingsArrayOutput() LogSettingsArrayOutput
	ToLogSettingsArrayOutputWithContext(context.Context) LogSettingsArrayOutput
}

type LogSettingsArray []LogSettingsInput

func (LogSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogSettings)(nil)).Elem()
}

func (i LogSettingsArray) ToLogSettingsArrayOutput() LogSettingsArrayOutput {
	return i.ToLogSettingsArrayOutputWithContext(context.Background())
}

func (i LogSettingsArray) ToLogSettingsArrayOutputWithContext(ctx context.Context) LogSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSettingsArrayOutput)
}

func (i LogSettingsArray) ToOutput(ctx context.Context) pulumix.Output[[]LogSettings] {
	return pulumix.Output[[]LogSettings]{
		OutputState: i.ToLogSettingsArrayOutputWithContext(ctx).OutputState,
	}
}

// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
type LogSettingsOutput struct{ *pulumi.OutputState }

func (LogSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSettings)(nil)).Elem()
}

func (o LogSettingsOutput) ToLogSettingsOutput() LogSettingsOutput {
	return o
}

func (o LogSettingsOutput) ToLogSettingsOutputWithContext(ctx context.Context) LogSettingsOutput {
	return o
}

func (o LogSettingsOutput) ToOutput(ctx context.Context) pulumix.Output[LogSettings] {
	return pulumix.Output[LogSettings]{
		OutputState: o.OutputState,
	}
}

// Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
func (o LogSettingsOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogSettings) *string { return v.Category }).(pulumi.StringPtrOutput)
}

// A value indicating whether this log is enabled.
func (o LogSettingsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LogSettings) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The retention policy for this log.
func (o LogSettingsOutput) RetentionPolicy() RetentionPolicyPtrOutput {
	return o.ApplyT(func(v LogSettings) *RetentionPolicy { return v.RetentionPolicy }).(RetentionPolicyPtrOutput)
}

type LogSettingsArrayOutput struct{ *pulumi.OutputState }

func (LogSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogSettings)(nil)).Elem()
}

func (o LogSettingsArrayOutput) ToLogSettingsArrayOutput() LogSettingsArrayOutput {
	return o
}

func (o LogSettingsArrayOutput) ToLogSettingsArrayOutputWithContext(ctx context.Context) LogSettingsArrayOutput {
	return o
}

func (o LogSettingsArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]LogSettings] {
	return pulumix.Output[[]LogSettings]{
		OutputState: o.OutputState,
	}
}

func (o LogSettingsArrayOutput) Index(i pulumi.IntInput) LogSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogSettings {
		return vs[0].([]LogSettings)[vs[1].(int)]
	}).(LogSettingsOutput)
}

// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
type LogSettingsResponse struct {
	// Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
	Category *string `pulumi:"category"`
	// A value indicating whether this log is enabled.
	Enabled bool `pulumi:"enabled"`
	// The retention policy for this log.
	RetentionPolicy *RetentionPolicyResponse `pulumi:"retentionPolicy"`
}

// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
type LogSettingsResponseOutput struct{ *pulumi.OutputState }

func (LogSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSettingsResponse)(nil)).Elem()
}

func (o LogSettingsResponseOutput) ToLogSettingsResponseOutput() LogSettingsResponseOutput {
	return o
}

func (o LogSettingsResponseOutput) ToLogSettingsResponseOutputWithContext(ctx context.Context) LogSettingsResponseOutput {
	return o
}

func (o LogSettingsResponseOutput) ToOutput(ctx context.Context) pulumix.Output[LogSettingsResponse] {
	return pulumix.Output[LogSettingsResponse]{
		OutputState: o.OutputState,
	}
}

// Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
func (o LogSettingsResponseOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogSettingsResponse) *string { return v.Category }).(pulumi.StringPtrOutput)
}

// A value indicating whether this log is enabled.
func (o LogSettingsResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LogSettingsResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The retention policy for this log.
func (o LogSettingsResponseOutput) RetentionPolicy() RetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v LogSettingsResponse) *RetentionPolicyResponse { return v.RetentionPolicy }).(RetentionPolicyResponsePtrOutput)
}

type LogSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (LogSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogSettingsResponse)(nil)).Elem()
}

func (o LogSettingsResponseArrayOutput) ToLogSettingsResponseArrayOutput() LogSettingsResponseArrayOutput {
	return o
}

func (o LogSettingsResponseArrayOutput) ToLogSettingsResponseArrayOutputWithContext(ctx context.Context) LogSettingsResponseArrayOutput {
	return o
}

func (o LogSettingsResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]LogSettingsResponse] {
	return pulumix.Output[[]LogSettingsResponse]{
		OutputState: o.OutputState,
	}
}

func (o LogSettingsResponseArrayOutput) Index(i pulumi.IntInput) LogSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogSettingsResponse {
		return vs[0].([]LogSettingsResponse)[vs[1].(int)]
	}).(LogSettingsResponseOutput)
}

// Specifies the retention policy for the log.
type RetentionPolicy struct {
	// The number of days for the retention in days. A value of 0 will retain the events indefinitely.
	Days int `pulumi:"days"`
	// A value indicating whether the retention policy is enabled.
	Enabled bool `pulumi:"enabled"`
}

// RetentionPolicyInput is an input type that accepts RetentionPolicyArgs and RetentionPolicyOutput values.
// You can construct a concrete instance of `RetentionPolicyInput` via:
//
//	RetentionPolicyArgs{...}
type RetentionPolicyInput interface {
	pulumi.Input

	ToRetentionPolicyOutput() RetentionPolicyOutput
	ToRetentionPolicyOutputWithContext(context.Context) RetentionPolicyOutput
}

// Specifies the retention policy for the log.
type RetentionPolicyArgs struct {
	// The number of days for the retention in days. A value of 0 will retain the events indefinitely.
	Days pulumi.IntInput `pulumi:"days"`
	// A value indicating whether the retention policy is enabled.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (RetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return i.ToRetentionPolicyOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput)
}

func (i RetentionPolicyArgs) ToOutput(ctx context.Context) pulumix.Output[RetentionPolicy] {
	return pulumix.Output[RetentionPolicy]{
		OutputState: i.ToRetentionPolicyOutputWithContext(ctx).OutputState,
	}
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput).ToRetentionPolicyPtrOutputWithContext(ctx)
}

// RetentionPolicyPtrInput is an input type that accepts RetentionPolicyArgs, RetentionPolicyPtr and RetentionPolicyPtrOutput values.
// You can construct a concrete instance of `RetentionPolicyPtrInput` via:
//
//	        RetentionPolicyArgs{...}
//
//	or:
//
//	        nil
type RetentionPolicyPtrInput interface {
	pulumi.Input

	ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput
	ToRetentionPolicyPtrOutputWithContext(context.Context) RetentionPolicyPtrOutput
}

type retentionPolicyPtrType RetentionPolicyArgs

func RetentionPolicyPtr(v *RetentionPolicyArgs) RetentionPolicyPtrInput {
	return (*retentionPolicyPtrType)(v)
}

func (*retentionPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyPtrOutput)
}

func (i *retentionPolicyPtrType) ToOutput(ctx context.Context) pulumix.Output[*RetentionPolicy] {
	return pulumix.Output[*RetentionPolicy]{
		OutputState: i.ToRetentionPolicyPtrOutputWithContext(ctx).OutputState,
	}
}

// Specifies the retention policy for the log.
type RetentionPolicyOutput struct{ *pulumi.OutputState }

func (RetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionPolicy) *RetentionPolicy {
		return &v
	}).(RetentionPolicyPtrOutput)
}

func (o RetentionPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[RetentionPolicy] {
	return pulumix.Output[RetentionPolicy]{
		OutputState: o.OutputState,
	}
}

// The number of days for the retention in days. A value of 0 will retain the events indefinitely.
func (o RetentionPolicyOutput) Days() pulumi.IntOutput {
	return o.ApplyT(func(v RetentionPolicy) int { return v.Days }).(pulumi.IntOutput)
}

// A value indicating whether the retention policy is enabled.
func (o RetentionPolicyOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RetentionPolicy) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type RetentionPolicyPtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*RetentionPolicy] {
	return pulumix.Output[*RetentionPolicy]{
		OutputState: o.OutputState,
	}
}

func (o RetentionPolicyPtrOutput) Elem() RetentionPolicyOutput {
	return o.ApplyT(func(v *RetentionPolicy) RetentionPolicy {
		if v != nil {
			return *v
		}
		var ret RetentionPolicy
		return ret
	}).(RetentionPolicyOutput)
}

// The number of days for the retention in days. A value of 0 will retain the events indefinitely.
func (o RetentionPolicyPtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.Days
	}).(pulumi.IntPtrOutput)
}

// A value indicating whether the retention policy is enabled.
func (o RetentionPolicyPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

// Specifies the retention policy for the log.
type RetentionPolicyResponse struct {
	// The number of days for the retention in days. A value of 0 will retain the events indefinitely.
	Days int `pulumi:"days"`
	// A value indicating whether the retention policy is enabled.
	Enabled bool `pulumi:"enabled"`
}

// Specifies the retention policy for the log.
type RetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutputWithContext(ctx context.Context) RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) ToOutput(ctx context.Context) pulumix.Output[RetentionPolicyResponse] {
	return pulumix.Output[RetentionPolicyResponse]{
		OutputState: o.OutputState,
	}
}

// The number of days for the retention in days. A value of 0 will retain the events indefinitely.
func (o RetentionPolicyResponseOutput) Days() pulumi.IntOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) int { return v.Days }).(pulumi.IntOutput)
}

// A value indicating whether the retention policy is enabled.
func (o RetentionPolicyResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type RetentionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*RetentionPolicyResponse] {
	return pulumix.Output[*RetentionPolicyResponse]{
		OutputState: o.OutputState,
	}
}

func (o RetentionPolicyResponsePtrOutput) Elem() RetentionPolicyResponseOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) RetentionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret RetentionPolicyResponse
		return ret
	}).(RetentionPolicyResponseOutput)
}

// The number of days for the retention in days. A value of 0 will retain the events indefinitely.
func (o RetentionPolicyResponsePtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Days
	}).(pulumi.IntPtrOutput)
}

// A value indicating whether the retention policy is enabled.
func (o RetentionPolicyResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LogSettingsOutput{})
	pulumi.RegisterOutputType(LogSettingsArrayOutput{})
	pulumi.RegisterOutputType(LogSettingsResponseOutput{})
	pulumi.RegisterOutputType(LogSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(RetentionPolicyOutput{})
	pulumi.RegisterOutputType(RetentionPolicyPtrOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponsePtrOutput{})
}
