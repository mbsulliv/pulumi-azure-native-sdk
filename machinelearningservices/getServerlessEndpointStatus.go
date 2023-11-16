// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure REST API version: 2023-08-01-preview.
func GetServerlessEndpointStatus(ctx *pulumi.Context, args *GetServerlessEndpointStatusArgs, opts ...pulumi.InvokeOption) (*GetServerlessEndpointStatusResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetServerlessEndpointStatusResult
	err := ctx.Invoke("azure-native:machinelearningservices:getServerlessEndpointStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetServerlessEndpointStatusArgs struct {
	// Serverless Endpoint name.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type GetServerlessEndpointStatusResult struct {
	// The model-specific metrics from the backing inference endpoint.
	Metrics map[string]string `pulumi:"metrics"`
}

func GetServerlessEndpointStatusOutput(ctx *pulumi.Context, args GetServerlessEndpointStatusOutputArgs, opts ...pulumi.InvokeOption) GetServerlessEndpointStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServerlessEndpointStatusResult, error) {
			args := v.(GetServerlessEndpointStatusArgs)
			r, err := GetServerlessEndpointStatus(ctx, &args, opts...)
			var s GetServerlessEndpointStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServerlessEndpointStatusResultOutput)
}

type GetServerlessEndpointStatusOutputArgs struct {
	// Serverless Endpoint name.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (GetServerlessEndpointStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerlessEndpointStatusArgs)(nil)).Elem()
}

type GetServerlessEndpointStatusResultOutput struct{ *pulumi.OutputState }

func (GetServerlessEndpointStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerlessEndpointStatusResult)(nil)).Elem()
}

func (o GetServerlessEndpointStatusResultOutput) ToGetServerlessEndpointStatusResultOutput() GetServerlessEndpointStatusResultOutput {
	return o
}

func (o GetServerlessEndpointStatusResultOutput) ToGetServerlessEndpointStatusResultOutputWithContext(ctx context.Context) GetServerlessEndpointStatusResultOutput {
	return o
}

// The model-specific metrics from the backing inference endpoint.
func (o GetServerlessEndpointStatusResultOutput) Metrics() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetServerlessEndpointStatusResult) map[string]string { return v.Metrics }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServerlessEndpointStatusResultOutput{})
}
