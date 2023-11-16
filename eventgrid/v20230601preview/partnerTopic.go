// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Event Grid Partner Topic.
type PartnerTopic struct {
	pulumi.CustomResourceState

	// Activation state of the partner topic.
	ActivationState pulumi.StringPtrOutput `pulumi:"activationState"`
	// Event Type information from the corresponding event channel.
	EventTypeInfo EventTypeInfoResponsePtrOutput `pulumi:"eventTypeInfo"`
	// Expiration time of the partner topic. If this timer expires while the partner topic is still never activated,
	// the partner topic and corresponding event channel are deleted.
	ExpirationTimeIfNotActivatedUtc pulumi.StringPtrOutput `pulumi:"expirationTimeIfNotActivatedUtc"`
	// Identity information for the Partner Topic resource.
	Identity IdentityInfoResponsePtrOutput `pulumi:"identity"`
	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Context or helpful message that can be used during the approval process by the subscriber.
	MessageForActivation pulumi.StringPtrOutput `pulumi:"messageForActivation"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The immutableId of the corresponding partner registration.
	PartnerRegistrationImmutableId pulumi.StringPtrOutput `pulumi:"partnerRegistrationImmutableId"`
	// Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
	// This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
	PartnerTopicFriendlyDescription pulumi.StringPtrOutput `pulumi:"partnerTopicFriendlyDescription"`
	// Provisioning state of the partner topic.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Source associated with this partner topic. This represents a unique partner resource.
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// The system metadata relating to Partner Topic resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPartnerTopic registers a new resource with the given unique name, arguments, and options.
func NewPartnerTopic(ctx *pulumi.Context,
	name string, args *PartnerTopicArgs, opts ...pulumi.ResourceOption) (*PartnerTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:PartnerTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:PartnerTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:PartnerTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20231215preview:PartnerTopic"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PartnerTopic
	err := ctx.RegisterResource("azure-native:eventgrid/v20230601preview:PartnerTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPartnerTopic gets an existing PartnerTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPartnerTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerTopicState, opts ...pulumi.ResourceOption) (*PartnerTopic, error) {
	var resource PartnerTopic
	err := ctx.ReadResource("azure-native:eventgrid/v20230601preview:PartnerTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PartnerTopic resources.
type partnerTopicState struct {
}

type PartnerTopicState struct {
}

func (PartnerTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerTopicState)(nil)).Elem()
}

type partnerTopicArgs struct {
	// Activation state of the partner topic.
	ActivationState *string `pulumi:"activationState"`
	// Event Type information from the corresponding event channel.
	EventTypeInfo *EventTypeInfo `pulumi:"eventTypeInfo"`
	// Expiration time of the partner topic. If this timer expires while the partner topic is still never activated,
	// the partner topic and corresponding event channel are deleted.
	ExpirationTimeIfNotActivatedUtc *string `pulumi:"expirationTimeIfNotActivatedUtc"`
	// Identity information for the Partner Topic resource.
	Identity *IdentityInfo `pulumi:"identity"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// Context or helpful message that can be used during the approval process by the subscriber.
	MessageForActivation *string `pulumi:"messageForActivation"`
	// The immutableId of the corresponding partner registration.
	PartnerRegistrationImmutableId *string `pulumi:"partnerRegistrationImmutableId"`
	// Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
	// This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
	PartnerTopicFriendlyDescription *string `pulumi:"partnerTopicFriendlyDescription"`
	// Name of the partner topic.
	PartnerTopicName *string `pulumi:"partnerTopicName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Source associated with this partner topic. This represents a unique partner resource.
	Source *string `pulumi:"source"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PartnerTopic resource.
type PartnerTopicArgs struct {
	// Activation state of the partner topic.
	ActivationState pulumi.StringPtrInput
	// Event Type information from the corresponding event channel.
	EventTypeInfo EventTypeInfoPtrInput
	// Expiration time of the partner topic. If this timer expires while the partner topic is still never activated,
	// the partner topic and corresponding event channel are deleted.
	ExpirationTimeIfNotActivatedUtc pulumi.StringPtrInput
	// Identity information for the Partner Topic resource.
	Identity IdentityInfoPtrInput
	// Location of the resource.
	Location pulumi.StringPtrInput
	// Context or helpful message that can be used during the approval process by the subscriber.
	MessageForActivation pulumi.StringPtrInput
	// The immutableId of the corresponding partner registration.
	PartnerRegistrationImmutableId pulumi.StringPtrInput
	// Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
	// This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
	PartnerTopicFriendlyDescription pulumi.StringPtrInput
	// Name of the partner topic.
	PartnerTopicName pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// Source associated with this partner topic. This represents a unique partner resource.
	Source pulumi.StringPtrInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
}

func (PartnerTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerTopicArgs)(nil)).Elem()
}

type PartnerTopicInput interface {
	pulumi.Input

	ToPartnerTopicOutput() PartnerTopicOutput
	ToPartnerTopicOutputWithContext(ctx context.Context) PartnerTopicOutput
}

func (*PartnerTopic) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerTopic)(nil)).Elem()
}

func (i *PartnerTopic) ToPartnerTopicOutput() PartnerTopicOutput {
	return i.ToPartnerTopicOutputWithContext(context.Background())
}

func (i *PartnerTopic) ToPartnerTopicOutputWithContext(ctx context.Context) PartnerTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerTopicOutput)
}

type PartnerTopicOutput struct{ *pulumi.OutputState }

func (PartnerTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerTopic)(nil)).Elem()
}

func (o PartnerTopicOutput) ToPartnerTopicOutput() PartnerTopicOutput {
	return o
}

func (o PartnerTopicOutput) ToPartnerTopicOutputWithContext(ctx context.Context) PartnerTopicOutput {
	return o
}

// Activation state of the partner topic.
func (o PartnerTopicOutput) ActivationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringPtrOutput { return v.ActivationState }).(pulumi.StringPtrOutput)
}

// Event Type information from the corresponding event channel.
func (o PartnerTopicOutput) EventTypeInfo() EventTypeInfoResponsePtrOutput {
	return o.ApplyT(func(v *PartnerTopic) EventTypeInfoResponsePtrOutput { return v.EventTypeInfo }).(EventTypeInfoResponsePtrOutput)
}

// Expiration time of the partner topic. If this timer expires while the partner topic is still never activated,
// the partner topic and corresponding event channel are deleted.
func (o PartnerTopicOutput) ExpirationTimeIfNotActivatedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringPtrOutput { return v.ExpirationTimeIfNotActivatedUtc }).(pulumi.StringPtrOutput)
}

// Identity information for the Partner Topic resource.
func (o PartnerTopicOutput) Identity() IdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v *PartnerTopic) IdentityInfoResponsePtrOutput { return v.Identity }).(IdentityInfoResponsePtrOutput)
}

// Location of the resource.
func (o PartnerTopicOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Context or helpful message that can be used during the approval process by the subscriber.
func (o PartnerTopicOutput) MessageForActivation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringPtrOutput { return v.MessageForActivation }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o PartnerTopicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The immutableId of the corresponding partner registration.
func (o PartnerTopicOutput) PartnerRegistrationImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringPtrOutput { return v.PartnerRegistrationImmutableId }).(pulumi.StringPtrOutput)
}

// Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
// This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
func (o PartnerTopicOutput) PartnerTopicFriendlyDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringPtrOutput { return v.PartnerTopicFriendlyDescription }).(pulumi.StringPtrOutput)
}

// Provisioning state of the partner topic.
func (o PartnerTopicOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Source associated with this partner topic. This represents a unique partner resource.
func (o PartnerTopicOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

// The system metadata relating to Partner Topic resource.
func (o PartnerTopicOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PartnerTopic) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags of the resource.
func (o PartnerTopicOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the resource.
func (o PartnerTopicOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PartnerTopicOutput{})
}
