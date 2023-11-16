// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231004preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a KafkaTopicMapResource
func LookupKafkaConnectorTopicMap(ctx *pulumi.Context, args *LookupKafkaConnectorTopicMapArgs, opts ...pulumi.InvokeOption) (*LookupKafkaConnectorTopicMapResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKafkaConnectorTopicMapResult
	err := ctx.Invoke("azure-native:iotoperationsmq/v20231004preview:getKafkaConnectorTopicMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupKafkaConnectorTopicMapArgs struct {
	// Name of MQ kafkaConnector resource
	KafkaConnectorName string `pulumi:"kafkaConnectorName"`
	// Name of MQ resource
	MqName string `pulumi:"mqName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of MQ kafka/topicMap resource
	TopicMapName string `pulumi:"topicMapName"`
}

// MQ kafkaConnector/topicMap resource
type LookupKafkaConnectorTopicMapResult struct {
	// The batching settings for kafka messages.
	Batching *KafkaTopicMapBatchingResponse `pulumi:"batching"`
	// The compression to use for kafka messages.
	Compression *string `pulumi:"compression"`
	// The flag to copy Mqtt properties.
	CopyMqttProperties *string `pulumi:"copyMqttProperties"`
	// Extended Location
	ExtendedLocation ExtendedLocationPropertyResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kafkaConnector CRD it refers to.
	KafkaConnectorRef string `pulumi:"kafkaConnectorRef"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The partition to use for Kafka.
	PartitionKeyProperty *string `pulumi:"partitionKeyProperty"`
	// The partition strategy to use for Kafka.
	PartitionStrategy *string `pulumi:"partitionStrategy"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// The route details for Kafka connector.
	Routes []KafkaRoutesResponse `pulumi:"routes"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupKafkaConnectorTopicMapResult
func (val *LookupKafkaConnectorTopicMapResult) Defaults() *LookupKafkaConnectorTopicMapResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Batching = tmp.Batching.Defaults()

	if tmp.Compression == nil {
		compression_ := "none"
		tmp.Compression = &compression_
	}
	if tmp.PartitionStrategy == nil {
		partitionStrategy_ := "default"
		tmp.PartitionStrategy = &partitionStrategy_
	}
	return &tmp
}

func LookupKafkaConnectorTopicMapOutput(ctx *pulumi.Context, args LookupKafkaConnectorTopicMapOutputArgs, opts ...pulumi.InvokeOption) LookupKafkaConnectorTopicMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKafkaConnectorTopicMapResult, error) {
			args := v.(LookupKafkaConnectorTopicMapArgs)
			r, err := LookupKafkaConnectorTopicMap(ctx, &args, opts...)
			var s LookupKafkaConnectorTopicMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKafkaConnectorTopicMapResultOutput)
}

type LookupKafkaConnectorTopicMapOutputArgs struct {
	// Name of MQ kafkaConnector resource
	KafkaConnectorName pulumi.StringInput `pulumi:"kafkaConnectorName"`
	// Name of MQ resource
	MqName pulumi.StringInput `pulumi:"mqName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of MQ kafka/topicMap resource
	TopicMapName pulumi.StringInput `pulumi:"topicMapName"`
}

func (LookupKafkaConnectorTopicMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKafkaConnectorTopicMapArgs)(nil)).Elem()
}

// MQ kafkaConnector/topicMap resource
type LookupKafkaConnectorTopicMapResultOutput struct{ *pulumi.OutputState }

func (LookupKafkaConnectorTopicMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKafkaConnectorTopicMapResult)(nil)).Elem()
}

func (o LookupKafkaConnectorTopicMapResultOutput) ToLookupKafkaConnectorTopicMapResultOutput() LookupKafkaConnectorTopicMapResultOutput {
	return o
}

func (o LookupKafkaConnectorTopicMapResultOutput) ToLookupKafkaConnectorTopicMapResultOutputWithContext(ctx context.Context) LookupKafkaConnectorTopicMapResultOutput {
	return o
}

// The batching settings for kafka messages.
func (o LookupKafkaConnectorTopicMapResultOutput) Batching() KafkaTopicMapBatchingResponsePtrOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) *KafkaTopicMapBatchingResponse { return v.Batching }).(KafkaTopicMapBatchingResponsePtrOutput)
}

// The compression to use for kafka messages.
func (o LookupKafkaConnectorTopicMapResultOutput) Compression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) *string { return v.Compression }).(pulumi.StringPtrOutput)
}

// The flag to copy Mqtt properties.
func (o LookupKafkaConnectorTopicMapResultOutput) CopyMqttProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) *string { return v.CopyMqttProperties }).(pulumi.StringPtrOutput)
}

// Extended Location
func (o LookupKafkaConnectorTopicMapResultOutput) ExtendedLocation() ExtendedLocationPropertyResponseOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) ExtendedLocationPropertyResponse { return v.ExtendedLocation }).(ExtendedLocationPropertyResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupKafkaConnectorTopicMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kafkaConnector CRD it refers to.
func (o LookupKafkaConnectorTopicMapResultOutput) KafkaConnectorRef() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) string { return v.KafkaConnectorRef }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupKafkaConnectorTopicMapResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupKafkaConnectorTopicMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) string { return v.Name }).(pulumi.StringOutput)
}

// The partition to use for Kafka.
func (o LookupKafkaConnectorTopicMapResultOutput) PartitionKeyProperty() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) *string { return v.PartitionKeyProperty }).(pulumi.StringPtrOutput)
}

// The partition strategy to use for Kafka.
func (o LookupKafkaConnectorTopicMapResultOutput) PartitionStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) *string { return v.PartitionStrategy }).(pulumi.StringPtrOutput)
}

// The status of the last operation.
func (o LookupKafkaConnectorTopicMapResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The route details for Kafka connector.
func (o LookupKafkaConnectorTopicMapResultOutput) Routes() KafkaRoutesResponseArrayOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) []KafkaRoutesResponse { return v.Routes }).(KafkaRoutesResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupKafkaConnectorTopicMapResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupKafkaConnectorTopicMapResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupKafkaConnectorTopicMapResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaConnectorTopicMapResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKafkaConnectorTopicMapResultOutput{})
}
