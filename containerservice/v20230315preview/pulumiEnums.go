// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230315preview

// The upgrade type.
// Full requires the KubernetesVersion property to be set.
// NodeImageOnly requires the KubernetesVersion property not to be set.
type ManagedClusterUpgradeType string

const (
	// Full upgrades the control plane and all agent pools of the target ManagedClusters.
	ManagedClusterUpgradeTypeFull = ManagedClusterUpgradeType("Full")
	// NodeImageOnly upgrades only the node images of the target ManagedClusters.
	ManagedClusterUpgradeTypeNodeImageOnly = ManagedClusterUpgradeType("NodeImageOnly")
)

func init() {
}