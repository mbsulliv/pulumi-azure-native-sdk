// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the upload location for the user to be able to upload the source.
// Azure REST API version: 2019-06-01-preview.
func ListRegistryBuildSourceUploadUrl(ctx *pulumi.Context, args *ListRegistryBuildSourceUploadUrlArgs, opts ...pulumi.InvokeOption) (*ListRegistryBuildSourceUploadUrlResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListRegistryBuildSourceUploadUrlResult
	err := ctx.Invoke("azure-native:containerregistry:listRegistryBuildSourceUploadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRegistryBuildSourceUploadUrlArgs struct {
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The properties of a response to source upload request.
type ListRegistryBuildSourceUploadUrlResult struct {
	// The relative path to the source. This is used to submit the subsequent queue build request.
	RelativePath *string `pulumi:"relativePath"`
	// The URL where the client can upload the source.
	UploadUrl *string `pulumi:"uploadUrl"`
}

func ListRegistryBuildSourceUploadUrlOutput(ctx *pulumi.Context, args ListRegistryBuildSourceUploadUrlOutputArgs, opts ...pulumi.InvokeOption) ListRegistryBuildSourceUploadUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRegistryBuildSourceUploadUrlResult, error) {
			args := v.(ListRegistryBuildSourceUploadUrlArgs)
			r, err := ListRegistryBuildSourceUploadUrl(ctx, &args, opts...)
			var s ListRegistryBuildSourceUploadUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRegistryBuildSourceUploadUrlResultOutput)
}

type ListRegistryBuildSourceUploadUrlOutputArgs struct {
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListRegistryBuildSourceUploadUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryBuildSourceUploadUrlArgs)(nil)).Elem()
}

// The properties of a response to source upload request.
type ListRegistryBuildSourceUploadUrlResultOutput struct{ *pulumi.OutputState }

func (ListRegistryBuildSourceUploadUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryBuildSourceUploadUrlResult)(nil)).Elem()
}

func (o ListRegistryBuildSourceUploadUrlResultOutput) ToListRegistryBuildSourceUploadUrlResultOutput() ListRegistryBuildSourceUploadUrlResultOutput {
	return o
}

func (o ListRegistryBuildSourceUploadUrlResultOutput) ToListRegistryBuildSourceUploadUrlResultOutputWithContext(ctx context.Context) ListRegistryBuildSourceUploadUrlResultOutput {
	return o
}

func (o ListRegistryBuildSourceUploadUrlResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListRegistryBuildSourceUploadUrlResult] {
	return pulumix.Output[ListRegistryBuildSourceUploadUrlResult]{
		OutputState: o.OutputState,
	}
}

// The relative path to the source. This is used to submit the subsequent queue build request.
func (o ListRegistryBuildSourceUploadUrlResultOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListRegistryBuildSourceUploadUrlResult) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

// The URL where the client can upload the source.
func (o ListRegistryBuildSourceUploadUrlResultOutput) UploadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListRegistryBuildSourceUploadUrlResult) *string { return v.UploadUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRegistryBuildSourceUploadUrlResultOutput{})
}
