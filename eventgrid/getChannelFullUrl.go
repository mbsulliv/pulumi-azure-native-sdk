// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the full endpoint URL of a partner destination channel.
// Azure REST API version: 2022-06-15.
//
// Other available API versions: 2023-06-01-preview, 2023-12-15-preview.
func GetChannelFullUrl(ctx *pulumi.Context, args *GetChannelFullUrlArgs, opts ...pulumi.InvokeOption) (*GetChannelFullUrlResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetChannelFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid:getChannelFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetChannelFullUrlArgs struct {
	// Name of the Channel.
	ChannelName string `pulumi:"channelName"`
	// Name of the partner namespace.
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	// The name of the resource group within the partners subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Full endpoint url of an event subscription
type GetChannelFullUrlResult struct {
	// The URL that represents the endpoint of the destination of an event subscription.
	EndpointUrl *string `pulumi:"endpointUrl"`
}

func GetChannelFullUrlOutput(ctx *pulumi.Context, args GetChannelFullUrlOutputArgs, opts ...pulumi.InvokeOption) GetChannelFullUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetChannelFullUrlResult, error) {
			args := v.(GetChannelFullUrlArgs)
			r, err := GetChannelFullUrl(ctx, &args, opts...)
			var s GetChannelFullUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetChannelFullUrlResultOutput)
}

type GetChannelFullUrlOutputArgs struct {
	// Name of the Channel.
	ChannelName pulumi.StringInput `pulumi:"channelName"`
	// Name of the partner namespace.
	PartnerNamespaceName pulumi.StringInput `pulumi:"partnerNamespaceName"`
	// The name of the resource group within the partners subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetChannelFullUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetChannelFullUrlArgs)(nil)).Elem()
}

// Full endpoint url of an event subscription
type GetChannelFullUrlResultOutput struct{ *pulumi.OutputState }

func (GetChannelFullUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetChannelFullUrlResult)(nil)).Elem()
}

func (o GetChannelFullUrlResultOutput) ToGetChannelFullUrlResultOutput() GetChannelFullUrlResultOutput {
	return o
}

func (o GetChannelFullUrlResultOutput) ToGetChannelFullUrlResultOutputWithContext(ctx context.Context) GetChannelFullUrlResultOutput {
	return o
}

// The URL that represents the endpoint of the destination of an event subscription.
func (o GetChannelFullUrlResultOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetChannelFullUrlResult) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetChannelFullUrlResultOutput{})
}
