// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231215preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get all delivery attributes for an event subscription.
func GetEventSubscriptionDeliveryAttributes(ctx *pulumi.Context, args *GetEventSubscriptionDeliveryAttributesArgs, opts ...pulumi.InvokeOption) (*GetEventSubscriptionDeliveryAttributesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetEventSubscriptionDeliveryAttributesResult
	err := ctx.Invoke("azure-native:eventgrid/v20231215preview:getEventSubscriptionDeliveryAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEventSubscriptionDeliveryAttributesArgs struct {
	// Name of the event subscription.
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	// The scope of the event subscription. The scope can be a subscription, or a resource group, or a top level resource belonging to a resource provider namespace, or an EventGrid topic. For example, use '/subscriptions/{subscriptionId}/' for a subscription, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for a resource group, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}' for a resource, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}' for an EventGrid topic.
	Scope string `pulumi:"scope"`
}

// Result of the Get delivery attributes operation.
type GetEventSubscriptionDeliveryAttributesResult struct {
	// A collection of DeliveryAttributeMapping
	Value []interface{} `pulumi:"value"`
}

func GetEventSubscriptionDeliveryAttributesOutput(ctx *pulumi.Context, args GetEventSubscriptionDeliveryAttributesOutputArgs, opts ...pulumi.InvokeOption) GetEventSubscriptionDeliveryAttributesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEventSubscriptionDeliveryAttributesResult, error) {
			args := v.(GetEventSubscriptionDeliveryAttributesArgs)
			r, err := GetEventSubscriptionDeliveryAttributes(ctx, &args, opts...)
			var s GetEventSubscriptionDeliveryAttributesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEventSubscriptionDeliveryAttributesResultOutput)
}

type GetEventSubscriptionDeliveryAttributesOutputArgs struct {
	// Name of the event subscription.
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	// The scope of the event subscription. The scope can be a subscription, or a resource group, or a top level resource belonging to a resource provider namespace, or an EventGrid topic. For example, use '/subscriptions/{subscriptionId}/' for a subscription, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for a resource group, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}' for a resource, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}' for an EventGrid topic.
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (GetEventSubscriptionDeliveryAttributesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEventSubscriptionDeliveryAttributesArgs)(nil)).Elem()
}

// Result of the Get delivery attributes operation.
type GetEventSubscriptionDeliveryAttributesResultOutput struct{ *pulumi.OutputState }

func (GetEventSubscriptionDeliveryAttributesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEventSubscriptionDeliveryAttributesResult)(nil)).Elem()
}

func (o GetEventSubscriptionDeliveryAttributesResultOutput) ToGetEventSubscriptionDeliveryAttributesResultOutput() GetEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetEventSubscriptionDeliveryAttributesResultOutput) ToGetEventSubscriptionDeliveryAttributesResultOutputWithContext(ctx context.Context) GetEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

// A collection of DeliveryAttributeMapping
func (o GetEventSubscriptionDeliveryAttributesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetEventSubscriptionDeliveryAttributesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEventSubscriptionDeliveryAttributesResultOutput{})
}