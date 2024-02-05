// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
type CertificateFormat string

const (
	// The certificate is a PFX (PKCS#12) formatted certificate or certificate chain.
	CertificateFormatPfx = CertificateFormat("Pfx")
	// The certificate is a base64-encoded X.509 certificate.
	CertificateFormatCer = CertificateFormat("Cer")
)

func (CertificateFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateFormat)(nil)).Elem()
}

func (e CertificateFormat) ToCertificateFormatOutput() CertificateFormatOutput {
	return pulumi.ToOutput(e).(CertificateFormatOutput)
}

func (e CertificateFormat) ToCertificateFormatOutputWithContext(ctx context.Context) CertificateFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CertificateFormatOutput)
}

func (e CertificateFormat) ToCertificateFormatPtrOutput() CertificateFormatPtrOutput {
	return e.ToCertificateFormatPtrOutputWithContext(context.Background())
}

func (e CertificateFormat) ToCertificateFormatPtrOutputWithContext(ctx context.Context) CertificateFormatPtrOutput {
	return CertificateFormat(e).ToCertificateFormatOutputWithContext(ctx).ToCertificateFormatPtrOutputWithContext(ctx)
}

func (e CertificateFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CertificateFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CertificateFormatOutput struct{ *pulumi.OutputState }

func (CertificateFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateFormat)(nil)).Elem()
}

func (o CertificateFormatOutput) ToCertificateFormatOutput() CertificateFormatOutput {
	return o
}

func (o CertificateFormatOutput) ToCertificateFormatOutputWithContext(ctx context.Context) CertificateFormatOutput {
	return o
}

func (o CertificateFormatOutput) ToCertificateFormatPtrOutput() CertificateFormatPtrOutput {
	return o.ToCertificateFormatPtrOutputWithContext(context.Background())
}

func (o CertificateFormatOutput) ToCertificateFormatPtrOutputWithContext(ctx context.Context) CertificateFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateFormat) *CertificateFormat {
		return &v
	}).(CertificateFormatPtrOutput)
}

func (o CertificateFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CertificateFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CertificateFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CertificateFormatPtrOutput struct{ *pulumi.OutputState }

func (CertificateFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateFormat)(nil)).Elem()
}

func (o CertificateFormatPtrOutput) ToCertificateFormatPtrOutput() CertificateFormatPtrOutput {
	return o
}

func (o CertificateFormatPtrOutput) ToCertificateFormatPtrOutputWithContext(ctx context.Context) CertificateFormatPtrOutput {
	return o
}

func (o CertificateFormatPtrOutput) Elem() CertificateFormatOutput {
	return o.ApplyT(func(v *CertificateFormat) CertificateFormat {
		if v != nil {
			return *v
		}
		var ret CertificateFormat
		return ret
	}).(CertificateFormatOutput)
}

func (o CertificateFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CertificateFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CertificateFormatInput is an input type that accepts values of the CertificateFormat enum
// A concrete instance of `CertificateFormatInput` can be one of the following:
//
//	CertificateFormatPfx
//	CertificateFormatCer
type CertificateFormatInput interface {
	pulumi.Input

	ToCertificateFormatOutput() CertificateFormatOutput
	ToCertificateFormatOutputWithContext(context.Context) CertificateFormatOutput
}

var certificateFormatPtrType = reflect.TypeOf((**CertificateFormat)(nil)).Elem()

type CertificateFormatPtrInput interface {
	pulumi.Input

	ToCertificateFormatPtrOutput() CertificateFormatPtrOutput
	ToCertificateFormatPtrOutputWithContext(context.Context) CertificateFormatPtrOutput
}

type certificateFormatPtr string

func CertificateFormatPtr(v string) CertificateFormatPtrInput {
	return (*certificateFormatPtr)(&v)
}

func (*certificateFormatPtr) ElementType() reflect.Type {
	return certificateFormatPtrType
}

func (in *certificateFormatPtr) ToCertificateFormatPtrOutput() CertificateFormatPtrOutput {
	return pulumi.ToOutput(in).(CertificateFormatPtrOutput)
}

func (in *certificateFormatPtr) ToCertificateFormatPtrOutputWithContext(ctx context.Context) CertificateFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CertificateFormatPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateFormatOutput{})
	pulumi.RegisterOutputType(CertificateFormatPtrOutput{})
}
