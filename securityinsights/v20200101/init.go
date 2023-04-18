// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101

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
	case "azure-native:securityinsights/v20200101:AADDataConnector":
		r = &AADDataConnector{}
	case "azure-native:securityinsights/v20200101:AATPDataConnector":
		r = &AATPDataConnector{}
	case "azure-native:securityinsights/v20200101:ASCDataConnector":
		r = &ASCDataConnector{}
	case "azure-native:securityinsights/v20200101:Action":
		r = &Action{}
	case "azure-native:securityinsights/v20200101:AwsCloudTrailDataConnector":
		r = &AwsCloudTrailDataConnector{}
	case "azure-native:securityinsights/v20200101:Bookmark":
		r = &Bookmark{}
	case "azure-native:securityinsights/v20200101:FusionAlertRule":
		r = &FusionAlertRule{}
	case "azure-native:securityinsights/v20200101:Incident":
		r = &Incident{}
	case "azure-native:securityinsights/v20200101:MCASDataConnector":
		r = &MCASDataConnector{}
	case "azure-native:securityinsights/v20200101:MDATPDataConnector":
		r = &MDATPDataConnector{}
	case "azure-native:securityinsights/v20200101:MicrosoftSecurityIncidentCreationAlertRule":
		r = &MicrosoftSecurityIncidentCreationAlertRule{}
	case "azure-native:securityinsights/v20200101:OfficeDataConnector":
		r = &OfficeDataConnector{}
	case "azure-native:securityinsights/v20200101:ScheduledAlertRule":
		r = &ScheduledAlertRule{}
	case "azure-native:securityinsights/v20200101:TIDataConnector":
		r = &TIDataConnector{}
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
		"securityinsights/v20200101",
		&module{version},
	)
}
