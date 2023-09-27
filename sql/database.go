// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A database resource.
// Azure REST API version: 2021-11-01. Prior API version in Azure Native 1.x: 2020-11-01-preview
type Database struct {
	pulumi.CustomResourceState

	// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
	AutoPauseDelay pulumi.IntPtrOutput `pulumi:"autoPauseDelay"`
	// Collation of the metadata catalog.
	CatalogCollation pulumi.StringPtrOutput `pulumi:"catalogCollation"`
	// The collation of the database.
	Collation pulumi.StringPtrOutput `pulumi:"collation"`
	// The creation date of the database (ISO8601 format).
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The storage account type used to store backups for this database.
	CurrentBackupStorageRedundancy pulumi.StringOutput `pulumi:"currentBackupStorageRedundancy"`
	// The current service level objective name of the database.
	CurrentServiceObjectiveName pulumi.StringOutput `pulumi:"currentServiceObjectiveName"`
	// The name and tier of the SKU.
	CurrentSku SkuResponseOutput `pulumi:"currentSku"`
	// The ID of the database.
	DatabaseId pulumi.StringOutput `pulumi:"databaseId"`
	// The default secondary region for this database.
	DefaultSecondaryLocation pulumi.StringOutput `pulumi:"defaultSecondaryLocation"`
	// This records the earliest start date and time that restore is available for this database (ISO8601 format).
	EarliestRestoreDate pulumi.StringOutput `pulumi:"earliestRestoreDate"`
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId pulumi.StringPtrOutput `pulumi:"elasticPoolId"`
	// Failover Group resource identifier that this database belongs to.
	FailoverGroupId pulumi.StringOutput `pulumi:"failoverGroupId"`
	// The Client id used for cross tenant per database CMK scenario
	FederatedClientId pulumi.StringPtrOutput `pulumi:"federatedClientId"`
	// The number of secondary replicas associated with the database that are used to provide high availability. Not applicable to a Hyperscale database within an elastic pool.
	HighAvailabilityReplicaCount pulumi.IntPtrOutput `pulumi:"highAvailabilityReplicaCount"`
	// The Azure Active Directory identity of the database.
	Identity DatabaseIdentityResponsePtrOutput `pulumi:"identity"`
	// Infra encryption is enabled for this database.
	IsInfraEncryptionEnabled pulumi.BoolOutput `pulumi:"isInfraEncryptionEnabled"`
	// Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
	IsLedgerOn pulumi.BoolPtrOutput `pulumi:"isLedgerOn"`
	// Kind of database. This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Maintenance configuration id assigned to the database. This configuration defines the period when the maintenance updates will occur.
	MaintenanceConfigurationId pulumi.StringPtrOutput `pulumi:"maintenanceConfigurationId"`
	// Resource that manages the database.
	ManagedBy pulumi.StringOutput `pulumi:"managedBy"`
	// The max log size for this database.
	MaxLogSizeBytes pulumi.Float64Output `pulumi:"maxLogSizeBytes"`
	// The max size of the database expressed in bytes.
	MaxSizeBytes pulumi.Float64PtrOutput `pulumi:"maxSizeBytes"`
	// Minimal capacity that database will always have allocated, if not paused
	MinCapacity pulumi.Float64PtrOutput `pulumi:"minCapacity"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The date when database was paused by user configuration or action(ISO8601 format). Null if the database is ready.
	PausedDate pulumi.StringOutput `pulumi:"pausedDate"`
	// The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region. Not applicable to a Hyperscale database within an elastic pool.
	ReadScale pulumi.StringPtrOutput `pulumi:"readScale"`
	// The storage account type to be used to store backups for this database.
	RequestedBackupStorageRedundancy pulumi.StringPtrOutput `pulumi:"requestedBackupStorageRedundancy"`
	// The requested service level objective name of the database.
	RequestedServiceObjectiveName pulumi.StringOutput `pulumi:"requestedServiceObjectiveName"`
	// The date when database was resumed by user action or database login (ISO8601 format). Null if the database is paused.
	ResumedDate pulumi.StringOutput `pulumi:"resumedDate"`
	// The secondary type of the database if it is a secondary.  Valid values are Geo and Named.
	SecondaryType pulumi.StringPtrOutput `pulumi:"secondaryType"`
	// The database SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The status of the database.
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql/v20140401:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20171001preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:Database"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Database
	err := ctx.RegisterResource("azure-native:sql:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azure-native:sql:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
}

type DatabaseState struct {
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
	AutoPauseDelay *int `pulumi:"autoPauseDelay"`
	// Collation of the metadata catalog.
	CatalogCollation *string `pulumi:"catalogCollation"`
	// The collation of the database.
	Collation *string `pulumi:"collation"`
	// Specifies the mode of database creation.
	//
	// Default: regular database creation.
	//
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
	//
	// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
	//
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
	//
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
	//
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//
	// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode *string `pulumi:"createMode"`
	// The name of the database.
	DatabaseName *string `pulumi:"databaseName"`
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId *string `pulumi:"elasticPoolId"`
	// The Client id used for cross tenant per database CMK scenario
	FederatedClientId *string `pulumi:"federatedClientId"`
	// The number of secondary replicas associated with the database that are used to provide high availability. Not applicable to a Hyperscale database within an elastic pool.
	HighAvailabilityReplicaCount *int `pulumi:"highAvailabilityReplicaCount"`
	// The Azure Active Directory identity of the database.
	Identity *DatabaseIdentity `pulumi:"identity"`
	// Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
	IsLedgerOn *bool `pulumi:"isLedgerOn"`
	// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType *string `pulumi:"licenseType"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The resource identifier of the long term retention backup associated with create operation of this database.
	LongTermRetentionBackupResourceId *string `pulumi:"longTermRetentionBackupResourceId"`
	// Maintenance configuration id assigned to the database. This configuration defines the period when the maintenance updates will occur.
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// The max size of the database expressed in bytes.
	MaxSizeBytes *float64 `pulumi:"maxSizeBytes"`
	// Minimal capacity that database will always have allocated, if not paused
	MinCapacity *float64 `pulumi:"minCapacity"`
	// The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region. Not applicable to a Hyperscale database within an elastic pool.
	ReadScale *string `pulumi:"readScale"`
	// The resource identifier of the recoverable database associated with create operation of this database.
	RecoverableDatabaseId *string `pulumi:"recoverableDatabaseId"`
	// The resource identifier of the recovery point associated with create operation of this database.
	RecoveryServicesRecoveryPointId *string `pulumi:"recoveryServicesRecoveryPointId"`
	// The storage account type to be used to store backups for this database.
	RequestedBackupStorageRedundancy *string `pulumi:"requestedBackupStorageRedundancy"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource identifier of the restorable dropped database associated with create operation of this database.
	RestorableDroppedDatabaseId *string `pulumi:"restorableDroppedDatabaseId"`
	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The name of the sample schema to apply when creating this database.
	SampleName *string `pulumi:"sampleName"`
	// The secondary type of the database if it is a secondary.  Valid values are Geo and Named.
	SecondaryType *string `pulumi:"secondaryType"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The database SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	Sku *Sku `pulumi:"sku"`
	// Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate *string `pulumi:"sourceDatabaseDeletionDate"`
	// The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId *string `pulumi:"sourceDatabaseId"`
	// The resource identifier of the source associated with the create operation of this database.
	//
	// This property is only supported for DataWarehouse edition and allows to restore across subscriptions.
	//
	// When sourceResourceId is specified, sourceDatabaseId, recoverableDatabaseId, restorableDroppedDatabaseId and sourceDatabaseDeletionDate must not be specified and CreateMode must be PointInTimeRestore, Restore or Recover.
	//
	// When createMode is PointInTimeRestore, sourceResourceId must be the resource ID of the existing database or existing sql pool, and restorePointInTime must be specified.
	//
	// When createMode is Restore, sourceResourceId must be the resource ID of restorable dropped database or restorable dropped sql pool.
	//
	// When createMode is Recover, sourceResourceId must be the resource ID of recoverable database or recoverable sql pool.
	//
	// When source subscription belongs to a different tenant than target subscription, “x-ms-authorization-auxiliary” header must contain authentication token for the source tenant. For more details about “x-ms-authorization-auxiliary” header see https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/authenticate-multi-tenant
	SourceResourceId *string `pulumi:"sourceResourceId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
	AutoPauseDelay pulumi.IntPtrInput
	// Collation of the metadata catalog.
	CatalogCollation pulumi.StringPtrInput
	// The collation of the database.
	Collation pulumi.StringPtrInput
	// Specifies the mode of database creation.
	//
	// Default: regular database creation.
	//
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
	//
	// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
	//
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
	//
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
	//
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//
	// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode pulumi.StringPtrInput
	// The name of the database.
	DatabaseName pulumi.StringPtrInput
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId pulumi.StringPtrInput
	// The Client id used for cross tenant per database CMK scenario
	FederatedClientId pulumi.StringPtrInput
	// The number of secondary replicas associated with the database that are used to provide high availability. Not applicable to a Hyperscale database within an elastic pool.
	HighAvailabilityReplicaCount pulumi.IntPtrInput
	// The Azure Active Directory identity of the database.
	Identity DatabaseIdentityPtrInput
	// Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
	IsLedgerOn pulumi.BoolPtrInput
	// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The resource identifier of the long term retention backup associated with create operation of this database.
	LongTermRetentionBackupResourceId pulumi.StringPtrInput
	// Maintenance configuration id assigned to the database. This configuration defines the period when the maintenance updates will occur.
	MaintenanceConfigurationId pulumi.StringPtrInput
	// The max size of the database expressed in bytes.
	MaxSizeBytes pulumi.Float64PtrInput
	// Minimal capacity that database will always have allocated, if not paused
	MinCapacity pulumi.Float64PtrInput
	// The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region. Not applicable to a Hyperscale database within an elastic pool.
	ReadScale pulumi.StringPtrInput
	// The resource identifier of the recoverable database associated with create operation of this database.
	RecoverableDatabaseId pulumi.StringPtrInput
	// The resource identifier of the recovery point associated with create operation of this database.
	RecoveryServicesRecoveryPointId pulumi.StringPtrInput
	// The storage account type to be used to store backups for this database.
	RequestedBackupStorageRedundancy pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The resource identifier of the restorable dropped database associated with create operation of this database.
	RestorableDroppedDatabaseId pulumi.StringPtrInput
	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
	RestorePointInTime pulumi.StringPtrInput
	// The name of the sample schema to apply when creating this database.
	SampleName pulumi.StringPtrInput
	// The secondary type of the database if it is a secondary.  Valid values are Geo and Named.
	SecondaryType pulumi.StringPtrInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The database SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	Sku SkuPtrInput
	// Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate pulumi.StringPtrInput
	// The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId pulumi.StringPtrInput
	// The resource identifier of the source associated with the create operation of this database.
	//
	// This property is only supported for DataWarehouse edition and allows to restore across subscriptions.
	//
	// When sourceResourceId is specified, sourceDatabaseId, recoverableDatabaseId, restorableDroppedDatabaseId and sourceDatabaseDeletionDate must not be specified and CreateMode must be PointInTimeRestore, Restore or Recover.
	//
	// When createMode is PointInTimeRestore, sourceResourceId must be the resource ID of the existing database or existing sql pool, and restorePointInTime must be specified.
	//
	// When createMode is Restore, sourceResourceId must be the resource ID of restorable dropped database or restorable dropped sql pool.
	//
	// When createMode is Recover, sourceResourceId must be the resource ID of recoverable database or recoverable sql pool.
	//
	// When source subscription belongs to a different tenant than target subscription, “x-ms-authorization-auxiliary” header must contain authentication token for the source tenant. For more details about “x-ms-authorization-auxiliary” header see https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/authenticate-multi-tenant
	SourceResourceId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

func (i *Database) ToOutput(ctx context.Context) pulumix.Output[*Database] {
	return pulumix.Output[*Database]{
		OutputState: i.ToDatabaseOutputWithContext(ctx).OutputState,
	}
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToOutput(ctx context.Context) pulumix.Output[*Database] {
	return pulumix.Output[*Database]{
		OutputState: o.OutputState,
	}
}

// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
func (o DatabaseOutput) AutoPauseDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.IntPtrOutput { return v.AutoPauseDelay }).(pulumi.IntPtrOutput)
}

// Collation of the metadata catalog.
func (o DatabaseOutput) CatalogCollation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.CatalogCollation }).(pulumi.StringPtrOutput)
}

// The collation of the database.
func (o DatabaseOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.Collation }).(pulumi.StringPtrOutput)
}

// The creation date of the database (ISO8601 format).
func (o DatabaseOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The storage account type used to store backups for this database.
func (o DatabaseOutput) CurrentBackupStorageRedundancy() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CurrentBackupStorageRedundancy }).(pulumi.StringOutput)
}

// The current service level objective name of the database.
func (o DatabaseOutput) CurrentServiceObjectiveName() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CurrentServiceObjectiveName }).(pulumi.StringOutput)
}

// The name and tier of the SKU.
func (o DatabaseOutput) CurrentSku() SkuResponseOutput {
	return o.ApplyT(func(v *Database) SkuResponseOutput { return v.CurrentSku }).(SkuResponseOutput)
}

// The ID of the database.
func (o DatabaseOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

// The default secondary region for this database.
func (o DatabaseOutput) DefaultSecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.DefaultSecondaryLocation }).(pulumi.StringOutput)
}

// This records the earliest start date and time that restore is available for this database (ISO8601 format).
func (o DatabaseOutput) EarliestRestoreDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.EarliestRestoreDate }).(pulumi.StringOutput)
}

// The resource identifier of the elastic pool containing this database.
func (o DatabaseOutput) ElasticPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.ElasticPoolId }).(pulumi.StringPtrOutput)
}

// Failover Group resource identifier that this database belongs to.
func (o DatabaseOutput) FailoverGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.FailoverGroupId }).(pulumi.StringOutput)
}

// The Client id used for cross tenant per database CMK scenario
func (o DatabaseOutput) FederatedClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.FederatedClientId }).(pulumi.StringPtrOutput)
}

// The number of secondary replicas associated with the database that are used to provide high availability. Not applicable to a Hyperscale database within an elastic pool.
func (o DatabaseOutput) HighAvailabilityReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.IntPtrOutput { return v.HighAvailabilityReplicaCount }).(pulumi.IntPtrOutput)
}

// The Azure Active Directory identity of the database.
func (o DatabaseOutput) Identity() DatabaseIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Database) DatabaseIdentityResponsePtrOutput { return v.Identity }).(DatabaseIdentityResponsePtrOutput)
}

// Infra encryption is enabled for this database.
func (o DatabaseOutput) IsInfraEncryptionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolOutput { return v.IsInfraEncryptionEnabled }).(pulumi.BoolOutput)
}

// Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
func (o DatabaseOutput) IsLedgerOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolPtrOutput { return v.IsLedgerOn }).(pulumi.BoolPtrOutput)
}

// Kind of database. This is metadata used for the Azure portal experience.
func (o DatabaseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
func (o DatabaseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o DatabaseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Maintenance configuration id assigned to the database. This configuration defines the period when the maintenance updates will occur.
func (o DatabaseOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

// Resource that manages the database.
func (o DatabaseOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

// The max log size for this database.
func (o DatabaseOutput) MaxLogSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *Database) pulumi.Float64Output { return v.MaxLogSizeBytes }).(pulumi.Float64Output)
}

// The max size of the database expressed in bytes.
func (o DatabaseOutput) MaxSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Database) pulumi.Float64PtrOutput { return v.MaxSizeBytes }).(pulumi.Float64PtrOutput)
}

// Minimal capacity that database will always have allocated, if not paused
func (o DatabaseOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Database) pulumi.Float64PtrOutput { return v.MinCapacity }).(pulumi.Float64PtrOutput)
}

// Resource name.
func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The date when database was paused by user configuration or action(ISO8601 format). Null if the database is ready.
func (o DatabaseOutput) PausedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.PausedDate }).(pulumi.StringOutput)
}

// The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region. Not applicable to a Hyperscale database within an elastic pool.
func (o DatabaseOutput) ReadScale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.ReadScale }).(pulumi.StringPtrOutput)
}

// The storage account type to be used to store backups for this database.
func (o DatabaseOutput) RequestedBackupStorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.RequestedBackupStorageRedundancy }).(pulumi.StringPtrOutput)
}

// The requested service level objective name of the database.
func (o DatabaseOutput) RequestedServiceObjectiveName() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.RequestedServiceObjectiveName }).(pulumi.StringOutput)
}

// The date when database was resumed by user action or database login (ISO8601 format). Null if the database is paused.
func (o DatabaseOutput) ResumedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ResumedDate }).(pulumi.StringOutput)
}

// The secondary type of the database if it is a secondary.  Valid values are Geo and Named.
func (o DatabaseOutput) SecondaryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.SecondaryType }).(pulumi.StringPtrOutput)
}

// The database SKU.
//
// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
func (o DatabaseOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Database) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The status of the database.
func (o DatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Resource tags.
func (o DatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Database) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o DatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
func (o DatabaseOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolPtrOutput { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseOutput{})
}
