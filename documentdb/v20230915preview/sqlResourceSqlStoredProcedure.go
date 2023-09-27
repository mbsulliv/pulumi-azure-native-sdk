// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230915preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An Azure Cosmos DB storedProcedure.
type SqlResourceSqlStoredProcedure struct {
	pulumi.CustomResourceState

	// Identity for the resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                      `pulumi:"name"`
	Resource SqlStoredProcedureGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlResourceSqlStoredProcedure registers a new resource with the given unique name, arguments, and options.
func NewSqlResourceSqlStoredProcedure(ctx *pulumi.Context,
	name string, args *SqlResourceSqlStoredProcedureArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlStoredProcedure, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230301preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230415:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915:SqlResourceSqlStoredProcedure"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SqlResourceSqlStoredProcedure
	err := ctx.RegisterResource("azure-native:documentdb/v20230915preview:SqlResourceSqlStoredProcedure", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlResourceSqlStoredProcedure gets an existing SqlResourceSqlStoredProcedure resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlResourceSqlStoredProcedure(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlStoredProcedureState, opts ...pulumi.ResourceOption) (*SqlResourceSqlStoredProcedure, error) {
	var resource SqlResourceSqlStoredProcedure
	err := ctx.ReadResource("azure-native:documentdb/v20230915preview:SqlResourceSqlStoredProcedure", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlResourceSqlStoredProcedure resources.
type sqlResourceSqlStoredProcedureState struct {
}

type SqlResourceSqlStoredProcedureState struct {
}

func (SqlResourceSqlStoredProcedureState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlStoredProcedureState)(nil)).Elem()
}

type sqlResourceSqlStoredProcedureArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB container name.
	ContainerName string `pulumi:"containerName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// Identity for the resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options *CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a storedProcedure
	Resource SqlStoredProcedureResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cosmos DB storedProcedure name.
	StoredProcedureName *string `pulumi:"storedProcedureName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SqlResourceSqlStoredProcedure resource.
type SqlResourceSqlStoredProcedureArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB container name.
	ContainerName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// Identity for the resource.
	Identity ManagedServiceIdentityPtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsPtrInput
	// The standard JSON format of a storedProcedure
	Resource SqlStoredProcedureResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Cosmos DB storedProcedure name.
	StoredProcedureName pulumi.StringPtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (SqlResourceSqlStoredProcedureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlStoredProcedureArgs)(nil)).Elem()
}

type SqlResourceSqlStoredProcedureInput interface {
	pulumi.Input

	ToSqlResourceSqlStoredProcedureOutput() SqlResourceSqlStoredProcedureOutput
	ToSqlResourceSqlStoredProcedureOutputWithContext(ctx context.Context) SqlResourceSqlStoredProcedureOutput
}

func (*SqlResourceSqlStoredProcedure) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlStoredProcedure)(nil)).Elem()
}

func (i *SqlResourceSqlStoredProcedure) ToSqlResourceSqlStoredProcedureOutput() SqlResourceSqlStoredProcedureOutput {
	return i.ToSqlResourceSqlStoredProcedureOutputWithContext(context.Background())
}

func (i *SqlResourceSqlStoredProcedure) ToSqlResourceSqlStoredProcedureOutputWithContext(ctx context.Context) SqlResourceSqlStoredProcedureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlStoredProcedureOutput)
}

func (i *SqlResourceSqlStoredProcedure) ToOutput(ctx context.Context) pulumix.Output[*SqlResourceSqlStoredProcedure] {
	return pulumix.Output[*SqlResourceSqlStoredProcedure]{
		OutputState: i.ToSqlResourceSqlStoredProcedureOutputWithContext(ctx).OutputState,
	}
}

type SqlResourceSqlStoredProcedureOutput struct{ *pulumi.OutputState }

func (SqlResourceSqlStoredProcedureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlStoredProcedure)(nil)).Elem()
}

func (o SqlResourceSqlStoredProcedureOutput) ToSqlResourceSqlStoredProcedureOutput() SqlResourceSqlStoredProcedureOutput {
	return o
}

func (o SqlResourceSqlStoredProcedureOutput) ToSqlResourceSqlStoredProcedureOutputWithContext(ctx context.Context) SqlResourceSqlStoredProcedureOutput {
	return o
}

func (o SqlResourceSqlStoredProcedureOutput) ToOutput(ctx context.Context) pulumix.Output[*SqlResourceSqlStoredProcedure] {
	return pulumix.Output[*SqlResourceSqlStoredProcedure]{
		OutputState: o.OutputState,
	}
}

// Identity for the resource.
func (o SqlResourceSqlStoredProcedureOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlStoredProcedure) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The location of the resource group to which the resource belongs.
func (o SqlResourceSqlStoredProcedureOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlStoredProcedure) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o SqlResourceSqlStoredProcedureOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlStoredProcedure) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlResourceSqlStoredProcedureOutput) Resource() SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlStoredProcedure) SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
		return v.Resource
	}).(SqlStoredProcedureGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o SqlResourceSqlStoredProcedureOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlResourceSqlStoredProcedure) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o SqlResourceSqlStoredProcedureOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlStoredProcedure) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlResourceSqlStoredProcedureOutput{})
}
