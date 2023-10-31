// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get deployment log file URL
func GetDeploymentLogFileUrl(ctx *pulumi.Context, args *GetDeploymentLogFileUrlArgs, opts ...pulumi.InvokeOption) (*GetDeploymentLogFileUrlResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetDeploymentLogFileUrlResult
	err := ctx.Invoke("azure-native:appplatform/v20231101preview:getDeploymentLogFileUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDeploymentLogFileUrlArgs struct {
	// The name of the App resource.
	AppName string `pulumi:"appName"`
	// The name of the Deployment resource.
	DeploymentName string `pulumi:"deploymentName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Log file URL payload
type GetDeploymentLogFileUrlResult struct {
	// URL of the log file
	Url string `pulumi:"url"`
}

func GetDeploymentLogFileUrlOutput(ctx *pulumi.Context, args GetDeploymentLogFileUrlOutputArgs, opts ...pulumi.InvokeOption) GetDeploymentLogFileUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDeploymentLogFileUrlResult, error) {
			args := v.(GetDeploymentLogFileUrlArgs)
			r, err := GetDeploymentLogFileUrl(ctx, &args, opts...)
			var s GetDeploymentLogFileUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDeploymentLogFileUrlResultOutput)
}

type GetDeploymentLogFileUrlOutputArgs struct {
	// The name of the App resource.
	AppName pulumi.StringInput `pulumi:"appName"`
	// The name of the Deployment resource.
	DeploymentName pulumi.StringInput `pulumi:"deploymentName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetDeploymentLogFileUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeploymentLogFileUrlArgs)(nil)).Elem()
}

// Log file URL payload
type GetDeploymentLogFileUrlResultOutput struct{ *pulumi.OutputState }

func (GetDeploymentLogFileUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeploymentLogFileUrlResult)(nil)).Elem()
}

func (o GetDeploymentLogFileUrlResultOutput) ToGetDeploymentLogFileUrlResultOutput() GetDeploymentLogFileUrlResultOutput {
	return o
}

func (o GetDeploymentLogFileUrlResultOutput) ToGetDeploymentLogFileUrlResultOutputWithContext(ctx context.Context) GetDeploymentLogFileUrlResultOutput {
	return o
}

func (o GetDeploymentLogFileUrlResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetDeploymentLogFileUrlResult] {
	return pulumix.Output[GetDeploymentLogFileUrlResult]{
		OutputState: o.OutputState,
	}
}

// URL of the log file
func (o GetDeploymentLogFileUrlResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeploymentLogFileUrlResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDeploymentLogFileUrlResultOutput{})
}