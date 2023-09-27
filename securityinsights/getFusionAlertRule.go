// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the alert rule.
// Azure REST API version: 2023-02-01.
func LookupFusionAlertRule(ctx *pulumi.Context, args *LookupFusionAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupFusionAlertRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFusionAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights:getFusionAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFusionAlertRuleArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId string `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Fusion alert rule.
type LookupFusionAlertRuleResult struct {
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
	// Expected value is 'Fusion'.
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

func LookupFusionAlertRuleOutput(ctx *pulumi.Context, args LookupFusionAlertRuleOutputArgs, opts ...pulumi.InvokeOption) LookupFusionAlertRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFusionAlertRuleResult, error) {
			args := v.(LookupFusionAlertRuleArgs)
			r, err := LookupFusionAlertRule(ctx, &args, opts...)
			var s LookupFusionAlertRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFusionAlertRuleResultOutput)
}

type LookupFusionAlertRuleOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId pulumi.StringInput `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupFusionAlertRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFusionAlertRuleArgs)(nil)).Elem()
}

// Represents Fusion alert rule.
type LookupFusionAlertRuleResultOutput struct{ *pulumi.OutputState }

func (LookupFusionAlertRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFusionAlertRuleResult)(nil)).Elem()
}

func (o LookupFusionAlertRuleResultOutput) ToLookupFusionAlertRuleResultOutput() LookupFusionAlertRuleResultOutput {
	return o
}

func (o LookupFusionAlertRuleResultOutput) ToLookupFusionAlertRuleResultOutputWithContext(ctx context.Context) LookupFusionAlertRuleResultOutput {
	return o
}

func (o LookupFusionAlertRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFusionAlertRuleResult] {
	return pulumix.Output[LookupFusionAlertRuleResult]{
		OutputState: o.OutputState,
	}
}

// The Name of the alert rule template used to create this rule.
func (o LookupFusionAlertRuleResultOutput) AlertRuleTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.AlertRuleTemplateName }).(pulumi.StringOutput)
}

// The description of the alert rule.
func (o LookupFusionAlertRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

// The display name for alerts created by this alert rule.
func (o LookupFusionAlertRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Determines whether this alert rule is enabled or disabled.
func (o LookupFusionAlertRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// Etag of the azure resource
func (o LookupFusionAlertRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupFusionAlertRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the alert rule
// Expected value is 'Fusion'.
func (o LookupFusionAlertRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The last time that this alert has been modified.
func (o LookupFusionAlertRuleResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupFusionAlertRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The severity for alerts created by this alert rule.
func (o LookupFusionAlertRuleResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.Severity }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFusionAlertRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tactics of the alert rule
func (o LookupFusionAlertRuleResultOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

// The techniques of the alert rule
func (o LookupFusionAlertRuleResultOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) []string { return v.Techniques }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFusionAlertRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFusionAlertRuleResultOutput{})
}
