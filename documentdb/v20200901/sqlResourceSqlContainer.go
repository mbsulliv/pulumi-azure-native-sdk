// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB container.
//
// Deprecated: Version 2020-09-01 will be removed in v2 of the provider.
type SqlResourceSqlContainer struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                `pulumi:"name"`
	Options  SqlContainerGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource SqlContainerGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlResourceSqlContainer registers a new resource with the given unique name, arguments, and options.
func NewSqlResourceSqlContainer(ctx *pulumi.Context,
	name string, args *SqlResourceSqlContainerArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Resource = args.Resource.ToSqlContainerResourceOutput().ApplyT(func(v SqlContainerResource) SqlContainerResource { return *v.Defaults() }).(SqlContainerResourceOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:SqlResourceSqlContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlResourceSqlContainer
	err := ctx.RegisterResource("azure-native:documentdb/v20200901:SqlResourceSqlContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlResourceSqlContainer gets an existing SqlResourceSqlContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlResourceSqlContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlContainerState, opts ...pulumi.ResourceOption) (*SqlResourceSqlContainer, error) {
	var resource SqlResourceSqlContainer
	err := ctx.ReadResource("azure-native:documentdb/v20200901:SqlResourceSqlContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlResourceSqlContainer resources.
type sqlResourceSqlContainerState struct {
}

type SqlResourceSqlContainerState struct {
}

func (SqlResourceSqlContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlContainerState)(nil)).Elem()
}

type sqlResourceSqlContainerArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB container name.
	ContainerName *string `pulumi:"containerName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options *CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a container
	Resource SqlContainerResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SqlResourceSqlContainer resource.
type SqlResourceSqlContainerArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB container name.
	ContainerName pulumi.StringPtrInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsPtrInput
	// The standard JSON format of a container
	Resource SqlContainerResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (SqlResourceSqlContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlContainerArgs)(nil)).Elem()
}

type SqlResourceSqlContainerInput interface {
	pulumi.Input

	ToSqlResourceSqlContainerOutput() SqlResourceSqlContainerOutput
	ToSqlResourceSqlContainerOutputWithContext(ctx context.Context) SqlResourceSqlContainerOutput
}

func (*SqlResourceSqlContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlContainer)(nil)).Elem()
}

func (i *SqlResourceSqlContainer) ToSqlResourceSqlContainerOutput() SqlResourceSqlContainerOutput {
	return i.ToSqlResourceSqlContainerOutputWithContext(context.Background())
}

func (i *SqlResourceSqlContainer) ToSqlResourceSqlContainerOutputWithContext(ctx context.Context) SqlResourceSqlContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlContainerOutput)
}

type SqlResourceSqlContainerOutput struct{ *pulumi.OutputState }

func (SqlResourceSqlContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlContainer)(nil)).Elem()
}

func (o SqlResourceSqlContainerOutput) ToSqlResourceSqlContainerOutput() SqlResourceSqlContainerOutput {
	return o
}

func (o SqlResourceSqlContainerOutput) ToSqlResourceSqlContainerOutputWithContext(ctx context.Context) SqlResourceSqlContainerOutput {
	return o
}

// The location of the resource group to which the resource belongs.
func (o SqlResourceSqlContainerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlContainer) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o SqlResourceSqlContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlResourceSqlContainerOutput) Options() SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlContainer) SqlContainerGetPropertiesResponseOptionsPtrOutput { return v.Options }).(SqlContainerGetPropertiesResponseOptionsPtrOutput)
}

func (o SqlResourceSqlContainerOutput) Resource() SqlContainerGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlContainer) SqlContainerGetPropertiesResponseResourcePtrOutput { return v.Resource }).(SqlContainerGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o SqlResourceSqlContainerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlResourceSqlContainer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o SqlResourceSqlContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlResourceSqlContainerOutput{})
}
