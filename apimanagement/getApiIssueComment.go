// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the issue Comment for an API specified by its identifier.
// Azure REST API version: 2022-08-01.
func LookupApiIssueComment(ctx *pulumi.Context, args *LookupApiIssueCommentArgs, opts ...pulumi.InvokeOption) (*LookupApiIssueCommentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApiIssueCommentResult
	err := ctx.Invoke("azure-native:apimanagement:getApiIssueComment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiIssueCommentArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Comment identifier within an Issue. Must be unique in the current Issue.
	CommentId string `pulumi:"commentId"`
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId string `pulumi:"issueId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Issue Comment Contract details.
type LookupApiIssueCommentResult struct {
	// Date and time when the comment was created.
	CreatedDate *string `pulumi:"createdDate"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Comment text.
	Text string `pulumi:"text"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// A resource identifier for the user who left the comment.
	UserId string `pulumi:"userId"`
}

func LookupApiIssueCommentOutput(ctx *pulumi.Context, args LookupApiIssueCommentOutputArgs, opts ...pulumi.InvokeOption) LookupApiIssueCommentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiIssueCommentResult, error) {
			args := v.(LookupApiIssueCommentArgs)
			r, err := LookupApiIssueComment(ctx, &args, opts...)
			var s LookupApiIssueCommentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiIssueCommentResultOutput)
}

type LookupApiIssueCommentOutputArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// Comment identifier within an Issue. Must be unique in the current Issue.
	CommentId pulumi.StringInput `pulumi:"commentId"`
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId pulumi.StringInput `pulumi:"issueId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiIssueCommentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiIssueCommentArgs)(nil)).Elem()
}

// Issue Comment Contract details.
type LookupApiIssueCommentResultOutput struct{ *pulumi.OutputState }

func (LookupApiIssueCommentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiIssueCommentResult)(nil)).Elem()
}

func (o LookupApiIssueCommentResultOutput) ToLookupApiIssueCommentResultOutput() LookupApiIssueCommentResultOutput {
	return o
}

func (o LookupApiIssueCommentResultOutput) ToLookupApiIssueCommentResultOutputWithContext(ctx context.Context) LookupApiIssueCommentResultOutput {
	return o
}

func (o LookupApiIssueCommentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupApiIssueCommentResult] {
	return pulumix.Output[LookupApiIssueCommentResult]{
		OutputState: o.OutputState,
	}
}

// Date and time when the comment was created.
func (o LookupApiIssueCommentResultOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiIssueCommentResult) *string { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApiIssueCommentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueCommentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupApiIssueCommentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueCommentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Comment text.
func (o LookupApiIssueCommentResultOutput) Text() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueCommentResult) string { return v.Text }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApiIssueCommentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueCommentResult) string { return v.Type }).(pulumi.StringOutput)
}

// A resource identifier for the user who left the comment.
func (o LookupApiIssueCommentResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueCommentResult) string { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiIssueCommentResultOutput{})
}
