// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of an event subscription of a domain.
func LookupDomainEventSubscription(ctx *pulumi.Context, args *LookupDomainEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupDomainEventSubscriptionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainEventSubscriptionResult
	err := ctx.Invoke("azure-native:eventgrid/v20230601preview:getDomainEventSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDomainEventSubscriptionArgs struct {
	// Name of the domain.
	DomainName string `pulumi:"domainName"`
	// Name of the event subscription to be found. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Event Subscription.
type LookupDomainEventSubscriptionResult struct {
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

// Defaults sets the appropriate defaults for LookupDomainEventSubscriptionResult
func (val *LookupDomainEventSubscriptionResult) Defaults() *LookupDomainEventSubscriptionResult {
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

func LookupDomainEventSubscriptionOutput(ctx *pulumi.Context, args LookupDomainEventSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupDomainEventSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainEventSubscriptionResult, error) {
			args := v.(LookupDomainEventSubscriptionArgs)
			r, err := LookupDomainEventSubscription(ctx, &args, opts...)
			var s LookupDomainEventSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainEventSubscriptionResultOutput)
}

type LookupDomainEventSubscriptionOutputArgs struct {
	// Name of the domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Name of the event subscription to be found. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainEventSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainEventSubscriptionArgs)(nil)).Elem()
}

// Event Subscription.
type LookupDomainEventSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupDomainEventSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainEventSubscriptionResult)(nil)).Elem()
}

func (o LookupDomainEventSubscriptionResultOutput) ToLookupDomainEventSubscriptionResultOutput() LookupDomainEventSubscriptionResultOutput {
	return o
}

func (o LookupDomainEventSubscriptionResultOutput) ToLookupDomainEventSubscriptionResultOutputWithContext(ctx context.Context) LookupDomainEventSubscriptionResultOutput {
	return o
}

// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
func (o LookupDomainEventSubscriptionResultOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *StorageBlobDeadLetterDestinationResponse {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
func (o LookupDomainEventSubscriptionResultOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *DeadLetterWithResourceIdentityResponse {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

// Information about the destination where events have to be delivered for the event subscription.
// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
func (o LookupDomainEventSubscriptionResultOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *DeliveryWithResourceIdentityResponse {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

// Information about the destination where events have to be delivered for the event subscription.
// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
func (o LookupDomainEventSubscriptionResultOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) interface{} { return v.Destination }).(pulumi.AnyOutput)
}

// The event delivery schema for the event subscription.
func (o LookupDomainEventSubscriptionResultOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *string { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

// Expiration time of the event subscription.
func (o LookupDomainEventSubscriptionResultOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *string { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

// Information about the filter for the event subscription.
func (o LookupDomainEventSubscriptionResultOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *EventSubscriptionFilterResponse { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

// Fully qualified identifier of the resource.
func (o LookupDomainEventSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of user defined labels.
func (o LookupDomainEventSubscriptionResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

// Name of the resource.
func (o LookupDomainEventSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the event subscription.
func (o LookupDomainEventSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
func (o LookupDomainEventSubscriptionResultOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *RetryPolicyResponse { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

// The system metadata relating to Event Subscription resource.
func (o LookupDomainEventSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Name of the topic of the event subscription.
func (o LookupDomainEventSubscriptionResultOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) string { return v.Topic }).(pulumi.StringOutput)
}

// Type of the resource.
func (o LookupDomainEventSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainEventSubscriptionResultOutput{})
}