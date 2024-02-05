// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the alert rule.
func LookupNrtAlertRule(ctx *pulumi.Context, args *LookupNrtAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupNrtAlertRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNrtAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20231201preview:getNrtAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNrtAlertRuleArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId string `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents NRT alert rule.
type LookupNrtAlertRuleResult struct {
	// The alert details override settings
	AlertDetailsOverride *AlertDetailsOverrideResponse `pulumi:"alertDetailsOverride"`
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName *string `pulumi:"alertRuleTemplateName"`
	// Dictionary of string key-value pairs of columns to be attached to the alert
	CustomDetails map[string]string `pulumi:"customDetails"`
	// The description of the alert rule.
	Description *string `pulumi:"description"`
	// The display name for alerts created by this alert rule.
	DisplayName string `pulumi:"displayName"`
	// Determines whether this alert rule is enabled or disabled.
	Enabled bool `pulumi:"enabled"`
	// Array of the entity mappings of the alert rule
	EntityMappings []EntityMappingResponse `pulumi:"entityMappings"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The event grouping settings.
	EventGroupingSettings *EventGroupingSettingsResponse `pulumi:"eventGroupingSettings"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The settings of the incidents that created from alerts triggered by this analytics rule
	IncidentConfiguration *IncidentConfigurationResponse `pulumi:"incidentConfiguration"`
	// The kind of the alert rule
	// Expected value is 'NRT'.
	Kind string `pulumi:"kind"`
	// The last time that this alert rule has been modified.
	LastModifiedUtc string `pulumi:"lastModifiedUtc"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The query that creates alerts for this rule.
	Query string `pulumi:"query"`
	// Array of the sentinel entity mappings of the alert rule
	SentinelEntitiesMappings []SentinelEntityMappingResponse `pulumi:"sentinelEntitiesMappings"`
	// The severity for alerts created by this alert rule.
	Severity string `pulumi:"severity"`
	// The sub-techniques of the alert rule
	SubTechniques []string `pulumi:"subTechniques"`
	// The suppression (in ISO 8601 duration format) to wait since last time this alert rule been triggered.
	SuppressionDuration string `pulumi:"suppressionDuration"`
	// Determines whether the suppression for this alert rule is enabled or disabled.
	SuppressionEnabled bool `pulumi:"suppressionEnabled"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tactics of the alert rule
	Tactics []string `pulumi:"tactics"`
	// The techniques of the alert rule
	Techniques []string `pulumi:"techniques"`
	// The version of the alert rule template used to create this rule - in format <a.b.c>, where all are numbers, for example 0 <1.0.2>
	TemplateVersion *string `pulumi:"templateVersion"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupNrtAlertRuleOutput(ctx *pulumi.Context, args LookupNrtAlertRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNrtAlertRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNrtAlertRuleResult, error) {
			args := v.(LookupNrtAlertRuleArgs)
			r, err := LookupNrtAlertRule(ctx, &args, opts...)
			var s LookupNrtAlertRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNrtAlertRuleResultOutput)
}

type LookupNrtAlertRuleOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId pulumi.StringInput `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupNrtAlertRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNrtAlertRuleArgs)(nil)).Elem()
}

// Represents NRT alert rule.
type LookupNrtAlertRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNrtAlertRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNrtAlertRuleResult)(nil)).Elem()
}

func (o LookupNrtAlertRuleResultOutput) ToLookupNrtAlertRuleResultOutput() LookupNrtAlertRuleResultOutput {
	return o
}

func (o LookupNrtAlertRuleResultOutput) ToLookupNrtAlertRuleResultOutputWithContext(ctx context.Context) LookupNrtAlertRuleResultOutput {
	return o
}

// The alert details override settings
func (o LookupNrtAlertRuleResultOutput) AlertDetailsOverride() AlertDetailsOverrideResponsePtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *AlertDetailsOverrideResponse { return v.AlertDetailsOverride }).(AlertDetailsOverrideResponsePtrOutput)
}

// The Name of the alert rule template used to create this rule.
func (o LookupNrtAlertRuleResultOutput) AlertRuleTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *string { return v.AlertRuleTemplateName }).(pulumi.StringPtrOutput)
}

// Dictionary of string key-value pairs of columns to be attached to the alert
func (o LookupNrtAlertRuleResultOutput) CustomDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) map[string]string { return v.CustomDetails }).(pulumi.StringMapOutput)
}

// The description of the alert rule.
func (o LookupNrtAlertRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name for alerts created by this alert rule.
func (o LookupNrtAlertRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Determines whether this alert rule is enabled or disabled.
func (o LookupNrtAlertRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// Array of the entity mappings of the alert rule
func (o LookupNrtAlertRuleResultOutput) EntityMappings() EntityMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) []EntityMappingResponse { return v.EntityMappings }).(EntityMappingResponseArrayOutput)
}

// Etag of the azure resource
func (o LookupNrtAlertRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The event grouping settings.
func (o LookupNrtAlertRuleResultOutput) EventGroupingSettings() EventGroupingSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *EventGroupingSettingsResponse { return v.EventGroupingSettings }).(EventGroupingSettingsResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupNrtAlertRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The settings of the incidents that created from alerts triggered by this analytics rule
func (o LookupNrtAlertRuleResultOutput) IncidentConfiguration() IncidentConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *IncidentConfigurationResponse { return v.IncidentConfiguration }).(IncidentConfigurationResponsePtrOutput)
}

// The kind of the alert rule
// Expected value is 'NRT'.
func (o LookupNrtAlertRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The last time that this alert rule has been modified.
func (o LookupNrtAlertRuleResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupNrtAlertRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The query that creates alerts for this rule.
func (o LookupNrtAlertRuleResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.Query }).(pulumi.StringOutput)
}

// Array of the sentinel entity mappings of the alert rule
func (o LookupNrtAlertRuleResultOutput) SentinelEntitiesMappings() SentinelEntityMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) []SentinelEntityMappingResponse { return v.SentinelEntitiesMappings }).(SentinelEntityMappingResponseArrayOutput)
}

// The severity for alerts created by this alert rule.
func (o LookupNrtAlertRuleResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.Severity }).(pulumi.StringOutput)
}

// The sub-techniques of the alert rule
func (o LookupNrtAlertRuleResultOutput) SubTechniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) []string { return v.SubTechniques }).(pulumi.StringArrayOutput)
}

// The suppression (in ISO 8601 duration format) to wait since last time this alert rule been triggered.
func (o LookupNrtAlertRuleResultOutput) SuppressionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.SuppressionDuration }).(pulumi.StringOutput)
}

// Determines whether the suppression for this alert rule is enabled or disabled.
func (o LookupNrtAlertRuleResultOutput) SuppressionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) bool { return v.SuppressionEnabled }).(pulumi.BoolOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNrtAlertRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tactics of the alert rule
func (o LookupNrtAlertRuleResultOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

// The techniques of the alert rule
func (o LookupNrtAlertRuleResultOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) []string { return v.Techniques }).(pulumi.StringArrayOutput)
}

// The version of the alert rule template used to create this rule - in format <a.b.c>, where all are numbers, for example 0 <1.0.2>
func (o LookupNrtAlertRuleResultOutput) TemplateVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *string { return v.TemplateVersion }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNrtAlertRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNrtAlertRuleResultOutput{})
}
