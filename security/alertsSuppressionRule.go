// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Describes the suppression rule
// Azure REST API version: 2019-01-01-preview. Prior API version in Azure Native 1.x: 2019-01-01-preview
type AlertsSuppressionRule struct {
	pulumi.CustomResourceState

	// Type of the alert to automatically suppress. For all alert types, use '*'
	AlertType pulumi.StringOutput `pulumi:"alertType"`
	// Any comment regarding the rule
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
	ExpirationDateUtc pulumi.StringPtrOutput `pulumi:"expirationDateUtc"`
	// The last time this rule was modified
	LastModifiedUtc pulumi.StringOutput `pulumi:"lastModifiedUtc"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The reason for dismissing the alert
	Reason pulumi.StringOutput `pulumi:"reason"`
	// Possible states of the rule
	State pulumi.StringOutput `pulumi:"state"`
	// The suppression conditions
	SuppressionAlertsScope SuppressionAlertsScopeResponsePtrOutput `pulumi:"suppressionAlertsScope"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAlertsSuppressionRule registers a new resource with the given unique name, arguments, and options.
func NewAlertsSuppressionRule(ctx *pulumi.Context,
	name string, args *AlertsSuppressionRuleArgs, opts ...pulumi.ResourceOption) (*AlertsSuppressionRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlertType == nil {
		return nil, errors.New("invalid value for required argument 'AlertType'")
	}
	if args.Reason == nil {
		return nil, errors.New("invalid value for required argument 'Reason'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security/v20190101preview:AlertsSuppressionRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AlertsSuppressionRule
	err := ctx.RegisterResource("azure-native:security:AlertsSuppressionRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertsSuppressionRule gets an existing AlertsSuppressionRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertsSuppressionRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertsSuppressionRuleState, opts ...pulumi.ResourceOption) (*AlertsSuppressionRule, error) {
	var resource AlertsSuppressionRule
	err := ctx.ReadResource("azure-native:security:AlertsSuppressionRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertsSuppressionRule resources.
type alertsSuppressionRuleState struct {
}

type AlertsSuppressionRuleState struct {
}

func (AlertsSuppressionRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertsSuppressionRuleState)(nil)).Elem()
}

type alertsSuppressionRuleArgs struct {
	// Type of the alert to automatically suppress. For all alert types, use '*'
	AlertType string `pulumi:"alertType"`
	// The unique name of the suppression alert rule
	AlertsSuppressionRuleName *string `pulumi:"alertsSuppressionRuleName"`
	// Any comment regarding the rule
	Comment *string `pulumi:"comment"`
	// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
	ExpirationDateUtc *string `pulumi:"expirationDateUtc"`
	// The reason for dismissing the alert
	Reason string `pulumi:"reason"`
	// Possible states of the rule
	State string `pulumi:"state"`
	// The suppression conditions
	SuppressionAlertsScope *SuppressionAlertsScope `pulumi:"suppressionAlertsScope"`
}

// The set of arguments for constructing a AlertsSuppressionRule resource.
type AlertsSuppressionRuleArgs struct {
	// Type of the alert to automatically suppress. For all alert types, use '*'
	AlertType pulumi.StringInput
	// The unique name of the suppression alert rule
	AlertsSuppressionRuleName pulumi.StringPtrInput
	// Any comment regarding the rule
	Comment pulumi.StringPtrInput
	// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
	ExpirationDateUtc pulumi.StringPtrInput
	// The reason for dismissing the alert
	Reason pulumi.StringInput
	// Possible states of the rule
	State pulumi.StringInput
	// The suppression conditions
	SuppressionAlertsScope SuppressionAlertsScopePtrInput
}

func (AlertsSuppressionRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertsSuppressionRuleArgs)(nil)).Elem()
}

type AlertsSuppressionRuleInput interface {
	pulumi.Input

	ToAlertsSuppressionRuleOutput() AlertsSuppressionRuleOutput
	ToAlertsSuppressionRuleOutputWithContext(ctx context.Context) AlertsSuppressionRuleOutput
}

func (*AlertsSuppressionRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertsSuppressionRule)(nil)).Elem()
}

func (i *AlertsSuppressionRule) ToAlertsSuppressionRuleOutput() AlertsSuppressionRuleOutput {
	return i.ToAlertsSuppressionRuleOutputWithContext(context.Background())
}

func (i *AlertsSuppressionRule) ToAlertsSuppressionRuleOutputWithContext(ctx context.Context) AlertsSuppressionRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsSuppressionRuleOutput)
}

func (i *AlertsSuppressionRule) ToOutput(ctx context.Context) pulumix.Output[*AlertsSuppressionRule] {
	return pulumix.Output[*AlertsSuppressionRule]{
		OutputState: i.ToAlertsSuppressionRuleOutputWithContext(ctx).OutputState,
	}
}

type AlertsSuppressionRuleOutput struct{ *pulumi.OutputState }

func (AlertsSuppressionRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertsSuppressionRule)(nil)).Elem()
}

func (o AlertsSuppressionRuleOutput) ToAlertsSuppressionRuleOutput() AlertsSuppressionRuleOutput {
	return o
}

func (o AlertsSuppressionRuleOutput) ToAlertsSuppressionRuleOutputWithContext(ctx context.Context) AlertsSuppressionRuleOutput {
	return o
}

func (o AlertsSuppressionRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*AlertsSuppressionRule] {
	return pulumix.Output[*AlertsSuppressionRule]{
		OutputState: o.OutputState,
	}
}

// Type of the alert to automatically suppress. For all alert types, use '*'
func (o AlertsSuppressionRuleOutput) AlertType() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertsSuppressionRule) pulumi.StringOutput { return v.AlertType }).(pulumi.StringOutput)
}

// Any comment regarding the rule
func (o AlertsSuppressionRuleOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertsSuppressionRule) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
func (o AlertsSuppressionRuleOutput) ExpirationDateUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertsSuppressionRule) pulumi.StringPtrOutput { return v.ExpirationDateUtc }).(pulumi.StringPtrOutput)
}

// The last time this rule was modified
func (o AlertsSuppressionRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertsSuppressionRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// Resource name
func (o AlertsSuppressionRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertsSuppressionRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The reason for dismissing the alert
func (o AlertsSuppressionRuleOutput) Reason() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertsSuppressionRule) pulumi.StringOutput { return v.Reason }).(pulumi.StringOutput)
}

// Possible states of the rule
func (o AlertsSuppressionRuleOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertsSuppressionRule) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The suppression conditions
func (o AlertsSuppressionRuleOutput) SuppressionAlertsScope() SuppressionAlertsScopeResponsePtrOutput {
	return o.ApplyT(func(v *AlertsSuppressionRule) SuppressionAlertsScopeResponsePtrOutput {
		return v.SuppressionAlertsScope
	}).(SuppressionAlertsScopeResponsePtrOutput)
}

// Resource type
func (o AlertsSuppressionRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertsSuppressionRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertsSuppressionRuleOutput{})
}
