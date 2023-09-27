// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A sensitivity label.
type ManagedDatabaseSensitivityLabel struct {
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

// NewManagedDatabaseSensitivityLabel registers a new resource with the given unique name, arguments, and options.
func NewManagedDatabaseSensitivityLabel(ctx *pulumi.Context,
	name string, args *ManagedDatabaseSensitivityLabelArgs, opts ...pulumi.ResourceOption) (*ManagedDatabaseSensitivityLabel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ColumnName == nil {
		return nil, errors.New("invalid value for required argument 'ColumnName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaName == nil {
		return nil, errors.New("invalid value for required argument 'SchemaName'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:ManagedDatabaseSensitivityLabel"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedDatabaseSensitivityLabel
	err := ctx.RegisterResource("azure-native:sql/v20221101preview:ManagedDatabaseSensitivityLabel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedDatabaseSensitivityLabel gets an existing ManagedDatabaseSensitivityLabel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedDatabaseSensitivityLabel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedDatabaseSensitivityLabelState, opts ...pulumi.ResourceOption) (*ManagedDatabaseSensitivityLabel, error) {
	var resource ManagedDatabaseSensitivityLabel
	err := ctx.ReadResource("azure-native:sql/v20221101preview:ManagedDatabaseSensitivityLabel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedDatabaseSensitivityLabel resources.
type managedDatabaseSensitivityLabelState struct {
}

type ManagedDatabaseSensitivityLabelState struct {
}

func (ManagedDatabaseSensitivityLabelState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedDatabaseSensitivityLabelState)(nil)).Elem()
}

type managedDatabaseSensitivityLabelArgs struct {
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
	LabelName *string `pulumi:"labelName"`
	// The name of the managed instance.
	ManagedInstanceName string                `pulumi:"managedInstanceName"`
	Rank                *SensitivityLabelRank `pulumi:"rank"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the schema.
	SchemaName string `pulumi:"schemaName"`
	// The source of the sensitivity label.
	SensitivityLabelSource *string `pulumi:"sensitivityLabelSource"`
	// The name of the table.
	TableName string `pulumi:"tableName"`
}

// The set of arguments for constructing a ManagedDatabaseSensitivityLabel resource.
type ManagedDatabaseSensitivityLabelArgs struct {
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
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput
	Rank                SensitivityLabelRankPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the schema.
	SchemaName pulumi.StringInput
	// The source of the sensitivity label.
	SensitivityLabelSource pulumi.StringPtrInput
	// The name of the table.
	TableName pulumi.StringInput
}

func (ManagedDatabaseSensitivityLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedDatabaseSensitivityLabelArgs)(nil)).Elem()
}

type ManagedDatabaseSensitivityLabelInput interface {
	pulumi.Input

	ToManagedDatabaseSensitivityLabelOutput() ManagedDatabaseSensitivityLabelOutput
	ToManagedDatabaseSensitivityLabelOutputWithContext(ctx context.Context) ManagedDatabaseSensitivityLabelOutput
}

func (*ManagedDatabaseSensitivityLabel) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDatabaseSensitivityLabel)(nil)).Elem()
}

func (i *ManagedDatabaseSensitivityLabel) ToManagedDatabaseSensitivityLabelOutput() ManagedDatabaseSensitivityLabelOutput {
	return i.ToManagedDatabaseSensitivityLabelOutputWithContext(context.Background())
}

func (i *ManagedDatabaseSensitivityLabel) ToManagedDatabaseSensitivityLabelOutputWithContext(ctx context.Context) ManagedDatabaseSensitivityLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDatabaseSensitivityLabelOutput)
}

func (i *ManagedDatabaseSensitivityLabel) ToOutput(ctx context.Context) pulumix.Output[*ManagedDatabaseSensitivityLabel] {
	return pulumix.Output[*ManagedDatabaseSensitivityLabel]{
		OutputState: i.ToManagedDatabaseSensitivityLabelOutputWithContext(ctx).OutputState,
	}
}

type ManagedDatabaseSensitivityLabelOutput struct{ *pulumi.OutputState }

func (ManagedDatabaseSensitivityLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDatabaseSensitivityLabel)(nil)).Elem()
}

func (o ManagedDatabaseSensitivityLabelOutput) ToManagedDatabaseSensitivityLabelOutput() ManagedDatabaseSensitivityLabelOutput {
	return o
}

func (o ManagedDatabaseSensitivityLabelOutput) ToManagedDatabaseSensitivityLabelOutputWithContext(ctx context.Context) ManagedDatabaseSensitivityLabelOutput {
	return o
}

func (o ManagedDatabaseSensitivityLabelOutput) ToOutput(ctx context.Context) pulumix.Output[*ManagedDatabaseSensitivityLabel] {
	return pulumix.Output[*ManagedDatabaseSensitivityLabel]{
		OutputState: o.OutputState,
	}
}

// The column name.
func (o ManagedDatabaseSensitivityLabelOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringOutput { return v.ColumnName }).(pulumi.StringOutput)
}

// The information type.
func (o ManagedDatabaseSensitivityLabelOutput) InformationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringPtrOutput { return v.InformationType }).(pulumi.StringPtrOutput)
}

// The information type ID.
func (o ManagedDatabaseSensitivityLabelOutput) InformationTypeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringPtrOutput { return v.InformationTypeId }).(pulumi.StringPtrOutput)
}

// Is sensitivity recommendation disabled. Applicable for recommended sensitivity label only. Specifies whether the sensitivity recommendation on this column is disabled (dismissed) or not.
func (o ManagedDatabaseSensitivityLabelOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.BoolOutput { return v.IsDisabled }).(pulumi.BoolOutput)
}

// The label ID.
func (o ManagedDatabaseSensitivityLabelOutput) LabelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringPtrOutput { return v.LabelId }).(pulumi.StringPtrOutput)
}

// The label name.
func (o ManagedDatabaseSensitivityLabelOutput) LabelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringPtrOutput { return v.LabelName }).(pulumi.StringPtrOutput)
}

// Resource that manages the sensitivity label.
func (o ManagedDatabaseSensitivityLabelOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

// Resource name.
func (o ManagedDatabaseSensitivityLabelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedDatabaseSensitivityLabelOutput) Rank() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringPtrOutput { return v.Rank }).(pulumi.StringPtrOutput)
}

// The schema name.
func (o ManagedDatabaseSensitivityLabelOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringOutput { return v.SchemaName }).(pulumi.StringOutput)
}

// The table name.
func (o ManagedDatabaseSensitivityLabelOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

// Resource type.
func (o ManagedDatabaseSensitivityLabelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedDatabaseSensitivityLabelOutput{})
}
