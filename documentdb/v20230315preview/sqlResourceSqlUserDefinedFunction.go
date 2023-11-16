// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230315preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB userDefinedFunction.
type SqlResourceSqlUserDefinedFunction struct {
	pulumi.CustomResourceState

	// Identity for the resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                          `pulumi:"name"`
	Resource SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlResourceSqlUserDefinedFunction registers a new resource with the given unique name, arguments, and options.
func NewSqlResourceSqlUserDefinedFunction(ctx *pulumi.Context,
	name string, args *SqlResourceSqlUserDefinedFunctionArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlUserDefinedFunction, error) {
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
			Type: pulumi.String("azure-native:documentdb:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230301preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230415:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:SqlResourceSqlUserDefinedFunction"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SqlResourceSqlUserDefinedFunction
	err := ctx.RegisterResource("azure-native:documentdb/v20230315preview:SqlResourceSqlUserDefinedFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlResourceSqlUserDefinedFunction gets an existing SqlResourceSqlUserDefinedFunction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlResourceSqlUserDefinedFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlUserDefinedFunctionState, opts ...pulumi.ResourceOption) (*SqlResourceSqlUserDefinedFunction, error) {
	var resource SqlResourceSqlUserDefinedFunction
	err := ctx.ReadResource("azure-native:documentdb/v20230315preview:SqlResourceSqlUserDefinedFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlResourceSqlUserDefinedFunction resources.
type sqlResourceSqlUserDefinedFunctionState struct {
}

type SqlResourceSqlUserDefinedFunctionState struct {
}

func (SqlResourceSqlUserDefinedFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlUserDefinedFunctionState)(nil)).Elem()
}

type sqlResourceSqlUserDefinedFunctionArgs struct {
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
	// The standard JSON format of a userDefinedFunction
	Resource SqlUserDefinedFunctionResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// Cosmos DB userDefinedFunction name.
	UserDefinedFunctionName *string `pulumi:"userDefinedFunctionName"`
}

// The set of arguments for constructing a SqlResourceSqlUserDefinedFunction resource.
type SqlResourceSqlUserDefinedFunctionArgs struct {
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
	// The standard JSON format of a userDefinedFunction
	Resource SqlUserDefinedFunctionResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// Cosmos DB userDefinedFunction name.
	UserDefinedFunctionName pulumi.StringPtrInput
}

func (SqlResourceSqlUserDefinedFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlUserDefinedFunctionArgs)(nil)).Elem()
}

type SqlResourceSqlUserDefinedFunctionInput interface {
	pulumi.Input

	ToSqlResourceSqlUserDefinedFunctionOutput() SqlResourceSqlUserDefinedFunctionOutput
	ToSqlResourceSqlUserDefinedFunctionOutputWithContext(ctx context.Context) SqlResourceSqlUserDefinedFunctionOutput
}

func (*SqlResourceSqlUserDefinedFunction) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlUserDefinedFunction)(nil)).Elem()
}

func (i *SqlResourceSqlUserDefinedFunction) ToSqlResourceSqlUserDefinedFunctionOutput() SqlResourceSqlUserDefinedFunctionOutput {
	return i.ToSqlResourceSqlUserDefinedFunctionOutputWithContext(context.Background())
}

func (i *SqlResourceSqlUserDefinedFunction) ToSqlResourceSqlUserDefinedFunctionOutputWithContext(ctx context.Context) SqlResourceSqlUserDefinedFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlUserDefinedFunctionOutput)
}

type SqlResourceSqlUserDefinedFunctionOutput struct{ *pulumi.OutputState }

func (SqlResourceSqlUserDefinedFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlUserDefinedFunction)(nil)).Elem()
}

func (o SqlResourceSqlUserDefinedFunctionOutput) ToSqlResourceSqlUserDefinedFunctionOutput() SqlResourceSqlUserDefinedFunctionOutput {
	return o
}

func (o SqlResourceSqlUserDefinedFunctionOutput) ToSqlResourceSqlUserDefinedFunctionOutputWithContext(ctx context.Context) SqlResourceSqlUserDefinedFunctionOutput {
	return o
}

// Identity for the resource.
func (o SqlResourceSqlUserDefinedFunctionOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlUserDefinedFunction) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The location of the resource group to which the resource belongs.
func (o SqlResourceSqlUserDefinedFunctionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlUserDefinedFunction) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o SqlResourceSqlUserDefinedFunctionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlUserDefinedFunction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlResourceSqlUserDefinedFunctionOutput) Resource() SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlUserDefinedFunction) SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
		return v.Resource
	}).(SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o SqlResourceSqlUserDefinedFunctionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlResourceSqlUserDefinedFunction) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o SqlResourceSqlUserDefinedFunctionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlUserDefinedFunction) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlResourceSqlUserDefinedFunctionOutput{})
}
