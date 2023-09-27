// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230502

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns a database.
func LookupReadWriteDatabase(ctx *pulumi.Context, args *LookupReadWriteDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupReadWriteDatabaseResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReadWriteDatabaseResult
	err := ctx.Invoke("azure-native:kusto/v20230502:getReadWriteDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReadWriteDatabaseArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing a read write database.
type LookupReadWriteDatabaseResult struct {
	// The time the data should be kept in cache for fast queries in TimeSpan.
	HotCachePeriod *string `pulumi:"hotCachePeriod"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Indicates whether the database is followed.
	IsFollowed bool `pulumi:"isFollowed"`
	// KeyVault properties for the database encryption.
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	// Kind of the database
	// Expected value is 'ReadWrite'.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The time the data should be kept before it stops being accessible to queries in TimeSpan.
	SoftDeletePeriod *string `pulumi:"softDeletePeriod"`
	// The statistics of the database.
	Statistics DatabaseStatisticsResponse `pulumi:"statistics"`
	// The database suspension details. If the database is suspended, this object contains information related to the database's suspension state.
	SuspensionDetails SuspensionDetailsResponse `pulumi:"suspensionDetails"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupReadWriteDatabaseOutput(ctx *pulumi.Context, args LookupReadWriteDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupReadWriteDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReadWriteDatabaseResult, error) {
			args := v.(LookupReadWriteDatabaseArgs)
			r, err := LookupReadWriteDatabase(ctx, &args, opts...)
			var s LookupReadWriteDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReadWriteDatabaseResultOutput)
}

type LookupReadWriteDatabaseOutputArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupReadWriteDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReadWriteDatabaseArgs)(nil)).Elem()
}

// Class representing a read write database.
type LookupReadWriteDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupReadWriteDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReadWriteDatabaseResult)(nil)).Elem()
}

func (o LookupReadWriteDatabaseResultOutput) ToLookupReadWriteDatabaseResultOutput() LookupReadWriteDatabaseResultOutput {
	return o
}

func (o LookupReadWriteDatabaseResultOutput) ToLookupReadWriteDatabaseResultOutputWithContext(ctx context.Context) LookupReadWriteDatabaseResultOutput {
	return o
}

func (o LookupReadWriteDatabaseResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupReadWriteDatabaseResult] {
	return pulumix.Output[LookupReadWriteDatabaseResult]{
		OutputState: o.OutputState,
	}
}

// The time the data should be kept in cache for fast queries in TimeSpan.
func (o LookupReadWriteDatabaseResultOutput) HotCachePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) *string { return v.HotCachePeriod }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupReadWriteDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates whether the database is followed.
func (o LookupReadWriteDatabaseResultOutput) IsFollowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) bool { return v.IsFollowed }).(pulumi.BoolOutput)
}

// KeyVault properties for the database encryption.
func (o LookupReadWriteDatabaseResultOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

// Kind of the database
// Expected value is 'ReadWrite'.
func (o LookupReadWriteDatabaseResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupReadWriteDatabaseResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupReadWriteDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o LookupReadWriteDatabaseResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The time the data should be kept before it stops being accessible to queries in TimeSpan.
func (o LookupReadWriteDatabaseResultOutput) SoftDeletePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) *string { return v.SoftDeletePeriod }).(pulumi.StringPtrOutput)
}

// The statistics of the database.
func (o LookupReadWriteDatabaseResultOutput) Statistics() DatabaseStatisticsResponseOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) DatabaseStatisticsResponse { return v.Statistics }).(DatabaseStatisticsResponseOutput)
}

// The database suspension details. If the database is suspended, this object contains information related to the database's suspension state.
func (o LookupReadWriteDatabaseResultOutput) SuspensionDetails() SuspensionDetailsResponseOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) SuspensionDetailsResponse { return v.SuspensionDetails }).(SuspensionDetailsResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupReadWriteDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReadWriteDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReadWriteDatabaseResultOutput{})
}
