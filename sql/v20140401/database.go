// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20140401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a database.
type Database struct {
	pulumi.CustomResourceState

	// The collation of the database. If createMode is not Default, this value is ignored.
	Collation pulumi.StringPtrOutput `pulumi:"collation"`
	// The containment state of the database.
	ContainmentState pulumi.Float64Output `pulumi:"containmentState"`
	// The creation date of the database (ISO8601 format).
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The current service level objective ID of the database. This is the ID of the service level objective that is currently active.
	CurrentServiceObjectiveId pulumi.StringOutput `pulumi:"currentServiceObjectiveId"`
	// The ID of the database.
	DatabaseId pulumi.StringOutput `pulumi:"databaseId"`
	// The default secondary region for this database.
	DefaultSecondaryLocation pulumi.StringOutput `pulumi:"defaultSecondaryLocation"`
	// This records the earliest start date and time that restore is available for this database (ISO8601 format).
	EarliestRestoreDate pulumi.StringOutput `pulumi:"earliestRestoreDate"`
	// The edition of the database. The DatabaseEditions enumeration contains all the valid editions. If createMode is NonReadableSecondary or OnlineSecondary, this value is ignored.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	Edition pulumi.StringPtrOutput `pulumi:"edition"`
	// The name of the elastic pool the database is in. If elasticPoolName and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveName is ignored. Not supported for DataWarehouse edition.
	ElasticPoolName pulumi.StringPtrOutput `pulumi:"elasticPoolName"`
	// The resource identifier of the failover group containing this database.
	FailoverGroupId pulumi.StringOutput `pulumi:"failoverGroupId"`
	// Kind of database.  This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The max size of the database expressed in bytes. If createMode is not Default, this value is ignored. To see possible values, query the capabilities API (/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationID}/capabilities) referred to by operationId: "Capabilities_ListByLocation."
	MaxSizeBytes pulumi.StringPtrOutput `pulumi:"maxSizeBytes"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Conditional. If the database is a geo-secondary, readScale indicates whether read-only connections are allowed to this database or not. Not supported for DataWarehouse edition.
	ReadScale pulumi.StringPtrOutput `pulumi:"readScale"`
	// The recommended indices for this database.
	RecommendedIndex RecommendedIndexResponseArrayOutput `pulumi:"recommendedIndex"`
	// The configured service level objective ID of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of currentServiceObjectiveId property. If requestedServiceObjectiveId and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveId overrides the value of requestedServiceObjectiveName.
	//
	// The list of SKUs may vary by region and support offer. To determine the service objective ids that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API.
	RequestedServiceObjectiveId pulumi.StringPtrOutput `pulumi:"requestedServiceObjectiveId"`
	// The name of the configured service level objective of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of serviceLevelObjective property.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	RequestedServiceObjectiveName pulumi.StringPtrOutput `pulumi:"requestedServiceObjectiveName"`
	// The current service level objective of the database.
	ServiceLevelObjective pulumi.StringOutput `pulumi:"serviceLevelObjective"`
	// The list of service tier advisors for this database. Expanded property
	ServiceTierAdvisors ServiceTierAdvisorResponseArrayOutput `pulumi:"serviceTierAdvisors"`
	// The status of the database.
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The transparent data encryption info for this database.
	TransparentDataEncryption TransparentDataEncryptionResponseArrayOutput `pulumi:"transparentDataEncryption"`
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
			Type: pulumi.String("azure-native:sql:Database"),
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
	err := ctx.RegisterResource("azure-native:sql/v20140401:Database", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:sql/v20140401:Database", name, id, state, &resource, opts...)
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
	// The collation of the database. If createMode is not Default, this value is ignored.
	Collation *string `pulumi:"collation"`
	// Specifies the mode of database creation.
	//
	// Default: regular database creation.
	//
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
	//
	// OnlineSecondary/NonReadableSecondary: creates a database as a (readable or nonreadable) secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
	//
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
	//
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
	//
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//
	// Copy, NonReadableSecondary, OnlineSecondary and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode *string `pulumi:"createMode"`
	// The name of the database to be operated on (updated or created).
	DatabaseName *string `pulumi:"databaseName"`
	// The edition of the database. The DatabaseEditions enumeration contains all the valid editions. If createMode is NonReadableSecondary or OnlineSecondary, this value is ignored.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	Edition *string `pulumi:"edition"`
	// The name of the elastic pool the database is in. If elasticPoolName and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveName is ignored. Not supported for DataWarehouse edition.
	ElasticPoolName *string `pulumi:"elasticPoolName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The max size of the database expressed in bytes. If createMode is not Default, this value is ignored. To see possible values, query the capabilities API (/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationID}/capabilities) referred to by operationId: "Capabilities_ListByLocation."
	MaxSizeBytes *string `pulumi:"maxSizeBytes"`
	// Conditional. If the database is a geo-secondary, readScale indicates whether read-only connections are allowed to this database or not. Not supported for DataWarehouse edition.
	ReadScale *string `pulumi:"readScale"`
	// Conditional. If createMode is RestoreLongTermRetentionBackup, then this value is required. Specifies the resource ID of the recovery point to restore from.
	RecoveryServicesRecoveryPointResourceId *string `pulumi:"recoveryServicesRecoveryPointResourceId"`
	// The configured service level objective ID of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of currentServiceObjectiveId property. If requestedServiceObjectiveId and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveId overrides the value of requestedServiceObjectiveName.
	//
	// The list of SKUs may vary by region and support offer. To determine the service objective ids that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API.
	RequestedServiceObjectiveId *string `pulumi:"requestedServiceObjectiveId"`
	// The name of the configured service level objective of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of serviceLevelObjective property.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	RequestedServiceObjectiveName *string `pulumi:"requestedServiceObjectiveName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Conditional. If createMode is PointInTimeRestore, this value is required. If createMode is Restore, this value is optional. Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. Must be greater than or equal to the source database's earliestRestoreDate value.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// Indicates the name of the sample schema to apply when creating this database. If createMode is not Default, this value is ignored. Not supported for DataWarehouse edition.
	SampleName *string `pulumi:"sampleName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// Conditional. If createMode is Restore and sourceDatabaseId is the deleted database's original resource id when it existed (as opposed to its current restorable dropped database id), then this value is required. Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate *string `pulumi:"sourceDatabaseDeletionDate"`
	// Conditional. If createMode is Copy, NonReadableSecondary, OnlineSecondary, PointInTimeRestore, Recovery, or Restore, then this value is required. Specifies the resource ID of the source database. If createMode is NonReadableSecondary or OnlineSecondary, the name of the source database must be the same as the new database being created.
	SourceDatabaseId *string `pulumi:"sourceDatabaseId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// The collation of the database. If createMode is not Default, this value is ignored.
	Collation pulumi.StringPtrInput
	// Specifies the mode of database creation.
	//
	// Default: regular database creation.
	//
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
	//
	// OnlineSecondary/NonReadableSecondary: creates a database as a (readable or nonreadable) secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
	//
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
	//
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
	//
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//
	// Copy, NonReadableSecondary, OnlineSecondary and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode pulumi.StringPtrInput
	// The name of the database to be operated on (updated or created).
	DatabaseName pulumi.StringPtrInput
	// The edition of the database. The DatabaseEditions enumeration contains all the valid editions. If createMode is NonReadableSecondary or OnlineSecondary, this value is ignored.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	Edition pulumi.StringPtrInput
	// The name of the elastic pool the database is in. If elasticPoolName and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveName is ignored. Not supported for DataWarehouse edition.
	ElasticPoolName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The max size of the database expressed in bytes. If createMode is not Default, this value is ignored. To see possible values, query the capabilities API (/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationID}/capabilities) referred to by operationId: "Capabilities_ListByLocation."
	MaxSizeBytes pulumi.StringPtrInput
	// Conditional. If the database is a geo-secondary, readScale indicates whether read-only connections are allowed to this database or not. Not supported for DataWarehouse edition.
	ReadScale pulumi.StringPtrInput
	// Conditional. If createMode is RestoreLongTermRetentionBackup, then this value is required. Specifies the resource ID of the recovery point to restore from.
	RecoveryServicesRecoveryPointResourceId pulumi.StringPtrInput
	// The configured service level objective ID of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of currentServiceObjectiveId property. If requestedServiceObjectiveId and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveId overrides the value of requestedServiceObjectiveName.
	//
	// The list of SKUs may vary by region and support offer. To determine the service objective ids that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API.
	RequestedServiceObjectiveId pulumi.StringPtrInput
	// The name of the configured service level objective of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of serviceLevelObjective property.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	RequestedServiceObjectiveName pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Conditional. If createMode is PointInTimeRestore, this value is required. If createMode is Restore, this value is optional. Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. Must be greater than or equal to the source database's earliestRestoreDate value.
	RestorePointInTime pulumi.StringPtrInput
	// Indicates the name of the sample schema to apply when creating this database. If createMode is not Default, this value is ignored. Not supported for DataWarehouse edition.
	SampleName pulumi.StringPtrInput
	// The name of the server.
	ServerName pulumi.StringInput
	// Conditional. If createMode is Restore and sourceDatabaseId is the deleted database's original resource id when it existed (as opposed to its current restorable dropped database id), then this value is required. Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate pulumi.StringPtrInput
	// Conditional. If createMode is Copy, NonReadableSecondary, OnlineSecondary, PointInTimeRestore, Recovery, or Restore, then this value is required. Specifies the resource ID of the source database. If createMode is NonReadableSecondary or OnlineSecondary, the name of the source database must be the same as the new database being created.
	SourceDatabaseId pulumi.StringPtrInput
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

// The collation of the database. If createMode is not Default, this value is ignored.
func (o DatabaseOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.Collation }).(pulumi.StringPtrOutput)
}

// The containment state of the database.
func (o DatabaseOutput) ContainmentState() pulumi.Float64Output {
	return o.ApplyT(func(v *Database) pulumi.Float64Output { return v.ContainmentState }).(pulumi.Float64Output)
}

// The creation date of the database (ISO8601 format).
func (o DatabaseOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The current service level objective ID of the database. This is the ID of the service level objective that is currently active.
func (o DatabaseOutput) CurrentServiceObjectiveId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CurrentServiceObjectiveId }).(pulumi.StringOutput)
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

// The edition of the database. The DatabaseEditions enumeration contains all the valid editions. If createMode is NonReadableSecondary or OnlineSecondary, this value is ignored.
//
// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
func (o DatabaseOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.Edition }).(pulumi.StringPtrOutput)
}

// The name of the elastic pool the database is in. If elasticPoolName and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveName is ignored. Not supported for DataWarehouse edition.
func (o DatabaseOutput) ElasticPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.ElasticPoolName }).(pulumi.StringPtrOutput)
}

// The resource identifier of the failover group containing this database.
func (o DatabaseOutput) FailoverGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.FailoverGroupId }).(pulumi.StringOutput)
}

// Kind of database.  This is metadata used for the Azure portal experience.
func (o DatabaseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o DatabaseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The max size of the database expressed in bytes. If createMode is not Default, this value is ignored. To see possible values, query the capabilities API (/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationID}/capabilities) referred to by operationId: "Capabilities_ListByLocation."
func (o DatabaseOutput) MaxSizeBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.MaxSizeBytes }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Conditional. If the database is a geo-secondary, readScale indicates whether read-only connections are allowed to this database or not. Not supported for DataWarehouse edition.
func (o DatabaseOutput) ReadScale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.ReadScale }).(pulumi.StringPtrOutput)
}

// The recommended indices for this database.
func (o DatabaseOutput) RecommendedIndex() RecommendedIndexResponseArrayOutput {
	return o.ApplyT(func(v *Database) RecommendedIndexResponseArrayOutput { return v.RecommendedIndex }).(RecommendedIndexResponseArrayOutput)
}

// The configured service level objective ID of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of currentServiceObjectiveId property. If requestedServiceObjectiveId and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveId overrides the value of requestedServiceObjectiveName.
//
// The list of SKUs may vary by region and support offer. To determine the service objective ids that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API.
func (o DatabaseOutput) RequestedServiceObjectiveId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.RequestedServiceObjectiveId }).(pulumi.StringPtrOutput)
}

// The name of the configured service level objective of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of serviceLevelObjective property.
//
// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
func (o DatabaseOutput) RequestedServiceObjectiveName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.RequestedServiceObjectiveName }).(pulumi.StringPtrOutput)
}

// The current service level objective of the database.
func (o DatabaseOutput) ServiceLevelObjective() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ServiceLevelObjective }).(pulumi.StringOutput)
}

// The list of service tier advisors for this database. Expanded property
func (o DatabaseOutput) ServiceTierAdvisors() ServiceTierAdvisorResponseArrayOutput {
	return o.ApplyT(func(v *Database) ServiceTierAdvisorResponseArrayOutput { return v.ServiceTierAdvisors }).(ServiceTierAdvisorResponseArrayOutput)
}

// The status of the database.
func (o DatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Resource tags.
func (o DatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Database) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The transparent data encryption info for this database.
func (o DatabaseOutput) TransparentDataEncryption() TransparentDataEncryptionResponseArrayOutput {
	return o.ApplyT(func(v *Database) TransparentDataEncryptionResponseArrayOutput { return v.TransparentDataEncryption }).(TransparentDataEncryptionResponseArrayOutput)
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
