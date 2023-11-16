// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package policyinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// A piece of evidence supporting the compliance state set in the attestation.
type AttestationEvidence struct {
	// The description for this piece of evidence.
	Description *string `pulumi:"description"`
	// The URI location of the evidence.
	SourceUri *string `pulumi:"sourceUri"`
}

// AttestationEvidenceInput is an input type that accepts AttestationEvidenceArgs and AttestationEvidenceOutput values.
// You can construct a concrete instance of `AttestationEvidenceInput` via:
//
//	AttestationEvidenceArgs{...}
type AttestationEvidenceInput interface {
	pulumi.Input

	ToAttestationEvidenceOutput() AttestationEvidenceOutput
	ToAttestationEvidenceOutputWithContext(context.Context) AttestationEvidenceOutput
}

// A piece of evidence supporting the compliance state set in the attestation.
type AttestationEvidenceArgs struct {
	// The description for this piece of evidence.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The URI location of the evidence.
	SourceUri pulumi.StringPtrInput `pulumi:"sourceUri"`
}

func (AttestationEvidenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestationEvidence)(nil)).Elem()
}

func (i AttestationEvidenceArgs) ToAttestationEvidenceOutput() AttestationEvidenceOutput {
	return i.ToAttestationEvidenceOutputWithContext(context.Background())
}

func (i AttestationEvidenceArgs) ToAttestationEvidenceOutputWithContext(ctx context.Context) AttestationEvidenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationEvidenceOutput)
}

// AttestationEvidenceArrayInput is an input type that accepts AttestationEvidenceArray and AttestationEvidenceArrayOutput values.
// You can construct a concrete instance of `AttestationEvidenceArrayInput` via:
//
//	AttestationEvidenceArray{ AttestationEvidenceArgs{...} }
type AttestationEvidenceArrayInput interface {
	pulumi.Input

	ToAttestationEvidenceArrayOutput() AttestationEvidenceArrayOutput
	ToAttestationEvidenceArrayOutputWithContext(context.Context) AttestationEvidenceArrayOutput
}

type AttestationEvidenceArray []AttestationEvidenceInput

func (AttestationEvidenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AttestationEvidence)(nil)).Elem()
}

func (i AttestationEvidenceArray) ToAttestationEvidenceArrayOutput() AttestationEvidenceArrayOutput {
	return i.ToAttestationEvidenceArrayOutputWithContext(context.Background())
}

func (i AttestationEvidenceArray) ToAttestationEvidenceArrayOutputWithContext(ctx context.Context) AttestationEvidenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationEvidenceArrayOutput)
}

// A piece of evidence supporting the compliance state set in the attestation.
type AttestationEvidenceOutput struct{ *pulumi.OutputState }

func (AttestationEvidenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestationEvidence)(nil)).Elem()
}

func (o AttestationEvidenceOutput) ToAttestationEvidenceOutput() AttestationEvidenceOutput {
	return o
}

func (o AttestationEvidenceOutput) ToAttestationEvidenceOutputWithContext(ctx context.Context) AttestationEvidenceOutput {
	return o
}

// The description for this piece of evidence.
func (o AttestationEvidenceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttestationEvidence) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The URI location of the evidence.
func (o AttestationEvidenceOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttestationEvidence) *string { return v.SourceUri }).(pulumi.StringPtrOutput)
}

type AttestationEvidenceArrayOutput struct{ *pulumi.OutputState }

func (AttestationEvidenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AttestationEvidence)(nil)).Elem()
}

func (o AttestationEvidenceArrayOutput) ToAttestationEvidenceArrayOutput() AttestationEvidenceArrayOutput {
	return o
}

func (o AttestationEvidenceArrayOutput) ToAttestationEvidenceArrayOutputWithContext(ctx context.Context) AttestationEvidenceArrayOutput {
	return o
}

func (o AttestationEvidenceArrayOutput) Index(i pulumi.IntInput) AttestationEvidenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AttestationEvidence {
		return vs[0].([]AttestationEvidence)[vs[1].(int)]
	}).(AttestationEvidenceOutput)
}

// A piece of evidence supporting the compliance state set in the attestation.
type AttestationEvidenceResponse struct {
	// The description for this piece of evidence.
	Description *string `pulumi:"description"`
	// The URI location of the evidence.
	SourceUri *string `pulumi:"sourceUri"`
}

// A piece of evidence supporting the compliance state set in the attestation.
type AttestationEvidenceResponseOutput struct{ *pulumi.OutputState }

func (AttestationEvidenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestationEvidenceResponse)(nil)).Elem()
}

func (o AttestationEvidenceResponseOutput) ToAttestationEvidenceResponseOutput() AttestationEvidenceResponseOutput {
	return o
}

func (o AttestationEvidenceResponseOutput) ToAttestationEvidenceResponseOutputWithContext(ctx context.Context) AttestationEvidenceResponseOutput {
	return o
}

// The description for this piece of evidence.
func (o AttestationEvidenceResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttestationEvidenceResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The URI location of the evidence.
func (o AttestationEvidenceResponseOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttestationEvidenceResponse) *string { return v.SourceUri }).(pulumi.StringPtrOutput)
}

type AttestationEvidenceResponseArrayOutput struct{ *pulumi.OutputState }

func (AttestationEvidenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AttestationEvidenceResponse)(nil)).Elem()
}

func (o AttestationEvidenceResponseArrayOutput) ToAttestationEvidenceResponseArrayOutput() AttestationEvidenceResponseArrayOutput {
	return o
}

func (o AttestationEvidenceResponseArrayOutput) ToAttestationEvidenceResponseArrayOutputWithContext(ctx context.Context) AttestationEvidenceResponseArrayOutput {
	return o
}

func (o AttestationEvidenceResponseArrayOutput) Index(i pulumi.IntInput) AttestationEvidenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AttestationEvidenceResponse {
		return vs[0].([]AttestationEvidenceResponse)[vs[1].(int)]
	}).(AttestationEvidenceResponseOutput)
}

// Error definition.
type ErrorDefinitionResponse struct {
	// Additional scenario specific error details.
	AdditionalInfo []TypedErrorInfoResponse `pulumi:"additionalInfo"`
	// Service specific error code which serves as the substatus for the HTTP error code.
	Code string `pulumi:"code"`
	// Internal error details.
	Details []ErrorDefinitionResponse `pulumi:"details"`
	// Description of the error.
	Message string `pulumi:"message"`
	// The target of the error.
	Target string `pulumi:"target"`
}

// Error definition.
type ErrorDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ErrorDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDefinitionResponse)(nil)).Elem()
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponseOutput() ErrorDefinitionResponseOutput {
	return o
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponseOutputWithContext(ctx context.Context) ErrorDefinitionResponseOutput {
	return o
}

// Additional scenario specific error details.
func (o ErrorDefinitionResponseOutput) AdditionalInfo() TypedErrorInfoResponseArrayOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) []TypedErrorInfoResponse { return v.AdditionalInfo }).(TypedErrorInfoResponseArrayOutput)
}

// Service specific error code which serves as the substatus for the HTTP error code.
func (o ErrorDefinitionResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Code }).(pulumi.StringOutput)
}

// Internal error details.
func (o ErrorDefinitionResponseOutput) Details() ErrorDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) []ErrorDefinitionResponse { return v.Details }).(ErrorDefinitionResponseArrayOutput)
}

// Description of the error.
func (o ErrorDefinitionResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Message }).(pulumi.StringOutput)
}

// The target of the error.
func (o ErrorDefinitionResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDefinitionResponse)(nil)).Elem()
}

func (o ErrorDefinitionResponseArrayOutput) ToErrorDefinitionResponseArrayOutput() ErrorDefinitionResponseArrayOutput {
	return o
}

func (o ErrorDefinitionResponseArrayOutput) ToErrorDefinitionResponseArrayOutputWithContext(ctx context.Context) ErrorDefinitionResponseArrayOutput {
	return o
}

func (o ErrorDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ErrorDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDefinitionResponse {
		return vs[0].([]ErrorDefinitionResponse)[vs[1].(int)]
	}).(ErrorDefinitionResponseOutput)
}

// Details of a single deployment created by the remediation.
type RemediationDeploymentResponse struct {
	// The time at which the remediation was created.
	CreatedOn string `pulumi:"createdOn"`
	// Resource ID of the template deployment that will remediate the resource.
	DeploymentId string `pulumi:"deploymentId"`
	// Error encountered while remediated the resource.
	Error ErrorDefinitionResponse `pulumi:"error"`
	// The time at which the remediation deployment was last updated.
	LastUpdatedOn string `pulumi:"lastUpdatedOn"`
	// Resource ID of the resource that is being remediated by the deployment.
	RemediatedResourceId string `pulumi:"remediatedResourceId"`
	// Location of the resource that is being remediated.
	ResourceLocation string `pulumi:"resourceLocation"`
	// Status of the remediation deployment.
	Status string `pulumi:"status"`
}

// Details of a single deployment created by the remediation.
type RemediationDeploymentResponseOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationDeploymentResponse)(nil)).Elem()
}

func (o RemediationDeploymentResponseOutput) ToRemediationDeploymentResponseOutput() RemediationDeploymentResponseOutput {
	return o
}

func (o RemediationDeploymentResponseOutput) ToRemediationDeploymentResponseOutputWithContext(ctx context.Context) RemediationDeploymentResponseOutput {
	return o
}

// The time at which the remediation was created.
func (o RemediationDeploymentResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

// Resource ID of the template deployment that will remediate the resource.
func (o RemediationDeploymentResponseOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.DeploymentId }).(pulumi.StringOutput)
}

// Error encountered while remediated the resource.
func (o RemediationDeploymentResponseOutput) Error() ErrorDefinitionResponseOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) ErrorDefinitionResponse { return v.Error }).(ErrorDefinitionResponseOutput)
}

// The time at which the remediation deployment was last updated.
func (o RemediationDeploymentResponseOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

// Resource ID of the resource that is being remediated by the deployment.
func (o RemediationDeploymentResponseOutput) RemediatedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.RemediatedResourceId }).(pulumi.StringOutput)
}

// Location of the resource that is being remediated.
func (o RemediationDeploymentResponseOutput) ResourceLocation() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.ResourceLocation }).(pulumi.StringOutput)
}

// Status of the remediation deployment.
func (o RemediationDeploymentResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.Status }).(pulumi.StringOutput)
}

type RemediationDeploymentResponseArrayOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RemediationDeploymentResponse)(nil)).Elem()
}

func (o RemediationDeploymentResponseArrayOutput) ToRemediationDeploymentResponseArrayOutput() RemediationDeploymentResponseArrayOutput {
	return o
}

func (o RemediationDeploymentResponseArrayOutput) ToRemediationDeploymentResponseArrayOutputWithContext(ctx context.Context) RemediationDeploymentResponseArrayOutput {
	return o
}

func (o RemediationDeploymentResponseArrayOutput) Index(i pulumi.IntInput) RemediationDeploymentResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RemediationDeploymentResponse {
		return vs[0].([]RemediationDeploymentResponse)[vs[1].(int)]
	}).(RemediationDeploymentResponseOutput)
}

// The deployment status summary for all deployments created by the remediation.
type RemediationDeploymentSummaryResponse struct {
	// The number of deployments required by the remediation that have failed.
	FailedDeployments int `pulumi:"failedDeployments"`
	// The number of deployments required by the remediation that have succeeded.
	SuccessfulDeployments int `pulumi:"successfulDeployments"`
	// The number of deployments required by the remediation.
	TotalDeployments int `pulumi:"totalDeployments"`
}

// The deployment status summary for all deployments created by the remediation.
type RemediationDeploymentSummaryResponseOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentSummaryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationDeploymentSummaryResponse)(nil)).Elem()
}

func (o RemediationDeploymentSummaryResponseOutput) ToRemediationDeploymentSummaryResponseOutput() RemediationDeploymentSummaryResponseOutput {
	return o
}

func (o RemediationDeploymentSummaryResponseOutput) ToRemediationDeploymentSummaryResponseOutputWithContext(ctx context.Context) RemediationDeploymentSummaryResponseOutput {
	return o
}

// The number of deployments required by the remediation that have failed.
func (o RemediationDeploymentSummaryResponseOutput) FailedDeployments() pulumi.IntOutput {
	return o.ApplyT(func(v RemediationDeploymentSummaryResponse) int { return v.FailedDeployments }).(pulumi.IntOutput)
}

// The number of deployments required by the remediation that have succeeded.
func (o RemediationDeploymentSummaryResponseOutput) SuccessfulDeployments() pulumi.IntOutput {
	return o.ApplyT(func(v RemediationDeploymentSummaryResponse) int { return v.SuccessfulDeployments }).(pulumi.IntOutput)
}

// The number of deployments required by the remediation.
func (o RemediationDeploymentSummaryResponseOutput) TotalDeployments() pulumi.IntOutput {
	return o.ApplyT(func(v RemediationDeploymentSummaryResponse) int { return v.TotalDeployments }).(pulumi.IntOutput)
}

// The filters that will be applied to determine which resources to remediate.
type RemediationFilters struct {
	// The resource locations that will be remediated.
	Locations []string `pulumi:"locations"`
}

// RemediationFiltersInput is an input type that accepts RemediationFiltersArgs and RemediationFiltersOutput values.
// You can construct a concrete instance of `RemediationFiltersInput` via:
//
//	RemediationFiltersArgs{...}
type RemediationFiltersInput interface {
	pulumi.Input

	ToRemediationFiltersOutput() RemediationFiltersOutput
	ToRemediationFiltersOutputWithContext(context.Context) RemediationFiltersOutput
}

// The filters that will be applied to determine which resources to remediate.
type RemediationFiltersArgs struct {
	// The resource locations that will be remediated.
	Locations pulumi.StringArrayInput `pulumi:"locations"`
}

func (RemediationFiltersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationFilters)(nil)).Elem()
}

func (i RemediationFiltersArgs) ToRemediationFiltersOutput() RemediationFiltersOutput {
	return i.ToRemediationFiltersOutputWithContext(context.Background())
}

func (i RemediationFiltersArgs) ToRemediationFiltersOutputWithContext(ctx context.Context) RemediationFiltersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationFiltersOutput)
}

func (i RemediationFiltersArgs) ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput {
	return i.ToRemediationFiltersPtrOutputWithContext(context.Background())
}

func (i RemediationFiltersArgs) ToRemediationFiltersPtrOutputWithContext(ctx context.Context) RemediationFiltersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationFiltersOutput).ToRemediationFiltersPtrOutputWithContext(ctx)
}

// RemediationFiltersPtrInput is an input type that accepts RemediationFiltersArgs, RemediationFiltersPtr and RemediationFiltersPtrOutput values.
// You can construct a concrete instance of `RemediationFiltersPtrInput` via:
//
//	        RemediationFiltersArgs{...}
//
//	or:
//
//	        nil
type RemediationFiltersPtrInput interface {
	pulumi.Input

	ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput
	ToRemediationFiltersPtrOutputWithContext(context.Context) RemediationFiltersPtrOutput
}

type remediationFiltersPtrType RemediationFiltersArgs

func RemediationFiltersPtr(v *RemediationFiltersArgs) RemediationFiltersPtrInput {
	return (*remediationFiltersPtrType)(v)
}

func (*remediationFiltersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationFilters)(nil)).Elem()
}

func (i *remediationFiltersPtrType) ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput {
	return i.ToRemediationFiltersPtrOutputWithContext(context.Background())
}

func (i *remediationFiltersPtrType) ToRemediationFiltersPtrOutputWithContext(ctx context.Context) RemediationFiltersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationFiltersPtrOutput)
}

// The filters that will be applied to determine which resources to remediate.
type RemediationFiltersOutput struct{ *pulumi.OutputState }

func (RemediationFiltersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationFilters)(nil)).Elem()
}

func (o RemediationFiltersOutput) ToRemediationFiltersOutput() RemediationFiltersOutput {
	return o
}

func (o RemediationFiltersOutput) ToRemediationFiltersOutputWithContext(ctx context.Context) RemediationFiltersOutput {
	return o
}

func (o RemediationFiltersOutput) ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput {
	return o.ToRemediationFiltersPtrOutputWithContext(context.Background())
}

func (o RemediationFiltersOutput) ToRemediationFiltersPtrOutputWithContext(ctx context.Context) RemediationFiltersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemediationFilters) *RemediationFilters {
		return &v
	}).(RemediationFiltersPtrOutput)
}

// The resource locations that will be remediated.
func (o RemediationFiltersOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RemediationFilters) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

type RemediationFiltersPtrOutput struct{ *pulumi.OutputState }

func (RemediationFiltersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationFilters)(nil)).Elem()
}

func (o RemediationFiltersPtrOutput) ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput {
	return o
}

func (o RemediationFiltersPtrOutput) ToRemediationFiltersPtrOutputWithContext(ctx context.Context) RemediationFiltersPtrOutput {
	return o
}

func (o RemediationFiltersPtrOutput) Elem() RemediationFiltersOutput {
	return o.ApplyT(func(v *RemediationFilters) RemediationFilters {
		if v != nil {
			return *v
		}
		var ret RemediationFilters
		return ret
	}).(RemediationFiltersOutput)
}

// The resource locations that will be remediated.
func (o RemediationFiltersPtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RemediationFilters) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

// The filters that will be applied to determine which resources to remediate.
type RemediationFiltersResponse struct {
	// The resource locations that will be remediated.
	Locations []string `pulumi:"locations"`
}

// The filters that will be applied to determine which resources to remediate.
type RemediationFiltersResponseOutput struct{ *pulumi.OutputState }

func (RemediationFiltersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationFiltersResponse)(nil)).Elem()
}

func (o RemediationFiltersResponseOutput) ToRemediationFiltersResponseOutput() RemediationFiltersResponseOutput {
	return o
}

func (o RemediationFiltersResponseOutput) ToRemediationFiltersResponseOutputWithContext(ctx context.Context) RemediationFiltersResponseOutput {
	return o
}

// The resource locations that will be remediated.
func (o RemediationFiltersResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RemediationFiltersResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

type RemediationFiltersResponsePtrOutput struct{ *pulumi.OutputState }

func (RemediationFiltersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationFiltersResponse)(nil)).Elem()
}

func (o RemediationFiltersResponsePtrOutput) ToRemediationFiltersResponsePtrOutput() RemediationFiltersResponsePtrOutput {
	return o
}

func (o RemediationFiltersResponsePtrOutput) ToRemediationFiltersResponsePtrOutputWithContext(ctx context.Context) RemediationFiltersResponsePtrOutput {
	return o
}

func (o RemediationFiltersResponsePtrOutput) Elem() RemediationFiltersResponseOutput {
	return o.ApplyT(func(v *RemediationFiltersResponse) RemediationFiltersResponse {
		if v != nil {
			return *v
		}
		var ret RemediationFiltersResponse
		return ret
	}).(RemediationFiltersResponseOutput)
}

// The resource locations that will be remediated.
func (o RemediationFiltersResponsePtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RemediationFiltersResponse) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

// The remediation failure threshold settings
type RemediationPropertiesFailureThreshold struct {
	// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage of failed remediation operations (i.e. failed deployments) exceeds this threshold.
	Percentage *float64 `pulumi:"percentage"`
}

// RemediationPropertiesFailureThresholdInput is an input type that accepts RemediationPropertiesFailureThresholdArgs and RemediationPropertiesFailureThresholdOutput values.
// You can construct a concrete instance of `RemediationPropertiesFailureThresholdInput` via:
//
//	RemediationPropertiesFailureThresholdArgs{...}
type RemediationPropertiesFailureThresholdInput interface {
	pulumi.Input

	ToRemediationPropertiesFailureThresholdOutput() RemediationPropertiesFailureThresholdOutput
	ToRemediationPropertiesFailureThresholdOutputWithContext(context.Context) RemediationPropertiesFailureThresholdOutput
}

// The remediation failure threshold settings
type RemediationPropertiesFailureThresholdArgs struct {
	// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage of failed remediation operations (i.e. failed deployments) exceeds this threshold.
	Percentage pulumi.Float64PtrInput `pulumi:"percentage"`
}

func (RemediationPropertiesFailureThresholdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationPropertiesFailureThreshold)(nil)).Elem()
}

func (i RemediationPropertiesFailureThresholdArgs) ToRemediationPropertiesFailureThresholdOutput() RemediationPropertiesFailureThresholdOutput {
	return i.ToRemediationPropertiesFailureThresholdOutputWithContext(context.Background())
}

func (i RemediationPropertiesFailureThresholdArgs) ToRemediationPropertiesFailureThresholdOutputWithContext(ctx context.Context) RemediationPropertiesFailureThresholdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationPropertiesFailureThresholdOutput)
}

func (i RemediationPropertiesFailureThresholdArgs) ToRemediationPropertiesFailureThresholdPtrOutput() RemediationPropertiesFailureThresholdPtrOutput {
	return i.ToRemediationPropertiesFailureThresholdPtrOutputWithContext(context.Background())
}

func (i RemediationPropertiesFailureThresholdArgs) ToRemediationPropertiesFailureThresholdPtrOutputWithContext(ctx context.Context) RemediationPropertiesFailureThresholdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationPropertiesFailureThresholdOutput).ToRemediationPropertiesFailureThresholdPtrOutputWithContext(ctx)
}

// RemediationPropertiesFailureThresholdPtrInput is an input type that accepts RemediationPropertiesFailureThresholdArgs, RemediationPropertiesFailureThresholdPtr and RemediationPropertiesFailureThresholdPtrOutput values.
// You can construct a concrete instance of `RemediationPropertiesFailureThresholdPtrInput` via:
//
//	        RemediationPropertiesFailureThresholdArgs{...}
//
//	or:
//
//	        nil
type RemediationPropertiesFailureThresholdPtrInput interface {
	pulumi.Input

	ToRemediationPropertiesFailureThresholdPtrOutput() RemediationPropertiesFailureThresholdPtrOutput
	ToRemediationPropertiesFailureThresholdPtrOutputWithContext(context.Context) RemediationPropertiesFailureThresholdPtrOutput
}

type remediationPropertiesFailureThresholdPtrType RemediationPropertiesFailureThresholdArgs

func RemediationPropertiesFailureThresholdPtr(v *RemediationPropertiesFailureThresholdArgs) RemediationPropertiesFailureThresholdPtrInput {
	return (*remediationPropertiesFailureThresholdPtrType)(v)
}

func (*remediationPropertiesFailureThresholdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationPropertiesFailureThreshold)(nil)).Elem()
}

func (i *remediationPropertiesFailureThresholdPtrType) ToRemediationPropertiesFailureThresholdPtrOutput() RemediationPropertiesFailureThresholdPtrOutput {
	return i.ToRemediationPropertiesFailureThresholdPtrOutputWithContext(context.Background())
}

func (i *remediationPropertiesFailureThresholdPtrType) ToRemediationPropertiesFailureThresholdPtrOutputWithContext(ctx context.Context) RemediationPropertiesFailureThresholdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationPropertiesFailureThresholdPtrOutput)
}

// The remediation failure threshold settings
type RemediationPropertiesFailureThresholdOutput struct{ *pulumi.OutputState }

func (RemediationPropertiesFailureThresholdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationPropertiesFailureThreshold)(nil)).Elem()
}

func (o RemediationPropertiesFailureThresholdOutput) ToRemediationPropertiesFailureThresholdOutput() RemediationPropertiesFailureThresholdOutput {
	return o
}

func (o RemediationPropertiesFailureThresholdOutput) ToRemediationPropertiesFailureThresholdOutputWithContext(ctx context.Context) RemediationPropertiesFailureThresholdOutput {
	return o
}

func (o RemediationPropertiesFailureThresholdOutput) ToRemediationPropertiesFailureThresholdPtrOutput() RemediationPropertiesFailureThresholdPtrOutput {
	return o.ToRemediationPropertiesFailureThresholdPtrOutputWithContext(context.Background())
}

func (o RemediationPropertiesFailureThresholdOutput) ToRemediationPropertiesFailureThresholdPtrOutputWithContext(ctx context.Context) RemediationPropertiesFailureThresholdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemediationPropertiesFailureThreshold) *RemediationPropertiesFailureThreshold {
		return &v
	}).(RemediationPropertiesFailureThresholdPtrOutput)
}

// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage of failed remediation operations (i.e. failed deployments) exceeds this threshold.
func (o RemediationPropertiesFailureThresholdOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RemediationPropertiesFailureThreshold) *float64 { return v.Percentage }).(pulumi.Float64PtrOutput)
}

type RemediationPropertiesFailureThresholdPtrOutput struct{ *pulumi.OutputState }

func (RemediationPropertiesFailureThresholdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationPropertiesFailureThreshold)(nil)).Elem()
}

func (o RemediationPropertiesFailureThresholdPtrOutput) ToRemediationPropertiesFailureThresholdPtrOutput() RemediationPropertiesFailureThresholdPtrOutput {
	return o
}

func (o RemediationPropertiesFailureThresholdPtrOutput) ToRemediationPropertiesFailureThresholdPtrOutputWithContext(ctx context.Context) RemediationPropertiesFailureThresholdPtrOutput {
	return o
}

func (o RemediationPropertiesFailureThresholdPtrOutput) Elem() RemediationPropertiesFailureThresholdOutput {
	return o.ApplyT(func(v *RemediationPropertiesFailureThreshold) RemediationPropertiesFailureThreshold {
		if v != nil {
			return *v
		}
		var ret RemediationPropertiesFailureThreshold
		return ret
	}).(RemediationPropertiesFailureThresholdOutput)
}

// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage of failed remediation operations (i.e. failed deployments) exceeds this threshold.
func (o RemediationPropertiesFailureThresholdPtrOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *RemediationPropertiesFailureThreshold) *float64 {
		if v == nil {
			return nil
		}
		return v.Percentage
	}).(pulumi.Float64PtrOutput)
}

// The remediation failure threshold settings
type RemediationPropertiesResponseFailureThreshold struct {
	// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage of failed remediation operations (i.e. failed deployments) exceeds this threshold.
	Percentage *float64 `pulumi:"percentage"`
}

// The remediation failure threshold settings
type RemediationPropertiesResponseFailureThresholdOutput struct{ *pulumi.OutputState }

func (RemediationPropertiesResponseFailureThresholdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationPropertiesResponseFailureThreshold)(nil)).Elem()
}

func (o RemediationPropertiesResponseFailureThresholdOutput) ToRemediationPropertiesResponseFailureThresholdOutput() RemediationPropertiesResponseFailureThresholdOutput {
	return o
}

func (o RemediationPropertiesResponseFailureThresholdOutput) ToRemediationPropertiesResponseFailureThresholdOutputWithContext(ctx context.Context) RemediationPropertiesResponseFailureThresholdOutput {
	return o
}

// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage of failed remediation operations (i.e. failed deployments) exceeds this threshold.
func (o RemediationPropertiesResponseFailureThresholdOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RemediationPropertiesResponseFailureThreshold) *float64 { return v.Percentage }).(pulumi.Float64PtrOutput)
}

type RemediationPropertiesResponseFailureThresholdPtrOutput struct{ *pulumi.OutputState }

func (RemediationPropertiesResponseFailureThresholdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationPropertiesResponseFailureThreshold)(nil)).Elem()
}

func (o RemediationPropertiesResponseFailureThresholdPtrOutput) ToRemediationPropertiesResponseFailureThresholdPtrOutput() RemediationPropertiesResponseFailureThresholdPtrOutput {
	return o
}

func (o RemediationPropertiesResponseFailureThresholdPtrOutput) ToRemediationPropertiesResponseFailureThresholdPtrOutputWithContext(ctx context.Context) RemediationPropertiesResponseFailureThresholdPtrOutput {
	return o
}

func (o RemediationPropertiesResponseFailureThresholdPtrOutput) Elem() RemediationPropertiesResponseFailureThresholdOutput {
	return o.ApplyT(func(v *RemediationPropertiesResponseFailureThreshold) RemediationPropertiesResponseFailureThreshold {
		if v != nil {
			return *v
		}
		var ret RemediationPropertiesResponseFailureThreshold
		return ret
	}).(RemediationPropertiesResponseFailureThresholdOutput)
}

// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage of failed remediation operations (i.e. failed deployments) exceeds this threshold.
func (o RemediationPropertiesResponseFailureThresholdPtrOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *RemediationPropertiesResponseFailureThreshold) *float64 {
		if v == nil {
			return nil
		}
		return v.Percentage
	}).(pulumi.Float64PtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

// Scenario specific error details.
type TypedErrorInfoResponse struct {
	// The scenario specific error details.
	Info interface{} `pulumi:"info"`
	// The type of included error details.
	Type string `pulumi:"type"`
}

// Scenario specific error details.
type TypedErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (TypedErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TypedErrorInfoResponse)(nil)).Elem()
}

func (o TypedErrorInfoResponseOutput) ToTypedErrorInfoResponseOutput() TypedErrorInfoResponseOutput {
	return o
}

func (o TypedErrorInfoResponseOutput) ToTypedErrorInfoResponseOutputWithContext(ctx context.Context) TypedErrorInfoResponseOutput {
	return o
}

// The scenario specific error details.
func (o TypedErrorInfoResponseOutput) Info() pulumi.AnyOutput {
	return o.ApplyT(func(v TypedErrorInfoResponse) interface{} { return v.Info }).(pulumi.AnyOutput)
}

// The type of included error details.
func (o TypedErrorInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TypedErrorInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type TypedErrorInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (TypedErrorInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TypedErrorInfoResponse)(nil)).Elem()
}

func (o TypedErrorInfoResponseArrayOutput) ToTypedErrorInfoResponseArrayOutput() TypedErrorInfoResponseArrayOutput {
	return o
}

func (o TypedErrorInfoResponseArrayOutput) ToTypedErrorInfoResponseArrayOutputWithContext(ctx context.Context) TypedErrorInfoResponseArrayOutput {
	return o
}

func (o TypedErrorInfoResponseArrayOutput) Index(i pulumi.IntInput) TypedErrorInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TypedErrorInfoResponse {
		return vs[0].([]TypedErrorInfoResponse)[vs[1].(int)]
	}).(TypedErrorInfoResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AttestationEvidenceOutput{})
	pulumi.RegisterOutputType(AttestationEvidenceArrayOutput{})
	pulumi.RegisterOutputType(AttestationEvidenceResponseOutput{})
	pulumi.RegisterOutputType(AttestationEvidenceResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ErrorDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentResponseOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentResponseArrayOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentSummaryResponseOutput{})
	pulumi.RegisterOutputType(RemediationFiltersOutput{})
	pulumi.RegisterOutputType(RemediationFiltersPtrOutput{})
	pulumi.RegisterOutputType(RemediationFiltersResponseOutput{})
	pulumi.RegisterOutputType(RemediationFiltersResponsePtrOutput{})
	pulumi.RegisterOutputType(RemediationPropertiesFailureThresholdOutput{})
	pulumi.RegisterOutputType(RemediationPropertiesFailureThresholdPtrOutput{})
	pulumi.RegisterOutputType(RemediationPropertiesResponseFailureThresholdOutput{})
	pulumi.RegisterOutputType(RemediationPropertiesResponseFailureThresholdPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TypedErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(TypedErrorInfoResponseArrayOutput{})
}
