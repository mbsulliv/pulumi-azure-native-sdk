// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the callback url for a trigger of a workflow version.
// Azure REST API version: 2016-06-01.
func ListWorkflowVersionCallbackUrl(ctx *pulumi.Context, args *ListWorkflowVersionCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListWorkflowVersionCallbackUrlResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWorkflowVersionCallbackUrlResult
	err := ctx.Invoke("azure-native:logic:listWorkflowVersionCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowVersionCallbackUrlArgs struct {
	// The key type.
	KeyType *KeyType `pulumi:"keyType"`
	// The expiry time.
	NotAfter *string `pulumi:"notAfter"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workflow trigger name.
	TriggerName string `pulumi:"triggerName"`
	// The workflow versionId.
	VersionId string `pulumi:"versionId"`
	// The workflow name.
	WorkflowName string `pulumi:"workflowName"`
}

// The workflow trigger callback URL.
type ListWorkflowVersionCallbackUrlResult struct {
	// Gets the workflow trigger callback URL base path.
	BasePath string `pulumi:"basePath"`
	// Gets the workflow trigger callback URL HTTP method.
	Method string `pulumi:"method"`
	// Gets the workflow trigger callback URL query parameters.
	Queries *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	// Gets the workflow trigger callback URL relative path.
	RelativePath string `pulumi:"relativePath"`
	// Gets the workflow trigger callback URL relative path parameters.
	RelativePathParameters []string `pulumi:"relativePathParameters"`
	// Gets the workflow trigger callback URL.
	Value string `pulumi:"value"`
}

func ListWorkflowVersionCallbackUrlOutput(ctx *pulumi.Context, args ListWorkflowVersionCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListWorkflowVersionCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkflowVersionCallbackUrlResult, error) {
			args := v.(ListWorkflowVersionCallbackUrlArgs)
			r, err := ListWorkflowVersionCallbackUrl(ctx, &args, opts...)
			var s ListWorkflowVersionCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkflowVersionCallbackUrlResultOutput)
}

type ListWorkflowVersionCallbackUrlOutputArgs struct {
	// The key type.
	KeyType KeyTypePtrInput `pulumi:"keyType"`
	// The expiry time.
	NotAfter pulumi.StringPtrInput `pulumi:"notAfter"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The workflow trigger name.
	TriggerName pulumi.StringInput `pulumi:"triggerName"`
	// The workflow versionId.
	VersionId pulumi.StringInput `pulumi:"versionId"`
	// The workflow name.
	WorkflowName pulumi.StringInput `pulumi:"workflowName"`
}

func (ListWorkflowVersionCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowVersionCallbackUrlArgs)(nil)).Elem()
}

// The workflow trigger callback URL.
type ListWorkflowVersionCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListWorkflowVersionCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowVersionCallbackUrlResult)(nil)).Elem()
}

func (o ListWorkflowVersionCallbackUrlResultOutput) ToListWorkflowVersionCallbackUrlResultOutput() ListWorkflowVersionCallbackUrlResultOutput {
	return o
}

func (o ListWorkflowVersionCallbackUrlResultOutput) ToListWorkflowVersionCallbackUrlResultOutputWithContext(ctx context.Context) ListWorkflowVersionCallbackUrlResultOutput {
	return o
}

// Gets the workflow trigger callback URL base path.
func (o ListWorkflowVersionCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowVersionCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

// Gets the workflow trigger callback URL HTTP method.
func (o ListWorkflowVersionCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowVersionCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

// Gets the workflow trigger callback URL query parameters.
func (o ListWorkflowVersionCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListWorkflowVersionCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

// Gets the workflow trigger callback URL relative path.
func (o ListWorkflowVersionCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowVersionCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

// Gets the workflow trigger callback URL relative path parameters.
func (o ListWorkflowVersionCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWorkflowVersionCallbackUrlResult) []string { return v.RelativePathParameters }).(pulumi.StringArrayOutput)
}

// Gets the workflow trigger callback URL.
func (o ListWorkflowVersionCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowVersionCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkflowVersionCallbackUrlResultOutput{})
}