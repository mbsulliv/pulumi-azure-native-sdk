// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240215preview

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
	case "azure-native:azurestackhci/v20240215preview:ArcSetting":
		r = &ArcSetting{}
	case "azure-native:azurestackhci/v20240215preview:Cluster":
		r = &Cluster{}
	case "azure-native:azurestackhci/v20240215preview:DeploymentSetting":
		r = &DeploymentSetting{}
	case "azure-native:azurestackhci/v20240215preview:Extension":
		r = &Extension{}
	case "azure-native:azurestackhci/v20240215preview:HciEdgeDevice":
		r = &HciEdgeDevice{}
	case "azure-native:azurestackhci/v20240215preview:SecuritySetting":
		r = &SecuritySetting{}
	case "azure-native:azurestackhci/v20240215preview:Update":
		r = &Update{}
	case "azure-native:azurestackhci/v20240215preview:UpdateRun":
		r = &UpdateRun{}
	case "azure-native:azurestackhci/v20240215preview:UpdateSummary":
		r = &UpdateSummary{}
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
		"azurestackhci/v20240215preview",
		&module{version},
	)
}
