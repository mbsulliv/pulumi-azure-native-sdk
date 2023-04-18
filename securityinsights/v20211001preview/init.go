// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211001preview

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
	case "azure-native:securityinsights/v20211001preview:AADDataConnector":
		r = &AADDataConnector{}
	case "azure-native:securityinsights/v20211001preview:AATPDataConnector":
		r = &AATPDataConnector{}
	case "azure-native:securityinsights/v20211001preview:ASCDataConnector":
		r = &ASCDataConnector{}
	case "azure-native:securityinsights/v20211001preview:Action":
		r = &Action{}
	case "azure-native:securityinsights/v20211001preview:ActivityCustomEntityQuery":
		r = &ActivityCustomEntityQuery{}
	case "azure-native:securityinsights/v20211001preview:Anomalies":
		r = &Anomalies{}
	case "azure-native:securityinsights/v20211001preview:AutomationRule":
		r = &AutomationRule{}
	case "azure-native:securityinsights/v20211001preview:AwsCloudTrailDataConnector":
		r = &AwsCloudTrailDataConnector{}
	case "azure-native:securityinsights/v20211001preview:AwsS3DataConnector":
		r = &AwsS3DataConnector{}
	case "azure-native:securityinsights/v20211001preview:Bookmark":
		r = &Bookmark{}
	case "azure-native:securityinsights/v20211001preview:BookmarkRelation":
		r = &BookmarkRelation{}
	case "azure-native:securityinsights/v20211001preview:CodelessApiPollingDataConnector":
		r = &CodelessApiPollingDataConnector{}
	case "azure-native:securityinsights/v20211001preview:CodelessUiDataConnector":
		r = &CodelessUiDataConnector{}
	case "azure-native:securityinsights/v20211001preview:Dynamics365DataConnector":
		r = &Dynamics365DataConnector{}
	case "azure-native:securityinsights/v20211001preview:EntityAnalytics":
		r = &EntityAnalytics{}
	case "azure-native:securityinsights/v20211001preview:EyesOn":
		r = &EyesOn{}
	case "azure-native:securityinsights/v20211001preview:FusionAlertRule":
		r = &FusionAlertRule{}
	case "azure-native:securityinsights/v20211001preview:Incident":
		r = &Incident{}
	case "azure-native:securityinsights/v20211001preview:IncidentComment":
		r = &IncidentComment{}
	case "azure-native:securityinsights/v20211001preview:IncidentRelation":
		r = &IncidentRelation{}
	case "azure-native:securityinsights/v20211001preview:MCASDataConnector":
		r = &MCASDataConnector{}
	case "azure-native:securityinsights/v20211001preview:MDATPDataConnector":
		r = &MDATPDataConnector{}
	case "azure-native:securityinsights/v20211001preview:MLBehaviorAnalyticsAlertRule":
		r = &MLBehaviorAnalyticsAlertRule{}
	case "azure-native:securityinsights/v20211001preview:MSTIDataConnector":
		r = &MSTIDataConnector{}
	case "azure-native:securityinsights/v20211001preview:MTPDataConnector":
		r = &MTPDataConnector{}
	case "azure-native:securityinsights/v20211001preview:Metadata":
		r = &Metadata{}
	case "azure-native:securityinsights/v20211001preview:MicrosoftSecurityIncidentCreationAlertRule":
		r = &MicrosoftSecurityIncidentCreationAlertRule{}
	case "azure-native:securityinsights/v20211001preview:NrtAlertRule":
		r = &NrtAlertRule{}
	case "azure-native:securityinsights/v20211001preview:Office365ProjectDataConnector":
		r = &Office365ProjectDataConnector{}
	case "azure-native:securityinsights/v20211001preview:OfficeATPDataConnector":
		r = &OfficeATPDataConnector{}
	case "azure-native:securityinsights/v20211001preview:OfficeDataConnector":
		r = &OfficeDataConnector{}
	case "azure-native:securityinsights/v20211001preview:OfficeIRMDataConnector":
		r = &OfficeIRMDataConnector{}
	case "azure-native:securityinsights/v20211001preview:OfficePowerBIDataConnector":
		r = &OfficePowerBIDataConnector{}
	case "azure-native:securityinsights/v20211001preview:ScheduledAlertRule":
		r = &ScheduledAlertRule{}
	case "azure-native:securityinsights/v20211001preview:SentinelOnboardingState":
		r = &SentinelOnboardingState{}
	case "azure-native:securityinsights/v20211001preview:SourceControl":
		r = &SourceControl{}
	case "azure-native:securityinsights/v20211001preview:TIDataConnector":
		r = &TIDataConnector{}
	case "azure-native:securityinsights/v20211001preview:ThreatIntelligenceAlertRule":
		r = &ThreatIntelligenceAlertRule{}
	case "azure-native:securityinsights/v20211001preview:ThreatIntelligenceIndicator":
		r = &ThreatIntelligenceIndicator{}
	case "azure-native:securityinsights/v20211001preview:TiTaxiiDataConnector":
		r = &TiTaxiiDataConnector{}
	case "azure-native:securityinsights/v20211001preview:Ueba":
		r = &Ueba{}
	case "azure-native:securityinsights/v20211001preview:Watchlist":
		r = &Watchlist{}
	case "azure-native:securityinsights/v20211001preview:WatchlistItem":
		r = &WatchlistItem{}
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
		"securityinsights/v20211001preview",
		&module{version},
	)
}
