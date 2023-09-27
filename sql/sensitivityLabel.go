// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A sensitivity label.
// Azure REST API version: 2021-11-01. Prior API version in Azure Native 1.x: 2020-11-01-preview
type SensitivityLabel struct {
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
	// Resource that manages the sensitivity label.
	ManagedBy pulumi.StringOutput `pulumi:"managedBy"`
	// Resource name.
	Name pulumi.StringOutput    `pulumi:"name"`
	Rank pulumi.StringPtrOutput `pulumi:"rank"`
	// The schema name.
	SchemaName pulumi.StringOutput `pulumi:"schemaName"`
	// The table name.
	TableName pulumi.StringOutput `pulumi:"tableName"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSensitivityLabel registers a new resource with the given unique name, arguments, and options.
func NewSensitivityLabel(ctx *pulumi.Context,
	name string, args *SensitivityLabelArgs, opts ...pulumi.ResourceOption) (*SensitivityLabel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ColumnName == nil {
		return nil, errors.New("invalid value for required argument 'ColumnName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaName == nil {
		return nil, errors.New("invalid value for required argument 'SchemaName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:SensitivityLabel"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SensitivityLabel
	err := ctx.RegisterResource("azure-native:sql:SensitivityLabel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSensitivityLabel gets an existing SensitivityLabel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSensitivityLabel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SensitivityLabelState, opts ...pulumi.ResourceOption) (*SensitivityLabel, error) {
	var resource SensitivityLabel
	err := ctx.ReadResource("azure-native:sql:SensitivityLabel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SensitivityLabel resources.
type sensitivityLabelState struct {
}

type SensitivityLabelState struct {
}

func (SensitivityLabelState) ElementType() reflect.Type {
	return reflect.TypeOf((*sensitivityLabelState)(nil)).Elem()
}

type sensitivityLabelArgs struct {
	// The name of the column.
	ColumnName string `pulumi:"columnName"`
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The information type.
	InformationType *string `pulumi:"informationType"`
	// The information type ID.
	InformationTypeId *string `pulumi:"informationTypeId"`
	// The label ID.
	LabelId *string `pulumi:"labelId"`
	// The label name.
	LabelName *string               `pulumi:"labelName"`
	Rank      *SensitivityLabelRank `pulumi:"rank"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the schema.
	SchemaName string `pulumi:"schemaName"`
	// The source of the sensitivity label.
	SensitivityLabelSource *string `pulumi:"sensitivityLabelSource"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the table.
	TableName string `pulumi:"tableName"`
}

// The set of arguments for constructing a SensitivityLabel resource.
type SensitivityLabelArgs struct {
	// The name of the column.
	ColumnName pulumi.StringInput
	// The name of the database.
	DatabaseName pulumi.StringInput
	// The information type.
	InformationType pulumi.StringPtrInput
	// The information type ID.
	InformationTypeId pulumi.StringPtrInput
	// The label ID.
	LabelId pulumi.StringPtrInput
	// The label name.
	LabelName pulumi.StringPtrInput
	Rank      SensitivityLabelRankPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the schema.
	SchemaName pulumi.StringInput
	// The source of the sensitivity label.
	SensitivityLabelSource pulumi.StringPtrInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The name of the table.
	TableName pulumi.StringInput
}

func (SensitivityLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sensitivityLabelArgs)(nil)).Elem()
}

type SensitivityLabelInput interface {
	pulumi.Input

	ToSensitivityLabelOutput() SensitivityLabelOutput
	ToSensitivityLabelOutputWithContext(ctx context.Context) SensitivityLabelOutput
}

func (*SensitivityLabel) ElementType() reflect.Type {
	return reflect.TypeOf((**SensitivityLabel)(nil)).Elem()
}

func (i *SensitivityLabel) ToSensitivityLabelOutput() SensitivityLabelOutput {
	return i.ToSensitivityLabelOutputWithContext(context.Background())
}

func (i *SensitivityLabel) ToSensitivityLabelOutputWithContext(ctx context.Context) SensitivityLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SensitivityLabelOutput)
}

func (i *SensitivityLabel) ToOutput(ctx context.Context) pulumix.Output[*SensitivityLabel] {
	return pulumix.Output[*SensitivityLabel]{
		OutputState: i.ToSensitivityLabelOutputWithContext(ctx).OutputState,
	}
}

type SensitivityLabelOutput struct{ *pulumi.OutputState }

func (SensitivityLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SensitivityLabel)(nil)).Elem()
}

func (o SensitivityLabelOutput) ToSensitivityLabelOutput() SensitivityLabelOutput {
	return o
}

func (o SensitivityLabelOutput) ToSensitivityLabelOutputWithContext(ctx context.Context) SensitivityLabelOutput {
	return o
}

func (o SensitivityLabelOutput) ToOutput(ctx context.Context) pulumix.Output[*SensitivityLabel] {
	return pulumix.Output[*SensitivityLabel]{
		OutputState: o.OutputState,
	}
}

// The column name.
func (o SensitivityLabelOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringOutput { return v.ColumnName }).(pulumi.StringOutput)
}

// The information type.
func (o SensitivityLabelOutput) InformationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringPtrOutput { return v.InformationType }).(pulumi.StringPtrOutput)
}

// The information type ID.
func (o SensitivityLabelOutput) InformationTypeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringPtrOutput { return v.InformationTypeId }).(pulumi.StringPtrOutput)
}

// Is sensitivity recommendation disabled. Applicable for recommended sensitivity label only. Specifies whether the sensitivity recommendation on this column is disabled (dismissed) or not.
func (o SensitivityLabelOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.BoolOutput { return v.IsDisabled }).(pulumi.BoolOutput)
}

// The label ID.
func (o SensitivityLabelOutput) LabelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringPtrOutput { return v.LabelId }).(pulumi.StringPtrOutput)
}

// The label name.
func (o SensitivityLabelOutput) LabelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringPtrOutput { return v.LabelName }).(pulumi.StringPtrOutput)
}

// Resource that manages the sensitivity label.
func (o SensitivityLabelOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

// Resource name.
func (o SensitivityLabelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SensitivityLabelOutput) Rank() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringPtrOutput { return v.Rank }).(pulumi.StringPtrOutput)
}

// The schema name.
func (o SensitivityLabelOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringOutput { return v.SchemaName }).(pulumi.StringOutput)
}

// The table name.
func (o SensitivityLabelOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

// Resource type.
func (o SensitivityLabelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SensitivityLabelOutput{})
}
