// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get Tenant Activity Log Alert rule.
func LookupTenantActivityLogAlert(ctx *pulumi.Context, args *LookupTenantActivityLogAlertArgs, opts ...pulumi.InvokeOption) (*LookupTenantActivityLogAlertResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTenantActivityLogAlertResult
	err := ctx.Invoke("azure-native:alertsmanagement/v20230401preview:getTenantActivityLogAlert", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupTenantActivityLogAlertArgs struct {
	// The name of the Tenant Activity Log Alert rule.
	AlertRuleName string `pulumi:"alertRuleName"`
	// The name of the management group. The name is case insensitive.
	ManagementGroupName string `pulumi:"managementGroupName"`
}

// A Tenant Activity Log Alert rule resource.
type LookupTenantActivityLogAlertResult struct {
	// The actions that will activate when the condition is met.
	Actions ActionListResponse `pulumi:"actions"`
	// The condition that will cause this alert to activate.
	Condition AlertRuleAllOfConditionResponse `pulumi:"condition"`
	// A description of this Activity Log Alert rule.
	Description *string `pulumi:"description"`
	// Indicates whether this Activity Log Alert rule is enabled. If an Activity Log Alert rule is not enabled, then none of its actions will be activated.
	Enabled *bool `pulumi:"enabled"`
	// The resource Id.
	Id string `pulumi:"id"`
	// The location of the resource. Since Azure Activity Log Alerts is a global service, the location of the rules should always be 'global'.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// A list of resource IDs that will be used as prefixes. The alert will only apply to Activity Log events with resource IDs that fall under one of these prefixes. This list must include at least one item.
	Scopes []string `pulumi:"scopes"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The tenant GUID. Must be provided for tenant-level and management group events rules.
	TenantScope *string `pulumi:"tenantScope"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupTenantActivityLogAlertResult
func (val *LookupTenantActivityLogAlertResult) Defaults() *LookupTenantActivityLogAlertResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Enabled == nil {
		enabled_ := true
		tmp.Enabled = &enabled_
	}
	if tmp.Location == nil {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

func LookupTenantActivityLogAlertOutput(ctx *pulumi.Context, args LookupTenantActivityLogAlertOutputArgs, opts ...pulumi.InvokeOption) LookupTenantActivityLogAlertResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTenantActivityLogAlertResult, error) {
			args := v.(LookupTenantActivityLogAlertArgs)
			r, err := LookupTenantActivityLogAlert(ctx, &args, opts...)
			var s LookupTenantActivityLogAlertResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTenantActivityLogAlertResultOutput)
}

type LookupTenantActivityLogAlertOutputArgs struct {
	// The name of the Tenant Activity Log Alert rule.
	AlertRuleName pulumi.StringInput `pulumi:"alertRuleName"`
	// The name of the management group. The name is case insensitive.
	ManagementGroupName pulumi.StringInput `pulumi:"managementGroupName"`
}

func (LookupTenantActivityLogAlertOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTenantActivityLogAlertArgs)(nil)).Elem()
}

// A Tenant Activity Log Alert rule resource.
type LookupTenantActivityLogAlertResultOutput struct{ *pulumi.OutputState }

func (LookupTenantActivityLogAlertResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTenantActivityLogAlertResult)(nil)).Elem()
}

func (o LookupTenantActivityLogAlertResultOutput) ToLookupTenantActivityLogAlertResultOutput() LookupTenantActivityLogAlertResultOutput {
	return o
}

func (o LookupTenantActivityLogAlertResultOutput) ToLookupTenantActivityLogAlertResultOutputWithContext(ctx context.Context) LookupTenantActivityLogAlertResultOutput {
	return o
}

// The actions that will activate when the condition is met.
func (o LookupTenantActivityLogAlertResultOutput) Actions() ActionListResponseOutput {
	return o.ApplyT(func(v LookupTenantActivityLogAlertResult) ActionListResponse { return v.Actions }).(ActionListResponseOutput)
}

// The condition that will cause this alert to activate.
func (o LookupTenantActivityLogAlertResultOutput) Condition() AlertRuleAllOfConditionResponseOutput {
	return o.ApplyT(func(v LookupTenantActivityLogAlertResult) AlertRuleAllOfConditionResponse { return v.Condition }).(AlertRuleAllOfConditionResponseOutput)
}

// A description of this Activity Log Alert rule.
func (o LookupTenantActivityLogAlertResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTenantActivityLogAlertResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates whether this Activity Log Alert rule is enabled. If an Activity Log Alert rule is not enabled, then none of its actions will be activated.
func (o LookupTenantActivityLogAlertResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTenantActivityLogAlertResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The resource Id.
func (o LookupTenantActivityLogAlertResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantActivityLogAlertResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource. Since Azure Activity Log Alerts is a global service, the location of the rules should always be 'global'.
func (o LookupTenantActivityLogAlertResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTenantActivityLogAlertResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupTenantActivityLogAlertResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantActivityLogAlertResult) string { return v.Name }).(pulumi.StringOutput)
}

// A list of resource IDs that will be used as prefixes. The alert will only apply to Activity Log events with resource IDs that fall under one of these prefixes. This list must include at least one item.
func (o LookupTenantActivityLogAlertResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTenantActivityLogAlertResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

// The tags of the resource.
func (o LookupTenantActivityLogAlertResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTenantActivityLogAlertResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The tenant GUID. Must be provided for tenant-level and management group events rules.
func (o LookupTenantActivityLogAlertResultOutput) TenantScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTenantActivityLogAlertResult) *string { return v.TenantScope }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o LookupTenantActivityLogAlertResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantActivityLogAlertResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTenantActivityLogAlertResultOutput{})
}
