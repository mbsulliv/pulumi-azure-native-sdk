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

// Get the full endpoint URL for an event subscription for domain.
// Azure REST API version: 2022-06-15.
func GetDomainEventSubscriptionFullUrl(ctx *pulumi.Context, args *GetDomainEventSubscriptionFullUrlArgs, opts ...pulumi.InvokeOption) (*GetDomainEventSubscriptionFullUrlResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetDomainEventSubscriptionFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid:getDomainEventSubscriptionFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDomainEventSubscriptionFullUrlArgs struct {
	// Name of the domain topic.
	DomainName string `pulumi:"domainName"`
	// Name of the event subscription.
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Full endpoint url of an event subscription
type GetDomainEventSubscriptionFullUrlResult struct {
	// The URL that represents the endpoint of the destination of an event subscription.
	EndpointUrl *string `pulumi:"endpointUrl"`
}

func GetDomainEventSubscriptionFullUrlOutput(ctx *pulumi.Context, args GetDomainEventSubscriptionFullUrlOutputArgs, opts ...pulumi.InvokeOption) GetDomainEventSubscriptionFullUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainEventSubscriptionFullUrlResult, error) {
			args := v.(GetDomainEventSubscriptionFullUrlArgs)
			r, err := GetDomainEventSubscriptionFullUrl(ctx, &args, opts...)
			var s GetDomainEventSubscriptionFullUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainEventSubscriptionFullUrlResultOutput)
}

type GetDomainEventSubscriptionFullUrlOutputArgs struct {
	// Name of the domain topic.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Name of the event subscription.
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetDomainEventSubscriptionFullUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainEventSubscriptionFullUrlArgs)(nil)).Elem()
}

// Full endpoint url of an event subscription
type GetDomainEventSubscriptionFullUrlResultOutput struct{ *pulumi.OutputState }

func (GetDomainEventSubscriptionFullUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainEventSubscriptionFullUrlResult)(nil)).Elem()
}

func (o GetDomainEventSubscriptionFullUrlResultOutput) ToGetDomainEventSubscriptionFullUrlResultOutput() GetDomainEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetDomainEventSubscriptionFullUrlResultOutput) ToGetDomainEventSubscriptionFullUrlResultOutputWithContext(ctx context.Context) GetDomainEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetDomainEventSubscriptionFullUrlResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetDomainEventSubscriptionFullUrlResult] {
	return pulumix.Output[GetDomainEventSubscriptionFullUrlResult]{
		OutputState: o.OutputState,
	}
}

// The URL that represents the endpoint of the destination of an event subscription.
func (o GetDomainEventSubscriptionFullUrlResultOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainEventSubscriptionFullUrlResult) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainEventSubscriptionFullUrlResultOutput{})
}
