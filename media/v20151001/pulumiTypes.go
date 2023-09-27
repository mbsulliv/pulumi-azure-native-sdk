// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20151001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// The properties for a Media Services REST API endpoint.
type ApiEndpointResponse struct {
	// The Media Services REST endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// The version of Media Services REST API.
	MajorVersion *string `pulumi:"majorVersion"`
}

// The properties for a Media Services REST API endpoint.
type ApiEndpointResponseOutput struct{ *pulumi.OutputState }

func (ApiEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEndpointResponse)(nil)).Elem()
}

func (o ApiEndpointResponseOutput) ToApiEndpointResponseOutput() ApiEndpointResponseOutput {
	return o
}

func (o ApiEndpointResponseOutput) ToApiEndpointResponseOutputWithContext(ctx context.Context) ApiEndpointResponseOutput {
	return o
}

func (o ApiEndpointResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ApiEndpointResponse] {
	return pulumix.Output[ApiEndpointResponse]{
		OutputState: o.OutputState,
	}
}

// The Media Services REST endpoint.
func (o ApiEndpointResponseOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEndpointResponse) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

// The version of Media Services REST API.
func (o ApiEndpointResponseOutput) MajorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEndpointResponse) *string { return v.MajorVersion }).(pulumi.StringPtrOutput)
}

type ApiEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (ApiEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiEndpointResponse)(nil)).Elem()
}

func (o ApiEndpointResponseArrayOutput) ToApiEndpointResponseArrayOutput() ApiEndpointResponseArrayOutput {
	return o
}

func (o ApiEndpointResponseArrayOutput) ToApiEndpointResponseArrayOutputWithContext(ctx context.Context) ApiEndpointResponseArrayOutput {
	return o
}

func (o ApiEndpointResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ApiEndpointResponse] {
	return pulumix.Output[[]ApiEndpointResponse]{
		OutputState: o.OutputState,
	}
}

func (o ApiEndpointResponseArrayOutput) Index(i pulumi.IntInput) ApiEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiEndpointResponse {
		return vs[0].([]ApiEndpointResponse)[vs[1].(int)]
	}).(ApiEndpointResponseOutput)
}

// The properties of a storage account associated with this resource.
type StorageAccount struct {
	// The id of the storage account resource. Media Services relies on tables and queues as well as blobs, so the primary storage account must be a Standard Storage account (either Microsoft.ClassicStorage or Microsoft.Storage). Blob only storage accounts can be added as secondary storage accounts (isPrimary false).
	Id string `pulumi:"id"`
	// Is this storage account resource the primary storage account for the Media Service resource. Blob only storage must set this to false.
	IsPrimary bool `pulumi:"isPrimary"`
}

// StorageAccountInput is an input type that accepts StorageAccountArgs and StorageAccountOutput values.
// You can construct a concrete instance of `StorageAccountInput` via:
//
//	StorageAccountArgs{...}
type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

// The properties of a storage account associated with this resource.
type StorageAccountArgs struct {
	// The id of the storage account resource. Media Services relies on tables and queues as well as blobs, so the primary storage account must be a Standard Storage account (either Microsoft.ClassicStorage or Microsoft.Storage). Blob only storage accounts can be added as secondary storage accounts (isPrimary false).
	Id pulumi.StringInput `pulumi:"id"`
	// Is this storage account resource the primary storage account for the Media Service resource. Blob only storage must set this to false.
	IsPrimary pulumi.BoolInput `pulumi:"isPrimary"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}

func (i StorageAccountArgs) ToOutput(ctx context.Context) pulumix.Output[StorageAccount] {
	return pulumix.Output[StorageAccount]{
		OutputState: i.ToStorageAccountOutputWithContext(ctx).OutputState,
	}
}

// StorageAccountArrayInput is an input type that accepts StorageAccountArray and StorageAccountArrayOutput values.
// You can construct a concrete instance of `StorageAccountArrayInput` via:
//
//	StorageAccountArray{ StorageAccountArgs{...} }
type StorageAccountArrayInput interface {
	pulumi.Input

	ToStorageAccountArrayOutput() StorageAccountArrayOutput
	ToStorageAccountArrayOutputWithContext(context.Context) StorageAccountArrayOutput
}

type StorageAccountArray []StorageAccountInput

func (StorageAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (i StorageAccountArray) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return i.ToStorageAccountArrayOutputWithContext(context.Background())
}

func (i StorageAccountArray) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountArrayOutput)
}

func (i StorageAccountArray) ToOutput(ctx context.Context) pulumix.Output[[]StorageAccount] {
	return pulumix.Output[[]StorageAccount]{
		OutputState: i.ToStorageAccountArrayOutputWithContext(ctx).OutputState,
	}
}

// The properties of a storage account associated with this resource.
type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToOutput(ctx context.Context) pulumix.Output[StorageAccount] {
	return pulumix.Output[StorageAccount]{
		OutputState: o.OutputState,
	}
}

// The id of the storage account resource. Media Services relies on tables and queues as well as blobs, so the primary storage account must be a Standard Storage account (either Microsoft.ClassicStorage or Microsoft.Storage). Blob only storage accounts can be added as secondary storage accounts (isPrimary false).
func (o StorageAccountOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.Id }).(pulumi.StringOutput)
}

// Is this storage account resource the primary storage account for the Media Service resource. Blob only storage must set this to false.
func (o StorageAccountOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v StorageAccount) bool { return v.IsPrimary }).(pulumi.BoolOutput)
}

type StorageAccountArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]StorageAccount] {
	return pulumix.Output[[]StorageAccount]{
		OutputState: o.OutputState,
	}
}

func (o StorageAccountArrayOutput) Index(i pulumi.IntInput) StorageAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccount {
		return vs[0].([]StorageAccount)[vs[1].(int)]
	}).(StorageAccountOutput)
}

// The properties of a storage account associated with this resource.
type StorageAccountResponse struct {
	// The id of the storage account resource. Media Services relies on tables and queues as well as blobs, so the primary storage account must be a Standard Storage account (either Microsoft.ClassicStorage or Microsoft.Storage). Blob only storage accounts can be added as secondary storage accounts (isPrimary false).
	Id string `pulumi:"id"`
	// Is this storage account resource the primary storage account for the Media Service resource. Blob only storage must set this to false.
	IsPrimary bool `pulumi:"isPrimary"`
}

// The properties of a storage account associated with this resource.
type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToOutput(ctx context.Context) pulumix.Output[StorageAccountResponse] {
	return pulumix.Output[StorageAccountResponse]{
		OutputState: o.OutputState,
	}
}

// The id of the storage account resource. Media Services relies on tables and queues as well as blobs, so the primary storage account must be a Standard Storage account (either Microsoft.ClassicStorage or Microsoft.Storage). Blob only storage accounts can be added as secondary storage accounts (isPrimary false).
func (o StorageAccountResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Id }).(pulumi.StringOutput)
}

// Is this storage account resource the primary storage account for the Media Service resource. Blob only storage must set this to false.
func (o StorageAccountResponseOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v StorageAccountResponse) bool { return v.IsPrimary }).(pulumi.BoolOutput)
}

type StorageAccountResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutputWithContext(ctx context.Context) StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]StorageAccountResponse] {
	return pulumix.Output[[]StorageAccountResponse]{
		OutputState: o.OutputState,
	}
}

func (o StorageAccountResponseArrayOutput) Index(i pulumi.IntInput) StorageAccountResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccountResponse {
		return vs[0].([]StorageAccountResponse)[vs[1].(int)]
	}).(StorageAccountResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiEndpointResponseOutput{})
	pulumi.RegisterOutputType(ApiEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseArrayOutput{})
}
