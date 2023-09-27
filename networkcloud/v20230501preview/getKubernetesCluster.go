// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get properties of the provided the Kubernetes cluster.
func LookupKubernetesCluster(ctx *pulumi.Context, args *LookupKubernetesClusterArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKubernetesClusterResult
	err := ctx.Invoke("azure-native:networkcloud/v20230501preview:getKubernetesCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupKubernetesClusterArgs struct {
	// The name of the Kubernetes cluster.
	KubernetesClusterName string `pulumi:"kubernetesClusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupKubernetesClusterResult struct {
	// The Azure Active Directory Integration properties.
	AadConfiguration *AadConfigurationResponse `pulumi:"aadConfiguration"`
	// The administrative credentials that will be applied to the control plane and agent pool nodes that do not specify their own values.
	AdministratorConfiguration *AdministratorConfigurationResponse `pulumi:"administratorConfiguration"`
	// The full list of network resource IDs that are attached to this cluster, including those attached only to specific agent pools.
	AttachedNetworkIds []string `pulumi:"attachedNetworkIds"`
	// The list of versions that this Kubernetes cluster can be upgraded to.
	AvailableUpgrades []AvailableUpgradeResponse `pulumi:"availableUpgrades"`
	// The resource ID of the Network Cloud cluster.
	ClusterId string `pulumi:"clusterId"`
	// The resource ID of the connected cluster set up when this Kubernetes cluster is created.
	ConnectedClusterId string `pulumi:"connectedClusterId"`
	// The current running version of Kubernetes on the control plane.
	ControlPlaneKubernetesVersion string `pulumi:"controlPlaneKubernetesVersion"`
	// The defining characteristics of the control plane for this Kubernetes Cluster.
	ControlPlaneNodeConfiguration ControlPlaneNodeConfigurationResponse `pulumi:"controlPlaneNodeConfiguration"`
	// The current status of the Kubernetes cluster.
	DetailedStatus string `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage string `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// The current feature settings.
	FeatureStatuses []FeatureStatusResponse `pulumi:"featureStatuses"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The agent pools that are created with this Kubernetes cluster for running critical system services and workloads. This data in this field is only used during creation, and the field will be empty following the creation of the Kubernetes Cluster. After creation, the management of agent pools is done using the agentPools sub-resource.
	InitialAgentPoolConfigurations []InitialAgentPoolConfigurationResponse `pulumi:"initialAgentPoolConfigurations"`
	// The Kubernetes version for this cluster. Accepts n.n, n.n.n, and n.n.n-n format. The interpreted version used will be resolved into this field after creation or update.
	KubernetesVersion string `pulumi:"kubernetesVersion"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The configuration of the managed resource group associated with the resource.
	ManagedResourceGroupConfiguration *ManagedResourceGroupConfigurationResponse `pulumi:"managedResourceGroupConfiguration"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The configuration of the Kubernetes cluster networking, including the attachment of networks that span the cluster.
	NetworkConfiguration NetworkConfigurationResponse `pulumi:"networkConfiguration"`
	// The details of the nodes in this cluster.
	Nodes []KubernetesClusterNodeResponse `pulumi:"nodes"`
	// The provisioning state of the Kubernetes cluster resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupKubernetesClusterResult
func (val *LookupKubernetesClusterResult) Defaults() *LookupKubernetesClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NetworkConfiguration = *tmp.NetworkConfiguration.Defaults()

	return &tmp
}

func LookupKubernetesClusterOutput(ctx *pulumi.Context, args LookupKubernetesClusterOutputArgs, opts ...pulumi.InvokeOption) LookupKubernetesClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKubernetesClusterResult, error) {
			args := v.(LookupKubernetesClusterArgs)
			r, err := LookupKubernetesCluster(ctx, &args, opts...)
			var s LookupKubernetesClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKubernetesClusterResultOutput)
}

type LookupKubernetesClusterOutputArgs struct {
	// The name of the Kubernetes cluster.
	KubernetesClusterName pulumi.StringInput `pulumi:"kubernetesClusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupKubernetesClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesClusterArgs)(nil)).Elem()
}

type LookupKubernetesClusterResultOutput struct{ *pulumi.OutputState }

func (LookupKubernetesClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesClusterResult)(nil)).Elem()
}

func (o LookupKubernetesClusterResultOutput) ToLookupKubernetesClusterResultOutput() LookupKubernetesClusterResultOutput {
	return o
}

func (o LookupKubernetesClusterResultOutput) ToLookupKubernetesClusterResultOutputWithContext(ctx context.Context) LookupKubernetesClusterResultOutput {
	return o
}

func (o LookupKubernetesClusterResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupKubernetesClusterResult] {
	return pulumix.Output[LookupKubernetesClusterResult]{
		OutputState: o.OutputState,
	}
}

// The Azure Active Directory Integration properties.
func (o LookupKubernetesClusterResultOutput) AadConfiguration() AadConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) *AadConfigurationResponse { return v.AadConfiguration }).(AadConfigurationResponsePtrOutput)
}

// The administrative credentials that will be applied to the control plane and agent pool nodes that do not specify their own values.
func (o LookupKubernetesClusterResultOutput) AdministratorConfiguration() AdministratorConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) *AdministratorConfigurationResponse {
		return v.AdministratorConfiguration
	}).(AdministratorConfigurationResponsePtrOutput)
}

// The full list of network resource IDs that are attached to this cluster, including those attached only to specific agent pools.
func (o LookupKubernetesClusterResultOutput) AttachedNetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) []string { return v.AttachedNetworkIds }).(pulumi.StringArrayOutput)
}

// The list of versions that this Kubernetes cluster can be upgraded to.
func (o LookupKubernetesClusterResultOutput) AvailableUpgrades() AvailableUpgradeResponseArrayOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) []AvailableUpgradeResponse { return v.AvailableUpgrades }).(AvailableUpgradeResponseArrayOutput)
}

// The resource ID of the Network Cloud cluster.
func (o LookupKubernetesClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The resource ID of the connected cluster set up when this Kubernetes cluster is created.
func (o LookupKubernetesClusterResultOutput) ConnectedClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.ConnectedClusterId }).(pulumi.StringOutput)
}

// The current running version of Kubernetes on the control plane.
func (o LookupKubernetesClusterResultOutput) ControlPlaneKubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.ControlPlaneKubernetesVersion }).(pulumi.StringOutput)
}

// The defining characteristics of the control plane for this Kubernetes Cluster.
func (o LookupKubernetesClusterResultOutput) ControlPlaneNodeConfiguration() ControlPlaneNodeConfigurationResponseOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) ControlPlaneNodeConfigurationResponse {
		return v.ControlPlaneNodeConfiguration
	}).(ControlPlaneNodeConfigurationResponseOutput)
}

// The current status of the Kubernetes cluster.
func (o LookupKubernetesClusterResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o LookupKubernetesClusterResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o LookupKubernetesClusterResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The current feature settings.
func (o LookupKubernetesClusterResultOutput) FeatureStatuses() FeatureStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) []FeatureStatusResponse { return v.FeatureStatuses }).(FeatureStatusResponseArrayOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupKubernetesClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The agent pools that are created with this Kubernetes cluster for running critical system services and workloads. This data in this field is only used during creation, and the field will be empty following the creation of the Kubernetes Cluster. After creation, the management of agent pools is done using the agentPools sub-resource.
func (o LookupKubernetesClusterResultOutput) InitialAgentPoolConfigurations() InitialAgentPoolConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) []InitialAgentPoolConfigurationResponse {
		return v.InitialAgentPoolConfigurations
	}).(InitialAgentPoolConfigurationResponseArrayOutput)
}

// The Kubernetes version for this cluster. Accepts n.n, n.n.n, and n.n.n-n format. The interpreted version used will be resolved into this field after creation or update.
func (o LookupKubernetesClusterResultOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.KubernetesVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupKubernetesClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// The configuration of the managed resource group associated with the resource.
func (o LookupKubernetesClusterResultOutput) ManagedResourceGroupConfiguration() ManagedResourceGroupConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) *ManagedResourceGroupConfigurationResponse {
		return v.ManagedResourceGroupConfiguration
	}).(ManagedResourceGroupConfigurationResponsePtrOutput)
}

// The name of the resource
func (o LookupKubernetesClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// The configuration of the Kubernetes cluster networking, including the attachment of networks that span the cluster.
func (o LookupKubernetesClusterResultOutput) NetworkConfiguration() NetworkConfigurationResponseOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) NetworkConfigurationResponse { return v.NetworkConfiguration }).(NetworkConfigurationResponseOutput)
}

// The details of the nodes in this cluster.
func (o LookupKubernetesClusterResultOutput) Nodes() KubernetesClusterNodeResponseArrayOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) []KubernetesClusterNodeResponse { return v.Nodes }).(KubernetesClusterNodeResponseArrayOutput)
}

// The provisioning state of the Kubernetes cluster resource.
func (o LookupKubernetesClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupKubernetesClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupKubernetesClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupKubernetesClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKubernetesClusterResultOutput{})
}
