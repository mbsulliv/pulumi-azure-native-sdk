// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Agent Pool.
type AgentPool struct {
	pulumi.CustomResourceState

	// The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.
	Count pulumi.IntPtrOutput `pulumi:"count"`
	// Whether to enable auto-scaler
	EnableAutoScaling pulumi.BoolPtrOutput `pulumi:"enableAutoScaling"`
	// This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption
	EnableEncryptionAtHost pulumi.BoolPtrOutput `pulumi:"enableEncryptionAtHost"`
	// See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.
	EnableFIPS pulumi.BoolPtrOutput `pulumi:"enableFIPS"`
	// Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.
	EnableNodePublicIP pulumi.BoolPtrOutput `pulumi:"enableNodePublicIP"`
	// Whether to enable UltraSSD
	EnableUltraSSD pulumi.BoolPtrOutput `pulumi:"enableUltraSSD"`
	// GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
	GpuInstanceProfile pulumi.StringPtrOutput `pulumi:"gpuInstanceProfile"`
	// The Kubelet configuration on the agent pool nodes.
	KubeletConfig KubeletConfigResponsePtrOutput `pulumi:"kubeletConfig"`
	// Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.
	KubeletDiskType pulumi.StringPtrOutput `pulumi:"kubeletDiskType"`
	// The OS configuration of Linux agent nodes.
	LinuxOSConfig LinuxOSConfigResponsePtrOutput `pulumi:"linuxOSConfig"`
	// The maximum number of nodes for auto-scaling
	MaxCount pulumi.IntPtrOutput `pulumi:"maxCount"`
	// The maximum number of pods that can run on a node.
	MaxPods pulumi.IntPtrOutput `pulumi:"maxPods"`
	// The minimum number of nodes for auto-scaling
	MinCount pulumi.IntPtrOutput `pulumi:"minCount"`
	// A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The version of node image
	NodeImageVersion pulumi.StringOutput `pulumi:"nodeImageVersion"`
	// The node labels to be persisted across all nodes in agent pool.
	NodeLabels pulumi.StringMapOutput `pulumi:"nodeLabels"`
	// This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}
	NodePublicIPPrefixID pulumi.StringPtrOutput `pulumi:"nodePublicIPPrefixID"`
	// The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints pulumi.StringArrayOutput `pulumi:"nodeTaints"`
	// As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool).
	OrchestratorVersion pulumi.StringPtrOutput `pulumi:"orchestratorVersion"`
	// OS Disk Size in GB to be used to specify the disk size for every machine in the master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB pulumi.IntPtrOutput `pulumi:"osDiskSizeGB"`
	// The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
	OsDiskType pulumi.StringPtrOutput `pulumi:"osDiskType"`
	// Specifies an OS SKU. This value must not be specified if OSType is Windows.
	OsSKU pulumi.StringPtrOutput `pulumi:"osSKU"`
	// The operating system type. The default is Linux.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	PodSubnetID pulumi.StringPtrOutput `pulumi:"podSubnetID"`
	// Describes whether the Agent Pool is Running or Stopped
	PowerState PowerStateResponseOutput `pulumi:"powerState"`
	// The current deployment or provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The ID for Proximity Placement Group.
	ProximityPlacementGroupID pulumi.StringPtrOutput `pulumi:"proximityPlacementGroupID"`
	// This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.
	ScaleDownMode pulumi.StringPtrOutput `pulumi:"scaleDownMode"`
	// This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.
	ScaleSetEvictionPolicy pulumi.StringPtrOutput `pulumi:"scaleSetEvictionPolicy"`
	// The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.
	ScaleSetPriority pulumi.StringPtrOutput `pulumi:"scaleSetPriority"`
	// Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)
	SpotMaxPrice pulumi.Float64PtrOutput `pulumi:"spotMaxPrice"`
	// The tags to be persisted on the agent pool virtual machine scale set.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Agent Pool.
	Type pulumi.StringOutput `pulumi:"type"`
	// Settings for upgrading the agentpool
	UpgradeSettings AgentPoolUpgradeSettingsResponsePtrOutput `pulumi:"upgradeSettings"`
	// VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions
	VmSize pulumi.StringPtrOutput `pulumi:"vmSize"`
	// If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	VnetSubnetID pulumi.StringPtrOutput `pulumi:"vnetSubnetID"`
}

// NewAgentPool registers a new resource with the given unique name, arguments, and options.
func NewAgentPool(ctx *pulumi.Context,
	name string, args *AgentPoolArgs, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191001:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210501:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211001:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211101preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220102preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220202preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220302preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220802preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220803preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221002preview:AgentPool"),
		},
	})
	opts = append(opts, aliases)
	var resource AgentPool
	err := ctx.RegisterResource("azure-native:containerservice/v20210701:AgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentPool gets an existing AgentPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentPoolState, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	var resource AgentPool
	err := ctx.ReadResource("azure-native:containerservice/v20210701:AgentPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentPool resources.
type agentPoolState struct {
}

type AgentPoolState struct {
}

func (AgentPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolState)(nil)).Elem()
}

type agentPoolArgs struct {
	// The name of the agent pool.
	AgentPoolName *string `pulumi:"agentPoolName"`
	// The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.
	Count *int `pulumi:"count"`
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
	// The Kubelet configuration on the agent pool nodes.
	KubeletConfig *KubeletConfig `pulumi:"kubeletConfig"`
	// Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.
	KubeletDiskType *string `pulumi:"kubeletDiskType"`
	// The OS configuration of Linux agent nodes.
	LinuxOSConfig *LinuxOSConfig `pulumi:"linuxOSConfig"`
	// The maximum number of nodes for auto-scaling
	MaxCount *int `pulumi:"maxCount"`
	// The maximum number of pods that can run on a node.
	MaxPods *int `pulumi:"maxPods"`
	// The minimum number of nodes for auto-scaling
	MinCount *int `pulumi:"minCount"`
	// A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
	Mode *string `pulumi:"mode"`
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
	// The ID for Proximity Placement Group.
	ProximityPlacementGroupID *string `pulumi:"proximityPlacementGroupID"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
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
	Type *string `pulumi:"type"`
	// Settings for upgrading the agentpool
	UpgradeSettings *AgentPoolUpgradeSettings `pulumi:"upgradeSettings"`
	// VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions
	VmSize *string `pulumi:"vmSize"`
	// If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	VnetSubnetID *string `pulumi:"vnetSubnetID"`
}

// The set of arguments for constructing a AgentPool resource.
type AgentPoolArgs struct {
	// The name of the agent pool.
	AgentPoolName pulumi.StringPtrInput
	// The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.
	AvailabilityZones pulumi.StringArrayInput
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.
	Count pulumi.IntPtrInput
	// Whether to enable auto-scaler
	EnableAutoScaling pulumi.BoolPtrInput
	// This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption
	EnableEncryptionAtHost pulumi.BoolPtrInput
	// See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.
	EnableFIPS pulumi.BoolPtrInput
	// Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.
	EnableNodePublicIP pulumi.BoolPtrInput
	// Whether to enable UltraSSD
	EnableUltraSSD pulumi.BoolPtrInput
	// GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
	GpuInstanceProfile pulumi.StringPtrInput
	// The Kubelet configuration on the agent pool nodes.
	KubeletConfig KubeletConfigPtrInput
	// Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.
	KubeletDiskType pulumi.StringPtrInput
	// The OS configuration of Linux agent nodes.
	LinuxOSConfig LinuxOSConfigPtrInput
	// The maximum number of nodes for auto-scaling
	MaxCount pulumi.IntPtrInput
	// The maximum number of pods that can run on a node.
	MaxPods pulumi.IntPtrInput
	// The minimum number of nodes for auto-scaling
	MinCount pulumi.IntPtrInput
	// A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
	Mode pulumi.StringPtrInput
	// The node labels to be persisted across all nodes in agent pool.
	NodeLabels pulumi.StringMapInput
	// This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}
	NodePublicIPPrefixID pulumi.StringPtrInput
	// The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints pulumi.StringArrayInput
	// As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool).
	OrchestratorVersion pulumi.StringPtrInput
	// OS Disk Size in GB to be used to specify the disk size for every machine in the master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB pulumi.IntPtrInput
	// The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
	OsDiskType pulumi.StringPtrInput
	// Specifies an OS SKU. This value must not be specified if OSType is Windows.
	OsSKU pulumi.StringPtrInput
	// The operating system type. The default is Linux.
	OsType pulumi.StringPtrInput
	// If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	PodSubnetID pulumi.StringPtrInput
	// The ID for Proximity Placement Group.
	ProximityPlacementGroupID pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput
	// This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.
	ScaleDownMode pulumi.StringPtrInput
	// This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.
	ScaleSetEvictionPolicy pulumi.StringPtrInput
	// The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.
	ScaleSetPriority pulumi.StringPtrInput
	// Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)
	SpotMaxPrice pulumi.Float64PtrInput
	// The tags to be persisted on the agent pool virtual machine scale set.
	Tags pulumi.StringMapInput
	// The type of Agent Pool.
	Type pulumi.StringPtrInput
	// Settings for upgrading the agentpool
	UpgradeSettings AgentPoolUpgradeSettingsPtrInput
	// VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions
	VmSize pulumi.StringPtrInput
	// If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	VnetSubnetID pulumi.StringPtrInput
}

func (AgentPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolArgs)(nil)).Elem()
}

type AgentPoolInput interface {
	pulumi.Input

	ToAgentPoolOutput() AgentPoolOutput
	ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput
}

func (*AgentPool) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPool)(nil)).Elem()
}

func (i *AgentPool) ToAgentPoolOutput() AgentPoolOutput {
	return i.ToAgentPoolOutputWithContext(context.Background())
}

func (i *AgentPool) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolOutput)
}

type AgentPoolOutput struct{ *pulumi.OutputState }

func (AgentPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPool)(nil)).Elem()
}

func (o AgentPoolOutput) ToAgentPoolOutput() AgentPoolOutput {
	return o
}

func (o AgentPoolOutput) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return o
}

// The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.
func (o AgentPoolOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.
func (o AgentPoolOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.Count }).(pulumi.IntPtrOutput)
}

// Whether to enable auto-scaler
func (o AgentPoolOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

// This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption
func (o AgentPoolOutput) EnableEncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableEncryptionAtHost }).(pulumi.BoolPtrOutput)
}

// See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.
func (o AgentPoolOutput) EnableFIPS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableFIPS }).(pulumi.BoolPtrOutput)
}

// Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.
func (o AgentPoolOutput) EnableNodePublicIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableNodePublicIP }).(pulumi.BoolPtrOutput)
}

// Whether to enable UltraSSD
func (o AgentPoolOutput) EnableUltraSSD() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableUltraSSD }).(pulumi.BoolPtrOutput)
}

// GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
func (o AgentPoolOutput) GpuInstanceProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.GpuInstanceProfile }).(pulumi.StringPtrOutput)
}

// The Kubelet configuration on the agent pool nodes.
func (o AgentPoolOutput) KubeletConfig() KubeletConfigResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) KubeletConfigResponsePtrOutput { return v.KubeletConfig }).(KubeletConfigResponsePtrOutput)
}

// Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.
func (o AgentPoolOutput) KubeletDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.KubeletDiskType }).(pulumi.StringPtrOutput)
}

// The OS configuration of Linux agent nodes.
func (o AgentPoolOutput) LinuxOSConfig() LinuxOSConfigResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) LinuxOSConfigResponsePtrOutput { return v.LinuxOSConfig }).(LinuxOSConfigResponsePtrOutput)
}

// The maximum number of nodes for auto-scaling
func (o AgentPoolOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MaxCount }).(pulumi.IntPtrOutput)
}

// The maximum number of pods that can run on a node.
func (o AgentPoolOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MaxPods }).(pulumi.IntPtrOutput)
}

// The minimum number of nodes for auto-scaling
func (o AgentPoolOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MinCount }).(pulumi.IntPtrOutput)
}

// A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
func (o AgentPoolOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o AgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The version of node image
func (o AgentPoolOutput) NodeImageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.NodeImageVersion }).(pulumi.StringOutput)
}

// The node labels to be persisted across all nodes in agent pool.
func (o AgentPoolOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringMapOutput { return v.NodeLabels }).(pulumi.StringMapOutput)
}

// This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}
func (o AgentPoolOutput) NodePublicIPPrefixID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.NodePublicIPPrefixID }).(pulumi.StringPtrOutput)
}

// The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
func (o AgentPoolOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringArrayOutput { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

// As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool).
func (o AgentPoolOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

// OS Disk Size in GB to be used to specify the disk size for every machine in the master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
func (o AgentPoolOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

// The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
func (o AgentPoolOutput) OsDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OsDiskType }).(pulumi.StringPtrOutput)
}

// Specifies an OS SKU. This value must not be specified if OSType is Windows.
func (o AgentPoolOutput) OsSKU() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OsSKU }).(pulumi.StringPtrOutput)
}

// The operating system type. The default is Linux.
func (o AgentPoolOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

// If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
func (o AgentPoolOutput) PodSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.PodSubnetID }).(pulumi.StringPtrOutput)
}

// Describes whether the Agent Pool is Running or Stopped
func (o AgentPoolOutput) PowerState() PowerStateResponseOutput {
	return o.ApplyT(func(v *AgentPool) PowerStateResponseOutput { return v.PowerState }).(PowerStateResponseOutput)
}

// The current deployment or provisioning state.
func (o AgentPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The ID for Proximity Placement Group.
func (o AgentPoolOutput) ProximityPlacementGroupID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.ProximityPlacementGroupID }).(pulumi.StringPtrOutput)
}

// This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.
func (o AgentPoolOutput) ScaleDownMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.ScaleDownMode }).(pulumi.StringPtrOutput)
}

// This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.
func (o AgentPoolOutput) ScaleSetEvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.ScaleSetEvictionPolicy }).(pulumi.StringPtrOutput)
}

// The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.
func (o AgentPoolOutput) ScaleSetPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.ScaleSetPriority }).(pulumi.StringPtrOutput)
}

// Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)
func (o AgentPoolOutput) SpotMaxPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.Float64PtrOutput { return v.SpotMaxPrice }).(pulumi.Float64PtrOutput)
}

// The tags to be persisted on the agent pool virtual machine scale set.
func (o AgentPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Agent Pool.
func (o AgentPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Settings for upgrading the agentpool
func (o AgentPoolOutput) UpgradeSettings() AgentPoolUpgradeSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) AgentPoolUpgradeSettingsResponsePtrOutput { return v.UpgradeSettings }).(AgentPoolUpgradeSettingsResponsePtrOutput)
}

// VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions
func (o AgentPoolOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.VmSize }).(pulumi.StringPtrOutput)
}

// If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
func (o AgentPoolOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentPoolOutput{})
}
