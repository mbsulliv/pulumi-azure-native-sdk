// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Auto-tracking configuration.
type AutoTrackingConfiguration string

const (
	AutoTrackingConfigurationDisabled = AutoTrackingConfiguration("disabled")
	AutoTrackingConfigurationXBand    = AutoTrackingConfiguration("xBand")
	AutoTrackingConfigurationSBand    = AutoTrackingConfiguration("sBand")
)

func (AutoTrackingConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoTrackingConfiguration)(nil)).Elem()
}

func (e AutoTrackingConfiguration) ToAutoTrackingConfigurationOutput() AutoTrackingConfigurationOutput {
	return pulumi.ToOutput(e).(AutoTrackingConfigurationOutput)
}

func (e AutoTrackingConfiguration) ToAutoTrackingConfigurationOutputWithContext(ctx context.Context) AutoTrackingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoTrackingConfigurationOutput)
}

func (e AutoTrackingConfiguration) ToAutoTrackingConfigurationPtrOutput() AutoTrackingConfigurationPtrOutput {
	return e.ToAutoTrackingConfigurationPtrOutputWithContext(context.Background())
}

func (e AutoTrackingConfiguration) ToAutoTrackingConfigurationPtrOutputWithContext(ctx context.Context) AutoTrackingConfigurationPtrOutput {
	return AutoTrackingConfiguration(e).ToAutoTrackingConfigurationOutputWithContext(ctx).ToAutoTrackingConfigurationPtrOutputWithContext(ctx)
}

func (e AutoTrackingConfiguration) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoTrackingConfiguration) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoTrackingConfiguration) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoTrackingConfiguration) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoTrackingConfigurationOutput struct{ *pulumi.OutputState }

func (AutoTrackingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoTrackingConfiguration)(nil)).Elem()
}

func (o AutoTrackingConfigurationOutput) ToAutoTrackingConfigurationOutput() AutoTrackingConfigurationOutput {
	return o
}

func (o AutoTrackingConfigurationOutput) ToAutoTrackingConfigurationOutputWithContext(ctx context.Context) AutoTrackingConfigurationOutput {
	return o
}

func (o AutoTrackingConfigurationOutput) ToAutoTrackingConfigurationPtrOutput() AutoTrackingConfigurationPtrOutput {
	return o.ToAutoTrackingConfigurationPtrOutputWithContext(context.Background())
}

func (o AutoTrackingConfigurationOutput) ToAutoTrackingConfigurationPtrOutputWithContext(ctx context.Context) AutoTrackingConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoTrackingConfiguration) *AutoTrackingConfiguration {
		return &v
	}).(AutoTrackingConfigurationPtrOutput)
}

func (o AutoTrackingConfigurationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoTrackingConfigurationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoTrackingConfiguration) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoTrackingConfigurationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoTrackingConfigurationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoTrackingConfiguration) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoTrackingConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AutoTrackingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoTrackingConfiguration)(nil)).Elem()
}

func (o AutoTrackingConfigurationPtrOutput) ToAutoTrackingConfigurationPtrOutput() AutoTrackingConfigurationPtrOutput {
	return o
}

func (o AutoTrackingConfigurationPtrOutput) ToAutoTrackingConfigurationPtrOutputWithContext(ctx context.Context) AutoTrackingConfigurationPtrOutput {
	return o
}

func (o AutoTrackingConfigurationPtrOutput) Elem() AutoTrackingConfigurationOutput {
	return o.ApplyT(func(v *AutoTrackingConfiguration) AutoTrackingConfiguration {
		if v != nil {
			return *v
		}
		var ret AutoTrackingConfiguration
		return ret
	}).(AutoTrackingConfigurationOutput)
}

func (o AutoTrackingConfigurationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoTrackingConfigurationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoTrackingConfiguration) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AutoTrackingConfigurationInput is an input type that accepts AutoTrackingConfigurationArgs and AutoTrackingConfigurationOutput values.
// You can construct a concrete instance of `AutoTrackingConfigurationInput` via:
//
//	AutoTrackingConfigurationArgs{...}
type AutoTrackingConfigurationInput interface {
	pulumi.Input

	ToAutoTrackingConfigurationOutput() AutoTrackingConfigurationOutput
	ToAutoTrackingConfigurationOutputWithContext(context.Context) AutoTrackingConfigurationOutput
}

var autoTrackingConfigurationPtrType = reflect.TypeOf((**AutoTrackingConfiguration)(nil)).Elem()

type AutoTrackingConfigurationPtrInput interface {
	pulumi.Input

	ToAutoTrackingConfigurationPtrOutput() AutoTrackingConfigurationPtrOutput
	ToAutoTrackingConfigurationPtrOutputWithContext(context.Context) AutoTrackingConfigurationPtrOutput
}

type autoTrackingConfigurationPtr string

func AutoTrackingConfigurationPtr(v string) AutoTrackingConfigurationPtrInput {
	return (*autoTrackingConfigurationPtr)(&v)
}

func (*autoTrackingConfigurationPtr) ElementType() reflect.Type {
	return autoTrackingConfigurationPtrType
}

func (in *autoTrackingConfigurationPtr) ToAutoTrackingConfigurationPtrOutput() AutoTrackingConfigurationPtrOutput {
	return pulumi.ToOutput(in).(AutoTrackingConfigurationPtrOutput)
}

func (in *autoTrackingConfigurationPtr) ToAutoTrackingConfigurationPtrOutputWithContext(ctx context.Context) AutoTrackingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoTrackingConfigurationPtrOutput)
}

func (in *autoTrackingConfigurationPtr) ToOutput(ctx context.Context) pulumix.Output[*AutoTrackingConfiguration] {
	return pulumix.Output[*AutoTrackingConfiguration]{
		OutputState: in.ToAutoTrackingConfigurationPtrOutputWithContext(ctx).OutputState,
	}
}

// Direction (uplink or downlink).
type Direction string

const (
	DirectionUplink   = Direction("uplink")
	DirectionDownlink = Direction("downlink")
)

// Polarization. e.g. (RHCP, LHCP).
type Polarization string

const (
	PolarizationRHCP             = Polarization("RHCP")
	PolarizationLHCP             = Polarization("LHCP")
	PolarizationLinearVertical   = Polarization("linearVertical")
	PolarizationLinearHorizontal = Polarization("linearHorizontal")
)

// Protocol either UDP or TCP.
type Protocol string

const (
	ProtocolTCP = Protocol("TCP")
	ProtocolUDP = Protocol("UDP")
)

func init() {
	pulumi.RegisterOutputType(AutoTrackingConfigurationOutput{})
	pulumi.RegisterOutputType(AutoTrackingConfigurationPtrOutput{})
}
