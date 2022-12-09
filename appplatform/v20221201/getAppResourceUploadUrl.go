// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource upload definition payload
func GetAppResourceUploadUrl(ctx *pulumi.Context, args *GetAppResourceUploadUrlArgs, opts ...pulumi.InvokeOption) (*GetAppResourceUploadUrlResult, error) {
	var rv GetAppResourceUploadUrlResult
	err := ctx.Invoke("azure-native:appplatform/v20221201:getAppResourceUploadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetAppResourceUploadUrlArgs struct {
	// The name of the App resource.
	AppName string `pulumi:"appName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Resource upload definition payload
type GetAppResourceUploadUrlResult struct {
	// Source relative path
	RelativePath *string `pulumi:"relativePath"`
	// Upload URL
	UploadUrl *string `pulumi:"uploadUrl"`
}

func GetAppResourceUploadUrlOutput(ctx *pulumi.Context, args GetAppResourceUploadUrlOutputArgs, opts ...pulumi.InvokeOption) GetAppResourceUploadUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppResourceUploadUrlResult, error) {
			args := v.(GetAppResourceUploadUrlArgs)
			r, err := GetAppResourceUploadUrl(ctx, &args, opts...)
			var s GetAppResourceUploadUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppResourceUploadUrlResultOutput)
}

type GetAppResourceUploadUrlOutputArgs struct {
	// The name of the App resource.
	AppName pulumi.StringInput `pulumi:"appName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetAppResourceUploadUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppResourceUploadUrlArgs)(nil)).Elem()
}

// Resource upload definition payload
type GetAppResourceUploadUrlResultOutput struct{ *pulumi.OutputState }

func (GetAppResourceUploadUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppResourceUploadUrlResult)(nil)).Elem()
}

func (o GetAppResourceUploadUrlResultOutput) ToGetAppResourceUploadUrlResultOutput() GetAppResourceUploadUrlResultOutput {
	return o
}

func (o GetAppResourceUploadUrlResultOutput) ToGetAppResourceUploadUrlResultOutputWithContext(ctx context.Context) GetAppResourceUploadUrlResultOutput {
	return o
}

// Source relative path
func (o GetAppResourceUploadUrlResultOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppResourceUploadUrlResult) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

// Upload URL
func (o GetAppResourceUploadUrlResultOutput) UploadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppResourceUploadUrlResult) *string { return v.UploadUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppResourceUploadUrlResultOutput{})
}
