// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets or sets the state of public network access.
type MasterSitePropertiesPublicNetworkAccess string

const (
	// NotSpecified value.
	MasterSitePropertiesPublicNetworkAccessNotSpecified = MasterSitePropertiesPublicNetworkAccess("NotSpecified")
	// Enabled value.
	MasterSitePropertiesPublicNetworkAccessEnabled = MasterSitePropertiesPublicNetworkAccess("Enabled")
	// Disabled value.
	MasterSitePropertiesPublicNetworkAccessDisabled = MasterSitePropertiesPublicNetworkAccess("Disabled")
)

func (MasterSitePropertiesPublicNetworkAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*MasterSitePropertiesPublicNetworkAccess)(nil)).Elem()
}

func (e MasterSitePropertiesPublicNetworkAccess) ToMasterSitePropertiesPublicNetworkAccessOutput() MasterSitePropertiesPublicNetworkAccessOutput {
	return pulumi.ToOutput(e).(MasterSitePropertiesPublicNetworkAccessOutput)
}

func (e MasterSitePropertiesPublicNetworkAccess) ToMasterSitePropertiesPublicNetworkAccessOutputWithContext(ctx context.Context) MasterSitePropertiesPublicNetworkAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MasterSitePropertiesPublicNetworkAccessOutput)
}

func (e MasterSitePropertiesPublicNetworkAccess) ToMasterSitePropertiesPublicNetworkAccessPtrOutput() MasterSitePropertiesPublicNetworkAccessPtrOutput {
	return e.ToMasterSitePropertiesPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (e MasterSitePropertiesPublicNetworkAccess) ToMasterSitePropertiesPublicNetworkAccessPtrOutputWithContext(ctx context.Context) MasterSitePropertiesPublicNetworkAccessPtrOutput {
	return MasterSitePropertiesPublicNetworkAccess(e).ToMasterSitePropertiesPublicNetworkAccessOutputWithContext(ctx).ToMasterSitePropertiesPublicNetworkAccessPtrOutputWithContext(ctx)
}

func (e MasterSitePropertiesPublicNetworkAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MasterSitePropertiesPublicNetworkAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MasterSitePropertiesPublicNetworkAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MasterSitePropertiesPublicNetworkAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MasterSitePropertiesPublicNetworkAccessOutput struct{ *pulumi.OutputState }

func (MasterSitePropertiesPublicNetworkAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MasterSitePropertiesPublicNetworkAccess)(nil)).Elem()
}

func (o MasterSitePropertiesPublicNetworkAccessOutput) ToMasterSitePropertiesPublicNetworkAccessOutput() MasterSitePropertiesPublicNetworkAccessOutput {
	return o
}

func (o MasterSitePropertiesPublicNetworkAccessOutput) ToMasterSitePropertiesPublicNetworkAccessOutputWithContext(ctx context.Context) MasterSitePropertiesPublicNetworkAccessOutput {
	return o
}

func (o MasterSitePropertiesPublicNetworkAccessOutput) ToMasterSitePropertiesPublicNetworkAccessPtrOutput() MasterSitePropertiesPublicNetworkAccessPtrOutput {
	return o.ToMasterSitePropertiesPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (o MasterSitePropertiesPublicNetworkAccessOutput) ToMasterSitePropertiesPublicNetworkAccessPtrOutputWithContext(ctx context.Context) MasterSitePropertiesPublicNetworkAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MasterSitePropertiesPublicNetworkAccess) *MasterSitePropertiesPublicNetworkAccess {
		return &v
	}).(MasterSitePropertiesPublicNetworkAccessPtrOutput)
}

func (o MasterSitePropertiesPublicNetworkAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MasterSitePropertiesPublicNetworkAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MasterSitePropertiesPublicNetworkAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MasterSitePropertiesPublicNetworkAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MasterSitePropertiesPublicNetworkAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MasterSitePropertiesPublicNetworkAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MasterSitePropertiesPublicNetworkAccessPtrOutput struct{ *pulumi.OutputState }

func (MasterSitePropertiesPublicNetworkAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MasterSitePropertiesPublicNetworkAccess)(nil)).Elem()
}

func (o MasterSitePropertiesPublicNetworkAccessPtrOutput) ToMasterSitePropertiesPublicNetworkAccessPtrOutput() MasterSitePropertiesPublicNetworkAccessPtrOutput {
	return o
}

func (o MasterSitePropertiesPublicNetworkAccessPtrOutput) ToMasterSitePropertiesPublicNetworkAccessPtrOutputWithContext(ctx context.Context) MasterSitePropertiesPublicNetworkAccessPtrOutput {
	return o
}

func (o MasterSitePropertiesPublicNetworkAccessPtrOutput) Elem() MasterSitePropertiesPublicNetworkAccessOutput {
	return o.ApplyT(func(v *MasterSitePropertiesPublicNetworkAccess) MasterSitePropertiesPublicNetworkAccess {
		if v != nil {
			return *v
		}
		var ret MasterSitePropertiesPublicNetworkAccess
		return ret
	}).(MasterSitePropertiesPublicNetworkAccessOutput)
}

func (o MasterSitePropertiesPublicNetworkAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MasterSitePropertiesPublicNetworkAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MasterSitePropertiesPublicNetworkAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MasterSitePropertiesPublicNetworkAccessInput is an input type that accepts values of the MasterSitePropertiesPublicNetworkAccess enum
// A concrete instance of `MasterSitePropertiesPublicNetworkAccessInput` can be one of the following:
//
//	MasterSitePropertiesPublicNetworkAccessNotSpecified
//	MasterSitePropertiesPublicNetworkAccessEnabled
//	MasterSitePropertiesPublicNetworkAccessDisabled
type MasterSitePropertiesPublicNetworkAccessInput interface {
	pulumi.Input

	ToMasterSitePropertiesPublicNetworkAccessOutput() MasterSitePropertiesPublicNetworkAccessOutput
	ToMasterSitePropertiesPublicNetworkAccessOutputWithContext(context.Context) MasterSitePropertiesPublicNetworkAccessOutput
}

var masterSitePropertiesPublicNetworkAccessPtrType = reflect.TypeOf((**MasterSitePropertiesPublicNetworkAccess)(nil)).Elem()

type MasterSitePropertiesPublicNetworkAccessPtrInput interface {
	pulumi.Input

	ToMasterSitePropertiesPublicNetworkAccessPtrOutput() MasterSitePropertiesPublicNetworkAccessPtrOutput
	ToMasterSitePropertiesPublicNetworkAccessPtrOutputWithContext(context.Context) MasterSitePropertiesPublicNetworkAccessPtrOutput
}

type masterSitePropertiesPublicNetworkAccessPtr string

func MasterSitePropertiesPublicNetworkAccessPtr(v string) MasterSitePropertiesPublicNetworkAccessPtrInput {
	return (*masterSitePropertiesPublicNetworkAccessPtr)(&v)
}

func (*masterSitePropertiesPublicNetworkAccessPtr) ElementType() reflect.Type {
	return masterSitePropertiesPublicNetworkAccessPtrType
}

func (in *masterSitePropertiesPublicNetworkAccessPtr) ToMasterSitePropertiesPublicNetworkAccessPtrOutput() MasterSitePropertiesPublicNetworkAccessPtrOutput {
	return pulumi.ToOutput(in).(MasterSitePropertiesPublicNetworkAccessPtrOutput)
}

func (in *masterSitePropertiesPublicNetworkAccessPtr) ToMasterSitePropertiesPublicNetworkAccessPtrOutputWithContext(ctx context.Context) MasterSitePropertiesPublicNetworkAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MasterSitePropertiesPublicNetworkAccessPtrOutput)
}

// state status
type PrivateLinkServiceConnectionStateStatus string

const (
	// Approved value.
	PrivateLinkServiceConnectionStateStatusApproved = PrivateLinkServiceConnectionStateStatus("Approved")
	// Pending value.
	PrivateLinkServiceConnectionStateStatusPending = PrivateLinkServiceConnectionStateStatus("Pending")
	// Rejected value.
	PrivateLinkServiceConnectionStateStatusRejected = PrivateLinkServiceConnectionStateStatus("Rejected")
	// Disconnected value.
	PrivateLinkServiceConnectionStateStatusDisconnected = PrivateLinkServiceConnectionStateStatus("Disconnected")
)

func (PrivateLinkServiceConnectionStateStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateStatus)(nil)).Elem()
}

func (e PrivateLinkServiceConnectionStateStatus) ToPrivateLinkServiceConnectionStateStatusOutput() PrivateLinkServiceConnectionStateStatusOutput {
	return pulumi.ToOutput(e).(PrivateLinkServiceConnectionStateStatusOutput)
}

func (e PrivateLinkServiceConnectionStateStatus) ToPrivateLinkServiceConnectionStateStatusOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateLinkServiceConnectionStateStatusOutput)
}

func (e PrivateLinkServiceConnectionStateStatus) ToPrivateLinkServiceConnectionStateStatusPtrOutput() PrivateLinkServiceConnectionStateStatusPtrOutput {
	return e.ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceConnectionStateStatus) ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateStatusPtrOutput {
	return PrivateLinkServiceConnectionStateStatus(e).ToPrivateLinkServiceConnectionStateStatusOutputWithContext(ctx).ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(ctx)
}

func (e PrivateLinkServiceConnectionStateStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStateStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStateStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceConnectionStateStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateLinkServiceConnectionStateStatusOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateStatus)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToPrivateLinkServiceConnectionStateStatusOutput() PrivateLinkServiceConnectionStateStatusOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToPrivateLinkServiceConnectionStateStatusOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateStatusOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToPrivateLinkServiceConnectionStateStatusPtrOutput() PrivateLinkServiceConnectionStateStatusPtrOutput {
	return o.ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateStatus) *PrivateLinkServiceConnectionStateStatus {
		return &v
	}).(PrivateLinkServiceConnectionStateStatusPtrOutput)
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceConnectionStateStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceConnectionStateStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateStatus)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateStatusPtrOutput) ToPrivateLinkServiceConnectionStateStatusPtrOutput() PrivateLinkServiceConnectionStateStatusPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateStatusPtrOutput) ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateStatusPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateStatusPtrOutput) Elem() PrivateLinkServiceConnectionStateStatusOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateStatus) PrivateLinkServiceConnectionStateStatus {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateStatus
		return ret
	}).(PrivateLinkServiceConnectionStateStatusOutput)
}

func (o PrivateLinkServiceConnectionStateStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateLinkServiceConnectionStateStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PrivateLinkServiceConnectionStateStatusInput is an input type that accepts values of the PrivateLinkServiceConnectionStateStatus enum
// A concrete instance of `PrivateLinkServiceConnectionStateStatusInput` can be one of the following:
//
//	PrivateLinkServiceConnectionStateStatusApproved
//	PrivateLinkServiceConnectionStateStatusPending
//	PrivateLinkServiceConnectionStateStatusRejected
//	PrivateLinkServiceConnectionStateStatusDisconnected
type PrivateLinkServiceConnectionStateStatusInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateStatusOutput() PrivateLinkServiceConnectionStateStatusOutput
	ToPrivateLinkServiceConnectionStateStatusOutputWithContext(context.Context) PrivateLinkServiceConnectionStateStatusOutput
}

var privateLinkServiceConnectionStateStatusPtrType = reflect.TypeOf((**PrivateLinkServiceConnectionStateStatus)(nil)).Elem()

type PrivateLinkServiceConnectionStateStatusPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateStatusPtrOutput() PrivateLinkServiceConnectionStateStatusPtrOutput
	ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateStatusPtrOutput
}

type privateLinkServiceConnectionStateStatusPtr string

func PrivateLinkServiceConnectionStateStatusPtr(v string) PrivateLinkServiceConnectionStateStatusPtrInput {
	return (*privateLinkServiceConnectionStateStatusPtr)(&v)
}

func (*privateLinkServiceConnectionStateStatusPtr) ElementType() reflect.Type {
	return privateLinkServiceConnectionStateStatusPtrType
}

func (in *privateLinkServiceConnectionStateStatusPtr) ToPrivateLinkServiceConnectionStateStatusPtrOutput() PrivateLinkServiceConnectionStateStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateLinkServiceConnectionStateStatusPtrOutput)
}

func (in *privateLinkServiceConnectionStateStatusPtr) ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateLinkServiceConnectionStateStatusPtrOutput)
}

// The status of the last operation.
type ProvisioningState string

const (
	// Created value.
	ProvisioningStateCreated = ProvisioningState("Created")
	// Updated value.
	ProvisioningStateUpdated = ProvisioningState("Updated")
	// Running value.
	ProvisioningStateRunning = ProvisioningState("Running")
	// Completed value.
	ProvisioningStateCompleted = ProvisioningState("Completed")
	// Failed value.
	ProvisioningStateFailed = ProvisioningState("Failed")
	// Succeeded value.
	ProvisioningStateSucceeded = ProvisioningState("Succeeded")
	// Canceled value.
	ProvisioningStateCanceled = ProvisioningState("Canceled")
)

func (ProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (e ProvisioningState) ToProvisioningStateOutput() ProvisioningStateOutput {
	return pulumi.ToOutput(e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return e.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return ProvisioningState(e).ToProvisioningStateOutputWithContext(ctx).ToProvisioningStatePtrOutputWithContext(ctx)
}

func (e ProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProvisioningStateOutput struct{ *pulumi.OutputState }

func (ProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStateOutput) ToProvisioningStateOutput() ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisioningState) *ProvisioningState {
		return &v
	}).(ProvisioningStatePtrOutput)
}

func (o ProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (ProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) Elem() ProvisioningStateOutput {
	return o.ApplyT(func(v *ProvisioningState) ProvisioningState {
		if v != nil {
			return *v
		}
		var ret ProvisioningState
		return ret
	}).(ProvisioningStateOutput)
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ProvisioningStateInput is an input type that accepts values of the ProvisioningState enum
// A concrete instance of `ProvisioningStateInput` can be one of the following:
//
//	ProvisioningStateCreated
//	ProvisioningStateUpdated
//	ProvisioningStateRunning
//	ProvisioningStateCompleted
//	ProvisioningStateFailed
//	ProvisioningStateSucceeded
//	ProvisioningStateCanceled
type ProvisioningStateInput interface {
	pulumi.Input

	ToProvisioningStateOutput() ProvisioningStateOutput
	ToProvisioningStateOutputWithContext(context.Context) ProvisioningStateOutput
}

var provisioningStatePtrType = reflect.TypeOf((**ProvisioningState)(nil)).Elem()

type ProvisioningStatePtrInput interface {
	pulumi.Input

	ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput
	ToProvisioningStatePtrOutputWithContext(context.Context) ProvisioningStatePtrOutput
}

type provisioningStatePtr string

func ProvisioningStatePtr(v string) ProvisioningStatePtrInput {
	return (*provisioningStatePtr)(&v)
}

func (*provisioningStatePtr) ElementType() reflect.Type {
	return provisioningStatePtrType
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(ProvisioningStatePtrOutput)
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProvisioningStatePtrOutput)
}

// Gets or sets the discovery scenario.
type SqlSitePropertiesDiscoveryScenario string

const (
	// Migrate value.
	SqlSitePropertiesDiscoveryScenarioMigrate = SqlSitePropertiesDiscoveryScenario("Migrate")
	// DR value.
	SqlSitePropertiesDiscoveryScenarioDR = SqlSitePropertiesDiscoveryScenario("DR")
)

func (SqlSitePropertiesDiscoveryScenario) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlSitePropertiesDiscoveryScenario)(nil)).Elem()
}

func (e SqlSitePropertiesDiscoveryScenario) ToSqlSitePropertiesDiscoveryScenarioOutput() SqlSitePropertiesDiscoveryScenarioOutput {
	return pulumi.ToOutput(e).(SqlSitePropertiesDiscoveryScenarioOutput)
}

func (e SqlSitePropertiesDiscoveryScenario) ToSqlSitePropertiesDiscoveryScenarioOutputWithContext(ctx context.Context) SqlSitePropertiesDiscoveryScenarioOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SqlSitePropertiesDiscoveryScenarioOutput)
}

func (e SqlSitePropertiesDiscoveryScenario) ToSqlSitePropertiesDiscoveryScenarioPtrOutput() SqlSitePropertiesDiscoveryScenarioPtrOutput {
	return e.ToSqlSitePropertiesDiscoveryScenarioPtrOutputWithContext(context.Background())
}

func (e SqlSitePropertiesDiscoveryScenario) ToSqlSitePropertiesDiscoveryScenarioPtrOutputWithContext(ctx context.Context) SqlSitePropertiesDiscoveryScenarioPtrOutput {
	return SqlSitePropertiesDiscoveryScenario(e).ToSqlSitePropertiesDiscoveryScenarioOutputWithContext(ctx).ToSqlSitePropertiesDiscoveryScenarioPtrOutputWithContext(ctx)
}

func (e SqlSitePropertiesDiscoveryScenario) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlSitePropertiesDiscoveryScenario) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlSitePropertiesDiscoveryScenario) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SqlSitePropertiesDiscoveryScenario) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SqlSitePropertiesDiscoveryScenarioOutput struct{ *pulumi.OutputState }

func (SqlSitePropertiesDiscoveryScenarioOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlSitePropertiesDiscoveryScenario)(nil)).Elem()
}

func (o SqlSitePropertiesDiscoveryScenarioOutput) ToSqlSitePropertiesDiscoveryScenarioOutput() SqlSitePropertiesDiscoveryScenarioOutput {
	return o
}

func (o SqlSitePropertiesDiscoveryScenarioOutput) ToSqlSitePropertiesDiscoveryScenarioOutputWithContext(ctx context.Context) SqlSitePropertiesDiscoveryScenarioOutput {
	return o
}

func (o SqlSitePropertiesDiscoveryScenarioOutput) ToSqlSitePropertiesDiscoveryScenarioPtrOutput() SqlSitePropertiesDiscoveryScenarioPtrOutput {
	return o.ToSqlSitePropertiesDiscoveryScenarioPtrOutputWithContext(context.Background())
}

func (o SqlSitePropertiesDiscoveryScenarioOutput) ToSqlSitePropertiesDiscoveryScenarioPtrOutputWithContext(ctx context.Context) SqlSitePropertiesDiscoveryScenarioPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlSitePropertiesDiscoveryScenario) *SqlSitePropertiesDiscoveryScenario {
		return &v
	}).(SqlSitePropertiesDiscoveryScenarioPtrOutput)
}

func (o SqlSitePropertiesDiscoveryScenarioOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SqlSitePropertiesDiscoveryScenarioOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlSitePropertiesDiscoveryScenario) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SqlSitePropertiesDiscoveryScenarioOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlSitePropertiesDiscoveryScenarioOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlSitePropertiesDiscoveryScenario) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SqlSitePropertiesDiscoveryScenarioPtrOutput struct{ *pulumi.OutputState }

func (SqlSitePropertiesDiscoveryScenarioPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlSitePropertiesDiscoveryScenario)(nil)).Elem()
}

func (o SqlSitePropertiesDiscoveryScenarioPtrOutput) ToSqlSitePropertiesDiscoveryScenarioPtrOutput() SqlSitePropertiesDiscoveryScenarioPtrOutput {
	return o
}

func (o SqlSitePropertiesDiscoveryScenarioPtrOutput) ToSqlSitePropertiesDiscoveryScenarioPtrOutputWithContext(ctx context.Context) SqlSitePropertiesDiscoveryScenarioPtrOutput {
	return o
}

func (o SqlSitePropertiesDiscoveryScenarioPtrOutput) Elem() SqlSitePropertiesDiscoveryScenarioOutput {
	return o.ApplyT(func(v *SqlSitePropertiesDiscoveryScenario) SqlSitePropertiesDiscoveryScenario {
		if v != nil {
			return *v
		}
		var ret SqlSitePropertiesDiscoveryScenario
		return ret
	}).(SqlSitePropertiesDiscoveryScenarioOutput)
}

func (o SqlSitePropertiesDiscoveryScenarioPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlSitePropertiesDiscoveryScenarioPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SqlSitePropertiesDiscoveryScenario) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SqlSitePropertiesDiscoveryScenarioInput is an input type that accepts values of the SqlSitePropertiesDiscoveryScenario enum
// A concrete instance of `SqlSitePropertiesDiscoveryScenarioInput` can be one of the following:
//
//	SqlSitePropertiesDiscoveryScenarioMigrate
//	SqlSitePropertiesDiscoveryScenarioDR
type SqlSitePropertiesDiscoveryScenarioInput interface {
	pulumi.Input

	ToSqlSitePropertiesDiscoveryScenarioOutput() SqlSitePropertiesDiscoveryScenarioOutput
	ToSqlSitePropertiesDiscoveryScenarioOutputWithContext(context.Context) SqlSitePropertiesDiscoveryScenarioOutput
}

var sqlSitePropertiesDiscoveryScenarioPtrType = reflect.TypeOf((**SqlSitePropertiesDiscoveryScenario)(nil)).Elem()

type SqlSitePropertiesDiscoveryScenarioPtrInput interface {
	pulumi.Input

	ToSqlSitePropertiesDiscoveryScenarioPtrOutput() SqlSitePropertiesDiscoveryScenarioPtrOutput
	ToSqlSitePropertiesDiscoveryScenarioPtrOutputWithContext(context.Context) SqlSitePropertiesDiscoveryScenarioPtrOutput
}

type sqlSitePropertiesDiscoveryScenarioPtr string

func SqlSitePropertiesDiscoveryScenarioPtr(v string) SqlSitePropertiesDiscoveryScenarioPtrInput {
	return (*sqlSitePropertiesDiscoveryScenarioPtr)(&v)
}

func (*sqlSitePropertiesDiscoveryScenarioPtr) ElementType() reflect.Type {
	return sqlSitePropertiesDiscoveryScenarioPtrType
}

func (in *sqlSitePropertiesDiscoveryScenarioPtr) ToSqlSitePropertiesDiscoveryScenarioPtrOutput() SqlSitePropertiesDiscoveryScenarioPtrOutput {
	return pulumi.ToOutput(in).(SqlSitePropertiesDiscoveryScenarioPtrOutput)
}

func (in *sqlSitePropertiesDiscoveryScenarioPtr) ToSqlSitePropertiesDiscoveryScenarioPtrOutputWithContext(ctx context.Context) SqlSitePropertiesDiscoveryScenarioPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SqlSitePropertiesDiscoveryScenarioPtrOutput)
}

// Gets or sets the discovery scenario.
type WebAppSitePropertiesDiscoveryScenario string

const (
	// Migrate value.
	WebAppSitePropertiesDiscoveryScenarioMigrate = WebAppSitePropertiesDiscoveryScenario("Migrate")
	// DR value.
	WebAppSitePropertiesDiscoveryScenarioDR = WebAppSitePropertiesDiscoveryScenario("DR")
)

func (WebAppSitePropertiesDiscoveryScenario) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppSitePropertiesDiscoveryScenario)(nil)).Elem()
}

func (e WebAppSitePropertiesDiscoveryScenario) ToWebAppSitePropertiesDiscoveryScenarioOutput() WebAppSitePropertiesDiscoveryScenarioOutput {
	return pulumi.ToOutput(e).(WebAppSitePropertiesDiscoveryScenarioOutput)
}

func (e WebAppSitePropertiesDiscoveryScenario) ToWebAppSitePropertiesDiscoveryScenarioOutputWithContext(ctx context.Context) WebAppSitePropertiesDiscoveryScenarioOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebAppSitePropertiesDiscoveryScenarioOutput)
}

func (e WebAppSitePropertiesDiscoveryScenario) ToWebAppSitePropertiesDiscoveryScenarioPtrOutput() WebAppSitePropertiesDiscoveryScenarioPtrOutput {
	return e.ToWebAppSitePropertiesDiscoveryScenarioPtrOutputWithContext(context.Background())
}

func (e WebAppSitePropertiesDiscoveryScenario) ToWebAppSitePropertiesDiscoveryScenarioPtrOutputWithContext(ctx context.Context) WebAppSitePropertiesDiscoveryScenarioPtrOutput {
	return WebAppSitePropertiesDiscoveryScenario(e).ToWebAppSitePropertiesDiscoveryScenarioOutputWithContext(ctx).ToWebAppSitePropertiesDiscoveryScenarioPtrOutputWithContext(ctx)
}

func (e WebAppSitePropertiesDiscoveryScenario) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebAppSitePropertiesDiscoveryScenario) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebAppSitePropertiesDiscoveryScenario) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebAppSitePropertiesDiscoveryScenario) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebAppSitePropertiesDiscoveryScenarioOutput struct{ *pulumi.OutputState }

func (WebAppSitePropertiesDiscoveryScenarioOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppSitePropertiesDiscoveryScenario)(nil)).Elem()
}

func (o WebAppSitePropertiesDiscoveryScenarioOutput) ToWebAppSitePropertiesDiscoveryScenarioOutput() WebAppSitePropertiesDiscoveryScenarioOutput {
	return o
}

func (o WebAppSitePropertiesDiscoveryScenarioOutput) ToWebAppSitePropertiesDiscoveryScenarioOutputWithContext(ctx context.Context) WebAppSitePropertiesDiscoveryScenarioOutput {
	return o
}

func (o WebAppSitePropertiesDiscoveryScenarioOutput) ToWebAppSitePropertiesDiscoveryScenarioPtrOutput() WebAppSitePropertiesDiscoveryScenarioPtrOutput {
	return o.ToWebAppSitePropertiesDiscoveryScenarioPtrOutputWithContext(context.Background())
}

func (o WebAppSitePropertiesDiscoveryScenarioOutput) ToWebAppSitePropertiesDiscoveryScenarioPtrOutputWithContext(ctx context.Context) WebAppSitePropertiesDiscoveryScenarioPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebAppSitePropertiesDiscoveryScenario) *WebAppSitePropertiesDiscoveryScenario {
		return &v
	}).(WebAppSitePropertiesDiscoveryScenarioPtrOutput)
}

func (o WebAppSitePropertiesDiscoveryScenarioOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebAppSitePropertiesDiscoveryScenarioOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebAppSitePropertiesDiscoveryScenario) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebAppSitePropertiesDiscoveryScenarioOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebAppSitePropertiesDiscoveryScenarioOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebAppSitePropertiesDiscoveryScenario) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebAppSitePropertiesDiscoveryScenarioPtrOutput struct{ *pulumi.OutputState }

func (WebAppSitePropertiesDiscoveryScenarioPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSitePropertiesDiscoveryScenario)(nil)).Elem()
}

func (o WebAppSitePropertiesDiscoveryScenarioPtrOutput) ToWebAppSitePropertiesDiscoveryScenarioPtrOutput() WebAppSitePropertiesDiscoveryScenarioPtrOutput {
	return o
}

func (o WebAppSitePropertiesDiscoveryScenarioPtrOutput) ToWebAppSitePropertiesDiscoveryScenarioPtrOutputWithContext(ctx context.Context) WebAppSitePropertiesDiscoveryScenarioPtrOutput {
	return o
}

func (o WebAppSitePropertiesDiscoveryScenarioPtrOutput) Elem() WebAppSitePropertiesDiscoveryScenarioOutput {
	return o.ApplyT(func(v *WebAppSitePropertiesDiscoveryScenario) WebAppSitePropertiesDiscoveryScenario {
		if v != nil {
			return *v
		}
		var ret WebAppSitePropertiesDiscoveryScenario
		return ret
	}).(WebAppSitePropertiesDiscoveryScenarioOutput)
}

func (o WebAppSitePropertiesDiscoveryScenarioPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebAppSitePropertiesDiscoveryScenarioPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebAppSitePropertiesDiscoveryScenario) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// WebAppSitePropertiesDiscoveryScenarioInput is an input type that accepts values of the WebAppSitePropertiesDiscoveryScenario enum
// A concrete instance of `WebAppSitePropertiesDiscoveryScenarioInput` can be one of the following:
//
//	WebAppSitePropertiesDiscoveryScenarioMigrate
//	WebAppSitePropertiesDiscoveryScenarioDR
type WebAppSitePropertiesDiscoveryScenarioInput interface {
	pulumi.Input

	ToWebAppSitePropertiesDiscoveryScenarioOutput() WebAppSitePropertiesDiscoveryScenarioOutput
	ToWebAppSitePropertiesDiscoveryScenarioOutputWithContext(context.Context) WebAppSitePropertiesDiscoveryScenarioOutput
}

var webAppSitePropertiesDiscoveryScenarioPtrType = reflect.TypeOf((**WebAppSitePropertiesDiscoveryScenario)(nil)).Elem()

type WebAppSitePropertiesDiscoveryScenarioPtrInput interface {
	pulumi.Input

	ToWebAppSitePropertiesDiscoveryScenarioPtrOutput() WebAppSitePropertiesDiscoveryScenarioPtrOutput
	ToWebAppSitePropertiesDiscoveryScenarioPtrOutputWithContext(context.Context) WebAppSitePropertiesDiscoveryScenarioPtrOutput
}

type webAppSitePropertiesDiscoveryScenarioPtr string

func WebAppSitePropertiesDiscoveryScenarioPtr(v string) WebAppSitePropertiesDiscoveryScenarioPtrInput {
	return (*webAppSitePropertiesDiscoveryScenarioPtr)(&v)
}

func (*webAppSitePropertiesDiscoveryScenarioPtr) ElementType() reflect.Type {
	return webAppSitePropertiesDiscoveryScenarioPtrType
}

func (in *webAppSitePropertiesDiscoveryScenarioPtr) ToWebAppSitePropertiesDiscoveryScenarioPtrOutput() WebAppSitePropertiesDiscoveryScenarioPtrOutput {
	return pulumi.ToOutput(in).(WebAppSitePropertiesDiscoveryScenarioPtrOutput)
}

func (in *webAppSitePropertiesDiscoveryScenarioPtr) ToWebAppSitePropertiesDiscoveryScenarioPtrOutputWithContext(ctx context.Context) WebAppSitePropertiesDiscoveryScenarioPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebAppSitePropertiesDiscoveryScenarioPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MasterSitePropertiesPublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(MasterSitePropertiesPublicNetworkAccessPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateStatusOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateStatusPtrOutput{})
	pulumi.RegisterOutputType(ProvisioningStateOutput{})
	pulumi.RegisterOutputType(ProvisioningStatePtrOutput{})
	pulumi.RegisterOutputType(SqlSitePropertiesDiscoveryScenarioOutput{})
	pulumi.RegisterOutputType(SqlSitePropertiesDiscoveryScenarioPtrOutput{})
	pulumi.RegisterOutputType(WebAppSitePropertiesDiscoveryScenarioOutput{})
	pulumi.RegisterOutputType(WebAppSitePropertiesDiscoveryScenarioPtrOutput{})
}
