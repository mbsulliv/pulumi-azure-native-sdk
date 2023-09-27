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

// Get all delivery attributes for an event subscription for domain topic.
// Azure REST API version: 2022-06-15.
func GetDomainTopicEventSubscriptionDeliveryAttributes(ctx *pulumi.Context, args *GetDomainTopicEventSubscriptionDeliveryAttributesArgs, opts ...pulumi.InvokeOption) (*GetDomainTopicEventSubscriptionDeliveryAttributesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetDomainTopicEventSubscriptionDeliveryAttributesResult
	err := ctx.Invoke("azure-native:eventgrid:getDomainTopicEventSubscriptionDeliveryAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDomainTopicEventSubscriptionDeliveryAttributesArgs struct {
	// Name of the top level domain.
	DomainName string `pulumi:"domainName"`
	// Name of the event subscription.
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the domain topic.
	TopicName string `pulumi:"topicName"`
}

// Result of the Get delivery attributes operation.
type GetDomainTopicEventSubscriptionDeliveryAttributesResult struct {
	// A collection of DeliveryAttributeMapping
	Value []interface{} `pulumi:"value"`
}

func GetDomainTopicEventSubscriptionDeliveryAttributesOutput(ctx *pulumi.Context, args GetDomainTopicEventSubscriptionDeliveryAttributesOutputArgs, opts ...pulumi.InvokeOption) GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainTopicEventSubscriptionDeliveryAttributesResult, error) {
			args := v.(GetDomainTopicEventSubscriptionDeliveryAttributesArgs)
			r, err := GetDomainTopicEventSubscriptionDeliveryAttributes(ctx, &args, opts...)
			var s GetDomainTopicEventSubscriptionDeliveryAttributesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput)
}

type GetDomainTopicEventSubscriptionDeliveryAttributesOutputArgs struct {
	// Name of the top level domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Name of the event subscription.
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the domain topic.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (GetDomainTopicEventSubscriptionDeliveryAttributesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainTopicEventSubscriptionDeliveryAttributesArgs)(nil)).Elem()
}

// Result of the Get delivery attributes operation.
type GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput struct{ *pulumi.OutputState }

func (GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainTopicEventSubscriptionDeliveryAttributesResult)(nil)).Elem()
}

func (o GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetDomainTopicEventSubscriptionDeliveryAttributesResultOutput() GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetDomainTopicEventSubscriptionDeliveryAttributesResultOutputWithContext(ctx context.Context) GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetDomainTopicEventSubscriptionDeliveryAttributesResult] {
	return pulumix.Output[GetDomainTopicEventSubscriptionDeliveryAttributesResult]{
		OutputState: o.OutputState,
	}
}

// A collection of DeliveryAttributeMapping
func (o GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetDomainTopicEventSubscriptionDeliveryAttributesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainTopicEventSubscriptionDeliveryAttributesResultOutput{})
}
