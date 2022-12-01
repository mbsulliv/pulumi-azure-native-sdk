// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Agent Pool.
type AgentPool struct {
	pulumi.CustomResourceState

	// Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 100 (inclusive) for user pools and in the range of 1 to 100 (inclusive) for system pools. The default value is 1.
	Count pulumi.IntPtrOutput `pulumi:"count"`
	// Whether to enable auto-scaler
	EnableAutoScaling pulumi.BoolPtrOutput `pulumi:"enableAutoScaling"`
	// Whether to enable EncryptionAtHost
	EnableEncryptionAtHost pulumi.BoolPtrOutput `pulumi:"enableEncryptionAtHost"`
	// Whether to use FIPS enabled OS
	EnableFIPS pulumi.BoolPtrOutput `pulumi:"enableFIPS"`
	// Enable public IP for nodes
	EnableNodePublicIP pulumi.BoolPtrOutput `pulumi:"enableNodePublicIP"`
	// GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU. Supported values are MIG1g, MIG2g, MIG3g, MIG4g and MIG7g.
	GpuInstanceProfile pulumi.StringPtrOutput `pulumi:"gpuInstanceProfile"`
	// KubeletConfig specifies the configuration of kubelet on agent nodes.
	KubeletConfig KubeletConfigResponsePtrOutput `pulumi:"kubeletConfig"`
	// KubeletDiskType determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage. Currently allows one value, OS, resulting in Kubelet using the OS disk for data.
	KubeletDiskType pulumi.StringPtrOutput `pulumi:"kubeletDiskType"`
	// LinuxOSConfig specifies the OS configuration of linux agent nodes.
	LinuxOSConfig LinuxOSConfigResponsePtrOutput `pulumi:"linuxOSConfig"`
	// Maximum number of nodes for auto-scaling
	MaxCount pulumi.IntPtrOutput `pulumi:"maxCount"`
	// Maximum number of pods that can run on a node.
	MaxPods pulumi.IntPtrOutput `pulumi:"maxPods"`
	// Minimum number of nodes for auto-scaling
	MinCount pulumi.IntPtrOutput `pulumi:"minCount"`
	// AgentPoolMode represents mode of an agent pool
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Version of node image
	NodeImageVersion pulumi.StringOutput `pulumi:"nodeImageVersion"`
	// Agent pool node labels to be persisted across all nodes in agent pool.
	NodeLabels pulumi.StringMapOutput `pulumi:"nodeLabels"`
	// Public IP Prefix ID. VM nodes use IPs assigned from this Public IP Prefix.
	NodePublicIPPrefixID pulumi.StringPtrOutput `pulumi:"nodePublicIPPrefixID"`
	// Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints pulumi.StringArrayOutput `pulumi:"nodeTaints"`
	// Version of orchestrator specified when creating the managed cluster.
	OrchestratorVersion pulumi.StringPtrOutput `pulumi:"orchestratorVersion"`
	// OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB pulumi.IntPtrOutput `pulumi:"osDiskSizeGB"`
	// OS disk type to be used for machines in a given agent pool. Allowed values are 'Ephemeral' and 'Managed'. If unspecified, defaults to 'Ephemeral' when the VM supports ephemeral OS and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation.
	OsDiskType pulumi.StringPtrOutput `pulumi:"osDiskType"`
	// OsSKU to be used to specify os sku. Choose from Ubuntu(default) and CBLMariner for Linux OSType. Not applicable to Windows OSType.
	OsSKU pulumi.StringPtrOutput `pulumi:"osSKU"`
	// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// Pod SubnetID specifies the VNet's subnet identifier for pods.
	PodSubnetID pulumi.StringPtrOutput `pulumi:"podSubnetID"`
	// Describes whether the Agent Pool is Running or Stopped
	PowerState PowerStateResponseOutput `pulumi:"powerState"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The ID for Proximity Placement Group.
	ProximityPlacementGroupID pulumi.StringPtrOutput `pulumi:"proximityPlacementGroupID"`
	// ScaleSetEvictionPolicy to be used to specify eviction policy for Spot virtual machine scale set. Default to Delete.
	ScaleSetEvictionPolicy pulumi.StringPtrOutput `pulumi:"scaleSetEvictionPolicy"`
	// ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
	ScaleSetPriority pulumi.StringPtrOutput `pulumi:"scaleSetPriority"`
	// SpotMaxPrice to be used to specify the maximum price you are willing to pay in US Dollars. Possible values are any decimal value greater than zero or -1 which indicates default price to be up-to on-demand.
	SpotMaxPrice pulumi.Float64PtrOutput `pulumi:"spotMaxPrice"`
	// Agent pool tags to be persisted on the agent pool virtual machine scale set.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// AgentPoolType represents types of an agent pool
	Type pulumi.StringOutput `pulumi:"type"`
	// Settings for upgrading the agentpool
	UpgradeSettings AgentPoolUpgradeSettingsResponsePtrOutput `pulumi:"upgradeSettings"`
	// Size of agent VMs.
	VmSize pulumi.StringPtrOutput `pulumi:"vmSize"`
	// VNet SubnetID specifies the VNet's subnet identifier for nodes and maybe pods
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
			Type: pulumi.String("azure-native:containerservice/v20210501:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210701:AgentPool"),
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
	err := ctx.RegisterResource("azure-native:containerservice/v20210301:AgentPool", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:containerservice/v20210301:AgentPool", name, id, state, &resource, opts...)
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
	// Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 100 (inclusive) for user pools and in the range of 1 to 100 (inclusive) for system pools. The default value is 1.
	Count *int `pulumi:"count"`
	// Whether to enable auto-scaler
	EnableAutoScaling *bool `pulumi:"enableAutoScaling"`
	// Whether to enable EncryptionAtHost
	EnableEncryptionAtHost *bool `pulumi:"enableEncryptionAtHost"`
	// Whether to use FIPS enabled OS
	EnableFIPS *bool `pulumi:"enableFIPS"`
	// Enable public IP for nodes
	EnableNodePublicIP *bool `pulumi:"enableNodePublicIP"`
	// GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU. Supported values are MIG1g, MIG2g, MIG3g, MIG4g and MIG7g.
	GpuInstanceProfile *string `pulumi:"gpuInstanceProfile"`
	// KubeletConfig specifies the configuration of kubelet on agent nodes.
	KubeletConfig *KubeletConfig `pulumi:"kubeletConfig"`
	// KubeletDiskType determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage. Currently allows one value, OS, resulting in Kubelet using the OS disk for data.
	KubeletDiskType *string `pulumi:"kubeletDiskType"`
	// LinuxOSConfig specifies the OS configuration of linux agent nodes.
	LinuxOSConfig *LinuxOSConfig `pulumi:"linuxOSConfig"`
	// Maximum number of nodes for auto-scaling
	MaxCount *int `pulumi:"maxCount"`
	// Maximum number of pods that can run on a node.
	MaxPods *int `pulumi:"maxPods"`
	// Minimum number of nodes for auto-scaling
	MinCount *int `pulumi:"minCount"`
	// AgentPoolMode represents mode of an agent pool
	Mode *string `pulumi:"mode"`
	// Agent pool node labels to be persisted across all nodes in agent pool.
	NodeLabels map[string]string `pulumi:"nodeLabels"`
	// Public IP Prefix ID. VM nodes use IPs assigned from this Public IP Prefix.
	NodePublicIPPrefixID *string `pulumi:"nodePublicIPPrefixID"`
	// Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints []string `pulumi:"nodeTaints"`
	// Version of orchestrator specified when creating the managed cluster.
	OrchestratorVersion *string `pulumi:"orchestratorVersion"`
	// OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB *int `pulumi:"osDiskSizeGB"`
	// OS disk type to be used for machines in a given agent pool. Allowed values are 'Ephemeral' and 'Managed'. If unspecified, defaults to 'Ephemeral' when the VM supports ephemeral OS and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation.
	OsDiskType *string `pulumi:"osDiskType"`
	// OsSKU to be used to specify os sku. Choose from Ubuntu(default) and CBLMariner for Linux OSType. Not applicable to Windows OSType.
	OsSKU *string `pulumi:"osSKU"`
	// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
	OsType *string `pulumi:"osType"`
	// Pod SubnetID specifies the VNet's subnet identifier for pods.
	PodSubnetID *string `pulumi:"podSubnetID"`
	// The ID for Proximity Placement Group.
	ProximityPlacementGroupID *string `pulumi:"proximityPlacementGroupID"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
	// ScaleSetEvictionPolicy to be used to specify eviction policy for Spot virtual machine scale set. Default to Delete.
	ScaleSetEvictionPolicy *string `pulumi:"scaleSetEvictionPolicy"`
	// ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
	ScaleSetPriority *string `pulumi:"scaleSetPriority"`
	// SpotMaxPrice to be used to specify the maximum price you are willing to pay in US Dollars. Possible values are any decimal value greater than zero or -1 which indicates default price to be up-to on-demand.
	SpotMaxPrice *float64 `pulumi:"spotMaxPrice"`
	// Agent pool tags to be persisted on the agent pool virtual machine scale set.
	Tags map[string]string `pulumi:"tags"`
	// AgentPoolType represents types of an agent pool
	Type *string `pulumi:"type"`
	// Settings for upgrading the agentpool
	UpgradeSettings *AgentPoolUpgradeSettings `pulumi:"upgradeSettings"`
	// Size of agent VMs.
	VmSize *string `pulumi:"vmSize"`
	// VNet SubnetID specifies the VNet's subnet identifier for nodes and maybe pods
	VnetSubnetID *string `pulumi:"vnetSubnetID"`
}

// The set of arguments for constructing a AgentPool resource.
type AgentPoolArgs struct {
	// The name of the agent pool.
	AgentPoolName pulumi.StringPtrInput
	// Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
	AvailabilityZones pulumi.StringArrayInput
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 100 (inclusive) for user pools and in the range of 1 to 100 (inclusive) for system pools. The default value is 1.
	Count pulumi.IntPtrInput
	// Whether to enable auto-scaler
	EnableAutoScaling pulumi.BoolPtrInput
	// Whether to enable EncryptionAtHost
	EnableEncryptionAtHost pulumi.BoolPtrInput
	// Whether to use FIPS enabled OS
	EnableFIPS pulumi.BoolPtrInput
	// Enable public IP for nodes
	EnableNodePublicIP pulumi.BoolPtrInput
	// GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU. Supported values are MIG1g, MIG2g, MIG3g, MIG4g and MIG7g.
	GpuInstanceProfile pulumi.StringPtrInput
	// KubeletConfig specifies the configuration of kubelet on agent nodes.
	KubeletConfig KubeletConfigPtrInput
	// KubeletDiskType determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage. Currently allows one value, OS, resulting in Kubelet using the OS disk for data.
	KubeletDiskType pulumi.StringPtrInput
	// LinuxOSConfig specifies the OS configuration of linux agent nodes.
	LinuxOSConfig LinuxOSConfigPtrInput
	// Maximum number of nodes for auto-scaling
	MaxCount pulumi.IntPtrInput
	// Maximum number of pods that can run on a node.
	MaxPods pulumi.IntPtrInput
	// Minimum number of nodes for auto-scaling
	MinCount pulumi.IntPtrInput
	// AgentPoolMode represents mode of an agent pool
	Mode pulumi.StringPtrInput
	// Agent pool node labels to be persisted across all nodes in agent pool.
	NodeLabels pulumi.StringMapInput
	// Public IP Prefix ID. VM nodes use IPs assigned from this Public IP Prefix.
	NodePublicIPPrefixID pulumi.StringPtrInput
	// Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints pulumi.StringArrayInput
	// Version of orchestrator specified when creating the managed cluster.
	OrchestratorVersion pulumi.StringPtrInput
	// OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB pulumi.IntPtrInput
	// OS disk type to be used for machines in a given agent pool. Allowed values are 'Ephemeral' and 'Managed'. If unspecified, defaults to 'Ephemeral' when the VM supports ephemeral OS and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation.
	OsDiskType pulumi.StringPtrInput
	// OsSKU to be used to specify os sku. Choose from Ubuntu(default) and CBLMariner for Linux OSType. Not applicable to Windows OSType.
	OsSKU pulumi.StringPtrInput
	// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
	OsType pulumi.StringPtrInput
	// Pod SubnetID specifies the VNet's subnet identifier for pods.
	PodSubnetID pulumi.StringPtrInput
	// The ID for Proximity Placement Group.
	ProximityPlacementGroupID pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput
	// ScaleSetEvictionPolicy to be used to specify eviction policy for Spot virtual machine scale set. Default to Delete.
	ScaleSetEvictionPolicy pulumi.StringPtrInput
	// ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
	ScaleSetPriority pulumi.StringPtrInput
	// SpotMaxPrice to be used to specify the maximum price you are willing to pay in US Dollars. Possible values are any decimal value greater than zero or -1 which indicates default price to be up-to on-demand.
	SpotMaxPrice pulumi.Float64PtrInput
	// Agent pool tags to be persisted on the agent pool virtual machine scale set.
	Tags pulumi.StringMapInput
	// AgentPoolType represents types of an agent pool
	Type pulumi.StringPtrInput
	// Settings for upgrading the agentpool
	UpgradeSettings AgentPoolUpgradeSettingsPtrInput
	// Size of agent VMs.
	VmSize pulumi.StringPtrInput
	// VNet SubnetID specifies the VNet's subnet identifier for nodes and maybe pods
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

// Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
func (o AgentPoolOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 100 (inclusive) for user pools and in the range of 1 to 100 (inclusive) for system pools. The default value is 1.
func (o AgentPoolOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.Count }).(pulumi.IntPtrOutput)
}

// Whether to enable auto-scaler
func (o AgentPoolOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

// Whether to enable EncryptionAtHost
func (o AgentPoolOutput) EnableEncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableEncryptionAtHost }).(pulumi.BoolPtrOutput)
}

// Whether to use FIPS enabled OS
func (o AgentPoolOutput) EnableFIPS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableFIPS }).(pulumi.BoolPtrOutput)
}

// Enable public IP for nodes
func (o AgentPoolOutput) EnableNodePublicIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableNodePublicIP }).(pulumi.BoolPtrOutput)
}

// GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU. Supported values are MIG1g, MIG2g, MIG3g, MIG4g and MIG7g.
func (o AgentPoolOutput) GpuInstanceProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.GpuInstanceProfile }).(pulumi.StringPtrOutput)
}

// KubeletConfig specifies the configuration of kubelet on agent nodes.
func (o AgentPoolOutput) KubeletConfig() KubeletConfigResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) KubeletConfigResponsePtrOutput { return v.KubeletConfig }).(KubeletConfigResponsePtrOutput)
}

// KubeletDiskType determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage. Currently allows one value, OS, resulting in Kubelet using the OS disk for data.
func (o AgentPoolOutput) KubeletDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.KubeletDiskType }).(pulumi.StringPtrOutput)
}

// LinuxOSConfig specifies the OS configuration of linux agent nodes.
func (o AgentPoolOutput) LinuxOSConfig() LinuxOSConfigResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) LinuxOSConfigResponsePtrOutput { return v.LinuxOSConfig }).(LinuxOSConfigResponsePtrOutput)
}

// Maximum number of nodes for auto-scaling
func (o AgentPoolOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MaxCount }).(pulumi.IntPtrOutput)
}

// Maximum number of pods that can run on a node.
func (o AgentPoolOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MaxPods }).(pulumi.IntPtrOutput)
}

// Minimum number of nodes for auto-scaling
func (o AgentPoolOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MinCount }).(pulumi.IntPtrOutput)
}

// AgentPoolMode represents mode of an agent pool
func (o AgentPoolOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o AgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Version of node image
func (o AgentPoolOutput) NodeImageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.NodeImageVersion }).(pulumi.StringOutput)
}

// Agent pool node labels to be persisted across all nodes in agent pool.
func (o AgentPoolOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringMapOutput { return v.NodeLabels }).(pulumi.StringMapOutput)
}

// Public IP Prefix ID. VM nodes use IPs assigned from this Public IP Prefix.
func (o AgentPoolOutput) NodePublicIPPrefixID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.NodePublicIPPrefixID }).(pulumi.StringPtrOutput)
}

// Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
func (o AgentPoolOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringArrayOutput { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

// Version of orchestrator specified when creating the managed cluster.
func (o AgentPoolOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

// OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
func (o AgentPoolOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

// OS disk type to be used for machines in a given agent pool. Allowed values are 'Ephemeral' and 'Managed'. If unspecified, defaults to 'Ephemeral' when the VM supports ephemeral OS and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation.
func (o AgentPoolOutput) OsDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OsDiskType }).(pulumi.StringPtrOutput)
}

// OsSKU to be used to specify os sku. Choose from Ubuntu(default) and CBLMariner for Linux OSType. Not applicable to Windows OSType.
func (o AgentPoolOutput) OsSKU() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OsSKU }).(pulumi.StringPtrOutput)
}

// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
func (o AgentPoolOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

// Pod SubnetID specifies the VNet's subnet identifier for pods.
func (o AgentPoolOutput) PodSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.PodSubnetID }).(pulumi.StringPtrOutput)
}

// Describes whether the Agent Pool is Running or Stopped
func (o AgentPoolOutput) PowerState() PowerStateResponseOutput {
	return o.ApplyT(func(v *AgentPool) PowerStateResponseOutput { return v.PowerState }).(PowerStateResponseOutput)
}

// The current deployment or provisioning state, which only appears in the response.
func (o AgentPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The ID for Proximity Placement Group.
func (o AgentPoolOutput) ProximityPlacementGroupID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.ProximityPlacementGroupID }).(pulumi.StringPtrOutput)
}

// ScaleSetEvictionPolicy to be used to specify eviction policy for Spot virtual machine scale set. Default to Delete.
func (o AgentPoolOutput) ScaleSetEvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.ScaleSetEvictionPolicy }).(pulumi.StringPtrOutput)
}

// ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
func (o AgentPoolOutput) ScaleSetPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.ScaleSetPriority }).(pulumi.StringPtrOutput)
}

// SpotMaxPrice to be used to specify the maximum price you are willing to pay in US Dollars. Possible values are any decimal value greater than zero or -1 which indicates default price to be up-to on-demand.
func (o AgentPoolOutput) SpotMaxPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.Float64PtrOutput { return v.SpotMaxPrice }).(pulumi.Float64PtrOutput)
}

// Agent pool tags to be persisted on the agent pool virtual machine scale set.
func (o AgentPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// AgentPoolType represents types of an agent pool
func (o AgentPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Settings for upgrading the agentpool
func (o AgentPoolOutput) UpgradeSettings() AgentPoolUpgradeSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) AgentPoolUpgradeSettingsResponsePtrOutput { return v.UpgradeSettings }).(AgentPoolUpgradeSettingsResponsePtrOutput)
}

// Size of agent VMs.
func (o AgentPoolOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.VmSize }).(pulumi.StringPtrOutput)
}

// VNet SubnetID specifies the VNet's subnet identifier for nodes and maybe pods
func (o AgentPoolOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentPoolOutput{})
}
