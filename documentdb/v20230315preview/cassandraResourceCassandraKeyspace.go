// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230315preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An Azure Cosmos DB Cassandra keyspace.
type CassandraResourceCassandraKeyspace struct {
	pulumi.CustomResourceState

	// Identity for the resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                     `pulumi:"name"`
	Options  CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource CassandraKeyspaceGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCassandraResourceCassandraKeyspace registers a new resource with the given unique name, arguments, and options.
func NewCassandraResourceCassandraKeyspace(ctx *pulumi.Context,
	name string, args *CassandraResourceCassandraKeyspaceArgs, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraKeyspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230301preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230415:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:CassandraResourceCassandraKeyspace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CassandraResourceCassandraKeyspace
	err := ctx.RegisterResource("azure-native:documentdb/v20230315preview:CassandraResourceCassandraKeyspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCassandraResourceCassandraKeyspace gets an existing CassandraResourceCassandraKeyspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCassandraResourceCassandraKeyspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraResourceCassandraKeyspaceState, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraKeyspace, error) {
	var resource CassandraResourceCassandraKeyspace
	err := ctx.ReadResource("azure-native:documentdb/v20230315preview:CassandraResourceCassandraKeyspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CassandraResourceCassandraKeyspace resources.
type cassandraResourceCassandraKeyspaceState struct {
}

type CassandraResourceCassandraKeyspaceState struct {
}

func (CassandraResourceCassandraKeyspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraKeyspaceState)(nil)).Elem()
}

type cassandraResourceCassandraKeyspaceArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Identity for the resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Cosmos DB keyspace name.
	KeyspaceName *string `pulumi:"keyspaceName"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options *CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a Cassandra keyspace
	Resource CassandraKeyspaceResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CassandraResourceCassandraKeyspace resource.
type CassandraResourceCassandraKeyspaceArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Identity for the resource.
	Identity ManagedServiceIdentityPtrInput
	// Cosmos DB keyspace name.
	KeyspaceName pulumi.StringPtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsPtrInput
	// The standard JSON format of a Cassandra keyspace
	Resource CassandraKeyspaceResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (CassandraResourceCassandraKeyspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraKeyspaceArgs)(nil)).Elem()
}

type CassandraResourceCassandraKeyspaceInput interface {
	pulumi.Input

	ToCassandraResourceCassandraKeyspaceOutput() CassandraResourceCassandraKeyspaceOutput
	ToCassandraResourceCassandraKeyspaceOutputWithContext(ctx context.Context) CassandraResourceCassandraKeyspaceOutput
}

func (*CassandraResourceCassandraKeyspace) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraResourceCassandraKeyspace)(nil)).Elem()
}

func (i *CassandraResourceCassandraKeyspace) ToCassandraResourceCassandraKeyspaceOutput() CassandraResourceCassandraKeyspaceOutput {
	return i.ToCassandraResourceCassandraKeyspaceOutputWithContext(context.Background())
}

func (i *CassandraResourceCassandraKeyspace) ToCassandraResourceCassandraKeyspaceOutputWithContext(ctx context.Context) CassandraResourceCassandraKeyspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraResourceCassandraKeyspaceOutput)
}

func (i *CassandraResourceCassandraKeyspace) ToOutput(ctx context.Context) pulumix.Output[*CassandraResourceCassandraKeyspace] {
	return pulumix.Output[*CassandraResourceCassandraKeyspace]{
		OutputState: i.ToCassandraResourceCassandraKeyspaceOutputWithContext(ctx).OutputState,
	}
}

type CassandraResourceCassandraKeyspaceOutput struct{ *pulumi.OutputState }

func (CassandraResourceCassandraKeyspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraResourceCassandraKeyspace)(nil)).Elem()
}

func (o CassandraResourceCassandraKeyspaceOutput) ToCassandraResourceCassandraKeyspaceOutput() CassandraResourceCassandraKeyspaceOutput {
	return o
}

func (o CassandraResourceCassandraKeyspaceOutput) ToCassandraResourceCassandraKeyspaceOutputWithContext(ctx context.Context) CassandraResourceCassandraKeyspaceOutput {
	return o
}

func (o CassandraResourceCassandraKeyspaceOutput) ToOutput(ctx context.Context) pulumix.Output[*CassandraResourceCassandraKeyspace] {
	return pulumix.Output[*CassandraResourceCassandraKeyspace]{
		OutputState: o.OutputState,
	}
}

// Identity for the resource.
func (o CassandraResourceCassandraKeyspaceOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraKeyspace) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The location of the resource group to which the resource belongs.
func (o CassandraResourceCassandraKeyspaceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraKeyspace) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o CassandraResourceCassandraKeyspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraKeyspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CassandraResourceCassandraKeyspaceOutput) Options() CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraKeyspace) CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput {
		return v.Options
	}).(CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput)
}

func (o CassandraResourceCassandraKeyspaceOutput) Resource() CassandraKeyspaceGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraKeyspace) CassandraKeyspaceGetPropertiesResponseResourcePtrOutput {
		return v.Resource
	}).(CassandraKeyspaceGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o CassandraResourceCassandraKeyspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraKeyspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o CassandraResourceCassandraKeyspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraKeyspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CassandraResourceCassandraKeyspaceOutput{})
}
