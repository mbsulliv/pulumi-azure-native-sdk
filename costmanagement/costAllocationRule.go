// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package costmanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The cost allocation rule model definition
// Azure REST API version: 2020-03-01-preview. Prior API version in Azure Native 1.x: 2020-03-01-preview
type CostAllocationRule struct {
	pulumi.CustomResourceState

	// Name of the rule. This is a read only value.
	Name pulumi.StringOutput `pulumi:"name"`
	// Cost allocation rule properties
	Properties CostAllocationRulePropertiesResponseOutput `pulumi:"properties"`
	// Resource type of the rule. This is a read only value of Microsoft.CostManagement/CostAllocationRule.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCostAllocationRule registers a new resource with the given unique name, arguments, and options.
func NewCostAllocationRule(ctx *pulumi.Context,
	name string, args *CostAllocationRuleArgs, opts ...pulumi.ResourceOption) (*CostAllocationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement/v20200301preview:CostAllocationRule"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20230801:CostAllocationRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CostAllocationRule
	err := ctx.RegisterResource("azure-native:costmanagement:CostAllocationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCostAllocationRule gets an existing CostAllocationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCostAllocationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CostAllocationRuleState, opts ...pulumi.ResourceOption) (*CostAllocationRule, error) {
	var resource CostAllocationRule
	err := ctx.ReadResource("azure-native:costmanagement:CostAllocationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CostAllocationRule resources.
type costAllocationRuleState struct {
}

type CostAllocationRuleState struct {
}

func (CostAllocationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*costAllocationRuleState)(nil)).Elem()
}

type costAllocationRuleArgs struct {
	// BillingAccount ID
	BillingAccountId string `pulumi:"billingAccountId"`
	// Cost allocation rule properties
	Properties *CostAllocationRuleProperties `pulumi:"properties"`
	// Cost allocation rule name. The name cannot include spaces or any non alphanumeric characters other than '_' and '-'. The max length is 260 characters.
	RuleName *string `pulumi:"ruleName"`
}

// The set of arguments for constructing a CostAllocationRule resource.
type CostAllocationRuleArgs struct {
	// BillingAccount ID
	BillingAccountId pulumi.StringInput
	// Cost allocation rule properties
	Properties CostAllocationRulePropertiesPtrInput
	// Cost allocation rule name. The name cannot include spaces or any non alphanumeric characters other than '_' and '-'. The max length is 260 characters.
	RuleName pulumi.StringPtrInput
}

func (CostAllocationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*costAllocationRuleArgs)(nil)).Elem()
}

type CostAllocationRuleInput interface {
	pulumi.Input

	ToCostAllocationRuleOutput() CostAllocationRuleOutput
	ToCostAllocationRuleOutputWithContext(ctx context.Context) CostAllocationRuleOutput
}

func (*CostAllocationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationRule)(nil)).Elem()
}

func (i *CostAllocationRule) ToCostAllocationRuleOutput() CostAllocationRuleOutput {
	return i.ToCostAllocationRuleOutputWithContext(context.Background())
}

func (i *CostAllocationRule) ToCostAllocationRuleOutputWithContext(ctx context.Context) CostAllocationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRuleOutput)
}

func (i *CostAllocationRule) ToOutput(ctx context.Context) pulumix.Output[*CostAllocationRule] {
	return pulumix.Output[*CostAllocationRule]{
		OutputState: i.ToCostAllocationRuleOutputWithContext(ctx).OutputState,
	}
}

type CostAllocationRuleOutput struct{ *pulumi.OutputState }

func (CostAllocationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationRule)(nil)).Elem()
}

func (o CostAllocationRuleOutput) ToCostAllocationRuleOutput() CostAllocationRuleOutput {
	return o
}

func (o CostAllocationRuleOutput) ToCostAllocationRuleOutputWithContext(ctx context.Context) CostAllocationRuleOutput {
	return o
}

func (o CostAllocationRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*CostAllocationRule] {
	return pulumix.Output[*CostAllocationRule]{
		OutputState: o.OutputState,
	}
}

// Name of the rule. This is a read only value.
func (o CostAllocationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CostAllocationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Cost allocation rule properties
func (o CostAllocationRuleOutput) Properties() CostAllocationRulePropertiesResponseOutput {
	return o.ApplyT(func(v *CostAllocationRule) CostAllocationRulePropertiesResponseOutput { return v.Properties }).(CostAllocationRulePropertiesResponseOutput)
}

// Resource type of the rule. This is a read only value of Microsoft.CostManagement/CostAllocationRule.
func (o CostAllocationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CostAllocationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CostAllocationRuleOutput{})
}
