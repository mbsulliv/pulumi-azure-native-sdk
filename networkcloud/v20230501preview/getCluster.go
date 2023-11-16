// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of the provided cluster.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:networkcloud/v20230501preview:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupClusterResult struct {
	// The rack definition that is intended to reflect only a single rack in a single rack cluster, or an aggregator rack in a multi-rack cluster.
	AggregatorOrSingleRackDefinition RackDefinitionResponse `pulumi:"aggregatorOrSingleRackDefinition"`
	// The resource ID of the Log Analytics Workspace that will be used for storing relevant logs.
	AnalyticsWorkspaceId *string `pulumi:"analyticsWorkspaceId"`
	// The list of cluster runtime version upgrades available for this cluster.
	AvailableUpgradeVersions []ClusterAvailableUpgradeVersionResponse `pulumi:"availableUpgradeVersions"`
	// The capacity supported by this cluster.
	ClusterCapacity ClusterCapacityResponse `pulumi:"clusterCapacity"`
	// The latest heartbeat status between the cluster manager and the cluster.
	ClusterConnectionStatus string `pulumi:"clusterConnectionStatus"`
	// The extended location (custom location) that represents the cluster's control plane location. This extended location is used to route the requests of child objects of the cluster that are handled by the platform operator.
	ClusterExtendedLocation ExtendedLocationResponse `pulumi:"clusterExtendedLocation"`
	// The customer-provided location information to identify where the cluster resides.
	ClusterLocation *string `pulumi:"clusterLocation"`
	// The latest connectivity status between cluster manager and the cluster.
	ClusterManagerConnectionStatus string `pulumi:"clusterManagerConnectionStatus"`
	// The resource ID of the cluster manager that manages this cluster. This is set by the Cluster Manager when the cluster is created.
	ClusterManagerId string `pulumi:"clusterManagerId"`
	// The service principal to be used by the cluster during Arc Appliance installation.
	ClusterServicePrincipal *ServicePrincipalInformationResponse `pulumi:"clusterServicePrincipal"`
	// The type of rack configuration for the cluster.
	ClusterType string `pulumi:"clusterType"`
	// The current runtime version of the cluster.
	ClusterVersion string `pulumi:"clusterVersion"`
	// The validation threshold indicating the allowable failures of compute machines during environment validation and deployment.
	ComputeDeploymentThreshold *ValidationThresholdResponse `pulumi:"computeDeploymentThreshold"`
	// The list of rack definitions for the compute racks in a multi-rack
	// cluster, or an empty list in a single-rack cluster.
	ComputeRackDefinitions []RackDefinitionResponse `pulumi:"computeRackDefinitions"`
	// The current detailed status of the cluster.
	DetailedStatus string `pulumi:"detailedStatus"`
	// The descriptive message about the detailed status.
	DetailedStatusMessage string `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster manager associated with the cluster.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Field Deprecated. This field will not be populated in an upcoming version. The extended location (custom location) that represents the Hybrid AKS control plane location. This extended location is used when creating provisioned clusters (Hybrid AKS clusters).
	HybridAksExtendedLocation ExtendedLocationResponse `pulumi:"hybridAksExtendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The configuration of the managed resource group associated with the resource.
	ManagedResourceGroupConfiguration *ManagedResourceGroupConfigurationResponse `pulumi:"managedResourceGroupConfiguration"`
	// The count of Manual Action Taken (MAT) events that have not been validated.
	ManualActionCount float64 `pulumi:"manualActionCount"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource ID of the Network Fabric associated with the cluster.
	NetworkFabricId string `pulumi:"networkFabricId"`
	// The provisioning state of the cluster.
	ProvisioningState string `pulumi:"provisioningState"`
	// The support end date of the runtime version of the cluster.
	SupportExpiryDate string `pulumi:"supportExpiryDate"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The list of workload resource IDs that are hosted within this cluster.
	WorkloadResourceIds []string `pulumi:"workloadResourceIds"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	// The name of the cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

// The rack definition that is intended to reflect only a single rack in a single rack cluster, or an aggregator rack in a multi-rack cluster.
func (o LookupClusterResultOutput) AggregatorOrSingleRackDefinition() RackDefinitionResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) RackDefinitionResponse { return v.AggregatorOrSingleRackDefinition }).(RackDefinitionResponseOutput)
}

// The resource ID of the Log Analytics Workspace that will be used for storing relevant logs.
func (o LookupClusterResultOutput) AnalyticsWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.AnalyticsWorkspaceId }).(pulumi.StringPtrOutput)
}

// The list of cluster runtime version upgrades available for this cluster.
func (o LookupClusterResultOutput) AvailableUpgradeVersions() ClusterAvailableUpgradeVersionResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []ClusterAvailableUpgradeVersionResponse {
		return v.AvailableUpgradeVersions
	}).(ClusterAvailableUpgradeVersionResponseArrayOutput)
}

// The capacity supported by this cluster.
func (o LookupClusterResultOutput) ClusterCapacity() ClusterCapacityResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) ClusterCapacityResponse { return v.ClusterCapacity }).(ClusterCapacityResponseOutput)
}

// The latest heartbeat status between the cluster manager and the cluster.
func (o LookupClusterResultOutput) ClusterConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterConnectionStatus }).(pulumi.StringOutput)
}

// The extended location (custom location) that represents the cluster's control plane location. This extended location is used to route the requests of child objects of the cluster that are handled by the platform operator.
func (o LookupClusterResultOutput) ClusterExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) ExtendedLocationResponse { return v.ClusterExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The customer-provided location information to identify where the cluster resides.
func (o LookupClusterResultOutput) ClusterLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ClusterLocation }).(pulumi.StringPtrOutput)
}

// The latest connectivity status between cluster manager and the cluster.
func (o LookupClusterResultOutput) ClusterManagerConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterManagerConnectionStatus }).(pulumi.StringOutput)
}

// The resource ID of the cluster manager that manages this cluster. This is set by the Cluster Manager when the cluster is created.
func (o LookupClusterResultOutput) ClusterManagerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterManagerId }).(pulumi.StringOutput)
}

// The service principal to be used by the cluster during Arc Appliance installation.
func (o LookupClusterResultOutput) ClusterServicePrincipal() ServicePrincipalInformationResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ServicePrincipalInformationResponse { return v.ClusterServicePrincipal }).(ServicePrincipalInformationResponsePtrOutput)
}

// The type of rack configuration for the cluster.
func (o LookupClusterResultOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterType }).(pulumi.StringOutput)
}

// The current runtime version of the cluster.
func (o LookupClusterResultOutput) ClusterVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterVersion }).(pulumi.StringOutput)
}

// The validation threshold indicating the allowable failures of compute machines during environment validation and deployment.
func (o LookupClusterResultOutput) ComputeDeploymentThreshold() ValidationThresholdResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ValidationThresholdResponse { return v.ComputeDeploymentThreshold }).(ValidationThresholdResponsePtrOutput)
}

// The list of rack definitions for the compute racks in a multi-rack
// cluster, or an empty list in a single-rack cluster.
func (o LookupClusterResultOutput) ComputeRackDefinitions() RackDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []RackDefinitionResponse { return v.ComputeRackDefinitions }).(RackDefinitionResponseArrayOutput)
}

// The current detailed status of the cluster.
func (o LookupClusterResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the detailed status.
func (o LookupClusterResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster manager associated with the cluster.
func (o LookupClusterResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Field Deprecated. This field will not be populated in an upcoming version. The extended location (custom location) that represents the Hybrid AKS control plane location. This extended location is used when creating provisioned clusters (Hybrid AKS clusters).
func (o LookupClusterResultOutput) HybridAksExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) ExtendedLocationResponse { return v.HybridAksExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// The configuration of the managed resource group associated with the resource.
func (o LookupClusterResultOutput) ManagedResourceGroupConfiguration() ManagedResourceGroupConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ManagedResourceGroupConfigurationResponse {
		return v.ManagedResourceGroupConfiguration
	}).(ManagedResourceGroupConfigurationResponsePtrOutput)
}

// The count of Manual Action Taken (MAT) events that have not been validated.
func (o LookupClusterResultOutput) ManualActionCount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupClusterResult) float64 { return v.ManualActionCount }).(pulumi.Float64Output)
}

// The name of the resource
func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource ID of the Network Fabric associated with the cluster.
func (o LookupClusterResultOutput) NetworkFabricId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.NetworkFabricId }).(pulumi.StringOutput)
}

// The provisioning state of the cluster.
func (o LookupClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The support end date of the runtime version of the cluster.
func (o LookupClusterResultOutput) SupportExpiryDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.SupportExpiryDate }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

// The list of workload resource IDs that are hosted within this cluster.
func (o LookupClusterResultOutput) WorkloadResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.WorkloadResourceIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
