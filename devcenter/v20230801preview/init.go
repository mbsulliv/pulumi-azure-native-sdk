// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

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
	case "azure-native:devcenter/v20230801preview:AttachedNetworkByDevCenter":
		r = &AttachedNetworkByDevCenter{}
	case "azure-native:devcenter/v20230801preview:Catalog":
		r = &Catalog{}
	case "azure-native:devcenter/v20230801preview:DevBoxDefinition":
		r = &DevBoxDefinition{}
	case "azure-native:devcenter/v20230801preview:DevCenter":
		r = &DevCenter{}
	case "azure-native:devcenter/v20230801preview:EnvironmentType":
		r = &EnvironmentType{}
	case "azure-native:devcenter/v20230801preview:Gallery":
		r = &Gallery{}
	case "azure-native:devcenter/v20230801preview:NetworkConnection":
		r = &NetworkConnection{}
	case "azure-native:devcenter/v20230801preview:Pool":
		r = &Pool{}
	case "azure-native:devcenter/v20230801preview:Project":
		r = &Project{}
	case "azure-native:devcenter/v20230801preview:ProjectEnvironmentType":
		r = &ProjectEnvironmentType{}
	case "azure-native:devcenter/v20230801preview:Schedule":
		r = &Schedule{}
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
		"devcenter/v20230801preview",
		&module{version},
	)
}
