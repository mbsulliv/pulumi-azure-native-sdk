// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Tenant Activity Log Alert rule resource.
type TenantActivityLogAlert struct {
	pulumi.CustomResourceState

	// The actions that will activate when the condition is met.
	Actions ActionListResponseOutput `pulumi:"actions"`
	// The condition that will cause this alert to activate.
	Condition AlertRuleAllOfConditionResponseOutput `pulumi:"condition"`
	// A description of this Activity Log Alert rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates whether this Activity Log Alert rule is enabled. If an Activity Log Alert rule is not enabled, then none of its actions will be activated.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The location of the resource. Since Azure Activity Log Alerts is a global service, the location of the rules should always be 'global'.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of resource IDs that will be used as prefixes. The alert will only apply to Activity Log events with resource IDs that fall under one of these prefixes. This list must include at least one item.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The tenant GUID. Must be provided for tenant-level and management group events rules.
	TenantScope pulumi.StringPtrOutput `pulumi:"tenantScope"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTenantActivityLogAlert registers a new resource with the given unique name, arguments, and options.
func NewTenantActivityLogAlert(ctx *pulumi.Context,
	name string, args *TenantActivityLogAlertArgs, opts ...pulumi.ResourceOption) (*TenantActivityLogAlert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.Condition == nil {
		return nil, errors.New("invalid value for required argument 'Condition'")
	}
	if args.ManagementGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupName'")
	}
	if args.Enabled == nil {
		args.Enabled = pulumi.BoolPtr(true)
	}
	if args.Location == nil {
		args.Location = pulumi.StringPtr("global")
	}
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TenantActivityLogAlert
	err := ctx.RegisterResource("azure-native:alertsmanagement/v20230401preview:TenantActivityLogAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTenantActivityLogAlert gets an existing TenantActivityLogAlert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTenantActivityLogAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantActivityLogAlertState, opts ...pulumi.ResourceOption) (*TenantActivityLogAlert, error) {
	var resource TenantActivityLogAlert
	err := ctx.ReadResource("azure-native:alertsmanagement/v20230401preview:TenantActivityLogAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TenantActivityLogAlert resources.
type tenantActivityLogAlertState struct {
}

type TenantActivityLogAlertState struct {
}

func (TenantActivityLogAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantActivityLogAlertState)(nil)).Elem()
}

type tenantActivityLogAlertArgs struct {
	// The actions that will activate when the condition is met.
	Actions ActionList `pulumi:"actions"`
	// The name of the Tenant Activity Log Alert rule.
	AlertRuleName *string `pulumi:"alertRuleName"`
	// The condition that will cause this alert to activate.
	Condition AlertRuleAllOfCondition `pulumi:"condition"`
	// A description of this Activity Log Alert rule.
	Description *string `pulumi:"description"`
	// Indicates whether this Activity Log Alert rule is enabled. If an Activity Log Alert rule is not enabled, then none of its actions will be activated.
	Enabled *bool `pulumi:"enabled"`
	// The location of the resource. Since Azure Activity Log Alerts is a global service, the location of the rules should always be 'global'.
	Location *string `pulumi:"location"`
	// The name of the management group. The name is case insensitive.
	ManagementGroupName string `pulumi:"managementGroupName"`
	// A list of resource IDs that will be used as prefixes. The alert will only apply to Activity Log events with resource IDs that fall under one of these prefixes. This list must include at least one item.
	Scopes []string `pulumi:"scopes"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The tenant GUID. Must be provided for tenant-level and management group events rules.
	TenantScope *string `pulumi:"tenantScope"`
}

// The set of arguments for constructing a TenantActivityLogAlert resource.
type TenantActivityLogAlertArgs struct {
	// The actions that will activate when the condition is met.
	Actions ActionListInput
	// The name of the Tenant Activity Log Alert rule.
	AlertRuleName pulumi.StringPtrInput
	// The condition that will cause this alert to activate.
	Condition AlertRuleAllOfConditionInput
	// A description of this Activity Log Alert rule.
	Description pulumi.StringPtrInput
	// Indicates whether this Activity Log Alert rule is enabled. If an Activity Log Alert rule is not enabled, then none of its actions will be activated.
	Enabled pulumi.BoolPtrInput
	// The location of the resource. Since Azure Activity Log Alerts is a global service, the location of the rules should always be 'global'.
	Location pulumi.StringPtrInput
	// The name of the management group. The name is case insensitive.
	ManagementGroupName pulumi.StringInput
	// A list of resource IDs that will be used as prefixes. The alert will only apply to Activity Log events with resource IDs that fall under one of these prefixes. This list must include at least one item.
	Scopes pulumi.StringArrayInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The tenant GUID. Must be provided for tenant-level and management group events rules.
	TenantScope pulumi.StringPtrInput
}

func (TenantActivityLogAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantActivityLogAlertArgs)(nil)).Elem()
}

type TenantActivityLogAlertInput interface {
	pulumi.Input

	ToTenantActivityLogAlertOutput() TenantActivityLogAlertOutput
	ToTenantActivityLogAlertOutputWithContext(ctx context.Context) TenantActivityLogAlertOutput
}

func (*TenantActivityLogAlert) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantActivityLogAlert)(nil)).Elem()
}

func (i *TenantActivityLogAlert) ToTenantActivityLogAlertOutput() TenantActivityLogAlertOutput {
	return i.ToTenantActivityLogAlertOutputWithContext(context.Background())
}

func (i *TenantActivityLogAlert) ToTenantActivityLogAlertOutputWithContext(ctx context.Context) TenantActivityLogAlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantActivityLogAlertOutput)
}

type TenantActivityLogAlertOutput struct{ *pulumi.OutputState }

func (TenantActivityLogAlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantActivityLogAlert)(nil)).Elem()
}

func (o TenantActivityLogAlertOutput) ToTenantActivityLogAlertOutput() TenantActivityLogAlertOutput {
	return o
}

func (o TenantActivityLogAlertOutput) ToTenantActivityLogAlertOutputWithContext(ctx context.Context) TenantActivityLogAlertOutput {
	return o
}

// The actions that will activate when the condition is met.
func (o TenantActivityLogAlertOutput) Actions() ActionListResponseOutput {
	return o.ApplyT(func(v *TenantActivityLogAlert) ActionListResponseOutput { return v.Actions }).(ActionListResponseOutput)
}

// The condition that will cause this alert to activate.
func (o TenantActivityLogAlertOutput) Condition() AlertRuleAllOfConditionResponseOutput {
	return o.ApplyT(func(v *TenantActivityLogAlert) AlertRuleAllOfConditionResponseOutput { return v.Condition }).(AlertRuleAllOfConditionResponseOutput)
}

// A description of this Activity Log Alert rule.
func (o TenantActivityLogAlertOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TenantActivityLogAlert) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates whether this Activity Log Alert rule is enabled. If an Activity Log Alert rule is not enabled, then none of its actions will be activated.
func (o TenantActivityLogAlertOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TenantActivityLogAlert) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The location of the resource. Since Azure Activity Log Alerts is a global service, the location of the rules should always be 'global'.
func (o TenantActivityLogAlertOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TenantActivityLogAlert) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o TenantActivityLogAlertOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantActivityLogAlert) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of resource IDs that will be used as prefixes. The alert will only apply to Activity Log events with resource IDs that fall under one of these prefixes. This list must include at least one item.
func (o TenantActivityLogAlertOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TenantActivityLogAlert) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// The tags of the resource.
func (o TenantActivityLogAlertOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TenantActivityLogAlert) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The tenant GUID. Must be provided for tenant-level and management group events rules.
func (o TenantActivityLogAlertOutput) TenantScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TenantActivityLogAlert) pulumi.StringPtrOutput { return v.TenantScope }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o TenantActivityLogAlertOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantActivityLogAlert) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TenantActivityLogAlertOutput{})
}