// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a database.
// Azure REST API version: 2022-12-29.
func LookupReadOnlyFollowingDatabase(ctx *pulumi.Context, args *LookupReadOnlyFollowingDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupReadOnlyFollowingDatabaseResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReadOnlyFollowingDatabaseResult
	err := ctx.Invoke("azure-native:kusto:getReadOnlyFollowingDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReadOnlyFollowingDatabaseArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing a read only following database.
type LookupReadOnlyFollowingDatabaseResult struct {
	// The name of the attached database configuration cluster
	AttachedDatabaseConfigurationName string `pulumi:"attachedDatabaseConfigurationName"`
	// The origin of the following setup.
	DatabaseShareOrigin string `pulumi:"databaseShareOrigin"`
	// The time the data should be kept in cache for fast queries in TimeSpan.
	HotCachePeriod *string `pulumi:"hotCachePeriod"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Kind of the database
	// Expected value is 'ReadOnlyFollowing'.
	Kind string `pulumi:"kind"`
	// The name of the leader cluster
	LeaderClusterResourceId string `pulumi:"leaderClusterResourceId"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The original database name, before databaseNameOverride or databaseNamePrefix where applied.
	OriginalDatabaseName string `pulumi:"originalDatabaseName"`
	// The principals modification kind of the database
	PrincipalsModificationKind string `pulumi:"principalsModificationKind"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The time the data should be kept before it stops being accessible to queries in TimeSpan.
	SoftDeletePeriod string `pulumi:"softDeletePeriod"`
	// The statistics of the database.
	Statistics DatabaseStatisticsResponse `pulumi:"statistics"`
	// Table level sharing specifications
	TableLevelSharingProperties TableLevelSharingPropertiesResponse `pulumi:"tableLevelSharingProperties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupReadOnlyFollowingDatabaseOutput(ctx *pulumi.Context, args LookupReadOnlyFollowingDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupReadOnlyFollowingDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReadOnlyFollowingDatabaseResult, error) {
			args := v.(LookupReadOnlyFollowingDatabaseArgs)
			r, err := LookupReadOnlyFollowingDatabase(ctx, &args, opts...)
			var s LookupReadOnlyFollowingDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReadOnlyFollowingDatabaseResultOutput)
}

type LookupReadOnlyFollowingDatabaseOutputArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupReadOnlyFollowingDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReadOnlyFollowingDatabaseArgs)(nil)).Elem()
}

// Class representing a read only following database.
type LookupReadOnlyFollowingDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupReadOnlyFollowingDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReadOnlyFollowingDatabaseResult)(nil)).Elem()
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) ToLookupReadOnlyFollowingDatabaseResultOutput() LookupReadOnlyFollowingDatabaseResultOutput {
	return o
}

func (o LookupReadOnlyFollowingDatabaseResultOutput) ToLookupReadOnlyFollowingDatabaseResultOutputWithContext(ctx context.Context) LookupReadOnlyFollowingDatabaseResultOutput {
	return o
}

// The name of the attached database configuration cluster
func (o LookupReadOnlyFollowingDatabaseResultOutput) AttachedDatabaseConfigurationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.AttachedDatabaseConfigurationName }).(pulumi.StringOutput)
}

// The origin of the following setup.
func (o LookupReadOnlyFollowingDatabaseResultOutput) DatabaseShareOrigin() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.DatabaseShareOrigin }).(pulumi.StringOutput)
}

// The time the data should be kept in cache for fast queries in TimeSpan.
func (o LookupReadOnlyFollowingDatabaseResultOutput) HotCachePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) *string { return v.HotCachePeriod }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupReadOnlyFollowingDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of the database
// Expected value is 'ReadOnlyFollowing'.
func (o LookupReadOnlyFollowingDatabaseResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the leader cluster
func (o LookupReadOnlyFollowingDatabaseResultOutput) LeaderClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.LeaderClusterResourceId }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupReadOnlyFollowingDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupReadOnlyFollowingDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// The original database name, before databaseNameOverride or databaseNamePrefix where applied.
func (o LookupReadOnlyFollowingDatabaseResultOutput) OriginalDatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.OriginalDatabaseName }).(pulumi.StringOutput)
}

// The principals modification kind of the database
func (o LookupReadOnlyFollowingDatabaseResultOutput) PrincipalsModificationKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.PrincipalsModificationKind }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o LookupReadOnlyFollowingDatabaseResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The time the data should be kept before it stops being accessible to queries in TimeSpan.
func (o LookupReadOnlyFollowingDatabaseResultOutput) SoftDeletePeriod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.SoftDeletePeriod }).(pulumi.StringOutput)
}

// The statistics of the database.
func (o LookupReadOnlyFollowingDatabaseResultOutput) Statistics() DatabaseStatisticsResponseOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) DatabaseStatisticsResponse { return v.Statistics }).(DatabaseStatisticsResponseOutput)
}

// Table level sharing specifications
func (o LookupReadOnlyFollowingDatabaseResultOutput) TableLevelSharingProperties() TableLevelSharingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) TableLevelSharingPropertiesResponse {
		return v.TableLevelSharingProperties
	}).(TableLevelSharingPropertiesResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupReadOnlyFollowingDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadOnlyFollowingDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReadOnlyFollowingDatabaseResultOutput{})
}
