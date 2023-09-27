// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Capture logs and metrics of Azure resources based on ARM tags.
func LookupTagRule(ctx *pulumi.Context, args *LookupTagRuleArgs, opts ...pulumi.InvokeOption) (*LookupTagRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTagRuleResult
	err := ctx.Invoke("azure-native:elastic/v20230601:getTagRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagRuleArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group to which the Elastic resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tag Rule Set resource name
	RuleSetName string `pulumi:"ruleSetName"`
}

// Capture logs and metrics of Azure resources based on ARM tags.
type LookupTagRuleResult struct {
	// The id of the rule set.
	Id string `pulumi:"id"`
	// Name of the rule set.
	Name string `pulumi:"name"`
	// Properties of the monitoring tag rules.
	Properties MonitoringTagRulesPropertiesResponse `pulumi:"properties"`
	// The system metadata relating to this resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the rule set.
	Type string `pulumi:"type"`
}

func LookupTagRuleOutput(ctx *pulumi.Context, args LookupTagRuleOutputArgs, opts ...pulumi.InvokeOption) LookupTagRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagRuleResult, error) {
			args := v.(LookupTagRuleArgs)
			r, err := LookupTagRule(ctx, &args, opts...)
			var s LookupTagRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagRuleResultOutput)
}

type LookupTagRuleOutputArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group to which the Elastic resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Tag Rule Set resource name
	RuleSetName pulumi.StringInput `pulumi:"ruleSetName"`
}

func (LookupTagRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagRuleArgs)(nil)).Elem()
}

// Capture logs and metrics of Azure resources based on ARM tags.
type LookupTagRuleResultOutput struct{ *pulumi.OutputState }

func (LookupTagRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagRuleResult)(nil)).Elem()
}

func (o LookupTagRuleResultOutput) ToLookupTagRuleResultOutput() LookupTagRuleResultOutput {
	return o
}

func (o LookupTagRuleResultOutput) ToLookupTagRuleResultOutputWithContext(ctx context.Context) LookupTagRuleResultOutput {
	return o
}

func (o LookupTagRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupTagRuleResult] {
	return pulumix.Output[LookupTagRuleResult]{
		OutputState: o.OutputState,
	}
}

// The id of the rule set.
func (o LookupTagRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the rule set.
func (o LookupTagRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the monitoring tag rules.
func (o LookupTagRuleResultOutput) Properties() MonitoringTagRulesPropertiesResponseOutput {
	return o.ApplyT(func(v LookupTagRuleResult) MonitoringTagRulesPropertiesResponse { return v.Properties }).(MonitoringTagRulesPropertiesResponseOutput)
}

// The system metadata relating to this resource
func (o LookupTagRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTagRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the rule set.
func (o LookupTagRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagRuleResultOutput{})
}
