// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240201

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
	case "azure-native:cdn/v20240201:AFDCustomDomain":
		r = &AFDCustomDomain{}
	case "azure-native:cdn/v20240201:AFDEndpoint":
		r = &AFDEndpoint{}
	case "azure-native:cdn/v20240201:AFDOrigin":
		r = &AFDOrigin{}
	case "azure-native:cdn/v20240201:AFDOriginGroup":
		r = &AFDOriginGroup{}
	case "azure-native:cdn/v20240201:CustomDomain":
		r = &CustomDomain{}
	case "azure-native:cdn/v20240201:Endpoint":
		r = &Endpoint{}
	case "azure-native:cdn/v20240201:Origin":
		r = &Origin{}
	case "azure-native:cdn/v20240201:OriginGroup":
		r = &OriginGroup{}
	case "azure-native:cdn/v20240201:Policy":
		r = &Policy{}
	case "azure-native:cdn/v20240201:Profile":
		r = &Profile{}
	case "azure-native:cdn/v20240201:Route":
		r = &Route{}
	case "azure-native:cdn/v20240201:Rule":
		r = &Rule{}
	case "azure-native:cdn/v20240201:RuleSet":
		r = &RuleSet{}
	case "azure-native:cdn/v20240201:Secret":
		r = &Secret{}
	case "azure-native:cdn/v20240201:SecurityPolicy":
		r = &SecurityPolicy{}
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
		"cdn/v20240201",
		&module{version},
	)
}
