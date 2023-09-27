// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Certificate product type.
type CertificateProductType string

const (
	CertificateProductTypeStandardDomainValidatedSsl         = CertificateProductType("StandardDomainValidatedSsl")
	CertificateProductTypeStandardDomainValidatedWildCardSsl = CertificateProductType("StandardDomainValidatedWildCardSsl")
)

func (CertificateProductType) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateProductType)(nil)).Elem()
}

func (e CertificateProductType) ToCertificateProductTypeOutput() CertificateProductTypeOutput {
	return pulumi.ToOutput(e).(CertificateProductTypeOutput)
}

func (e CertificateProductType) ToCertificateProductTypeOutputWithContext(ctx context.Context) CertificateProductTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CertificateProductTypeOutput)
}

func (e CertificateProductType) ToCertificateProductTypePtrOutput() CertificateProductTypePtrOutput {
	return e.ToCertificateProductTypePtrOutputWithContext(context.Background())
}

func (e CertificateProductType) ToCertificateProductTypePtrOutputWithContext(ctx context.Context) CertificateProductTypePtrOutput {
	return CertificateProductType(e).ToCertificateProductTypeOutputWithContext(ctx).ToCertificateProductTypePtrOutputWithContext(ctx)
}

func (e CertificateProductType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateProductType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateProductType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CertificateProductType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CertificateProductTypeOutput struct{ *pulumi.OutputState }

func (CertificateProductTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateProductType)(nil)).Elem()
}

func (o CertificateProductTypeOutput) ToCertificateProductTypeOutput() CertificateProductTypeOutput {
	return o
}

func (o CertificateProductTypeOutput) ToCertificateProductTypeOutputWithContext(ctx context.Context) CertificateProductTypeOutput {
	return o
}

func (o CertificateProductTypeOutput) ToCertificateProductTypePtrOutput() CertificateProductTypePtrOutput {
	return o.ToCertificateProductTypePtrOutputWithContext(context.Background())
}

func (o CertificateProductTypeOutput) ToCertificateProductTypePtrOutputWithContext(ctx context.Context) CertificateProductTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateProductType) *CertificateProductType {
		return &v
	}).(CertificateProductTypePtrOutput)
}

func (o CertificateProductTypeOutput) ToOutput(ctx context.Context) pulumix.Output[CertificateProductType] {
	return pulumix.Output[CertificateProductType]{
		OutputState: o.OutputState,
	}
}

func (o CertificateProductTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CertificateProductTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateProductType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CertificateProductTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateProductTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateProductType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CertificateProductTypePtrOutput struct{ *pulumi.OutputState }

func (CertificateProductTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateProductType)(nil)).Elem()
}

func (o CertificateProductTypePtrOutput) ToCertificateProductTypePtrOutput() CertificateProductTypePtrOutput {
	return o
}

func (o CertificateProductTypePtrOutput) ToCertificateProductTypePtrOutputWithContext(ctx context.Context) CertificateProductTypePtrOutput {
	return o
}

func (o CertificateProductTypePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*CertificateProductType] {
	return pulumix.Output[*CertificateProductType]{
		OutputState: o.OutputState,
	}
}

func (o CertificateProductTypePtrOutput) Elem() CertificateProductTypeOutput {
	return o.ApplyT(func(v *CertificateProductType) CertificateProductType {
		if v != nil {
			return *v
		}
		var ret CertificateProductType
		return ret
	}).(CertificateProductTypeOutput)
}

func (o CertificateProductTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateProductTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CertificateProductType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CertificateProductTypeInput is an input type that accepts CertificateProductTypeArgs and CertificateProductTypeOutput values.
// You can construct a concrete instance of `CertificateProductTypeInput` via:
//
//	CertificateProductTypeArgs{...}
type CertificateProductTypeInput interface {
	pulumi.Input

	ToCertificateProductTypeOutput() CertificateProductTypeOutput
	ToCertificateProductTypeOutputWithContext(context.Context) CertificateProductTypeOutput
}

var certificateProductTypePtrType = reflect.TypeOf((**CertificateProductType)(nil)).Elem()

type CertificateProductTypePtrInput interface {
	pulumi.Input

	ToCertificateProductTypePtrOutput() CertificateProductTypePtrOutput
	ToCertificateProductTypePtrOutputWithContext(context.Context) CertificateProductTypePtrOutput
}

type certificateProductTypePtr string

func CertificateProductTypePtr(v string) CertificateProductTypePtrInput {
	return (*certificateProductTypePtr)(&v)
}

func (*certificateProductTypePtr) ElementType() reflect.Type {
	return certificateProductTypePtrType
}

func (in *certificateProductTypePtr) ToCertificateProductTypePtrOutput() CertificateProductTypePtrOutput {
	return pulumi.ToOutput(in).(CertificateProductTypePtrOutput)
}

func (in *certificateProductTypePtr) ToCertificateProductTypePtrOutputWithContext(ctx context.Context) CertificateProductTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CertificateProductTypePtrOutput)
}

func (in *certificateProductTypePtr) ToOutput(ctx context.Context) pulumix.Output[*CertificateProductType] {
	return pulumix.Output[*CertificateProductType]{
		OutputState: in.ToCertificateProductTypePtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterOutputType(CertificateProductTypeOutput{})
	pulumi.RegisterOutputType(CertificateProductTypePtrOutput{})
}
