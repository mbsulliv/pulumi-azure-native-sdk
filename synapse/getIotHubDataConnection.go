// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a data connection.
// Azure REST API version: 2021-06-01-preview.
func LookupIotHubDataConnection(ctx *pulumi.Context, args *LookupIotHubDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupIotHubDataConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIotHubDataConnectionResult
	err := ctx.Invoke("azure-native:synapse:getIotHubDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotHubDataConnectionArgs struct {
	// The name of the data connection.
	DataConnectionName string `pulumi:"dataConnectionName"`
	// The name of the database in the Kusto pool.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the Kusto pool.
	KustoPoolName string `pulumi:"kustoPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Class representing an iot hub data connection.
type LookupIotHubDataConnectionResult struct {
	// The iot hub consumer group.
	ConsumerGroup string `pulumi:"consumerGroup"`
	// The data format of the message. Optionally the data format can be added to each message.
	DataFormat *string `pulumi:"dataFormat"`
	// System properties of the iot hub
	EventSystemProperties []string `pulumi:"eventSystemProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The resource ID of the Iot hub to be used to create a data connection.
	IotHubResourceId string `pulumi:"iotHubResourceId"`
	// Kind of the endpoint for the data connection
	// Expected value is 'IotHub'.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The name of the share access policy
	SharedAccessPolicyName string `pulumi:"sharedAccessPolicyName"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName *string `pulumi:"tableName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupIotHubDataConnectionOutput(ctx *pulumi.Context, args LookupIotHubDataConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupIotHubDataConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotHubDataConnectionResult, error) {
			args := v.(LookupIotHubDataConnectionArgs)
			r, err := LookupIotHubDataConnection(ctx, &args, opts...)
			var s LookupIotHubDataConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotHubDataConnectionResultOutput)
}

type LookupIotHubDataConnectionOutputArgs struct {
	// The name of the data connection.
	DataConnectionName pulumi.StringInput `pulumi:"dataConnectionName"`
	// The name of the database in the Kusto pool.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the Kusto pool.
	KustoPoolName pulumi.StringInput `pulumi:"kustoPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIotHubDataConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubDataConnectionArgs)(nil)).Elem()
}

// Class representing an iot hub data connection.
type LookupIotHubDataConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupIotHubDataConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubDataConnectionResult)(nil)).Elem()
}

func (o LookupIotHubDataConnectionResultOutput) ToLookupIotHubDataConnectionResultOutput() LookupIotHubDataConnectionResultOutput {
	return o
}

func (o LookupIotHubDataConnectionResultOutput) ToLookupIotHubDataConnectionResultOutputWithContext(ctx context.Context) LookupIotHubDataConnectionResultOutput {
	return o
}

// The iot hub consumer group.
func (o LookupIotHubDataConnectionResultOutput) ConsumerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.ConsumerGroup }).(pulumi.StringOutput)
}

// The data format of the message. Optionally the data format can be added to each message.
func (o LookupIotHubDataConnectionResultOutput) DataFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) *string { return v.DataFormat }).(pulumi.StringPtrOutput)
}

// System properties of the iot hub
func (o LookupIotHubDataConnectionResultOutput) EventSystemProperties() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) []string { return v.EventSystemProperties }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupIotHubDataConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource ID of the Iot hub to be used to create a data connection.
func (o LookupIotHubDataConnectionResultOutput) IotHubResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.IotHubResourceId }).(pulumi.StringOutput)
}

// Kind of the endpoint for the data connection
// Expected value is 'IotHub'.
func (o LookupIotHubDataConnectionResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupIotHubDataConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
func (o LookupIotHubDataConnectionResultOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) *string { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupIotHubDataConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o LookupIotHubDataConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The name of the share access policy
func (o LookupIotHubDataConnectionResultOutput) SharedAccessPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.SharedAccessPolicyName }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupIotHubDataConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The table where the data should be ingested. Optionally the table information can be added to each message.
func (o LookupIotHubDataConnectionResultOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) *string { return v.TableName }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupIotHubDataConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubDataConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotHubDataConnectionResultOutput{})
}
