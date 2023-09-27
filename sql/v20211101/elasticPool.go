// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An elastic pool.
type ElasticPool struct {
	pulumi.CustomResourceState

	// The creation date of the elastic pool (ISO8601 format).
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The number of secondary replicas associated with the elastic pool that are used to provide high availability. Applicable only to Hyperscale elastic pools.
	HighAvailabilityReplicaCount pulumi.IntPtrOutput `pulumi:"highAvailabilityReplicaCount"`
	// Kind of elastic pool. This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The license type to apply for this elastic pool.
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Maintenance configuration id assigned to the elastic pool. This configuration defines the period when the maintenance updates will will occur.
	MaintenanceConfigurationId pulumi.StringPtrOutput `pulumi:"maintenanceConfigurationId"`
	// The storage limit for the database elastic pool in bytes.
	MaxSizeBytes pulumi.Float64PtrOutput `pulumi:"maxSizeBytes"`
	// Minimal capacity that serverless pool will not shrink below, if not paused
	MinCapacity pulumi.Float64PtrOutput `pulumi:"minCapacity"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The per database settings for the elastic pool.
	PerDatabaseSettings ElasticPoolPerDatabaseSettingsResponsePtrOutput `pulumi:"perDatabaseSettings"`
	// The elastic pool SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The state of the elastic pool.
	State pulumi.StringOutput `pulumi:"state"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewElasticPool registers a new resource with the given unique name, arguments, and options.
func NewElasticPool(ctx *pulumi.Context,
	name string, args *ElasticPoolArgs, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20171001preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:ElasticPool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ElasticPool
	err := ctx.RegisterResource("azure-native:sql/v20211101:ElasticPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticPool gets an existing ElasticPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticPoolState, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	var resource ElasticPool
	err := ctx.ReadResource("azure-native:sql/v20211101:ElasticPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticPool resources.
type elasticPoolState struct {
}

type ElasticPoolState struct {
}

func (ElasticPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticPoolState)(nil)).Elem()
}

type elasticPoolArgs struct {
	// The name of the elastic pool.
	ElasticPoolName *string `pulumi:"elasticPoolName"`
	// The number of secondary replicas associated with the elastic pool that are used to provide high availability. Applicable only to Hyperscale elastic pools.
	HighAvailabilityReplicaCount *int `pulumi:"highAvailabilityReplicaCount"`
	// The license type to apply for this elastic pool.
	LicenseType *string `pulumi:"licenseType"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Maintenance configuration id assigned to the elastic pool. This configuration defines the period when the maintenance updates will will occur.
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// The storage limit for the database elastic pool in bytes.
	MaxSizeBytes *float64 `pulumi:"maxSizeBytes"`
	// Minimal capacity that serverless pool will not shrink below, if not paused
	MinCapacity *float64 `pulumi:"minCapacity"`
	// The per database settings for the elastic pool.
	PerDatabaseSettings *ElasticPoolPerDatabaseSettings `pulumi:"perDatabaseSettings"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The elastic pool SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a ElasticPool resource.
type ElasticPoolArgs struct {
	// The name of the elastic pool.
	ElasticPoolName pulumi.StringPtrInput
	// The number of secondary replicas associated with the elastic pool that are used to provide high availability. Applicable only to Hyperscale elastic pools.
	HighAvailabilityReplicaCount pulumi.IntPtrInput
	// The license type to apply for this elastic pool.
	LicenseType pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Maintenance configuration id assigned to the elastic pool. This configuration defines the period when the maintenance updates will will occur.
	MaintenanceConfigurationId pulumi.StringPtrInput
	// The storage limit for the database elastic pool in bytes.
	MaxSizeBytes pulumi.Float64PtrInput
	// Minimal capacity that serverless pool will not shrink below, if not paused
	MinCapacity pulumi.Float64PtrInput
	// The per database settings for the elastic pool.
	PerDatabaseSettings ElasticPoolPerDatabaseSettingsPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The elastic pool SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (ElasticPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticPoolArgs)(nil)).Elem()
}

type ElasticPoolInput interface {
	pulumi.Input

	ToElasticPoolOutput() ElasticPoolOutput
	ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput
}

func (*ElasticPool) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPool)(nil)).Elem()
}

func (i *ElasticPool) ToElasticPoolOutput() ElasticPoolOutput {
	return i.ToElasticPoolOutputWithContext(context.Background())
}

func (i *ElasticPool) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolOutput)
}

func (i *ElasticPool) ToOutput(ctx context.Context) pulumix.Output[*ElasticPool] {
	return pulumix.Output[*ElasticPool]{
		OutputState: i.ToElasticPoolOutputWithContext(ctx).OutputState,
	}
}

type ElasticPoolOutput struct{ *pulumi.OutputState }

func (ElasticPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPool)(nil)).Elem()
}

func (o ElasticPoolOutput) ToElasticPoolOutput() ElasticPoolOutput {
	return o
}

func (o ElasticPoolOutput) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return o
}

func (o ElasticPoolOutput) ToOutput(ctx context.Context) pulumix.Output[*ElasticPool] {
	return pulumix.Output[*ElasticPool]{
		OutputState: o.OutputState,
	}
}

// The creation date of the elastic pool (ISO8601 format).
func (o ElasticPoolOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The number of secondary replicas associated with the elastic pool that are used to provide high availability. Applicable only to Hyperscale elastic pools.
func (o ElasticPoolOutput) HighAvailabilityReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.IntPtrOutput { return v.HighAvailabilityReplicaCount }).(pulumi.IntPtrOutput)
}

// Kind of elastic pool. This is metadata used for the Azure portal experience.
func (o ElasticPoolOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The license type to apply for this elastic pool.
func (o ElasticPoolOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o ElasticPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Maintenance configuration id assigned to the elastic pool. This configuration defines the period when the maintenance updates will will occur.
func (o ElasticPoolOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringPtrOutput { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

// The storage limit for the database elastic pool in bytes.
func (o ElasticPoolOutput) MaxSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.Float64PtrOutput { return v.MaxSizeBytes }).(pulumi.Float64PtrOutput)
}

// Minimal capacity that serverless pool will not shrink below, if not paused
func (o ElasticPoolOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.Float64PtrOutput { return v.MinCapacity }).(pulumi.Float64PtrOutput)
}

// Resource name.
func (o ElasticPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The per database settings for the elastic pool.
func (o ElasticPoolOutput) PerDatabaseSettings() ElasticPoolPerDatabaseSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ElasticPool) ElasticPoolPerDatabaseSettingsResponsePtrOutput { return v.PerDatabaseSettings }).(ElasticPoolPerDatabaseSettingsResponsePtrOutput)
}

// The elastic pool SKU.
//
// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
func (o ElasticPoolOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ElasticPool) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The state of the elastic pool.
func (o ElasticPoolOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Resource tags.
func (o ElasticPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ElasticPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
func (o ElasticPoolOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.BoolPtrOutput { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticPoolOutput{})
}
