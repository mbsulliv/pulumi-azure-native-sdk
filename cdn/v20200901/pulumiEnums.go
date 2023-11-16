// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Supported protocols for the customer's endpoint.
type AFDEndpointProtocols string

const (
	AFDEndpointProtocolsHttp  = AFDEndpointProtocols("Http")
	AFDEndpointProtocolsHttps = AFDEndpointProtocols("Https")
)

// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
type AfdQueryStringCachingBehavior string

const (
	AfdQueryStringCachingBehaviorIgnoreQueryString = AfdQueryStringCachingBehavior("IgnoreQueryString")
	AfdQueryStringCachingBehaviorUseQueryString    = AfdQueryStringCachingBehavior("UseQueryString")
	AfdQueryStringCachingBehaviorNotSet            = AfdQueryStringCachingBehavior("NotSet")
)

func (AfdQueryStringCachingBehavior) ElementType() reflect.Type {
	return reflect.TypeOf((*AfdQueryStringCachingBehavior)(nil)).Elem()
}

func (e AfdQueryStringCachingBehavior) ToAfdQueryStringCachingBehaviorOutput() AfdQueryStringCachingBehaviorOutput {
	return pulumi.ToOutput(e).(AfdQueryStringCachingBehaviorOutput)
}

func (e AfdQueryStringCachingBehavior) ToAfdQueryStringCachingBehaviorOutputWithContext(ctx context.Context) AfdQueryStringCachingBehaviorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AfdQueryStringCachingBehaviorOutput)
}

func (e AfdQueryStringCachingBehavior) ToAfdQueryStringCachingBehaviorPtrOutput() AfdQueryStringCachingBehaviorPtrOutput {
	return e.ToAfdQueryStringCachingBehaviorPtrOutputWithContext(context.Background())
}

func (e AfdQueryStringCachingBehavior) ToAfdQueryStringCachingBehaviorPtrOutputWithContext(ctx context.Context) AfdQueryStringCachingBehaviorPtrOutput {
	return AfdQueryStringCachingBehavior(e).ToAfdQueryStringCachingBehaviorOutputWithContext(ctx).ToAfdQueryStringCachingBehaviorPtrOutputWithContext(ctx)
}

func (e AfdQueryStringCachingBehavior) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AfdQueryStringCachingBehavior) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AfdQueryStringCachingBehavior) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AfdQueryStringCachingBehavior) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AfdQueryStringCachingBehaviorOutput struct{ *pulumi.OutputState }

func (AfdQueryStringCachingBehaviorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AfdQueryStringCachingBehavior)(nil)).Elem()
}

func (o AfdQueryStringCachingBehaviorOutput) ToAfdQueryStringCachingBehaviorOutput() AfdQueryStringCachingBehaviorOutput {
	return o
}

func (o AfdQueryStringCachingBehaviorOutput) ToAfdQueryStringCachingBehaviorOutputWithContext(ctx context.Context) AfdQueryStringCachingBehaviorOutput {
	return o
}

func (o AfdQueryStringCachingBehaviorOutput) ToAfdQueryStringCachingBehaviorPtrOutput() AfdQueryStringCachingBehaviorPtrOutput {
	return o.ToAfdQueryStringCachingBehaviorPtrOutputWithContext(context.Background())
}

func (o AfdQueryStringCachingBehaviorOutput) ToAfdQueryStringCachingBehaviorPtrOutputWithContext(ctx context.Context) AfdQueryStringCachingBehaviorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AfdQueryStringCachingBehavior) *AfdQueryStringCachingBehavior {
		return &v
	}).(AfdQueryStringCachingBehaviorPtrOutput)
}

func (o AfdQueryStringCachingBehaviorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AfdQueryStringCachingBehaviorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AfdQueryStringCachingBehavior) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AfdQueryStringCachingBehaviorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AfdQueryStringCachingBehaviorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AfdQueryStringCachingBehavior) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AfdQueryStringCachingBehaviorPtrOutput struct{ *pulumi.OutputState }

func (AfdQueryStringCachingBehaviorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AfdQueryStringCachingBehavior)(nil)).Elem()
}

func (o AfdQueryStringCachingBehaviorPtrOutput) ToAfdQueryStringCachingBehaviorPtrOutput() AfdQueryStringCachingBehaviorPtrOutput {
	return o
}

func (o AfdQueryStringCachingBehaviorPtrOutput) ToAfdQueryStringCachingBehaviorPtrOutputWithContext(ctx context.Context) AfdQueryStringCachingBehaviorPtrOutput {
	return o
}

func (o AfdQueryStringCachingBehaviorPtrOutput) Elem() AfdQueryStringCachingBehaviorOutput {
	return o.ApplyT(func(v *AfdQueryStringCachingBehavior) AfdQueryStringCachingBehavior {
		if v != nil {
			return *v
		}
		var ret AfdQueryStringCachingBehavior
		return ret
	}).(AfdQueryStringCachingBehaviorOutput)
}

func (o AfdQueryStringCachingBehaviorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AfdQueryStringCachingBehaviorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AfdQueryStringCachingBehavior) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AfdQueryStringCachingBehaviorInput is an input type that accepts AfdQueryStringCachingBehaviorArgs and AfdQueryStringCachingBehaviorOutput values.
// You can construct a concrete instance of `AfdQueryStringCachingBehaviorInput` via:
//
//	AfdQueryStringCachingBehaviorArgs{...}
type AfdQueryStringCachingBehaviorInput interface {
	pulumi.Input

	ToAfdQueryStringCachingBehaviorOutput() AfdQueryStringCachingBehaviorOutput
	ToAfdQueryStringCachingBehaviorOutputWithContext(context.Context) AfdQueryStringCachingBehaviorOutput
}

var afdQueryStringCachingBehaviorPtrType = reflect.TypeOf((**AfdQueryStringCachingBehavior)(nil)).Elem()

type AfdQueryStringCachingBehaviorPtrInput interface {
	pulumi.Input

	ToAfdQueryStringCachingBehaviorPtrOutput() AfdQueryStringCachingBehaviorPtrOutput
	ToAfdQueryStringCachingBehaviorPtrOutputWithContext(context.Context) AfdQueryStringCachingBehaviorPtrOutput
}

type afdQueryStringCachingBehaviorPtr string

func AfdQueryStringCachingBehaviorPtr(v string) AfdQueryStringCachingBehaviorPtrInput {
	return (*afdQueryStringCachingBehaviorPtr)(&v)
}

func (*afdQueryStringCachingBehaviorPtr) ElementType() reflect.Type {
	return afdQueryStringCachingBehaviorPtrType
}

func (in *afdQueryStringCachingBehaviorPtr) ToAfdQueryStringCachingBehaviorPtrOutput() AfdQueryStringCachingBehaviorPtrOutput {
	return pulumi.ToOutput(in).(AfdQueryStringCachingBehaviorPtrOutput)
}

func (in *afdQueryStringCachingBehaviorPtr) ToAfdQueryStringCachingBehaviorPtrOutputWithContext(ctx context.Context) AfdQueryStringCachingBehaviorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AfdQueryStringCachingBehaviorPtrOutput)
}

func (in *afdQueryStringCachingBehaviorPtr) ToOutput(ctx context.Context) pulumix.Output[*AfdQueryStringCachingBehavior] {
	return pulumix.Output[*AfdQueryStringCachingBehavior]{
		OutputState: in.ToAfdQueryStringCachingBehaviorPtrOutputWithContext(ctx).OutputState,
	}
}

// Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
type EnabledState string

const (
	EnabledStateEnabled  = EnabledState("Enabled")
	EnabledStateDisabled = EnabledState("Disabled")
)

// Protocol this rule will use when forwarding traffic to backends.
type ForwardingProtocol string

const (
	ForwardingProtocolHttpOnly     = ForwardingProtocol("HttpOnly")
	ForwardingProtocolHttpsOnly    = ForwardingProtocol("HttpsOnly")
	ForwardingProtocolMatchRequest = ForwardingProtocol("MatchRequest")
)

// The type of health probe request that is made.
type HealthProbeRequestType string

const (
	HealthProbeRequestTypeNotSet = HealthProbeRequestType("NotSet")
	HealthProbeRequestTypeGET    = HealthProbeRequestType("GET")
	HealthProbeRequestTypeHEAD   = HealthProbeRequestType("HEAD")
)

func (HealthProbeRequestType) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthProbeRequestType)(nil)).Elem()
}

func (e HealthProbeRequestType) ToHealthProbeRequestTypeOutput() HealthProbeRequestTypeOutput {
	return pulumi.ToOutput(e).(HealthProbeRequestTypeOutput)
}

func (e HealthProbeRequestType) ToHealthProbeRequestTypeOutputWithContext(ctx context.Context) HealthProbeRequestTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HealthProbeRequestTypeOutput)
}

func (e HealthProbeRequestType) ToHealthProbeRequestTypePtrOutput() HealthProbeRequestTypePtrOutput {
	return e.ToHealthProbeRequestTypePtrOutputWithContext(context.Background())
}

func (e HealthProbeRequestType) ToHealthProbeRequestTypePtrOutputWithContext(ctx context.Context) HealthProbeRequestTypePtrOutput {
	return HealthProbeRequestType(e).ToHealthProbeRequestTypeOutputWithContext(ctx).ToHealthProbeRequestTypePtrOutputWithContext(ctx)
}

func (e HealthProbeRequestType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HealthProbeRequestType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HealthProbeRequestType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HealthProbeRequestType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HealthProbeRequestTypeOutput struct{ *pulumi.OutputState }

func (HealthProbeRequestTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthProbeRequestType)(nil)).Elem()
}

func (o HealthProbeRequestTypeOutput) ToHealthProbeRequestTypeOutput() HealthProbeRequestTypeOutput {
	return o
}

func (o HealthProbeRequestTypeOutput) ToHealthProbeRequestTypeOutputWithContext(ctx context.Context) HealthProbeRequestTypeOutput {
	return o
}

func (o HealthProbeRequestTypeOutput) ToHealthProbeRequestTypePtrOutput() HealthProbeRequestTypePtrOutput {
	return o.ToHealthProbeRequestTypePtrOutputWithContext(context.Background())
}

func (o HealthProbeRequestTypeOutput) ToHealthProbeRequestTypePtrOutputWithContext(ctx context.Context) HealthProbeRequestTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HealthProbeRequestType) *HealthProbeRequestType {
		return &v
	}).(HealthProbeRequestTypePtrOutput)
}

func (o HealthProbeRequestTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HealthProbeRequestTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HealthProbeRequestType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HealthProbeRequestTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HealthProbeRequestTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HealthProbeRequestType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HealthProbeRequestTypePtrOutput struct{ *pulumi.OutputState }

func (HealthProbeRequestTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthProbeRequestType)(nil)).Elem()
}

func (o HealthProbeRequestTypePtrOutput) ToHealthProbeRequestTypePtrOutput() HealthProbeRequestTypePtrOutput {
	return o
}

func (o HealthProbeRequestTypePtrOutput) ToHealthProbeRequestTypePtrOutputWithContext(ctx context.Context) HealthProbeRequestTypePtrOutput {
	return o
}

func (o HealthProbeRequestTypePtrOutput) Elem() HealthProbeRequestTypeOutput {
	return o.ApplyT(func(v *HealthProbeRequestType) HealthProbeRequestType {
		if v != nil {
			return *v
		}
		var ret HealthProbeRequestType
		return ret
	}).(HealthProbeRequestTypeOutput)
}

func (o HealthProbeRequestTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HealthProbeRequestTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HealthProbeRequestType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// HealthProbeRequestTypeInput is an input type that accepts HealthProbeRequestTypeArgs and HealthProbeRequestTypeOutput values.
// You can construct a concrete instance of `HealthProbeRequestTypeInput` via:
//
//	HealthProbeRequestTypeArgs{...}
type HealthProbeRequestTypeInput interface {
	pulumi.Input

	ToHealthProbeRequestTypeOutput() HealthProbeRequestTypeOutput
	ToHealthProbeRequestTypeOutputWithContext(context.Context) HealthProbeRequestTypeOutput
}

var healthProbeRequestTypePtrType = reflect.TypeOf((**HealthProbeRequestType)(nil)).Elem()

type HealthProbeRequestTypePtrInput interface {
	pulumi.Input

	ToHealthProbeRequestTypePtrOutput() HealthProbeRequestTypePtrOutput
	ToHealthProbeRequestTypePtrOutputWithContext(context.Context) HealthProbeRequestTypePtrOutput
}

type healthProbeRequestTypePtr string

func HealthProbeRequestTypePtr(v string) HealthProbeRequestTypePtrInput {
	return (*healthProbeRequestTypePtr)(&v)
}

func (*healthProbeRequestTypePtr) ElementType() reflect.Type {
	return healthProbeRequestTypePtrType
}

func (in *healthProbeRequestTypePtr) ToHealthProbeRequestTypePtrOutput() HealthProbeRequestTypePtrOutput {
	return pulumi.ToOutput(in).(HealthProbeRequestTypePtrOutput)
}

func (in *healthProbeRequestTypePtr) ToHealthProbeRequestTypePtrOutputWithContext(ctx context.Context) HealthProbeRequestTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HealthProbeRequestTypePtrOutput)
}

func (in *healthProbeRequestTypePtr) ToOutput(ctx context.Context) pulumix.Output[*HealthProbeRequestType] {
	return pulumix.Output[*HealthProbeRequestType]{
		OutputState: in.ToHealthProbeRequestTypePtrOutputWithContext(ctx).OutputState,
	}
}

// Whether to automatically redirect HTTP traffic to HTTPS traffic. Note that this is a easy way to set up this rule and it will be the first rule that gets executed.
type HttpsRedirect string

const (
	HttpsRedirectEnabled  = HttpsRedirect("Enabled")
	HttpsRedirectDisabled = HttpsRedirect("Disabled")
)

// whether this route will be linked to the default endpoint domain.
type LinkToDefaultDomain string

const (
	LinkToDefaultDomainEnabled  = LinkToDefaultDomain("Enabled")
	LinkToDefaultDomainDisabled = LinkToDefaultDomain("Disabled")
)

// Protocol to use for health probe.
type ProbeProtocol string

const (
	ProbeProtocolNotSet = ProbeProtocol("NotSet")
	ProbeProtocolHttp   = ProbeProtocol("Http")
	ProbeProtocolHttps  = ProbeProtocol("Https")
)

func (ProbeProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*ProbeProtocol)(nil)).Elem()
}

func (e ProbeProtocol) ToProbeProtocolOutput() ProbeProtocolOutput {
	return pulumi.ToOutput(e).(ProbeProtocolOutput)
}

func (e ProbeProtocol) ToProbeProtocolOutputWithContext(ctx context.Context) ProbeProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProbeProtocolOutput)
}

func (e ProbeProtocol) ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput {
	return e.ToProbeProtocolPtrOutputWithContext(context.Background())
}

func (e ProbeProtocol) ToProbeProtocolPtrOutputWithContext(ctx context.Context) ProbeProtocolPtrOutput {
	return ProbeProtocol(e).ToProbeProtocolOutputWithContext(ctx).ToProbeProtocolPtrOutputWithContext(ctx)
}

func (e ProbeProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProbeProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProbeProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProbeProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProbeProtocolOutput struct{ *pulumi.OutputState }

func (ProbeProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProbeProtocol)(nil)).Elem()
}

func (o ProbeProtocolOutput) ToProbeProtocolOutput() ProbeProtocolOutput {
	return o
}

func (o ProbeProtocolOutput) ToProbeProtocolOutputWithContext(ctx context.Context) ProbeProtocolOutput {
	return o
}

func (o ProbeProtocolOutput) ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput {
	return o.ToProbeProtocolPtrOutputWithContext(context.Background())
}

func (o ProbeProtocolOutput) ToProbeProtocolPtrOutputWithContext(ctx context.Context) ProbeProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProbeProtocol) *ProbeProtocol {
		return &v
	}).(ProbeProtocolPtrOutput)
}

func (o ProbeProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProbeProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProbeProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProbeProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProbeProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProbeProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProbeProtocolPtrOutput struct{ *pulumi.OutputState }

func (ProbeProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProbeProtocol)(nil)).Elem()
}

func (o ProbeProtocolPtrOutput) ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput {
	return o
}

func (o ProbeProtocolPtrOutput) ToProbeProtocolPtrOutputWithContext(ctx context.Context) ProbeProtocolPtrOutput {
	return o
}

func (o ProbeProtocolPtrOutput) Elem() ProbeProtocolOutput {
	return o.ApplyT(func(v *ProbeProtocol) ProbeProtocol {
		if v != nil {
			return *v
		}
		var ret ProbeProtocol
		return ret
	}).(ProbeProtocolOutput)
}

func (o ProbeProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProbeProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProbeProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ProbeProtocolInput is an input type that accepts ProbeProtocolArgs and ProbeProtocolOutput values.
// You can construct a concrete instance of `ProbeProtocolInput` via:
//
//	ProbeProtocolArgs{...}
type ProbeProtocolInput interface {
	pulumi.Input

	ToProbeProtocolOutput() ProbeProtocolOutput
	ToProbeProtocolOutputWithContext(context.Context) ProbeProtocolOutput
}

var probeProtocolPtrType = reflect.TypeOf((**ProbeProtocol)(nil)).Elem()

type ProbeProtocolPtrInput interface {
	pulumi.Input

	ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput
	ToProbeProtocolPtrOutputWithContext(context.Context) ProbeProtocolPtrOutput
}

type probeProtocolPtr string

func ProbeProtocolPtr(v string) ProbeProtocolPtrInput {
	return (*probeProtocolPtr)(&v)
}

func (*probeProtocolPtr) ElementType() reflect.Type {
	return probeProtocolPtrType
}

func (in *probeProtocolPtr) ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput {
	return pulumi.ToOutput(in).(ProbeProtocolPtrOutput)
}

func (in *probeProtocolPtr) ToProbeProtocolPtrOutputWithContext(ctx context.Context) ProbeProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProbeProtocolPtrOutput)
}

func (in *probeProtocolPtr) ToOutput(ctx context.Context) pulumix.Output[*ProbeProtocol] {
	return pulumix.Output[*ProbeProtocol]{
		OutputState: in.ToProbeProtocolPtrOutputWithContext(ctx).OutputState,
	}
}

// Type of response errors for real user requests for which origin will be deemed unhealthy
type ResponseBasedDetectedErrorTypes string

const (
	ResponseBasedDetectedErrorTypesNone             = ResponseBasedDetectedErrorTypes("None")
	ResponseBasedDetectedErrorTypesTcpErrorsOnly    = ResponseBasedDetectedErrorTypes("TcpErrorsOnly")
	ResponseBasedDetectedErrorTypesTcpAndHttpErrors = ResponseBasedDetectedErrorTypes("TcpAndHttpErrors")
)

func (ResponseBasedDetectedErrorTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseBasedDetectedErrorTypes)(nil)).Elem()
}

func (e ResponseBasedDetectedErrorTypes) ToResponseBasedDetectedErrorTypesOutput() ResponseBasedDetectedErrorTypesOutput {
	return pulumi.ToOutput(e).(ResponseBasedDetectedErrorTypesOutput)
}

func (e ResponseBasedDetectedErrorTypes) ToResponseBasedDetectedErrorTypesOutputWithContext(ctx context.Context) ResponseBasedDetectedErrorTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResponseBasedDetectedErrorTypesOutput)
}

func (e ResponseBasedDetectedErrorTypes) ToResponseBasedDetectedErrorTypesPtrOutput() ResponseBasedDetectedErrorTypesPtrOutput {
	return e.ToResponseBasedDetectedErrorTypesPtrOutputWithContext(context.Background())
}

func (e ResponseBasedDetectedErrorTypes) ToResponseBasedDetectedErrorTypesPtrOutputWithContext(ctx context.Context) ResponseBasedDetectedErrorTypesPtrOutput {
	return ResponseBasedDetectedErrorTypes(e).ToResponseBasedDetectedErrorTypesOutputWithContext(ctx).ToResponseBasedDetectedErrorTypesPtrOutputWithContext(ctx)
}

func (e ResponseBasedDetectedErrorTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResponseBasedDetectedErrorTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResponseBasedDetectedErrorTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResponseBasedDetectedErrorTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResponseBasedDetectedErrorTypesOutput struct{ *pulumi.OutputState }

func (ResponseBasedDetectedErrorTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseBasedDetectedErrorTypes)(nil)).Elem()
}

func (o ResponseBasedDetectedErrorTypesOutput) ToResponseBasedDetectedErrorTypesOutput() ResponseBasedDetectedErrorTypesOutput {
	return o
}

func (o ResponseBasedDetectedErrorTypesOutput) ToResponseBasedDetectedErrorTypesOutputWithContext(ctx context.Context) ResponseBasedDetectedErrorTypesOutput {
	return o
}

func (o ResponseBasedDetectedErrorTypesOutput) ToResponseBasedDetectedErrorTypesPtrOutput() ResponseBasedDetectedErrorTypesPtrOutput {
	return o.ToResponseBasedDetectedErrorTypesPtrOutputWithContext(context.Background())
}

func (o ResponseBasedDetectedErrorTypesOutput) ToResponseBasedDetectedErrorTypesPtrOutputWithContext(ctx context.Context) ResponseBasedDetectedErrorTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResponseBasedDetectedErrorTypes) *ResponseBasedDetectedErrorTypes {
		return &v
	}).(ResponseBasedDetectedErrorTypesPtrOutput)
}

func (o ResponseBasedDetectedErrorTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResponseBasedDetectedErrorTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResponseBasedDetectedErrorTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResponseBasedDetectedErrorTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResponseBasedDetectedErrorTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResponseBasedDetectedErrorTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResponseBasedDetectedErrorTypesPtrOutput struct{ *pulumi.OutputState }

func (ResponseBasedDetectedErrorTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponseBasedDetectedErrorTypes)(nil)).Elem()
}

func (o ResponseBasedDetectedErrorTypesPtrOutput) ToResponseBasedDetectedErrorTypesPtrOutput() ResponseBasedDetectedErrorTypesPtrOutput {
	return o
}

func (o ResponseBasedDetectedErrorTypesPtrOutput) ToResponseBasedDetectedErrorTypesPtrOutputWithContext(ctx context.Context) ResponseBasedDetectedErrorTypesPtrOutput {
	return o
}

func (o ResponseBasedDetectedErrorTypesPtrOutput) Elem() ResponseBasedDetectedErrorTypesOutput {
	return o.ApplyT(func(v *ResponseBasedDetectedErrorTypes) ResponseBasedDetectedErrorTypes {
		if v != nil {
			return *v
		}
		var ret ResponseBasedDetectedErrorTypes
		return ret
	}).(ResponseBasedDetectedErrorTypesOutput)
}

func (o ResponseBasedDetectedErrorTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResponseBasedDetectedErrorTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResponseBasedDetectedErrorTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ResponseBasedDetectedErrorTypesInput is an input type that accepts ResponseBasedDetectedErrorTypesArgs and ResponseBasedDetectedErrorTypesOutput values.
// You can construct a concrete instance of `ResponseBasedDetectedErrorTypesInput` via:
//
//	ResponseBasedDetectedErrorTypesArgs{...}
type ResponseBasedDetectedErrorTypesInput interface {
	pulumi.Input

	ToResponseBasedDetectedErrorTypesOutput() ResponseBasedDetectedErrorTypesOutput
	ToResponseBasedDetectedErrorTypesOutputWithContext(context.Context) ResponseBasedDetectedErrorTypesOutput
}

var responseBasedDetectedErrorTypesPtrType = reflect.TypeOf((**ResponseBasedDetectedErrorTypes)(nil)).Elem()

type ResponseBasedDetectedErrorTypesPtrInput interface {
	pulumi.Input

	ToResponseBasedDetectedErrorTypesPtrOutput() ResponseBasedDetectedErrorTypesPtrOutput
	ToResponseBasedDetectedErrorTypesPtrOutputWithContext(context.Context) ResponseBasedDetectedErrorTypesPtrOutput
}

type responseBasedDetectedErrorTypesPtr string

func ResponseBasedDetectedErrorTypesPtr(v string) ResponseBasedDetectedErrorTypesPtrInput {
	return (*responseBasedDetectedErrorTypesPtr)(&v)
}

func (*responseBasedDetectedErrorTypesPtr) ElementType() reflect.Type {
	return responseBasedDetectedErrorTypesPtrType
}

func (in *responseBasedDetectedErrorTypesPtr) ToResponseBasedDetectedErrorTypesPtrOutput() ResponseBasedDetectedErrorTypesPtrOutput {
	return pulumi.ToOutput(in).(ResponseBasedDetectedErrorTypesPtrOutput)
}

func (in *responseBasedDetectedErrorTypesPtr) ToResponseBasedDetectedErrorTypesPtrOutputWithContext(ctx context.Context) ResponseBasedDetectedErrorTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResponseBasedDetectedErrorTypesPtrOutput)
}

func (in *responseBasedDetectedErrorTypesPtr) ToOutput(ctx context.Context) pulumix.Output[*ResponseBasedDetectedErrorTypes] {
	return pulumix.Output[*ResponseBasedDetectedErrorTypes]{
		OutputState: in.ToResponseBasedDetectedErrorTypesPtrOutputWithContext(ctx).OutputState,
	}
}

// Name of the pricing tier.
type SkuName string

const (
	SkuName_Standard_Verizon                   = SkuName("Standard_Verizon")
	SkuName_Premium_Verizon                    = SkuName("Premium_Verizon")
	SkuName_Custom_Verizon                     = SkuName("Custom_Verizon")
	SkuName_Standard_Akamai                    = SkuName("Standard_Akamai")
	SkuName_Standard_ChinaCdn                  = SkuName("Standard_ChinaCdn")
	SkuName_Standard_Microsoft                 = SkuName("Standard_Microsoft")
	SkuName_Premium_ChinaCdn                   = SkuName("Premium_ChinaCdn")
	SkuName_Standard_AzureFrontDoor            = SkuName("Standard_AzureFrontDoor")
	SkuName_Premium_AzureFrontDoor             = SkuName("Premium_AzureFrontDoor")
	SkuName_Standard_955BandWidth_ChinaCdn     = SkuName("Standard_955BandWidth_ChinaCdn")
	SkuName_Standard_AvgBandWidth_ChinaCdn     = SkuName("Standard_AvgBandWidth_ChinaCdn")
	SkuName_StandardPlus_ChinaCdn              = SkuName("StandardPlus_ChinaCdn")
	SkuName_StandardPlus_955BandWidth_ChinaCdn = SkuName("StandardPlus_955BandWidth_ChinaCdn")
	SkuName_StandardPlus_AvgBandWidth_ChinaCdn = SkuName("StandardPlus_AvgBandWidth_ChinaCdn")
)

func init() {
	pulumi.RegisterOutputType(AfdQueryStringCachingBehaviorOutput{})
	pulumi.RegisterOutputType(AfdQueryStringCachingBehaviorPtrOutput{})
	pulumi.RegisterOutputType(HealthProbeRequestTypeOutput{})
	pulumi.RegisterOutputType(HealthProbeRequestTypePtrOutput{})
	pulumi.RegisterOutputType(ProbeProtocolOutput{})
	pulumi.RegisterOutputType(ProbeProtocolPtrOutput{})
	pulumi.RegisterOutputType(ResponseBasedDetectedErrorTypesOutput{})
	pulumi.RegisterOutputType(ResponseBasedDetectedErrorTypesPtrOutput{})
}
