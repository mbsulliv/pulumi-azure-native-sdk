// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220615

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Event Subscription
type EventSubscription struct {
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

// NewEventSubscription registers a new resource with the given unique name, arguments, and options.
func NewEventSubscription(ctx *pulumi.Context,
	name string, args *EventSubscriptionArgs, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
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
			Type: pulumi.String("azure-native:eventgrid:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20170615preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20170915preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180101:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180501preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180915preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190101:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190201preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190601:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200101preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200601:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211201:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20230601preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20231215preview:EventSubscription"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EventSubscription
	err := ctx.RegisterResource("azure-native:eventgrid/v20220615:EventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventSubscription gets an existing EventSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventSubscriptionState, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	var resource EventSubscription
	err := ctx.ReadResource("azure-native:eventgrid/v20220615:EventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventSubscription resources.
type eventSubscriptionState struct {
}

type EventSubscriptionState struct {
}

func (EventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionState)(nil)).Elem()
}

type eventSubscriptionArgs struct {
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
	// Name of the event subscription. Event subscription names must be between 3 and 64 characters in length and should use alphanumeric letters only.
	EventSubscriptionName *string `pulumi:"eventSubscriptionName"`
	// Expiration time of the event subscription.
	ExpirationTimeUtc *string `pulumi:"expirationTimeUtc"`
	// Information about the filter for the event subscription.
	Filter *EventSubscriptionFilter `pulumi:"filter"`
	// List of user defined labels.
	Labels []string `pulumi:"labels"`
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy *RetryPolicy `pulumi:"retryPolicy"`
	// The identifier of the resource to which the event subscription needs to be created or updated. The scope can be a subscription, or a resource group, or a top level resource belonging to a resource provider namespace, or an EventGrid topic. For example, use '/subscriptions/{subscriptionId}/' for a subscription, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for a resource group, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}' for a resource, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}' for an EventGrid topic.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a EventSubscription resource.
type EventSubscriptionArgs struct {
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
	// Name of the event subscription. Event subscription names must be between 3 and 64 characters in length and should use alphanumeric letters only.
	EventSubscriptionName pulumi.StringPtrInput
	// Expiration time of the event subscription.
	ExpirationTimeUtc pulumi.StringPtrInput
	// Information about the filter for the event subscription.
	Filter EventSubscriptionFilterPtrInput
	// List of user defined labels.
	Labels pulumi.StringArrayInput
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy RetryPolicyPtrInput
	// The identifier of the resource to which the event subscription needs to be created or updated. The scope can be a subscription, or a resource group, or a top level resource belonging to a resource provider namespace, or an EventGrid topic. For example, use '/subscriptions/{subscriptionId}/' for a subscription, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for a resource group, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}' for a resource, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}' for an EventGrid topic.
	Scope pulumi.StringInput
}

func (EventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionArgs)(nil)).Elem()
}

type EventSubscriptionInput interface {
	pulumi.Input

	ToEventSubscriptionOutput() EventSubscriptionOutput
	ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput
}

func (*EventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscription)(nil)).Elem()
}

func (i *EventSubscription) ToEventSubscriptionOutput() EventSubscriptionOutput {
	return i.ToEventSubscriptionOutputWithContext(context.Background())
}

func (i *EventSubscription) ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionOutput)
}

type EventSubscriptionOutput struct{ *pulumi.OutputState }

func (EventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscription)(nil)).Elem()
}

func (o EventSubscriptionOutput) ToEventSubscriptionOutput() EventSubscriptionOutput {
	return o
}

func (o EventSubscriptionOutput) ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput {
	return o
}

// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
func (o EventSubscriptionOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v *EventSubscription) StorageBlobDeadLetterDestinationResponsePtrOutput {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
func (o EventSubscriptionOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *EventSubscription) DeadLetterWithResourceIdentityResponsePtrOutput {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

// Information about the destination where events have to be delivered for the event subscription.
// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
func (o EventSubscriptionOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *EventSubscription) DeliveryWithResourceIdentityResponsePtrOutput {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

// Information about the destination where events have to be delivered for the event subscription.
// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
func (o EventSubscriptionOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.AnyOutput { return v.Destination }).(pulumi.AnyOutput)
}

// The event delivery schema for the event subscription.
func (o EventSubscriptionOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringPtrOutput { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

// Expiration time of the event subscription.
func (o EventSubscriptionOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringPtrOutput { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

// Information about the filter for the event subscription.
func (o EventSubscriptionOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v *EventSubscription) EventSubscriptionFilterResponsePtrOutput { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

// List of user defined labels.
func (o EventSubscriptionOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// Name of the resource.
func (o EventSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the event subscription.
func (o EventSubscriptionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
func (o EventSubscriptionOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v *EventSubscription) RetryPolicyResponsePtrOutput { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

// The system metadata relating to Event Subscription resource.
func (o EventSubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EventSubscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Name of the topic of the event subscription.
func (o EventSubscriptionOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

// Type of the resource.
func (o EventSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventSubscriptionOutput{})
}
