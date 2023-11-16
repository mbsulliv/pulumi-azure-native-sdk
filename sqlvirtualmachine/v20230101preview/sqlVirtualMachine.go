// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A SQL virtual machine.
type SqlVirtualMachine struct {
	pulumi.CustomResourceState

	// SQL best practices Assessment Settings.
	AssessmentSettings AssessmentSettingsResponsePtrOutput `pulumi:"assessmentSettings"`
	// Auto backup settings for SQL Server.
	AutoBackupSettings AutoBackupSettingsResponsePtrOutput `pulumi:"autoBackupSettings"`
	// Auto patching settings for applying critical security updates to SQL virtual machine.
	AutoPatchingSettings AutoPatchingSettingsResponsePtrOutput `pulumi:"autoPatchingSettings"`
	// Enable automatic upgrade of Sql IaaS extension Agent.
	EnableAutomaticUpgrade pulumi.BoolPtrOutput `pulumi:"enableAutomaticUpgrade"`
	// DO NOT USE. This value will be deprecated. Azure Active Directory identity of the server.
	Identity ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	// Key vault credential settings.
	KeyVaultCredentialSettings KeyVaultCredentialSettingsResponsePtrOutput `pulumi:"keyVaultCredentialSettings"`
	// SQL IaaS Agent least privilege mode.
	LeastPrivilegeMode pulumi.StringPtrOutput `pulumi:"leastPrivilegeMode"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state to track the async operation status.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// SQL Server configuration management settings.
	ServerConfigurationsManagementSettings ServerConfigurationsManagementSettingsResponsePtrOutput `pulumi:"serverConfigurationsManagementSettings"`
	// SQL image offer. Examples include SQL2016-WS2016, SQL2017-WS2016.
	SqlImageOffer pulumi.StringPtrOutput `pulumi:"sqlImageOffer"`
	// SQL Server edition type.
	SqlImageSku pulumi.StringPtrOutput `pulumi:"sqlImageSku"`
	// SQL Server Management type. NOTE: This parameter is not used anymore. API will automatically detect the Sql Management, refrain from using it.
	SqlManagement pulumi.StringPtrOutput `pulumi:"sqlManagement"`
	// SQL Server license type.
	SqlServerLicenseType pulumi.StringPtrOutput `pulumi:"sqlServerLicenseType"`
	// ARM resource id of the SQL virtual machine group this SQL virtual machine is or will be part of.
	SqlVirtualMachineGroupResourceId pulumi.StringPtrOutput `pulumi:"sqlVirtualMachineGroupResourceId"`
	// Storage Configuration Settings.
	StorageConfigurationSettings StorageConfigurationSettingsResponsePtrOutput `pulumi:"storageConfigurationSettings"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Troubleshooting status
	TroubleshootingStatus TroubleshootingStatusResponseOutput `pulumi:"troubleshootingStatus"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// ARM Resource id of underlying virtual machine created from SQL marketplace image.
	VirtualMachineResourceId pulumi.StringPtrOutput `pulumi:"virtualMachineResourceId"`
	// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
	WsfcDomainCredentials WsfcDomainCredentialsResponsePtrOutput `pulumi:"wsfcDomainCredentials"`
	// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
	WsfcStaticIp pulumi.StringPtrOutput `pulumi:"wsfcStaticIp"`
}

// NewSqlVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewSqlVirtualMachine(ctx *pulumi.Context,
	name string, args *SqlVirtualMachineArgs, opts ...pulumi.ResourceOption) (*SqlVirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.AutoPatchingSettings != nil {
		args.AutoPatchingSettings = args.AutoPatchingSettings.ToAutoPatchingSettingsPtrOutput().ApplyT(func(v *AutoPatchingSettings) *AutoPatchingSettings { return v.Defaults() }).(AutoPatchingSettingsPtrOutput)
	}
	if args.EnableAutomaticUpgrade == nil {
		args.EnableAutomaticUpgrade = pulumi.BoolPtr(false)
	}
	if args.LeastPrivilegeMode == nil {
		args.LeastPrivilegeMode = pulumi.StringPtr("NotSet")
	}
	if args.StorageConfigurationSettings != nil {
		args.StorageConfigurationSettings = args.StorageConfigurationSettings.ToStorageConfigurationSettingsPtrOutput().ApplyT(func(v *StorageConfigurationSettings) *StorageConfigurationSettings { return v.Defaults() }).(StorageConfigurationSettingsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine:SqlVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20170301preview:SqlVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20211101preview:SqlVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220201:SqlVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220201preview:SqlVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220701preview:SqlVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220801preview:SqlVirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SqlVirtualMachine
	err := ctx.RegisterResource("azure-native:sqlvirtualmachine/v20230101preview:SqlVirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlVirtualMachine gets an existing SqlVirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlVirtualMachineState, opts ...pulumi.ResourceOption) (*SqlVirtualMachine, error) {
	var resource SqlVirtualMachine
	err := ctx.ReadResource("azure-native:sqlvirtualmachine/v20230101preview:SqlVirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlVirtualMachine resources.
type sqlVirtualMachineState struct {
}

type SqlVirtualMachineState struct {
}

func (SqlVirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlVirtualMachineState)(nil)).Elem()
}

type sqlVirtualMachineArgs struct {
	// SQL best practices Assessment Settings.
	AssessmentSettings *AssessmentSettings `pulumi:"assessmentSettings"`
	// Auto backup settings for SQL Server.
	AutoBackupSettings *AutoBackupSettings `pulumi:"autoBackupSettings"`
	// Auto patching settings for applying critical security updates to SQL virtual machine.
	AutoPatchingSettings *AutoPatchingSettings `pulumi:"autoPatchingSettings"`
	// Enable automatic upgrade of Sql IaaS extension Agent.
	EnableAutomaticUpgrade *bool `pulumi:"enableAutomaticUpgrade"`
	// DO NOT USE. This value will be deprecated. Azure Active Directory identity of the server.
	Identity *ResourceIdentity `pulumi:"identity"`
	// Key vault credential settings.
	KeyVaultCredentialSettings *KeyVaultCredentialSettings `pulumi:"keyVaultCredentialSettings"`
	// SQL IaaS Agent least privilege mode.
	LeastPrivilegeMode *string `pulumi:"leastPrivilegeMode"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SQL Server configuration management settings.
	ServerConfigurationsManagementSettings *ServerConfigurationsManagementSettings `pulumi:"serverConfigurationsManagementSettings"`
	// SQL image offer. Examples include SQL2016-WS2016, SQL2017-WS2016.
	SqlImageOffer *string `pulumi:"sqlImageOffer"`
	// SQL Server edition type.
	SqlImageSku *string `pulumi:"sqlImageSku"`
	// SQL Server Management type. NOTE: This parameter is not used anymore. API will automatically detect the Sql Management, refrain from using it.
	SqlManagement *string `pulumi:"sqlManagement"`
	// SQL Server license type.
	SqlServerLicenseType *string `pulumi:"sqlServerLicenseType"`
	// ARM resource id of the SQL virtual machine group this SQL virtual machine is or will be part of.
	SqlVirtualMachineGroupResourceId *string `pulumi:"sqlVirtualMachineGroupResourceId"`
	// Name of the SQL virtual machine.
	SqlVirtualMachineName *string `pulumi:"sqlVirtualMachineName"`
	// Storage Configuration Settings.
	StorageConfigurationSettings *StorageConfigurationSettings `pulumi:"storageConfigurationSettings"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// ARM Resource id of underlying virtual machine created from SQL marketplace image.
	VirtualMachineResourceId *string `pulumi:"virtualMachineResourceId"`
	// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
	WsfcDomainCredentials *WsfcDomainCredentials `pulumi:"wsfcDomainCredentials"`
	// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
	WsfcStaticIp *string `pulumi:"wsfcStaticIp"`
}

// The set of arguments for constructing a SqlVirtualMachine resource.
type SqlVirtualMachineArgs struct {
	// SQL best practices Assessment Settings.
	AssessmentSettings AssessmentSettingsPtrInput
	// Auto backup settings for SQL Server.
	AutoBackupSettings AutoBackupSettingsPtrInput
	// Auto patching settings for applying critical security updates to SQL virtual machine.
	AutoPatchingSettings AutoPatchingSettingsPtrInput
	// Enable automatic upgrade of Sql IaaS extension Agent.
	EnableAutomaticUpgrade pulumi.BoolPtrInput
	// DO NOT USE. This value will be deprecated. Azure Active Directory identity of the server.
	Identity ResourceIdentityPtrInput
	// Key vault credential settings.
	KeyVaultCredentialSettings KeyVaultCredentialSettingsPtrInput
	// SQL IaaS Agent least privilege mode.
	LeastPrivilegeMode pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// SQL Server configuration management settings.
	ServerConfigurationsManagementSettings ServerConfigurationsManagementSettingsPtrInput
	// SQL image offer. Examples include SQL2016-WS2016, SQL2017-WS2016.
	SqlImageOffer pulumi.StringPtrInput
	// SQL Server edition type.
	SqlImageSku pulumi.StringPtrInput
	// SQL Server Management type. NOTE: This parameter is not used anymore. API will automatically detect the Sql Management, refrain from using it.
	SqlManagement pulumi.StringPtrInput
	// SQL Server license type.
	SqlServerLicenseType pulumi.StringPtrInput
	// ARM resource id of the SQL virtual machine group this SQL virtual machine is or will be part of.
	SqlVirtualMachineGroupResourceId pulumi.StringPtrInput
	// Name of the SQL virtual machine.
	SqlVirtualMachineName pulumi.StringPtrInput
	// Storage Configuration Settings.
	StorageConfigurationSettings StorageConfigurationSettingsPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// ARM Resource id of underlying virtual machine created from SQL marketplace image.
	VirtualMachineResourceId pulumi.StringPtrInput
	// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
	WsfcDomainCredentials WsfcDomainCredentialsPtrInput
	// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
	WsfcStaticIp pulumi.StringPtrInput
}

func (SqlVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlVirtualMachineArgs)(nil)).Elem()
}

type SqlVirtualMachineInput interface {
	pulumi.Input

	ToSqlVirtualMachineOutput() SqlVirtualMachineOutput
	ToSqlVirtualMachineOutputWithContext(ctx context.Context) SqlVirtualMachineOutput
}

func (*SqlVirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlVirtualMachine)(nil)).Elem()
}

func (i *SqlVirtualMachine) ToSqlVirtualMachineOutput() SqlVirtualMachineOutput {
	return i.ToSqlVirtualMachineOutputWithContext(context.Background())
}

func (i *SqlVirtualMachine) ToSqlVirtualMachineOutputWithContext(ctx context.Context) SqlVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlVirtualMachineOutput)
}

type SqlVirtualMachineOutput struct{ *pulumi.OutputState }

func (SqlVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlVirtualMachine)(nil)).Elem()
}

func (o SqlVirtualMachineOutput) ToSqlVirtualMachineOutput() SqlVirtualMachineOutput {
	return o
}

func (o SqlVirtualMachineOutput) ToSqlVirtualMachineOutputWithContext(ctx context.Context) SqlVirtualMachineOutput {
	return o
}

// SQL best practices Assessment Settings.
func (o SqlVirtualMachineOutput) AssessmentSettings() AssessmentSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) AssessmentSettingsResponsePtrOutput { return v.AssessmentSettings }).(AssessmentSettingsResponsePtrOutput)
}

// Auto backup settings for SQL Server.
func (o SqlVirtualMachineOutput) AutoBackupSettings() AutoBackupSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) AutoBackupSettingsResponsePtrOutput { return v.AutoBackupSettings }).(AutoBackupSettingsResponsePtrOutput)
}

// Auto patching settings for applying critical security updates to SQL virtual machine.
func (o SqlVirtualMachineOutput) AutoPatchingSettings() AutoPatchingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) AutoPatchingSettingsResponsePtrOutput { return v.AutoPatchingSettings }).(AutoPatchingSettingsResponsePtrOutput)
}

// Enable automatic upgrade of Sql IaaS extension Agent.
func (o SqlVirtualMachineOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.BoolPtrOutput { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

// DO NOT USE. This value will be deprecated. Azure Active Directory identity of the server.
func (o SqlVirtualMachineOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// Key vault credential settings.
func (o SqlVirtualMachineOutput) KeyVaultCredentialSettings() KeyVaultCredentialSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) KeyVaultCredentialSettingsResponsePtrOutput {
		return v.KeyVaultCredentialSettings
	}).(KeyVaultCredentialSettingsResponsePtrOutput)
}

// SQL IaaS Agent least privilege mode.
func (o SqlVirtualMachineOutput) LeastPrivilegeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.LeastPrivilegeMode }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o SqlVirtualMachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o SqlVirtualMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state to track the async operation status.
func (o SqlVirtualMachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// SQL Server configuration management settings.
func (o SqlVirtualMachineOutput) ServerConfigurationsManagementSettings() ServerConfigurationsManagementSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) ServerConfigurationsManagementSettingsResponsePtrOutput {
		return v.ServerConfigurationsManagementSettings
	}).(ServerConfigurationsManagementSettingsResponsePtrOutput)
}

// SQL image offer. Examples include SQL2016-WS2016, SQL2017-WS2016.
func (o SqlVirtualMachineOutput) SqlImageOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.SqlImageOffer }).(pulumi.StringPtrOutput)
}

// SQL Server edition type.
func (o SqlVirtualMachineOutput) SqlImageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.SqlImageSku }).(pulumi.StringPtrOutput)
}

// SQL Server Management type. NOTE: This parameter is not used anymore. API will automatically detect the Sql Management, refrain from using it.
func (o SqlVirtualMachineOutput) SqlManagement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.SqlManagement }).(pulumi.StringPtrOutput)
}

// SQL Server license type.
func (o SqlVirtualMachineOutput) SqlServerLicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.SqlServerLicenseType }).(pulumi.StringPtrOutput)
}

// ARM resource id of the SQL virtual machine group this SQL virtual machine is or will be part of.
func (o SqlVirtualMachineOutput) SqlVirtualMachineGroupResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.SqlVirtualMachineGroupResourceId }).(pulumi.StringPtrOutput)
}

// Storage Configuration Settings.
func (o SqlVirtualMachineOutput) StorageConfigurationSettings() StorageConfigurationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) StorageConfigurationSettingsResponsePtrOutput {
		return v.StorageConfigurationSettings
	}).(StorageConfigurationSettingsResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o SqlVirtualMachineOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SqlVirtualMachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Troubleshooting status
func (o SqlVirtualMachineOutput) TroubleshootingStatus() TroubleshootingStatusResponseOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) TroubleshootingStatusResponseOutput { return v.TroubleshootingStatus }).(TroubleshootingStatusResponseOutput)
}

// Resource type.
func (o SqlVirtualMachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// ARM Resource id of underlying virtual machine created from SQL marketplace image.
func (o SqlVirtualMachineOutput) VirtualMachineResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.VirtualMachineResourceId }).(pulumi.StringPtrOutput)
}

// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
func (o SqlVirtualMachineOutput) WsfcDomainCredentials() WsfcDomainCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) WsfcDomainCredentialsResponsePtrOutput { return v.WsfcDomainCredentials }).(WsfcDomainCredentialsResponsePtrOutput)
}

// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
func (o SqlVirtualMachineOutput) WsfcStaticIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.WsfcStaticIp }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlVirtualMachineOutput{})
}
