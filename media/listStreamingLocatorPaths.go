// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// List Paths supported by this Streaming Locator
// Azure REST API version: 2023-01-01.
func ListStreamingLocatorPaths(ctx *pulumi.Context, args *ListStreamingLocatorPathsArgs, opts ...pulumi.InvokeOption) (*ListStreamingLocatorPathsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListStreamingLocatorPathsResult
	err := ctx.Invoke("azure-native:media:listStreamingLocatorPaths", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStreamingLocatorPathsArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Streaming Locator name.
	StreamingLocatorName string `pulumi:"streamingLocatorName"`
}

// Class of response for listPaths action
type ListStreamingLocatorPathsResult struct {
	// Download Paths supported by current Streaming Locator
	DownloadPaths []string `pulumi:"downloadPaths"`
	// Streaming Paths supported by current Streaming Locator
	StreamingPaths []StreamingPathResponse `pulumi:"streamingPaths"`
}

func ListStreamingLocatorPathsOutput(ctx *pulumi.Context, args ListStreamingLocatorPathsOutputArgs, opts ...pulumi.InvokeOption) ListStreamingLocatorPathsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStreamingLocatorPathsResult, error) {
			args := v.(ListStreamingLocatorPathsArgs)
			r, err := ListStreamingLocatorPaths(ctx, &args, opts...)
			var s ListStreamingLocatorPathsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStreamingLocatorPathsResultOutput)
}

type ListStreamingLocatorPathsOutputArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Streaming Locator name.
	StreamingLocatorName pulumi.StringInput `pulumi:"streamingLocatorName"`
}

func (ListStreamingLocatorPathsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStreamingLocatorPathsArgs)(nil)).Elem()
}

// Class of response for listPaths action
type ListStreamingLocatorPathsResultOutput struct{ *pulumi.OutputState }

func (ListStreamingLocatorPathsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStreamingLocatorPathsResult)(nil)).Elem()
}

func (o ListStreamingLocatorPathsResultOutput) ToListStreamingLocatorPathsResultOutput() ListStreamingLocatorPathsResultOutput {
	return o
}

func (o ListStreamingLocatorPathsResultOutput) ToListStreamingLocatorPathsResultOutputWithContext(ctx context.Context) ListStreamingLocatorPathsResultOutput {
	return o
}

func (o ListStreamingLocatorPathsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListStreamingLocatorPathsResult] {
	return pulumix.Output[ListStreamingLocatorPathsResult]{
		OutputState: o.OutputState,
	}
}

// Download Paths supported by current Streaming Locator
func (o ListStreamingLocatorPathsResultOutput) DownloadPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListStreamingLocatorPathsResult) []string { return v.DownloadPaths }).(pulumi.StringArrayOutput)
}

// Streaming Paths supported by current Streaming Locator
func (o ListStreamingLocatorPathsResultOutput) StreamingPaths() StreamingPathResponseArrayOutput {
	return o.ApplyT(func(v ListStreamingLocatorPathsResult) []StreamingPathResponse { return v.StreamingPaths }).(StreamingPathResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStreamingLocatorPathsResultOutput{})
}
