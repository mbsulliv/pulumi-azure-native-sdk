// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of an event subscription of a namespace topic.
// Azure REST API version: 2023-06-01-preview.
func LookupNamespaceTopicEventSubscription(ctx *pulumi.Context, args *LookupNamespaceTopicEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceTopicEventSubscriptionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNamespaceTopicEventSubscriptionResult
	err := ctx.Invoke("azure-native:eventgrid:getNamespaceTopicEventSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceTopicEventSubscriptionArgs struct {
	// Name of the event subscription to be created. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	// Name of the namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the namespace topic.
	TopicName string `pulumi:"topicName"`
}

// Event Subscription.
type LookupNamespaceTopicEventSubscriptionResult struct {
	// Information about the delivery configuration of the event subscription.
	DeliveryConfiguration *DeliveryConfigurationResponse `pulumi:"deliveryConfiguration"`
	// The event delivery schema for the event subscription.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// Information about the filter for the event subscription.
	FiltersConfiguration *FiltersConfigurationResponse `pulumi:"filtersConfiguration"`
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// Provisioning state of the event subscription.
	ProvisioningState string `pulumi:"provisioningState"`
	// The system metadata relating to Event Subscription resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

func LookupNamespaceTopicEventSubscriptionOutput(ctx *pulumi.Context, args LookupNamespaceTopicEventSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceTopicEventSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceTopicEventSubscriptionResult, error) {
			args := v.(LookupNamespaceTopicEventSubscriptionArgs)
			r, err := LookupNamespaceTopicEventSubscription(ctx, &args, opts...)
			var s LookupNamespaceTopicEventSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceTopicEventSubscriptionResultOutput)
}

type LookupNamespaceTopicEventSubscriptionOutputArgs struct {
	// Name of the event subscription to be created. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	// Name of the namespace.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the namespace topic.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (LookupNamespaceTopicEventSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceTopicEventSubscriptionArgs)(nil)).Elem()
}

// Event Subscription.
type LookupNamespaceTopicEventSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceTopicEventSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceTopicEventSubscriptionResult)(nil)).Elem()
}

func (o LookupNamespaceTopicEventSubscriptionResultOutput) ToLookupNamespaceTopicEventSubscriptionResultOutput() LookupNamespaceTopicEventSubscriptionResultOutput {
	return o
}

func (o LookupNamespaceTopicEventSubscriptionResultOutput) ToLookupNamespaceTopicEventSubscriptionResultOutputWithContext(ctx context.Context) LookupNamespaceTopicEventSubscriptionResultOutput {
	return o
}

// Information about the delivery configuration of the event subscription.
func (o LookupNamespaceTopicEventSubscriptionResultOutput) DeliveryConfiguration() DeliveryConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceTopicEventSubscriptionResult) *DeliveryConfigurationResponse {
		return v.DeliveryConfiguration
	}).(DeliveryConfigurationResponsePtrOutput)
}

// The event delivery schema for the event subscription.
func (o LookupNamespaceTopicEventSubscriptionResultOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceTopicEventSubscriptionResult) *string { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

// Information about the filter for the event subscription.
func (o LookupNamespaceTopicEventSubscriptionResultOutput) FiltersConfiguration() FiltersConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceTopicEventSubscriptionResult) *FiltersConfigurationResponse {
		return v.FiltersConfiguration
	}).(FiltersConfigurationResponsePtrOutput)
}

// Fully qualified identifier of the resource.
func (o LookupNamespaceTopicEventSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceTopicEventSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the resource.
func (o LookupNamespaceTopicEventSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceTopicEventSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the event subscription.
func (o LookupNamespaceTopicEventSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceTopicEventSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata relating to Event Subscription resource.
func (o LookupNamespaceTopicEventSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNamespaceTopicEventSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the resource.
func (o LookupNamespaceTopicEventSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceTopicEventSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceTopicEventSubscriptionResultOutput{})
}