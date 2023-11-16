// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets secrets related to Machine Learning compute (storage keys, service credentials, etc).
// Azure REST API version: 2023-04-01.
//
// Other available API versions: 2022-01-01-preview, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01.
func ListComputeKeys(ctx *pulumi.Context, args *ListComputeKeysArgs, opts ...pulumi.InvokeOption) (*ListComputeKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListComputeKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices:listComputeKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListComputeKeysArgs struct {
	// Name of the Azure Machine Learning compute.
	ComputeName string `pulumi:"computeName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Secrets related to a Machine Learning compute. Might differ for every type of compute.
type ListComputeKeysResult struct {
	// The type of compute
	ComputeType string `pulumi:"computeType"`
}

func ListComputeKeysOutput(ctx *pulumi.Context, args ListComputeKeysOutputArgs, opts ...pulumi.InvokeOption) ListComputeKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListComputeKeysResult, error) {
			args := v.(ListComputeKeysArgs)
			r, err := ListComputeKeys(ctx, &args, opts...)
			var s ListComputeKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListComputeKeysResultOutput)
}

type ListComputeKeysOutputArgs struct {
	// Name of the Azure Machine Learning compute.
	ComputeName pulumi.StringInput `pulumi:"computeName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListComputeKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListComputeKeysArgs)(nil)).Elem()
}

// Secrets related to a Machine Learning compute. Might differ for every type of compute.
type ListComputeKeysResultOutput struct{ *pulumi.OutputState }

func (ListComputeKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListComputeKeysResult)(nil)).Elem()
}

func (o ListComputeKeysResultOutput) ToListComputeKeysResultOutput() ListComputeKeysResultOutput {
	return o
}

func (o ListComputeKeysResultOutput) ToListComputeKeysResultOutputWithContext(ctx context.Context) ListComputeKeysResultOutput {
	return o
}

// The type of compute
func (o ListComputeKeysResultOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v ListComputeKeysResult) string { return v.ComputeType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListComputeKeysResultOutput{})
}
