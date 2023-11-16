// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an extended server's blob auditing policy.
func LookupExtendedServerBlobAuditingPolicy(ctx *pulumi.Context, args *LookupExtendedServerBlobAuditingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupExtendedServerBlobAuditingPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExtendedServerBlobAuditingPolicyResult
	err := ctx.Invoke("azure-native:sql/v20230201preview:getExtendedServerBlobAuditingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtendedServerBlobAuditingPolicyArgs struct {
	// The name of the blob auditing policy.
	BlobAuditingPolicyName string `pulumi:"blobAuditingPolicyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// An extended server blob auditing policy.
type LookupExtendedServerBlobAuditingPolicyResult struct {
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
	// Resource ID.
	Id string `pulumi:"id"`
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
	// Resource name.
	Name string `pulumi:"name"`
	// Specifies condition of where clause when creating an audit.
	PredicateExpression *string `pulumi:"predicateExpression"`
	// Specifies the amount of time in milliseconds that can elapse before audit actions are forced to be processed.
	// The default minimum value is 1000 (1 second). The maximum is 2,147,483,647.
	QueueDelayMs *int `pulumi:"queueDelayMs"`
	// Specifies the number of days to keep in the audit logs in the storage account.
	RetentionDays *int `pulumi:"retentionDays"`
	// Specifies the state of the audit. If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled are required.
	State string `pulumi:"state"`
	// Specifies the blob storage subscription Id.
	StorageAccountSubscriptionId *string `pulumi:"storageAccountSubscriptionId"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled is required.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupExtendedServerBlobAuditingPolicyOutput(ctx *pulumi.Context, args LookupExtendedServerBlobAuditingPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupExtendedServerBlobAuditingPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExtendedServerBlobAuditingPolicyResult, error) {
			args := v.(LookupExtendedServerBlobAuditingPolicyArgs)
			r, err := LookupExtendedServerBlobAuditingPolicy(ctx, &args, opts...)
			var s LookupExtendedServerBlobAuditingPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExtendedServerBlobAuditingPolicyResultOutput)
}

type LookupExtendedServerBlobAuditingPolicyOutputArgs struct {
	// The name of the blob auditing policy.
	BlobAuditingPolicyName pulumi.StringInput `pulumi:"blobAuditingPolicyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupExtendedServerBlobAuditingPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtendedServerBlobAuditingPolicyArgs)(nil)).Elem()
}

// An extended server blob auditing policy.
type LookupExtendedServerBlobAuditingPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupExtendedServerBlobAuditingPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtendedServerBlobAuditingPolicyResult)(nil)).Elem()
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) ToLookupExtendedServerBlobAuditingPolicyResultOutput() LookupExtendedServerBlobAuditingPolicyResultOutput {
	return o
}

func (o LookupExtendedServerBlobAuditingPolicyResultOutput) ToLookupExtendedServerBlobAuditingPolicyResultOutputWithContext(ctx context.Context) LookupExtendedServerBlobAuditingPolicyResultOutput {
	return o
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
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) AuditActionsAndGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) []string { return v.AuditActionsAndGroups }).(pulumi.StringArrayOutput)
}

// Resource ID.
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) string { return v.Id }).(pulumi.StringOutput)
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
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) IsAzureMonitorTargetEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *bool { return v.IsAzureMonitorTargetEnabled }).(pulumi.BoolPtrOutput)
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
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) IsDevopsAuditEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *bool { return v.IsDevopsAuditEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies whether Managed Identity is used to access blob storage
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) IsManagedIdentityInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *bool { return v.IsManagedIdentityInUse }).(pulumi.BoolPtrOutput)
}

// Specifies whether storageAccountAccessKey value is the storage's secondary key.
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) IsStorageSecondaryKeyInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *bool { return v.IsStorageSecondaryKeyInUse }).(pulumi.BoolPtrOutput)
}

// Resource name.
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies condition of where clause when creating an audit.
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) PredicateExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *string { return v.PredicateExpression }).(pulumi.StringPtrOutput)
}

// Specifies the amount of time in milliseconds that can elapse before audit actions are forced to be processed.
// The default minimum value is 1000 (1 second). The maximum is 2,147,483,647.
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) QueueDelayMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *int { return v.QueueDelayMs }).(pulumi.IntPtrOutput)
}

// Specifies the number of days to keep in the audit logs in the storage account.
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Specifies the state of the audit. If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled are required.
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

// Specifies the blob storage subscription Id.
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) StorageAccountSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *string { return v.StorageAccountSubscriptionId }).(pulumi.StringPtrOutput)
}

// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled is required.
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) *string { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupExtendedServerBlobAuditingPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtendedServerBlobAuditingPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExtendedServerBlobAuditingPolicyResultOutput{})
}
