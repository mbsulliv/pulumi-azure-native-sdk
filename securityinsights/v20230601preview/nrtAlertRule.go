// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents NRT alert rule.
type NrtAlertRule struct {
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
	// Expected value is 'NRT'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The last time that this alert rule has been modified.
	LastModifiedUtc pulumi.StringOutput `pulumi:"lastModifiedUtc"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The query that creates alerts for this rule.
	Query pulumi.StringOutput `pulumi:"query"`
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
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNrtAlertRule registers a new resource with the given unique name, arguments, and options.
func NewNrtAlertRule(ctx *pulumi.Context,
	name string, args *NrtAlertRuleArgs, opts ...pulumi.ResourceOption) (*NrtAlertRule, error) {
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
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("NRT")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:NrtAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240301:NrtAlertRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NrtAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20230601preview:NrtAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNrtAlertRule gets an existing NrtAlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNrtAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NrtAlertRuleState, opts ...pulumi.ResourceOption) (*NrtAlertRule, error) {
	var resource NrtAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20230601preview:NrtAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NrtAlertRule resources.
type nrtAlertRuleState struct {
}

type NrtAlertRuleState struct {
}

func (NrtAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*nrtAlertRuleState)(nil)).Elem()
}

type nrtAlertRuleArgs struct {
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
	// Expected value is 'NRT'.
	Kind string `pulumi:"kind"`
	// The query that creates alerts for this rule.
	Query string `pulumi:"query"`
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
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a NrtAlertRule resource.
type NrtAlertRuleArgs struct {
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
	// Expected value is 'NRT'.
	Kind pulumi.StringInput
	// The query that creates alerts for this rule.
	Query pulumi.StringInput
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
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (NrtAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nrtAlertRuleArgs)(nil)).Elem()
}

type NrtAlertRuleInput interface {
	pulumi.Input

	ToNrtAlertRuleOutput() NrtAlertRuleOutput
	ToNrtAlertRuleOutputWithContext(ctx context.Context) NrtAlertRuleOutput
}

func (*NrtAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NrtAlertRule)(nil)).Elem()
}

func (i *NrtAlertRule) ToNrtAlertRuleOutput() NrtAlertRuleOutput {
	return i.ToNrtAlertRuleOutputWithContext(context.Background())
}

func (i *NrtAlertRule) ToNrtAlertRuleOutputWithContext(ctx context.Context) NrtAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NrtAlertRuleOutput)
}

type NrtAlertRuleOutput struct{ *pulumi.OutputState }

func (NrtAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NrtAlertRule)(nil)).Elem()
}

func (o NrtAlertRuleOutput) ToNrtAlertRuleOutput() NrtAlertRuleOutput {
	return o
}

func (o NrtAlertRuleOutput) ToNrtAlertRuleOutputWithContext(ctx context.Context) NrtAlertRuleOutput {
	return o
}

// The alert details override settings
func (o NrtAlertRuleOutput) AlertDetailsOverride() AlertDetailsOverrideResponsePtrOutput {
	return o.ApplyT(func(v *NrtAlertRule) AlertDetailsOverrideResponsePtrOutput { return v.AlertDetailsOverride }).(AlertDetailsOverrideResponsePtrOutput)
}

// The Name of the alert rule template used to create this rule.
func (o NrtAlertRuleOutput) AlertRuleTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringPtrOutput { return v.AlertRuleTemplateName }).(pulumi.StringPtrOutput)
}

// Dictionary of string key-value pairs of columns to be attached to the alert
func (o NrtAlertRuleOutput) CustomDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringMapOutput { return v.CustomDetails }).(pulumi.StringMapOutput)
}

// The description of the alert rule.
func (o NrtAlertRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name for alerts created by this alert rule.
func (o NrtAlertRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Determines whether this alert rule is enabled or disabled.
func (o NrtAlertRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Array of the entity mappings of the alert rule
func (o NrtAlertRuleOutput) EntityMappings() EntityMappingResponseArrayOutput {
	return o.ApplyT(func(v *NrtAlertRule) EntityMappingResponseArrayOutput { return v.EntityMappings }).(EntityMappingResponseArrayOutput)
}

// Etag of the azure resource
func (o NrtAlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The event grouping settings.
func (o NrtAlertRuleOutput) EventGroupingSettings() EventGroupingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *NrtAlertRule) EventGroupingSettingsResponsePtrOutput { return v.EventGroupingSettings }).(EventGroupingSettingsResponsePtrOutput)
}

// The settings of the incidents that created from alerts triggered by this analytics rule
func (o NrtAlertRuleOutput) IncidentConfiguration() IncidentConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *NrtAlertRule) IncidentConfigurationResponsePtrOutput { return v.IncidentConfiguration }).(IncidentConfigurationResponsePtrOutput)
}

// The kind of the alert rule
// Expected value is 'NRT'.
func (o NrtAlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The last time that this alert rule has been modified.
func (o NrtAlertRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o NrtAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The query that creates alerts for this rule.
func (o NrtAlertRuleOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

// Array of the sentinel entity mappings of the alert rule
func (o NrtAlertRuleOutput) SentinelEntitiesMappings() SentinelEntityMappingResponseArrayOutput {
	return o.ApplyT(func(v *NrtAlertRule) SentinelEntityMappingResponseArrayOutput { return v.SentinelEntitiesMappings }).(SentinelEntityMappingResponseArrayOutput)
}

// The severity for alerts created by this alert rule.
func (o NrtAlertRuleOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// The suppression (in ISO 8601 duration format) to wait since last time this alert rule been triggered.
func (o NrtAlertRuleOutput) SuppressionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.SuppressionDuration }).(pulumi.StringOutput)
}

// Determines whether the suppression for this alert rule is enabled or disabled.
func (o NrtAlertRuleOutput) SuppressionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.BoolOutput { return v.SuppressionEnabled }).(pulumi.BoolOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o NrtAlertRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NrtAlertRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tactics of the alert rule
func (o NrtAlertRuleOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

// The techniques of the alert rule
func (o NrtAlertRuleOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringArrayOutput { return v.Techniques }).(pulumi.StringArrayOutput)
}

// The version of the alert rule template used to create this rule - in format <a.b.c>, where all are numbers, for example 0 <1.0.2>
func (o NrtAlertRuleOutput) TemplateVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringPtrOutput { return v.TemplateVersion }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NrtAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NrtAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NrtAlertRuleOutput{})
}
