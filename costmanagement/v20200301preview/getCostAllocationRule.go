// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a cost allocation rule by rule name and billing account or enterprise enrollment.
func LookupCostAllocationRule(ctx *pulumi.Context, args *LookupCostAllocationRuleArgs, opts ...pulumi.InvokeOption) (*LookupCostAllocationRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCostAllocationRuleResult
	err := ctx.Invoke("azure-native:costmanagement/v20200301preview:getCostAllocationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCostAllocationRuleArgs struct {
	// BillingAccount ID
	BillingAccountId string `pulumi:"billingAccountId"`
	// Cost allocation rule name. The name cannot include spaces or any non alphanumeric characters other than '_' and '-'. The max length is 260 characters.
	RuleName string `pulumi:"ruleName"`
}

// The cost allocation rule model definition
type LookupCostAllocationRuleResult struct {
	// Azure Resource Manager Id for the rule. This is a read ony value.
	Id string `pulumi:"id"`
	// Name of the rule. This is a read only value.
	Name string `pulumi:"name"`
	// Cost allocation rule properties
	Properties CostAllocationRulePropertiesResponse `pulumi:"properties"`
	// Resource type of the rule. This is a read only value of Microsoft.CostManagement/CostAllocationRule.
	Type string `pulumi:"type"`
}

func LookupCostAllocationRuleOutput(ctx *pulumi.Context, args LookupCostAllocationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupCostAllocationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCostAllocationRuleResult, error) {
			args := v.(LookupCostAllocationRuleArgs)
			r, err := LookupCostAllocationRule(ctx, &args, opts...)
			var s LookupCostAllocationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCostAllocationRuleResultOutput)
}

type LookupCostAllocationRuleOutputArgs struct {
	// BillingAccount ID
	BillingAccountId pulumi.StringInput `pulumi:"billingAccountId"`
	// Cost allocation rule name. The name cannot include spaces or any non alphanumeric characters other than '_' and '-'. The max length is 260 characters.
	RuleName pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupCostAllocationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCostAllocationRuleArgs)(nil)).Elem()
}

// The cost allocation rule model definition
type LookupCostAllocationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupCostAllocationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCostAllocationRuleResult)(nil)).Elem()
}

func (o LookupCostAllocationRuleResultOutput) ToLookupCostAllocationRuleResultOutput() LookupCostAllocationRuleResultOutput {
	return o
}

func (o LookupCostAllocationRuleResultOutput) ToLookupCostAllocationRuleResultOutputWithContext(ctx context.Context) LookupCostAllocationRuleResultOutput {
	return o
}

func (o LookupCostAllocationRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCostAllocationRuleResult] {
	return pulumix.Output[LookupCostAllocationRuleResult]{
		OutputState: o.OutputState,
	}
}

// Azure Resource Manager Id for the rule. This is a read ony value.
func (o LookupCostAllocationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCostAllocationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the rule. This is a read only value.
func (o LookupCostAllocationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCostAllocationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Cost allocation rule properties
func (o LookupCostAllocationRuleResultOutput) Properties() CostAllocationRulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupCostAllocationRuleResult) CostAllocationRulePropertiesResponse { return v.Properties }).(CostAllocationRulePropertiesResponseOutput)
}

// Resource type of the rule. This is a read only value of Microsoft.CostManagement/CostAllocationRule.
func (o LookupCostAllocationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCostAllocationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCostAllocationRuleResultOutput{})
}
