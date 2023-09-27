// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns a description for the specified topic.
func LookupTopic(ctx *pulumi.Context, args *LookupTopicArgs, opts ...pulumi.InvokeOption) (*LookupTopicResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTopicResult
	err := ctx.Invoke("azure-native:servicebus/v20221001preview:getTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicArgs struct {
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The topic name.
	TopicName string `pulumi:"topicName"`
}

// Description of topic resource.
type LookupTopicResult struct {
	// Last time the message was sent, or a request was received, for this topic.
	AccessedAt string `pulumi:"accessedAt"`
	// ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// Message count details
	CountDetails MessageCountDetailsResponse `pulumi:"countDetails"`
	// Exact time the message was created.
	CreatedAt string `pulumi:"createdAt"`
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
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Maximum size (in KB) of the message payload that can be accepted by the topic. This property is only used in Premium today and default is 1024.
	MaxMessageSizeInKilobytes *float64 `pulumi:"maxMessageSizeInKilobytes"`
	// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
	MaxSizeInMegabytes *int `pulumi:"maxSizeInMegabytes"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Value indicating if this topic requires duplicate detection.
	RequiresDuplicateDetection *bool `pulumi:"requiresDuplicateDetection"`
	// Size of the topic, in bytes.
	SizeInBytes float64 `pulumi:"sizeInBytes"`
	// Enumerates the possible values for the status of a messaging entity.
	Status *string `pulumi:"status"`
	// Number of subscriptions.
	SubscriptionCount int `pulumi:"subscriptionCount"`
	// Value that indicates whether the topic supports ordering.
	SupportOrdering *bool `pulumi:"supportOrdering"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type string `pulumi:"type"`
	// The exact time the message was updated.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupTopicOutput(ctx *pulumi.Context, args LookupTopicOutputArgs, opts ...pulumi.InvokeOption) LookupTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTopicResult, error) {
			args := v.(LookupTopicArgs)
			r, err := LookupTopic(ctx, &args, opts...)
			var s LookupTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTopicResultOutput)
}

type LookupTopicOutputArgs struct {
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The topic name.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (LookupTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicArgs)(nil)).Elem()
}

// Description of topic resource.
type LookupTopicResultOutput struct{ *pulumi.OutputState }

func (LookupTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicResult)(nil)).Elem()
}

func (o LookupTopicResultOutput) ToLookupTopicResultOutput() LookupTopicResultOutput {
	return o
}

func (o LookupTopicResultOutput) ToLookupTopicResultOutputWithContext(ctx context.Context) LookupTopicResultOutput {
	return o
}

func (o LookupTopicResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupTopicResult] {
	return pulumix.Output[LookupTopicResult]{
		OutputState: o.OutputState,
	}
}

// Last time the message was sent, or a request was received, for this topic.
func (o LookupTopicResultOutput) AccessedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.AccessedAt }).(pulumi.StringOutput)
}

// ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
func (o LookupTopicResultOutput) AutoDeleteOnIdle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.AutoDeleteOnIdle }).(pulumi.StringPtrOutput)
}

// Message count details
func (o LookupTopicResultOutput) CountDetails() MessageCountDetailsResponseOutput {
	return o.ApplyT(func(v LookupTopicResult) MessageCountDetailsResponse { return v.CountDetails }).(MessageCountDetailsResponseOutput)
}

// Exact time the message was created.
func (o LookupTopicResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
func (o LookupTopicResultOutput) DefaultMessageTimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.DefaultMessageTimeToLive }).(pulumi.StringPtrOutput)
}

// ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
func (o LookupTopicResultOutput) DuplicateDetectionHistoryTimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.DuplicateDetectionHistoryTimeWindow }).(pulumi.StringPtrOutput)
}

// Value that indicates whether server-side batched operations are enabled.
func (o LookupTopicResultOutput) EnableBatchedOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.EnableBatchedOperations }).(pulumi.BoolPtrOutput)
}

// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
func (o LookupTopicResultOutput) EnableExpress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.EnableExpress }).(pulumi.BoolPtrOutput)
}

// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
func (o LookupTopicResultOutput) EnablePartitioning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.EnablePartitioning }).(pulumi.BoolPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupTopicResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Location }).(pulumi.StringOutput)
}

// Maximum size (in KB) of the message payload that can be accepted by the topic. This property is only used in Premium today and default is 1024.
func (o LookupTopicResultOutput) MaxMessageSizeInKilobytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *float64 { return v.MaxMessageSizeInKilobytes }).(pulumi.Float64PtrOutput)
}

// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
func (o LookupTopicResultOutput) MaxSizeInMegabytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *int { return v.MaxSizeInMegabytes }).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o LookupTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

// Value indicating if this topic requires duplicate detection.
func (o LookupTopicResultOutput) RequiresDuplicateDetection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.RequiresDuplicateDetection }).(pulumi.BoolPtrOutput)
}

// Size of the topic, in bytes.
func (o LookupTopicResultOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupTopicResult) float64 { return v.SizeInBytes }).(pulumi.Float64Output)
}

// Enumerates the possible values for the status of a messaging entity.
func (o LookupTopicResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Number of subscriptions.
func (o LookupTopicResultOutput) SubscriptionCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTopicResult) int { return v.SubscriptionCount }).(pulumi.IntOutput)
}

// Value that indicates whether the topic supports ordering.
func (o LookupTopicResultOutput) SupportOrdering() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.SupportOrdering }).(pulumi.BoolPtrOutput)
}

// The system meta data relating to this resource.
func (o LookupTopicResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTopicResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o LookupTopicResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Type }).(pulumi.StringOutput)
}

// The exact time the message was updated.
func (o LookupTopicResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicResultOutput{})
}
