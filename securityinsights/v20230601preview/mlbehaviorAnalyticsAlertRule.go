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

// Represents MLBehaviorAnalytics alert rule.
type MLBehaviorAnalyticsAlertRule struct {
	pulumi.CustomResourceState

	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName pulumi.StringOutput `pulumi:"alertRuleTemplateName"`
	// The description of the alert rule.
	Description pulumi.StringOutput `pulumi:"description"`
	// The display name for alerts created by this alert rule.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Determines whether this alert rule is enabled or disabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the alert rule
	// Expected value is 'MLBehaviorAnalytics'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The last time that this alert rule has been modified.
	LastModifiedUtc pulumi.StringOutput `pulumi:"lastModifiedUtc"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The severity for alerts created by this alert rule.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tactics of the alert rule
	Tactics pulumi.StringArrayOutput `pulumi:"tactics"`
	// The techniques of the alert rule
	Techniques pulumi.StringArrayOutput `pulumi:"techniques"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMLBehaviorAnalyticsAlertRule registers a new resource with the given unique name, arguments, and options.
func NewMLBehaviorAnalyticsAlertRule(ctx *pulumi.Context,
	name string, args *MLBehaviorAnalyticsAlertRuleArgs, opts ...pulumi.ResourceOption) (*MLBehaviorAnalyticsAlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlertRuleTemplateName == nil {
		return nil, errors.New("invalid value for required argument 'AlertRuleTemplateName'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("MLBehaviorAnalytics")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:MLBehaviorAnalyticsAlertRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MLBehaviorAnalyticsAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20230601preview:MLBehaviorAnalyticsAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMLBehaviorAnalyticsAlertRule gets an existing MLBehaviorAnalyticsAlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMLBehaviorAnalyticsAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MLBehaviorAnalyticsAlertRuleState, opts ...pulumi.ResourceOption) (*MLBehaviorAnalyticsAlertRule, error) {
	var resource MLBehaviorAnalyticsAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20230601preview:MLBehaviorAnalyticsAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MLBehaviorAnalyticsAlertRule resources.
type mlbehaviorAnalyticsAlertRuleState struct {
}

type MLBehaviorAnalyticsAlertRuleState struct {
}

func (MLBehaviorAnalyticsAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*mlbehaviorAnalyticsAlertRuleState)(nil)).Elem()
}

type mlbehaviorAnalyticsAlertRuleArgs struct {
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName string `pulumi:"alertRuleTemplateName"`
	// Determines whether this alert rule is enabled or disabled.
	Enabled bool `pulumi:"enabled"`
	// The kind of the alert rule
	// Expected value is 'MLBehaviorAnalytics'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId *string `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a MLBehaviorAnalyticsAlertRule resource.
type MLBehaviorAnalyticsAlertRuleArgs struct {
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName pulumi.StringInput
	// Determines whether this alert rule is enabled or disabled.
	Enabled pulumi.BoolInput
	// The kind of the alert rule
	// Expected value is 'MLBehaviorAnalytics'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Alert rule ID
	RuleId pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (MLBehaviorAnalyticsAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mlbehaviorAnalyticsAlertRuleArgs)(nil)).Elem()
}

type MLBehaviorAnalyticsAlertRuleInput interface {
	pulumi.Input

	ToMLBehaviorAnalyticsAlertRuleOutput() MLBehaviorAnalyticsAlertRuleOutput
	ToMLBehaviorAnalyticsAlertRuleOutputWithContext(ctx context.Context) MLBehaviorAnalyticsAlertRuleOutput
}

func (*MLBehaviorAnalyticsAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**MLBehaviorAnalyticsAlertRule)(nil)).Elem()
}

func (i *MLBehaviorAnalyticsAlertRule) ToMLBehaviorAnalyticsAlertRuleOutput() MLBehaviorAnalyticsAlertRuleOutput {
	return i.ToMLBehaviorAnalyticsAlertRuleOutputWithContext(context.Background())
}

func (i *MLBehaviorAnalyticsAlertRule) ToMLBehaviorAnalyticsAlertRuleOutputWithContext(ctx context.Context) MLBehaviorAnalyticsAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MLBehaviorAnalyticsAlertRuleOutput)
}

type MLBehaviorAnalyticsAlertRuleOutput struct{ *pulumi.OutputState }

func (MLBehaviorAnalyticsAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MLBehaviorAnalyticsAlertRule)(nil)).Elem()
}

func (o MLBehaviorAnalyticsAlertRuleOutput) ToMLBehaviorAnalyticsAlertRuleOutput() MLBehaviorAnalyticsAlertRuleOutput {
	return o
}

func (o MLBehaviorAnalyticsAlertRuleOutput) ToMLBehaviorAnalyticsAlertRuleOutputWithContext(ctx context.Context) MLBehaviorAnalyticsAlertRuleOutput {
	return o
}

// The Name of the alert rule template used to create this rule.
func (o MLBehaviorAnalyticsAlertRuleOutput) AlertRuleTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.AlertRuleTemplateName }).(pulumi.StringOutput)
}

// The description of the alert rule.
func (o MLBehaviorAnalyticsAlertRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The display name for alerts created by this alert rule.
func (o MLBehaviorAnalyticsAlertRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Determines whether this alert rule is enabled or disabled.
func (o MLBehaviorAnalyticsAlertRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Etag of the azure resource
func (o MLBehaviorAnalyticsAlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the alert rule
// Expected value is 'MLBehaviorAnalytics'.
func (o MLBehaviorAnalyticsAlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The last time that this alert rule has been modified.
func (o MLBehaviorAnalyticsAlertRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o MLBehaviorAnalyticsAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The severity for alerts created by this alert rule.
func (o MLBehaviorAnalyticsAlertRuleOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MLBehaviorAnalyticsAlertRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tactics of the alert rule
func (o MLBehaviorAnalyticsAlertRuleOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

// The techniques of the alert rule
func (o MLBehaviorAnalyticsAlertRuleOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringArrayOutput { return v.Techniques }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MLBehaviorAnalyticsAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MLBehaviorAnalyticsAlertRuleOutput{})
}
