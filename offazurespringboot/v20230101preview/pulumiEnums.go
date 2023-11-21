// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

// The resource provisioning state.
type ProvisioningState string

const (
	ProvisioningStateUnknown      = ProvisioningState("Unknown")
	ProvisioningStateSucceeded    = ProvisioningState("Succeeded")
	ProvisioningStateFailed       = ProvisioningState("Failed")
	ProvisioningStateCanceled     = ProvisioningState("Canceled")
	ProvisioningStateAccepted     = ProvisioningState("Accepted")
	ProvisioningStateProvisioning = ProvisioningState("Provisioning")
	ProvisioningStateDeleting     = ProvisioningState("Deleting")
)

func init() {
}