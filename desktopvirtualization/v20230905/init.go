// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230905

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
	case "azure-native:desktopvirtualization/v20230905:Application":
		r = &Application{}
	case "azure-native:desktopvirtualization/v20230905:ApplicationGroup":
		r = &ApplicationGroup{}
	case "azure-native:desktopvirtualization/v20230905:HostPool":
		r = &HostPool{}
	case "azure-native:desktopvirtualization/v20230905:MSIXPackage":
		r = &MSIXPackage{}
	case "azure-native:desktopvirtualization/v20230905:PrivateEndpointConnectionByHostPool":
		r = &PrivateEndpointConnectionByHostPool{}
	case "azure-native:desktopvirtualization/v20230905:PrivateEndpointConnectionByWorkspace":
		r = &PrivateEndpointConnectionByWorkspace{}
	case "azure-native:desktopvirtualization/v20230905:ScalingPlan":
		r = &ScalingPlan{}
	case "azure-native:desktopvirtualization/v20230905:ScalingPlanPersonalSchedule":
		r = &ScalingPlanPersonalSchedule{}
	case "azure-native:desktopvirtualization/v20230905:ScalingPlanPooledSchedule":
		r = &ScalingPlanPooledSchedule{}
	case "azure-native:desktopvirtualization/v20230905:Workspace":
		r = &Workspace{}
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
		"desktopvirtualization/v20230905",
		&module{version},
	)
}
