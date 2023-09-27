// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Capture logs and metrics of Azure resources based on ARM tags.
type SubAccountTagRule struct {
	pulumi.CustomResourceState

	// Name of the rule set.
	Name pulumi.StringOutput `pulumi:"name"`
	// Definition of the properties for a TagRules resource.
	Properties MonitoringTagRulesPropertiesResponseOutput `pulumi:"properties"`
	// The system metadata relating to this resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the rule set.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSubAccountTagRule registers a new resource with the given unique name, arguments, and options.
func NewSubAccountTagRule(ctx *pulumi.Context,
	name string, args *SubAccountTagRuleArgs, opts ...pulumi.ResourceOption) (*SubAccountTagRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MonitorName == nil {
		return nil, errors.New("invalid value for required argument 'MonitorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubAccountName == nil {
		return nil, errors.New("invalid value for required argument 'SubAccountName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logz:SubAccountTagRule"),
		},
		{
			Type: pulumi.String("azure-native:logz/v20201001:SubAccountTagRule"),
		},
		{
			Type: pulumi.String("azure-native:logz/v20201001preview:SubAccountTagRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SubAccountTagRule
	err := ctx.RegisterResource("azure-native:logz/v20220101preview:SubAccountTagRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubAccountTagRule gets an existing SubAccountTagRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubAccountTagRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubAccountTagRuleState, opts ...pulumi.ResourceOption) (*SubAccountTagRule, error) {
	var resource SubAccountTagRule
	err := ctx.ReadResource("azure-native:logz/v20220101preview:SubAccountTagRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubAccountTagRule resources.
type subAccountTagRuleState struct {
}

type SubAccountTagRuleState struct {
}

func (SubAccountTagRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*subAccountTagRuleState)(nil)).Elem()
}

type subAccountTagRuleArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// Definition of the properties for a TagRules resource.
	Properties *MonitoringTagRulesProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RuleSetName       *string `pulumi:"ruleSetName"`
	// Sub Account resource name
	SubAccountName string `pulumi:"subAccountName"`
}

// The set of arguments for constructing a SubAccountTagRule resource.
type SubAccountTagRuleArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput
	// Definition of the properties for a TagRules resource.
	Properties MonitoringTagRulesPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	RuleSetName       pulumi.StringPtrInput
	// Sub Account resource name
	SubAccountName pulumi.StringInput
}

func (SubAccountTagRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subAccountTagRuleArgs)(nil)).Elem()
}

type SubAccountTagRuleInput interface {
	pulumi.Input

	ToSubAccountTagRuleOutput() SubAccountTagRuleOutput
	ToSubAccountTagRuleOutputWithContext(ctx context.Context) SubAccountTagRuleOutput
}

func (*SubAccountTagRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SubAccountTagRule)(nil)).Elem()
}

func (i *SubAccountTagRule) ToSubAccountTagRuleOutput() SubAccountTagRuleOutput {
	return i.ToSubAccountTagRuleOutputWithContext(context.Background())
}

func (i *SubAccountTagRule) ToSubAccountTagRuleOutputWithContext(ctx context.Context) SubAccountTagRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubAccountTagRuleOutput)
}

func (i *SubAccountTagRule) ToOutput(ctx context.Context) pulumix.Output[*SubAccountTagRule] {
	return pulumix.Output[*SubAccountTagRule]{
		OutputState: i.ToSubAccountTagRuleOutputWithContext(ctx).OutputState,
	}
}

type SubAccountTagRuleOutput struct{ *pulumi.OutputState }

func (SubAccountTagRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubAccountTagRule)(nil)).Elem()
}

func (o SubAccountTagRuleOutput) ToSubAccountTagRuleOutput() SubAccountTagRuleOutput {
	return o
}

func (o SubAccountTagRuleOutput) ToSubAccountTagRuleOutputWithContext(ctx context.Context) SubAccountTagRuleOutput {
	return o
}

func (o SubAccountTagRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*SubAccountTagRule] {
	return pulumix.Output[*SubAccountTagRule]{
		OutputState: o.OutputState,
	}
}

// Name of the rule set.
func (o SubAccountTagRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SubAccountTagRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Definition of the properties for a TagRules resource.
func (o SubAccountTagRuleOutput) Properties() MonitoringTagRulesPropertiesResponseOutput {
	return o.ApplyT(func(v *SubAccountTagRule) MonitoringTagRulesPropertiesResponseOutput { return v.Properties }).(MonitoringTagRulesPropertiesResponseOutput)
}

// The system metadata relating to this resource
func (o SubAccountTagRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SubAccountTagRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the rule set.
func (o SubAccountTagRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SubAccountTagRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SubAccountTagRuleOutput{})
}
