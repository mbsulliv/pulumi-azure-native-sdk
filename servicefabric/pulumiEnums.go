// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The network traffic is allowed or denied.
type Access string

const (
	AccessAllow = Access("allow")
	AccessDeny  = Access("deny")
)

// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0. Only applies when **clusterUpgradeMode** is set to 'Automatic'.
type ClusterUpgradeCadence string

const (
	// Cluster upgrade starts immediately after a new version is rolled out. Recommended for Test/Dev clusters.
	ClusterUpgradeCadenceWave0 = ClusterUpgradeCadence("Wave0")
	// Cluster upgrade starts 7 days after a new version is rolled out. Recommended for Pre-prod clusters.
	ClusterUpgradeCadenceWave1 = ClusterUpgradeCadence("Wave1")
	// Cluster upgrade starts 14 days after a new version is rolled out. Recommended for Production clusters.
	ClusterUpgradeCadenceWave2 = ClusterUpgradeCadence("Wave2")
)

// The upgrade mode of the cluster when new Service Fabric runtime version is available.
type ClusterUpgradeMode string

const (
	// The cluster will be automatically upgraded to the latest Service Fabric runtime version, **clusterUpgradeCadence** will determine when the upgrade starts after the new version becomes available.
	ClusterUpgradeModeAutomatic = ClusterUpgradeMode("Automatic")
	// The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
	ClusterUpgradeModeManual = ClusterUpgradeMode("Manual")
)

// Network security rule direction.
type Direction string

const (
	DirectionInbound  = Direction("inbound")
	DirectionOutbound = Direction("outbound")
)

// Managed data disk type. Specifies the storage account type for the managed disk
type DiskType string

const (
	// Standard HDD locally redundant storage. Best for backup, non-critical, and infrequent access.
	DiskType_Standard_LRS = DiskType("Standard_LRS")
	// Standard SSD locally redundant storage. Best for web servers, lightly used enterprise applications and dev/test.
	DiskType_StandardSSD_LRS = DiskType("StandardSSD_LRS")
	// Premium SSD locally redundant storage. Best for production and performance sensitive workloads.
	DiskType_Premium_LRS = DiskType("Premium_LRS")
)

// Specifies the eviction policy for virtual machines in a SPOT node type. Default is Delete.
type EvictionPolicyType string

const (
	// Eviction policy will be Delete for SPOT vms.
	EvictionPolicyTypeDelete = EvictionPolicyType("Delete")
	// Eviction policy will be Deallocate for SPOT vms.
	EvictionPolicyTypeDeallocate = EvictionPolicyType("Deallocate")
)

// The compensating action to perform when a Monitored upgrade encounters monitoring policy or health policy violations. Invalid indicates the failure action is invalid. Rollback specifies that the upgrade will start rolling back automatically. Manual indicates that the upgrade will switch to UnmonitoredManual upgrade mode.
type FailureAction string

const (
	// Indicates that a rollback of the upgrade will be performed by Service Fabric if the upgrade fails.
	FailureActionRollback = FailureAction("Rollback")
	// Indicates that a manual repair will need to be performed by the administrator if the upgrade fails. Service Fabric will not proceed to the next upgrade domain automatically.
	FailureActionManual = FailureAction("Manual")
)

// The IP address type of this frontend configuration. If omitted the default value is IPv4.
type IPAddressType string

const (
	// IPv4 address type.
	IPAddressTypeIPv4 = IPAddressType("IPv4")
	// IPv6 address type.
	IPAddressTypeIPv6 = IPAddressType("IPv6")
)

// Available cluster add-on features
type ManagedClusterAddOnFeature string

const (
	// Dns service
	ManagedClusterAddOnFeatureDnsService = ManagedClusterAddOnFeature("DnsService")
	// Backup and restore service
	ManagedClusterAddOnFeatureBackupRestoreService = ManagedClusterAddOnFeature("BackupRestoreService")
	// Resource monitor service
	ManagedClusterAddOnFeatureResourceMonitorService = ManagedClusterAddOnFeature("ResourceMonitorService")
)

// The type of managed identity for the resource.
type ManagedIdentityType string

const (
	// Indicates that no identity is associated with the resource.
	ManagedIdentityTypeNone = ManagedIdentityType("None")
	// Indicates that system assigned identity is associated with the resource.
	ManagedIdentityTypeSystemAssigned = ManagedIdentityType("SystemAssigned")
	// Indicates that user assigned identity is associated with the resource.
	ManagedIdentityTypeUserAssigned = ManagedIdentityType("UserAssigned")
	// Indicates that both system assigned and user assigned identity are associated with the resource.
	ManagedIdentityType_SystemAssigned_UserAssigned = ManagedIdentityType("SystemAssigned, UserAssigned")
)

func (ManagedIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityType)(nil)).Elem()
}

func (e ManagedIdentityType) ToManagedIdentityTypeOutput() ManagedIdentityTypeOutput {
	return pulumi.ToOutput(e).(ManagedIdentityTypeOutput)
}

func (e ManagedIdentityType) ToManagedIdentityTypeOutputWithContext(ctx context.Context) ManagedIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedIdentityTypeOutput)
}

func (e ManagedIdentityType) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return e.ToManagedIdentityTypePtrOutputWithContext(context.Background())
}

func (e ManagedIdentityType) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return ManagedIdentityType(e).ToManagedIdentityTypeOutputWithContext(ctx).ToManagedIdentityTypePtrOutputWithContext(ctx)
}

func (e ManagedIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedIdentityTypeOutput struct{ *pulumi.OutputState }

func (ManagedIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityType)(nil)).Elem()
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypeOutput() ManagedIdentityTypeOutput {
	return o
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypeOutputWithContext(ctx context.Context) ManagedIdentityTypeOutput {
	return o
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return o.ToManagedIdentityTypePtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityType) *ManagedIdentityType {
		return &v
	}).(ManagedIdentityTypePtrOutput)
}

func (o ManagedIdentityTypeOutput) ToOutput(ctx context.Context) pulumix.Output[ManagedIdentityType] {
	return pulumix.Output[ManagedIdentityType]{
		OutputState: o.OutputState,
	}
}

func (o ManagedIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityType)(nil)).Elem()
}

func (o ManagedIdentityTypePtrOutput) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return o
}

func (o ManagedIdentityTypePtrOutput) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return o
}

func (o ManagedIdentityTypePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ManagedIdentityType] {
	return pulumix.Output[*ManagedIdentityType]{
		OutputState: o.OutputState,
	}
}

func (o ManagedIdentityTypePtrOutput) Elem() ManagedIdentityTypeOutput {
	return o.ApplyT(func(v *ManagedIdentityType) ManagedIdentityType {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityType
		return ret
	}).(ManagedIdentityTypeOutput)
}

func (o ManagedIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ManagedIdentityTypeInput is an input type that accepts ManagedIdentityTypeArgs and ManagedIdentityTypeOutput values.
// You can construct a concrete instance of `ManagedIdentityTypeInput` via:
//
//	ManagedIdentityTypeArgs{...}
type ManagedIdentityTypeInput interface {
	pulumi.Input

	ToManagedIdentityTypeOutput() ManagedIdentityTypeOutput
	ToManagedIdentityTypeOutputWithContext(context.Context) ManagedIdentityTypeOutput
}

var managedIdentityTypePtrType = reflect.TypeOf((**ManagedIdentityType)(nil)).Elem()

type ManagedIdentityTypePtrInput interface {
	pulumi.Input

	ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput
	ToManagedIdentityTypePtrOutputWithContext(context.Context) ManagedIdentityTypePtrOutput
}

type managedIdentityTypePtr string

func ManagedIdentityTypePtr(v string) ManagedIdentityTypePtrInput {
	return (*managedIdentityTypePtr)(&v)
}

func (*managedIdentityTypePtr) ElementType() reflect.Type {
	return managedIdentityTypePtrType
}

func (in *managedIdentityTypePtr) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ManagedIdentityTypePtrOutput)
}

func (in *managedIdentityTypePtr) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedIdentityTypePtrOutput)
}

func (in *managedIdentityTypePtr) ToOutput(ctx context.Context) pulumix.Output[*ManagedIdentityType] {
	return pulumix.Output[*ManagedIdentityType]{
		OutputState: in.ToManagedIdentityTypePtrOutputWithContext(ctx).OutputState,
	}
}

// Specifies the move cost for the service.
type MoveCost string

const (
	// Zero move cost. This value is zero.
	MoveCostZero = MoveCost("Zero")
	// Specifies the move cost of the service as Low. The value is 1.
	MoveCostLow = MoveCost("Low")
	// Specifies the move cost of the service as Medium. The value is 2.
	MoveCostMedium = MoveCost("Medium")
	// Specifies the move cost of the service as High. The value is 3.
	MoveCostHigh = MoveCost("High")
)

// Network protocol this rule applies to.
type NsgProtocol string

const (
	NsgProtocolHttp  = NsgProtocol("http")
	NsgProtocolHttps = NsgProtocol("https")
	NsgProtocolTcp   = NsgProtocol("tcp")
	NsgProtocolUdp   = NsgProtocol("udp")
	NsgProtocolIcmp  = NsgProtocol("icmp")
	NsgProtocolAh    = NsgProtocol("ah")
	NsgProtocolEsp   = NsgProtocol("esp")
)

// Specifies how the service is partitioned.
type PartitionScheme string

const (
	// Indicates that the partition is based on string names, and is a SingletonPartitionScheme object, The value is 0.
	PartitionSchemeSingleton = PartitionScheme("Singleton")
	// Indicates that the partition is based on Int64 key ranges, and is a UniformInt64RangePartitionScheme object. The value is 1.
	PartitionSchemeUniformInt64Range = PartitionScheme("UniformInt64Range")
	// Indicates that the partition is based on string names, and is a NamedPartitionScheme object. The value is 2.
	PartitionSchemeNamed = PartitionScheme("Named")
)

// Enable or Disable apply network policies on private end point in the subnet.
type PrivateEndpointNetworkPolicies string

const (
	PrivateEndpointNetworkPoliciesEnabled  = PrivateEndpointNetworkPolicies("enabled")
	PrivateEndpointNetworkPoliciesDisabled = PrivateEndpointNetworkPolicies("disabled")
)

// Enable or Disable apply network policies on private link service in the subnet.
type PrivateLinkServiceNetworkPolicies string

const (
	PrivateLinkServiceNetworkPoliciesEnabled  = PrivateLinkServiceNetworkPolicies("enabled")
	PrivateLinkServiceNetworkPoliciesDisabled = PrivateLinkServiceNetworkPolicies("disabled")
)

// the reference to the load balancer probe used by the load balancing rule.
type ProbeProtocol string

const (
	ProbeProtocolTcp   = ProbeProtocol("tcp")
	ProbeProtocolHttp  = ProbeProtocol("http")
	ProbeProtocolHttps = ProbeProtocol("https")
)

// The reference to the transport protocol used by the load balancing rule.
type Protocol string

const (
	ProtocolTcp = Protocol("tcp")
	ProtocolUdp = Protocol("udp")
)

// The mode used to monitor health during a rolling upgrade. The values are Monitored, and UnmonitoredAuto.
type RollingUpgradeMode string

const (
	// The upgrade will stop after completing each upgrade domain and automatically monitor health before proceeding. The value is 0.
	RollingUpgradeModeMonitored = RollingUpgradeMode("Monitored")
	// The upgrade will proceed automatically without performing any health monitoring. The value is 1.
	RollingUpgradeModeUnmonitoredAuto = RollingUpgradeMode("UnmonitoredAuto")
)

// Specifies the security type of the nodeType. Only TrustedLaunch is currently supported
type SecurityType string

const (
	// Trusted Launch is a security type that secures generation 2 virtual machines.
	SecurityTypeTrustedLaunch = SecurityType("TrustedLaunch")
)

// The ServiceCorrelationScheme which describes the relationship between this service and the service specified via ServiceName.
type ServiceCorrelationScheme string

const (
	// Aligned affinity ensures that the primaries of the partitions of the affinitized services are collocated on the same nodes. This is the default and is the same as selecting the Affinity scheme. The value is 0.
	ServiceCorrelationSchemeAlignedAffinity = ServiceCorrelationScheme("AlignedAffinity")
	// Non-Aligned affinity guarantees that all replicas of each service will be placed on the same nodes. Unlike Aligned Affinity, this does not guarantee that replicas of particular role will be collocated. The value is 1.
	ServiceCorrelationSchemeNonAlignedAffinity = ServiceCorrelationScheme("NonAlignedAffinity")
)

// The kind of service (Stateless or Stateful).
type ServiceKind string

const (
	// Does not use Service Fabric to make its state highly available or reliable. The value is 0.
	ServiceKindStateless = ServiceKind("Stateless")
	// Uses Service Fabric to make its state or part of its state highly available and reliable. The value is 1.
	ServiceKindStateful = ServiceKind("Stateful")
)

// The service load metric relative weight, compared to other metrics configured for this service, as a number.
type ServiceLoadMetricWeight string

const (
	// Disables resource balancing for this metric. This value is zero.
	ServiceLoadMetricWeightZero = ServiceLoadMetricWeight("Zero")
	// Specifies the metric weight of the service load as Low. The value is 1.
	ServiceLoadMetricWeightLow = ServiceLoadMetricWeight("Low")
	// Specifies the metric weight of the service load as Medium. The value is 2.
	ServiceLoadMetricWeightMedium = ServiceLoadMetricWeight("Medium")
	// Specifies the metric weight of the service load as High. The value is 3.
	ServiceLoadMetricWeightHigh = ServiceLoadMetricWeight("High")
)

// The activation Mode of the service package
type ServicePackageActivationMode string

const (
	// Indicates the application package activation mode will use shared process.
	ServicePackageActivationModeSharedProcess = ServicePackageActivationMode("SharedProcess")
	// Indicates the application package activation mode will use exclusive process.
	ServicePackageActivationModeExclusiveProcess = ServicePackageActivationMode("ExclusiveProcess")
)

// The type of placement policy for a service fabric service. Following are the possible values.
type ServicePlacementPolicyType string

const (
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementInvalidDomainPolicyDescription, which indicates that a particular fault or upgrade domain cannot be used for placement of this service. The value is 0.
	ServicePlacementPolicyTypeInvalidDomain = ServicePlacementPolicyType("InvalidDomain")
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementRequireDomainDistributionPolicyDescription indicating that the replicas of the service must be placed in a specific domain. The value is 1.
	ServicePlacementPolicyTypeRequiredDomain = ServicePlacementPolicyType("RequiredDomain")
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementPreferPrimaryDomainPolicyDescription, which indicates that if possible the Primary replica for the partitions of the service should be located in a particular domain as an optimization. The value is 2.
	ServicePlacementPolicyTypePreferredPrimaryDomain = ServicePlacementPolicyType("PreferredPrimaryDomain")
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementRequireDomainDistributionPolicyDescription, indicating that the system will disallow placement of any two replicas from the same partition in the same domain at any time. The value is 3.
	ServicePlacementPolicyTypeRequiredDomainDistribution = ServicePlacementPolicyType("RequiredDomainDistribution")
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementNonPartiallyPlaceServicePolicyDescription, which indicates that if possible all replicas of a particular partition of the service should be placed atomically. The value is 4.
	ServicePlacementPolicyTypeNonPartiallyPlaceService = ServicePlacementPolicyType("NonPartiallyPlaceService")
)

// Specifies the mechanism associated with this scaling policy.
type ServiceScalingMechanismKind string

const (
	// Represents a scaling mechanism for adding or removing instances of stateless service partition. The value is 0.
	ServiceScalingMechanismKindScalePartitionInstanceCount = ServiceScalingMechanismKind("ScalePartitionInstanceCount")
	// Represents a scaling mechanism for adding or removing named partitions of a stateless service. The value is 1.
	ServiceScalingMechanismKindAddRemoveIncrementalNamedPartition = ServiceScalingMechanismKind("AddRemoveIncrementalNamedPartition")
)

// Specifies the trigger associated with this scaling policy.
type ServiceScalingTriggerKind string

const (
	// Represents a scaling trigger related to an average load of a metric/resource of a partition. The value is 0.
	ServiceScalingTriggerKindAveragePartitionLoadTrigger = ServiceScalingTriggerKind("AveragePartitionLoadTrigger")
	// Represents a scaling policy related to an average load of a metric/resource of a service. The value is 1.
	ServiceScalingTriggerKindAverageServiceLoadTrigger = ServiceScalingTriggerKind("AverageServiceLoadTrigger")
)

// Sku Name.
type SkuName string

const (
	// Basic requires a minimum of 3 nodes and allows only 1 node type.
	SkuNameBasic = SkuName("Basic")
	// Requires a minimum of 5 nodes and allows 1 or more node type.
	SkuNameStandard = SkuName("Standard")
)

// action to be performed on the vms before bootstrapping the service fabric runtime.
type VmSetupAction string

const (
	// Enable windows containers feature.
	VmSetupActionEnableContainers = VmSetupAction("EnableContainers")
	// Enables windows HyperV feature.
	VmSetupActionEnableHyperV = VmSetupAction("EnableHyperV")
)

// Indicates the update mode for Cross Az clusters.
type ZonalUpdateMode string

const (
	// The cluster will use 5 upgrade domains for Cross Az Node types.
	ZonalUpdateModeStandard = ZonalUpdateMode("Standard")
	// The cluster will use a maximum of 3 upgrade domains per zone instead of 5 for Cross Az Node types for faster deployments.
	ZonalUpdateModeFast = ZonalUpdateMode("Fast")
)

func init() {
	pulumi.RegisterOutputType(ManagedIdentityTypeOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTypePtrOutput{})
}
