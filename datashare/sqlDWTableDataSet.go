// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A SQL DW table data set.
// Azure REST API version: 2021-08-01. Prior API version in Azure Native 1.x: 2020-09-01
type SqlDWTableDataSet struct {
	pulumi.CustomResourceState

	// Unique id for identifying a data set resource
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// DataWarehouse name of the source data set
	DataWarehouseName pulumi.StringOutput `pulumi:"dataWarehouseName"`
	// Kind of data set.
	// Expected value is 'SqlDWTable'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Schema of the table. Default value is dbo.
	SchemaName pulumi.StringOutput `pulumi:"schemaName"`
	// Resource id of SQL server
	SqlServerResourceId pulumi.StringOutput `pulumi:"sqlServerResourceId"`
	// System Data of the Azure resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// SQL DW table name.
	TableName pulumi.StringOutput `pulumi:"tableName"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlDWTableDataSet registers a new resource with the given unique name, arguments, and options.
func NewSqlDWTableDataSet(ctx *pulumi.Context,
	name string, args *SqlDWTableDataSetArgs, opts ...pulumi.ResourceOption) (*SqlDWTableDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DataWarehouseName == nil {
		return nil, errors.New("invalid value for required argument 'DataWarehouseName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaName == nil {
		return nil, errors.New("invalid value for required argument 'SchemaName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.SqlServerResourceId == nil {
		return nil, errors.New("invalid value for required argument 'SqlServerResourceId'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	args.Kind = pulumi.String("SqlDWTable")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:SqlDWTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:SqlDWTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:SqlDWTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:SqlDWTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:SqlDWTableDataSet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SqlDWTableDataSet
	err := ctx.RegisterResource("azure-native:datashare:SqlDWTableDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlDWTableDataSet gets an existing SqlDWTableDataSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlDWTableDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlDWTableDataSetState, opts ...pulumi.ResourceOption) (*SqlDWTableDataSet, error) {
	var resource SqlDWTableDataSet
	err := ctx.ReadResource("azure-native:datashare:SqlDWTableDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlDWTableDataSet resources.
type sqlDWTableDataSetState struct {
}

type SqlDWTableDataSetState struct {
}

func (SqlDWTableDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDWTableDataSetState)(nil)).Elem()
}

type sqlDWTableDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName *string `pulumi:"dataSetName"`
	// DataWarehouse name of the source data set
	DataWarehouseName string `pulumi:"dataWarehouseName"`
	// Kind of data set.
	// Expected value is 'SqlDWTable'.
	Kind string `pulumi:"kind"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Schema of the table. Default value is dbo.
	SchemaName string `pulumi:"schemaName"`
	// The name of the share to add the data set to.
	ShareName string `pulumi:"shareName"`
	// Resource id of SQL server
	SqlServerResourceId string `pulumi:"sqlServerResourceId"`
	// SQL DW table name.
	TableName string `pulumi:"tableName"`
}

// The set of arguments for constructing a SqlDWTableDataSet resource.
type SqlDWTableDataSetArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// The name of the dataSet.
	DataSetName pulumi.StringPtrInput
	// DataWarehouse name of the source data set
	DataWarehouseName pulumi.StringInput
	// Kind of data set.
	// Expected value is 'SqlDWTable'.
	Kind pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Schema of the table. Default value is dbo.
	SchemaName pulumi.StringInput
	// The name of the share to add the data set to.
	ShareName pulumi.StringInput
	// Resource id of SQL server
	SqlServerResourceId pulumi.StringInput
	// SQL DW table name.
	TableName pulumi.StringInput
}

func (SqlDWTableDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDWTableDataSetArgs)(nil)).Elem()
}

type SqlDWTableDataSetInput interface {
	pulumi.Input

	ToSqlDWTableDataSetOutput() SqlDWTableDataSetOutput
	ToSqlDWTableDataSetOutputWithContext(ctx context.Context) SqlDWTableDataSetOutput
}

func (*SqlDWTableDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDWTableDataSet)(nil)).Elem()
}

func (i *SqlDWTableDataSet) ToSqlDWTableDataSetOutput() SqlDWTableDataSetOutput {
	return i.ToSqlDWTableDataSetOutputWithContext(context.Background())
}

func (i *SqlDWTableDataSet) ToSqlDWTableDataSetOutputWithContext(ctx context.Context) SqlDWTableDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDWTableDataSetOutput)
}

func (i *SqlDWTableDataSet) ToOutput(ctx context.Context) pulumix.Output[*SqlDWTableDataSet] {
	return pulumix.Output[*SqlDWTableDataSet]{
		OutputState: i.ToSqlDWTableDataSetOutputWithContext(ctx).OutputState,
	}
}

type SqlDWTableDataSetOutput struct{ *pulumi.OutputState }

func (SqlDWTableDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDWTableDataSet)(nil)).Elem()
}

func (o SqlDWTableDataSetOutput) ToSqlDWTableDataSetOutput() SqlDWTableDataSetOutput {
	return o
}

func (o SqlDWTableDataSetOutput) ToSqlDWTableDataSetOutputWithContext(ctx context.Context) SqlDWTableDataSetOutput {
	return o
}

func (o SqlDWTableDataSetOutput) ToOutput(ctx context.Context) pulumix.Output[*SqlDWTableDataSet] {
	return pulumix.Output[*SqlDWTableDataSet]{
		OutputState: o.OutputState,
	}
}

// Unique id for identifying a data set resource
func (o SqlDWTableDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

// DataWarehouse name of the source data set
func (o SqlDWTableDataSetOutput) DataWarehouseName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.DataWarehouseName }).(pulumi.StringOutput)
}

// Kind of data set.
// Expected value is 'SqlDWTable'.
func (o SqlDWTableDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o SqlDWTableDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Schema of the table. Default value is dbo.
func (o SqlDWTableDataSetOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.SchemaName }).(pulumi.StringOutput)
}

// Resource id of SQL server
func (o SqlDWTableDataSetOutput) SqlServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.SqlServerResourceId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o SqlDWTableDataSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// SQL DW table name.
func (o SqlDWTableDataSetOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

// Type of the azure resource
func (o SqlDWTableDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlDWTableDataSetOutput{})
}
