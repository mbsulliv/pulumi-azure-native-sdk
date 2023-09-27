// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieve a Prometheus rule group definition.
func LookupPrometheusRuleGroup(ctx *pulumi.Context, args *LookupPrometheusRuleGroupArgs, opts ...pulumi.InvokeOption) (*LookupPrometheusRuleGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrometheusRuleGroupResult
	err := ctx.Invoke("azure-native:alertsmanagement/v20230301:getPrometheusRuleGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrometheusRuleGroupArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the rule group.
	RuleGroupName string `pulumi:"ruleGroupName"`
}

// The Prometheus rule group resource.
type LookupPrometheusRuleGroupResult struct {
	// Apply rule to data from a specific cluster.
	ClusterName *string `pulumi:"clusterName"`
	// Rule group description.
	Description *string `pulumi:"description"`
	// Enable/disable rule group.
	Enabled *bool `pulumi:"enabled"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The interval in which to run the Prometheus rule group represented in ISO 8601 duration format. Should be between 1 and 15 minutes
	Interval *string `pulumi:"interval"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Defines the rules in the Prometheus rule group.
	Rules []PrometheusRuleResponse `pulumi:"rules"`
	// Target Azure Monitor workspaces resource ids. This api-version is currently limited to creating with one scope. This may change in future.
	Scopes []string `pulumi:"scopes"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupPrometheusRuleGroupOutput(ctx *pulumi.Context, args LookupPrometheusRuleGroupOutputArgs, opts ...pulumi.InvokeOption) LookupPrometheusRuleGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrometheusRuleGroupResult, error) {
			args := v.(LookupPrometheusRuleGroupArgs)
			r, err := LookupPrometheusRuleGroup(ctx, &args, opts...)
			var s LookupPrometheusRuleGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrometheusRuleGroupResultOutput)
}

type LookupPrometheusRuleGroupOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the rule group.
	RuleGroupName pulumi.StringInput `pulumi:"ruleGroupName"`
}

func (LookupPrometheusRuleGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrometheusRuleGroupArgs)(nil)).Elem()
}

// The Prometheus rule group resource.
type LookupPrometheusRuleGroupResultOutput struct{ *pulumi.OutputState }

func (LookupPrometheusRuleGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrometheusRuleGroupResult)(nil)).Elem()
}

func (o LookupPrometheusRuleGroupResultOutput) ToLookupPrometheusRuleGroupResultOutput() LookupPrometheusRuleGroupResultOutput {
	return o
}

func (o LookupPrometheusRuleGroupResultOutput) ToLookupPrometheusRuleGroupResultOutputWithContext(ctx context.Context) LookupPrometheusRuleGroupResultOutput {
	return o
}

func (o LookupPrometheusRuleGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPrometheusRuleGroupResult] {
	return pulumix.Output[LookupPrometheusRuleGroupResult]{
		OutputState: o.OutputState,
	}
}

// Apply rule to data from a specific cluster.
func (o LookupPrometheusRuleGroupResultOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) *string { return v.ClusterName }).(pulumi.StringPtrOutput)
}

// Rule group description.
func (o LookupPrometheusRuleGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Enable/disable rule group.
func (o LookupPrometheusRuleGroupResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPrometheusRuleGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The interval in which to run the Prometheus rule group represented in ISO 8601 duration format. Should be between 1 and 15 minutes
func (o LookupPrometheusRuleGroupResultOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) *string { return v.Interval }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupPrometheusRuleGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPrometheusRuleGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Defines the rules in the Prometheus rule group.
func (o LookupPrometheusRuleGroupResultOutput) Rules() PrometheusRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) []PrometheusRuleResponse { return v.Rules }).(PrometheusRuleResponseArrayOutput)
}

// Target Azure Monitor workspaces resource ids. This api-version is currently limited to creating with one scope. This may change in future.
func (o LookupPrometheusRuleGroupResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupPrometheusRuleGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupPrometheusRuleGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPrometheusRuleGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrometheusRuleGroupResultOutput{})
}
