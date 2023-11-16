// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a bookmark relation.
func LookupBookmarkRelation(ctx *pulumi.Context, args *LookupBookmarkRelationArgs, opts ...pulumi.InvokeOption) (*LookupBookmarkRelationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBookmarkRelationResult
	err := ctx.Invoke("azure-native:securityinsights/v20230601preview:getBookmarkRelation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBookmarkRelationArgs struct {
	// Bookmark ID
	BookmarkId string `pulumi:"bookmarkId"`
	// Relation Name
	RelationName string `pulumi:"relationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents a relation between two resources
type LookupBookmarkRelationResult struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource ID of the related resource
	RelatedResourceId string `pulumi:"relatedResourceId"`
	// The resource kind of the related resource
	RelatedResourceKind string `pulumi:"relatedResourceKind"`
	// The name of the related resource
	RelatedResourceName string `pulumi:"relatedResourceName"`
	// The resource type of the related resource
	RelatedResourceType string `pulumi:"relatedResourceType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupBookmarkRelationOutput(ctx *pulumi.Context, args LookupBookmarkRelationOutputArgs, opts ...pulumi.InvokeOption) LookupBookmarkRelationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBookmarkRelationResult, error) {
			args := v.(LookupBookmarkRelationArgs)
			r, err := LookupBookmarkRelation(ctx, &args, opts...)
			var s LookupBookmarkRelationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBookmarkRelationResultOutput)
}

type LookupBookmarkRelationOutputArgs struct {
	// Bookmark ID
	BookmarkId pulumi.StringInput `pulumi:"bookmarkId"`
	// Relation Name
	RelationName pulumi.StringInput `pulumi:"relationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupBookmarkRelationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBookmarkRelationArgs)(nil)).Elem()
}

// Represents a relation between two resources
type LookupBookmarkRelationResultOutput struct{ *pulumi.OutputState }

func (LookupBookmarkRelationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBookmarkRelationResult)(nil)).Elem()
}

func (o LookupBookmarkRelationResultOutput) ToLookupBookmarkRelationResultOutput() LookupBookmarkRelationResultOutput {
	return o
}

func (o LookupBookmarkRelationResultOutput) ToLookupBookmarkRelationResultOutputWithContext(ctx context.Context) LookupBookmarkRelationResultOutput {
	return o
}

// Etag of the azure resource
func (o LookupBookmarkRelationResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupBookmarkRelationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupBookmarkRelationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource ID of the related resource
func (o LookupBookmarkRelationResultOutput) RelatedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.RelatedResourceId }).(pulumi.StringOutput)
}

// The resource kind of the related resource
func (o LookupBookmarkRelationResultOutput) RelatedResourceKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.RelatedResourceKind }).(pulumi.StringOutput)
}

// The name of the related resource
func (o LookupBookmarkRelationResultOutput) RelatedResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.RelatedResourceName }).(pulumi.StringOutput)
}

// The resource type of the related resource
func (o LookupBookmarkRelationResultOutput) RelatedResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.RelatedResourceType }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupBookmarkRelationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBookmarkRelationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBookmarkRelationResultOutput{})
}
