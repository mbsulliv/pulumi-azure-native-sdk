// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package timeseriesinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the event source with the specified name in the specified environment.
// Azure REST API version: 2020-05-15.
func LookupEventHubEventSource(ctx *pulumi.Context, args *LookupEventHubEventSourceArgs, opts ...pulumi.InvokeOption) (*LookupEventHubEventSourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEventHubEventSourceResult
	err := ctx.Invoke("azure-native:timeseriesinsights:getEventHubEventSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubEventSourceArgs struct {
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName string `pulumi:"environmentName"`
	// The name of the Time Series Insights event source associated with the specified environment.
	EventSourceName string `pulumi:"eventSourceName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An event source that receives its data from an Azure EventHub.
type LookupEventHubEventSourceResult struct {
	// The name of the event hub's consumer group that holds the partitions from which events will be read.
	ConsumerGroupName string `pulumi:"consumerGroupName"`
	// The time the resource was created.
	CreationTime string `pulumi:"creationTime"`
	// The name of the event hub.
	EventHubName string `pulumi:"eventHubName"`
	// The resource id of the event source in Azure Resource Manager.
	EventSourceResourceId string `pulumi:"eventSourceResourceId"`
	// Resource Id
	Id string `pulumi:"id"`
	// The name of the SAS key that grants the Time Series Insights service access to the event hub. The shared access policies for this key must grant 'Listen' permissions to the event hub.
	KeyName string `pulumi:"keyName"`
	// The kind of the event source.
	// Expected value is 'Microsoft.EventHub'.
	Kind string `pulumi:"kind"`
	// An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
	LocalTimestamp *LocalTimestampResponse `pulumi:"localTimestamp"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The name of the service bus that contains the event hub.
	ServiceBusNamespace string `pulumi:"serviceBusNamespace"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// ISO8601 UTC datetime with seconds precision (milliseconds are optional), specifying the date and time that will be the starting point for Events to be consumed.
	Time *string `pulumi:"time"`
	// The event property that will be used as the event source's timestamp. If a value isn't specified for timestampPropertyName, or if null or empty-string is specified, the event creation time will be used.
	TimestampPropertyName *string `pulumi:"timestampPropertyName"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupEventHubEventSourceOutput(ctx *pulumi.Context, args LookupEventHubEventSourceOutputArgs, opts ...pulumi.InvokeOption) LookupEventHubEventSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventHubEventSourceResult, error) {
			args := v.(LookupEventHubEventSourceArgs)
			r, err := LookupEventHubEventSource(ctx, &args, opts...)
			var s LookupEventHubEventSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventHubEventSourceResultOutput)
}

type LookupEventHubEventSourceOutputArgs struct {
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	// The name of the Time Series Insights event source associated with the specified environment.
	EventSourceName pulumi.StringInput `pulumi:"eventSourceName"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventHubEventSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubEventSourceArgs)(nil)).Elem()
}

// An event source that receives its data from an Azure EventHub.
type LookupEventHubEventSourceResultOutput struct{ *pulumi.OutputState }

func (LookupEventHubEventSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubEventSourceResult)(nil)).Elem()
}

func (o LookupEventHubEventSourceResultOutput) ToLookupEventHubEventSourceResultOutput() LookupEventHubEventSourceResultOutput {
	return o
}

func (o LookupEventHubEventSourceResultOutput) ToLookupEventHubEventSourceResultOutputWithContext(ctx context.Context) LookupEventHubEventSourceResultOutput {
	return o
}

func (o LookupEventHubEventSourceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupEventHubEventSourceResult] {
	return pulumix.Output[LookupEventHubEventSourceResult]{
		OutputState: o.OutputState,
	}
}

// The name of the event hub's consumer group that holds the partitions from which events will be read.
func (o LookupEventHubEventSourceResultOutput) ConsumerGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.ConsumerGroupName }).(pulumi.StringOutput)
}

// The time the resource was created.
func (o LookupEventHubEventSourceResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// The name of the event hub.
func (o LookupEventHubEventSourceResultOutput) EventHubName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.EventHubName }).(pulumi.StringOutput)
}

// The resource id of the event source in Azure Resource Manager.
func (o LookupEventHubEventSourceResultOutput) EventSourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.EventSourceResourceId }).(pulumi.StringOutput)
}

// Resource Id
func (o LookupEventHubEventSourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the SAS key that grants the Time Series Insights service access to the event hub. The shared access policies for this key must grant 'Listen' permissions to the event hub.
func (o LookupEventHubEventSourceResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.KeyName }).(pulumi.StringOutput)
}

// The kind of the event source.
// Expected value is 'Microsoft.EventHub'.
func (o LookupEventHubEventSourceResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.Kind }).(pulumi.StringOutput)
}

// An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
func (o LookupEventHubEventSourceResultOutput) LocalTimestamp() LocalTimestampResponsePtrOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) *LocalTimestampResponse { return v.LocalTimestamp }).(LocalTimestampResponsePtrOutput)
}

// Resource location
func (o LookupEventHubEventSourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupEventHubEventSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupEventHubEventSourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The name of the service bus that contains the event hub.
func (o LookupEventHubEventSourceResultOutput) ServiceBusNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.ServiceBusNamespace }).(pulumi.StringOutput)
}

// Resource tags
func (o LookupEventHubEventSourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// ISO8601 UTC datetime with seconds precision (milliseconds are optional), specifying the date and time that will be the starting point for Events to be consumed.
func (o LookupEventHubEventSourceResultOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) *string { return v.Time }).(pulumi.StringPtrOutput)
}

// The event property that will be used as the event source's timestamp. If a value isn't specified for timestampPropertyName, or if null or empty-string is specified, the event creation time will be used.
func (o LookupEventHubEventSourceResultOutput) TimestampPropertyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) *string { return v.TimestampPropertyName }).(pulumi.StringPtrOutput)
}

// Resource type
func (o LookupEventHubEventSourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubEventSourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventHubEventSourceResultOutput{})
}
