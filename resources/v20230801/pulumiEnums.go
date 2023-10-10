// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801

// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
type CleanupOptions string

const (
	CleanupOptionsAlways       = CleanupOptions("Always")
	CleanupOptionsOnSuccess    = CleanupOptions("OnSuccess")
	CleanupOptionsOnExpiration = CleanupOptions("OnExpiration")
)

// Type of the managed identity.
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeUserAssigned = ManagedServiceIdentityType("UserAssigned")
)

// Type of the script.
type ScriptType string

const (
	ScriptTypeAzurePowerShell = ScriptType("AzurePowerShell")
	ScriptTypeAzureCLI        = ScriptType("AzureCLI")
)

func init() {
}
