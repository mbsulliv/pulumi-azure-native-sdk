// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Actions to invoke when the alert fires.
type Action struct {
	// Action Group resource Id to invoke when the alert fires.
	ActionGroupId *string `pulumi:"actionGroupId"`
	// The properties of a webhook object.
	WebHookProperties map[string]string `pulumi:"webHookProperties"`
}

// ActionInput is an input type that accepts ActionArgs and ActionOutput values.
// You can construct a concrete instance of `ActionInput` via:
//
//	ActionArgs{...}
type ActionInput interface {
	pulumi.Input

	ToActionOutput() ActionOutput
	ToActionOutputWithContext(context.Context) ActionOutput
}

// Actions to invoke when the alert fires.
type ActionArgs struct {
	// Action Group resource Id to invoke when the alert fires.
	ActionGroupId pulumi.StringPtrInput `pulumi:"actionGroupId"`
	// The properties of a webhook object.
	WebHookProperties pulumi.StringMapInput `pulumi:"webHookProperties"`
}

func (ActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (i ActionArgs) ToActionOutput() ActionOutput {
	return i.ToActionOutputWithContext(context.Background())
}

func (i ActionArgs) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionOutput)
}

// ActionArrayInput is an input type that accepts ActionArray and ActionArrayOutput values.
// You can construct a concrete instance of `ActionArrayInput` via:
//
//	ActionArray{ ActionArgs{...} }
type ActionArrayInput interface {
	pulumi.Input

	ToActionArrayOutput() ActionArrayOutput
	ToActionArrayOutputWithContext(context.Context) ActionArrayOutput
}

type ActionArray []ActionInput

func (ActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Action)(nil)).Elem()
}

func (i ActionArray) ToActionArrayOutput() ActionArrayOutput {
	return i.ToActionArrayOutputWithContext(context.Background())
}

func (i ActionArray) ToActionArrayOutputWithContext(ctx context.Context) ActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionArrayOutput)
}

// Actions to invoke when the alert fires.
type ActionOutput struct{ *pulumi.OutputState }

func (ActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (o ActionOutput) ToActionOutput() ActionOutput {
	return o
}

func (o ActionOutput) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return o
}

// Action Group resource Id to invoke when the alert fires.
func (o ActionOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Action) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

// The properties of a webhook object.
func (o ActionOutput) WebHookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v Action) map[string]string { return v.WebHookProperties }).(pulumi.StringMapOutput)
}

type ActionArrayOutput struct{ *pulumi.OutputState }

func (ActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Action)(nil)).Elem()
}

func (o ActionArrayOutput) ToActionArrayOutput() ActionArrayOutput {
	return o
}

func (o ActionArrayOutput) ToActionArrayOutputWithContext(ctx context.Context) ActionArrayOutput {
	return o
}

func (o ActionArrayOutput) Index(i pulumi.IntInput) ActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Action {
		return vs[0].([]Action)[vs[1].(int)]
	}).(ActionOutput)
}

// Actions to invoke when the alert fires.
type ActionResponse struct {
	// Action Group resource Id to invoke when the alert fires.
	ActionGroupId *string `pulumi:"actionGroupId"`
	// The properties of a webhook object.
	WebHookProperties map[string]string `pulumi:"webHookProperties"`
}

// Actions to invoke when the alert fires.
type ActionResponseOutput struct{ *pulumi.OutputState }

func (ActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionResponse)(nil)).Elem()
}

func (o ActionResponseOutput) ToActionResponseOutput() ActionResponseOutput {
	return o
}

func (o ActionResponseOutput) ToActionResponseOutputWithContext(ctx context.Context) ActionResponseOutput {
	return o
}

// Action Group resource Id to invoke when the alert fires.
func (o ActionResponseOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionResponse) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

// The properties of a webhook object.
func (o ActionResponseOutput) WebHookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ActionResponse) map[string]string { return v.WebHookProperties }).(pulumi.StringMapOutput)
}

type ActionResponseArrayOutput struct{ *pulumi.OutputState }

func (ActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActionResponse)(nil)).Elem()
}

func (o ActionResponseArrayOutput) ToActionResponseArrayOutput() ActionResponseArrayOutput {
	return o
}

func (o ActionResponseArrayOutput) ToActionResponseArrayOutputWithContext(ctx context.Context) ActionResponseArrayOutput {
	return o
}

func (o ActionResponseArrayOutput) Index(i pulumi.IntInput) ActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActionResponse {
		return vs[0].([]ActionResponse)[vs[1].(int)]
	}).(ActionResponseOutput)
}

// A condition of the scheduled query rule.
type Condition struct {
	// List of Dimensions conditions
	Dimensions []Dimension `pulumi:"dimensions"`
	// The minimum number of violations required within the selected lookback time window required to raise an alert.
	FailingPeriods *ConditionFailingPeriods `pulumi:"failingPeriods"`
	// The column containing the metric measure number.
	MetricMeasureColumn *string `pulumi:"metricMeasureColumn"`
	// The criteria operator.
	Operator string `pulumi:"operator"`
	// Log query alert
	Query *string `pulumi:"query"`
	// The column containing the resource id. The content of the column must be a uri formatted as resource id
	ResourceIdColumn *string `pulumi:"resourceIdColumn"`
	// the criteria threshold value that activates the alert.
	Threshold float64 `pulumi:"threshold"`
	// Aggregation type
	TimeAggregation string `pulumi:"timeAggregation"`
}

// Defaults sets the appropriate defaults for Condition
func (val *Condition) Defaults() *Condition {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.FailingPeriods = tmp.FailingPeriods.Defaults()

	return &tmp
}

// ConditionInput is an input type that accepts ConditionArgs and ConditionOutput values.
// You can construct a concrete instance of `ConditionInput` via:
//
//	ConditionArgs{...}
type ConditionInput interface {
	pulumi.Input

	ToConditionOutput() ConditionOutput
	ToConditionOutputWithContext(context.Context) ConditionOutput
}

// A condition of the scheduled query rule.
type ConditionArgs struct {
	// List of Dimensions conditions
	Dimensions DimensionArrayInput `pulumi:"dimensions"`
	// The minimum number of violations required within the selected lookback time window required to raise an alert.
	FailingPeriods ConditionFailingPeriodsPtrInput `pulumi:"failingPeriods"`
	// The column containing the metric measure number.
	MetricMeasureColumn pulumi.StringPtrInput `pulumi:"metricMeasureColumn"`
	// The criteria operator.
	Operator pulumi.StringInput `pulumi:"operator"`
	// Log query alert
	Query pulumi.StringPtrInput `pulumi:"query"`
	// The column containing the resource id. The content of the column must be a uri formatted as resource id
	ResourceIdColumn pulumi.StringPtrInput `pulumi:"resourceIdColumn"`
	// the criteria threshold value that activates the alert.
	Threshold pulumi.Float64Input `pulumi:"threshold"`
	// Aggregation type
	TimeAggregation pulumi.StringInput `pulumi:"timeAggregation"`
}

// Defaults sets the appropriate defaults for ConditionArgs
func (val *ConditionArgs) Defaults() *ConditionArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (ConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Condition)(nil)).Elem()
}

func (i ConditionArgs) ToConditionOutput() ConditionOutput {
	return i.ToConditionOutputWithContext(context.Background())
}

func (i ConditionArgs) ToConditionOutputWithContext(ctx context.Context) ConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionOutput)
}

// ConditionArrayInput is an input type that accepts ConditionArray and ConditionArrayOutput values.
// You can construct a concrete instance of `ConditionArrayInput` via:
//
//	ConditionArray{ ConditionArgs{...} }
type ConditionArrayInput interface {
	pulumi.Input

	ToConditionArrayOutput() ConditionArrayOutput
	ToConditionArrayOutputWithContext(context.Context) ConditionArrayOutput
}

type ConditionArray []ConditionInput

func (ConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Condition)(nil)).Elem()
}

func (i ConditionArray) ToConditionArrayOutput() ConditionArrayOutput {
	return i.ToConditionArrayOutputWithContext(context.Background())
}

func (i ConditionArray) ToConditionArrayOutputWithContext(ctx context.Context) ConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionArrayOutput)
}

// A condition of the scheduled query rule.
type ConditionOutput struct{ *pulumi.OutputState }

func (ConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Condition)(nil)).Elem()
}

func (o ConditionOutput) ToConditionOutput() ConditionOutput {
	return o
}

func (o ConditionOutput) ToConditionOutputWithContext(ctx context.Context) ConditionOutput {
	return o
}

// List of Dimensions conditions
func (o ConditionOutput) Dimensions() DimensionArrayOutput {
	return o.ApplyT(func(v Condition) []Dimension { return v.Dimensions }).(DimensionArrayOutput)
}

// The minimum number of violations required within the selected lookback time window required to raise an alert.
func (o ConditionOutput) FailingPeriods() ConditionFailingPeriodsPtrOutput {
	return o.ApplyT(func(v Condition) *ConditionFailingPeriods { return v.FailingPeriods }).(ConditionFailingPeriodsPtrOutput)
}

// The column containing the metric measure number.
func (o ConditionOutput) MetricMeasureColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.MetricMeasureColumn }).(pulumi.StringPtrOutput)
}

// The criteria operator.
func (o ConditionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v Condition) string { return v.Operator }).(pulumi.StringOutput)
}

// Log query alert
func (o ConditionOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.Query }).(pulumi.StringPtrOutput)
}

// The column containing the resource id. The content of the column must be a uri formatted as resource id
func (o ConditionOutput) ResourceIdColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.ResourceIdColumn }).(pulumi.StringPtrOutput)
}

// the criteria threshold value that activates the alert.
func (o ConditionOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v Condition) float64 { return v.Threshold }).(pulumi.Float64Output)
}

// Aggregation type
func (o ConditionOutput) TimeAggregation() pulumi.StringOutput {
	return o.ApplyT(func(v Condition) string { return v.TimeAggregation }).(pulumi.StringOutput)
}

type ConditionArrayOutput struct{ *pulumi.OutputState }

func (ConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Condition)(nil)).Elem()
}

func (o ConditionArrayOutput) ToConditionArrayOutput() ConditionArrayOutput {
	return o
}

func (o ConditionArrayOutput) ToConditionArrayOutputWithContext(ctx context.Context) ConditionArrayOutput {
	return o
}

func (o ConditionArrayOutput) Index(i pulumi.IntInput) ConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Condition {
		return vs[0].([]Condition)[vs[1].(int)]
	}).(ConditionOutput)
}

// The minimum number of violations required within the selected lookback time window required to raise an alert.
type ConditionFailingPeriods struct {
	// The number of violations to trigger an alert. Should be smaller or equal to numberOfEvaluationPeriods. Default value is 1
	MinFailingPeriodsToAlert *float64 `pulumi:"minFailingPeriodsToAlert"`
	// The number of aggregated lookback points. The lookback time window is calculated based on the aggregation granularity (windowSize) and the selected number of aggregated points. Default value is 1
	NumberOfEvaluationPeriods *float64 `pulumi:"numberOfEvaluationPeriods"`
}

// Defaults sets the appropriate defaults for ConditionFailingPeriods
func (val *ConditionFailingPeriods) Defaults() *ConditionFailingPeriods {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.MinFailingPeriodsToAlert == nil {
		minFailingPeriodsToAlert_ := 1.0
		tmp.MinFailingPeriodsToAlert = &minFailingPeriodsToAlert_
	}
	if tmp.NumberOfEvaluationPeriods == nil {
		numberOfEvaluationPeriods_ := 1.0
		tmp.NumberOfEvaluationPeriods = &numberOfEvaluationPeriods_
	}
	return &tmp
}

// ConditionFailingPeriodsInput is an input type that accepts ConditionFailingPeriodsArgs and ConditionFailingPeriodsOutput values.
// You can construct a concrete instance of `ConditionFailingPeriodsInput` via:
//
//	ConditionFailingPeriodsArgs{...}
type ConditionFailingPeriodsInput interface {
	pulumi.Input

	ToConditionFailingPeriodsOutput() ConditionFailingPeriodsOutput
	ToConditionFailingPeriodsOutputWithContext(context.Context) ConditionFailingPeriodsOutput
}

// The minimum number of violations required within the selected lookback time window required to raise an alert.
type ConditionFailingPeriodsArgs struct {
	// The number of violations to trigger an alert. Should be smaller or equal to numberOfEvaluationPeriods. Default value is 1
	MinFailingPeriodsToAlert pulumi.Float64PtrInput `pulumi:"minFailingPeriodsToAlert"`
	// The number of aggregated lookback points. The lookback time window is calculated based on the aggregation granularity (windowSize) and the selected number of aggregated points. Default value is 1
	NumberOfEvaluationPeriods pulumi.Float64PtrInput `pulumi:"numberOfEvaluationPeriods"`
}

// Defaults sets the appropriate defaults for ConditionFailingPeriodsArgs
func (val *ConditionFailingPeriodsArgs) Defaults() *ConditionFailingPeriodsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.MinFailingPeriodsToAlert == nil {
		tmp.MinFailingPeriodsToAlert = pulumi.Float64Ptr(1.0)
	}
	if tmp.NumberOfEvaluationPeriods == nil {
		tmp.NumberOfEvaluationPeriods = pulumi.Float64Ptr(1.0)
	}
	return &tmp
}
func (ConditionFailingPeriodsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionFailingPeriods)(nil)).Elem()
}

func (i ConditionFailingPeriodsArgs) ToConditionFailingPeriodsOutput() ConditionFailingPeriodsOutput {
	return i.ToConditionFailingPeriodsOutputWithContext(context.Background())
}

func (i ConditionFailingPeriodsArgs) ToConditionFailingPeriodsOutputWithContext(ctx context.Context) ConditionFailingPeriodsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionFailingPeriodsOutput)
}

func (i ConditionFailingPeriodsArgs) ToConditionFailingPeriodsPtrOutput() ConditionFailingPeriodsPtrOutput {
	return i.ToConditionFailingPeriodsPtrOutputWithContext(context.Background())
}

func (i ConditionFailingPeriodsArgs) ToConditionFailingPeriodsPtrOutputWithContext(ctx context.Context) ConditionFailingPeriodsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionFailingPeriodsOutput).ToConditionFailingPeriodsPtrOutputWithContext(ctx)
}

// ConditionFailingPeriodsPtrInput is an input type that accepts ConditionFailingPeriodsArgs, ConditionFailingPeriodsPtr and ConditionFailingPeriodsPtrOutput values.
// You can construct a concrete instance of `ConditionFailingPeriodsPtrInput` via:
//
//	        ConditionFailingPeriodsArgs{...}
//
//	or:
//
//	        nil
type ConditionFailingPeriodsPtrInput interface {
	pulumi.Input

	ToConditionFailingPeriodsPtrOutput() ConditionFailingPeriodsPtrOutput
	ToConditionFailingPeriodsPtrOutputWithContext(context.Context) ConditionFailingPeriodsPtrOutput
}

type conditionFailingPeriodsPtrType ConditionFailingPeriodsArgs

func ConditionFailingPeriodsPtr(v *ConditionFailingPeriodsArgs) ConditionFailingPeriodsPtrInput {
	return (*conditionFailingPeriodsPtrType)(v)
}

func (*conditionFailingPeriodsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionFailingPeriods)(nil)).Elem()
}

func (i *conditionFailingPeriodsPtrType) ToConditionFailingPeriodsPtrOutput() ConditionFailingPeriodsPtrOutput {
	return i.ToConditionFailingPeriodsPtrOutputWithContext(context.Background())
}

func (i *conditionFailingPeriodsPtrType) ToConditionFailingPeriodsPtrOutputWithContext(ctx context.Context) ConditionFailingPeriodsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionFailingPeriodsPtrOutput)
}

// The minimum number of violations required within the selected lookback time window required to raise an alert.
type ConditionFailingPeriodsOutput struct{ *pulumi.OutputState }

func (ConditionFailingPeriodsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionFailingPeriods)(nil)).Elem()
}

func (o ConditionFailingPeriodsOutput) ToConditionFailingPeriodsOutput() ConditionFailingPeriodsOutput {
	return o
}

func (o ConditionFailingPeriodsOutput) ToConditionFailingPeriodsOutputWithContext(ctx context.Context) ConditionFailingPeriodsOutput {
	return o
}

func (o ConditionFailingPeriodsOutput) ToConditionFailingPeriodsPtrOutput() ConditionFailingPeriodsPtrOutput {
	return o.ToConditionFailingPeriodsPtrOutputWithContext(context.Background())
}

func (o ConditionFailingPeriodsOutput) ToConditionFailingPeriodsPtrOutputWithContext(ctx context.Context) ConditionFailingPeriodsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConditionFailingPeriods) *ConditionFailingPeriods {
		return &v
	}).(ConditionFailingPeriodsPtrOutput)
}

// The number of violations to trigger an alert. Should be smaller or equal to numberOfEvaluationPeriods. Default value is 1
func (o ConditionFailingPeriodsOutput) MinFailingPeriodsToAlert() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConditionFailingPeriods) *float64 { return v.MinFailingPeriodsToAlert }).(pulumi.Float64PtrOutput)
}

// The number of aggregated lookback points. The lookback time window is calculated based on the aggregation granularity (windowSize) and the selected number of aggregated points. Default value is 1
func (o ConditionFailingPeriodsOutput) NumberOfEvaluationPeriods() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConditionFailingPeriods) *float64 { return v.NumberOfEvaluationPeriods }).(pulumi.Float64PtrOutput)
}

type ConditionFailingPeriodsPtrOutput struct{ *pulumi.OutputState }

func (ConditionFailingPeriodsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionFailingPeriods)(nil)).Elem()
}

func (o ConditionFailingPeriodsPtrOutput) ToConditionFailingPeriodsPtrOutput() ConditionFailingPeriodsPtrOutput {
	return o
}

func (o ConditionFailingPeriodsPtrOutput) ToConditionFailingPeriodsPtrOutputWithContext(ctx context.Context) ConditionFailingPeriodsPtrOutput {
	return o
}

func (o ConditionFailingPeriodsPtrOutput) Elem() ConditionFailingPeriodsOutput {
	return o.ApplyT(func(v *ConditionFailingPeriods) ConditionFailingPeriods {
		if v != nil {
			return *v
		}
		var ret ConditionFailingPeriods
		return ret
	}).(ConditionFailingPeriodsOutput)
}

// The number of violations to trigger an alert. Should be smaller or equal to numberOfEvaluationPeriods. Default value is 1
func (o ConditionFailingPeriodsPtrOutput) MinFailingPeriodsToAlert() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConditionFailingPeriods) *float64 {
		if v == nil {
			return nil
		}
		return v.MinFailingPeriodsToAlert
	}).(pulumi.Float64PtrOutput)
}

// The number of aggregated lookback points. The lookback time window is calculated based on the aggregation granularity (windowSize) and the selected number of aggregated points. Default value is 1
func (o ConditionFailingPeriodsPtrOutput) NumberOfEvaluationPeriods() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConditionFailingPeriods) *float64 {
		if v == nil {
			return nil
		}
		return v.NumberOfEvaluationPeriods
	}).(pulumi.Float64PtrOutput)
}

// A condition of the scheduled query rule.
type ConditionResponse struct {
	// List of Dimensions conditions
	Dimensions []DimensionResponse `pulumi:"dimensions"`
	// The minimum number of violations required within the selected lookback time window required to raise an alert.
	FailingPeriods *ConditionResponseFailingPeriods `pulumi:"failingPeriods"`
	// The column containing the metric measure number.
	MetricMeasureColumn *string `pulumi:"metricMeasureColumn"`
	// The criteria operator.
	Operator string `pulumi:"operator"`
	// Log query alert
	Query *string `pulumi:"query"`
	// The column containing the resource id. The content of the column must be a uri formatted as resource id
	ResourceIdColumn *string `pulumi:"resourceIdColumn"`
	// the criteria threshold value that activates the alert.
	Threshold float64 `pulumi:"threshold"`
	// Aggregation type
	TimeAggregation string `pulumi:"timeAggregation"`
}

// Defaults sets the appropriate defaults for ConditionResponse
func (val *ConditionResponse) Defaults() *ConditionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.FailingPeriods = tmp.FailingPeriods.Defaults()

	return &tmp
}

// A condition of the scheduled query rule.
type ConditionResponseOutput struct{ *pulumi.OutputState }

func (ConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionResponse)(nil)).Elem()
}

func (o ConditionResponseOutput) ToConditionResponseOutput() ConditionResponseOutput {
	return o
}

func (o ConditionResponseOutput) ToConditionResponseOutputWithContext(ctx context.Context) ConditionResponseOutput {
	return o
}

// List of Dimensions conditions
func (o ConditionResponseOutput) Dimensions() DimensionResponseArrayOutput {
	return o.ApplyT(func(v ConditionResponse) []DimensionResponse { return v.Dimensions }).(DimensionResponseArrayOutput)
}

// The minimum number of violations required within the selected lookback time window required to raise an alert.
func (o ConditionResponseOutput) FailingPeriods() ConditionResponseFailingPeriodsPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *ConditionResponseFailingPeriods { return v.FailingPeriods }).(ConditionResponseFailingPeriodsPtrOutput)
}

// The column containing the metric measure number.
func (o ConditionResponseOutput) MetricMeasureColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.MetricMeasureColumn }).(pulumi.StringPtrOutput)
}

// The criteria operator.
func (o ConditionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ConditionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

// Log query alert
func (o ConditionResponseOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.Query }).(pulumi.StringPtrOutput)
}

// The column containing the resource id. The content of the column must be a uri formatted as resource id
func (o ConditionResponseOutput) ResourceIdColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.ResourceIdColumn }).(pulumi.StringPtrOutput)
}

// the criteria threshold value that activates the alert.
func (o ConditionResponseOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v ConditionResponse) float64 { return v.Threshold }).(pulumi.Float64Output)
}

// Aggregation type
func (o ConditionResponseOutput) TimeAggregation() pulumi.StringOutput {
	return o.ApplyT(func(v ConditionResponse) string { return v.TimeAggregation }).(pulumi.StringOutput)
}

type ConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (ConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConditionResponse)(nil)).Elem()
}

func (o ConditionResponseArrayOutput) ToConditionResponseArrayOutput() ConditionResponseArrayOutput {
	return o
}

func (o ConditionResponseArrayOutput) ToConditionResponseArrayOutputWithContext(ctx context.Context) ConditionResponseArrayOutput {
	return o
}

func (o ConditionResponseArrayOutput) Index(i pulumi.IntInput) ConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConditionResponse {
		return vs[0].([]ConditionResponse)[vs[1].(int)]
	}).(ConditionResponseOutput)
}

// The minimum number of violations required within the selected lookback time window required to raise an alert.
type ConditionResponseFailingPeriods struct {
	// The number of violations to trigger an alert. Should be smaller or equal to numberOfEvaluationPeriods. Default value is 1
	MinFailingPeriodsToAlert *float64 `pulumi:"minFailingPeriodsToAlert"`
	// The number of aggregated lookback points. The lookback time window is calculated based on the aggregation granularity (windowSize) and the selected number of aggregated points. Default value is 1
	NumberOfEvaluationPeriods *float64 `pulumi:"numberOfEvaluationPeriods"`
}

// Defaults sets the appropriate defaults for ConditionResponseFailingPeriods
func (val *ConditionResponseFailingPeriods) Defaults() *ConditionResponseFailingPeriods {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.MinFailingPeriodsToAlert == nil {
		minFailingPeriodsToAlert_ := 1.0
		tmp.MinFailingPeriodsToAlert = &minFailingPeriodsToAlert_
	}
	if tmp.NumberOfEvaluationPeriods == nil {
		numberOfEvaluationPeriods_ := 1.0
		tmp.NumberOfEvaluationPeriods = &numberOfEvaluationPeriods_
	}
	return &tmp
}

// The minimum number of violations required within the selected lookback time window required to raise an alert.
type ConditionResponseFailingPeriodsOutput struct{ *pulumi.OutputState }

func (ConditionResponseFailingPeriodsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionResponseFailingPeriods)(nil)).Elem()
}

func (o ConditionResponseFailingPeriodsOutput) ToConditionResponseFailingPeriodsOutput() ConditionResponseFailingPeriodsOutput {
	return o
}

func (o ConditionResponseFailingPeriodsOutput) ToConditionResponseFailingPeriodsOutputWithContext(ctx context.Context) ConditionResponseFailingPeriodsOutput {
	return o
}

// The number of violations to trigger an alert. Should be smaller or equal to numberOfEvaluationPeriods. Default value is 1
func (o ConditionResponseFailingPeriodsOutput) MinFailingPeriodsToAlert() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConditionResponseFailingPeriods) *float64 { return v.MinFailingPeriodsToAlert }).(pulumi.Float64PtrOutput)
}

// The number of aggregated lookback points. The lookback time window is calculated based on the aggregation granularity (windowSize) and the selected number of aggregated points. Default value is 1
func (o ConditionResponseFailingPeriodsOutput) NumberOfEvaluationPeriods() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConditionResponseFailingPeriods) *float64 { return v.NumberOfEvaluationPeriods }).(pulumi.Float64PtrOutput)
}

type ConditionResponseFailingPeriodsPtrOutput struct{ *pulumi.OutputState }

func (ConditionResponseFailingPeriodsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionResponseFailingPeriods)(nil)).Elem()
}

func (o ConditionResponseFailingPeriodsPtrOutput) ToConditionResponseFailingPeriodsPtrOutput() ConditionResponseFailingPeriodsPtrOutput {
	return o
}

func (o ConditionResponseFailingPeriodsPtrOutput) ToConditionResponseFailingPeriodsPtrOutputWithContext(ctx context.Context) ConditionResponseFailingPeriodsPtrOutput {
	return o
}

func (o ConditionResponseFailingPeriodsPtrOutput) Elem() ConditionResponseFailingPeriodsOutput {
	return o.ApplyT(func(v *ConditionResponseFailingPeriods) ConditionResponseFailingPeriods {
		if v != nil {
			return *v
		}
		var ret ConditionResponseFailingPeriods
		return ret
	}).(ConditionResponseFailingPeriodsOutput)
}

// The number of violations to trigger an alert. Should be smaller or equal to numberOfEvaluationPeriods. Default value is 1
func (o ConditionResponseFailingPeriodsPtrOutput) MinFailingPeriodsToAlert() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConditionResponseFailingPeriods) *float64 {
		if v == nil {
			return nil
		}
		return v.MinFailingPeriodsToAlert
	}).(pulumi.Float64PtrOutput)
}

// The number of aggregated lookback points. The lookback time window is calculated based on the aggregation granularity (windowSize) and the selected number of aggregated points. Default value is 1
func (o ConditionResponseFailingPeriodsPtrOutput) NumberOfEvaluationPeriods() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConditionResponseFailingPeriods) *float64 {
		if v == nil {
			return nil
		}
		return v.NumberOfEvaluationPeriods
	}).(pulumi.Float64PtrOutput)
}

// Dimension splitting and filtering definition
type Dimension struct {
	// Name of the dimension
	Name string `pulumi:"name"`
	// Operator for dimension values
	Operator string `pulumi:"operator"`
	// List of dimension values
	Values []string `pulumi:"values"`
}

// DimensionInput is an input type that accepts DimensionArgs and DimensionOutput values.
// You can construct a concrete instance of `DimensionInput` via:
//
//	DimensionArgs{...}
type DimensionInput interface {
	pulumi.Input

	ToDimensionOutput() DimensionOutput
	ToDimensionOutputWithContext(context.Context) DimensionOutput
}

// Dimension splitting and filtering definition
type DimensionArgs struct {
	// Name of the dimension
	Name pulumi.StringInput `pulumi:"name"`
	// Operator for dimension values
	Operator pulumi.StringInput `pulumi:"operator"`
	// List of dimension values
	Values pulumi.StringArrayInput `pulumi:"values"`
}

func (DimensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Dimension)(nil)).Elem()
}

func (i DimensionArgs) ToDimensionOutput() DimensionOutput {
	return i.ToDimensionOutputWithContext(context.Background())
}

func (i DimensionArgs) ToDimensionOutputWithContext(ctx context.Context) DimensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DimensionOutput)
}

// DimensionArrayInput is an input type that accepts DimensionArray and DimensionArrayOutput values.
// You can construct a concrete instance of `DimensionArrayInput` via:
//
//	DimensionArray{ DimensionArgs{...} }
type DimensionArrayInput interface {
	pulumi.Input

	ToDimensionArrayOutput() DimensionArrayOutput
	ToDimensionArrayOutputWithContext(context.Context) DimensionArrayOutput
}

type DimensionArray []DimensionInput

func (DimensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Dimension)(nil)).Elem()
}

func (i DimensionArray) ToDimensionArrayOutput() DimensionArrayOutput {
	return i.ToDimensionArrayOutputWithContext(context.Background())
}

func (i DimensionArray) ToDimensionArrayOutputWithContext(ctx context.Context) DimensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DimensionArrayOutput)
}

// Dimension splitting and filtering definition
type DimensionOutput struct{ *pulumi.OutputState }

func (DimensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Dimension)(nil)).Elem()
}

func (o DimensionOutput) ToDimensionOutput() DimensionOutput {
	return o
}

func (o DimensionOutput) ToDimensionOutputWithContext(ctx context.Context) DimensionOutput {
	return o
}

// Name of the dimension
func (o DimensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Dimension) string { return v.Name }).(pulumi.StringOutput)
}

// Operator for dimension values
func (o DimensionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v Dimension) string { return v.Operator }).(pulumi.StringOutput)
}

// List of dimension values
func (o DimensionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Dimension) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type DimensionArrayOutput struct{ *pulumi.OutputState }

func (DimensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Dimension)(nil)).Elem()
}

func (o DimensionArrayOutput) ToDimensionArrayOutput() DimensionArrayOutput {
	return o
}

func (o DimensionArrayOutput) ToDimensionArrayOutputWithContext(ctx context.Context) DimensionArrayOutput {
	return o
}

func (o DimensionArrayOutput) Index(i pulumi.IntInput) DimensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Dimension {
		return vs[0].([]Dimension)[vs[1].(int)]
	}).(DimensionOutput)
}

// Dimension splitting and filtering definition
type DimensionResponse struct {
	// Name of the dimension
	Name string `pulumi:"name"`
	// Operator for dimension values
	Operator string `pulumi:"operator"`
	// List of dimension values
	Values []string `pulumi:"values"`
}

// Dimension splitting and filtering definition
type DimensionResponseOutput struct{ *pulumi.OutputState }

func (DimensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DimensionResponse)(nil)).Elem()
}

func (o DimensionResponseOutput) ToDimensionResponseOutput() DimensionResponseOutput {
	return o
}

func (o DimensionResponseOutput) ToDimensionResponseOutputWithContext(ctx context.Context) DimensionResponseOutput {
	return o
}

// Name of the dimension
func (o DimensionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DimensionResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Operator for dimension values
func (o DimensionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v DimensionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

// List of dimension values
func (o DimensionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DimensionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type DimensionResponseArrayOutput struct{ *pulumi.OutputState }

func (DimensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DimensionResponse)(nil)).Elem()
}

func (o DimensionResponseArrayOutput) ToDimensionResponseArrayOutput() DimensionResponseArrayOutput {
	return o
}

func (o DimensionResponseArrayOutput) ToDimensionResponseArrayOutputWithContext(ctx context.Context) DimensionResponseArrayOutput {
	return o
}

func (o DimensionResponseArrayOutput) Index(i pulumi.IntInput) DimensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DimensionResponse {
		return vs[0].([]DimensionResponse)[vs[1].(int)]
	}).(DimensionResponseOutput)
}

// The rule criteria that defines the conditions of the scheduled query rule.
type ScheduledQueryRuleCriteria struct {
	// A list of conditions to evaluate against the specified scopes
	AllOf []Condition `pulumi:"allOf"`
}

// ScheduledQueryRuleCriteriaInput is an input type that accepts ScheduledQueryRuleCriteriaArgs and ScheduledQueryRuleCriteriaOutput values.
// You can construct a concrete instance of `ScheduledQueryRuleCriteriaInput` via:
//
//	ScheduledQueryRuleCriteriaArgs{...}
type ScheduledQueryRuleCriteriaInput interface {
	pulumi.Input

	ToScheduledQueryRuleCriteriaOutput() ScheduledQueryRuleCriteriaOutput
	ToScheduledQueryRuleCriteriaOutputWithContext(context.Context) ScheduledQueryRuleCriteriaOutput
}

// The rule criteria that defines the conditions of the scheduled query rule.
type ScheduledQueryRuleCriteriaArgs struct {
	// A list of conditions to evaluate against the specified scopes
	AllOf ConditionArrayInput `pulumi:"allOf"`
}

func (ScheduledQueryRuleCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRuleCriteria)(nil)).Elem()
}

func (i ScheduledQueryRuleCriteriaArgs) ToScheduledQueryRuleCriteriaOutput() ScheduledQueryRuleCriteriaOutput {
	return i.ToScheduledQueryRuleCriteriaOutputWithContext(context.Background())
}

func (i ScheduledQueryRuleCriteriaArgs) ToScheduledQueryRuleCriteriaOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRuleCriteriaOutput)
}

// The rule criteria that defines the conditions of the scheduled query rule.
type ScheduledQueryRuleCriteriaOutput struct{ *pulumi.OutputState }

func (ScheduledQueryRuleCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRuleCriteria)(nil)).Elem()
}

func (o ScheduledQueryRuleCriteriaOutput) ToScheduledQueryRuleCriteriaOutput() ScheduledQueryRuleCriteriaOutput {
	return o
}

func (o ScheduledQueryRuleCriteriaOutput) ToScheduledQueryRuleCriteriaOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaOutput {
	return o
}

// A list of conditions to evaluate against the specified scopes
func (o ScheduledQueryRuleCriteriaOutput) AllOf() ConditionArrayOutput {
	return o.ApplyT(func(v ScheduledQueryRuleCriteria) []Condition { return v.AllOf }).(ConditionArrayOutput)
}

// The rule criteria that defines the conditions of the scheduled query rule.
type ScheduledQueryRuleCriteriaResponse struct {
	// A list of conditions to evaluate against the specified scopes
	AllOf []ConditionResponse `pulumi:"allOf"`
}

// The rule criteria that defines the conditions of the scheduled query rule.
type ScheduledQueryRuleCriteriaResponseOutput struct{ *pulumi.OutputState }

func (ScheduledQueryRuleCriteriaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRuleCriteriaResponse)(nil)).Elem()
}

func (o ScheduledQueryRuleCriteriaResponseOutput) ToScheduledQueryRuleCriteriaResponseOutput() ScheduledQueryRuleCriteriaResponseOutput {
	return o
}

func (o ScheduledQueryRuleCriteriaResponseOutput) ToScheduledQueryRuleCriteriaResponseOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaResponseOutput {
	return o
}

// A list of conditions to evaluate against the specified scopes
func (o ScheduledQueryRuleCriteriaResponseOutput) AllOf() ConditionResponseArrayOutput {
	return o.ApplyT(func(v ScheduledQueryRuleCriteriaResponse) []ConditionResponse { return v.AllOf }).(ConditionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionArrayOutput{})
	pulumi.RegisterOutputType(ActionResponseOutput{})
	pulumi.RegisterOutputType(ActionResponseArrayOutput{})
	pulumi.RegisterOutputType(ConditionOutput{})
	pulumi.RegisterOutputType(ConditionArrayOutput{})
	pulumi.RegisterOutputType(ConditionFailingPeriodsOutput{})
	pulumi.RegisterOutputType(ConditionFailingPeriodsPtrOutput{})
	pulumi.RegisterOutputType(ConditionResponseOutput{})
	pulumi.RegisterOutputType(ConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(ConditionResponseFailingPeriodsOutput{})
	pulumi.RegisterOutputType(ConditionResponseFailingPeriodsPtrOutput{})
	pulumi.RegisterOutputType(DimensionOutput{})
	pulumi.RegisterOutputType(DimensionArrayOutput{})
	pulumi.RegisterOutputType(DimensionResponseOutput{})
	pulumi.RegisterOutputType(DimensionResponseArrayOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRuleCriteriaOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRuleCriteriaResponseOutput{})
}
