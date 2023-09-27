// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230406

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// Describes the billing related details of the AzureBareMetalStorageInstance.
type StorageBillingProperties struct {
	// the SKU type that is provisioned
	AzureBareMetalStorageInstanceSize *string `pulumi:"azureBareMetalStorageInstanceSize"`
	// the billing mode for the storage instance
	BillingMode *string `pulumi:"billingMode"`
}

// StorageBillingPropertiesInput is an input type that accepts StorageBillingPropertiesArgs and StorageBillingPropertiesOutput values.
// You can construct a concrete instance of `StorageBillingPropertiesInput` via:
//
//	StorageBillingPropertiesArgs{...}
type StorageBillingPropertiesInput interface {
	pulumi.Input

	ToStorageBillingPropertiesOutput() StorageBillingPropertiesOutput
	ToStorageBillingPropertiesOutputWithContext(context.Context) StorageBillingPropertiesOutput
}

// Describes the billing related details of the AzureBareMetalStorageInstance.
type StorageBillingPropertiesArgs struct {
	// the SKU type that is provisioned
	AzureBareMetalStorageInstanceSize pulumi.StringPtrInput `pulumi:"azureBareMetalStorageInstanceSize"`
	// the billing mode for the storage instance
	BillingMode pulumi.StringPtrInput `pulumi:"billingMode"`
}

func (StorageBillingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBillingProperties)(nil)).Elem()
}

func (i StorageBillingPropertiesArgs) ToStorageBillingPropertiesOutput() StorageBillingPropertiesOutput {
	return i.ToStorageBillingPropertiesOutputWithContext(context.Background())
}

func (i StorageBillingPropertiesArgs) ToStorageBillingPropertiesOutputWithContext(ctx context.Context) StorageBillingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBillingPropertiesOutput)
}

func (i StorageBillingPropertiesArgs) ToOutput(ctx context.Context) pulumix.Output[StorageBillingProperties] {
	return pulumix.Output[StorageBillingProperties]{
		OutputState: i.ToStorageBillingPropertiesOutputWithContext(ctx).OutputState,
	}
}

func (i StorageBillingPropertiesArgs) ToStorageBillingPropertiesPtrOutput() StorageBillingPropertiesPtrOutput {
	return i.ToStorageBillingPropertiesPtrOutputWithContext(context.Background())
}

func (i StorageBillingPropertiesArgs) ToStorageBillingPropertiesPtrOutputWithContext(ctx context.Context) StorageBillingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBillingPropertiesOutput).ToStorageBillingPropertiesPtrOutputWithContext(ctx)
}

// StorageBillingPropertiesPtrInput is an input type that accepts StorageBillingPropertiesArgs, StorageBillingPropertiesPtr and StorageBillingPropertiesPtrOutput values.
// You can construct a concrete instance of `StorageBillingPropertiesPtrInput` via:
//
//	        StorageBillingPropertiesArgs{...}
//
//	or:
//
//	        nil
type StorageBillingPropertiesPtrInput interface {
	pulumi.Input

	ToStorageBillingPropertiesPtrOutput() StorageBillingPropertiesPtrOutput
	ToStorageBillingPropertiesPtrOutputWithContext(context.Context) StorageBillingPropertiesPtrOutput
}

type storageBillingPropertiesPtrType StorageBillingPropertiesArgs

func StorageBillingPropertiesPtr(v *StorageBillingPropertiesArgs) StorageBillingPropertiesPtrInput {
	return (*storageBillingPropertiesPtrType)(v)
}

func (*storageBillingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBillingProperties)(nil)).Elem()
}

func (i *storageBillingPropertiesPtrType) ToStorageBillingPropertiesPtrOutput() StorageBillingPropertiesPtrOutput {
	return i.ToStorageBillingPropertiesPtrOutputWithContext(context.Background())
}

func (i *storageBillingPropertiesPtrType) ToStorageBillingPropertiesPtrOutputWithContext(ctx context.Context) StorageBillingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBillingPropertiesPtrOutput)
}

func (i *storageBillingPropertiesPtrType) ToOutput(ctx context.Context) pulumix.Output[*StorageBillingProperties] {
	return pulumix.Output[*StorageBillingProperties]{
		OutputState: i.ToStorageBillingPropertiesPtrOutputWithContext(ctx).OutputState,
	}
}

// Describes the billing related details of the AzureBareMetalStorageInstance.
type StorageBillingPropertiesOutput struct{ *pulumi.OutputState }

func (StorageBillingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBillingProperties)(nil)).Elem()
}

func (o StorageBillingPropertiesOutput) ToStorageBillingPropertiesOutput() StorageBillingPropertiesOutput {
	return o
}

func (o StorageBillingPropertiesOutput) ToStorageBillingPropertiesOutputWithContext(ctx context.Context) StorageBillingPropertiesOutput {
	return o
}

func (o StorageBillingPropertiesOutput) ToStorageBillingPropertiesPtrOutput() StorageBillingPropertiesPtrOutput {
	return o.ToStorageBillingPropertiesPtrOutputWithContext(context.Background())
}

func (o StorageBillingPropertiesOutput) ToStorageBillingPropertiesPtrOutputWithContext(ctx context.Context) StorageBillingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageBillingProperties) *StorageBillingProperties {
		return &v
	}).(StorageBillingPropertiesPtrOutput)
}

func (o StorageBillingPropertiesOutput) ToOutput(ctx context.Context) pulumix.Output[StorageBillingProperties] {
	return pulumix.Output[StorageBillingProperties]{
		OutputState: o.OutputState,
	}
}

// the SKU type that is provisioned
func (o StorageBillingPropertiesOutput) AzureBareMetalStorageInstanceSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageBillingProperties) *string { return v.AzureBareMetalStorageInstanceSize }).(pulumi.StringPtrOutput)
}

// the billing mode for the storage instance
func (o StorageBillingPropertiesOutput) BillingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageBillingProperties) *string { return v.BillingMode }).(pulumi.StringPtrOutput)
}

type StorageBillingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (StorageBillingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBillingProperties)(nil)).Elem()
}

func (o StorageBillingPropertiesPtrOutput) ToStorageBillingPropertiesPtrOutput() StorageBillingPropertiesPtrOutput {
	return o
}

func (o StorageBillingPropertiesPtrOutput) ToStorageBillingPropertiesPtrOutputWithContext(ctx context.Context) StorageBillingPropertiesPtrOutput {
	return o
}

func (o StorageBillingPropertiesPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*StorageBillingProperties] {
	return pulumix.Output[*StorageBillingProperties]{
		OutputState: o.OutputState,
	}
}

func (o StorageBillingPropertiesPtrOutput) Elem() StorageBillingPropertiesOutput {
	return o.ApplyT(func(v *StorageBillingProperties) StorageBillingProperties {
		if v != nil {
			return *v
		}
		var ret StorageBillingProperties
		return ret
	}).(StorageBillingPropertiesOutput)
}

// the SKU type that is provisioned
func (o StorageBillingPropertiesPtrOutput) AzureBareMetalStorageInstanceSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBillingProperties) *string {
		if v == nil {
			return nil
		}
		return v.AzureBareMetalStorageInstanceSize
	}).(pulumi.StringPtrOutput)
}

// the billing mode for the storage instance
func (o StorageBillingPropertiesPtrOutput) BillingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBillingProperties) *string {
		if v == nil {
			return nil
		}
		return v.BillingMode
	}).(pulumi.StringPtrOutput)
}

// Describes the billing related details of the AzureBareMetalStorageInstance.
type StorageBillingPropertiesResponse struct {
	// the SKU type that is provisioned
	AzureBareMetalStorageInstanceSize *string `pulumi:"azureBareMetalStorageInstanceSize"`
	// the billing mode for the storage instance
	BillingMode *string `pulumi:"billingMode"`
}

// Describes the billing related details of the AzureBareMetalStorageInstance.
type StorageBillingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StorageBillingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBillingPropertiesResponse)(nil)).Elem()
}

func (o StorageBillingPropertiesResponseOutput) ToStorageBillingPropertiesResponseOutput() StorageBillingPropertiesResponseOutput {
	return o
}

func (o StorageBillingPropertiesResponseOutput) ToStorageBillingPropertiesResponseOutputWithContext(ctx context.Context) StorageBillingPropertiesResponseOutput {
	return o
}

func (o StorageBillingPropertiesResponseOutput) ToOutput(ctx context.Context) pulumix.Output[StorageBillingPropertiesResponse] {
	return pulumix.Output[StorageBillingPropertiesResponse]{
		OutputState: o.OutputState,
	}
}

// the SKU type that is provisioned
func (o StorageBillingPropertiesResponseOutput) AzureBareMetalStorageInstanceSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageBillingPropertiesResponse) *string { return v.AzureBareMetalStorageInstanceSize }).(pulumi.StringPtrOutput)
}

// the billing mode for the storage instance
func (o StorageBillingPropertiesResponseOutput) BillingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageBillingPropertiesResponse) *string { return v.BillingMode }).(pulumi.StringPtrOutput)
}

type StorageBillingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageBillingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBillingPropertiesResponse)(nil)).Elem()
}

func (o StorageBillingPropertiesResponsePtrOutput) ToStorageBillingPropertiesResponsePtrOutput() StorageBillingPropertiesResponsePtrOutput {
	return o
}

func (o StorageBillingPropertiesResponsePtrOutput) ToStorageBillingPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageBillingPropertiesResponsePtrOutput {
	return o
}

func (o StorageBillingPropertiesResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*StorageBillingPropertiesResponse] {
	return pulumix.Output[*StorageBillingPropertiesResponse]{
		OutputState: o.OutputState,
	}
}

func (o StorageBillingPropertiesResponsePtrOutput) Elem() StorageBillingPropertiesResponseOutput {
	return o.ApplyT(func(v *StorageBillingPropertiesResponse) StorageBillingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret StorageBillingPropertiesResponse
		return ret
	}).(StorageBillingPropertiesResponseOutput)
}

// the SKU type that is provisioned
func (o StorageBillingPropertiesResponsePtrOutput) AzureBareMetalStorageInstanceSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBillingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AzureBareMetalStorageInstanceSize
	}).(pulumi.StringPtrOutput)
}

// the billing mode for the storage instance
func (o StorageBillingPropertiesResponsePtrOutput) BillingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBillingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.BillingMode
	}).(pulumi.StringPtrOutput)
}

// described the storage properties of the azure baremetalstorage instance
type StorageProperties struct {
	// the kind of storage instance
	Generation *string `pulumi:"generation"`
	// the hardware type of the storage instance
	HardwareType *string `pulumi:"hardwareType"`
	// the offering type for which the resource is getting provisioned
	OfferingType *string `pulumi:"offeringType"`
	// State of provisioning of the AzureBareMetalStorageInstance
	ProvisioningState *string `pulumi:"provisioningState"`
	// the billing related information for the resource
	StorageBillingProperties *StorageBillingProperties `pulumi:"storageBillingProperties"`
	// the storage protocol for which the resource is getting provisioned
	StorageType *string `pulumi:"storageType"`
	// the workload for which the resource is getting provisioned
	WorkloadType *string `pulumi:"workloadType"`
}

// StoragePropertiesInput is an input type that accepts StoragePropertiesArgs and StoragePropertiesOutput values.
// You can construct a concrete instance of `StoragePropertiesInput` via:
//
//	StoragePropertiesArgs{...}
type StoragePropertiesInput interface {
	pulumi.Input

	ToStoragePropertiesOutput() StoragePropertiesOutput
	ToStoragePropertiesOutputWithContext(context.Context) StoragePropertiesOutput
}

// described the storage properties of the azure baremetalstorage instance
type StoragePropertiesArgs struct {
	// the kind of storage instance
	Generation pulumi.StringPtrInput `pulumi:"generation"`
	// the hardware type of the storage instance
	HardwareType pulumi.StringPtrInput `pulumi:"hardwareType"`
	// the offering type for which the resource is getting provisioned
	OfferingType pulumi.StringPtrInput `pulumi:"offeringType"`
	// State of provisioning of the AzureBareMetalStorageInstance
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	// the billing related information for the resource
	StorageBillingProperties StorageBillingPropertiesPtrInput `pulumi:"storageBillingProperties"`
	// the storage protocol for which the resource is getting provisioned
	StorageType pulumi.StringPtrInput `pulumi:"storageType"`
	// the workload for which the resource is getting provisioned
	WorkloadType pulumi.StringPtrInput `pulumi:"workloadType"`
}

func (StoragePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProperties)(nil)).Elem()
}

func (i StoragePropertiesArgs) ToStoragePropertiesOutput() StoragePropertiesOutput {
	return i.ToStoragePropertiesOutputWithContext(context.Background())
}

func (i StoragePropertiesArgs) ToStoragePropertiesOutputWithContext(ctx context.Context) StoragePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoragePropertiesOutput)
}

func (i StoragePropertiesArgs) ToOutput(ctx context.Context) pulumix.Output[StorageProperties] {
	return pulumix.Output[StorageProperties]{
		OutputState: i.ToStoragePropertiesOutputWithContext(ctx).OutputState,
	}
}

func (i StoragePropertiesArgs) ToStoragePropertiesPtrOutput() StoragePropertiesPtrOutput {
	return i.ToStoragePropertiesPtrOutputWithContext(context.Background())
}

func (i StoragePropertiesArgs) ToStoragePropertiesPtrOutputWithContext(ctx context.Context) StoragePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoragePropertiesOutput).ToStoragePropertiesPtrOutputWithContext(ctx)
}

// StoragePropertiesPtrInput is an input type that accepts StoragePropertiesArgs, StoragePropertiesPtr and StoragePropertiesPtrOutput values.
// You can construct a concrete instance of `StoragePropertiesPtrInput` via:
//
//	        StoragePropertiesArgs{...}
//
//	or:
//
//	        nil
type StoragePropertiesPtrInput interface {
	pulumi.Input

	ToStoragePropertiesPtrOutput() StoragePropertiesPtrOutput
	ToStoragePropertiesPtrOutputWithContext(context.Context) StoragePropertiesPtrOutput
}

type storagePropertiesPtrType StoragePropertiesArgs

func StoragePropertiesPtr(v *StoragePropertiesArgs) StoragePropertiesPtrInput {
	return (*storagePropertiesPtrType)(v)
}

func (*storagePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProperties)(nil)).Elem()
}

func (i *storagePropertiesPtrType) ToStoragePropertiesPtrOutput() StoragePropertiesPtrOutput {
	return i.ToStoragePropertiesPtrOutputWithContext(context.Background())
}

func (i *storagePropertiesPtrType) ToStoragePropertiesPtrOutputWithContext(ctx context.Context) StoragePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoragePropertiesPtrOutput)
}

func (i *storagePropertiesPtrType) ToOutput(ctx context.Context) pulumix.Output[*StorageProperties] {
	return pulumix.Output[*StorageProperties]{
		OutputState: i.ToStoragePropertiesPtrOutputWithContext(ctx).OutputState,
	}
}

// described the storage properties of the azure baremetalstorage instance
type StoragePropertiesOutput struct{ *pulumi.OutputState }

func (StoragePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProperties)(nil)).Elem()
}

func (o StoragePropertiesOutput) ToStoragePropertiesOutput() StoragePropertiesOutput {
	return o
}

func (o StoragePropertiesOutput) ToStoragePropertiesOutputWithContext(ctx context.Context) StoragePropertiesOutput {
	return o
}

func (o StoragePropertiesOutput) ToStoragePropertiesPtrOutput() StoragePropertiesPtrOutput {
	return o.ToStoragePropertiesPtrOutputWithContext(context.Background())
}

func (o StoragePropertiesOutput) ToStoragePropertiesPtrOutputWithContext(ctx context.Context) StoragePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProperties) *StorageProperties {
		return &v
	}).(StoragePropertiesPtrOutput)
}

func (o StoragePropertiesOutput) ToOutput(ctx context.Context) pulumix.Output[StorageProperties] {
	return pulumix.Output[StorageProperties]{
		OutputState: o.OutputState,
	}
}

// the kind of storage instance
func (o StoragePropertiesOutput) Generation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProperties) *string { return v.Generation }).(pulumi.StringPtrOutput)
}

// the hardware type of the storage instance
func (o StoragePropertiesOutput) HardwareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProperties) *string { return v.HardwareType }).(pulumi.StringPtrOutput)
}

// the offering type for which the resource is getting provisioned
func (o StoragePropertiesOutput) OfferingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProperties) *string { return v.OfferingType }).(pulumi.StringPtrOutput)
}

// State of provisioning of the AzureBareMetalStorageInstance
func (o StoragePropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// the billing related information for the resource
func (o StoragePropertiesOutput) StorageBillingProperties() StorageBillingPropertiesPtrOutput {
	return o.ApplyT(func(v StorageProperties) *StorageBillingProperties { return v.StorageBillingProperties }).(StorageBillingPropertiesPtrOutput)
}

// the storage protocol for which the resource is getting provisioned
func (o StoragePropertiesOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProperties) *string { return v.StorageType }).(pulumi.StringPtrOutput)
}

// the workload for which the resource is getting provisioned
func (o StoragePropertiesOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProperties) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type StoragePropertiesPtrOutput struct{ *pulumi.OutputState }

func (StoragePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProperties)(nil)).Elem()
}

func (o StoragePropertiesPtrOutput) ToStoragePropertiesPtrOutput() StoragePropertiesPtrOutput {
	return o
}

func (o StoragePropertiesPtrOutput) ToStoragePropertiesPtrOutputWithContext(ctx context.Context) StoragePropertiesPtrOutput {
	return o
}

func (o StoragePropertiesPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*StorageProperties] {
	return pulumix.Output[*StorageProperties]{
		OutputState: o.OutputState,
	}
}

func (o StoragePropertiesPtrOutput) Elem() StoragePropertiesOutput {
	return o.ApplyT(func(v *StorageProperties) StorageProperties {
		if v != nil {
			return *v
		}
		var ret StorageProperties
		return ret
	}).(StoragePropertiesOutput)
}

// the kind of storage instance
func (o StoragePropertiesPtrOutput) Generation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProperties) *string {
		if v == nil {
			return nil
		}
		return v.Generation
	}).(pulumi.StringPtrOutput)
}

// the hardware type of the storage instance
func (o StoragePropertiesPtrOutput) HardwareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProperties) *string {
		if v == nil {
			return nil
		}
		return v.HardwareType
	}).(pulumi.StringPtrOutput)
}

// the offering type for which the resource is getting provisioned
func (o StoragePropertiesPtrOutput) OfferingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProperties) *string {
		if v == nil {
			return nil
		}
		return v.OfferingType
	}).(pulumi.StringPtrOutput)
}

// State of provisioning of the AzureBareMetalStorageInstance
func (o StoragePropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

// the billing related information for the resource
func (o StoragePropertiesPtrOutput) StorageBillingProperties() StorageBillingPropertiesPtrOutput {
	return o.ApplyT(func(v *StorageProperties) *StorageBillingProperties {
		if v == nil {
			return nil
		}
		return v.StorageBillingProperties
	}).(StorageBillingPropertiesPtrOutput)
}

// the storage protocol for which the resource is getting provisioned
func (o StoragePropertiesPtrOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProperties) *string {
		if v == nil {
			return nil
		}
		return v.StorageType
	}).(pulumi.StringPtrOutput)
}

// the workload for which the resource is getting provisioned
func (o StoragePropertiesPtrOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProperties) *string {
		if v == nil {
			return nil
		}
		return v.WorkloadType
	}).(pulumi.StringPtrOutput)
}

// described the storage properties of the azure baremetalstorage instance
type StoragePropertiesResponse struct {
	// the kind of storage instance
	Generation *string `pulumi:"generation"`
	// the hardware type of the storage instance
	HardwareType *string `pulumi:"hardwareType"`
	// the offering type for which the resource is getting provisioned
	OfferingType *string `pulumi:"offeringType"`
	// State of provisioning of the AzureBareMetalStorageInstance
	ProvisioningState *string `pulumi:"provisioningState"`
	// the billing related information for the resource
	StorageBillingProperties *StorageBillingPropertiesResponse `pulumi:"storageBillingProperties"`
	// the storage protocol for which the resource is getting provisioned
	StorageType *string `pulumi:"storageType"`
	// the workload for which the resource is getting provisioned
	WorkloadType *string `pulumi:"workloadType"`
}

// described the storage properties of the azure baremetalstorage instance
type StoragePropertiesResponseOutput struct{ *pulumi.OutputState }

func (StoragePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StoragePropertiesResponse)(nil)).Elem()
}

func (o StoragePropertiesResponseOutput) ToStoragePropertiesResponseOutput() StoragePropertiesResponseOutput {
	return o
}

func (o StoragePropertiesResponseOutput) ToStoragePropertiesResponseOutputWithContext(ctx context.Context) StoragePropertiesResponseOutput {
	return o
}

func (o StoragePropertiesResponseOutput) ToOutput(ctx context.Context) pulumix.Output[StoragePropertiesResponse] {
	return pulumix.Output[StoragePropertiesResponse]{
		OutputState: o.OutputState,
	}
}

// the kind of storage instance
func (o StoragePropertiesResponseOutput) Generation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoragePropertiesResponse) *string { return v.Generation }).(pulumi.StringPtrOutput)
}

// the hardware type of the storage instance
func (o StoragePropertiesResponseOutput) HardwareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoragePropertiesResponse) *string { return v.HardwareType }).(pulumi.StringPtrOutput)
}

// the offering type for which the resource is getting provisioned
func (o StoragePropertiesResponseOutput) OfferingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoragePropertiesResponse) *string { return v.OfferingType }).(pulumi.StringPtrOutput)
}

// State of provisioning of the AzureBareMetalStorageInstance
func (o StoragePropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoragePropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// the billing related information for the resource
func (o StoragePropertiesResponseOutput) StorageBillingProperties() StorageBillingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v StoragePropertiesResponse) *StorageBillingPropertiesResponse { return v.StorageBillingProperties }).(StorageBillingPropertiesResponsePtrOutput)
}

// the storage protocol for which the resource is getting provisioned
func (o StoragePropertiesResponseOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoragePropertiesResponse) *string { return v.StorageType }).(pulumi.StringPtrOutput)
}

// the workload for which the resource is getting provisioned
func (o StoragePropertiesResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoragePropertiesResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type StoragePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (StoragePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StoragePropertiesResponse)(nil)).Elem()
}

func (o StoragePropertiesResponsePtrOutput) ToStoragePropertiesResponsePtrOutput() StoragePropertiesResponsePtrOutput {
	return o
}

func (o StoragePropertiesResponsePtrOutput) ToStoragePropertiesResponsePtrOutputWithContext(ctx context.Context) StoragePropertiesResponsePtrOutput {
	return o
}

func (o StoragePropertiesResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*StoragePropertiesResponse] {
	return pulumix.Output[*StoragePropertiesResponse]{
		OutputState: o.OutputState,
	}
}

func (o StoragePropertiesResponsePtrOutput) Elem() StoragePropertiesResponseOutput {
	return o.ApplyT(func(v *StoragePropertiesResponse) StoragePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret StoragePropertiesResponse
		return ret
	}).(StoragePropertiesResponseOutput)
}

// the kind of storage instance
func (o StoragePropertiesResponsePtrOutput) Generation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Generation
	}).(pulumi.StringPtrOutput)
}

// the hardware type of the storage instance
func (o StoragePropertiesResponsePtrOutput) HardwareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.HardwareType
	}).(pulumi.StringPtrOutput)
}

// the offering type for which the resource is getting provisioned
func (o StoragePropertiesResponsePtrOutput) OfferingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.OfferingType
	}).(pulumi.StringPtrOutput)
}

// State of provisioning of the AzureBareMetalStorageInstance
func (o StoragePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

// the billing related information for the resource
func (o StoragePropertiesResponsePtrOutput) StorageBillingProperties() StorageBillingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *StoragePropertiesResponse) *StorageBillingPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.StorageBillingProperties
	}).(StorageBillingPropertiesResponsePtrOutput)
}

// the storage protocol for which the resource is getting provisioned
func (o StoragePropertiesResponsePtrOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageType
	}).(pulumi.StringPtrOutput)
}

// the workload for which the resource is getting provisioned
func (o StoragePropertiesResponsePtrOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.WorkloadType
	}).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(StorageBillingPropertiesOutput{})
	pulumi.RegisterOutputType(StorageBillingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(StorageBillingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageBillingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(StoragePropertiesOutput{})
	pulumi.RegisterOutputType(StoragePropertiesPtrOutput{})
	pulumi.RegisterOutputType(StoragePropertiesResponseOutput{})
	pulumi.RegisterOutputType(StoragePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
