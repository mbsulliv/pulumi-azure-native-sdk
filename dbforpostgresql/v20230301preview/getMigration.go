// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets details of a migration.
func LookupMigration(ctx *pulumi.Context, args *LookupMigrationArgs, opts ...pulumi.InvokeOption) (*LookupMigrationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMigrationResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20230301preview:getMigration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMigrationArgs struct {
	// The name of the migration.
	MigrationName string `pulumi:"migrationName"`
	// The resource group name of the target database server.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The subscription ID of the target database server.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The name of the target database server.
	TargetDbServerName string `pulumi:"targetDbServerName"`
}

// Represents a migration resource.
type LookupMigrationResult struct {
	// To trigger cancel for entire migration we need to send this flag as True
	Cancel *string `pulumi:"cancel"`
	// Current status of migration
	CurrentStatus MigrationStatusResponse `pulumi:"currentStatus"`
	// When you want to trigger cancel for specific databases send cancel flag as True and database names in this array
	DbsToCancelMigrationOn []string `pulumi:"dbsToCancelMigrationOn"`
	// Number of databases to migrate
	DbsToMigrate []string `pulumi:"dbsToMigrate"`
	// When you want to trigger cutover for specific databases send triggerCutover flag as True and database names in this array
	DbsToTriggerCutoverOn []string `pulumi:"dbsToTriggerCutoverOn"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// ID for migration, a GUID.
	MigrationId string `pulumi:"migrationId"`
	// There are two types of migration modes Online and Offline
	MigrationMode *string `pulumi:"migrationMode"`
	// End time in UTC for migration window
	MigrationWindowEndTimeInUtc *string `pulumi:"migrationWindowEndTimeInUtc"`
	// Start time in UTC for migration window
	MigrationWindowStartTimeInUtc *string `pulumi:"migrationWindowStartTimeInUtc"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Indicates whether the databases on the target server can be overwritten, if already present. If set to False, the migration workflow will wait for a confirmation, if it detects that the database already exists.
	OverwriteDbsInTarget *string `pulumi:"overwriteDbsInTarget"`
	// Indicates whether to setup LogicalReplicationOnSourceDb, if needed
	SetupLogicalReplicationOnSourceDbIfNeeded *string `pulumi:"setupLogicalReplicationOnSourceDbIfNeeded"`
	// Source server fully qualified domain name or ip. It is a optional value, if customer provide it, dms will always use it for connection
	SourceDbServerFullyQualifiedDomainName *string `pulumi:"sourceDbServerFullyQualifiedDomainName"`
	// Metadata of the source database server
	SourceDbServerMetadata DbServerMetadataResponse `pulumi:"sourceDbServerMetadata"`
	// ResourceId of the source database server
	SourceDbServerResourceId *string `pulumi:"sourceDbServerResourceId"`
	// Indicates whether the data migration should start right away
	StartDataMigration *string `pulumi:"startDataMigration"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Target server fully qualified domain name or ip. It is a optional value, if customer provide it, dms will always use it for connection
	TargetDbServerFullyQualifiedDomainName *string `pulumi:"targetDbServerFullyQualifiedDomainName"`
	// Metadata of the target database server
	TargetDbServerMetadata DbServerMetadataResponse `pulumi:"targetDbServerMetadata"`
	// ResourceId of the source database server
	TargetDbServerResourceId string `pulumi:"targetDbServerResourceId"`
	// To trigger cutover for entire migration we need to send this flag as True
	TriggerCutover *string `pulumi:"triggerCutover"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupMigrationOutput(ctx *pulumi.Context, args LookupMigrationOutputArgs, opts ...pulumi.InvokeOption) LookupMigrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMigrationResult, error) {
			args := v.(LookupMigrationArgs)
			r, err := LookupMigration(ctx, &args, opts...)
			var s LookupMigrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMigrationResultOutput)
}

type LookupMigrationOutputArgs struct {
	// The name of the migration.
	MigrationName pulumi.StringInput `pulumi:"migrationName"`
	// The resource group name of the target database server.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The subscription ID of the target database server.
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
	// The name of the target database server.
	TargetDbServerName pulumi.StringInput `pulumi:"targetDbServerName"`
}

func (LookupMigrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrationArgs)(nil)).Elem()
}

// Represents a migration resource.
type LookupMigrationResultOutput struct{ *pulumi.OutputState }

func (LookupMigrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrationResult)(nil)).Elem()
}

func (o LookupMigrationResultOutput) ToLookupMigrationResultOutput() LookupMigrationResultOutput {
	return o
}

func (o LookupMigrationResultOutput) ToLookupMigrationResultOutputWithContext(ctx context.Context) LookupMigrationResultOutput {
	return o
}

func (o LookupMigrationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMigrationResult] {
	return pulumix.Output[LookupMigrationResult]{
		OutputState: o.OutputState,
	}
}

// To trigger cancel for entire migration we need to send this flag as True
func (o LookupMigrationResultOutput) Cancel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.Cancel }).(pulumi.StringPtrOutput)
}

// Current status of migration
func (o LookupMigrationResultOutput) CurrentStatus() MigrationStatusResponseOutput {
	return o.ApplyT(func(v LookupMigrationResult) MigrationStatusResponse { return v.CurrentStatus }).(MigrationStatusResponseOutput)
}

// When you want to trigger cancel for specific databases send cancel flag as True and database names in this array
func (o LookupMigrationResultOutput) DbsToCancelMigrationOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMigrationResult) []string { return v.DbsToCancelMigrationOn }).(pulumi.StringArrayOutput)
}

// Number of databases to migrate
func (o LookupMigrationResultOutput) DbsToMigrate() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMigrationResult) []string { return v.DbsToMigrate }).(pulumi.StringArrayOutput)
}

// When you want to trigger cutover for specific databases send triggerCutover flag as True and database names in this array
func (o LookupMigrationResultOutput) DbsToTriggerCutoverOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMigrationResult) []string { return v.DbsToTriggerCutoverOn }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupMigrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupMigrationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.Location }).(pulumi.StringOutput)
}

// ID for migration, a GUID.
func (o LookupMigrationResultOutput) MigrationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.MigrationId }).(pulumi.StringOutput)
}

// There are two types of migration modes Online and Offline
func (o LookupMigrationResultOutput) MigrationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.MigrationMode }).(pulumi.StringPtrOutput)
}

// End time in UTC for migration window
func (o LookupMigrationResultOutput) MigrationWindowEndTimeInUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.MigrationWindowEndTimeInUtc }).(pulumi.StringPtrOutput)
}

// Start time in UTC for migration window
func (o LookupMigrationResultOutput) MigrationWindowStartTimeInUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.MigrationWindowStartTimeInUtc }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupMigrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether the databases on the target server can be overwritten, if already present. If set to False, the migration workflow will wait for a confirmation, if it detects that the database already exists.
func (o LookupMigrationResultOutput) OverwriteDbsInTarget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.OverwriteDbsInTarget }).(pulumi.StringPtrOutput)
}

// Indicates whether to setup LogicalReplicationOnSourceDb, if needed
func (o LookupMigrationResultOutput) SetupLogicalReplicationOnSourceDbIfNeeded() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.SetupLogicalReplicationOnSourceDbIfNeeded }).(pulumi.StringPtrOutput)
}

// Source server fully qualified domain name or ip. It is a optional value, if customer provide it, dms will always use it for connection
func (o LookupMigrationResultOutput) SourceDbServerFullyQualifiedDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.SourceDbServerFullyQualifiedDomainName }).(pulumi.StringPtrOutput)
}

// Metadata of the source database server
func (o LookupMigrationResultOutput) SourceDbServerMetadata() DbServerMetadataResponseOutput {
	return o.ApplyT(func(v LookupMigrationResult) DbServerMetadataResponse { return v.SourceDbServerMetadata }).(DbServerMetadataResponseOutput)
}

// ResourceId of the source database server
func (o LookupMigrationResultOutput) SourceDbServerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.SourceDbServerResourceId }).(pulumi.StringPtrOutput)
}

// Indicates whether the data migration should start right away
func (o LookupMigrationResultOutput) StartDataMigration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.StartDataMigration }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMigrationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMigrationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupMigrationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMigrationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Target server fully qualified domain name or ip. It is a optional value, if customer provide it, dms will always use it for connection
func (o LookupMigrationResultOutput) TargetDbServerFullyQualifiedDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.TargetDbServerFullyQualifiedDomainName }).(pulumi.StringPtrOutput)
}

// Metadata of the target database server
func (o LookupMigrationResultOutput) TargetDbServerMetadata() DbServerMetadataResponseOutput {
	return o.ApplyT(func(v LookupMigrationResult) DbServerMetadataResponse { return v.TargetDbServerMetadata }).(DbServerMetadataResponseOutput)
}

// ResourceId of the source database server
func (o LookupMigrationResultOutput) TargetDbServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.TargetDbServerResourceId }).(pulumi.StringOutput)
}

// To trigger cutover for entire migration we need to send this flag as True
func (o LookupMigrationResultOutput) TriggerCutover() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.TriggerCutover }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMigrationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMigrationResultOutput{})
}
