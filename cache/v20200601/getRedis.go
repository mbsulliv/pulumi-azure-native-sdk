// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a Redis cache (resource description).
func LookupRedis(ctx *pulumi.Context, args *LookupRedisArgs, opts ...pulumi.InvokeOption) (*LookupRedisResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRedisResult
	err := ctx.Invoke("azure-native:cache/v20200601:getRedis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRedisArgs struct {
	// The name of the Redis cache.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A single Redis item in List or Get Operation.
type LookupRedisResult struct {
	// The keys of the Redis cache - not set if this object is not the response to Create or Update redis cache
	AccessKeys RedisAccessKeysResponse `pulumi:"accessKeys"`
	// Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort *bool `pulumi:"enableNonSslPort"`
	// Redis host name.
	HostName string `pulumi:"hostName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// List of the Redis instances associated with the cache
	Instances []RedisInstanceDetailsResponse `pulumi:"instances"`
	// List of the linked servers associated with the cache
	LinkedServers []RedisLinkedServerResponse `pulumi:"linkedServers"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1', '1.2')
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// Resource name.
	Name string `pulumi:"name"`
	// Redis non-SSL port.
	Port int `pulumi:"port"`
	// List of private endpoint connection associated with the specified redis cache
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Redis instance provisioning status.
	ProvisioningState string `pulumi:"provisioningState"`
	// Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method. Default value is 'Enabled'
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
	RedisConfiguration *RedisCommonPropertiesResponseRedisConfiguration `pulumi:"redisConfiguration"`
	// Redis version.
	RedisVersion string `pulumi:"redisVersion"`
	// The number of replicas to be created per master.
	ReplicasPerMaster *int `pulumi:"replicasPerMaster"`
	// The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int `pulumi:"shardCount"`
	// The SKU of the Redis cache to deploy.
	Sku SkuResponse `pulumi:"sku"`
	// Redis SSL port.
	SslPort int `pulumi:"sslPort"`
	// Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual Network; auto assigned by default.
	StaticIP *string `pulumi:"staticIP"`
	// The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetId *string `pulumi:"subnetId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A dictionary of tenant settings
	TenantSettings map[string]string `pulumi:"tenantSettings"`
	// Resource type.
	Type string `pulumi:"type"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

// Defaults sets the appropriate defaults for LookupRedisResult
func (val *LookupRedisResult) Defaults() *LookupRedisResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EnableNonSslPort == nil {
		enableNonSslPort_ := false
		tmp.EnableNonSslPort = &enableNonSslPort_
	}
	if tmp.PublicNetworkAccess == nil {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	return &tmp
}

func LookupRedisOutput(ctx *pulumi.Context, args LookupRedisOutputArgs, opts ...pulumi.InvokeOption) LookupRedisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRedisResult, error) {
			args := v.(LookupRedisArgs)
			r, err := LookupRedis(ctx, &args, opts...)
			var s LookupRedisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRedisResultOutput)
}

type LookupRedisOutputArgs struct {
	// The name of the Redis cache.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRedisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisArgs)(nil)).Elem()
}

// A single Redis item in List or Get Operation.
type LookupRedisResultOutput struct{ *pulumi.OutputState }

func (LookupRedisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisResult)(nil)).Elem()
}

func (o LookupRedisResultOutput) ToLookupRedisResultOutput() LookupRedisResultOutput {
	return o
}

func (o LookupRedisResultOutput) ToLookupRedisResultOutputWithContext(ctx context.Context) LookupRedisResultOutput {
	return o
}

func (o LookupRedisResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRedisResult] {
	return pulumix.Output[LookupRedisResult]{
		OutputState: o.OutputState,
	}
}

// The keys of the Redis cache - not set if this object is not the response to Create or Update redis cache
func (o LookupRedisResultOutput) AccessKeys() RedisAccessKeysResponseOutput {
	return o.ApplyT(func(v LookupRedisResult) RedisAccessKeysResponse { return v.AccessKeys }).(RedisAccessKeysResponseOutput)
}

// Specifies whether the non-ssl Redis server port (6379) is enabled.
func (o LookupRedisResultOutput) EnableNonSslPort() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *bool { return v.EnableNonSslPort }).(pulumi.BoolPtrOutput)
}

// Redis host name.
func (o LookupRedisResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.HostName }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupRedisResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of the Redis instances associated with the cache
func (o LookupRedisResultOutput) Instances() RedisInstanceDetailsResponseArrayOutput {
	return o.ApplyT(func(v LookupRedisResult) []RedisInstanceDetailsResponse { return v.Instances }).(RedisInstanceDetailsResponseArrayOutput)
}

// List of the linked servers associated with the cache
func (o LookupRedisResultOutput) LinkedServers() RedisLinkedServerResponseArrayOutput {
	return o.ApplyT(func(v LookupRedisResult) []RedisLinkedServerResponse { return v.LinkedServers }).(RedisLinkedServerResponseArrayOutput)
}

// The geo-location where the resource lives
func (o LookupRedisResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.Location }).(pulumi.StringOutput)
}

// Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1', '1.2')
func (o LookupRedisResultOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *string { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupRedisResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.Name }).(pulumi.StringOutput)
}

// Redis non-SSL port.
func (o LookupRedisResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRedisResult) int { return v.Port }).(pulumi.IntOutput)
}

// List of private endpoint connection associated with the specified redis cache
func (o LookupRedisResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupRedisResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Redis instance provisioning status.
func (o LookupRedisResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method. Default value is 'Enabled'
func (o LookupRedisResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
func (o LookupRedisResultOutput) RedisConfiguration() RedisCommonPropertiesResponseRedisConfigurationPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *RedisCommonPropertiesResponseRedisConfiguration {
		return v.RedisConfiguration
	}).(RedisCommonPropertiesResponseRedisConfigurationPtrOutput)
}

// Redis version.
func (o LookupRedisResultOutput) RedisVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.RedisVersion }).(pulumi.StringOutput)
}

// The number of replicas to be created per master.
func (o LookupRedisResultOutput) ReplicasPerMaster() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *int { return v.ReplicasPerMaster }).(pulumi.IntPtrOutput)
}

// The number of shards to be created on a Premium Cluster Cache.
func (o LookupRedisResultOutput) ShardCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *int { return v.ShardCount }).(pulumi.IntPtrOutput)
}

// The SKU of the Redis cache to deploy.
func (o LookupRedisResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupRedisResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// Redis SSL port.
func (o LookupRedisResultOutput) SslPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRedisResult) int { return v.SslPort }).(pulumi.IntOutput)
}

// Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual Network; auto assigned by default.
func (o LookupRedisResultOutput) StaticIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *string { return v.StaticIP }).(pulumi.StringPtrOutput)
}

// The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
func (o LookupRedisResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupRedisResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRedisResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A dictionary of tenant settings
func (o LookupRedisResultOutput) TenantSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRedisResult) map[string]string { return v.TenantSettings }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupRedisResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.Type }).(pulumi.StringOutput)
}

// A list of availability zones denoting where the resource needs to come from.
func (o LookupRedisResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRedisResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRedisResultOutput{})
}
