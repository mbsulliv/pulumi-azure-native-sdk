// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Which authentication method Cassandra should use to authenticate clients. 'None' turns off authentication, so should not be used except in emergencies. 'Cassandra' is the default password based authentication. The default is 'Cassandra'.
type AuthenticationMethod string

const (
	AuthenticationMethodNone      = AuthenticationMethod("None")
	AuthenticationMethodCassandra = AuthenticationMethod("Cassandra")
)

func (AuthenticationMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationMethod)(nil)).Elem()
}

func (e AuthenticationMethod) ToAuthenticationMethodOutput() AuthenticationMethodOutput {
	return pulumi.ToOutput(e).(AuthenticationMethodOutput)
}

func (e AuthenticationMethod) ToAuthenticationMethodOutputWithContext(ctx context.Context) AuthenticationMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuthenticationMethodOutput)
}

func (e AuthenticationMethod) ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput {
	return e.ToAuthenticationMethodPtrOutputWithContext(context.Background())
}

func (e AuthenticationMethod) ToAuthenticationMethodPtrOutputWithContext(ctx context.Context) AuthenticationMethodPtrOutput {
	return AuthenticationMethod(e).ToAuthenticationMethodOutputWithContext(ctx).ToAuthenticationMethodPtrOutputWithContext(ctx)
}

func (e AuthenticationMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuthenticationMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuthenticationMethodOutput struct{ *pulumi.OutputState }

func (AuthenticationMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationMethod)(nil)).Elem()
}

func (o AuthenticationMethodOutput) ToAuthenticationMethodOutput() AuthenticationMethodOutput {
	return o
}

func (o AuthenticationMethodOutput) ToAuthenticationMethodOutputWithContext(ctx context.Context) AuthenticationMethodOutput {
	return o
}

func (o AuthenticationMethodOutput) ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput {
	return o.ToAuthenticationMethodPtrOutputWithContext(context.Background())
}

func (o AuthenticationMethodOutput) ToAuthenticationMethodPtrOutputWithContext(ctx context.Context) AuthenticationMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthenticationMethod) *AuthenticationMethod {
		return &v
	}).(AuthenticationMethodPtrOutput)
}

func (o AuthenticationMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuthenticationMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuthenticationMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuthenticationMethodPtrOutput struct{ *pulumi.OutputState }

func (AuthenticationMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationMethod)(nil)).Elem()
}

func (o AuthenticationMethodPtrOutput) ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput {
	return o
}

func (o AuthenticationMethodPtrOutput) ToAuthenticationMethodPtrOutputWithContext(ctx context.Context) AuthenticationMethodPtrOutput {
	return o
}

func (o AuthenticationMethodPtrOutput) Elem() AuthenticationMethodOutput {
	return o.ApplyT(func(v *AuthenticationMethod) AuthenticationMethod {
		if v != nil {
			return *v
		}
		var ret AuthenticationMethod
		return ret
	}).(AuthenticationMethodOutput)
}

func (o AuthenticationMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuthenticationMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AuthenticationMethodInput is an input type that accepts values of the AuthenticationMethod enum
// A concrete instance of `AuthenticationMethodInput` can be one of the following:
//
//	AuthenticationMethodNone
//	AuthenticationMethodCassandra
type AuthenticationMethodInput interface {
	pulumi.Input

	ToAuthenticationMethodOutput() AuthenticationMethodOutput
	ToAuthenticationMethodOutputWithContext(context.Context) AuthenticationMethodOutput
}

var authenticationMethodPtrType = reflect.TypeOf((**AuthenticationMethod)(nil)).Elem()

type AuthenticationMethodPtrInput interface {
	pulumi.Input

	ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput
	ToAuthenticationMethodPtrOutputWithContext(context.Context) AuthenticationMethodPtrOutput
}

type authenticationMethodPtr string

func AuthenticationMethodPtr(v string) AuthenticationMethodPtrInput {
	return (*authenticationMethodPtr)(&v)
}

func (*authenticationMethodPtr) ElementType() reflect.Type {
	return authenticationMethodPtrType
}

func (in *authenticationMethodPtr) ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput {
	return pulumi.ToOutput(in).(AuthenticationMethodPtrOutput)
}

func (in *authenticationMethodPtr) ToAuthenticationMethodPtrOutputWithContext(ctx context.Context) AuthenticationMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuthenticationMethodPtrOutput)
}

// The status of the resource at the time the operation was called.
type ManagedCassandraProvisioningState string

const (
	ManagedCassandraProvisioningStateCreating  = ManagedCassandraProvisioningState("Creating")
	ManagedCassandraProvisioningStateUpdating  = ManagedCassandraProvisioningState("Updating")
	ManagedCassandraProvisioningStateDeleting  = ManagedCassandraProvisioningState("Deleting")
	ManagedCassandraProvisioningStateSucceeded = ManagedCassandraProvisioningState("Succeeded")
	ManagedCassandraProvisioningStateFailed    = ManagedCassandraProvisioningState("Failed")
	ManagedCassandraProvisioningStateCanceled  = ManagedCassandraProvisioningState("Canceled")
)

func (ManagedCassandraProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedCassandraProvisioningState)(nil)).Elem()
}

func (e ManagedCassandraProvisioningState) ToManagedCassandraProvisioningStateOutput() ManagedCassandraProvisioningStateOutput {
	return pulumi.ToOutput(e).(ManagedCassandraProvisioningStateOutput)
}

func (e ManagedCassandraProvisioningState) ToManagedCassandraProvisioningStateOutputWithContext(ctx context.Context) ManagedCassandraProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedCassandraProvisioningStateOutput)
}

func (e ManagedCassandraProvisioningState) ToManagedCassandraProvisioningStatePtrOutput() ManagedCassandraProvisioningStatePtrOutput {
	return e.ToManagedCassandraProvisioningStatePtrOutputWithContext(context.Background())
}

func (e ManagedCassandraProvisioningState) ToManagedCassandraProvisioningStatePtrOutputWithContext(ctx context.Context) ManagedCassandraProvisioningStatePtrOutput {
	return ManagedCassandraProvisioningState(e).ToManagedCassandraProvisioningStateOutputWithContext(ctx).ToManagedCassandraProvisioningStatePtrOutputWithContext(ctx)
}

func (e ManagedCassandraProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedCassandraProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedCassandraProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedCassandraProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedCassandraProvisioningStateOutput struct{ *pulumi.OutputState }

func (ManagedCassandraProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedCassandraProvisioningState)(nil)).Elem()
}

func (o ManagedCassandraProvisioningStateOutput) ToManagedCassandraProvisioningStateOutput() ManagedCassandraProvisioningStateOutput {
	return o
}

func (o ManagedCassandraProvisioningStateOutput) ToManagedCassandraProvisioningStateOutputWithContext(ctx context.Context) ManagedCassandraProvisioningStateOutput {
	return o
}

func (o ManagedCassandraProvisioningStateOutput) ToManagedCassandraProvisioningStatePtrOutput() ManagedCassandraProvisioningStatePtrOutput {
	return o.ToManagedCassandraProvisioningStatePtrOutputWithContext(context.Background())
}

func (o ManagedCassandraProvisioningStateOutput) ToManagedCassandraProvisioningStatePtrOutputWithContext(ctx context.Context) ManagedCassandraProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedCassandraProvisioningState) *ManagedCassandraProvisioningState {
		return &v
	}).(ManagedCassandraProvisioningStatePtrOutput)
}

func (o ManagedCassandraProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedCassandraProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedCassandraProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedCassandraProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedCassandraProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedCassandraProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedCassandraProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (ManagedCassandraProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCassandraProvisioningState)(nil)).Elem()
}

func (o ManagedCassandraProvisioningStatePtrOutput) ToManagedCassandraProvisioningStatePtrOutput() ManagedCassandraProvisioningStatePtrOutput {
	return o
}

func (o ManagedCassandraProvisioningStatePtrOutput) ToManagedCassandraProvisioningStatePtrOutputWithContext(ctx context.Context) ManagedCassandraProvisioningStatePtrOutput {
	return o
}

func (o ManagedCassandraProvisioningStatePtrOutput) Elem() ManagedCassandraProvisioningStateOutput {
	return o.ApplyT(func(v *ManagedCassandraProvisioningState) ManagedCassandraProvisioningState {
		if v != nil {
			return *v
		}
		var ret ManagedCassandraProvisioningState
		return ret
	}).(ManagedCassandraProvisioningStateOutput)
}

func (o ManagedCassandraProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedCassandraProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedCassandraProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ManagedCassandraProvisioningStateInput is an input type that accepts values of the ManagedCassandraProvisioningState enum
// A concrete instance of `ManagedCassandraProvisioningStateInput` can be one of the following:
//
//	ManagedCassandraProvisioningStateCreating
//	ManagedCassandraProvisioningStateUpdating
//	ManagedCassandraProvisioningStateDeleting
//	ManagedCassandraProvisioningStateSucceeded
//	ManagedCassandraProvisioningStateFailed
//	ManagedCassandraProvisioningStateCanceled
type ManagedCassandraProvisioningStateInput interface {
	pulumi.Input

	ToManagedCassandraProvisioningStateOutput() ManagedCassandraProvisioningStateOutput
	ToManagedCassandraProvisioningStateOutputWithContext(context.Context) ManagedCassandraProvisioningStateOutput
}

var managedCassandraProvisioningStatePtrType = reflect.TypeOf((**ManagedCassandraProvisioningState)(nil)).Elem()

type ManagedCassandraProvisioningStatePtrInput interface {
	pulumi.Input

	ToManagedCassandraProvisioningStatePtrOutput() ManagedCassandraProvisioningStatePtrOutput
	ToManagedCassandraProvisioningStatePtrOutputWithContext(context.Context) ManagedCassandraProvisioningStatePtrOutput
}

type managedCassandraProvisioningStatePtr string

func ManagedCassandraProvisioningStatePtr(v string) ManagedCassandraProvisioningStatePtrInput {
	return (*managedCassandraProvisioningStatePtr)(&v)
}

func (*managedCassandraProvisioningStatePtr) ElementType() reflect.Type {
	return managedCassandraProvisioningStatePtrType
}

func (in *managedCassandraProvisioningStatePtr) ToManagedCassandraProvisioningStatePtrOutput() ManagedCassandraProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(ManagedCassandraProvisioningStatePtrOutput)
}

func (in *managedCassandraProvisioningStatePtr) ToManagedCassandraProvisioningStatePtrOutputWithContext(ctx context.Context) ManagedCassandraProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedCassandraProvisioningStatePtrOutput)
}

// The type of identity used for the resource. The type 'SystemAssigned,UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the service.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned,UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ResourceIdentityTypeInput is an input type that accepts values of the ResourceIdentityType enum
// A concrete instance of `ResourceIdentityTypeInput` can be one of the following:
//
//	ResourceIdentityTypeSystemAssigned
//	ResourceIdentityTypeUserAssigned
//	ResourceIdentityType_SystemAssigned_UserAssigned
//	ResourceIdentityTypeNone
type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthenticationMethodOutput{})
	pulumi.RegisterOutputType(AuthenticationMethodPtrOutput{})
	pulumi.RegisterOutputType(ManagedCassandraProvisioningStateOutput{})
	pulumi.RegisterOutputType(ManagedCassandraProvisioningStatePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
