// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The server key type like 'ServiceManaged', 'AzureKeyVault'.
type ServerKeyType string

const (
	ServerKeyTypeServiceManaged = ServerKeyType("ServiceManaged")
	ServerKeyTypeAzureKeyVault  = ServerKeyType("AzureKeyVault")
)

func (ServerKeyType) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerKeyType)(nil)).Elem()
}

func (e ServerKeyType) ToServerKeyTypeOutput() ServerKeyTypeOutput {
	return pulumi.ToOutput(e).(ServerKeyTypeOutput)
}

func (e ServerKeyType) ToServerKeyTypeOutputWithContext(ctx context.Context) ServerKeyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServerKeyTypeOutput)
}

func (e ServerKeyType) ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput {
	return e.ToServerKeyTypePtrOutputWithContext(context.Background())
}

func (e ServerKeyType) ToServerKeyTypePtrOutputWithContext(ctx context.Context) ServerKeyTypePtrOutput {
	return ServerKeyType(e).ToServerKeyTypeOutputWithContext(ctx).ToServerKeyTypePtrOutputWithContext(ctx)
}

func (e ServerKeyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerKeyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerKeyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServerKeyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServerKeyTypeOutput struct{ *pulumi.OutputState }

func (ServerKeyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerKeyType)(nil)).Elem()
}

func (o ServerKeyTypeOutput) ToServerKeyTypeOutput() ServerKeyTypeOutput {
	return o
}

func (o ServerKeyTypeOutput) ToServerKeyTypeOutputWithContext(ctx context.Context) ServerKeyTypeOutput {
	return o
}

func (o ServerKeyTypeOutput) ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput {
	return o.ToServerKeyTypePtrOutputWithContext(context.Background())
}

func (o ServerKeyTypeOutput) ToServerKeyTypePtrOutputWithContext(ctx context.Context) ServerKeyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerKeyType) *ServerKeyType {
		return &v
	}).(ServerKeyTypePtrOutput)
}

func (o ServerKeyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServerKeyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerKeyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServerKeyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerKeyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerKeyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServerKeyTypePtrOutput struct{ *pulumi.OutputState }

func (ServerKeyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerKeyType)(nil)).Elem()
}

func (o ServerKeyTypePtrOutput) ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput {
	return o
}

func (o ServerKeyTypePtrOutput) ToServerKeyTypePtrOutputWithContext(ctx context.Context) ServerKeyTypePtrOutput {
	return o
}

func (o ServerKeyTypePtrOutput) Elem() ServerKeyTypeOutput {
	return o.ApplyT(func(v *ServerKeyType) ServerKeyType {
		if v != nil {
			return *v
		}
		var ret ServerKeyType
		return ret
	}).(ServerKeyTypeOutput)
}

func (o ServerKeyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerKeyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServerKeyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ServerKeyTypeInput is an input type that accepts values of the ServerKeyType enum
// A concrete instance of `ServerKeyTypeInput` can be one of the following:
//
//	ServerKeyTypeServiceManaged
//	ServerKeyTypeAzureKeyVault
type ServerKeyTypeInput interface {
	pulumi.Input

	ToServerKeyTypeOutput() ServerKeyTypeOutput
	ToServerKeyTypeOutputWithContext(context.Context) ServerKeyTypeOutput
}

var serverKeyTypePtrType = reflect.TypeOf((**ServerKeyType)(nil)).Elem()

type ServerKeyTypePtrInput interface {
	pulumi.Input

	ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput
	ToServerKeyTypePtrOutputWithContext(context.Context) ServerKeyTypePtrOutput
}

type serverKeyTypePtr string

func ServerKeyTypePtr(v string) ServerKeyTypePtrInput {
	return (*serverKeyTypePtr)(&v)
}

func (*serverKeyTypePtr) ElementType() reflect.Type {
	return serverKeyTypePtrType
}

func (in *serverKeyTypePtr) ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput {
	return pulumi.ToOutput(in).(ServerKeyTypePtrOutput)
}

func (in *serverKeyTypePtr) ToServerKeyTypePtrOutputWithContext(ctx context.Context) ServerKeyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServerKeyTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerKeyTypeOutput{})
	pulumi.RegisterOutputType(ServerKeyTypePtrOutput{})
}
