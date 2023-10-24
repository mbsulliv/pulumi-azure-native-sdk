// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logz

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Capture metrics of Azure resources based on ARM tags.
// Azure REST API version: 2022-01-01-preview. Prior API version in Azure Native 1.x: 2022-01-01-preview.
type MetricsSourceTagRule struct {
	pulumi.CustomResourceState

	// Name of the rule set.
	Name pulumi.StringOutput `pulumi:"name"`
	// Definition of the properties for a TagRules resource.
	Properties MetricsTagRulesPropertiesResponseOutput `pulumi:"properties"`
	// The system metadata relating to this resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the rule set.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMetricsSourceTagRule registers a new resource with the given unique name, arguments, and options.
func NewMetricsSourceTagRule(ctx *pulumi.Context,
	name string, args *MetricsSourceTagRuleArgs, opts ...pulumi.ResourceOption) (*MetricsSourceTagRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetricsSourceName == nil {
		return nil, errors.New("invalid value for required argument 'MetricsSourceName'")
	}
	if args.MonitorName == nil {
		return nil, errors.New("invalid value for required argument 'MonitorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logz/v20220101preview:MetricsSourceTagRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MetricsSourceTagRule
	err := ctx.RegisterResource("azure-native:logz:MetricsSourceTagRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricsSourceTagRule gets an existing MetricsSourceTagRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricsSourceTagRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricsSourceTagRuleState, opts ...pulumi.ResourceOption) (*MetricsSourceTagRule, error) {
	var resource MetricsSourceTagRule
	err := ctx.ReadResource("azure-native:logz:MetricsSourceTagRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricsSourceTagRule resources.
type metricsSourceTagRuleState struct {
}

type MetricsSourceTagRuleState struct {
}

func (MetricsSourceTagRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricsSourceTagRuleState)(nil)).Elem()
}

type metricsSourceTagRuleArgs struct {
	// Metrics Account resource name
	MetricsSourceName string `pulumi:"metricsSourceName"`
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// Definition of the properties for a TagRules resource.
	Properties *MetricsTagRulesProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RuleSetName       *string `pulumi:"ruleSetName"`
}

// The set of arguments for constructing a MetricsSourceTagRule resource.
type MetricsSourceTagRuleArgs struct {
	// Metrics Account resource name
	MetricsSourceName pulumi.StringInput
	// Monitor resource name
	MonitorName pulumi.StringInput
	// Definition of the properties for a TagRules resource.
	Properties MetricsTagRulesPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	RuleSetName       pulumi.StringPtrInput
}

func (MetricsSourceTagRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricsSourceTagRuleArgs)(nil)).Elem()
}

type MetricsSourceTagRuleInput interface {
	pulumi.Input

	ToMetricsSourceTagRuleOutput() MetricsSourceTagRuleOutput
	ToMetricsSourceTagRuleOutputWithContext(ctx context.Context) MetricsSourceTagRuleOutput
}

func (*MetricsSourceTagRule) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricsSourceTagRule)(nil)).Elem()
}

func (i *MetricsSourceTagRule) ToMetricsSourceTagRuleOutput() MetricsSourceTagRuleOutput {
	return i.ToMetricsSourceTagRuleOutputWithContext(context.Background())
}

func (i *MetricsSourceTagRule) ToMetricsSourceTagRuleOutputWithContext(ctx context.Context) MetricsSourceTagRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricsSourceTagRuleOutput)
}

func (i *MetricsSourceTagRule) ToOutput(ctx context.Context) pulumix.Output[*MetricsSourceTagRule] {
	return pulumix.Output[*MetricsSourceTagRule]{
		OutputState: i.ToMetricsSourceTagRuleOutputWithContext(ctx).OutputState,
	}
}

type MetricsSourceTagRuleOutput struct{ *pulumi.OutputState }

func (MetricsSourceTagRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricsSourceTagRule)(nil)).Elem()
}

func (o MetricsSourceTagRuleOutput) ToMetricsSourceTagRuleOutput() MetricsSourceTagRuleOutput {
	return o
}

func (o MetricsSourceTagRuleOutput) ToMetricsSourceTagRuleOutputWithContext(ctx context.Context) MetricsSourceTagRuleOutput {
	return o
}

func (o MetricsSourceTagRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*MetricsSourceTagRule] {
	return pulumix.Output[*MetricsSourceTagRule]{
		OutputState: o.OutputState,
	}
}

// Name of the rule set.
func (o MetricsSourceTagRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsSourceTagRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Definition of the properties for a TagRules resource.
func (o MetricsSourceTagRuleOutput) Properties() MetricsTagRulesPropertiesResponseOutput {
	return o.ApplyT(func(v *MetricsSourceTagRule) MetricsTagRulesPropertiesResponseOutput { return v.Properties }).(MetricsTagRulesPropertiesResponseOutput)
}

// The system metadata relating to this resource
func (o MetricsSourceTagRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MetricsSourceTagRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the rule set.
func (o MetricsSourceTagRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsSourceTagRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MetricsSourceTagRuleOutput{})
}
