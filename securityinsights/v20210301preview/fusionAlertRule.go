// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Fusion alert rule.
type FusionAlertRule struct {
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
	// Expected value is 'Fusion'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The last time that this alert has been modified.
	LastModifiedUtc pulumi.StringOutput `pulumi:"lastModifiedUtc"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The severity for alerts created by this alert rule.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tactics of the alert rule
	Tactics pulumi.StringArrayOutput `pulumi:"tactics"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFusionAlertRule registers a new resource with the given unique name, arguments, and options.
func NewFusionAlertRule(ctx *pulumi.Context,
	name string, args *FusionAlertRuleArgs, opts ...pulumi.ResourceOption) (*FusionAlertRule, error) {
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
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("Fusion")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240301:FusionAlertRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FusionAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:FusionAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAlertRule gets an existing FusionAlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAlertRuleState, opts ...pulumi.ResourceOption) (*FusionAlertRule, error) {
	var resource FusionAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:FusionAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAlertRule resources.
type fusionAlertRuleState struct {
}

type FusionAlertRuleState struct {
}

func (FusionAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAlertRuleState)(nil)).Elem()
}

type fusionAlertRuleArgs struct {
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName string `pulumi:"alertRuleTemplateName"`
	// Determines whether this alert rule is enabled or disabled.
	Enabled bool `pulumi:"enabled"`
	// The kind of the alert rule
	// Expected value is 'Fusion'.
	Kind string `pulumi:"kind"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId *string `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a FusionAlertRule resource.
type FusionAlertRuleArgs struct {
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName pulumi.StringInput
	// Determines whether this alert rule is enabled or disabled.
	Enabled pulumi.BoolInput
	// The kind of the alert rule
	// Expected value is 'Fusion'.
	Kind pulumi.StringInput
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Alert rule ID
	RuleId pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (FusionAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAlertRuleArgs)(nil)).Elem()
}

type FusionAlertRuleInput interface {
	pulumi.Input

	ToFusionAlertRuleOutput() FusionAlertRuleOutput
	ToFusionAlertRuleOutputWithContext(ctx context.Context) FusionAlertRuleOutput
}

func (*FusionAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAlertRule)(nil)).Elem()
}

func (i *FusionAlertRule) ToFusionAlertRuleOutput() FusionAlertRuleOutput {
	return i.ToFusionAlertRuleOutputWithContext(context.Background())
}

func (i *FusionAlertRule) ToFusionAlertRuleOutputWithContext(ctx context.Context) FusionAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAlertRuleOutput)
}

type FusionAlertRuleOutput struct{ *pulumi.OutputState }

func (FusionAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAlertRule)(nil)).Elem()
}

func (o FusionAlertRuleOutput) ToFusionAlertRuleOutput() FusionAlertRuleOutput {
	return o
}

func (o FusionAlertRuleOutput) ToFusionAlertRuleOutputWithContext(ctx context.Context) FusionAlertRuleOutput {
	return o
}

// The Name of the alert rule template used to create this rule.
func (o FusionAlertRuleOutput) AlertRuleTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.AlertRuleTemplateName }).(pulumi.StringOutput)
}

// The description of the alert rule.
func (o FusionAlertRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The display name for alerts created by this alert rule.
func (o FusionAlertRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Determines whether this alert rule is enabled or disabled.
func (o FusionAlertRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Etag of the azure resource
func (o FusionAlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the alert rule
// Expected value is 'Fusion'.
func (o FusionAlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The last time that this alert has been modified.
func (o FusionAlertRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// Azure resource name
func (o FusionAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The severity for alerts created by this alert rule.
func (o FusionAlertRuleOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FusionAlertRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FusionAlertRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tactics of the alert rule
func (o FusionAlertRuleOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

// Azure resource type
func (o FusionAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FusionAlertRuleOutput{})
}
