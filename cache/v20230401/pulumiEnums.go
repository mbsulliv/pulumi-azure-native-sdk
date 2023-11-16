// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Day of the week when a cache can be patched.
type DayOfWeek string

const (
	DayOfWeekMonday    = DayOfWeek("Monday")
	DayOfWeekTuesday   = DayOfWeek("Tuesday")
	DayOfWeekWednesday = DayOfWeek("Wednesday")
	DayOfWeekThursday  = DayOfWeek("Thursday")
	DayOfWeekFriday    = DayOfWeek("Friday")
	DayOfWeekSaturday  = DayOfWeek("Saturday")
	DayOfWeekSunday    = DayOfWeek("Sunday")
	DayOfWeekEveryday  = DayOfWeek("Everyday")
	DayOfWeekWeekend   = DayOfWeek("Weekend")
)

func (DayOfWeek) ElementType() reflect.Type {
	return reflect.TypeOf((*DayOfWeek)(nil)).Elem()
}

func (e DayOfWeek) ToDayOfWeekOutput() DayOfWeekOutput {
	return pulumi.ToOutput(e).(DayOfWeekOutput)
}

func (e DayOfWeek) ToDayOfWeekOutputWithContext(ctx context.Context) DayOfWeekOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DayOfWeekOutput)
}

func (e DayOfWeek) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return e.ToDayOfWeekPtrOutputWithContext(context.Background())
}

func (e DayOfWeek) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return DayOfWeek(e).ToDayOfWeekOutputWithContext(ctx).ToDayOfWeekPtrOutputWithContext(ctx)
}

func (e DayOfWeek) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DayOfWeek) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DayOfWeek) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DayOfWeek) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DayOfWeekOutput struct{ *pulumi.OutputState }

func (DayOfWeekOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DayOfWeek)(nil)).Elem()
}

func (o DayOfWeekOutput) ToDayOfWeekOutput() DayOfWeekOutput {
	return o
}

func (o DayOfWeekOutput) ToDayOfWeekOutputWithContext(ctx context.Context) DayOfWeekOutput {
	return o
}

func (o DayOfWeekOutput) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return o.ToDayOfWeekPtrOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DayOfWeek) *DayOfWeek {
		return &v
	}).(DayOfWeekPtrOutput)
}

func (o DayOfWeekOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DayOfWeek) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DayOfWeekOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DayOfWeek) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DayOfWeekPtrOutput struct{ *pulumi.OutputState }

func (DayOfWeekPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DayOfWeek)(nil)).Elem()
}

func (o DayOfWeekPtrOutput) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return o
}

func (o DayOfWeekPtrOutput) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return o
}

func (o DayOfWeekPtrOutput) Elem() DayOfWeekOutput {
	return o.ApplyT(func(v *DayOfWeek) DayOfWeek {
		if v != nil {
			return *v
		}
		var ret DayOfWeek
		return ret
	}).(DayOfWeekOutput)
}

func (o DayOfWeekPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DayOfWeekPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DayOfWeek) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DayOfWeekInput is an input type that accepts DayOfWeekArgs and DayOfWeekOutput values.
// You can construct a concrete instance of `DayOfWeekInput` via:
//
//	DayOfWeekArgs{...}
type DayOfWeekInput interface {
	pulumi.Input

	ToDayOfWeekOutput() DayOfWeekOutput
	ToDayOfWeekOutputWithContext(context.Context) DayOfWeekOutput
}

var dayOfWeekPtrType = reflect.TypeOf((**DayOfWeek)(nil)).Elem()

type DayOfWeekPtrInput interface {
	pulumi.Input

	ToDayOfWeekPtrOutput() DayOfWeekPtrOutput
	ToDayOfWeekPtrOutputWithContext(context.Context) DayOfWeekPtrOutput
}

type dayOfWeekPtr string

func DayOfWeekPtr(v string) DayOfWeekPtrInput {
	return (*dayOfWeekPtr)(&v)
}

func (*dayOfWeekPtr) ElementType() reflect.Type {
	return dayOfWeekPtrType
}

func (in *dayOfWeekPtr) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return pulumi.ToOutput(in).(DayOfWeekPtrOutput)
}

func (in *dayOfWeekPtr) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DayOfWeekPtrOutput)
}

func (in *dayOfWeekPtr) ToOutput(ctx context.Context) pulumix.Output[*DayOfWeek] {
	return pulumix.Output[*DayOfWeek]{
		OutputState: in.ToDayOfWeekPtrOutputWithContext(ctx).OutputState,
	}
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned, UserAssigned")
)

// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

// Whether or not public endpoint access is allowed for this cache.  Value is optional, but if passed in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method. Default value is 'Enabled'. Note: This setting is important for caches with private endpoints. It has *no effect* on caches that are joined to, or injected into, a virtual network subnet.
type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

// Role of the linked server.
type ReplicationRole string

const (
	ReplicationRolePrimary   = ReplicationRole("Primary")
	ReplicationRoleSecondary = ReplicationRole("Secondary")
)

func (ReplicationRole) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationRole)(nil)).Elem()
}

func (e ReplicationRole) ToReplicationRoleOutput() ReplicationRoleOutput {
	return pulumi.ToOutput(e).(ReplicationRoleOutput)
}

func (e ReplicationRole) ToReplicationRoleOutputWithContext(ctx context.Context) ReplicationRoleOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ReplicationRoleOutput)
}

func (e ReplicationRole) ToReplicationRolePtrOutput() ReplicationRolePtrOutput {
	return e.ToReplicationRolePtrOutputWithContext(context.Background())
}

func (e ReplicationRole) ToReplicationRolePtrOutputWithContext(ctx context.Context) ReplicationRolePtrOutput {
	return ReplicationRole(e).ToReplicationRoleOutputWithContext(ctx).ToReplicationRolePtrOutputWithContext(ctx)
}

func (e ReplicationRole) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReplicationRole) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReplicationRole) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReplicationRole) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ReplicationRoleOutput struct{ *pulumi.OutputState }

func (ReplicationRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationRole)(nil)).Elem()
}

func (o ReplicationRoleOutput) ToReplicationRoleOutput() ReplicationRoleOutput {
	return o
}

func (o ReplicationRoleOutput) ToReplicationRoleOutputWithContext(ctx context.Context) ReplicationRoleOutput {
	return o
}

func (o ReplicationRoleOutput) ToReplicationRolePtrOutput() ReplicationRolePtrOutput {
	return o.ToReplicationRolePtrOutputWithContext(context.Background())
}

func (o ReplicationRoleOutput) ToReplicationRolePtrOutputWithContext(ctx context.Context) ReplicationRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicationRole) *ReplicationRole {
		return &v
	}).(ReplicationRolePtrOutput)
}

func (o ReplicationRoleOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ReplicationRoleOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReplicationRole) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ReplicationRoleOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReplicationRoleOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReplicationRole) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ReplicationRolePtrOutput struct{ *pulumi.OutputState }

func (ReplicationRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationRole)(nil)).Elem()
}

func (o ReplicationRolePtrOutput) ToReplicationRolePtrOutput() ReplicationRolePtrOutput {
	return o
}

func (o ReplicationRolePtrOutput) ToReplicationRolePtrOutputWithContext(ctx context.Context) ReplicationRolePtrOutput {
	return o
}

func (o ReplicationRolePtrOutput) Elem() ReplicationRoleOutput {
	return o.ApplyT(func(v *ReplicationRole) ReplicationRole {
		if v != nil {
			return *v
		}
		var ret ReplicationRole
		return ret
	}).(ReplicationRoleOutput)
}

func (o ReplicationRolePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReplicationRolePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ReplicationRole) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ReplicationRoleInput is an input type that accepts ReplicationRoleArgs and ReplicationRoleOutput values.
// You can construct a concrete instance of `ReplicationRoleInput` via:
//
//	ReplicationRoleArgs{...}
type ReplicationRoleInput interface {
	pulumi.Input

	ToReplicationRoleOutput() ReplicationRoleOutput
	ToReplicationRoleOutputWithContext(context.Context) ReplicationRoleOutput
}

var replicationRolePtrType = reflect.TypeOf((**ReplicationRole)(nil)).Elem()

type ReplicationRolePtrInput interface {
	pulumi.Input

	ToReplicationRolePtrOutput() ReplicationRolePtrOutput
	ToReplicationRolePtrOutputWithContext(context.Context) ReplicationRolePtrOutput
}

type replicationRolePtr string

func ReplicationRolePtr(v string) ReplicationRolePtrInput {
	return (*replicationRolePtr)(&v)
}

func (*replicationRolePtr) ElementType() reflect.Type {
	return replicationRolePtrType
}

func (in *replicationRolePtr) ToReplicationRolePtrOutput() ReplicationRolePtrOutput {
	return pulumi.ToOutput(in).(ReplicationRolePtrOutput)
}

func (in *replicationRolePtr) ToReplicationRolePtrOutputWithContext(ctx context.Context) ReplicationRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ReplicationRolePtrOutput)
}

func (in *replicationRolePtr) ToOutput(ctx context.Context) pulumix.Output[*ReplicationRole] {
	return pulumix.Output[*ReplicationRole]{
		OutputState: in.ToReplicationRolePtrOutputWithContext(ctx).OutputState,
	}
}

// The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
type SkuFamily string

const (
	SkuFamilyC = SkuFamily("C")
	SkuFamilyP = SkuFamily("P")
)

// The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
type SkuName string

const (
	SkuNameBasic    = SkuName("Basic")
	SkuNameStandard = SkuName("Standard")
	SkuNamePremium  = SkuName("Premium")
)

// Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1', '1.2')
type TlsVersion string

const (
	TlsVersion_1_0 = TlsVersion("1.0")
	TlsVersion_1_1 = TlsVersion("1.1")
	TlsVersion_1_2 = TlsVersion("1.2")
)

func init() {
	pulumi.RegisterOutputType(DayOfWeekOutput{})
	pulumi.RegisterOutputType(DayOfWeekPtrOutput{})
	pulumi.RegisterOutputType(ReplicationRoleOutput{})
	pulumi.RegisterOutputType(ReplicationRolePtrOutput{})
}
