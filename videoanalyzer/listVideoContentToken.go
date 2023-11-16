// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package videoanalyzer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Generates a streaming token which can be used for accessing content from video content URLs, for a video resource with the given name.
// Azure REST API version: 2021-11-01-preview.
func ListVideoContentToken(ctx *pulumi.Context, args *ListVideoContentTokenArgs, opts ...pulumi.InvokeOption) (*ListVideoContentTokenResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListVideoContentTokenResult
	err := ctx.Invoke("azure-native:videoanalyzer:listVideoContentToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVideoContentTokenArgs struct {
	// The Azure Video Analyzer account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Video name.
	VideoName string `pulumi:"videoName"`
}

// "Video content token grants access to the video content URLs."
type ListVideoContentTokenResult struct {
	// The content token expiration date in ISO8601 format (eg. 2021-01-01T00:00:00Z).
	ExpirationDate string `pulumi:"expirationDate"`
	// The content token value to be added to the video content URL as the value for the "token" query string parameter. The token is specific to a single video.
	Token string `pulumi:"token"`
}

func ListVideoContentTokenOutput(ctx *pulumi.Context, args ListVideoContentTokenOutputArgs, opts ...pulumi.InvokeOption) ListVideoContentTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListVideoContentTokenResult, error) {
			args := v.(ListVideoContentTokenArgs)
			r, err := ListVideoContentToken(ctx, &args, opts...)
			var s ListVideoContentTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListVideoContentTokenResultOutput)
}

type ListVideoContentTokenOutputArgs struct {
	// The Azure Video Analyzer account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Video name.
	VideoName pulumi.StringInput `pulumi:"videoName"`
}

func (ListVideoContentTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVideoContentTokenArgs)(nil)).Elem()
}

// "Video content token grants access to the video content URLs."
type ListVideoContentTokenResultOutput struct{ *pulumi.OutputState }

func (ListVideoContentTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVideoContentTokenResult)(nil)).Elem()
}

func (o ListVideoContentTokenResultOutput) ToListVideoContentTokenResultOutput() ListVideoContentTokenResultOutput {
	return o
}

func (o ListVideoContentTokenResultOutput) ToListVideoContentTokenResultOutputWithContext(ctx context.Context) ListVideoContentTokenResultOutput {
	return o
}

// The content token expiration date in ISO8601 format (eg. 2021-01-01T00:00:00Z).
func (o ListVideoContentTokenResultOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v ListVideoContentTokenResult) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

// The content token value to be added to the video content URL as the value for the "token" query string parameter. The token is specific to a single video.
func (o ListVideoContentTokenResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v ListVideoContentTokenResult) string { return v.Token }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVideoContentTokenResultOutput{})
}
