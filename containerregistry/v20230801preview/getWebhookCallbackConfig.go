// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the configuration of service URI and custom headers for the webhook.
func GetWebhookCallbackConfig(ctx *pulumi.Context, args *GetWebhookCallbackConfigArgs, opts ...pulumi.InvokeOption) (*GetWebhookCallbackConfigResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetWebhookCallbackConfigResult
	err := ctx.Invoke("azure-native:containerregistry/v20230801preview:getWebhookCallbackConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetWebhookCallbackConfigArgs struct {
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the webhook.
	WebhookName string `pulumi:"webhookName"`
}

// The configuration of service URI and custom headers for the webhook.
type GetWebhookCallbackConfigResult struct {
	// Custom headers that will be added to the webhook notifications.
	CustomHeaders map[string]string `pulumi:"customHeaders"`
	// The service URI for the webhook to post notifications.
	ServiceUri string `pulumi:"serviceUri"`
}

func GetWebhookCallbackConfigOutput(ctx *pulumi.Context, args GetWebhookCallbackConfigOutputArgs, opts ...pulumi.InvokeOption) GetWebhookCallbackConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWebhookCallbackConfigResult, error) {
			args := v.(GetWebhookCallbackConfigArgs)
			r, err := GetWebhookCallbackConfig(ctx, &args, opts...)
			var s GetWebhookCallbackConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWebhookCallbackConfigResultOutput)
}

type GetWebhookCallbackConfigOutputArgs struct {
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the webhook.
	WebhookName pulumi.StringInput `pulumi:"webhookName"`
}

func (GetWebhookCallbackConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebhookCallbackConfigArgs)(nil)).Elem()
}

// The configuration of service URI and custom headers for the webhook.
type GetWebhookCallbackConfigResultOutput struct{ *pulumi.OutputState }

func (GetWebhookCallbackConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebhookCallbackConfigResult)(nil)).Elem()
}

func (o GetWebhookCallbackConfigResultOutput) ToGetWebhookCallbackConfigResultOutput() GetWebhookCallbackConfigResultOutput {
	return o
}

func (o GetWebhookCallbackConfigResultOutput) ToGetWebhookCallbackConfigResultOutputWithContext(ctx context.Context) GetWebhookCallbackConfigResultOutput {
	return o
}

func (o GetWebhookCallbackConfigResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetWebhookCallbackConfigResult] {
	return pulumix.Output[GetWebhookCallbackConfigResult]{
		OutputState: o.OutputState,
	}
}

// Custom headers that will be added to the webhook notifications.
func (o GetWebhookCallbackConfigResultOutput) CustomHeaders() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetWebhookCallbackConfigResult) map[string]string { return v.CustomHeaders }).(pulumi.StringMapOutput)
}

// The service URI for the webhook to post notifications.
func (o GetWebhookCallbackConfigResultOutput) ServiceUri() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebhookCallbackConfigResult) string { return v.ServiceUri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWebhookCallbackConfigResultOutput{})
}
