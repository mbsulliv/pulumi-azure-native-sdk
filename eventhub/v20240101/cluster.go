// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Single Event Hubs Cluster resource in List or Get operations.
type Cluster struct {
	pulumi.CustomResourceState

	// The UTC time when the Event Hubs Cluster was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The metric ID of the cluster resource. Provided by the service and not modifiable by the user.
	MetricId pulumi.StringOutput `pulumi:"metricId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the Cluster.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Properties of the cluster SKU.
	Sku ClusterSkuResponsePtrOutput `pulumi:"sku"`
	// Status of the Cluster resource
	Status pulumi.StringOutput `pulumi:"status"`
	// A value that indicates whether Scaling is Supported.
	SupportsScaling pulumi.BoolPtrOutput `pulumi:"supportsScaling"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The UTC time when the Event Hubs Cluster was last updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Properties of the cluster upgrade preferences.
	UpgradePreferences UpgradePreferencesResponsePtrOutput `pulumi:"upgradePreferences"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventhub:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20180101preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210601preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20211101:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20221001preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20230101preview:Cluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:eventhub/v20240101:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:eventhub/v20240101:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The name of the Event Hubs Cluster.
	ClusterName *string `pulumi:"clusterName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Properties of the cluster SKU.
	Sku *ClusterSku `pulumi:"sku"`
	// A value that indicates whether Scaling is Supported.
	SupportsScaling *bool `pulumi:"supportsScaling"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Properties of the cluster upgrade preferences.
	UpgradePreferences *UpgradePreferences `pulumi:"upgradePreferences"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The name of the Event Hubs Cluster.
	ClusterName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
	// Properties of the cluster SKU.
	Sku ClusterSkuPtrInput
	// A value that indicates whether Scaling is Supported.
	SupportsScaling pulumi.BoolPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Properties of the cluster upgrade preferences.
	UpgradePreferences UpgradePreferencesPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

// The UTC time when the Event Hubs Cluster was created.
func (o ClusterOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Resource location.
func (o ClusterOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The metric ID of the cluster resource. Provided by the service and not modifiable by the user.
func (o ClusterOutput) MetricId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.MetricId }).(pulumi.StringOutput)
}

// The name of the resource
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the Cluster.
func (o ClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Properties of the cluster SKU.
func (o ClusterOutput) Sku() ClusterSkuResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterSkuResponsePtrOutput { return v.Sku }).(ClusterSkuResponsePtrOutput)
}

// Status of the Cluster resource
func (o ClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A value that indicates whether Scaling is Supported.
func (o ClusterOutput) SupportsScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.SupportsScaling }).(pulumi.BoolPtrOutput)
}

// The system meta data relating to this resource.
func (o ClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Cluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The UTC time when the Event Hubs Cluster was last updated.
func (o ClusterOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Properties of the cluster upgrade preferences.
func (o ClusterOutput) UpgradePreferences() UpgradePreferencesResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) UpgradePreferencesResponsePtrOutput { return v.UpgradePreferences }).(UpgradePreferencesResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
