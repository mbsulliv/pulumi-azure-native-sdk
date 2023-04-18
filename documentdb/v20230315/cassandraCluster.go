// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230315

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Representation of a managed Cassandra cluster.
type CassandraCluster struct {
	pulumi.CustomResourceState

	// Identity for the resource.
	Identity ManagedCassandraManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of a managed Cassandra cluster.
	Properties ClusterResourceResponsePropertiesOutput `pulumi:"properties"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCassandraCluster registers a new resource with the given unique name, arguments, and options.
func NewCassandraCluster(ctx *pulumi.Context,
	name string, args *CassandraClusterArgs, opts ...pulumi.ResourceOption) (*CassandraCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:CassandraCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:CassandraCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource CassandraCluster
	err := ctx.RegisterResource("azure-native:documentdb/v20230315:CassandraCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCassandraCluster gets an existing CassandraCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCassandraCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraClusterState, opts ...pulumi.ResourceOption) (*CassandraCluster, error) {
	var resource CassandraCluster
	err := ctx.ReadResource("azure-native:documentdb/v20230315:CassandraCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CassandraCluster resources.
type cassandraClusterState struct {
}

type CassandraClusterState struct {
}

func (CassandraClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraClusterState)(nil)).Elem()
}

type cassandraClusterArgs struct {
	// Managed Cassandra cluster name.
	ClusterName *string `pulumi:"clusterName"`
	// Identity for the resource.
	Identity *ManagedCassandraManagedServiceIdentity `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// Properties of a managed Cassandra cluster.
	Properties *ClusterResourceProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CassandraCluster resource.
type CassandraClusterArgs struct {
	// Managed Cassandra cluster name.
	ClusterName pulumi.StringPtrInput
	// Identity for the resource.
	Identity ManagedCassandraManagedServiceIdentityPtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// Properties of a managed Cassandra cluster.
	Properties ClusterResourcePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (CassandraClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraClusterArgs)(nil)).Elem()
}

type CassandraClusterInput interface {
	pulumi.Input

	ToCassandraClusterOutput() CassandraClusterOutput
	ToCassandraClusterOutputWithContext(ctx context.Context) CassandraClusterOutput
}

func (*CassandraCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraCluster)(nil)).Elem()
}

func (i *CassandraCluster) ToCassandraClusterOutput() CassandraClusterOutput {
	return i.ToCassandraClusterOutputWithContext(context.Background())
}

func (i *CassandraCluster) ToCassandraClusterOutputWithContext(ctx context.Context) CassandraClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraClusterOutput)
}

type CassandraClusterOutput struct{ *pulumi.OutputState }

func (CassandraClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraCluster)(nil)).Elem()
}

func (o CassandraClusterOutput) ToCassandraClusterOutput() CassandraClusterOutput {
	return o
}

func (o CassandraClusterOutput) ToCassandraClusterOutputWithContext(ctx context.Context) CassandraClusterOutput {
	return o
}

// Identity for the resource.
func (o CassandraClusterOutput) Identity() ManagedCassandraManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *CassandraCluster) ManagedCassandraManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedCassandraManagedServiceIdentityResponsePtrOutput)
}

// The location of the resource group to which the resource belongs.
func (o CassandraClusterOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraCluster) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o CassandraClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CassandraCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of a managed Cassandra cluster.
func (o CassandraClusterOutput) Properties() ClusterResourceResponsePropertiesOutput {
	return o.ApplyT(func(v *CassandraCluster) ClusterResourceResponsePropertiesOutput { return v.Properties }).(ClusterResourceResponsePropertiesOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o CassandraClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CassandraCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o CassandraClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CassandraCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CassandraClusterOutput{})
}
