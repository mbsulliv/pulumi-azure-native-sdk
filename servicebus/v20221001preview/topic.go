// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description of topic resource.
type Topic struct {
	pulumi.CustomResourceState

	// Last time the message was sent, or a request was received, for this topic.
	AccessedAt pulumi.StringOutput `pulumi:"accessedAt"`
	// ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrOutput `pulumi:"autoDeleteOnIdle"`
	// Message count details
	CountDetails MessageCountDetailsResponseOutput `pulumi:"countDetails"`
	// Exact time the message was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive pulumi.StringPtrOutput `pulumi:"defaultMessageTimeToLive"`
	// ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrOutput `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations pulumi.BoolPtrOutput `pulumi:"enableBatchedOperations"`
	// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
	EnableExpress pulumi.BoolPtrOutput `pulumi:"enableExpress"`
	// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
	EnablePartitioning pulumi.BoolPtrOutput `pulumi:"enablePartitioning"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Maximum size (in KB) of the message payload that can be accepted by the topic. This property is only used in Premium today and default is 1024.
	MaxMessageSizeInKilobytes pulumi.Float64PtrOutput `pulumi:"maxMessageSizeInKilobytes"`
	// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
	MaxSizeInMegabytes pulumi.IntPtrOutput `pulumi:"maxSizeInMegabytes"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Value indicating if this topic requires duplicate detection.
	RequiresDuplicateDetection pulumi.BoolPtrOutput `pulumi:"requiresDuplicateDetection"`
	// Size of the topic, in bytes.
	SizeInBytes pulumi.Float64Output `pulumi:"sizeInBytes"`
	// Enumerates the possible values for the status of a messaging entity.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Number of subscriptions.
	SubscriptionCount pulumi.IntOutput `pulumi:"subscriptionCount"`
	// Value that indicates whether the topic supports ordering.
	SupportOrdering pulumi.BoolPtrOutput `pulumi:"supportOrdering"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type pulumi.StringOutput `pulumi:"type"`
	// The exact time the message was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20140901:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20211101:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20220101preview:Topic"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Topic
	err := ctx.RegisterResource("azure-native:servicebus/v20221001preview:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("azure-native:servicebus/v20221001preview:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
}

type TopicState struct {
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive *string `pulumi:"defaultMessageTimeToLive"`
	// ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow *string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
	EnableExpress *bool `pulumi:"enableExpress"`
	// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
	EnablePartitioning *bool `pulumi:"enablePartitioning"`
	// Maximum size (in KB) of the message payload that can be accepted by the topic. This property is only used in Premium today and default is 1024.
	MaxMessageSizeInKilobytes *float64 `pulumi:"maxMessageSizeInKilobytes"`
	// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
	MaxSizeInMegabytes *int `pulumi:"maxSizeInMegabytes"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Value indicating if this topic requires duplicate detection.
	RequiresDuplicateDetection *bool `pulumi:"requiresDuplicateDetection"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Enumerates the possible values for the status of a messaging entity.
	Status *EntityStatus `pulumi:"status"`
	// Value that indicates whether the topic supports ordering.
	SupportOrdering *bool `pulumi:"supportOrdering"`
	// The topic name.
	TopicName *string `pulumi:"topicName"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive pulumi.StringPtrInput
	// ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations pulumi.BoolPtrInput
	// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
	EnableExpress pulumi.BoolPtrInput
	// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
	EnablePartitioning pulumi.BoolPtrInput
	// Maximum size (in KB) of the message payload that can be accepted by the topic. This property is only used in Premium today and default is 1024.
	MaxMessageSizeInKilobytes pulumi.Float64PtrInput
	// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
	MaxSizeInMegabytes pulumi.IntPtrInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// Value indicating if this topic requires duplicate detection.
	RequiresDuplicateDetection pulumi.BoolPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Enumerates the possible values for the status of a messaging entity.
	Status EntityStatusPtrInput
	// Value that indicates whether the topic supports ordering.
	SupportOrdering pulumi.BoolPtrInput
	// The topic name.
	TopicName pulumi.StringPtrInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

func (i *Topic) ToOutput(ctx context.Context) pulumix.Output[*Topic] {
	return pulumix.Output[*Topic]{
		OutputState: i.ToTopicOutputWithContext(ctx).OutputState,
	}
}

type TopicOutput struct{ *pulumi.OutputState }

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

func (o TopicOutput) ToOutput(ctx context.Context) pulumix.Output[*Topic] {
	return pulumix.Output[*Topic]{
		OutputState: o.OutputState,
	}
}

// Last time the message was sent, or a request was received, for this topic.
func (o TopicOutput) AccessedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.AccessedAt }).(pulumi.StringOutput)
}

// ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
func (o TopicOutput) AutoDeleteOnIdle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.AutoDeleteOnIdle }).(pulumi.StringPtrOutput)
}

// Message count details
func (o TopicOutput) CountDetails() MessageCountDetailsResponseOutput {
	return o.ApplyT(func(v *Topic) MessageCountDetailsResponseOutput { return v.CountDetails }).(MessageCountDetailsResponseOutput)
}

// Exact time the message was created.
func (o TopicOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
func (o TopicOutput) DefaultMessageTimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.DefaultMessageTimeToLive }).(pulumi.StringPtrOutput)
}

// ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
func (o TopicOutput) DuplicateDetectionHistoryTimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.DuplicateDetectionHistoryTimeWindow }).(pulumi.StringPtrOutput)
}

// Value that indicates whether server-side batched operations are enabled.
func (o TopicOutput) EnableBatchedOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.EnableBatchedOperations }).(pulumi.BoolPtrOutput)
}

// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
func (o TopicOutput) EnableExpress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.EnableExpress }).(pulumi.BoolPtrOutput)
}

// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
func (o TopicOutput) EnablePartitioning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.EnablePartitioning }).(pulumi.BoolPtrOutput)
}

// The geo-location where the resource lives
func (o TopicOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Maximum size (in KB) of the message payload that can be accepted by the topic. This property is only used in Premium today and default is 1024.
func (o TopicOutput) MaxMessageSizeInKilobytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.Float64PtrOutput { return v.MaxMessageSizeInKilobytes }).(pulumi.Float64PtrOutput)
}

// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
func (o TopicOutput) MaxSizeInMegabytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntPtrOutput { return v.MaxSizeInMegabytes }).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o TopicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Value indicating if this topic requires duplicate detection.
func (o TopicOutput) RequiresDuplicateDetection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.RequiresDuplicateDetection }).(pulumi.BoolPtrOutput)
}

// Size of the topic, in bytes.
func (o TopicOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *Topic) pulumi.Float64Output { return v.SizeInBytes }).(pulumi.Float64Output)
}

// Enumerates the possible values for the status of a messaging entity.
func (o TopicOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// Number of subscriptions.
func (o TopicOutput) SubscriptionCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.SubscriptionCount }).(pulumi.IntOutput)
}

// Value that indicates whether the topic supports ordering.
func (o TopicOutput) SupportOrdering() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.SupportOrdering }).(pulumi.BoolPtrOutput)
}

// The system meta data relating to this resource.
func (o TopicOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Topic) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o TopicOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The exact time the message was updated.
func (o TopicOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicOutput{})
}
