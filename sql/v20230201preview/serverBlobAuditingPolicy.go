// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A server blob auditing policy.
type ServerBlobAuditingPolicy struct {
	pulumi.CustomResourceState

	// Specifies the Actions-Groups and Actions to audit.
	//
	// The recommended set of action groups to use is the following combination - this will audit all the queries and stored procedures executed against the database, as well as successful and failed logins:
	//
	// BATCH_COMPLETED_GROUP,
	// SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP,
	// FAILED_DATABASE_AUTHENTICATION_GROUP.
	//
	// This above combination is also the set that is configured by default when enabling auditing from the Azure portal.
	//
	// The supported action groups to audit are (note: choose only specific groups that cover your auditing needs. Using unnecessary groups could lead to very large quantities of audit records):
	//
	// APPLICATION_ROLE_CHANGE_PASSWORD_GROUP
	// BACKUP_RESTORE_GROUP
	// DATABASE_LOGOUT_GROUP
	// DATABASE_OBJECT_CHANGE_GROUP
	// DATABASE_OBJECT_OWNERSHIP_CHANGE_GROUP
	// DATABASE_OBJECT_PERMISSION_CHANGE_GROUP
	// DATABASE_OPERATION_GROUP
	// DATABASE_PERMISSION_CHANGE_GROUP
	// DATABASE_PRINCIPAL_CHANGE_GROUP
	// DATABASE_PRINCIPAL_IMPERSONATION_GROUP
	// DATABASE_ROLE_MEMBER_CHANGE_GROUP
	// FAILED_DATABASE_AUTHENTICATION_GROUP
	// SCHEMA_OBJECT_ACCESS_GROUP
	// SCHEMA_OBJECT_CHANGE_GROUP
	// SCHEMA_OBJECT_OWNERSHIP_CHANGE_GROUP
	// SCHEMA_OBJECT_PERMISSION_CHANGE_GROUP
	// SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP
	// USER_CHANGE_PASSWORD_GROUP
	// BATCH_STARTED_GROUP
	// BATCH_COMPLETED_GROUP
	// DBCC_GROUP
	// DATABASE_OWNERSHIP_CHANGE_GROUP
	// DATABASE_CHANGE_GROUP
	// LEDGER_OPERATION_GROUP
	//
	// These are groups that cover all sql statements and stored procedures executed against the database, and should not be used in combination with other groups as this will result in duplicate audit logs.
	//
	// For more information, see [Database-Level Audit Action Groups](https://docs.microsoft.com/en-us/sql/relational-databases/security/auditing/sql-server-audit-action-groups-and-actions#database-level-audit-action-groups).
	//
	// For Database auditing policy, specific Actions can also be specified (note that Actions cannot be specified for Server auditing policy). The supported actions to audit are:
	// SELECT
	// UPDATE
	// INSERT
	// DELETE
	// EXECUTE
	// RECEIVE
	// REFERENCES
	//
	// The general form for defining an action to be audited is:
	// {action} ON {object} BY {principal}
	//
	// Note that <object> in the above format can refer to an object like a table, view, or stored procedure, or an entire database or schema. For the latter cases, the forms DATABASE::{db_name} and SCHEMA::{schema_name} are used, respectively.
	//
	// For example:
	// SELECT on dbo.myTable by public
	// SELECT on DATABASE::myDatabase by public
	// SELECT on SCHEMA::mySchema by public
	//
	// For more information, see [Database-Level Audit Actions](https://docs.microsoft.com/en-us/sql/relational-databases/security/auditing/sql-server-audit-action-groups-and-actions#database-level-audit-actions)
	AuditActionsAndGroups pulumi.StringArrayOutput `pulumi:"auditActionsAndGroups"`
	// Specifies whether audit events are sent to Azure Monitor.
	// In order to send the events to Azure Monitor, specify 'State' as 'Enabled' and 'IsAzureMonitorTargetEnabled' as true.
	//
	// When using REST API to configure auditing, Diagnostic Settings with 'SQLSecurityAuditEvents' diagnostic logs category on the database should be also created.
	// Note that for server level audit you should use the 'master' database as {databaseName}.
	//
	// Diagnostic Settings URI format:
	// PUT https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/providers/microsoft.insights/diagnosticSettings/{settingsName}?api-version=2017-05-01-preview
	//
	// For more information, see [Diagnostic Settings REST API](https://go.microsoft.com/fwlink/?linkid=2033207)
	// or [Diagnostic Settings PowerShell](https://go.microsoft.com/fwlink/?linkid=2033043)
	IsAzureMonitorTargetEnabled pulumi.BoolPtrOutput `pulumi:"isAzureMonitorTargetEnabled"`
	// Specifies the state of devops audit. If state is Enabled, devops logs will be sent to Azure Monitor.
	// In order to send the events to Azure Monitor, specify 'State' as 'Enabled', 'IsAzureMonitorTargetEnabled' as true and 'IsDevopsAuditEnabled' as true
	//
	// When using REST API to configure auditing, Diagnostic Settings with 'DevOpsOperationsAudit' diagnostic logs category on the master database should also be created.
	//
	// Diagnostic Settings URI format:
	// PUT https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Sql/servers/{serverName}/databases/master/providers/microsoft.insights/diagnosticSettings/{settingsName}?api-version=2017-05-01-preview
	//
	// For more information, see [Diagnostic Settings REST API](https://go.microsoft.com/fwlink/?linkid=2033207)
	// or [Diagnostic Settings PowerShell](https://go.microsoft.com/fwlink/?linkid=2033043)
	IsDevopsAuditEnabled pulumi.BoolPtrOutput `pulumi:"isDevopsAuditEnabled"`
	// Specifies whether Managed Identity is used to access blob storage
	IsManagedIdentityInUse pulumi.BoolPtrOutput `pulumi:"isManagedIdentityInUse"`
	// Specifies whether storageAccountAccessKey value is the storage's secondary key.
	IsStorageSecondaryKeyInUse pulumi.BoolPtrOutput `pulumi:"isStorageSecondaryKeyInUse"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the amount of time in milliseconds that can elapse before audit actions are forced to be processed.
	// The default minimum value is 1000 (1 second). The maximum is 2,147,483,647.
	QueueDelayMs pulumi.IntPtrOutput `pulumi:"queueDelayMs"`
	// Specifies the number of days to keep in the audit logs in the storage account.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
	// Specifies the state of the audit. If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled are required.
	State pulumi.StringOutput `pulumi:"state"`
	// Specifies the blob storage subscription Id.
	StorageAccountSubscriptionId pulumi.StringPtrOutput `pulumi:"storageAccountSubscriptionId"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled is required.
	StorageEndpoint pulumi.StringPtrOutput `pulumi:"storageEndpoint"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServerBlobAuditingPolicy registers a new resource with the given unique name, arguments, and options.
func NewServerBlobAuditingPolicy(ctx *pulumi.Context,
	name string, args *ServerBlobAuditingPolicyArgs, opts ...pulumi.ResourceOption) (*ServerBlobAuditingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ServerBlobAuditingPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServerBlobAuditingPolicy
	err := ctx.RegisterResource("azure-native:sql/v20230201preview:ServerBlobAuditingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerBlobAuditingPolicy gets an existing ServerBlobAuditingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerBlobAuditingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerBlobAuditingPolicyState, opts ...pulumi.ResourceOption) (*ServerBlobAuditingPolicy, error) {
	var resource ServerBlobAuditingPolicy
	err := ctx.ReadResource("azure-native:sql/v20230201preview:ServerBlobAuditingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerBlobAuditingPolicy resources.
type serverBlobAuditingPolicyState struct {
}

type ServerBlobAuditingPolicyState struct {
}

func (ServerBlobAuditingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverBlobAuditingPolicyState)(nil)).Elem()
}

type serverBlobAuditingPolicyArgs struct {
	// Specifies the Actions-Groups and Actions to audit.
	//
	// The recommended set of action groups to use is the following combination - this will audit all the queries and stored procedures executed against the database, as well as successful and failed logins:
	//
	// BATCH_COMPLETED_GROUP,
	// SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP,
	// FAILED_DATABASE_AUTHENTICATION_GROUP.
	//
	// This above combination is also the set that is configured by default when enabling auditing from the Azure portal.
	//
	// The supported action groups to audit are (note: choose only specific groups that cover your auditing needs. Using unnecessary groups could lead to very large quantities of audit records):
	//
	// APPLICATION_ROLE_CHANGE_PASSWORD_GROUP
	// BACKUP_RESTORE_GROUP
	// DATABASE_LOGOUT_GROUP
	// DATABASE_OBJECT_CHANGE_GROUP
	// DATABASE_OBJECT_OWNERSHIP_CHANGE_GROUP
	// DATABASE_OBJECT_PERMISSION_CHANGE_GROUP
	// DATABASE_OPERATION_GROUP
	// DATABASE_PERMISSION_CHANGE_GROUP
	// DATABASE_PRINCIPAL_CHANGE_GROUP
	// DATABASE_PRINCIPAL_IMPERSONATION_GROUP
	// DATABASE_ROLE_MEMBER_CHANGE_GROUP
	// FAILED_DATABASE_AUTHENTICATION_GROUP
	// SCHEMA_OBJECT_ACCESS_GROUP
	// SCHEMA_OBJECT_CHANGE_GROUP
	// SCHEMA_OBJECT_OWNERSHIP_CHANGE_GROUP
	// SCHEMA_OBJECT_PERMISSION_CHANGE_GROUP
	// SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP
	// USER_CHANGE_PASSWORD_GROUP
	// BATCH_STARTED_GROUP
	// BATCH_COMPLETED_GROUP
	// DBCC_GROUP
	// DATABASE_OWNERSHIP_CHANGE_GROUP
	// DATABASE_CHANGE_GROUP
	// LEDGER_OPERATION_GROUP
	//
	// These are groups that cover all sql statements and stored procedures executed against the database, and should not be used in combination with other groups as this will result in duplicate audit logs.
	//
	// For more information, see [Database-Level Audit Action Groups](https://docs.microsoft.com/en-us/sql/relational-databases/security/auditing/sql-server-audit-action-groups-and-actions#database-level-audit-action-groups).
	//
	// For Database auditing policy, specific Actions can also be specified (note that Actions cannot be specified for Server auditing policy). The supported actions to audit are:
	// SELECT
	// UPDATE
	// INSERT
	// DELETE
	// EXECUTE
	// RECEIVE
	// REFERENCES
	//
	// The general form for defining an action to be audited is:
	// {action} ON {object} BY {principal}
	//
	// Note that <object> in the above format can refer to an object like a table, view, or stored procedure, or an entire database or schema. For the latter cases, the forms DATABASE::{db_name} and SCHEMA::{schema_name} are used, respectively.
	//
	// For example:
	// SELECT on dbo.myTable by public
	// SELECT on DATABASE::myDatabase by public
	// SELECT on SCHEMA::mySchema by public
	//
	// For more information, see [Database-Level Audit Actions](https://docs.microsoft.com/en-us/sql/relational-databases/security/auditing/sql-server-audit-action-groups-and-actions#database-level-audit-actions)
	AuditActionsAndGroups []string `pulumi:"auditActionsAndGroups"`
	// The name of the blob auditing policy.
	BlobAuditingPolicyName *string `pulumi:"blobAuditingPolicyName"`
	// Specifies whether audit events are sent to Azure Monitor.
	// In order to send the events to Azure Monitor, specify 'State' as 'Enabled' and 'IsAzureMonitorTargetEnabled' as true.
	//
	// When using REST API to configure auditing, Diagnostic Settings with 'SQLSecurityAuditEvents' diagnostic logs category on the database should be also created.
	// Note that for server level audit you should use the 'master' database as {databaseName}.
	//
	// Diagnostic Settings URI format:
	// PUT https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/providers/microsoft.insights/diagnosticSettings/{settingsName}?api-version=2017-05-01-preview
	//
	// For more information, see [Diagnostic Settings REST API](https://go.microsoft.com/fwlink/?linkid=2033207)
	// or [Diagnostic Settings PowerShell](https://go.microsoft.com/fwlink/?linkid=2033043)
	IsAzureMonitorTargetEnabled *bool `pulumi:"isAzureMonitorTargetEnabled"`
	// Specifies the state of devops audit. If state is Enabled, devops logs will be sent to Azure Monitor.
	// In order to send the events to Azure Monitor, specify 'State' as 'Enabled', 'IsAzureMonitorTargetEnabled' as true and 'IsDevopsAuditEnabled' as true
	//
	// When using REST API to configure auditing, Diagnostic Settings with 'DevOpsOperationsAudit' diagnostic logs category on the master database should also be created.
	//
	// Diagnostic Settings URI format:
	// PUT https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Sql/servers/{serverName}/databases/master/providers/microsoft.insights/diagnosticSettings/{settingsName}?api-version=2017-05-01-preview
	//
	// For more information, see [Diagnostic Settings REST API](https://go.microsoft.com/fwlink/?linkid=2033207)
	// or [Diagnostic Settings PowerShell](https://go.microsoft.com/fwlink/?linkid=2033043)
	IsDevopsAuditEnabled *bool `pulumi:"isDevopsAuditEnabled"`
	// Specifies whether Managed Identity is used to access blob storage
	IsManagedIdentityInUse *bool `pulumi:"isManagedIdentityInUse"`
	// Specifies whether storageAccountAccessKey value is the storage's secondary key.
	IsStorageSecondaryKeyInUse *bool `pulumi:"isStorageSecondaryKeyInUse"`
	// Specifies the amount of time in milliseconds that can elapse before audit actions are forced to be processed.
	// The default minimum value is 1000 (1 second). The maximum is 2,147,483,647.
	QueueDelayMs *int `pulumi:"queueDelayMs"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the number of days to keep in the audit logs in the storage account.
	RetentionDays *int `pulumi:"retentionDays"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// Specifies the state of the audit. If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled are required.
	State BlobAuditingPolicyState `pulumi:"state"`
	// Specifies the identifier key of the auditing storage account.
	// If state is Enabled and storageEndpoint is specified, not specifying the storageAccountAccessKey will use SQL server system-assigned managed identity to access the storage.
	// Prerequisites for using managed identity authentication:
	// 1. Assign SQL Server a system-assigned managed identity in Azure Active Directory (AAD).
	// 2. Grant SQL Server identity access to the storage account by adding 'Storage Blob Data Contributor' RBAC role to the server identity.
	//    For more information, see [Auditing to storage using Managed Identity authentication](https://go.microsoft.com/fwlink/?linkid=2114355)
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage subscription Id.
	StorageAccountSubscriptionId *string `pulumi:"storageAccountSubscriptionId"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled is required.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
}

// The set of arguments for constructing a ServerBlobAuditingPolicy resource.
type ServerBlobAuditingPolicyArgs struct {
	// Specifies the Actions-Groups and Actions to audit.
	//
	// The recommended set of action groups to use is the following combination - this will audit all the queries and stored procedures executed against the database, as well as successful and failed logins:
	//
	// BATCH_COMPLETED_GROUP,
	// SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP,
	// FAILED_DATABASE_AUTHENTICATION_GROUP.
	//
	// This above combination is also the set that is configured by default when enabling auditing from the Azure portal.
	//
	// The supported action groups to audit are (note: choose only specific groups that cover your auditing needs. Using unnecessary groups could lead to very large quantities of audit records):
	//
	// APPLICATION_ROLE_CHANGE_PASSWORD_GROUP
	// BACKUP_RESTORE_GROUP
	// DATABASE_LOGOUT_GROUP
	// DATABASE_OBJECT_CHANGE_GROUP
	// DATABASE_OBJECT_OWNERSHIP_CHANGE_GROUP
	// DATABASE_OBJECT_PERMISSION_CHANGE_GROUP
	// DATABASE_OPERATION_GROUP
	// DATABASE_PERMISSION_CHANGE_GROUP
	// DATABASE_PRINCIPAL_CHANGE_GROUP
	// DATABASE_PRINCIPAL_IMPERSONATION_GROUP
	// DATABASE_ROLE_MEMBER_CHANGE_GROUP
	// FAILED_DATABASE_AUTHENTICATION_GROUP
	// SCHEMA_OBJECT_ACCESS_GROUP
	// SCHEMA_OBJECT_CHANGE_GROUP
	// SCHEMA_OBJECT_OWNERSHIP_CHANGE_GROUP
	// SCHEMA_OBJECT_PERMISSION_CHANGE_GROUP
	// SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP
	// USER_CHANGE_PASSWORD_GROUP
	// BATCH_STARTED_GROUP
	// BATCH_COMPLETED_GROUP
	// DBCC_GROUP
	// DATABASE_OWNERSHIP_CHANGE_GROUP
	// DATABASE_CHANGE_GROUP
	// LEDGER_OPERATION_GROUP
	//
	// These are groups that cover all sql statements and stored procedures executed against the database, and should not be used in combination with other groups as this will result in duplicate audit logs.
	//
	// For more information, see [Database-Level Audit Action Groups](https://docs.microsoft.com/en-us/sql/relational-databases/security/auditing/sql-server-audit-action-groups-and-actions#database-level-audit-action-groups).
	//
	// For Database auditing policy, specific Actions can also be specified (note that Actions cannot be specified for Server auditing policy). The supported actions to audit are:
	// SELECT
	// UPDATE
	// INSERT
	// DELETE
	// EXECUTE
	// RECEIVE
	// REFERENCES
	//
	// The general form for defining an action to be audited is:
	// {action} ON {object} BY {principal}
	//
	// Note that <object> in the above format can refer to an object like a table, view, or stored procedure, or an entire database or schema. For the latter cases, the forms DATABASE::{db_name} and SCHEMA::{schema_name} are used, respectively.
	//
	// For example:
	// SELECT on dbo.myTable by public
	// SELECT on DATABASE::myDatabase by public
	// SELECT on SCHEMA::mySchema by public
	//
	// For more information, see [Database-Level Audit Actions](https://docs.microsoft.com/en-us/sql/relational-databases/security/auditing/sql-server-audit-action-groups-and-actions#database-level-audit-actions)
	AuditActionsAndGroups pulumi.StringArrayInput
	// The name of the blob auditing policy.
	BlobAuditingPolicyName pulumi.StringPtrInput
	// Specifies whether audit events are sent to Azure Monitor.
	// In order to send the events to Azure Monitor, specify 'State' as 'Enabled' and 'IsAzureMonitorTargetEnabled' as true.
	//
	// When using REST API to configure auditing, Diagnostic Settings with 'SQLSecurityAuditEvents' diagnostic logs category on the database should be also created.
	// Note that for server level audit you should use the 'master' database as {databaseName}.
	//
	// Diagnostic Settings URI format:
	// PUT https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/providers/microsoft.insights/diagnosticSettings/{settingsName}?api-version=2017-05-01-preview
	//
	// For more information, see [Diagnostic Settings REST API](https://go.microsoft.com/fwlink/?linkid=2033207)
	// or [Diagnostic Settings PowerShell](https://go.microsoft.com/fwlink/?linkid=2033043)
	IsAzureMonitorTargetEnabled pulumi.BoolPtrInput
	// Specifies the state of devops audit. If state is Enabled, devops logs will be sent to Azure Monitor.
	// In order to send the events to Azure Monitor, specify 'State' as 'Enabled', 'IsAzureMonitorTargetEnabled' as true and 'IsDevopsAuditEnabled' as true
	//
	// When using REST API to configure auditing, Diagnostic Settings with 'DevOpsOperationsAudit' diagnostic logs category on the master database should also be created.
	//
	// Diagnostic Settings URI format:
	// PUT https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Sql/servers/{serverName}/databases/master/providers/microsoft.insights/diagnosticSettings/{settingsName}?api-version=2017-05-01-preview
	//
	// For more information, see [Diagnostic Settings REST API](https://go.microsoft.com/fwlink/?linkid=2033207)
	// or [Diagnostic Settings PowerShell](https://go.microsoft.com/fwlink/?linkid=2033043)
	IsDevopsAuditEnabled pulumi.BoolPtrInput
	// Specifies whether Managed Identity is used to access blob storage
	IsManagedIdentityInUse pulumi.BoolPtrInput
	// Specifies whether storageAccountAccessKey value is the storage's secondary key.
	IsStorageSecondaryKeyInUse pulumi.BoolPtrInput
	// Specifies the amount of time in milliseconds that can elapse before audit actions are forced to be processed.
	// The default minimum value is 1000 (1 second). The maximum is 2,147,483,647.
	QueueDelayMs pulumi.IntPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Specifies the number of days to keep in the audit logs in the storage account.
	RetentionDays pulumi.IntPtrInput
	// The name of the server.
	ServerName pulumi.StringInput
	// Specifies the state of the audit. If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled are required.
	State BlobAuditingPolicyStateInput
	// Specifies the identifier key of the auditing storage account.
	// If state is Enabled and storageEndpoint is specified, not specifying the storageAccountAccessKey will use SQL server system-assigned managed identity to access the storage.
	// Prerequisites for using managed identity authentication:
	// 1. Assign SQL Server a system-assigned managed identity in Azure Active Directory (AAD).
	// 2. Grant SQL Server identity access to the storage account by adding 'Storage Blob Data Contributor' RBAC role to the server identity.
	//    For more information, see [Auditing to storage using Managed Identity authentication](https://go.microsoft.com/fwlink/?linkid=2114355)
	StorageAccountAccessKey pulumi.StringPtrInput
	// Specifies the blob storage subscription Id.
	StorageAccountSubscriptionId pulumi.StringPtrInput
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled is required.
	StorageEndpoint pulumi.StringPtrInput
}

func (ServerBlobAuditingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverBlobAuditingPolicyArgs)(nil)).Elem()
}

type ServerBlobAuditingPolicyInput interface {
	pulumi.Input

	ToServerBlobAuditingPolicyOutput() ServerBlobAuditingPolicyOutput
	ToServerBlobAuditingPolicyOutputWithContext(ctx context.Context) ServerBlobAuditingPolicyOutput
}

func (*ServerBlobAuditingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerBlobAuditingPolicy)(nil)).Elem()
}

func (i *ServerBlobAuditingPolicy) ToServerBlobAuditingPolicyOutput() ServerBlobAuditingPolicyOutput {
	return i.ToServerBlobAuditingPolicyOutputWithContext(context.Background())
}

func (i *ServerBlobAuditingPolicy) ToServerBlobAuditingPolicyOutputWithContext(ctx context.Context) ServerBlobAuditingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerBlobAuditingPolicyOutput)
}

func (i *ServerBlobAuditingPolicy) ToOutput(ctx context.Context) pulumix.Output[*ServerBlobAuditingPolicy] {
	return pulumix.Output[*ServerBlobAuditingPolicy]{
		OutputState: i.ToServerBlobAuditingPolicyOutputWithContext(ctx).OutputState,
	}
}

type ServerBlobAuditingPolicyOutput struct{ *pulumi.OutputState }

func (ServerBlobAuditingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerBlobAuditingPolicy)(nil)).Elem()
}

func (o ServerBlobAuditingPolicyOutput) ToServerBlobAuditingPolicyOutput() ServerBlobAuditingPolicyOutput {
	return o
}

func (o ServerBlobAuditingPolicyOutput) ToServerBlobAuditingPolicyOutputWithContext(ctx context.Context) ServerBlobAuditingPolicyOutput {
	return o
}

func (o ServerBlobAuditingPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*ServerBlobAuditingPolicy] {
	return pulumix.Output[*ServerBlobAuditingPolicy]{
		OutputState: o.OutputState,
	}
}

// Specifies the Actions-Groups and Actions to audit.
//
// The recommended set of action groups to use is the following combination - this will audit all the queries and stored procedures executed against the database, as well as successful and failed logins:
//
// BATCH_COMPLETED_GROUP,
// SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP,
// FAILED_DATABASE_AUTHENTICATION_GROUP.
//
// This above combination is also the set that is configured by default when enabling auditing from the Azure portal.
//
// The supported action groups to audit are (note: choose only specific groups that cover your auditing needs. Using unnecessary groups could lead to very large quantities of audit records):
//
// APPLICATION_ROLE_CHANGE_PASSWORD_GROUP
// BACKUP_RESTORE_GROUP
// DATABASE_LOGOUT_GROUP
// DATABASE_OBJECT_CHANGE_GROUP
// DATABASE_OBJECT_OWNERSHIP_CHANGE_GROUP
// DATABASE_OBJECT_PERMISSION_CHANGE_GROUP
// DATABASE_OPERATION_GROUP
// DATABASE_PERMISSION_CHANGE_GROUP
// DATABASE_PRINCIPAL_CHANGE_GROUP
// DATABASE_PRINCIPAL_IMPERSONATION_GROUP
// DATABASE_ROLE_MEMBER_CHANGE_GROUP
// FAILED_DATABASE_AUTHENTICATION_GROUP
// SCHEMA_OBJECT_ACCESS_GROUP
// SCHEMA_OBJECT_CHANGE_GROUP
// SCHEMA_OBJECT_OWNERSHIP_CHANGE_GROUP
// SCHEMA_OBJECT_PERMISSION_CHANGE_GROUP
// SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP
// USER_CHANGE_PASSWORD_GROUP
// BATCH_STARTED_GROUP
// BATCH_COMPLETED_GROUP
// DBCC_GROUP
// DATABASE_OWNERSHIP_CHANGE_GROUP
// DATABASE_CHANGE_GROUP
// LEDGER_OPERATION_GROUP
//
// These are groups that cover all sql statements and stored procedures executed against the database, and should not be used in combination with other groups as this will result in duplicate audit logs.
//
// For more information, see [Database-Level Audit Action Groups](https://docs.microsoft.com/en-us/sql/relational-databases/security/auditing/sql-server-audit-action-groups-and-actions#database-level-audit-action-groups).
//
// For Database auditing policy, specific Actions can also be specified (note that Actions cannot be specified for Server auditing policy). The supported actions to audit are:
// SELECT
// UPDATE
// INSERT
// DELETE
// EXECUTE
// RECEIVE
// REFERENCES
//
// The general form for defining an action to be audited is:
// {action} ON {object} BY {principal}
//
// Note that <object> in the above format can refer to an object like a table, view, or stored procedure, or an entire database or schema. For the latter cases, the forms DATABASE::{db_name} and SCHEMA::{schema_name} are used, respectively.
//
// For example:
// SELECT on dbo.myTable by public
// SELECT on DATABASE::myDatabase by public
// SELECT on SCHEMA::mySchema by public
//
// For more information, see [Database-Level Audit Actions](https://docs.microsoft.com/en-us/sql/relational-databases/security/auditing/sql-server-audit-action-groups-and-actions#database-level-audit-actions)
func (o ServerBlobAuditingPolicyOutput) AuditActionsAndGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.StringArrayOutput { return v.AuditActionsAndGroups }).(pulumi.StringArrayOutput)
}

// Specifies whether audit events are sent to Azure Monitor.
// In order to send the events to Azure Monitor, specify 'State' as 'Enabled' and 'IsAzureMonitorTargetEnabled' as true.
//
// When using REST API to configure auditing, Diagnostic Settings with 'SQLSecurityAuditEvents' diagnostic logs category on the database should be also created.
// Note that for server level audit you should use the 'master' database as {databaseName}.
//
// Diagnostic Settings URI format:
// PUT https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/providers/microsoft.insights/diagnosticSettings/{settingsName}?api-version=2017-05-01-preview
//
// For more information, see [Diagnostic Settings REST API](https://go.microsoft.com/fwlink/?linkid=2033207)
// or [Diagnostic Settings PowerShell](https://go.microsoft.com/fwlink/?linkid=2033043)
func (o ServerBlobAuditingPolicyOutput) IsAzureMonitorTargetEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsAzureMonitorTargetEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies the state of devops audit. If state is Enabled, devops logs will be sent to Azure Monitor.
// In order to send the events to Azure Monitor, specify 'State' as 'Enabled', 'IsAzureMonitorTargetEnabled' as true and 'IsDevopsAuditEnabled' as true
//
// When using REST API to configure auditing, Diagnostic Settings with 'DevOpsOperationsAudit' diagnostic logs category on the master database should also be created.
//
// Diagnostic Settings URI format:
// PUT https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Sql/servers/{serverName}/databases/master/providers/microsoft.insights/diagnosticSettings/{settingsName}?api-version=2017-05-01-preview
//
// For more information, see [Diagnostic Settings REST API](https://go.microsoft.com/fwlink/?linkid=2033207)
// or [Diagnostic Settings PowerShell](https://go.microsoft.com/fwlink/?linkid=2033043)
func (o ServerBlobAuditingPolicyOutput) IsDevopsAuditEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsDevopsAuditEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies whether Managed Identity is used to access blob storage
func (o ServerBlobAuditingPolicyOutput) IsManagedIdentityInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsManagedIdentityInUse }).(pulumi.BoolPtrOutput)
}

// Specifies whether storageAccountAccessKey value is the storage's secondary key.
func (o ServerBlobAuditingPolicyOutput) IsStorageSecondaryKeyInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsStorageSecondaryKeyInUse }).(pulumi.BoolPtrOutput)
}

// Resource name.
func (o ServerBlobAuditingPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the amount of time in milliseconds that can elapse before audit actions are forced to be processed.
// The default minimum value is 1000 (1 second). The maximum is 2,147,483,647.
func (o ServerBlobAuditingPolicyOutput) QueueDelayMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.IntPtrOutput { return v.QueueDelayMs }).(pulumi.IntPtrOutput)
}

// Specifies the number of days to keep in the audit logs in the storage account.
func (o ServerBlobAuditingPolicyOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Specifies the state of the audit. If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled are required.
func (o ServerBlobAuditingPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Specifies the blob storage subscription Id.
func (o ServerBlobAuditingPolicyOutput) StorageAccountSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.StringPtrOutput { return v.StorageAccountSubscriptionId }).(pulumi.StringPtrOutput)
}

// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled is required.
func (o ServerBlobAuditingPolicyOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.StringPtrOutput { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o ServerBlobAuditingPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerBlobAuditingPolicyOutput{})
}
