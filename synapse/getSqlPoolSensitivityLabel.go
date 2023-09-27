// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the sensitivity label of a given column
// Azure REST API version: 2021-06-01.
func LookupSqlPoolSensitivityLabel(ctx *pulumi.Context, args *LookupSqlPoolSensitivityLabelArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolSensitivityLabelResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlPoolSensitivityLabelResult
	err := ctx.Invoke("azure-native:synapse:getSqlPoolSensitivityLabel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolSensitivityLabelArgs struct {
	// The name of the column.
	ColumnName string `pulumi:"columnName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the schema.
	SchemaName string `pulumi:"schemaName"`
	// The source of the sensitivity label.
	SensitivityLabelSource string `pulumi:"sensitivityLabelSource"`
	// SQL pool name
	SqlPoolName string `pulumi:"sqlPoolName"`
	// The name of the table.
	TableName string `pulumi:"tableName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// A sensitivity label.
type LookupSqlPoolSensitivityLabelResult struct {
	// The column name.
	ColumnName string `pulumi:"columnName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
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
	// managed by
	ManagedBy string `pulumi:"managedBy"`
	// The name of the resource
	Name string  `pulumi:"name"`
	Rank *string `pulumi:"rank"`
	// The schema name.
	SchemaName string `pulumi:"schemaName"`
	// The table name.
	TableName string `pulumi:"tableName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSqlPoolSensitivityLabelOutput(ctx *pulumi.Context, args LookupSqlPoolSensitivityLabelOutputArgs, opts ...pulumi.InvokeOption) LookupSqlPoolSensitivityLabelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlPoolSensitivityLabelResult, error) {
			args := v.(LookupSqlPoolSensitivityLabelArgs)
			r, err := LookupSqlPoolSensitivityLabel(ctx, &args, opts...)
			var s LookupSqlPoolSensitivityLabelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlPoolSensitivityLabelResultOutput)
}

type LookupSqlPoolSensitivityLabelOutputArgs struct {
	// The name of the column.
	ColumnName pulumi.StringInput `pulumi:"columnName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the schema.
	SchemaName pulumi.StringInput `pulumi:"schemaName"`
	// The source of the sensitivity label.
	SensitivityLabelSource pulumi.StringInput `pulumi:"sensitivityLabelSource"`
	// SQL pool name
	SqlPoolName pulumi.StringInput `pulumi:"sqlPoolName"`
	// The name of the table.
	TableName pulumi.StringInput `pulumi:"tableName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSqlPoolSensitivityLabelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolSensitivityLabelArgs)(nil)).Elem()
}

// A sensitivity label.
type LookupSqlPoolSensitivityLabelResultOutput struct{ *pulumi.OutputState }

func (LookupSqlPoolSensitivityLabelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolSensitivityLabelResult)(nil)).Elem()
}

func (o LookupSqlPoolSensitivityLabelResultOutput) ToLookupSqlPoolSensitivityLabelResultOutput() LookupSqlPoolSensitivityLabelResultOutput {
	return o
}

func (o LookupSqlPoolSensitivityLabelResultOutput) ToLookupSqlPoolSensitivityLabelResultOutputWithContext(ctx context.Context) LookupSqlPoolSensitivityLabelResultOutput {
	return o
}

func (o LookupSqlPoolSensitivityLabelResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSqlPoolSensitivityLabelResult] {
	return pulumix.Output[LookupSqlPoolSensitivityLabelResult]{
		OutputState: o.OutputState,
	}
}

// The column name.
func (o LookupSqlPoolSensitivityLabelResultOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.ColumnName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSqlPoolSensitivityLabelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.Id }).(pulumi.StringOutput)
}

// The information type.
func (o LookupSqlPoolSensitivityLabelResultOutput) InformationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) *string { return v.InformationType }).(pulumi.StringPtrOutput)
}

// The information type ID.
func (o LookupSqlPoolSensitivityLabelResultOutput) InformationTypeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) *string { return v.InformationTypeId }).(pulumi.StringPtrOutput)
}

// Is sensitivity recommendation disabled. Applicable for recommended sensitivity label only. Specifies whether the sensitivity recommendation on this column is disabled (dismissed) or not.
func (o LookupSqlPoolSensitivityLabelResultOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) bool { return v.IsDisabled }).(pulumi.BoolOutput)
}

// The label ID.
func (o LookupSqlPoolSensitivityLabelResultOutput) LabelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) *string { return v.LabelId }).(pulumi.StringPtrOutput)
}

// The label name.
func (o LookupSqlPoolSensitivityLabelResultOutput) LabelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) *string { return v.LabelName }).(pulumi.StringPtrOutput)
}

// managed by
func (o LookupSqlPoolSensitivityLabelResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSqlPoolSensitivityLabelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlPoolSensitivityLabelResultOutput) Rank() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) *string { return v.Rank }).(pulumi.StringPtrOutput)
}

// The schema name.
func (o LookupSqlPoolSensitivityLabelResultOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.SchemaName }).(pulumi.StringOutput)
}

// The table name.
func (o LookupSqlPoolSensitivityLabelResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.TableName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSqlPoolSensitivityLabelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlPoolSensitivityLabelResultOutput{})
}
