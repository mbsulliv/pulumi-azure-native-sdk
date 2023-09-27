// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the properties of the specified storage account credential name.
func LookupStorageAccountCredential(ctx *pulumi.Context, args *LookupStorageAccountCredentialArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountCredentialResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStorageAccountCredentialResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:getStorageAccountCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountCredentialArgs struct {
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of storage account credential to be fetched.
	StorageAccountCredentialName string `pulumi:"storageAccountCredentialName"`
}

// The storage account credential.
type LookupStorageAccountCredentialResult struct {
	// The details of the storage account password.
	AccessKey *AsymmetricEncryptedSecretResponse `pulumi:"accessKey"`
	// The storage endpoint
	EndPoint string `pulumi:"endPoint"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The name of the object.
	Name string `pulumi:"name"`
	// Signifies whether SSL needs to be enabled or not.
	SslStatus string `pulumi:"sslStatus"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
	// The count of volumes using this storage account credential.
	VolumesCount int `pulumi:"volumesCount"`
}

func LookupStorageAccountCredentialOutput(ctx *pulumi.Context, args LookupStorageAccountCredentialOutputArgs, opts ...pulumi.InvokeOption) LookupStorageAccountCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageAccountCredentialResult, error) {
			args := v.(LookupStorageAccountCredentialArgs)
			r, err := LookupStorageAccountCredential(ctx, &args, opts...)
			var s LookupStorageAccountCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageAccountCredentialResultOutput)
}

type LookupStorageAccountCredentialOutputArgs struct {
	// The manager name
	ManagerName pulumi.StringInput `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of storage account credential to be fetched.
	StorageAccountCredentialName pulumi.StringInput `pulumi:"storageAccountCredentialName"`
}

func (LookupStorageAccountCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountCredentialArgs)(nil)).Elem()
}

// The storage account credential.
type LookupStorageAccountCredentialResultOutput struct{ *pulumi.OutputState }

func (LookupStorageAccountCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountCredentialResult)(nil)).Elem()
}

func (o LookupStorageAccountCredentialResultOutput) ToLookupStorageAccountCredentialResultOutput() LookupStorageAccountCredentialResultOutput {
	return o
}

func (o LookupStorageAccountCredentialResultOutput) ToLookupStorageAccountCredentialResultOutputWithContext(ctx context.Context) LookupStorageAccountCredentialResultOutput {
	return o
}

func (o LookupStorageAccountCredentialResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupStorageAccountCredentialResult] {
	return pulumix.Output[LookupStorageAccountCredentialResult]{
		OutputState: o.OutputState,
	}
}

// The details of the storage account password.
func (o LookupStorageAccountCredentialResultOutput) AccessKey() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *AsymmetricEncryptedSecretResponse { return v.AccessKey }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

// The storage endpoint
func (o LookupStorageAccountCredentialResultOutput) EndPoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.EndPoint }).(pulumi.StringOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupStorageAccountCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Kind of the object. Currently only Series8000 is supported
func (o LookupStorageAccountCredentialResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The name of the object.
func (o LookupStorageAccountCredentialResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Name }).(pulumi.StringOutput)
}

// Signifies whether SSL needs to be enabled or not.
func (o LookupStorageAccountCredentialResultOutput) SslStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.SslStatus }).(pulumi.StringOutput)
}

// The hierarchical type of the object.
func (o LookupStorageAccountCredentialResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Type }).(pulumi.StringOutput)
}

// The count of volumes using this storage account credential.
func (o LookupStorageAccountCredentialResultOutput) VolumesCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) int { return v.VolumesCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageAccountCredentialResultOutput{})
}
