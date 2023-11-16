// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List the API keys for the transaction node.
func ListTransactionNodeApiKeys(ctx *pulumi.Context, args *ListTransactionNodeApiKeysArgs, opts ...pulumi.InvokeOption) (*ListTransactionNodeApiKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListTransactionNodeApiKeysResult
	err := ctx.Invoke("azure-native:blockchain/v20180601preview:listTransactionNodeApiKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTransactionNodeApiKeysArgs struct {
	// Blockchain member name.
	BlockchainMemberName string `pulumi:"blockchainMemberName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Transaction node name.
	TransactionNodeName string `pulumi:"transactionNodeName"`
}

// Collection of the API key payload which is exposed in the response of the resource provider.
type ListTransactionNodeApiKeysResult struct {
	// Gets or sets the collection of API key.
	Keys []ApiKeyResponse `pulumi:"keys"`
}

func ListTransactionNodeApiKeysOutput(ctx *pulumi.Context, args ListTransactionNodeApiKeysOutputArgs, opts ...pulumi.InvokeOption) ListTransactionNodeApiKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListTransactionNodeApiKeysResult, error) {
			args := v.(ListTransactionNodeApiKeysArgs)
			r, err := ListTransactionNodeApiKeys(ctx, &args, opts...)
			var s ListTransactionNodeApiKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListTransactionNodeApiKeysResultOutput)
}

type ListTransactionNodeApiKeysOutputArgs struct {
	// Blockchain member name.
	BlockchainMemberName pulumi.StringInput `pulumi:"blockchainMemberName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Transaction node name.
	TransactionNodeName pulumi.StringInput `pulumi:"transactionNodeName"`
}

func (ListTransactionNodeApiKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTransactionNodeApiKeysArgs)(nil)).Elem()
}

// Collection of the API key payload which is exposed in the response of the resource provider.
type ListTransactionNodeApiKeysResultOutput struct{ *pulumi.OutputState }

func (ListTransactionNodeApiKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTransactionNodeApiKeysResult)(nil)).Elem()
}

func (o ListTransactionNodeApiKeysResultOutput) ToListTransactionNodeApiKeysResultOutput() ListTransactionNodeApiKeysResultOutput {
	return o
}

func (o ListTransactionNodeApiKeysResultOutput) ToListTransactionNodeApiKeysResultOutputWithContext(ctx context.Context) ListTransactionNodeApiKeysResultOutput {
	return o
}

// Gets or sets the collection of API key.
func (o ListTransactionNodeApiKeysResultOutput) Keys() ApiKeyResponseArrayOutput {
	return o.ApplyT(func(v ListTransactionNodeApiKeysResult) []ApiKeyResponse { return v.Keys }).(ApiKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListTransactionNodeApiKeysResultOutput{})
}
