// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the details (e.g IP address, port etc) of all the compute nodes in the compute.
func ListComputeNodes(ctx *pulumi.Context, args *ListComputeNodesArgs, opts ...pulumi.InvokeOption) (*ListComputeNodesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListComputeNodesResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401:listComputeNodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListComputeNodesArgs struct {
	// Name of the Azure Machine Learning compute.
	ComputeName string `pulumi:"computeName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Result of AmlCompute Nodes
type ListComputeNodesResult struct {
	// The continuation token.
	NextLink string `pulumi:"nextLink"`
	// The collection of returned AmlCompute nodes details.
	Nodes []AmlComputeNodeInformationResponse `pulumi:"nodes"`
}

func ListComputeNodesOutput(ctx *pulumi.Context, args ListComputeNodesOutputArgs, opts ...pulumi.InvokeOption) ListComputeNodesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListComputeNodesResult, error) {
			args := v.(ListComputeNodesArgs)
			r, err := ListComputeNodes(ctx, &args, opts...)
			var s ListComputeNodesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListComputeNodesResultOutput)
}

type ListComputeNodesOutputArgs struct {
	// Name of the Azure Machine Learning compute.
	ComputeName pulumi.StringInput `pulumi:"computeName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListComputeNodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListComputeNodesArgs)(nil)).Elem()
}

// Result of AmlCompute Nodes
type ListComputeNodesResultOutput struct{ *pulumi.OutputState }

func (ListComputeNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListComputeNodesResult)(nil)).Elem()
}

func (o ListComputeNodesResultOutput) ToListComputeNodesResultOutput() ListComputeNodesResultOutput {
	return o
}

func (o ListComputeNodesResultOutput) ToListComputeNodesResultOutputWithContext(ctx context.Context) ListComputeNodesResultOutput {
	return o
}

func (o ListComputeNodesResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListComputeNodesResult] {
	return pulumix.Output[ListComputeNodesResult]{
		OutputState: o.OutputState,
	}
}

// The continuation token.
func (o ListComputeNodesResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListComputeNodesResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// The collection of returned AmlCompute nodes details.
func (o ListComputeNodesResultOutput) Nodes() AmlComputeNodeInformationResponseArrayOutput {
	return o.ApplyT(func(v ListComputeNodesResult) []AmlComputeNodeInformationResponse { return v.Nodes }).(AmlComputeNodeInformationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListComputeNodesResultOutput{})
}
