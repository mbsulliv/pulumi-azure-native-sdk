// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List Content Keys used by this Streaming Locator
func ListStreamingLocatorContentKeys(ctx *pulumi.Context, args *ListStreamingLocatorContentKeysArgs, opts ...pulumi.InvokeOption) (*ListStreamingLocatorContentKeysResult, error) {
	var rv ListStreamingLocatorContentKeysResult
	err := ctx.Invoke("azure-native:media/v20230101:listStreamingLocatorContentKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStreamingLocatorContentKeysArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Streaming Locator name.
	StreamingLocatorName string `pulumi:"streamingLocatorName"`
}

// Class of response for listContentKeys action
type ListStreamingLocatorContentKeysResult struct {
	// ContentKeys used by current Streaming Locator
	ContentKeys []StreamingLocatorContentKeyResponse `pulumi:"contentKeys"`
}

func ListStreamingLocatorContentKeysOutput(ctx *pulumi.Context, args ListStreamingLocatorContentKeysOutputArgs, opts ...pulumi.InvokeOption) ListStreamingLocatorContentKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStreamingLocatorContentKeysResult, error) {
			args := v.(ListStreamingLocatorContentKeysArgs)
			r, err := ListStreamingLocatorContentKeys(ctx, &args, opts...)
			var s ListStreamingLocatorContentKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStreamingLocatorContentKeysResultOutput)
}

type ListStreamingLocatorContentKeysOutputArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Streaming Locator name.
	StreamingLocatorName pulumi.StringInput `pulumi:"streamingLocatorName"`
}

func (ListStreamingLocatorContentKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStreamingLocatorContentKeysArgs)(nil)).Elem()
}

// Class of response for listContentKeys action
type ListStreamingLocatorContentKeysResultOutput struct{ *pulumi.OutputState }

func (ListStreamingLocatorContentKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStreamingLocatorContentKeysResult)(nil)).Elem()
}

func (o ListStreamingLocatorContentKeysResultOutput) ToListStreamingLocatorContentKeysResultOutput() ListStreamingLocatorContentKeysResultOutput {
	return o
}

func (o ListStreamingLocatorContentKeysResultOutput) ToListStreamingLocatorContentKeysResultOutputWithContext(ctx context.Context) ListStreamingLocatorContentKeysResultOutput {
	return o
}

// ContentKeys used by current Streaming Locator
func (o ListStreamingLocatorContentKeysResultOutput) ContentKeys() StreamingLocatorContentKeyResponseArrayOutput {
	return o.ApplyT(func(v ListStreamingLocatorContentKeysResult) []StreamingLocatorContentKeyResponse {
		return v.ContentKeys
	}).(StreamingLocatorContentKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStreamingLocatorContentKeysResultOutput{})
}
