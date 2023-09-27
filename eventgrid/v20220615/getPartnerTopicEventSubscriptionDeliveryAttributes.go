// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get all delivery attributes for an event subscription of a partner topic.
func GetPartnerTopicEventSubscriptionDeliveryAttributes(ctx *pulumi.Context, args *GetPartnerTopicEventSubscriptionDeliveryAttributesArgs, opts ...pulumi.InvokeOption) (*GetPartnerTopicEventSubscriptionDeliveryAttributesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetPartnerTopicEventSubscriptionDeliveryAttributesResult
	err := ctx.Invoke("azure-native:eventgrid/v20220615:getPartnerTopicEventSubscriptionDeliveryAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetPartnerTopicEventSubscriptionDeliveryAttributesArgs struct {
	// Name of the event subscription to be created. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	// Name of the partner topic.
	PartnerTopicName string `pulumi:"partnerTopicName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Result of the Get delivery attributes operation.
type GetPartnerTopicEventSubscriptionDeliveryAttributesResult struct {
	// A collection of DeliveryAttributeMapping
	Value []interface{} `pulumi:"value"`
}

func GetPartnerTopicEventSubscriptionDeliveryAttributesOutput(ctx *pulumi.Context, args GetPartnerTopicEventSubscriptionDeliveryAttributesOutputArgs, opts ...pulumi.InvokeOption) GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPartnerTopicEventSubscriptionDeliveryAttributesResult, error) {
			args := v.(GetPartnerTopicEventSubscriptionDeliveryAttributesArgs)
			r, err := GetPartnerTopicEventSubscriptionDeliveryAttributes(ctx, &args, opts...)
			var s GetPartnerTopicEventSubscriptionDeliveryAttributesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput)
}

type GetPartnerTopicEventSubscriptionDeliveryAttributesOutputArgs struct {
	// Name of the event subscription to be created. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	// Name of the partner topic.
	PartnerTopicName pulumi.StringInput `pulumi:"partnerTopicName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetPartnerTopicEventSubscriptionDeliveryAttributesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPartnerTopicEventSubscriptionDeliveryAttributesArgs)(nil)).Elem()
}

// Result of the Get delivery attributes operation.
type GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput struct{ *pulumi.OutputState }

func (GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPartnerTopicEventSubscriptionDeliveryAttributesResult)(nil)).Elem()
}

func (o GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput() GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetPartnerTopicEventSubscriptionDeliveryAttributesResultOutputWithContext(ctx context.Context) GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetPartnerTopicEventSubscriptionDeliveryAttributesResult] {
	return pulumix.Output[GetPartnerTopicEventSubscriptionDeliveryAttributesResult]{
		OutputState: o.OutputState,
	}
}

// A collection of DeliveryAttributeMapping
func (o GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetPartnerTopicEventSubscriptionDeliveryAttributesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPartnerTopicEventSubscriptionDeliveryAttributesResultOutput{})
}
