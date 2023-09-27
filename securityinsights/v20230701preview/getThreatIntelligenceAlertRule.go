// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the alert rule.
func LookupThreatIntelligenceAlertRule(ctx *pulumi.Context, args *LookupThreatIntelligenceAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupThreatIntelligenceAlertRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupThreatIntelligenceAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20230701preview:getThreatIntelligenceAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupThreatIntelligenceAlertRuleArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId string `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Threat Intelligence alert rule.
type LookupThreatIntelligenceAlertRuleResult struct {
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName string `pulumi:"alertRuleTemplateName"`
	// The description of the alert rule.
	Description string `pulumi:"description"`
	// The display name for alerts created by this alert rule.
	DisplayName string `pulumi:"displayName"`
	// Determines whether this alert rule is enabled or disabled.
	Enabled bool `pulumi:"enabled"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the alert rule
	// Expected value is 'ThreatIntelligence'.
	Kind string `pulumi:"kind"`
	// The last time that this alert has been modified.
	LastModifiedUtc string `pulumi:"lastModifiedUtc"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The severity for alerts created by this alert rule.
	Severity string `pulumi:"severity"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tactics of the alert rule
	Tactics []string `pulumi:"tactics"`
	// The techniques of the alert rule
	Techniques []string `pulumi:"techniques"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupThreatIntelligenceAlertRuleOutput(ctx *pulumi.Context, args LookupThreatIntelligenceAlertRuleOutputArgs, opts ...pulumi.InvokeOption) LookupThreatIntelligenceAlertRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupThreatIntelligenceAlertRuleResult, error) {
			args := v.(LookupThreatIntelligenceAlertRuleArgs)
			r, err := LookupThreatIntelligenceAlertRule(ctx, &args, opts...)
			var s LookupThreatIntelligenceAlertRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupThreatIntelligenceAlertRuleResultOutput)
}

type LookupThreatIntelligenceAlertRuleOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId pulumi.StringInput `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupThreatIntelligenceAlertRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThreatIntelligenceAlertRuleArgs)(nil)).Elem()
}

// Represents Threat Intelligence alert rule.
type LookupThreatIntelligenceAlertRuleResultOutput struct{ *pulumi.OutputState }

func (LookupThreatIntelligenceAlertRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThreatIntelligenceAlertRuleResult)(nil)).Elem()
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) ToLookupThreatIntelligenceAlertRuleResultOutput() LookupThreatIntelligenceAlertRuleResultOutput {
	return o
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) ToLookupThreatIntelligenceAlertRuleResultOutputWithContext(ctx context.Context) LookupThreatIntelligenceAlertRuleResultOutput {
	return o
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupThreatIntelligenceAlertRuleResult] {
	return pulumix.Output[LookupThreatIntelligenceAlertRuleResult]{
		OutputState: o.OutputState,
	}
}

// The Name of the alert rule template used to create this rule.
func (o LookupThreatIntelligenceAlertRuleResultOutput) AlertRuleTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.AlertRuleTemplateName }).(pulumi.StringOutput)
}

// The description of the alert rule.
func (o LookupThreatIntelligenceAlertRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

// The display name for alerts created by this alert rule.
func (o LookupThreatIntelligenceAlertRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Determines whether this alert rule is enabled or disabled.
func (o LookupThreatIntelligenceAlertRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// Etag of the azure resource
func (o LookupThreatIntelligenceAlertRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupThreatIntelligenceAlertRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the alert rule
// Expected value is 'ThreatIntelligence'.
func (o LookupThreatIntelligenceAlertRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The last time that this alert has been modified.
func (o LookupThreatIntelligenceAlertRuleResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupThreatIntelligenceAlertRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The severity for alerts created by this alert rule.
func (o LookupThreatIntelligenceAlertRuleResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.Severity }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupThreatIntelligenceAlertRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tactics of the alert rule
func (o LookupThreatIntelligenceAlertRuleResultOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

// The techniques of the alert rule
func (o LookupThreatIntelligenceAlertRuleResultOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) []string { return v.Techniques }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupThreatIntelligenceAlertRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupThreatIntelligenceAlertRuleResultOutput{})
}
