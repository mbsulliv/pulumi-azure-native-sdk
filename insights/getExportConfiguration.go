// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the Continuous Export configuration for this export id.
// Azure REST API version: 2015-05-01.
func LookupExportConfiguration(ctx *pulumi.Context, args *LookupExportConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupExportConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExportConfigurationResult
	err := ctx.Invoke("azure-native:insights:getExportConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExportConfigurationArgs struct {
	// The Continuous Export configuration ID. This is unique within a Application Insights component.
	ExportId string `pulumi:"exportId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName string `pulumi:"resourceName"`
}

// Properties that define a Continuous Export configuration.
type LookupExportConfigurationResult struct {
	// The name of the Application Insights component.
	ApplicationName string `pulumi:"applicationName"`
	// The name of the destination storage container.
	ContainerName string `pulumi:"containerName"`
	// The name of destination account.
	DestinationAccountId string `pulumi:"destinationAccountId"`
	// The destination account location ID.
	DestinationStorageLocationId string `pulumi:"destinationStorageLocationId"`
	// The destination storage account subscription ID.
	DestinationStorageSubscriptionId string `pulumi:"destinationStorageSubscriptionId"`
	// The destination type.
	DestinationType string `pulumi:"destinationType"`
	// The unique ID of the export configuration inside an Application Insights component. It is auto generated when the Continuous Export configuration is created.
	ExportId string `pulumi:"exportId"`
	// This indicates current Continuous Export configuration status. The possible values are 'Preparing', 'Success', 'Failure'.
	ExportStatus string `pulumi:"exportStatus"`
	// The instrumentation key of the Application Insights component.
	InstrumentationKey string `pulumi:"instrumentationKey"`
	// This will be 'true' if the Continuous Export configuration is enabled, otherwise it will be 'false'.
	IsUserEnabled string `pulumi:"isUserEnabled"`
	// The last time the Continuous Export configuration started failing.
	LastGapTime string `pulumi:"lastGapTime"`
	// The last time data was successfully delivered to the destination storage container for this Continuous Export configuration.
	LastSuccessTime string `pulumi:"lastSuccessTime"`
	// Last time the Continuous Export configuration was updated.
	LastUserUpdate string `pulumi:"lastUserUpdate"`
	// Deprecated
	NotificationQueueEnabled *string `pulumi:"notificationQueueEnabled"`
	// This is the reason the Continuous Export configuration started failing. It can be 'AzureStorageNotFound' or 'AzureStorageAccessDenied'.
	PermanentErrorReason string `pulumi:"permanentErrorReason"`
	// This comma separated list of document types that will be exported. The possible values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
	RecordTypes *string `pulumi:"recordTypes"`
	// The resource group of the Application Insights component.
	ResourceGroup string `pulumi:"resourceGroup"`
	// The name of the destination storage account.
	StorageName string `pulumi:"storageName"`
	// The subscription of the Application Insights component.
	SubscriptionId string `pulumi:"subscriptionId"`
}

func LookupExportConfigurationOutput(ctx *pulumi.Context, args LookupExportConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupExportConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExportConfigurationResult, error) {
			args := v.(LookupExportConfigurationArgs)
			r, err := LookupExportConfiguration(ctx, &args, opts...)
			var s LookupExportConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExportConfigurationResultOutput)
}

type LookupExportConfigurationOutputArgs struct {
	// The Continuous Export configuration ID. This is unique within a Application Insights component.
	ExportId pulumi.StringInput `pulumi:"exportId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupExportConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportConfigurationArgs)(nil)).Elem()
}

// Properties that define a Continuous Export configuration.
type LookupExportConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupExportConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportConfigurationResult)(nil)).Elem()
}

func (o LookupExportConfigurationResultOutput) ToLookupExportConfigurationResultOutput() LookupExportConfigurationResultOutput {
	return o
}

func (o LookupExportConfigurationResultOutput) ToLookupExportConfigurationResultOutputWithContext(ctx context.Context) LookupExportConfigurationResultOutput {
	return o
}

// The name of the Application Insights component.
func (o LookupExportConfigurationResultOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.ApplicationName }).(pulumi.StringOutput)
}

// The name of the destination storage container.
func (o LookupExportConfigurationResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

// The name of destination account.
func (o LookupExportConfigurationResultOutput) DestinationAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.DestinationAccountId }).(pulumi.StringOutput)
}

// The destination account location ID.
func (o LookupExportConfigurationResultOutput) DestinationStorageLocationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.DestinationStorageLocationId }).(pulumi.StringOutput)
}

// The destination storage account subscription ID.
func (o LookupExportConfigurationResultOutput) DestinationStorageSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.DestinationStorageSubscriptionId }).(pulumi.StringOutput)
}

// The destination type.
func (o LookupExportConfigurationResultOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.DestinationType }).(pulumi.StringOutput)
}

// The unique ID of the export configuration inside an Application Insights component. It is auto generated when the Continuous Export configuration is created.
func (o LookupExportConfigurationResultOutput) ExportId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.ExportId }).(pulumi.StringOutput)
}

// This indicates current Continuous Export configuration status. The possible values are 'Preparing', 'Success', 'Failure'.
func (o LookupExportConfigurationResultOutput) ExportStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.ExportStatus }).(pulumi.StringOutput)
}

// The instrumentation key of the Application Insights component.
func (o LookupExportConfigurationResultOutput) InstrumentationKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.InstrumentationKey }).(pulumi.StringOutput)
}

// This will be 'true' if the Continuous Export configuration is enabled, otherwise it will be 'false'.
func (o LookupExportConfigurationResultOutput) IsUserEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.IsUserEnabled }).(pulumi.StringOutput)
}

// The last time the Continuous Export configuration started failing.
func (o LookupExportConfigurationResultOutput) LastGapTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.LastGapTime }).(pulumi.StringOutput)
}

// The last time data was successfully delivered to the destination storage container for this Continuous Export configuration.
func (o LookupExportConfigurationResultOutput) LastSuccessTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.LastSuccessTime }).(pulumi.StringOutput)
}

// Last time the Continuous Export configuration was updated.
func (o LookupExportConfigurationResultOutput) LastUserUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.LastUserUpdate }).(pulumi.StringOutput)
}

// Deprecated
func (o LookupExportConfigurationResultOutput) NotificationQueueEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) *string { return v.NotificationQueueEnabled }).(pulumi.StringPtrOutput)
}

// This is the reason the Continuous Export configuration started failing. It can be 'AzureStorageNotFound' or 'AzureStorageAccessDenied'.
func (o LookupExportConfigurationResultOutput) PermanentErrorReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.PermanentErrorReason }).(pulumi.StringOutput)
}

// This comma separated list of document types that will be exported. The possible values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
func (o LookupExportConfigurationResultOutput) RecordTypes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) *string { return v.RecordTypes }).(pulumi.StringPtrOutput)
}

// The resource group of the Application Insights component.
func (o LookupExportConfigurationResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

// The name of the destination storage account.
func (o LookupExportConfigurationResultOutput) StorageName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.StorageName }).(pulumi.StringOutput)
}

// The subscription of the Application Insights component.
func (o LookupExportConfigurationResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExportConfigurationResultOutput{})
}
