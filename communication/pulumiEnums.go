// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package communication

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes how a Domains resource is being managed.
type DomainManagement string

const (
	DomainManagementAzureManaged                    = DomainManagement("AzureManaged")
	DomainManagementCustomerManaged                 = DomainManagement("CustomerManaged")
	DomainManagementCustomerManagedInExchangeOnline = DomainManagement("CustomerManagedInExchangeOnline")
)

func (DomainManagement) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainManagement)(nil)).Elem()
}

func (e DomainManagement) ToDomainManagementOutput() DomainManagementOutput {
	return pulumi.ToOutput(e).(DomainManagementOutput)
}

func (e DomainManagement) ToDomainManagementOutputWithContext(ctx context.Context) DomainManagementOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DomainManagementOutput)
}

func (e DomainManagement) ToDomainManagementPtrOutput() DomainManagementPtrOutput {
	return e.ToDomainManagementPtrOutputWithContext(context.Background())
}

func (e DomainManagement) ToDomainManagementPtrOutputWithContext(ctx context.Context) DomainManagementPtrOutput {
	return DomainManagement(e).ToDomainManagementOutputWithContext(ctx).ToDomainManagementPtrOutputWithContext(ctx)
}

func (e DomainManagement) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DomainManagement) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DomainManagement) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DomainManagement) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DomainManagementOutput struct{ *pulumi.OutputState }

func (DomainManagementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainManagement)(nil)).Elem()
}

func (o DomainManagementOutput) ToDomainManagementOutput() DomainManagementOutput {
	return o
}

func (o DomainManagementOutput) ToDomainManagementOutputWithContext(ctx context.Context) DomainManagementOutput {
	return o
}

func (o DomainManagementOutput) ToDomainManagementPtrOutput() DomainManagementPtrOutput {
	return o.ToDomainManagementPtrOutputWithContext(context.Background())
}

func (o DomainManagementOutput) ToDomainManagementPtrOutputWithContext(ctx context.Context) DomainManagementPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainManagement) *DomainManagement {
		return &v
	}).(DomainManagementPtrOutput)
}

func (o DomainManagementOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DomainManagementOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DomainManagement) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DomainManagementOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DomainManagementOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DomainManagement) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DomainManagementPtrOutput struct{ *pulumi.OutputState }

func (DomainManagementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainManagement)(nil)).Elem()
}

func (o DomainManagementPtrOutput) ToDomainManagementPtrOutput() DomainManagementPtrOutput {
	return o
}

func (o DomainManagementPtrOutput) ToDomainManagementPtrOutputWithContext(ctx context.Context) DomainManagementPtrOutput {
	return o
}

func (o DomainManagementPtrOutput) Elem() DomainManagementOutput {
	return o.ApplyT(func(v *DomainManagement) DomainManagement {
		if v != nil {
			return *v
		}
		var ret DomainManagement
		return ret
	}).(DomainManagementOutput)
}

func (o DomainManagementPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DomainManagementPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DomainManagement) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DomainManagementInput is an input type that accepts values of the DomainManagement enum
// A concrete instance of `DomainManagementInput` can be one of the following:
//
//	DomainManagementAzureManaged
//	DomainManagementCustomerManaged
//	DomainManagementCustomerManagedInExchangeOnline
type DomainManagementInput interface {
	pulumi.Input

	ToDomainManagementOutput() DomainManagementOutput
	ToDomainManagementOutputWithContext(context.Context) DomainManagementOutput
}

var domainManagementPtrType = reflect.TypeOf((**DomainManagement)(nil)).Elem()

type DomainManagementPtrInput interface {
	pulumi.Input

	ToDomainManagementPtrOutput() DomainManagementPtrOutput
	ToDomainManagementPtrOutputWithContext(context.Context) DomainManagementPtrOutput
}

type domainManagementPtr string

func DomainManagementPtr(v string) DomainManagementPtrInput {
	return (*domainManagementPtr)(&v)
}

func (*domainManagementPtr) ElementType() reflect.Type {
	return domainManagementPtrType
}

func (in *domainManagementPtr) ToDomainManagementPtrOutput() DomainManagementPtrOutput {
	return pulumi.ToOutput(in).(DomainManagementPtrOutput)
}

func (in *domainManagementPtr) ToDomainManagementPtrOutputWithContext(ctx context.Context) DomainManagementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DomainManagementPtrOutput)
}

// Describes whether user engagement tracking is enabled or disabled.
type UserEngagementTracking string

const (
	UserEngagementTrackingDisabled = UserEngagementTracking("Disabled")
	UserEngagementTrackingEnabled  = UserEngagementTracking("Enabled")
)

func (UserEngagementTracking) ElementType() reflect.Type {
	return reflect.TypeOf((*UserEngagementTracking)(nil)).Elem()
}

func (e UserEngagementTracking) ToUserEngagementTrackingOutput() UserEngagementTrackingOutput {
	return pulumi.ToOutput(e).(UserEngagementTrackingOutput)
}

func (e UserEngagementTracking) ToUserEngagementTrackingOutputWithContext(ctx context.Context) UserEngagementTrackingOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UserEngagementTrackingOutput)
}

func (e UserEngagementTracking) ToUserEngagementTrackingPtrOutput() UserEngagementTrackingPtrOutput {
	return e.ToUserEngagementTrackingPtrOutputWithContext(context.Background())
}

func (e UserEngagementTracking) ToUserEngagementTrackingPtrOutputWithContext(ctx context.Context) UserEngagementTrackingPtrOutput {
	return UserEngagementTracking(e).ToUserEngagementTrackingOutputWithContext(ctx).ToUserEngagementTrackingPtrOutputWithContext(ctx)
}

func (e UserEngagementTracking) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserEngagementTracking) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserEngagementTracking) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UserEngagementTracking) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UserEngagementTrackingOutput struct{ *pulumi.OutputState }

func (UserEngagementTrackingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserEngagementTracking)(nil)).Elem()
}

func (o UserEngagementTrackingOutput) ToUserEngagementTrackingOutput() UserEngagementTrackingOutput {
	return o
}

func (o UserEngagementTrackingOutput) ToUserEngagementTrackingOutputWithContext(ctx context.Context) UserEngagementTrackingOutput {
	return o
}

func (o UserEngagementTrackingOutput) ToUserEngagementTrackingPtrOutput() UserEngagementTrackingPtrOutput {
	return o.ToUserEngagementTrackingPtrOutputWithContext(context.Background())
}

func (o UserEngagementTrackingOutput) ToUserEngagementTrackingPtrOutputWithContext(ctx context.Context) UserEngagementTrackingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserEngagementTracking) *UserEngagementTracking {
		return &v
	}).(UserEngagementTrackingPtrOutput)
}

func (o UserEngagementTrackingOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UserEngagementTrackingOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserEngagementTracking) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UserEngagementTrackingOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserEngagementTrackingOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserEngagementTracking) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UserEngagementTrackingPtrOutput struct{ *pulumi.OutputState }

func (UserEngagementTrackingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserEngagementTracking)(nil)).Elem()
}

func (o UserEngagementTrackingPtrOutput) ToUserEngagementTrackingPtrOutput() UserEngagementTrackingPtrOutput {
	return o
}

func (o UserEngagementTrackingPtrOutput) ToUserEngagementTrackingPtrOutputWithContext(ctx context.Context) UserEngagementTrackingPtrOutput {
	return o
}

func (o UserEngagementTrackingPtrOutput) Elem() UserEngagementTrackingOutput {
	return o.ApplyT(func(v *UserEngagementTracking) UserEngagementTracking {
		if v != nil {
			return *v
		}
		var ret UserEngagementTracking
		return ret
	}).(UserEngagementTrackingOutput)
}

func (o UserEngagementTrackingPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserEngagementTrackingPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UserEngagementTracking) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// UserEngagementTrackingInput is an input type that accepts values of the UserEngagementTracking enum
// A concrete instance of `UserEngagementTrackingInput` can be one of the following:
//
//	UserEngagementTrackingDisabled
//	UserEngagementTrackingEnabled
type UserEngagementTrackingInput interface {
	pulumi.Input

	ToUserEngagementTrackingOutput() UserEngagementTrackingOutput
	ToUserEngagementTrackingOutputWithContext(context.Context) UserEngagementTrackingOutput
}

var userEngagementTrackingPtrType = reflect.TypeOf((**UserEngagementTracking)(nil)).Elem()

type UserEngagementTrackingPtrInput interface {
	pulumi.Input

	ToUserEngagementTrackingPtrOutput() UserEngagementTrackingPtrOutput
	ToUserEngagementTrackingPtrOutputWithContext(context.Context) UserEngagementTrackingPtrOutput
}

type userEngagementTrackingPtr string

func UserEngagementTrackingPtr(v string) UserEngagementTrackingPtrInput {
	return (*userEngagementTrackingPtr)(&v)
}

func (*userEngagementTrackingPtr) ElementType() reflect.Type {
	return userEngagementTrackingPtrType
}

func (in *userEngagementTrackingPtr) ToUserEngagementTrackingPtrOutput() UserEngagementTrackingPtrOutput {
	return pulumi.ToOutput(in).(UserEngagementTrackingPtrOutput)
}

func (in *userEngagementTrackingPtr) ToUserEngagementTrackingPtrOutputWithContext(ctx context.Context) UserEngagementTrackingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UserEngagementTrackingPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainManagementOutput{})
	pulumi.RegisterOutputType(DomainManagementPtrOutput{})
	pulumi.RegisterOutputType(UserEngagementTrackingOutput{})
	pulumi.RegisterOutputType(UserEngagementTrackingPtrOutput{})
}
