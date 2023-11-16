// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Monitor struct {
	pulumi.CustomResourceState

	Identity IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	Location pulumi.StringOutput                 `pulumi:"location"`
	// Name of the monitor resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties specific to the monitor resource.
	Properties MonitorPropertiesResponseOutput `pulumi:"properties"`
	Sku        ResourceSkuResponsePtrOutput    `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput   `pulumi:"tags"`
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
			Type: pulumi.String("azure-native:datadog:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:datadog/v20200201preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:datadog/v20210301:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:datadog/v20220601:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:datadog/v20230101:Monitor"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Monitor
	err := ctx.RegisterResource("azure-native:datadog/v20220801:Monitor", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:datadog/v20220801:Monitor", name, id, state, &resource, opts...)
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
	Identity *IdentityProperties `pulumi:"identity"`
	Location *string             `pulumi:"location"`
	// Monitor resource name
	MonitorName *string `pulumi:"monitorName"`
	// Properties specific to the monitor resource.
	Properties *MonitorProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               *ResourceSku      `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Monitor resource.
type MonitorArgs struct {
	Identity IdentityPropertiesPtrInput
	Location pulumi.StringPtrInput
	// Monitor resource name
	MonitorName pulumi.StringPtrInput
	// Properties specific to the monitor resource.
	Properties MonitorPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	Sku               ResourceSkuPtrInput
	Tags              pulumi.StringMapInput
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

func (o MonitorOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Monitor) IdentityPropertiesResponsePtrOutput { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o MonitorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the monitor resource.
func (o MonitorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties specific to the monitor resource.
func (o MonitorOutput) Properties() MonitorPropertiesResponseOutput {
	return o.ApplyT(func(v *Monitor) MonitorPropertiesResponseOutput { return v.Properties }).(MonitorPropertiesResponseOutput)
}

func (o MonitorOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v *Monitor) ResourceSkuResponsePtrOutput { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o MonitorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Monitor) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

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
