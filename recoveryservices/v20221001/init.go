// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:recoveryservices/v20221001:ReplicationFabric":
		r = &ReplicationFabric{}
	case "azure-native:recoveryservices/v20221001:ReplicationMigrationItem":
		r = &ReplicationMigrationItem{}
	case "azure-native:recoveryservices/v20221001:ReplicationNetworkMapping":
		r = &ReplicationNetworkMapping{}
	case "azure-native:recoveryservices/v20221001:ReplicationPolicy":
		r = &ReplicationPolicy{}
	case "azure-native:recoveryservices/v20221001:ReplicationProtectedItem":
		r = &ReplicationProtectedItem{}
	case "azure-native:recoveryservices/v20221001:ReplicationProtectionContainerMapping":
		r = &ReplicationProtectionContainerMapping{}
	case "azure-native:recoveryservices/v20221001:ReplicationRecoveryPlan":
		r = &ReplicationRecoveryPlan{}
	case "azure-native:recoveryservices/v20221001:ReplicationRecoveryServicesProvider":
		r = &ReplicationRecoveryServicesProvider{}
	case "azure-native:recoveryservices/v20221001:ReplicationStorageClassificationMapping":
		r = &ReplicationStorageClassificationMapping{}
	case "azure-native:recoveryservices/v20221001:ReplicationvCenter":
		r = &ReplicationvCenter{}
	case "azure-native:recoveryservices/v20221001:Vault":
		r = &Vault{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := pulumiazurenativesdk.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"recoveryservices/v20221001",
		&module{version},
	)
}
