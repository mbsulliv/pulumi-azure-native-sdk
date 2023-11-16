// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents scheduled alert rule.
type ScheduledAlertRule struct {
	pulumi.CustomResourceState

	// The alert details override settings
	AlertDetailsOverride AlertDetailsOverrideResponsePtrOutput `pulumi:"alertDetailsOverride"`
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName pulumi.StringPtrOutput `pulumi:"alertRuleTemplateName"`
	// Dictionary of string key-value pairs of columns to be attached to the alert
	CustomDetails pulumi.StringMapOutput `pulumi:"customDetails"`
	// The description of the alert rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name for alerts created by this alert rule.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Determines whether this alert rule is enabled or disabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Array of the entity mappings of the alert rule
	EntityMappings EntityMappingResponseArrayOutput `pulumi:"entityMappings"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The event grouping settings.
	EventGroupingSettings EventGroupingSettingsResponsePtrOutput `pulumi:"eventGroupingSettings"`
	// The settings of the incidents that created from alerts triggered by this analytics rule
	IncidentConfiguration IncidentConfigurationResponsePtrOutput `pulumi:"incidentConfiguration"`
	// The kind of the alert rule
	// Expected value is 'Scheduled'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The last time that this alert rule has been modified.
	LastModifiedUtc pulumi.StringOutput `pulumi:"lastModifiedUtc"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The query that creates alerts for this rule.
	Query pulumi.StringOutput `pulumi:"query"`
	// The frequency (in ISO 8601 duration format) for this alert rule to run.
	QueryFrequency pulumi.StringOutput `pulumi:"queryFrequency"`
	// The period (in ISO 8601 duration format) that this alert rule looks at.
	QueryPeriod pulumi.StringOutput `pulumi:"queryPeriod"`
	// Array of the sentinel entity mappings of the alert rule
	SentinelEntitiesMappings SentinelEntityMappingResponseArrayOutput `pulumi:"sentinelEntitiesMappings"`
	// The severity for alerts created by this alert rule.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// The suppression (in ISO 8601 duration format) to wait since last time this alert rule been triggered.
	SuppressionDuration pulumi.StringOutput `pulumi:"suppressionDuration"`
	// Determines whether the suppression for this alert rule is enabled or disabled.
	SuppressionEnabled pulumi.BoolOutput `pulumi:"suppressionEnabled"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tactics of the alert rule
	Tactics pulumi.StringArrayOutput `pulumi:"tactics"`
	// The techniques of the alert rule
	Techniques pulumi.StringArrayOutput `pulumi:"techniques"`
	// The version of the alert rule template used to create this rule - in format <a.b.c>, where all are numbers, for example 0 <1.0.2>
	TemplateVersion pulumi.StringPtrOutput `pulumi:"templateVersion"`
	// The operation against the threshold that triggers alert rule.
	TriggerOperator pulumi.StringOutput `pulumi:"triggerOperator"`
	// The threshold triggers this alert rule.
	TriggerThreshold pulumi.IntOutput `pulumi:"triggerThreshold"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewScheduledAlertRule registers a new resource with the given unique name, arguments, and options.
func NewScheduledAlertRule(ctx *pulumi.Context,
	name string, args *ScheduledAlertRuleArgs, opts ...pulumi.ResourceOption) (*ScheduledAlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	if args.QueryFrequency == nil {
		return nil, errors.New("invalid value for required argument 'QueryFrequency'")
	}
	if args.QueryPeriod == nil {
		return nil, errors.New("invalid value for required argument 'QueryPeriod'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	if args.SuppressionDuration == nil {
		return nil, errors.New("invalid value for required argument 'SuppressionDuration'")
	}
	if args.SuppressionEnabled == nil {
		return nil, errors.New("invalid value for required argument 'SuppressionEnabled'")
	}
	if args.TriggerOperator == nil {
		return nil, errors.New("invalid value for required argument 'TriggerOperator'")
	}
	if args.TriggerThreshold == nil {
		return nil, errors.New("invalid value for required argument 'TriggerThreshold'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("Scheduled")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:ScheduledAlertRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ScheduledAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20230901preview:ScheduledAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledAlertRule gets an existing ScheduledAlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledAlertRuleState, opts ...pulumi.ResourceOption) (*ScheduledAlertRule, error) {
	var resource ScheduledAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20230901preview:ScheduledAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledAlertRule resources.
type scheduledAlertRuleState struct {
}

type ScheduledAlertRuleState struct {
}

func (ScheduledAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledAlertRuleState)(nil)).Elem()
}

type scheduledAlertRuleArgs struct {
	// The alert details override settings
	AlertDetailsOverride *AlertDetailsOverride `pulumi:"alertDetailsOverride"`
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
	EntityMappings []EntityMapping `pulumi:"entityMappings"`
	// The event grouping settings.
	EventGroupingSettings *EventGroupingSettings `pulumi:"eventGroupingSettings"`
	// The settings of the incidents that created from alerts triggered by this analytics rule
	IncidentConfiguration *IncidentConfiguration `pulumi:"incidentConfiguration"`
	// The kind of the alert rule
	// Expected value is 'Scheduled'.
	Kind string `pulumi:"kind"`
	// The query that creates alerts for this rule.
	Query string `pulumi:"query"`
	// The frequency (in ISO 8601 duration format) for this alert rule to run.
	QueryFrequency string `pulumi:"queryFrequency"`
	// The period (in ISO 8601 duration format) that this alert rule looks at.
	QueryPeriod string `pulumi:"queryPeriod"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId *string `pulumi:"ruleId"`
	// Array of the sentinel entity mappings of the alert rule
	SentinelEntitiesMappings []SentinelEntityMapping `pulumi:"sentinelEntitiesMappings"`
	// The severity for alerts created by this alert rule.
	Severity string `pulumi:"severity"`
	// The suppression (in ISO 8601 duration format) to wait since last time this alert rule been triggered.
	SuppressionDuration string `pulumi:"suppressionDuration"`
	// Determines whether the suppression for this alert rule is enabled or disabled.
	SuppressionEnabled bool `pulumi:"suppressionEnabled"`
	// The tactics of the alert rule
	Tactics []string `pulumi:"tactics"`
	// The techniques of the alert rule
	Techniques []string `pulumi:"techniques"`
	// The version of the alert rule template used to create this rule - in format <a.b.c>, where all are numbers, for example 0 <1.0.2>
	TemplateVersion *string `pulumi:"templateVersion"`
	// The operation against the threshold that triggers alert rule.
	TriggerOperator TriggerOperator `pulumi:"triggerOperator"`
	// The threshold triggers this alert rule.
	TriggerThreshold int `pulumi:"triggerThreshold"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ScheduledAlertRule resource.
type ScheduledAlertRuleArgs struct {
	// The alert details override settings
	AlertDetailsOverride AlertDetailsOverridePtrInput
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName pulumi.StringPtrInput
	// Dictionary of string key-value pairs of columns to be attached to the alert
	CustomDetails pulumi.StringMapInput
	// The description of the alert rule.
	Description pulumi.StringPtrInput
	// The display name for alerts created by this alert rule.
	DisplayName pulumi.StringInput
	// Determines whether this alert rule is enabled or disabled.
	Enabled pulumi.BoolInput
	// Array of the entity mappings of the alert rule
	EntityMappings EntityMappingArrayInput
	// The event grouping settings.
	EventGroupingSettings EventGroupingSettingsPtrInput
	// The settings of the incidents that created from alerts triggered by this analytics rule
	IncidentConfiguration IncidentConfigurationPtrInput
	// The kind of the alert rule
	// Expected value is 'Scheduled'.
	Kind pulumi.StringInput
	// The query that creates alerts for this rule.
	Query pulumi.StringInput
	// The frequency (in ISO 8601 duration format) for this alert rule to run.
	QueryFrequency pulumi.StringInput
	// The period (in ISO 8601 duration format) that this alert rule looks at.
	QueryPeriod pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Alert rule ID
	RuleId pulumi.StringPtrInput
	// Array of the sentinel entity mappings of the alert rule
	SentinelEntitiesMappings SentinelEntityMappingArrayInput
	// The severity for alerts created by this alert rule.
	Severity pulumi.StringInput
	// The suppression (in ISO 8601 duration format) to wait since last time this alert rule been triggered.
	SuppressionDuration pulumi.StringInput
	// Determines whether the suppression for this alert rule is enabled or disabled.
	SuppressionEnabled pulumi.BoolInput
	// The tactics of the alert rule
	Tactics pulumi.StringArrayInput
	// The techniques of the alert rule
	Techniques pulumi.StringArrayInput
	// The version of the alert rule template used to create this rule - in format <a.b.c>, where all are numbers, for example 0 <1.0.2>
	TemplateVersion pulumi.StringPtrInput
	// The operation against the threshold that triggers alert rule.
	TriggerOperator TriggerOperatorInput
	// The threshold triggers this alert rule.
	TriggerThreshold pulumi.IntInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (ScheduledAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledAlertRuleArgs)(nil)).Elem()
}

type ScheduledAlertRuleInput interface {
	pulumi.Input

	ToScheduledAlertRuleOutput() ScheduledAlertRuleOutput
	ToScheduledAlertRuleOutputWithContext(ctx context.Context) ScheduledAlertRuleOutput
}

func (*ScheduledAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledAlertRule)(nil)).Elem()
}

func (i *ScheduledAlertRule) ToScheduledAlertRuleOutput() ScheduledAlertRuleOutput {
	return i.ToScheduledAlertRuleOutputWithContext(context.Background())
}

func (i *ScheduledAlertRule) ToScheduledAlertRuleOutputWithContext(ctx context.Context) ScheduledAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledAlertRuleOutput)
}

type ScheduledAlertRuleOutput struct{ *pulumi.OutputState }

func (ScheduledAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledAlertRule)(nil)).Elem()
}

func (o ScheduledAlertRuleOutput) ToScheduledAlertRuleOutput() ScheduledAlertRuleOutput {
	return o
}

func (o ScheduledAlertRuleOutput) ToScheduledAlertRuleOutputWithContext(ctx context.Context) ScheduledAlertRuleOutput {
	return o
}

// The alert details override settings
func (o ScheduledAlertRuleOutput) AlertDetailsOverride() AlertDetailsOverrideResponsePtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) AlertDetailsOverrideResponsePtrOutput { return v.AlertDetailsOverride }).(AlertDetailsOverrideResponsePtrOutput)
}

// The Name of the alert rule template used to create this rule.
func (o ScheduledAlertRuleOutput) AlertRuleTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringPtrOutput { return v.AlertRuleTemplateName }).(pulumi.StringPtrOutput)
}

// Dictionary of string key-value pairs of columns to be attached to the alert
func (o ScheduledAlertRuleOutput) CustomDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringMapOutput { return v.CustomDetails }).(pulumi.StringMapOutput)
}

// The description of the alert rule.
func (o ScheduledAlertRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name for alerts created by this alert rule.
func (o ScheduledAlertRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Determines whether this alert rule is enabled or disabled.
func (o ScheduledAlertRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Array of the entity mappings of the alert rule
func (o ScheduledAlertRuleOutput) EntityMappings() EntityMappingResponseArrayOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) EntityMappingResponseArrayOutput { return v.EntityMappings }).(EntityMappingResponseArrayOutput)
}

// Etag of the azure resource
func (o ScheduledAlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The event grouping settings.
func (o ScheduledAlertRuleOutput) EventGroupingSettings() EventGroupingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) EventGroupingSettingsResponsePtrOutput { return v.EventGroupingSettings }).(EventGroupingSettingsResponsePtrOutput)
}

// The settings of the incidents that created from alerts triggered by this analytics rule
func (o ScheduledAlertRuleOutput) IncidentConfiguration() IncidentConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) IncidentConfigurationResponsePtrOutput { return v.IncidentConfiguration }).(IncidentConfigurationResponsePtrOutput)
}

// The kind of the alert rule
// Expected value is 'Scheduled'.
func (o ScheduledAlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The last time that this alert rule has been modified.
func (o ScheduledAlertRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o ScheduledAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The query that creates alerts for this rule.
func (o ScheduledAlertRuleOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

// The frequency (in ISO 8601 duration format) for this alert rule to run.
func (o ScheduledAlertRuleOutput) QueryFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.QueryFrequency }).(pulumi.StringOutput)
}

// The period (in ISO 8601 duration format) that this alert rule looks at.
func (o ScheduledAlertRuleOutput) QueryPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.QueryPeriod }).(pulumi.StringOutput)
}

// Array of the sentinel entity mappings of the alert rule
func (o ScheduledAlertRuleOutput) SentinelEntitiesMappings() SentinelEntityMappingResponseArrayOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) SentinelEntityMappingResponseArrayOutput {
		return v.SentinelEntitiesMappings
	}).(SentinelEntityMappingResponseArrayOutput)
}

// The severity for alerts created by this alert rule.
func (o ScheduledAlertRuleOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// The suppression (in ISO 8601 duration format) to wait since last time this alert rule been triggered.
func (o ScheduledAlertRuleOutput) SuppressionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.SuppressionDuration }).(pulumi.StringOutput)
}

// Determines whether the suppression for this alert rule is enabled or disabled.
func (o ScheduledAlertRuleOutput) SuppressionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.BoolOutput { return v.SuppressionEnabled }).(pulumi.BoolOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ScheduledAlertRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tactics of the alert rule
func (o ScheduledAlertRuleOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

// The techniques of the alert rule
func (o ScheduledAlertRuleOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringArrayOutput { return v.Techniques }).(pulumi.StringArrayOutput)
}

// The version of the alert rule template used to create this rule - in format <a.b.c>, where all are numbers, for example 0 <1.0.2>
func (o ScheduledAlertRuleOutput) TemplateVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringPtrOutput { return v.TemplateVersion }).(pulumi.StringPtrOutput)
}

// The operation against the threshold that triggers alert rule.
func (o ScheduledAlertRuleOutput) TriggerOperator() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.TriggerOperator }).(pulumi.StringOutput)
}

// The threshold triggers this alert rule.
func (o ScheduledAlertRuleOutput) TriggerThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.IntOutput { return v.TriggerThreshold }).(pulumi.IntOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ScheduledAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduledAlertRuleOutput{})
}
