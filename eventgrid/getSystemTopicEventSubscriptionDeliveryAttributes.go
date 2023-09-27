// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get all delivery attributes for an event subscription.
// Azure REST API version: 2022-06-15.
func GetSystemTopicEventSubscriptionDeliveryAttributes(ctx *pulumi.Context, args *GetSystemTopicEventSubscriptionDeliveryAttributesArgs, opts ...pulumi.InvokeOption) (*GetSystemTopicEventSubscriptionDeliveryAttributesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetSystemTopicEventSubscriptionDeliveryAttributesResult
	err := ctx.Invoke("azure-native:eventgrid:getSystemTopicEventSubscriptionDeliveryAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSystemTopicEventSubscriptionDeliveryAttributesArgs struct {
	// Name of the event subscription to be created. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the system topic.
	SystemTopicName string `pulumi:"systemTopicName"`
}

// Result of the Get delivery attributes operation.
type GetSystemTopicEventSubscriptionDeliveryAttributesResult struct {
	// A collection of DeliveryAttributeMapping
	Value []interface{} `pulumi:"value"`
}

func GetSystemTopicEventSubscriptionDeliveryAttributesOutput(ctx *pulumi.Context, args GetSystemTopicEventSubscriptionDeliveryAttributesOutputArgs, opts ...pulumi.InvokeOption) GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemTopicEventSubscriptionDeliveryAttributesResult, error) {
			args := v.(GetSystemTopicEventSubscriptionDeliveryAttributesArgs)
			r, err := GetSystemTopicEventSubscriptionDeliveryAttributes(ctx, &args, opts...)
			var s GetSystemTopicEventSubscriptionDeliveryAttributesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput)
}

type GetSystemTopicEventSubscriptionDeliveryAttributesOutputArgs struct {
	// Name of the event subscription to be created. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the system topic.
	SystemTopicName pulumi.StringInput `pulumi:"systemTopicName"`
}

func (GetSystemTopicEventSubscriptionDeliveryAttributesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemTopicEventSubscriptionDeliveryAttributesArgs)(nil)).Elem()
}

// Result of the Get delivery attributes operation.
type GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput struct{ *pulumi.OutputState }

func (GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemTopicEventSubscriptionDeliveryAttributesResult)(nil)).Elem()
}

func (o GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetSystemTopicEventSubscriptionDeliveryAttributesResultOutput() GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetSystemTopicEventSubscriptionDeliveryAttributesResultOutputWithContext(ctx context.Context) GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetSystemTopicEventSubscriptionDeliveryAttributesResult] {
	return pulumix.Output[GetSystemTopicEventSubscriptionDeliveryAttributesResult]{
		OutputState: o.OutputState,
	}
}

// A collection of DeliveryAttributeMapping
func (o GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetSystemTopicEventSubscriptionDeliveryAttributesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemTopicEventSubscriptionDeliveryAttributesResultOutput{})
}
