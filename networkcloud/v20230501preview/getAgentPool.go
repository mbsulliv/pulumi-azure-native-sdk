// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of the provided Kubernetes cluster agent pool.
func LookupAgentPool(ctx *pulumi.Context, args *LookupAgentPoolArgs, opts ...pulumi.InvokeOption) (*LookupAgentPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAgentPoolResult
	err := ctx.Invoke("azure-native:networkcloud/v20230501preview:getAgentPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAgentPoolArgs struct {
	// The name of the Kubernetes cluster agent pool.
	AgentPoolName string `pulumi:"agentPoolName"`
	// The name of the Kubernetes cluster.
	KubernetesClusterName string `pulumi:"kubernetesClusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupAgentPoolResult struct {
	// The administrator credentials to be used for the nodes in this agent pool.
	AdministratorConfiguration *AdministratorConfigurationResponse `pulumi:"administratorConfiguration"`
	// The configurations that will be applied to each agent in this agent pool.
	AgentOptions *AgentOptionsResponse `pulumi:"agentOptions"`
	// The configuration of networks being attached to the agent pool for use by the workloads that run on this Kubernetes cluster.
	AttachedNetworkConfiguration *AttachedNetworkConfigurationResponse `pulumi:"attachedNetworkConfiguration"`
	// The list of availability zones of the Network Cloud cluster used for the provisioning of nodes in this agent pool. If not specified, all availability zones will be used.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The number of virtual machines that use this configuration.
	Count float64 `pulumi:"count"`
	// The current status of the agent pool.
	DetailedStatus string `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage string `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The Kubernetes version running in this agent pool.
	KubernetesVersion string `pulumi:"kubernetesVersion"`
	// The labels applied to the nodes in this agent pool.
	Labels []KubernetesLabelResponse `pulumi:"labels"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The selection of how this agent pool is utilized, either as a system pool or a user pool. System pools run the features and critical services for the Kubernetes Cluster, while user pools are dedicated to user workloads. Every Kubernetes cluster must contain at least one system node pool with at least one node.
	Mode string `pulumi:"mode"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the agent pool.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The taints applied to the nodes in this agent pool.
	Taints []KubernetesLabelResponse `pulumi:"taints"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The configuration of the agent pool.
	UpgradeSettings *AgentPoolUpgradeSettingsResponse `pulumi:"upgradeSettings"`
	// The name of the VM SKU that determines the size of resources allocated for node VMs.
	VmSkuName string `pulumi:"vmSkuName"`
}

// Defaults sets the appropriate defaults for LookupAgentPoolResult
func (val *LookupAgentPoolResult) Defaults() *LookupAgentPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.AgentOptions = tmp.AgentOptions.Defaults()

	tmp.UpgradeSettings = tmp.UpgradeSettings.Defaults()

	return &tmp
}

func LookupAgentPoolOutput(ctx *pulumi.Context, args LookupAgentPoolOutputArgs, opts ...pulumi.InvokeOption) LookupAgentPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAgentPoolResult, error) {
			args := v.(LookupAgentPoolArgs)
			r, err := LookupAgentPool(ctx, &args, opts...)
			var s LookupAgentPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAgentPoolResultOutput)
}

type LookupAgentPoolOutputArgs struct {
	// The name of the Kubernetes cluster agent pool.
	AgentPoolName pulumi.StringInput `pulumi:"agentPoolName"`
	// The name of the Kubernetes cluster.
	KubernetesClusterName pulumi.StringInput `pulumi:"kubernetesClusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAgentPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentPoolArgs)(nil)).Elem()
}

type LookupAgentPoolResultOutput struct{ *pulumi.OutputState }

func (LookupAgentPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentPoolResult)(nil)).Elem()
}

func (o LookupAgentPoolResultOutput) ToLookupAgentPoolResultOutput() LookupAgentPoolResultOutput {
	return o
}

func (o LookupAgentPoolResultOutput) ToLookupAgentPoolResultOutputWithContext(ctx context.Context) LookupAgentPoolResultOutput {
	return o
}

// The administrator credentials to be used for the nodes in this agent pool.
func (o LookupAgentPoolResultOutput) AdministratorConfiguration() AdministratorConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *AdministratorConfigurationResponse { return v.AdministratorConfiguration }).(AdministratorConfigurationResponsePtrOutput)
}

// The configurations that will be applied to each agent in this agent pool.
func (o LookupAgentPoolResultOutput) AgentOptions() AgentOptionsResponsePtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *AgentOptionsResponse { return v.AgentOptions }).(AgentOptionsResponsePtrOutput)
}

// The configuration of networks being attached to the agent pool for use by the workloads that run on this Kubernetes cluster.
func (o LookupAgentPoolResultOutput) AttachedNetworkConfiguration() AttachedNetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *AttachedNetworkConfigurationResponse {
		return v.AttachedNetworkConfiguration
	}).(AttachedNetworkConfigurationResponsePtrOutput)
}

// The list of availability zones of the Network Cloud cluster used for the provisioning of nodes in this agent pool. If not specified, all availability zones will be used.
func (o LookupAgentPoolResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// The number of virtual machines that use this configuration.
func (o LookupAgentPoolResultOutput) Count() pulumi.Float64Output {
	return o.ApplyT(func(v LookupAgentPoolResult) float64 { return v.Count }).(pulumi.Float64Output)
}

// The current status of the agent pool.
func (o LookupAgentPoolResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o LookupAgentPoolResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o LookupAgentPoolResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupAgentPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Kubernetes version running in this agent pool.
func (o LookupAgentPoolResultOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.KubernetesVersion }).(pulumi.StringOutput)
}

// The labels applied to the nodes in this agent pool.
func (o LookupAgentPoolResultOutput) Labels() KubernetesLabelResponseArrayOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) []KubernetesLabelResponse { return v.Labels }).(KubernetesLabelResponseArrayOutput)
}

// The geo-location where the resource lives
func (o LookupAgentPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

// The selection of how this agent pool is utilized, either as a system pool or a user pool. System pools run the features and critical services for the Kubernetes Cluster, while user pools are dedicated to user workloads. Every Kubernetes cluster must contain at least one system node pool with at least one node.
func (o LookupAgentPoolResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Mode }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAgentPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the agent pool.
func (o LookupAgentPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAgentPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAgentPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The taints applied to the nodes in this agent pool.
func (o LookupAgentPoolResultOutput) Taints() KubernetesLabelResponseArrayOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) []KubernetesLabelResponse { return v.Taints }).(KubernetesLabelResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAgentPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

// The configuration of the agent pool.
func (o LookupAgentPoolResultOutput) UpgradeSettings() AgentPoolUpgradeSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *AgentPoolUpgradeSettingsResponse { return v.UpgradeSettings }).(AgentPoolUpgradeSettingsResponsePtrOutput)
}

// The name of the VM SKU that determines the size of resources allocated for node VMs.
func (o LookupAgentPoolResultOutput) VmSkuName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.VmSkuName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAgentPoolResultOutput{})
}
