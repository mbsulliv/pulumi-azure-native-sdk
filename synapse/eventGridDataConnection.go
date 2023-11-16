// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing an Event Grid data connection.
// Azure REST API version: 2021-06-01-preview.
type EventGridDataConnection struct {
	pulumi.CustomResourceState

	// The name of blob storage event type to process.
	BlobStorageEventType pulumi.StringPtrOutput `pulumi:"blobStorageEventType"`
	// The event hub consumer group.
	ConsumerGroup pulumi.StringOutput `pulumi:"consumerGroup"`
	// The data format of the message. Optionally the data format can be added to each message.
	DataFormat pulumi.StringPtrOutput `pulumi:"dataFormat"`
	// The resource ID where the event grid is configured to send events.
	EventHubResourceId pulumi.StringOutput `pulumi:"eventHubResourceId"`
	// A Boolean value that, if set to true, indicates that ingestion should ignore the first record of every file
	IgnoreFirstRecord pulumi.BoolPtrOutput `pulumi:"ignoreFirstRecord"`
	// Kind of the endpoint for the data connection
	// Expected value is 'EventGrid'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
	MappingRuleName pulumi.StringPtrOutput `pulumi:"mappingRuleName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource ID of the storage account where the data resides.
	StorageAccountResourceId pulumi.StringOutput `pulumi:"storageAccountResourceId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName pulumi.StringPtrOutput `pulumi:"tableName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEventGridDataConnection registers a new resource with the given unique name, arguments, and options.
func NewEventGridDataConnection(ctx *pulumi.Context,
	name string, args *EventGridDataConnectionArgs, opts ...pulumi.ResourceOption) (*EventGridDataConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.KustoPoolName == nil {
		return nil, errors.New("invalid value for required argument 'KustoPoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountResourceId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountResourceId'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("EventGrid")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:EventGridDataConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EventGridDataConnection
	err := ctx.RegisterResource("azure-native:synapse:EventGridDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventGridDataConnection gets an existing EventGridDataConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventGridDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventGridDataConnectionState, opts ...pulumi.ResourceOption) (*EventGridDataConnection, error) {
	var resource EventGridDataConnection
	err := ctx.ReadResource("azure-native:synapse:EventGridDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventGridDataConnection resources.
type eventGridDataConnectionState struct {
}

type EventGridDataConnectionState struct {
}

func (EventGridDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventGridDataConnectionState)(nil)).Elem()
}

type eventGridDataConnectionArgs struct {
	// The name of blob storage event type to process.
	BlobStorageEventType *string `pulumi:"blobStorageEventType"`
	// The event hub consumer group.
	ConsumerGroup string `pulumi:"consumerGroup"`
	// The name of the data connection.
	DataConnectionName *string `pulumi:"dataConnectionName"`
	// The data format of the message. Optionally the data format can be added to each message.
	DataFormat *string `pulumi:"dataFormat"`
	// The name of the database in the Kusto pool.
	DatabaseName string `pulumi:"databaseName"`
	// The resource ID where the event grid is configured to send events.
	EventHubResourceId string `pulumi:"eventHubResourceId"`
	// A Boolean value that, if set to true, indicates that ingestion should ignore the first record of every file
	IgnoreFirstRecord *bool `pulumi:"ignoreFirstRecord"`
	// Kind of the endpoint for the data connection
	// Expected value is 'EventGrid'.
	Kind string `pulumi:"kind"`
	// The name of the Kusto pool.
	KustoPoolName string `pulumi:"kustoPoolName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource ID of the storage account where the data resides.
	StorageAccountResourceId string `pulumi:"storageAccountResourceId"`
	// The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName *string `pulumi:"tableName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a EventGridDataConnection resource.
type EventGridDataConnectionArgs struct {
	// The name of blob storage event type to process.
	BlobStorageEventType pulumi.StringPtrInput
	// The event hub consumer group.
	ConsumerGroup pulumi.StringInput
	// The name of the data connection.
	DataConnectionName pulumi.StringPtrInput
	// The data format of the message. Optionally the data format can be added to each message.
	DataFormat pulumi.StringPtrInput
	// The name of the database in the Kusto pool.
	DatabaseName pulumi.StringInput
	// The resource ID where the event grid is configured to send events.
	EventHubResourceId pulumi.StringInput
	// A Boolean value that, if set to true, indicates that ingestion should ignore the first record of every file
	IgnoreFirstRecord pulumi.BoolPtrInput
	// Kind of the endpoint for the data connection
	// Expected value is 'EventGrid'.
	Kind pulumi.StringInput
	// The name of the Kusto pool.
	KustoPoolName pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
	MappingRuleName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The resource ID of the storage account where the data resides.
	StorageAccountResourceId pulumi.StringInput
	// The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (EventGridDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventGridDataConnectionArgs)(nil)).Elem()
}

type EventGridDataConnectionInput interface {
	pulumi.Input

	ToEventGridDataConnectionOutput() EventGridDataConnectionOutput
	ToEventGridDataConnectionOutputWithContext(ctx context.Context) EventGridDataConnectionOutput
}

func (*EventGridDataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**EventGridDataConnection)(nil)).Elem()
}

func (i *EventGridDataConnection) ToEventGridDataConnectionOutput() EventGridDataConnectionOutput {
	return i.ToEventGridDataConnectionOutputWithContext(context.Background())
}

func (i *EventGridDataConnection) ToEventGridDataConnectionOutputWithContext(ctx context.Context) EventGridDataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGridDataConnectionOutput)
}

type EventGridDataConnectionOutput struct{ *pulumi.OutputState }

func (EventGridDataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventGridDataConnection)(nil)).Elem()
}

func (o EventGridDataConnectionOutput) ToEventGridDataConnectionOutput() EventGridDataConnectionOutput {
	return o
}

func (o EventGridDataConnectionOutput) ToEventGridDataConnectionOutputWithContext(ctx context.Context) EventGridDataConnectionOutput {
	return o
}

// The name of blob storage event type to process.
func (o EventGridDataConnectionOutput) BlobStorageEventType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringPtrOutput { return v.BlobStorageEventType }).(pulumi.StringPtrOutput)
}

// The event hub consumer group.
func (o EventGridDataConnectionOutput) ConsumerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.ConsumerGroup }).(pulumi.StringOutput)
}

// The data format of the message. Optionally the data format can be added to each message.
func (o EventGridDataConnectionOutput) DataFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringPtrOutput { return v.DataFormat }).(pulumi.StringPtrOutput)
}

// The resource ID where the event grid is configured to send events.
func (o EventGridDataConnectionOutput) EventHubResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.EventHubResourceId }).(pulumi.StringOutput)
}

// A Boolean value that, if set to true, indicates that ingestion should ignore the first record of every file
func (o EventGridDataConnectionOutput) IgnoreFirstRecord() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.BoolPtrOutput { return v.IgnoreFirstRecord }).(pulumi.BoolPtrOutput)
}

// Kind of the endpoint for the data connection
// Expected value is 'EventGrid'.
func (o EventGridDataConnectionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o EventGridDataConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
func (o EventGridDataConnectionOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringPtrOutput { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o EventGridDataConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o EventGridDataConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource ID of the storage account where the data resides.
func (o EventGridDataConnectionOutput) StorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.StorageAccountResourceId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o EventGridDataConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EventGridDataConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The table where the data should be ingested. Optionally the table information can be added to each message.
func (o EventGridDataConnectionOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringPtrOutput { return v.TableName }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EventGridDataConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventGridDataConnectionOutput{})
}
