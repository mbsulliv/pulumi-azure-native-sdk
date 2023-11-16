// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231003preview

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
	case "azure-native:hybridcompute/v20231003preview:License":
		r = &License{}
	case "azure-native:hybridcompute/v20231003preview:LicenseProfile":
		r = &LicenseProfile{}
	case "azure-native:hybridcompute/v20231003preview:Machine":
		r = &Machine{}
	case "azure-native:hybridcompute/v20231003preview:MachineExtension":
		r = &MachineExtension{}
	case "azure-native:hybridcompute/v20231003preview:MachineRunCommand":
		r = &MachineRunCommand{}
	case "azure-native:hybridcompute/v20231003preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:hybridcompute/v20231003preview:PrivateLinkScope":
		r = &PrivateLinkScope{}
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
		"hybridcompute/v20231003preview",
		&module{version},
	)
}
