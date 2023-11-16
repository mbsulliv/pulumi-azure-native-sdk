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
func LookupEventHubDataConnection(ctx *pulumi.Context, args *LookupEventHubDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEventHubDataConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEventHubDataConnectionResult
	err := ctx.Invoke("azure-native:kusto/v20230502:getEventHubDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupEventHubDataConnectionArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the data connection.
	DataConnectionName string `pulumi:"dataConnectionName"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing an event hub data connection.
type LookupEventHubDataConnectionResult struct {
	// The event hub messages compression type
	Compression *string `pulumi:"compression"`
	// The event hub consumer group.
	ConsumerGroup string `pulumi:"consumerGroup"`
	// The data format of the message. Optionally the data format can be added to each message.
	DataFormat *string `pulumi:"dataFormat"`
	// Indication for database routing information from the data connection, by default only database routing information is allowed
	DatabaseRouting *string `pulumi:"databaseRouting"`
	// The resource ID of the event hub to be used to create a data connection.
	EventHubResourceId string `pulumi:"eventHubResourceId"`
	// System properties of the event hub
	EventSystemProperties []string `pulumi:"eventSystemProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Kind of the endpoint for the data connection
	// Expected value is 'EventHub'.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The object ID of the managedIdentityResourceId
	ManagedIdentityObjectId string `pulumi:"managedIdentityObjectId"`
	// The resource ID of a managed identity (system or user assigned) to be used to authenticate with event hub.
	ManagedIdentityResourceId *string `pulumi:"managedIdentityResourceId"`
	// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// When defined, the data connection retrieves existing Event hub events created since the Retrieval start date. It can only retrieve events retained by the Event hub, based on its retention period.
	RetrievalStartDate *string `pulumi:"retrievalStartDate"`
	// The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName *string `pulumi:"tableName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupEventHubDataConnectionResult
func (val *LookupEventHubDataConnectionResult) Defaults() *LookupEventHubDataConnectionResult {
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

func LookupEventHubDataConnectionOutput(ctx *pulumi.Context, args LookupEventHubDataConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupEventHubDataConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventHubDataConnectionResult, error) {
			args := v.(LookupEventHubDataConnectionArgs)
			r, err := LookupEventHubDataConnection(ctx, &args, opts...)
			var s LookupEventHubDataConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventHubDataConnectionResultOutput)
}

type LookupEventHubDataConnectionOutputArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the data connection.
	DataConnectionName pulumi.StringInput `pulumi:"dataConnectionName"`
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventHubDataConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubDataConnectionArgs)(nil)).Elem()
}

// Class representing an event hub data connection.
type LookupEventHubDataConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupEventHubDataConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubDataConnectionResult)(nil)).Elem()
}

func (o LookupEventHubDataConnectionResultOutput) ToLookupEventHubDataConnectionResultOutput() LookupEventHubDataConnectionResultOutput {
	return o
}

func (o LookupEventHubDataConnectionResultOutput) ToLookupEventHubDataConnectionResultOutputWithContext(ctx context.Context) LookupEventHubDataConnectionResultOutput {
	return o
}

// The event hub messages compression type
func (o LookupEventHubDataConnectionResultOutput) Compression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) *string { return v.Compression }).(pulumi.StringPtrOutput)
}

// The event hub consumer group.
func (o LookupEventHubDataConnectionResultOutput) ConsumerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.ConsumerGroup }).(pulumi.StringOutput)
}

// The data format of the message. Optionally the data format can be added to each message.
func (o LookupEventHubDataConnectionResultOutput) DataFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) *string { return v.DataFormat }).(pulumi.StringPtrOutput)
}

// Indication for database routing information from the data connection, by default only database routing information is allowed
func (o LookupEventHubDataConnectionResultOutput) DatabaseRouting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) *string { return v.DatabaseRouting }).(pulumi.StringPtrOutput)
}

// The resource ID of the event hub to be used to create a data connection.
func (o LookupEventHubDataConnectionResultOutput) EventHubResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.EventHubResourceId }).(pulumi.StringOutput)
}

// System properties of the event hub
func (o LookupEventHubDataConnectionResultOutput) EventSystemProperties() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) []string { return v.EventSystemProperties }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEventHubDataConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of the endpoint for the data connection
// Expected value is 'EventHub'.
func (o LookupEventHubDataConnectionResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupEventHubDataConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The object ID of the managedIdentityResourceId
func (o LookupEventHubDataConnectionResultOutput) ManagedIdentityObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.ManagedIdentityObjectId }).(pulumi.StringOutput)
}

// The resource ID of a managed identity (system or user assigned) to be used to authenticate with event hub.
func (o LookupEventHubDataConnectionResultOutput) ManagedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) *string { return v.ManagedIdentityResourceId }).(pulumi.StringPtrOutput)
}

// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
func (o LookupEventHubDataConnectionResultOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) *string { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupEventHubDataConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o LookupEventHubDataConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// When defined, the data connection retrieves existing Event hub events created since the Retrieval start date. It can only retrieve events retained by the Event hub, based on its retention period.
func (o LookupEventHubDataConnectionResultOutput) RetrievalStartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) *string { return v.RetrievalStartDate }).(pulumi.StringPtrOutput)
}

// The table where the data should be ingested. Optionally the table information can be added to each message.
func (o LookupEventHubDataConnectionResultOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) *string { return v.TableName }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEventHubDataConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubDataConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventHubDataConnectionResultOutput{})
}
