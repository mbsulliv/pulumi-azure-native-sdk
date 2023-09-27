// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get dismiss rule, with name: {alertsSuppressionRuleName}, for the given subscription
func LookupAlertsSuppressionRule(ctx *pulumi.Context, args *LookupAlertsSuppressionRuleArgs, opts ...pulumi.InvokeOption) (*LookupAlertsSuppressionRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAlertsSuppressionRuleResult
	err := ctx.Invoke("azure-native:security/v20190101preview:getAlertsSuppressionRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAlertsSuppressionRuleArgs struct {
	// The unique name of the suppression alert rule
	AlertsSuppressionRuleName string `pulumi:"alertsSuppressionRuleName"`
}

// Describes the suppression rule
type LookupAlertsSuppressionRuleResult struct {
	// Type of the alert to automatically suppress. For all alert types, use '*'
	AlertType string `pulumi:"alertType"`
	// Any comment regarding the rule
	Comment *string `pulumi:"comment"`
	// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
	ExpirationDateUtc *string `pulumi:"expirationDateUtc"`
	// Resource Id
	Id string `pulumi:"id"`
	// The last time this rule was modified
	LastModifiedUtc string `pulumi:"lastModifiedUtc"`
	// Resource name
	Name string `pulumi:"name"`
	// The reason for dismissing the alert
	Reason string `pulumi:"reason"`
	// Possible states of the rule
	State string `pulumi:"state"`
	// The suppression conditions
	SuppressionAlertsScope *SuppressionAlertsScopeResponse `pulumi:"suppressionAlertsScope"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupAlertsSuppressionRuleOutput(ctx *pulumi.Context, args LookupAlertsSuppressionRuleOutputArgs, opts ...pulumi.InvokeOption) LookupAlertsSuppressionRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAlertsSuppressionRuleResult, error) {
			args := v.(LookupAlertsSuppressionRuleArgs)
			r, err := LookupAlertsSuppressionRule(ctx, &args, opts...)
			var s LookupAlertsSuppressionRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAlertsSuppressionRuleResultOutput)
}

type LookupAlertsSuppressionRuleOutputArgs struct {
	// The unique name of the suppression alert rule
	AlertsSuppressionRuleName pulumi.StringInput `pulumi:"alertsSuppressionRuleName"`
}

func (LookupAlertsSuppressionRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlertsSuppressionRuleArgs)(nil)).Elem()
}

// Describes the suppression rule
type LookupAlertsSuppressionRuleResultOutput struct{ *pulumi.OutputState }

func (LookupAlertsSuppressionRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlertsSuppressionRuleResult)(nil)).Elem()
}

func (o LookupAlertsSuppressionRuleResultOutput) ToLookupAlertsSuppressionRuleResultOutput() LookupAlertsSuppressionRuleResultOutput {
	return o
}

func (o LookupAlertsSuppressionRuleResultOutput) ToLookupAlertsSuppressionRuleResultOutputWithContext(ctx context.Context) LookupAlertsSuppressionRuleResultOutput {
	return o
}

func (o LookupAlertsSuppressionRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAlertsSuppressionRuleResult] {
	return pulumix.Output[LookupAlertsSuppressionRuleResult]{
		OutputState: o.OutputState,
	}
}

// Type of the alert to automatically suppress. For all alert types, use '*'
func (o LookupAlertsSuppressionRuleResultOutput) AlertType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.AlertType }).(pulumi.StringOutput)
}

// Any comment regarding the rule
func (o LookupAlertsSuppressionRuleResultOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) *string { return v.Comment }).(pulumi.StringPtrOutput)
}

// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
func (o LookupAlertsSuppressionRuleResultOutput) ExpirationDateUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) *string { return v.ExpirationDateUtc }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupAlertsSuppressionRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The last time this rule was modified
func (o LookupAlertsSuppressionRuleResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// Resource name
func (o LookupAlertsSuppressionRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The reason for dismissing the alert
func (o LookupAlertsSuppressionRuleResultOutput) Reason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.Reason }).(pulumi.StringOutput)
}

// Possible states of the rule
func (o LookupAlertsSuppressionRuleResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.State }).(pulumi.StringOutput)
}

// The suppression conditions
func (o LookupAlertsSuppressionRuleResultOutput) SuppressionAlertsScope() SuppressionAlertsScopeResponsePtrOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) *SuppressionAlertsScopeResponse {
		return v.SuppressionAlertsScope
	}).(SuppressionAlertsScopeResponsePtrOutput)
}

// Resource type
func (o LookupAlertsSuppressionRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlertsSuppressionRuleResultOutput{})
}
