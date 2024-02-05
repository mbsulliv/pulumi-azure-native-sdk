// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Kind of data set.
type DataSetKind string

const (
	DataSetKindBlob                         = DataSetKind("Blob")
	DataSetKindContainer                    = DataSetKind("Container")
	DataSetKindBlobFolder                   = DataSetKind("BlobFolder")
	DataSetKindAdlsGen2FileSystem           = DataSetKind("AdlsGen2FileSystem")
	DataSetKindAdlsGen2Folder               = DataSetKind("AdlsGen2Folder")
	DataSetKindAdlsGen2File                 = DataSetKind("AdlsGen2File")
	DataSetKindAdlsGen1Folder               = DataSetKind("AdlsGen1Folder")
	DataSetKindAdlsGen1File                 = DataSetKind("AdlsGen1File")
	DataSetKindKustoCluster                 = DataSetKind("KustoCluster")
	DataSetKindKustoDatabase                = DataSetKind("KustoDatabase")
	DataSetKindKustoTable                   = DataSetKind("KustoTable")
	DataSetKindSqlDBTable                   = DataSetKind("SqlDBTable")
	DataSetKindSqlDWTable                   = DataSetKind("SqlDWTable")
	DataSetKindSynapseWorkspaceSqlPoolTable = DataSetKind("SynapseWorkspaceSqlPoolTable")
)

// Kind of data set mapping.
type DataSetMappingKind string

const (
	DataSetMappingKindBlob                         = DataSetMappingKind("Blob")
	DataSetMappingKindContainer                    = DataSetMappingKind("Container")
	DataSetMappingKindBlobFolder                   = DataSetMappingKind("BlobFolder")
	DataSetMappingKindAdlsGen2FileSystem           = DataSetMappingKind("AdlsGen2FileSystem")
	DataSetMappingKindAdlsGen2Folder               = DataSetMappingKind("AdlsGen2Folder")
	DataSetMappingKindAdlsGen2File                 = DataSetMappingKind("AdlsGen2File")
	DataSetMappingKindKustoCluster                 = DataSetMappingKind("KustoCluster")
	DataSetMappingKindKustoDatabase                = DataSetMappingKind("KustoDatabase")
	DataSetMappingKindKustoTable                   = DataSetMappingKind("KustoTable")
	DataSetMappingKindSqlDBTable                   = DataSetMappingKind("SqlDBTable")
	DataSetMappingKindSqlDWTable                   = DataSetMappingKind("SqlDWTable")
	DataSetMappingKindSynapseWorkspaceSqlPoolTable = DataSetMappingKind("SynapseWorkspaceSqlPoolTable")
)

// File output type
type OutputType string

const (
	OutputTypeCsv     = OutputType("Csv")
	OutputTypeParquet = OutputType("Parquet")
)

func (OutputType) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputType)(nil)).Elem()
}

func (e OutputType) ToOutputTypeOutput() OutputTypeOutput {
	return pulumi.ToOutput(e).(OutputTypeOutput)
}

func (e OutputType) ToOutputTypeOutputWithContext(ctx context.Context) OutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OutputTypeOutput)
}

func (e OutputType) ToOutputTypePtrOutput() OutputTypePtrOutput {
	return e.ToOutputTypePtrOutputWithContext(context.Background())
}

func (e OutputType) ToOutputTypePtrOutputWithContext(ctx context.Context) OutputTypePtrOutput {
	return OutputType(e).ToOutputTypeOutputWithContext(ctx).ToOutputTypePtrOutputWithContext(ctx)
}

func (e OutputType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OutputType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OutputType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OutputType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OutputTypeOutput struct{ *pulumi.OutputState }

func (OutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputType)(nil)).Elem()
}

func (o OutputTypeOutput) ToOutputTypeOutput() OutputTypeOutput {
	return o
}

func (o OutputTypeOutput) ToOutputTypeOutputWithContext(ctx context.Context) OutputTypeOutput {
	return o
}

func (o OutputTypeOutput) ToOutputTypePtrOutput() OutputTypePtrOutput {
	return o.ToOutputTypePtrOutputWithContext(context.Background())
}

func (o OutputTypeOutput) ToOutputTypePtrOutputWithContext(ctx context.Context) OutputTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OutputType) *OutputType {
		return &v
	}).(OutputTypePtrOutput)
}

func (o OutputTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OutputTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutputType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OutputTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutputTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutputType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OutputTypePtrOutput struct{ *pulumi.OutputState }

func (OutputTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputType)(nil)).Elem()
}

func (o OutputTypePtrOutput) ToOutputTypePtrOutput() OutputTypePtrOutput {
	return o
}

func (o OutputTypePtrOutput) ToOutputTypePtrOutputWithContext(ctx context.Context) OutputTypePtrOutput {
	return o
}

func (o OutputTypePtrOutput) Elem() OutputTypeOutput {
	return o.ApplyT(func(v *OutputType) OutputType {
		if v != nil {
			return *v
		}
		var ret OutputType
		return ret
	}).(OutputTypeOutput)
}

func (o OutputTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutputTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OutputType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// OutputTypeInput is an input type that accepts values of the OutputType enum
// A concrete instance of `OutputTypeInput` can be one of the following:
//
//	OutputTypeCsv
//	OutputTypeParquet
type OutputTypeInput interface {
	pulumi.Input

	ToOutputTypeOutput() OutputTypeOutput
	ToOutputTypeOutputWithContext(context.Context) OutputTypeOutput
}

var outputTypePtrType = reflect.TypeOf((**OutputType)(nil)).Elem()

type OutputTypePtrInput interface {
	pulumi.Input

	ToOutputTypePtrOutput() OutputTypePtrOutput
	ToOutputTypePtrOutputWithContext(context.Context) OutputTypePtrOutput
}

type outputTypePtr string

func OutputTypePtr(v string) OutputTypePtrInput {
	return (*outputTypePtr)(&v)
}

func (*outputTypePtr) ElementType() reflect.Type {
	return outputTypePtrType
}

func (in *outputTypePtr) ToOutputTypePtrOutput() OutputTypePtrOutput {
	return pulumi.ToOutput(in).(OutputTypePtrOutput)
}

func (in *outputTypePtr) ToOutputTypePtrOutputWithContext(ctx context.Context) OutputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OutputTypePtrOutput)
}

// Recurrence Interval
type RecurrenceInterval string

const (
	RecurrenceIntervalHour = RecurrenceInterval("Hour")
	RecurrenceIntervalDay  = RecurrenceInterval("Day")
)

func (RecurrenceInterval) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceInterval)(nil)).Elem()
}

func (e RecurrenceInterval) ToRecurrenceIntervalOutput() RecurrenceIntervalOutput {
	return pulumi.ToOutput(e).(RecurrenceIntervalOutput)
}

func (e RecurrenceInterval) ToRecurrenceIntervalOutputWithContext(ctx context.Context) RecurrenceIntervalOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RecurrenceIntervalOutput)
}

func (e RecurrenceInterval) ToRecurrenceIntervalPtrOutput() RecurrenceIntervalPtrOutput {
	return e.ToRecurrenceIntervalPtrOutputWithContext(context.Background())
}

func (e RecurrenceInterval) ToRecurrenceIntervalPtrOutputWithContext(ctx context.Context) RecurrenceIntervalPtrOutput {
	return RecurrenceInterval(e).ToRecurrenceIntervalOutputWithContext(ctx).ToRecurrenceIntervalPtrOutputWithContext(ctx)
}

func (e RecurrenceInterval) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurrenceInterval) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurrenceInterval) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RecurrenceInterval) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RecurrenceIntervalOutput struct{ *pulumi.OutputState }

func (RecurrenceIntervalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceInterval)(nil)).Elem()
}

func (o RecurrenceIntervalOutput) ToRecurrenceIntervalOutput() RecurrenceIntervalOutput {
	return o
}

func (o RecurrenceIntervalOutput) ToRecurrenceIntervalOutputWithContext(ctx context.Context) RecurrenceIntervalOutput {
	return o
}

func (o RecurrenceIntervalOutput) ToRecurrenceIntervalPtrOutput() RecurrenceIntervalPtrOutput {
	return o.ToRecurrenceIntervalPtrOutputWithContext(context.Background())
}

func (o RecurrenceIntervalOutput) ToRecurrenceIntervalPtrOutputWithContext(ctx context.Context) RecurrenceIntervalPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecurrenceInterval) *RecurrenceInterval {
		return &v
	}).(RecurrenceIntervalPtrOutput)
}

func (o RecurrenceIntervalOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RecurrenceIntervalOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecurrenceInterval) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RecurrenceIntervalOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecurrenceIntervalOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecurrenceInterval) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RecurrenceIntervalPtrOutput struct{ *pulumi.OutputState }

func (RecurrenceIntervalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrenceInterval)(nil)).Elem()
}

func (o RecurrenceIntervalPtrOutput) ToRecurrenceIntervalPtrOutput() RecurrenceIntervalPtrOutput {
	return o
}

func (o RecurrenceIntervalPtrOutput) ToRecurrenceIntervalPtrOutputWithContext(ctx context.Context) RecurrenceIntervalPtrOutput {
	return o
}

func (o RecurrenceIntervalPtrOutput) Elem() RecurrenceIntervalOutput {
	return o.ApplyT(func(v *RecurrenceInterval) RecurrenceInterval {
		if v != nil {
			return *v
		}
		var ret RecurrenceInterval
		return ret
	}).(RecurrenceIntervalOutput)
}

func (o RecurrenceIntervalPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecurrenceIntervalPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RecurrenceInterval) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// RecurrenceIntervalInput is an input type that accepts values of the RecurrenceInterval enum
// A concrete instance of `RecurrenceIntervalInput` can be one of the following:
//
//	RecurrenceIntervalHour
//	RecurrenceIntervalDay
type RecurrenceIntervalInput interface {
	pulumi.Input

	ToRecurrenceIntervalOutput() RecurrenceIntervalOutput
	ToRecurrenceIntervalOutputWithContext(context.Context) RecurrenceIntervalOutput
}

var recurrenceIntervalPtrType = reflect.TypeOf((**RecurrenceInterval)(nil)).Elem()

type RecurrenceIntervalPtrInput interface {
	pulumi.Input

	ToRecurrenceIntervalPtrOutput() RecurrenceIntervalPtrOutput
	ToRecurrenceIntervalPtrOutputWithContext(context.Context) RecurrenceIntervalPtrOutput
}

type recurrenceIntervalPtr string

func RecurrenceIntervalPtr(v string) RecurrenceIntervalPtrInput {
	return (*recurrenceIntervalPtr)(&v)
}

func (*recurrenceIntervalPtr) ElementType() reflect.Type {
	return recurrenceIntervalPtrType
}

func (in *recurrenceIntervalPtr) ToRecurrenceIntervalPtrOutput() RecurrenceIntervalPtrOutput {
	return pulumi.ToOutput(in).(RecurrenceIntervalPtrOutput)
}

func (in *recurrenceIntervalPtr) ToRecurrenceIntervalPtrOutputWithContext(ctx context.Context) RecurrenceIntervalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RecurrenceIntervalPtrOutput)
}

// Share kind.
type ShareKind string

const (
	ShareKindCopyBased = ShareKind("CopyBased")
	ShareKindInPlace   = ShareKind("InPlace")
)

func (ShareKind) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareKind)(nil)).Elem()
}

func (e ShareKind) ToShareKindOutput() ShareKindOutput {
	return pulumi.ToOutput(e).(ShareKindOutput)
}

func (e ShareKind) ToShareKindOutputWithContext(ctx context.Context) ShareKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ShareKindOutput)
}

func (e ShareKind) ToShareKindPtrOutput() ShareKindPtrOutput {
	return e.ToShareKindPtrOutputWithContext(context.Background())
}

func (e ShareKind) ToShareKindPtrOutputWithContext(ctx context.Context) ShareKindPtrOutput {
	return ShareKind(e).ToShareKindOutputWithContext(ctx).ToShareKindPtrOutputWithContext(ctx)
}

func (e ShareKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShareKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShareKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ShareKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ShareKindOutput struct{ *pulumi.OutputState }

func (ShareKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareKind)(nil)).Elem()
}

func (o ShareKindOutput) ToShareKindOutput() ShareKindOutput {
	return o
}

func (o ShareKindOutput) ToShareKindOutputWithContext(ctx context.Context) ShareKindOutput {
	return o
}

func (o ShareKindOutput) ToShareKindPtrOutput() ShareKindPtrOutput {
	return o.ToShareKindPtrOutputWithContext(context.Background())
}

func (o ShareKindOutput) ToShareKindPtrOutputWithContext(ctx context.Context) ShareKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ShareKind) *ShareKind {
		return &v
	}).(ShareKindPtrOutput)
}

func (o ShareKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ShareKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ShareKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ShareKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ShareKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ShareKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ShareKindPtrOutput struct{ *pulumi.OutputState }

func (ShareKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareKind)(nil)).Elem()
}

func (o ShareKindPtrOutput) ToShareKindPtrOutput() ShareKindPtrOutput {
	return o
}

func (o ShareKindPtrOutput) ToShareKindPtrOutputWithContext(ctx context.Context) ShareKindPtrOutput {
	return o
}

func (o ShareKindPtrOutput) Elem() ShareKindOutput {
	return o.ApplyT(func(v *ShareKind) ShareKind {
		if v != nil {
			return *v
		}
		var ret ShareKind
		return ret
	}).(ShareKindOutput)
}

func (o ShareKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ShareKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ShareKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ShareKindInput is an input type that accepts values of the ShareKind enum
// A concrete instance of `ShareKindInput` can be one of the following:
//
//	ShareKindCopyBased
//	ShareKindInPlace
type ShareKindInput interface {
	pulumi.Input

	ToShareKindOutput() ShareKindOutput
	ToShareKindOutputWithContext(context.Context) ShareKindOutput
}

var shareKindPtrType = reflect.TypeOf((**ShareKind)(nil)).Elem()

type ShareKindPtrInput interface {
	pulumi.Input

	ToShareKindPtrOutput() ShareKindPtrOutput
	ToShareKindPtrOutputWithContext(context.Context) ShareKindPtrOutput
}

type shareKindPtr string

func ShareKindPtr(v string) ShareKindPtrInput {
	return (*shareKindPtr)(&v)
}

func (*shareKindPtr) ElementType() reflect.Type {
	return shareKindPtrType
}

func (in *shareKindPtr) ToShareKindPtrOutput() ShareKindPtrOutput {
	return pulumi.ToOutput(in).(ShareKindPtrOutput)
}

func (in *shareKindPtr) ToShareKindPtrOutputWithContext(ctx context.Context) ShareKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ShareKindPtrOutput)
}

// Synchronization mode
type SynchronizationMode string

const (
	SynchronizationModeIncremental = SynchronizationMode("Incremental")
	SynchronizationModeFullSync    = SynchronizationMode("FullSync")
)

func (SynchronizationMode) ElementType() reflect.Type {
	return reflect.TypeOf((*SynchronizationMode)(nil)).Elem()
}

func (e SynchronizationMode) ToSynchronizationModeOutput() SynchronizationModeOutput {
	return pulumi.ToOutput(e).(SynchronizationModeOutput)
}

func (e SynchronizationMode) ToSynchronizationModeOutputWithContext(ctx context.Context) SynchronizationModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SynchronizationModeOutput)
}

func (e SynchronizationMode) ToSynchronizationModePtrOutput() SynchronizationModePtrOutput {
	return e.ToSynchronizationModePtrOutputWithContext(context.Background())
}

func (e SynchronizationMode) ToSynchronizationModePtrOutputWithContext(ctx context.Context) SynchronizationModePtrOutput {
	return SynchronizationMode(e).ToSynchronizationModeOutputWithContext(ctx).ToSynchronizationModePtrOutputWithContext(ctx)
}

func (e SynchronizationMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SynchronizationMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SynchronizationMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SynchronizationMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SynchronizationModeOutput struct{ *pulumi.OutputState }

func (SynchronizationModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SynchronizationMode)(nil)).Elem()
}

func (o SynchronizationModeOutput) ToSynchronizationModeOutput() SynchronizationModeOutput {
	return o
}

func (o SynchronizationModeOutput) ToSynchronizationModeOutputWithContext(ctx context.Context) SynchronizationModeOutput {
	return o
}

func (o SynchronizationModeOutput) ToSynchronizationModePtrOutput() SynchronizationModePtrOutput {
	return o.ToSynchronizationModePtrOutputWithContext(context.Background())
}

func (o SynchronizationModeOutput) ToSynchronizationModePtrOutputWithContext(ctx context.Context) SynchronizationModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SynchronizationMode) *SynchronizationMode {
		return &v
	}).(SynchronizationModePtrOutput)
}

func (o SynchronizationModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SynchronizationModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SynchronizationMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SynchronizationModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SynchronizationModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SynchronizationMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SynchronizationModePtrOutput struct{ *pulumi.OutputState }

func (SynchronizationModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SynchronizationMode)(nil)).Elem()
}

func (o SynchronizationModePtrOutput) ToSynchronizationModePtrOutput() SynchronizationModePtrOutput {
	return o
}

func (o SynchronizationModePtrOutput) ToSynchronizationModePtrOutputWithContext(ctx context.Context) SynchronizationModePtrOutput {
	return o
}

func (o SynchronizationModePtrOutput) Elem() SynchronizationModeOutput {
	return o.ApplyT(func(v *SynchronizationMode) SynchronizationMode {
		if v != nil {
			return *v
		}
		var ret SynchronizationMode
		return ret
	}).(SynchronizationModeOutput)
}

func (o SynchronizationModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SynchronizationModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SynchronizationMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SynchronizationModeInput is an input type that accepts values of the SynchronizationMode enum
// A concrete instance of `SynchronizationModeInput` can be one of the following:
//
//	SynchronizationModeIncremental
//	SynchronizationModeFullSync
type SynchronizationModeInput interface {
	pulumi.Input

	ToSynchronizationModeOutput() SynchronizationModeOutput
	ToSynchronizationModeOutputWithContext(context.Context) SynchronizationModeOutput
}

var synchronizationModePtrType = reflect.TypeOf((**SynchronizationMode)(nil)).Elem()

type SynchronizationModePtrInput interface {
	pulumi.Input

	ToSynchronizationModePtrOutput() SynchronizationModePtrOutput
	ToSynchronizationModePtrOutputWithContext(context.Context) SynchronizationModePtrOutput
}

type synchronizationModePtr string

func SynchronizationModePtr(v string) SynchronizationModePtrInput {
	return (*synchronizationModePtr)(&v)
}

func (*synchronizationModePtr) ElementType() reflect.Type {
	return synchronizationModePtrType
}

func (in *synchronizationModePtr) ToSynchronizationModePtrOutput() SynchronizationModePtrOutput {
	return pulumi.ToOutput(in).(SynchronizationModePtrOutput)
}

func (in *synchronizationModePtr) ToSynchronizationModePtrOutputWithContext(ctx context.Context) SynchronizationModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SynchronizationModePtrOutput)
}

// Kind of synchronization setting.
type SynchronizationSettingKind string

const (
	SynchronizationSettingKindScheduleBased = SynchronizationSettingKind("ScheduleBased")
)

// Kind of synchronization on trigger.
type TriggerKind string

const (
	TriggerKindScheduleBased = TriggerKind("ScheduleBased")
)

// Identity Type
type Type string

const (
	TypeSystemAssigned = Type("SystemAssigned")
)

func (Type) ElementType() reflect.Type {
	return reflect.TypeOf((*Type)(nil)).Elem()
}

func (e Type) ToTypeOutput() TypeOutput {
	return pulumi.ToOutput(e).(TypeOutput)
}

func (e Type) ToTypeOutputWithContext(ctx context.Context) TypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TypeOutput)
}

func (e Type) ToTypePtrOutput() TypePtrOutput {
	return e.ToTypePtrOutputWithContext(context.Background())
}

func (e Type) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return Type(e).ToTypeOutputWithContext(ctx).ToTypePtrOutputWithContext(ctx)
}

func (e Type) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Type) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Type) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Type) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TypeOutput struct{ *pulumi.OutputState }

func (TypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Type)(nil)).Elem()
}

func (o TypeOutput) ToTypeOutput() TypeOutput {
	return o
}

func (o TypeOutput) ToTypeOutputWithContext(ctx context.Context) TypeOutput {
	return o
}

func (o TypeOutput) ToTypePtrOutput() TypePtrOutput {
	return o.ToTypePtrOutputWithContext(context.Background())
}

func (o TypeOutput) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Type) *Type {
		return &v
	}).(TypePtrOutput)
}

func (o TypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Type) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Type) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TypePtrOutput struct{ *pulumi.OutputState }

func (TypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Type)(nil)).Elem()
}

func (o TypePtrOutput) ToTypePtrOutput() TypePtrOutput {
	return o
}

func (o TypePtrOutput) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return o
}

func (o TypePtrOutput) Elem() TypeOutput {
	return o.ApplyT(func(v *Type) Type {
		if v != nil {
			return *v
		}
		var ret Type
		return ret
	}).(TypeOutput)
}

func (o TypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Type) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TypeInput is an input type that accepts values of the Type enum
// A concrete instance of `TypeInput` can be one of the following:
//
//	TypeSystemAssigned
type TypeInput interface {
	pulumi.Input

	ToTypeOutput() TypeOutput
	ToTypeOutputWithContext(context.Context) TypeOutput
}

var typePtrType = reflect.TypeOf((**Type)(nil)).Elem()

type TypePtrInput interface {
	pulumi.Input

	ToTypePtrOutput() TypePtrOutput
	ToTypePtrOutputWithContext(context.Context) TypePtrOutput
}

type typePtr string

func TypePtr(v string) TypePtrInput {
	return (*typePtr)(&v)
}

func (*typePtr) ElementType() reflect.Type {
	return typePtrType
}

func (in *typePtr) ToTypePtrOutput() TypePtrOutput {
	return pulumi.ToOutput(in).(TypePtrOutput)
}

func (in *typePtr) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(OutputTypeOutput{})
	pulumi.RegisterOutputType(OutputTypePtrOutput{})
	pulumi.RegisterOutputType(RecurrenceIntervalOutput{})
	pulumi.RegisterOutputType(RecurrenceIntervalPtrOutput{})
	pulumi.RegisterOutputType(ShareKindOutput{})
	pulumi.RegisterOutputType(ShareKindPtrOutput{})
	pulumi.RegisterOutputType(SynchronizationModeOutput{})
	pulumi.RegisterOutputType(SynchronizationModePtrOutput{})
	pulumi.RegisterOutputType(TypeOutput{})
	pulumi.RegisterOutputType(TypePtrOutput{})
}
