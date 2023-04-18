// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a watchlist, without its watchlist items.
func LookupWatchlist(ctx *pulumi.Context, args *LookupWatchlistArgs, opts ...pulumi.InvokeOption) (*LookupWatchlistResult, error) {
	var rv LookupWatchlistResult
	err := ctx.Invoke("azure-native:securityinsights/v20230401preview:getWatchlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWatchlistArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Watchlist Alias
	WatchlistAlias string `pulumi:"watchlistAlias"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents a Watchlist in Azure Security Insights.
type LookupWatchlistResult struct {
	// The content type of the raw content. Example : text/csv or text/tsv
	ContentType *string `pulumi:"contentType"`
	// The time the watchlist was created
	Created *string `pulumi:"created"`
	// Describes a user that created the watchlist
	CreatedBy *WatchlistUserInfoResponse `pulumi:"createdBy"`
	// The default duration of a watchlist (in ISO 8601 duration format)
	DefaultDuration *string `pulumi:"defaultDuration"`
	// A description of the watchlist
	Description *string `pulumi:"description"`
	// The display name of the watchlist
	DisplayName string `pulumi:"displayName"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// A flag that indicates if the watchlist is deleted or not
	IsDeleted *bool `pulumi:"isDeleted"`
	// The search key is used to optimize query performance when using watchlists for joins with other data. For example, enable a column with IP addresses to be the designated SearchKey field, then use this field as the key field when joining to other event data by IP address.
	ItemsSearchKey string `pulumi:"itemsSearchKey"`
	// List of labels relevant to this watchlist
	Labels []string `pulumi:"labels"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The number of lines in a csv/tsv content to skip before the header
	NumberOfLinesToSkip *int `pulumi:"numberOfLinesToSkip"`
	// The provider of the watchlist
	Provider string `pulumi:"provider"`
	// The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
	RawContent *string `pulumi:"rawContent"`
	// The filename of the watchlist, called 'source'
	Source *string `pulumi:"source"`
	// The sourceType of the watchlist
	SourceType *string `pulumi:"sourceType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tenantId where the watchlist belongs to
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The last time the watchlist was updated
	Updated *string `pulumi:"updated"`
	// Describes a user that updated the watchlist
	UpdatedBy *WatchlistUserInfoResponse `pulumi:"updatedBy"`
	// The status of the Watchlist upload : New, InProgress or Complete. Pls note : When a Watchlist upload status is equal to InProgress, the Watchlist cannot be deleted
	UploadStatus *string `pulumi:"uploadStatus"`
	// The alias of the watchlist
	WatchlistAlias *string `pulumi:"watchlistAlias"`
	// The id (a Guid) of the watchlist
	WatchlistId *string `pulumi:"watchlistId"`
	// The type of the watchlist
	WatchlistType *string `pulumi:"watchlistType"`
}

func LookupWatchlistOutput(ctx *pulumi.Context, args LookupWatchlistOutputArgs, opts ...pulumi.InvokeOption) LookupWatchlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWatchlistResult, error) {
			args := v.(LookupWatchlistArgs)
			r, err := LookupWatchlist(ctx, &args, opts...)
			var s LookupWatchlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWatchlistResultOutput)
}

type LookupWatchlistOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Watchlist Alias
	WatchlistAlias pulumi.StringInput `pulumi:"watchlistAlias"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWatchlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWatchlistArgs)(nil)).Elem()
}

// Represents a Watchlist in Azure Security Insights.
type LookupWatchlistResultOutput struct{ *pulumi.OutputState }

func (LookupWatchlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWatchlistResult)(nil)).Elem()
}

func (o LookupWatchlistResultOutput) ToLookupWatchlistResultOutput() LookupWatchlistResultOutput {
	return o
}

func (o LookupWatchlistResultOutput) ToLookupWatchlistResultOutputWithContext(ctx context.Context) LookupWatchlistResultOutput {
	return o
}

// The content type of the raw content. Example : text/csv or text/tsv
func (o LookupWatchlistResultOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

// The time the watchlist was created
func (o LookupWatchlistResultOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.Created }).(pulumi.StringPtrOutput)
}

// Describes a user that created the watchlist
func (o LookupWatchlistResultOutput) CreatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *WatchlistUserInfoResponse { return v.CreatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

// The default duration of a watchlist (in ISO 8601 duration format)
func (o LookupWatchlistResultOutput) DefaultDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.DefaultDuration }).(pulumi.StringPtrOutput)
}

// A description of the watchlist
func (o LookupWatchlistResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the watchlist
func (o LookupWatchlistResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o LookupWatchlistResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWatchlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A flag that indicates if the watchlist is deleted or not
func (o LookupWatchlistResultOutput) IsDeleted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *bool { return v.IsDeleted }).(pulumi.BoolPtrOutput)
}

// The search key is used to optimize query performance when using watchlists for joins with other data. For example, enable a column with IP addresses to be the designated SearchKey field, then use this field as the key field when joining to other event data by IP address.
func (o LookupWatchlistResultOutput) ItemsSearchKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistResult) string { return v.ItemsSearchKey }).(pulumi.StringOutput)
}

// List of labels relevant to this watchlist
func (o LookupWatchlistResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWatchlistResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o LookupWatchlistResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistResult) string { return v.Name }).(pulumi.StringOutput)
}

// The number of lines in a csv/tsv content to skip before the header
func (o LookupWatchlistResultOutput) NumberOfLinesToSkip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *int { return v.NumberOfLinesToSkip }).(pulumi.IntPtrOutput)
}

// The provider of the watchlist
func (o LookupWatchlistResultOutput) Provider() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistResult) string { return v.Provider }).(pulumi.StringOutput)
}

// The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
func (o LookupWatchlistResultOutput) RawContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.RawContent }).(pulumi.StringPtrOutput)
}

// The filename of the watchlist, called 'source'
func (o LookupWatchlistResultOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.Source }).(pulumi.StringPtrOutput)
}

// The sourceType of the watchlist
func (o LookupWatchlistResultOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.SourceType }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWatchlistResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWatchlistResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenantId where the watchlist belongs to
func (o LookupWatchlistResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWatchlistResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistResult) string { return v.Type }).(pulumi.StringOutput)
}

// The last time the watchlist was updated
func (o LookupWatchlistResultOutput) Updated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.Updated }).(pulumi.StringPtrOutput)
}

// Describes a user that updated the watchlist
func (o LookupWatchlistResultOutput) UpdatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *WatchlistUserInfoResponse { return v.UpdatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

// The status of the Watchlist upload : New, InProgress or Complete. Pls note : When a Watchlist upload status is equal to InProgress, the Watchlist cannot be deleted
func (o LookupWatchlistResultOutput) UploadStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.UploadStatus }).(pulumi.StringPtrOutput)
}

// The alias of the watchlist
func (o LookupWatchlistResultOutput) WatchlistAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.WatchlistAlias }).(pulumi.StringPtrOutput)
}

// The id (a Guid) of the watchlist
func (o LookupWatchlistResultOutput) WatchlistId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.WatchlistId }).(pulumi.StringPtrOutput)
}

// The type of the watchlist
func (o LookupWatchlistResultOutput) WatchlistType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistResult) *string { return v.WatchlistType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWatchlistResultOutput{})
}
