// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Whether to enable Azure front door
type AzureFrontDoorEnabled string

const (
	AzureFrontDoorEnabledEnabled  = AzureFrontDoorEnabled("Enabled")
	AzureFrontDoorEnabledDisabled = AzureFrontDoorEnabled("Disabled")
)

// Tier of the server SKU
type DatabaseTier string

const (
	DatabaseTierBurstable       = DatabaseTier("Burstable")
	DatabaseTierGeneralPurpose  = DatabaseTier("GeneralPurpose")
	DatabaseTierMemoryOptimized = DatabaseTier("MemoryOptimized")
)

func (DatabaseTier) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseTier)(nil)).Elem()
}

func (e DatabaseTier) ToDatabaseTierOutput() DatabaseTierOutput {
	return pulumi.ToOutput(e).(DatabaseTierOutput)
}

func (e DatabaseTier) ToDatabaseTierOutputWithContext(ctx context.Context) DatabaseTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatabaseTierOutput)
}

func (e DatabaseTier) ToDatabaseTierPtrOutput() DatabaseTierPtrOutput {
	return e.ToDatabaseTierPtrOutputWithContext(context.Background())
}

func (e DatabaseTier) ToDatabaseTierPtrOutputWithContext(ctx context.Context) DatabaseTierPtrOutput {
	return DatabaseTier(e).ToDatabaseTierOutputWithContext(ctx).ToDatabaseTierPtrOutputWithContext(ctx)
}

func (e DatabaseTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatabaseTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatabaseTierOutput struct{ *pulumi.OutputState }

func (DatabaseTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseTier)(nil)).Elem()
}

func (o DatabaseTierOutput) ToDatabaseTierOutput() DatabaseTierOutput {
	return o
}

func (o DatabaseTierOutput) ToDatabaseTierOutputWithContext(ctx context.Context) DatabaseTierOutput {
	return o
}

func (o DatabaseTierOutput) ToDatabaseTierPtrOutput() DatabaseTierPtrOutput {
	return o.ToDatabaseTierPtrOutputWithContext(context.Background())
}

func (o DatabaseTierOutput) ToDatabaseTierPtrOutputWithContext(ctx context.Context) DatabaseTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseTier) *DatabaseTier {
		return &v
	}).(DatabaseTierPtrOutput)
}

func (o DatabaseTierOutput) ToOutput(ctx context.Context) pulumix.Output[DatabaseTier] {
	return pulumix.Output[DatabaseTier]{
		OutputState: o.OutputState,
	}
}

func (o DatabaseTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatabaseTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatabaseTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatabaseTierPtrOutput struct{ *pulumi.OutputState }

func (DatabaseTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseTier)(nil)).Elem()
}

func (o DatabaseTierPtrOutput) ToDatabaseTierPtrOutput() DatabaseTierPtrOutput {
	return o
}

func (o DatabaseTierPtrOutput) ToDatabaseTierPtrOutputWithContext(ctx context.Context) DatabaseTierPtrOutput {
	return o
}

func (o DatabaseTierPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*DatabaseTier] {
	return pulumix.Output[*DatabaseTier]{
		OutputState: o.OutputState,
	}
}

func (o DatabaseTierPtrOutput) Elem() DatabaseTierOutput {
	return o.ApplyT(func(v *DatabaseTier) DatabaseTier {
		if v != nil {
			return *v
		}
		var ret DatabaseTier
		return ret
	}).(DatabaseTierOutput)
}

func (o DatabaseTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatabaseTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DatabaseTierInput is an input type that accepts DatabaseTierArgs and DatabaseTierOutput values.
// You can construct a concrete instance of `DatabaseTierInput` via:
//
//	DatabaseTierArgs{...}
type DatabaseTierInput interface {
	pulumi.Input

	ToDatabaseTierOutput() DatabaseTierOutput
	ToDatabaseTierOutputWithContext(context.Context) DatabaseTierOutput
}

var databaseTierPtrType = reflect.TypeOf((**DatabaseTier)(nil)).Elem()

type DatabaseTierPtrInput interface {
	pulumi.Input

	ToDatabaseTierPtrOutput() DatabaseTierPtrOutput
	ToDatabaseTierPtrOutputWithContext(context.Context) DatabaseTierPtrOutput
}

type databaseTierPtr string

func DatabaseTierPtr(v string) DatabaseTierPtrInput {
	return (*databaseTierPtr)(&v)
}

func (*databaseTierPtr) ElementType() reflect.Type {
	return databaseTierPtrType
}

func (in *databaseTierPtr) ToDatabaseTierPtrOutput() DatabaseTierPtrOutput {
	return pulumi.ToOutput(in).(DatabaseTierPtrOutput)
}

func (in *databaseTierPtr) ToDatabaseTierPtrOutputWithContext(ctx context.Context) DatabaseTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatabaseTierPtrOutput)
}

func (in *databaseTierPtr) ToOutput(ctx context.Context) pulumix.Output[*DatabaseTier] {
	return pulumix.Output[*DatabaseTier]{
		OutputState: in.ToDatabaseTierPtrOutputWithContext(ctx).OutputState,
	}
}

// Database type
type DatabaseType string

const (
	DatabaseTypeMySql = DatabaseType("MySql")
)

// Storage type
type DiskStorageType string

const (
	DiskStorageType_Premium_LRS     = DiskStorageType("Premium_LRS")
	DiskStorageType_Standard_LRS    = DiskStorageType("Standard_LRS")
	DiskStorageType_StandardSSD_LRS = DiskStorageType("StandardSSD_LRS")
)

func (DiskStorageType) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskStorageType)(nil)).Elem()
}

func (e DiskStorageType) ToDiskStorageTypeOutput() DiskStorageTypeOutput {
	return pulumi.ToOutput(e).(DiskStorageTypeOutput)
}

func (e DiskStorageType) ToDiskStorageTypeOutputWithContext(ctx context.Context) DiskStorageTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiskStorageTypeOutput)
}

func (e DiskStorageType) ToDiskStorageTypePtrOutput() DiskStorageTypePtrOutput {
	return e.ToDiskStorageTypePtrOutputWithContext(context.Background())
}

func (e DiskStorageType) ToDiskStorageTypePtrOutputWithContext(ctx context.Context) DiskStorageTypePtrOutput {
	return DiskStorageType(e).ToDiskStorageTypeOutputWithContext(ctx).ToDiskStorageTypePtrOutputWithContext(ctx)
}

func (e DiskStorageType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskStorageType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskStorageType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiskStorageType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiskStorageTypeOutput struct{ *pulumi.OutputState }

func (DiskStorageTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskStorageType)(nil)).Elem()
}

func (o DiskStorageTypeOutput) ToDiskStorageTypeOutput() DiskStorageTypeOutput {
	return o
}

func (o DiskStorageTypeOutput) ToDiskStorageTypeOutputWithContext(ctx context.Context) DiskStorageTypeOutput {
	return o
}

func (o DiskStorageTypeOutput) ToDiskStorageTypePtrOutput() DiskStorageTypePtrOutput {
	return o.ToDiskStorageTypePtrOutputWithContext(context.Background())
}

func (o DiskStorageTypeOutput) ToDiskStorageTypePtrOutputWithContext(ctx context.Context) DiskStorageTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskStorageType) *DiskStorageType {
		return &v
	}).(DiskStorageTypePtrOutput)
}

func (o DiskStorageTypeOutput) ToOutput(ctx context.Context) pulumix.Output[DiskStorageType] {
	return pulumix.Output[DiskStorageType]{
		OutputState: o.OutputState,
	}
}

func (o DiskStorageTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiskStorageTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskStorageType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiskStorageTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskStorageTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskStorageType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiskStorageTypePtrOutput struct{ *pulumi.OutputState }

func (DiskStorageTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskStorageType)(nil)).Elem()
}

func (o DiskStorageTypePtrOutput) ToDiskStorageTypePtrOutput() DiskStorageTypePtrOutput {
	return o
}

func (o DiskStorageTypePtrOutput) ToDiskStorageTypePtrOutputWithContext(ctx context.Context) DiskStorageTypePtrOutput {
	return o
}

func (o DiskStorageTypePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*DiskStorageType] {
	return pulumix.Output[*DiskStorageType]{
		OutputState: o.OutputState,
	}
}

func (o DiskStorageTypePtrOutput) Elem() DiskStorageTypeOutput {
	return o.ApplyT(func(v *DiskStorageType) DiskStorageType {
		if v != nil {
			return *v
		}
		var ret DiskStorageType
		return ret
	}).(DiskStorageTypeOutput)
}

func (o DiskStorageTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskStorageTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiskStorageType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DiskStorageTypeInput is an input type that accepts DiskStorageTypeArgs and DiskStorageTypeOutput values.
// You can construct a concrete instance of `DiskStorageTypeInput` via:
//
//	DiskStorageTypeArgs{...}
type DiskStorageTypeInput interface {
	pulumi.Input

	ToDiskStorageTypeOutput() DiskStorageTypeOutput
	ToDiskStorageTypeOutputWithContext(context.Context) DiskStorageTypeOutput
}

var diskStorageTypePtrType = reflect.TypeOf((**DiskStorageType)(nil)).Elem()

type DiskStorageTypePtrInput interface {
	pulumi.Input

	ToDiskStorageTypePtrOutput() DiskStorageTypePtrOutput
	ToDiskStorageTypePtrOutputWithContext(context.Context) DiskStorageTypePtrOutput
}

type diskStorageTypePtr string

func DiskStorageTypePtr(v string) DiskStorageTypePtrInput {
	return (*diskStorageTypePtr)(&v)
}

func (*diskStorageTypePtr) ElementType() reflect.Type {
	return diskStorageTypePtrType
}

func (in *diskStorageTypePtr) ToDiskStorageTypePtrOutput() DiskStorageTypePtrOutput {
	return pulumi.ToOutput(in).(DiskStorageTypePtrOutput)
}

func (in *diskStorageTypePtr) ToDiskStorageTypePtrOutputWithContext(ctx context.Context) DiskStorageTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiskStorageTypePtrOutput)
}

func (in *diskStorageTypePtr) ToOutput(ctx context.Context) pulumix.Output[*DiskStorageType] {
	return pulumix.Output[*DiskStorageType]{
		OutputState: in.ToDiskStorageTypePtrOutputWithContext(ctx).OutputState,
	}
}

// Whether to enable Azure backup for the workload
type EnableBackup string

const (
	EnableBackupEnabled  = EnableBackup("Enabled")
	EnableBackupDisabled = EnableBackup("Disabled")
)

// Whether to enable SSL enforcement on the database
type EnableSslEnforcement string

const (
	EnableSslEnforcementEnabled  = EnableSslEnforcement("Enabled")
	EnableSslEnforcementDisabled = EnableSslEnforcement("Disabled")
)

// File share backing storage type
type FileShareStorageType string

const (
	FileShareStorageType_Standard_LRS = FileShareStorageType("Standard_LRS")
	FileShareStorageType_Standard_GRS = FileShareStorageType("Standard_GRS")
	FileShareStorageType_Standard_ZRS = FileShareStorageType("Standard_ZRS")
	FileShareStorageType_Premium_LRS  = FileShareStorageType("Premium_LRS")
)

// Share type
type FileShareType string

const (
	FileShareTypeNfsOnController = FileShareType("NfsOnController")
	FileShareTypeAzureFiles      = FileShareType("AzureFiles")
)

// Whether to enable HA for the server
type HAEnabled string

const (
	HAEnabledEnabled  = HAEnabled("Enabled")
	HAEnabledDisabled = HAEnabled("Disabled")
)

// Load balancer type
type LoadBalancerType string

const (
	LoadBalancerTypeApplicationGateway = LoadBalancerType("ApplicationGateway")
	LoadBalancerTypeLoadBalancer       = LoadBalancerType("LoadBalancer")
)

// Type of manage identity
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone         = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeUserAssigned = ManagedServiceIdentityType("UserAssigned")
)

// OS image offer
type OSImageOffer string

const (
	OSImageOfferUbuntuServer = OSImageOffer("UbuntuServer")
)

// OS image publisher
type OSImagePublisher string

const (
	OSImagePublisherCanonical = OSImagePublisher("Canonical")
)

// OS image sku
type OSImageSku string

const (
	OSImageSku_18_04_LTS = OSImageSku("18.04-LTS")
	OSImageSku_16_04_LTS = OSImageSku("16.04-LTS")
)

// OS image version
type OSImageVersion string

const (
	OSImageVersionLatest = OSImageVersion("latest")
)

// PHP version
type PHPVersion string

const (
	PHPVersion_7_2 = PHPVersion("7.2")
	PHPVersion_7_3 = PHPVersion("7.3")
	PHPVersion_7_4 = PHPVersion("7.4")
)

// Cache family
type RedisCacheFamily string

const (
	RedisCacheFamilyC = RedisCacheFamily("C")
	RedisCacheFamilyP = RedisCacheFamily("P")
)

// The DB scale method.
type SAPDatabaseScaleMethod string

const (
	SAPDatabaseScaleMethodScaleUp = SAPDatabaseScaleMethod("ScaleUp")
)

// The database type. Eg: HANA, DB2, etc
type SAPDatabaseType string

const (
	SAPDatabaseTypeHANA = SAPDatabaseType("HANA")
	SAPDatabaseTypeDB2  = SAPDatabaseType("DB2")
)

// The deployment type. Eg: SingleServer/ThreeTier
type SAPDeploymentType string

const (
	SAPDeploymentTypeSingleServer = SAPDeploymentType("SingleServer")
	SAPDeploymentTypeThreeTier    = SAPDeploymentType("ThreeTier")
)

// Defines the environment type - Production/Non Production.
type SAPEnvironmentType string

const (
	SAPEnvironmentTypeNonProd = SAPEnvironmentType("NonProd")
	SAPEnvironmentTypeProd    = SAPEnvironmentType("Prod")
)

// The high availability type.
type SAPHighAvailabilityType string

const (
	SAPHighAvailabilityTypeAvailabilitySet  = SAPHighAvailabilityType("AvailabilitySet")
	SAPHighAvailabilityTypeAvailabilityZone = SAPHighAvailabilityType("AvailabilityZone")
)

// Defines the SAP Product type.
type SAPProductType string

const (
	SAPProductTypeECC    = SAPProductType("ECC")
	SAPProductTypeS4HANA = SAPProductType("S4HANA")
	SAPProductTypeOther  = SAPProductType("Other")
)

// Search type
type SearchType string

const (
	SearchTypeElastic = SearchType("Elastic")
)

// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
type SkuTier string

const (
	SkuTierFree     = SkuTier("Free")
	SkuTierBasic    = SkuTier("Basic")
	SkuTierStandard = SkuTier("Standard")
	SkuTierPremium  = SkuTier("Premium")
)

func (SkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (e SkuTier) ToSkuTierOutput() SkuTierOutput {
	return pulumi.ToOutput(e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return e.ToSkuTierPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return SkuTier(e).ToSkuTierOutputWithContext(ctx).ToSkuTierPtrOutputWithContext(ctx)
}

func (e SkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuTierOutput struct{ *pulumi.OutputState }

func (SkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (o SkuTierOutput) ToSkuTierOutput() SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o.ToSkuTierPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuTier) *SkuTier {
		return &v
	}).(SkuTierPtrOutput)
}

func (o SkuTierOutput) ToOutput(ctx context.Context) pulumix.Output[SkuTier] {
	return pulumix.Output[SkuTier]{
		OutputState: o.OutputState,
	}
}

func (o SkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuTierPtrOutput struct{ *pulumi.OutputState }

func (SkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuTier)(nil)).Elem()
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*SkuTier] {
	return pulumix.Output[*SkuTier]{
		OutputState: o.OutputState,
	}
}

func (o SkuTierPtrOutput) Elem() SkuTierOutput {
	return o.ApplyT(func(v *SkuTier) SkuTier {
		if v != nil {
			return *v
		}
		var ret SkuTier
		return ret
	}).(SkuTierOutput)
}

func (o SkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SkuTierInput is an input type that accepts SkuTierArgs and SkuTierOutput values.
// You can construct a concrete instance of `SkuTierInput` via:
//
//	SkuTierArgs{...}
type SkuTierInput interface {
	pulumi.Input

	ToSkuTierOutput() SkuTierOutput
	ToSkuTierOutputWithContext(context.Context) SkuTierOutput
}

var skuTierPtrType = reflect.TypeOf((**SkuTier)(nil)).Elem()

type SkuTierPtrInput interface {
	pulumi.Input

	ToSkuTierPtrOutput() SkuTierPtrOutput
	ToSkuTierPtrOutputWithContext(context.Context) SkuTierPtrOutput
}

type skuTierPtr string

func SkuTierPtr(v string) SkuTierPtrInput {
	return (*skuTierPtr)(&v)
}

func (*skuTierPtr) ElementType() reflect.Type {
	return skuTierPtrType
}

func (in *skuTierPtr) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return pulumi.ToOutput(in).(SkuTierPtrOutput)
}

func (in *skuTierPtr) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuTierPtrOutput)
}

func (in *skuTierPtr) ToOutput(ctx context.Context) pulumix.Output[*SkuTier] {
	return pulumix.Output[*SkuTier]{
		OutputState: in.ToSkuTierPtrOutputWithContext(ctx).OutputState,
	}
}

// Application version
type WordpressVersions string

const (
	WordpressVersions_5_4_3 = WordpressVersions("5.4.3")
	WordpressVersions_5_4_2 = WordpressVersions("5.4.2")
	WordpressVersions_5_4_1 = WordpressVersions("5.4.1")
	WordpressVersions_5_4   = WordpressVersions("5.4")
)

// Indicates which kind of php workload this resource represent e.g WordPress
type WorkloadKind string

const (
	WorkloadKindWordPress = WorkloadKind("WordPress")
)

func init() {
	pulumi.RegisterOutputType(DatabaseTierOutput{})
	pulumi.RegisterOutputType(DatabaseTierPtrOutput{})
	pulumi.RegisterOutputType(DiskStorageTypeOutput{})
	pulumi.RegisterOutputType(DiskStorageTypePtrOutput{})
	pulumi.RegisterOutputType(SkuTierOutput{})
	pulumi.RegisterOutputType(SkuTierPtrOutput{})
}
