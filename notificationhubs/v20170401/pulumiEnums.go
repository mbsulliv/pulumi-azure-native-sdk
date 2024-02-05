// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessRights string

const (
	AccessRightsManage = AccessRights("Manage")
	AccessRightsSend   = AccessRights("Send")
	AccessRightsListen = AccessRights("Listen")
)

func (AccessRights) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRights)(nil)).Elem()
}

func (e AccessRights) ToAccessRightsOutput() AccessRightsOutput {
	return pulumi.ToOutput(e).(AccessRightsOutput)
}

func (e AccessRights) ToAccessRightsOutputWithContext(ctx context.Context) AccessRightsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessRightsOutput)
}

func (e AccessRights) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return e.ToAccessRightsPtrOutputWithContext(context.Background())
}

func (e AccessRights) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return AccessRights(e).ToAccessRightsOutputWithContext(ctx).ToAccessRightsPtrOutputWithContext(ctx)
}

func (e AccessRights) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessRights) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessRights) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessRights) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessRightsOutput struct{ *pulumi.OutputState }

func (AccessRightsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRights)(nil)).Elem()
}

func (o AccessRightsOutput) ToAccessRightsOutput() AccessRightsOutput {
	return o
}

func (o AccessRightsOutput) ToAccessRightsOutputWithContext(ctx context.Context) AccessRightsOutput {
	return o
}

func (o AccessRightsOutput) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return o.ToAccessRightsPtrOutputWithContext(context.Background())
}

func (o AccessRightsOutput) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessRights) *AccessRights {
		return &v
	}).(AccessRightsPtrOutput)
}

func (o AccessRightsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessRightsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessRights) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessRightsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessRightsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessRights) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessRightsPtrOutput struct{ *pulumi.OutputState }

func (AccessRightsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessRights)(nil)).Elem()
}

func (o AccessRightsPtrOutput) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return o
}

func (o AccessRightsPtrOutput) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return o
}

func (o AccessRightsPtrOutput) Elem() AccessRightsOutput {
	return o.ApplyT(func(v *AccessRights) AccessRights {
		if v != nil {
			return *v
		}
		var ret AccessRights
		return ret
	}).(AccessRightsOutput)
}

func (o AccessRightsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessRightsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessRights) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AccessRightsInput is an input type that accepts values of the AccessRights enum
// A concrete instance of `AccessRightsInput` can be one of the following:
//
//	AccessRightsManage
//	AccessRightsSend
//	AccessRightsListen
type AccessRightsInput interface {
	pulumi.Input

	ToAccessRightsOutput() AccessRightsOutput
	ToAccessRightsOutputWithContext(context.Context) AccessRightsOutput
}

var accessRightsPtrType = reflect.TypeOf((**AccessRights)(nil)).Elem()

type AccessRightsPtrInput interface {
	pulumi.Input

	ToAccessRightsPtrOutput() AccessRightsPtrOutput
	ToAccessRightsPtrOutputWithContext(context.Context) AccessRightsPtrOutput
}

type accessRightsPtr string

func AccessRightsPtr(v string) AccessRightsPtrInput {
	return (*accessRightsPtr)(&v)
}

func (*accessRightsPtr) ElementType() reflect.Type {
	return accessRightsPtrType
}

func (in *accessRightsPtr) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return pulumi.ToOutput(in).(AccessRightsPtrOutput)
}

func (in *accessRightsPtr) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessRightsPtrOutput)
}

// AccessRightsArrayInput is an input type that accepts AccessRightsArray and AccessRightsArrayOutput values.
// You can construct a concrete instance of `AccessRightsArrayInput` via:
//
//	AccessRightsArray{ AccessRightsArgs{...} }
type AccessRightsArrayInput interface {
	pulumi.Input

	ToAccessRightsArrayOutput() AccessRightsArrayOutput
	ToAccessRightsArrayOutputWithContext(context.Context) AccessRightsArrayOutput
}

type AccessRightsArray []AccessRights

func (AccessRightsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessRights)(nil)).Elem()
}

func (i AccessRightsArray) ToAccessRightsArrayOutput() AccessRightsArrayOutput {
	return i.ToAccessRightsArrayOutputWithContext(context.Background())
}

func (i AccessRightsArray) ToAccessRightsArrayOutputWithContext(ctx context.Context) AccessRightsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessRightsArrayOutput)
}

type AccessRightsArrayOutput struct{ *pulumi.OutputState }

func (AccessRightsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessRights)(nil)).Elem()
}

func (o AccessRightsArrayOutput) ToAccessRightsArrayOutput() AccessRightsArrayOutput {
	return o
}

func (o AccessRightsArrayOutput) ToAccessRightsArrayOutputWithContext(ctx context.Context) AccessRightsArrayOutput {
	return o
}

func (o AccessRightsArrayOutput) Index(i pulumi.IntInput) AccessRightsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessRights {
		return vs[0].([]AccessRights)[vs[1].(int)]
	}).(AccessRightsOutput)
}

// The namespace type.
type NamespaceType string

const (
	NamespaceTypeMessaging       = NamespaceType("Messaging")
	NamespaceTypeNotificationHub = NamespaceType("NotificationHub")
)

func (NamespaceType) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceType)(nil)).Elem()
}

func (e NamespaceType) ToNamespaceTypeOutput() NamespaceTypeOutput {
	return pulumi.ToOutput(e).(NamespaceTypeOutput)
}

func (e NamespaceType) ToNamespaceTypeOutputWithContext(ctx context.Context) NamespaceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NamespaceTypeOutput)
}

func (e NamespaceType) ToNamespaceTypePtrOutput() NamespaceTypePtrOutput {
	return e.ToNamespaceTypePtrOutputWithContext(context.Background())
}

func (e NamespaceType) ToNamespaceTypePtrOutputWithContext(ctx context.Context) NamespaceTypePtrOutput {
	return NamespaceType(e).ToNamespaceTypeOutputWithContext(ctx).ToNamespaceTypePtrOutputWithContext(ctx)
}

func (e NamespaceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NamespaceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NamespaceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NamespaceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NamespaceTypeOutput struct{ *pulumi.OutputState }

func (NamespaceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceType)(nil)).Elem()
}

func (o NamespaceTypeOutput) ToNamespaceTypeOutput() NamespaceTypeOutput {
	return o
}

func (o NamespaceTypeOutput) ToNamespaceTypeOutputWithContext(ctx context.Context) NamespaceTypeOutput {
	return o
}

func (o NamespaceTypeOutput) ToNamespaceTypePtrOutput() NamespaceTypePtrOutput {
	return o.ToNamespaceTypePtrOutputWithContext(context.Background())
}

func (o NamespaceTypeOutput) ToNamespaceTypePtrOutputWithContext(ctx context.Context) NamespaceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NamespaceType) *NamespaceType {
		return &v
	}).(NamespaceTypePtrOutput)
}

func (o NamespaceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NamespaceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NamespaceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NamespaceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NamespaceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NamespaceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NamespaceTypePtrOutput struct{ *pulumi.OutputState }

func (NamespaceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceType)(nil)).Elem()
}

func (o NamespaceTypePtrOutput) ToNamespaceTypePtrOutput() NamespaceTypePtrOutput {
	return o
}

func (o NamespaceTypePtrOutput) ToNamespaceTypePtrOutputWithContext(ctx context.Context) NamespaceTypePtrOutput {
	return o
}

func (o NamespaceTypePtrOutput) Elem() NamespaceTypeOutput {
	return o.ApplyT(func(v *NamespaceType) NamespaceType {
		if v != nil {
			return *v
		}
		var ret NamespaceType
		return ret
	}).(NamespaceTypeOutput)
}

func (o NamespaceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NamespaceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NamespaceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// NamespaceTypeInput is an input type that accepts values of the NamespaceType enum
// A concrete instance of `NamespaceTypeInput` can be one of the following:
//
//	NamespaceTypeMessaging
//	NamespaceTypeNotificationHub
type NamespaceTypeInput interface {
	pulumi.Input

	ToNamespaceTypeOutput() NamespaceTypeOutput
	ToNamespaceTypeOutputWithContext(context.Context) NamespaceTypeOutput
}

var namespaceTypePtrType = reflect.TypeOf((**NamespaceType)(nil)).Elem()

type NamespaceTypePtrInput interface {
	pulumi.Input

	ToNamespaceTypePtrOutput() NamespaceTypePtrOutput
	ToNamespaceTypePtrOutputWithContext(context.Context) NamespaceTypePtrOutput
}

type namespaceTypePtr string

func NamespaceTypePtr(v string) NamespaceTypePtrInput {
	return (*namespaceTypePtr)(&v)
}

func (*namespaceTypePtr) ElementType() reflect.Type {
	return namespaceTypePtrType
}

func (in *namespaceTypePtr) ToNamespaceTypePtrOutput() NamespaceTypePtrOutput {
	return pulumi.ToOutput(in).(NamespaceTypePtrOutput)
}

func (in *namespaceTypePtr) ToNamespaceTypePtrOutputWithContext(ctx context.Context) NamespaceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NamespaceTypePtrOutput)
}

// Name of the notification hub sku
type SkuName string

const (
	SkuNameFree     = SkuName("Free")
	SkuNameBasic    = SkuName("Basic")
	SkuNameStandard = SkuName("Standard")
)

func (SkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (e SkuName) ToSkuNameOutput() SkuNameOutput {
	return pulumi.ToOutput(e).(SkuNameOutput)
}

func (e SkuName) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameOutput)
}

func (e SkuName) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return e.ToSkuNamePtrOutputWithContext(context.Background())
}

func (e SkuName) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return SkuName(e).ToSkuNameOutputWithContext(ctx).ToSkuNamePtrOutputWithContext(ctx)
}

func (e SkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameOutput struct{ *pulumi.OutputState }

func (SkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (o SkuNameOutput) ToSkuNameOutput() SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o.ToSkuNamePtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuName) *SkuName {
		return &v
	}).(SkuNamePtrOutput)
}

func (o SkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNamePtrOutput struct{ *pulumi.OutputState }

func (SkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuName)(nil)).Elem()
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) Elem() SkuNameOutput {
	return o.ApplyT(func(v *SkuName) SkuName {
		if v != nil {
			return *v
		}
		var ret SkuName
		return ret
	}).(SkuNameOutput)
}

func (o SkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SkuNameInput is an input type that accepts values of the SkuName enum
// A concrete instance of `SkuNameInput` can be one of the following:
//
//	SkuNameFree
//	SkuNameBasic
//	SkuNameStandard
type SkuNameInput interface {
	pulumi.Input

	ToSkuNameOutput() SkuNameOutput
	ToSkuNameOutputWithContext(context.Context) SkuNameOutput
}

var skuNamePtrType = reflect.TypeOf((**SkuName)(nil)).Elem()

type SkuNamePtrInput interface {
	pulumi.Input

	ToSkuNamePtrOutput() SkuNamePtrOutput
	ToSkuNamePtrOutputWithContext(context.Context) SkuNamePtrOutput
}

type skuNamePtr string

func SkuNamePtr(v string) SkuNamePtrInput {
	return (*skuNamePtr)(&v)
}

func (*skuNamePtr) ElementType() reflect.Type {
	return skuNamePtrType
}

func (in *skuNamePtr) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return pulumi.ToOutput(in).(SkuNamePtrOutput)
}

func (in *skuNamePtr) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNamePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessRightsOutput{})
	pulumi.RegisterOutputType(AccessRightsPtrOutput{})
	pulumi.RegisterOutputType(AccessRightsArrayOutput{})
	pulumi.RegisterOutputType(NamespaceTypeOutput{})
	pulumi.RegisterOutputType(NamespaceTypePtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
}
