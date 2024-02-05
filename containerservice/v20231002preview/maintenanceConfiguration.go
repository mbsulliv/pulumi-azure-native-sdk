// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231002preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// See [planned maintenance](https://docs.microsoft.com/azure/aks/planned-maintenance) for more information about planned maintenance.
type MaintenanceConfiguration struct {
	pulumi.CustomResourceState

	// Maintenance window for the maintenance configuration.
	MaintenanceWindow MaintenanceWindowResponsePtrOutput `pulumi:"maintenanceWindow"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Time slots on which upgrade is not allowed.
	NotAllowedTime TimeSpanResponseArrayOutput `pulumi:"notAllowedTime"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// If two array entries specify the same day of the week, the applied configuration is the union of times in both entries.
	TimeInWeek TimeInWeekResponseArrayOutput `pulumi:"timeInWeek"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMaintenanceConfiguration registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceConfiguration(ctx *pulumi.Context,
	name string, args *MaintenanceConfigurationArgs, opts ...pulumi.ResourceOption) (*MaintenanceConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.MaintenanceWindow != nil {
		args.MaintenanceWindow = args.MaintenanceWindow.ToMaintenanceWindowPtrOutput().ApplyT(func(v *MaintenanceWindow) *MaintenanceWindow { return v.Defaults() }).(MaintenanceWindowPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201201:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210201:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210301:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210501:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210701:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211001:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211101preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220101:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220102preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220201:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220202preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220301:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220302preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220401:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220601:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220701:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220802preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220803preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220901:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221002preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221101:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221102preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230101:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230102preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230201:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230202preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230301:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230302preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230401:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230402preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230501:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230502preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230601:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230602preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230701:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230702preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230801:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230802preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230901:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230902preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20231001:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20231101:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20231102preview:MaintenanceConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MaintenanceConfiguration
	err := ctx.RegisterResource("azure-native:containerservice/v20231002preview:MaintenanceConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceConfiguration gets an existing MaintenanceConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceConfigurationState, opts ...pulumi.ResourceOption) (*MaintenanceConfiguration, error) {
	var resource MaintenanceConfiguration
	err := ctx.ReadResource("azure-native:containerservice/v20231002preview:MaintenanceConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceConfiguration resources.
type maintenanceConfigurationState struct {
}

type MaintenanceConfigurationState struct {
}

func (MaintenanceConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceConfigurationState)(nil)).Elem()
}

type maintenanceConfigurationArgs struct {
	// The name of the maintenance configuration.
	ConfigName *string `pulumi:"configName"`
	// Maintenance window for the maintenance configuration.
	MaintenanceWindow *MaintenanceWindow `pulumi:"maintenanceWindow"`
	// Time slots on which upgrade is not allowed.
	NotAllowedTime []TimeSpan `pulumi:"notAllowedTime"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
	// If two array entries specify the same day of the week, the applied configuration is the union of times in both entries.
	TimeInWeek []TimeInWeek `pulumi:"timeInWeek"`
}

// The set of arguments for constructing a MaintenanceConfiguration resource.
type MaintenanceConfigurationArgs struct {
	// The name of the maintenance configuration.
	ConfigName pulumi.StringPtrInput
	// Maintenance window for the maintenance configuration.
	MaintenanceWindow MaintenanceWindowPtrInput
	// Time slots on which upgrade is not allowed.
	NotAllowedTime TimeSpanArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput
	// If two array entries specify the same day of the week, the applied configuration is the union of times in both entries.
	TimeInWeek TimeInWeekArrayInput
}

func (MaintenanceConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceConfigurationArgs)(nil)).Elem()
}

type MaintenanceConfigurationInput interface {
	pulumi.Input

	ToMaintenanceConfigurationOutput() MaintenanceConfigurationOutput
	ToMaintenanceConfigurationOutputWithContext(ctx context.Context) MaintenanceConfigurationOutput
}

func (*MaintenanceConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceConfiguration)(nil)).Elem()
}

func (i *MaintenanceConfiguration) ToMaintenanceConfigurationOutput() MaintenanceConfigurationOutput {
	return i.ToMaintenanceConfigurationOutputWithContext(context.Background())
}

func (i *MaintenanceConfiguration) ToMaintenanceConfigurationOutputWithContext(ctx context.Context) MaintenanceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceConfigurationOutput)
}

type MaintenanceConfigurationOutput struct{ *pulumi.OutputState }

func (MaintenanceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceConfiguration)(nil)).Elem()
}

func (o MaintenanceConfigurationOutput) ToMaintenanceConfigurationOutput() MaintenanceConfigurationOutput {
	return o
}

func (o MaintenanceConfigurationOutput) ToMaintenanceConfigurationOutputWithContext(ctx context.Context) MaintenanceConfigurationOutput {
	return o
}

// Maintenance window for the maintenance configuration.
func (o MaintenanceConfigurationOutput) MaintenanceWindow() MaintenanceWindowResponsePtrOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) MaintenanceWindowResponsePtrOutput { return v.MaintenanceWindow }).(MaintenanceWindowResponsePtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o MaintenanceConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Time slots on which upgrade is not allowed.
func (o MaintenanceConfigurationOutput) NotAllowedTime() TimeSpanResponseArrayOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) TimeSpanResponseArrayOutput { return v.NotAllowedTime }).(TimeSpanResponseArrayOutput)
}

// The system metadata relating to this resource.
func (o MaintenanceConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// If two array entries specify the same day of the week, the applied configuration is the union of times in both entries.
func (o MaintenanceConfigurationOutput) TimeInWeek() TimeInWeekResponseArrayOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) TimeInWeekResponseArrayOutput { return v.TimeInWeek }).(TimeInWeekResponseArrayOutput)
}

// Resource type
func (o MaintenanceConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MaintenanceConfigurationOutput{})
}
