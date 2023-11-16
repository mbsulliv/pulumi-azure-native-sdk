// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the Issue for an API specified by its identifier.
func LookupApiIssue(ctx *pulumi.Context, args *LookupApiIssueArgs, opts ...pulumi.InvokeOption) (*LookupApiIssueResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApiIssueResult
	err := ctx.Invoke("azure-native:apimanagement/v20220801:getApiIssue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiIssueArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Expand the comment attachments.
	ExpandCommentsAttachments *bool `pulumi:"expandCommentsAttachments"`
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId string `pulumi:"issueId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Issue Contract details.
type LookupApiIssueResult struct {
	// A resource identifier for the API the issue was created for.
	ApiId *string `pulumi:"apiId"`
	// Date and time when the issue was created.
	CreatedDate *string `pulumi:"createdDate"`
	// Text describing the issue.
	Description string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Status of the issue.
	State *string `pulumi:"state"`
	// The issue title.
	Title string `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// A resource identifier for the user created the issue.
	UserId string `pulumi:"userId"`
}

func LookupApiIssueOutput(ctx *pulumi.Context, args LookupApiIssueOutputArgs, opts ...pulumi.InvokeOption) LookupApiIssueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiIssueResult, error) {
			args := v.(LookupApiIssueArgs)
			r, err := LookupApiIssue(ctx, &args, opts...)
			var s LookupApiIssueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiIssueResultOutput)
}

type LookupApiIssueOutputArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// Expand the comment attachments.
	ExpandCommentsAttachments pulumi.BoolPtrInput `pulumi:"expandCommentsAttachments"`
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId pulumi.StringInput `pulumi:"issueId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiIssueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiIssueArgs)(nil)).Elem()
}

// Issue Contract details.
type LookupApiIssueResultOutput struct{ *pulumi.OutputState }

func (LookupApiIssueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiIssueResult)(nil)).Elem()
}

func (o LookupApiIssueResultOutput) ToLookupApiIssueResultOutput() LookupApiIssueResultOutput {
	return o
}

func (o LookupApiIssueResultOutput) ToLookupApiIssueResultOutputWithContext(ctx context.Context) LookupApiIssueResultOutput {
	return o
}

// A resource identifier for the API the issue was created for.
func (o LookupApiIssueResultOutput) ApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiIssueResult) *string { return v.ApiId }).(pulumi.StringPtrOutput)
}

// Date and time when the issue was created.
func (o LookupApiIssueResultOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiIssueResult) *string { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

// Text describing the issue.
func (o LookupApiIssueResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueResult) string { return v.Description }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApiIssueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupApiIssueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueResult) string { return v.Name }).(pulumi.StringOutput)
}

// Status of the issue.
func (o LookupApiIssueResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiIssueResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// The issue title.
func (o LookupApiIssueResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueResult) string { return v.Title }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApiIssueResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueResult) string { return v.Type }).(pulumi.StringOutput)
}

// A resource identifier for the user created the issue.
func (o LookupApiIssueResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueResult) string { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiIssueResultOutput{})
}
