// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210303preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The properties of the EventHubConsumerGroupInfo object.
type IotHubResourceEventHubConsumerGroup struct {
	pulumi.CustomResourceState

	// The etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The Event Hub-compatible consumer group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The tags.
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	// the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIotHubResourceEventHubConsumerGroup registers a new resource with the given unique name, arguments, and options.
func NewIotHubResourceEventHubConsumerGroup(ctx *pulumi.Context,
	name string, args *IotHubResourceEventHubConsumerGroupArgs, opts ...pulumi.ResourceOption) (*IotHubResourceEventHubConsumerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventHubEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EventHubEndpointName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devices:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20160203:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20170119:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20170701:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180122:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180401:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20181201preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190322:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190322preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190701preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20191104:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200301:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200401:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200615:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200710preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200801:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210201preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210331:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210702:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210702preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20220430preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20221115preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20230630:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20230630preview:IotHubResourceEventHubConsumerGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IotHubResourceEventHubConsumerGroup
	err := ctx.RegisterResource("azure-native:devices/v20210303preview:IotHubResourceEventHubConsumerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotHubResourceEventHubConsumerGroup gets an existing IotHubResourceEventHubConsumerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotHubResourceEventHubConsumerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubResourceEventHubConsumerGroupState, opts ...pulumi.ResourceOption) (*IotHubResourceEventHubConsumerGroup, error) {
	var resource IotHubResourceEventHubConsumerGroup
	err := ctx.ReadResource("azure-native:devices/v20210303preview:IotHubResourceEventHubConsumerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotHubResourceEventHubConsumerGroup resources.
type iotHubResourceEventHubConsumerGroupState struct {
}

type IotHubResourceEventHubConsumerGroupState struct {
}

func (IotHubResourceEventHubConsumerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubResourceEventHubConsumerGroupState)(nil)).Elem()
}

type iotHubResourceEventHubConsumerGroupArgs struct {
	// The name of the Event Hub-compatible endpoint in the IoT hub.
	EventHubEndpointName string `pulumi:"eventHubEndpointName"`
	// The name of the consumer group to add.
	Name *string `pulumi:"name"`
	// The EventHub consumer group name.
	Properties *EventHubConsumerGroupName `pulumi:"properties"`
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a IotHubResourceEventHubConsumerGroup resource.
type IotHubResourceEventHubConsumerGroupArgs struct {
	// The name of the Event Hub-compatible endpoint in the IoT hub.
	EventHubEndpointName pulumi.StringInput
	// The name of the consumer group to add.
	Name pulumi.StringPtrInput
	// The EventHub consumer group name.
	Properties EventHubConsumerGroupNamePtrInput
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName pulumi.StringInput
	// The name of the IoT hub.
	ResourceName pulumi.StringInput
}

func (IotHubResourceEventHubConsumerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubResourceEventHubConsumerGroupArgs)(nil)).Elem()
}

type IotHubResourceEventHubConsumerGroupInput interface {
	pulumi.Input

	ToIotHubResourceEventHubConsumerGroupOutput() IotHubResourceEventHubConsumerGroupOutput
	ToIotHubResourceEventHubConsumerGroupOutputWithContext(ctx context.Context) IotHubResourceEventHubConsumerGroupOutput
}

func (*IotHubResourceEventHubConsumerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubResourceEventHubConsumerGroup)(nil)).Elem()
}

func (i *IotHubResourceEventHubConsumerGroup) ToIotHubResourceEventHubConsumerGroupOutput() IotHubResourceEventHubConsumerGroupOutput {
	return i.ToIotHubResourceEventHubConsumerGroupOutputWithContext(context.Background())
}

func (i *IotHubResourceEventHubConsumerGroup) ToIotHubResourceEventHubConsumerGroupOutputWithContext(ctx context.Context) IotHubResourceEventHubConsumerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubResourceEventHubConsumerGroupOutput)
}

func (i *IotHubResourceEventHubConsumerGroup) ToOutput(ctx context.Context) pulumix.Output[*IotHubResourceEventHubConsumerGroup] {
	return pulumix.Output[*IotHubResourceEventHubConsumerGroup]{
		OutputState: i.ToIotHubResourceEventHubConsumerGroupOutputWithContext(ctx).OutputState,
	}
}

type IotHubResourceEventHubConsumerGroupOutput struct{ *pulumi.OutputState }

func (IotHubResourceEventHubConsumerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubResourceEventHubConsumerGroup)(nil)).Elem()
}

func (o IotHubResourceEventHubConsumerGroupOutput) ToIotHubResourceEventHubConsumerGroupOutput() IotHubResourceEventHubConsumerGroupOutput {
	return o
}

func (o IotHubResourceEventHubConsumerGroupOutput) ToIotHubResourceEventHubConsumerGroupOutputWithContext(ctx context.Context) IotHubResourceEventHubConsumerGroupOutput {
	return o
}

func (o IotHubResourceEventHubConsumerGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*IotHubResourceEventHubConsumerGroup] {
	return pulumix.Output[*IotHubResourceEventHubConsumerGroup]{
		OutputState: o.OutputState,
	}
}

// The etag.
func (o IotHubResourceEventHubConsumerGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHubResourceEventHubConsumerGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The Event Hub-compatible consumer group name.
func (o IotHubResourceEventHubConsumerGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHubResourceEventHubConsumerGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The tags.
func (o IotHubResourceEventHubConsumerGroupOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IotHubResourceEventHubConsumerGroup) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

// the resource type.
func (o IotHubResourceEventHubConsumerGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHubResourceEventHubConsumerGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IotHubResourceEventHubConsumerGroupOutput{})
}
