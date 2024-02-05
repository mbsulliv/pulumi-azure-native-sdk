// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package peering

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of direct peering.
type DirectPeeringType string

const (
	DirectPeeringTypeEdge                 = DirectPeeringType("Edge")
	DirectPeeringTypeTransit              = DirectPeeringType("Transit")
	DirectPeeringTypeCdn                  = DirectPeeringType("Cdn")
	DirectPeeringTypeInternal             = DirectPeeringType("Internal")
	DirectPeeringTypeIx                   = DirectPeeringType("Ix")
	DirectPeeringTypeIxRs                 = DirectPeeringType("IxRs")
	DirectPeeringTypeVoice                = DirectPeeringType("Voice")
	DirectPeeringTypeEdgeZoneForOperators = DirectPeeringType("EdgeZoneForOperators")
)

func (DirectPeeringType) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectPeeringType)(nil)).Elem()
}

func (e DirectPeeringType) ToDirectPeeringTypeOutput() DirectPeeringTypeOutput {
	return pulumi.ToOutput(e).(DirectPeeringTypeOutput)
}

func (e DirectPeeringType) ToDirectPeeringTypeOutputWithContext(ctx context.Context) DirectPeeringTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DirectPeeringTypeOutput)
}

func (e DirectPeeringType) ToDirectPeeringTypePtrOutput() DirectPeeringTypePtrOutput {
	return e.ToDirectPeeringTypePtrOutputWithContext(context.Background())
}

func (e DirectPeeringType) ToDirectPeeringTypePtrOutputWithContext(ctx context.Context) DirectPeeringTypePtrOutput {
	return DirectPeeringType(e).ToDirectPeeringTypeOutputWithContext(ctx).ToDirectPeeringTypePtrOutputWithContext(ctx)
}

func (e DirectPeeringType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DirectPeeringType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DirectPeeringType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DirectPeeringType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DirectPeeringTypeOutput struct{ *pulumi.OutputState }

func (DirectPeeringTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectPeeringType)(nil)).Elem()
}

func (o DirectPeeringTypeOutput) ToDirectPeeringTypeOutput() DirectPeeringTypeOutput {
	return o
}

func (o DirectPeeringTypeOutput) ToDirectPeeringTypeOutputWithContext(ctx context.Context) DirectPeeringTypeOutput {
	return o
}

func (o DirectPeeringTypeOutput) ToDirectPeeringTypePtrOutput() DirectPeeringTypePtrOutput {
	return o.ToDirectPeeringTypePtrOutputWithContext(context.Background())
}

func (o DirectPeeringTypeOutput) ToDirectPeeringTypePtrOutputWithContext(ctx context.Context) DirectPeeringTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DirectPeeringType) *DirectPeeringType {
		return &v
	}).(DirectPeeringTypePtrOutput)
}

func (o DirectPeeringTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DirectPeeringTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DirectPeeringType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DirectPeeringTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DirectPeeringTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DirectPeeringType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DirectPeeringTypePtrOutput struct{ *pulumi.OutputState }

func (DirectPeeringTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectPeeringType)(nil)).Elem()
}

func (o DirectPeeringTypePtrOutput) ToDirectPeeringTypePtrOutput() DirectPeeringTypePtrOutput {
	return o
}

func (o DirectPeeringTypePtrOutput) ToDirectPeeringTypePtrOutputWithContext(ctx context.Context) DirectPeeringTypePtrOutput {
	return o
}

func (o DirectPeeringTypePtrOutput) Elem() DirectPeeringTypeOutput {
	return o.ApplyT(func(v *DirectPeeringType) DirectPeeringType {
		if v != nil {
			return *v
		}
		var ret DirectPeeringType
		return ret
	}).(DirectPeeringTypeOutput)
}

func (o DirectPeeringTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DirectPeeringTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DirectPeeringType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DirectPeeringTypeInput is an input type that accepts values of the DirectPeeringType enum
// A concrete instance of `DirectPeeringTypeInput` can be one of the following:
//
//	DirectPeeringTypeEdge
//	DirectPeeringTypeTransit
//	DirectPeeringTypeCdn
//	DirectPeeringTypeInternal
//	DirectPeeringTypeIx
//	DirectPeeringTypeIxRs
//	DirectPeeringTypeVoice
//	DirectPeeringTypeEdgeZoneForOperators
type DirectPeeringTypeInput interface {
	pulumi.Input

	ToDirectPeeringTypeOutput() DirectPeeringTypeOutput
	ToDirectPeeringTypeOutputWithContext(context.Context) DirectPeeringTypeOutput
}

var directPeeringTypePtrType = reflect.TypeOf((**DirectPeeringType)(nil)).Elem()

type DirectPeeringTypePtrInput interface {
	pulumi.Input

	ToDirectPeeringTypePtrOutput() DirectPeeringTypePtrOutput
	ToDirectPeeringTypePtrOutputWithContext(context.Context) DirectPeeringTypePtrOutput
}

type directPeeringTypePtr string

func DirectPeeringTypePtr(v string) DirectPeeringTypePtrInput {
	return (*directPeeringTypePtr)(&v)
}

func (*directPeeringTypePtr) ElementType() reflect.Type {
	return directPeeringTypePtrType
}

func (in *directPeeringTypePtr) ToDirectPeeringTypePtrOutput() DirectPeeringTypePtrOutput {
	return pulumi.ToOutput(in).(DirectPeeringTypePtrOutput)
}

func (in *directPeeringTypePtr) ToDirectPeeringTypePtrOutputWithContext(ctx context.Context) DirectPeeringTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DirectPeeringTypePtrOutput)
}

// The kind of the peering.
type Kind string

const (
	KindDirect   = Kind("Direct")
	KindExchange = Kind("Exchange")
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
//	KindDirect
//	KindExchange
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

// The role of the contact.
type Role string

const (
	RoleNoc        = Role("Noc")
	RolePolicy     = Role("Policy")
	RoleTechnical  = Role("Technical")
	RoleService    = Role("Service")
	RoleEscalation = Role("Escalation")
	RoleOther      = Role("Other")
)

func (Role) ElementType() reflect.Type {
	return reflect.TypeOf((*Role)(nil)).Elem()
}

func (e Role) ToRoleOutput() RoleOutput {
	return pulumi.ToOutput(e).(RoleOutput)
}

func (e Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RoleOutput)
}

func (e Role) ToRolePtrOutput() RolePtrOutput {
	return e.ToRolePtrOutputWithContext(context.Background())
}

func (e Role) ToRolePtrOutputWithContext(ctx context.Context) RolePtrOutput {
	return Role(e).ToRoleOutputWithContext(ctx).ToRolePtrOutputWithContext(ctx)
}

func (e Role) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Role) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Role) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Role) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RoleOutput struct{ *pulumi.OutputState }

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Role)(nil)).Elem()
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

func (o RoleOutput) ToRolePtrOutput() RolePtrOutput {
	return o.ToRolePtrOutputWithContext(context.Background())
}

func (o RoleOutput) ToRolePtrOutputWithContext(ctx context.Context) RolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Role) *Role {
		return &v
	}).(RolePtrOutput)
}

func (o RoleOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RoleOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Role) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RoleOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoleOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Role) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RolePtrOutput struct{ *pulumi.OutputState }

func (RolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (o RolePtrOutput) ToRolePtrOutput() RolePtrOutput {
	return o
}

func (o RolePtrOutput) ToRolePtrOutputWithContext(ctx context.Context) RolePtrOutput {
	return o
}

func (o RolePtrOutput) Elem() RoleOutput {
	return o.ApplyT(func(v *Role) Role {
		if v != nil {
			return *v
		}
		var ret Role
		return ret
	}).(RoleOutput)
}

func (o RolePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RolePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Role) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// RoleInput is an input type that accepts values of the Role enum
// A concrete instance of `RoleInput` can be one of the following:
//
//	RoleNoc
//	RolePolicy
//	RoleTechnical
//	RoleService
//	RoleEscalation
//	RoleOther
type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(context.Context) RoleOutput
}

var rolePtrType = reflect.TypeOf((**Role)(nil)).Elem()

type RolePtrInput interface {
	pulumi.Input

	ToRolePtrOutput() RolePtrOutput
	ToRolePtrOutputWithContext(context.Context) RolePtrOutput
}

type rolePtr string

func RolePtr(v string) RolePtrInput {
	return (*rolePtr)(&v)
}

func (*rolePtr) ElementType() reflect.Type {
	return rolePtrType
}

func (in *rolePtr) ToRolePtrOutput() RolePtrOutput {
	return pulumi.ToOutput(in).(RolePtrOutput)
}

func (in *rolePtr) ToRolePtrOutputWithContext(ctx context.Context) RolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RolePtrOutput)
}

// The field indicating if Microsoft provides session ip addresses.
type SessionAddressProvider string

const (
	SessionAddressProviderMicrosoft = SessionAddressProvider("Microsoft")
	SessionAddressProviderPeer      = SessionAddressProvider("Peer")
)

func (SessionAddressProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*SessionAddressProvider)(nil)).Elem()
}

func (e SessionAddressProvider) ToSessionAddressProviderOutput() SessionAddressProviderOutput {
	return pulumi.ToOutput(e).(SessionAddressProviderOutput)
}

func (e SessionAddressProvider) ToSessionAddressProviderOutputWithContext(ctx context.Context) SessionAddressProviderOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SessionAddressProviderOutput)
}

func (e SessionAddressProvider) ToSessionAddressProviderPtrOutput() SessionAddressProviderPtrOutput {
	return e.ToSessionAddressProviderPtrOutputWithContext(context.Background())
}

func (e SessionAddressProvider) ToSessionAddressProviderPtrOutputWithContext(ctx context.Context) SessionAddressProviderPtrOutput {
	return SessionAddressProvider(e).ToSessionAddressProviderOutputWithContext(ctx).ToSessionAddressProviderPtrOutputWithContext(ctx)
}

func (e SessionAddressProvider) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SessionAddressProvider) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SessionAddressProvider) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SessionAddressProvider) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SessionAddressProviderOutput struct{ *pulumi.OutputState }

func (SessionAddressProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SessionAddressProvider)(nil)).Elem()
}

func (o SessionAddressProviderOutput) ToSessionAddressProviderOutput() SessionAddressProviderOutput {
	return o
}

func (o SessionAddressProviderOutput) ToSessionAddressProviderOutputWithContext(ctx context.Context) SessionAddressProviderOutput {
	return o
}

func (o SessionAddressProviderOutput) ToSessionAddressProviderPtrOutput() SessionAddressProviderPtrOutput {
	return o.ToSessionAddressProviderPtrOutputWithContext(context.Background())
}

func (o SessionAddressProviderOutput) ToSessionAddressProviderPtrOutputWithContext(ctx context.Context) SessionAddressProviderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SessionAddressProvider) *SessionAddressProvider {
		return &v
	}).(SessionAddressProviderPtrOutput)
}

func (o SessionAddressProviderOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SessionAddressProviderOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SessionAddressProvider) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SessionAddressProviderOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SessionAddressProviderOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SessionAddressProvider) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SessionAddressProviderPtrOutput struct{ *pulumi.OutputState }

func (SessionAddressProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SessionAddressProvider)(nil)).Elem()
}

func (o SessionAddressProviderPtrOutput) ToSessionAddressProviderPtrOutput() SessionAddressProviderPtrOutput {
	return o
}

func (o SessionAddressProviderPtrOutput) ToSessionAddressProviderPtrOutputWithContext(ctx context.Context) SessionAddressProviderPtrOutput {
	return o
}

func (o SessionAddressProviderPtrOutput) Elem() SessionAddressProviderOutput {
	return o.ApplyT(func(v *SessionAddressProvider) SessionAddressProvider {
		if v != nil {
			return *v
		}
		var ret SessionAddressProvider
		return ret
	}).(SessionAddressProviderOutput)
}

func (o SessionAddressProviderPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SessionAddressProviderPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SessionAddressProvider) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SessionAddressProviderInput is an input type that accepts values of the SessionAddressProvider enum
// A concrete instance of `SessionAddressProviderInput` can be one of the following:
//
//	SessionAddressProviderMicrosoft
//	SessionAddressProviderPeer
type SessionAddressProviderInput interface {
	pulumi.Input

	ToSessionAddressProviderOutput() SessionAddressProviderOutput
	ToSessionAddressProviderOutputWithContext(context.Context) SessionAddressProviderOutput
}

var sessionAddressProviderPtrType = reflect.TypeOf((**SessionAddressProvider)(nil)).Elem()

type SessionAddressProviderPtrInput interface {
	pulumi.Input

	ToSessionAddressProviderPtrOutput() SessionAddressProviderPtrOutput
	ToSessionAddressProviderPtrOutputWithContext(context.Context) SessionAddressProviderPtrOutput
}

type sessionAddressProviderPtr string

func SessionAddressProviderPtr(v string) SessionAddressProviderPtrInput {
	return (*sessionAddressProviderPtr)(&v)
}

func (*sessionAddressProviderPtr) ElementType() reflect.Type {
	return sessionAddressProviderPtrType
}

func (in *sessionAddressProviderPtr) ToSessionAddressProviderPtrOutput() SessionAddressProviderPtrOutput {
	return pulumi.ToOutput(in).(SessionAddressProviderPtrOutput)
}

func (in *sessionAddressProviderPtr) ToSessionAddressProviderPtrOutputWithContext(ctx context.Context) SessionAddressProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SessionAddressProviderPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DirectPeeringTypeOutput{})
	pulumi.RegisterOutputType(DirectPeeringTypePtrOutput{})
	pulumi.RegisterOutputType(KindOutput{})
	pulumi.RegisterOutputType(KindPtrOutput{})
	pulumi.RegisterOutputType(RoleOutput{})
	pulumi.RegisterOutputType(RolePtrOutput{})
	pulumi.RegisterOutputType(SessionAddressProviderOutput{})
	pulumi.RegisterOutputType(SessionAddressProviderPtrOutput{})
}
