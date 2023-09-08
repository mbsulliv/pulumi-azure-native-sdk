// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230606

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
	case "azure-native:offazure/v20230606:HypervClusterControllerCluster":
		r = &HypervClusterControllerCluster{}
	case "azure-native:offazure/v20230606:HypervHostController":
		r = &HypervHostController{}
	case "azure-native:offazure/v20230606:HypervSitesController":
		r = &HypervSitesController{}
	case "azure-native:offazure/v20230606:ImportSitesController":
		r = &ImportSitesController{}
	case "azure-native:offazure/v20230606:MasterSitesController":
		r = &MasterSitesController{}
	case "azure-native:offazure/v20230606:PrivateEndpointConnectionController":
		r = &PrivateEndpointConnectionController{}
	case "azure-native:offazure/v20230606:ServerSitesController":
		r = &ServerSitesController{}
	case "azure-native:offazure/v20230606:SitesController":
		r = &SitesController{}
	case "azure-native:offazure/v20230606:SqlDiscoverySiteDataSourceController":
		r = &SqlDiscoverySiteDataSourceController{}
	case "azure-native:offazure/v20230606:SqlSitesController":
		r = &SqlSitesController{}
	case "azure-native:offazure/v20230606:VcenterController":
		r = &VcenterController{}
	case "azure-native:offazure/v20230606:WebAppDiscoverySiteDataSourcesController":
		r = &WebAppDiscoverySiteDataSourcesController{}
	case "azure-native:offazure/v20230606:WebAppSitesController":
		r = &WebAppSitesController{}
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
		"offazure/v20230606",
		&module{version},
	)
}