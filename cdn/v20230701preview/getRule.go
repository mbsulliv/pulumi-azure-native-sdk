// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an existing delivery rule within a rule set.
func LookupRule(ctx *pulumi.Context, args *LookupRuleArgs, opts ...pulumi.InvokeOption) (*LookupRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRuleResult
	err := ctx.Invoke("azure-native:cdn/v20230701preview:getRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRuleArgs struct {
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the delivery rule which is unique within the endpoint.
	RuleName string `pulumi:"ruleName"`
	// Name of the rule set under the profile.
	RuleSetName string `pulumi:"ruleSetName"`
}

// Friendly Rules name mapping to the any Rules or secret related information.
type LookupRuleResult struct {
	// A list of actions that are executed when all the conditions of a rule are satisfied.
	Actions []interface{} `pulumi:"actions"`
	// A list of conditions that must be matched for the actions to be executed
	Conditions       []interface{} `pulumi:"conditions"`
	DeploymentStatus string        `pulumi:"deploymentStatus"`
	// Resource ID.
	Id string `pulumi:"id"`
	// If this rule is a match should the rules engine continue running the remaining rules or stop. If not present, defaults to Continue.
	MatchProcessingBehavior *string `pulumi:"matchProcessingBehavior"`
	// Resource name.
	Name string `pulumi:"name"`
	// The order in which the rules are applied for the endpoint. Possible values {0,1,2,3,………}. A rule with a lesser order will be applied before a rule with a greater order. Rule with order 0 is a special rule. It does not require any condition and actions listed in it will always be applied.
	Order int `pulumi:"order"`
	// Provisioning status
	ProvisioningState string `pulumi:"provisioningState"`
	// The name of the rule set containing the rule.
	RuleSetName string `pulumi:"ruleSetName"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupRuleResult
func (val *LookupRuleResult) Defaults() *LookupRuleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.MatchProcessingBehavior == nil {
		matchProcessingBehavior_ := "Continue"
		tmp.MatchProcessingBehavior = &matchProcessingBehavior_
	}
	return &tmp
}

func LookupRuleOutput(ctx *pulumi.Context, args LookupRuleOutputArgs, opts ...pulumi.InvokeOption) LookupRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRuleResult, error) {
			args := v.(LookupRuleArgs)
			r, err := LookupRule(ctx, &args, opts...)
			var s LookupRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRuleResultOutput)
}

type LookupRuleOutputArgs struct {
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the delivery rule which is unique within the endpoint.
	RuleName pulumi.StringInput `pulumi:"ruleName"`
	// Name of the rule set under the profile.
	RuleSetName pulumi.StringInput `pulumi:"ruleSetName"`
}

func (LookupRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleArgs)(nil)).Elem()
}

// Friendly Rules name mapping to the any Rules or secret related information.
type LookupRuleResultOutput struct{ *pulumi.OutputState }

func (LookupRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleResult)(nil)).Elem()
}

func (o LookupRuleResultOutput) ToLookupRuleResultOutput() LookupRuleResultOutput {
	return o
}

func (o LookupRuleResultOutput) ToLookupRuleResultOutputWithContext(ctx context.Context) LookupRuleResultOutput {
	return o
}

func (o LookupRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRuleResult] {
	return pulumix.Output[LookupRuleResult]{
		OutputState: o.OutputState,
	}
}

// A list of actions that are executed when all the conditions of a rule are satisfied.
func (o LookupRuleResultOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupRuleResult) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

// A list of conditions that must be matched for the actions to be executed
func (o LookupRuleResultOutput) Conditions() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupRuleResult) []interface{} { return v.Conditions }).(pulumi.ArrayOutput)
}

func (o LookupRuleResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// If this rule is a match should the rules engine continue running the remaining rules or stop. If not present, defaults to Continue.
func (o LookupRuleResultOutput) MatchProcessingBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleResult) *string { return v.MatchProcessingBehavior }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The order in which the rules are applied for the endpoint. Possible values {0,1,2,3,………}. A rule with a lesser order will be applied before a rule with a greater order. Rule with order 0 is a special rule. It does not require any condition and actions listed in it will always be applied.
func (o LookupRuleResultOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRuleResult) int { return v.Order }).(pulumi.IntOutput)
}

// Provisioning status
func (o LookupRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The name of the rule set containing the rule.
func (o LookupRuleResultOutput) RuleSetName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.RuleSetName }).(pulumi.StringOutput)
}

// Read only system data
func (o LookupRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRuleResultOutput{})
}
