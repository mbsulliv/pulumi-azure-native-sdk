// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a DataSetMapping in a shareSubscription
func LookupSqlDWTableDataSetMapping(ctx *pulumi.Context, args *LookupSqlDWTableDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupSqlDWTableDataSetMappingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlDWTableDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getSqlDWTableDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlDWTableDataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSetMapping.
	DataSetMappingName string `pulumi:"dataSetMappingName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}

// A SQL DW Table data set mapping.
type LookupSqlDWTableDataSetMappingResult struct {
	// The id of the source data set.
	DataSetId string `pulumi:"dataSetId"`
	// Gets the status of the data set mapping.
	DataSetMappingStatus string `pulumi:"dataSetMappingStatus"`
	// DataWarehouse name of the source data set
	DataWarehouseName string `pulumi:"dataWarehouseName"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set mapping.
	// Expected value is 'SqlDWTable'.
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Provisioning state of the data set mapping.
	ProvisioningState string `pulumi:"provisioningState"`
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

func LookupSqlDWTableDataSetMappingOutput(ctx *pulumi.Context, args LookupSqlDWTableDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupSqlDWTableDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlDWTableDataSetMappingResult, error) {
			args := v.(LookupSqlDWTableDataSetMappingArgs)
			r, err := LookupSqlDWTableDataSetMapping(ctx, &args, opts...)
			var s LookupSqlDWTableDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlDWTableDataSetMappingResultOutput)
}

type LookupSqlDWTableDataSetMappingOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the dataSetMapping.
	DataSetMappingName pulumi.StringInput `pulumi:"dataSetMappingName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupSqlDWTableDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDWTableDataSetMappingArgs)(nil)).Elem()
}

// A SQL DW Table data set mapping.
type LookupSqlDWTableDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupSqlDWTableDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDWTableDataSetMappingResult)(nil)).Elem()
}

func (o LookupSqlDWTableDataSetMappingResultOutput) ToLookupSqlDWTableDataSetMappingResultOutput() LookupSqlDWTableDataSetMappingResultOutput {
	return o
}

func (o LookupSqlDWTableDataSetMappingResultOutput) ToLookupSqlDWTableDataSetMappingResultOutputWithContext(ctx context.Context) LookupSqlDWTableDataSetMappingResultOutput {
	return o
}

// The id of the source data set.
func (o LookupSqlDWTableDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

// Gets the status of the data set mapping.
func (o LookupSqlDWTableDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

// DataWarehouse name of the source data set
func (o LookupSqlDWTableDataSetMappingResultOutput) DataWarehouseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.DataWarehouseName }).(pulumi.StringOutput)
}

// The resource id of the azure resource
func (o LookupSqlDWTableDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of data set mapping.
// Expected value is 'SqlDWTable'.
func (o LookupSqlDWTableDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o LookupSqlDWTableDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the data set mapping.
func (o LookupSqlDWTableDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Schema of the table. Default value is dbo.
func (o LookupSqlDWTableDataSetMappingResultOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.SchemaName }).(pulumi.StringOutput)
}

// Resource id of SQL server
func (o LookupSqlDWTableDataSetMappingResultOutput) SqlServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.SqlServerResourceId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o LookupSqlDWTableDataSetMappingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// SQL DW table name.
func (o LookupSqlDWTableDataSetMappingResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.TableName }).(pulumi.StringOutput)
}

// Type of the azure resource
func (o LookupSqlDWTableDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlDWTableDataSetMappingResultOutput{})
}
