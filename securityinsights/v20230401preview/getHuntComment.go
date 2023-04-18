// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a hunt comment
func LookupHuntComment(ctx *pulumi.Context, args *LookupHuntCommentArgs, opts ...pulumi.InvokeOption) (*LookupHuntCommentResult, error) {
	var rv LookupHuntCommentResult
	err := ctx.Invoke("azure-native:securityinsights/v20230401preview:getHuntComment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHuntCommentArgs struct {
	// The hunt comment id (GUID)
	HuntCommentId string `pulumi:"huntCommentId"`
	// The hunt id (GUID)
	HuntId string `pulumi:"huntId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents a Hunt Comment in Azure Security Insights
type LookupHuntCommentResult struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The message for the comment
	Message string `pulumi:"message"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupHuntCommentOutput(ctx *pulumi.Context, args LookupHuntCommentOutputArgs, opts ...pulumi.InvokeOption) LookupHuntCommentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHuntCommentResult, error) {
			args := v.(LookupHuntCommentArgs)
			r, err := LookupHuntComment(ctx, &args, opts...)
			var s LookupHuntCommentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHuntCommentResultOutput)
}

type LookupHuntCommentOutputArgs struct {
	// The hunt comment id (GUID)
	HuntCommentId pulumi.StringInput `pulumi:"huntCommentId"`
	// The hunt id (GUID)
	HuntId pulumi.StringInput `pulumi:"huntId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupHuntCommentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHuntCommentArgs)(nil)).Elem()
}

// Represents a Hunt Comment in Azure Security Insights
type LookupHuntCommentResultOutput struct{ *pulumi.OutputState }

func (LookupHuntCommentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHuntCommentResult)(nil)).Elem()
}

func (o LookupHuntCommentResultOutput) ToLookupHuntCommentResultOutput() LookupHuntCommentResultOutput {
	return o
}

func (o LookupHuntCommentResultOutput) ToLookupHuntCommentResultOutputWithContext(ctx context.Context) LookupHuntCommentResultOutput {
	return o
}

// Etag of the azure resource
func (o LookupHuntCommentResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHuntCommentResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupHuntCommentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntCommentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The message for the comment
func (o LookupHuntCommentResultOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntCommentResult) string { return v.Message }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupHuntCommentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntCommentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupHuntCommentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHuntCommentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupHuntCommentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntCommentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHuntCommentResultOutput{})
}
