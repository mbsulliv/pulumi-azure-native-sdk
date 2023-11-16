// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// The login password for the container registry.
type RegistryPasswordResponse struct {
	// The password name.
	Name *string `pulumi:"name"`
	// The password value.
	Value *string `pulumi:"value"`
}

// The login password for the container registry.
type RegistryPasswordResponseOutput struct{ *pulumi.OutputState }

func (RegistryPasswordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryPasswordResponse)(nil)).Elem()
}

func (o RegistryPasswordResponseOutput) ToRegistryPasswordResponseOutput() RegistryPasswordResponseOutput {
	return o
}

func (o RegistryPasswordResponseOutput) ToRegistryPasswordResponseOutputWithContext(ctx context.Context) RegistryPasswordResponseOutput {
	return o
}

// The password name.
func (o RegistryPasswordResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryPasswordResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The password value.
func (o RegistryPasswordResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryPasswordResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type RegistryPasswordResponseArrayOutput struct{ *pulumi.OutputState }

func (RegistryPasswordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryPasswordResponse)(nil)).Elem()
}

func (o RegistryPasswordResponseArrayOutput) ToRegistryPasswordResponseArrayOutput() RegistryPasswordResponseArrayOutput {
	return o
}

func (o RegistryPasswordResponseArrayOutput) ToRegistryPasswordResponseArrayOutputWithContext(ctx context.Context) RegistryPasswordResponseArrayOutput {
	return o
}

func (o RegistryPasswordResponseArrayOutput) Index(i pulumi.IntInput) RegistryPasswordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegistryPasswordResponse {
		return vs[0].([]RegistryPasswordResponse)[vs[1].(int)]
	}).(RegistryPasswordResponseOutput)
}

// The SKU of a container registry.
type Sku struct {
	// The SKU name of the container registry. Required for registry creation. Allowed value: Basic.
	Name string `pulumi:"name"`
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

// The SKU of a container registry.
type SkuArgs struct {
	// The SKU name of the container registry. Required for registry creation. Allowed value: Basic.
	Name pulumi.StringInput `pulumi:"name"`
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

// The SKU of a container registry.
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

// The SKU name of the container registry. Required for registry creation. Allowed value: Basic.
func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

// The SKU of a container registry.
type SkuResponse struct {
	// The SKU name of the container registry. Required for registry creation. Allowed value: Basic.
	Name string `pulumi:"name"`
	// The SKU tier based on the SKU name.
	Tier string `pulumi:"tier"`
}

// The SKU of a container registry.
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

// The SKU name of the container registry. Required for registry creation. Allowed value: Basic.
func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The SKU tier based on the SKU name.
func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

// The parameters of a storage account for a container registry.
type StorageAccountParameters struct {
	// The access key to the storage account.
	AccessKey string `pulumi:"accessKey"`
	// The name of the storage account.
	Name string `pulumi:"name"`
}

// StorageAccountParametersInput is an input type that accepts StorageAccountParametersArgs and StorageAccountParametersOutput values.
// You can construct a concrete instance of `StorageAccountParametersInput` via:
//
//	StorageAccountParametersArgs{...}
type StorageAccountParametersInput interface {
	pulumi.Input

	ToStorageAccountParametersOutput() StorageAccountParametersOutput
	ToStorageAccountParametersOutputWithContext(context.Context) StorageAccountParametersOutput
}

// The parameters of a storage account for a container registry.
type StorageAccountParametersArgs struct {
	// The access key to the storage account.
	AccessKey pulumi.StringInput `pulumi:"accessKey"`
	// The name of the storage account.
	Name pulumi.StringInput `pulumi:"name"`
}

func (StorageAccountParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountParameters)(nil)).Elem()
}

func (i StorageAccountParametersArgs) ToStorageAccountParametersOutput() StorageAccountParametersOutput {
	return i.ToStorageAccountParametersOutputWithContext(context.Background())
}

func (i StorageAccountParametersArgs) ToStorageAccountParametersOutputWithContext(ctx context.Context) StorageAccountParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountParametersOutput)
}

// The parameters of a storage account for a container registry.
type StorageAccountParametersOutput struct{ *pulumi.OutputState }

func (StorageAccountParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountParameters)(nil)).Elem()
}

func (o StorageAccountParametersOutput) ToStorageAccountParametersOutput() StorageAccountParametersOutput {
	return o
}

func (o StorageAccountParametersOutput) ToStorageAccountParametersOutputWithContext(ctx context.Context) StorageAccountParametersOutput {
	return o
}

// The access key to the storage account.
func (o StorageAccountParametersOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountParameters) string { return v.AccessKey }).(pulumi.StringOutput)
}

// The name of the storage account.
func (o StorageAccountParametersOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountParameters) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of a storage account for a container registry.
type StorageAccountPropertiesResponse struct {
	// The name of the storage account.
	Name *string `pulumi:"name"`
}

// The properties of a storage account for a container registry.
type StorageAccountPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountPropertiesResponse)(nil)).Elem()
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput {
	return o
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutputWithContext(ctx context.Context) StorageAccountPropertiesResponseOutput {
	return o
}

// The name of the storage account.
func (o StorageAccountPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type StorageAccountPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountPropertiesResponse)(nil)).Elem()
}

func (o StorageAccountPropertiesResponsePtrOutput) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return o
}

func (o StorageAccountPropertiesResponsePtrOutput) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return o
}

func (o StorageAccountPropertiesResponsePtrOutput) Elem() StorageAccountPropertiesResponseOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) StorageAccountPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret StorageAccountPropertiesResponse
		return ret
	}).(StorageAccountPropertiesResponseOutput)
}

// The name of the storage account.
func (o StorageAccountPropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryPasswordResponseOutput{})
	pulumi.RegisterOutputType(RegistryPasswordResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountParametersOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponsePtrOutput{})
}
