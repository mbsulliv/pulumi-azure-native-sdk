// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Values can be systemAssignedIdentity or userAssignedIdentity
type IdentityType string

const (
	IdentityTypeSystemAssignedIdentity    = IdentityType("systemAssignedIdentity")
	IdentityTypeUserAssignedIdentity      = IdentityType("userAssignedIdentity")
	IdentityTypeDelegatedResourceIdentity = IdentityType("delegatedResourceIdentity")
)

func (IdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (e IdentityType) ToIdentityTypeOutput() IdentityTypeOutput {
	return pulumi.ToOutput(e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return e.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (e IdentityType) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return IdentityType(e).ToIdentityTypeOutputWithContext(ctx).ToIdentityTypePtrOutputWithContext(ctx)
}

func (e IdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IdentityTypeOutput struct{ *pulumi.OutputState }

func (IdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (o IdentityTypeOutput) ToIdentityTypeOutput() IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityType) *IdentityType {
		return &v
	}).(IdentityTypePtrOutput)
}

func (o IdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentityTypePtrOutput struct{ *pulumi.OutputState }

func (IdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityType)(nil)).Elem()
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) Elem() IdentityTypeOutput {
	return o.ApplyT(func(v *IdentityType) IdentityType {
		if v != nil {
			return *v
		}
		var ret IdentityType
		return ret
	}).(IdentityTypeOutput)
}

func (o IdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IdentityTypeInput is an input type that accepts values of the IdentityType enum
// A concrete instance of `IdentityTypeInput` can be one of the following:
//
//	IdentityTypeSystemAssignedIdentity
//	IdentityTypeUserAssignedIdentity
//	IdentityTypeDelegatedResourceIdentity
type IdentityTypeInput interface {
	pulumi.Input

	ToIdentityTypeOutput() IdentityTypeOutput
	ToIdentityTypeOutputWithContext(context.Context) IdentityTypeOutput
}

var identityTypePtrType = reflect.TypeOf((**IdentityType)(nil)).Elem()

type IdentityTypePtrInput interface {
	pulumi.Input

	ToIdentityTypePtrOutput() IdentityTypePtrOutput
	ToIdentityTypePtrOutputWithContext(context.Context) IdentityTypePtrOutput
}

type identityTypePtr string

func IdentityTypePtr(v string) IdentityTypePtrInput {
	return (*identityTypePtr)(&v)
}

func (*identityTypePtr) ElementType() reflect.Type {
	return identityTypePtrType
}

func (in *identityTypePtr) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return pulumi.ToOutput(in).(IdentityTypePtrOutput)
}

func (in *identityTypePtr) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IdentityTypePtrOutput)
}

// Values are enabled and disabled.
type InfrastructureEncryption string

const (
	InfrastructureEncryptionEnabled  = InfrastructureEncryption("enabled")
	InfrastructureEncryptionDisabled = InfrastructureEncryption("disabled")
)

func (InfrastructureEncryption) ElementType() reflect.Type {
	return reflect.TypeOf((*InfrastructureEncryption)(nil)).Elem()
}

func (e InfrastructureEncryption) ToInfrastructureEncryptionOutput() InfrastructureEncryptionOutput {
	return pulumi.ToOutput(e).(InfrastructureEncryptionOutput)
}

func (e InfrastructureEncryption) ToInfrastructureEncryptionOutputWithContext(ctx context.Context) InfrastructureEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InfrastructureEncryptionOutput)
}

func (e InfrastructureEncryption) ToInfrastructureEncryptionPtrOutput() InfrastructureEncryptionPtrOutput {
	return e.ToInfrastructureEncryptionPtrOutputWithContext(context.Background())
}

func (e InfrastructureEncryption) ToInfrastructureEncryptionPtrOutputWithContext(ctx context.Context) InfrastructureEncryptionPtrOutput {
	return InfrastructureEncryption(e).ToInfrastructureEncryptionOutputWithContext(ctx).ToInfrastructureEncryptionPtrOutputWithContext(ctx)
}

func (e InfrastructureEncryption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InfrastructureEncryption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InfrastructureEncryption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InfrastructureEncryption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InfrastructureEncryptionOutput struct{ *pulumi.OutputState }

func (InfrastructureEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InfrastructureEncryption)(nil)).Elem()
}

func (o InfrastructureEncryptionOutput) ToInfrastructureEncryptionOutput() InfrastructureEncryptionOutput {
	return o
}

func (o InfrastructureEncryptionOutput) ToInfrastructureEncryptionOutputWithContext(ctx context.Context) InfrastructureEncryptionOutput {
	return o
}

func (o InfrastructureEncryptionOutput) ToInfrastructureEncryptionPtrOutput() InfrastructureEncryptionPtrOutput {
	return o.ToInfrastructureEncryptionPtrOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionOutput) ToInfrastructureEncryptionPtrOutputWithContext(ctx context.Context) InfrastructureEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InfrastructureEncryption) *InfrastructureEncryption {
		return &v
	}).(InfrastructureEncryptionPtrOutput)
}

func (o InfrastructureEncryptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InfrastructureEncryption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InfrastructureEncryptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InfrastructureEncryption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InfrastructureEncryptionPtrOutput struct{ *pulumi.OutputState }

func (InfrastructureEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InfrastructureEncryption)(nil)).Elem()
}

func (o InfrastructureEncryptionPtrOutput) ToInfrastructureEncryptionPtrOutput() InfrastructureEncryptionPtrOutput {
	return o
}

func (o InfrastructureEncryptionPtrOutput) ToInfrastructureEncryptionPtrOutputWithContext(ctx context.Context) InfrastructureEncryptionPtrOutput {
	return o
}

func (o InfrastructureEncryptionPtrOutput) Elem() InfrastructureEncryptionOutput {
	return o.ApplyT(func(v *InfrastructureEncryption) InfrastructureEncryption {
		if v != nil {
			return *v
		}
		var ret InfrastructureEncryption
		return ret
	}).(InfrastructureEncryptionOutput)
}

func (o InfrastructureEncryptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InfrastructureEncryption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// InfrastructureEncryptionInput is an input type that accepts values of the InfrastructureEncryption enum
// A concrete instance of `InfrastructureEncryptionInput` can be one of the following:
//
//	InfrastructureEncryptionEnabled
//	InfrastructureEncryptionDisabled
type InfrastructureEncryptionInput interface {
	pulumi.Input

	ToInfrastructureEncryptionOutput() InfrastructureEncryptionOutput
	ToInfrastructureEncryptionOutputWithContext(context.Context) InfrastructureEncryptionOutput
}

var infrastructureEncryptionPtrType = reflect.TypeOf((**InfrastructureEncryption)(nil)).Elem()

type InfrastructureEncryptionPtrInput interface {
	pulumi.Input

	ToInfrastructureEncryptionPtrOutput() InfrastructureEncryptionPtrOutput
	ToInfrastructureEncryptionPtrOutputWithContext(context.Context) InfrastructureEncryptionPtrOutput
}

type infrastructureEncryptionPtr string

func InfrastructureEncryptionPtr(v string) InfrastructureEncryptionPtrInput {
	return (*infrastructureEncryptionPtr)(&v)
}

func (*infrastructureEncryptionPtr) ElementType() reflect.Type {
	return infrastructureEncryptionPtrType
}

func (in *infrastructureEncryptionPtr) ToInfrastructureEncryptionPtrOutput() InfrastructureEncryptionPtrOutput {
	return pulumi.ToOutput(in).(InfrastructureEncryptionPtrOutput)
}

func (in *infrastructureEncryptionPtr) ToInfrastructureEncryptionPtrOutputWithContext(ctx context.Context) InfrastructureEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InfrastructureEncryptionPtrOutput)
}

// Get or Set Kind property.
type Kind string

const (
	KindGen1 = Kind("Gen1")
	KindGen2 = Kind("Gen2")
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
//	KindGen1
//	KindGen2
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

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned, UserAssigned")
)

func (ManagedServiceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return e.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return ManagedServiceIdentityType(e).ToManagedServiceIdentityTypeOutputWithContext(ctx).ToManagedServiceIdentityTypePtrOutputWithContext(ctx)
}

func (e ManagedServiceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedServiceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentityType) *ManagedServiceIdentityType {
		return &v
	}).(ManagedServiceIdentityTypePtrOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedServiceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) Elem() ManagedServiceIdentityTypeOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityType) ManagedServiceIdentityType {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityType
		return ret
	}).(ManagedServiceIdentityTypeOutput)
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedServiceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ManagedServiceIdentityTypeInput is an input type that accepts values of the ManagedServiceIdentityType enum
// A concrete instance of `ManagedServiceIdentityTypeInput` can be one of the following:
//
//	ManagedServiceIdentityTypeNone
//	ManagedServiceIdentityTypeSystemAssigned
//	ManagedServiceIdentityTypeUserAssigned
//	ManagedServiceIdentityType_SystemAssigned_UserAssigned
type ManagedServiceIdentityTypeInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput
	ToManagedServiceIdentityTypeOutputWithContext(context.Context) ManagedServiceIdentityTypeOutput
}

var managedServiceIdentityTypePtrType = reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()

type ManagedServiceIdentityTypePtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput
	ToManagedServiceIdentityTypePtrOutputWithContext(context.Context) ManagedServiceIdentityTypePtrOutput
}

type managedServiceIdentityTypePtr string

func ManagedServiceIdentityTypePtr(v string) ManagedServiceIdentityTypePtrInput {
	return (*managedServiceIdentityTypePtr)(&v)
}

func (*managedServiceIdentityTypePtr) ElementType() reflect.Type {
	return managedServiceIdentityTypePtrType
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ManagedServiceIdentityTypePtrOutput)
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedServiceIdentityTypePtrOutput)
}

// The name of the SKU, in standard format (such as S0).
type Name string

const (
	NameS0 = Name("S0")
	NameS1 = Name("S1")
	NameG2 = Name("G2")
)

func (Name) ElementType() reflect.Type {
	return reflect.TypeOf((*Name)(nil)).Elem()
}

func (e Name) ToNameOutput() NameOutput {
	return pulumi.ToOutput(e).(NameOutput)
}

func (e Name) ToNameOutputWithContext(ctx context.Context) NameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NameOutput)
}

func (e Name) ToNamePtrOutput() NamePtrOutput {
	return e.ToNamePtrOutputWithContext(context.Background())
}

func (e Name) ToNamePtrOutputWithContext(ctx context.Context) NamePtrOutput {
	return Name(e).ToNameOutputWithContext(ctx).ToNamePtrOutputWithContext(ctx)
}

func (e Name) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Name) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Name) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Name) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NameOutput struct{ *pulumi.OutputState }

func (NameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Name)(nil)).Elem()
}

func (o NameOutput) ToNameOutput() NameOutput {
	return o
}

func (o NameOutput) ToNameOutputWithContext(ctx context.Context) NameOutput {
	return o
}

func (o NameOutput) ToNamePtrOutput() NamePtrOutput {
	return o.ToNamePtrOutputWithContext(context.Background())
}

func (o NameOutput) ToNamePtrOutputWithContext(ctx context.Context) NamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Name) *Name {
		return &v
	}).(NamePtrOutput)
}

func (o NameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Name) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Name) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NamePtrOutput struct{ *pulumi.OutputState }

func (NamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Name)(nil)).Elem()
}

func (o NamePtrOutput) ToNamePtrOutput() NamePtrOutput {
	return o
}

func (o NamePtrOutput) ToNamePtrOutputWithContext(ctx context.Context) NamePtrOutput {
	return o
}

func (o NamePtrOutput) Elem() NameOutput {
	return o.ApplyT(func(v *Name) Name {
		if v != nil {
			return *v
		}
		var ret Name
		return ret
	}).(NameOutput)
}

func (o NamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Name) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// NameInput is an input type that accepts values of the Name enum
// A concrete instance of `NameInput` can be one of the following:
//
//	NameS0
//	NameS1
//	NameG2
type NameInput interface {
	pulumi.Input

	ToNameOutput() NameOutput
	ToNameOutputWithContext(context.Context) NameOutput
}

var namePtrType = reflect.TypeOf((**Name)(nil)).Elem()

type NamePtrInput interface {
	pulumi.Input

	ToNamePtrOutput() NamePtrOutput
	ToNamePtrOutputWithContext(context.Context) NamePtrOutput
}

type namePtr string

func NamePtr(v string) NamePtrInput {
	return (*namePtr)(&v)
}

func (*namePtr) ElementType() reflect.Type {
	return namePtrType
}

func (in *namePtr) ToNamePtrOutput() NamePtrOutput {
	return pulumi.ToOutput(in).(NamePtrOutput)
}

func (in *namePtr) ToNamePtrOutputWithContext(ctx context.Context) NamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NamePtrOutput)
}

// The Map account key to use for signing. Picking `primaryKey` or `secondaryKey` will use the Map account Shared Keys, and using `managedIdentity` will use the auto-renewed private key to sign the SAS.
type SigningKey string

const (
	SigningKeyPrimaryKey      = SigningKey("primaryKey")
	SigningKeySecondaryKey    = SigningKey("secondaryKey")
	SigningKeyManagedIdentity = SigningKey("managedIdentity")
)

func (SigningKey) ElementType() reflect.Type {
	return reflect.TypeOf((*SigningKey)(nil)).Elem()
}

func (e SigningKey) ToSigningKeyOutput() SigningKeyOutput {
	return pulumi.ToOutput(e).(SigningKeyOutput)
}

func (e SigningKey) ToSigningKeyOutputWithContext(ctx context.Context) SigningKeyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SigningKeyOutput)
}

func (e SigningKey) ToSigningKeyPtrOutput() SigningKeyPtrOutput {
	return e.ToSigningKeyPtrOutputWithContext(context.Background())
}

func (e SigningKey) ToSigningKeyPtrOutputWithContext(ctx context.Context) SigningKeyPtrOutput {
	return SigningKey(e).ToSigningKeyOutputWithContext(ctx).ToSigningKeyPtrOutputWithContext(ctx)
}

func (e SigningKey) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SigningKey) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SigningKey) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SigningKey) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SigningKeyOutput struct{ *pulumi.OutputState }

func (SigningKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SigningKey)(nil)).Elem()
}

func (o SigningKeyOutput) ToSigningKeyOutput() SigningKeyOutput {
	return o
}

func (o SigningKeyOutput) ToSigningKeyOutputWithContext(ctx context.Context) SigningKeyOutput {
	return o
}

func (o SigningKeyOutput) ToSigningKeyPtrOutput() SigningKeyPtrOutput {
	return o.ToSigningKeyPtrOutputWithContext(context.Background())
}

func (o SigningKeyOutput) ToSigningKeyPtrOutputWithContext(ctx context.Context) SigningKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SigningKey) *SigningKey {
		return &v
	}).(SigningKeyPtrOutput)
}

func (o SigningKeyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SigningKeyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SigningKey) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SigningKeyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SigningKeyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SigningKey) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SigningKeyPtrOutput struct{ *pulumi.OutputState }

func (SigningKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SigningKey)(nil)).Elem()
}

func (o SigningKeyPtrOutput) ToSigningKeyPtrOutput() SigningKeyPtrOutput {
	return o
}

func (o SigningKeyPtrOutput) ToSigningKeyPtrOutputWithContext(ctx context.Context) SigningKeyPtrOutput {
	return o
}

func (o SigningKeyPtrOutput) Elem() SigningKeyOutput {
	return o.ApplyT(func(v *SigningKey) SigningKey {
		if v != nil {
			return *v
		}
		var ret SigningKey
		return ret
	}).(SigningKeyOutput)
}

func (o SigningKeyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SigningKeyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SigningKey) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SigningKeyInput is an input type that accepts values of the SigningKey enum
// A concrete instance of `SigningKeyInput` can be one of the following:
//
//	SigningKeyPrimaryKey
//	SigningKeySecondaryKey
//	SigningKeyManagedIdentity
type SigningKeyInput interface {
	pulumi.Input

	ToSigningKeyOutput() SigningKeyOutput
	ToSigningKeyOutputWithContext(context.Context) SigningKeyOutput
}

var signingKeyPtrType = reflect.TypeOf((**SigningKey)(nil)).Elem()

type SigningKeyPtrInput interface {
	pulumi.Input

	ToSigningKeyPtrOutput() SigningKeyPtrOutput
	ToSigningKeyPtrOutputWithContext(context.Context) SigningKeyPtrOutput
}

type signingKeyPtr string

func SigningKeyPtr(v string) SigningKeyPtrInput {
	return (*signingKeyPtr)(&v)
}

func (*signingKeyPtr) ElementType() reflect.Type {
	return signingKeyPtrType
}

func (in *signingKeyPtr) ToSigningKeyPtrOutput() SigningKeyPtrOutput {
	return pulumi.ToOutput(in).(SigningKeyPtrOutput)
}

func (in *signingKeyPtr) ToSigningKeyPtrOutputWithContext(ctx context.Context) SigningKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SigningKeyPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(InfrastructureEncryptionOutput{})
	pulumi.RegisterOutputType(InfrastructureEncryptionPtrOutput{})
	pulumi.RegisterOutputType(KindOutput{})
	pulumi.RegisterOutputType(KindPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(NameOutput{})
	pulumi.RegisterOutputType(NamePtrOutput{})
	pulumi.RegisterOutputType(SigningKeyOutput{})
	pulumi.RegisterOutputType(SigningKeyPtrOutput{})
}
