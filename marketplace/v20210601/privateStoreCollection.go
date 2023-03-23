// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Collection data structure.
type PrivateStoreCollection struct {
	pulumi.CustomResourceState

	// Indicating whether all subscriptions are selected (=true) or not (=false).
	AllSubscriptions pulumi.BoolPtrOutput `pulumi:"allSubscriptions"`
	// Gets or sets the association with Commercial's Billing Account.
	Claim pulumi.StringPtrOutput `pulumi:"claim"`
	// Gets collection Id.
	CollectionId pulumi.StringOutput `pulumi:"collectionId"`
	// Gets or sets collection name.
	CollectionName pulumi.StringPtrOutput `pulumi:"collectionName"`
	// Indicating whether the collection is enabled or disabled.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the number of offers associated with the collection.
	NumberOfOffers pulumi.Float64Output `pulumi:"numberOfOffers"`
	// Gets or sets subscription ids list. Empty list indicates all subscriptions are selected, null indicates no update is done, explicit list indicates the explicit selected subscriptions. On insert, null is considered as bad request
	SubscriptionsList pulumi.StringArrayOutput `pulumi:"subscriptionsList"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateStoreCollection registers a new resource with the given unique name, arguments, and options.
func NewPrivateStoreCollection(ctx *pulumi.Context,
	name string, args *PrivateStoreCollectionArgs, opts ...pulumi.ResourceOption) (*PrivateStoreCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateStoreId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateStoreId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:marketplace:PrivateStoreCollection"),
		},
		{
			Type: pulumi.String("azure-native:marketplace/v20211201:PrivateStoreCollection"),
		},
		{
			Type: pulumi.String("azure-native:marketplace/v20220301:PrivateStoreCollection"),
		},
		{
			Type: pulumi.String("azure-native:marketplace/v20220901:PrivateStoreCollection"),
		},
		{
			Type: pulumi.String("azure-native:marketplace/v20230101:PrivateStoreCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateStoreCollection
	err := ctx.RegisterResource("azure-native:marketplace/v20210601:PrivateStoreCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateStoreCollection gets an existing PrivateStoreCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateStoreCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateStoreCollectionState, opts ...pulumi.ResourceOption) (*PrivateStoreCollection, error) {
	var resource PrivateStoreCollection
	err := ctx.ReadResource("azure-native:marketplace/v20210601:PrivateStoreCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateStoreCollection resources.
type privateStoreCollectionState struct {
}

type PrivateStoreCollectionState struct {
}

func (PrivateStoreCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateStoreCollectionState)(nil)).Elem()
}

type privateStoreCollectionArgs struct {
	// Indicating whether all subscriptions are selected (=true) or not (=false).
	AllSubscriptions *bool `pulumi:"allSubscriptions"`
	// Gets or sets the association with Commercial's Billing Account.
	Claim *string `pulumi:"claim"`
	// The collection ID
	CollectionId *string `pulumi:"collectionId"`
	// Gets or sets collection name.
	CollectionName *string `pulumi:"collectionName"`
	// Indicating whether the collection is enabled or disabled.
	Enabled *bool `pulumi:"enabled"`
	// The store ID - must use the tenant ID
	PrivateStoreId string `pulumi:"privateStoreId"`
	// Gets or sets subscription ids list. Empty list indicates all subscriptions are selected, null indicates no update is done, explicit list indicates the explicit selected subscriptions. On insert, null is considered as bad request
	SubscriptionsList []string `pulumi:"subscriptionsList"`
}

// The set of arguments for constructing a PrivateStoreCollection resource.
type PrivateStoreCollectionArgs struct {
	// Indicating whether all subscriptions are selected (=true) or not (=false).
	AllSubscriptions pulumi.BoolPtrInput
	// Gets or sets the association with Commercial's Billing Account.
	Claim pulumi.StringPtrInput
	// The collection ID
	CollectionId pulumi.StringPtrInput
	// Gets or sets collection name.
	CollectionName pulumi.StringPtrInput
	// Indicating whether the collection is enabled or disabled.
	Enabled pulumi.BoolPtrInput
	// The store ID - must use the tenant ID
	PrivateStoreId pulumi.StringInput
	// Gets or sets subscription ids list. Empty list indicates all subscriptions are selected, null indicates no update is done, explicit list indicates the explicit selected subscriptions. On insert, null is considered as bad request
	SubscriptionsList pulumi.StringArrayInput
}

func (PrivateStoreCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateStoreCollectionArgs)(nil)).Elem()
}

type PrivateStoreCollectionInput interface {
	pulumi.Input

	ToPrivateStoreCollectionOutput() PrivateStoreCollectionOutput
	ToPrivateStoreCollectionOutputWithContext(ctx context.Context) PrivateStoreCollectionOutput
}

func (*PrivateStoreCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateStoreCollection)(nil)).Elem()
}

func (i *PrivateStoreCollection) ToPrivateStoreCollectionOutput() PrivateStoreCollectionOutput {
	return i.ToPrivateStoreCollectionOutputWithContext(context.Background())
}

func (i *PrivateStoreCollection) ToPrivateStoreCollectionOutputWithContext(ctx context.Context) PrivateStoreCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateStoreCollectionOutput)
}

type PrivateStoreCollectionOutput struct{ *pulumi.OutputState }

func (PrivateStoreCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateStoreCollection)(nil)).Elem()
}

func (o PrivateStoreCollectionOutput) ToPrivateStoreCollectionOutput() PrivateStoreCollectionOutput {
	return o
}

func (o PrivateStoreCollectionOutput) ToPrivateStoreCollectionOutputWithContext(ctx context.Context) PrivateStoreCollectionOutput {
	return o
}

// Indicating whether all subscriptions are selected (=true) or not (=false).
func (o PrivateStoreCollectionOutput) AllSubscriptions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PrivateStoreCollection) pulumi.BoolPtrOutput { return v.AllSubscriptions }).(pulumi.BoolPtrOutput)
}

// Gets or sets the association with Commercial's Billing Account.
func (o PrivateStoreCollectionOutput) Claim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateStoreCollection) pulumi.StringPtrOutput { return v.Claim }).(pulumi.StringPtrOutput)
}

// Gets collection Id.
func (o PrivateStoreCollectionOutput) CollectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateStoreCollection) pulumi.StringOutput { return v.CollectionId }).(pulumi.StringOutput)
}

// Gets or sets collection name.
func (o PrivateStoreCollectionOutput) CollectionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateStoreCollection) pulumi.StringPtrOutput { return v.CollectionName }).(pulumi.StringPtrOutput)
}

// Indicating whether the collection is enabled or disabled.
func (o PrivateStoreCollectionOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PrivateStoreCollection) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The name of the resource.
func (o PrivateStoreCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateStoreCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the number of offers associated with the collection.
func (o PrivateStoreCollectionOutput) NumberOfOffers() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateStoreCollection) pulumi.Float64Output { return v.NumberOfOffers }).(pulumi.Float64Output)
}

// Gets or sets subscription ids list. Empty list indicates all subscriptions are selected, null indicates no update is done, explicit list indicates the explicit selected subscriptions. On insert, null is considered as bad request
func (o PrivateStoreCollectionOutput) SubscriptionsList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateStoreCollection) pulumi.StringArrayOutput { return v.SubscriptionsList }).(pulumi.StringArrayOutput)
}

// Metadata pertaining to creation and last modification of the resource
func (o PrivateStoreCollectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateStoreCollection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o PrivateStoreCollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateStoreCollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateStoreCollectionOutput{})
}
