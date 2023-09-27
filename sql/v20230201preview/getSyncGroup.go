// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a sync group.
func LookupSyncGroup(ctx *pulumi.Context, args *LookupSyncGroupArgs, opts ...pulumi.InvokeOption) (*LookupSyncGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSyncGroupResult
	err := ctx.Invoke("azure-native:sql/v20230201preview:getSyncGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncGroupArgs struct {
	// The name of the database on which the sync group is hosted.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the sync group.
	SyncGroupName string `pulumi:"syncGroupName"`
}

// An Azure SQL Database sync group.
type LookupSyncGroupResult struct {
	// Conflict logging retention period.
	ConflictLoggingRetentionInDays *int `pulumi:"conflictLoggingRetentionInDays"`
	// Conflict resolution policy of the sync group.
	ConflictResolutionPolicy *string `pulumi:"conflictResolutionPolicy"`
	// If conflict logging is enabled.
	EnableConflictLogging *bool `pulumi:"enableConflictLogging"`
	// User name for the sync group hub database credential.
	HubDatabaseUserName *string `pulumi:"hubDatabaseUserName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Sync interval of the sync group.
	Interval *int `pulumi:"interval"`
	// Last sync time of the sync group.
	LastSyncTime string `pulumi:"lastSyncTime"`
	// Resource name.
	Name string `pulumi:"name"`
	// Private endpoint name of the sync group if use private link connection is enabled.
	PrivateEndpointName string `pulumi:"privateEndpointName"`
	// Sync schema of the sync group.
	Schema *SyncGroupSchemaResponse `pulumi:"schema"`
	// The name and capacity of the SKU.
	Sku *SkuResponse `pulumi:"sku"`
	// ARM resource id of the sync database in the sync group.
	SyncDatabaseId *string `pulumi:"syncDatabaseId"`
	// Sync state of the sync group.
	SyncState string `pulumi:"syncState"`
	// Resource type.
	Type string `pulumi:"type"`
	// If use private link connection is enabled.
	UsePrivateLinkConnection *bool `pulumi:"usePrivateLinkConnection"`
}

func LookupSyncGroupOutput(ctx *pulumi.Context, args LookupSyncGroupOutputArgs, opts ...pulumi.InvokeOption) LookupSyncGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSyncGroupResult, error) {
			args := v.(LookupSyncGroupArgs)
			r, err := LookupSyncGroup(ctx, &args, opts...)
			var s LookupSyncGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSyncGroupResultOutput)
}

type LookupSyncGroupOutputArgs struct {
	// The name of the database on which the sync group is hosted.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
	// The name of the sync group.
	SyncGroupName pulumi.StringInput `pulumi:"syncGroupName"`
}

func (LookupSyncGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncGroupArgs)(nil)).Elem()
}

// An Azure SQL Database sync group.
type LookupSyncGroupResultOutput struct{ *pulumi.OutputState }

func (LookupSyncGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSyncGroupResult)(nil)).Elem()
}

func (o LookupSyncGroupResultOutput) ToLookupSyncGroupResultOutput() LookupSyncGroupResultOutput {
	return o
}

func (o LookupSyncGroupResultOutput) ToLookupSyncGroupResultOutputWithContext(ctx context.Context) LookupSyncGroupResultOutput {
	return o
}

func (o LookupSyncGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSyncGroupResult] {
	return pulumix.Output[LookupSyncGroupResult]{
		OutputState: o.OutputState,
	}
}

// Conflict logging retention period.
func (o LookupSyncGroupResultOutput) ConflictLoggingRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *int { return v.ConflictLoggingRetentionInDays }).(pulumi.IntPtrOutput)
}

// Conflict resolution policy of the sync group.
func (o LookupSyncGroupResultOutput) ConflictResolutionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *string { return v.ConflictResolutionPolicy }).(pulumi.StringPtrOutput)
}

// If conflict logging is enabled.
func (o LookupSyncGroupResultOutput) EnableConflictLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *bool { return v.EnableConflictLogging }).(pulumi.BoolPtrOutput)
}

// User name for the sync group hub database credential.
func (o LookupSyncGroupResultOutput) HubDatabaseUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *string { return v.HubDatabaseUserName }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupSyncGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Sync interval of the sync group.
func (o LookupSyncGroupResultOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

// Last sync time of the sync group.
func (o LookupSyncGroupResultOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.LastSyncTime }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupSyncGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint name of the sync group if use private link connection is enabled.
func (o LookupSyncGroupResultOutput) PrivateEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.PrivateEndpointName }).(pulumi.StringOutput)
}

// Sync schema of the sync group.
func (o LookupSyncGroupResultOutput) Schema() SyncGroupSchemaResponsePtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *SyncGroupSchemaResponse { return v.Schema }).(SyncGroupSchemaResponsePtrOutput)
}

// The name and capacity of the SKU.
func (o LookupSyncGroupResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// ARM resource id of the sync database in the sync group.
func (o LookupSyncGroupResultOutput) SyncDatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *string { return v.SyncDatabaseId }).(pulumi.StringPtrOutput)
}

// Sync state of the sync group.
func (o LookupSyncGroupResultOutput) SyncState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.SyncState }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupSyncGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// If use private link connection is enabled.
func (o LookupSyncGroupResultOutput) UsePrivateLinkConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSyncGroupResult) *bool { return v.UsePrivateLinkConnection }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSyncGroupResultOutput{})
}
