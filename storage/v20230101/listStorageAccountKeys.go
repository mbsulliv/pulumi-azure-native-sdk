// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Lists the access keys or Kerberos keys (if active directory enabled) for the specified storage account.
func ListStorageAccountKeys(ctx *pulumi.Context, args *ListStorageAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListStorageAccountKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListStorageAccountKeysResult
	err := ctx.Invoke("azure-native:storage/v20230101:listStorageAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStorageAccountKeysArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// Specifies type of the key to be listed. Possible value is kerb.
	Expand *string `pulumi:"expand"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The response from the ListKeys operation.
type ListStorageAccountKeysResult struct {
	// Gets the list of storage account keys and their properties for the specified storage account.
	Keys []StorageAccountKeyResponse `pulumi:"keys"`
}

func ListStorageAccountKeysOutput(ctx *pulumi.Context, args ListStorageAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListStorageAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStorageAccountKeysResult, error) {
			args := v.(ListStorageAccountKeysArgs)
			r, err := ListStorageAccountKeys(ctx, &args, opts...)
			var s ListStorageAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStorageAccountKeysResultOutput)
}

type ListStorageAccountKeysOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Specifies type of the key to be listed. Possible value is kerb.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStorageAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountKeysArgs)(nil)).Elem()
}

// The response from the ListKeys operation.
type ListStorageAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListStorageAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountKeysResult)(nil)).Elem()
}

func (o ListStorageAccountKeysResultOutput) ToListStorageAccountKeysResultOutput() ListStorageAccountKeysResultOutput {
	return o
}

func (o ListStorageAccountKeysResultOutput) ToListStorageAccountKeysResultOutputWithContext(ctx context.Context) ListStorageAccountKeysResultOutput {
	return o
}

func (o ListStorageAccountKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListStorageAccountKeysResult] {
	return pulumix.Output[ListStorageAccountKeysResult]{
		OutputState: o.OutputState,
	}
}

// Gets the list of storage account keys and their properties for the specified storage account.
func (o ListStorageAccountKeysResultOutput) Keys() StorageAccountKeyResponseArrayOutput {
	return o.ApplyT(func(v ListStorageAccountKeysResult) []StorageAccountKeyResponse { return v.Keys }).(StorageAccountKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStorageAccountKeysResultOutput{})
}
