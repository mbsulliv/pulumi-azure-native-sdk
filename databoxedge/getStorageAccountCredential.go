// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of the specified storage account credential.
// Azure REST API version: 2022-03-01.
//
// Other available API versions: 2023-01-01-preview, 2023-07-01.
func LookupStorageAccountCredential(ctx *pulumi.Context, args *LookupStorageAccountCredentialArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountCredentialResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStorageAccountCredentialResult
	err := ctx.Invoke("azure-native:databoxedge:getStorageAccountCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountCredentialArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The storage account credential name.
	Name string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The storage account credential.
type LookupStorageAccountCredentialResult struct {
	// Encrypted storage key.
	AccountKey *AsymmetricEncryptedSecretResponse `pulumi:"accountKey"`
	// Type of storage accessed on the storage account.
	AccountType string `pulumi:"accountType"`
	// Alias for the storage account.
	Alias string `pulumi:"alias"`
	// Blob end point for private clouds.
	BlobDomainName *string `pulumi:"blobDomainName"`
	// Connection string for the storage account. Use this string if username and account key are not specified.
	ConnectionString *string `pulumi:"connectionString"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// The object name.
	Name string `pulumi:"name"`
	// Signifies whether SSL needs to be enabled or not.
	SslStatus string `pulumi:"sslStatus"`
	// Id of the storage account.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// Metadata pertaining to creation and last modification of StorageAccountCredential
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
	// Username for the storage account.
	UserName *string `pulumi:"userName"`
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
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The storage account credential name.
	Name pulumi.StringInput `pulumi:"name"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
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

// Encrypted storage key.
func (o LookupStorageAccountCredentialResultOutput) AccountKey() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *AsymmetricEncryptedSecretResponse { return v.AccountKey }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

// Type of storage accessed on the storage account.
func (o LookupStorageAccountCredentialResultOutput) AccountType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.AccountType }).(pulumi.StringOutput)
}

// Alias for the storage account.
func (o LookupStorageAccountCredentialResultOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Alias }).(pulumi.StringOutput)
}

// Blob end point for private clouds.
func (o LookupStorageAccountCredentialResultOutput) BlobDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *string { return v.BlobDomainName }).(pulumi.StringPtrOutput)
}

// Connection string for the storage account. Use this string if username and account key are not specified.
func (o LookupStorageAccountCredentialResultOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupStorageAccountCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

// The object name.
func (o LookupStorageAccountCredentialResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Name }).(pulumi.StringOutput)
}

// Signifies whether SSL needs to be enabled or not.
func (o LookupStorageAccountCredentialResultOutput) SslStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.SslStatus }).(pulumi.StringOutput)
}

// Id of the storage account.
func (o LookupStorageAccountCredentialResultOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of StorageAccountCredential
func (o LookupStorageAccountCredentialResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o LookupStorageAccountCredentialResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) string { return v.Type }).(pulumi.StringOutput)
}

// Username for the storage account.
func (o LookupStorageAccountCredentialResultOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountCredentialResult) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageAccountCredentialResultOutput{})
}
