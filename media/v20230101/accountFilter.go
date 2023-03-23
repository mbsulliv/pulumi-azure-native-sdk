// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Account Filter.
type AccountFilter struct {
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

// NewAccountFilter registers a new resource with the given unique name, arguments, and options.
func NewAccountFilter(ctx *pulumi.Context,
	name string, args *AccountFilterArgs, opts ...pulumi.ResourceOption) (*AccountFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:AccountFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:AccountFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:AccountFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:AccountFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:AccountFilter"),
		},
		{
			Type: pulumi.String("azure-native:media/v20220801:AccountFilter"),
		},
	})
	opts = append(opts, aliases)
	var resource AccountFilter
	err := ctx.RegisterResource("azure-native:media/v20230101:AccountFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountFilter gets an existing AccountFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountFilterState, opts ...pulumi.ResourceOption) (*AccountFilter, error) {
	var resource AccountFilter
	err := ctx.ReadResource("azure-native:media/v20230101:AccountFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountFilter resources.
type accountFilterState struct {
}

type AccountFilterState struct {
}

func (AccountFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountFilterState)(nil)).Elem()
}

type accountFilterArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Account Filter name
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

// The set of arguments for constructing a AccountFilter resource.
type AccountFilterArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The Account Filter name
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

func (AccountFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountFilterArgs)(nil)).Elem()
}

type AccountFilterInput interface {
	pulumi.Input

	ToAccountFilterOutput() AccountFilterOutput
	ToAccountFilterOutputWithContext(ctx context.Context) AccountFilterOutput
}

func (*AccountFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountFilter)(nil)).Elem()
}

func (i *AccountFilter) ToAccountFilterOutput() AccountFilterOutput {
	return i.ToAccountFilterOutputWithContext(context.Background())
}

func (i *AccountFilter) ToAccountFilterOutputWithContext(ctx context.Context) AccountFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountFilterOutput)
}

type AccountFilterOutput struct{ *pulumi.OutputState }

func (AccountFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountFilter)(nil)).Elem()
}

func (o AccountFilterOutput) ToAccountFilterOutput() AccountFilterOutput {
	return o
}

func (o AccountFilterOutput) ToAccountFilterOutputWithContext(ctx context.Context) AccountFilterOutput {
	return o
}

// The first quality.
func (o AccountFilterOutput) FirstQuality() FirstQualityResponsePtrOutput {
	return o.ApplyT(func(v *AccountFilter) FirstQualityResponsePtrOutput { return v.FirstQuality }).(FirstQualityResponsePtrOutput)
}

// The name of the resource
func (o AccountFilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountFilter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The presentation time range.
func (o AccountFilterOutput) PresentationTimeRange() PresentationTimeRangeResponsePtrOutput {
	return o.ApplyT(func(v *AccountFilter) PresentationTimeRangeResponsePtrOutput { return v.PresentationTimeRange }).(PresentationTimeRangeResponsePtrOutput)
}

// The system metadata relating to this resource.
func (o AccountFilterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AccountFilter) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tracks selection conditions.
func (o AccountFilterOutput) Tracks() FilterTrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v *AccountFilter) FilterTrackSelectionResponseArrayOutput { return v.Tracks }).(FilterTrackSelectionResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AccountFilterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountFilter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountFilterOutput{})
}
