// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230915

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB Table.
type TableResourceTable struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                         `pulumi:"name"`
	Options  TableGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource TableGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTableResourceTable registers a new resource with the given unique name, arguments, and options.
func NewTableResourceTable(ctx *pulumi.Context,
	name string, args *TableResourceTableArgs, opts ...pulumi.ResourceOption) (*TableResourceTable, error) {
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
			Type: pulumi.String("azure-native:documentdb:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230301preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230415:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:TableResourceTable"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TableResourceTable
	err := ctx.RegisterResource("azure-native:documentdb/v20230915:TableResourceTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTableResourceTable gets an existing TableResourceTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTableResourceTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableResourceTableState, opts ...pulumi.ResourceOption) (*TableResourceTable, error) {
	var resource TableResourceTable
	err := ctx.ReadResource("azure-native:documentdb/v20230915:TableResourceTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TableResourceTable resources.
type tableResourceTableState struct {
}

type TableResourceTableState struct {
}

func (TableResourceTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableResourceTableState)(nil)).Elem()
}

type tableResourceTableArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options *CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a Table
	Resource TableResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cosmos DB table name.
	TableName *string `pulumi:"tableName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a TableResourceTable resource.
type TableResourceTableArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsPtrInput
	// The standard JSON format of a Table
	Resource TableResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Cosmos DB table name.
	TableName pulumi.StringPtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (TableResourceTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableResourceTableArgs)(nil)).Elem()
}

type TableResourceTableInput interface {
	pulumi.Input

	ToTableResourceTableOutput() TableResourceTableOutput
	ToTableResourceTableOutputWithContext(ctx context.Context) TableResourceTableOutput
}

func (*TableResourceTable) ElementType() reflect.Type {
	return reflect.TypeOf((**TableResourceTable)(nil)).Elem()
}

func (i *TableResourceTable) ToTableResourceTableOutput() TableResourceTableOutput {
	return i.ToTableResourceTableOutputWithContext(context.Background())
}

func (i *TableResourceTable) ToTableResourceTableOutputWithContext(ctx context.Context) TableResourceTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableResourceTableOutput)
}

type TableResourceTableOutput struct{ *pulumi.OutputState }

func (TableResourceTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableResourceTable)(nil)).Elem()
}

func (o TableResourceTableOutput) ToTableResourceTableOutput() TableResourceTableOutput {
	return o
}

func (o TableResourceTableOutput) ToTableResourceTableOutputWithContext(ctx context.Context) TableResourceTableOutput {
	return o
}

// The location of the resource group to which the resource belongs.
func (o TableResourceTableOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TableResourceTable) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o TableResourceTableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TableResourceTable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TableResourceTableOutput) Options() TableGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v *TableResourceTable) TableGetPropertiesResponseOptionsPtrOutput { return v.Options }).(TableGetPropertiesResponseOptionsPtrOutput)
}

func (o TableResourceTableOutput) Resource() TableGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *TableResourceTable) TableGetPropertiesResponseResourcePtrOutput { return v.Resource }).(TableGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o TableResourceTableOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TableResourceTable) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o TableResourceTableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TableResourceTable) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TableResourceTableOutput{})
}
