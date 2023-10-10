// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Capture logs and metrics of Azure resources based on ARM tags.
type TagRule struct {
	pulumi.CustomResourceState

	// Name of the rule set.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the monitoring tag rules.
	Properties MonitoringTagRulesPropertiesResponseOutput `pulumi:"properties"`
	// The system metadata relating to this resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the rule set.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTagRule registers a new resource with the given unique name, arguments, and options.
func NewTagRule(ctx *pulumi.Context,
	name string, args *TagRuleArgs, opts ...pulumi.ResourceOption) (*TagRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MonitorName == nil {
		return nil, errors.New("invalid value for required argument 'MonitorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:elastic:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20200701:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20200701preview:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20210901preview:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20211001preview:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20220505preview:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20220701preview:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20220901preview:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20230201preview:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20230501preview:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20230615preview:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20230701preview:TagRule"),
		},
		{
			Type: pulumi.String("azure-native:elastic/v20231001preview:TagRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TagRule
	err := ctx.RegisterResource("azure-native:elastic/v20230601:TagRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagRule gets an existing TagRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagRuleState, opts ...pulumi.ResourceOption) (*TagRule, error) {
	var resource TagRule
	err := ctx.ReadResource("azure-native:elastic/v20230601:TagRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagRule resources.
type tagRuleState struct {
}

type TagRuleState struct {
}

func (TagRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagRuleState)(nil)).Elem()
}

type tagRuleArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// Properties of the monitoring tag rules.
	Properties *MonitoringTagRulesProperties `pulumi:"properties"`
	// The name of the resource group to which the Elastic resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tag Rule Set resource name
	RuleSetName *string `pulumi:"ruleSetName"`
}

// The set of arguments for constructing a TagRule resource.
type TagRuleArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput
	// Properties of the monitoring tag rules.
	Properties MonitoringTagRulesPropertiesPtrInput
	// The name of the resource group to which the Elastic resource belongs.
	ResourceGroupName pulumi.StringInput
	// Tag Rule Set resource name
	RuleSetName pulumi.StringPtrInput
}

func (TagRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagRuleArgs)(nil)).Elem()
}

type TagRuleInput interface {
	pulumi.Input

	ToTagRuleOutput() TagRuleOutput
	ToTagRuleOutputWithContext(ctx context.Context) TagRuleOutput
}

func (*TagRule) ElementType() reflect.Type {
	return reflect.TypeOf((**TagRule)(nil)).Elem()
}

func (i *TagRule) ToTagRuleOutput() TagRuleOutput {
	return i.ToTagRuleOutputWithContext(context.Background())
}

func (i *TagRule) ToTagRuleOutputWithContext(ctx context.Context) TagRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagRuleOutput)
}

func (i *TagRule) ToOutput(ctx context.Context) pulumix.Output[*TagRule] {
	return pulumix.Output[*TagRule]{
		OutputState: i.ToTagRuleOutputWithContext(ctx).OutputState,
	}
}

type TagRuleOutput struct{ *pulumi.OutputState }

func (TagRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagRule)(nil)).Elem()
}

func (o TagRuleOutput) ToTagRuleOutput() TagRuleOutput {
	return o
}

func (o TagRuleOutput) ToTagRuleOutputWithContext(ctx context.Context) TagRuleOutput {
	return o
}

func (o TagRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*TagRule] {
	return pulumix.Output[*TagRule]{
		OutputState: o.OutputState,
	}
}

// Name of the rule set.
func (o TagRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the monitoring tag rules.
func (o TagRuleOutput) Properties() MonitoringTagRulesPropertiesResponseOutput {
	return o.ApplyT(func(v *TagRule) MonitoringTagRulesPropertiesResponseOutput { return v.Properties }).(MonitoringTagRulesPropertiesResponseOutput)
}

// The system metadata relating to this resource
func (o TagRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TagRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the rule set.
func (o TagRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TagRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TagRuleOutput{})
}
