// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Alert rule.
//
// Deprecated: Please use one of the variants: FusionAlertRule, MicrosoftSecurityIncidentCreationAlertRule, ScheduledAlertRule.
type AlertRule struct {
	pulumi.CustomResourceState

	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The alert rule kind
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAlertRule registers a new resource with the given unique name, arguments, and options.
func NewAlertRule(ctx *pulumi.Context,
	name string, args *AlertRuleArgs, opts ...pulumi.ResourceOption) (*AlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:AlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource AlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20221101:AlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertRule gets an existing AlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertRuleState, opts ...pulumi.ResourceOption) (*AlertRule, error) {
	var resource AlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20221101:AlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertRule resources.
type alertRuleState struct {
}

type AlertRuleState struct {
}

func (AlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleState)(nil)).Elem()
}

type alertRuleArgs struct {
	// The alert rule kind
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId *string `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a AlertRule resource.
type AlertRuleArgs struct {
	// The alert rule kind
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Alert rule ID
	RuleId pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (AlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleArgs)(nil)).Elem()
}

type AlertRuleInput interface {
	pulumi.Input

	ToAlertRuleOutput() AlertRuleOutput
	ToAlertRuleOutputWithContext(ctx context.Context) AlertRuleOutput
}

func (*AlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRule)(nil)).Elem()
}

func (i *AlertRule) ToAlertRuleOutput() AlertRuleOutput {
	return i.ToAlertRuleOutputWithContext(context.Background())
}

func (i *AlertRule) ToAlertRuleOutputWithContext(ctx context.Context) AlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleOutput)
}

type AlertRuleOutput struct{ *pulumi.OutputState }

func (AlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRule)(nil)).Elem()
}

func (o AlertRuleOutput) ToAlertRuleOutput() AlertRuleOutput {
	return o
}

func (o AlertRuleOutput) ToAlertRuleOutputWithContext(ctx context.Context) AlertRuleOutput {
	return o
}

// Etag of the azure resource
func (o AlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The alert rule kind
func (o AlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o AlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AlertRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AlertRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertRuleOutput{})
}
