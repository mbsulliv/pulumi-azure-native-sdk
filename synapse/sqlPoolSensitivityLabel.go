// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A sensitivity label.
// Azure REST API version: 2021-06-01. Prior API version in Azure Native 1.x: 2021-03-01
type SqlPoolSensitivityLabel struct {
	pulumi.CustomResourceState

	// The column name.
	ColumnName pulumi.StringOutput `pulumi:"columnName"`
	// The information type.
	InformationType pulumi.StringPtrOutput `pulumi:"informationType"`
	// The information type ID.
	InformationTypeId pulumi.StringPtrOutput `pulumi:"informationTypeId"`
	// Is sensitivity recommendation disabled. Applicable for recommended sensitivity label only. Specifies whether the sensitivity recommendation on this column is disabled (dismissed) or not.
	IsDisabled pulumi.BoolOutput `pulumi:"isDisabled"`
	// The label ID.
	LabelId pulumi.StringPtrOutput `pulumi:"labelId"`
	// The label name.
	LabelName pulumi.StringPtrOutput `pulumi:"labelName"`
	// managed by
	ManagedBy pulumi.StringOutput `pulumi:"managedBy"`
	// The name of the resource
	Name pulumi.StringOutput    `pulumi:"name"`
	Rank pulumi.StringPtrOutput `pulumi:"rank"`
	// The schema name.
	SchemaName pulumi.StringOutput `pulumi:"schemaName"`
	// The table name.
	TableName pulumi.StringOutput `pulumi:"tableName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlPoolSensitivityLabel registers a new resource with the given unique name, arguments, and options.
func NewSqlPoolSensitivityLabel(ctx *pulumi.Context,
	name string, args *SqlPoolSensitivityLabelArgs, opts ...pulumi.ResourceOption) (*SqlPoolSensitivityLabel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ColumnName == nil {
		return nil, errors.New("invalid value for required argument 'ColumnName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaName == nil {
		return nil, errors.New("invalid value for required argument 'SchemaName'")
	}
	if args.SqlPoolName == nil {
		return nil, errors.New("invalid value for required argument 'SqlPoolName'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:SqlPoolSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:SqlPoolSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:SqlPoolSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:SqlPoolSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:SqlPoolSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:SqlPoolSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:SqlPoolSensitivityLabel"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SqlPoolSensitivityLabel
	err := ctx.RegisterResource("azure-native:synapse:SqlPoolSensitivityLabel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlPoolSensitivityLabel gets an existing SqlPoolSensitivityLabel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlPoolSensitivityLabel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlPoolSensitivityLabelState, opts ...pulumi.ResourceOption) (*SqlPoolSensitivityLabel, error) {
	var resource SqlPoolSensitivityLabel
	err := ctx.ReadResource("azure-native:synapse:SqlPoolSensitivityLabel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlPoolSensitivityLabel resources.
type sqlPoolSensitivityLabelState struct {
}

type SqlPoolSensitivityLabelState struct {
}

func (SqlPoolSensitivityLabelState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolSensitivityLabelState)(nil)).Elem()
}

type sqlPoolSensitivityLabelArgs struct {
	// The name of the column.
	ColumnName string `pulumi:"columnName"`
	// The information type.
	InformationType *string `pulumi:"informationType"`
	// The information type ID.
	InformationTypeId *string `pulumi:"informationTypeId"`
	// The label ID.
	LabelId *string `pulumi:"labelId"`
	// The label name.
	LabelName *string               `pulumi:"labelName"`
	Rank      *SensitivityLabelRank `pulumi:"rank"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the schema.
	SchemaName string `pulumi:"schemaName"`
	// The source of the sensitivity label.
	SensitivityLabelSource *string `pulumi:"sensitivityLabelSource"`
	// SQL pool name
	SqlPoolName string `pulumi:"sqlPoolName"`
	// The name of the table.
	TableName string `pulumi:"tableName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a SqlPoolSensitivityLabel resource.
type SqlPoolSensitivityLabelArgs struct {
	// The name of the column.
	ColumnName pulumi.StringInput
	// The information type.
	InformationType pulumi.StringPtrInput
	// The information type ID.
	InformationTypeId pulumi.StringPtrInput
	// The label ID.
	LabelId pulumi.StringPtrInput
	// The label name.
	LabelName pulumi.StringPtrInput
	Rank      SensitivityLabelRankPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the schema.
	SchemaName pulumi.StringInput
	// The source of the sensitivity label.
	SensitivityLabelSource pulumi.StringPtrInput
	// SQL pool name
	SqlPoolName pulumi.StringInput
	// The name of the table.
	TableName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (SqlPoolSensitivityLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolSensitivityLabelArgs)(nil)).Elem()
}

type SqlPoolSensitivityLabelInput interface {
	pulumi.Input

	ToSqlPoolSensitivityLabelOutput() SqlPoolSensitivityLabelOutput
	ToSqlPoolSensitivityLabelOutputWithContext(ctx context.Context) SqlPoolSensitivityLabelOutput
}

func (*SqlPoolSensitivityLabel) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlPoolSensitivityLabel)(nil)).Elem()
}

func (i *SqlPoolSensitivityLabel) ToSqlPoolSensitivityLabelOutput() SqlPoolSensitivityLabelOutput {
	return i.ToSqlPoolSensitivityLabelOutputWithContext(context.Background())
}

func (i *SqlPoolSensitivityLabel) ToSqlPoolSensitivityLabelOutputWithContext(ctx context.Context) SqlPoolSensitivityLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolSensitivityLabelOutput)
}

func (i *SqlPoolSensitivityLabel) ToOutput(ctx context.Context) pulumix.Output[*SqlPoolSensitivityLabel] {
	return pulumix.Output[*SqlPoolSensitivityLabel]{
		OutputState: i.ToSqlPoolSensitivityLabelOutputWithContext(ctx).OutputState,
	}
}

type SqlPoolSensitivityLabelOutput struct{ *pulumi.OutputState }

func (SqlPoolSensitivityLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlPoolSensitivityLabel)(nil)).Elem()
}

func (o SqlPoolSensitivityLabelOutput) ToSqlPoolSensitivityLabelOutput() SqlPoolSensitivityLabelOutput {
	return o
}

func (o SqlPoolSensitivityLabelOutput) ToSqlPoolSensitivityLabelOutputWithContext(ctx context.Context) SqlPoolSensitivityLabelOutput {
	return o
}

func (o SqlPoolSensitivityLabelOutput) ToOutput(ctx context.Context) pulumix.Output[*SqlPoolSensitivityLabel] {
	return pulumix.Output[*SqlPoolSensitivityLabel]{
		OutputState: o.OutputState,
	}
}

// The column name.
func (o SqlPoolSensitivityLabelOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPoolSensitivityLabel) pulumi.StringOutput { return v.ColumnName }).(pulumi.StringOutput)
}

// The information type.
func (o SqlPoolSensitivityLabelOutput) InformationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPoolSensitivityLabel) pulumi.StringPtrOutput { return v.InformationType }).(pulumi.StringPtrOutput)
}

// The information type ID.
func (o SqlPoolSensitivityLabelOutput) InformationTypeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPoolSensitivityLabel) pulumi.StringPtrOutput { return v.InformationTypeId }).(pulumi.StringPtrOutput)
}

// Is sensitivity recommendation disabled. Applicable for recommended sensitivity label only. Specifies whether the sensitivity recommendation on this column is disabled (dismissed) or not.
func (o SqlPoolSensitivityLabelOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *SqlPoolSensitivityLabel) pulumi.BoolOutput { return v.IsDisabled }).(pulumi.BoolOutput)
}

// The label ID.
func (o SqlPoolSensitivityLabelOutput) LabelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPoolSensitivityLabel) pulumi.StringPtrOutput { return v.LabelId }).(pulumi.StringPtrOutput)
}

// The label name.
func (o SqlPoolSensitivityLabelOutput) LabelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPoolSensitivityLabel) pulumi.StringPtrOutput { return v.LabelName }).(pulumi.StringPtrOutput)
}

// managed by
func (o SqlPoolSensitivityLabelOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPoolSensitivityLabel) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

// The name of the resource
func (o SqlPoolSensitivityLabelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPoolSensitivityLabel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlPoolSensitivityLabelOutput) Rank() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPoolSensitivityLabel) pulumi.StringPtrOutput { return v.Rank }).(pulumi.StringPtrOutput)
}

// The schema name.
func (o SqlPoolSensitivityLabelOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPoolSensitivityLabel) pulumi.StringOutput { return v.SchemaName }).(pulumi.StringOutput)
}

// The table name.
func (o SqlPoolSensitivityLabelOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPoolSensitivityLabel) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SqlPoolSensitivityLabelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPoolSensitivityLabel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlPoolSensitivityLabelOutput{})
}
