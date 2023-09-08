// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a specific governanceAssignment for the requested scope by AssignmentKey
// Azure REST API version: 2022-01-01-preview.
func LookupGovernanceAssignment(ctx *pulumi.Context, args *LookupGovernanceAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupGovernanceAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGovernanceAssignmentResult
	err := ctx.Invoke("azure-native:security:getGovernanceAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGovernanceAssignmentArgs struct {
	// The Assessment Key - A unique key for the assessment type
	AssessmentName string `pulumi:"assessmentName"`
	// The governance assignment key - the assessment key of the required governance assignment
	AssignmentKey string `pulumi:"assignmentKey"`
	// The scope of the Governance assignments. Valid scopes are: subscription (format: 'subscriptions/{subscriptionId}'), or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
	Scope string `pulumi:"scope"`
}

// Governance assignment over a given scope
type LookupGovernanceAssignmentResult struct {
	// The additional data for the governance assignment - e.g. links to ticket (optional), see example
	AdditionalData *GovernanceAssignmentAdditionalDataResponse `pulumi:"additionalData"`
	// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
	GovernanceEmailNotification *GovernanceEmailNotificationResponse `pulumi:"governanceEmailNotification"`
	// Resource Id
	Id string `pulumi:"id"`
	// Defines whether there is a grace period on the governance assignment
	IsGracePeriod *bool `pulumi:"isGracePeriod"`
	// Resource name
	Name string `pulumi:"name"`
	// The Owner for the governance assignment - e.g. user@contoso.com - see example
	Owner *string `pulumi:"owner"`
	// The remediation due-date - after this date Secure Score will be affected (in case of  active grace-period)
	RemediationDueDate string `pulumi:"remediationDueDate"`
	// The ETA (estimated time of arrival) for remediation (optional), see example
	RemediationEta *RemediationEtaResponse `pulumi:"remediationEta"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupGovernanceAssignmentOutput(ctx *pulumi.Context, args LookupGovernanceAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupGovernanceAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGovernanceAssignmentResult, error) {
			args := v.(LookupGovernanceAssignmentArgs)
			r, err := LookupGovernanceAssignment(ctx, &args, opts...)
			var s LookupGovernanceAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGovernanceAssignmentResultOutput)
}

type LookupGovernanceAssignmentOutputArgs struct {
	// The Assessment Key - A unique key for the assessment type
	AssessmentName pulumi.StringInput `pulumi:"assessmentName"`
	// The governance assignment key - the assessment key of the required governance assignment
	AssignmentKey pulumi.StringInput `pulumi:"assignmentKey"`
	// The scope of the Governance assignments. Valid scopes are: subscription (format: 'subscriptions/{subscriptionId}'), or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupGovernanceAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGovernanceAssignmentArgs)(nil)).Elem()
}

// Governance assignment over a given scope
type LookupGovernanceAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupGovernanceAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGovernanceAssignmentResult)(nil)).Elem()
}

func (o LookupGovernanceAssignmentResultOutput) ToLookupGovernanceAssignmentResultOutput() LookupGovernanceAssignmentResultOutput {
	return o
}

func (o LookupGovernanceAssignmentResultOutput) ToLookupGovernanceAssignmentResultOutputWithContext(ctx context.Context) LookupGovernanceAssignmentResultOutput {
	return o
}

// The additional data for the governance assignment - e.g. links to ticket (optional), see example
func (o LookupGovernanceAssignmentResultOutput) AdditionalData() GovernanceAssignmentAdditionalDataResponsePtrOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) *GovernanceAssignmentAdditionalDataResponse {
		return v.AdditionalData
	}).(GovernanceAssignmentAdditionalDataResponsePtrOutput)
}

// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
func (o LookupGovernanceAssignmentResultOutput) GovernanceEmailNotification() GovernanceEmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) *GovernanceEmailNotificationResponse {
		return v.GovernanceEmailNotification
	}).(GovernanceEmailNotificationResponsePtrOutput)
}

// Resource Id
func (o LookupGovernanceAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Defines whether there is a grace period on the governance assignment
func (o LookupGovernanceAssignmentResultOutput) IsGracePeriod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) *bool { return v.IsGracePeriod }).(pulumi.BoolPtrOutput)
}

// Resource name
func (o LookupGovernanceAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Owner for the governance assignment - e.g. user@contoso.com - see example
func (o LookupGovernanceAssignmentResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

// The remediation due-date - after this date Secure Score will be affected (in case of  active grace-period)
func (o LookupGovernanceAssignmentResultOutput) RemediationDueDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) string { return v.RemediationDueDate }).(pulumi.StringOutput)
}

// The ETA (estimated time of arrival) for remediation (optional), see example
func (o LookupGovernanceAssignmentResultOutput) RemediationEta() RemediationEtaResponsePtrOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) *RemediationEtaResponse { return v.RemediationEta }).(RemediationEtaResponsePtrOutput)
}

// Resource type
func (o LookupGovernanceAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGovernanceAssignmentResultOutput{})
}