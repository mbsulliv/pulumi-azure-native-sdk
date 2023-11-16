// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the full endpoint URL for a nested event subscription for domain topic.
func GetDomainTopicEventSubscriptionFullUrl(ctx *pulumi.Context, args *GetDomainTopicEventSubscriptionFullUrlArgs, opts ...pulumi.InvokeOption) (*GetDomainTopicEventSubscriptionFullUrlResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetDomainTopicEventSubscriptionFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid/v20230601preview:getDomainTopicEventSubscriptionFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDomainTopicEventSubscriptionFullUrlArgs struct {
	// Name of the top level domain.
	DomainName string `pulumi:"domainName"`
	// Name of the event subscription.
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the domain topic.
	TopicName string `pulumi:"topicName"`
}

// Full endpoint url of an event subscription
type GetDomainTopicEventSubscriptionFullUrlResult struct {
	// The URL that represents the endpoint of the destination of an event subscription.
	EndpointUrl *string `pulumi:"endpointUrl"`
}

func GetDomainTopicEventSubscriptionFullUrlOutput(ctx *pulumi.Context, args GetDomainTopicEventSubscriptionFullUrlOutputArgs, opts ...pulumi.InvokeOption) GetDomainTopicEventSubscriptionFullUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainTopicEventSubscriptionFullUrlResult, error) {
			args := v.(GetDomainTopicEventSubscriptionFullUrlArgs)
			r, err := GetDomainTopicEventSubscriptionFullUrl(ctx, &args, opts...)
			var s GetDomainTopicEventSubscriptionFullUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainTopicEventSubscriptionFullUrlResultOutput)
}

type GetDomainTopicEventSubscriptionFullUrlOutputArgs struct {
	// Name of the top level domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Name of the event subscription.
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the domain topic.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (GetDomainTopicEventSubscriptionFullUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainTopicEventSubscriptionFullUrlArgs)(nil)).Elem()
}

// Full endpoint url of an event subscription
type GetDomainTopicEventSubscriptionFullUrlResultOutput struct{ *pulumi.OutputState }

func (GetDomainTopicEventSubscriptionFullUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainTopicEventSubscriptionFullUrlResult)(nil)).Elem()
}

func (o GetDomainTopicEventSubscriptionFullUrlResultOutput) ToGetDomainTopicEventSubscriptionFullUrlResultOutput() GetDomainTopicEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetDomainTopicEventSubscriptionFullUrlResultOutput) ToGetDomainTopicEventSubscriptionFullUrlResultOutputWithContext(ctx context.Context) GetDomainTopicEventSubscriptionFullUrlResultOutput {
	return o
}

// The URL that represents the endpoint of the destination of an event subscription.
func (o GetDomainTopicEventSubscriptionFullUrlResultOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainTopicEventSubscriptionFullUrlResult) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainTopicEventSubscriptionFullUrlResultOutput{})
}
