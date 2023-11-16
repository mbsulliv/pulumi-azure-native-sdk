// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Represents the SKU name and Azure pricing tier for auto scale v-core resource.
type AutoScaleVCoreSku struct {
	// The capacity of an auto scale v-core resource.
	Capacity *int `pulumi:"capacity"`
	// Name of the SKU level.
	Name string `pulumi:"name"`
	// The name of the Azure pricing tier to which the SKU applies.
	Tier *string `pulumi:"tier"`
}

// AutoScaleVCoreSkuInput is an input type that accepts AutoScaleVCoreSkuArgs and AutoScaleVCoreSkuOutput values.
// You can construct a concrete instance of `AutoScaleVCoreSkuInput` via:
//
//	AutoScaleVCoreSkuArgs{...}
type AutoScaleVCoreSkuInput interface {
	pulumi.Input

	ToAutoScaleVCoreSkuOutput() AutoScaleVCoreSkuOutput
	ToAutoScaleVCoreSkuOutputWithContext(context.Context) AutoScaleVCoreSkuOutput
}

// Represents the SKU name and Azure pricing tier for auto scale v-core resource.
type AutoScaleVCoreSkuArgs struct {
	// The capacity of an auto scale v-core resource.
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	// Name of the SKU level.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Azure pricing tier to which the SKU applies.
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (AutoScaleVCoreSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleVCoreSku)(nil)).Elem()
}

func (i AutoScaleVCoreSkuArgs) ToAutoScaleVCoreSkuOutput() AutoScaleVCoreSkuOutput {
	return i.ToAutoScaleVCoreSkuOutputWithContext(context.Background())
}

func (i AutoScaleVCoreSkuArgs) ToAutoScaleVCoreSkuOutputWithContext(ctx context.Context) AutoScaleVCoreSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleVCoreSkuOutput)
}

// Represents the SKU name and Azure pricing tier for auto scale v-core resource.
type AutoScaleVCoreSkuOutput struct{ *pulumi.OutputState }

func (AutoScaleVCoreSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleVCoreSku)(nil)).Elem()
}

func (o AutoScaleVCoreSkuOutput) ToAutoScaleVCoreSkuOutput() AutoScaleVCoreSkuOutput {
	return o
}

func (o AutoScaleVCoreSkuOutput) ToAutoScaleVCoreSkuOutputWithContext(ctx context.Context) AutoScaleVCoreSkuOutput {
	return o
}

// The capacity of an auto scale v-core resource.
func (o AutoScaleVCoreSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleVCoreSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// Name of the SKU level.
func (o AutoScaleVCoreSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScaleVCoreSku) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the Azure pricing tier to which the SKU applies.
func (o AutoScaleVCoreSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoScaleVCoreSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

// Represents the SKU name and Azure pricing tier for auto scale v-core resource.
type AutoScaleVCoreSkuResponse struct {
	// The capacity of an auto scale v-core resource.
	Capacity *int `pulumi:"capacity"`
	// Name of the SKU level.
	Name string `pulumi:"name"`
	// The name of the Azure pricing tier to which the SKU applies.
	Tier *string `pulumi:"tier"`
}

// Represents the SKU name and Azure pricing tier for auto scale v-core resource.
type AutoScaleVCoreSkuResponseOutput struct{ *pulumi.OutputState }

func (AutoScaleVCoreSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleVCoreSkuResponse)(nil)).Elem()
}

func (o AutoScaleVCoreSkuResponseOutput) ToAutoScaleVCoreSkuResponseOutput() AutoScaleVCoreSkuResponseOutput {
	return o
}

func (o AutoScaleVCoreSkuResponseOutput) ToAutoScaleVCoreSkuResponseOutputWithContext(ctx context.Context) AutoScaleVCoreSkuResponseOutput {
	return o
}

// The capacity of an auto scale v-core resource.
func (o AutoScaleVCoreSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleVCoreSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// Name of the SKU level.
func (o AutoScaleVCoreSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScaleVCoreSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the Azure pricing tier to which the SKU applies.
func (o AutoScaleVCoreSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoScaleVCoreSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

// Represents the SKU name and Azure pricing tier for PowerBI Dedicated capacity resource.
type CapacitySku struct {
	// The capacity of the SKU.
	Capacity *int `pulumi:"capacity"`
	// Name of the SKU level.
	Name string `pulumi:"name"`
	// The name of the Azure pricing tier to which the SKU applies.
	Tier *string `pulumi:"tier"`
}

// CapacitySkuInput is an input type that accepts CapacitySkuArgs and CapacitySkuOutput values.
// You can construct a concrete instance of `CapacitySkuInput` via:
//
//	CapacitySkuArgs{...}
type CapacitySkuInput interface {
	pulumi.Input

	ToCapacitySkuOutput() CapacitySkuOutput
	ToCapacitySkuOutputWithContext(context.Context) CapacitySkuOutput
}

// Represents the SKU name and Azure pricing tier for PowerBI Dedicated capacity resource.
type CapacitySkuArgs struct {
	// The capacity of the SKU.
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	// Name of the SKU level.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Azure pricing tier to which the SKU applies.
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (CapacitySkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacitySku)(nil)).Elem()
}

func (i CapacitySkuArgs) ToCapacitySkuOutput() CapacitySkuOutput {
	return i.ToCapacitySkuOutputWithContext(context.Background())
}

func (i CapacitySkuArgs) ToCapacitySkuOutputWithContext(ctx context.Context) CapacitySkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacitySkuOutput)
}

// Represents the SKU name and Azure pricing tier for PowerBI Dedicated capacity resource.
type CapacitySkuOutput struct{ *pulumi.OutputState }

func (CapacitySkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacitySku)(nil)).Elem()
}

func (o CapacitySkuOutput) ToCapacitySkuOutput() CapacitySkuOutput {
	return o
}

func (o CapacitySkuOutput) ToCapacitySkuOutputWithContext(ctx context.Context) CapacitySkuOutput {
	return o
}

// The capacity of the SKU.
func (o CapacitySkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CapacitySku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// Name of the SKU level.
func (o CapacitySkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CapacitySku) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the Azure pricing tier to which the SKU applies.
func (o CapacitySkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapacitySku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

// Represents the SKU name and Azure pricing tier for PowerBI Dedicated capacity resource.
type CapacitySkuResponse struct {
	// The capacity of the SKU.
	Capacity *int `pulumi:"capacity"`
	// Name of the SKU level.
	Name string `pulumi:"name"`
	// The name of the Azure pricing tier to which the SKU applies.
	Tier *string `pulumi:"tier"`
}

// Represents the SKU name and Azure pricing tier for PowerBI Dedicated capacity resource.
type CapacitySkuResponseOutput struct{ *pulumi.OutputState }

func (CapacitySkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacitySkuResponse)(nil)).Elem()
}

func (o CapacitySkuResponseOutput) ToCapacitySkuResponseOutput() CapacitySkuResponseOutput {
	return o
}

func (o CapacitySkuResponseOutput) ToCapacitySkuResponseOutputWithContext(ctx context.Context) CapacitySkuResponseOutput {
	return o
}

// The capacity of the SKU.
func (o CapacitySkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CapacitySkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// Name of the SKU level.
func (o CapacitySkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CapacitySkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the Azure pricing tier to which the SKU applies.
func (o CapacitySkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapacitySkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

// An array of administrator user identities
type DedicatedCapacityAdministrators struct {
	// An array of administrator user identities.
	Members []string `pulumi:"members"`
}

// DedicatedCapacityAdministratorsInput is an input type that accepts DedicatedCapacityAdministratorsArgs and DedicatedCapacityAdministratorsOutput values.
// You can construct a concrete instance of `DedicatedCapacityAdministratorsInput` via:
//
//	DedicatedCapacityAdministratorsArgs{...}
type DedicatedCapacityAdministratorsInput interface {
	pulumi.Input

	ToDedicatedCapacityAdministratorsOutput() DedicatedCapacityAdministratorsOutput
	ToDedicatedCapacityAdministratorsOutputWithContext(context.Context) DedicatedCapacityAdministratorsOutput
}

// An array of administrator user identities
type DedicatedCapacityAdministratorsArgs struct {
	// An array of administrator user identities.
	Members pulumi.StringArrayInput `pulumi:"members"`
}

func (DedicatedCapacityAdministratorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedCapacityAdministrators)(nil)).Elem()
}

func (i DedicatedCapacityAdministratorsArgs) ToDedicatedCapacityAdministratorsOutput() DedicatedCapacityAdministratorsOutput {
	return i.ToDedicatedCapacityAdministratorsOutputWithContext(context.Background())
}

func (i DedicatedCapacityAdministratorsArgs) ToDedicatedCapacityAdministratorsOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsOutput)
}

func (i DedicatedCapacityAdministratorsArgs) ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput {
	return i.ToDedicatedCapacityAdministratorsPtrOutputWithContext(context.Background())
}

func (i DedicatedCapacityAdministratorsArgs) ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsOutput).ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx)
}

// DedicatedCapacityAdministratorsPtrInput is an input type that accepts DedicatedCapacityAdministratorsArgs, DedicatedCapacityAdministratorsPtr and DedicatedCapacityAdministratorsPtrOutput values.
// You can construct a concrete instance of `DedicatedCapacityAdministratorsPtrInput` via:
//
//	        DedicatedCapacityAdministratorsArgs{...}
//
//	or:
//
//	        nil
type DedicatedCapacityAdministratorsPtrInput interface {
	pulumi.Input

	ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput
	ToDedicatedCapacityAdministratorsPtrOutputWithContext(context.Context) DedicatedCapacityAdministratorsPtrOutput
}

type dedicatedCapacityAdministratorsPtrType DedicatedCapacityAdministratorsArgs

func DedicatedCapacityAdministratorsPtr(v *DedicatedCapacityAdministratorsArgs) DedicatedCapacityAdministratorsPtrInput {
	return (*dedicatedCapacityAdministratorsPtrType)(v)
}

func (*dedicatedCapacityAdministratorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCapacityAdministrators)(nil)).Elem()
}

func (i *dedicatedCapacityAdministratorsPtrType) ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput {
	return i.ToDedicatedCapacityAdministratorsPtrOutputWithContext(context.Background())
}

func (i *dedicatedCapacityAdministratorsPtrType) ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsPtrOutput)
}

// An array of administrator user identities
type DedicatedCapacityAdministratorsOutput struct{ *pulumi.OutputState }

func (DedicatedCapacityAdministratorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedCapacityAdministrators)(nil)).Elem()
}

func (o DedicatedCapacityAdministratorsOutput) ToDedicatedCapacityAdministratorsOutput() DedicatedCapacityAdministratorsOutput {
	return o
}

func (o DedicatedCapacityAdministratorsOutput) ToDedicatedCapacityAdministratorsOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsOutput {
	return o
}

func (o DedicatedCapacityAdministratorsOutput) ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput {
	return o.ToDedicatedCapacityAdministratorsPtrOutputWithContext(context.Background())
}

func (o DedicatedCapacityAdministratorsOutput) ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DedicatedCapacityAdministrators) *DedicatedCapacityAdministrators {
		return &v
	}).(DedicatedCapacityAdministratorsPtrOutput)
}

// An array of administrator user identities.
func (o DedicatedCapacityAdministratorsOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DedicatedCapacityAdministrators) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type DedicatedCapacityAdministratorsPtrOutput struct{ *pulumi.OutputState }

func (DedicatedCapacityAdministratorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCapacityAdministrators)(nil)).Elem()
}

func (o DedicatedCapacityAdministratorsPtrOutput) ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput {
	return o
}

func (o DedicatedCapacityAdministratorsPtrOutput) ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsPtrOutput {
	return o
}

func (o DedicatedCapacityAdministratorsPtrOutput) Elem() DedicatedCapacityAdministratorsOutput {
	return o.ApplyT(func(v *DedicatedCapacityAdministrators) DedicatedCapacityAdministrators {
		if v != nil {
			return *v
		}
		var ret DedicatedCapacityAdministrators
		return ret
	}).(DedicatedCapacityAdministratorsOutput)
}

// An array of administrator user identities.
func (o DedicatedCapacityAdministratorsPtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DedicatedCapacityAdministrators) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

// An array of administrator user identities
type DedicatedCapacityAdministratorsResponse struct {
	// An array of administrator user identities.
	Members []string `pulumi:"members"`
}

// An array of administrator user identities
type DedicatedCapacityAdministratorsResponseOutput struct{ *pulumi.OutputState }

func (DedicatedCapacityAdministratorsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedCapacityAdministratorsResponse)(nil)).Elem()
}

func (o DedicatedCapacityAdministratorsResponseOutput) ToDedicatedCapacityAdministratorsResponseOutput() DedicatedCapacityAdministratorsResponseOutput {
	return o
}

func (o DedicatedCapacityAdministratorsResponseOutput) ToDedicatedCapacityAdministratorsResponseOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponseOutput {
	return o
}

// An array of administrator user identities.
func (o DedicatedCapacityAdministratorsResponseOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DedicatedCapacityAdministratorsResponse) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type DedicatedCapacityAdministratorsResponsePtrOutput struct{ *pulumi.OutputState }

func (DedicatedCapacityAdministratorsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCapacityAdministratorsResponse)(nil)).Elem()
}

func (o DedicatedCapacityAdministratorsResponsePtrOutput) ToDedicatedCapacityAdministratorsResponsePtrOutput() DedicatedCapacityAdministratorsResponsePtrOutput {
	return o
}

func (o DedicatedCapacityAdministratorsResponsePtrOutput) ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponsePtrOutput {
	return o
}

func (o DedicatedCapacityAdministratorsResponsePtrOutput) Elem() DedicatedCapacityAdministratorsResponseOutput {
	return o.ApplyT(func(v *DedicatedCapacityAdministratorsResponse) DedicatedCapacityAdministratorsResponse {
		if v != nil {
			return *v
		}
		var ret DedicatedCapacityAdministratorsResponse
		return ret
	}).(DedicatedCapacityAdministratorsResponseOutput)
}

// An array of administrator user identities.
func (o DedicatedCapacityAdministratorsResponsePtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DedicatedCapacityAdministratorsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC)
	CreatedAt *string `pulumi:"createdAt"`
	// An identifier for the identity that created the resource
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// An identifier for the identity that last modified the resource
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// SystemDataInput is an input type that accepts SystemDataArgs and SystemDataOutput values.
// You can construct a concrete instance of `SystemDataInput` via:
//
//	SystemDataArgs{...}
type SystemDataInput interface {
	pulumi.Input

	ToSystemDataOutput() SystemDataOutput
	ToSystemDataOutputWithContext(context.Context) SystemDataOutput
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataArgs struct {
	// The timestamp of resource creation (UTC)
	CreatedAt pulumi.StringPtrInput `pulumi:"createdAt"`
	// An identifier for the identity that created the resource
	CreatedBy pulumi.StringPtrInput `pulumi:"createdBy"`
	// The type of identity that created the resource
	CreatedByType pulumi.StringPtrInput `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	// An identifier for the identity that last modified the resource
	LastModifiedBy pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemData)(nil)).Elem()
}

func (i SystemDataArgs) ToSystemDataOutput() SystemDataOutput {
	return i.ToSystemDataOutputWithContext(context.Background())
}

func (i SystemDataArgs) ToSystemDataOutputWithContext(ctx context.Context) SystemDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataOutput)
}

func (i SystemDataArgs) ToSystemDataPtrOutput() SystemDataPtrOutput {
	return i.ToSystemDataPtrOutputWithContext(context.Background())
}

func (i SystemDataArgs) ToSystemDataPtrOutputWithContext(ctx context.Context) SystemDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataOutput).ToSystemDataPtrOutputWithContext(ctx)
}

// SystemDataPtrInput is an input type that accepts SystemDataArgs, SystemDataPtr and SystemDataPtrOutput values.
// You can construct a concrete instance of `SystemDataPtrInput` via:
//
//	        SystemDataArgs{...}
//
//	or:
//
//	        nil
type SystemDataPtrInput interface {
	pulumi.Input

	ToSystemDataPtrOutput() SystemDataPtrOutput
	ToSystemDataPtrOutputWithContext(context.Context) SystemDataPtrOutput
}

type systemDataPtrType SystemDataArgs

func SystemDataPtr(v *SystemDataArgs) SystemDataPtrInput {
	return (*systemDataPtrType)(v)
}

func (*systemDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemData)(nil)).Elem()
}

func (i *systemDataPtrType) ToSystemDataPtrOutput() SystemDataPtrOutput {
	return i.ToSystemDataPtrOutputWithContext(context.Background())
}

func (i *systemDataPtrType) ToSystemDataPtrOutputWithContext(ctx context.Context) SystemDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataOutput struct{ *pulumi.OutputState }

func (SystemDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemData)(nil)).Elem()
}

func (o SystemDataOutput) ToSystemDataOutput() SystemDataOutput {
	return o
}

func (o SystemDataOutput) ToSystemDataOutputWithContext(ctx context.Context) SystemDataOutput {
	return o
}

func (o SystemDataOutput) ToSystemDataPtrOutput() SystemDataPtrOutput {
	return o.ToSystemDataPtrOutputWithContext(context.Background())
}

func (o SystemDataOutput) ToSystemDataPtrOutputWithContext(ctx context.Context) SystemDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemData) *SystemData {
		return &v
	}).(SystemDataPtrOutput)
}

// The timestamp of resource creation (UTC)
func (o SystemDataOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemData) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// An identifier for the identity that created the resource
func (o SystemDataOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemData) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource
func (o SystemDataOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemData) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemData) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// An identifier for the identity that last modified the resource
func (o SystemDataOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemData) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource
func (o SystemDataOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemData) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataPtrOutput struct{ *pulumi.OutputState }

func (SystemDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemData)(nil)).Elem()
}

func (o SystemDataPtrOutput) ToSystemDataPtrOutput() SystemDataPtrOutput {
	return o
}

func (o SystemDataPtrOutput) ToSystemDataPtrOutputWithContext(ctx context.Context) SystemDataPtrOutput {
	return o
}

func (o SystemDataPtrOutput) Elem() SystemDataOutput {
	return o.ApplyT(func(v *SystemData) SystemData {
		if v != nil {
			return *v
		}
		var ret SystemData
		return ret
	}).(SystemDataOutput)
}

// The timestamp of resource creation (UTC)
func (o SystemDataPtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemData) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

// An identifier for the identity that created the resource
func (o SystemDataPtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemData) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource
func (o SystemDataPtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemData) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataPtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemData) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

// An identifier for the identity that last modified the resource
func (o SystemDataPtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemData) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource
func (o SystemDataPtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemData) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC)
	CreatedAt *string `pulumi:"createdAt"`
	// An identifier for the identity that created the resource
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// An identifier for the identity that last modified the resource
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource
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

// The timestamp of resource creation (UTC)
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// An identifier for the identity that created the resource
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// An identifier for the identity that last modified the resource
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

// The timestamp of resource creation (UTC)
func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

// An identifier for the identity that created the resource
func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource
func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

// An identifier for the identity that last modified the resource
func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource
func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AutoScaleVCoreSkuOutput{})
	pulumi.RegisterOutputType(AutoScaleVCoreSkuResponseOutput{})
	pulumi.RegisterOutputType(CapacitySkuOutput{})
	pulumi.RegisterOutputType(CapacitySkuResponseOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsPtrOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsResponseOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataOutput{})
	pulumi.RegisterOutputType(SystemDataPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
