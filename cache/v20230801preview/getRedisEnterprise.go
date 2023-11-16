// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a RedisEnterprise cluster
func LookupRedisEnterprise(ctx *pulumi.Context, args *LookupRedisEnterpriseArgs, opts ...pulumi.InvokeOption) (*LookupRedisEnterpriseResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRedisEnterpriseResult
	err := ctx.Invoke("azure-native:cache/v20230801preview:getRedisEnterprise", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRedisEnterpriseArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes the RedisEnterprise cluster
type LookupRedisEnterpriseResult struct {
	// Encryption-at-rest configuration for the cluster.
	Encryption *ClusterPropertiesResponseEncryption `pulumi:"encryption"`
	// DNS name of the cluster endpoint
	HostName string `pulumi:"hostName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The identity of the resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of private endpoint connections associated with the specified RedisEnterprise cluster
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Current provisioning status of the cluster
	ProvisioningState string `pulumi:"provisioningState"`
	// Version of redis the cluster supports, e.g. '6'
	RedisVersion string `pulumi:"redisVersion"`
	// Current resource status of the cluster
	ResourceState string `pulumi:"resourceState"`
	// The SKU to create, which affects price, performance, and features.
	Sku EnterpriseSkuResponse `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The Availability Zones where this cluster will be deployed.
	Zones []string `pulumi:"zones"`
}

func LookupRedisEnterpriseOutput(ctx *pulumi.Context, args LookupRedisEnterpriseOutputArgs, opts ...pulumi.InvokeOption) LookupRedisEnterpriseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRedisEnterpriseResult, error) {
			args := v.(LookupRedisEnterpriseArgs)
			r, err := LookupRedisEnterprise(ctx, &args, opts...)
			var s LookupRedisEnterpriseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRedisEnterpriseResultOutput)
}

type LookupRedisEnterpriseOutputArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRedisEnterpriseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisEnterpriseArgs)(nil)).Elem()
}

// Describes the RedisEnterprise cluster
type LookupRedisEnterpriseResultOutput struct{ *pulumi.OutputState }

func (LookupRedisEnterpriseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisEnterpriseResult)(nil)).Elem()
}

func (o LookupRedisEnterpriseResultOutput) ToLookupRedisEnterpriseResultOutput() LookupRedisEnterpriseResultOutput {
	return o
}

func (o LookupRedisEnterpriseResultOutput) ToLookupRedisEnterpriseResultOutputWithContext(ctx context.Context) LookupRedisEnterpriseResultOutput {
	return o
}

// Encryption-at-rest configuration for the cluster.
func (o LookupRedisEnterpriseResultOutput) Encryption() ClusterPropertiesResponseEncryptionPtrOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) *ClusterPropertiesResponseEncryption { return v.Encryption }).(ClusterPropertiesResponseEncryptionPtrOutput)
}

// DNS name of the cluster endpoint
func (o LookupRedisEnterpriseResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.HostName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRedisEnterpriseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the resource.
func (o LookupRedisEnterpriseResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupRedisEnterpriseResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.Location }).(pulumi.StringOutput)
}

// The minimum TLS version for the cluster to support, e.g. '1.2'
func (o LookupRedisEnterpriseResultOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) *string { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupRedisEnterpriseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connections associated with the specified RedisEnterprise cluster
func (o LookupRedisEnterpriseResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Current provisioning status of the cluster
func (o LookupRedisEnterpriseResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Version of redis the cluster supports, e.g. '6'
func (o LookupRedisEnterpriseResultOutput) RedisVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.RedisVersion }).(pulumi.StringOutput)
}

// Current resource status of the cluster
func (o LookupRedisEnterpriseResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

// The SKU to create, which affects price, performance, and features.
func (o LookupRedisEnterpriseResultOutput) Sku() EnterpriseSkuResponseOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) EnterpriseSkuResponse { return v.Sku }).(EnterpriseSkuResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRedisEnterpriseResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupRedisEnterpriseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRedisEnterpriseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.Type }).(pulumi.StringOutput)
}

// The Availability Zones where this cluster will be deployed.
func (o LookupRedisEnterpriseResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRedisEnterpriseResultOutput{})
}
