// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Interval value in minutes used to create LogAnalytics call rate logs.
type IntervalInMins string

const (
	IntervalInMinsThreeMins  = IntervalInMins("ThreeMins")
	IntervalInMinsFiveMins   = IntervalInMins("FiveMins")
	IntervalInMinsThirtyMins = IntervalInMins("ThirtyMins")
	IntervalInMinsSixtyMins  = IntervalInMins("SixtyMins")
)

func (IntervalInMins) ElementType() reflect.Type {
	return reflect.TypeOf((*IntervalInMins)(nil)).Elem()
}

func (e IntervalInMins) ToIntervalInMinsOutput() IntervalInMinsOutput {
	return pulumi.ToOutput(e).(IntervalInMinsOutput)
}

func (e IntervalInMins) ToIntervalInMinsOutputWithContext(ctx context.Context) IntervalInMinsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IntervalInMinsOutput)
}

func (e IntervalInMins) ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput {
	return e.ToIntervalInMinsPtrOutputWithContext(context.Background())
}

func (e IntervalInMins) ToIntervalInMinsPtrOutputWithContext(ctx context.Context) IntervalInMinsPtrOutput {
	return IntervalInMins(e).ToIntervalInMinsOutputWithContext(ctx).ToIntervalInMinsPtrOutputWithContext(ctx)
}

func (e IntervalInMins) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntervalInMins) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntervalInMins) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IntervalInMins) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IntervalInMinsOutput struct{ *pulumi.OutputState }

func (IntervalInMinsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntervalInMins)(nil)).Elem()
}

func (o IntervalInMinsOutput) ToIntervalInMinsOutput() IntervalInMinsOutput {
	return o
}

func (o IntervalInMinsOutput) ToIntervalInMinsOutputWithContext(ctx context.Context) IntervalInMinsOutput {
	return o
}

func (o IntervalInMinsOutput) ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput {
	return o.ToIntervalInMinsPtrOutputWithContext(context.Background())
}

func (o IntervalInMinsOutput) ToIntervalInMinsPtrOutputWithContext(ctx context.Context) IntervalInMinsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntervalInMins) *IntervalInMins {
		return &v
	}).(IntervalInMinsPtrOutput)
}

func (o IntervalInMinsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IntervalInMinsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntervalInMins) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IntervalInMinsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntervalInMinsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntervalInMins) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IntervalInMinsPtrOutput struct{ *pulumi.OutputState }

func (IntervalInMinsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntervalInMins)(nil)).Elem()
}

func (o IntervalInMinsPtrOutput) ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput {
	return o
}

func (o IntervalInMinsPtrOutput) ToIntervalInMinsPtrOutputWithContext(ctx context.Context) IntervalInMinsPtrOutput {
	return o
}

func (o IntervalInMinsPtrOutput) Elem() IntervalInMinsOutput {
	return o.ApplyT(func(v *IntervalInMins) IntervalInMins {
		if v != nil {
			return *v
		}
		var ret IntervalInMins
		return ret
	}).(IntervalInMinsOutput)
}

func (o IntervalInMinsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntervalInMinsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IntervalInMins) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IntervalInMinsInput is an input type that accepts IntervalInMinsArgs and IntervalInMinsOutput values.
// You can construct a concrete instance of `IntervalInMinsInput` via:
//
//	IntervalInMinsArgs{...}
type IntervalInMinsInput interface {
	pulumi.Input

	ToIntervalInMinsOutput() IntervalInMinsOutput
	ToIntervalInMinsOutputWithContext(context.Context) IntervalInMinsOutput
}

var intervalInMinsPtrType = reflect.TypeOf((**IntervalInMins)(nil)).Elem()

type IntervalInMinsPtrInput interface {
	pulumi.Input

	ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput
	ToIntervalInMinsPtrOutputWithContext(context.Context) IntervalInMinsPtrOutput
}

type intervalInMinsPtr string

func IntervalInMinsPtr(v string) IntervalInMinsPtrInput {
	return (*intervalInMinsPtr)(&v)
}

func (*intervalInMinsPtr) ElementType() reflect.Type {
	return intervalInMinsPtrType
}

func (in *intervalInMinsPtr) ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput {
	return pulumi.ToOutput(in).(IntervalInMinsPtrOutput)
}

func (in *intervalInMinsPtr) ToIntervalInMinsPtrOutputWithContext(ctx context.Context) IntervalInMinsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IntervalInMinsPtrOutput)
}

func (in *intervalInMinsPtr) ToOutput(ctx context.Context) pulumix.Output[*IntervalInMins] {
	return pulumix.Output[*IntervalInMins]{
		OutputState: in.ToIntervalInMinsPtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterOutputType(IntervalInMinsOutput{})
	pulumi.RegisterOutputType(IntervalInMinsPtrOutput{})
}
