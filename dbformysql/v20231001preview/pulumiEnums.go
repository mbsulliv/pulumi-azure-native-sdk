// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The mode to create a new MySQL server.
type CreateMode string

const (
	CreateModeDefault            = CreateMode("Default")
	CreateModePointInTimeRestore = CreateMode("PointInTimeRestore")
	CreateModeReplica            = CreateMode("Replica")
	CreateModeGeoRestore         = CreateMode("GeoRestore")
)

// The key type, AzureKeyVault for enable cmk, SystemManaged for disable cmk.
type DataEncryptionType string

const (
	DataEncryptionTypeAzureKeyVault = DataEncryptionType("AzureKeyVault")
	DataEncryptionTypeSystemManaged = DataEncryptionType("SystemManaged")
)

func (DataEncryptionType) ElementType() reflect.Type {
	return reflect.TypeOf((*DataEncryptionType)(nil)).Elem()
}

func (e DataEncryptionType) ToDataEncryptionTypeOutput() DataEncryptionTypeOutput {
	return pulumi.ToOutput(e).(DataEncryptionTypeOutput)
}

func (e DataEncryptionType) ToDataEncryptionTypeOutputWithContext(ctx context.Context) DataEncryptionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataEncryptionTypeOutput)
}

func (e DataEncryptionType) ToDataEncryptionTypePtrOutput() DataEncryptionTypePtrOutput {
	return e.ToDataEncryptionTypePtrOutputWithContext(context.Background())
}

func (e DataEncryptionType) ToDataEncryptionTypePtrOutputWithContext(ctx context.Context) DataEncryptionTypePtrOutput {
	return DataEncryptionType(e).ToDataEncryptionTypeOutputWithContext(ctx).ToDataEncryptionTypePtrOutputWithContext(ctx)
}

func (e DataEncryptionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataEncryptionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataEncryptionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataEncryptionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataEncryptionTypeOutput struct{ *pulumi.OutputState }

func (DataEncryptionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataEncryptionType)(nil)).Elem()
}

func (o DataEncryptionTypeOutput) ToDataEncryptionTypeOutput() DataEncryptionTypeOutput {
	return o
}

func (o DataEncryptionTypeOutput) ToDataEncryptionTypeOutputWithContext(ctx context.Context) DataEncryptionTypeOutput {
	return o
}

func (o DataEncryptionTypeOutput) ToDataEncryptionTypePtrOutput() DataEncryptionTypePtrOutput {
	return o.ToDataEncryptionTypePtrOutputWithContext(context.Background())
}

func (o DataEncryptionTypeOutput) ToDataEncryptionTypePtrOutputWithContext(ctx context.Context) DataEncryptionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataEncryptionType) *DataEncryptionType {
		return &v
	}).(DataEncryptionTypePtrOutput)
}

func (o DataEncryptionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataEncryptionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataEncryptionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataEncryptionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataEncryptionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataEncryptionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataEncryptionTypePtrOutput struct{ *pulumi.OutputState }

func (DataEncryptionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataEncryptionType)(nil)).Elem()
}

func (o DataEncryptionTypePtrOutput) ToDataEncryptionTypePtrOutput() DataEncryptionTypePtrOutput {
	return o
}

func (o DataEncryptionTypePtrOutput) ToDataEncryptionTypePtrOutputWithContext(ctx context.Context) DataEncryptionTypePtrOutput {
	return o
}

func (o DataEncryptionTypePtrOutput) Elem() DataEncryptionTypeOutput {
	return o.ApplyT(func(v *DataEncryptionType) DataEncryptionType {
		if v != nil {
			return *v
		}
		var ret DataEncryptionType
		return ret
	}).(DataEncryptionTypeOutput)
}

func (o DataEncryptionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataEncryptionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataEncryptionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DataEncryptionTypeInput is an input type that accepts DataEncryptionTypeArgs and DataEncryptionTypeOutput values.
// You can construct a concrete instance of `DataEncryptionTypeInput` via:
//
//	DataEncryptionTypeArgs{...}
type DataEncryptionTypeInput interface {
	pulumi.Input

	ToDataEncryptionTypeOutput() DataEncryptionTypeOutput
	ToDataEncryptionTypeOutputWithContext(context.Context) DataEncryptionTypeOutput
}

var dataEncryptionTypePtrType = reflect.TypeOf((**DataEncryptionType)(nil)).Elem()

type DataEncryptionTypePtrInput interface {
	pulumi.Input

	ToDataEncryptionTypePtrOutput() DataEncryptionTypePtrOutput
	ToDataEncryptionTypePtrOutputWithContext(context.Context) DataEncryptionTypePtrOutput
}

type dataEncryptionTypePtr string

func DataEncryptionTypePtr(v string) DataEncryptionTypePtrInput {
	return (*dataEncryptionTypePtr)(&v)
}

func (*dataEncryptionTypePtr) ElementType() reflect.Type {
	return dataEncryptionTypePtrType
}

func (in *dataEncryptionTypePtr) ToDataEncryptionTypePtrOutput() DataEncryptionTypePtrOutput {
	return pulumi.ToOutput(in).(DataEncryptionTypePtrOutput)
}

func (in *dataEncryptionTypePtr) ToDataEncryptionTypePtrOutputWithContext(ctx context.Context) DataEncryptionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataEncryptionTypePtrOutput)
}

func (in *dataEncryptionTypePtr) ToOutput(ctx context.Context) pulumix.Output[*DataEncryptionType] {
	return pulumix.Output[*DataEncryptionType]{
		OutputState: in.ToDataEncryptionTypePtrOutputWithContext(ctx).OutputState,
	}
}

// Enable Log On Disk or not.
type EnableStatusEnum string

const (
	EnableStatusEnumEnabled  = EnableStatusEnum("Enabled")
	EnableStatusEnumDisabled = EnableStatusEnum("Disabled")
)

// High availability mode for a server.
type HighAvailabilityMode string

const (
	HighAvailabilityModeDisabled      = HighAvailabilityMode("Disabled")
	HighAvailabilityModeZoneRedundant = HighAvailabilityMode("ZoneRedundant")
	HighAvailabilityModeSameZone      = HighAvailabilityMode("SameZone")
)

// Storage type of import source.
type ImportSourceStorageType string

const (
	ImportSourceStorageTypeAzureBlob = ImportSourceStorageType("AzureBlob")
)

// Type of managed service identity.
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeUserAssigned = ManagedServiceIdentityType("UserAssigned")
)

// The replication role.
type ReplicationRole string

const (
	ReplicationRoleNone    = ReplicationRole("None")
	ReplicationRoleSource  = ReplicationRole("Source")
	ReplicationRoleReplica = ReplicationRole("Replica")
)

// The tier of the particular SKU, e.g. GeneralPurpose.
type ServerSkuTier string

const (
	ServerSkuTierBurstable       = ServerSkuTier("Burstable")
	ServerSkuTierGeneralPurpose  = ServerSkuTier("GeneralPurpose")
	ServerSkuTierMemoryOptimized = ServerSkuTier("MemoryOptimized")
)

// Server version.
type ServerVersion string

const (
	ServerVersion_5_7    = ServerVersion("5.7")
	ServerVersion_8_0_21 = ServerVersion("8.0.21")
)

func init() {
	pulumi.RegisterOutputType(DataEncryptionTypeOutput{})
	pulumi.RegisterOutputType(DataEncryptionTypePtrOutput{})
}