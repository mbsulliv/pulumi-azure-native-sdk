// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230915preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the access keys for the specified Azure Cosmos DB database account.
func ListDatabaseAccountKeys(ctx *pulumi.Context, args *ListDatabaseAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListDatabaseAccountKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDatabaseAccountKeysResult
	err := ctx.Invoke("azure-native:documentdb/v20230915preview:listDatabaseAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabaseAccountKeysArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The access keys for the given database account.
type ListDatabaseAccountKeysResult struct {
	// Base 64 encoded value of the primary read-write key.
	PrimaryMasterKey string `pulumi:"primaryMasterKey"`
	// Base 64 encoded value of the primary read-only key.
	PrimaryReadonlyMasterKey string `pulumi:"primaryReadonlyMasterKey"`
	// Base 64 encoded value of the secondary read-write key.
	SecondaryMasterKey string `pulumi:"secondaryMasterKey"`
	// Base 64 encoded value of the secondary read-only key.
	SecondaryReadonlyMasterKey string `pulumi:"secondaryReadonlyMasterKey"`
}

func ListDatabaseAccountKeysOutput(ctx *pulumi.Context, args ListDatabaseAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListDatabaseAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDatabaseAccountKeysResult, error) {
			args := v.(ListDatabaseAccountKeysArgs)
			r, err := ListDatabaseAccountKeys(ctx, &args, opts...)
			var s ListDatabaseAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDatabaseAccountKeysResultOutput)
}

type ListDatabaseAccountKeysOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDatabaseAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseAccountKeysArgs)(nil)).Elem()
}

// The access keys for the given database account.
type ListDatabaseAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListDatabaseAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseAccountKeysResult)(nil)).Elem()
}

func (o ListDatabaseAccountKeysResultOutput) ToListDatabaseAccountKeysResultOutput() ListDatabaseAccountKeysResultOutput {
	return o
}

func (o ListDatabaseAccountKeysResultOutput) ToListDatabaseAccountKeysResultOutputWithContext(ctx context.Context) ListDatabaseAccountKeysResultOutput {
	return o
}

// Base 64 encoded value of the primary read-write key.
func (o ListDatabaseAccountKeysResultOutput) PrimaryMasterKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseAccountKeysResult) string { return v.PrimaryMasterKey }).(pulumi.StringOutput)
}

// Base 64 encoded value of the primary read-only key.
func (o ListDatabaseAccountKeysResultOutput) PrimaryReadonlyMasterKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseAccountKeysResult) string { return v.PrimaryReadonlyMasterKey }).(pulumi.StringOutput)
}

// Base 64 encoded value of the secondary read-write key.
func (o ListDatabaseAccountKeysResultOutput) SecondaryMasterKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseAccountKeysResult) string { return v.SecondaryMasterKey }).(pulumi.StringOutput)
}

// Base 64 encoded value of the secondary read-only key.
func (o ListDatabaseAccountKeysResultOutput) SecondaryReadonlyMasterKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseAccountKeysResult) string { return v.SecondaryReadonlyMasterKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabaseAccountKeysResultOutput{})
}
