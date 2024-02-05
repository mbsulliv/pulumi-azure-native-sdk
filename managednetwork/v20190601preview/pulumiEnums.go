// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Responsibility role under which this Managed Network Group will be created
type Kind string

const (
	KindConnectivity = Kind("Connectivity")
)

func (Kind) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (e Kind) ToKindOutput() KindOutput {
	return pulumi.ToOutput(e).(KindOutput)
}

func (e Kind) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KindOutput)
}

func (e Kind) ToKindPtrOutput() KindPtrOutput {
	return e.ToKindPtrOutputWithContext(context.Background())
}

func (e Kind) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return Kind(e).ToKindOutputWithContext(ctx).ToKindPtrOutputWithContext(ctx)
}

func (e Kind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Kind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KindOutput struct{ *pulumi.OutputState }

func (KindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (o KindOutput) ToKindOutput() KindOutput {
	return o
}

func (o KindOutput) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return o
}

func (o KindOutput) ToKindPtrOutput() KindPtrOutput {
	return o.ToKindPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Kind) *Kind {
		return &v
	}).(KindPtrOutput)
}

func (o KindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KindPtrOutput struct{ *pulumi.OutputState }

func (KindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kind)(nil)).Elem()
}

func (o KindPtrOutput) ToKindPtrOutput() KindPtrOutput {
	return o
}

func (o KindPtrOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o
}

func (o KindPtrOutput) Elem() KindOutput {
	return o.ApplyT(func(v *Kind) Kind {
		if v != nil {
			return *v
		}
		var ret Kind
		return ret
	}).(KindOutput)
}

func (o KindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Kind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// KindInput is an input type that accepts values of the Kind enum
// A concrete instance of `KindInput` can be one of the following:
//
//	KindConnectivity
type KindInput interface {
	pulumi.Input

	ToKindOutput() KindOutput
	ToKindOutputWithContext(context.Context) KindOutput
}

var kindPtrType = reflect.TypeOf((**Kind)(nil)).Elem()

type KindPtrInput interface {
	pulumi.Input

	ToKindPtrOutput() KindPtrOutput
	ToKindPtrOutputWithContext(context.Context) KindPtrOutput
}

type kindPtr string

func KindPtr(v string) KindPtrInput {
	return (*kindPtr)(&v)
}

func (*kindPtr) ElementType() reflect.Type {
	return kindPtrType
}

func (in *kindPtr) ToKindPtrOutput() KindPtrOutput {
	return pulumi.ToOutput(in).(KindPtrOutput)
}

func (in *kindPtr) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KindPtrOutput)
}

// Gets or sets the connectivity type of a network structure policy
type Type string

const (
	TypeHubAndSpokeTopology = Type("HubAndSpokeTopology")
	TypeMeshTopology        = Type("MeshTopology")
)

func (Type) ElementType() reflect.Type {
	return reflect.TypeOf((*Type)(nil)).Elem()
}

func (e Type) ToTypeOutput() TypeOutput {
	return pulumi.ToOutput(e).(TypeOutput)
}

func (e Type) ToTypeOutputWithContext(ctx context.Context) TypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TypeOutput)
}

func (e Type) ToTypePtrOutput() TypePtrOutput {
	return e.ToTypePtrOutputWithContext(context.Background())
}

func (e Type) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return Type(e).ToTypeOutputWithContext(ctx).ToTypePtrOutputWithContext(ctx)
}

func (e Type) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Type) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Type) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Type) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TypeOutput struct{ *pulumi.OutputState }

func (TypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Type)(nil)).Elem()
}

func (o TypeOutput) ToTypeOutput() TypeOutput {
	return o
}

func (o TypeOutput) ToTypeOutputWithContext(ctx context.Context) TypeOutput {
	return o
}

func (o TypeOutput) ToTypePtrOutput() TypePtrOutput {
	return o.ToTypePtrOutputWithContext(context.Background())
}

func (o TypeOutput) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Type) *Type {
		return &v
	}).(TypePtrOutput)
}

func (o TypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Type) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Type) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TypePtrOutput struct{ *pulumi.OutputState }

func (TypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Type)(nil)).Elem()
}

func (o TypePtrOutput) ToTypePtrOutput() TypePtrOutput {
	return o
}

func (o TypePtrOutput) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return o
}

func (o TypePtrOutput) Elem() TypeOutput {
	return o.ApplyT(func(v *Type) Type {
		if v != nil {
			return *v
		}
		var ret Type
		return ret
	}).(TypeOutput)
}

func (o TypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Type) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TypeInput is an input type that accepts values of the Type enum
// A concrete instance of `TypeInput` can be one of the following:
//
//	TypeHubAndSpokeTopology
//	TypeMeshTopology
type TypeInput interface {
	pulumi.Input

	ToTypeOutput() TypeOutput
	ToTypeOutputWithContext(context.Context) TypeOutput
}

var typePtrType = reflect.TypeOf((**Type)(nil)).Elem()

type TypePtrInput interface {
	pulumi.Input

	ToTypePtrOutput() TypePtrOutput
	ToTypePtrOutputWithContext(context.Context) TypePtrOutput
}

type typePtr string

func TypePtr(v string) TypePtrInput {
	return (*typePtr)(&v)
}

func (*typePtr) ElementType() reflect.Type {
	return typePtrType
}

func (in *typePtr) ToTypePtrOutput() TypePtrOutput {
	return pulumi.ToOutput(in).(TypePtrOutput)
}

func (in *typePtr) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(KindOutput{})
	pulumi.RegisterOutputType(KindPtrOutput{})
	pulumi.RegisterOutputType(TypeOutput{})
	pulumi.RegisterOutputType(TypePtrOutput{})
}
