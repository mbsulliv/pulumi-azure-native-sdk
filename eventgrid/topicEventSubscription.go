// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Event Subscription
// Azure REST API version: 2022-06-15. Prior API version in Azure Native 1.x: 2021-10-15-preview.
//
// Other available API versions: 2023-06-01-preview.
type TopicEventSubscription struct {
	pulumi.CustomResourceState

	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterDestination StorageBlobDeadLetterDestinationResponsePtrOutput `pulumi:"deadLetterDestination"`
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterWithResourceIdentity DeadLetterWithResourceIdentityResponsePtrOutput `pulumi:"deadLetterWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeliveryWithResourceIdentity DeliveryWithResourceIdentityResponsePtrOutput `pulumi:"deliveryWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	Destination pulumi.AnyOutput `pulumi:"destination"`
	// The event delivery schema for the event subscription.
	EventDeliverySchema pulumi.StringPtrOutput `pulumi:"eventDeliverySchema"`
	// Expiration time of the event subscription.
	ExpirationTimeUtc pulumi.StringPtrOutput `pulumi:"expirationTimeUtc"`
	// Information about the filter for the event subscription.
	Filter EventSubscriptionFilterResponsePtrOutput `pulumi:"filter"`
	// List of user defined labels.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the event subscription.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy RetryPolicyResponsePtrOutput `pulumi:"retryPolicy"`
	// The system metadata relating to Event Subscription resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Name of the topic of the event subscription.
	Topic pulumi.StringOutput `pulumi:"topic"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTopicEventSubscription registers a new resource with the given unique name, arguments, and options.
func NewTopicEventSubscription(ctx *pulumi.Context,
	name string, args *TopicEventSubscriptionArgs, opts ...pulumi.ResourceOption) (*TopicEventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	if args.EventDeliverySchema == nil {
		args.EventDeliverySchema = pulumi.StringPtr("EventGridSchema")
	}
	if args.Filter != nil {
		args.Filter = args.Filter.ToEventSubscriptionFilterPtrOutput().ApplyT(func(v *EventSubscriptionFilter) *EventSubscriptionFilter { return v.Defaults() }).(EventSubscriptionFilterPtrOutput)
	}
	if args.RetryPolicy != nil {
		args.RetryPolicy = args.RetryPolicy.ToRetryPolicyPtrOutput().ApplyT(func(v *RetryPolicy) *RetryPolicy { return v.Defaults() }).(RetryPolicyPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:TopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:TopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20230601preview:TopicEventSubscription"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TopicEventSubscription
	err := ctx.RegisterResource("azure-native:eventgrid:TopicEventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicEventSubscription gets an existing TopicEventSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicEventSubscriptionState, opts ...pulumi.ResourceOption) (*TopicEventSubscription, error) {
	var resource TopicEventSubscription
	err := ctx.ReadResource("azure-native:eventgrid:TopicEventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicEventSubscription resources.
type topicEventSubscriptionState struct {
}

type TopicEventSubscriptionState struct {
}

func (TopicEventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicEventSubscriptionState)(nil)).Elem()
}

type topicEventSubscriptionArgs struct {
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterDestination *StorageBlobDeadLetterDestination `pulumi:"deadLetterDestination"`
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterWithResourceIdentity *DeadLetterWithResourceIdentity `pulumi:"deadLetterWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeliveryWithResourceIdentity *DeliveryWithResourceIdentity `pulumi:"deliveryWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	Destination interface{} `pulumi:"destination"`
	// The event delivery schema for the event subscription.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// Name of the event subscription to be created. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName *string `pulumi:"eventSubscriptionName"`
	// Expiration time of the event subscription.
	ExpirationTimeUtc *string `pulumi:"expirationTimeUtc"`
	// Information about the filter for the event subscription.
	Filter *EventSubscriptionFilter `pulumi:"filter"`
	// List of user defined labels.
	Labels []string `pulumi:"labels"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy *RetryPolicy `pulumi:"retryPolicy"`
	// Name of the domain topic.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a TopicEventSubscription resource.
type TopicEventSubscriptionArgs struct {
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterDestination StorageBlobDeadLetterDestinationPtrInput
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterWithResourceIdentity DeadLetterWithResourceIdentityPtrInput
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeliveryWithResourceIdentity DeliveryWithResourceIdentityPtrInput
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	Destination pulumi.Input
	// The event delivery schema for the event subscription.
	EventDeliverySchema pulumi.StringPtrInput
	// Name of the event subscription to be created. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName pulumi.StringPtrInput
	// Expiration time of the event subscription.
	ExpirationTimeUtc pulumi.StringPtrInput
	// Information about the filter for the event subscription.
	Filter EventSubscriptionFilterPtrInput
	// List of user defined labels.
	Labels pulumi.StringArrayInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy RetryPolicyPtrInput
	// Name of the domain topic.
	TopicName pulumi.StringInput
}

func (TopicEventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicEventSubscriptionArgs)(nil)).Elem()
}

type TopicEventSubscriptionInput interface {
	pulumi.Input

	ToTopicEventSubscriptionOutput() TopicEventSubscriptionOutput
	ToTopicEventSubscriptionOutputWithContext(ctx context.Context) TopicEventSubscriptionOutput
}

func (*TopicEventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicEventSubscription)(nil)).Elem()
}

func (i *TopicEventSubscription) ToTopicEventSubscriptionOutput() TopicEventSubscriptionOutput {
	return i.ToTopicEventSubscriptionOutputWithContext(context.Background())
}

func (i *TopicEventSubscription) ToTopicEventSubscriptionOutputWithContext(ctx context.Context) TopicEventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicEventSubscriptionOutput)
}

func (i *TopicEventSubscription) ToOutput(ctx context.Context) pulumix.Output[*TopicEventSubscription] {
	return pulumix.Output[*TopicEventSubscription]{
		OutputState: i.ToTopicEventSubscriptionOutputWithContext(ctx).OutputState,
	}
}

type TopicEventSubscriptionOutput struct{ *pulumi.OutputState }

func (TopicEventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicEventSubscription)(nil)).Elem()
}

func (o TopicEventSubscriptionOutput) ToTopicEventSubscriptionOutput() TopicEventSubscriptionOutput {
	return o
}

func (o TopicEventSubscriptionOutput) ToTopicEventSubscriptionOutputWithContext(ctx context.Context) TopicEventSubscriptionOutput {
	return o
}

func (o TopicEventSubscriptionOutput) ToOutput(ctx context.Context) pulumix.Output[*TopicEventSubscription] {
	return pulumix.Output[*TopicEventSubscription]{
		OutputState: o.OutputState,
	}
}

// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
func (o TopicEventSubscriptionOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) StorageBlobDeadLetterDestinationResponsePtrOutput {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
func (o TopicEventSubscriptionOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) DeadLetterWithResourceIdentityResponsePtrOutput {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

// Information about the destination where events have to be delivered for the event subscription.
// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
func (o TopicEventSubscriptionOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) DeliveryWithResourceIdentityResponsePtrOutput {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

// Information about the destination where events have to be delivered for the event subscription.
// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
func (o TopicEventSubscriptionOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.AnyOutput { return v.Destination }).(pulumi.AnyOutput)
}

// The event delivery schema for the event subscription.
func (o TopicEventSubscriptionOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringPtrOutput { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

// Expiration time of the event subscription.
func (o TopicEventSubscriptionOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringPtrOutput { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

// Information about the filter for the event subscription.
func (o TopicEventSubscriptionOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) EventSubscriptionFilterResponsePtrOutput { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

// List of user defined labels.
func (o TopicEventSubscriptionOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// Name of the resource.
func (o TopicEventSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the event subscription.
func (o TopicEventSubscriptionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
func (o TopicEventSubscriptionOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) RetryPolicyResponsePtrOutput { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

// The system metadata relating to Event Subscription resource.
func (o TopicEventSubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TopicEventSubscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Name of the topic of the event subscription.
func (o TopicEventSubscriptionOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

// Type of the resource.
func (o TopicEventSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicEventSubscriptionOutput{})
}
