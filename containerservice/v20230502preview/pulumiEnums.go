// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230502preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
type AgentPoolMode string

const (
	// System agent pools are primarily for hosting critical system pods such as CoreDNS and metrics-server. System agent pools osType must be Linux. System agent pools VM SKU must have at least 2vCPUs and 4GB of memory.
	AgentPoolModeSystem = AgentPoolMode("System")
	// User agent pools are primarily for hosting your application pods.
	AgentPoolModeUser = AgentPoolMode("User")
)

// The type of Agent Pool.
type AgentPoolType string

const (
	// Create an Agent Pool backed by a Virtual Machine Scale Set.
	AgentPoolTypeVirtualMachineScaleSets = AgentPoolType("VirtualMachineScaleSets")
	// Use of this is strongly discouraged.
	AgentPoolTypeAvailabilitySet = AgentPoolType("AvailabilitySet")
)

// The type of the managed inbound Load Balancer BackendPool.
type BackendPoolType string

const (
	// The type of the managed inbound Load Balancer BackendPool. https://cloud-provider-azure.sigs.k8s.io/topics/loadbalancer/#configure-load-balancer-backend.
	BackendPoolTypeNodeIPConfiguration = BackendPoolType("NodeIPConfiguration")
	// The type of the managed inbound Load Balancer BackendPool. https://cloud-provider-azure.sigs.k8s.io/topics/loadbalancer/#configure-load-balancer-backend.
	BackendPoolTypeNodeIP = BackendPoolType("NodeIP")
)

// Tells whether the cluster is Running or Stopped
type Code string

const (
	// The cluster is running.
	CodeRunning = Code("Running")
	// The cluster is stopped.
	CodeStopped = Code("Stopped")
)

// The private link service connection status.
type ConnectionStatus string

const (
	ConnectionStatusPending      = ConnectionStatus("Pending")
	ConnectionStatusApproved     = ConnectionStatus("Approved")
	ConnectionStatusRejected     = ConnectionStatus("Rejected")
	ConnectionStatusDisconnected = ConnectionStatus("Disconnected")
)

// The list of control plane upgrade override settings.
type ControlPlaneUpgradeOverride string

const (
	// Upgrade the cluster control plane version without checking for recent Kubernetes deprecations usage.
	ControlPlaneUpgradeOverrideIgnoreKubernetesDeprecations = ControlPlaneUpgradeOverride("IgnoreKubernetesDeprecations")
)

// Controls which resource value autoscaler will change. Default value is RequestsAndLimits.
type ControlledValues string

const (
	// Autoscaler will control resource requests and limits.
	ControlledValuesRequestsAndLimits = ControlledValues("RequestsAndLimits")
	// Autoscaler will control resource requests only.
	ControlledValuesRequestsOnly = ControlledValues("RequestsOnly")
)

// If not specified, the default is 'random'. See [expanders](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders) for more information.
type Expander string

const (
	// Selects the node group that will have the least idle CPU (if tied, unused memory) after scale-up. This is useful when you have different classes of nodes, for example, high CPU or high memory nodes, and only want to expand those when there are pending pods that need a lot of those resources.
	Expander_Least_Waste = Expander("least-waste")
	// Selects the node group that would be able to schedule the most pods when scaling up. This is useful when you are using nodeSelector to make sure certain pods land on certain nodes. Note that this won't cause the autoscaler to select bigger nodes vs. smaller, as it can add multiple smaller nodes at once.
	Expander_Most_Pods = Expander("most-pods")
	// Selects the node group that has the highest priority assigned by the user. It's configuration is described in more details [here](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/expander/priority/readme.md).
	ExpanderPriority = Expander("priority")
	// Used when you don't have a particular need for the node groups to scale differently.
	ExpanderRandom = Expander("random")
)

// The type of the extended location.
type ExtendedLocationTypes string

const (
	ExtendedLocationTypesEdgeZone = ExtendedLocationTypes("EdgeZone")
)

// GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
type GPUInstanceProfile string

const (
	GPUInstanceProfileMIG1g = GPUInstanceProfile("MIG1g")
	GPUInstanceProfileMIG2g = GPUInstanceProfile("MIG2g")
	GPUInstanceProfileMIG3g = GPUInstanceProfile("MIG3g")
	GPUInstanceProfileMIG4g = GPUInstanceProfile("MIG4g")
	GPUInstanceProfileMIG7g = GPUInstanceProfile("MIG7g")
)

// The IP version to use for cluster networking and IP assignment.
type IpFamily string

const (
	IpFamilyIPv4 = IpFamily("IPv4")
	IpFamilyIPv6 = IpFamily("IPv6")
)

// IPVS scheduler, for more information please see http://www.linuxvirtualserver.org/docs/scheduling.html.
type IpvsScheduler string

const (
	// Round Robin
	IpvsSchedulerRoundRobin = IpvsScheduler("RoundRobin")
	// Least Connection
	IpvsSchedulerLeastConnection = IpvsScheduler("LeastConnection")
)

// Mode of an ingress gateway.
type IstioIngressGatewayMode string

const (
	// The ingress gateway is assigned a public IP address and is publicly accessible.
	IstioIngressGatewayModeExternal = IstioIngressGatewayMode("External")
	// The ingress gateway is assigned an internal IP address and cannot is accessed publicly.
	IstioIngressGatewayModeInternal = IstioIngressGatewayMode("Internal")
)

// Network access of key vault. The possible values are `Public` and `Private`. `Public` means the key vault allows public access from all networks. `Private` means the key vault disables public access and enables private link. The default value is `Public`.
type KeyVaultNetworkAccessTypes string

const (
	KeyVaultNetworkAccessTypesPublic  = KeyVaultNetworkAccessTypes("Public")
	KeyVaultNetworkAccessTypesPrivate = KeyVaultNetworkAccessTypes("Private")
)

// Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.
type KubeletDiskType string

const (
	// Kubelet will use the OS disk for its data.
	KubeletDiskTypeOS = KubeletDiskType("OS")
	// Kubelet will use the temporary disk for its data.
	KubeletDiskTypeTemporary = KubeletDiskType("Temporary")
)

// The support plan for the Managed Cluster. If unspecified, the default is 'KubernetesOfficial'.
type KubernetesSupportPlan string

const (
	// Support for the version is the same as for the open source Kubernetes offering. Official Kubernetes open source community support versions for 1 year after release.
	KubernetesSupportPlanKubernetesOfficial = KubernetesSupportPlan("KubernetesOfficial")
	// Support for the version extended past the KubernetesOfficial support of 1 year. AKS continues to patch CVEs for another 1 year, for a total of 2 years of support.
	KubernetesSupportPlanAKSLongTermSupport = KubernetesSupportPlan("AKSLongTermSupport")
)

// The guardrails level to be used. By default, Guardrails is enabled for all namespaces except those that AKS excludes via systemExcludedNamespaces
type Level string

const (
	LevelOff         = Level("Off")
	LevelWarning     = Level("Warning")
	LevelEnforcement = Level("Enforcement")
)

// The license type to use for Windows VMs. See [Azure Hybrid User Benefits](https://azure.microsoft.com/pricing/hybrid-benefit/faq/) for more details.
type LicenseType string

const (
	// No additional licensing is applied.
	LicenseTypeNone = LicenseType("None")
	// Enables Azure Hybrid User Benefits for Windows VMs.
	LicenseType_Windows_Server = LicenseType("Windows_Server")
)

// The default is 'standard'. See [Azure Load Balancer SKUs](https://docs.microsoft.com/azure/load-balancer/skus) for more information about the differences between load balancer SKUs.
type LoadBalancerSku string

const (
	// Use a a standard Load Balancer. This is the recommended Load Balancer SKU. For more information about on working with the load balancer in the managed cluster, see the [standard Load Balancer](https://docs.microsoft.com/azure/aks/load-balancer-standard) article.
	LoadBalancerSkuStandard = LoadBalancerSku("standard")
	// Use a basic Load Balancer with limited functionality.
	LoadBalancerSkuBasic = LoadBalancerSku("basic")
)

// The name of a managed cluster SKU.
type ManagedClusterSKUName string

const (
	// Base option for the AKS control plane.
	ManagedClusterSKUNameBase = ManagedClusterSKUName("Base")
)

// If not specified, the default is 'Free'. See [AKS Pricing Tier](https://learn.microsoft.com/azure/aks/free-standard-pricing-tiers) for more details.
type ManagedClusterSKUTier string

const (
	// Cluster has premium capabilities in addition to all of the capabilities included in 'Standard'. Premium enables selection of LongTermSupport (aka.ms/aks/lts) for certain Kubernetes versions.
	ManagedClusterSKUTierPremium = ManagedClusterSKUTier("Premium")
	// Recommended for mission-critical and production workloads. Includes Kubernetes control plane autoscaling, workload-intensive testing, and up to 5,000 nodes per cluster. Guarantees 99.95% availability of the Kubernetes API server endpoint for clusters that use Availability Zones and 99.9% of availability for clusters that don't use Availability Zones.
	ManagedClusterSKUTierStandard = ManagedClusterSKUTier("Standard")
	// The cluster management is free, but charged for VM, storage, and networking usage. Best for experimenting, learning, simple testing, or workloads with fewer than 10 nodes. Not recommended for production use cases.
	ManagedClusterSKUTierFree = ManagedClusterSKUTier("Free")
)

// Specify which proxy mode to use ('IPTABLES' or 'IPVS')
type Mode string

const (
	// IPTables proxy mode
	ModeIPTABLES = Mode("IPTABLES")
	// IPVS proxy mode. Must be using Kubernetes version >= 1.22.
	ModeIPVS = Mode("IPVS")
)

// Network dataplane used in the Kubernetes cluster.
type NetworkDataplane string

const (
	// Use Azure network dataplane.
	NetworkDataplaneAzure = NetworkDataplane("azure")
	// Use Cilium network dataplane. See [Azure CNI Powered by Cilium](https://learn.microsoft.com/azure/aks/azure-cni-powered-by-cilium) for more information.
	NetworkDataplaneCilium = NetworkDataplane("cilium")
)

// This cannot be specified if networkPlugin is anything other than 'azure'.
type NetworkMode string

const (
	// No bridge is created. Intra-VM Pod to Pod communication is through IP routes created by Azure CNI. See [Transparent Mode](https://docs.microsoft.com/azure/aks/faq#transparent-mode) for more information.
	NetworkModeTransparent = NetworkMode("transparent")
	// This is no longer supported
	NetworkModeBridge = NetworkMode("bridge")
)

// Network plugin used for building the Kubernetes network.
type NetworkPlugin string

const (
	// Use the Azure CNI network plugin. See [Azure CNI (advanced) networking](https://docs.microsoft.com/azure/aks/concepts-network#azure-cni-advanced-networking) for more information.
	NetworkPluginAzure = NetworkPlugin("azure")
	// Use the Kubenet network plugin. See [Kubenet (basic) networking](https://docs.microsoft.com/azure/aks/concepts-network#kubenet-basic-networking) for more information.
	NetworkPluginKubenet = NetworkPlugin("kubenet")
	// Do not use a network plugin. A custom CNI will need to be installed after cluster creation for networking functionality.
	NetworkPluginNone = NetworkPlugin("none")
)

// Network plugin mode used for building the Kubernetes network.
type NetworkPluginMode string

const (
	// Pods are given IPs from the PodCIDR address space but use Azure Routing Domains rather than Kubenet reference plugins host-local and bridge.
	NetworkPluginModeOverlay = NetworkPluginMode("overlay")
)

// Network policy used for building the Kubernetes network.
type NetworkPolicy string

const (
	// Use Calico network policies. See [differences between Azure and Calico policies](https://docs.microsoft.com/azure/aks/use-network-policies#differences-between-azure-and-calico-policies-and-their-capabilities) for more information.
	NetworkPolicyCalico = NetworkPolicy("calico")
	// Use Azure network policies. See [differences between Azure and Calico policies](https://docs.microsoft.com/azure/aks/use-network-policies#differences-between-azure-and-calico-policies-and-their-capabilities) for more information.
	NetworkPolicyAzure = NetworkPolicy("azure")
	// Use Cilium to enforce network policies. This requires networkDataplane to be 'cilium'.
	NetworkPolicyCilium = NetworkPolicy("cilium")
)

// The default is Unmanaged, but may change to either NodeImage or SecurityPatch at GA.
type NodeOSUpgradeChannel string

const (
	// No attempt to update your machines OS will be made either by OS or by rolling VHDs. This means you are responsible for your security updates
	NodeOSUpgradeChannelNone = NodeOSUpgradeChannel("None")
	// OS updates will be applied automatically through the OS built-in patching infrastructure. Newly scaled in machines will be unpatched initially, and will be patched at some later time by the OS's infrastructure. Behavior of this option depends on the OS in question. Ubuntu and Mariner apply security patches through unattended upgrade roughly once a day around 06:00 UTC. Windows does not apply security patches automatically and so for them this option is equivalent to None till further notice
	NodeOSUpgradeChannelUnmanaged = NodeOSUpgradeChannel("Unmanaged")
	// AKS will update the nodes VHD with patches from the image maintainer labelled "security only" on a regular basis. Where possible, patches will also be applied without reimaging to existing nodes. Some patches, such as kernel patches, cannot be applied to existing nodes without disruption. For such patches, the VHD will be updated, and machines will be rolling reimaged to that VHD following maintenance windows and surge settings. This option incurs the extra cost of hosting the VHDs in your node resource group.
	NodeOSUpgradeChannelSecurityPatch = NodeOSUpgradeChannel("SecurityPatch")
	// AKS will update the nodes with a newly patched VHD containing security fixes and bugfixes on a weekly cadence. With the VHD update machines will be rolling reimaged to that VHD following maintenance windows and surge settings. No extra VHD cost is incurred when choosing this option as AKS hosts the images.
	NodeOSUpgradeChannelNodeImage = NodeOSUpgradeChannel("NodeImage")
)

// The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
type OSDiskType string

const (
	// Azure replicates the operating system disk for a virtual machine to Azure storage to avoid data loss should the VM need to be relocated to another host. Since containers aren't designed to have local state persisted, this behavior offers limited value while providing some drawbacks, including slower node provisioning and higher read/write latency.
	OSDiskTypeManaged = OSDiskType("Managed")
	// Ephemeral OS disks are stored only on the host machine, just like a temporary disk. This provides lower read/write latency, along with faster node scaling and cluster upgrades.
	OSDiskTypeEphemeral = OSDiskType("Ephemeral")
)

// Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.
type OSSKU string

const (
	// Use Ubuntu as the OS for node images.
	OSSKUUbuntu = OSSKU("Ubuntu")
	// Deprecated OSSKU. Microsoft recommends that new deployments choose 'AzureLinux' instead.
	OSSKUMariner = OSSKU("Mariner")
	// Use AzureLinux as the OS for node images. Azure Linux is a container-optimized Linux distro built by Microsoft, visit https://aka.ms/azurelinux for more information.
	OSSKUAzureLinux = OSSKU("AzureLinux")
	// Deprecated OSSKU. Microsoft recommends that new deployments choose 'AzureLinux' instead.
	OSSKUCBLMariner = OSSKU("CBLMariner")
	// Use Windows2019 as the OS for node images. Unsupported for system node pools. Windows2019 only supports Windows2019 containers; it cannot run Windows2022 containers and vice versa.
	OSSKUWindows2019 = OSSKU("Windows2019")
	// Use Windows2022 as the OS for node images. Unsupported for system node pools. Windows2022 only supports Windows2022 containers; it cannot run Windows2019 containers and vice versa.
	OSSKUWindows2022 = OSSKU("Windows2022")
)

// The operating system type. The default is Linux.
type OSType string

const (
	// Use Linux.
	OSTypeLinux = OSType("Linux")
	// Use Windows.
	OSTypeWindows = OSType("Windows")
)

// This can only be set at cluster creation time and cannot be changed later. For more information see [egress outbound type](https://docs.microsoft.com/azure/aks/egress-outboundtype).
type OutboundType string

const (
	// The load balancer is used for egress through an AKS assigned public IP. This supports Kubernetes services of type 'loadBalancer'. For more information see [outbound type loadbalancer](https://docs.microsoft.com/azure/aks/egress-outboundtype#outbound-type-of-loadbalancer).
	OutboundTypeLoadBalancer = OutboundType("loadBalancer")
	// Egress paths must be defined by the user. This is an advanced scenario and requires proper network configuration. For more information see [outbound type userDefinedRouting](https://docs.microsoft.com/azure/aks/egress-outboundtype#outbound-type-of-userdefinedrouting).
	OutboundTypeUserDefinedRouting = OutboundType("userDefinedRouting")
	// The AKS-managed NAT gateway is used for egress.
	OutboundTypeManagedNATGateway = OutboundType("managedNATGateway")
	// The user-assigned NAT gateway associated to the cluster subnet is used for egress. This is an advanced scenario and requires proper network configuration.
	OutboundTypeUserAssignedNATGateway = OutboundType("userAssignedNATGateway")
)

// The network protocol of the port.
type Protocol string

const (
	// TCP protocol.
	ProtocolTCP = Protocol("TCP")
	// UDP protocol.
	ProtocolUDP = Protocol("UDP")
)

// Allow or deny public network access for AKS
type PublicNetworkAccess string

const (
	// Inbound/Outbound to the managedCluster is allowed.
	PublicNetworkAccessEnabled = PublicNetworkAccess("Enabled")
	// Inbound traffic to managedCluster is disabled, traffic from managedCluster is allowed.
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
	// Inbound/Outbound traffic is managed by Microsoft.Network/NetworkSecurityPerimeters.
	PublicNetworkAccessSecuredByPerimeter = PublicNetworkAccess("SecuredByPerimeter")
)

// For more information see [use managed identities in AKS](https://docs.microsoft.com/azure/aks/use-managed-identity).
type ResourceIdentityType string

const (
	// Use an implicitly created system assigned managed identity to manage cluster resources. Master components in the control plane such as kube-controller-manager will use the system assigned managed identity to manipulate Azure resources.
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	// Use a user-specified identity to manage cluster resources. Master components in the control plane such as kube-controller-manager will use the specified user assigned managed identity to manipulate Azure resources.
	ResourceIdentityTypeUserAssigned = ResourceIdentityType("UserAssigned")
	// Do not use a managed identity for the Managed Cluster, service principal will be used instead.
	ResourceIdentityTypeNone = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ResourceIdentityTypeInput is an input type that accepts ResourceIdentityTypeArgs and ResourceIdentityTypeOutput values.
// You can construct a concrete instance of `ResourceIdentityTypeInput` via:
//
//	ResourceIdentityTypeArgs{...}
type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToOutput(ctx context.Context) pulumix.Output[*ResourceIdentityType] {
	return pulumix.Output[*ResourceIdentityType]{
		OutputState: in.ToResourceIdentityTypePtrOutputWithContext(ctx).OutputState,
	}
}

// The restriction level applied to the cluster's node resource group
type RestrictionLevel string

const (
	// All RBAC permissions are allowed on the managed node resource group
	RestrictionLevelUnrestricted = RestrictionLevel("Unrestricted")
	// Only */read RBAC permissions allowed on the managed node resource group
	RestrictionLevelReadOnly = RestrictionLevel("ReadOnly")
)

// This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.
type ScaleDownMode string

const (
	// Create new instances during scale up and remove instances during scale down.
	ScaleDownModeDelete = ScaleDownMode("Delete")
	// Attempt to start deallocated instances (if they exist) during scale up and deallocate instances during scale down.
	ScaleDownModeDeallocate = ScaleDownMode("Deallocate")
)

// This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.
type ScaleSetEvictionPolicy string

const (
	// Nodes in the underlying Scale Set of the node pool are deleted when they're evicted.
	ScaleSetEvictionPolicyDelete = ScaleSetEvictionPolicy("Delete")
	// Nodes in the underlying Scale Set of the node pool are set to the stopped-deallocated state upon eviction. Nodes in the stopped-deallocated state count against your compute quota and can cause issues with cluster scaling or upgrading.
	ScaleSetEvictionPolicyDeallocate = ScaleSetEvictionPolicy("Deallocate")
)

// The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.
type ScaleSetPriority string

const (
	// Spot priority VMs will be used. There is no SLA for spot nodes. See [spot on AKS](https://docs.microsoft.com/azure/aks/spot-node-pool) for more information.
	ScaleSetPrioritySpot = ScaleSetPriority("Spot")
	// Regular VMs will be used.
	ScaleSetPriorityRegular = ScaleSetPriority("Regular")
)

// Mode of the service mesh.
type ServiceMeshMode string

const (
	// Istio deployed as an AKS addon.
	ServiceMeshModeIstio = ServiceMeshMode("Istio")
	// Mesh is disabled.
	ServiceMeshModeDisabled = ServiceMeshMode("Disabled")
)

// The type of a snapshot. The default is NodePool.
type SnapshotType string

const (
	// The snapshot is a snapshot of a node pool.
	SnapshotTypeNodePool = SnapshotType("NodePool")
	// The snapshot is a snapshot of a managed cluster.
	SnapshotTypeManagedCluster = SnapshotType("ManagedCluster")
)

// Specifies on which instance of the allowed days specified in daysOfWeek the maintenance occurs.
type Type string

const (
	// First.
	TypeFirst = Type("First")
	// Second.
	TypeSecond = Type("Second")
	// Third.
	TypeThird = Type("Third")
	// Fourth.
	TypeFourth = Type("Fourth")
	// Last.
	TypeLast = Type("Last")
)

// Each update mode level is a superset of the lower levels. Off<Initial<Recreate<=Auto. For example: if UpdateMode is Initial, it means VPA sets the recommended resources in the VerticalPodAutoscaler Custom Resource (from UpdateMode Off) and also assigns resources on pod creation (from Initial). The default value is Off.
type UpdateMode string

const (
	// Autoscaler never changes pod resources but provides recommendations.
	UpdateModeOff = UpdateMode("Off")
	// Autoscaler only assigns resources on pod creation and doesn't change them during the lifetime of the pod.
	UpdateModeInitial = UpdateMode("Initial")
	// Autoscaler assigns resources on pod creation and updates pods that need further scaling during their lifetime by deleting and recreating.
	UpdateModeRecreate = UpdateMode("Recreate")
	// Autoscaler chooses the update mode. Autoscaler currently does the same as Recreate. In the future, it may take advantage of restart-free mechanisms once they are available.
	UpdateModeAuto = UpdateMode("Auto")
)

// For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel).
type UpgradeChannel string

const (
	// Automatically upgrade the cluster to the latest supported patch release on the latest supported minor version. In cases where the cluster is at a version of Kubernetes that is at an N-2 minor version where N is the latest supported minor version, the cluster first upgrades to the latest supported patch version on N-1 minor version. For example, if a cluster is running version 1.17.7 and versions 1.17.9, 1.18.4, 1.18.6, and 1.19.1 are available, your cluster first is upgraded to 1.18.6, then is upgraded to 1.19.1.
	UpgradeChannelRapid = UpgradeChannel("rapid")
	// Automatically upgrade the cluster to the latest supported patch release on minor version N-1, where N is the latest supported minor version. For example, if a cluster is running version 1.17.7 and versions 1.17.9, 1.18.4, 1.18.6, and 1.19.1 are available, your cluster is upgraded to 1.18.6.
	UpgradeChannelStable = UpgradeChannel("stable")
	// Automatically upgrade the cluster to the latest supported patch version when it becomes available while keeping the minor version the same. For example, if a cluster is running version 1.17.7 and versions 1.17.9, 1.18.4, 1.18.6, and 1.19.1 are available, your cluster is upgraded to 1.17.9.
	UpgradeChannelPatch = UpgradeChannel("patch")
	// Automatically upgrade the node image to the latest version available. Consider using nodeOSUpgradeChannel instead as that allows you to configure node OS patching separate from Kubernetes version patching
	UpgradeChannel_Node_Image = UpgradeChannel("node-image")
	// Disables auto-upgrades and keeps the cluster at its current version of Kubernetes.
	UpgradeChannelNone = UpgradeChannel("none")
)

// The day of the week.
type WeekDay string

const (
	WeekDaySunday    = WeekDay("Sunday")
	WeekDayMonday    = WeekDay("Monday")
	WeekDayTuesday   = WeekDay("Tuesday")
	WeekDayWednesday = WeekDay("Wednesday")
	WeekDayThursday  = WeekDay("Thursday")
	WeekDayFriday    = WeekDay("Friday")
	WeekDaySaturday  = WeekDay("Saturday")
)

// Determines the type of workload a node can run.
type WorkloadRuntime string

const (
	// Nodes will use Kubelet to run standard OCI container workloads.
	WorkloadRuntimeOCIContainer = WorkloadRuntime("OCIContainer")
	// Nodes will use Krustlet to run WASM workloads using the WASI provider (Preview).
	WorkloadRuntimeWasmWasi = WorkloadRuntime("WasmWasi")
	// Nodes can use (Kata + Cloud Hypervisor + Hyper-V) to enable Nested VM-based pods (Preview). Due to the use Hyper-V, AKS node OS itself is a nested VM (the root OS) of Hyper-V. Thus it can only be used with VM series that support Nested Virtualization such as Dv3 series.
	WorkloadRuntimeKataMshvVmIsolation = WorkloadRuntime("KataMshvVmIsolation")
)

func init() {
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
