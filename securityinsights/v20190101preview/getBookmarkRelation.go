// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

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
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getBookmarkRelation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBookmarkRelationArgs struct {
	// Bookmark ID
	BookmarkId string `pulumi:"bookmarkId"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// Relation Name
	RelationName string `pulumi:"relationName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents a relation between two resources
type LookupBookmarkRelationResult struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Azure resource Id
	Id string `pulumi:"id"`
	// Azure resource name
	Name string `pulumi:"name"`
	// The resource ID of the related resource
	RelatedResourceId string `pulumi:"relatedResourceId"`
	// The resource kind of the related resource
	RelatedResourceKind string `pulumi:"relatedResourceKind"`
	// The name of the related resource
	RelatedResourceName string `pulumi:"relatedResourceName"`
	// The resource type of the related resource
	RelatedResourceType string `pulumi:"relatedResourceType"`
	// Azure resource type
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
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput `pulumi:"operationalInsightsResourceProvider"`
	// Relation Name
	RelationName pulumi.StringInput `pulumi:"relationName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
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

// Azure resource Id
func (o LookupBookmarkRelationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Azure resource name
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

// Azure resource type
func (o LookupBookmarkRelationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBookmarkRelationResultOutput{})
}
