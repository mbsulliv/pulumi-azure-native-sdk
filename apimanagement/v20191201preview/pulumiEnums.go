// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Logger type.
type LoggerType string

const (
	// Azure Event Hub as log destination.
	LoggerTypeAzureEventHub = LoggerType("azureEventHub")
	// Azure Application Insights as log destination.
	LoggerTypeApplicationInsights = LoggerType("applicationInsights")
)

func (LoggerType) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggerType)(nil)).Elem()
}

func (e LoggerType) ToLoggerTypeOutput() LoggerTypeOutput {
	return pulumi.ToOutput(e).(LoggerTypeOutput)
}

func (e LoggerType) ToLoggerTypeOutputWithContext(ctx context.Context) LoggerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LoggerTypeOutput)
}

func (e LoggerType) ToLoggerTypePtrOutput() LoggerTypePtrOutput {
	return e.ToLoggerTypePtrOutputWithContext(context.Background())
}

func (e LoggerType) ToLoggerTypePtrOutputWithContext(ctx context.Context) LoggerTypePtrOutput {
	return LoggerType(e).ToLoggerTypeOutputWithContext(ctx).ToLoggerTypePtrOutputWithContext(ctx)
}

func (e LoggerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoggerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoggerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LoggerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LoggerTypeOutput struct{ *pulumi.OutputState }

func (LoggerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggerType)(nil)).Elem()
}

func (o LoggerTypeOutput) ToLoggerTypeOutput() LoggerTypeOutput {
	return o
}

func (o LoggerTypeOutput) ToLoggerTypeOutputWithContext(ctx context.Context) LoggerTypeOutput {
	return o
}

func (o LoggerTypeOutput) ToLoggerTypePtrOutput() LoggerTypePtrOutput {
	return o.ToLoggerTypePtrOutputWithContext(context.Background())
}

func (o LoggerTypeOutput) ToLoggerTypePtrOutputWithContext(ctx context.Context) LoggerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoggerType) *LoggerType {
		return &v
	}).(LoggerTypePtrOutput)
}

func (o LoggerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LoggerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoggerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LoggerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoggerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoggerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LoggerTypePtrOutput struct{ *pulumi.OutputState }

func (LoggerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggerType)(nil)).Elem()
}

func (o LoggerTypePtrOutput) ToLoggerTypePtrOutput() LoggerTypePtrOutput {
	return o
}

func (o LoggerTypePtrOutput) ToLoggerTypePtrOutputWithContext(ctx context.Context) LoggerTypePtrOutput {
	return o
}

func (o LoggerTypePtrOutput) Elem() LoggerTypeOutput {
	return o.ApplyT(func(v *LoggerType) LoggerType {
		if v != nil {
			return *v
		}
		var ret LoggerType
		return ret
	}).(LoggerTypeOutput)
}

func (o LoggerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoggerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LoggerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// LoggerTypeInput is an input type that accepts values of the LoggerType enum
// A concrete instance of `LoggerTypeInput` can be one of the following:
//
//	LoggerTypeAzureEventHub
//	LoggerTypeApplicationInsights
type LoggerTypeInput interface {
	pulumi.Input

	ToLoggerTypeOutput() LoggerTypeOutput
	ToLoggerTypeOutputWithContext(context.Context) LoggerTypeOutput
}

var loggerTypePtrType = reflect.TypeOf((**LoggerType)(nil)).Elem()

type LoggerTypePtrInput interface {
	pulumi.Input

	ToLoggerTypePtrOutput() LoggerTypePtrOutput
	ToLoggerTypePtrOutputWithContext(context.Context) LoggerTypePtrOutput
}

type loggerTypePtr string

func LoggerTypePtr(v string) LoggerTypePtrInput {
	return (*loggerTypePtr)(&v)
}

func (*loggerTypePtr) ElementType() reflect.Type {
	return loggerTypePtrType
}

func (in *loggerTypePtr) ToLoggerTypePtrOutput() LoggerTypePtrOutput {
	return pulumi.ToOutput(in).(LoggerTypePtrOutput)
}

func (in *loggerTypePtr) ToLoggerTypePtrOutputWithContext(ctx context.Context) LoggerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LoggerTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LoggerTypeOutput{})
	pulumi.RegisterOutputType(LoggerTypePtrOutput{})
}
