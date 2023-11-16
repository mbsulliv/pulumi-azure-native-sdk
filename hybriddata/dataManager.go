// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybriddata

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The DataManager resource.
// Azure REST API version: 2019-06-01. Prior API version in Azure Native 1.x: 2019-06-01.
type DataManager struct {
	pulumi.CustomResourceState

	// Etag of the Resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East
	// US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo
	// region is specified on update the request will succeed.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The sku type.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource
	// (across resource groups).
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataManager registers a new resource with the given unique name, arguments, and options.
func NewDataManager(ctx *pulumi.Context,
	name string, args *DataManagerArgs, opts ...pulumi.ResourceOption) (*DataManager, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybriddata/v20160601:DataManager"),
		},
		{
			Type: pulumi.String("azure-native:hybriddata/v20190601:DataManager"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DataManager
	err := ctx.RegisterResource("azure-native:hybriddata:DataManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataManager gets an existing DataManager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataManagerState, opts ...pulumi.ResourceOption) (*DataManager, error) {
	var resource DataManager
	err := ctx.ReadResource("azure-native:hybriddata:DataManager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataManager resources.
type dataManagerState struct {
}

type DataManagerState struct {
}

func (DataManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataManagerState)(nil)).Elem()
}

type dataManagerArgs struct {
	// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	DataManagerName *string `pulumi:"dataManagerName"`
	// The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East
	// US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo
	// region is specified on update the request will succeed.
	Location *string `pulumi:"location"`
	// The Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku type.
	Sku *Sku `pulumi:"sku"`
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource
	// (across resource groups).
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DataManager resource.
type DataManagerArgs struct {
	// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	DataManagerName pulumi.StringPtrInput
	// The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East
	// US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo
	// region is specified on update the request will succeed.
	Location pulumi.StringPtrInput
	// The Resource Group Name
	ResourceGroupName pulumi.StringInput
	// The sku type.
	Sku SkuPtrInput
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource
	// (across resource groups).
	Tags pulumi.StringMapInput
}

func (DataManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataManagerArgs)(nil)).Elem()
}

type DataManagerInput interface {
	pulumi.Input

	ToDataManagerOutput() DataManagerOutput
	ToDataManagerOutputWithContext(ctx context.Context) DataManagerOutput
}

func (*DataManager) ElementType() reflect.Type {
	return reflect.TypeOf((**DataManager)(nil)).Elem()
}

func (i *DataManager) ToDataManagerOutput() DataManagerOutput {
	return i.ToDataManagerOutputWithContext(context.Background())
}

func (i *DataManager) ToDataManagerOutputWithContext(ctx context.Context) DataManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataManagerOutput)
}

type DataManagerOutput struct{ *pulumi.OutputState }

func (DataManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataManager)(nil)).Elem()
}

func (o DataManagerOutput) ToDataManagerOutput() DataManagerOutput {
	return o
}

func (o DataManagerOutput) ToDataManagerOutputWithContext(ctx context.Context) DataManagerOutput {
	return o
}

// Etag of the Resource.
func (o DataManagerOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataManager) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East
// US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo
// region is specified on update the request will succeed.
func (o DataManagerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataManager) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The Resource Name.
func (o DataManagerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataManager) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The sku type.
func (o DataManagerOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *DataManager) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource
// (across resource groups).
func (o DataManagerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataManager) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The Resource type.
func (o DataManagerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataManager) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataManagerOutput{})
}
