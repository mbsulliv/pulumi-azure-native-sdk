// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cache

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes the RedisEnterprise cluster
// Azure REST API version: 2023-03-01-preview. Prior API version in Azure Native 1.x: 2021-03-01.
//
// Other available API versions: 2020-10-01-preview, 2023-07-01, 2023-08-01-preview, 2023-10-01-preview, 2023-11-01.
type RedisEnterprise struct {
	pulumi.CustomResourceState

	// Encryption-at-rest configuration for the cluster.
	Encryption ClusterPropertiesResponseEncryptionPtrOutput `pulumi:"encryption"`
	// DNS name of the cluster endpoint
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The identity of the resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion pulumi.StringPtrOutput `pulumi:"minimumTlsVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of private endpoint connections associated with the specified RedisEnterprise cluster
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Current provisioning status of the cluster
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Version of redis the cluster supports, e.g. '6'
	RedisVersion pulumi.StringOutput `pulumi:"redisVersion"`
	// Current resource status of the cluster
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The SKU to create, which affects price, performance, and features.
	Sku EnterpriseSkuResponseOutput `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The Availability Zones where this cluster will be deployed.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewRedisEnterprise registers a new resource with the given unique name, arguments, and options.
func NewRedisEnterprise(ctx *pulumi.Context,
	name string, args *RedisEnterpriseArgs, opts ...pulumi.ResourceOption) (*RedisEnterprise, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cache/v20201001preview:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210201preview:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210301:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210801:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220101:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20221101preview:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20230301preview:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20230701:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20230801preview:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20231001preview:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20231101:RedisEnterprise"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RedisEnterprise
	err := ctx.RegisterResource("azure-native:cache:RedisEnterprise", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRedisEnterprise gets an existing RedisEnterprise resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRedisEnterprise(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RedisEnterpriseState, opts ...pulumi.ResourceOption) (*RedisEnterprise, error) {
	var resource RedisEnterprise
	err := ctx.ReadResource("azure-native:cache:RedisEnterprise", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RedisEnterprise resources.
type redisEnterpriseState struct {
}

type RedisEnterpriseState struct {
}

func (RedisEnterpriseState) ElementType() reflect.Type {
	return reflect.TypeOf((*redisEnterpriseState)(nil)).Elem()
}

type redisEnterpriseArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName *string `pulumi:"clusterName"`
	// Encryption-at-rest configuration for the cluster.
	Encryption *ClusterPropertiesEncryption `pulumi:"encryption"`
	// The identity of the resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU to create, which affects price, performance, and features.
	Sku EnterpriseSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The Availability Zones where this cluster will be deployed.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a RedisEnterprise resource.
type RedisEnterpriseArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName pulumi.StringPtrInput
	// Encryption-at-rest configuration for the cluster.
	Encryption ClusterPropertiesEncryptionPtrInput
	// The identity of the resource.
	Identity ManagedServiceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The SKU to create, which affects price, performance, and features.
	Sku EnterpriseSkuInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The Availability Zones where this cluster will be deployed.
	Zones pulumi.StringArrayInput
}

func (RedisEnterpriseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*redisEnterpriseArgs)(nil)).Elem()
}

type RedisEnterpriseInput interface {
	pulumi.Input

	ToRedisEnterpriseOutput() RedisEnterpriseOutput
	ToRedisEnterpriseOutputWithContext(ctx context.Context) RedisEnterpriseOutput
}

func (*RedisEnterprise) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisEnterprise)(nil)).Elem()
}

func (i *RedisEnterprise) ToRedisEnterpriseOutput() RedisEnterpriseOutput {
	return i.ToRedisEnterpriseOutputWithContext(context.Background())
}

func (i *RedisEnterprise) ToRedisEnterpriseOutputWithContext(ctx context.Context) RedisEnterpriseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisEnterpriseOutput)
}

type RedisEnterpriseOutput struct{ *pulumi.OutputState }

func (RedisEnterpriseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisEnterprise)(nil)).Elem()
}

func (o RedisEnterpriseOutput) ToRedisEnterpriseOutput() RedisEnterpriseOutput {
	return o
}

func (o RedisEnterpriseOutput) ToRedisEnterpriseOutputWithContext(ctx context.Context) RedisEnterpriseOutput {
	return o
}

// Encryption-at-rest configuration for the cluster.
func (o RedisEnterpriseOutput) Encryption() ClusterPropertiesResponseEncryptionPtrOutput {
	return o.ApplyT(func(v *RedisEnterprise) ClusterPropertiesResponseEncryptionPtrOutput { return v.Encryption }).(ClusterPropertiesResponseEncryptionPtrOutput)
}

// DNS name of the cluster endpoint
func (o RedisEnterpriseOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// The identity of the resource.
func (o RedisEnterpriseOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *RedisEnterprise) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o RedisEnterpriseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The minimum TLS version for the cluster to support, e.g. '1.2'
func (o RedisEnterpriseOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringPtrOutput { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o RedisEnterpriseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connections associated with the specified RedisEnterprise cluster
func (o RedisEnterpriseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *RedisEnterprise) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Current provisioning status of the cluster
func (o RedisEnterpriseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Version of redis the cluster supports, e.g. '6'
func (o RedisEnterpriseOutput) RedisVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.RedisVersion }).(pulumi.StringOutput)
}

// Current resource status of the cluster
func (o RedisEnterpriseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// The SKU to create, which affects price, performance, and features.
func (o RedisEnterpriseOutput) Sku() EnterpriseSkuResponseOutput {
	return o.ApplyT(func(v *RedisEnterprise) EnterpriseSkuResponseOutput { return v.Sku }).(EnterpriseSkuResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RedisEnterpriseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RedisEnterprise) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o RedisEnterpriseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RedisEnterpriseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The Availability Zones where this cluster will be deployed.
func (o RedisEnterpriseOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(RedisEnterpriseOutput{})
}
