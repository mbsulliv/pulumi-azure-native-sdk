// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Agent Pool.
func LookupAgentPool(ctx *pulumi.Context, args *LookupAgentPoolArgs, opts ...pulumi.InvokeOption) (*LookupAgentPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAgentPoolResult
	err := ctx.Invoke("azure-native:containerservice/v20210801:getAgentPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAgentPoolArgs struct {
	// The name of the agent pool.
	AgentPoolName string `pulumi:"agentPoolName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
}

// Agent Pool.
type LookupAgentPoolResult struct {
	// The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Desired Number of agents (VMs) specified to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.
	Count *int `pulumi:"count"`
	// CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using a snapshot.
	CreationData *CreationDataResponse `pulumi:"creationData"`
	// Whether to enable auto-scaler
	EnableAutoScaling *bool `pulumi:"enableAutoScaling"`
	// This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption
	EnableEncryptionAtHost *bool `pulumi:"enableEncryptionAtHost"`
	// See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.
	EnableFIPS *bool `pulumi:"enableFIPS"`
	// Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.
	EnableNodePublicIP *bool `pulumi:"enableNodePublicIP"`
	// Whether to enable UltraSSD
	EnableUltraSSD *bool `pulumi:"enableUltraSSD"`
	// GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
	GpuInstanceProfile *string `pulumi:"gpuInstanceProfile"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The Kubelet configuration on the agent pool nodes.
	KubeletConfig *KubeletConfigResponse `pulumi:"kubeletConfig"`
	// Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.
	KubeletDiskType *string `pulumi:"kubeletDiskType"`
	// The OS configuration of Linux agent nodes.
	LinuxOSConfig *LinuxOSConfigResponse `pulumi:"linuxOSConfig"`
	// The maximum number of nodes for auto-scaling
	MaxCount *int `pulumi:"maxCount"`
	// The maximum number of pods that can run on a node.
	MaxPods *int `pulumi:"maxPods"`
	// The minimum number of nodes for auto-scaling
	MinCount *int `pulumi:"minCount"`
	// A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
	Mode *string `pulumi:"mode"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `pulumi:"name"`
	// The version of node image
	NodeImageVersion string `pulumi:"nodeImageVersion"`
	// The node labels to be persisted across all nodes in agent pool.
	NodeLabels map[string]string `pulumi:"nodeLabels"`
	// This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}
	NodePublicIPPrefixID *string `pulumi:"nodePublicIPPrefixID"`
	// The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints []string `pulumi:"nodeTaints"`
	// As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool).
	OrchestratorVersion *string `pulumi:"orchestratorVersion"`
	// OS Disk Size in GB to be used to specify the disk size for every machine in the master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB *int `pulumi:"osDiskSizeGB"`
	// The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
	OsDiskType *string `pulumi:"osDiskType"`
	// Specifies an OS SKU. This value must not be specified if OSType is Windows.
	OsSKU *string `pulumi:"osSKU"`
	// The operating system type. The default is Linux.
	OsType *string `pulumi:"osType"`
	// If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	PodSubnetID *string `pulumi:"podSubnetID"`
	// Describes whether the Agent Pool is Running or Stopped
	PowerState PowerStateResponse `pulumi:"powerState"`
	// The current deployment or provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// The ID for Proximity Placement Group.
	ProximityPlacementGroupID *string `pulumi:"proximityPlacementGroupID"`
	// This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.
	ScaleDownMode *string `pulumi:"scaleDownMode"`
	// This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.
	ScaleSetEvictionPolicy *string `pulumi:"scaleSetEvictionPolicy"`
	// The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.
	ScaleSetPriority *string `pulumi:"scaleSetPriority"`
	// Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)
	SpotMaxPrice *float64 `pulumi:"spotMaxPrice"`
	// The tags to be persisted on the agent pool virtual machine scale set.
	Tags map[string]string `pulumi:"tags"`
	// The type of Agent Pool.
	Type string `pulumi:"type"`
	// Settings for upgrading the agentpool
	UpgradeSettings *AgentPoolUpgradeSettingsResponse `pulumi:"upgradeSettings"`
	// VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions
	VmSize *string `pulumi:"vmSize"`
	// If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	VnetSubnetID *string `pulumi:"vnetSubnetID"`
	// Determines the type of workload a node can run.
	WorkloadRuntime *string `pulumi:"workloadRuntime"`
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
	// The name of the agent pool.
	AgentPoolName pulumi.StringInput `pulumi:"agentPoolName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupAgentPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentPoolArgs)(nil)).Elem()
}

// Agent Pool.
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

func (o LookupAgentPoolResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAgentPoolResult] {
	return pulumix.Output[LookupAgentPoolResult]{
		OutputState: o.OutputState,
	}
}

// The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.
func (o LookupAgentPoolResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// Desired Number of agents (VMs) specified to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.
func (o LookupAgentPoolResultOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.Count }).(pulumi.IntPtrOutput)
}

// CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using a snapshot.
func (o LookupAgentPoolResultOutput) CreationData() CreationDataResponsePtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *CreationDataResponse { return v.CreationData }).(CreationDataResponsePtrOutput)
}

// Whether to enable auto-scaler
func (o LookupAgentPoolResultOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *bool { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

// This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption
func (o LookupAgentPoolResultOutput) EnableEncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *bool { return v.EnableEncryptionAtHost }).(pulumi.BoolPtrOutput)
}

// See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.
func (o LookupAgentPoolResultOutput) EnableFIPS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *bool { return v.EnableFIPS }).(pulumi.BoolPtrOutput)
}

// Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.
func (o LookupAgentPoolResultOutput) EnableNodePublicIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *bool { return v.EnableNodePublicIP }).(pulumi.BoolPtrOutput)
}

// Whether to enable UltraSSD
func (o LookupAgentPoolResultOutput) EnableUltraSSD() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *bool { return v.EnableUltraSSD }).(pulumi.BoolPtrOutput)
}

// GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
func (o LookupAgentPoolResultOutput) GpuInstanceProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.GpuInstanceProfile }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupAgentPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Kubelet configuration on the agent pool nodes.
func (o LookupAgentPoolResultOutput) KubeletConfig() KubeletConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *KubeletConfigResponse { return v.KubeletConfig }).(KubeletConfigResponsePtrOutput)
}

// Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.
func (o LookupAgentPoolResultOutput) KubeletDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.KubeletDiskType }).(pulumi.StringPtrOutput)
}

// The OS configuration of Linux agent nodes.
func (o LookupAgentPoolResultOutput) LinuxOSConfig() LinuxOSConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *LinuxOSConfigResponse { return v.LinuxOSConfig }).(LinuxOSConfigResponsePtrOutput)
}

// The maximum number of nodes for auto-scaling
func (o LookupAgentPoolResultOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

// The maximum number of pods that can run on a node.
func (o LookupAgentPoolResultOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

// The minimum number of nodes for auto-scaling
func (o LookupAgentPoolResultOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

// A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
func (o LookupAgentPoolResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupAgentPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// The version of node image
func (o LookupAgentPoolResultOutput) NodeImageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.NodeImageVersion }).(pulumi.StringOutput)
}

// The node labels to be persisted across all nodes in agent pool.
func (o LookupAgentPoolResultOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

// This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}
func (o LookupAgentPoolResultOutput) NodePublicIPPrefixID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.NodePublicIPPrefixID }).(pulumi.StringPtrOutput)
}

// The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
func (o LookupAgentPoolResultOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

// As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool).
func (o LookupAgentPoolResultOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

// OS Disk Size in GB to be used to specify the disk size for every machine in the master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
func (o LookupAgentPoolResultOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

// The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
func (o LookupAgentPoolResultOutput) OsDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.OsDiskType }).(pulumi.StringPtrOutput)
}

// Specifies an OS SKU. This value must not be specified if OSType is Windows.
func (o LookupAgentPoolResultOutput) OsSKU() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.OsSKU }).(pulumi.StringPtrOutput)
}

// The operating system type. The default is Linux.
func (o LookupAgentPoolResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

// If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
func (o LookupAgentPoolResultOutput) PodSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.PodSubnetID }).(pulumi.StringPtrOutput)
}

// Describes whether the Agent Pool is Running or Stopped
func (o LookupAgentPoolResultOutput) PowerState() PowerStateResponseOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) PowerStateResponse { return v.PowerState }).(PowerStateResponseOutput)
}

// The current deployment or provisioning state.
func (o LookupAgentPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The ID for Proximity Placement Group.
func (o LookupAgentPoolResultOutput) ProximityPlacementGroupID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.ProximityPlacementGroupID }).(pulumi.StringPtrOutput)
}

// This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.
func (o LookupAgentPoolResultOutput) ScaleDownMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.ScaleDownMode }).(pulumi.StringPtrOutput)
}

// This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.
func (o LookupAgentPoolResultOutput) ScaleSetEvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.ScaleSetEvictionPolicy }).(pulumi.StringPtrOutput)
}

// The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.
func (o LookupAgentPoolResultOutput) ScaleSetPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.ScaleSetPriority }).(pulumi.StringPtrOutput)
}

// Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)
func (o LookupAgentPoolResultOutput) SpotMaxPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *float64 { return v.SpotMaxPrice }).(pulumi.Float64PtrOutput)
}

// The tags to be persisted on the agent pool virtual machine scale set.
func (o LookupAgentPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Agent Pool.
func (o LookupAgentPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

// Settings for upgrading the agentpool
func (o LookupAgentPoolResultOutput) UpgradeSettings() AgentPoolUpgradeSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *AgentPoolUpgradeSettingsResponse { return v.UpgradeSettings }).(AgentPoolUpgradeSettingsResponsePtrOutput)
}

// VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions
func (o LookupAgentPoolResultOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

// If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
func (o LookupAgentPoolResultOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

// Determines the type of workload a node can run.
func (o LookupAgentPoolResultOutput) WorkloadRuntime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.WorkloadRuntime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAgentPoolResultOutput{})
}
