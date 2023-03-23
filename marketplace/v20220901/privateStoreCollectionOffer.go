// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The privateStore offer data structure.
type PrivateStoreCollectionOffer struct {
	pulumi.CustomResourceState

	// Private store offer creation date
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Identifier for purposes of race condition
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Icon File Uris
	IconFileUris pulumi.StringMapOutput `pulumi:"iconFileUris"`
	// Private store offer modification date
	ModifiedAt pulumi.StringOutput `pulumi:"modifiedAt"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// It will be displayed prominently in the marketplace
	OfferDisplayName pulumi.StringOutput `pulumi:"offerDisplayName"`
	// Offer plans
	Plans PlanResponseArrayOutput `pulumi:"plans"`
	// Private store unique id
	PrivateStoreId pulumi.StringOutput `pulumi:"privateStoreId"`
	// Publisher name that will be displayed prominently in the marketplace
	PublisherDisplayName pulumi.StringOutput `pulumi:"publisherDisplayName"`
	// Plan ids limitation for this offer
	SpecificPlanIdsLimitation pulumi.StringArrayOutput `pulumi:"specificPlanIdsLimitation"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Offers unique id
	UniqueOfferId pulumi.StringOutput `pulumi:"uniqueOfferId"`
	// Indicating whether the offer was not updated to db (true = not updated). If the allow list is identical to the existed one in db, the offer would not be updated.
	UpdateSuppressedDueIdempotence pulumi.BoolPtrOutput `pulumi:"updateSuppressedDueIdempotence"`
}

// NewPrivateStoreCollectionOffer registers a new resource with the given unique name, arguments, and options.
func NewPrivateStoreCollectionOffer(ctx *pulumi.Context,
	name string, args *PrivateStoreCollectionOfferArgs, opts ...pulumi.ResourceOption) (*PrivateStoreCollectionOffer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CollectionId == nil {
		return nil, errors.New("invalid value for required argument 'CollectionId'")
	}
	if args.PrivateStoreId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateStoreId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:marketplace:PrivateStoreCollectionOffer"),
		},
		{
			Type: pulumi.String("azure-native:marketplace/v20210601:PrivateStoreCollectionOffer"),
		},
		{
			Type: pulumi.String("azure-native:marketplace/v20211201:PrivateStoreCollectionOffer"),
		},
		{
			Type: pulumi.String("azure-native:marketplace/v20220301:PrivateStoreCollectionOffer"),
		},
		{
			Type: pulumi.String("azure-native:marketplace/v20230101:PrivateStoreCollectionOffer"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateStoreCollectionOffer
	err := ctx.RegisterResource("azure-native:marketplace/v20220901:PrivateStoreCollectionOffer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateStoreCollectionOffer gets an existing PrivateStoreCollectionOffer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateStoreCollectionOffer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateStoreCollectionOfferState, opts ...pulumi.ResourceOption) (*PrivateStoreCollectionOffer, error) {
	var resource PrivateStoreCollectionOffer
	err := ctx.ReadResource("azure-native:marketplace/v20220901:PrivateStoreCollectionOffer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateStoreCollectionOffer resources.
type privateStoreCollectionOfferState struct {
}

type PrivateStoreCollectionOfferState struct {
}

func (PrivateStoreCollectionOfferState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateStoreCollectionOfferState)(nil)).Elem()
}

type privateStoreCollectionOfferArgs struct {
	// The collection ID
	CollectionId string `pulumi:"collectionId"`
	// Identifier for purposes of race condition
	ETag *string `pulumi:"eTag"`
	// Icon File Uris
	IconFileUris map[string]string `pulumi:"iconFileUris"`
	// The offer ID to update or delete
	OfferId *string `pulumi:"offerId"`
	// Offer plans
	Plans []Plan `pulumi:"plans"`
	// The store ID - must use the tenant ID
	PrivateStoreId string `pulumi:"privateStoreId"`
	// Plan ids limitation for this offer
	SpecificPlanIdsLimitation []string `pulumi:"specificPlanIdsLimitation"`
	// Indicating whether the offer was not updated to db (true = not updated). If the allow list is identical to the existed one in db, the offer would not be updated.
	UpdateSuppressedDueIdempotence *bool `pulumi:"updateSuppressedDueIdempotence"`
}

// The set of arguments for constructing a PrivateStoreCollectionOffer resource.
type PrivateStoreCollectionOfferArgs struct {
	// The collection ID
	CollectionId pulumi.StringInput
	// Identifier for purposes of race condition
	ETag pulumi.StringPtrInput
	// Icon File Uris
	IconFileUris pulumi.StringMapInput
	// The offer ID to update or delete
	OfferId pulumi.StringPtrInput
	// Offer plans
	Plans PlanArrayInput
	// The store ID - must use the tenant ID
	PrivateStoreId pulumi.StringInput
	// Plan ids limitation for this offer
	SpecificPlanIdsLimitation pulumi.StringArrayInput
	// Indicating whether the offer was not updated to db (true = not updated). If the allow list is identical to the existed one in db, the offer would not be updated.
	UpdateSuppressedDueIdempotence pulumi.BoolPtrInput
}

func (PrivateStoreCollectionOfferArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateStoreCollectionOfferArgs)(nil)).Elem()
}

type PrivateStoreCollectionOfferInput interface {
	pulumi.Input

	ToPrivateStoreCollectionOfferOutput() PrivateStoreCollectionOfferOutput
	ToPrivateStoreCollectionOfferOutputWithContext(ctx context.Context) PrivateStoreCollectionOfferOutput
}

func (*PrivateStoreCollectionOffer) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateStoreCollectionOffer)(nil)).Elem()
}

func (i *PrivateStoreCollectionOffer) ToPrivateStoreCollectionOfferOutput() PrivateStoreCollectionOfferOutput {
	return i.ToPrivateStoreCollectionOfferOutputWithContext(context.Background())
}

func (i *PrivateStoreCollectionOffer) ToPrivateStoreCollectionOfferOutputWithContext(ctx context.Context) PrivateStoreCollectionOfferOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateStoreCollectionOfferOutput)
}

type PrivateStoreCollectionOfferOutput struct{ *pulumi.OutputState }

func (PrivateStoreCollectionOfferOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateStoreCollectionOffer)(nil)).Elem()
}

func (o PrivateStoreCollectionOfferOutput) ToPrivateStoreCollectionOfferOutput() PrivateStoreCollectionOfferOutput {
	return o
}

func (o PrivateStoreCollectionOfferOutput) ToPrivateStoreCollectionOfferOutputWithContext(ctx context.Context) PrivateStoreCollectionOfferOutput {
	return o
}

// Private store offer creation date
func (o PrivateStoreCollectionOfferOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Identifier for purposes of race condition
func (o PrivateStoreCollectionOfferOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// Icon File Uris
func (o PrivateStoreCollectionOfferOutput) IconFileUris() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) pulumi.StringMapOutput { return v.IconFileUris }).(pulumi.StringMapOutput)
}

// Private store offer modification date
func (o PrivateStoreCollectionOfferOutput) ModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) pulumi.StringOutput { return v.ModifiedAt }).(pulumi.StringOutput)
}

// The name of the resource.
func (o PrivateStoreCollectionOfferOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// It will be displayed prominently in the marketplace
func (o PrivateStoreCollectionOfferOutput) OfferDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) pulumi.StringOutput { return v.OfferDisplayName }).(pulumi.StringOutput)
}

// Offer plans
func (o PrivateStoreCollectionOfferOutput) Plans() PlanResponseArrayOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) PlanResponseArrayOutput { return v.Plans }).(PlanResponseArrayOutput)
}

// Private store unique id
func (o PrivateStoreCollectionOfferOutput) PrivateStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) pulumi.StringOutput { return v.PrivateStoreId }).(pulumi.StringOutput)
}

// Publisher name that will be displayed prominently in the marketplace
func (o PrivateStoreCollectionOfferOutput) PublisherDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) pulumi.StringOutput { return v.PublisherDisplayName }).(pulumi.StringOutput)
}

// Plan ids limitation for this offer
func (o PrivateStoreCollectionOfferOutput) SpecificPlanIdsLimitation() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) pulumi.StringArrayOutput { return v.SpecificPlanIdsLimitation }).(pulumi.StringArrayOutput)
}

// Metadata pertaining to creation and last modification of the resource
func (o PrivateStoreCollectionOfferOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o PrivateStoreCollectionOfferOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Offers unique id
func (o PrivateStoreCollectionOfferOutput) UniqueOfferId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) pulumi.StringOutput { return v.UniqueOfferId }).(pulumi.StringOutput)
}

// Indicating whether the offer was not updated to db (true = not updated). If the allow list is identical to the existed one in db, the offer would not be updated.
func (o PrivateStoreCollectionOfferOutput) UpdateSuppressedDueIdempotence() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PrivateStoreCollectionOffer) pulumi.BoolPtrOutput { return v.UpdateSuppressedDueIdempotence }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateStoreCollectionOfferOutput{})
}
