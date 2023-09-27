// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package webpubsub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a hub setting.
// Azure REST API version: 2023-02-01.
func LookupWebPubSubHub(ctx *pulumi.Context, args *LookupWebPubSubHubArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubHubResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebPubSubHubResult
	err := ctx.Invoke("azure-native:webpubsub:getWebPubSubHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebPubSubHubArgs struct {
	// The hub name.
	HubName string `pulumi:"hubName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// A hub setting
type LookupWebPubSubHubResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Properties of a hub.
	Properties WebPubSubHubPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupWebPubSubHubResult
func (val *LookupWebPubSubHubResult) Defaults() *LookupWebPubSubHubResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupWebPubSubHubOutput(ctx *pulumi.Context, args LookupWebPubSubHubOutputArgs, opts ...pulumi.InvokeOption) LookupWebPubSubHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebPubSubHubResult, error) {
			args := v.(LookupWebPubSubHubArgs)
			r, err := LookupWebPubSubHub(ctx, &args, opts...)
			var s LookupWebPubSubHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebPubSubHubResultOutput)
}

type LookupWebPubSubHubOutputArgs struct {
	// The hub name.
	HubName pulumi.StringInput `pulumi:"hubName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWebPubSubHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubHubArgs)(nil)).Elem()
}

// A hub setting
type LookupWebPubSubHubResultOutput struct{ *pulumi.OutputState }

func (LookupWebPubSubHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubHubResult)(nil)).Elem()
}

func (o LookupWebPubSubHubResultOutput) ToLookupWebPubSubHubResultOutput() LookupWebPubSubHubResultOutput {
	return o
}

func (o LookupWebPubSubHubResultOutput) ToLookupWebPubSubHubResultOutputWithContext(ctx context.Context) LookupWebPubSubHubResultOutput {
	return o
}

func (o LookupWebPubSubHubResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWebPubSubHubResult] {
	return pulumix.Output[LookupWebPubSubHubResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource Id for the resource.
func (o LookupWebPubSubHubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubHubResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupWebPubSubHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubHubResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of a hub.
func (o LookupWebPubSubHubResultOutput) Properties() WebPubSubHubPropertiesResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubHubResult) WebPubSubHubPropertiesResponse { return v.Properties }).(WebPubSubHubPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupWebPubSubHubResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubHubResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
func (o LookupWebPubSubHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubHubResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebPubSubHubResultOutput{})
}
