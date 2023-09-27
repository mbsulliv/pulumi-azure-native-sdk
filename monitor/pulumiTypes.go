// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// The Data Collection Rule and Endpoint used for ingestion by default.
type AzureMonitorWorkspaceResponseDefaultIngestionSettings struct {
	// The Azure resource Id of the default data collection endpoint for this Azure Monitor Workspace.
	DataCollectionEndpointResourceId string `pulumi:"dataCollectionEndpointResourceId"`
	// The Azure resource Id of the default data collection rule for this Azure Monitor Workspace.
	DataCollectionRuleResourceId string `pulumi:"dataCollectionRuleResourceId"`
}

// The Data Collection Rule and Endpoint used for ingestion by default.
type AzureMonitorWorkspaceResponseDefaultIngestionSettingsOutput struct{ *pulumi.OutputState }

func (AzureMonitorWorkspaceResponseDefaultIngestionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorWorkspaceResponseDefaultIngestionSettings)(nil)).Elem()
}

func (o AzureMonitorWorkspaceResponseDefaultIngestionSettingsOutput) ToAzureMonitorWorkspaceResponseDefaultIngestionSettingsOutput() AzureMonitorWorkspaceResponseDefaultIngestionSettingsOutput {
	return o
}

func (o AzureMonitorWorkspaceResponseDefaultIngestionSettingsOutput) ToAzureMonitorWorkspaceResponseDefaultIngestionSettingsOutputWithContext(ctx context.Context) AzureMonitorWorkspaceResponseDefaultIngestionSettingsOutput {
	return o
}

func (o AzureMonitorWorkspaceResponseDefaultIngestionSettingsOutput) ToOutput(ctx context.Context) pulumix.Output[AzureMonitorWorkspaceResponseDefaultIngestionSettings] {
	return pulumix.Output[AzureMonitorWorkspaceResponseDefaultIngestionSettings]{
		OutputState: o.OutputState,
	}
}

// The Azure resource Id of the default data collection endpoint for this Azure Monitor Workspace.
func (o AzureMonitorWorkspaceResponseDefaultIngestionSettingsOutput) DataCollectionEndpointResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v AzureMonitorWorkspaceResponseDefaultIngestionSettings) string {
		return v.DataCollectionEndpointResourceId
	}).(pulumi.StringOutput)
}

// The Azure resource Id of the default data collection rule for this Azure Monitor Workspace.
func (o AzureMonitorWorkspaceResponseDefaultIngestionSettingsOutput) DataCollectionRuleResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v AzureMonitorWorkspaceResponseDefaultIngestionSettings) string {
		return v.DataCollectionRuleResourceId
	}).(pulumi.StringOutput)
}

// Properties related to the metrics container in the Azure Monitor Workspace
type AzureMonitorWorkspaceResponseMetrics struct {
	// An internal identifier for the metrics container. Only to be used by the system
	InternalId string `pulumi:"internalId"`
	// The Prometheus query endpoint for the Azure Monitor Workspace
	PrometheusQueryEndpoint string `pulumi:"prometheusQueryEndpoint"`
}

// Properties related to the metrics container in the Azure Monitor Workspace
type AzureMonitorWorkspaceResponseMetricsOutput struct{ *pulumi.OutputState }

func (AzureMonitorWorkspaceResponseMetricsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorWorkspaceResponseMetrics)(nil)).Elem()
}

func (o AzureMonitorWorkspaceResponseMetricsOutput) ToAzureMonitorWorkspaceResponseMetricsOutput() AzureMonitorWorkspaceResponseMetricsOutput {
	return o
}

func (o AzureMonitorWorkspaceResponseMetricsOutput) ToAzureMonitorWorkspaceResponseMetricsOutputWithContext(ctx context.Context) AzureMonitorWorkspaceResponseMetricsOutput {
	return o
}

func (o AzureMonitorWorkspaceResponseMetricsOutput) ToOutput(ctx context.Context) pulumix.Output[AzureMonitorWorkspaceResponseMetrics] {
	return pulumix.Output[AzureMonitorWorkspaceResponseMetrics]{
		OutputState: o.OutputState,
	}
}

// An internal identifier for the metrics container. Only to be used by the system
func (o AzureMonitorWorkspaceResponseMetricsOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v AzureMonitorWorkspaceResponseMetrics) string { return v.InternalId }).(pulumi.StringOutput)
}

// The Prometheus query endpoint for the Azure Monitor Workspace
func (o AzureMonitorWorkspaceResponseMetricsOutput) PrometheusQueryEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v AzureMonitorWorkspaceResponseMetrics) string { return v.PrometheusQueryEndpoint }).(pulumi.StringOutput)
}

// The private endpoint connection resource.
type PrivateEndpointConnectionResponse struct {
	// The group ids for the private endpoint resource.
	GroupIds []string `pulumi:"groupIds"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The private endpoint resource.
	PrivateEndpoint *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// The private endpoint connection resource.
type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToOutput(ctx context.Context) pulumix.Output[PrivateEndpointConnectionResponse] {
	return pulumix.Output[PrivateEndpointConnectionResponse]{
		OutputState: o.OutputState,
	}
}

// The group ids for the private endpoint resource.
func (o PrivateEndpointConnectionResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The private endpoint resource.
func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]PrivateEndpointConnectionResponse] {
	return pulumix.Output[[]PrivateEndpointConnectionResponse]{
		OutputState: o.OutputState,
	}
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

// The private endpoint resource.
type PrivateEndpointResponse struct {
	// The ARM identifier for private endpoint.
	Id string `pulumi:"id"`
}

// The private endpoint resource.
type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToOutput(ctx context.Context) pulumix.Output[PrivateEndpointResponse] {
	return pulumix.Output[PrivateEndpointResponse]{
		OutputState: o.OutputState,
	}
}

// The ARM identifier for private endpoint.
func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpointResponse] {
	return pulumix.Output[*PrivateEndpointResponse]{
		OutputState: o.OutputState,
	}
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

// The ARM identifier for private endpoint.
func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
type PrivateLinkServiceConnectionStateResponse struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string `pulumi:"actionsRequired"`
	// The reason for approval/rejection of the connection.
	Description *string `pulumi:"description"`
	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *string `pulumi:"status"`
}

// A collection of information about the state of the connection between service consumer and provider.
type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToOutput(ctx context.Context) pulumix.Output[PrivateLinkServiceConnectionStateResponse] {
	return pulumix.Output[PrivateLinkServiceConnectionStateResponse]{
		OutputState: o.OutputState,
	}
}

// A message indicating if changes on the service provider require any updates on the consumer.
func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

// The reason for approval/rejection of the connection.
func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToOutput(ctx context.Context) pulumix.Output[SystemDataResponse] {
	return pulumix.Output[SystemDataResponse]{
		OutputState: o.OutputState,
	}
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureMonitorWorkspaceResponseDefaultIngestionSettingsOutput{})
	pulumi.RegisterOutputType(AzureMonitorWorkspaceResponseMetricsOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
