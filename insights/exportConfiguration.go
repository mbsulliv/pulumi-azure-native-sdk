// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Properties that define a Continuous Export configuration.
// Azure REST API version: 2015-05-01. Prior API version in Azure Native 1.x: 2015-05-01.
type ExportConfiguration struct {
	pulumi.CustomResourceState

	// The name of the Application Insights component.
	ApplicationName pulumi.StringOutput `pulumi:"applicationName"`
	// The name of the destination storage container.
	ContainerName pulumi.StringOutput `pulumi:"containerName"`
	// The name of destination account.
	DestinationAccountId pulumi.StringOutput `pulumi:"destinationAccountId"`
	// The destination account location ID.
	DestinationStorageLocationId pulumi.StringOutput `pulumi:"destinationStorageLocationId"`
	// The destination storage account subscription ID.
	DestinationStorageSubscriptionId pulumi.StringOutput `pulumi:"destinationStorageSubscriptionId"`
	// The destination type.
	DestinationType pulumi.StringOutput `pulumi:"destinationType"`
	// The unique ID of the export configuration inside an Application Insights component. It is auto generated when the Continuous Export configuration is created.
	ExportId pulumi.StringOutput `pulumi:"exportId"`
	// This indicates current Continuous Export configuration status. The possible values are 'Preparing', 'Success', 'Failure'.
	ExportStatus pulumi.StringOutput `pulumi:"exportStatus"`
	// The instrumentation key of the Application Insights component.
	InstrumentationKey pulumi.StringOutput `pulumi:"instrumentationKey"`
	// This will be 'true' if the Continuous Export configuration is enabled, otherwise it will be 'false'.
	IsUserEnabled pulumi.StringOutput `pulumi:"isUserEnabled"`
	// The last time the Continuous Export configuration started failing.
	LastGapTime pulumi.StringOutput `pulumi:"lastGapTime"`
	// The last time data was successfully delivered to the destination storage container for this Continuous Export configuration.
	LastSuccessTime pulumi.StringOutput `pulumi:"lastSuccessTime"`
	// Last time the Continuous Export configuration was updated.
	LastUserUpdate pulumi.StringOutput `pulumi:"lastUserUpdate"`
	// Deprecated
	NotificationQueueEnabled pulumi.StringPtrOutput `pulumi:"notificationQueueEnabled"`
	// This is the reason the Continuous Export configuration started failing. It can be 'AzureStorageNotFound' or 'AzureStorageAccessDenied'.
	PermanentErrorReason pulumi.StringOutput `pulumi:"permanentErrorReason"`
	// This comma separated list of document types that will be exported. The possible values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
	RecordTypes pulumi.StringPtrOutput `pulumi:"recordTypes"`
	// The resource group of the Application Insights component.
	ResourceGroup pulumi.StringOutput `pulumi:"resourceGroup"`
	// The name of the destination storage account.
	StorageName pulumi.StringOutput `pulumi:"storageName"`
	// The subscription of the Application Insights component.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
}

// NewExportConfiguration registers a new resource with the given unique name, arguments, and options.
func NewExportConfiguration(ctx *pulumi.Context,
	name string, args *ExportConfigurationArgs, opts ...pulumi.ResourceOption) (*ExportConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20150501:ExportConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ExportConfiguration
	err := ctx.RegisterResource("azure-native:insights:ExportConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExportConfiguration gets an existing ExportConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExportConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExportConfigurationState, opts ...pulumi.ResourceOption) (*ExportConfiguration, error) {
	var resource ExportConfiguration
	err := ctx.ReadResource("azure-native:insights:ExportConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExportConfiguration resources.
type exportConfigurationState struct {
}

type ExportConfigurationState struct {
}

func (ExportConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*exportConfigurationState)(nil)).Elem()
}

type exportConfigurationArgs struct {
	// The name of destination storage account.
	DestinationAccountId *string `pulumi:"destinationAccountId"`
	// The SAS URL for the destination storage container. It must grant write permission.
	DestinationAddress *string `pulumi:"destinationAddress"`
	// The location ID of the destination storage container.
	DestinationStorageLocationId *string `pulumi:"destinationStorageLocationId"`
	// The subscription ID of the destination storage container.
	DestinationStorageSubscriptionId *string `pulumi:"destinationStorageSubscriptionId"`
	// The Continuous Export destination type. This has to be 'Blob'.
	DestinationType *string `pulumi:"destinationType"`
	// The Continuous Export configuration ID. This is unique within a Application Insights component.
	ExportId *string `pulumi:"exportId"`
	// Set to 'true' to create a Continuous Export configuration as enabled, otherwise set it to 'false'.
	IsEnabled *string `pulumi:"isEnabled"`
	// Deprecated
	NotificationQueueEnabled *string `pulumi:"notificationQueueEnabled"`
	// Deprecated
	NotificationQueueUri *string `pulumi:"notificationQueueUri"`
	// The document types to be exported, as comma separated values. Allowed values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
	RecordTypes *string `pulumi:"recordTypes"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ExportConfiguration resource.
type ExportConfigurationArgs struct {
	// The name of destination storage account.
	DestinationAccountId pulumi.StringPtrInput
	// The SAS URL for the destination storage container. It must grant write permission.
	DestinationAddress pulumi.StringPtrInput
	// The location ID of the destination storage container.
	DestinationStorageLocationId pulumi.StringPtrInput
	// The subscription ID of the destination storage container.
	DestinationStorageSubscriptionId pulumi.StringPtrInput
	// The Continuous Export destination type. This has to be 'Blob'.
	DestinationType pulumi.StringPtrInput
	// The Continuous Export configuration ID. This is unique within a Application Insights component.
	ExportId pulumi.StringPtrInput
	// Set to 'true' to create a Continuous Export configuration as enabled, otherwise set it to 'false'.
	IsEnabled pulumi.StringPtrInput
	// Deprecated
	NotificationQueueEnabled pulumi.StringPtrInput
	// Deprecated
	NotificationQueueUri pulumi.StringPtrInput
	// The document types to be exported, as comma separated values. Allowed values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
	RecordTypes pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringInput
}

func (ExportConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*exportConfigurationArgs)(nil)).Elem()
}

type ExportConfigurationInput interface {
	pulumi.Input

	ToExportConfigurationOutput() ExportConfigurationOutput
	ToExportConfigurationOutputWithContext(ctx context.Context) ExportConfigurationOutput
}

func (*ExportConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportConfiguration)(nil)).Elem()
}

func (i *ExportConfiguration) ToExportConfigurationOutput() ExportConfigurationOutput {
	return i.ToExportConfigurationOutputWithContext(context.Background())
}

func (i *ExportConfiguration) ToExportConfigurationOutputWithContext(ctx context.Context) ExportConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportConfigurationOutput)
}

type ExportConfigurationOutput struct{ *pulumi.OutputState }

func (ExportConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportConfiguration)(nil)).Elem()
}

func (o ExportConfigurationOutput) ToExportConfigurationOutput() ExportConfigurationOutput {
	return o
}

func (o ExportConfigurationOutput) ToExportConfigurationOutputWithContext(ctx context.Context) ExportConfigurationOutput {
	return o
}

// The name of the Application Insights component.
func (o ExportConfigurationOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.ApplicationName }).(pulumi.StringOutput)
}

// The name of the destination storage container.
func (o ExportConfigurationOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.ContainerName }).(pulumi.StringOutput)
}

// The name of destination account.
func (o ExportConfigurationOutput) DestinationAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.DestinationAccountId }).(pulumi.StringOutput)
}

// The destination account location ID.
func (o ExportConfigurationOutput) DestinationStorageLocationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.DestinationStorageLocationId }).(pulumi.StringOutput)
}

// The destination storage account subscription ID.
func (o ExportConfigurationOutput) DestinationStorageSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.DestinationStorageSubscriptionId }).(pulumi.StringOutput)
}

// The destination type.
func (o ExportConfigurationOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

// The unique ID of the export configuration inside an Application Insights component. It is auto generated when the Continuous Export configuration is created.
func (o ExportConfigurationOutput) ExportId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.ExportId }).(pulumi.StringOutput)
}

// This indicates current Continuous Export configuration status. The possible values are 'Preparing', 'Success', 'Failure'.
func (o ExportConfigurationOutput) ExportStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.ExportStatus }).(pulumi.StringOutput)
}

// The instrumentation key of the Application Insights component.
func (o ExportConfigurationOutput) InstrumentationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.InstrumentationKey }).(pulumi.StringOutput)
}

// This will be 'true' if the Continuous Export configuration is enabled, otherwise it will be 'false'.
func (o ExportConfigurationOutput) IsUserEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.IsUserEnabled }).(pulumi.StringOutput)
}

// The last time the Continuous Export configuration started failing.
func (o ExportConfigurationOutput) LastGapTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.LastGapTime }).(pulumi.StringOutput)
}

// The last time data was successfully delivered to the destination storage container for this Continuous Export configuration.
func (o ExportConfigurationOutput) LastSuccessTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.LastSuccessTime }).(pulumi.StringOutput)
}

// Last time the Continuous Export configuration was updated.
func (o ExportConfigurationOutput) LastUserUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.LastUserUpdate }).(pulumi.StringOutput)
}

// Deprecated
func (o ExportConfigurationOutput) NotificationQueueEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringPtrOutput { return v.NotificationQueueEnabled }).(pulumi.StringPtrOutput)
}

// This is the reason the Continuous Export configuration started failing. It can be 'AzureStorageNotFound' or 'AzureStorageAccessDenied'.
func (o ExportConfigurationOutput) PermanentErrorReason() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.PermanentErrorReason }).(pulumi.StringOutput)
}

// This comma separated list of document types that will be exported. The possible values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
func (o ExportConfigurationOutput) RecordTypes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringPtrOutput { return v.RecordTypes }).(pulumi.StringPtrOutput)
}

// The resource group of the Application Insights component.
func (o ExportConfigurationOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

// The name of the destination storage account.
func (o ExportConfigurationOutput) StorageName() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.StorageName }).(pulumi.StringOutput)
}

// The subscription of the Application Insights component.
func (o ExportConfigurationOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExportConfigurationOutput{})
}
