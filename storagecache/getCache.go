// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storagecache

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns a cache.
// Azure REST API version: 2023-05-01.
func LookupCache(ctx *pulumi.Context, args *LookupCacheArgs, opts ...pulumi.InvokeOption) (*LookupCacheResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCacheResult
	err := ctx.Invoke("azure-native:storagecache:getCache", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupCacheArgs struct {
	// Name of cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
	CacheName string `pulumi:"cacheName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A cache instance. Follows Azure Resource Manager standards: https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md
type LookupCacheResult struct {
	// The size of this Cache, in GB.
	CacheSizeGB *int `pulumi:"cacheSizeGB"`
	// Specifies Directory Services settings of the cache.
	DirectoryServicesSettings *CacheDirectorySettingsResponse `pulumi:"directoryServicesSettings"`
	// Specifies encryption settings of the cache.
	EncryptionSettings *CacheEncryptionSettingsResponse `pulumi:"encryptionSettings"`
	// Health of the cache.
	Health CacheHealthResponse `pulumi:"health"`
	// Resource ID of the cache.
	Id string `pulumi:"id"`
	// The identity of the cache, if configured.
	Identity *CacheIdentityResponse `pulumi:"identity"`
	// Region name string.
	Location *string `pulumi:"location"`
	// Array of IPv4 addresses that can be used by clients mounting this cache.
	MountAddresses []string `pulumi:"mountAddresses"`
	// Name of cache.
	Name string `pulumi:"name"`
	// Specifies network settings of the cache.
	NetworkSettings *CacheNetworkSettingsResponse `pulumi:"networkSettings"`
	// Specifies the priming jobs defined in the cache.
	PrimingJobs []PrimingJobResponse `pulumi:"primingJobs"`
	// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState string `pulumi:"provisioningState"`
	// Specifies security settings of the cache.
	SecuritySettings *CacheSecuritySettingsResponse `pulumi:"securitySettings"`
	// SKU for the cache.
	Sku *CacheResponseSku `pulumi:"sku"`
	// Specifies the space allocation percentage for each storage target in the cache.
	SpaceAllocation []StorageTargetSpaceAllocationResponse `pulumi:"spaceAllocation"`
	// Subnet used for the cache.
	Subnet *string `pulumi:"subnet"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Type of the cache; Microsoft.StorageCache/Cache
	Type string `pulumi:"type"`
	// Upgrade settings of the cache.
	UpgradeSettings *CacheUpgradeSettingsResponse `pulumi:"upgradeSettings"`
	// Upgrade status of the cache.
	UpgradeStatus CacheUpgradeStatusResponse `pulumi:"upgradeStatus"`
	// Availability zones for resources. This field should only contain a single element in the array.
	Zones []string `pulumi:"zones"`
}

// Defaults sets the appropriate defaults for LookupCacheResult
func (val *LookupCacheResult) Defaults() *LookupCacheResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DirectoryServicesSettings = tmp.DirectoryServicesSettings.Defaults()

	tmp.NetworkSettings = tmp.NetworkSettings.Defaults()

	return &tmp
}

func LookupCacheOutput(ctx *pulumi.Context, args LookupCacheOutputArgs, opts ...pulumi.InvokeOption) LookupCacheResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCacheResult, error) {
			args := v.(LookupCacheArgs)
			r, err := LookupCache(ctx, &args, opts...)
			var s LookupCacheResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCacheResultOutput)
}

type LookupCacheOutputArgs struct {
	// Name of cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
	CacheName pulumi.StringInput `pulumi:"cacheName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCacheOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheArgs)(nil)).Elem()
}

// A cache instance. Follows Azure Resource Manager standards: https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md
type LookupCacheResultOutput struct{ *pulumi.OutputState }

func (LookupCacheResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheResult)(nil)).Elem()
}

func (o LookupCacheResultOutput) ToLookupCacheResultOutput() LookupCacheResultOutput {
	return o
}

func (o LookupCacheResultOutput) ToLookupCacheResultOutputWithContext(ctx context.Context) LookupCacheResultOutput {
	return o
}

func (o LookupCacheResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCacheResult] {
	return pulumix.Output[LookupCacheResult]{
		OutputState: o.OutputState,
	}
}

// The size of this Cache, in GB.
func (o LookupCacheResultOutput) CacheSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *int { return v.CacheSizeGB }).(pulumi.IntPtrOutput)
}

// Specifies Directory Services settings of the cache.
func (o LookupCacheResultOutput) DirectoryServicesSettings() CacheDirectorySettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheDirectorySettingsResponse { return v.DirectoryServicesSettings }).(CacheDirectorySettingsResponsePtrOutput)
}

// Specifies encryption settings of the cache.
func (o LookupCacheResultOutput) EncryptionSettings() CacheEncryptionSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheEncryptionSettingsResponse { return v.EncryptionSettings }).(CacheEncryptionSettingsResponsePtrOutput)
}

// Health of the cache.
func (o LookupCacheResultOutput) Health() CacheHealthResponseOutput {
	return o.ApplyT(func(v LookupCacheResult) CacheHealthResponse { return v.Health }).(CacheHealthResponseOutput)
}

// Resource ID of the cache.
func (o LookupCacheResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the cache, if configured.
func (o LookupCacheResultOutput) Identity() CacheIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheIdentityResponse { return v.Identity }).(CacheIdentityResponsePtrOutput)
}

// Region name string.
func (o LookupCacheResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Array of IPv4 addresses that can be used by clients mounting this cache.
func (o LookupCacheResultOutput) MountAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCacheResult) []string { return v.MountAddresses }).(pulumi.StringArrayOutput)
}

// Name of cache.
func (o LookupCacheResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies network settings of the cache.
func (o LookupCacheResultOutput) NetworkSettings() CacheNetworkSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheNetworkSettingsResponse { return v.NetworkSettings }).(CacheNetworkSettingsResponsePtrOutput)
}

// Specifies the priming jobs defined in the cache.
func (o LookupCacheResultOutput) PrimingJobs() PrimingJobResponseArrayOutput {
	return o.ApplyT(func(v LookupCacheResult) []PrimingJobResponse { return v.PrimingJobs }).(PrimingJobResponseArrayOutput)
}

// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
func (o LookupCacheResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specifies security settings of the cache.
func (o LookupCacheResultOutput) SecuritySettings() CacheSecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheSecuritySettingsResponse { return v.SecuritySettings }).(CacheSecuritySettingsResponsePtrOutput)
}

// SKU for the cache.
func (o LookupCacheResultOutput) Sku() CacheResponseSkuPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheResponseSku { return v.Sku }).(CacheResponseSkuPtrOutput)
}

// Specifies the space allocation percentage for each storage target in the cache.
func (o LookupCacheResultOutput) SpaceAllocation() StorageTargetSpaceAllocationResponseArrayOutput {
	return o.ApplyT(func(v LookupCacheResult) []StorageTargetSpaceAllocationResponse { return v.SpaceAllocation }).(StorageTargetSpaceAllocationResponseArrayOutput)
}

// Subnet used for the cache.
func (o LookupCacheResultOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

// The system meta data relating to this resource.
func (o LookupCacheResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCacheResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupCacheResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCacheResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the cache; Microsoft.StorageCache/Cache
func (o LookupCacheResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.Type }).(pulumi.StringOutput)
}

// Upgrade settings of the cache.
func (o LookupCacheResultOutput) UpgradeSettings() CacheUpgradeSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheUpgradeSettingsResponse { return v.UpgradeSettings }).(CacheUpgradeSettingsResponsePtrOutput)
}

// Upgrade status of the cache.
func (o LookupCacheResultOutput) UpgradeStatus() CacheUpgradeStatusResponseOutput {
	return o.ApplyT(func(v LookupCacheResult) CacheUpgradeStatusResponse { return v.UpgradeStatus }).(CacheUpgradeStatusResponseOutput)
}

// Availability zones for resources. This field should only contain a single element in the array.
func (o LookupCacheResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCacheResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCacheResultOutput{})
}
