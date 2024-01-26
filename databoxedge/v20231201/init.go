// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231201

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
	case "azure-native:databoxedge/v20231201:ArcAddon":
		r = &ArcAddon{}
	case "azure-native:databoxedge/v20231201:BandwidthSchedule":
		r = &BandwidthSchedule{}
	case "azure-native:databoxedge/v20231201:CloudEdgeManagementRole":
		r = &CloudEdgeManagementRole{}
	case "azure-native:databoxedge/v20231201:Container":
		r = &Container{}
	case "azure-native:databoxedge/v20231201:Device":
		r = &Device{}
	case "azure-native:databoxedge/v20231201:FileEventTrigger":
		r = &FileEventTrigger{}
	case "azure-native:databoxedge/v20231201:IoTAddon":
		r = &IoTAddon{}
	case "azure-native:databoxedge/v20231201:IoTRole":
		r = &IoTRole{}
	case "azure-native:databoxedge/v20231201:KubernetesRole":
		r = &KubernetesRole{}
	case "azure-native:databoxedge/v20231201:MECRole":
		r = &MECRole{}
	case "azure-native:databoxedge/v20231201:MonitoringConfig":
		r = &MonitoringConfig{}
	case "azure-native:databoxedge/v20231201:Order":
		r = &Order{}
	case "azure-native:databoxedge/v20231201:PeriodicTimerEventTrigger":
		r = &PeriodicTimerEventTrigger{}
	case "azure-native:databoxedge/v20231201:Share":
		r = &Share{}
	case "azure-native:databoxedge/v20231201:StorageAccount":
		r = &StorageAccount{}
	case "azure-native:databoxedge/v20231201:StorageAccountCredential":
		r = &StorageAccountCredential{}
	case "azure-native:databoxedge/v20231201:User":
		r = &User{}
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
		"databoxedge/v20231201",
		&module{version},
	)
}
