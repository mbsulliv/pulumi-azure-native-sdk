// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Single item in List or Get Event Hub operation
type EventHub struct {
	pulumi.CustomResourceState

	// Properties of capture description
	CaptureDescription CaptureDescriptionResponsePtrOutput `pulumi:"captureDescription"`
	// Exact time the Event Hub was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	MessageRetentionInDays pulumi.Float64PtrOutput `pulumi:"messageRetentionInDays"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
	PartitionCount pulumi.Float64PtrOutput `pulumi:"partitionCount"`
	// Current number of shards on the Event Hub.
	PartitionIds pulumi.StringArrayOutput `pulumi:"partitionIds"`
	// Event Hub retention settings
	RetentionDescription RetentionDescriptionResponsePtrOutput `pulumi:"retentionDescription"`
	// Enumerates the possible values for the status of the Event Hub.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type pulumi.StringOutput `pulumi:"type"`
	// The exact time the message was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewEventHub registers a new resource with the given unique name, arguments, and options.
func NewEventHub(ctx *pulumi.Context,
	name string, args *EventHubArgs, opts ...pulumi.ResourceOption) (*EventHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventhub:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20140901:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20150801:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20170401:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20180101preview:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210101preview:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210601preview:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20211101:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20221001preview:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20230101preview:EventHub"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EventHub
	err := ctx.RegisterResource("azure-native:eventhub/v20240101:EventHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventHub gets an existing EventHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubState, opts ...pulumi.ResourceOption) (*EventHub, error) {
	var resource EventHub
	err := ctx.ReadResource("azure-native:eventhub/v20240101:EventHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventHub resources.
type eventHubState struct {
}

type EventHubState struct {
}

func (EventHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubState)(nil)).Elem()
}

type eventHubArgs struct {
	// Properties of capture description
	CaptureDescription *CaptureDescription `pulumi:"captureDescription"`
	// The Event Hub name
	EventHubName *string `pulumi:"eventHubName"`
	// Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	MessageRetentionInDays *float64 `pulumi:"messageRetentionInDays"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
	PartitionCount *float64 `pulumi:"partitionCount"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Event Hub retention settings
	RetentionDescription *RetentionDescription `pulumi:"retentionDescription"`
	// Enumerates the possible values for the status of the Event Hub.
	Status *EntityStatus `pulumi:"status"`
}

// The set of arguments for constructing a EventHub resource.
type EventHubArgs struct {
	// Properties of capture description
	CaptureDescription CaptureDescriptionPtrInput
	// The Event Hub name
	EventHubName pulumi.StringPtrInput
	// Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	MessageRetentionInDays pulumi.Float64PtrInput
	// The Namespace name
	NamespaceName pulumi.StringInput
	// Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
	PartitionCount pulumi.Float64PtrInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
	// Event Hub retention settings
	RetentionDescription RetentionDescriptionPtrInput
	// Enumerates the possible values for the status of the Event Hub.
	Status EntityStatusPtrInput
}

func (EventHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubArgs)(nil)).Elem()
}

type EventHubInput interface {
	pulumi.Input

	ToEventHubOutput() EventHubOutput
	ToEventHubOutputWithContext(ctx context.Context) EventHubOutput
}

func (*EventHub) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHub)(nil)).Elem()
}

func (i *EventHub) ToEventHubOutput() EventHubOutput {
	return i.ToEventHubOutputWithContext(context.Background())
}

func (i *EventHub) ToEventHubOutputWithContext(ctx context.Context) EventHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubOutput)
}

type EventHubOutput struct{ *pulumi.OutputState }

func (EventHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHub)(nil)).Elem()
}

func (o EventHubOutput) ToEventHubOutput() EventHubOutput {
	return o
}

func (o EventHubOutput) ToEventHubOutputWithContext(ctx context.Context) EventHubOutput {
	return o
}

// Properties of capture description
func (o EventHubOutput) CaptureDescription() CaptureDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *EventHub) CaptureDescriptionResponsePtrOutput { return v.CaptureDescription }).(CaptureDescriptionResponsePtrOutput)
}

// Exact time the Event Hub was created.
func (o EventHubOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o EventHubOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Number of days to retain the events for this Event Hub, value should be 1 to 7 days
func (o EventHubOutput) MessageRetentionInDays() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EventHub) pulumi.Float64PtrOutput { return v.MessageRetentionInDays }).(pulumi.Float64PtrOutput)
}

// The name of the resource
func (o EventHubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
func (o EventHubOutput) PartitionCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EventHub) pulumi.Float64PtrOutput { return v.PartitionCount }).(pulumi.Float64PtrOutput)
}

// Current number of shards on the Event Hub.
func (o EventHubOutput) PartitionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringArrayOutput { return v.PartitionIds }).(pulumi.StringArrayOutput)
}

// Event Hub retention settings
func (o EventHubOutput) RetentionDescription() RetentionDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *EventHub) RetentionDescriptionResponsePtrOutput { return v.RetentionDescription }).(RetentionDescriptionResponsePtrOutput)
}

// Enumerates the possible values for the status of the Event Hub.
func (o EventHubOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// The system meta data relating to this resource.
func (o EventHubOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EventHub) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o EventHubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The exact time the message was updated.
func (o EventHubOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventHubOutput{})
}
