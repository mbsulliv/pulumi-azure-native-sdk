// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a database.
func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabaseResult
	err := ctx.Invoke("azure-native:sql/v20211101:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A database resource.
type LookupDatabaseResult struct {
	// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
	AutoPauseDelay *int `pulumi:"autoPauseDelay"`
	// Collation of the metadata catalog.
	CatalogCollation *string `pulumi:"catalogCollation"`
	// The collation of the database.
	Collation *string `pulumi:"collation"`
	// The creation date of the database (ISO8601 format).
	CreationDate string `pulumi:"creationDate"`
	// The storage account type used to store backups for this database.
	CurrentBackupStorageRedundancy string `pulumi:"currentBackupStorageRedundancy"`
	// The current service level objective name of the database.
	CurrentServiceObjectiveName string `pulumi:"currentServiceObjectiveName"`
	// The name and tier of the SKU.
	CurrentSku SkuResponse `pulumi:"currentSku"`
	// The ID of the database.
	DatabaseId string `pulumi:"databaseId"`
	// The default secondary region for this database.
	DefaultSecondaryLocation string `pulumi:"defaultSecondaryLocation"`
	// This records the earliest start date and time that restore is available for this database (ISO8601 format).
	EarliestRestoreDate string `pulumi:"earliestRestoreDate"`
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId *string `pulumi:"elasticPoolId"`
	// Failover Group resource identifier that this database belongs to.
	FailoverGroupId string `pulumi:"failoverGroupId"`
	// The Client id used for cross tenant per database CMK scenario
	FederatedClientId *string `pulumi:"federatedClientId"`
	// The number of secondary replicas associated with the database that are used to provide high availability. Not applicable to a Hyperscale database within an elastic pool.
	HighAvailabilityReplicaCount *int `pulumi:"highAvailabilityReplicaCount"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The Azure Active Directory identity of the database.
	Identity *DatabaseIdentityResponse `pulumi:"identity"`
	// Infra encryption is enabled for this database.
	IsInfraEncryptionEnabled bool `pulumi:"isInfraEncryptionEnabled"`
	// Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
	IsLedgerOn *bool `pulumi:"isLedgerOn"`
	// Kind of database. This is metadata used for the Azure portal experience.
	Kind string `pulumi:"kind"`
	// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType *string `pulumi:"licenseType"`
	// Resource location.
	Location string `pulumi:"location"`
	// Maintenance configuration id assigned to the database. This configuration defines the period when the maintenance updates will occur.
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// Resource that manages the database.
	ManagedBy string `pulumi:"managedBy"`
	// The max log size for this database.
	MaxLogSizeBytes float64 `pulumi:"maxLogSizeBytes"`
	// The max size of the database expressed in bytes.
	MaxSizeBytes *float64 `pulumi:"maxSizeBytes"`
	// Minimal capacity that database will always have allocated, if not paused
	MinCapacity *float64 `pulumi:"minCapacity"`
	// Resource name.
	Name string `pulumi:"name"`
	// The date when database was paused by user configuration or action(ISO8601 format). Null if the database is ready.
	PausedDate string `pulumi:"pausedDate"`
	// The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region. Not applicable to a Hyperscale database within an elastic pool.
	ReadScale *string `pulumi:"readScale"`
	// The storage account type to be used to store backups for this database.
	RequestedBackupStorageRedundancy *string `pulumi:"requestedBackupStorageRedundancy"`
	// The requested service level objective name of the database.
	RequestedServiceObjectiveName string `pulumi:"requestedServiceObjectiveName"`
	// The date when database was resumed by user action or database login (ISO8601 format). Null if the database is paused.
	ResumedDate string `pulumi:"resumedDate"`
	// The secondary type of the database if it is a secondary.  Valid values are Geo and Named.
	SecondaryType *string `pulumi:"secondaryType"`
	// The database SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	Sku *SkuResponse `pulumi:"sku"`
	// The status of the database.
	Status string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

func LookupDatabaseOutput(ctx *pulumi.Context, args LookupDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseResult, error) {
			args := v.(LookupDatabaseArgs)
			r, err := LookupDatabase(ctx, &args, opts...)
			var s LookupDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseResultOutput)
}

type LookupDatabaseOutputArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseArgs)(nil)).Elem()
}

// A database resource.
type LookupDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseResult)(nil)).Elem()
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutput() LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutputWithContext(ctx context.Context) LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDatabaseResult] {
	return pulumix.Output[LookupDatabaseResult]{
		OutputState: o.OutputState,
	}
}

// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
func (o LookupDatabaseResultOutput) AutoPauseDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *int { return v.AutoPauseDelay }).(pulumi.IntPtrOutput)
}

// Collation of the metadata catalog.
func (o LookupDatabaseResultOutput) CatalogCollation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.CatalogCollation }).(pulumi.StringPtrOutput)
}

// The collation of the database.
func (o LookupDatabaseResultOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

// The creation date of the database (ISO8601 format).
func (o LookupDatabaseResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// The storage account type used to store backups for this database.
func (o LookupDatabaseResultOutput) CurrentBackupStorageRedundancy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.CurrentBackupStorageRedundancy }).(pulumi.StringOutput)
}

// The current service level objective name of the database.
func (o LookupDatabaseResultOutput) CurrentServiceObjectiveName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.CurrentServiceObjectiveName }).(pulumi.StringOutput)
}

// The name and tier of the SKU.
func (o LookupDatabaseResultOutput) CurrentSku() SkuResponseOutput {
	return o.ApplyT(func(v LookupDatabaseResult) SkuResponse { return v.CurrentSku }).(SkuResponseOutput)
}

// The ID of the database.
func (o LookupDatabaseResultOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.DatabaseId }).(pulumi.StringOutput)
}

// The default secondary region for this database.
func (o LookupDatabaseResultOutput) DefaultSecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.DefaultSecondaryLocation }).(pulumi.StringOutput)
}

// This records the earliest start date and time that restore is available for this database (ISO8601 format).
func (o LookupDatabaseResultOutput) EarliestRestoreDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.EarliestRestoreDate }).(pulumi.StringOutput)
}

// The resource identifier of the elastic pool containing this database.
func (o LookupDatabaseResultOutput) ElasticPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.ElasticPoolId }).(pulumi.StringPtrOutput)
}

// Failover Group resource identifier that this database belongs to.
func (o LookupDatabaseResultOutput) FailoverGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.FailoverGroupId }).(pulumi.StringOutput)
}

// The Client id used for cross tenant per database CMK scenario
func (o LookupDatabaseResultOutput) FederatedClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.FederatedClientId }).(pulumi.StringPtrOutput)
}

// The number of secondary replicas associated with the database that are used to provide high availability. Not applicable to a Hyperscale database within an elastic pool.
func (o LookupDatabaseResultOutput) HighAvailabilityReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *int { return v.HighAvailabilityReplicaCount }).(pulumi.IntPtrOutput)
}

// Resource ID.
func (o LookupDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Azure Active Directory identity of the database.
func (o LookupDatabaseResultOutput) Identity() DatabaseIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *DatabaseIdentityResponse { return v.Identity }).(DatabaseIdentityResponsePtrOutput)
}

// Infra encryption is enabled for this database.
func (o LookupDatabaseResultOutput) IsInfraEncryptionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatabaseResult) bool { return v.IsInfraEncryptionEnabled }).(pulumi.BoolOutput)
}

// Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
func (o LookupDatabaseResultOutput) IsLedgerOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *bool { return v.IsLedgerOn }).(pulumi.BoolPtrOutput)
}

// Kind of database. This is metadata used for the Azure portal experience.
func (o LookupDatabaseResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
func (o LookupDatabaseResultOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupDatabaseResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Location }).(pulumi.StringOutput)
}

// Maintenance configuration id assigned to the database. This configuration defines the period when the maintenance updates will occur.
func (o LookupDatabaseResultOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

// Resource that manages the database.
func (o LookupDatabaseResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

// The max log size for this database.
func (o LookupDatabaseResultOutput) MaxLogSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDatabaseResult) float64 { return v.MaxLogSizeBytes }).(pulumi.Float64Output)
}

// The max size of the database expressed in bytes.
func (o LookupDatabaseResultOutput) MaxSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *float64 { return v.MaxSizeBytes }).(pulumi.Float64PtrOutput)
}

// Minimal capacity that database will always have allocated, if not paused
func (o LookupDatabaseResultOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *float64 { return v.MinCapacity }).(pulumi.Float64PtrOutput)
}

// Resource name.
func (o LookupDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// The date when database was paused by user configuration or action(ISO8601 format). Null if the database is ready.
func (o LookupDatabaseResultOutput) PausedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.PausedDate }).(pulumi.StringOutput)
}

// The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region. Not applicable to a Hyperscale database within an elastic pool.
func (o LookupDatabaseResultOutput) ReadScale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.ReadScale }).(pulumi.StringPtrOutput)
}

// The storage account type to be used to store backups for this database.
func (o LookupDatabaseResultOutput) RequestedBackupStorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.RequestedBackupStorageRedundancy }).(pulumi.StringPtrOutput)
}

// The requested service level objective name of the database.
func (o LookupDatabaseResultOutput) RequestedServiceObjectiveName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.RequestedServiceObjectiveName }).(pulumi.StringOutput)
}

// The date when database was resumed by user action or database login (ISO8601 format). Null if the database is paused.
func (o LookupDatabaseResultOutput) ResumedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ResumedDate }).(pulumi.StringOutput)
}

// The secondary type of the database if it is a secondary.  Valid values are Geo and Named.
func (o LookupDatabaseResultOutput) SecondaryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.SecondaryType }).(pulumi.StringPtrOutput)
}

// The database SKU.
//
// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
func (o LookupDatabaseResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// The status of the database.
func (o LookupDatabaseResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Status }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
func (o LookupDatabaseResultOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *bool { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseResultOutput{})
}
