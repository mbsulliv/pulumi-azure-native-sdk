// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a DataSet in a share
// Azure REST API version: 2021-08-01.
func LookupSqlDWTableDataSet(ctx *pulumi.Context, args *LookupSqlDWTableDataSetArgs, opts ...pulumi.InvokeOption) (*LookupSqlDWTableDataSetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlDWTableDataSetResult
	err := ctx.Invoke("azure-native:datashare:getSqlDWTableDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlDWTableDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName string `pulumi:"dataSetName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName string `pulumi:"shareName"`
}

// A SQL DW table data set.
type LookupSqlDWTableDataSetResult struct {
	// Unique id for identifying a data set resource
	DataSetId string `pulumi:"dataSetId"`
	// DataWarehouse name of the source data set
	DataWarehouseName string `pulumi:"dataWarehouseName"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set.
	// Expected value is 'SqlDWTable'.
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Schema of the table. Default value is dbo.
	SchemaName string `pulumi:"schemaName"`
	// Resource id of SQL server
	SqlServerResourceId string `pulumi:"sqlServerResourceId"`
	// System Data of the Azure resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// SQL DW table name.
	TableName string `pulumi:"tableName"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}

func LookupSqlDWTableDataSetOutput(ctx *pulumi.Context, args LookupSqlDWTableDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupSqlDWTableDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlDWTableDataSetResult, error) {
			args := v.(LookupSqlDWTableDataSetArgs)
			r, err := LookupSqlDWTableDataSet(ctx, &args, opts...)
			var s LookupSqlDWTableDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlDWTableDataSetResultOutput)
}

type LookupSqlDWTableDataSetOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName pulumi.StringInput `pulumi:"dataSetName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName pulumi.StringInput `pulumi:"shareName"`
}

func (LookupSqlDWTableDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDWTableDataSetArgs)(nil)).Elem()
}

// A SQL DW table data set.
type LookupSqlDWTableDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupSqlDWTableDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDWTableDataSetResult)(nil)).Elem()
}

func (o LookupSqlDWTableDataSetResultOutput) ToLookupSqlDWTableDataSetResultOutput() LookupSqlDWTableDataSetResultOutput {
	return o
}

func (o LookupSqlDWTableDataSetResultOutput) ToLookupSqlDWTableDataSetResultOutputWithContext(ctx context.Context) LookupSqlDWTableDataSetResultOutput {
	return o
}

// Unique id for identifying a data set resource
func (o LookupSqlDWTableDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

// DataWarehouse name of the source data set
func (o LookupSqlDWTableDataSetResultOutput) DataWarehouseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.DataWarehouseName }).(pulumi.StringOutput)
}

// The resource id of the azure resource
func (o LookupSqlDWTableDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of data set.
// Expected value is 'SqlDWTable'.
func (o LookupSqlDWTableDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o LookupSqlDWTableDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// Schema of the table. Default value is dbo.
func (o LookupSqlDWTableDataSetResultOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.SchemaName }).(pulumi.StringOutput)
}

// Resource id of SQL server
func (o LookupSqlDWTableDataSetResultOutput) SqlServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.SqlServerResourceId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o LookupSqlDWTableDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// SQL DW table name.
func (o LookupSqlDWTableDataSetResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.TableName }).(pulumi.StringOutput)
}

// Type of the azure resource
func (o LookupSqlDWTableDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlDWTableDataSetResultOutput{})
}
