// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Monitoring Setting resource
type MonitoringSetting struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the Monitoring Setting resource
	Properties MonitoringSettingPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMonitoringSetting registers a new resource with the given unique name, arguments, and options.
func NewMonitoringSetting(ctx *pulumi.Context,
	name string, args *MonitoringSettingArgs, opts ...pulumi.ResourceOption) (*MonitoringSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210601preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:MonitoringSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource MonitoringSetting
	err := ctx.RegisterResource("azure-native:appplatform/v20200701:MonitoringSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitoringSetting gets an existing MonitoringSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitoringSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitoringSettingState, opts ...pulumi.ResourceOption) (*MonitoringSetting, error) {
	var resource MonitoringSetting
	err := ctx.ReadResource("azure-native:appplatform/v20200701:MonitoringSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MonitoringSetting resources.
type monitoringSettingState struct {
}

type MonitoringSettingState struct {
}

func (MonitoringSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringSettingState)(nil)).Elem()
}

type monitoringSettingArgs struct {
	// Properties of the Monitoring Setting resource
	Properties *MonitoringSettingProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a MonitoringSetting resource.
type MonitoringSettingArgs struct {
	// Properties of the Monitoring Setting resource
	Properties MonitoringSettingPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (MonitoringSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringSettingArgs)(nil)).Elem()
}

type MonitoringSettingInput interface {
	pulumi.Input

	ToMonitoringSettingOutput() MonitoringSettingOutput
	ToMonitoringSettingOutputWithContext(ctx context.Context) MonitoringSettingOutput
}

func (*MonitoringSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSetting)(nil)).Elem()
}

func (i *MonitoringSetting) ToMonitoringSettingOutput() MonitoringSettingOutput {
	return i.ToMonitoringSettingOutputWithContext(context.Background())
}

func (i *MonitoringSetting) ToMonitoringSettingOutputWithContext(ctx context.Context) MonitoringSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringSettingOutput)
}

type MonitoringSettingOutput struct{ *pulumi.OutputState }

func (MonitoringSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSetting)(nil)).Elem()
}

func (o MonitoringSettingOutput) ToMonitoringSettingOutput() MonitoringSettingOutput {
	return o
}

func (o MonitoringSettingOutput) ToMonitoringSettingOutputWithContext(ctx context.Context) MonitoringSettingOutput {
	return o
}

// The name of the resource.
func (o MonitoringSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the Monitoring Setting resource
func (o MonitoringSettingOutput) Properties() MonitoringSettingPropertiesResponseOutput {
	return o.ApplyT(func(v *MonitoringSetting) MonitoringSettingPropertiesResponseOutput { return v.Properties }).(MonitoringSettingPropertiesResponseOutput)
}

// The type of the resource.
func (o MonitoringSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MonitoringSettingOutput{})
}
