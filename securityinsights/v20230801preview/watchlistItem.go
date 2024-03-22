// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Watchlist item in Azure Security Insights.
type WatchlistItem struct {
	pulumi.CustomResourceState

	// The time the watchlist item was created
	Created pulumi.StringPtrOutput `pulumi:"created"`
	// Describes a user that created the watchlist item
	CreatedBy WatchlistUserInfoResponsePtrOutput `pulumi:"createdBy"`
	// key-value pairs for a watchlist item entity mapping
	EntityMapping pulumi.AnyOutput `pulumi:"entityMapping"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// A flag that indicates if the watchlist item is deleted or not
	IsDeleted pulumi.BoolPtrOutput `pulumi:"isDeleted"`
	// key-value pairs for a watchlist item
	ItemsKeyValue pulumi.AnyOutput `pulumi:"itemsKeyValue"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tenantId to which the watchlist item belongs to
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The last time the watchlist item was updated
	Updated pulumi.StringPtrOutput `pulumi:"updated"`
	// Describes a user that updated the watchlist item
	UpdatedBy WatchlistUserInfoResponsePtrOutput `pulumi:"updatedBy"`
	// The id (a Guid) of the watchlist item
	WatchlistItemId pulumi.StringPtrOutput `pulumi:"watchlistItemId"`
	// The type of the watchlist item
	WatchlistItemType pulumi.StringPtrOutput `pulumi:"watchlistItemType"`
}

// NewWatchlistItem registers a new resource with the given unique name, arguments, and options.
func NewWatchlistItem(ctx *pulumi.Context,
	name string, args *WatchlistItemArgs, opts ...pulumi.ResourceOption) (*WatchlistItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ItemsKeyValue == nil {
		return nil, errors.New("invalid value for required argument 'ItemsKeyValue'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WatchlistAlias == nil {
		return nil, errors.New("invalid value for required argument 'WatchlistAlias'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210401:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240301:WatchlistItem"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WatchlistItem
	err := ctx.RegisterResource("azure-native:securityinsights/v20230801preview:WatchlistItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWatchlistItem gets an existing WatchlistItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWatchlistItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WatchlistItemState, opts ...pulumi.ResourceOption) (*WatchlistItem, error) {
	var resource WatchlistItem
	err := ctx.ReadResource("azure-native:securityinsights/v20230801preview:WatchlistItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WatchlistItem resources.
type watchlistItemState struct {
}

type WatchlistItemState struct {
}

func (WatchlistItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*watchlistItemState)(nil)).Elem()
}

type watchlistItemArgs struct {
	// The time the watchlist item was created
	Created *string `pulumi:"created"`
	// Describes a user that created the watchlist item
	CreatedBy *WatchlistUserInfo `pulumi:"createdBy"`
	// key-value pairs for a watchlist item entity mapping
	EntityMapping interface{} `pulumi:"entityMapping"`
	// A flag that indicates if the watchlist item is deleted or not
	IsDeleted *bool `pulumi:"isDeleted"`
	// key-value pairs for a watchlist item
	ItemsKeyValue interface{} `pulumi:"itemsKeyValue"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenantId to which the watchlist item belongs to
	TenantId *string `pulumi:"tenantId"`
	// The last time the watchlist item was updated
	Updated *string `pulumi:"updated"`
	// Describes a user that updated the watchlist item
	UpdatedBy *WatchlistUserInfo `pulumi:"updatedBy"`
	// Watchlist Alias
	WatchlistAlias string `pulumi:"watchlistAlias"`
	// The id (a Guid) of the watchlist item
	WatchlistItemId *string `pulumi:"watchlistItemId"`
	// The type of the watchlist item
	WatchlistItemType *string `pulumi:"watchlistItemType"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a WatchlistItem resource.
type WatchlistItemArgs struct {
	// The time the watchlist item was created
	Created pulumi.StringPtrInput
	// Describes a user that created the watchlist item
	CreatedBy WatchlistUserInfoPtrInput
	// key-value pairs for a watchlist item entity mapping
	EntityMapping pulumi.Input
	// A flag that indicates if the watchlist item is deleted or not
	IsDeleted pulumi.BoolPtrInput
	// key-value pairs for a watchlist item
	ItemsKeyValue pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenantId to which the watchlist item belongs to
	TenantId pulumi.StringPtrInput
	// The last time the watchlist item was updated
	Updated pulumi.StringPtrInput
	// Describes a user that updated the watchlist item
	UpdatedBy WatchlistUserInfoPtrInput
	// Watchlist Alias
	WatchlistAlias pulumi.StringInput
	// The id (a Guid) of the watchlist item
	WatchlistItemId pulumi.StringPtrInput
	// The type of the watchlist item
	WatchlistItemType pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (WatchlistItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*watchlistItemArgs)(nil)).Elem()
}

type WatchlistItemInput interface {
	pulumi.Input

	ToWatchlistItemOutput() WatchlistItemOutput
	ToWatchlistItemOutputWithContext(ctx context.Context) WatchlistItemOutput
}

func (*WatchlistItem) ElementType() reflect.Type {
	return reflect.TypeOf((**WatchlistItem)(nil)).Elem()
}

func (i *WatchlistItem) ToWatchlistItemOutput() WatchlistItemOutput {
	return i.ToWatchlistItemOutputWithContext(context.Background())
}

func (i *WatchlistItem) ToWatchlistItemOutputWithContext(ctx context.Context) WatchlistItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistItemOutput)
}

type WatchlistItemOutput struct{ *pulumi.OutputState }

func (WatchlistItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WatchlistItem)(nil)).Elem()
}

func (o WatchlistItemOutput) ToWatchlistItemOutput() WatchlistItemOutput {
	return o
}

func (o WatchlistItemOutput) ToWatchlistItemOutputWithContext(ctx context.Context) WatchlistItemOutput {
	return o
}

// The time the watchlist item was created
func (o WatchlistItemOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringPtrOutput { return v.Created }).(pulumi.StringPtrOutput)
}

// Describes a user that created the watchlist item
func (o WatchlistItemOutput) CreatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v *WatchlistItem) WatchlistUserInfoResponsePtrOutput { return v.CreatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

// key-value pairs for a watchlist item entity mapping
func (o WatchlistItemOutput) EntityMapping() pulumi.AnyOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.AnyOutput { return v.EntityMapping }).(pulumi.AnyOutput)
}

// Etag of the azure resource
func (o WatchlistItemOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// A flag that indicates if the watchlist item is deleted or not
func (o WatchlistItemOutput) IsDeleted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.BoolPtrOutput { return v.IsDeleted }).(pulumi.BoolPtrOutput)
}

// key-value pairs for a watchlist item
func (o WatchlistItemOutput) ItemsKeyValue() pulumi.AnyOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.AnyOutput { return v.ItemsKeyValue }).(pulumi.AnyOutput)
}

// The name of the resource
func (o WatchlistItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WatchlistItemOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WatchlistItem) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenantId to which the watchlist item belongs to
func (o WatchlistItemOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WatchlistItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The last time the watchlist item was updated
func (o WatchlistItemOutput) Updated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringPtrOutput { return v.Updated }).(pulumi.StringPtrOutput)
}

// Describes a user that updated the watchlist item
func (o WatchlistItemOutput) UpdatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v *WatchlistItem) WatchlistUserInfoResponsePtrOutput { return v.UpdatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

// The id (a Guid) of the watchlist item
func (o WatchlistItemOutput) WatchlistItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringPtrOutput { return v.WatchlistItemId }).(pulumi.StringPtrOutput)
}

// The type of the watchlist item
func (o WatchlistItemOutput) WatchlistItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringPtrOutput { return v.WatchlistItemType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WatchlistItemOutput{})
}
