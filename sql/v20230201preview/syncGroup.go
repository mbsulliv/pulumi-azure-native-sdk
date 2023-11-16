// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure SQL Database sync group.
type SyncGroup struct {
	pulumi.CustomResourceState

	// Conflict logging retention period.
	ConflictLoggingRetentionInDays pulumi.IntPtrOutput `pulumi:"conflictLoggingRetentionInDays"`
	// Conflict resolution policy of the sync group.
	ConflictResolutionPolicy pulumi.StringPtrOutput `pulumi:"conflictResolutionPolicy"`
	// If conflict logging is enabled.
	EnableConflictLogging pulumi.BoolPtrOutput `pulumi:"enableConflictLogging"`
	// User name for the sync group hub database credential.
	HubDatabaseUserName pulumi.StringPtrOutput `pulumi:"hubDatabaseUserName"`
	// Sync interval of the sync group.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// Last sync time of the sync group.
	LastSyncTime pulumi.StringOutput `pulumi:"lastSyncTime"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Private endpoint name of the sync group if use private link connection is enabled.
	PrivateEndpointName pulumi.StringOutput `pulumi:"privateEndpointName"`
	// Sync schema of the sync group.
	Schema SyncGroupSchemaResponsePtrOutput `pulumi:"schema"`
	// The name and capacity of the SKU.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// ARM resource id of the sync database in the sync group.
	SyncDatabaseId pulumi.StringPtrOutput `pulumi:"syncDatabaseId"`
	// Sync state of the sync group.
	SyncState pulumi.StringOutput `pulumi:"syncState"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// If use private link connection is enabled.
	UsePrivateLinkConnection pulumi.BoolPtrOutput `pulumi:"usePrivateLinkConnection"`
}

// NewSyncGroup registers a new resource with the given unique name, arguments, and options.
func NewSyncGroup(ctx *pulumi.Context,
	name string, args *SyncGroupArgs, opts ...pulumi.ResourceOption) (*SyncGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:SyncGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SyncGroup
	err := ctx.RegisterResource("azure-native:sql/v20230201preview:SyncGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncGroup gets an existing SyncGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncGroupState, opts ...pulumi.ResourceOption) (*SyncGroup, error) {
	var resource SyncGroup
	err := ctx.ReadResource("azure-native:sql/v20230201preview:SyncGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncGroup resources.
type syncGroupState struct {
}

type SyncGroupState struct {
}

func (SyncGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncGroupState)(nil)).Elem()
}

type syncGroupArgs struct {
	// Conflict logging retention period.
	ConflictLoggingRetentionInDays *int `pulumi:"conflictLoggingRetentionInDays"`
	// Conflict resolution policy of the sync group.
	ConflictResolutionPolicy *string `pulumi:"conflictResolutionPolicy"`
	// The name of the database on which the sync group is hosted.
	DatabaseName string `pulumi:"databaseName"`
	// If conflict logging is enabled.
	EnableConflictLogging *bool `pulumi:"enableConflictLogging"`
	// Password for the sync group hub database credential.
	HubDatabasePassword *string `pulumi:"hubDatabasePassword"`
	// User name for the sync group hub database credential.
	HubDatabaseUserName *string `pulumi:"hubDatabaseUserName"`
	// Sync interval of the sync group.
	Interval *int `pulumi:"interval"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sync schema of the sync group.
	Schema *SyncGroupSchema `pulumi:"schema"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name and capacity of the SKU.
	Sku *Sku `pulumi:"sku"`
	// ARM resource id of the sync database in the sync group.
	SyncDatabaseId *string `pulumi:"syncDatabaseId"`
	// The name of the sync group.
	SyncGroupName *string `pulumi:"syncGroupName"`
	// If use private link connection is enabled.
	UsePrivateLinkConnection *bool `pulumi:"usePrivateLinkConnection"`
}

// The set of arguments for constructing a SyncGroup resource.
type SyncGroupArgs struct {
	// Conflict logging retention period.
	ConflictLoggingRetentionInDays pulumi.IntPtrInput
	// Conflict resolution policy of the sync group.
	ConflictResolutionPolicy pulumi.StringPtrInput
	// The name of the database on which the sync group is hosted.
	DatabaseName pulumi.StringInput
	// If conflict logging is enabled.
	EnableConflictLogging pulumi.BoolPtrInput
	// Password for the sync group hub database credential.
	HubDatabasePassword pulumi.StringPtrInput
	// User name for the sync group hub database credential.
	HubDatabaseUserName pulumi.StringPtrInput
	// Sync interval of the sync group.
	Interval pulumi.IntPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Sync schema of the sync group.
	Schema SyncGroupSchemaPtrInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The name and capacity of the SKU.
	Sku SkuPtrInput
	// ARM resource id of the sync database in the sync group.
	SyncDatabaseId pulumi.StringPtrInput
	// The name of the sync group.
	SyncGroupName pulumi.StringPtrInput
	// If use private link connection is enabled.
	UsePrivateLinkConnection pulumi.BoolPtrInput
}

func (SyncGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncGroupArgs)(nil)).Elem()
}

type SyncGroupInput interface {
	pulumi.Input

	ToSyncGroupOutput() SyncGroupOutput
	ToSyncGroupOutputWithContext(ctx context.Context) SyncGroupOutput
}

func (*SyncGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroup)(nil)).Elem()
}

func (i *SyncGroup) ToSyncGroupOutput() SyncGroupOutput {
	return i.ToSyncGroupOutputWithContext(context.Background())
}

func (i *SyncGroup) ToSyncGroupOutputWithContext(ctx context.Context) SyncGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupOutput)
}

type SyncGroupOutput struct{ *pulumi.OutputState }

func (SyncGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroup)(nil)).Elem()
}

func (o SyncGroupOutput) ToSyncGroupOutput() SyncGroupOutput {
	return o
}

func (o SyncGroupOutput) ToSyncGroupOutputWithContext(ctx context.Context) SyncGroupOutput {
	return o
}

// Conflict logging retention period.
func (o SyncGroupOutput) ConflictLoggingRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.IntPtrOutput { return v.ConflictLoggingRetentionInDays }).(pulumi.IntPtrOutput)
}

// Conflict resolution policy of the sync group.
func (o SyncGroupOutput) ConflictResolutionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.StringPtrOutput { return v.ConflictResolutionPolicy }).(pulumi.StringPtrOutput)
}

// If conflict logging is enabled.
func (o SyncGroupOutput) EnableConflictLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.BoolPtrOutput { return v.EnableConflictLogging }).(pulumi.BoolPtrOutput)
}

// User name for the sync group hub database credential.
func (o SyncGroupOutput) HubDatabaseUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.StringPtrOutput { return v.HubDatabaseUserName }).(pulumi.StringPtrOutput)
}

// Sync interval of the sync group.
func (o SyncGroupOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

// Last sync time of the sync group.
func (o SyncGroupOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.StringOutput { return v.LastSyncTime }).(pulumi.StringOutput)
}

// Resource name.
func (o SyncGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint name of the sync group if use private link connection is enabled.
func (o SyncGroupOutput) PrivateEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.StringOutput { return v.PrivateEndpointName }).(pulumi.StringOutput)
}

// Sync schema of the sync group.
func (o SyncGroupOutput) Schema() SyncGroupSchemaResponsePtrOutput {
	return o.ApplyT(func(v *SyncGroup) SyncGroupSchemaResponsePtrOutput { return v.Schema }).(SyncGroupSchemaResponsePtrOutput)
}

// The name and capacity of the SKU.
func (o SyncGroupOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *SyncGroup) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// ARM resource id of the sync database in the sync group.
func (o SyncGroupOutput) SyncDatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.StringPtrOutput { return v.SyncDatabaseId }).(pulumi.StringPtrOutput)
}

// Sync state of the sync group.
func (o SyncGroupOutput) SyncState() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.StringOutput { return v.SyncState }).(pulumi.StringOutput)
}

// Resource type.
func (o SyncGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// If use private link connection is enabled.
func (o SyncGroupOutput) UsePrivateLinkConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SyncGroup) pulumi.BoolPtrOutput { return v.UsePrivateLinkConnection }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SyncGroupOutput{})
}
