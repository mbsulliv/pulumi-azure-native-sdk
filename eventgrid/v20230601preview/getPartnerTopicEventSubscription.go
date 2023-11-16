// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of an event subscription of a partner topic.
func LookupPartnerTopicEventSubscription(ctx *pulumi.Context, args *LookupPartnerTopicEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupPartnerTopicEventSubscriptionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPartnerTopicEventSubscriptionResult
	err := ctx.Invoke("azure-native:eventgrid/v20230601preview:getPartnerTopicEventSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPartnerTopicEventSubscriptionArgs struct {
	// Name of the event subscription to be found. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	// Name of the partner topic.
	PartnerTopicName string `pulumi:"partnerTopicName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Event Subscription.
type LookupPartnerTopicEventSubscriptionResult struct {
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterDestination *StorageBlobDeadLetterDestinationResponse `pulumi:"deadLetterDestination"`
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterWithResourceIdentity *DeadLetterWithResourceIdentityResponse `pulumi:"deadLetterWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeliveryWithResourceIdentity *DeliveryWithResourceIdentityResponse `pulumi:"deliveryWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	Destination interface{} `pulumi:"destination"`
	// The event delivery schema for the event subscription.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// Expiration time of the event subscription.
	ExpirationTimeUtc *string `pulumi:"expirationTimeUtc"`
	// Information about the filter for the event subscription.
	Filter *EventSubscriptionFilterResponse `pulumi:"filter"`
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// List of user defined labels.
	Labels []string `pulumi:"labels"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// Provisioning state of the event subscription.
	ProvisioningState string `pulumi:"provisioningState"`
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy *RetryPolicyResponse `pulumi:"retryPolicy"`
	// The system metadata relating to Event Subscription resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Name of the topic of the event subscription.
	Topic string `pulumi:"topic"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupPartnerTopicEventSubscriptionResult
func (val *LookupPartnerTopicEventSubscriptionResult) Defaults() *LookupPartnerTopicEventSubscriptionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EventDeliverySchema == nil {
		eventDeliverySchema_ := "EventGridSchema"
		tmp.EventDeliverySchema = &eventDeliverySchema_
	}
	tmp.Filter = tmp.Filter.Defaults()

	tmp.RetryPolicy = tmp.RetryPolicy.Defaults()

	return &tmp
}

func LookupPartnerTopicEventSubscriptionOutput(ctx *pulumi.Context, args LookupPartnerTopicEventSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerTopicEventSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerTopicEventSubscriptionResult, error) {
			args := v.(LookupPartnerTopicEventSubscriptionArgs)
			r, err := LookupPartnerTopicEventSubscription(ctx, &args, opts...)
			var s LookupPartnerTopicEventSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerTopicEventSubscriptionResultOutput)
}

type LookupPartnerTopicEventSubscriptionOutputArgs struct {
	// Name of the event subscription to be found. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	// Name of the partner topic.
	PartnerTopicName pulumi.StringInput `pulumi:"partnerTopicName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPartnerTopicEventSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerTopicEventSubscriptionArgs)(nil)).Elem()
}

// Event Subscription.
type LookupPartnerTopicEventSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerTopicEventSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerTopicEventSubscriptionResult)(nil)).Elem()
}

func (o LookupPartnerTopicEventSubscriptionResultOutput) ToLookupPartnerTopicEventSubscriptionResultOutput() LookupPartnerTopicEventSubscriptionResultOutput {
	return o
}

func (o LookupPartnerTopicEventSubscriptionResultOutput) ToLookupPartnerTopicEventSubscriptionResultOutputWithContext(ctx context.Context) LookupPartnerTopicEventSubscriptionResultOutput {
	return o
}

// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
func (o LookupPartnerTopicEventSubscriptionResultOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) *StorageBlobDeadLetterDestinationResponse {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
func (o LookupPartnerTopicEventSubscriptionResultOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) *DeadLetterWithResourceIdentityResponse {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

// Information about the destination where events have to be delivered for the event subscription.
// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
func (o LookupPartnerTopicEventSubscriptionResultOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) *DeliveryWithResourceIdentityResponse {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

// Information about the destination where events have to be delivered for the event subscription.
// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
func (o LookupPartnerTopicEventSubscriptionResultOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) interface{} { return v.Destination }).(pulumi.AnyOutput)
}

// The event delivery schema for the event subscription.
func (o LookupPartnerTopicEventSubscriptionResultOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) *string { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

// Expiration time of the event subscription.
func (o LookupPartnerTopicEventSubscriptionResultOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) *string { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

// Information about the filter for the event subscription.
func (o LookupPartnerTopicEventSubscriptionResultOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) *EventSubscriptionFilterResponse { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

// Fully qualified identifier of the resource.
func (o LookupPartnerTopicEventSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of user defined labels.
func (o LookupPartnerTopicEventSubscriptionResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

// Name of the resource.
func (o LookupPartnerTopicEventSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the event subscription.
func (o LookupPartnerTopicEventSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
func (o LookupPartnerTopicEventSubscriptionResultOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) *RetryPolicyResponse { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

// The system metadata relating to Event Subscription resource.
func (o LookupPartnerTopicEventSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Name of the topic of the event subscription.
func (o LookupPartnerTopicEventSubscriptionResultOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) string { return v.Topic }).(pulumi.StringOutput)
}

// Type of the resource.
func (o LookupPartnerTopicEventSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicEventSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerTopicEventSubscriptionResultOutput{})
}
