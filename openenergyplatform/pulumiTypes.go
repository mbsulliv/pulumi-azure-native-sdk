// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openenergyplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// The list of Energy services resource's Data Partition Names.
type DataPartitionNames struct {
	Name *string `pulumi:"name"`
}

// DataPartitionNamesInput is an input type that accepts DataPartitionNamesArgs and DataPartitionNamesOutput values.
// You can construct a concrete instance of `DataPartitionNamesInput` via:
//
//	DataPartitionNamesArgs{...}
type DataPartitionNamesInput interface {
	pulumi.Input

	ToDataPartitionNamesOutput() DataPartitionNamesOutput
	ToDataPartitionNamesOutputWithContext(context.Context) DataPartitionNamesOutput
}

// The list of Energy services resource's Data Partition Names.
type DataPartitionNamesArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DataPartitionNamesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPartitionNames)(nil)).Elem()
}

func (i DataPartitionNamesArgs) ToDataPartitionNamesOutput() DataPartitionNamesOutput {
	return i.ToDataPartitionNamesOutputWithContext(context.Background())
}

func (i DataPartitionNamesArgs) ToDataPartitionNamesOutputWithContext(ctx context.Context) DataPartitionNamesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPartitionNamesOutput)
}

func (i DataPartitionNamesArgs) ToOutput(ctx context.Context) pulumix.Output[DataPartitionNames] {
	return pulumix.Output[DataPartitionNames]{
		OutputState: i.ToDataPartitionNamesOutputWithContext(ctx).OutputState,
	}
}

// DataPartitionNamesArrayInput is an input type that accepts DataPartitionNamesArray and DataPartitionNamesArrayOutput values.
// You can construct a concrete instance of `DataPartitionNamesArrayInput` via:
//
//	DataPartitionNamesArray{ DataPartitionNamesArgs{...} }
type DataPartitionNamesArrayInput interface {
	pulumi.Input

	ToDataPartitionNamesArrayOutput() DataPartitionNamesArrayOutput
	ToDataPartitionNamesArrayOutputWithContext(context.Context) DataPartitionNamesArrayOutput
}

type DataPartitionNamesArray []DataPartitionNamesInput

func (DataPartitionNamesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataPartitionNames)(nil)).Elem()
}

func (i DataPartitionNamesArray) ToDataPartitionNamesArrayOutput() DataPartitionNamesArrayOutput {
	return i.ToDataPartitionNamesArrayOutputWithContext(context.Background())
}

func (i DataPartitionNamesArray) ToDataPartitionNamesArrayOutputWithContext(ctx context.Context) DataPartitionNamesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPartitionNamesArrayOutput)
}

func (i DataPartitionNamesArray) ToOutput(ctx context.Context) pulumix.Output[[]DataPartitionNames] {
	return pulumix.Output[[]DataPartitionNames]{
		OutputState: i.ToDataPartitionNamesArrayOutputWithContext(ctx).OutputState,
	}
}

// The list of Energy services resource's Data Partition Names.
type DataPartitionNamesOutput struct{ *pulumi.OutputState }

func (DataPartitionNamesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPartitionNames)(nil)).Elem()
}

func (o DataPartitionNamesOutput) ToDataPartitionNamesOutput() DataPartitionNamesOutput {
	return o
}

func (o DataPartitionNamesOutput) ToDataPartitionNamesOutputWithContext(ctx context.Context) DataPartitionNamesOutput {
	return o
}

func (o DataPartitionNamesOutput) ToOutput(ctx context.Context) pulumix.Output[DataPartitionNames] {
	return pulumix.Output[DataPartitionNames]{
		OutputState: o.OutputState,
	}
}

func (o DataPartitionNamesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataPartitionNames) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DataPartitionNamesArrayOutput struct{ *pulumi.OutputState }

func (DataPartitionNamesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataPartitionNames)(nil)).Elem()
}

func (o DataPartitionNamesArrayOutput) ToDataPartitionNamesArrayOutput() DataPartitionNamesArrayOutput {
	return o
}

func (o DataPartitionNamesArrayOutput) ToDataPartitionNamesArrayOutputWithContext(ctx context.Context) DataPartitionNamesArrayOutput {
	return o
}

func (o DataPartitionNamesArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]DataPartitionNames] {
	return pulumix.Output[[]DataPartitionNames]{
		OutputState: o.OutputState,
	}
}

func (o DataPartitionNamesArrayOutput) Index(i pulumi.IntInput) DataPartitionNamesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataPartitionNames {
		return vs[0].([]DataPartitionNames)[vs[1].(int)]
	}).(DataPartitionNamesOutput)
}

// The list of Energy services resource's Data Partition Names.
type DataPartitionNamesResponse struct {
	Name *string `pulumi:"name"`
}

// The list of Energy services resource's Data Partition Names.
type DataPartitionNamesResponseOutput struct{ *pulumi.OutputState }

func (DataPartitionNamesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPartitionNamesResponse)(nil)).Elem()
}

func (o DataPartitionNamesResponseOutput) ToDataPartitionNamesResponseOutput() DataPartitionNamesResponseOutput {
	return o
}

func (o DataPartitionNamesResponseOutput) ToDataPartitionNamesResponseOutputWithContext(ctx context.Context) DataPartitionNamesResponseOutput {
	return o
}

func (o DataPartitionNamesResponseOutput) ToOutput(ctx context.Context) pulumix.Output[DataPartitionNamesResponse] {
	return pulumix.Output[DataPartitionNamesResponse]{
		OutputState: o.OutputState,
	}
}

func (o DataPartitionNamesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataPartitionNamesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DataPartitionNamesResponseArrayOutput struct{ *pulumi.OutputState }

func (DataPartitionNamesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataPartitionNamesResponse)(nil)).Elem()
}

func (o DataPartitionNamesResponseArrayOutput) ToDataPartitionNamesResponseArrayOutput() DataPartitionNamesResponseArrayOutput {
	return o
}

func (o DataPartitionNamesResponseArrayOutput) ToDataPartitionNamesResponseArrayOutputWithContext(ctx context.Context) DataPartitionNamesResponseArrayOutput {
	return o
}

func (o DataPartitionNamesResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]DataPartitionNamesResponse] {
	return pulumix.Output[[]DataPartitionNamesResponse]{
		OutputState: o.OutputState,
	}
}

func (o DataPartitionNamesResponseArrayOutput) Index(i pulumi.IntInput) DataPartitionNamesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataPartitionNamesResponse {
		return vs[0].([]DataPartitionNamesResponse)[vs[1].(int)]
	}).(DataPartitionNamesResponseOutput)
}

// Defines the properties of an individual data partition.
type DataPartitionPropertiesResponse struct {
	// Name of the data partition
	Name *string `pulumi:"name"`
	// Name of the data partition
	ProvisioningState *string `pulumi:"provisioningState"`
}

// Defines the properties of an individual data partition.
type DataPartitionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DataPartitionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPartitionPropertiesResponse)(nil)).Elem()
}

func (o DataPartitionPropertiesResponseOutput) ToDataPartitionPropertiesResponseOutput() DataPartitionPropertiesResponseOutput {
	return o
}

func (o DataPartitionPropertiesResponseOutput) ToDataPartitionPropertiesResponseOutputWithContext(ctx context.Context) DataPartitionPropertiesResponseOutput {
	return o
}

func (o DataPartitionPropertiesResponseOutput) ToOutput(ctx context.Context) pulumix.Output[DataPartitionPropertiesResponse] {
	return pulumix.Output[DataPartitionPropertiesResponse]{
		OutputState: o.OutputState,
	}
}

// Name of the data partition
func (o DataPartitionPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataPartitionPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Name of the data partition
func (o DataPartitionPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataPartitionPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type DataPartitionPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (DataPartitionPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataPartitionPropertiesResponse)(nil)).Elem()
}

func (o DataPartitionPropertiesResponseArrayOutput) ToDataPartitionPropertiesResponseArrayOutput() DataPartitionPropertiesResponseArrayOutput {
	return o
}

func (o DataPartitionPropertiesResponseArrayOutput) ToDataPartitionPropertiesResponseArrayOutputWithContext(ctx context.Context) DataPartitionPropertiesResponseArrayOutput {
	return o
}

func (o DataPartitionPropertiesResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]DataPartitionPropertiesResponse] {
	return pulumix.Output[[]DataPartitionPropertiesResponse]{
		OutputState: o.OutputState,
	}
}

func (o DataPartitionPropertiesResponseArrayOutput) Index(i pulumi.IntInput) DataPartitionPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataPartitionPropertiesResponse {
		return vs[0].([]DataPartitionPropertiesResponse)[vs[1].(int)]
	}).(DataPartitionPropertiesResponseOutput)
}

type EnergyServiceProperties struct {
	AuthAppId          *string              `pulumi:"authAppId"`
	DataPartitionNames []DataPartitionNames `pulumi:"dataPartitionNames"`
}

// EnergyServicePropertiesInput is an input type that accepts EnergyServicePropertiesArgs and EnergyServicePropertiesOutput values.
// You can construct a concrete instance of `EnergyServicePropertiesInput` via:
//
//	EnergyServicePropertiesArgs{...}
type EnergyServicePropertiesInput interface {
	pulumi.Input

	ToEnergyServicePropertiesOutput() EnergyServicePropertiesOutput
	ToEnergyServicePropertiesOutputWithContext(context.Context) EnergyServicePropertiesOutput
}

type EnergyServicePropertiesArgs struct {
	AuthAppId          pulumi.StringPtrInput        `pulumi:"authAppId"`
	DataPartitionNames DataPartitionNamesArrayInput `pulumi:"dataPartitionNames"`
}

func (EnergyServicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnergyServiceProperties)(nil)).Elem()
}

func (i EnergyServicePropertiesArgs) ToEnergyServicePropertiesOutput() EnergyServicePropertiesOutput {
	return i.ToEnergyServicePropertiesOutputWithContext(context.Background())
}

func (i EnergyServicePropertiesArgs) ToEnergyServicePropertiesOutputWithContext(ctx context.Context) EnergyServicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnergyServicePropertiesOutput)
}

func (i EnergyServicePropertiesArgs) ToOutput(ctx context.Context) pulumix.Output[EnergyServiceProperties] {
	return pulumix.Output[EnergyServiceProperties]{
		OutputState: i.ToEnergyServicePropertiesOutputWithContext(ctx).OutputState,
	}
}

func (i EnergyServicePropertiesArgs) ToEnergyServicePropertiesPtrOutput() EnergyServicePropertiesPtrOutput {
	return i.ToEnergyServicePropertiesPtrOutputWithContext(context.Background())
}

func (i EnergyServicePropertiesArgs) ToEnergyServicePropertiesPtrOutputWithContext(ctx context.Context) EnergyServicePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnergyServicePropertiesOutput).ToEnergyServicePropertiesPtrOutputWithContext(ctx)
}

// EnergyServicePropertiesPtrInput is an input type that accepts EnergyServicePropertiesArgs, EnergyServicePropertiesPtr and EnergyServicePropertiesPtrOutput values.
// You can construct a concrete instance of `EnergyServicePropertiesPtrInput` via:
//
//	        EnergyServicePropertiesArgs{...}
//
//	or:
//
//	        nil
type EnergyServicePropertiesPtrInput interface {
	pulumi.Input

	ToEnergyServicePropertiesPtrOutput() EnergyServicePropertiesPtrOutput
	ToEnergyServicePropertiesPtrOutputWithContext(context.Context) EnergyServicePropertiesPtrOutput
}

type energyServicePropertiesPtrType EnergyServicePropertiesArgs

func EnergyServicePropertiesPtr(v *EnergyServicePropertiesArgs) EnergyServicePropertiesPtrInput {
	return (*energyServicePropertiesPtrType)(v)
}

func (*energyServicePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnergyServiceProperties)(nil)).Elem()
}

func (i *energyServicePropertiesPtrType) ToEnergyServicePropertiesPtrOutput() EnergyServicePropertiesPtrOutput {
	return i.ToEnergyServicePropertiesPtrOutputWithContext(context.Background())
}

func (i *energyServicePropertiesPtrType) ToEnergyServicePropertiesPtrOutputWithContext(ctx context.Context) EnergyServicePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnergyServicePropertiesPtrOutput)
}

func (i *energyServicePropertiesPtrType) ToOutput(ctx context.Context) pulumix.Output[*EnergyServiceProperties] {
	return pulumix.Output[*EnergyServiceProperties]{
		OutputState: i.ToEnergyServicePropertiesPtrOutputWithContext(ctx).OutputState,
	}
}

type EnergyServicePropertiesOutput struct{ *pulumi.OutputState }

func (EnergyServicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnergyServiceProperties)(nil)).Elem()
}

func (o EnergyServicePropertiesOutput) ToEnergyServicePropertiesOutput() EnergyServicePropertiesOutput {
	return o
}

func (o EnergyServicePropertiesOutput) ToEnergyServicePropertiesOutputWithContext(ctx context.Context) EnergyServicePropertiesOutput {
	return o
}

func (o EnergyServicePropertiesOutput) ToEnergyServicePropertiesPtrOutput() EnergyServicePropertiesPtrOutput {
	return o.ToEnergyServicePropertiesPtrOutputWithContext(context.Background())
}

func (o EnergyServicePropertiesOutput) ToEnergyServicePropertiesPtrOutputWithContext(ctx context.Context) EnergyServicePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnergyServiceProperties) *EnergyServiceProperties {
		return &v
	}).(EnergyServicePropertiesPtrOutput)
}

func (o EnergyServicePropertiesOutput) ToOutput(ctx context.Context) pulumix.Output[EnergyServiceProperties] {
	return pulumix.Output[EnergyServiceProperties]{
		OutputState: o.OutputState,
	}
}

func (o EnergyServicePropertiesOutput) AuthAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnergyServiceProperties) *string { return v.AuthAppId }).(pulumi.StringPtrOutput)
}

func (o EnergyServicePropertiesOutput) DataPartitionNames() DataPartitionNamesArrayOutput {
	return o.ApplyT(func(v EnergyServiceProperties) []DataPartitionNames { return v.DataPartitionNames }).(DataPartitionNamesArrayOutput)
}

type EnergyServicePropertiesPtrOutput struct{ *pulumi.OutputState }

func (EnergyServicePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnergyServiceProperties)(nil)).Elem()
}

func (o EnergyServicePropertiesPtrOutput) ToEnergyServicePropertiesPtrOutput() EnergyServicePropertiesPtrOutput {
	return o
}

func (o EnergyServicePropertiesPtrOutput) ToEnergyServicePropertiesPtrOutputWithContext(ctx context.Context) EnergyServicePropertiesPtrOutput {
	return o
}

func (o EnergyServicePropertiesPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*EnergyServiceProperties] {
	return pulumix.Output[*EnergyServiceProperties]{
		OutputState: o.OutputState,
	}
}

func (o EnergyServicePropertiesPtrOutput) Elem() EnergyServicePropertiesOutput {
	return o.ApplyT(func(v *EnergyServiceProperties) EnergyServiceProperties {
		if v != nil {
			return *v
		}
		var ret EnergyServiceProperties
		return ret
	}).(EnergyServicePropertiesOutput)
}

func (o EnergyServicePropertiesPtrOutput) AuthAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnergyServiceProperties) *string {
		if v == nil {
			return nil
		}
		return v.AuthAppId
	}).(pulumi.StringPtrOutput)
}

func (o EnergyServicePropertiesPtrOutput) DataPartitionNames() DataPartitionNamesArrayOutput {
	return o.ApplyT(func(v *EnergyServiceProperties) []DataPartitionNames {
		if v == nil {
			return nil
		}
		return v.DataPartitionNames
	}).(DataPartitionNamesArrayOutput)
}

type EnergyServicePropertiesResponse struct {
	AuthAppId          *string                      `pulumi:"authAppId"`
	DataPartitionNames []DataPartitionNamesResponse `pulumi:"dataPartitionNames"`
	DnsName            string                       `pulumi:"dnsName"`
	ProvisioningState  string                       `pulumi:"provisioningState"`
}

type EnergyServicePropertiesResponseOutput struct{ *pulumi.OutputState }

func (EnergyServicePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnergyServicePropertiesResponse)(nil)).Elem()
}

func (o EnergyServicePropertiesResponseOutput) ToEnergyServicePropertiesResponseOutput() EnergyServicePropertiesResponseOutput {
	return o
}

func (o EnergyServicePropertiesResponseOutput) ToEnergyServicePropertiesResponseOutputWithContext(ctx context.Context) EnergyServicePropertiesResponseOutput {
	return o
}

func (o EnergyServicePropertiesResponseOutput) ToOutput(ctx context.Context) pulumix.Output[EnergyServicePropertiesResponse] {
	return pulumix.Output[EnergyServicePropertiesResponse]{
		OutputState: o.OutputState,
	}
}

func (o EnergyServicePropertiesResponseOutput) AuthAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnergyServicePropertiesResponse) *string { return v.AuthAppId }).(pulumi.StringPtrOutput)
}

func (o EnergyServicePropertiesResponseOutput) DataPartitionNames() DataPartitionNamesResponseArrayOutput {
	return o.ApplyT(func(v EnergyServicePropertiesResponse) []DataPartitionNamesResponse { return v.DataPartitionNames }).(DataPartitionNamesResponseArrayOutput)
}

func (o EnergyServicePropertiesResponseOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v EnergyServicePropertiesResponse) string { return v.DnsName }).(pulumi.StringOutput)
}

func (o EnergyServicePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v EnergyServicePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToOutput(ctx context.Context) pulumix.Output[SystemDataResponse] {
	return pulumix.Output[SystemDataResponse]{
		OutputState: o.OutputState,
	}
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DataPartitionNamesOutput{})
	pulumi.RegisterOutputType(DataPartitionNamesArrayOutput{})
	pulumi.RegisterOutputType(DataPartitionNamesResponseOutput{})
	pulumi.RegisterOutputType(DataPartitionNamesResponseArrayOutput{})
	pulumi.RegisterOutputType(DataPartitionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DataPartitionPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(EnergyServicePropertiesOutput{})
	pulumi.RegisterOutputType(EnergyServicePropertiesPtrOutput{})
	pulumi.RegisterOutputType(EnergyServicePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
