// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// The supported disk size details for a disk type.
type DiskDetailsResponse struct {
	// The disk tier, e.g. P10, E10.
	DiskTier *string `pulumi:"diskTier"`
	// The disk Iops.
	IopsReadWrite *float64 `pulumi:"iopsReadWrite"`
	// The maximum supported disk count.
	MaximumSupportedDiskCount *float64 `pulumi:"maximumSupportedDiskCount"`
	// The disk provisioned throughput in MBps.
	MbpsReadWrite *float64 `pulumi:"mbpsReadWrite"`
	// The minimum supported disk count.
	MinimumSupportedDiskCount *float64 `pulumi:"minimumSupportedDiskCount"`
	// The disk size in GB.
	SizeGB *float64 `pulumi:"sizeGB"`
	// The disk sku.
	Sku *DiskSkuResponse `pulumi:"sku"`
}

// The supported disk size details for a disk type.
type DiskDetailsResponseOutput struct{ *pulumi.OutputState }

func (DiskDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskDetailsResponse)(nil)).Elem()
}

func (o DiskDetailsResponseOutput) ToDiskDetailsResponseOutput() DiskDetailsResponseOutput {
	return o
}

func (o DiskDetailsResponseOutput) ToDiskDetailsResponseOutputWithContext(ctx context.Context) DiskDetailsResponseOutput {
	return o
}

func (o DiskDetailsResponseOutput) ToOutput(ctx context.Context) pulumix.Output[DiskDetailsResponse] {
	return pulumix.Output[DiskDetailsResponse]{
		OutputState: o.OutputState,
	}
}

// The disk tier, e.g. P10, E10.
func (o DiskDetailsResponseOutput) DiskTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskDetailsResponse) *string { return v.DiskTier }).(pulumi.StringPtrOutput)
}

// The disk Iops.
func (o DiskDetailsResponseOutput) IopsReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DiskDetailsResponse) *float64 { return v.IopsReadWrite }).(pulumi.Float64PtrOutput)
}

// The maximum supported disk count.
func (o DiskDetailsResponseOutput) MaximumSupportedDiskCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DiskDetailsResponse) *float64 { return v.MaximumSupportedDiskCount }).(pulumi.Float64PtrOutput)
}

// The disk provisioned throughput in MBps.
func (o DiskDetailsResponseOutput) MbpsReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DiskDetailsResponse) *float64 { return v.MbpsReadWrite }).(pulumi.Float64PtrOutput)
}

// The minimum supported disk count.
func (o DiskDetailsResponseOutput) MinimumSupportedDiskCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DiskDetailsResponse) *float64 { return v.MinimumSupportedDiskCount }).(pulumi.Float64PtrOutput)
}

// The disk size in GB.
func (o DiskDetailsResponseOutput) SizeGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DiskDetailsResponse) *float64 { return v.SizeGB }).(pulumi.Float64PtrOutput)
}

// The disk sku.
func (o DiskDetailsResponseOutput) Sku() DiskSkuResponsePtrOutput {
	return o.ApplyT(func(v DiskDetailsResponse) *DiskSkuResponse { return v.Sku }).(DiskSkuResponsePtrOutput)
}

type DiskDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskDetailsResponse)(nil)).Elem()
}

func (o DiskDetailsResponseArrayOutput) ToDiskDetailsResponseArrayOutput() DiskDetailsResponseArrayOutput {
	return o
}

func (o DiskDetailsResponseArrayOutput) ToDiskDetailsResponseArrayOutputWithContext(ctx context.Context) DiskDetailsResponseArrayOutput {
	return o
}

func (o DiskDetailsResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]DiskDetailsResponse] {
	return pulumix.Output[[]DiskDetailsResponse]{
		OutputState: o.OutputState,
	}
}

func (o DiskDetailsResponseArrayOutput) Index(i pulumi.IntInput) DiskDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskDetailsResponse {
		return vs[0].([]DiskDetailsResponse)[vs[1].(int)]
	}).(DiskDetailsResponseOutput)
}

// The disk sku.
type DiskSkuResponse struct {
	// Defines the disk sku name.
	Name *string `pulumi:"name"`
}

// The disk sku.
type DiskSkuResponseOutput struct{ *pulumi.OutputState }

func (DiskSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSkuResponse)(nil)).Elem()
}

func (o DiskSkuResponseOutput) ToDiskSkuResponseOutput() DiskSkuResponseOutput {
	return o
}

func (o DiskSkuResponseOutput) ToDiskSkuResponseOutputWithContext(ctx context.Context) DiskSkuResponseOutput {
	return o
}

func (o DiskSkuResponseOutput) ToOutput(ctx context.Context) pulumix.Output[DiskSkuResponse] {
	return pulumix.Output[DiskSkuResponse]{
		OutputState: o.OutputState,
	}
}

// Defines the disk sku name.
func (o DiskSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DiskSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskSkuResponse)(nil)).Elem()
}

func (o DiskSkuResponsePtrOutput) ToDiskSkuResponsePtrOutput() DiskSkuResponsePtrOutput {
	return o
}

func (o DiskSkuResponsePtrOutput) ToDiskSkuResponsePtrOutputWithContext(ctx context.Context) DiskSkuResponsePtrOutput {
	return o
}

func (o DiskSkuResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*DiskSkuResponse] {
	return pulumix.Output[*DiskSkuResponse]{
		OutputState: o.OutputState,
	}
}

func (o DiskSkuResponsePtrOutput) Elem() DiskSkuResponseOutput {
	return o.ApplyT(func(v *DiskSkuResponse) DiskSkuResponse {
		if v != nil {
			return *v
		}
		var ret DiskSkuResponse
		return ret
	}).(DiskSkuResponseOutput)
}

// Defines the disk sku name.
func (o DiskSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// The disk configuration required for the selected volume.
type DiskVolumeConfigurationResponse struct {
	// The total number of disks required for the concerned volume.
	Count *float64 `pulumi:"count"`
	// The disk size in GB.
	SizeGB *float64 `pulumi:"sizeGB"`
	// The disk SKU details.
	Sku *DiskSkuResponse `pulumi:"sku"`
}

// The disk configuration required for the selected volume.
type DiskVolumeConfigurationResponseOutput struct{ *pulumi.OutputState }

func (DiskVolumeConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskVolumeConfigurationResponse)(nil)).Elem()
}

func (o DiskVolumeConfigurationResponseOutput) ToDiskVolumeConfigurationResponseOutput() DiskVolumeConfigurationResponseOutput {
	return o
}

func (o DiskVolumeConfigurationResponseOutput) ToDiskVolumeConfigurationResponseOutputWithContext(ctx context.Context) DiskVolumeConfigurationResponseOutput {
	return o
}

func (o DiskVolumeConfigurationResponseOutput) ToOutput(ctx context.Context) pulumix.Output[DiskVolumeConfigurationResponse] {
	return pulumix.Output[DiskVolumeConfigurationResponse]{
		OutputState: o.OutputState,
	}
}

// The total number of disks required for the concerned volume.
func (o DiskVolumeConfigurationResponseOutput) Count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DiskVolumeConfigurationResponse) *float64 { return v.Count }).(pulumi.Float64PtrOutput)
}

// The disk size in GB.
func (o DiskVolumeConfigurationResponseOutput) SizeGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DiskVolumeConfigurationResponse) *float64 { return v.SizeGB }).(pulumi.Float64PtrOutput)
}

// The disk SKU details.
func (o DiskVolumeConfigurationResponseOutput) Sku() DiskSkuResponsePtrOutput {
	return o.ApplyT(func(v DiskVolumeConfigurationResponse) *DiskSkuResponse { return v.Sku }).(DiskSkuResponsePtrOutput)
}

type DiskVolumeConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskVolumeConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskVolumeConfigurationResponse)(nil)).Elem()
}

func (o DiskVolumeConfigurationResponsePtrOutput) ToDiskVolumeConfigurationResponsePtrOutput() DiskVolumeConfigurationResponsePtrOutput {
	return o
}

func (o DiskVolumeConfigurationResponsePtrOutput) ToDiskVolumeConfigurationResponsePtrOutputWithContext(ctx context.Context) DiskVolumeConfigurationResponsePtrOutput {
	return o
}

func (o DiskVolumeConfigurationResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*DiskVolumeConfigurationResponse] {
	return pulumix.Output[*DiskVolumeConfigurationResponse]{
		OutputState: o.OutputState,
	}
}

func (o DiskVolumeConfigurationResponsePtrOutput) Elem() DiskVolumeConfigurationResponseOutput {
	return o.ApplyT(func(v *DiskVolumeConfigurationResponse) DiskVolumeConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret DiskVolumeConfigurationResponse
		return ret
	}).(DiskVolumeConfigurationResponseOutput)
}

// The total number of disks required for the concerned volume.
func (o DiskVolumeConfigurationResponsePtrOutput) Count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DiskVolumeConfigurationResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.Float64PtrOutput)
}

// The disk size in GB.
func (o DiskVolumeConfigurationResponsePtrOutput) SizeGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DiskVolumeConfigurationResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SizeGB
	}).(pulumi.Float64PtrOutput)
}

// The disk SKU details.
func (o DiskVolumeConfigurationResponsePtrOutput) Sku() DiskSkuResponsePtrOutput {
	return o.ApplyT(func(v *DiskVolumeConfigurationResponse) *DiskSkuResponse {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(DiskSkuResponsePtrOutput)
}

// The SAP Availability Zone Pair.
type SAPAvailabilityZonePairResponse struct {
	// The zone A.
	ZoneA *float64 `pulumi:"zoneA"`
	// The zone B.
	ZoneB *float64 `pulumi:"zoneB"`
}

// The SAP Availability Zone Pair.
type SAPAvailabilityZonePairResponseOutput struct{ *pulumi.OutputState }

func (SAPAvailabilityZonePairResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SAPAvailabilityZonePairResponse)(nil)).Elem()
}

func (o SAPAvailabilityZonePairResponseOutput) ToSAPAvailabilityZonePairResponseOutput() SAPAvailabilityZonePairResponseOutput {
	return o
}

func (o SAPAvailabilityZonePairResponseOutput) ToSAPAvailabilityZonePairResponseOutputWithContext(ctx context.Context) SAPAvailabilityZonePairResponseOutput {
	return o
}

func (o SAPAvailabilityZonePairResponseOutput) ToOutput(ctx context.Context) pulumix.Output[SAPAvailabilityZonePairResponse] {
	return pulumix.Output[SAPAvailabilityZonePairResponse]{
		OutputState: o.OutputState,
	}
}

// The zone A.
func (o SAPAvailabilityZonePairResponseOutput) ZoneA() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SAPAvailabilityZonePairResponse) *float64 { return v.ZoneA }).(pulumi.Float64PtrOutput)
}

// The zone B.
func (o SAPAvailabilityZonePairResponseOutput) ZoneB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SAPAvailabilityZonePairResponse) *float64 { return v.ZoneB }).(pulumi.Float64PtrOutput)
}

type SAPAvailabilityZonePairResponseArrayOutput struct{ *pulumi.OutputState }

func (SAPAvailabilityZonePairResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SAPAvailabilityZonePairResponse)(nil)).Elem()
}

func (o SAPAvailabilityZonePairResponseArrayOutput) ToSAPAvailabilityZonePairResponseArrayOutput() SAPAvailabilityZonePairResponseArrayOutput {
	return o
}

func (o SAPAvailabilityZonePairResponseArrayOutput) ToSAPAvailabilityZonePairResponseArrayOutputWithContext(ctx context.Context) SAPAvailabilityZonePairResponseArrayOutput {
	return o
}

func (o SAPAvailabilityZonePairResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]SAPAvailabilityZonePairResponse] {
	return pulumix.Output[[]SAPAvailabilityZonePairResponse]{
		OutputState: o.OutputState,
	}
}

func (o SAPAvailabilityZonePairResponseArrayOutput) Index(i pulumi.IntInput) SAPAvailabilityZonePairResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SAPAvailabilityZonePairResponse {
		return vs[0].([]SAPAvailabilityZonePairResponse)[vs[1].(int)]
	}).(SAPAvailabilityZonePairResponseOutput)
}

// The SAP Disk Configuration contains 'recommended disk' details and list of supported disks detail for a volume type.
type SAPDiskConfigurationResponse struct {
	// The recommended disk details for a given VM Sku.
	RecommendedConfiguration *DiskVolumeConfigurationResponse `pulumi:"recommendedConfiguration"`
	// The list of supported disks for a given VM Sku.
	SupportedConfigurations []DiskDetailsResponse `pulumi:"supportedConfigurations"`
}

// The SAP Disk Configuration contains 'recommended disk' details and list of supported disks detail for a volume type.
type SAPDiskConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SAPDiskConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SAPDiskConfigurationResponse)(nil)).Elem()
}

func (o SAPDiskConfigurationResponseOutput) ToSAPDiskConfigurationResponseOutput() SAPDiskConfigurationResponseOutput {
	return o
}

func (o SAPDiskConfigurationResponseOutput) ToSAPDiskConfigurationResponseOutputWithContext(ctx context.Context) SAPDiskConfigurationResponseOutput {
	return o
}

func (o SAPDiskConfigurationResponseOutput) ToOutput(ctx context.Context) pulumix.Output[SAPDiskConfigurationResponse] {
	return pulumix.Output[SAPDiskConfigurationResponse]{
		OutputState: o.OutputState,
	}
}

// The recommended disk details for a given VM Sku.
func (o SAPDiskConfigurationResponseOutput) RecommendedConfiguration() DiskVolumeConfigurationResponsePtrOutput {
	return o.ApplyT(func(v SAPDiskConfigurationResponse) *DiskVolumeConfigurationResponse {
		return v.RecommendedConfiguration
	}).(DiskVolumeConfigurationResponsePtrOutput)
}

// The list of supported disks for a given VM Sku.
func (o SAPDiskConfigurationResponseOutput) SupportedConfigurations() DiskDetailsResponseArrayOutput {
	return o.ApplyT(func(v SAPDiskConfigurationResponse) []DiskDetailsResponse { return v.SupportedConfigurations }).(DiskDetailsResponseArrayOutput)
}

type SAPDiskConfigurationResponseMapOutput struct{ *pulumi.OutputState }

func (SAPDiskConfigurationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SAPDiskConfigurationResponse)(nil)).Elem()
}

func (o SAPDiskConfigurationResponseMapOutput) ToSAPDiskConfigurationResponseMapOutput() SAPDiskConfigurationResponseMapOutput {
	return o
}

func (o SAPDiskConfigurationResponseMapOutput) ToSAPDiskConfigurationResponseMapOutputWithContext(ctx context.Context) SAPDiskConfigurationResponseMapOutput {
	return o
}

func (o SAPDiskConfigurationResponseMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]SAPDiskConfigurationResponse] {
	return pulumix.Output[map[string]SAPDiskConfigurationResponse]{
		OutputState: o.OutputState,
	}
}

func (o SAPDiskConfigurationResponseMapOutput) MapIndex(k pulumi.StringInput) SAPDiskConfigurationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SAPDiskConfigurationResponse {
		return vs[0].(map[string]SAPDiskConfigurationResponse)[vs[1].(string)]
	}).(SAPDiskConfigurationResponseOutput)
}

// The SAP supported SKU.
type SAPSupportedSkuResponse struct {
	// True if the Sku is certified for App server in the SAP system.
	IsAppServerCertified *bool `pulumi:"isAppServerCertified"`
	// True if the Sku is certified for Database server in the SAP system.
	IsDatabaseCertified *bool `pulumi:"isDatabaseCertified"`
	// The VM Sku.
	VmSku *string `pulumi:"vmSku"`
}

// The SAP supported SKU.
type SAPSupportedSkuResponseOutput struct{ *pulumi.OutputState }

func (SAPSupportedSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SAPSupportedSkuResponse)(nil)).Elem()
}

func (o SAPSupportedSkuResponseOutput) ToSAPSupportedSkuResponseOutput() SAPSupportedSkuResponseOutput {
	return o
}

func (o SAPSupportedSkuResponseOutput) ToSAPSupportedSkuResponseOutputWithContext(ctx context.Context) SAPSupportedSkuResponseOutput {
	return o
}

func (o SAPSupportedSkuResponseOutput) ToOutput(ctx context.Context) pulumix.Output[SAPSupportedSkuResponse] {
	return pulumix.Output[SAPSupportedSkuResponse]{
		OutputState: o.OutputState,
	}
}

// True if the Sku is certified for App server in the SAP system.
func (o SAPSupportedSkuResponseOutput) IsAppServerCertified() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SAPSupportedSkuResponse) *bool { return v.IsAppServerCertified }).(pulumi.BoolPtrOutput)
}

// True if the Sku is certified for Database server in the SAP system.
func (o SAPSupportedSkuResponseOutput) IsDatabaseCertified() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SAPSupportedSkuResponse) *bool { return v.IsDatabaseCertified }).(pulumi.BoolPtrOutput)
}

// The VM Sku.
func (o SAPSupportedSkuResponseOutput) VmSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SAPSupportedSkuResponse) *string { return v.VmSku }).(pulumi.StringPtrOutput)
}

type SAPSupportedSkuResponseArrayOutput struct{ *pulumi.OutputState }

func (SAPSupportedSkuResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SAPSupportedSkuResponse)(nil)).Elem()
}

func (o SAPSupportedSkuResponseArrayOutput) ToSAPSupportedSkuResponseArrayOutput() SAPSupportedSkuResponseArrayOutput {
	return o
}

func (o SAPSupportedSkuResponseArrayOutput) ToSAPSupportedSkuResponseArrayOutputWithContext(ctx context.Context) SAPSupportedSkuResponseArrayOutput {
	return o
}

func (o SAPSupportedSkuResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]SAPSupportedSkuResponse] {
	return pulumix.Output[[]SAPSupportedSkuResponse]{
		OutputState: o.OutputState,
	}
}

func (o SAPSupportedSkuResponseArrayOutput) Index(i pulumi.IntInput) SAPSupportedSkuResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SAPSupportedSkuResponse {
		return vs[0].([]SAPSupportedSkuResponse)[vs[1].(int)]
	}).(SAPSupportedSkuResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(DiskDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(DiskSkuResponseOutput{})
	pulumi.RegisterOutputType(DiskSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(DiskVolumeConfigurationResponseOutput{})
	pulumi.RegisterOutputType(DiskVolumeConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SAPAvailabilityZonePairResponseOutput{})
	pulumi.RegisterOutputType(SAPAvailabilityZonePairResponseArrayOutput{})
	pulumi.RegisterOutputType(SAPDiskConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SAPDiskConfigurationResponseMapOutput{})
	pulumi.RegisterOutputType(SAPSupportedSkuResponseOutput{})
	pulumi.RegisterOutputType(SAPSupportedSkuResponseArrayOutput{})
}
