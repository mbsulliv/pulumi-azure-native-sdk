// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get remote debugging config.
func GetDeploymentRemoteDebuggingConfig(ctx *pulumi.Context, args *GetDeploymentRemoteDebuggingConfigArgs, opts ...pulumi.InvokeOption) (*GetDeploymentRemoteDebuggingConfigResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetDeploymentRemoteDebuggingConfigResult
	err := ctx.Invoke("azure-native:appplatform/v20230501preview:getDeploymentRemoteDebuggingConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDeploymentRemoteDebuggingConfigArgs struct {
	// The name of the App resource.
	AppName string `pulumi:"appName"`
	// The name of the Deployment resource.
	DeploymentName string `pulumi:"deploymentName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Remote debugging config.
type GetDeploymentRemoteDebuggingConfigResult struct {
	// Indicate if remote debugging is enabled
	Enabled *bool `pulumi:"enabled"`
	// Application debugging port
	Port *int `pulumi:"port"`
}

func GetDeploymentRemoteDebuggingConfigOutput(ctx *pulumi.Context, args GetDeploymentRemoteDebuggingConfigOutputArgs, opts ...pulumi.InvokeOption) GetDeploymentRemoteDebuggingConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDeploymentRemoteDebuggingConfigResult, error) {
			args := v.(GetDeploymentRemoteDebuggingConfigArgs)
			r, err := GetDeploymentRemoteDebuggingConfig(ctx, &args, opts...)
			var s GetDeploymentRemoteDebuggingConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDeploymentRemoteDebuggingConfigResultOutput)
}

type GetDeploymentRemoteDebuggingConfigOutputArgs struct {
	// The name of the App resource.
	AppName pulumi.StringInput `pulumi:"appName"`
	// The name of the Deployment resource.
	DeploymentName pulumi.StringInput `pulumi:"deploymentName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetDeploymentRemoteDebuggingConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeploymentRemoteDebuggingConfigArgs)(nil)).Elem()
}

// Remote debugging config.
type GetDeploymentRemoteDebuggingConfigResultOutput struct{ *pulumi.OutputState }

func (GetDeploymentRemoteDebuggingConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeploymentRemoteDebuggingConfigResult)(nil)).Elem()
}

func (o GetDeploymentRemoteDebuggingConfigResultOutput) ToGetDeploymentRemoteDebuggingConfigResultOutput() GetDeploymentRemoteDebuggingConfigResultOutput {
	return o
}

func (o GetDeploymentRemoteDebuggingConfigResultOutput) ToGetDeploymentRemoteDebuggingConfigResultOutputWithContext(ctx context.Context) GetDeploymentRemoteDebuggingConfigResultOutput {
	return o
}

func (o GetDeploymentRemoteDebuggingConfigResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetDeploymentRemoteDebuggingConfigResult] {
	return pulumix.Output[GetDeploymentRemoteDebuggingConfigResult]{
		OutputState: o.OutputState,
	}
}

// Indicate if remote debugging is enabled
func (o GetDeploymentRemoteDebuggingConfigResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetDeploymentRemoteDebuggingConfigResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Application debugging port
func (o GetDeploymentRemoteDebuggingConfigResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDeploymentRemoteDebuggingConfigResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDeploymentRemoteDebuggingConfigResultOutput{})
}
