// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an existing remediation at subscription scope.
func LookupRemediationAtSubscription(ctx *pulumi.Context, args *LookupRemediationAtSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupRemediationAtSubscriptionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRemediationAtSubscriptionResult
	err := ctx.Invoke("azure-native:policyinsights/v20211001:getRemediationAtSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemediationAtSubscriptionArgs struct {
	// The name of the remediation.
	RemediationName string `pulumi:"remediationName"`
}

// The remediation definition.
type LookupRemediationAtSubscriptionResult struct {
	// The remediation correlation Id. Can be used to find events related to the remediation in the activity log.
	CorrelationId string `pulumi:"correlationId"`
	// The time at which the remediation was created.
	CreatedOn string `pulumi:"createdOn"`
	// The deployment status summary for all deployments created by the remediation.
	DeploymentStatus RemediationDeploymentSummaryResponse `pulumi:"deploymentStatus"`
	// The remediation failure threshold settings
	FailureThreshold *RemediationPropertiesResponseFailureThreshold `pulumi:"failureThreshold"`
	// The filters that will be applied to determine which resources to remediate.
	Filters *RemediationFiltersResponse `pulumi:"filters"`
	// The ID of the remediation.
	Id string `pulumi:"id"`
	// The time at which the remediation was last updated.
	LastUpdatedOn string `pulumi:"lastUpdatedOn"`
	// The name of the remediation.
	Name string `pulumi:"name"`
	// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation. If not provided, the default parallel deployments value is used.
	ParallelDeployments *int `pulumi:"parallelDeployments"`
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId *string `pulumi:"policyAssignmentId"`
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
	// The status of the remediation. This refers to the entire remediation task, not individual deployments. Allowed values are Evaluating, Canceled, Cancelling, Failed, Complete, or Succeeded.
	ProvisioningState string `pulumi:"provisioningState"`
	// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource count is used.
	ResourceCount *int `pulumi:"resourceCount"`
	// The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
	ResourceDiscoveryMode *string `pulumi:"resourceDiscoveryMode"`
	// The remediation status message. Provides additional details regarding the state of the remediation.
	StatusMessage string `pulumi:"statusMessage"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the remediation.
	Type string `pulumi:"type"`
}

func LookupRemediationAtSubscriptionOutput(ctx *pulumi.Context, args LookupRemediationAtSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupRemediationAtSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemediationAtSubscriptionResult, error) {
			args := v.(LookupRemediationAtSubscriptionArgs)
			r, err := LookupRemediationAtSubscription(ctx, &args, opts...)
			var s LookupRemediationAtSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemediationAtSubscriptionResultOutput)
}

type LookupRemediationAtSubscriptionOutputArgs struct {
	// The name of the remediation.
	RemediationName pulumi.StringInput `pulumi:"remediationName"`
}

func (LookupRemediationAtSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtSubscriptionArgs)(nil)).Elem()
}

// The remediation definition.
type LookupRemediationAtSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupRemediationAtSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtSubscriptionResult)(nil)).Elem()
}

func (o LookupRemediationAtSubscriptionResultOutput) ToLookupRemediationAtSubscriptionResultOutput() LookupRemediationAtSubscriptionResultOutput {
	return o
}

func (o LookupRemediationAtSubscriptionResultOutput) ToLookupRemediationAtSubscriptionResultOutputWithContext(ctx context.Context) LookupRemediationAtSubscriptionResultOutput {
	return o
}

func (o LookupRemediationAtSubscriptionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRemediationAtSubscriptionResult] {
	return pulumix.Output[LookupRemediationAtSubscriptionResult]{
		OutputState: o.OutputState,
	}
}

// The remediation correlation Id. Can be used to find events related to the remediation in the activity log.
func (o LookupRemediationAtSubscriptionResultOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.CorrelationId }).(pulumi.StringOutput)
}

// The time at which the remediation was created.
func (o LookupRemediationAtSubscriptionResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

// The deployment status summary for all deployments created by the remediation.
func (o LookupRemediationAtSubscriptionResultOutput) DeploymentStatus() RemediationDeploymentSummaryResponseOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) RemediationDeploymentSummaryResponse {
		return v.DeploymentStatus
	}).(RemediationDeploymentSummaryResponseOutput)
}

// The remediation failure threshold settings
func (o LookupRemediationAtSubscriptionResultOutput) FailureThreshold() RemediationPropertiesResponseFailureThresholdPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *RemediationPropertiesResponseFailureThreshold {
		return v.FailureThreshold
	}).(RemediationPropertiesResponseFailureThresholdPtrOutput)
}

// The filters that will be applied to determine which resources to remediate.
func (o LookupRemediationAtSubscriptionResultOutput) Filters() RemediationFiltersResponsePtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *RemediationFiltersResponse { return v.Filters }).(RemediationFiltersResponsePtrOutput)
}

// The ID of the remediation.
func (o LookupRemediationAtSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The time at which the remediation was last updated.
func (o LookupRemediationAtSubscriptionResultOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

// The name of the remediation.
func (o LookupRemediationAtSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation. If not provided, the default parallel deployments value is used.
func (o LookupRemediationAtSubscriptionResultOutput) ParallelDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *int { return v.ParallelDeployments }).(pulumi.IntPtrOutput)
}

// The resource ID of the policy assignment that should be remediated.
func (o LookupRemediationAtSubscriptionResultOutput) PolicyAssignmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *string { return v.PolicyAssignmentId }).(pulumi.StringPtrOutput)
}

// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
func (o LookupRemediationAtSubscriptionResultOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

// The status of the remediation. This refers to the entire remediation task, not individual deployments. Allowed values are Evaluating, Canceled, Cancelling, Failed, Complete, or Succeeded.
func (o LookupRemediationAtSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource count is used.
func (o LookupRemediationAtSubscriptionResultOutput) ResourceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *int { return v.ResourceCount }).(pulumi.IntPtrOutput)
}

// The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
func (o LookupRemediationAtSubscriptionResultOutput) ResourceDiscoveryMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) *string { return v.ResourceDiscoveryMode }).(pulumi.StringPtrOutput)
}

// The remediation status message. Provides additional details regarding the state of the remediation.
func (o LookupRemediationAtSubscriptionResultOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.StatusMessage }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRemediationAtSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the remediation.
func (o LookupRemediationAtSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemediationAtSubscriptionResultOutput{})
}
