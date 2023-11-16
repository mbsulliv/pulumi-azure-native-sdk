// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a bookmark.
func LookupBookmark(ctx *pulumi.Context, args *LookupBookmarkArgs, opts ...pulumi.InvokeOption) (*LookupBookmarkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBookmarkResult
	err := ctx.Invoke("azure-native:securityinsights/v20230801preview:getBookmark", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBookmarkArgs struct {
	// Bookmark ID
	BookmarkId string `pulumi:"bookmarkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents a bookmark in Azure Security Insights.
type LookupBookmarkResult struct {
	// The time the bookmark was created
	Created *string `pulumi:"created"`
	// Describes a user that created the bookmark
	CreatedBy *UserInfoResponse `pulumi:"createdBy"`
	// The display name of the bookmark
	DisplayName string `pulumi:"displayName"`
	// Describes the entity mappings of the bookmark
	EntityMappings []BookmarkEntityMappingsResponse `pulumi:"entityMappings"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The bookmark event time
	EventTime *string `pulumi:"eventTime"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Describes an incident that relates to bookmark
	IncidentInfo *IncidentInfoResponse `pulumi:"incidentInfo"`
	// List of labels relevant to this bookmark
	Labels []string `pulumi:"labels"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The notes of the bookmark
	Notes *string `pulumi:"notes"`
	// The query of the bookmark.
	Query string `pulumi:"query"`
	// The end time for the query
	QueryEndTime *string `pulumi:"queryEndTime"`
	// The query result of the bookmark.
	QueryResult *string `pulumi:"queryResult"`
	// The start time for the query
	QueryStartTime *string `pulumi:"queryStartTime"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// A list of relevant mitre attacks
	Tactics []string `pulumi:"tactics"`
	// A list of relevant mitre techniques
	Techniques []string `pulumi:"techniques"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The last time the bookmark was updated
	Updated *string `pulumi:"updated"`
	// Describes a user that updated the bookmark
	UpdatedBy *UserInfoResponse `pulumi:"updatedBy"`
}

func LookupBookmarkOutput(ctx *pulumi.Context, args LookupBookmarkOutputArgs, opts ...pulumi.InvokeOption) LookupBookmarkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBookmarkResult, error) {
			args := v.(LookupBookmarkArgs)
			r, err := LookupBookmark(ctx, &args, opts...)
			var s LookupBookmarkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBookmarkResultOutput)
}

type LookupBookmarkOutputArgs struct {
	// Bookmark ID
	BookmarkId pulumi.StringInput `pulumi:"bookmarkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupBookmarkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBookmarkArgs)(nil)).Elem()
}

// Represents a bookmark in Azure Security Insights.
type LookupBookmarkResultOutput struct{ *pulumi.OutputState }

func (LookupBookmarkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBookmarkResult)(nil)).Elem()
}

func (o LookupBookmarkResultOutput) ToLookupBookmarkResultOutput() LookupBookmarkResultOutput {
	return o
}

func (o LookupBookmarkResultOutput) ToLookupBookmarkResultOutputWithContext(ctx context.Context) LookupBookmarkResultOutput {
	return o
}

// The time the bookmark was created
func (o LookupBookmarkResultOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.Created }).(pulumi.StringPtrOutput)
}

// Describes a user that created the bookmark
func (o LookupBookmarkResultOutput) CreatedBy() UserInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *UserInfoResponse { return v.CreatedBy }).(UserInfoResponsePtrOutput)
}

// The display name of the bookmark
func (o LookupBookmarkResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Describes the entity mappings of the bookmark
func (o LookupBookmarkResultOutput) EntityMappings() BookmarkEntityMappingsResponseArrayOutput {
	return o.ApplyT(func(v LookupBookmarkResult) []BookmarkEntityMappingsResponse { return v.EntityMappings }).(BookmarkEntityMappingsResponseArrayOutput)
}

// Etag of the azure resource
func (o LookupBookmarkResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The bookmark event time
func (o LookupBookmarkResultOutput) EventTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.EventTime }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupBookmarkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkResult) string { return v.Id }).(pulumi.StringOutput)
}

// Describes an incident that relates to bookmark
func (o LookupBookmarkResultOutput) IncidentInfo() IncidentInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *IncidentInfoResponse { return v.IncidentInfo }).(IncidentInfoResponsePtrOutput)
}

// List of labels relevant to this bookmark
func (o LookupBookmarkResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBookmarkResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o LookupBookmarkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The notes of the bookmark
func (o LookupBookmarkResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// The query of the bookmark.
func (o LookupBookmarkResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkResult) string { return v.Query }).(pulumi.StringOutput)
}

// The end time for the query
func (o LookupBookmarkResultOutput) QueryEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.QueryEndTime }).(pulumi.StringPtrOutput)
}

// The query result of the bookmark.
func (o LookupBookmarkResultOutput) QueryResult() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.QueryResult }).(pulumi.StringPtrOutput)
}

// The start time for the query
func (o LookupBookmarkResultOutput) QueryStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.QueryStartTime }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupBookmarkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBookmarkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// A list of relevant mitre attacks
func (o LookupBookmarkResultOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBookmarkResult) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

// A list of relevant mitre techniques
func (o LookupBookmarkResultOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBookmarkResult) []string { return v.Techniques }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBookmarkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkResult) string { return v.Type }).(pulumi.StringOutput)
}

// The last time the bookmark was updated
func (o LookupBookmarkResultOutput) Updated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.Updated }).(pulumi.StringPtrOutput)
}

// Describes a user that updated the bookmark
func (o LookupBookmarkResultOutput) UpdatedBy() UserInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *UserInfoResponse { return v.UpdatedBy }).(UserInfoResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBookmarkResultOutput{})
}
