// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description of subscription resource.
type Subscription struct {
	pulumi.CustomResourceState

	// Last time there was a receive request to this subscription.
	AccessedAt pulumi.StringOutput `pulumi:"accessedAt"`
	// ISO 8061 timeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrOutput `pulumi:"autoDeleteOnIdle"`
	// Properties specific to client affine subscriptions.
	ClientAffineProperties SBClientAffinePropertiesResponsePtrOutput `pulumi:"clientAffineProperties"`
	// Message count details
	CountDetails MessageCountDetailsResponseOutput `pulumi:"countDetails"`
	// Exact time the message was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Value that indicates whether a subscription has dead letter support on filter evaluation exceptions.
	DeadLetteringOnFilterEvaluationExceptions pulumi.BoolPtrOutput `pulumi:"deadLetteringOnFilterEvaluationExceptions"`
	// Value that indicates whether a subscription has dead letter support when a message expires.
	DeadLetteringOnMessageExpiration pulumi.BoolPtrOutput `pulumi:"deadLetteringOnMessageExpiration"`
	// ISO 8061 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive pulumi.StringPtrOutput `pulumi:"defaultMessageTimeToLive"`
	// ISO 8601 timeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrOutput `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations pulumi.BoolPtrOutput `pulumi:"enableBatchedOperations"`
	// Queue/Topic name to forward the Dead Letter message
	ForwardDeadLetteredMessagesTo pulumi.StringPtrOutput `pulumi:"forwardDeadLetteredMessagesTo"`
	// Queue/Topic name to forward the messages
	ForwardTo pulumi.StringPtrOutput `pulumi:"forwardTo"`
	// Value that indicates whether the subscription has an affinity to the client id.
	IsClientAffine pulumi.BoolPtrOutput `pulumi:"isClientAffine"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// ISO 8061 lock duration timespan for the subscription. The default value is 1 minute.
	LockDuration pulumi.StringPtrOutput `pulumi:"lockDuration"`
	// Number of maximum deliveries.
	MaxDeliveryCount pulumi.IntPtrOutput `pulumi:"maxDeliveryCount"`
	// Number of messages.
	MessageCount pulumi.Float64Output `pulumi:"messageCount"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Value indicating if a subscription supports the concept of sessions.
	RequiresSession pulumi.BoolPtrOutput `pulumi:"requiresSession"`
	// Enumerates the possible values for the status of a messaging entity.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type pulumi.StringOutput `pulumi:"type"`
	// The exact time the message was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20140901:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20211101:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20221001preview:Subscription"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Subscription
	err := ctx.RegisterResource("azure-native:servicebus/v20220101preview:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("azure-native:servicebus/v20220101preview:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
}

type SubscriptionState struct {
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	// ISO 8061 timeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// Properties specific to client affine subscriptions.
	ClientAffineProperties *SBClientAffineProperties `pulumi:"clientAffineProperties"`
	// Value that indicates whether a subscription has dead letter support on filter evaluation exceptions.
	DeadLetteringOnFilterEvaluationExceptions *bool `pulumi:"deadLetteringOnFilterEvaluationExceptions"`
	// Value that indicates whether a subscription has dead letter support when a message expires.
	DeadLetteringOnMessageExpiration *bool `pulumi:"deadLetteringOnMessageExpiration"`
	// ISO 8061 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive *string `pulumi:"defaultMessageTimeToLive"`
	// ISO 8601 timeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow *string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// Queue/Topic name to forward the Dead Letter message
	ForwardDeadLetteredMessagesTo *string `pulumi:"forwardDeadLetteredMessagesTo"`
	// Queue/Topic name to forward the messages
	ForwardTo *string `pulumi:"forwardTo"`
	// Value that indicates whether the subscription has an affinity to the client id.
	IsClientAffine *bool `pulumi:"isClientAffine"`
	// ISO 8061 lock duration timespan for the subscription. The default value is 1 minute.
	LockDuration *string `pulumi:"lockDuration"`
	// Number of maximum deliveries.
	MaxDeliveryCount *int `pulumi:"maxDeliveryCount"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Value indicating if a subscription supports the concept of sessions.
	RequiresSession *bool `pulumi:"requiresSession"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Enumerates the possible values for the status of a messaging entity.
	Status *EntityStatus `pulumi:"status"`
	// The subscription name.
	SubscriptionName *string `pulumi:"subscriptionName"`
	// The topic name.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// ISO 8061 timeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// Properties specific to client affine subscriptions.
	ClientAffineProperties SBClientAffinePropertiesPtrInput
	// Value that indicates whether a subscription has dead letter support on filter evaluation exceptions.
	DeadLetteringOnFilterEvaluationExceptions pulumi.BoolPtrInput
	// Value that indicates whether a subscription has dead letter support when a message expires.
	DeadLetteringOnMessageExpiration pulumi.BoolPtrInput
	// ISO 8061 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive pulumi.StringPtrInput
	// ISO 8601 timeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations pulumi.BoolPtrInput
	// Queue/Topic name to forward the Dead Letter message
	ForwardDeadLetteredMessagesTo pulumi.StringPtrInput
	// Queue/Topic name to forward the messages
	ForwardTo pulumi.StringPtrInput
	// Value that indicates whether the subscription has an affinity to the client id.
	IsClientAffine pulumi.BoolPtrInput
	// ISO 8061 lock duration timespan for the subscription. The default value is 1 minute.
	LockDuration pulumi.StringPtrInput
	// Number of maximum deliveries.
	MaxDeliveryCount pulumi.IntPtrInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// Value indicating if a subscription supports the concept of sessions.
	RequiresSession pulumi.BoolPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Enumerates the possible values for the status of a messaging entity.
	Status EntityStatusPtrInput
	// The subscription name.
	SubscriptionName pulumi.StringPtrInput
	// The topic name.
	TopicName pulumi.StringInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

type SubscriptionInput interface {
	pulumi.Input

	ToSubscriptionOutput() SubscriptionOutput
	ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput
}

func (*Subscription) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (i *Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

func (i *Subscription) ToOutput(ctx context.Context) pulumix.Output[*Subscription] {
	return pulumix.Output[*Subscription]{
		OutputState: i.ToSubscriptionOutputWithContext(ctx).OutputState,
	}
}

type SubscriptionOutput struct{ *pulumi.OutputState }

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToOutput(ctx context.Context) pulumix.Output[*Subscription] {
	return pulumix.Output[*Subscription]{
		OutputState: o.OutputState,
	}
}

// Last time there was a receive request to this subscription.
func (o SubscriptionOutput) AccessedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.AccessedAt }).(pulumi.StringOutput)
}

// ISO 8061 timeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
func (o SubscriptionOutput) AutoDeleteOnIdle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.AutoDeleteOnIdle }).(pulumi.StringPtrOutput)
}

// Properties specific to client affine subscriptions.
func (o SubscriptionOutput) ClientAffineProperties() SBClientAffinePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Subscription) SBClientAffinePropertiesResponsePtrOutput { return v.ClientAffineProperties }).(SBClientAffinePropertiesResponsePtrOutput)
}

// Message count details
func (o SubscriptionOutput) CountDetails() MessageCountDetailsResponseOutput {
	return o.ApplyT(func(v *Subscription) MessageCountDetailsResponseOutput { return v.CountDetails }).(MessageCountDetailsResponseOutput)
}

// Exact time the message was created.
func (o SubscriptionOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Value that indicates whether a subscription has dead letter support on filter evaluation exceptions.
func (o SubscriptionOutput) DeadLetteringOnFilterEvaluationExceptions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.DeadLetteringOnFilterEvaluationExceptions }).(pulumi.BoolPtrOutput)
}

// Value that indicates whether a subscription has dead letter support when a message expires.
func (o SubscriptionOutput) DeadLetteringOnMessageExpiration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.DeadLetteringOnMessageExpiration }).(pulumi.BoolPtrOutput)
}

// ISO 8061 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
func (o SubscriptionOutput) DefaultMessageTimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.DefaultMessageTimeToLive }).(pulumi.StringPtrOutput)
}

// ISO 8601 timeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
func (o SubscriptionOutput) DuplicateDetectionHistoryTimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.DuplicateDetectionHistoryTimeWindow }).(pulumi.StringPtrOutput)
}

// Value that indicates whether server-side batched operations are enabled.
func (o SubscriptionOutput) EnableBatchedOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.EnableBatchedOperations }).(pulumi.BoolPtrOutput)
}

// Queue/Topic name to forward the Dead Letter message
func (o SubscriptionOutput) ForwardDeadLetteredMessagesTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.ForwardDeadLetteredMessagesTo }).(pulumi.StringPtrOutput)
}

// Queue/Topic name to forward the messages
func (o SubscriptionOutput) ForwardTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.ForwardTo }).(pulumi.StringPtrOutput)
}

// Value that indicates whether the subscription has an affinity to the client id.
func (o SubscriptionOutput) IsClientAffine() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.IsClientAffine }).(pulumi.BoolPtrOutput)
}

// The geo-location where the resource lives
func (o SubscriptionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// ISO 8061 lock duration timespan for the subscription. The default value is 1 minute.
func (o SubscriptionOutput) LockDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.LockDuration }).(pulumi.StringPtrOutput)
}

// Number of maximum deliveries.
func (o SubscriptionOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.IntPtrOutput { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

// Number of messages.
func (o SubscriptionOutput) MessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v *Subscription) pulumi.Float64Output { return v.MessageCount }).(pulumi.Float64Output)
}

// The name of the resource
func (o SubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Value indicating if a subscription supports the concept of sessions.
func (o SubscriptionOutput) RequiresSession() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.RequiresSession }).(pulumi.BoolPtrOutput)
}

// Enumerates the possible values for the status of a messaging entity.
func (o SubscriptionOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// The system meta data relating to this resource.
func (o SubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Subscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o SubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The exact time the message was updated.
func (o SubscriptionOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionOutput{})
}
