// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// The pair of customer secret.
type CustomerSecret struct {
	// The encryption algorithm used to encrypt data.
	Algorithm SupportedAlgorithm `pulumi:"algorithm"`
	// The identifier to the data service input object which this secret corresponds to.
	KeyIdentifier string `pulumi:"keyIdentifier"`
	// It contains the encrypted customer secret.
	KeyValue string `pulumi:"keyValue"`
}

// CustomerSecretInput is an input type that accepts CustomerSecretArgs and CustomerSecretOutput values.
// You can construct a concrete instance of `CustomerSecretInput` via:
//
//	CustomerSecretArgs{...}
type CustomerSecretInput interface {
	pulumi.Input

	ToCustomerSecretOutput() CustomerSecretOutput
	ToCustomerSecretOutputWithContext(context.Context) CustomerSecretOutput
}

// The pair of customer secret.
type CustomerSecretArgs struct {
	// The encryption algorithm used to encrypt data.
	Algorithm SupportedAlgorithmInput `pulumi:"algorithm"`
	// The identifier to the data service input object which this secret corresponds to.
	KeyIdentifier pulumi.StringInput `pulumi:"keyIdentifier"`
	// It contains the encrypted customer secret.
	KeyValue pulumi.StringInput `pulumi:"keyValue"`
}

func (CustomerSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSecret)(nil)).Elem()
}

func (i CustomerSecretArgs) ToCustomerSecretOutput() CustomerSecretOutput {
	return i.ToCustomerSecretOutputWithContext(context.Background())
}

func (i CustomerSecretArgs) ToCustomerSecretOutputWithContext(ctx context.Context) CustomerSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSecretOutput)
}

func (i CustomerSecretArgs) ToOutput(ctx context.Context) pulumix.Output[CustomerSecret] {
	return pulumix.Output[CustomerSecret]{
		OutputState: i.ToCustomerSecretOutputWithContext(ctx).OutputState,
	}
}

// CustomerSecretArrayInput is an input type that accepts CustomerSecretArray and CustomerSecretArrayOutput values.
// You can construct a concrete instance of `CustomerSecretArrayInput` via:
//
//	CustomerSecretArray{ CustomerSecretArgs{...} }
type CustomerSecretArrayInput interface {
	pulumi.Input

	ToCustomerSecretArrayOutput() CustomerSecretArrayOutput
	ToCustomerSecretArrayOutputWithContext(context.Context) CustomerSecretArrayOutput
}

type CustomerSecretArray []CustomerSecretInput

func (CustomerSecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomerSecret)(nil)).Elem()
}

func (i CustomerSecretArray) ToCustomerSecretArrayOutput() CustomerSecretArrayOutput {
	return i.ToCustomerSecretArrayOutputWithContext(context.Background())
}

func (i CustomerSecretArray) ToCustomerSecretArrayOutputWithContext(ctx context.Context) CustomerSecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSecretArrayOutput)
}

func (i CustomerSecretArray) ToOutput(ctx context.Context) pulumix.Output[[]CustomerSecret] {
	return pulumix.Output[[]CustomerSecret]{
		OutputState: i.ToCustomerSecretArrayOutputWithContext(ctx).OutputState,
	}
}

// The pair of customer secret.
type CustomerSecretOutput struct{ *pulumi.OutputState }

func (CustomerSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSecret)(nil)).Elem()
}

func (o CustomerSecretOutput) ToCustomerSecretOutput() CustomerSecretOutput {
	return o
}

func (o CustomerSecretOutput) ToCustomerSecretOutputWithContext(ctx context.Context) CustomerSecretOutput {
	return o
}

func (o CustomerSecretOutput) ToOutput(ctx context.Context) pulumix.Output[CustomerSecret] {
	return pulumix.Output[CustomerSecret]{
		OutputState: o.OutputState,
	}
}

// The encryption algorithm used to encrypt data.
func (o CustomerSecretOutput) Algorithm() SupportedAlgorithmOutput {
	return o.ApplyT(func(v CustomerSecret) SupportedAlgorithm { return v.Algorithm }).(SupportedAlgorithmOutput)
}

// The identifier to the data service input object which this secret corresponds to.
func (o CustomerSecretOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerSecret) string { return v.KeyIdentifier }).(pulumi.StringOutput)
}

// It contains the encrypted customer secret.
func (o CustomerSecretOutput) KeyValue() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerSecret) string { return v.KeyValue }).(pulumi.StringOutput)
}

type CustomerSecretArrayOutput struct{ *pulumi.OutputState }

func (CustomerSecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomerSecret)(nil)).Elem()
}

func (o CustomerSecretArrayOutput) ToCustomerSecretArrayOutput() CustomerSecretArrayOutput {
	return o
}

func (o CustomerSecretArrayOutput) ToCustomerSecretArrayOutputWithContext(ctx context.Context) CustomerSecretArrayOutput {
	return o
}

func (o CustomerSecretArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]CustomerSecret] {
	return pulumix.Output[[]CustomerSecret]{
		OutputState: o.OutputState,
	}
}

func (o CustomerSecretArrayOutput) Index(i pulumi.IntInput) CustomerSecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomerSecret {
		return vs[0].([]CustomerSecret)[vs[1].(int)]
	}).(CustomerSecretOutput)
}

// The pair of customer secret.
type CustomerSecretResponse struct {
	// The encryption algorithm used to encrypt data.
	Algorithm string `pulumi:"algorithm"`
	// The identifier to the data service input object which this secret corresponds to.
	KeyIdentifier string `pulumi:"keyIdentifier"`
	// It contains the encrypted customer secret.
	KeyValue string `pulumi:"keyValue"`
}

// The pair of customer secret.
type CustomerSecretResponseOutput struct{ *pulumi.OutputState }

func (CustomerSecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSecretResponse)(nil)).Elem()
}

func (o CustomerSecretResponseOutput) ToCustomerSecretResponseOutput() CustomerSecretResponseOutput {
	return o
}

func (o CustomerSecretResponseOutput) ToCustomerSecretResponseOutputWithContext(ctx context.Context) CustomerSecretResponseOutput {
	return o
}

func (o CustomerSecretResponseOutput) ToOutput(ctx context.Context) pulumix.Output[CustomerSecretResponse] {
	return pulumix.Output[CustomerSecretResponse]{
		OutputState: o.OutputState,
	}
}

// The encryption algorithm used to encrypt data.
func (o CustomerSecretResponseOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerSecretResponse) string { return v.Algorithm }).(pulumi.StringOutput)
}

// The identifier to the data service input object which this secret corresponds to.
func (o CustomerSecretResponseOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerSecretResponse) string { return v.KeyIdentifier }).(pulumi.StringOutput)
}

// It contains the encrypted customer secret.
func (o CustomerSecretResponseOutput) KeyValue() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerSecretResponse) string { return v.KeyValue }).(pulumi.StringOutput)
}

type CustomerSecretResponseArrayOutput struct{ *pulumi.OutputState }

func (CustomerSecretResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomerSecretResponse)(nil)).Elem()
}

func (o CustomerSecretResponseArrayOutput) ToCustomerSecretResponseArrayOutput() CustomerSecretResponseArrayOutput {
	return o
}

func (o CustomerSecretResponseArrayOutput) ToCustomerSecretResponseArrayOutputWithContext(ctx context.Context) CustomerSecretResponseArrayOutput {
	return o
}

func (o CustomerSecretResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]CustomerSecretResponse] {
	return pulumix.Output[[]CustomerSecretResponse]{
		OutputState: o.OutputState,
	}
}

func (o CustomerSecretResponseArrayOutput) Index(i pulumi.IntInput) CustomerSecretResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomerSecretResponse {
		return vs[0].([]CustomerSecretResponse)[vs[1].(int)]
	}).(CustomerSecretResponseOutput)
}

// Schedule for the job run.
type Schedule struct {
	// Name of the schedule.
	Name *string `pulumi:"name"`
	// A list of repetition intervals in ISO 8601 format.
	PolicyList []string `pulumi:"policyList"`
}

// ScheduleInput is an input type that accepts ScheduleArgs and ScheduleOutput values.
// You can construct a concrete instance of `ScheduleInput` via:
//
//	ScheduleArgs{...}
type ScheduleInput interface {
	pulumi.Input

	ToScheduleOutput() ScheduleOutput
	ToScheduleOutputWithContext(context.Context) ScheduleOutput
}

// Schedule for the job run.
type ScheduleArgs struct {
	// Name of the schedule.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A list of repetition intervals in ISO 8601 format.
	PolicyList pulumi.StringArrayInput `pulumi:"policyList"`
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (i ScheduleArgs) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i ScheduleArgs) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}

func (i ScheduleArgs) ToOutput(ctx context.Context) pulumix.Output[Schedule] {
	return pulumix.Output[Schedule]{
		OutputState: i.ToScheduleOutputWithContext(ctx).OutputState,
	}
}

// ScheduleArrayInput is an input type that accepts ScheduleArray and ScheduleArrayOutput values.
// You can construct a concrete instance of `ScheduleArrayInput` via:
//
//	ScheduleArray{ ScheduleArgs{...} }
type ScheduleArrayInput interface {
	pulumi.Input

	ToScheduleArrayOutput() ScheduleArrayOutput
	ToScheduleArrayOutputWithContext(context.Context) ScheduleArrayOutput
}

type ScheduleArray []ScheduleInput

func (ScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Schedule)(nil)).Elem()
}

func (i ScheduleArray) ToScheduleArrayOutput() ScheduleArrayOutput {
	return i.ToScheduleArrayOutputWithContext(context.Background())
}

func (i ScheduleArray) ToScheduleArrayOutputWithContext(ctx context.Context) ScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleArrayOutput)
}

func (i ScheduleArray) ToOutput(ctx context.Context) pulumix.Output[[]Schedule] {
	return pulumix.Output[[]Schedule]{
		OutputState: i.ToScheduleArrayOutputWithContext(ctx).OutputState,
	}
}

// Schedule for the job run.
type ScheduleOutput struct{ *pulumi.OutputState }

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToOutput(ctx context.Context) pulumix.Output[Schedule] {
	return pulumix.Output[Schedule]{
		OutputState: o.OutputState,
	}
}

// Name of the schedule.
func (o ScheduleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Schedule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of repetition intervals in ISO 8601 format.
func (o ScheduleOutput) PolicyList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Schedule) []string { return v.PolicyList }).(pulumi.StringArrayOutput)
}

type ScheduleArrayOutput struct{ *pulumi.OutputState }

func (ScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Schedule)(nil)).Elem()
}

func (o ScheduleArrayOutput) ToScheduleArrayOutput() ScheduleArrayOutput {
	return o
}

func (o ScheduleArrayOutput) ToScheduleArrayOutputWithContext(ctx context.Context) ScheduleArrayOutput {
	return o
}

func (o ScheduleArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]Schedule] {
	return pulumix.Output[[]Schedule]{
		OutputState: o.OutputState,
	}
}

func (o ScheduleArrayOutput) Index(i pulumi.IntInput) ScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Schedule {
		return vs[0].([]Schedule)[vs[1].(int)]
	}).(ScheduleOutput)
}

// Schedule for the job run.
type ScheduleResponse struct {
	// Name of the schedule.
	Name *string `pulumi:"name"`
	// A list of repetition intervals in ISO 8601 format.
	PolicyList []string `pulumi:"policyList"`
}

// Schedule for the job run.
type ScheduleResponseOutput struct{ *pulumi.OutputState }

func (ScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponseOutput) ToScheduleResponseOutput() ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) ToScheduleResponseOutputWithContext(ctx context.Context) ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ScheduleResponse] {
	return pulumix.Output[ScheduleResponse]{
		OutputState: o.OutputState,
	}
}

// Name of the schedule.
func (o ScheduleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of repetition intervals in ISO 8601 format.
func (o ScheduleResponseOutput) PolicyList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScheduleResponse) []string { return v.PolicyList }).(pulumi.StringArrayOutput)
}

type ScheduleResponseArrayOutput struct{ *pulumi.OutputState }

func (ScheduleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponseArrayOutput) ToScheduleResponseArrayOutput() ScheduleResponseArrayOutput {
	return o
}

func (o ScheduleResponseArrayOutput) ToScheduleResponseArrayOutputWithContext(ctx context.Context) ScheduleResponseArrayOutput {
	return o
}

func (o ScheduleResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ScheduleResponse] {
	return pulumix.Output[[]ScheduleResponse]{
		OutputState: o.OutputState,
	}
}

func (o ScheduleResponseArrayOutput) Index(i pulumi.IntInput) ScheduleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleResponse {
		return vs[0].([]ScheduleResponse)[vs[1].(int)]
	}).(ScheduleResponseOutput)
}

// The sku type.
type Sku struct {
	// The sku name. Required for data manager creation, optional for update.
	Name *string `pulumi:"name"`
	// The sku tier. This is based on the SKU name.
	Tier *string `pulumi:"tier"`
}

// SkuInput is an input type that accepts SkuArgs and SkuOutput values.
// You can construct a concrete instance of `SkuInput` via:
//
//	SkuArgs{...}
type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

// The sku type.
type SkuArgs struct {
	// The sku name. Required for data manager creation, optional for update.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The sku tier. This is based on the SKU name.
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToOutput(ctx context.Context) pulumix.Output[Sku] {
	return pulumix.Output[Sku]{
		OutputState: i.ToSkuOutputWithContext(ctx).OutputState,
	}
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}

// SkuPtrInput is an input type that accepts SkuArgs, SkuPtr and SkuPtrOutput values.
// You can construct a concrete instance of `SkuPtrInput` via:
//
//	        SkuArgs{...}
//
//	or:
//
//	        nil
type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

func (i *skuPtrType) ToOutput(ctx context.Context) pulumix.Output[*Sku] {
	return pulumix.Output[*Sku]{
		OutputState: i.ToSkuPtrOutputWithContext(ctx).OutputState,
	}
}

// The sku type.
type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) ToOutput(ctx context.Context) pulumix.Output[Sku] {
	return pulumix.Output[Sku]{
		OutputState: o.OutputState,
	}
}

// The sku name. Required for data manager creation, optional for update.
func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The sku tier. This is based on the SKU name.
func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*Sku] {
	return pulumix.Output[*Sku]{
		OutputState: o.OutputState,
	}
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

// The sku name. Required for data manager creation, optional for update.
func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// The sku tier. This is based on the SKU name.
func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

// The sku type.
type SkuResponse struct {
	// The sku name. Required for data manager creation, optional for update.
	Name *string `pulumi:"name"`
	// The sku tier. This is based on the SKU name.
	Tier *string `pulumi:"tier"`
}

// The sku type.
type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToOutput(ctx context.Context) pulumix.Output[SkuResponse] {
	return pulumix.Output[SkuResponse]{
		OutputState: o.OutputState,
	}
}

// The sku name. Required for data manager creation, optional for update.
func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The sku tier. This is based on the SKU name.
func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*SkuResponse] {
	return pulumix.Output[*SkuResponse]{
		OutputState: o.OutputState,
	}
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

// The sku name. Required for data manager creation, optional for update.
func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// The sku tier. This is based on the SKU name.
func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomerSecretOutput{})
	pulumi.RegisterOutputType(CustomerSecretArrayOutput{})
	pulumi.RegisterOutputType(CustomerSecretResponseOutput{})
	pulumi.RegisterOutputType(CustomerSecretResponseArrayOutput{})
	pulumi.RegisterOutputType(ScheduleOutput{})
	pulumi.RegisterOutputType(ScheduleArrayOutput{})
	pulumi.RegisterOutputType(ScheduleResponseOutput{})
	pulumi.RegisterOutputType(ScheduleResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
