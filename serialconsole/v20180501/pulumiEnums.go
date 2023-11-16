// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Specifies whether the port is enabled for a serial console connection.
type SerialPortStateEnum string

const (
	SerialPortStateEnumEnabled  = SerialPortStateEnum("enabled")
	SerialPortStateEnumDisabled = SerialPortStateEnum("disabled")
)

func (SerialPortStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SerialPortStateEnum)(nil)).Elem()
}

func (e SerialPortStateEnum) ToSerialPortStateEnumOutput() SerialPortStateEnumOutput {
	return pulumi.ToOutput(e).(SerialPortStateEnumOutput)
}

func (e SerialPortStateEnum) ToSerialPortStateEnumOutputWithContext(ctx context.Context) SerialPortStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SerialPortStateEnumOutput)
}

func (e SerialPortStateEnum) ToSerialPortStateEnumPtrOutput() SerialPortStateEnumPtrOutput {
	return e.ToSerialPortStateEnumPtrOutputWithContext(context.Background())
}

func (e SerialPortStateEnum) ToSerialPortStateEnumPtrOutputWithContext(ctx context.Context) SerialPortStateEnumPtrOutput {
	return SerialPortStateEnum(e).ToSerialPortStateEnumOutputWithContext(ctx).ToSerialPortStateEnumPtrOutputWithContext(ctx)
}

func (e SerialPortStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SerialPortStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SerialPortStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SerialPortStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SerialPortStateEnumOutput struct{ *pulumi.OutputState }

func (SerialPortStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SerialPortStateEnum)(nil)).Elem()
}

func (o SerialPortStateEnumOutput) ToSerialPortStateEnumOutput() SerialPortStateEnumOutput {
	return o
}

func (o SerialPortStateEnumOutput) ToSerialPortStateEnumOutputWithContext(ctx context.Context) SerialPortStateEnumOutput {
	return o
}

func (o SerialPortStateEnumOutput) ToSerialPortStateEnumPtrOutput() SerialPortStateEnumPtrOutput {
	return o.ToSerialPortStateEnumPtrOutputWithContext(context.Background())
}

func (o SerialPortStateEnumOutput) ToSerialPortStateEnumPtrOutputWithContext(ctx context.Context) SerialPortStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SerialPortStateEnum) *SerialPortStateEnum {
		return &v
	}).(SerialPortStateEnumPtrOutput)
}

func (o SerialPortStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SerialPortStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SerialPortStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SerialPortStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SerialPortStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SerialPortStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SerialPortStateEnumPtrOutput struct{ *pulumi.OutputState }

func (SerialPortStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SerialPortStateEnum)(nil)).Elem()
}

func (o SerialPortStateEnumPtrOutput) ToSerialPortStateEnumPtrOutput() SerialPortStateEnumPtrOutput {
	return o
}

func (o SerialPortStateEnumPtrOutput) ToSerialPortStateEnumPtrOutputWithContext(ctx context.Context) SerialPortStateEnumPtrOutput {
	return o
}

func (o SerialPortStateEnumPtrOutput) Elem() SerialPortStateEnumOutput {
	return o.ApplyT(func(v *SerialPortStateEnum) SerialPortStateEnum {
		if v != nil {
			return *v
		}
		var ret SerialPortStateEnum
		return ret
	}).(SerialPortStateEnumOutput)
}

func (o SerialPortStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SerialPortStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SerialPortStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SerialPortStateEnumInput is an input type that accepts SerialPortStateEnumArgs and SerialPortStateEnumOutput values.
// You can construct a concrete instance of `SerialPortStateEnumInput` via:
//
//	SerialPortStateEnumArgs{...}
type SerialPortStateEnumInput interface {
	pulumi.Input

	ToSerialPortStateEnumOutput() SerialPortStateEnumOutput
	ToSerialPortStateEnumOutputWithContext(context.Context) SerialPortStateEnumOutput
}

var serialPortStateEnumPtrType = reflect.TypeOf((**SerialPortStateEnum)(nil)).Elem()

type SerialPortStateEnumPtrInput interface {
	pulumi.Input

	ToSerialPortStateEnumPtrOutput() SerialPortStateEnumPtrOutput
	ToSerialPortStateEnumPtrOutputWithContext(context.Context) SerialPortStateEnumPtrOutput
}

type serialPortStateEnumPtr string

func SerialPortStateEnumPtr(v string) SerialPortStateEnumPtrInput {
	return (*serialPortStateEnumPtr)(&v)
}

func (*serialPortStateEnumPtr) ElementType() reflect.Type {
	return serialPortStateEnumPtrType
}

func (in *serialPortStateEnumPtr) ToSerialPortStateEnumPtrOutput() SerialPortStateEnumPtrOutput {
	return pulumi.ToOutput(in).(SerialPortStateEnumPtrOutput)
}

func (in *serialPortStateEnumPtr) ToSerialPortStateEnumPtrOutputWithContext(ctx context.Context) SerialPortStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SerialPortStateEnumPtrOutput)
}

func (in *serialPortStateEnumPtr) ToOutput(ctx context.Context) pulumix.Output[*SerialPortStateEnum] {
	return pulumix.Output[*SerialPortStateEnum]{
		OutputState: in.ToSerialPortStateEnumPtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterOutputType(SerialPortStateEnumOutput{})
	pulumi.RegisterOutputType(SerialPortStateEnumPtrOutput{})
}
