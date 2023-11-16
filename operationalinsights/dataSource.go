// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasources under OMS Workspace.
// Azure REST API version: 2020-08-01. Prior API version in Azure Native 1.x: 2020-08-01.
//
// Other available API versions: 2015-11-01-preview.
type DataSource struct {
	pulumi.CustomResourceState

	// The ETag of the data source.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the DataSource.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The data source properties in raw json format, each kind of data source have it's own schema.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataSource registers a new resource with the given unique name, arguments, and options.
func NewDataSource(ctx *pulumi.Context,
	name string, args *DataSourceArgs, opts ...pulumi.ResourceOption) (*DataSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights/v20151101preview:DataSource"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:DataSource"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:DataSource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DataSource
	err := ctx.RegisterResource("azure-native:operationalinsights:DataSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSource gets an existing DataSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSourceState, opts ...pulumi.ResourceOption) (*DataSource, error) {
	var resource DataSource
	err := ctx.ReadResource("azure-native:operationalinsights:DataSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSource resources.
type dataSourceState struct {
}

type DataSourceState struct {
}

func (DataSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceState)(nil)).Elem()
}

type dataSourceArgs struct {
	// The name of the datasource resource.
	DataSourceName *string `pulumi:"dataSourceName"`
	// The kind of the DataSource.
	Kind string `pulumi:"kind"`
	// The data source properties in raw json format, each kind of data source have it's own schema.
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a DataSource resource.
type DataSourceArgs struct {
	// The name of the datasource resource.
	DataSourceName pulumi.StringPtrInput
	// The kind of the DataSource.
	Kind pulumi.StringInput
	// The data source properties in raw json format, each kind of data source have it's own schema.
	Properties pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (DataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceArgs)(nil)).Elem()
}

type DataSourceInput interface {
	pulumi.Input

	ToDataSourceOutput() DataSourceOutput
	ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput
}

func (*DataSource) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil)).Elem()
}

func (i *DataSource) ToDataSourceOutput() DataSourceOutput {
	return i.ToDataSourceOutputWithContext(context.Background())
}

func (i *DataSource) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceOutput)
}

type DataSourceOutput struct{ *pulumi.OutputState }

func (DataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil)).Elem()
}

func (o DataSourceOutput) ToDataSourceOutput() DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return o
}

// The ETag of the data source.
func (o DataSourceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the DataSource.
func (o DataSourceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o DataSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The data source properties in raw json format, each kind of data source have it's own schema.
func (o DataSourceOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *DataSource) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Resource tags.
func (o DataSourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataSourceOutput{})
}
