// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get an resource upload URL for build service, which may be artifacts or source archive.
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview.
func GetBuildServiceResourceUploadUrl(ctx *pulumi.Context, args *GetBuildServiceResourceUploadUrlArgs, opts ...pulumi.InvokeOption) (*GetBuildServiceResourceUploadUrlResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetBuildServiceResourceUploadUrlResult
	err := ctx.Invoke("azure-native:appplatform:getBuildServiceResourceUploadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBuildServiceResourceUploadUrlArgs struct {
	// The name of the build service resource.
	BuildServiceName string `pulumi:"buildServiceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Resource upload definition payload
type GetBuildServiceResourceUploadUrlResult struct {
	// Source relative path
	RelativePath *string `pulumi:"relativePath"`
	// Upload URL
	UploadUrl *string `pulumi:"uploadUrl"`
}

func GetBuildServiceResourceUploadUrlOutput(ctx *pulumi.Context, args GetBuildServiceResourceUploadUrlOutputArgs, opts ...pulumi.InvokeOption) GetBuildServiceResourceUploadUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBuildServiceResourceUploadUrlResult, error) {
			args := v.(GetBuildServiceResourceUploadUrlArgs)
			r, err := GetBuildServiceResourceUploadUrl(ctx, &args, opts...)
			var s GetBuildServiceResourceUploadUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBuildServiceResourceUploadUrlResultOutput)
}

type GetBuildServiceResourceUploadUrlOutputArgs struct {
	// The name of the build service resource.
	BuildServiceName pulumi.StringInput `pulumi:"buildServiceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetBuildServiceResourceUploadUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBuildServiceResourceUploadUrlArgs)(nil)).Elem()
}

// Resource upload definition payload
type GetBuildServiceResourceUploadUrlResultOutput struct{ *pulumi.OutputState }

func (GetBuildServiceResourceUploadUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBuildServiceResourceUploadUrlResult)(nil)).Elem()
}

func (o GetBuildServiceResourceUploadUrlResultOutput) ToGetBuildServiceResourceUploadUrlResultOutput() GetBuildServiceResourceUploadUrlResultOutput {
	return o
}

func (o GetBuildServiceResourceUploadUrlResultOutput) ToGetBuildServiceResourceUploadUrlResultOutputWithContext(ctx context.Context) GetBuildServiceResourceUploadUrlResultOutput {
	return o
}

func (o GetBuildServiceResourceUploadUrlResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetBuildServiceResourceUploadUrlResult] {
	return pulumix.Output[GetBuildServiceResourceUploadUrlResult]{
		OutputState: o.OutputState,
	}
}

// Source relative path
func (o GetBuildServiceResourceUploadUrlResultOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBuildServiceResourceUploadUrlResult) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

// Upload URL
func (o GetBuildServiceResourceUploadUrlResultOutput) UploadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBuildServiceResourceUploadUrlResult) *string { return v.UploadUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBuildServiceResourceUploadUrlResultOutput{})
}
