// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230502

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a data connection.
func LookupEventGridDataConnection(ctx *pulumi.Context, args *LookupEventGridDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEventGridDataConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEventGridDataConnectionResult
	err := ctx.Invoke("azure-native:kusto/v20230502:getEventGridDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupEventGridDataConnectionArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the data connection.
	DataConnectionName string `pulumi:"dataConnectionName"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing an Event Grid data connection.
type LookupEventGridDataConnectionResult struct {
	// The name of blob storage event type to process.
	BlobStorageEventType *string `pulumi:"blobStorageEventType"`
	// The event hub consumer group.
	ConsumerGroup string `pulumi:"consumerGroup"`
	// The data format of the message. Optionally the data format can be added to each message.
	DataFormat *string `pulumi:"dataFormat"`
	// Indication for database routing information from the data connection, by default only database routing information is allowed
	DatabaseRouting *string `pulumi:"databaseRouting"`
	// The resource ID of the event grid that is subscribed to the storage account events.
	EventGridResourceId *string `pulumi:"eventGridResourceId"`
	// The resource ID where the event grid is configured to send events.
	EventHubResourceId string `pulumi:"eventHubResourceId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// A Boolean value that, if set to true, indicates that ingestion should ignore the first record of every file
	IgnoreFirstRecord *bool `pulumi:"ignoreFirstRecord"`
	// Kind of the endpoint for the data connection
	// Expected value is 'EventGrid'.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The object ID of managedIdentityResourceId
	ManagedIdentityObjectId string `pulumi:"managedIdentityObjectId"`
	// The resource ID of a managed identity (system or user assigned) to be used to authenticate with event hub and storage account.
	ManagedIdentityResourceId *string `pulumi:"managedIdentityResourceId"`
	// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource ID of the storage account where the data resides.
	StorageAccountResourceId string `pulumi:"storageAccountResourceId"`
	// The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName *string `pulumi:"tableName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupEventGridDataConnectionResult
func (val *LookupEventGridDataConnectionResult) Defaults() *LookupEventGridDataConnectionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.DatabaseRouting == nil {
		databaseRouting_ := "Single"
		tmp.DatabaseRouting = &databaseRouting_
	}
	return &tmp
}

func LookupEventGridDataConnectionOutput(ctx *pulumi.Context, args LookupEventGridDataConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupEventGridDataConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventGridDataConnectionResult, error) {
			args := v.(LookupEventGridDataConnectionArgs)
			r, err := LookupEventGridDataConnection(ctx, &args, opts...)
			var s LookupEventGridDataConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventGridDataConnectionResultOutput)
}

type LookupEventGridDataConnectionOutputArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the data connection.
	DataConnectionName pulumi.StringInput `pulumi:"dataConnectionName"`
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventGridDataConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventGridDataConnectionArgs)(nil)).Elem()
}

// Class representing an Event Grid data connection.
type LookupEventGridDataConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupEventGridDataConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventGridDataConnectionResult)(nil)).Elem()
}

func (o LookupEventGridDataConnectionResultOutput) ToLookupEventGridDataConnectionResultOutput() LookupEventGridDataConnectionResultOutput {
	return o
}

func (o LookupEventGridDataConnectionResultOutput) ToLookupEventGridDataConnectionResultOutputWithContext(ctx context.Context) LookupEventGridDataConnectionResultOutput {
	return o
}

// The name of blob storage event type to process.
func (o LookupEventGridDataConnectionResultOutput) BlobStorageEventType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *string { return v.BlobStorageEventType }).(pulumi.StringPtrOutput)
}

// The event hub consumer group.
func (o LookupEventGridDataConnectionResultOutput) ConsumerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.ConsumerGroup }).(pulumi.StringOutput)
}

// The data format of the message. Optionally the data format can be added to each message.
func (o LookupEventGridDataConnectionResultOutput) DataFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *string { return v.DataFormat }).(pulumi.StringPtrOutput)
}

// Indication for database routing information from the data connection, by default only database routing information is allowed
func (o LookupEventGridDataConnectionResultOutput) DatabaseRouting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *string { return v.DatabaseRouting }).(pulumi.StringPtrOutput)
}

// The resource ID of the event grid that is subscribed to the storage account events.
func (o LookupEventGridDataConnectionResultOutput) EventGridResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *string { return v.EventGridResourceId }).(pulumi.StringPtrOutput)
}

// The resource ID where the event grid is configured to send events.
func (o LookupEventGridDataConnectionResultOutput) EventHubResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.EventHubResourceId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEventGridDataConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// A Boolean value that, if set to true, indicates that ingestion should ignore the first record of every file
func (o LookupEventGridDataConnectionResultOutput) IgnoreFirstRecord() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *bool { return v.IgnoreFirstRecord }).(pulumi.BoolPtrOutput)
}

// Kind of the endpoint for the data connection
// Expected value is 'EventGrid'.
func (o LookupEventGridDataConnectionResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupEventGridDataConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The object ID of managedIdentityResourceId
func (o LookupEventGridDataConnectionResultOutput) ManagedIdentityObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.ManagedIdentityObjectId }).(pulumi.StringOutput)
}

// The resource ID of a managed identity (system or user assigned) to be used to authenticate with event hub and storage account.
func (o LookupEventGridDataConnectionResultOutput) ManagedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *string { return v.ManagedIdentityResourceId }).(pulumi.StringPtrOutput)
}

// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
func (o LookupEventGridDataConnectionResultOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *string { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupEventGridDataConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o LookupEventGridDataConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource ID of the storage account where the data resides.
func (o LookupEventGridDataConnectionResultOutput) StorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.StorageAccountResourceId }).(pulumi.StringOutput)
}

// The table where the data should be ingested. Optionally the table information can be added to each message.
func (o LookupEventGridDataConnectionResultOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) *string { return v.TableName }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEventGridDataConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventGridDataConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventGridDataConnectionResultOutput{})
}
