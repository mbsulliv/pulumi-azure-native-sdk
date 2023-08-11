// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

// Values can be systemAssignedIdentity or userAssignedIdentity
type IdentityType string

const (
	IdentityTypeSystemAssignedIdentity    = IdentityType("systemAssignedIdentity")
	IdentityTypeUserAssignedIdentity      = IdentityType("userAssignedIdentity")
	IdentityTypeDelegatedResourceIdentity = IdentityType("delegatedResourceIdentity")
)

// Values are enabled and disabled.
type InfrastructureEncryption string

const (
	InfrastructureEncryptionEnabled  = InfrastructureEncryption("enabled")
	InfrastructureEncryptionDisabled = InfrastructureEncryption("disabled")
)

// Get or Set Kind property.
type Kind string

const (
	KindGen1 = Kind("Gen1")
	KindGen2 = Kind("Gen2")
)

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned, UserAssigned")
)

// The name of the SKU, in standard format (such as S0).
type Name string

const (
	NameS0 = Name("S0")
	NameS1 = Name("S1")
	NameG2 = Name("G2")
)

// The Map account key to use for signing. Picking `primaryKey` or `secondaryKey` will use the Map account Shared Keys, and using `managedIdentity` will use the auto-renewed private key to sign the SAS.
type SigningKey string

const (
	SigningKeyPrimaryKey      = SigningKey("primaryKey")
	SigningKeySecondaryKey    = SigningKey("secondaryKey")
	SigningKeyManagedIdentity = SigningKey("managedIdentity")
)

func init() {
}
