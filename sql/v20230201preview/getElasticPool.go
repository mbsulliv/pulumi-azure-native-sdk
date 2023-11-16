// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an elastic pool.
func LookupElasticPool(ctx *pulumi.Context, args *LookupElasticPoolArgs, opts ...pulumi.InvokeOption) (*LookupElasticPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupElasticPoolResult
	err := ctx.Invoke("azure-native:sql/v20230201preview:getElasticPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupElasticPoolArgs struct {
	// The name of the elastic pool.
	ElasticPoolName string `pulumi:"elasticPoolName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// An elastic pool.
type LookupElasticPoolResult struct {
	// Specifies the availability zone the pool's primary replica is pinned to.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The creation date of the elastic pool (ISO8601 format).
	CreationDate string `pulumi:"creationDate"`
	// The number of secondary replicas associated with the elastic pool that are used to provide high availability. Applicable only to Hyperscale elastic pools.
	HighAvailabilityReplicaCount *int `pulumi:"highAvailabilityReplicaCount"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Kind of elastic pool. This is metadata used for the Azure portal experience.
	Kind string `pulumi:"kind"`
	// The license type to apply for this elastic pool.
	LicenseType *string `pulumi:"licenseType"`
	// Resource location.
	Location string `pulumi:"location"`
	// Maintenance configuration id assigned to the elastic pool. This configuration defines the period when the maintenance updates will will occur.
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// The storage limit for the database elastic pool in bytes.
	MaxSizeBytes *float64 `pulumi:"maxSizeBytes"`
	// Minimal capacity that serverless pool will not shrink below, if not paused
	MinCapacity *float64 `pulumi:"minCapacity"`
	// Resource name.
	Name string `pulumi:"name"`
	// The per database settings for the elastic pool.
	PerDatabaseSettings *ElasticPoolPerDatabaseSettingsResponse `pulumi:"perDatabaseSettings"`
	// Type of enclave requested on the elastic pool.
	PreferredEnclaveType *string `pulumi:"preferredEnclaveType"`
	// The elastic pool SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
	Sku *SkuResponse `pulumi:"sku"`
	// The state of the elastic pool.
	State string `pulumi:"state"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

func LookupElasticPoolOutput(ctx *pulumi.Context, args LookupElasticPoolOutputArgs, opts ...pulumi.InvokeOption) LookupElasticPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupElasticPoolResult, error) {
			args := v.(LookupElasticPoolArgs)
			r, err := LookupElasticPool(ctx, &args, opts...)
			var s LookupElasticPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupElasticPoolResultOutput)
}

type LookupElasticPoolOutputArgs struct {
	// The name of the elastic pool.
	ElasticPoolName pulumi.StringInput `pulumi:"elasticPoolName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupElasticPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticPoolArgs)(nil)).Elem()
}

// An elastic pool.
type LookupElasticPoolResultOutput struct{ *pulumi.OutputState }

func (LookupElasticPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticPoolResult)(nil)).Elem()
}

func (o LookupElasticPoolResultOutput) ToLookupElasticPoolResultOutput() LookupElasticPoolResultOutput {
	return o
}

func (o LookupElasticPoolResultOutput) ToLookupElasticPoolResultOutputWithContext(ctx context.Context) LookupElasticPoolResultOutput {
	return o
}

// Specifies the availability zone the pool's primary replica is pinned to.
func (o LookupElasticPoolResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// The creation date of the elastic pool (ISO8601 format).
func (o LookupElasticPoolResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// The number of secondary replicas associated with the elastic pool that are used to provide high availability. Applicable only to Hyperscale elastic pools.
func (o LookupElasticPoolResultOutput) HighAvailabilityReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *int { return v.HighAvailabilityReplicaCount }).(pulumi.IntPtrOutput)
}

// Resource ID.
func (o LookupElasticPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of elastic pool. This is metadata used for the Azure portal experience.
func (o LookupElasticPoolResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The license type to apply for this elastic pool.
func (o LookupElasticPoolResultOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupElasticPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

// Maintenance configuration id assigned to the elastic pool. This configuration defines the period when the maintenance updates will will occur.
func (o LookupElasticPoolResultOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *string { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

// The storage limit for the database elastic pool in bytes.
func (o LookupElasticPoolResultOutput) MaxSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *float64 { return v.MaxSizeBytes }).(pulumi.Float64PtrOutput)
}

// Minimal capacity that serverless pool will not shrink below, if not paused
func (o LookupElasticPoolResultOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *float64 { return v.MinCapacity }).(pulumi.Float64PtrOutput)
}

// Resource name.
func (o LookupElasticPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// The per database settings for the elastic pool.
func (o LookupElasticPoolResultOutput) PerDatabaseSettings() ElasticPoolPerDatabaseSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *ElasticPoolPerDatabaseSettingsResponse { return v.PerDatabaseSettings }).(ElasticPoolPerDatabaseSettingsResponsePtrOutput)
}

// Type of enclave requested on the elastic pool.
func (o LookupElasticPoolResultOutput) PreferredEnclaveType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *string { return v.PreferredEnclaveType }).(pulumi.StringPtrOutput)
}

// The elastic pool SKU.
//
// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
func (o LookupElasticPoolResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// The state of the elastic pool.
func (o LookupElasticPoolResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.State }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupElasticPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupElasticPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

// Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
func (o LookupElasticPoolResultOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *bool { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupElasticPoolResultOutput{})
}
