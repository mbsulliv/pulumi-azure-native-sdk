// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Class representing an event hub connection.
// Azure REST API version: 2018-09-07-preview. Prior API version in Azure Native 1.x: 2018-09-07-preview.
type EventHubConnection struct {
	pulumi.CustomResourceState

	// The event hub consumer group.
	ConsumerGroup pulumi.StringOutput `pulumi:"consumerGroup"`
	// The data format of the message. Optionally the data format can be added to each message.
	DataFormat pulumi.StringPtrOutput `pulumi:"dataFormat"`
	// The resource ID of the event hub to be used to create a data connection.
	EventHubResourceId pulumi.StringOutput `pulumi:"eventHubResourceId"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
	MappingRuleName pulumi.StringPtrOutput `pulumi:"mappingRuleName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName pulumi.StringPtrOutput `pulumi:"tableName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEventHubConnection registers a new resource with the given unique name, arguments, and options.
func NewEventHubConnection(ctx *pulumi.Context,
	name string, args *EventHubConnectionArgs, opts ...pulumi.ResourceOption) (*EventHubConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ConsumerGroup == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroup'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.EventHubResourceId == nil {
		return nil, errors.New("invalid value for required argument 'EventHubResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto/v20170907privatepreview:EventHubConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20180907preview:EventHubConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EventHubConnection
	err := ctx.RegisterResource("azure-native:kusto:EventHubConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventHubConnection gets an existing EventHubConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHubConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubConnectionState, opts ...pulumi.ResourceOption) (*EventHubConnection, error) {
	var resource EventHubConnection
	err := ctx.ReadResource("azure-native:kusto:EventHubConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventHubConnection resources.
type eventHubConnectionState struct {
}

type EventHubConnectionState struct {
}

func (EventHubConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubConnectionState)(nil)).Elem()
}

type eventHubConnectionArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The event hub consumer group.
	ConsumerGroup string `pulumi:"consumerGroup"`
	// The data format of the message. Optionally the data format can be added to each message.
	DataFormat *string `pulumi:"dataFormat"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the event hub connection.
	EventHubConnectionName *string `pulumi:"eventHubConnectionName"`
	// The resource ID of the event hub to be used to create a data connection.
	EventHubResourceId string `pulumi:"eventHubResourceId"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName *string `pulumi:"tableName"`
}

// The set of arguments for constructing a EventHubConnection resource.
type EventHubConnectionArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// The event hub consumer group.
	ConsumerGroup pulumi.StringInput
	// The data format of the message. Optionally the data format can be added to each message.
	DataFormat pulumi.StringPtrInput
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringInput
	// The name of the event hub connection.
	EventHubConnectionName pulumi.StringPtrInput
	// The resource ID of the event hub to be used to create a data connection.
	EventHubResourceId pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
	MappingRuleName pulumi.StringPtrInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
	// The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName pulumi.StringPtrInput
}

func (EventHubConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubConnectionArgs)(nil)).Elem()
}

type EventHubConnectionInput interface {
	pulumi.Input

	ToEventHubConnectionOutput() EventHubConnectionOutput
	ToEventHubConnectionOutputWithContext(ctx context.Context) EventHubConnectionOutput
}

func (*EventHubConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubConnection)(nil)).Elem()
}

func (i *EventHubConnection) ToEventHubConnectionOutput() EventHubConnectionOutput {
	return i.ToEventHubConnectionOutputWithContext(context.Background())
}

func (i *EventHubConnection) ToEventHubConnectionOutputWithContext(ctx context.Context) EventHubConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubConnectionOutput)
}

func (i *EventHubConnection) ToOutput(ctx context.Context) pulumix.Output[*EventHubConnection] {
	return pulumix.Output[*EventHubConnection]{
		OutputState: i.ToEventHubConnectionOutputWithContext(ctx).OutputState,
	}
}

type EventHubConnectionOutput struct{ *pulumi.OutputState }

func (EventHubConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubConnection)(nil)).Elem()
}

func (o EventHubConnectionOutput) ToEventHubConnectionOutput() EventHubConnectionOutput {
	return o
}

func (o EventHubConnectionOutput) ToEventHubConnectionOutputWithContext(ctx context.Context) EventHubConnectionOutput {
	return o
}

func (o EventHubConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*EventHubConnection] {
	return pulumix.Output[*EventHubConnection]{
		OutputState: o.OutputState,
	}
}

// The event hub consumer group.
func (o EventHubConnectionOutput) ConsumerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubConnection) pulumi.StringOutput { return v.ConsumerGroup }).(pulumi.StringOutput)
}

// The data format of the message. Optionally the data format can be added to each message.
func (o EventHubConnectionOutput) DataFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubConnection) pulumi.StringPtrOutput { return v.DataFormat }).(pulumi.StringPtrOutput)
}

// The resource ID of the event hub to be used to create a data connection.
func (o EventHubConnectionOutput) EventHubResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubConnection) pulumi.StringOutput { return v.EventHubResourceId }).(pulumi.StringOutput)
}

// Resource location.
func (o EventHubConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
func (o EventHubConnectionOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubConnection) pulumi.StringPtrOutput { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o EventHubConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The table where the data should be ingested. Optionally the table information can be added to each message.
func (o EventHubConnectionOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubConnection) pulumi.StringPtrOutput { return v.TableName }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EventHubConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventHubConnectionOutput{})
}
