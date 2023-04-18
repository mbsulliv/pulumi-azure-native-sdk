// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

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
	case "azure-native:securityinsights/v20230201preview:AADDataConnector":
		r = &AADDataConnector{}
	case "azure-native:securityinsights/v20230201preview:AATPDataConnector":
		r = &AATPDataConnector{}
	case "azure-native:securityinsights/v20230201preview:ASCDataConnector":
		r = &ASCDataConnector{}
	case "azure-native:securityinsights/v20230201preview:Action":
		r = &Action{}
	case "azure-native:securityinsights/v20230201preview:ActivityCustomEntityQuery":
		r = &ActivityCustomEntityQuery{}
	case "azure-native:securityinsights/v20230201preview:Anomalies":
		r = &Anomalies{}
	case "azure-native:securityinsights/v20230201preview:AnomalySecurityMLAnalyticsSettings":
		r = &AnomalySecurityMLAnalyticsSettings{}
	case "azure-native:securityinsights/v20230201preview:AutomationRule":
		r = &AutomationRule{}
	case "azure-native:securityinsights/v20230201preview:AwsCloudTrailDataConnector":
		r = &AwsCloudTrailDataConnector{}
	case "azure-native:securityinsights/v20230201preview:AwsS3DataConnector":
		r = &AwsS3DataConnector{}
	case "azure-native:securityinsights/v20230201preview:Bookmark":
		r = &Bookmark{}
	case "azure-native:securityinsights/v20230201preview:BookmarkRelation":
		r = &BookmarkRelation{}
	case "azure-native:securityinsights/v20230201preview:CodelessApiPollingDataConnector":
		r = &CodelessApiPollingDataConnector{}
	case "azure-native:securityinsights/v20230201preview:CodelessUiDataConnector":
		r = &CodelessUiDataConnector{}
	case "azure-native:securityinsights/v20230201preview:Dynamics365DataConnector":
		r = &Dynamics365DataConnector{}
	case "azure-native:securityinsights/v20230201preview:EntityAnalytics":
		r = &EntityAnalytics{}
	case "azure-native:securityinsights/v20230201preview:EyesOn":
		r = &EyesOn{}
	case "azure-native:securityinsights/v20230201preview:FileImport":
		r = &FileImport{}
	case "azure-native:securityinsights/v20230201preview:FusionAlertRule":
		r = &FusionAlertRule{}
	case "azure-native:securityinsights/v20230201preview:Incident":
		r = &Incident{}
	case "azure-native:securityinsights/v20230201preview:IncidentComment":
		r = &IncidentComment{}
	case "azure-native:securityinsights/v20230201preview:IncidentRelation":
		r = &IncidentRelation{}
	case "azure-native:securityinsights/v20230201preview:IncidentTask":
		r = &IncidentTask{}
	case "azure-native:securityinsights/v20230201preview:IoTDataConnector":
		r = &IoTDataConnector{}
	case "azure-native:securityinsights/v20230201preview:MCASDataConnector":
		r = &MCASDataConnector{}
	case "azure-native:securityinsights/v20230201preview:MDATPDataConnector":
		r = &MDATPDataConnector{}
	case "azure-native:securityinsights/v20230201preview:MLBehaviorAnalyticsAlertRule":
		r = &MLBehaviorAnalyticsAlertRule{}
	case "azure-native:securityinsights/v20230201preview:MSTIDataConnector":
		r = &MSTIDataConnector{}
	case "azure-native:securityinsights/v20230201preview:MTPDataConnector":
		r = &MTPDataConnector{}
	case "azure-native:securityinsights/v20230201preview:Metadata":
		r = &Metadata{}
	case "azure-native:securityinsights/v20230201preview:MicrosoftPurviewInformationProtectionDataConnector":
		r = &MicrosoftPurviewInformationProtectionDataConnector{}
	case "azure-native:securityinsights/v20230201preview:MicrosoftSecurityIncidentCreationAlertRule":
		r = &MicrosoftSecurityIncidentCreationAlertRule{}
	case "azure-native:securityinsights/v20230201preview:NrtAlertRule":
		r = &NrtAlertRule{}
	case "azure-native:securityinsights/v20230201preview:Office365ProjectDataConnector":
		r = &Office365ProjectDataConnector{}
	case "azure-native:securityinsights/v20230201preview:OfficeATPDataConnector":
		r = &OfficeATPDataConnector{}
	case "azure-native:securityinsights/v20230201preview:OfficeDataConnector":
		r = &OfficeDataConnector{}
	case "azure-native:securityinsights/v20230201preview:OfficeIRMDataConnector":
		r = &OfficeIRMDataConnector{}
	case "azure-native:securityinsights/v20230201preview:OfficePowerBIDataConnector":
		r = &OfficePowerBIDataConnector{}
	case "azure-native:securityinsights/v20230201preview:ScheduledAlertRule":
		r = &ScheduledAlertRule{}
	case "azure-native:securityinsights/v20230201preview:SentinelOnboardingState":
		r = &SentinelOnboardingState{}
	case "azure-native:securityinsights/v20230201preview:SourceControl":
		r = &SourceControl{}
	case "azure-native:securityinsights/v20230201preview:TIDataConnector":
		r = &TIDataConnector{}
	case "azure-native:securityinsights/v20230201preview:ThreatIntelligenceAlertRule":
		r = &ThreatIntelligenceAlertRule{}
	case "azure-native:securityinsights/v20230201preview:ThreatIntelligenceIndicator":
		r = &ThreatIntelligenceIndicator{}
	case "azure-native:securityinsights/v20230201preview:TiTaxiiDataConnector":
		r = &TiTaxiiDataConnector{}
	case "azure-native:securityinsights/v20230201preview:Ueba":
		r = &Ueba{}
	case "azure-native:securityinsights/v20230201preview:Watchlist":
		r = &Watchlist{}
	case "azure-native:securityinsights/v20230201preview:WatchlistItem":
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
		"securityinsights/v20230201preview",
		&module{version},
	)
}
