// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Lists the access key for the specified configuration store.
func ListConfigurationStoreKeys(ctx *pulumi.Context, args *ListConfigurationStoreKeysArgs, opts ...pulumi.InvokeOption) (*ListConfigurationStoreKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListConfigurationStoreKeysResult
	err := ctx.Invoke("azure-native:appconfiguration/v20230301:listConfigurationStoreKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConfigurationStoreKeysArgs struct {
	// The name of the configuration store.
	ConfigStoreName string `pulumi:"configStoreName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains a nextLink element, the value of the nextLink element will include a skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string `pulumi:"skipToken"`
}

// The result of a request to list API keys.
type ListConfigurationStoreKeysResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string `pulumi:"nextLink"`
	// The collection value.
	Value []ApiKeyResponse `pulumi:"value"`
}

func ListConfigurationStoreKeysOutput(ctx *pulumi.Context, args ListConfigurationStoreKeysOutputArgs, opts ...pulumi.InvokeOption) ListConfigurationStoreKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListConfigurationStoreKeysResult, error) {
			args := v.(ListConfigurationStoreKeysArgs)
			r, err := ListConfigurationStoreKeys(ctx, &args, opts...)
			var s ListConfigurationStoreKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListConfigurationStoreKeysResultOutput)
}

type ListConfigurationStoreKeysOutputArgs struct {
	// The name of the configuration store.
	ConfigStoreName pulumi.StringInput `pulumi:"configStoreName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains a nextLink element, the value of the nextLink element will include a skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListConfigurationStoreKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigurationStoreKeysArgs)(nil)).Elem()
}

// The result of a request to list API keys.
type ListConfigurationStoreKeysResultOutput struct{ *pulumi.OutputState }

func (ListConfigurationStoreKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigurationStoreKeysResult)(nil)).Elem()
}

func (o ListConfigurationStoreKeysResultOutput) ToListConfigurationStoreKeysResultOutput() ListConfigurationStoreKeysResultOutput {
	return o
}

func (o ListConfigurationStoreKeysResultOutput) ToListConfigurationStoreKeysResultOutputWithContext(ctx context.Context) ListConfigurationStoreKeysResultOutput {
	return o
}

func (o ListConfigurationStoreKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListConfigurationStoreKeysResult] {
	return pulumix.Output[ListConfigurationStoreKeysResult]{
		OutputState: o.OutputState,
	}
}

// The URI that can be used to request the next set of paged results.
func (o ListConfigurationStoreKeysResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListConfigurationStoreKeysResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// The collection value.
func (o ListConfigurationStoreKeysResultOutput) Value() ApiKeyResponseArrayOutput {
	return o.ApplyT(func(v ListConfigurationStoreKeysResult) []ApiKeyResponse { return v.Value }).(ApiKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConfigurationStoreKeysResultOutput{})
}
