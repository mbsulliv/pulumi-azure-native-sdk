// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the automation rule.
func LookupAutomationRule(ctx *pulumi.Context, args *LookupAutomationRuleArgs, opts ...pulumi.InvokeOption) (*LookupAutomationRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAutomationRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20240301:getAutomationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutomationRuleArgs struct {
	// Automation rule ID
	AutomationRuleId string `pulumi:"automationRuleId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type LookupAutomationRuleResult struct {
	// The actions to execute when the automation rule is triggered.
	Actions []interface{} `pulumi:"actions"`
	// Information on the client (user or application) that made some action
	CreatedBy ClientInfoResponse `pulumi:"createdBy"`
	// The time the automation rule was created.
	CreatedTimeUtc string `pulumi:"createdTimeUtc"`
	// The display name of the automation rule.
	DisplayName string `pulumi:"displayName"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Information on the client (user or application) that made some action
	LastModifiedBy ClientInfoResponse `pulumi:"lastModifiedBy"`
	// The last time the automation rule was updated.
	LastModifiedTimeUtc string `pulumi:"lastModifiedTimeUtc"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The order of execution of the automation rule.
	Order int `pulumi:"order"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Describes automation rule triggering logic.
	TriggeringLogic AutomationRuleTriggeringLogicResponse `pulumi:"triggeringLogic"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAutomationRuleOutput(ctx *pulumi.Context, args LookupAutomationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupAutomationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutomationRuleResult, error) {
			args := v.(LookupAutomationRuleArgs)
			r, err := LookupAutomationRule(ctx, &args, opts...)
			var s LookupAutomationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutomationRuleResultOutput)
}

type LookupAutomationRuleOutputArgs struct {
	// Automation rule ID
	AutomationRuleId pulumi.StringInput `pulumi:"automationRuleId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAutomationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutomationRuleArgs)(nil)).Elem()
}

type LookupAutomationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupAutomationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutomationRuleResult)(nil)).Elem()
}

func (o LookupAutomationRuleResultOutput) ToLookupAutomationRuleResultOutput() LookupAutomationRuleResultOutput {
	return o
}

func (o LookupAutomationRuleResultOutput) ToLookupAutomationRuleResultOutputWithContext(ctx context.Context) LookupAutomationRuleResultOutput {
	return o
}

// The actions to execute when the automation rule is triggered.
func (o LookupAutomationRuleResultOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

// Information on the client (user or application) that made some action
func (o LookupAutomationRuleResultOutput) CreatedBy() ClientInfoResponseOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) ClientInfoResponse { return v.CreatedBy }).(ClientInfoResponseOutput)
}

// The time the automation rule was created.
func (o LookupAutomationRuleResultOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) string { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

// The display name of the automation rule.
func (o LookupAutomationRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o LookupAutomationRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAutomationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Information on the client (user or application) that made some action
func (o LookupAutomationRuleResultOutput) LastModifiedBy() ClientInfoResponseOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) ClientInfoResponse { return v.LastModifiedBy }).(ClientInfoResponseOutput)
}

// The last time the automation rule was updated.
func (o LookupAutomationRuleResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAutomationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The order of execution of the automation rule.
func (o LookupAutomationRuleResultOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) int { return v.Order }).(pulumi.IntOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAutomationRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Describes automation rule triggering logic.
func (o LookupAutomationRuleResultOutput) TriggeringLogic() AutomationRuleTriggeringLogicResponseOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) AutomationRuleTriggeringLogicResponse { return v.TriggeringLogic }).(AutomationRuleTriggeringLogicResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAutomationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutomationRuleResultOutput{})
}
