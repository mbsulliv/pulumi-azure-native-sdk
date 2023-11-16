// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	// The rack definition that is intended to reflect only a single rack in a single rack cluster, or an aggregator rack in a multi-rack cluster.
	AggregatorOrSingleRackDefinition RackDefinitionResponseOutput `pulumi:"aggregatorOrSingleRackDefinition"`
	// The resource ID of the Log Analytics Workspace that will be used for storing relevant logs.
	AnalyticsWorkspaceId pulumi.StringPtrOutput `pulumi:"analyticsWorkspaceId"`
	// The list of cluster runtime version upgrades available for this cluster.
	AvailableUpgradeVersions ClusterAvailableUpgradeVersionResponseArrayOutput `pulumi:"availableUpgradeVersions"`
	// The capacity supported by this cluster.
	ClusterCapacity ClusterCapacityResponseOutput `pulumi:"clusterCapacity"`
	// The latest heartbeat status between the cluster manager and the cluster.
	ClusterConnectionStatus pulumi.StringOutput `pulumi:"clusterConnectionStatus"`
	// The extended location (custom location) that represents the cluster's control plane location. This extended location is used to route the requests of child objects of the cluster that are handled by the platform operator.
	ClusterExtendedLocation ExtendedLocationResponseOutput `pulumi:"clusterExtendedLocation"`
	// The customer-provided location information to identify where the cluster resides.
	ClusterLocation pulumi.StringPtrOutput `pulumi:"clusterLocation"`
	// The latest connectivity status between cluster manager and the cluster.
	ClusterManagerConnectionStatus pulumi.StringOutput `pulumi:"clusterManagerConnectionStatus"`
	// The resource ID of the cluster manager that manages this cluster. This is set by the Cluster Manager when the cluster is created.
	ClusterManagerId pulumi.StringOutput `pulumi:"clusterManagerId"`
	// The service principal to be used by the cluster during Arc Appliance installation.
	ClusterServicePrincipal ServicePrincipalInformationResponsePtrOutput `pulumi:"clusterServicePrincipal"`
	// The type of rack configuration for the cluster.
	ClusterType pulumi.StringOutput `pulumi:"clusterType"`
	// The current runtime version of the cluster.
	ClusterVersion pulumi.StringOutput `pulumi:"clusterVersion"`
	// The validation threshold indicating the allowable failures of compute machines during environment validation and deployment.
	ComputeDeploymentThreshold ValidationThresholdResponsePtrOutput `pulumi:"computeDeploymentThreshold"`
	// The list of rack definitions for the compute racks in a multi-rack
	// cluster, or an empty list in a single-rack cluster.
	ComputeRackDefinitions RackDefinitionResponseArrayOutput `pulumi:"computeRackDefinitions"`
	// The current detailed status of the cluster.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// The descriptive message about the detailed status.
	DetailedStatusMessage pulumi.StringOutput `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster manager associated with the cluster.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// Field Deprecated. This field will not be populated in an upcoming version. The extended location (custom location) that represents the Hybrid AKS control plane location. This extended location is used when creating provisioned clusters (Hybrid AKS clusters).
	HybridAksExtendedLocation ExtendedLocationResponseOutput `pulumi:"hybridAksExtendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The configuration of the managed resource group associated with the resource.
	ManagedResourceGroupConfiguration ManagedResourceGroupConfigurationResponsePtrOutput `pulumi:"managedResourceGroupConfiguration"`
	// The count of Manual Action Taken (MAT) events that have not been validated.
	ManualActionCount pulumi.Float64Output `pulumi:"manualActionCount"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource ID of the Network Fabric associated with the cluster.
	NetworkFabricId pulumi.StringOutput `pulumi:"networkFabricId"`
	// The provisioning state of the cluster.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The support end date of the runtime version of the cluster.
	SupportExpiryDate pulumi.StringOutput `pulumi:"supportExpiryDate"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of workload resource IDs that are hosted within this cluster.
	WorkloadResourceIds pulumi.StringArrayOutput `pulumi:"workloadResourceIds"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AggregatorOrSingleRackDefinition == nil {
		return nil, errors.New("invalid value for required argument 'AggregatorOrSingleRackDefinition'")
	}
	if args.ClusterType == nil {
		return nil, errors.New("invalid value for required argument 'ClusterType'")
	}
	if args.ClusterVersion == nil {
		return nil, errors.New("invalid value for required argument 'ClusterVersion'")
	}
	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.NetworkFabricId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkFabricId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkcloud:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20221212preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20230701:Cluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:networkcloud/v20230501preview:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:networkcloud/v20230501preview:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The rack definition that is intended to reflect only a single rack in a single rack cluster, or an aggregator rack in a multi-rack cluster.
	AggregatorOrSingleRackDefinition RackDefinition `pulumi:"aggregatorOrSingleRackDefinition"`
	// The resource ID of the Log Analytics Workspace that will be used for storing relevant logs.
	AnalyticsWorkspaceId *string `pulumi:"analyticsWorkspaceId"`
	// The customer-provided location information to identify where the cluster resides.
	ClusterLocation *string `pulumi:"clusterLocation"`
	// The name of the cluster.
	ClusterName *string `pulumi:"clusterName"`
	// The service principal to be used by the cluster during Arc Appliance installation.
	ClusterServicePrincipal *ServicePrincipalInformation `pulumi:"clusterServicePrincipal"`
	// The type of rack configuration for the cluster.
	ClusterType string `pulumi:"clusterType"`
	// The current runtime version of the cluster.
	ClusterVersion string `pulumi:"clusterVersion"`
	// The validation threshold indicating the allowable failures of compute machines during environment validation and deployment.
	ComputeDeploymentThreshold *ValidationThreshold `pulumi:"computeDeploymentThreshold"`
	// The list of rack definitions for the compute racks in a multi-rack
	// cluster, or an empty list in a single-rack cluster.
	ComputeRackDefinitions []RackDefinition `pulumi:"computeRackDefinitions"`
	// The extended location of the cluster manager associated with the cluster.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The configuration of the managed resource group associated with the resource.
	ManagedResourceGroupConfiguration *ManagedResourceGroupConfiguration `pulumi:"managedResourceGroupConfiguration"`
	// The resource ID of the Network Fabric associated with the cluster.
	NetworkFabricId string `pulumi:"networkFabricId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The rack definition that is intended to reflect only a single rack in a single rack cluster, or an aggregator rack in a multi-rack cluster.
	AggregatorOrSingleRackDefinition RackDefinitionInput
	// The resource ID of the Log Analytics Workspace that will be used for storing relevant logs.
	AnalyticsWorkspaceId pulumi.StringPtrInput
	// The customer-provided location information to identify where the cluster resides.
	ClusterLocation pulumi.StringPtrInput
	// The name of the cluster.
	ClusterName pulumi.StringPtrInput
	// The service principal to be used by the cluster during Arc Appliance installation.
	ClusterServicePrincipal ServicePrincipalInformationPtrInput
	// The type of rack configuration for the cluster.
	ClusterType pulumi.StringInput
	// The current runtime version of the cluster.
	ClusterVersion pulumi.StringInput
	// The validation threshold indicating the allowable failures of compute machines during environment validation and deployment.
	ComputeDeploymentThreshold ValidationThresholdPtrInput
	// The list of rack definitions for the compute racks in a multi-rack
	// cluster, or an empty list in a single-rack cluster.
	ComputeRackDefinitions RackDefinitionArrayInput
	// The extended location of the cluster manager associated with the cluster.
	ExtendedLocation ExtendedLocationInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The configuration of the managed resource group associated with the resource.
	ManagedResourceGroupConfiguration ManagedResourceGroupConfigurationPtrInput
	// The resource ID of the Network Fabric associated with the cluster.
	NetworkFabricId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

// The rack definition that is intended to reflect only a single rack in a single rack cluster, or an aggregator rack in a multi-rack cluster.
func (o ClusterOutput) AggregatorOrSingleRackDefinition() RackDefinitionResponseOutput {
	return o.ApplyT(func(v *Cluster) RackDefinitionResponseOutput { return v.AggregatorOrSingleRackDefinition }).(RackDefinitionResponseOutput)
}

// The resource ID of the Log Analytics Workspace that will be used for storing relevant logs.
func (o ClusterOutput) AnalyticsWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.AnalyticsWorkspaceId }).(pulumi.StringPtrOutput)
}

// The list of cluster runtime version upgrades available for this cluster.
func (o ClusterOutput) AvailableUpgradeVersions() ClusterAvailableUpgradeVersionResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterAvailableUpgradeVersionResponseArrayOutput { return v.AvailableUpgradeVersions }).(ClusterAvailableUpgradeVersionResponseArrayOutput)
}

// The capacity supported by this cluster.
func (o ClusterOutput) ClusterCapacity() ClusterCapacityResponseOutput {
	return o.ApplyT(func(v *Cluster) ClusterCapacityResponseOutput { return v.ClusterCapacity }).(ClusterCapacityResponseOutput)
}

// The latest heartbeat status between the cluster manager and the cluster.
func (o ClusterOutput) ClusterConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterConnectionStatus }).(pulumi.StringOutput)
}

// The extended location (custom location) that represents the cluster's control plane location. This extended location is used to route the requests of child objects of the cluster that are handled by the platform operator.
func (o ClusterOutput) ClusterExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *Cluster) ExtendedLocationResponseOutput { return v.ClusterExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The customer-provided location information to identify where the cluster resides.
func (o ClusterOutput) ClusterLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ClusterLocation }).(pulumi.StringPtrOutput)
}

// The latest connectivity status between cluster manager and the cluster.
func (o ClusterOutput) ClusterManagerConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterManagerConnectionStatus }).(pulumi.StringOutput)
}

// The resource ID of the cluster manager that manages this cluster. This is set by the Cluster Manager when the cluster is created.
func (o ClusterOutput) ClusterManagerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterManagerId }).(pulumi.StringOutput)
}

// The service principal to be used by the cluster during Arc Appliance installation.
func (o ClusterOutput) ClusterServicePrincipal() ServicePrincipalInformationResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ServicePrincipalInformationResponsePtrOutput { return v.ClusterServicePrincipal }).(ServicePrincipalInformationResponsePtrOutput)
}

// The type of rack configuration for the cluster.
func (o ClusterOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterType }).(pulumi.StringOutput)
}

// The current runtime version of the cluster.
func (o ClusterOutput) ClusterVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterVersion }).(pulumi.StringOutput)
}

// The validation threshold indicating the allowable failures of compute machines during environment validation and deployment.
func (o ClusterOutput) ComputeDeploymentThreshold() ValidationThresholdResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ValidationThresholdResponsePtrOutput { return v.ComputeDeploymentThreshold }).(ValidationThresholdResponsePtrOutput)
}

// The list of rack definitions for the compute racks in a multi-rack
// cluster, or an empty list in a single-rack cluster.
func (o ClusterOutput) ComputeRackDefinitions() RackDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) RackDefinitionResponseArrayOutput { return v.ComputeRackDefinitions }).(RackDefinitionResponseArrayOutput)
}

// The current detailed status of the cluster.
func (o ClusterOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the detailed status.
func (o ClusterOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster manager associated with the cluster.
func (o ClusterOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *Cluster) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Field Deprecated. This field will not be populated in an upcoming version. The extended location (custom location) that represents the Hybrid AKS control plane location. This extended location is used when creating provisioned clusters (Hybrid AKS clusters).
func (o ClusterOutput) HybridAksExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *Cluster) ExtendedLocationResponseOutput { return v.HybridAksExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The geo-location where the resource lives
func (o ClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The configuration of the managed resource group associated with the resource.
func (o ClusterOutput) ManagedResourceGroupConfiguration() ManagedResourceGroupConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ManagedResourceGroupConfigurationResponsePtrOutput {
		return v.ManagedResourceGroupConfiguration
	}).(ManagedResourceGroupConfigurationResponsePtrOutput)
}

// The count of Manual Action Taken (MAT) events that have not been validated.
func (o ClusterOutput) ManualActionCount() pulumi.Float64Output {
	return o.ApplyT(func(v *Cluster) pulumi.Float64Output { return v.ManualActionCount }).(pulumi.Float64Output)
}

// The name of the resource
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource ID of the Network Fabric associated with the cluster.
func (o ClusterOutput) NetworkFabricId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.NetworkFabricId }).(pulumi.StringOutput)
}

// The provisioning state of the cluster.
func (o ClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The support end date of the runtime version of the cluster.
func (o ClusterOutput) SupportExpiryDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.SupportExpiryDate }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Cluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The list of workload resource IDs that are hosted within this cluster.
func (o ClusterOutput) WorkloadResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.WorkloadResourceIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
