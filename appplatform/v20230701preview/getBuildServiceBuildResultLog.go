// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a KPack build result log download URL.
func GetBuildServiceBuildResultLog(ctx *pulumi.Context, args *GetBuildServiceBuildResultLogArgs, opts ...pulumi.InvokeOption) (*GetBuildServiceBuildResultLogResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetBuildServiceBuildResultLogResult
	err := ctx.Invoke("azure-native:appplatform/v20230701preview:getBuildServiceBuildResultLog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBuildServiceBuildResultLogArgs struct {
	// The name of the build resource.
	BuildName string `pulumi:"buildName"`
	// The name of the build result resource.
	BuildResultName string `pulumi:"buildResultName"`
	// The name of the build service resource.
	BuildServiceName string `pulumi:"buildServiceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Build result log resource properties payload
type GetBuildServiceBuildResultLogResult struct {
	// The public download URL of this build result log
	BlobUrl *string `pulumi:"blobUrl"`
}

func GetBuildServiceBuildResultLogOutput(ctx *pulumi.Context, args GetBuildServiceBuildResultLogOutputArgs, opts ...pulumi.InvokeOption) GetBuildServiceBuildResultLogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBuildServiceBuildResultLogResult, error) {
			args := v.(GetBuildServiceBuildResultLogArgs)
			r, err := GetBuildServiceBuildResultLog(ctx, &args, opts...)
			var s GetBuildServiceBuildResultLogResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBuildServiceBuildResultLogResultOutput)
}

type GetBuildServiceBuildResultLogOutputArgs struct {
	// The name of the build resource.
	BuildName pulumi.StringInput `pulumi:"buildName"`
	// The name of the build result resource.
	BuildResultName pulumi.StringInput `pulumi:"buildResultName"`
	// The name of the build service resource.
	BuildServiceName pulumi.StringInput `pulumi:"buildServiceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetBuildServiceBuildResultLogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBuildServiceBuildResultLogArgs)(nil)).Elem()
}

// Build result log resource properties payload
type GetBuildServiceBuildResultLogResultOutput struct{ *pulumi.OutputState }

func (GetBuildServiceBuildResultLogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBuildServiceBuildResultLogResult)(nil)).Elem()
}

func (o GetBuildServiceBuildResultLogResultOutput) ToGetBuildServiceBuildResultLogResultOutput() GetBuildServiceBuildResultLogResultOutput {
	return o
}

func (o GetBuildServiceBuildResultLogResultOutput) ToGetBuildServiceBuildResultLogResultOutputWithContext(ctx context.Context) GetBuildServiceBuildResultLogResultOutput {
	return o
}

// The public download URL of this build result log
func (o GetBuildServiceBuildResultLogResultOutput) BlobUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBuildServiceBuildResultLogResult) *string { return v.BlobUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBuildServiceBuildResultLogResultOutput{})
}
