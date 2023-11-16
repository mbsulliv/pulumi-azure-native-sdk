// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Asset Filter.
type AssetFilter struct {
	pulumi.CustomResourceState

	// The first quality.
	FirstQuality FirstQualityResponsePtrOutput `pulumi:"firstQuality"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The presentation time range.
	PresentationTimeRange PresentationTimeRangeResponsePtrOutput `pulumi:"presentationTimeRange"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tracks selection conditions.
	Tracks FilterTrackSelectionResponseArrayOutput `pulumi:"tracks"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAssetFilter registers a new resource with the given unique name, arguments, and options.
func NewAssetFilter(ctx *pulumi.Context,
	name string, args *AssetFilterArgs, opts ...pulumi.ResourceOption) (*AssetFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.AssetName == nil {
		return nil, errors.New("invalid value for required argument 'AssetName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:AssetFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:AssetFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:AssetFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:AssetFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:AssetFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20220801:AssetFilter"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AssetFilter
	err := ctx.RegisterResource("azure-native:media/v20230101:AssetFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssetFilter gets an existing AssetFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssetFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssetFilterState, opts ...pulumi.ResourceOption) (*AssetFilter, error) {
	var resource AssetFilter
	err := ctx.ReadResource("azure-native:media/v20230101:AssetFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssetFilter resources.
type assetFilterState struct {
}

type AssetFilterState struct {
}

func (AssetFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*assetFilterState)(nil)).Elem()
}

type assetFilterArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Asset name.
	AssetName string `pulumi:"assetName"`
	// The Asset Filter name
	FilterName *string `pulumi:"filterName"`
	// The first quality.
	FirstQuality *FirstQuality `pulumi:"firstQuality"`
	// The presentation time range.
	PresentationTimeRange *PresentationTimeRange `pulumi:"presentationTimeRange"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tracks selection conditions.
	Tracks []FilterTrackSelection `pulumi:"tracks"`
}

// The set of arguments for constructing a AssetFilter resource.
type AssetFilterArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The Asset name.
	AssetName pulumi.StringInput
	// The Asset Filter name
	FilterName pulumi.StringPtrInput
	// The first quality.
	FirstQuality FirstQualityPtrInput
	// The presentation time range.
	PresentationTimeRange PresentationTimeRangePtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The tracks selection conditions.
	Tracks FilterTrackSelectionArrayInput
}

func (AssetFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assetFilterArgs)(nil)).Elem()
}

type AssetFilterInput interface {
	pulumi.Input

	ToAssetFilterOutput() AssetFilterOutput
	ToAssetFilterOutputWithContext(ctx context.Context) AssetFilterOutput
}

func (*AssetFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**AssetFilter)(nil)).Elem()
}

func (i *AssetFilter) ToAssetFilterOutput() AssetFilterOutput {
	return i.ToAssetFilterOutputWithContext(context.Background())
}

func (i *AssetFilter) ToAssetFilterOutputWithContext(ctx context.Context) AssetFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetFilterOutput)
}

type AssetFilterOutput struct{ *pulumi.OutputState }

func (AssetFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssetFilter)(nil)).Elem()
}

func (o AssetFilterOutput) ToAssetFilterOutput() AssetFilterOutput {
	return o
}

func (o AssetFilterOutput) ToAssetFilterOutputWithContext(ctx context.Context) AssetFilterOutput {
	return o
}

// The first quality.
func (o AssetFilterOutput) FirstQuality() FirstQualityResponsePtrOutput {
	return o.ApplyT(func(v *AssetFilter) FirstQualityResponsePtrOutput { return v.FirstQuality }).(FirstQualityResponsePtrOutput)
}

// The name of the resource
func (o AssetFilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AssetFilter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The presentation time range.
func (o AssetFilterOutput) PresentationTimeRange() PresentationTimeRangeResponsePtrOutput {
	return o.ApplyT(func(v *AssetFilter) PresentationTimeRangeResponsePtrOutput { return v.PresentationTimeRange }).(PresentationTimeRangeResponsePtrOutput)
}

// The system metadata relating to this resource.
func (o AssetFilterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AssetFilter) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tracks selection conditions.
func (o AssetFilterOutput) Tracks() FilterTrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v *AssetFilter) FilterTrackSelectionResponseArrayOutput { return v.Tracks }).(FilterTrackSelectionResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AssetFilterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AssetFilter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssetFilterOutput{})
}
