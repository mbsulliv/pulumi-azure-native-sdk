// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MetricsConfiguration struct {
	pulumi.CustomResourceState

	// The interval in minutes by which metrics will be collected.
	CollectionInterval pulumi.Float64Output `pulumi:"collectionInterval"`
	// The more detailed status of the metrics configuration.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage pulumi.StringOutput `pulumi:"detailedStatusMessage"`
	// The list of metrics that are available for the cluster but disabled at the moment.
	DisabledMetrics pulumi.StringArrayOutput `pulumi:"disabledMetrics"`
	// The list of metric names that have been chosen to be enabled in addition to the core set of enabled metrics.
	EnabledMetrics pulumi.StringArrayOutput `pulumi:"enabledMetrics"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the metrics configuration.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMetricsConfiguration registers a new resource with the given unique name, arguments, and options.
func NewMetricsConfiguration(ctx *pulumi.Context,
	name string, args *MetricsConfigurationArgs, opts ...pulumi.ResourceOption) (*MetricsConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.CollectionInterval == nil {
		return nil, errors.New("invalid value for required argument 'CollectionInterval'")
	}
	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkcloud:MetricsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20221212preview:MetricsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20230501preview:MetricsConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MetricsConfiguration
	err := ctx.RegisterResource("azure-native:networkcloud/v20230701:MetricsConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricsConfiguration gets an existing MetricsConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricsConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricsConfigurationState, opts ...pulumi.ResourceOption) (*MetricsConfiguration, error) {
	var resource MetricsConfiguration
	err := ctx.ReadResource("azure-native:networkcloud/v20230701:MetricsConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricsConfiguration resources.
type metricsConfigurationState struct {
}

type MetricsConfigurationState struct {
}

func (MetricsConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricsConfigurationState)(nil)).Elem()
}

type metricsConfigurationArgs struct {
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The interval in minutes by which metrics will be collected.
	CollectionInterval float64 `pulumi:"collectionInterval"`
	// The list of metric names that have been chosen to be enabled in addition to the core set of enabled metrics.
	EnabledMetrics []string `pulumi:"enabledMetrics"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the metrics configuration for the cluster.
	MetricsConfigurationName *string `pulumi:"metricsConfigurationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MetricsConfiguration resource.
type MetricsConfigurationArgs struct {
	// The name of the cluster.
	ClusterName pulumi.StringInput
	// The interval in minutes by which metrics will be collected.
	CollectionInterval pulumi.Float64Input
	// The list of metric names that have been chosen to be enabled in addition to the core set of enabled metrics.
	EnabledMetrics pulumi.StringArrayInput
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the metrics configuration for the cluster.
	MetricsConfigurationName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (MetricsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricsConfigurationArgs)(nil)).Elem()
}

type MetricsConfigurationInput interface {
	pulumi.Input

	ToMetricsConfigurationOutput() MetricsConfigurationOutput
	ToMetricsConfigurationOutputWithContext(ctx context.Context) MetricsConfigurationOutput
}

func (*MetricsConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricsConfiguration)(nil)).Elem()
}

func (i *MetricsConfiguration) ToMetricsConfigurationOutput() MetricsConfigurationOutput {
	return i.ToMetricsConfigurationOutputWithContext(context.Background())
}

func (i *MetricsConfiguration) ToMetricsConfigurationOutputWithContext(ctx context.Context) MetricsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricsConfigurationOutput)
}

type MetricsConfigurationOutput struct{ *pulumi.OutputState }

func (MetricsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricsConfiguration)(nil)).Elem()
}

func (o MetricsConfigurationOutput) ToMetricsConfigurationOutput() MetricsConfigurationOutput {
	return o
}

func (o MetricsConfigurationOutput) ToMetricsConfigurationOutputWithContext(ctx context.Context) MetricsConfigurationOutput {
	return o
}

// The interval in minutes by which metrics will be collected.
func (o MetricsConfigurationOutput) CollectionInterval() pulumi.Float64Output {
	return o.ApplyT(func(v *MetricsConfiguration) pulumi.Float64Output { return v.CollectionInterval }).(pulumi.Float64Output)
}

// The more detailed status of the metrics configuration.
func (o MetricsConfigurationOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsConfiguration) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o MetricsConfigurationOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsConfiguration) pulumi.StringOutput { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The list of metrics that are available for the cluster but disabled at the moment.
func (o MetricsConfigurationOutput) DisabledMetrics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MetricsConfiguration) pulumi.StringArrayOutput { return v.DisabledMetrics }).(pulumi.StringArrayOutput)
}

// The list of metric names that have been chosen to be enabled in addition to the core set of enabled metrics.
func (o MetricsConfigurationOutput) EnabledMetrics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MetricsConfiguration) pulumi.StringArrayOutput { return v.EnabledMetrics }).(pulumi.StringArrayOutput)
}

// The extended location of the cluster associated with the resource.
func (o MetricsConfigurationOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *MetricsConfiguration) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The geo-location where the resource lives
func (o MetricsConfigurationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsConfiguration) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o MetricsConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the metrics configuration.
func (o MetricsConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MetricsConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MetricsConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o MetricsConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MetricsConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MetricsConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MetricsConfigurationOutput{})
}
