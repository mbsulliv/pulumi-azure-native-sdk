// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Access mode for storage
type AccessMode string

const (
	AccessModeReadOnly  = AccessMode("ReadOnly")
	AccessModeReadWrite = AccessMode("ReadWrite")
)

// Allow or Deny rules to determine for incoming IP. Note: Rules can only consist of ALL Allow or ALL Deny
type Action string

const (
	ActionAllow = Action("Allow")
	ActionDeny  = Action("Deny")
)

// ActiveRevisionsMode controls how active revisions are handled for the Container app:
// <list><item>Multiple: multiple revisions can be active.</item><item>Single: Only one revision can be active at a time. Revision weights can not be used in this mode. If no value if provided, this is the default.</item></list>
type ActiveRevisionsMode string

const (
	ActiveRevisionsModeMultiple = ActiveRevisionsMode("Multiple")
	ActiveRevisionsModeSingle   = ActiveRevisionsMode("Single")
)

// Sticky Session Affinity
type Affinity string

const (
	AffinitySticky = Affinity("sticky")
	AffinityNone   = Affinity("none")
)

// Tells Dapr which protocol your application is using. Valid options are http and grpc. Default is http
type AppProtocol string

const (
	AppProtocolHttp = AppProtocol("http")
	AppProtocolGrpc = AppProtocol("grpc")
)

// Custom Domain binding type.
type BindingType string

const (
	BindingTypeDisabled   = BindingType("Disabled")
	BindingTypeSniEnabled = BindingType("SniEnabled")
)

// The method that should be used to authenticate the user.
type ClientCredentialMethod string

const (
	ClientCredentialMethodClientSecretPost = ClientCredentialMethod("ClientSecretPost")
)

func (ClientCredentialMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCredentialMethod)(nil)).Elem()
}

func (e ClientCredentialMethod) ToClientCredentialMethodOutput() ClientCredentialMethodOutput {
	return pulumi.ToOutput(e).(ClientCredentialMethodOutput)
}

func (e ClientCredentialMethod) ToClientCredentialMethodOutputWithContext(ctx context.Context) ClientCredentialMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClientCredentialMethodOutput)
}

func (e ClientCredentialMethod) ToClientCredentialMethodPtrOutput() ClientCredentialMethodPtrOutput {
	return e.ToClientCredentialMethodPtrOutputWithContext(context.Background())
}

func (e ClientCredentialMethod) ToClientCredentialMethodPtrOutputWithContext(ctx context.Context) ClientCredentialMethodPtrOutput {
	return ClientCredentialMethod(e).ToClientCredentialMethodOutputWithContext(ctx).ToClientCredentialMethodPtrOutputWithContext(ctx)
}

func (e ClientCredentialMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClientCredentialMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClientCredentialMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClientCredentialMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClientCredentialMethodOutput struct{ *pulumi.OutputState }

func (ClientCredentialMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCredentialMethod)(nil)).Elem()
}

func (o ClientCredentialMethodOutput) ToClientCredentialMethodOutput() ClientCredentialMethodOutput {
	return o
}

func (o ClientCredentialMethodOutput) ToClientCredentialMethodOutputWithContext(ctx context.Context) ClientCredentialMethodOutput {
	return o
}

func (o ClientCredentialMethodOutput) ToClientCredentialMethodPtrOutput() ClientCredentialMethodPtrOutput {
	return o.ToClientCredentialMethodPtrOutputWithContext(context.Background())
}

func (o ClientCredentialMethodOutput) ToClientCredentialMethodPtrOutputWithContext(ctx context.Context) ClientCredentialMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClientCredentialMethod) *ClientCredentialMethod {
		return &v
	}).(ClientCredentialMethodPtrOutput)
}

func (o ClientCredentialMethodOutput) ToOutput(ctx context.Context) pulumix.Output[ClientCredentialMethod] {
	return pulumix.Output[ClientCredentialMethod]{
		OutputState: o.OutputState,
	}
}

func (o ClientCredentialMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClientCredentialMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClientCredentialMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClientCredentialMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClientCredentialMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClientCredentialMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClientCredentialMethodPtrOutput struct{ *pulumi.OutputState }

func (ClientCredentialMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientCredentialMethod)(nil)).Elem()
}

func (o ClientCredentialMethodPtrOutput) ToClientCredentialMethodPtrOutput() ClientCredentialMethodPtrOutput {
	return o
}

func (o ClientCredentialMethodPtrOutput) ToClientCredentialMethodPtrOutputWithContext(ctx context.Context) ClientCredentialMethodPtrOutput {
	return o
}

func (o ClientCredentialMethodPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ClientCredentialMethod] {
	return pulumix.Output[*ClientCredentialMethod]{
		OutputState: o.OutputState,
	}
}

func (o ClientCredentialMethodPtrOutput) Elem() ClientCredentialMethodOutput {
	return o.ApplyT(func(v *ClientCredentialMethod) ClientCredentialMethod {
		if v != nil {
			return *v
		}
		var ret ClientCredentialMethod
		return ret
	}).(ClientCredentialMethodOutput)
}

func (o ClientCredentialMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClientCredentialMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClientCredentialMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ClientCredentialMethodInput is an input type that accepts ClientCredentialMethodArgs and ClientCredentialMethodOutput values.
// You can construct a concrete instance of `ClientCredentialMethodInput` via:
//
//	ClientCredentialMethodArgs{...}
type ClientCredentialMethodInput interface {
	pulumi.Input

	ToClientCredentialMethodOutput() ClientCredentialMethodOutput
	ToClientCredentialMethodOutputWithContext(context.Context) ClientCredentialMethodOutput
}

var clientCredentialMethodPtrType = reflect.TypeOf((**ClientCredentialMethod)(nil)).Elem()

type ClientCredentialMethodPtrInput interface {
	pulumi.Input

	ToClientCredentialMethodPtrOutput() ClientCredentialMethodPtrOutput
	ToClientCredentialMethodPtrOutputWithContext(context.Context) ClientCredentialMethodPtrOutput
}

type clientCredentialMethodPtr string

func ClientCredentialMethodPtr(v string) ClientCredentialMethodPtrInput {
	return (*clientCredentialMethodPtr)(&v)
}

func (*clientCredentialMethodPtr) ElementType() reflect.Type {
	return clientCredentialMethodPtrType
}

func (in *clientCredentialMethodPtr) ToClientCredentialMethodPtrOutput() ClientCredentialMethodPtrOutput {
	return pulumi.ToOutput(in).(ClientCredentialMethodPtrOutput)
}

func (in *clientCredentialMethodPtr) ToClientCredentialMethodPtrOutputWithContext(ctx context.Context) ClientCredentialMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClientCredentialMethodPtrOutput)
}

func (in *clientCredentialMethodPtr) ToOutput(ctx context.Context) pulumix.Output[*ClientCredentialMethod] {
	return pulumix.Output[*ClientCredentialMethod]{
		OutputState: in.ToClientCredentialMethodPtrOutputWithContext(ctx).OutputState,
	}
}

// The convention used when determining the session cookie's expiration.
type CookieExpirationConvention string

const (
	CookieExpirationConventionFixedTime               = CookieExpirationConvention("FixedTime")
	CookieExpirationConventionIdentityProviderDerived = CookieExpirationConvention("IdentityProviderDerived")
)

func (CookieExpirationConvention) ElementType() reflect.Type {
	return reflect.TypeOf((*CookieExpirationConvention)(nil)).Elem()
}

func (e CookieExpirationConvention) ToCookieExpirationConventionOutput() CookieExpirationConventionOutput {
	return pulumi.ToOutput(e).(CookieExpirationConventionOutput)
}

func (e CookieExpirationConvention) ToCookieExpirationConventionOutputWithContext(ctx context.Context) CookieExpirationConventionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CookieExpirationConventionOutput)
}

func (e CookieExpirationConvention) ToCookieExpirationConventionPtrOutput() CookieExpirationConventionPtrOutput {
	return e.ToCookieExpirationConventionPtrOutputWithContext(context.Background())
}

func (e CookieExpirationConvention) ToCookieExpirationConventionPtrOutputWithContext(ctx context.Context) CookieExpirationConventionPtrOutput {
	return CookieExpirationConvention(e).ToCookieExpirationConventionOutputWithContext(ctx).ToCookieExpirationConventionPtrOutputWithContext(ctx)
}

func (e CookieExpirationConvention) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CookieExpirationConvention) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CookieExpirationConvention) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CookieExpirationConvention) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CookieExpirationConventionOutput struct{ *pulumi.OutputState }

func (CookieExpirationConventionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CookieExpirationConvention)(nil)).Elem()
}

func (o CookieExpirationConventionOutput) ToCookieExpirationConventionOutput() CookieExpirationConventionOutput {
	return o
}

func (o CookieExpirationConventionOutput) ToCookieExpirationConventionOutputWithContext(ctx context.Context) CookieExpirationConventionOutput {
	return o
}

func (o CookieExpirationConventionOutput) ToCookieExpirationConventionPtrOutput() CookieExpirationConventionPtrOutput {
	return o.ToCookieExpirationConventionPtrOutputWithContext(context.Background())
}

func (o CookieExpirationConventionOutput) ToCookieExpirationConventionPtrOutputWithContext(ctx context.Context) CookieExpirationConventionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CookieExpirationConvention) *CookieExpirationConvention {
		return &v
	}).(CookieExpirationConventionPtrOutput)
}

func (o CookieExpirationConventionOutput) ToOutput(ctx context.Context) pulumix.Output[CookieExpirationConvention] {
	return pulumix.Output[CookieExpirationConvention]{
		OutputState: o.OutputState,
	}
}

func (o CookieExpirationConventionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CookieExpirationConventionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CookieExpirationConvention) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CookieExpirationConventionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CookieExpirationConventionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CookieExpirationConvention) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CookieExpirationConventionPtrOutput struct{ *pulumi.OutputState }

func (CookieExpirationConventionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CookieExpirationConvention)(nil)).Elem()
}

func (o CookieExpirationConventionPtrOutput) ToCookieExpirationConventionPtrOutput() CookieExpirationConventionPtrOutput {
	return o
}

func (o CookieExpirationConventionPtrOutput) ToCookieExpirationConventionPtrOutputWithContext(ctx context.Context) CookieExpirationConventionPtrOutput {
	return o
}

func (o CookieExpirationConventionPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*CookieExpirationConvention] {
	return pulumix.Output[*CookieExpirationConvention]{
		OutputState: o.OutputState,
	}
}

func (o CookieExpirationConventionPtrOutput) Elem() CookieExpirationConventionOutput {
	return o.ApplyT(func(v *CookieExpirationConvention) CookieExpirationConvention {
		if v != nil {
			return *v
		}
		var ret CookieExpirationConvention
		return ret
	}).(CookieExpirationConventionOutput)
}

func (o CookieExpirationConventionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CookieExpirationConventionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CookieExpirationConvention) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CookieExpirationConventionInput is an input type that accepts CookieExpirationConventionArgs and CookieExpirationConventionOutput values.
// You can construct a concrete instance of `CookieExpirationConventionInput` via:
//
//	CookieExpirationConventionArgs{...}
type CookieExpirationConventionInput interface {
	pulumi.Input

	ToCookieExpirationConventionOutput() CookieExpirationConventionOutput
	ToCookieExpirationConventionOutputWithContext(context.Context) CookieExpirationConventionOutput
}

var cookieExpirationConventionPtrType = reflect.TypeOf((**CookieExpirationConvention)(nil)).Elem()

type CookieExpirationConventionPtrInput interface {
	pulumi.Input

	ToCookieExpirationConventionPtrOutput() CookieExpirationConventionPtrOutput
	ToCookieExpirationConventionPtrOutputWithContext(context.Context) CookieExpirationConventionPtrOutput
}

type cookieExpirationConventionPtr string

func CookieExpirationConventionPtr(v string) CookieExpirationConventionPtrInput {
	return (*cookieExpirationConventionPtr)(&v)
}

func (*cookieExpirationConventionPtr) ElementType() reflect.Type {
	return cookieExpirationConventionPtrType
}

func (in *cookieExpirationConventionPtr) ToCookieExpirationConventionPtrOutput() CookieExpirationConventionPtrOutput {
	return pulumi.ToOutput(in).(CookieExpirationConventionPtrOutput)
}

func (in *cookieExpirationConventionPtr) ToCookieExpirationConventionPtrOutputWithContext(ctx context.Context) CookieExpirationConventionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CookieExpirationConventionPtrOutput)
}

func (in *cookieExpirationConventionPtr) ToOutput(ctx context.Context) pulumix.Output[*CookieExpirationConvention] {
	return pulumix.Output[*CookieExpirationConvention]{
		OutputState: in.ToCookieExpirationConventionPtrOutputWithContext(ctx).OutputState,
	}
}

// The type of the extended location.
type ExtendedLocationTypes string

const (
	ExtendedLocationTypesCustomLocation = ExtendedLocationTypes("CustomLocation")
)

// The convention used to determine the url of the request made.
type ForwardProxyConvention string

const (
	ForwardProxyConventionNoProxy  = ForwardProxyConvention("NoProxy")
	ForwardProxyConventionStandard = ForwardProxyConvention("Standard")
	ForwardProxyConventionCustom   = ForwardProxyConvention("Custom")
)

func (ForwardProxyConvention) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardProxyConvention)(nil)).Elem()
}

func (e ForwardProxyConvention) ToForwardProxyConventionOutput() ForwardProxyConventionOutput {
	return pulumi.ToOutput(e).(ForwardProxyConventionOutput)
}

func (e ForwardProxyConvention) ToForwardProxyConventionOutputWithContext(ctx context.Context) ForwardProxyConventionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ForwardProxyConventionOutput)
}

func (e ForwardProxyConvention) ToForwardProxyConventionPtrOutput() ForwardProxyConventionPtrOutput {
	return e.ToForwardProxyConventionPtrOutputWithContext(context.Background())
}

func (e ForwardProxyConvention) ToForwardProxyConventionPtrOutputWithContext(ctx context.Context) ForwardProxyConventionPtrOutput {
	return ForwardProxyConvention(e).ToForwardProxyConventionOutputWithContext(ctx).ToForwardProxyConventionPtrOutputWithContext(ctx)
}

func (e ForwardProxyConvention) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ForwardProxyConvention) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ForwardProxyConvention) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ForwardProxyConvention) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ForwardProxyConventionOutput struct{ *pulumi.OutputState }

func (ForwardProxyConventionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardProxyConvention)(nil)).Elem()
}

func (o ForwardProxyConventionOutput) ToForwardProxyConventionOutput() ForwardProxyConventionOutput {
	return o
}

func (o ForwardProxyConventionOutput) ToForwardProxyConventionOutputWithContext(ctx context.Context) ForwardProxyConventionOutput {
	return o
}

func (o ForwardProxyConventionOutput) ToForwardProxyConventionPtrOutput() ForwardProxyConventionPtrOutput {
	return o.ToForwardProxyConventionPtrOutputWithContext(context.Background())
}

func (o ForwardProxyConventionOutput) ToForwardProxyConventionPtrOutputWithContext(ctx context.Context) ForwardProxyConventionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ForwardProxyConvention) *ForwardProxyConvention {
		return &v
	}).(ForwardProxyConventionPtrOutput)
}

func (o ForwardProxyConventionOutput) ToOutput(ctx context.Context) pulumix.Output[ForwardProxyConvention] {
	return pulumix.Output[ForwardProxyConvention]{
		OutputState: o.OutputState,
	}
}

func (o ForwardProxyConventionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ForwardProxyConventionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ForwardProxyConvention) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ForwardProxyConventionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ForwardProxyConventionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ForwardProxyConvention) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ForwardProxyConventionPtrOutput struct{ *pulumi.OutputState }

func (ForwardProxyConventionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardProxyConvention)(nil)).Elem()
}

func (o ForwardProxyConventionPtrOutput) ToForwardProxyConventionPtrOutput() ForwardProxyConventionPtrOutput {
	return o
}

func (o ForwardProxyConventionPtrOutput) ToForwardProxyConventionPtrOutputWithContext(ctx context.Context) ForwardProxyConventionPtrOutput {
	return o
}

func (o ForwardProxyConventionPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ForwardProxyConvention] {
	return pulumix.Output[*ForwardProxyConvention]{
		OutputState: o.OutputState,
	}
}

func (o ForwardProxyConventionPtrOutput) Elem() ForwardProxyConventionOutput {
	return o.ApplyT(func(v *ForwardProxyConvention) ForwardProxyConvention {
		if v != nil {
			return *v
		}
		var ret ForwardProxyConvention
		return ret
	}).(ForwardProxyConventionOutput)
}

func (o ForwardProxyConventionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ForwardProxyConventionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ForwardProxyConvention) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ForwardProxyConventionInput is an input type that accepts ForwardProxyConventionArgs and ForwardProxyConventionOutput values.
// You can construct a concrete instance of `ForwardProxyConventionInput` via:
//
//	ForwardProxyConventionArgs{...}
type ForwardProxyConventionInput interface {
	pulumi.Input

	ToForwardProxyConventionOutput() ForwardProxyConventionOutput
	ToForwardProxyConventionOutputWithContext(context.Context) ForwardProxyConventionOutput
}

var forwardProxyConventionPtrType = reflect.TypeOf((**ForwardProxyConvention)(nil)).Elem()

type ForwardProxyConventionPtrInput interface {
	pulumi.Input

	ToForwardProxyConventionPtrOutput() ForwardProxyConventionPtrOutput
	ToForwardProxyConventionPtrOutputWithContext(context.Context) ForwardProxyConventionPtrOutput
}

type forwardProxyConventionPtr string

func ForwardProxyConventionPtr(v string) ForwardProxyConventionPtrInput {
	return (*forwardProxyConventionPtr)(&v)
}

func (*forwardProxyConventionPtr) ElementType() reflect.Type {
	return forwardProxyConventionPtrType
}

func (in *forwardProxyConventionPtr) ToForwardProxyConventionPtrOutput() ForwardProxyConventionPtrOutput {
	return pulumi.ToOutput(in).(ForwardProxyConventionPtrOutput)
}

func (in *forwardProxyConventionPtr) ToForwardProxyConventionPtrOutputWithContext(ctx context.Context) ForwardProxyConventionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ForwardProxyConventionPtrOutput)
}

func (in *forwardProxyConventionPtr) ToOutput(ctx context.Context) pulumix.Output[*ForwardProxyConvention] {
	return pulumix.Output[*ForwardProxyConvention]{
		OutputState: in.ToForwardProxyConventionPtrOutputWithContext(ctx).OutputState,
	}
}

// Client certificate mode for mTLS authentication. Ignore indicates server drops client certificate on forwarding. Accept indicates server forwards client certificate but does not require a client certificate. Require indicates server requires a client certificate.
type IngressClientCertificateMode string

const (
	IngressClientCertificateModeIgnore  = IngressClientCertificateMode("ignore")
	IngressClientCertificateModeAccept  = IngressClientCertificateMode("accept")
	IngressClientCertificateModeRequire = IngressClientCertificateMode("require")
)

// Ingress transport protocol
type IngressTransportMethod string

const (
	IngressTransportMethodAuto  = IngressTransportMethod("auto")
	IngressTransportMethodHttp  = IngressTransportMethod("http")
	IngressTransportMethodHttp2 = IngressTransportMethod("http2")
	IngressTransportMethodTcp   = IngressTransportMethod("tcp")
)

// Sets the log level for the Dapr sidecar. Allowed values are debug, info, warn, error. Default is info.
type LogLevel string

const (
	LogLevelInfo  = LogLevel("info")
	LogLevelDebug = LogLevel("debug")
	LogLevelWarn  = LogLevel("warn")
	LogLevelError = LogLevel("error")
)

// Selected type of domain control validation for managed certificates.
type ManagedCertificateDomainControlValidation string

const (
	ManagedCertificateDomainControlValidationCNAME = ManagedCertificateDomainControlValidation("CNAME")
	ManagedCertificateDomainControlValidationHTTP  = ManagedCertificateDomainControlValidation("HTTP")
	ManagedCertificateDomainControlValidationTXT   = ManagedCertificateDomainControlValidation("TXT")
)

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned,UserAssigned")
)

// Scheme to use for connecting to the host. Defaults to HTTP.
type Scheme string

const (
	SchemeHTTP  = Scheme("HTTP")
	SchemeHTTPS = Scheme("HTTPS")
)

// Storage type for the volume. If not provided, use EmptyDir.
type StorageType string

const (
	StorageTypeAzureFile = StorageType("AzureFile")
	StorageTypeEmptyDir  = StorageType("EmptyDir")
	StorageTypeSecret    = StorageType("Secret")
)

// Trigger type of the job
type TriggerType string

const (
	TriggerTypeSchedule = TriggerType("Schedule")
	TriggerTypeEvent    = TriggerType("Event")
	TriggerTypeManual   = TriggerType("Manual")
)

// The type of probe.
type Type string

const (
	TypeLiveness  = Type("Liveness")
	TypeReadiness = Type("Readiness")
	TypeStartup   = Type("Startup")
)

// The action to take when an unauthenticated client attempts to access the app.
type UnauthenticatedClientActionV2 string

const (
	UnauthenticatedClientActionV2RedirectToLoginPage = UnauthenticatedClientActionV2("RedirectToLoginPage")
	UnauthenticatedClientActionV2AllowAnonymous      = UnauthenticatedClientActionV2("AllowAnonymous")
	UnauthenticatedClientActionV2Return401           = UnauthenticatedClientActionV2("Return401")
	UnauthenticatedClientActionV2Return403           = UnauthenticatedClientActionV2("Return403")
)

func (UnauthenticatedClientActionV2) ElementType() reflect.Type {
	return reflect.TypeOf((*UnauthenticatedClientActionV2)(nil)).Elem()
}

func (e UnauthenticatedClientActionV2) ToUnauthenticatedClientActionV2Output() UnauthenticatedClientActionV2Output {
	return pulumi.ToOutput(e).(UnauthenticatedClientActionV2Output)
}

func (e UnauthenticatedClientActionV2) ToUnauthenticatedClientActionV2OutputWithContext(ctx context.Context) UnauthenticatedClientActionV2Output {
	return pulumi.ToOutputWithContext(ctx, e).(UnauthenticatedClientActionV2Output)
}

func (e UnauthenticatedClientActionV2) ToUnauthenticatedClientActionV2PtrOutput() UnauthenticatedClientActionV2PtrOutput {
	return e.ToUnauthenticatedClientActionV2PtrOutputWithContext(context.Background())
}

func (e UnauthenticatedClientActionV2) ToUnauthenticatedClientActionV2PtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionV2PtrOutput {
	return UnauthenticatedClientActionV2(e).ToUnauthenticatedClientActionV2OutputWithContext(ctx).ToUnauthenticatedClientActionV2PtrOutputWithContext(ctx)
}

func (e UnauthenticatedClientActionV2) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UnauthenticatedClientActionV2) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UnauthenticatedClientActionV2) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UnauthenticatedClientActionV2) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UnauthenticatedClientActionV2Output struct{ *pulumi.OutputState }

func (UnauthenticatedClientActionV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((*UnauthenticatedClientActionV2)(nil)).Elem()
}

func (o UnauthenticatedClientActionV2Output) ToUnauthenticatedClientActionV2Output() UnauthenticatedClientActionV2Output {
	return o
}

func (o UnauthenticatedClientActionV2Output) ToUnauthenticatedClientActionV2OutputWithContext(ctx context.Context) UnauthenticatedClientActionV2Output {
	return o
}

func (o UnauthenticatedClientActionV2Output) ToUnauthenticatedClientActionV2PtrOutput() UnauthenticatedClientActionV2PtrOutput {
	return o.ToUnauthenticatedClientActionV2PtrOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionV2Output) ToUnauthenticatedClientActionV2PtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionV2PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UnauthenticatedClientActionV2) *UnauthenticatedClientActionV2 {
		return &v
	}).(UnauthenticatedClientActionV2PtrOutput)
}

func (o UnauthenticatedClientActionV2Output) ToOutput(ctx context.Context) pulumix.Output[UnauthenticatedClientActionV2] {
	return pulumix.Output[UnauthenticatedClientActionV2]{
		OutputState: o.OutputState,
	}
}

func (o UnauthenticatedClientActionV2Output) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionV2Output) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UnauthenticatedClientActionV2) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UnauthenticatedClientActionV2Output) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionV2Output) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UnauthenticatedClientActionV2) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UnauthenticatedClientActionV2PtrOutput struct{ *pulumi.OutputState }

func (UnauthenticatedClientActionV2PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UnauthenticatedClientActionV2)(nil)).Elem()
}

func (o UnauthenticatedClientActionV2PtrOutput) ToUnauthenticatedClientActionV2PtrOutput() UnauthenticatedClientActionV2PtrOutput {
	return o
}

func (o UnauthenticatedClientActionV2PtrOutput) ToUnauthenticatedClientActionV2PtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionV2PtrOutput {
	return o
}

func (o UnauthenticatedClientActionV2PtrOutput) ToOutput(ctx context.Context) pulumix.Output[*UnauthenticatedClientActionV2] {
	return pulumix.Output[*UnauthenticatedClientActionV2]{
		OutputState: o.OutputState,
	}
}

func (o UnauthenticatedClientActionV2PtrOutput) Elem() UnauthenticatedClientActionV2Output {
	return o.ApplyT(func(v *UnauthenticatedClientActionV2) UnauthenticatedClientActionV2 {
		if v != nil {
			return *v
		}
		var ret UnauthenticatedClientActionV2
		return ret
	}).(UnauthenticatedClientActionV2Output)
}

func (o UnauthenticatedClientActionV2PtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionV2PtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UnauthenticatedClientActionV2) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// UnauthenticatedClientActionV2Input is an input type that accepts UnauthenticatedClientActionV2Args and UnauthenticatedClientActionV2Output values.
// You can construct a concrete instance of `UnauthenticatedClientActionV2Input` via:
//
//	UnauthenticatedClientActionV2Args{...}
type UnauthenticatedClientActionV2Input interface {
	pulumi.Input

	ToUnauthenticatedClientActionV2Output() UnauthenticatedClientActionV2Output
	ToUnauthenticatedClientActionV2OutputWithContext(context.Context) UnauthenticatedClientActionV2Output
}

var unauthenticatedClientActionV2PtrType = reflect.TypeOf((**UnauthenticatedClientActionV2)(nil)).Elem()

type UnauthenticatedClientActionV2PtrInput interface {
	pulumi.Input

	ToUnauthenticatedClientActionV2PtrOutput() UnauthenticatedClientActionV2PtrOutput
	ToUnauthenticatedClientActionV2PtrOutputWithContext(context.Context) UnauthenticatedClientActionV2PtrOutput
}

type unauthenticatedClientActionV2Ptr string

func UnauthenticatedClientActionV2Ptr(v string) UnauthenticatedClientActionV2PtrInput {
	return (*unauthenticatedClientActionV2Ptr)(&v)
}

func (*unauthenticatedClientActionV2Ptr) ElementType() reflect.Type {
	return unauthenticatedClientActionV2PtrType
}

func (in *unauthenticatedClientActionV2Ptr) ToUnauthenticatedClientActionV2PtrOutput() UnauthenticatedClientActionV2PtrOutput {
	return pulumi.ToOutput(in).(UnauthenticatedClientActionV2PtrOutput)
}

func (in *unauthenticatedClientActionV2Ptr) ToUnauthenticatedClientActionV2PtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionV2PtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UnauthenticatedClientActionV2PtrOutput)
}

func (in *unauthenticatedClientActionV2Ptr) ToOutput(ctx context.Context) pulumix.Output[*UnauthenticatedClientActionV2] {
	return pulumix.Output[*UnauthenticatedClientActionV2]{
		OutputState: in.ToUnauthenticatedClientActionV2PtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterOutputType(ClientCredentialMethodOutput{})
	pulumi.RegisterOutputType(ClientCredentialMethodPtrOutput{})
	pulumi.RegisterOutputType(CookieExpirationConventionOutput{})
	pulumi.RegisterOutputType(CookieExpirationConventionPtrOutput{})
	pulumi.RegisterOutputType(ForwardProxyConventionOutput{})
	pulumi.RegisterOutputType(ForwardProxyConventionPtrOutput{})
	pulumi.RegisterOutputType(UnauthenticatedClientActionV2Output{})
	pulumi.RegisterOutputType(UnauthenticatedClientActionV2PtrOutput{})
}
