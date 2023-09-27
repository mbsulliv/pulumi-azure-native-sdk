// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a server's blob auditing policy.
// Azure REST API version: 2021-11-01.
func LookupServerBlobAuditingPolicy(ctx *pulumi.Context, args *LookupServerBlobAuditingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupServerBlobAuditingPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServerBlobAuditingPolicyResult
	err := ctx.Invoke("azure-native:sql:getServerBlobAuditingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerBlobAuditingPolicyArgs struct {
	// The name of the blob auditing policy.
	BlobAuditingPolicyName string `pulumi:"blobAuditingPolicyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A server blob auditing policy.
type LookupServerBlobAuditingPolicyResult struct {
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

func LookupServerBlobAuditingPolicyOutput(ctx *pulumi.Context, args LookupServerBlobAuditingPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupServerBlobAuditingPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerBlobAuditingPolicyResult, error) {
			args := v.(LookupServerBlobAuditingPolicyArgs)
			r, err := LookupServerBlobAuditingPolicy(ctx, &args, opts...)
			var s LookupServerBlobAuditingPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerBlobAuditingPolicyResultOutput)
}

type LookupServerBlobAuditingPolicyOutputArgs struct {
	// The name of the blob auditing policy.
	BlobAuditingPolicyName pulumi.StringInput `pulumi:"blobAuditingPolicyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerBlobAuditingPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerBlobAuditingPolicyArgs)(nil)).Elem()
}

// A server blob auditing policy.
type LookupServerBlobAuditingPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupServerBlobAuditingPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerBlobAuditingPolicyResult)(nil)).Elem()
}

func (o LookupServerBlobAuditingPolicyResultOutput) ToLookupServerBlobAuditingPolicyResultOutput() LookupServerBlobAuditingPolicyResultOutput {
	return o
}

func (o LookupServerBlobAuditingPolicyResultOutput) ToLookupServerBlobAuditingPolicyResultOutputWithContext(ctx context.Context) LookupServerBlobAuditingPolicyResultOutput {
	return o
}

func (o LookupServerBlobAuditingPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupServerBlobAuditingPolicyResult] {
	return pulumix.Output[LookupServerBlobAuditingPolicyResult]{
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
func (o LookupServerBlobAuditingPolicyResultOutput) AuditActionsAndGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) []string { return v.AuditActionsAndGroups }).(pulumi.StringArrayOutput)
}

// Resource ID.
func (o LookupServerBlobAuditingPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) string { return v.Id }).(pulumi.StringOutput)
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
func (o LookupServerBlobAuditingPolicyResultOutput) IsAzureMonitorTargetEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *bool { return v.IsAzureMonitorTargetEnabled }).(pulumi.BoolPtrOutput)
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
func (o LookupServerBlobAuditingPolicyResultOutput) IsDevopsAuditEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *bool { return v.IsDevopsAuditEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies whether Managed Identity is used to access blob storage
func (o LookupServerBlobAuditingPolicyResultOutput) IsManagedIdentityInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *bool { return v.IsManagedIdentityInUse }).(pulumi.BoolPtrOutput)
}

// Specifies whether storageAccountAccessKey value is the storage's secondary key.
func (o LookupServerBlobAuditingPolicyResultOutput) IsStorageSecondaryKeyInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *bool { return v.IsStorageSecondaryKeyInUse }).(pulumi.BoolPtrOutput)
}

// Resource name.
func (o LookupServerBlobAuditingPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the amount of time in milliseconds that can elapse before audit actions are forced to be processed.
// The default minimum value is 1000 (1 second). The maximum is 2,147,483,647.
func (o LookupServerBlobAuditingPolicyResultOutput) QueueDelayMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *int { return v.QueueDelayMs }).(pulumi.IntPtrOutput)
}

// Specifies the number of days to keep in the audit logs in the storage account.
func (o LookupServerBlobAuditingPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Specifies the state of the audit. If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled are required.
func (o LookupServerBlobAuditingPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

// Specifies the blob storage subscription Id.
func (o LookupServerBlobAuditingPolicyResultOutput) StorageAccountSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *string { return v.StorageAccountSubscriptionId }).(pulumi.StringPtrOutput)
}

// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled is required.
func (o LookupServerBlobAuditingPolicyResultOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) *string { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupServerBlobAuditingPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerBlobAuditingPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerBlobAuditingPolicyResultOutput{})
}
