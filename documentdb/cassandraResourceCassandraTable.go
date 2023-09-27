// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package documentdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An Azure Cosmos DB Cassandra table.
// Azure REST API version: 2023-04-15. Prior API version in Azure Native 1.x: 2021-03-15
type CassandraResourceCassandraTable struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                  `pulumi:"name"`
	Options  CassandraTableGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource CassandraTableGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCassandraResourceCassandraTable registers a new resource with the given unique name, arguments, and options.
func NewCassandraResourceCassandraTable(ctx *pulumi.Context,
	name string, args *CassandraResourceCassandraTableArgs, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.KeyspaceName == nil {
		return nil, errors.New("invalid value for required argument 'KeyspaceName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230301preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230415:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:CassandraResourceCassandraTable"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CassandraResourceCassandraTable
	err := ctx.RegisterResource("azure-native:documentdb:CassandraResourceCassandraTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCassandraResourceCassandraTable gets an existing CassandraResourceCassandraTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCassandraResourceCassandraTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraResourceCassandraTableState, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraTable, error) {
	var resource CassandraResourceCassandraTable
	err := ctx.ReadResource("azure-native:documentdb:CassandraResourceCassandraTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CassandraResourceCassandraTable resources.
type cassandraResourceCassandraTableState struct {
}

type CassandraResourceCassandraTableState struct {
}

func (CassandraResourceCassandraTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraTableState)(nil)).Elem()
}

type cassandraResourceCassandraTableArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB keyspace name.
	KeyspaceName string `pulumi:"keyspaceName"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options *CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a Cassandra table
	Resource CassandraTableResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cosmos DB table name.
	TableName *string `pulumi:"tableName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CassandraResourceCassandraTable resource.
type CassandraResourceCassandraTableArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB keyspace name.
	KeyspaceName pulumi.StringInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsPtrInput
	// The standard JSON format of a Cassandra table
	Resource CassandraTableResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Cosmos DB table name.
	TableName pulumi.StringPtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (CassandraResourceCassandraTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraTableArgs)(nil)).Elem()
}

type CassandraResourceCassandraTableInput interface {
	pulumi.Input

	ToCassandraResourceCassandraTableOutput() CassandraResourceCassandraTableOutput
	ToCassandraResourceCassandraTableOutputWithContext(ctx context.Context) CassandraResourceCassandraTableOutput
}

func (*CassandraResourceCassandraTable) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraResourceCassandraTable)(nil)).Elem()
}

func (i *CassandraResourceCassandraTable) ToCassandraResourceCassandraTableOutput() CassandraResourceCassandraTableOutput {
	return i.ToCassandraResourceCassandraTableOutputWithContext(context.Background())
}

func (i *CassandraResourceCassandraTable) ToCassandraResourceCassandraTableOutputWithContext(ctx context.Context) CassandraResourceCassandraTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraResourceCassandraTableOutput)
}

func (i *CassandraResourceCassandraTable) ToOutput(ctx context.Context) pulumix.Output[*CassandraResourceCassandraTable] {
	return pulumix.Output[*CassandraResourceCassandraTable]{
		OutputState: i.ToCassandraResourceCassandraTableOutputWithContext(ctx).OutputState,
	}
}

type CassandraResourceCassandraTableOutput struct{ *pulumi.OutputState }

func (CassandraResourceCassandraTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraResourceCassandraTable)(nil)).Elem()
}

func (o CassandraResourceCassandraTableOutput) ToCassandraResourceCassandraTableOutput() CassandraResourceCassandraTableOutput {
	return o
}

func (o CassandraResourceCassandraTableOutput) ToCassandraResourceCassandraTableOutputWithContext(ctx context.Context) CassandraResourceCassandraTableOutput {
	return o
}

func (o CassandraResourceCassandraTableOutput) ToOutput(ctx context.Context) pulumix.Output[*CassandraResourceCassandraTable] {
	return pulumix.Output[*CassandraResourceCassandraTable]{
		OutputState: o.OutputState,
	}
}

// The location of the resource group to which the resource belongs.
func (o CassandraResourceCassandraTableOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraTable) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o CassandraResourceCassandraTableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraTable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CassandraResourceCassandraTableOutput) Options() CassandraTableGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraTable) CassandraTableGetPropertiesResponseOptionsPtrOutput {
		return v.Options
	}).(CassandraTableGetPropertiesResponseOptionsPtrOutput)
}

func (o CassandraResourceCassandraTableOutput) Resource() CassandraTableGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraTable) CassandraTableGetPropertiesResponseResourcePtrOutput {
		return v.Resource
	}).(CassandraTableGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o CassandraResourceCassandraTableOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraTable) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o CassandraResourceCassandraTableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CassandraResourceCassandraTable) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CassandraResourceCassandraTableOutput{})
}
