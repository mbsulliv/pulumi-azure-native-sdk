// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Lists a workflow run expression trace.
// Azure REST API version: 2019-05-01.
func ListWorkflowRunActionRepetitionExpressionTraces(ctx *pulumi.Context, args *ListWorkflowRunActionRepetitionExpressionTracesArgs, opts ...pulumi.InvokeOption) (*ListWorkflowRunActionRepetitionExpressionTracesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWorkflowRunActionRepetitionExpressionTracesResult
	err := ctx.Invoke("azure-native:logic:listWorkflowRunActionRepetitionExpressionTraces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowRunActionRepetitionExpressionTracesArgs struct {
	// The workflow action name.
	ActionName string `pulumi:"actionName"`
	// The workflow repetition.
	RepetitionName string `pulumi:"repetitionName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workflow run name.
	RunName string `pulumi:"runName"`
	// The workflow name.
	WorkflowName string `pulumi:"workflowName"`
}

// The expression traces.
type ListWorkflowRunActionRepetitionExpressionTracesResult struct {
	Inputs []ExpressionRootResponse `pulumi:"inputs"`
}

func ListWorkflowRunActionRepetitionExpressionTracesOutput(ctx *pulumi.Context, args ListWorkflowRunActionRepetitionExpressionTracesOutputArgs, opts ...pulumi.InvokeOption) ListWorkflowRunActionRepetitionExpressionTracesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkflowRunActionRepetitionExpressionTracesResult, error) {
			args := v.(ListWorkflowRunActionRepetitionExpressionTracesArgs)
			r, err := ListWorkflowRunActionRepetitionExpressionTraces(ctx, &args, opts...)
			var s ListWorkflowRunActionRepetitionExpressionTracesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkflowRunActionRepetitionExpressionTracesResultOutput)
}

type ListWorkflowRunActionRepetitionExpressionTracesOutputArgs struct {
	// The workflow action name.
	ActionName pulumi.StringInput `pulumi:"actionName"`
	// The workflow repetition.
	RepetitionName pulumi.StringInput `pulumi:"repetitionName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The workflow run name.
	RunName pulumi.StringInput `pulumi:"runName"`
	// The workflow name.
	WorkflowName pulumi.StringInput `pulumi:"workflowName"`
}

func (ListWorkflowRunActionRepetitionExpressionTracesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowRunActionRepetitionExpressionTracesArgs)(nil)).Elem()
}

// The expression traces.
type ListWorkflowRunActionRepetitionExpressionTracesResultOutput struct{ *pulumi.OutputState }

func (ListWorkflowRunActionRepetitionExpressionTracesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowRunActionRepetitionExpressionTracesResult)(nil)).Elem()
}

func (o ListWorkflowRunActionRepetitionExpressionTracesResultOutput) ToListWorkflowRunActionRepetitionExpressionTracesResultOutput() ListWorkflowRunActionRepetitionExpressionTracesResultOutput {
	return o
}

func (o ListWorkflowRunActionRepetitionExpressionTracesResultOutput) ToListWorkflowRunActionRepetitionExpressionTracesResultOutputWithContext(ctx context.Context) ListWorkflowRunActionRepetitionExpressionTracesResultOutput {
	return o
}

func (o ListWorkflowRunActionRepetitionExpressionTracesResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWorkflowRunActionRepetitionExpressionTracesResult] {
	return pulumix.Output[ListWorkflowRunActionRepetitionExpressionTracesResult]{
		OutputState: o.OutputState,
	}
}

func (o ListWorkflowRunActionRepetitionExpressionTracesResultOutput) Inputs() ExpressionRootResponseArrayOutput {
	return o.ApplyT(func(v ListWorkflowRunActionRepetitionExpressionTracesResult) []ExpressionRootResponse {
		return v.Inputs
	}).(ExpressionRootResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkflowRunActionRepetitionExpressionTracesResultOutput{})
}
