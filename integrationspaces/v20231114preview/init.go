// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231114preview

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
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
	case "azure-native:integrationspaces/v20231114preview:Application":
		r = &Application{}
	case "azure-native:integrationspaces/v20231114preview:ApplicationResource":
		r = &ApplicationResource{}
	case "azure-native:integrationspaces/v20231114preview:BusinessProcess":
		r = &BusinessProcess{}
	case "azure-native:integrationspaces/v20231114preview:InfrastructureResource":
		r = &InfrastructureResource{}
	case "azure-native:integrationspaces/v20231114preview:Space":
		r = &Space{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := utilities.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"integrationspaces/v20231114preview",
		&module{version},
	)
}