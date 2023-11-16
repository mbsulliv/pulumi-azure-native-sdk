// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns supported virtual machine sizes for resize
func GetComputeAllowedResizeSizes(ctx *pulumi.Context, args *GetComputeAllowedResizeSizesArgs, opts ...pulumi.InvokeOption) (*GetComputeAllowedResizeSizesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetComputeAllowedResizeSizesResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230801preview:getComputeAllowedResizeSizes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetComputeAllowedResizeSizesArgs struct {
	// Name of the Azure Machine Learning compute.
	ComputeName string `pulumi:"computeName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The List Virtual Machine size operation response.
type GetComputeAllowedResizeSizesResult struct {
	// The list of virtual machine sizes supported by AmlCompute.
	Value []VirtualMachineSizeResponse `pulumi:"value"`
}

func GetComputeAllowedResizeSizesOutput(ctx *pulumi.Context, args GetComputeAllowedResizeSizesOutputArgs, opts ...pulumi.InvokeOption) GetComputeAllowedResizeSizesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetComputeAllowedResizeSizesResult, error) {
			args := v.(GetComputeAllowedResizeSizesArgs)
			r, err := GetComputeAllowedResizeSizes(ctx, &args, opts...)
			var s GetComputeAllowedResizeSizesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetComputeAllowedResizeSizesResultOutput)
}

type GetComputeAllowedResizeSizesOutputArgs struct {
	// Name of the Azure Machine Learning compute.
	ComputeName pulumi.StringInput `pulumi:"computeName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (GetComputeAllowedResizeSizesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetComputeAllowedResizeSizesArgs)(nil)).Elem()
}

// The List Virtual Machine size operation response.
type GetComputeAllowedResizeSizesResultOutput struct{ *pulumi.OutputState }

func (GetComputeAllowedResizeSizesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetComputeAllowedResizeSizesResult)(nil)).Elem()
}

func (o GetComputeAllowedResizeSizesResultOutput) ToGetComputeAllowedResizeSizesResultOutput() GetComputeAllowedResizeSizesResultOutput {
	return o
}

func (o GetComputeAllowedResizeSizesResultOutput) ToGetComputeAllowedResizeSizesResultOutputWithContext(ctx context.Context) GetComputeAllowedResizeSizesResultOutput {
	return o
}

// The list of virtual machine sizes supported by AmlCompute.
func (o GetComputeAllowedResizeSizesResultOutput) Value() VirtualMachineSizeResponseArrayOutput {
	return o.ApplyT(func(v GetComputeAllowedResizeSizesResult) []VirtualMachineSizeResponse { return v.Value }).(VirtualMachineSizeResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetComputeAllowedResizeSizesResultOutput{})
}
