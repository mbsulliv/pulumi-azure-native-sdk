// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve an scheduled query rule definition.
// Azure REST API version: 2023-03-15-preview.
//
// Other available API versions: 2018-04-16, 2020-05-01-preview, 2022-08-01-preview.
func LookupScheduledQueryRule(ctx *pulumi.Context, args *LookupScheduledQueryRuleArgs, opts ...pulumi.InvokeOption) (*LookupScheduledQueryRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupScheduledQueryRuleResult
	err := ctx.Invoke("azure-native:insights:getScheduledQueryRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledQueryRuleArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the rule.
	RuleName string `pulumi:"ruleName"`
}

// The scheduled query rule resource.
type LookupScheduledQueryRuleResult struct {
	// Actions to invoke when the alert fires.
	Actions *ActionsResponse `pulumi:"actions"`
	// The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.
	AutoMitigate *bool `pulumi:"autoMitigate"`
	// The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.
	CheckWorkspaceAlertsStorageConfigured *bool `pulumi:"checkWorkspaceAlertsStorageConfigured"`
	// The api-version used when creating this alert rule
	CreatedWithApiVersion string `pulumi:"createdWithApiVersion"`
	// The rule criteria that defines the conditions of the scheduled query rule.
	Criteria ScheduledQueryRuleCriteriaResponse `pulumi:"criteria"`
	// The description of the scheduled query rule.
	Description *string `pulumi:"description"`
	// The display name of the alert rule
	DisplayName *string `pulumi:"displayName"`
	// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled bool `pulumi:"enabled"`
	// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag string `pulumi:"etag"`
	// How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.
	EvaluationFrequency *string `pulumi:"evaluationFrequency"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// True if alert rule is legacy Log Analytic rule
	IsLegacyLogAnalyticsRule bool `pulumi:"isLegacyLogAnalyticsRule"`
	// The flag which indicates whether this scheduled query rule has been configured to be stored in the customer's storage. The default is false.
	IsWorkspaceAlertsStorageConfigured bool `pulumi:"isWorkspaceAlertsStorageConfigured"`
	// Indicates the type of scheduled query rule. The default is LogAlert.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.
	MuteActionsDuration *string `pulumi:"muteActionsDuration"`
	// The name of the resource
	Name string `pulumi:"name"`
	// If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.
	OverrideQueryTimeRange *string `pulumi:"overrideQueryTimeRange"`
	// Defines the configuration for resolving fired alerts. Relevant only for rules of the kind LogAlert.
	RuleResolveConfiguration *RuleResolveConfigurationResponse `pulumi:"ruleResolveConfiguration"`
	// The list of resource id's that this scheduled query rule is scoped to.
	Scopes []string `pulumi:"scopes"`
	// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.
	Severity *float64 `pulumi:"severity"`
	// The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.
	SkipQueryValidation *bool `pulumi:"skipQueryValidation"`
	// SystemData of ScheduledQueryRule.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert
	TargetResourceTypes []string `pulumi:"targetResourceTypes"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.
	WindowSize *string `pulumi:"windowSize"`
}

func LookupScheduledQueryRuleOutput(ctx *pulumi.Context, args LookupScheduledQueryRuleOutputArgs, opts ...pulumi.InvokeOption) LookupScheduledQueryRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduledQueryRuleResult, error) {
			args := v.(LookupScheduledQueryRuleArgs)
			r, err := LookupScheduledQueryRule(ctx, &args, opts...)
			var s LookupScheduledQueryRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduledQueryRuleResultOutput)
}

type LookupScheduledQueryRuleOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the rule.
	RuleName pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupScheduledQueryRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledQueryRuleArgs)(nil)).Elem()
}

// The scheduled query rule resource.
type LookupScheduledQueryRuleResultOutput struct{ *pulumi.OutputState }

func (LookupScheduledQueryRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledQueryRuleResult)(nil)).Elem()
}

func (o LookupScheduledQueryRuleResultOutput) ToLookupScheduledQueryRuleResultOutput() LookupScheduledQueryRuleResultOutput {
	return o
}

func (o LookupScheduledQueryRuleResultOutput) ToLookupScheduledQueryRuleResultOutputWithContext(ctx context.Context) LookupScheduledQueryRuleResultOutput {
	return o
}

// Actions to invoke when the alert fires.
func (o LookupScheduledQueryRuleResultOutput) Actions() ActionsResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *ActionsResponse { return v.Actions }).(ActionsResponsePtrOutput)
}

// The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.
func (o LookupScheduledQueryRuleResultOutput) AutoMitigate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *bool { return v.AutoMitigate }).(pulumi.BoolPtrOutput)
}

// The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.
func (o LookupScheduledQueryRuleResultOutput) CheckWorkspaceAlertsStorageConfigured() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *bool { return v.CheckWorkspaceAlertsStorageConfigured }).(pulumi.BoolPtrOutput)
}

// The api-version used when creating this alert rule
func (o LookupScheduledQueryRuleResultOutput) CreatedWithApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.CreatedWithApiVersion }).(pulumi.StringOutput)
}

// The rule criteria that defines the conditions of the scheduled query rule.
func (o LookupScheduledQueryRuleResultOutput) Criteria() ScheduledQueryRuleCriteriaResponseOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) ScheduledQueryRuleCriteriaResponse { return v.Criteria }).(ScheduledQueryRuleCriteriaResponseOutput)
}

// The description of the scheduled query rule.
func (o LookupScheduledQueryRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the alert rule
func (o LookupScheduledQueryRuleResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
func (o LookupScheduledQueryRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o LookupScheduledQueryRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

// How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.
func (o LookupScheduledQueryRuleResultOutput) EvaluationFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.EvaluationFrequency }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupScheduledQueryRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the resource.
func (o LookupScheduledQueryRuleResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// True if alert rule is legacy Log Analytic rule
func (o LookupScheduledQueryRuleResultOutput) IsLegacyLogAnalyticsRule() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) bool { return v.IsLegacyLogAnalyticsRule }).(pulumi.BoolOutput)
}

// The flag which indicates whether this scheduled query rule has been configured to be stored in the customer's storage. The default is false.
func (o LookupScheduledQueryRuleResultOutput) IsWorkspaceAlertsStorageConfigured() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) bool { return v.IsWorkspaceAlertsStorageConfigured }).(pulumi.BoolOutput)
}

// Indicates the type of scheduled query rule. The default is LogAlert.
func (o LookupScheduledQueryRuleResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupScheduledQueryRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.
func (o LookupScheduledQueryRuleResultOutput) MuteActionsDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.MuteActionsDuration }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupScheduledQueryRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.
func (o LookupScheduledQueryRuleResultOutput) OverrideQueryTimeRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.OverrideQueryTimeRange }).(pulumi.StringPtrOutput)
}

// Defines the configuration for resolving fired alerts. Relevant only for rules of the kind LogAlert.
func (o LookupScheduledQueryRuleResultOutput) RuleResolveConfiguration() RuleResolveConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *RuleResolveConfigurationResponse {
		return v.RuleResolveConfiguration
	}).(RuleResolveConfigurationResponsePtrOutput)
}

// The list of resource id's that this scheduled query rule is scoped to.
func (o LookupScheduledQueryRuleResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.
func (o LookupScheduledQueryRuleResultOutput) Severity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *float64 { return v.Severity }).(pulumi.Float64PtrOutput)
}

// The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.
func (o LookupScheduledQueryRuleResultOutput) SkipQueryValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *bool { return v.SkipQueryValidation }).(pulumi.BoolPtrOutput)
}

// SystemData of ScheduledQueryRule.
func (o LookupScheduledQueryRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupScheduledQueryRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert
func (o LookupScheduledQueryRuleResultOutput) TargetResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) []string { return v.TargetResourceTypes }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupScheduledQueryRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.
func (o LookupScheduledQueryRuleResultOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduledQueryRuleResultOutput{})
}
