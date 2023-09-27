// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A managed database resource.
type ManagedDatabase struct {
	pulumi.CustomResourceState

	// Collation of the metadata catalog.
	CatalogCollation pulumi.StringPtrOutput `pulumi:"catalogCollation"`
	// Collation of the managed database.
	Collation pulumi.StringPtrOutput `pulumi:"collation"`
	// Creation date of the database.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Geo paired region.
	DefaultSecondaryLocation pulumi.StringOutput `pulumi:"defaultSecondaryLocation"`
	// Earliest restore point in time for point in time restore.
	EarliestRestorePoint pulumi.StringOutput `pulumi:"earliestRestorePoint"`
	// Instance Failover Group resource identifier that this managed database belongs to.
	FailoverGroupId pulumi.StringOutput `pulumi:"failoverGroupId"`
	// Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
	IsLedgerOn pulumi.BoolPtrOutput `pulumi:"isLedgerOn"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Status of the database.
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedDatabase registers a new resource with the given unique name, arguments, and options.
func NewManagedDatabase(ctx *pulumi.Context,
	name string, args *ManagedDatabaseArgs, opts ...pulumi.ResourceOption) (*ManagedDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:ManagedDatabase"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedDatabase
	err := ctx.RegisterResource("azure-native:sql/v20221101preview:ManagedDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedDatabase gets an existing ManagedDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedDatabaseState, opts ...pulumi.ResourceOption) (*ManagedDatabase, error) {
	var resource ManagedDatabase
	err := ctx.ReadResource("azure-native:sql/v20221101preview:ManagedDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedDatabase resources.
type managedDatabaseState struct {
}

type ManagedDatabaseState struct {
}

func (ManagedDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedDatabaseState)(nil)).Elem()
}

type managedDatabaseArgs struct {
	// Whether to auto complete restore of this managed database.
	AutoCompleteRestore *bool `pulumi:"autoCompleteRestore"`
	// Collation of the metadata catalog.
	CatalogCollation *string `pulumi:"catalogCollation"`
	// Collation of the managed database.
	Collation *string `pulumi:"collation"`
	// Managed database create mode. PointInTimeRestore: Create a database by restoring a point in time backup of an existing database. SourceDatabaseName, SourceManagedInstanceName and PointInTime must be specified. RestoreExternalBackup: Create a database by restoring from external backup files. Collation, StorageContainerUri and StorageContainerSasToken must be specified. Recovery: Creates a database by restoring a geo-replicated backup. RecoverableDatabaseId must be specified as the recoverable database resource ID to restore. RestoreLongTermRetentionBackup: Create a database by restoring from a long term retention backup (longTermRetentionBackupResourceId required).
	CreateMode *string `pulumi:"createMode"`
	// The restorable cross-subscription dropped database resource id to restore when creating this database.
	CrossSubscriptionRestorableDroppedDatabaseId *string `pulumi:"crossSubscriptionRestorableDroppedDatabaseId"`
	// The resource identifier of the cross-subscription source database associated with create operation of this database.
	CrossSubscriptionSourceDatabaseId *string `pulumi:"crossSubscriptionSourceDatabaseId"`
	// Target managed instance id used in cross-subscription restore.
	CrossSubscriptionTargetManagedInstanceId *string `pulumi:"crossSubscriptionTargetManagedInstanceId"`
	// The name of the database.
	DatabaseName *string `pulumi:"databaseName"`
	// Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
	IsLedgerOn *bool `pulumi:"isLedgerOn"`
	// Last backup file name for restore of this managed database.
	LastBackupName *string `pulumi:"lastBackupName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the Long Term Retention backup to be used for restore of this managed database.
	LongTermRetentionBackupResourceId *string `pulumi:"longTermRetentionBackupResourceId"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The resource identifier of the recoverable database associated with create operation of this database.
	RecoverableDatabaseId *string `pulumi:"recoverableDatabaseId"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The restorable dropped database resource id to restore when creating this database.
	RestorableDroppedDatabaseId *string `pulumi:"restorableDroppedDatabaseId"`
	// Conditional. If createMode is PointInTimeRestore, this value is required. Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId *string `pulumi:"sourceDatabaseId"`
	// Conditional. If createMode is RestoreExternalBackup, this value is used. Specifies the identity used for storage container authentication. Can be 'SharedAccessSignature' or 'ManagedIdentity'; if not specified 'SharedAccessSignature' is assumed.
	StorageContainerIdentity *string `pulumi:"storageContainerIdentity"`
	// Conditional. If createMode is RestoreExternalBackup and storageContainerIdentity is not ManagedIdentity, this value is required. Specifies the storage container sas token.
	StorageContainerSasToken *string `pulumi:"storageContainerSasToken"`
	// Conditional. If createMode is RestoreExternalBackup, this value is required. Specifies the uri of the storage container where backups for this restore are stored.
	StorageContainerUri *string `pulumi:"storageContainerUri"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ManagedDatabase resource.
type ManagedDatabaseArgs struct {
	// Whether to auto complete restore of this managed database.
	AutoCompleteRestore pulumi.BoolPtrInput
	// Collation of the metadata catalog.
	CatalogCollation pulumi.StringPtrInput
	// Collation of the managed database.
	Collation pulumi.StringPtrInput
	// Managed database create mode. PointInTimeRestore: Create a database by restoring a point in time backup of an existing database. SourceDatabaseName, SourceManagedInstanceName and PointInTime must be specified. RestoreExternalBackup: Create a database by restoring from external backup files. Collation, StorageContainerUri and StorageContainerSasToken must be specified. Recovery: Creates a database by restoring a geo-replicated backup. RecoverableDatabaseId must be specified as the recoverable database resource ID to restore. RestoreLongTermRetentionBackup: Create a database by restoring from a long term retention backup (longTermRetentionBackupResourceId required).
	CreateMode pulumi.StringPtrInput
	// The restorable cross-subscription dropped database resource id to restore when creating this database.
	CrossSubscriptionRestorableDroppedDatabaseId pulumi.StringPtrInput
	// The resource identifier of the cross-subscription source database associated with create operation of this database.
	CrossSubscriptionSourceDatabaseId pulumi.StringPtrInput
	// Target managed instance id used in cross-subscription restore.
	CrossSubscriptionTargetManagedInstanceId pulumi.StringPtrInput
	// The name of the database.
	DatabaseName pulumi.StringPtrInput
	// Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
	IsLedgerOn pulumi.BoolPtrInput
	// Last backup file name for restore of this managed database.
	LastBackupName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the Long Term Retention backup to be used for restore of this managed database.
	LongTermRetentionBackupResourceId pulumi.StringPtrInput
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput
	// The resource identifier of the recoverable database associated with create operation of this database.
	RecoverableDatabaseId pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The restorable dropped database resource id to restore when creating this database.
	RestorableDroppedDatabaseId pulumi.StringPtrInput
	// Conditional. If createMode is PointInTimeRestore, this value is required. Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
	RestorePointInTime pulumi.StringPtrInput
	// The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId pulumi.StringPtrInput
	// Conditional. If createMode is RestoreExternalBackup, this value is used. Specifies the identity used for storage container authentication. Can be 'SharedAccessSignature' or 'ManagedIdentity'; if not specified 'SharedAccessSignature' is assumed.
	StorageContainerIdentity pulumi.StringPtrInput
	// Conditional. If createMode is RestoreExternalBackup and storageContainerIdentity is not ManagedIdentity, this value is required. Specifies the storage container sas token.
	StorageContainerSasToken pulumi.StringPtrInput
	// Conditional. If createMode is RestoreExternalBackup, this value is required. Specifies the uri of the storage container where backups for this restore are stored.
	StorageContainerUri pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ManagedDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedDatabaseArgs)(nil)).Elem()
}

type ManagedDatabaseInput interface {
	pulumi.Input

	ToManagedDatabaseOutput() ManagedDatabaseOutput
	ToManagedDatabaseOutputWithContext(ctx context.Context) ManagedDatabaseOutput
}

func (*ManagedDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDatabase)(nil)).Elem()
}

func (i *ManagedDatabase) ToManagedDatabaseOutput() ManagedDatabaseOutput {
	return i.ToManagedDatabaseOutputWithContext(context.Background())
}

func (i *ManagedDatabase) ToManagedDatabaseOutputWithContext(ctx context.Context) ManagedDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDatabaseOutput)
}

func (i *ManagedDatabase) ToOutput(ctx context.Context) pulumix.Output[*ManagedDatabase] {
	return pulumix.Output[*ManagedDatabase]{
		OutputState: i.ToManagedDatabaseOutputWithContext(ctx).OutputState,
	}
}

type ManagedDatabaseOutput struct{ *pulumi.OutputState }

func (ManagedDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDatabase)(nil)).Elem()
}

func (o ManagedDatabaseOutput) ToManagedDatabaseOutput() ManagedDatabaseOutput {
	return o
}

func (o ManagedDatabaseOutput) ToManagedDatabaseOutputWithContext(ctx context.Context) ManagedDatabaseOutput {
	return o
}

func (o ManagedDatabaseOutput) ToOutput(ctx context.Context) pulumix.Output[*ManagedDatabase] {
	return pulumix.Output[*ManagedDatabase]{
		OutputState: o.OutputState,
	}
}

// Collation of the metadata catalog.
func (o ManagedDatabaseOutput) CatalogCollation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringPtrOutput { return v.CatalogCollation }).(pulumi.StringPtrOutput)
}

// Collation of the managed database.
func (o ManagedDatabaseOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringPtrOutput { return v.Collation }).(pulumi.StringPtrOutput)
}

// Creation date of the database.
func (o ManagedDatabaseOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// Geo paired region.
func (o ManagedDatabaseOutput) DefaultSecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.DefaultSecondaryLocation }).(pulumi.StringOutput)
}

// Earliest restore point in time for point in time restore.
func (o ManagedDatabaseOutput) EarliestRestorePoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.EarliestRestorePoint }).(pulumi.StringOutput)
}

// Instance Failover Group resource identifier that this managed database belongs to.
func (o ManagedDatabaseOutput) FailoverGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.FailoverGroupId }).(pulumi.StringOutput)
}

// Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
func (o ManagedDatabaseOutput) IsLedgerOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.BoolPtrOutput { return v.IsLedgerOn }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o ManagedDatabaseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o ManagedDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Status of the database.
func (o ManagedDatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Resource tags.
func (o ManagedDatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ManagedDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedDatabaseOutput{})
}
