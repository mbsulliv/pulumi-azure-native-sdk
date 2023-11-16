// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists a workflow run expression trace.
// Azure REST API version: 2019-05-01.
//
// Other available API versions: 2016-06-01, 2018-07-01-preview.
func ListWorkflowRunActionExpressionTraces(ctx *pulumi.Context, args *ListWorkflowRunActionExpressionTracesArgs, opts ...pulumi.InvokeOption) (*ListWorkflowRunActionExpressionTracesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWorkflowRunActionExpressionTracesResult
	err := ctx.Invoke("azure-native:logic:listWorkflowRunActionExpressionTraces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowRunActionExpressionTracesArgs struct {
	// The workflow action name.
	ActionName string `pulumi:"actionName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workflow run name.
	RunName string `pulumi:"runName"`
	// The workflow name.
	WorkflowName string `pulumi:"workflowName"`
}

// The expression traces.
type ListWorkflowRunActionExpressionTracesResult struct {
	Inputs []ExpressionRootResponse `pulumi:"inputs"`
}

func ListWorkflowRunActionExpressionTracesOutput(ctx *pulumi.Context, args ListWorkflowRunActionExpressionTracesOutputArgs, opts ...pulumi.InvokeOption) ListWorkflowRunActionExpressionTracesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkflowRunActionExpressionTracesResult, error) {
			args := v.(ListWorkflowRunActionExpressionTracesArgs)
			r, err := ListWorkflowRunActionExpressionTraces(ctx, &args, opts...)
			var s ListWorkflowRunActionExpressionTracesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkflowRunActionExpressionTracesResultOutput)
}

type ListWorkflowRunActionExpressionTracesOutputArgs struct {
	// The workflow action name.
	ActionName pulumi.StringInput `pulumi:"actionName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The workflow run name.
	RunName pulumi.StringInput `pulumi:"runName"`
	// The workflow name.
	WorkflowName pulumi.StringInput `pulumi:"workflowName"`
}

func (ListWorkflowRunActionExpressionTracesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowRunActionExpressionTracesArgs)(nil)).Elem()
}

// The expression traces.
type ListWorkflowRunActionExpressionTracesResultOutput struct{ *pulumi.OutputState }

func (ListWorkflowRunActionExpressionTracesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowRunActionExpressionTracesResult)(nil)).Elem()
}

func (o ListWorkflowRunActionExpressionTracesResultOutput) ToListWorkflowRunActionExpressionTracesResultOutput() ListWorkflowRunActionExpressionTracesResultOutput {
	return o
}

func (o ListWorkflowRunActionExpressionTracesResultOutput) ToListWorkflowRunActionExpressionTracesResultOutputWithContext(ctx context.Context) ListWorkflowRunActionExpressionTracesResultOutput {
	return o
}

func (o ListWorkflowRunActionExpressionTracesResultOutput) Inputs() ExpressionRootResponseArrayOutput {
	return o.ApplyT(func(v ListWorkflowRunActionExpressionTracesResult) []ExpressionRootResponse { return v.Inputs }).(ExpressionRootResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkflowRunActionExpressionTracesResultOutput{})
}
