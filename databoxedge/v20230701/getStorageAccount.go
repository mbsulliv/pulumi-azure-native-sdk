// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Storage Account on the  Data Box Edge/Gateway device.
func LookupStorageAccount(ctx *pulumi.Context, args *LookupStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStorageAccountResult
	err := ctx.Invoke("azure-native:databoxedge/v20230701:getStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The storage account name.
	StorageAccountName string `pulumi:"storageAccountName"`
}

// Represents a Storage Account on the  Data Box Edge/Gateway device.
type LookupStorageAccountResult struct {
	// BlobEndpoint of Storage Account
	BlobEndpoint string `pulumi:"blobEndpoint"`
	// The Container Count. Present only for Storage Accounts with DataPolicy set to Cloud.
	ContainerCount int `pulumi:"containerCount"`
	// Data policy of the storage Account.
	DataPolicy string `pulumi:"dataPolicy"`
	// Description for the storage Account.
	Description *string `pulumi:"description"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// The object name.
	Name string `pulumi:"name"`
	// Storage Account Credential Id
	StorageAccountCredentialId *string `pulumi:"storageAccountCredentialId"`
	// Current status of the storage account
	StorageAccountStatus *string `pulumi:"storageAccountStatus"`
	// Metadata pertaining to creation and last modification of StorageAccount
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
}

func LookupStorageAccountOutput(ctx *pulumi.Context, args LookupStorageAccountOutputArgs, opts ...pulumi.InvokeOption) LookupStorageAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageAccountResult, error) {
			args := v.(LookupStorageAccountArgs)
			r, err := LookupStorageAccount(ctx, &args, opts...)
			var s LookupStorageAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageAccountResultOutput)
}

type LookupStorageAccountOutputArgs struct {
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The storage account name.
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
}

func (LookupStorageAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountArgs)(nil)).Elem()
}

// Represents a Storage Account on the  Data Box Edge/Gateway device.
type LookupStorageAccountResultOutput struct{ *pulumi.OutputState }

func (LookupStorageAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountResult)(nil)).Elem()
}

func (o LookupStorageAccountResultOutput) ToLookupStorageAccountResultOutput() LookupStorageAccountResultOutput {
	return o
}

func (o LookupStorageAccountResultOutput) ToLookupStorageAccountResultOutputWithContext(ctx context.Context) LookupStorageAccountResultOutput {
	return o
}

// BlobEndpoint of Storage Account
func (o LookupStorageAccountResultOutput) BlobEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.BlobEndpoint }).(pulumi.StringOutput)
}

// The Container Count. Present only for Storage Accounts with DataPolicy set to Cloud.
func (o LookupStorageAccountResultOutput) ContainerCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) int { return v.ContainerCount }).(pulumi.IntOutput)
}

// Data policy of the storage Account.
func (o LookupStorageAccountResultOutput) DataPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.DataPolicy }).(pulumi.StringOutput)
}

// Description for the storage Account.
func (o LookupStorageAccountResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupStorageAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The object name.
func (o LookupStorageAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// Storage Account Credential Id
func (o LookupStorageAccountResultOutput) StorageAccountCredentialId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.StorageAccountCredentialId }).(pulumi.StringPtrOutput)
}

// Current status of the storage account
func (o LookupStorageAccountResultOutput) StorageAccountStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.StorageAccountStatus }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of StorageAccount
func (o LookupStorageAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o LookupStorageAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageAccountResultOutput{})
}
