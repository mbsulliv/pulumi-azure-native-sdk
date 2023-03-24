// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801preview

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
	case "azure-native:sql/v20220801preview:BackupShortTermRetentionPolicy":
		r = &BackupShortTermRetentionPolicy{}
	case "azure-native:sql/v20220801preview:DataMaskingPolicy":
		r = &DataMaskingPolicy{}
	case "azure-native:sql/v20220801preview:Database":
		r = &Database{}
	case "azure-native:sql/v20220801preview:DatabaseAdvisor":
		r = &DatabaseAdvisor{}
	case "azure-native:sql/v20220801preview:DatabaseBlobAuditingPolicy":
		r = &DatabaseBlobAuditingPolicy{}
	case "azure-native:sql/v20220801preview:DatabaseSecurityAlertPolicy":
		r = &DatabaseSecurityAlertPolicy{}
	case "azure-native:sql/v20220801preview:DatabaseSqlVulnerabilityAssessmentRuleBaseline":
		r = &DatabaseSqlVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20220801preview:DatabaseVulnerabilityAssessment":
		r = &DatabaseVulnerabilityAssessment{}
	case "azure-native:sql/v20220801preview:DatabaseVulnerabilityAssessmentRuleBaseline":
		r = &DatabaseVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20220801preview:DistributedAvailabilityGroup":
		r = &DistributedAvailabilityGroup{}
	case "azure-native:sql/v20220801preview:ElasticPool":
		r = &ElasticPool{}
	case "azure-native:sql/v20220801preview:EncryptionProtector":
		r = &EncryptionProtector{}
	case "azure-native:sql/v20220801preview:ExtendedDatabaseBlobAuditingPolicy":
		r = &ExtendedDatabaseBlobAuditingPolicy{}
	case "azure-native:sql/v20220801preview:ExtendedServerBlobAuditingPolicy":
		r = &ExtendedServerBlobAuditingPolicy{}
	case "azure-native:sql/v20220801preview:FailoverGroup":
		r = &FailoverGroup{}
	case "azure-native:sql/v20220801preview:FirewallRule":
		r = &FirewallRule{}
	case "azure-native:sql/v20220801preview:GeoBackupPolicy":
		r = &GeoBackupPolicy{}
	case "azure-native:sql/v20220801preview:IPv6FirewallRule":
		r = &IPv6FirewallRule{}
	case "azure-native:sql/v20220801preview:InstanceFailoverGroup":
		r = &InstanceFailoverGroup{}
	case "azure-native:sql/v20220801preview:InstancePool":
		r = &InstancePool{}
	case "azure-native:sql/v20220801preview:Job":
		r = &Job{}
	case "azure-native:sql/v20220801preview:JobAgent":
		r = &JobAgent{}
	case "azure-native:sql/v20220801preview:JobCredential":
		r = &JobCredential{}
	case "azure-native:sql/v20220801preview:JobStep":
		r = &JobStep{}
	case "azure-native:sql/v20220801preview:JobTargetGroup":
		r = &JobTargetGroup{}
	case "azure-native:sql/v20220801preview:LongTermRetentionPolicy":
		r = &LongTermRetentionPolicy{}
	case "azure-native:sql/v20220801preview:ManagedDatabase":
		r = &ManagedDatabase{}
	case "azure-native:sql/v20220801preview:ManagedDatabaseSensitivityLabel":
		r = &ManagedDatabaseSensitivityLabel{}
	case "azure-native:sql/v20220801preview:ManagedDatabaseVulnerabilityAssessment":
		r = &ManagedDatabaseVulnerabilityAssessment{}
	case "azure-native:sql/v20220801preview:ManagedDatabaseVulnerabilityAssessmentRuleBaseline":
		r = &ManagedDatabaseVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20220801preview:ManagedInstance":
		r = &ManagedInstance{}
	case "azure-native:sql/v20220801preview:ManagedInstanceAdministrator":
		r = &ManagedInstanceAdministrator{}
	case "azure-native:sql/v20220801preview:ManagedInstanceAzureADOnlyAuthentication":
		r = &ManagedInstanceAzureADOnlyAuthentication{}
	case "azure-native:sql/v20220801preview:ManagedInstanceKey":
		r = &ManagedInstanceKey{}
	case "azure-native:sql/v20220801preview:ManagedInstanceLongTermRetentionPolicy":
		r = &ManagedInstanceLongTermRetentionPolicy{}
	case "azure-native:sql/v20220801preview:ManagedInstancePrivateEndpointConnection":
		r = &ManagedInstancePrivateEndpointConnection{}
	case "azure-native:sql/v20220801preview:ManagedInstanceVulnerabilityAssessment":
		r = &ManagedInstanceVulnerabilityAssessment{}
	case "azure-native:sql/v20220801preview:ManagedServerDnsAlias":
		r = &ManagedServerDnsAlias{}
	case "azure-native:sql/v20220801preview:OutboundFirewallRule":
		r = &OutboundFirewallRule{}
	case "azure-native:sql/v20220801preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:sql/v20220801preview:SensitivityLabel":
		r = &SensitivityLabel{}
	case "azure-native:sql/v20220801preview:Server":
		r = &Server{}
	case "azure-native:sql/v20220801preview:ServerAdvisor":
		r = &ServerAdvisor{}
	case "azure-native:sql/v20220801preview:ServerAzureADAdministrator":
		r = &ServerAzureADAdministrator{}
	case "azure-native:sql/v20220801preview:ServerAzureADOnlyAuthentication":
		r = &ServerAzureADOnlyAuthentication{}
	case "azure-native:sql/v20220801preview:ServerBlobAuditingPolicy":
		r = &ServerBlobAuditingPolicy{}
	case "azure-native:sql/v20220801preview:ServerDnsAlias":
		r = &ServerDnsAlias{}
	case "azure-native:sql/v20220801preview:ServerKey":
		r = &ServerKey{}
	case "azure-native:sql/v20220801preview:ServerSecurityAlertPolicy":
		r = &ServerSecurityAlertPolicy{}
	case "azure-native:sql/v20220801preview:ServerTrustCertificate":
		r = &ServerTrustCertificate{}
	case "azure-native:sql/v20220801preview:ServerTrustGroup":
		r = &ServerTrustGroup{}
	case "azure-native:sql/v20220801preview:ServerVulnerabilityAssessment":
		r = &ServerVulnerabilityAssessment{}
	case "azure-native:sql/v20220801preview:SqlVulnerabilityAssessmentRuleBaseline":
		r = &SqlVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20220801preview:SqlVulnerabilityAssessmentsSetting":
		r = &SqlVulnerabilityAssessmentsSetting{}
	case "azure-native:sql/v20220801preview:StartStopManagedInstanceSchedule":
		r = &StartStopManagedInstanceSchedule{}
	case "azure-native:sql/v20220801preview:SyncAgent":
		r = &SyncAgent{}
	case "azure-native:sql/v20220801preview:SyncGroup":
		r = &SyncGroup{}
	case "azure-native:sql/v20220801preview:SyncMember":
		r = &SyncMember{}
	case "azure-native:sql/v20220801preview:TransparentDataEncryption":
		r = &TransparentDataEncryption{}
	case "azure-native:sql/v20220801preview:VirtualNetworkRule":
		r = &VirtualNetworkRule{}
	case "azure-native:sql/v20220801preview:WorkloadClassifier":
		r = &WorkloadClassifier{}
	case "azure-native:sql/v20220801preview:WorkloadGroup":
		r = &WorkloadGroup{}
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
		"sql/v20220801preview",
		&module{version},
	)
}