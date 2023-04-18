// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A managed Cassandra data center.
type CassandraDataCenter struct {
	pulumi.CustomResourceState

	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of a managed Cassandra data center.
	Properties DataCenterResourceResponsePropertiesOutput `pulumi:"properties"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCassandraDataCenter registers a new resource with the given unique name, arguments, and options.
func NewCassandraDataCenter(ctx *pulumi.Context,
	name string, args *CassandraDataCenterArgs, opts ...pulumi.ResourceOption) (*CassandraDataCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:CassandraDataCenter"),
		},
	})
	opts = append(opts, aliases)
	var resource CassandraDataCenter
	err := ctx.RegisterResource("azure-native:documentdb/v20210401preview:CassandraDataCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCassandraDataCenter gets an existing CassandraDataCenter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCassandraDataCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraDataCenterState, opts ...pulumi.ResourceOption) (*CassandraDataCenter, error) {
	var resource CassandraDataCenter
	err := ctx.ReadResource("azure-native:documentdb/v20210401preview:CassandraDataCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CassandraDataCenter resources.
type cassandraDataCenterState struct {
}

type CassandraDataCenterState struct {
}

func (CassandraDataCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraDataCenterState)(nil)).Elem()
}

type cassandraDataCenterArgs struct {
	// Managed Cassandra cluster name.
	ClusterName string `pulumi:"clusterName"`
	// Data center name in a managed Cassandra cluster.
	DataCenterName *string `pulumi:"dataCenterName"`
	// Properties of a managed Cassandra data center.
	Properties *DataCenterResourceProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a CassandraDataCenter resource.
type CassandraDataCenterArgs struct {
	// Managed Cassandra cluster name.
	ClusterName pulumi.StringInput
	// Data center name in a managed Cassandra cluster.
	DataCenterName pulumi.StringPtrInput
	// Properties of a managed Cassandra data center.
	Properties DataCenterResourcePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (CassandraDataCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraDataCenterArgs)(nil)).Elem()
}

type CassandraDataCenterInput interface {
	pulumi.Input

	ToCassandraDataCenterOutput() CassandraDataCenterOutput
	ToCassandraDataCenterOutputWithContext(ctx context.Context) CassandraDataCenterOutput
}

func (*CassandraDataCenter) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraDataCenter)(nil)).Elem()
}

func (i *CassandraDataCenter) ToCassandraDataCenterOutput() CassandraDataCenterOutput {
	return i.ToCassandraDataCenterOutputWithContext(context.Background())
}

func (i *CassandraDataCenter) ToCassandraDataCenterOutputWithContext(ctx context.Context) CassandraDataCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraDataCenterOutput)
}

type CassandraDataCenterOutput struct{ *pulumi.OutputState }

func (CassandraDataCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraDataCenter)(nil)).Elem()
}

func (o CassandraDataCenterOutput) ToCassandraDataCenterOutput() CassandraDataCenterOutput {
	return o
}

func (o CassandraDataCenterOutput) ToCassandraDataCenterOutputWithContext(ctx context.Context) CassandraDataCenterOutput {
	return o
}

// The name of the database account.
func (o CassandraDataCenterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CassandraDataCenter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of a managed Cassandra data center.
func (o CassandraDataCenterOutput) Properties() DataCenterResourceResponsePropertiesOutput {
	return o.ApplyT(func(v *CassandraDataCenter) DataCenterResourceResponsePropertiesOutput { return v.Properties }).(DataCenterResourceResponsePropertiesOutput)
}

// The type of Azure resource.
func (o CassandraDataCenterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CassandraDataCenter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CassandraDataCenterOutput{})
}
