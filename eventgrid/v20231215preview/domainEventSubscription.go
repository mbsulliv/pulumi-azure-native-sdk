// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231215preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Event Subscription.
type DomainEventSubscription struct {
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

// NewDomainEventSubscription registers a new resource with the given unique name, arguments, and options.
func NewDomainEventSubscription(ctx *pulumi.Context,
	name string, args *DomainEventSubscriptionArgs, opts ...pulumi.ResourceOption) (*DomainEventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
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
			Type: pulumi.String("azure-native:eventgrid:DomainEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:DomainEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:DomainEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20230601preview:DomainEventSubscription"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DomainEventSubscription
	err := ctx.RegisterResource("azure-native:eventgrid/v20231215preview:DomainEventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainEventSubscription gets an existing DomainEventSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainEventSubscriptionState, opts ...pulumi.ResourceOption) (*DomainEventSubscription, error) {
	var resource DomainEventSubscription
	err := ctx.ReadResource("azure-native:eventgrid/v20231215preview:DomainEventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainEventSubscription resources.
type domainEventSubscriptionState struct {
}

type DomainEventSubscriptionState struct {
}

func (DomainEventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainEventSubscriptionState)(nil)).Elem()
}

type domainEventSubscriptionArgs struct {
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
	// Name of the domain topic.
	DomainName string `pulumi:"domainName"`
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
}

// The set of arguments for constructing a DomainEventSubscription resource.
type DomainEventSubscriptionArgs struct {
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
	// Name of the domain topic.
	DomainName pulumi.StringInput
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
}

func (DomainEventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainEventSubscriptionArgs)(nil)).Elem()
}

type DomainEventSubscriptionInput interface {
	pulumi.Input

	ToDomainEventSubscriptionOutput() DomainEventSubscriptionOutput
	ToDomainEventSubscriptionOutputWithContext(ctx context.Context) DomainEventSubscriptionOutput
}

func (*DomainEventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainEventSubscription)(nil)).Elem()
}

func (i *DomainEventSubscription) ToDomainEventSubscriptionOutput() DomainEventSubscriptionOutput {
	return i.ToDomainEventSubscriptionOutputWithContext(context.Background())
}

func (i *DomainEventSubscription) ToDomainEventSubscriptionOutputWithContext(ctx context.Context) DomainEventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainEventSubscriptionOutput)
}

type DomainEventSubscriptionOutput struct{ *pulumi.OutputState }

func (DomainEventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainEventSubscription)(nil)).Elem()
}

func (o DomainEventSubscriptionOutput) ToDomainEventSubscriptionOutput() DomainEventSubscriptionOutput {
	return o
}

func (o DomainEventSubscriptionOutput) ToDomainEventSubscriptionOutputWithContext(ctx context.Context) DomainEventSubscriptionOutput {
	return o
}

// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
func (o DomainEventSubscriptionOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) StorageBlobDeadLetterDestinationResponsePtrOutput {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
func (o DomainEventSubscriptionOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) DeadLetterWithResourceIdentityResponsePtrOutput {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

// Information about the destination where events have to be delivered for the event subscription.
// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
func (o DomainEventSubscriptionOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) DeliveryWithResourceIdentityResponsePtrOutput {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

// Information about the destination where events have to be delivered for the event subscription.
// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
func (o DomainEventSubscriptionOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.AnyOutput { return v.Destination }).(pulumi.AnyOutput)
}

// The event delivery schema for the event subscription.
func (o DomainEventSubscriptionOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringPtrOutput { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

// Expiration time of the event subscription.
func (o DomainEventSubscriptionOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringPtrOutput { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

// Information about the filter for the event subscription.
func (o DomainEventSubscriptionOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) EventSubscriptionFilterResponsePtrOutput { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

// List of user defined labels.
func (o DomainEventSubscriptionOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// Name of the resource.
func (o DomainEventSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the event subscription.
func (o DomainEventSubscriptionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
func (o DomainEventSubscriptionOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) RetryPolicyResponsePtrOutput { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

// The system metadata relating to Event Subscription resource.
func (o DomainEventSubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DomainEventSubscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Name of the topic of the event subscription.
func (o DomainEventSubscriptionOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

// Type of the resource.
func (o DomainEventSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainEventSubscriptionOutput{})
}
