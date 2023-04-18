// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

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
	case "azure-native:databoxedge/v20190701:BandwidthSchedule":
		r = &BandwidthSchedule{}
	case "azure-native:databoxedge/v20190701:Device":
		r = &Device{}
	case "azure-native:databoxedge/v20190701:FileEventTrigger":
		r = &FileEventTrigger{}
	case "azure-native:databoxedge/v20190701:IoTRole":
		r = &IoTRole{}
	case "azure-native:databoxedge/v20190701:Order":
		r = &Order{}
	case "azure-native:databoxedge/v20190701:PeriodicTimerEventTrigger":
		r = &PeriodicTimerEventTrigger{}
	case "azure-native:databoxedge/v20190701:Share":
		r = &Share{}
	case "azure-native:databoxedge/v20190701:StorageAccountCredential":
		r = &StorageAccountCredential{}
	case "azure-native:databoxedge/v20190701:User":
		r = &User{}
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
		"databoxedge/v20190701",
		&module{version},
	)
}
