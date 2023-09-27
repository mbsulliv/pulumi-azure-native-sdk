// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200320

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A cluster resource
type Cluster struct {
	pulumi.CustomResourceState

	// The identity
	ClusterId pulumi.IntOutput `pulumi:"clusterId"`
	// The cluster size
	ClusterSize pulumi.IntOutput `pulumi:"clusterSize"`
	// The hosts
	Hosts pulumi.StringArrayOutput `pulumi:"hosts"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The state of the cluster provisioning
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The cluster SKU
	Sku SkuResponseOutput `pulumi:"sku"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterSize == nil {
		return nil, errors.New("invalid value for required argument 'ClusterSize'")
	}
	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230301:Cluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:avs/v20200320:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:avs/v20200320:Cluster", name, id, state, &resource, opts...)
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
	// Name of the cluster in the private cloud
	ClusterName *string `pulumi:"clusterName"`
	// The cluster size
	ClusterSize int `pulumi:"clusterSize"`
	// The name of the private cloud.
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The cluster SKU
	Sku Sku `pulumi:"sku"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Name of the cluster in the private cloud
	ClusterName pulumi.StringPtrInput
	// The cluster size
	ClusterSize pulumi.IntInput
	// The name of the private cloud.
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The cluster SKU
	Sku SkuInput
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

func (i *Cluster) ToOutput(ctx context.Context) pulumix.Output[*Cluster] {
	return pulumix.Output[*Cluster]{
		OutputState: i.ToClusterOutputWithContext(ctx).OutputState,
	}
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

func (o ClusterOutput) ToOutput(ctx context.Context) pulumix.Output[*Cluster] {
	return pulumix.Output[*Cluster]{
		OutputState: o.OutputState,
	}
}

// The identity
func (o ClusterOutput) ClusterId() pulumi.IntOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntOutput { return v.ClusterId }).(pulumi.IntOutput)
}

// The cluster size
func (o ClusterOutput) ClusterSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntOutput { return v.ClusterSize }).(pulumi.IntOutput)
}

// The hosts
func (o ClusterOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.Hosts }).(pulumi.StringArrayOutput)
}

// Resource name.
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The state of the cluster provisioning
func (o ClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The cluster SKU
func (o ClusterOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Cluster) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// Resource type.
func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
