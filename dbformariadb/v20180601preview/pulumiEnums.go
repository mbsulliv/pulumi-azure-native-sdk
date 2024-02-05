// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The mode to create a new server.
type CreateMode string

const (
	CreateModeDefault            = CreateMode("Default")
	CreateModePointInTimeRestore = CreateMode("PointInTimeRestore")
	CreateModeGeoRestore         = CreateMode("GeoRestore")
	CreateModeReplica            = CreateMode("Replica")
)

// Enable Geo-redundant or not for server backup.
type GeoRedundantBackup string

const (
	GeoRedundantBackupEnabled  = GeoRedundantBackup("Enabled")
	GeoRedundantBackupDisabled = GeoRedundantBackup("Disabled")
)

func (GeoRedundantBackup) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoRedundantBackup)(nil)).Elem()
}

func (e GeoRedundantBackup) ToGeoRedundantBackupOutput() GeoRedundantBackupOutput {
	return pulumi.ToOutput(e).(GeoRedundantBackupOutput)
}

func (e GeoRedundantBackup) ToGeoRedundantBackupOutputWithContext(ctx context.Context) GeoRedundantBackupOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GeoRedundantBackupOutput)
}

func (e GeoRedundantBackup) ToGeoRedundantBackupPtrOutput() GeoRedundantBackupPtrOutput {
	return e.ToGeoRedundantBackupPtrOutputWithContext(context.Background())
}

func (e GeoRedundantBackup) ToGeoRedundantBackupPtrOutputWithContext(ctx context.Context) GeoRedundantBackupPtrOutput {
	return GeoRedundantBackup(e).ToGeoRedundantBackupOutputWithContext(ctx).ToGeoRedundantBackupPtrOutputWithContext(ctx)
}

func (e GeoRedundantBackup) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GeoRedundantBackup) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GeoRedundantBackup) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GeoRedundantBackup) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GeoRedundantBackupOutput struct{ *pulumi.OutputState }

func (GeoRedundantBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoRedundantBackup)(nil)).Elem()
}

func (o GeoRedundantBackupOutput) ToGeoRedundantBackupOutput() GeoRedundantBackupOutput {
	return o
}

func (o GeoRedundantBackupOutput) ToGeoRedundantBackupOutputWithContext(ctx context.Context) GeoRedundantBackupOutput {
	return o
}

func (o GeoRedundantBackupOutput) ToGeoRedundantBackupPtrOutput() GeoRedundantBackupPtrOutput {
	return o.ToGeoRedundantBackupPtrOutputWithContext(context.Background())
}

func (o GeoRedundantBackupOutput) ToGeoRedundantBackupPtrOutputWithContext(ctx context.Context) GeoRedundantBackupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GeoRedundantBackup) *GeoRedundantBackup {
		return &v
	}).(GeoRedundantBackupPtrOutput)
}

func (o GeoRedundantBackupOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GeoRedundantBackupOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GeoRedundantBackup) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GeoRedundantBackupOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GeoRedundantBackupOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GeoRedundantBackup) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GeoRedundantBackupPtrOutput struct{ *pulumi.OutputState }

func (GeoRedundantBackupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeoRedundantBackup)(nil)).Elem()
}

func (o GeoRedundantBackupPtrOutput) ToGeoRedundantBackupPtrOutput() GeoRedundantBackupPtrOutput {
	return o
}

func (o GeoRedundantBackupPtrOutput) ToGeoRedundantBackupPtrOutputWithContext(ctx context.Context) GeoRedundantBackupPtrOutput {
	return o
}

func (o GeoRedundantBackupPtrOutput) Elem() GeoRedundantBackupOutput {
	return o.ApplyT(func(v *GeoRedundantBackup) GeoRedundantBackup {
		if v != nil {
			return *v
		}
		var ret GeoRedundantBackup
		return ret
	}).(GeoRedundantBackupOutput)
}

func (o GeoRedundantBackupPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GeoRedundantBackupPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GeoRedundantBackup) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// GeoRedundantBackupInput is an input type that accepts values of the GeoRedundantBackup enum
// A concrete instance of `GeoRedundantBackupInput` can be one of the following:
//
//	GeoRedundantBackupEnabled
//	GeoRedundantBackupDisabled
type GeoRedundantBackupInput interface {
	pulumi.Input

	ToGeoRedundantBackupOutput() GeoRedundantBackupOutput
	ToGeoRedundantBackupOutputWithContext(context.Context) GeoRedundantBackupOutput
}

var geoRedundantBackupPtrType = reflect.TypeOf((**GeoRedundantBackup)(nil)).Elem()

type GeoRedundantBackupPtrInput interface {
	pulumi.Input

	ToGeoRedundantBackupPtrOutput() GeoRedundantBackupPtrOutput
	ToGeoRedundantBackupPtrOutputWithContext(context.Context) GeoRedundantBackupPtrOutput
}

type geoRedundantBackupPtr string

func GeoRedundantBackupPtr(v string) GeoRedundantBackupPtrInput {
	return (*geoRedundantBackupPtr)(&v)
}

func (*geoRedundantBackupPtr) ElementType() reflect.Type {
	return geoRedundantBackupPtrType
}

func (in *geoRedundantBackupPtr) ToGeoRedundantBackupPtrOutput() GeoRedundantBackupPtrOutput {
	return pulumi.ToOutput(in).(GeoRedundantBackupPtrOutput)
}

func (in *geoRedundantBackupPtr) ToGeoRedundantBackupPtrOutputWithContext(ctx context.Context) GeoRedundantBackupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GeoRedundantBackupPtrOutput)
}

// Enforce a minimal Tls version for the server.
type MinimalTlsVersionEnum string

const (
	MinimalTlsVersionEnum_TLS1_0                = MinimalTlsVersionEnum("TLS1_0")
	MinimalTlsVersionEnum_TLS1_1                = MinimalTlsVersionEnum("TLS1_1")
	MinimalTlsVersionEnum_TLS1_2                = MinimalTlsVersionEnum("TLS1_2")
	MinimalTlsVersionEnumTLSEnforcementDisabled = MinimalTlsVersionEnum("TLSEnforcementDisabled")
)

func (MinimalTlsVersionEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*MinimalTlsVersionEnum)(nil)).Elem()
}

func (e MinimalTlsVersionEnum) ToMinimalTlsVersionEnumOutput() MinimalTlsVersionEnumOutput {
	return pulumi.ToOutput(e).(MinimalTlsVersionEnumOutput)
}

func (e MinimalTlsVersionEnum) ToMinimalTlsVersionEnumOutputWithContext(ctx context.Context) MinimalTlsVersionEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MinimalTlsVersionEnumOutput)
}

func (e MinimalTlsVersionEnum) ToMinimalTlsVersionEnumPtrOutput() MinimalTlsVersionEnumPtrOutput {
	return e.ToMinimalTlsVersionEnumPtrOutputWithContext(context.Background())
}

func (e MinimalTlsVersionEnum) ToMinimalTlsVersionEnumPtrOutputWithContext(ctx context.Context) MinimalTlsVersionEnumPtrOutput {
	return MinimalTlsVersionEnum(e).ToMinimalTlsVersionEnumOutputWithContext(ctx).ToMinimalTlsVersionEnumPtrOutputWithContext(ctx)
}

func (e MinimalTlsVersionEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MinimalTlsVersionEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MinimalTlsVersionEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MinimalTlsVersionEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MinimalTlsVersionEnumOutput struct{ *pulumi.OutputState }

func (MinimalTlsVersionEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MinimalTlsVersionEnum)(nil)).Elem()
}

func (o MinimalTlsVersionEnumOutput) ToMinimalTlsVersionEnumOutput() MinimalTlsVersionEnumOutput {
	return o
}

func (o MinimalTlsVersionEnumOutput) ToMinimalTlsVersionEnumOutputWithContext(ctx context.Context) MinimalTlsVersionEnumOutput {
	return o
}

func (o MinimalTlsVersionEnumOutput) ToMinimalTlsVersionEnumPtrOutput() MinimalTlsVersionEnumPtrOutput {
	return o.ToMinimalTlsVersionEnumPtrOutputWithContext(context.Background())
}

func (o MinimalTlsVersionEnumOutput) ToMinimalTlsVersionEnumPtrOutputWithContext(ctx context.Context) MinimalTlsVersionEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MinimalTlsVersionEnum) *MinimalTlsVersionEnum {
		return &v
	}).(MinimalTlsVersionEnumPtrOutput)
}

func (o MinimalTlsVersionEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MinimalTlsVersionEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MinimalTlsVersionEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MinimalTlsVersionEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MinimalTlsVersionEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MinimalTlsVersionEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MinimalTlsVersionEnumPtrOutput struct{ *pulumi.OutputState }

func (MinimalTlsVersionEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MinimalTlsVersionEnum)(nil)).Elem()
}

func (o MinimalTlsVersionEnumPtrOutput) ToMinimalTlsVersionEnumPtrOutput() MinimalTlsVersionEnumPtrOutput {
	return o
}

func (o MinimalTlsVersionEnumPtrOutput) ToMinimalTlsVersionEnumPtrOutputWithContext(ctx context.Context) MinimalTlsVersionEnumPtrOutput {
	return o
}

func (o MinimalTlsVersionEnumPtrOutput) Elem() MinimalTlsVersionEnumOutput {
	return o.ApplyT(func(v *MinimalTlsVersionEnum) MinimalTlsVersionEnum {
		if v != nil {
			return *v
		}
		var ret MinimalTlsVersionEnum
		return ret
	}).(MinimalTlsVersionEnumOutput)
}

func (o MinimalTlsVersionEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MinimalTlsVersionEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MinimalTlsVersionEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MinimalTlsVersionEnumInput is an input type that accepts values of the MinimalTlsVersionEnum enum
// A concrete instance of `MinimalTlsVersionEnumInput` can be one of the following:
//
//	MinimalTlsVersionEnum_TLS1_0
//	MinimalTlsVersionEnum_TLS1_1
//	MinimalTlsVersionEnum_TLS1_2
//	MinimalTlsVersionEnumTLSEnforcementDisabled
type MinimalTlsVersionEnumInput interface {
	pulumi.Input

	ToMinimalTlsVersionEnumOutput() MinimalTlsVersionEnumOutput
	ToMinimalTlsVersionEnumOutputWithContext(context.Context) MinimalTlsVersionEnumOutput
}

var minimalTlsVersionEnumPtrType = reflect.TypeOf((**MinimalTlsVersionEnum)(nil)).Elem()

type MinimalTlsVersionEnumPtrInput interface {
	pulumi.Input

	ToMinimalTlsVersionEnumPtrOutput() MinimalTlsVersionEnumPtrOutput
	ToMinimalTlsVersionEnumPtrOutputWithContext(context.Context) MinimalTlsVersionEnumPtrOutput
}

type minimalTlsVersionEnumPtr string

func MinimalTlsVersionEnumPtr(v string) MinimalTlsVersionEnumPtrInput {
	return (*minimalTlsVersionEnumPtr)(&v)
}

func (*minimalTlsVersionEnumPtr) ElementType() reflect.Type {
	return minimalTlsVersionEnumPtrType
}

func (in *minimalTlsVersionEnumPtr) ToMinimalTlsVersionEnumPtrOutput() MinimalTlsVersionEnumPtrOutput {
	return pulumi.ToOutput(in).(MinimalTlsVersionEnumPtrOutput)
}

func (in *minimalTlsVersionEnumPtr) ToMinimalTlsVersionEnumPtrOutputWithContext(ctx context.Context) MinimalTlsVersionEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MinimalTlsVersionEnumPtrOutput)
}

// Server version.
type ServerVersion string

const (
	ServerVersion_10_2 = ServerVersion("10.2")
	ServerVersion_10_3 = ServerVersion("10.3")
)

func (ServerVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerVersion)(nil)).Elem()
}

func (e ServerVersion) ToServerVersionOutput() ServerVersionOutput {
	return pulumi.ToOutput(e).(ServerVersionOutput)
}

func (e ServerVersion) ToServerVersionOutputWithContext(ctx context.Context) ServerVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServerVersionOutput)
}

func (e ServerVersion) ToServerVersionPtrOutput() ServerVersionPtrOutput {
	return e.ToServerVersionPtrOutputWithContext(context.Background())
}

func (e ServerVersion) ToServerVersionPtrOutputWithContext(ctx context.Context) ServerVersionPtrOutput {
	return ServerVersion(e).ToServerVersionOutputWithContext(ctx).ToServerVersionPtrOutputWithContext(ctx)
}

func (e ServerVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServerVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServerVersionOutput struct{ *pulumi.OutputState }

func (ServerVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerVersion)(nil)).Elem()
}

func (o ServerVersionOutput) ToServerVersionOutput() ServerVersionOutput {
	return o
}

func (o ServerVersionOutput) ToServerVersionOutputWithContext(ctx context.Context) ServerVersionOutput {
	return o
}

func (o ServerVersionOutput) ToServerVersionPtrOutput() ServerVersionPtrOutput {
	return o.ToServerVersionPtrOutputWithContext(context.Background())
}

func (o ServerVersionOutput) ToServerVersionPtrOutputWithContext(ctx context.Context) ServerVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerVersion) *ServerVersion {
		return &v
	}).(ServerVersionPtrOutput)
}

func (o ServerVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServerVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServerVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServerVersionPtrOutput struct{ *pulumi.OutputState }

func (ServerVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerVersion)(nil)).Elem()
}

func (o ServerVersionPtrOutput) ToServerVersionPtrOutput() ServerVersionPtrOutput {
	return o
}

func (o ServerVersionPtrOutput) ToServerVersionPtrOutputWithContext(ctx context.Context) ServerVersionPtrOutput {
	return o
}

func (o ServerVersionPtrOutput) Elem() ServerVersionOutput {
	return o.ApplyT(func(v *ServerVersion) ServerVersion {
		if v != nil {
			return *v
		}
		var ret ServerVersion
		return ret
	}).(ServerVersionOutput)
}

func (o ServerVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServerVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ServerVersionInput is an input type that accepts values of the ServerVersion enum
// A concrete instance of `ServerVersionInput` can be one of the following:
//
//	ServerVersion_10_2
//	ServerVersion_10_3
type ServerVersionInput interface {
	pulumi.Input

	ToServerVersionOutput() ServerVersionOutput
	ToServerVersionOutputWithContext(context.Context) ServerVersionOutput
}

var serverVersionPtrType = reflect.TypeOf((**ServerVersion)(nil)).Elem()

type ServerVersionPtrInput interface {
	pulumi.Input

	ToServerVersionPtrOutput() ServerVersionPtrOutput
	ToServerVersionPtrOutputWithContext(context.Context) ServerVersionPtrOutput
}

type serverVersionPtr string

func ServerVersionPtr(v string) ServerVersionPtrInput {
	return (*serverVersionPtr)(&v)
}

func (*serverVersionPtr) ElementType() reflect.Type {
	return serverVersionPtrType
}

func (in *serverVersionPtr) ToServerVersionPtrOutput() ServerVersionPtrOutput {
	return pulumi.ToOutput(in).(ServerVersionPtrOutput)
}

func (in *serverVersionPtr) ToServerVersionPtrOutputWithContext(ctx context.Context) ServerVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServerVersionPtrOutput)
}

// The tier of the particular SKU, e.g. Basic.
type SkuTier string

const (
	SkuTierBasic           = SkuTier("Basic")
	SkuTierGeneralPurpose  = SkuTier("GeneralPurpose")
	SkuTierMemoryOptimized = SkuTier("MemoryOptimized")
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

// SkuTierInput is an input type that accepts values of the SkuTier enum
// A concrete instance of `SkuTierInput` can be one of the following:
//
//	SkuTierBasic
//	SkuTierGeneralPurpose
//	SkuTierMemoryOptimized
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

// Enable ssl enforcement or not when connect to server.
type SslEnforcementEnum string

const (
	SslEnforcementEnumEnabled  = SslEnforcementEnum("Enabled")
	SslEnforcementEnumDisabled = SslEnforcementEnum("Disabled")
)

func (SslEnforcementEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SslEnforcementEnum)(nil)).Elem()
}

func (e SslEnforcementEnum) ToSslEnforcementEnumOutput() SslEnforcementEnumOutput {
	return pulumi.ToOutput(e).(SslEnforcementEnumOutput)
}

func (e SslEnforcementEnum) ToSslEnforcementEnumOutputWithContext(ctx context.Context) SslEnforcementEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SslEnforcementEnumOutput)
}

func (e SslEnforcementEnum) ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput {
	return e.ToSslEnforcementEnumPtrOutputWithContext(context.Background())
}

func (e SslEnforcementEnum) ToSslEnforcementEnumPtrOutputWithContext(ctx context.Context) SslEnforcementEnumPtrOutput {
	return SslEnforcementEnum(e).ToSslEnforcementEnumOutputWithContext(ctx).ToSslEnforcementEnumPtrOutputWithContext(ctx)
}

func (e SslEnforcementEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslEnforcementEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslEnforcementEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SslEnforcementEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SslEnforcementEnumOutput struct{ *pulumi.OutputState }

func (SslEnforcementEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslEnforcementEnum)(nil)).Elem()
}

func (o SslEnforcementEnumOutput) ToSslEnforcementEnumOutput() SslEnforcementEnumOutput {
	return o
}

func (o SslEnforcementEnumOutput) ToSslEnforcementEnumOutputWithContext(ctx context.Context) SslEnforcementEnumOutput {
	return o
}

func (o SslEnforcementEnumOutput) ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput {
	return o.ToSslEnforcementEnumPtrOutputWithContext(context.Background())
}

func (o SslEnforcementEnumOutput) ToSslEnforcementEnumPtrOutputWithContext(ctx context.Context) SslEnforcementEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslEnforcementEnum) *SslEnforcementEnum {
		return &v
	}).(SslEnforcementEnumPtrOutput)
}

func (o SslEnforcementEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SslEnforcementEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslEnforcementEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SslEnforcementEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslEnforcementEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslEnforcementEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SslEnforcementEnumPtrOutput struct{ *pulumi.OutputState }

func (SslEnforcementEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslEnforcementEnum)(nil)).Elem()
}

func (o SslEnforcementEnumPtrOutput) ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput {
	return o
}

func (o SslEnforcementEnumPtrOutput) ToSslEnforcementEnumPtrOutputWithContext(ctx context.Context) SslEnforcementEnumPtrOutput {
	return o
}

func (o SslEnforcementEnumPtrOutput) Elem() SslEnforcementEnumOutput {
	return o.ApplyT(func(v *SslEnforcementEnum) SslEnforcementEnum {
		if v != nil {
			return *v
		}
		var ret SslEnforcementEnum
		return ret
	}).(SslEnforcementEnumOutput)
}

func (o SslEnforcementEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslEnforcementEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SslEnforcementEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SslEnforcementEnumInput is an input type that accepts values of the SslEnforcementEnum enum
// A concrete instance of `SslEnforcementEnumInput` can be one of the following:
//
//	SslEnforcementEnumEnabled
//	SslEnforcementEnumDisabled
type SslEnforcementEnumInput interface {
	pulumi.Input

	ToSslEnforcementEnumOutput() SslEnforcementEnumOutput
	ToSslEnforcementEnumOutputWithContext(context.Context) SslEnforcementEnumOutput
}

var sslEnforcementEnumPtrType = reflect.TypeOf((**SslEnforcementEnum)(nil)).Elem()

type SslEnforcementEnumPtrInput interface {
	pulumi.Input

	ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput
	ToSslEnforcementEnumPtrOutputWithContext(context.Context) SslEnforcementEnumPtrOutput
}

type sslEnforcementEnumPtr string

func SslEnforcementEnumPtr(v string) SslEnforcementEnumPtrInput {
	return (*sslEnforcementEnumPtr)(&v)
}

func (*sslEnforcementEnumPtr) ElementType() reflect.Type {
	return sslEnforcementEnumPtrType
}

func (in *sslEnforcementEnumPtr) ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput {
	return pulumi.ToOutput(in).(SslEnforcementEnumPtrOutput)
}

func (in *sslEnforcementEnumPtr) ToSslEnforcementEnumPtrOutputWithContext(ctx context.Context) SslEnforcementEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SslEnforcementEnumPtrOutput)
}

// Enable Storage Auto Grow.
type StorageAutogrow string

const (
	StorageAutogrowEnabled  = StorageAutogrow("Enabled")
	StorageAutogrowDisabled = StorageAutogrow("Disabled")
)

func (StorageAutogrow) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAutogrow)(nil)).Elem()
}

func (e StorageAutogrow) ToStorageAutogrowOutput() StorageAutogrowOutput {
	return pulumi.ToOutput(e).(StorageAutogrowOutput)
}

func (e StorageAutogrow) ToStorageAutogrowOutputWithContext(ctx context.Context) StorageAutogrowOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageAutogrowOutput)
}

func (e StorageAutogrow) ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput {
	return e.ToStorageAutogrowPtrOutputWithContext(context.Background())
}

func (e StorageAutogrow) ToStorageAutogrowPtrOutputWithContext(ctx context.Context) StorageAutogrowPtrOutput {
	return StorageAutogrow(e).ToStorageAutogrowOutputWithContext(ctx).ToStorageAutogrowPtrOutputWithContext(ctx)
}

func (e StorageAutogrow) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAutogrow) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAutogrow) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageAutogrow) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageAutogrowOutput struct{ *pulumi.OutputState }

func (StorageAutogrowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAutogrow)(nil)).Elem()
}

func (o StorageAutogrowOutput) ToStorageAutogrowOutput() StorageAutogrowOutput {
	return o
}

func (o StorageAutogrowOutput) ToStorageAutogrowOutputWithContext(ctx context.Context) StorageAutogrowOutput {
	return o
}

func (o StorageAutogrowOutput) ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput {
	return o.ToStorageAutogrowPtrOutputWithContext(context.Background())
}

func (o StorageAutogrowOutput) ToStorageAutogrowPtrOutputWithContext(ctx context.Context) StorageAutogrowPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAutogrow) *StorageAutogrow {
		return &v
	}).(StorageAutogrowPtrOutput)
}

func (o StorageAutogrowOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageAutogrowOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAutogrow) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageAutogrowOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAutogrowOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAutogrow) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageAutogrowPtrOutput struct{ *pulumi.OutputState }

func (StorageAutogrowPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAutogrow)(nil)).Elem()
}

func (o StorageAutogrowPtrOutput) ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput {
	return o
}

func (o StorageAutogrowPtrOutput) ToStorageAutogrowPtrOutputWithContext(ctx context.Context) StorageAutogrowPtrOutput {
	return o
}

func (o StorageAutogrowPtrOutput) Elem() StorageAutogrowOutput {
	return o.ApplyT(func(v *StorageAutogrow) StorageAutogrow {
		if v != nil {
			return *v
		}
		var ret StorageAutogrow
		return ret
	}).(StorageAutogrowOutput)
}

func (o StorageAutogrowPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAutogrowPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageAutogrow) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// StorageAutogrowInput is an input type that accepts values of the StorageAutogrow enum
// A concrete instance of `StorageAutogrowInput` can be one of the following:
//
//	StorageAutogrowEnabled
//	StorageAutogrowDisabled
type StorageAutogrowInput interface {
	pulumi.Input

	ToStorageAutogrowOutput() StorageAutogrowOutput
	ToStorageAutogrowOutputWithContext(context.Context) StorageAutogrowOutput
}

var storageAutogrowPtrType = reflect.TypeOf((**StorageAutogrow)(nil)).Elem()

type StorageAutogrowPtrInput interface {
	pulumi.Input

	ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput
	ToStorageAutogrowPtrOutputWithContext(context.Context) StorageAutogrowPtrOutput
}

type storageAutogrowPtr string

func StorageAutogrowPtr(v string) StorageAutogrowPtrInput {
	return (*storageAutogrowPtr)(&v)
}

func (*storageAutogrowPtr) ElementType() reflect.Type {
	return storageAutogrowPtrType
}

func (in *storageAutogrowPtr) ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput {
	return pulumi.ToOutput(in).(StorageAutogrowPtrOutput)
}

func (in *storageAutogrowPtr) ToStorageAutogrowPtrOutputWithContext(ctx context.Context) StorageAutogrowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageAutogrowPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GeoRedundantBackupOutput{})
	pulumi.RegisterOutputType(GeoRedundantBackupPtrOutput{})
	pulumi.RegisterOutputType(MinimalTlsVersionEnumOutput{})
	pulumi.RegisterOutputType(MinimalTlsVersionEnumPtrOutput{})
	pulumi.RegisterOutputType(ServerVersionOutput{})
	pulumi.RegisterOutputType(ServerVersionPtrOutput{})
	pulumi.RegisterOutputType(SkuTierOutput{})
	pulumi.RegisterOutputType(SkuTierPtrOutput{})
	pulumi.RegisterOutputType(SslEnforcementEnumOutput{})
	pulumi.RegisterOutputType(SslEnforcementEnumPtrOutput{})
	pulumi.RegisterOutputType(StorageAutogrowOutput{})
	pulumi.RegisterOutputType(StorageAutogrowPtrOutput{})
}
