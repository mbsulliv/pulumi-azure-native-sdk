// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get metadata information on an assessment type in a specific subscription
// Azure REST API version: 2021-06-01.
func LookupAssessmentMetadataInSubscription(ctx *pulumi.Context, args *LookupAssessmentMetadataInSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupAssessmentMetadataInSubscriptionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAssessmentMetadataInSubscriptionResult
	err := ctx.Invoke("azure-native:security:getAssessmentMetadataInSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssessmentMetadataInSubscriptionArgs struct {
	// The Assessment Key - Unique key for the assessment type
	AssessmentMetadataName string `pulumi:"assessmentMetadataName"`
}

// Security assessment metadata response
type LookupAssessmentMetadataInSubscriptionResult struct {
	// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
	AssessmentType string   `pulumi:"assessmentType"`
	Categories     []string `pulumi:"categories"`
	// Human readable description of the assessment
	Description *string `pulumi:"description"`
	// User friendly display name of the assessment
	DisplayName string `pulumi:"displayName"`
	// Resource Id
	Id string `pulumi:"id"`
	// The implementation effort required to remediate this assessment
	ImplementationEffort *string `pulumi:"implementationEffort"`
	// Resource name
	Name string `pulumi:"name"`
	// Describes the partner that created the assessment
	PartnerData            *SecurityAssessmentMetadataPartnerDataResponse `pulumi:"partnerData"`
	PlannedDeprecationDate *string                                        `pulumi:"plannedDeprecationDate"`
	// Azure resource ID of the policy definition that turns this assessment calculation on
	PolicyDefinitionId string `pulumi:"policyDefinitionId"`
	// True if this assessment is in preview release status
	Preview      *bool                                                             `pulumi:"preview"`
	PublishDates *SecurityAssessmentMetadataPropertiesResponseResponsePublishDates `pulumi:"publishDates"`
	// Human readable description of what you should do to mitigate this security issue
	RemediationDescription *string `pulumi:"remediationDescription"`
	// The severity level of the assessment
	Severity   string   `pulumi:"severity"`
	Tactics    []string `pulumi:"tactics"`
	Techniques []string `pulumi:"techniques"`
	Threats    []string `pulumi:"threats"`
	// Resource type
	Type string `pulumi:"type"`
	// The user impact of the assessment
	UserImpact *string `pulumi:"userImpact"`
}

func LookupAssessmentMetadataInSubscriptionOutput(ctx *pulumi.Context, args LookupAssessmentMetadataInSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupAssessmentMetadataInSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssessmentMetadataInSubscriptionResult, error) {
			args := v.(LookupAssessmentMetadataInSubscriptionArgs)
			r, err := LookupAssessmentMetadataInSubscription(ctx, &args, opts...)
			var s LookupAssessmentMetadataInSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssessmentMetadataInSubscriptionResultOutput)
}

type LookupAssessmentMetadataInSubscriptionOutputArgs struct {
	// The Assessment Key - Unique key for the assessment type
	AssessmentMetadataName pulumi.StringInput `pulumi:"assessmentMetadataName"`
}

func (LookupAssessmentMetadataInSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssessmentMetadataInSubscriptionArgs)(nil)).Elem()
}

// Security assessment metadata response
type LookupAssessmentMetadataInSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupAssessmentMetadataInSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssessmentMetadataInSubscriptionResult)(nil)).Elem()
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) ToLookupAssessmentMetadataInSubscriptionResultOutput() LookupAssessmentMetadataInSubscriptionResultOutput {
	return o
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) ToLookupAssessmentMetadataInSubscriptionResultOutputWithContext(ctx context.Context) LookupAssessmentMetadataInSubscriptionResultOutput {
	return o
}

// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
func (o LookupAssessmentMetadataInSubscriptionResultOutput) AssessmentType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.AssessmentType }).(pulumi.StringOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

// Human readable description of the assessment
func (o LookupAssessmentMetadataInSubscriptionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// User friendly display name of the assessment
func (o LookupAssessmentMetadataInSubscriptionResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource Id
func (o LookupAssessmentMetadataInSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The implementation effort required to remediate this assessment
func (o LookupAssessmentMetadataInSubscriptionResultOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *string { return v.ImplementationEffort }).(pulumi.StringPtrOutput)
}

// Resource name
func (o LookupAssessmentMetadataInSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Describes the partner that created the assessment
func (o LookupAssessmentMetadataInSubscriptionResultOutput) PartnerData() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *SecurityAssessmentMetadataPartnerDataResponse {
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataResponsePtrOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) PlannedDeprecationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *string { return v.PlannedDeprecationDate }).(pulumi.StringPtrOutput)
}

// Azure resource ID of the policy definition that turns this assessment calculation on
func (o LookupAssessmentMetadataInSubscriptionResultOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

// True if this assessment is in preview release status
func (o LookupAssessmentMetadataInSubscriptionResultOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *bool { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) PublishDates() SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *SecurityAssessmentMetadataPropertiesResponseResponsePublishDates {
		return v.PublishDates
	}).(SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput)
}

// Human readable description of what you should do to mitigate this security issue
func (o LookupAssessmentMetadataInSubscriptionResultOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *string { return v.RemediationDescription }).(pulumi.StringPtrOutput)
}

// The severity level of the assessment
func (o LookupAssessmentMetadataInSubscriptionResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.Severity }).(pulumi.StringOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) []string { return v.Techniques }).(pulumi.StringArrayOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) []string { return v.Threats }).(pulumi.StringArrayOutput)
}

// Resource type
func (o LookupAssessmentMetadataInSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

// The user impact of the assessment
func (o LookupAssessmentMetadataInSubscriptionResultOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *string { return v.UserImpact }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssessmentMetadataInSubscriptionResultOutput{})
}
