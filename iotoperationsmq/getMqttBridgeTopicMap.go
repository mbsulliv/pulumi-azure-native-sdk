// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotoperationsmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a MqttBridgeTopicMapResource
// Azure REST API version: 2023-10-04-preview.
func LookupMqttBridgeTopicMap(ctx *pulumi.Context, args *LookupMqttBridgeTopicMapArgs, opts ...pulumi.InvokeOption) (*LookupMqttBridgeTopicMapResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMqttBridgeTopicMapResult
	err := ctx.Invoke("azure-native:iotoperationsmq:getMqttBridgeTopicMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMqttBridgeTopicMapArgs struct {
	// Name of MQ resource
	MqName string `pulumi:"mqName"`
	// Name of MQ mqttBridgeConnector resource
	MqttBridgeConnectorName string `pulumi:"mqttBridgeConnectorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of MQ mqttBridgeTopicMap resource
	TopicMapName string `pulumi:"topicMapName"`
}

// MQ mqttBridgeTopicMap resource
type LookupMqttBridgeTopicMapResult struct {
	// Extended Location
	ExtendedLocation ExtendedLocationPropertyResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The MqttBridgeConnector CRD it refers to.
	MqttBridgeConnectorRef string `pulumi:"mqttBridgeConnectorRef"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// The route details for MqttBridge connector.
	Routes []MqttBridgeRoutesResponse `pulumi:"routes"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupMqttBridgeTopicMapOutput(ctx *pulumi.Context, args LookupMqttBridgeTopicMapOutputArgs, opts ...pulumi.InvokeOption) LookupMqttBridgeTopicMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMqttBridgeTopicMapResult, error) {
			args := v.(LookupMqttBridgeTopicMapArgs)
			r, err := LookupMqttBridgeTopicMap(ctx, &args, opts...)
			var s LookupMqttBridgeTopicMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMqttBridgeTopicMapResultOutput)
}

type LookupMqttBridgeTopicMapOutputArgs struct {
	// Name of MQ resource
	MqName pulumi.StringInput `pulumi:"mqName"`
	// Name of MQ mqttBridgeConnector resource
	MqttBridgeConnectorName pulumi.StringInput `pulumi:"mqttBridgeConnectorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of MQ mqttBridgeTopicMap resource
	TopicMapName pulumi.StringInput `pulumi:"topicMapName"`
}

func (LookupMqttBridgeTopicMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMqttBridgeTopicMapArgs)(nil)).Elem()
}

// MQ mqttBridgeTopicMap resource
type LookupMqttBridgeTopicMapResultOutput struct{ *pulumi.OutputState }

func (LookupMqttBridgeTopicMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMqttBridgeTopicMapResult)(nil)).Elem()
}

func (o LookupMqttBridgeTopicMapResultOutput) ToLookupMqttBridgeTopicMapResultOutput() LookupMqttBridgeTopicMapResultOutput {
	return o
}

func (o LookupMqttBridgeTopicMapResultOutput) ToLookupMqttBridgeTopicMapResultOutputWithContext(ctx context.Context) LookupMqttBridgeTopicMapResultOutput {
	return o
}

// Extended Location
func (o LookupMqttBridgeTopicMapResultOutput) ExtendedLocation() ExtendedLocationPropertyResponseOutput {
	return o.ApplyT(func(v LookupMqttBridgeTopicMapResult) ExtendedLocationPropertyResponse { return v.ExtendedLocation }).(ExtendedLocationPropertyResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupMqttBridgeTopicMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMqttBridgeTopicMapResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupMqttBridgeTopicMapResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMqttBridgeTopicMapResult) string { return v.Location }).(pulumi.StringOutput)
}

// The MqttBridgeConnector CRD it refers to.
func (o LookupMqttBridgeTopicMapResultOutput) MqttBridgeConnectorRef() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMqttBridgeTopicMapResult) string { return v.MqttBridgeConnectorRef }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMqttBridgeTopicMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMqttBridgeTopicMapResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupMqttBridgeTopicMapResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMqttBridgeTopicMapResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The route details for MqttBridge connector.
func (o LookupMqttBridgeTopicMapResultOutput) Routes() MqttBridgeRoutesResponseArrayOutput {
	return o.ApplyT(func(v LookupMqttBridgeTopicMapResult) []MqttBridgeRoutesResponse { return v.Routes }).(MqttBridgeRoutesResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMqttBridgeTopicMapResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMqttBridgeTopicMapResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupMqttBridgeTopicMapResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMqttBridgeTopicMapResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMqttBridgeTopicMapResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMqttBridgeTopicMapResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMqttBridgeTopicMapResultOutput{})
}
