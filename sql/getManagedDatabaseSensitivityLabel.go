// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the sensitivity label of a given column
// Azure REST API version: 2021-11-01.
func LookupManagedDatabaseSensitivityLabel(ctx *pulumi.Context, args *LookupManagedDatabaseSensitivityLabelArgs, opts ...pulumi.InvokeOption) (*LookupManagedDatabaseSensitivityLabelResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedDatabaseSensitivityLabelResult
	err := ctx.Invoke("azure-native:sql:getManagedDatabaseSensitivityLabel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedDatabaseSensitivityLabelArgs struct {
	// The name of the column.
	ColumnName string `pulumi:"columnName"`
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the schema.
	SchemaName string `pulumi:"schemaName"`
	// The source of the sensitivity label.
	SensitivityLabelSource string `pulumi:"sensitivityLabelSource"`
	// The name of the table.
	TableName string `pulumi:"tableName"`
}

// A sensitivity label.
type LookupManagedDatabaseSensitivityLabelResult struct {
	// The column name.
	ColumnName string `pulumi:"columnName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The information type.
	InformationType *string `pulumi:"informationType"`
	// The information type ID.
	InformationTypeId *string `pulumi:"informationTypeId"`
	// Is sensitivity recommendation disabled. Applicable for recommended sensitivity label only. Specifies whether the sensitivity recommendation on this column is disabled (dismissed) or not.
	IsDisabled bool `pulumi:"isDisabled"`
	// The label ID.
	LabelId *string `pulumi:"labelId"`
	// The label name.
	LabelName *string `pulumi:"labelName"`
	// Resource that manages the sensitivity label.
	ManagedBy string `pulumi:"managedBy"`
	// Resource name.
	Name string  `pulumi:"name"`
	Rank *string `pulumi:"rank"`
	// The schema name.
	SchemaName string `pulumi:"schemaName"`
	// The table name.
	TableName string `pulumi:"tableName"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupManagedDatabaseSensitivityLabelOutput(ctx *pulumi.Context, args LookupManagedDatabaseSensitivityLabelOutputArgs, opts ...pulumi.InvokeOption) LookupManagedDatabaseSensitivityLabelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedDatabaseSensitivityLabelResult, error) {
			args := v.(LookupManagedDatabaseSensitivityLabelArgs)
			r, err := LookupManagedDatabaseSensitivityLabel(ctx, &args, opts...)
			var s LookupManagedDatabaseSensitivityLabelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedDatabaseSensitivityLabelResultOutput)
}

type LookupManagedDatabaseSensitivityLabelOutputArgs struct {
	// The name of the column.
	ColumnName pulumi.StringInput `pulumi:"columnName"`
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the schema.
	SchemaName pulumi.StringInput `pulumi:"schemaName"`
	// The source of the sensitivity label.
	SensitivityLabelSource pulumi.StringInput `pulumi:"sensitivityLabelSource"`
	// The name of the table.
	TableName pulumi.StringInput `pulumi:"tableName"`
}

func (LookupManagedDatabaseSensitivityLabelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedDatabaseSensitivityLabelArgs)(nil)).Elem()
}

// A sensitivity label.
type LookupManagedDatabaseSensitivityLabelResultOutput struct{ *pulumi.OutputState }

func (LookupManagedDatabaseSensitivityLabelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedDatabaseSensitivityLabelResult)(nil)).Elem()
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) ToLookupManagedDatabaseSensitivityLabelResultOutput() LookupManagedDatabaseSensitivityLabelResultOutput {
	return o
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) ToLookupManagedDatabaseSensitivityLabelResultOutputWithContext(ctx context.Context) LookupManagedDatabaseSensitivityLabelResultOutput {
	return o
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupManagedDatabaseSensitivityLabelResult] {
	return pulumix.Output[LookupManagedDatabaseSensitivityLabelResult]{
		OutputState: o.OutputState,
	}
}

// The column name.
func (o LookupManagedDatabaseSensitivityLabelResultOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.ColumnName }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupManagedDatabaseSensitivityLabelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.Id }).(pulumi.StringOutput)
}

// The information type.
func (o LookupManagedDatabaseSensitivityLabelResultOutput) InformationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) *string { return v.InformationType }).(pulumi.StringPtrOutput)
}

// The information type ID.
func (o LookupManagedDatabaseSensitivityLabelResultOutput) InformationTypeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) *string { return v.InformationTypeId }).(pulumi.StringPtrOutput)
}

// Is sensitivity recommendation disabled. Applicable for recommended sensitivity label only. Specifies whether the sensitivity recommendation on this column is disabled (dismissed) or not.
func (o LookupManagedDatabaseSensitivityLabelResultOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) bool { return v.IsDisabled }).(pulumi.BoolOutput)
}

// The label ID.
func (o LookupManagedDatabaseSensitivityLabelResultOutput) LabelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) *string { return v.LabelId }).(pulumi.StringPtrOutput)
}

// The label name.
func (o LookupManagedDatabaseSensitivityLabelResultOutput) LabelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) *string { return v.LabelName }).(pulumi.StringPtrOutput)
}

// Resource that manages the sensitivity label.
func (o LookupManagedDatabaseSensitivityLabelResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupManagedDatabaseSensitivityLabelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) Rank() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) *string { return v.Rank }).(pulumi.StringPtrOutput)
}

// The schema name.
func (o LookupManagedDatabaseSensitivityLabelResultOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.SchemaName }).(pulumi.StringOutput)
}

// The table name.
func (o LookupManagedDatabaseSensitivityLabelResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.TableName }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupManagedDatabaseSensitivityLabelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedDatabaseSensitivityLabelResultOutput{})
}
