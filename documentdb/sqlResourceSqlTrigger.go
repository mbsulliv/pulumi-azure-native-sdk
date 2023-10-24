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

// An Azure Cosmos DB trigger.
// Azure REST API version: 2023-04-15. Prior API version in Azure Native 1.x: 2021-03-15.
//
// Other available API versions: 2019-08-01, 2023-03-15-preview, 2023-09-15, 2023-09-15-preview.
type SqlResourceSqlTrigger struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                              `pulumi:"name"`
	Resource SqlTriggerGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlResourceSqlTrigger registers a new resource with the given unique name, arguments, and options.
func NewSqlResourceSqlTrigger(ctx *pulumi.Context,
	name string, args *SqlResourceSqlTriggerArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlTrigger, error) {
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
			Type: pulumi.String("azure-native:documentdb/v20190801:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230301preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230415:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915:SqlResourceSqlTrigger"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:SqlResourceSqlTrigger"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SqlResourceSqlTrigger
	err := ctx.RegisterResource("azure-native:documentdb:SqlResourceSqlTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlResourceSqlTrigger gets an existing SqlResourceSqlTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlResourceSqlTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlTriggerState, opts ...pulumi.ResourceOption) (*SqlResourceSqlTrigger, error) {
	var resource SqlResourceSqlTrigger
	err := ctx.ReadResource("azure-native:documentdb:SqlResourceSqlTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlResourceSqlTrigger resources.
type sqlResourceSqlTriggerState struct {
}

type SqlResourceSqlTriggerState struct {
}

func (SqlResourceSqlTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlTriggerState)(nil)).Elem()
}

type sqlResourceSqlTriggerArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB container name.
	ContainerName string `pulumi:"containerName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options *CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a trigger
	Resource SqlTriggerResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// Cosmos DB trigger name.
	TriggerName *string `pulumi:"triggerName"`
}

// The set of arguments for constructing a SqlResourceSqlTrigger resource.
type SqlResourceSqlTriggerArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB container name.
	ContainerName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsPtrInput
	// The standard JSON format of a trigger
	Resource SqlTriggerResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// Cosmos DB trigger name.
	TriggerName pulumi.StringPtrInput
}

func (SqlResourceSqlTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlTriggerArgs)(nil)).Elem()
}

type SqlResourceSqlTriggerInput interface {
	pulumi.Input

	ToSqlResourceSqlTriggerOutput() SqlResourceSqlTriggerOutput
	ToSqlResourceSqlTriggerOutputWithContext(ctx context.Context) SqlResourceSqlTriggerOutput
}

func (*SqlResourceSqlTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlTrigger)(nil)).Elem()
}

func (i *SqlResourceSqlTrigger) ToSqlResourceSqlTriggerOutput() SqlResourceSqlTriggerOutput {
	return i.ToSqlResourceSqlTriggerOutputWithContext(context.Background())
}

func (i *SqlResourceSqlTrigger) ToSqlResourceSqlTriggerOutputWithContext(ctx context.Context) SqlResourceSqlTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlTriggerOutput)
}

func (i *SqlResourceSqlTrigger) ToOutput(ctx context.Context) pulumix.Output[*SqlResourceSqlTrigger] {
	return pulumix.Output[*SqlResourceSqlTrigger]{
		OutputState: i.ToSqlResourceSqlTriggerOutputWithContext(ctx).OutputState,
	}
}

type SqlResourceSqlTriggerOutput struct{ *pulumi.OutputState }

func (SqlResourceSqlTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlTrigger)(nil)).Elem()
}

func (o SqlResourceSqlTriggerOutput) ToSqlResourceSqlTriggerOutput() SqlResourceSqlTriggerOutput {
	return o
}

func (o SqlResourceSqlTriggerOutput) ToSqlResourceSqlTriggerOutputWithContext(ctx context.Context) SqlResourceSqlTriggerOutput {
	return o
}

func (o SqlResourceSqlTriggerOutput) ToOutput(ctx context.Context) pulumix.Output[*SqlResourceSqlTrigger] {
	return pulumix.Output[*SqlResourceSqlTrigger]{
		OutputState: o.OutputState,
	}
}

// The location of the resource group to which the resource belongs.
func (o SqlResourceSqlTriggerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlTrigger) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o SqlResourceSqlTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlTrigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlResourceSqlTriggerOutput) Resource() SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlTrigger) SqlTriggerGetPropertiesResponseResourcePtrOutput { return v.Resource }).(SqlTriggerGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o SqlResourceSqlTriggerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlResourceSqlTrigger) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o SqlResourceSqlTriggerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlTrigger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlResourceSqlTriggerOutput{})
}
