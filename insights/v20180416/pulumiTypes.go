// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180416

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// Specify action need to be taken when rule type is Alert
type AlertingAction struct {
	// Azure action group reference.
	AznsAction *AzNsActionGroup `pulumi:"aznsAction"`
	// Specifies the action. Supported values - AlertingAction, LogToMetricAction
	// Expected value is 'Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.AlertingAction'.
	OdataType string `pulumi:"odataType"`
	// Severity of the alert
	Severity string `pulumi:"severity"`
	// time (in minutes) for which Alerts should be throttled or suppressed.
	ThrottlingInMin *int `pulumi:"throttlingInMin"`
	// The trigger condition that results in the alert rule being.
	Trigger TriggerCondition `pulumi:"trigger"`
}

// Specify action need to be taken when rule type is Alert
type AlertingActionResponse struct {
	// Azure action group reference.
	AznsAction *AzNsActionGroupResponse `pulumi:"aznsAction"`
	// Specifies the action. Supported values - AlertingAction, LogToMetricAction
	// Expected value is 'Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.AlertingAction'.
	OdataType string `pulumi:"odataType"`
	// Severity of the alert
	Severity string `pulumi:"severity"`
	// time (in minutes) for which Alerts should be throttled or suppressed.
	ThrottlingInMin *int `pulumi:"throttlingInMin"`
	// The trigger condition that results in the alert rule being.
	Trigger TriggerConditionResponse `pulumi:"trigger"`
}

// Azure action group
type AzNsActionGroup struct {
	// Azure Action Group reference.
	ActionGroup []string `pulumi:"actionGroup"`
	// Custom payload to be sent for all webhook URI in Azure action group
	CustomWebhookPayload *string `pulumi:"customWebhookPayload"`
	// Custom subject override for all email ids in Azure action group
	EmailSubject *string `pulumi:"emailSubject"`
}

// Azure action group
type AzNsActionGroupResponse struct {
	// Azure Action Group reference.
	ActionGroup []string `pulumi:"actionGroup"`
	// Custom payload to be sent for all webhook URI in Azure action group
	CustomWebhookPayload *string `pulumi:"customWebhookPayload"`
	// Custom subject override for all email ids in Azure action group
	EmailSubject *string `pulumi:"emailSubject"`
}

// Specifies the criteria for converting log to metric.
type Criteria struct {
	// List of Dimensions for creating metric
	Dimensions []Dimension `pulumi:"dimensions"`
	// Name of the metric
	MetricName string `pulumi:"metricName"`
}

// Specifies the criteria for converting log to metric.
type CriteriaResponse struct {
	// List of Dimensions for creating metric
	Dimensions []DimensionResponse `pulumi:"dimensions"`
	// Name of the metric
	MetricName string `pulumi:"metricName"`
}

// Specifies the criteria for converting log to metric.
type Dimension struct {
	// Name of the dimension
	Name string `pulumi:"name"`
	// Operator for dimension values
	Operator string `pulumi:"operator"`
	// List of dimension values
	Values []string `pulumi:"values"`
}

// Specifies the criteria for converting log to metric.
type DimensionResponse struct {
	// Name of the dimension
	Name string `pulumi:"name"`
	// Operator for dimension values
	Operator string `pulumi:"operator"`
	// List of dimension values
	Values []string `pulumi:"values"`
}

// A log metrics trigger descriptor.
type LogMetricTrigger struct {
	// Evaluation of metric on a particular column
	MetricColumn *string `pulumi:"metricColumn"`
	// Metric Trigger Type - 'Consecutive' or 'Total'
	MetricTriggerType *string `pulumi:"metricTriggerType"`
	// The threshold of the metric trigger.
	Threshold *float64 `pulumi:"threshold"`
	// Evaluation operation for Metric -'GreaterThan' or 'LessThan' or 'Equal'.
	ThresholdOperator *string `pulumi:"thresholdOperator"`
}

// A log metrics trigger descriptor.
type LogMetricTriggerResponse struct {
	// Evaluation of metric on a particular column
	MetricColumn *string `pulumi:"metricColumn"`
	// Metric Trigger Type - 'Consecutive' or 'Total'
	MetricTriggerType *string `pulumi:"metricTriggerType"`
	// The threshold of the metric trigger.
	Threshold *float64 `pulumi:"threshold"`
	// Evaluation operation for Metric -'GreaterThan' or 'LessThan' or 'Equal'.
	ThresholdOperator *string `pulumi:"thresholdOperator"`
}

// Specify action need to be taken when rule type is converting log to metric
type LogToMetricAction struct {
	// Criteria of Metric
	Criteria []Criteria `pulumi:"criteria"`
	// Specifies the action. Supported values - AlertingAction, LogToMetricAction
	// Expected value is 'Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.LogToMetricAction'.
	OdataType string `pulumi:"odataType"`
}

// Specify action need to be taken when rule type is converting log to metric
type LogToMetricActionResponse struct {
	// Criteria of Metric
	Criteria []CriteriaResponse `pulumi:"criteria"`
	// Specifies the action. Supported values - AlertingAction, LogToMetricAction
	// Expected value is 'Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.LogToMetricAction'.
	OdataType string `pulumi:"odataType"`
}

// Defines how often to run the search and the time interval.
type Schedule struct {
	// frequency (in minutes) at which rule condition should be evaluated.
	FrequencyInMinutes int `pulumi:"frequencyInMinutes"`
	// Time window for which data needs to be fetched for query (should be greater than or equal to frequencyInMinutes).
	TimeWindowInMinutes int `pulumi:"timeWindowInMinutes"`
}

// ScheduleInput is an input type that accepts ScheduleArgs and ScheduleOutput values.
// You can construct a concrete instance of `ScheduleInput` via:
//
//	ScheduleArgs{...}
type ScheduleInput interface {
	pulumi.Input

	ToScheduleOutput() ScheduleOutput
	ToScheduleOutputWithContext(context.Context) ScheduleOutput
}

// Defines how often to run the search and the time interval.
type ScheduleArgs struct {
	// frequency (in minutes) at which rule condition should be evaluated.
	FrequencyInMinutes pulumi.IntInput `pulumi:"frequencyInMinutes"`
	// Time window for which data needs to be fetched for query (should be greater than or equal to frequencyInMinutes).
	TimeWindowInMinutes pulumi.IntInput `pulumi:"timeWindowInMinutes"`
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (i ScheduleArgs) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i ScheduleArgs) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}

func (i ScheduleArgs) ToOutput(ctx context.Context) pulumix.Output[Schedule] {
	return pulumix.Output[Schedule]{
		OutputState: i.ToScheduleOutputWithContext(ctx).OutputState,
	}
}

func (i ScheduleArgs) ToSchedulePtrOutput() SchedulePtrOutput {
	return i.ToSchedulePtrOutputWithContext(context.Background())
}

func (i ScheduleArgs) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput).ToSchedulePtrOutputWithContext(ctx)
}

// SchedulePtrInput is an input type that accepts ScheduleArgs, SchedulePtr and SchedulePtrOutput values.
// You can construct a concrete instance of `SchedulePtrInput` via:
//
//	        ScheduleArgs{...}
//
//	or:
//
//	        nil
type SchedulePtrInput interface {
	pulumi.Input

	ToSchedulePtrOutput() SchedulePtrOutput
	ToSchedulePtrOutputWithContext(context.Context) SchedulePtrOutput
}

type schedulePtrType ScheduleArgs

func SchedulePtr(v *ScheduleArgs) SchedulePtrInput {
	return (*schedulePtrType)(v)
}

func (*schedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (i *schedulePtrType) ToSchedulePtrOutput() SchedulePtrOutput {
	return i.ToSchedulePtrOutputWithContext(context.Background())
}

func (i *schedulePtrType) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulePtrOutput)
}

func (i *schedulePtrType) ToOutput(ctx context.Context) pulumix.Output[*Schedule] {
	return pulumix.Output[*Schedule]{
		OutputState: i.ToSchedulePtrOutputWithContext(ctx).OutputState,
	}
}

// Defines how often to run the search and the time interval.
type ScheduleOutput struct{ *pulumi.OutputState }

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToSchedulePtrOutput() SchedulePtrOutput {
	return o.ToSchedulePtrOutputWithContext(context.Background())
}

func (o ScheduleOutput) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Schedule) *Schedule {
		return &v
	}).(SchedulePtrOutput)
}

func (o ScheduleOutput) ToOutput(ctx context.Context) pulumix.Output[Schedule] {
	return pulumix.Output[Schedule]{
		OutputState: o.OutputState,
	}
}

// frequency (in minutes) at which rule condition should be evaluated.
func (o ScheduleOutput) FrequencyInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v Schedule) int { return v.FrequencyInMinutes }).(pulumi.IntOutput)
}

// Time window for which data needs to be fetched for query (should be greater than or equal to frequencyInMinutes).
func (o ScheduleOutput) TimeWindowInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v Schedule) int { return v.TimeWindowInMinutes }).(pulumi.IntOutput)
}

type SchedulePtrOutput struct{ *pulumi.OutputState }

func (SchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (o SchedulePtrOutput) ToSchedulePtrOutput() SchedulePtrOutput {
	return o
}

func (o SchedulePtrOutput) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return o
}

func (o SchedulePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*Schedule] {
	return pulumix.Output[*Schedule]{
		OutputState: o.OutputState,
	}
}

func (o SchedulePtrOutput) Elem() ScheduleOutput {
	return o.ApplyT(func(v *Schedule) Schedule {
		if v != nil {
			return *v
		}
		var ret Schedule
		return ret
	}).(ScheduleOutput)
}

// frequency (in minutes) at which rule condition should be evaluated.
func (o SchedulePtrOutput) FrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Schedule) *int {
		if v == nil {
			return nil
		}
		return &v.FrequencyInMinutes
	}).(pulumi.IntPtrOutput)
}

// Time window for which data needs to be fetched for query (should be greater than or equal to frequencyInMinutes).
func (o SchedulePtrOutput) TimeWindowInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Schedule) *int {
		if v == nil {
			return nil
		}
		return &v.TimeWindowInMinutes
	}).(pulumi.IntPtrOutput)
}

// Defines how often to run the search and the time interval.
type ScheduleResponse struct {
	// frequency (in minutes) at which rule condition should be evaluated.
	FrequencyInMinutes int `pulumi:"frequencyInMinutes"`
	// Time window for which data needs to be fetched for query (should be greater than or equal to frequencyInMinutes).
	TimeWindowInMinutes int `pulumi:"timeWindowInMinutes"`
}

// Defines how often to run the search and the time interval.
type ScheduleResponseOutput struct{ *pulumi.OutputState }

func (ScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponseOutput) ToScheduleResponseOutput() ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) ToScheduleResponseOutputWithContext(ctx context.Context) ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ScheduleResponse] {
	return pulumix.Output[ScheduleResponse]{
		OutputState: o.OutputState,
	}
}

// frequency (in minutes) at which rule condition should be evaluated.
func (o ScheduleResponseOutput) FrequencyInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleResponse) int { return v.FrequencyInMinutes }).(pulumi.IntOutput)
}

// Time window for which data needs to be fetched for query (should be greater than or equal to frequencyInMinutes).
func (o ScheduleResponseOutput) TimeWindowInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleResponse) int { return v.TimeWindowInMinutes }).(pulumi.IntOutput)
}

type ScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (ScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponsePtrOutput) ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput {
	return o
}

func (o ScheduleResponsePtrOutput) ToScheduleResponsePtrOutputWithContext(ctx context.Context) ScheduleResponsePtrOutput {
	return o
}

func (o ScheduleResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ScheduleResponse] {
	return pulumix.Output[*ScheduleResponse]{
		OutputState: o.OutputState,
	}
}

func (o ScheduleResponsePtrOutput) Elem() ScheduleResponseOutput {
	return o.ApplyT(func(v *ScheduleResponse) ScheduleResponse {
		if v != nil {
			return *v
		}
		var ret ScheduleResponse
		return ret
	}).(ScheduleResponseOutput)
}

// frequency (in minutes) at which rule condition should be evaluated.
func (o ScheduleResponsePtrOutput) FrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.FrequencyInMinutes
	}).(pulumi.IntPtrOutput)
}

// Time window for which data needs to be fetched for query (should be greater than or equal to frequencyInMinutes).
func (o ScheduleResponsePtrOutput) TimeWindowInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TimeWindowInMinutes
	}).(pulumi.IntPtrOutput)
}

// Specifies the log search query.
type Source struct {
	// List of  Resource referred into query
	AuthorizedResources []string `pulumi:"authorizedResources"`
	// The resource uri over which log search query is to be run.
	DataSourceId string `pulumi:"dataSourceId"`
	// Log search query. Required for action type - AlertingAction
	Query *string `pulumi:"query"`
	// Set value to 'ResultCount' .
	QueryType *string `pulumi:"queryType"`
}

// SourceInput is an input type that accepts SourceArgs and SourceOutput values.
// You can construct a concrete instance of `SourceInput` via:
//
//	SourceArgs{...}
type SourceInput interface {
	pulumi.Input

	ToSourceOutput() SourceOutput
	ToSourceOutputWithContext(context.Context) SourceOutput
}

// Specifies the log search query.
type SourceArgs struct {
	// List of  Resource referred into query
	AuthorizedResources pulumi.StringArrayInput `pulumi:"authorizedResources"`
	// The resource uri over which log search query is to be run.
	DataSourceId pulumi.StringInput `pulumi:"dataSourceId"`
	// Log search query. Required for action type - AlertingAction
	Query pulumi.StringPtrInput `pulumi:"query"`
	// Set value to 'ResultCount' .
	QueryType pulumi.StringPtrInput `pulumi:"queryType"`
}

func (SourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (i SourceArgs) ToSourceOutput() SourceOutput {
	return i.ToSourceOutputWithContext(context.Background())
}

func (i SourceArgs) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutput)
}

func (i SourceArgs) ToOutput(ctx context.Context) pulumix.Output[Source] {
	return pulumix.Output[Source]{
		OutputState: i.ToSourceOutputWithContext(ctx).OutputState,
	}
}

// Specifies the log search query.
type SourceOutput struct{ *pulumi.OutputState }

func (SourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (o SourceOutput) ToSourceOutput() SourceOutput {
	return o
}

func (o SourceOutput) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return o
}

func (o SourceOutput) ToOutput(ctx context.Context) pulumix.Output[Source] {
	return pulumix.Output[Source]{
		OutputState: o.OutputState,
	}
}

// List of  Resource referred into query
func (o SourceOutput) AuthorizedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Source) []string { return v.AuthorizedResources }).(pulumi.StringArrayOutput)
}

// The resource uri over which log search query is to be run.
func (o SourceOutput) DataSourceId() pulumi.StringOutput {
	return o.ApplyT(func(v Source) string { return v.DataSourceId }).(pulumi.StringOutput)
}

// Log search query. Required for action type - AlertingAction
func (o SourceOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Source) *string { return v.Query }).(pulumi.StringPtrOutput)
}

// Set value to 'ResultCount' .
func (o SourceOutput) QueryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Source) *string { return v.QueryType }).(pulumi.StringPtrOutput)
}

// Specifies the log search query.
type SourceResponse struct {
	// List of  Resource referred into query
	AuthorizedResources []string `pulumi:"authorizedResources"`
	// The resource uri over which log search query is to be run.
	DataSourceId string `pulumi:"dataSourceId"`
	// Log search query. Required for action type - AlertingAction
	Query *string `pulumi:"query"`
	// Set value to 'ResultCount' .
	QueryType *string `pulumi:"queryType"`
}

// Specifies the log search query.
type SourceResponseOutput struct{ *pulumi.OutputState }

func (SourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceResponse)(nil)).Elem()
}

func (o SourceResponseOutput) ToSourceResponseOutput() SourceResponseOutput {
	return o
}

func (o SourceResponseOutput) ToSourceResponseOutputWithContext(ctx context.Context) SourceResponseOutput {
	return o
}

func (o SourceResponseOutput) ToOutput(ctx context.Context) pulumix.Output[SourceResponse] {
	return pulumix.Output[SourceResponse]{
		OutputState: o.OutputState,
	}
}

// List of  Resource referred into query
func (o SourceResponseOutput) AuthorizedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SourceResponse) []string { return v.AuthorizedResources }).(pulumi.StringArrayOutput)
}

// The resource uri over which log search query is to be run.
func (o SourceResponseOutput) DataSourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SourceResponse) string { return v.DataSourceId }).(pulumi.StringOutput)
}

// Log search query. Required for action type - AlertingAction
func (o SourceResponseOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceResponse) *string { return v.Query }).(pulumi.StringPtrOutput)
}

// Set value to 'ResultCount' .
func (o SourceResponseOutput) QueryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceResponse) *string { return v.QueryType }).(pulumi.StringPtrOutput)
}

// The condition that results in the Log Search rule.
type TriggerCondition struct {
	// Trigger condition for metric query rule
	MetricTrigger *LogMetricTrigger `pulumi:"metricTrigger"`
	// Result or count threshold based on which rule should be triggered.
	Threshold float64 `pulumi:"threshold"`
	// Evaluation operation for rule - 'GreaterThan' or 'LessThan.
	ThresholdOperator string `pulumi:"thresholdOperator"`
}

// The condition that results in the Log Search rule.
type TriggerConditionResponse struct {
	// Trigger condition for metric query rule
	MetricTrigger *LogMetricTriggerResponse `pulumi:"metricTrigger"`
	// Result or count threshold based on which rule should be triggered.
	Threshold float64 `pulumi:"threshold"`
	// Evaluation operation for rule - 'GreaterThan' or 'LessThan.
	ThresholdOperator string `pulumi:"thresholdOperator"`
}

func init() {
	pulumi.RegisterOutputType(ScheduleOutput{})
	pulumi.RegisterOutputType(SchedulePtrOutput{})
	pulumi.RegisterOutputType(ScheduleResponseOutput{})
	pulumi.RegisterOutputType(ScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceOutput{})
	pulumi.RegisterOutputType(SourceResponseOutput{})
}
