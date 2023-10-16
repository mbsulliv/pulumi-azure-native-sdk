// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

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
	case "azure-native:elasticsan/v20230101:ElasticSan":
		r = &ElasticSan{}
	case "azure-native:elasticsan/v20230101:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:elasticsan/v20230101:Volume":
		r = &Volume{}
	case "azure-native:elasticsan/v20230101:VolumeGroup":
		r = &VolumeGroup{}
	case "azure-native:elasticsan/v20230101:VolumeSnapshot":
		r = &VolumeSnapshot{}
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
		"elasticsan/v20230101",
		&module{version},
	)
}