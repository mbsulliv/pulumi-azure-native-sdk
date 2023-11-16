// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package policyinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an existing remediation at management group scope.
// Azure REST API version: 2021-10-01.
//
// Other available API versions: 2018-07-01-preview.
func LookupRemediationAtManagementGroup(ctx *pulumi.Context, args *LookupRemediationAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupRemediationAtManagementGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRemediationAtManagementGroupResult
	err := ctx.Invoke("azure-native:policyinsights:getRemediationAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemediationAtManagementGroupArgs struct {
	// Management group ID.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// The namespace for Microsoft Management RP; only "Microsoft.Management" is allowed.
	ManagementGroupsNamespace string `pulumi:"managementGroupsNamespace"`
	// The name of the remediation.
	RemediationName string `pulumi:"remediationName"`
}

// The remediation definition.
type LookupRemediationAtManagementGroupResult struct {
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

func LookupRemediationAtManagementGroupOutput(ctx *pulumi.Context, args LookupRemediationAtManagementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupRemediationAtManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemediationAtManagementGroupResult, error) {
			args := v.(LookupRemediationAtManagementGroupArgs)
			r, err := LookupRemediationAtManagementGroup(ctx, &args, opts...)
			var s LookupRemediationAtManagementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemediationAtManagementGroupResultOutput)
}

type LookupRemediationAtManagementGroupOutputArgs struct {
	// Management group ID.
	ManagementGroupId pulumi.StringInput `pulumi:"managementGroupId"`
	// The namespace for Microsoft Management RP; only "Microsoft.Management" is allowed.
	ManagementGroupsNamespace pulumi.StringInput `pulumi:"managementGroupsNamespace"`
	// The name of the remediation.
	RemediationName pulumi.StringInput `pulumi:"remediationName"`
}

func (LookupRemediationAtManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtManagementGroupArgs)(nil)).Elem()
}

// The remediation definition.
type LookupRemediationAtManagementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupRemediationAtManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationAtManagementGroupResult)(nil)).Elem()
}

func (o LookupRemediationAtManagementGroupResultOutput) ToLookupRemediationAtManagementGroupResultOutput() LookupRemediationAtManagementGroupResultOutput {
	return o
}

func (o LookupRemediationAtManagementGroupResultOutput) ToLookupRemediationAtManagementGroupResultOutputWithContext(ctx context.Context) LookupRemediationAtManagementGroupResultOutput {
	return o
}

// The remediation correlation Id. Can be used to find events related to the remediation in the activity log.
func (o LookupRemediationAtManagementGroupResultOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.CorrelationId }).(pulumi.StringOutput)
}

// The time at which the remediation was created.
func (o LookupRemediationAtManagementGroupResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

// The deployment status summary for all deployments created by the remediation.
func (o LookupRemediationAtManagementGroupResultOutput) DeploymentStatus() RemediationDeploymentSummaryResponseOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) RemediationDeploymentSummaryResponse {
		return v.DeploymentStatus
	}).(RemediationDeploymentSummaryResponseOutput)
}

// The remediation failure threshold settings
func (o LookupRemediationAtManagementGroupResultOutput) FailureThreshold() RemediationPropertiesResponseFailureThresholdPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) *RemediationPropertiesResponseFailureThreshold {
		return v.FailureThreshold
	}).(RemediationPropertiesResponseFailureThresholdPtrOutput)
}

// The filters that will be applied to determine which resources to remediate.
func (o LookupRemediationAtManagementGroupResultOutput) Filters() RemediationFiltersResponsePtrOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) *RemediationFiltersResponse { return v.Filters }).(RemediationFiltersResponsePtrOutput)
}

// The ID of the remediation.
func (o LookupRemediationAtManagementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The time at which the remediation was last updated.
func (o LookupRemediationAtManagementGroupResultOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

// The name of the remediation.
func (o LookupRemediationAtManagementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation. If not provided, the default parallel deployments value is used.
func (o LookupRemediationAtManagementGroupResultOutput) ParallelDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) *int { return v.ParallelDeployments }).(pulumi.IntPtrOutput)
}

// The resource ID of the policy assignment that should be remediated.
func (o LookupRemediationAtManagementGroupResultOutput) PolicyAssignmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) *string { return v.PolicyAssignmentId }).(pulumi.StringPtrOutput)
}

// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
func (o LookupRemediationAtManagementGroupResultOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

// The status of the remediation. This refers to the entire remediation task, not individual deployments. Allowed values are Evaluating, Canceled, Cancelling, Failed, Complete, or Succeeded.
func (o LookupRemediationAtManagementGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource count is used.
func (o LookupRemediationAtManagementGroupResultOutput) ResourceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) *int { return v.ResourceCount }).(pulumi.IntPtrOutput)
}

// The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
func (o LookupRemediationAtManagementGroupResultOutput) ResourceDiscoveryMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) *string { return v.ResourceDiscoveryMode }).(pulumi.StringPtrOutput)
}

// The remediation status message. Provides additional details regarding the state of the remediation.
func (o LookupRemediationAtManagementGroupResultOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.StatusMessage }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRemediationAtManagementGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the remediation.
func (o LookupRemediationAtManagementGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemediationAtManagementGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemediationAtManagementGroupResultOutput{})
}
