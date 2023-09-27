// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170426

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a Prediction in the hub.
func LookupPrediction(ctx *pulumi.Context, args *LookupPredictionArgs, opts ...pulumi.InvokeOption) (*LookupPredictionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPredictionResult
	err := ctx.Invoke("azure-native:customerinsights/v20170426:getPrediction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPredictionArgs struct {
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the Prediction.
	PredictionName string `pulumi:"predictionName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The prediction resource format.
type LookupPredictionResult struct {
	// Whether do auto analyze.
	AutoAnalyze bool `pulumi:"autoAnalyze"`
	// Description of the prediction.
	Description map[string]string `pulumi:"description"`
	// Display name of the prediction.
	DisplayName map[string]string `pulumi:"displayName"`
	// The prediction grades.
	Grades []PredictionResponseGrades `pulumi:"grades"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Interaction types involved in the prediction.
	InvolvedInteractionTypes []string `pulumi:"involvedInteractionTypes"`
	// KPI types involved in the prediction.
	InvolvedKpiTypes []string `pulumi:"involvedKpiTypes"`
	// Relationships involved in the prediction.
	InvolvedRelationships []string `pulumi:"involvedRelationships"`
	// Definition of the link mapping of prediction.
	Mappings PredictionResponseMappings `pulumi:"mappings"`
	// Resource name.
	Name string `pulumi:"name"`
	// Negative outcome expression.
	NegativeOutcomeExpression string `pulumi:"negativeOutcomeExpression"`
	// Positive outcome expression.
	PositiveOutcomeExpression string `pulumi:"positiveOutcomeExpression"`
	// Name of the prediction.
	PredictionName *string `pulumi:"predictionName"`
	// Primary profile type.
	PrimaryProfileType string `pulumi:"primaryProfileType"`
	// Provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Scope expression.
	ScopeExpression string `pulumi:"scopeExpression"`
	// Score label.
	ScoreLabel string `pulumi:"scoreLabel"`
	// System generated entities.
	SystemGeneratedEntities PredictionResponseSystemGeneratedEntities `pulumi:"systemGeneratedEntities"`
	// The hub name.
	TenantId string `pulumi:"tenantId"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupPredictionOutput(ctx *pulumi.Context, args LookupPredictionOutputArgs, opts ...pulumi.InvokeOption) LookupPredictionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPredictionResult, error) {
			args := v.(LookupPredictionArgs)
			r, err := LookupPrediction(ctx, &args, opts...)
			var s LookupPredictionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPredictionResultOutput)
}

type LookupPredictionOutputArgs struct {
	// The name of the hub.
	HubName pulumi.StringInput `pulumi:"hubName"`
	// The name of the Prediction.
	PredictionName pulumi.StringInput `pulumi:"predictionName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPredictionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPredictionArgs)(nil)).Elem()
}

// The prediction resource format.
type LookupPredictionResultOutput struct{ *pulumi.OutputState }

func (LookupPredictionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPredictionResult)(nil)).Elem()
}

func (o LookupPredictionResultOutput) ToLookupPredictionResultOutput() LookupPredictionResultOutput {
	return o
}

func (o LookupPredictionResultOutput) ToLookupPredictionResultOutputWithContext(ctx context.Context) LookupPredictionResultOutput {
	return o
}

func (o LookupPredictionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPredictionResult] {
	return pulumix.Output[LookupPredictionResult]{
		OutputState: o.OutputState,
	}
}

// Whether do auto analyze.
func (o LookupPredictionResultOutput) AutoAnalyze() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPredictionResult) bool { return v.AutoAnalyze }).(pulumi.BoolOutput)
}

// Description of the prediction.
func (o LookupPredictionResultOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPredictionResult) map[string]string { return v.Description }).(pulumi.StringMapOutput)
}

// Display name of the prediction.
func (o LookupPredictionResultOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPredictionResult) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

// The prediction grades.
func (o LookupPredictionResultOutput) Grades() PredictionResponseGradesArrayOutput {
	return o.ApplyT(func(v LookupPredictionResult) []PredictionResponseGrades { return v.Grades }).(PredictionResponseGradesArrayOutput)
}

// Resource ID.
func (o LookupPredictionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Interaction types involved in the prediction.
func (o LookupPredictionResultOutput) InvolvedInteractionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPredictionResult) []string { return v.InvolvedInteractionTypes }).(pulumi.StringArrayOutput)
}

// KPI types involved in the prediction.
func (o LookupPredictionResultOutput) InvolvedKpiTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPredictionResult) []string { return v.InvolvedKpiTypes }).(pulumi.StringArrayOutput)
}

// Relationships involved in the prediction.
func (o LookupPredictionResultOutput) InvolvedRelationships() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPredictionResult) []string { return v.InvolvedRelationships }).(pulumi.StringArrayOutput)
}

// Definition of the link mapping of prediction.
func (o LookupPredictionResultOutput) Mappings() PredictionResponseMappingsOutput {
	return o.ApplyT(func(v LookupPredictionResult) PredictionResponseMappings { return v.Mappings }).(PredictionResponseMappingsOutput)
}

// Resource name.
func (o LookupPredictionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Negative outcome expression.
func (o LookupPredictionResultOutput) NegativeOutcomeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.NegativeOutcomeExpression }).(pulumi.StringOutput)
}

// Positive outcome expression.
func (o LookupPredictionResultOutput) PositiveOutcomeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.PositiveOutcomeExpression }).(pulumi.StringOutput)
}

// Name of the prediction.
func (o LookupPredictionResultOutput) PredictionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPredictionResult) *string { return v.PredictionName }).(pulumi.StringPtrOutput)
}

// Primary profile type.
func (o LookupPredictionResultOutput) PrimaryProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.PrimaryProfileType }).(pulumi.StringOutput)
}

// Provisioning state.
func (o LookupPredictionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Scope expression.
func (o LookupPredictionResultOutput) ScopeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.ScopeExpression }).(pulumi.StringOutput)
}

// Score label.
func (o LookupPredictionResultOutput) ScoreLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.ScoreLabel }).(pulumi.StringOutput)
}

// System generated entities.
func (o LookupPredictionResultOutput) SystemGeneratedEntities() PredictionResponseSystemGeneratedEntitiesOutput {
	return o.ApplyT(func(v LookupPredictionResult) PredictionResponseSystemGeneratedEntities {
		return v.SystemGeneratedEntities
	}).(PredictionResponseSystemGeneratedEntitiesOutput)
}

// The hub name.
func (o LookupPredictionResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupPredictionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPredictionResultOutput{})
}
