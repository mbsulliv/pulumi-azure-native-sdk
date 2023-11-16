// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the full endpoint URL for an event subscription of a system topic.
// Azure REST API version: 2022-06-15.
//
// Other available API versions: 2023-06-01-preview, 2023-12-15-preview.
func GetSystemTopicEventSubscriptionFullUrl(ctx *pulumi.Context, args *GetSystemTopicEventSubscriptionFullUrlArgs, opts ...pulumi.InvokeOption) (*GetSystemTopicEventSubscriptionFullUrlResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetSystemTopicEventSubscriptionFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid:getSystemTopicEventSubscriptionFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSystemTopicEventSubscriptionFullUrlArgs struct {
	// Name of the event subscription to be created. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the system topic.
	SystemTopicName string `pulumi:"systemTopicName"`
}

// Full endpoint url of an event subscription
type GetSystemTopicEventSubscriptionFullUrlResult struct {
	// The URL that represents the endpoint of the destination of an event subscription.
	EndpointUrl *string `pulumi:"endpointUrl"`
}

func GetSystemTopicEventSubscriptionFullUrlOutput(ctx *pulumi.Context, args GetSystemTopicEventSubscriptionFullUrlOutputArgs, opts ...pulumi.InvokeOption) GetSystemTopicEventSubscriptionFullUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemTopicEventSubscriptionFullUrlResult, error) {
			args := v.(GetSystemTopicEventSubscriptionFullUrlArgs)
			r, err := GetSystemTopicEventSubscriptionFullUrl(ctx, &args, opts...)
			var s GetSystemTopicEventSubscriptionFullUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemTopicEventSubscriptionFullUrlResultOutput)
}

type GetSystemTopicEventSubscriptionFullUrlOutputArgs struct {
	// Name of the event subscription to be created. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the system topic.
	SystemTopicName pulumi.StringInput `pulumi:"systemTopicName"`
}

func (GetSystemTopicEventSubscriptionFullUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemTopicEventSubscriptionFullUrlArgs)(nil)).Elem()
}

// Full endpoint url of an event subscription
type GetSystemTopicEventSubscriptionFullUrlResultOutput struct{ *pulumi.OutputState }

func (GetSystemTopicEventSubscriptionFullUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemTopicEventSubscriptionFullUrlResult)(nil)).Elem()
}

func (o GetSystemTopicEventSubscriptionFullUrlResultOutput) ToGetSystemTopicEventSubscriptionFullUrlResultOutput() GetSystemTopicEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetSystemTopicEventSubscriptionFullUrlResultOutput) ToGetSystemTopicEventSubscriptionFullUrlResultOutputWithContext(ctx context.Context) GetSystemTopicEventSubscriptionFullUrlResultOutput {
	return o
}

// The URL that represents the endpoint of the destination of an event subscription.
func (o GetSystemTopicEventSubscriptionFullUrlResultOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemTopicEventSubscriptionFullUrlResult) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemTopicEventSubscriptionFullUrlResultOutput{})
}
