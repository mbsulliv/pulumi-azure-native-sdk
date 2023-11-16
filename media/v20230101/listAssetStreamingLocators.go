// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists Streaming Locators which are associated with this asset.
func ListAssetStreamingLocators(ctx *pulumi.Context, args *ListAssetStreamingLocatorsArgs, opts ...pulumi.InvokeOption) (*ListAssetStreamingLocatorsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListAssetStreamingLocatorsResult
	err := ctx.Invoke("azure-native:media/v20230101:listAssetStreamingLocators", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAssetStreamingLocatorsArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Asset name.
	AssetName string `pulumi:"assetName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Streaming Locators associated with this Asset.
type ListAssetStreamingLocatorsResult struct {
	// The list of Streaming Locators.
	StreamingLocators []AssetStreamingLocatorResponse `pulumi:"streamingLocators"`
}

func ListAssetStreamingLocatorsOutput(ctx *pulumi.Context, args ListAssetStreamingLocatorsOutputArgs, opts ...pulumi.InvokeOption) ListAssetStreamingLocatorsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAssetStreamingLocatorsResult, error) {
			args := v.(ListAssetStreamingLocatorsArgs)
			r, err := ListAssetStreamingLocators(ctx, &args, opts...)
			var s ListAssetStreamingLocatorsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAssetStreamingLocatorsResultOutput)
}

type ListAssetStreamingLocatorsOutputArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The Asset name.
	AssetName pulumi.StringInput `pulumi:"assetName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListAssetStreamingLocatorsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAssetStreamingLocatorsArgs)(nil)).Elem()
}

// The Streaming Locators associated with this Asset.
type ListAssetStreamingLocatorsResultOutput struct{ *pulumi.OutputState }

func (ListAssetStreamingLocatorsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAssetStreamingLocatorsResult)(nil)).Elem()
}

func (o ListAssetStreamingLocatorsResultOutput) ToListAssetStreamingLocatorsResultOutput() ListAssetStreamingLocatorsResultOutput {
	return o
}

func (o ListAssetStreamingLocatorsResultOutput) ToListAssetStreamingLocatorsResultOutputWithContext(ctx context.Context) ListAssetStreamingLocatorsResultOutput {
	return o
}

// The list of Streaming Locators.
func (o ListAssetStreamingLocatorsResultOutput) StreamingLocators() AssetStreamingLocatorResponseArrayOutput {
	return o.ApplyT(func(v ListAssetStreamingLocatorsResult) []AssetStreamingLocatorResponse { return v.StreamingLocators }).(AssetStreamingLocatorResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAssetStreamingLocatorsResultOutput{})
}
