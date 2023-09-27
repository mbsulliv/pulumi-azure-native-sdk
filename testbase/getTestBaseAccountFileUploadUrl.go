// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package testbase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the file upload URL of a Test Base Account.
// Azure REST API version: 2022-04-01-preview.
func GetTestBaseAccountFileUploadUrl(ctx *pulumi.Context, args *GetTestBaseAccountFileUploadUrlArgs, opts ...pulumi.InvokeOption) (*GetTestBaseAccountFileUploadUrlResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetTestBaseAccountFileUploadUrlResult
	err := ctx.Invoke("azure-native:testbase:getTestBaseAccountFileUploadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTestBaseAccountFileUploadUrlArgs struct {
	// The custom file name of the uploaded blob.
	BlobName *string `pulumi:"blobName"`
	// The name of the resource group that contains the resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource name of the Test Base Account.
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}

// The URL response
type GetTestBaseAccountFileUploadUrlResult struct {
	// The blob path of the uploaded package. It will be used as the 'blobPath' property of PackageResource.
	BlobPath string `pulumi:"blobPath"`
	// The URL used for uploading the package.
	UploadUrl string `pulumi:"uploadUrl"`
}

func GetTestBaseAccountFileUploadUrlOutput(ctx *pulumi.Context, args GetTestBaseAccountFileUploadUrlOutputArgs, opts ...pulumi.InvokeOption) GetTestBaseAccountFileUploadUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTestBaseAccountFileUploadUrlResult, error) {
			args := v.(GetTestBaseAccountFileUploadUrlArgs)
			r, err := GetTestBaseAccountFileUploadUrl(ctx, &args, opts...)
			var s GetTestBaseAccountFileUploadUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTestBaseAccountFileUploadUrlResultOutput)
}

type GetTestBaseAccountFileUploadUrlOutputArgs struct {
	// The custom file name of the uploaded blob.
	BlobName pulumi.StringPtrInput `pulumi:"blobName"`
	// The name of the resource group that contains the resource.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The resource name of the Test Base Account.
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
}

func (GetTestBaseAccountFileUploadUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestBaseAccountFileUploadUrlArgs)(nil)).Elem()
}

// The URL response
type GetTestBaseAccountFileUploadUrlResultOutput struct{ *pulumi.OutputState }

func (GetTestBaseAccountFileUploadUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestBaseAccountFileUploadUrlResult)(nil)).Elem()
}

func (o GetTestBaseAccountFileUploadUrlResultOutput) ToGetTestBaseAccountFileUploadUrlResultOutput() GetTestBaseAccountFileUploadUrlResultOutput {
	return o
}

func (o GetTestBaseAccountFileUploadUrlResultOutput) ToGetTestBaseAccountFileUploadUrlResultOutputWithContext(ctx context.Context) GetTestBaseAccountFileUploadUrlResultOutput {
	return o
}

func (o GetTestBaseAccountFileUploadUrlResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetTestBaseAccountFileUploadUrlResult] {
	return pulumix.Output[GetTestBaseAccountFileUploadUrlResult]{
		OutputState: o.OutputState,
	}
}

// The blob path of the uploaded package. It will be used as the 'blobPath' property of PackageResource.
func (o GetTestBaseAccountFileUploadUrlResultOutput) BlobPath() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestBaseAccountFileUploadUrlResult) string { return v.BlobPath }).(pulumi.StringOutput)
}

// The URL used for uploading the package.
func (o GetTestBaseAccountFileUploadUrlResultOutput) UploadUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestBaseAccountFileUploadUrlResult) string { return v.UploadUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTestBaseAccountFileUploadUrlResultOutput{})
}
