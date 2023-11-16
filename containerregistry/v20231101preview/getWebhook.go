// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of the specified webhook.
func LookupWebhook(ctx *pulumi.Context, args *LookupWebhookArgs, opts ...pulumi.InvokeOption) (*LookupWebhookResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebhookResult
	err := ctx.Invoke("azure-native:containerregistry/v20231101preview:getWebhook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebhookArgs struct {
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the webhook.
	WebhookName string `pulumi:"webhookName"`
}

// An object that represents a webhook for a container registry.
type LookupWebhookResult struct {
	// The list of actions that trigger the webhook to post notifications.
	Actions []string `pulumi:"actions"`
	// The resource ID.
	Id string `pulumi:"id"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The provisioning state of the webhook at the time the operation was called.
	ProvisioningState string `pulumi:"provisioningState"`
	// The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
	Scope *string `pulumi:"scope"`
	// The status of the webhook at the time the operation was called.
	Status *string `pulumi:"status"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupWebhookOutput(ctx *pulumi.Context, args LookupWebhookOutputArgs, opts ...pulumi.InvokeOption) LookupWebhookResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebhookResult, error) {
			args := v.(LookupWebhookArgs)
			r, err := LookupWebhook(ctx, &args, opts...)
			var s LookupWebhookResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebhookResultOutput)
}

type LookupWebhookOutputArgs struct {
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the webhook.
	WebhookName pulumi.StringInput `pulumi:"webhookName"`
}

func (LookupWebhookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebhookArgs)(nil)).Elem()
}

// An object that represents a webhook for a container registry.
type LookupWebhookResultOutput struct{ *pulumi.OutputState }

func (LookupWebhookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebhookResult)(nil)).Elem()
}

func (o LookupWebhookResultOutput) ToLookupWebhookResultOutput() LookupWebhookResultOutput {
	return o
}

func (o LookupWebhookResultOutput) ToLookupWebhookResultOutputWithContext(ctx context.Context) LookupWebhookResultOutput {
	return o
}

// The list of actions that trigger the webhook to post notifications.
func (o LookupWebhookResultOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebhookResult) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

// The resource ID.
func (o LookupWebhookResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource. This cannot be changed after the resource is created.
func (o LookupWebhookResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupWebhookResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the webhook at the time the operation was called.
func (o LookupWebhookResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
func (o LookupWebhookResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// The status of the webhook at the time the operation was called.
func (o LookupWebhookResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupWebhookResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebhookResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tags of the resource.
func (o LookupWebhookResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebhookResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupWebhookResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebhookResultOutput{})
}
