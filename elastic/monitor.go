// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elastic

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Monitor resource.
// Azure REST API version: 2023-06-01. Prior API version in Azure Native 1.x: 2020-07-01.
//
// Other available API versions: 2023-06-15-preview, 2023-07-01-preview, 2023-10-01-preview.
type Monitor struct {
	pulumi.CustomResourceState

	// Identity properties of the monitor resource.
	Identity IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	// The location of the monitor resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the monitor resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the monitor resource.
	Properties MonitorPropertiesResponseOutput `pulumi:"properties"`
	// SKU of the monitor resource.
	Sku ResourceSkuResponsePtrOutput `pulumi:"sku"`
	// The system metadata relating to this resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tags of the monitor resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the monitor resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMonitor registers a new resource with the given unique name, arguments, and options.
func NewMonitor(ctx *pulumi.Context,
	name string, args *MonitorArgs, opts ...pulumi.ResourceOption) (*Monitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:elastic/v20200701:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20200701preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20210901preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20211001preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20220505preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20220701preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20220901preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20230201preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20230501preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20230601:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20230615preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20230701preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20231001preview:Monitor"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Monitor
	err := ctx.RegisterResource("azure-native:elastic:Monitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitor gets an existing Monitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitorState, opts ...pulumi.ResourceOption) (*Monitor, error) {
	var resource Monitor
	err := ctx.ReadResource("azure-native:elastic:Monitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Monitor resources.
type monitorState struct {
}

type MonitorState struct {
}

func (MonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorState)(nil)).Elem()
}

type monitorArgs struct {
	// Identity properties of the monitor resource.
	Identity *IdentityProperties `pulumi:"identity"`
	// The location of the monitor resource
	Location *string `pulumi:"location"`
	// Monitor resource name
	MonitorName *string `pulumi:"monitorName"`
	// Properties of the monitor resource.
	Properties *MonitorProperties `pulumi:"properties"`
	// The name of the resource group to which the Elastic resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU of the monitor resource.
	Sku *ResourceSku `pulumi:"sku"`
	// The tags of the monitor resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Monitor resource.
type MonitorArgs struct {
	// Identity properties of the monitor resource.
	Identity IdentityPropertiesPtrInput
	// The location of the monitor resource
	Location pulumi.StringPtrInput
	// Monitor resource name
	MonitorName pulumi.StringPtrInput
	// Properties of the monitor resource.
	Properties MonitorPropertiesPtrInput
	// The name of the resource group to which the Elastic resource belongs.
	ResourceGroupName pulumi.StringInput
	// SKU of the monitor resource.
	Sku ResourceSkuPtrInput
	// The tags of the monitor resource.
	Tags pulumi.StringMapInput
}

func (MonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorArgs)(nil)).Elem()
}

type MonitorInput interface {
	pulumi.Input

	ToMonitorOutput() MonitorOutput
	ToMonitorOutputWithContext(ctx context.Context) MonitorOutput
}

func (*Monitor) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (i *Monitor) ToMonitorOutput() MonitorOutput {
	return i.ToMonitorOutputWithContext(context.Background())
}

func (i *Monitor) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorOutput)
}

type MonitorOutput struct{ *pulumi.OutputState }

func (MonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (o MonitorOutput) ToMonitorOutput() MonitorOutput {
	return o
}

func (o MonitorOutput) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return o
}

// Identity properties of the monitor resource.
func (o MonitorOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Monitor) IdentityPropertiesResponsePtrOutput { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

// The location of the monitor resource
func (o MonitorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the monitor resource.
func (o MonitorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the monitor resource.
func (o MonitorOutput) Properties() MonitorPropertiesResponseOutput {
	return o.ApplyT(func(v *Monitor) MonitorPropertiesResponseOutput { return v.Properties }).(MonitorPropertiesResponseOutput)
}

// SKU of the monitor resource.
func (o MonitorOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v *Monitor) ResourceSkuResponsePtrOutput { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// The system metadata relating to this resource
func (o MonitorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Monitor) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tags of the monitor resource.
func (o MonitorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the monitor resource.
func (o MonitorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MonitorOutput{})
}
