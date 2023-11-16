// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logz

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Capture metrics of Azure resources based on ARM tags.
// Azure REST API version: 2022-01-01-preview.
func LookupMetricsSourceTagRule(ctx *pulumi.Context, args *LookupMetricsSourceTagRuleArgs, opts ...pulumi.InvokeOption) (*LookupMetricsSourceTagRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMetricsSourceTagRuleResult
	err := ctx.Invoke("azure-native:logz:getMetricsSourceTagRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMetricsSourceTagRuleArgs struct {
	// Metrics Account resource name
	MetricsSourceName string `pulumi:"metricsSourceName"`
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleSetName       string `pulumi:"ruleSetName"`
}

// Capture metrics of Azure resources based on ARM tags.
type LookupMetricsSourceTagRuleResult struct {
	// The id of the rule set.
	Id string `pulumi:"id"`
	// Name of the rule set.
	Name string `pulumi:"name"`
	// Definition of the properties for a TagRules resource.
	Properties MetricsTagRulesPropertiesResponse `pulumi:"properties"`
	// The system metadata relating to this resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the rule set.
	Type string `pulumi:"type"`
}

func LookupMetricsSourceTagRuleOutput(ctx *pulumi.Context, args LookupMetricsSourceTagRuleOutputArgs, opts ...pulumi.InvokeOption) LookupMetricsSourceTagRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMetricsSourceTagRuleResult, error) {
			args := v.(LookupMetricsSourceTagRuleArgs)
			r, err := LookupMetricsSourceTagRule(ctx, &args, opts...)
			var s LookupMetricsSourceTagRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMetricsSourceTagRuleResultOutput)
}

type LookupMetricsSourceTagRuleOutputArgs struct {
	// Metrics Account resource name
	MetricsSourceName pulumi.StringInput `pulumi:"metricsSourceName"`
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleSetName       pulumi.StringInput `pulumi:"ruleSetName"`
}

func (LookupMetricsSourceTagRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricsSourceTagRuleArgs)(nil)).Elem()
}

// Capture metrics of Azure resources based on ARM tags.
type LookupMetricsSourceTagRuleResultOutput struct{ *pulumi.OutputState }

func (LookupMetricsSourceTagRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricsSourceTagRuleResult)(nil)).Elem()
}

func (o LookupMetricsSourceTagRuleResultOutput) ToLookupMetricsSourceTagRuleResultOutput() LookupMetricsSourceTagRuleResultOutput {
	return o
}

func (o LookupMetricsSourceTagRuleResultOutput) ToLookupMetricsSourceTagRuleResultOutputWithContext(ctx context.Context) LookupMetricsSourceTagRuleResultOutput {
	return o
}

// The id of the rule set.
func (o LookupMetricsSourceTagRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceTagRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the rule set.
func (o LookupMetricsSourceTagRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceTagRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Definition of the properties for a TagRules resource.
func (o LookupMetricsSourceTagRuleResultOutput) Properties() MetricsTagRulesPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMetricsSourceTagRuleResult) MetricsTagRulesPropertiesResponse { return v.Properties }).(MetricsTagRulesPropertiesResponseOutput)
}

// The system metadata relating to this resource
func (o LookupMetricsSourceTagRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMetricsSourceTagRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the rule set.
func (o LookupMetricsSourceTagRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceTagRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetricsSourceTagRuleResultOutput{})
}
