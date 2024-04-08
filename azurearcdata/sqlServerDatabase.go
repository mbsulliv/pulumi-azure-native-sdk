// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurearcdata

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Arc Sql Server database
// Azure REST API version: 2023-01-15-preview.
//
// Other available API versions: 2024-01-01.
type SqlServerDatabase struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of Arc Sql Server database
	Properties SqlServerDatabaseResourcePropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlServerDatabase registers a new resource with the given unique name, arguments, and options.
func NewSqlServerDatabase(ctx *pulumi.Context,
	name string, args *SqlServerDatabaseArgs, opts ...pulumi.ResourceOption) (*SqlServerDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SqlServerInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'SqlServerInstanceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurearcdata/v20220615preview:SqlServerDatabase"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20230115preview:SqlServerDatabase"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20240101:SqlServerDatabase"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SqlServerDatabase
	err := ctx.RegisterResource("azure-native:azurearcdata:SqlServerDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlServerDatabase gets an existing SqlServerDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlServerDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlServerDatabaseState, opts ...pulumi.ResourceOption) (*SqlServerDatabase, error) {
	var resource SqlServerDatabase
	err := ctx.ReadResource("azure-native:azurearcdata:SqlServerDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlServerDatabase resources.
type sqlServerDatabaseState struct {
}

type SqlServerDatabaseState struct {
}

func (SqlServerDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerDatabaseState)(nil)).Elem()
}

type sqlServerDatabaseArgs struct {
	// Name of the database
	DatabaseName *string `pulumi:"databaseName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Properties of Arc Sql Server database
	Properties SqlServerDatabaseResourceProperties `pulumi:"properties"`
	// The name of the Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of SQL Server Instance
	SqlServerInstanceName string `pulumi:"sqlServerInstanceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SqlServerDatabase resource.
type SqlServerDatabaseArgs struct {
	// Name of the database
	DatabaseName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Properties of Arc Sql Server database
	Properties SqlServerDatabaseResourcePropertiesInput
	// The name of the Azure resource group
	ResourceGroupName pulumi.StringInput
	// Name of SQL Server Instance
	SqlServerInstanceName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SqlServerDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerDatabaseArgs)(nil)).Elem()
}

type SqlServerDatabaseInput interface {
	pulumi.Input

	ToSqlServerDatabaseOutput() SqlServerDatabaseOutput
	ToSqlServerDatabaseOutputWithContext(ctx context.Context) SqlServerDatabaseOutput
}

func (*SqlServerDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerDatabase)(nil)).Elem()
}

func (i *SqlServerDatabase) ToSqlServerDatabaseOutput() SqlServerDatabaseOutput {
	return i.ToSqlServerDatabaseOutputWithContext(context.Background())
}

func (i *SqlServerDatabase) ToSqlServerDatabaseOutputWithContext(ctx context.Context) SqlServerDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerDatabaseOutput)
}

type SqlServerDatabaseOutput struct{ *pulumi.OutputState }

func (SqlServerDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerDatabase)(nil)).Elem()
}

func (o SqlServerDatabaseOutput) ToSqlServerDatabaseOutput() SqlServerDatabaseOutput {
	return o
}

func (o SqlServerDatabaseOutput) ToSqlServerDatabaseOutputWithContext(ctx context.Context) SqlServerDatabaseOutput {
	return o
}

// The geo-location where the resource lives
func (o SqlServerDatabaseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServerDatabase) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SqlServerDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServerDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of Arc Sql Server database
func (o SqlServerDatabaseOutput) Properties() SqlServerDatabaseResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *SqlServerDatabase) SqlServerDatabaseResourcePropertiesResponseOutput { return v.Properties }).(SqlServerDatabaseResourcePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SqlServerDatabaseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlServerDatabase) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SqlServerDatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlServerDatabase) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SqlServerDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServerDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlServerDatabaseOutput{})
}
