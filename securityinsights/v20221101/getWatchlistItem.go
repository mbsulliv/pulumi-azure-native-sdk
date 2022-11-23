// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Watchlist Item in Azure Security Insights.
func LookupWatchlistItem(ctx *pulumi.Context, args *LookupWatchlistItemArgs, opts ...pulumi.InvokeOption) (*LookupWatchlistItemResult, error) {
	var rv LookupWatchlistItemResult
	err := ctx.Invoke("azure-native:securityinsights/v20221101:getWatchlistItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWatchlistItemArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The watchlist alias
	WatchlistAlias string `pulumi:"watchlistAlias"`
	// The watchlist item id (GUID)
	WatchlistItemId string `pulumi:"watchlistItemId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents a Watchlist Item in Azure Security Insights.
type LookupWatchlistItemResult struct {
	// The time the watchlist item was created
	Created *string `pulumi:"created"`
	// Describes a user that created the watchlist item
	CreatedBy *WatchlistUserInfoResponse `pulumi:"createdBy"`
	// key-value pairs for a watchlist item entity mapping
	EntityMapping interface{} `pulumi:"entityMapping"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// A flag that indicates if the watchlist item is deleted or not
	IsDeleted *bool `pulumi:"isDeleted"`
	// key-value pairs for a watchlist item
	ItemsKeyValue interface{} `pulumi:"itemsKeyValue"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tenantId to which the watchlist item belongs to
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The last time the watchlist item was updated
	Updated *string `pulumi:"updated"`
	// Describes a user that updated the watchlist item
	UpdatedBy *WatchlistUserInfoResponse `pulumi:"updatedBy"`
	// The id (a Guid) of the watchlist item
	WatchlistItemId *string `pulumi:"watchlistItemId"`
	// The type of the watchlist item
	WatchlistItemType *string `pulumi:"watchlistItemType"`
}

func LookupWatchlistItemOutput(ctx *pulumi.Context, args LookupWatchlistItemOutputArgs, opts ...pulumi.InvokeOption) LookupWatchlistItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWatchlistItemResult, error) {
			args := v.(LookupWatchlistItemArgs)
			r, err := LookupWatchlistItem(ctx, &args, opts...)
			var s LookupWatchlistItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWatchlistItemResultOutput)
}

type LookupWatchlistItemOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The watchlist alias
	WatchlistAlias pulumi.StringInput `pulumi:"watchlistAlias"`
	// The watchlist item id (GUID)
	WatchlistItemId pulumi.StringInput `pulumi:"watchlistItemId"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWatchlistItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWatchlistItemArgs)(nil)).Elem()
}

// Represents a Watchlist Item in Azure Security Insights.
type LookupWatchlistItemResultOutput struct{ *pulumi.OutputState }

func (LookupWatchlistItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWatchlistItemResult)(nil)).Elem()
}

func (o LookupWatchlistItemResultOutput) ToLookupWatchlistItemResultOutput() LookupWatchlistItemResultOutput {
	return o
}

func (o LookupWatchlistItemResultOutput) ToLookupWatchlistItemResultOutputWithContext(ctx context.Context) LookupWatchlistItemResultOutput {
	return o
}

// The time the watchlist item was created
func (o LookupWatchlistItemResultOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *string { return v.Created }).(pulumi.StringPtrOutput)
}

// Describes a user that created the watchlist item
func (o LookupWatchlistItemResultOutput) CreatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *WatchlistUserInfoResponse { return v.CreatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

// key-value pairs for a watchlist item entity mapping
func (o LookupWatchlistItemResultOutput) EntityMapping() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) interface{} { return v.EntityMapping }).(pulumi.AnyOutput)
}

// Etag of the azure resource
func (o LookupWatchlistItemResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWatchlistItemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) string { return v.Id }).(pulumi.StringOutput)
}

// A flag that indicates if the watchlist item is deleted or not
func (o LookupWatchlistItemResultOutput) IsDeleted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *bool { return v.IsDeleted }).(pulumi.BoolPtrOutput)
}

// key-value pairs for a watchlist item
func (o LookupWatchlistItemResultOutput) ItemsKeyValue() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) interface{} { return v.ItemsKeyValue }).(pulumi.AnyOutput)
}

// The name of the resource
func (o LookupWatchlistItemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWatchlistItemResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenantId to which the watchlist item belongs to
func (o LookupWatchlistItemResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWatchlistItemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) string { return v.Type }).(pulumi.StringOutput)
}

// The last time the watchlist item was updated
func (o LookupWatchlistItemResultOutput) Updated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *string { return v.Updated }).(pulumi.StringPtrOutput)
}

// Describes a user that updated the watchlist item
func (o LookupWatchlistItemResultOutput) UpdatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *WatchlistUserInfoResponse { return v.UpdatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

// The id (a Guid) of the watchlist item
func (o LookupWatchlistItemResultOutput) WatchlistItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *string { return v.WatchlistItemId }).(pulumi.StringPtrOutput)
}

// The type of the watchlist item
func (o LookupWatchlistItemResultOutput) WatchlistItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatchlistItemResult) *string { return v.WatchlistItemType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWatchlistItemResultOutput{})
}
