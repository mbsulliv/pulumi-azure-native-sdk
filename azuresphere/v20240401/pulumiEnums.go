// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Flag to define if the user allows for crash dump collection.
type AllowCrashDumpCollection string

const (
	// Crash dump collection enabled
	AllowCrashDumpCollectionEnabled = AllowCrashDumpCollection("Enabled")
	// Crash dump collection disabled
	AllowCrashDumpCollectionDisabled = AllowCrashDumpCollection("Disabled")
)

func (AllowCrashDumpCollection) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowCrashDumpCollection)(nil)).Elem()
}

func (e AllowCrashDumpCollection) ToAllowCrashDumpCollectionOutput() AllowCrashDumpCollectionOutput {
	return pulumi.ToOutput(e).(AllowCrashDumpCollectionOutput)
}

func (e AllowCrashDumpCollection) ToAllowCrashDumpCollectionOutputWithContext(ctx context.Context) AllowCrashDumpCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AllowCrashDumpCollectionOutput)
}

func (e AllowCrashDumpCollection) ToAllowCrashDumpCollectionPtrOutput() AllowCrashDumpCollectionPtrOutput {
	return e.ToAllowCrashDumpCollectionPtrOutputWithContext(context.Background())
}

func (e AllowCrashDumpCollection) ToAllowCrashDumpCollectionPtrOutputWithContext(ctx context.Context) AllowCrashDumpCollectionPtrOutput {
	return AllowCrashDumpCollection(e).ToAllowCrashDumpCollectionOutputWithContext(ctx).ToAllowCrashDumpCollectionPtrOutputWithContext(ctx)
}

func (e AllowCrashDumpCollection) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AllowCrashDumpCollection) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AllowCrashDumpCollection) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AllowCrashDumpCollection) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AllowCrashDumpCollectionOutput struct{ *pulumi.OutputState }

func (AllowCrashDumpCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowCrashDumpCollection)(nil)).Elem()
}

func (o AllowCrashDumpCollectionOutput) ToAllowCrashDumpCollectionOutput() AllowCrashDumpCollectionOutput {
	return o
}

func (o AllowCrashDumpCollectionOutput) ToAllowCrashDumpCollectionOutputWithContext(ctx context.Context) AllowCrashDumpCollectionOutput {
	return o
}

func (o AllowCrashDumpCollectionOutput) ToAllowCrashDumpCollectionPtrOutput() AllowCrashDumpCollectionPtrOutput {
	return o.ToAllowCrashDumpCollectionPtrOutputWithContext(context.Background())
}

func (o AllowCrashDumpCollectionOutput) ToAllowCrashDumpCollectionPtrOutputWithContext(ctx context.Context) AllowCrashDumpCollectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AllowCrashDumpCollection) *AllowCrashDumpCollection {
		return &v
	}).(AllowCrashDumpCollectionPtrOutput)
}

func (o AllowCrashDumpCollectionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AllowCrashDumpCollectionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AllowCrashDumpCollection) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AllowCrashDumpCollectionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AllowCrashDumpCollectionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AllowCrashDumpCollection) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AllowCrashDumpCollectionPtrOutput struct{ *pulumi.OutputState }

func (AllowCrashDumpCollectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowCrashDumpCollection)(nil)).Elem()
}

func (o AllowCrashDumpCollectionPtrOutput) ToAllowCrashDumpCollectionPtrOutput() AllowCrashDumpCollectionPtrOutput {
	return o
}

func (o AllowCrashDumpCollectionPtrOutput) ToAllowCrashDumpCollectionPtrOutputWithContext(ctx context.Context) AllowCrashDumpCollectionPtrOutput {
	return o
}

func (o AllowCrashDumpCollectionPtrOutput) Elem() AllowCrashDumpCollectionOutput {
	return o.ApplyT(func(v *AllowCrashDumpCollection) AllowCrashDumpCollection {
		if v != nil {
			return *v
		}
		var ret AllowCrashDumpCollection
		return ret
	}).(AllowCrashDumpCollectionOutput)
}

func (o AllowCrashDumpCollectionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AllowCrashDumpCollectionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AllowCrashDumpCollection) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AllowCrashDumpCollectionInput is an input type that accepts values of the AllowCrashDumpCollection enum
// A concrete instance of `AllowCrashDumpCollectionInput` can be one of the following:
//
//	AllowCrashDumpCollectionEnabled
//	AllowCrashDumpCollectionDisabled
type AllowCrashDumpCollectionInput interface {
	pulumi.Input

	ToAllowCrashDumpCollectionOutput() AllowCrashDumpCollectionOutput
	ToAllowCrashDumpCollectionOutputWithContext(context.Context) AllowCrashDumpCollectionOutput
}

var allowCrashDumpCollectionPtrType = reflect.TypeOf((**AllowCrashDumpCollection)(nil)).Elem()

type AllowCrashDumpCollectionPtrInput interface {
	pulumi.Input

	ToAllowCrashDumpCollectionPtrOutput() AllowCrashDumpCollectionPtrOutput
	ToAllowCrashDumpCollectionPtrOutputWithContext(context.Context) AllowCrashDumpCollectionPtrOutput
}

type allowCrashDumpCollectionPtr string

func AllowCrashDumpCollectionPtr(v string) AllowCrashDumpCollectionPtrInput {
	return (*allowCrashDumpCollectionPtr)(&v)
}

func (*allowCrashDumpCollectionPtr) ElementType() reflect.Type {
	return allowCrashDumpCollectionPtrType
}

func (in *allowCrashDumpCollectionPtr) ToAllowCrashDumpCollectionPtrOutput() AllowCrashDumpCollectionPtrOutput {
	return pulumi.ToOutput(in).(AllowCrashDumpCollectionPtrOutput)
}

func (in *allowCrashDumpCollectionPtr) ToAllowCrashDumpCollectionPtrOutputWithContext(ctx context.Context) AllowCrashDumpCollectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AllowCrashDumpCollectionPtrOutput)
}

// Operating system feed type of the device group.
type OSFeedType string

const (
	// Retail OS feed type.
	OSFeedTypeRetail = OSFeedType("Retail")
	// Retail evaluation OS feed type.
	OSFeedTypeRetailEval = OSFeedType("RetailEval")
)

func (OSFeedType) ElementType() reflect.Type {
	return reflect.TypeOf((*OSFeedType)(nil)).Elem()
}

func (e OSFeedType) ToOSFeedTypeOutput() OSFeedTypeOutput {
	return pulumi.ToOutput(e).(OSFeedTypeOutput)
}

func (e OSFeedType) ToOSFeedTypeOutputWithContext(ctx context.Context) OSFeedTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OSFeedTypeOutput)
}

func (e OSFeedType) ToOSFeedTypePtrOutput() OSFeedTypePtrOutput {
	return e.ToOSFeedTypePtrOutputWithContext(context.Background())
}

func (e OSFeedType) ToOSFeedTypePtrOutputWithContext(ctx context.Context) OSFeedTypePtrOutput {
	return OSFeedType(e).ToOSFeedTypeOutputWithContext(ctx).ToOSFeedTypePtrOutputWithContext(ctx)
}

func (e OSFeedType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OSFeedType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OSFeedType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OSFeedType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OSFeedTypeOutput struct{ *pulumi.OutputState }

func (OSFeedTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSFeedType)(nil)).Elem()
}

func (o OSFeedTypeOutput) ToOSFeedTypeOutput() OSFeedTypeOutput {
	return o
}

func (o OSFeedTypeOutput) ToOSFeedTypeOutputWithContext(ctx context.Context) OSFeedTypeOutput {
	return o
}

func (o OSFeedTypeOutput) ToOSFeedTypePtrOutput() OSFeedTypePtrOutput {
	return o.ToOSFeedTypePtrOutputWithContext(context.Background())
}

func (o OSFeedTypeOutput) ToOSFeedTypePtrOutputWithContext(ctx context.Context) OSFeedTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSFeedType) *OSFeedType {
		return &v
	}).(OSFeedTypePtrOutput)
}

func (o OSFeedTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OSFeedTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OSFeedType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OSFeedTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OSFeedTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OSFeedType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OSFeedTypePtrOutput struct{ *pulumi.OutputState }

func (OSFeedTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSFeedType)(nil)).Elem()
}

func (o OSFeedTypePtrOutput) ToOSFeedTypePtrOutput() OSFeedTypePtrOutput {
	return o
}

func (o OSFeedTypePtrOutput) ToOSFeedTypePtrOutputWithContext(ctx context.Context) OSFeedTypePtrOutput {
	return o
}

func (o OSFeedTypePtrOutput) Elem() OSFeedTypeOutput {
	return o.ApplyT(func(v *OSFeedType) OSFeedType {
		if v != nil {
			return *v
		}
		var ret OSFeedType
		return ret
	}).(OSFeedTypeOutput)
}

func (o OSFeedTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OSFeedTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OSFeedType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// OSFeedTypeInput is an input type that accepts values of the OSFeedType enum
// A concrete instance of `OSFeedTypeInput` can be one of the following:
//
//	OSFeedTypeRetail
//	OSFeedTypeRetailEval
type OSFeedTypeInput interface {
	pulumi.Input

	ToOSFeedTypeOutput() OSFeedTypeOutput
	ToOSFeedTypeOutputWithContext(context.Context) OSFeedTypeOutput
}

var osfeedTypePtrType = reflect.TypeOf((**OSFeedType)(nil)).Elem()

type OSFeedTypePtrInput interface {
	pulumi.Input

	ToOSFeedTypePtrOutput() OSFeedTypePtrOutput
	ToOSFeedTypePtrOutputWithContext(context.Context) OSFeedTypePtrOutput
}

type osfeedTypePtr string

func OSFeedTypePtr(v string) OSFeedTypePtrInput {
	return (*osfeedTypePtr)(&v)
}

func (*osfeedTypePtr) ElementType() reflect.Type {
	return osfeedTypePtrType
}

func (in *osfeedTypePtr) ToOSFeedTypePtrOutput() OSFeedTypePtrOutput {
	return pulumi.ToOutput(in).(OSFeedTypePtrOutput)
}

func (in *osfeedTypePtr) ToOSFeedTypePtrOutputWithContext(ctx context.Context) OSFeedTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OSFeedTypePtrOutput)
}

// Regional data boundary for an image
type RegionalDataBoundary string

const (
	// No data boundary
	RegionalDataBoundaryNone = RegionalDataBoundary("None")
	// EU data boundary
	RegionalDataBoundaryEU = RegionalDataBoundary("EU")
)

func (RegionalDataBoundary) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionalDataBoundary)(nil)).Elem()
}

func (e RegionalDataBoundary) ToRegionalDataBoundaryOutput() RegionalDataBoundaryOutput {
	return pulumi.ToOutput(e).(RegionalDataBoundaryOutput)
}

func (e RegionalDataBoundary) ToRegionalDataBoundaryOutputWithContext(ctx context.Context) RegionalDataBoundaryOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RegionalDataBoundaryOutput)
}

func (e RegionalDataBoundary) ToRegionalDataBoundaryPtrOutput() RegionalDataBoundaryPtrOutput {
	return e.ToRegionalDataBoundaryPtrOutputWithContext(context.Background())
}

func (e RegionalDataBoundary) ToRegionalDataBoundaryPtrOutputWithContext(ctx context.Context) RegionalDataBoundaryPtrOutput {
	return RegionalDataBoundary(e).ToRegionalDataBoundaryOutputWithContext(ctx).ToRegionalDataBoundaryPtrOutputWithContext(ctx)
}

func (e RegionalDataBoundary) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RegionalDataBoundary) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RegionalDataBoundary) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RegionalDataBoundary) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RegionalDataBoundaryOutput struct{ *pulumi.OutputState }

func (RegionalDataBoundaryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionalDataBoundary)(nil)).Elem()
}

func (o RegionalDataBoundaryOutput) ToRegionalDataBoundaryOutput() RegionalDataBoundaryOutput {
	return o
}

func (o RegionalDataBoundaryOutput) ToRegionalDataBoundaryOutputWithContext(ctx context.Context) RegionalDataBoundaryOutput {
	return o
}

func (o RegionalDataBoundaryOutput) ToRegionalDataBoundaryPtrOutput() RegionalDataBoundaryPtrOutput {
	return o.ToRegionalDataBoundaryPtrOutputWithContext(context.Background())
}

func (o RegionalDataBoundaryOutput) ToRegionalDataBoundaryPtrOutputWithContext(ctx context.Context) RegionalDataBoundaryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegionalDataBoundary) *RegionalDataBoundary {
		return &v
	}).(RegionalDataBoundaryPtrOutput)
}

func (o RegionalDataBoundaryOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RegionalDataBoundaryOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RegionalDataBoundary) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RegionalDataBoundaryOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RegionalDataBoundaryOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RegionalDataBoundary) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RegionalDataBoundaryPtrOutput struct{ *pulumi.OutputState }

func (RegionalDataBoundaryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionalDataBoundary)(nil)).Elem()
}

func (o RegionalDataBoundaryPtrOutput) ToRegionalDataBoundaryPtrOutput() RegionalDataBoundaryPtrOutput {
	return o
}

func (o RegionalDataBoundaryPtrOutput) ToRegionalDataBoundaryPtrOutputWithContext(ctx context.Context) RegionalDataBoundaryPtrOutput {
	return o
}

func (o RegionalDataBoundaryPtrOutput) Elem() RegionalDataBoundaryOutput {
	return o.ApplyT(func(v *RegionalDataBoundary) RegionalDataBoundary {
		if v != nil {
			return *v
		}
		var ret RegionalDataBoundary
		return ret
	}).(RegionalDataBoundaryOutput)
}

func (o RegionalDataBoundaryPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RegionalDataBoundaryPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RegionalDataBoundary) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// RegionalDataBoundaryInput is an input type that accepts values of the RegionalDataBoundary enum
// A concrete instance of `RegionalDataBoundaryInput` can be one of the following:
//
//	RegionalDataBoundaryNone
//	RegionalDataBoundaryEU
type RegionalDataBoundaryInput interface {
	pulumi.Input

	ToRegionalDataBoundaryOutput() RegionalDataBoundaryOutput
	ToRegionalDataBoundaryOutputWithContext(context.Context) RegionalDataBoundaryOutput
}

var regionalDataBoundaryPtrType = reflect.TypeOf((**RegionalDataBoundary)(nil)).Elem()

type RegionalDataBoundaryPtrInput interface {
	pulumi.Input

	ToRegionalDataBoundaryPtrOutput() RegionalDataBoundaryPtrOutput
	ToRegionalDataBoundaryPtrOutputWithContext(context.Context) RegionalDataBoundaryPtrOutput
}

type regionalDataBoundaryPtr string

func RegionalDataBoundaryPtr(v string) RegionalDataBoundaryPtrInput {
	return (*regionalDataBoundaryPtr)(&v)
}

func (*regionalDataBoundaryPtr) ElementType() reflect.Type {
	return regionalDataBoundaryPtrType
}

func (in *regionalDataBoundaryPtr) ToRegionalDataBoundaryPtrOutput() RegionalDataBoundaryPtrOutput {
	return pulumi.ToOutput(in).(RegionalDataBoundaryPtrOutput)
}

func (in *regionalDataBoundaryPtr) ToRegionalDataBoundaryPtrOutputWithContext(ctx context.Context) RegionalDataBoundaryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RegionalDataBoundaryPtrOutput)
}

// Update policy of the device group.
type UpdatePolicy string

const (
	// Update all policy.
	UpdatePolicyUpdateAll = UpdatePolicy("UpdateAll")
	// No update for 3rd party app policy.
	UpdatePolicyNo3rdPartyAppUpdates = UpdatePolicy("No3rdPartyAppUpdates")
)

func (UpdatePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdatePolicy)(nil)).Elem()
}

func (e UpdatePolicy) ToUpdatePolicyOutput() UpdatePolicyOutput {
	return pulumi.ToOutput(e).(UpdatePolicyOutput)
}

func (e UpdatePolicy) ToUpdatePolicyOutputWithContext(ctx context.Context) UpdatePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UpdatePolicyOutput)
}

func (e UpdatePolicy) ToUpdatePolicyPtrOutput() UpdatePolicyPtrOutput {
	return e.ToUpdatePolicyPtrOutputWithContext(context.Background())
}

func (e UpdatePolicy) ToUpdatePolicyPtrOutputWithContext(ctx context.Context) UpdatePolicyPtrOutput {
	return UpdatePolicy(e).ToUpdatePolicyOutputWithContext(ctx).ToUpdatePolicyPtrOutputWithContext(ctx)
}

func (e UpdatePolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UpdatePolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UpdatePolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UpdatePolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UpdatePolicyOutput struct{ *pulumi.OutputState }

func (UpdatePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdatePolicy)(nil)).Elem()
}

func (o UpdatePolicyOutput) ToUpdatePolicyOutput() UpdatePolicyOutput {
	return o
}

func (o UpdatePolicyOutput) ToUpdatePolicyOutputWithContext(ctx context.Context) UpdatePolicyOutput {
	return o
}

func (o UpdatePolicyOutput) ToUpdatePolicyPtrOutput() UpdatePolicyPtrOutput {
	return o.ToUpdatePolicyPtrOutputWithContext(context.Background())
}

func (o UpdatePolicyOutput) ToUpdatePolicyPtrOutputWithContext(ctx context.Context) UpdatePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UpdatePolicy) *UpdatePolicy {
		return &v
	}).(UpdatePolicyPtrOutput)
}

func (o UpdatePolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UpdatePolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UpdatePolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UpdatePolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UpdatePolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UpdatePolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UpdatePolicyPtrOutput struct{ *pulumi.OutputState }

func (UpdatePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdatePolicy)(nil)).Elem()
}

func (o UpdatePolicyPtrOutput) ToUpdatePolicyPtrOutput() UpdatePolicyPtrOutput {
	return o
}

func (o UpdatePolicyPtrOutput) ToUpdatePolicyPtrOutputWithContext(ctx context.Context) UpdatePolicyPtrOutput {
	return o
}

func (o UpdatePolicyPtrOutput) Elem() UpdatePolicyOutput {
	return o.ApplyT(func(v *UpdatePolicy) UpdatePolicy {
		if v != nil {
			return *v
		}
		var ret UpdatePolicy
		return ret
	}).(UpdatePolicyOutput)
}

func (o UpdatePolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UpdatePolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UpdatePolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// UpdatePolicyInput is an input type that accepts values of the UpdatePolicy enum
// A concrete instance of `UpdatePolicyInput` can be one of the following:
//
//	UpdatePolicyUpdateAll
//	UpdatePolicyNo3rdPartyAppUpdates
type UpdatePolicyInput interface {
	pulumi.Input

	ToUpdatePolicyOutput() UpdatePolicyOutput
	ToUpdatePolicyOutputWithContext(context.Context) UpdatePolicyOutput
}

var updatePolicyPtrType = reflect.TypeOf((**UpdatePolicy)(nil)).Elem()

type UpdatePolicyPtrInput interface {
	pulumi.Input

	ToUpdatePolicyPtrOutput() UpdatePolicyPtrOutput
	ToUpdatePolicyPtrOutputWithContext(context.Context) UpdatePolicyPtrOutput
}

type updatePolicyPtr string

func UpdatePolicyPtr(v string) UpdatePolicyPtrInput {
	return (*updatePolicyPtr)(&v)
}

func (*updatePolicyPtr) ElementType() reflect.Type {
	return updatePolicyPtrType
}

func (in *updatePolicyPtr) ToUpdatePolicyPtrOutput() UpdatePolicyPtrOutput {
	return pulumi.ToOutput(in).(UpdatePolicyPtrOutput)
}

func (in *updatePolicyPtr) ToUpdatePolicyPtrOutputWithContext(ctx context.Context) UpdatePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UpdatePolicyPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AllowCrashDumpCollectionOutput{})
	pulumi.RegisterOutputType(AllowCrashDumpCollectionPtrOutput{})
	pulumi.RegisterOutputType(OSFeedTypeOutput{})
	pulumi.RegisterOutputType(OSFeedTypePtrOutput{})
	pulumi.RegisterOutputType(RegionalDataBoundaryOutput{})
	pulumi.RegisterOutputType(RegionalDataBoundaryPtrOutput{})
	pulumi.RegisterOutputType(UpdatePolicyOutput{})
	pulumi.RegisterOutputType(UpdatePolicyPtrOutput{})
}
