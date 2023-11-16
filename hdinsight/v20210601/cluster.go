// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The HDInsight cluster.
type Cluster struct {
	pulumi.CustomResourceState

	// The ETag for the resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The identity of the cluster, if configured.
	Identity ClusterIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the cluster.
	Properties ClusterGetPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The availability zones.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
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
	if args.Properties != nil {
		args.Properties = args.Properties.ToClusterCreatePropertiesPtrOutput().ApplyT(func(v *ClusterCreateProperties) *ClusterCreateProperties { return v.Defaults() }).(ClusterCreatePropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hdinsight:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:hdinsight/v20150301preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:hdinsight/v20180601preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:hdinsight/v20230415preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:hdinsight/v20230815preview:Cluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:hdinsight/v20210601:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:hdinsight/v20210601:Cluster", name, id, state, &resource, opts...)
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
	// The name of the cluster.
	ClusterName *string `pulumi:"clusterName"`
	// The identity of the cluster, if configured.
	Identity *ClusterIdentity `pulumi:"identity"`
	// The location of the cluster.
	Location *string `pulumi:"location"`
	// The cluster create parameters.
	Properties *ClusterCreateProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The availability zones.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The name of the cluster.
	ClusterName pulumi.StringPtrInput
	// The identity of the cluster, if configured.
	Identity ClusterIdentityPtrInput
	// The location of the cluster.
	Location pulumi.StringPtrInput
	// The cluster create parameters.
	Properties ClusterCreatePropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The availability zones.
	Zones pulumi.StringArrayInput
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

// The ETag for the resource
func (o ClusterOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The identity of the cluster, if configured.
func (o ClusterOutput) Identity() ClusterIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterIdentityResponsePtrOutput { return v.Identity }).(ClusterIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o ClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The properties of the cluster.
func (o ClusterOutput) Properties() ClusterGetPropertiesResponseOutput {
	return o.ApplyT(func(v *Cluster) ClusterGetPropertiesResponseOutput { return v.Properties }).(ClusterGetPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
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

// The availability zones.
func (o ClusterOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
