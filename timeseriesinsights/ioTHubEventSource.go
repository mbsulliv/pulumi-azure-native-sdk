// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package timeseriesinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An event source that receives its data from an Azure IoTHub.
// Azure REST API version: 2020-05-15. Prior API version in Azure Native 1.x: 2020-05-15.
type IoTHubEventSource struct {
	pulumi.CustomResourceState

	// The name of the iot hub's consumer group that holds the partitions from which events will be read.
	ConsumerGroupName pulumi.StringOutput `pulumi:"consumerGroupName"`
	// The time the resource was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The resource id of the event source in Azure Resource Manager.
	EventSourceResourceId pulumi.StringOutput `pulumi:"eventSourceResourceId"`
	// The name of the iot hub.
	IotHubName pulumi.StringOutput `pulumi:"iotHubName"`
	// The name of the Shared Access Policy key that grants the Time Series Insights service access to the iot hub. This shared access policy key must grant 'service connect' permissions to the iot hub.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
	// The kind of the event source.
	// Expected value is 'Microsoft.IoTHub'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
	LocalTimestamp LocalTimestampResponsePtrOutput `pulumi:"localTimestamp"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// ISO8601 UTC datetime with seconds precision (milliseconds are optional), specifying the date and time that will be the starting point for Events to be consumed.
	Time pulumi.StringPtrOutput `pulumi:"time"`
	// The event property that will be used as the event source's timestamp. If a value isn't specified for timestampPropertyName, or if null or empty-string is specified, the event creation time will be used.
	TimestampPropertyName pulumi.StringPtrOutput `pulumi:"timestampPropertyName"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIoTHubEventSource registers a new resource with the given unique name, arguments, and options.
func NewIoTHubEventSource(ctx *pulumi.Context,
	name string, args *IoTHubEventSourceArgs, opts ...pulumi.ResourceOption) (*IoTHubEventSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroupName'")
	}
	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.EventSourceResourceId == nil {
		return nil, errors.New("invalid value for required argument 'EventSourceResourceId'")
	}
	if args.IotHubName == nil {
		return nil, errors.New("invalid value for required argument 'IotHubName'")
	}
	if args.KeyName == nil {
		return nil, errors.New("invalid value for required argument 'KeyName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SharedAccessKey == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessKey'")
	}
	args.Kind = pulumi.String("Microsoft.IoTHub")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20170228preview:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20180815preview:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20200515:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:IoTHubEventSource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IoTHubEventSource
	err := ctx.RegisterResource("azure-native:timeseriesinsights:IoTHubEventSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIoTHubEventSource gets an existing IoTHubEventSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIoTHubEventSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IoTHubEventSourceState, opts ...pulumi.ResourceOption) (*IoTHubEventSource, error) {
	var resource IoTHubEventSource
	err := ctx.ReadResource("azure-native:timeseriesinsights:IoTHubEventSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IoTHubEventSource resources.
type ioTHubEventSourceState struct {
}

type IoTHubEventSourceState struct {
}

func (IoTHubEventSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTHubEventSourceState)(nil)).Elem()
}

type ioTHubEventSourceArgs struct {
	// The name of the iot hub's consumer group that holds the partitions from which events will be read.
	ConsumerGroupName string `pulumi:"consumerGroupName"`
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName string `pulumi:"environmentName"`
	// Name of the event source.
	EventSourceName *string `pulumi:"eventSourceName"`
	// The resource id of the event source in Azure Resource Manager.
	EventSourceResourceId string `pulumi:"eventSourceResourceId"`
	// The name of the iot hub.
	IotHubName string `pulumi:"iotHubName"`
	// The name of the Shared Access Policy key that grants the Time Series Insights service access to the iot hub. This shared access policy key must grant 'service connect' permissions to the iot hub.
	KeyName string `pulumi:"keyName"`
	// The kind of the event source.
	// Expected value is 'Microsoft.IoTHub'.
	Kind string `pulumi:"kind"`
	// An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
	LocalTimestamp *LocalTimestamp `pulumi:"localTimestamp"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The value of the Shared Access Policy key that grants the Time Series Insights service read access to the iot hub. This property is not shown in event source responses.
	SharedAccessKey string `pulumi:"sharedAccessKey"`
	// Key-value pairs of additional properties for the resource.
	Tags map[string]string `pulumi:"tags"`
	// ISO8601 UTC datetime with seconds precision (milliseconds are optional), specifying the date and time that will be the starting point for Events to be consumed.
	Time *string `pulumi:"time"`
	// The event property that will be used as the event source's timestamp. If a value isn't specified for timestampPropertyName, or if null or empty-string is specified, the event creation time will be used.
	TimestampPropertyName *string `pulumi:"timestampPropertyName"`
	// The type of the ingressStartAt, It can be "EarliestAvailable", "EventSourceCreationTime", "CustomEnqueuedTime".
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a IoTHubEventSource resource.
type IoTHubEventSourceArgs struct {
	// The name of the iot hub's consumer group that holds the partitions from which events will be read.
	ConsumerGroupName pulumi.StringInput
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName pulumi.StringInput
	// Name of the event source.
	EventSourceName pulumi.StringPtrInput
	// The resource id of the event source in Azure Resource Manager.
	EventSourceResourceId pulumi.StringInput
	// The name of the iot hub.
	IotHubName pulumi.StringInput
	// The name of the Shared Access Policy key that grants the Time Series Insights service access to the iot hub. This shared access policy key must grant 'service connect' permissions to the iot hub.
	KeyName pulumi.StringInput
	// The kind of the event source.
	// Expected value is 'Microsoft.IoTHub'.
	Kind pulumi.StringInput
	// An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
	LocalTimestamp LocalTimestampPtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// The value of the Shared Access Policy key that grants the Time Series Insights service read access to the iot hub. This property is not shown in event source responses.
	SharedAccessKey pulumi.StringInput
	// Key-value pairs of additional properties for the resource.
	Tags pulumi.StringMapInput
	// ISO8601 UTC datetime with seconds precision (milliseconds are optional), specifying the date and time that will be the starting point for Events to be consumed.
	Time pulumi.StringPtrInput
	// The event property that will be used as the event source's timestamp. If a value isn't specified for timestampPropertyName, or if null or empty-string is specified, the event creation time will be used.
	TimestampPropertyName pulumi.StringPtrInput
	// The type of the ingressStartAt, It can be "EarliestAvailable", "EventSourceCreationTime", "CustomEnqueuedTime".
	Type pulumi.StringPtrInput
}

func (IoTHubEventSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTHubEventSourceArgs)(nil)).Elem()
}

type IoTHubEventSourceInput interface {
	pulumi.Input

	ToIoTHubEventSourceOutput() IoTHubEventSourceOutput
	ToIoTHubEventSourceOutputWithContext(ctx context.Context) IoTHubEventSourceOutput
}

func (*IoTHubEventSource) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTHubEventSource)(nil)).Elem()
}

func (i *IoTHubEventSource) ToIoTHubEventSourceOutput() IoTHubEventSourceOutput {
	return i.ToIoTHubEventSourceOutputWithContext(context.Background())
}

func (i *IoTHubEventSource) ToIoTHubEventSourceOutputWithContext(ctx context.Context) IoTHubEventSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTHubEventSourceOutput)
}

type IoTHubEventSourceOutput struct{ *pulumi.OutputState }

func (IoTHubEventSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTHubEventSource)(nil)).Elem()
}

func (o IoTHubEventSourceOutput) ToIoTHubEventSourceOutput() IoTHubEventSourceOutput {
	return o
}

func (o IoTHubEventSourceOutput) ToIoTHubEventSourceOutputWithContext(ctx context.Context) IoTHubEventSourceOutput {
	return o
}

// The name of the iot hub's consumer group that holds the partitions from which events will be read.
func (o IoTHubEventSourceOutput) ConsumerGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.ConsumerGroupName }).(pulumi.StringOutput)
}

// The time the resource was created.
func (o IoTHubEventSourceOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The resource id of the event source in Azure Resource Manager.
func (o IoTHubEventSourceOutput) EventSourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.EventSourceResourceId }).(pulumi.StringOutput)
}

// The name of the iot hub.
func (o IoTHubEventSourceOutput) IotHubName() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.IotHubName }).(pulumi.StringOutput)
}

// The name of the Shared Access Policy key that grants the Time Series Insights service access to the iot hub. This shared access policy key must grant 'service connect' permissions to the iot hub.
func (o IoTHubEventSourceOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

// The kind of the event source.
// Expected value is 'Microsoft.IoTHub'.
func (o IoTHubEventSourceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
func (o IoTHubEventSourceOutput) LocalTimestamp() LocalTimestampResponsePtrOutput {
	return o.ApplyT(func(v *IoTHubEventSource) LocalTimestampResponsePtrOutput { return v.LocalTimestamp }).(LocalTimestampResponsePtrOutput)
}

// Resource location
func (o IoTHubEventSourceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o IoTHubEventSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o IoTHubEventSourceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags
func (o IoTHubEventSourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// ISO8601 UTC datetime with seconds precision (milliseconds are optional), specifying the date and time that will be the starting point for Events to be consumed.
func (o IoTHubEventSourceOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringPtrOutput { return v.Time }).(pulumi.StringPtrOutput)
}

// The event property that will be used as the event source's timestamp. If a value isn't specified for timestampPropertyName, or if null or empty-string is specified, the event creation time will be used.
func (o IoTHubEventSourceOutput) TimestampPropertyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringPtrOutput { return v.TimestampPropertyName }).(pulumi.StringPtrOutput)
}

// Resource type
func (o IoTHubEventSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IoTHubEventSourceOutput{})
}
