// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20151106

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB Cassandra keyspace.
//
// Deprecated: Version 2015-11-06 will be removed in v2 of the provider.
type DatabaseAccountCassandraKeyspace struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabaseAccountCassandraKeyspace registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccountCassandraKeyspace(ctx *pulumi.Context,
	name string, args *DatabaseAccountCassandraKeyspaceArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountCassandraKeyspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:DatabaseAccountCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:DatabaseAccountCassandraKeyspace"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountCassandraKeyspace
	err := ctx.RegisterResource("azure-native:documentdb/v20151106:DatabaseAccountCassandraKeyspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAccountCassandraKeyspace gets an existing DatabaseAccountCassandraKeyspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAccountCassandraKeyspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountCassandraKeyspaceState, opts ...pulumi.ResourceOption) (*DatabaseAccountCassandraKeyspace, error) {
	var resource DatabaseAccountCassandraKeyspace
	err := ctx.ReadResource("azure-native:documentdb/v20151106:DatabaseAccountCassandraKeyspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccountCassandraKeyspace resources.
type databaseAccountCassandraKeyspaceState struct {
}

type DatabaseAccountCassandraKeyspaceState struct {
}

func (DatabaseAccountCassandraKeyspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountCassandraKeyspaceState)(nil)).Elem()
}

type databaseAccountCassandraKeyspaceArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB keyspace name.
	KeyspaceName *string `pulumi:"keyspaceName"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options map[string]string `pulumi:"options"`
	// The standard JSON format of a Cassandra keyspace
	Resource CassandraKeyspaceResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DatabaseAccountCassandraKeyspace resource.
type DatabaseAccountCassandraKeyspaceArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB keyspace name.
	KeyspaceName pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options pulumi.StringMapInput
	// The standard JSON format of a Cassandra keyspace
	Resource CassandraKeyspaceResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountCassandraKeyspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountCassandraKeyspaceArgs)(nil)).Elem()
}

type DatabaseAccountCassandraKeyspaceInput interface {
	pulumi.Input

	ToDatabaseAccountCassandraKeyspaceOutput() DatabaseAccountCassandraKeyspaceOutput
	ToDatabaseAccountCassandraKeyspaceOutputWithContext(ctx context.Context) DatabaseAccountCassandraKeyspaceOutput
}

func (*DatabaseAccountCassandraKeyspace) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountCassandraKeyspace)(nil)).Elem()
}

func (i *DatabaseAccountCassandraKeyspace) ToDatabaseAccountCassandraKeyspaceOutput() DatabaseAccountCassandraKeyspaceOutput {
	return i.ToDatabaseAccountCassandraKeyspaceOutputWithContext(context.Background())
}

func (i *DatabaseAccountCassandraKeyspace) ToDatabaseAccountCassandraKeyspaceOutputWithContext(ctx context.Context) DatabaseAccountCassandraKeyspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountCassandraKeyspaceOutput)
}

type DatabaseAccountCassandraKeyspaceOutput struct{ *pulumi.OutputState }

func (DatabaseAccountCassandraKeyspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountCassandraKeyspace)(nil)).Elem()
}

func (o DatabaseAccountCassandraKeyspaceOutput) ToDatabaseAccountCassandraKeyspaceOutput() DatabaseAccountCassandraKeyspaceOutput {
	return o
}

func (o DatabaseAccountCassandraKeyspaceOutput) ToDatabaseAccountCassandraKeyspaceOutputWithContext(ctx context.Context) DatabaseAccountCassandraKeyspaceOutput {
	return o
}

// The location of the resource group to which the resource belongs.
func (o DatabaseAccountCassandraKeyspaceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraKeyspace) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the database account.
func (o DatabaseAccountCassandraKeyspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraKeyspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o DatabaseAccountCassandraKeyspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraKeyspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o DatabaseAccountCassandraKeyspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraKeyspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountCassandraKeyspaceOutput{})
}
