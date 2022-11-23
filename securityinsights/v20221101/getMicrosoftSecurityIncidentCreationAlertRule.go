// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents MicrosoftSecurityIncidentCreation rule.
func LookupMicrosoftSecurityIncidentCreationAlertRule(ctx *pulumi.Context, args *LookupMicrosoftSecurityIncidentCreationAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupMicrosoftSecurityIncidentCreationAlertRuleResult, error) {
	var rv LookupMicrosoftSecurityIncidentCreationAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20221101:getMicrosoftSecurityIncidentCreationAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMicrosoftSecurityIncidentCreationAlertRuleArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId string `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents MicrosoftSecurityIncidentCreation rule.
type LookupMicrosoftSecurityIncidentCreationAlertRuleResult struct {
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName *string `pulumi:"alertRuleTemplateName"`
	// The description of the alert rule.
	Description *string `pulumi:"description"`
	// The display name for alerts created by this alert rule.
	DisplayName string `pulumi:"displayName"`
	// the alerts' displayNames on which the cases will not be generated
	DisplayNamesExcludeFilter []string `pulumi:"displayNamesExcludeFilter"`
	// the alerts' displayNames on which the cases will be generated
	DisplayNamesFilter []string `pulumi:"displayNamesFilter"`
	// Determines whether this alert rule is enabled or disabled.
	Enabled bool `pulumi:"enabled"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the alert rule
	// Expected value is 'MicrosoftSecurityIncidentCreation'.
	Kind string `pulumi:"kind"`
	// The last time that this alert has been modified.
	LastModifiedUtc string `pulumi:"lastModifiedUtc"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The alerts' productName on which the cases will be generated
	ProductFilter string `pulumi:"productFilter"`
	// the alerts' severities on which the cases will be generated
	SeveritiesFilter []string `pulumi:"severitiesFilter"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupMicrosoftSecurityIncidentCreationAlertRuleOutput(ctx *pulumi.Context, args LookupMicrosoftSecurityIncidentCreationAlertRuleOutputArgs, opts ...pulumi.InvokeOption) LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMicrosoftSecurityIncidentCreationAlertRuleResult, error) {
			args := v.(LookupMicrosoftSecurityIncidentCreationAlertRuleArgs)
			r, err := LookupMicrosoftSecurityIncidentCreationAlertRule(ctx, &args, opts...)
			var s LookupMicrosoftSecurityIncidentCreationAlertRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput)
}

type LookupMicrosoftSecurityIncidentCreationAlertRuleOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId pulumi.StringInput `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMicrosoftSecurityIncidentCreationAlertRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMicrosoftSecurityIncidentCreationAlertRuleArgs)(nil)).Elem()
}

// Represents MicrosoftSecurityIncidentCreation rule.
type LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput struct{ *pulumi.OutputState }

func (LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMicrosoftSecurityIncidentCreationAlertRuleResult)(nil)).Elem()
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) ToLookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput() LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput {
	return o
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) ToLookupMicrosoftSecurityIncidentCreationAlertRuleResultOutputWithContext(ctx context.Context) LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput {
	return o
}

// The Name of the alert rule template used to create this rule.
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) AlertRuleTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) *string { return v.AlertRuleTemplateName }).(pulumi.StringPtrOutput)
}

// The description of the alert rule.
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name for alerts created by this alert rule.
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// the alerts' displayNames on which the cases will not be generated
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) DisplayNamesExcludeFilter() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) []string {
		return v.DisplayNamesExcludeFilter
	}).(pulumi.StringArrayOutput)
}

// the alerts' displayNames on which the cases will be generated
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) DisplayNamesFilter() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) []string { return v.DisplayNamesFilter }).(pulumi.StringArrayOutput)
}

// Determines whether this alert rule is enabled or disabled.
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// Etag of the azure resource
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the alert rule
// Expected value is 'MicrosoftSecurityIncidentCreation'.
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The last time that this alert has been modified.
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The alerts' productName on which the cases will be generated
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) ProductFilter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.ProductFilter }).(pulumi.StringOutput)
}

// the alerts' severities on which the cases will be generated
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) SeveritiesFilter() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) []string { return v.SeveritiesFilter }).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput{})
}
