// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a Watchlist in Azure Security Insights.
type Watchlist struct {
	pulumi.CustomResourceState

	// The content type of the raw content. Example : text/csv or text/tsv
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// The time the watchlist was created
	Created pulumi.StringPtrOutput `pulumi:"created"`
	// Describes a user that created the watchlist
	CreatedBy WatchlistUserInfoResponsePtrOutput `pulumi:"createdBy"`
	// The default duration of a watchlist (in ISO 8601 duration format)
	DefaultDuration pulumi.StringPtrOutput `pulumi:"defaultDuration"`
	// A description of the watchlist
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the watchlist
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// A flag that indicates if the watchlist is deleted or not
	IsDeleted pulumi.BoolPtrOutput `pulumi:"isDeleted"`
	// The search key is used to optimize query performance when using watchlists for joins with other data. For example, enable a column with IP addresses to be the designated SearchKey field, then use this field as the key field when joining to other event data by IP address.
	ItemsSearchKey pulumi.StringOutput `pulumi:"itemsSearchKey"`
	// List of labels relevant to this watchlist
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of lines in a csv/tsv content to skip before the header
	NumberOfLinesToSkip pulumi.IntPtrOutput `pulumi:"numberOfLinesToSkip"`
	// The provider of the watchlist
	Provider pulumi.StringOutput `pulumi:"provider"`
	// The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
	RawContent pulumi.StringPtrOutput `pulumi:"rawContent"`
	// The filename of the watchlist, called 'source'
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// The sourceType of the watchlist
	SourceType pulumi.StringPtrOutput `pulumi:"sourceType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tenantId where the watchlist belongs to
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The last time the watchlist was updated
	Updated pulumi.StringPtrOutput `pulumi:"updated"`
	// Describes a user that updated the watchlist
	UpdatedBy WatchlistUserInfoResponsePtrOutput `pulumi:"updatedBy"`
	// The status of the Watchlist upload : New, InProgress or Complete. Pls note : When a Watchlist upload status is equal to InProgress, the Watchlist cannot be deleted
	UploadStatus pulumi.StringPtrOutput `pulumi:"uploadStatus"`
	// The alias of the watchlist
	WatchlistAlias pulumi.StringPtrOutput `pulumi:"watchlistAlias"`
	// The id (a Guid) of the watchlist
	WatchlistId pulumi.StringPtrOutput `pulumi:"watchlistId"`
	// The type of the watchlist
	WatchlistType pulumi.StringPtrOutput `pulumi:"watchlistType"`
}

// NewWatchlist registers a new resource with the given unique name, arguments, and options.
func NewWatchlist(ctx *pulumi.Context,
	name string, args *WatchlistArgs, opts ...pulumi.ResourceOption) (*Watchlist, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ItemsSearchKey == nil {
		return nil, errors.New("invalid value for required argument 'ItemsSearchKey'")
	}
	if args.Provider == nil {
		return nil, errors.New("invalid value for required argument 'Provider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210401:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:Watchlist"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Watchlist
	err := ctx.RegisterResource("azure-native:securityinsights/v20230701preview:Watchlist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWatchlist gets an existing Watchlist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWatchlist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WatchlistState, opts ...pulumi.ResourceOption) (*Watchlist, error) {
	var resource Watchlist
	err := ctx.ReadResource("azure-native:securityinsights/v20230701preview:Watchlist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Watchlist resources.
type watchlistState struct {
}

type WatchlistState struct {
}

func (WatchlistState) ElementType() reflect.Type {
	return reflect.TypeOf((*watchlistState)(nil)).Elem()
}

type watchlistArgs struct {
	// The content type of the raw content. Example : text/csv or text/tsv
	ContentType *string `pulumi:"contentType"`
	// The time the watchlist was created
	Created *string `pulumi:"created"`
	// Describes a user that created the watchlist
	CreatedBy *WatchlistUserInfo `pulumi:"createdBy"`
	// The default duration of a watchlist (in ISO 8601 duration format)
	DefaultDuration *string `pulumi:"defaultDuration"`
	// A description of the watchlist
	Description *string `pulumi:"description"`
	// The display name of the watchlist
	DisplayName string `pulumi:"displayName"`
	// A flag that indicates if the watchlist is deleted or not
	IsDeleted *bool `pulumi:"isDeleted"`
	// The search key is used to optimize query performance when using watchlists for joins with other data. For example, enable a column with IP addresses to be the designated SearchKey field, then use this field as the key field when joining to other event data by IP address.
	ItemsSearchKey string `pulumi:"itemsSearchKey"`
	// List of labels relevant to this watchlist
	Labels []string `pulumi:"labels"`
	// The number of lines in a csv/tsv content to skip before the header
	NumberOfLinesToSkip *int `pulumi:"numberOfLinesToSkip"`
	// The provider of the watchlist
	Provider string `pulumi:"provider"`
	// The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
	RawContent *string `pulumi:"rawContent"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The filename of the watchlist, called 'source'
	Source *string `pulumi:"source"`
	// The sourceType of the watchlist
	SourceType *string `pulumi:"sourceType"`
	// The tenantId where the watchlist belongs to
	TenantId *string `pulumi:"tenantId"`
	// The last time the watchlist was updated
	Updated *string `pulumi:"updated"`
	// Describes a user that updated the watchlist
	UpdatedBy *WatchlistUserInfo `pulumi:"updatedBy"`
	// The status of the Watchlist upload : New, InProgress or Complete. Pls note : When a Watchlist upload status is equal to InProgress, the Watchlist cannot be deleted
	UploadStatus *string `pulumi:"uploadStatus"`
	// The alias of the watchlist
	WatchlistAlias *string `pulumi:"watchlistAlias"`
	// The id (a Guid) of the watchlist
	WatchlistId *string `pulumi:"watchlistId"`
	// The type of the watchlist
	WatchlistType *string `pulumi:"watchlistType"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Watchlist resource.
type WatchlistArgs struct {
	// The content type of the raw content. Example : text/csv or text/tsv
	ContentType pulumi.StringPtrInput
	// The time the watchlist was created
	Created pulumi.StringPtrInput
	// Describes a user that created the watchlist
	CreatedBy WatchlistUserInfoPtrInput
	// The default duration of a watchlist (in ISO 8601 duration format)
	DefaultDuration pulumi.StringPtrInput
	// A description of the watchlist
	Description pulumi.StringPtrInput
	// The display name of the watchlist
	DisplayName pulumi.StringInput
	// A flag that indicates if the watchlist is deleted or not
	IsDeleted pulumi.BoolPtrInput
	// The search key is used to optimize query performance when using watchlists for joins with other data. For example, enable a column with IP addresses to be the designated SearchKey field, then use this field as the key field when joining to other event data by IP address.
	ItemsSearchKey pulumi.StringInput
	// List of labels relevant to this watchlist
	Labels pulumi.StringArrayInput
	// The number of lines in a csv/tsv content to skip before the header
	NumberOfLinesToSkip pulumi.IntPtrInput
	// The provider of the watchlist
	Provider pulumi.StringInput
	// The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
	RawContent pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The filename of the watchlist, called 'source'
	Source pulumi.StringPtrInput
	// The sourceType of the watchlist
	SourceType pulumi.StringPtrInput
	// The tenantId where the watchlist belongs to
	TenantId pulumi.StringPtrInput
	// The last time the watchlist was updated
	Updated pulumi.StringPtrInput
	// Describes a user that updated the watchlist
	UpdatedBy WatchlistUserInfoPtrInput
	// The status of the Watchlist upload : New, InProgress or Complete. Pls note : When a Watchlist upload status is equal to InProgress, the Watchlist cannot be deleted
	UploadStatus pulumi.StringPtrInput
	// The alias of the watchlist
	WatchlistAlias pulumi.StringPtrInput
	// The id (a Guid) of the watchlist
	WatchlistId pulumi.StringPtrInput
	// The type of the watchlist
	WatchlistType pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (WatchlistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*watchlistArgs)(nil)).Elem()
}

type WatchlistInput interface {
	pulumi.Input

	ToWatchlistOutput() WatchlistOutput
	ToWatchlistOutputWithContext(ctx context.Context) WatchlistOutput
}

func (*Watchlist) ElementType() reflect.Type {
	return reflect.TypeOf((**Watchlist)(nil)).Elem()
}

func (i *Watchlist) ToWatchlistOutput() WatchlistOutput {
	return i.ToWatchlistOutputWithContext(context.Background())
}

func (i *Watchlist) ToWatchlistOutputWithContext(ctx context.Context) WatchlistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistOutput)
}

func (i *Watchlist) ToOutput(ctx context.Context) pulumix.Output[*Watchlist] {
	return pulumix.Output[*Watchlist]{
		OutputState: i.ToWatchlistOutputWithContext(ctx).OutputState,
	}
}

type WatchlistOutput struct{ *pulumi.OutputState }

func (WatchlistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Watchlist)(nil)).Elem()
}

func (o WatchlistOutput) ToWatchlistOutput() WatchlistOutput {
	return o
}

func (o WatchlistOutput) ToWatchlistOutputWithContext(ctx context.Context) WatchlistOutput {
	return o
}

func (o WatchlistOutput) ToOutput(ctx context.Context) pulumix.Output[*Watchlist] {
	return pulumix.Output[*Watchlist]{
		OutputState: o.OutputState,
	}
}

// The content type of the raw content. Example : text/csv or text/tsv
func (o WatchlistOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

// The time the watchlist was created
func (o WatchlistOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.Created }).(pulumi.StringPtrOutput)
}

// Describes a user that created the watchlist
func (o WatchlistOutput) CreatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v *Watchlist) WatchlistUserInfoResponsePtrOutput { return v.CreatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

// The default duration of a watchlist (in ISO 8601 duration format)
func (o WatchlistOutput) DefaultDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.DefaultDuration }).(pulumi.StringPtrOutput)
}

// A description of the watchlist
func (o WatchlistOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the watchlist
func (o WatchlistOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o WatchlistOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// A flag that indicates if the watchlist is deleted or not
func (o WatchlistOutput) IsDeleted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.BoolPtrOutput { return v.IsDeleted }).(pulumi.BoolPtrOutput)
}

// The search key is used to optimize query performance when using watchlists for joins with other data. For example, enable a column with IP addresses to be the designated SearchKey field, then use this field as the key field when joining to other event data by IP address.
func (o WatchlistOutput) ItemsSearchKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringOutput { return v.ItemsSearchKey }).(pulumi.StringOutput)
}

// List of labels relevant to this watchlist
func (o WatchlistOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o WatchlistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of lines in a csv/tsv content to skip before the header
func (o WatchlistOutput) NumberOfLinesToSkip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.IntPtrOutput { return v.NumberOfLinesToSkip }).(pulumi.IntPtrOutput)
}

// The provider of the watchlist
func (o WatchlistOutput) Provider() pulumi.StringOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringOutput { return v.Provider }).(pulumi.StringOutput)
}

// The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
func (o WatchlistOutput) RawContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.RawContent }).(pulumi.StringPtrOutput)
}

// The filename of the watchlist, called 'source'
func (o WatchlistOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

// The sourceType of the watchlist
func (o WatchlistOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.SourceType }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WatchlistOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Watchlist) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenantId where the watchlist belongs to
func (o WatchlistOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WatchlistOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The last time the watchlist was updated
func (o WatchlistOutput) Updated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.Updated }).(pulumi.StringPtrOutput)
}

// Describes a user that updated the watchlist
func (o WatchlistOutput) UpdatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v *Watchlist) WatchlistUserInfoResponsePtrOutput { return v.UpdatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

// The status of the Watchlist upload : New, InProgress or Complete. Pls note : When a Watchlist upload status is equal to InProgress, the Watchlist cannot be deleted
func (o WatchlistOutput) UploadStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.UploadStatus }).(pulumi.StringPtrOutput)
}

// The alias of the watchlist
func (o WatchlistOutput) WatchlistAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.WatchlistAlias }).(pulumi.StringPtrOutput)
}

// The id (a Guid) of the watchlist
func (o WatchlistOutput) WatchlistId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.WatchlistId }).(pulumi.StringPtrOutput)
}

// The type of the watchlist
func (o WatchlistOutput) WatchlistType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.WatchlistType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WatchlistOutput{})
}
