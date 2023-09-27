// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220430preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a consumer group from the Event Hub-compatible device-to-cloud endpoint for an IoT hub.
func LookupIotHubResourceEventHubConsumerGroup(ctx *pulumi.Context, args *LookupIotHubResourceEventHubConsumerGroupArgs, opts ...pulumi.InvokeOption) (*LookupIotHubResourceEventHubConsumerGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIotHubResourceEventHubConsumerGroupResult
	err := ctx.Invoke("azure-native:devices/v20220430preview:getIotHubResourceEventHubConsumerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotHubResourceEventHubConsumerGroupArgs struct {
	// The name of the Event Hub-compatible endpoint in the IoT hub.
	EventHubEndpointName string `pulumi:"eventHubEndpointName"`
	// The name of the consumer group to retrieve.
	Name string `pulumi:"name"`
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName string `pulumi:"resourceName"`
}

// The properties of the EventHubConsumerGroupInfo object.
type LookupIotHubResourceEventHubConsumerGroupResult struct {
	// The etag.
	Etag string `pulumi:"etag"`
	// The Event Hub-compatible consumer group identifier.
	Id string `pulumi:"id"`
	// The Event Hub-compatible consumer group name.
	Name string `pulumi:"name"`
	// The tags.
	Properties interface{} `pulumi:"properties"`
	// the resource type.
	Type string `pulumi:"type"`
}

func LookupIotHubResourceEventHubConsumerGroupOutput(ctx *pulumi.Context, args LookupIotHubResourceEventHubConsumerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupIotHubResourceEventHubConsumerGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotHubResourceEventHubConsumerGroupResult, error) {
			args := v.(LookupIotHubResourceEventHubConsumerGroupArgs)
			r, err := LookupIotHubResourceEventHubConsumerGroup(ctx, &args, opts...)
			var s LookupIotHubResourceEventHubConsumerGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotHubResourceEventHubConsumerGroupResultOutput)
}

type LookupIotHubResourceEventHubConsumerGroupOutputArgs struct {
	// The name of the Event Hub-compatible endpoint in the IoT hub.
	EventHubEndpointName pulumi.StringInput `pulumi:"eventHubEndpointName"`
	// The name of the consumer group to retrieve.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupIotHubResourceEventHubConsumerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubResourceEventHubConsumerGroupArgs)(nil)).Elem()
}

// The properties of the EventHubConsumerGroupInfo object.
type LookupIotHubResourceEventHubConsumerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupIotHubResourceEventHubConsumerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubResourceEventHubConsumerGroupResult)(nil)).Elem()
}

func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) ToLookupIotHubResourceEventHubConsumerGroupResultOutput() LookupIotHubResourceEventHubConsumerGroupResultOutput {
	return o
}

func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) ToLookupIotHubResourceEventHubConsumerGroupResultOutputWithContext(ctx context.Context) LookupIotHubResourceEventHubConsumerGroupResultOutput {
	return o
}

func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIotHubResourceEventHubConsumerGroupResult] {
	return pulumix.Output[LookupIotHubResourceEventHubConsumerGroupResult]{
		OutputState: o.OutputState,
	}
}

// The etag.
func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceEventHubConsumerGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The Event Hub-compatible consumer group identifier.
func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceEventHubConsumerGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Event Hub-compatible consumer group name.
func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceEventHubConsumerGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The tags.
func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIotHubResourceEventHubConsumerGroupResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// the resource type.
func (o LookupIotHubResourceEventHubConsumerGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceEventHubConsumerGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotHubResourceEventHubConsumerGroupResultOutput{})
}
