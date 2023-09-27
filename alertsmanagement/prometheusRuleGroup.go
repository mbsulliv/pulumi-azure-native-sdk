// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alertsmanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Prometheus rule group resource.
// Azure REST API version: 2023-03-01.
type PrometheusRuleGroup struct {
	pulumi.CustomResourceState

	// Apply rule to data from a specific cluster.
	ClusterName pulumi.StringPtrOutput `pulumi:"clusterName"`
	// Rule group description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Enable/disable rule group.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The interval in which to run the Prometheus rule group represented in ISO 8601 duration format. Should be between 1 and 15 minutes
	Interval pulumi.StringPtrOutput `pulumi:"interval"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Defines the rules in the Prometheus rule group.
	Rules PrometheusRuleResponseArrayOutput `pulumi:"rules"`
	// Target Azure Monitor workspaces resource ids. This api-version is currently limited to creating with one scope. This may change in future.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrometheusRuleGroup registers a new resource with the given unique name, arguments, and options.
func NewPrometheusRuleGroup(ctx *pulumi.Context,
	name string, args *PrometheusRuleGroupArgs, opts ...pulumi.ResourceOption) (*PrometheusRuleGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20210722preview:PrometheusRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20230301:PrometheusRuleGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrometheusRuleGroup
	err := ctx.RegisterResource("azure-native:alertsmanagement:PrometheusRuleGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrometheusRuleGroup gets an existing PrometheusRuleGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrometheusRuleGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrometheusRuleGroupState, opts ...pulumi.ResourceOption) (*PrometheusRuleGroup, error) {
	var resource PrometheusRuleGroup
	err := ctx.ReadResource("azure-native:alertsmanagement:PrometheusRuleGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrometheusRuleGroup resources.
type prometheusRuleGroupState struct {
}

type PrometheusRuleGroupState struct {
}

func (PrometheusRuleGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusRuleGroupState)(nil)).Elem()
}

type prometheusRuleGroupArgs struct {
	// Apply rule to data from a specific cluster.
	ClusterName *string `pulumi:"clusterName"`
	// Rule group description.
	Description *string `pulumi:"description"`
	// Enable/disable rule group.
	Enabled *bool `pulumi:"enabled"`
	// The interval in which to run the Prometheus rule group represented in ISO 8601 duration format. Should be between 1 and 15 minutes
	Interval *string `pulumi:"interval"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the rule group.
	RuleGroupName *string `pulumi:"ruleGroupName"`
	// Defines the rules in the Prometheus rule group.
	Rules []PrometheusRule `pulumi:"rules"`
	// Target Azure Monitor workspaces resource ids. This api-version is currently limited to creating with one scope. This may change in future.
	Scopes []string `pulumi:"scopes"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrometheusRuleGroup resource.
type PrometheusRuleGroupArgs struct {
	// Apply rule to data from a specific cluster.
	ClusterName pulumi.StringPtrInput
	// Rule group description.
	Description pulumi.StringPtrInput
	// Enable/disable rule group.
	Enabled pulumi.BoolPtrInput
	// The interval in which to run the Prometheus rule group represented in ISO 8601 duration format. Should be between 1 and 15 minutes
	Interval pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the rule group.
	RuleGroupName pulumi.StringPtrInput
	// Defines the rules in the Prometheus rule group.
	Rules PrometheusRuleArrayInput
	// Target Azure Monitor workspaces resource ids. This api-version is currently limited to creating with one scope. This may change in future.
	Scopes pulumi.StringArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (PrometheusRuleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusRuleGroupArgs)(nil)).Elem()
}

type PrometheusRuleGroupInput interface {
	pulumi.Input

	ToPrometheusRuleGroupOutput() PrometheusRuleGroupOutput
	ToPrometheusRuleGroupOutputWithContext(ctx context.Context) PrometheusRuleGroupOutput
}

func (*PrometheusRuleGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusRuleGroup)(nil)).Elem()
}

func (i *PrometheusRuleGroup) ToPrometheusRuleGroupOutput() PrometheusRuleGroupOutput {
	return i.ToPrometheusRuleGroupOutputWithContext(context.Background())
}

func (i *PrometheusRuleGroup) ToPrometheusRuleGroupOutputWithContext(ctx context.Context) PrometheusRuleGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleGroupOutput)
}

func (i *PrometheusRuleGroup) ToOutput(ctx context.Context) pulumix.Output[*PrometheusRuleGroup] {
	return pulumix.Output[*PrometheusRuleGroup]{
		OutputState: i.ToPrometheusRuleGroupOutputWithContext(ctx).OutputState,
	}
}

type PrometheusRuleGroupOutput struct{ *pulumi.OutputState }

func (PrometheusRuleGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusRuleGroup)(nil)).Elem()
}

func (o PrometheusRuleGroupOutput) ToPrometheusRuleGroupOutput() PrometheusRuleGroupOutput {
	return o
}

func (o PrometheusRuleGroupOutput) ToPrometheusRuleGroupOutputWithContext(ctx context.Context) PrometheusRuleGroupOutput {
	return o
}

func (o PrometheusRuleGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*PrometheusRuleGroup] {
	return pulumix.Output[*PrometheusRuleGroup]{
		OutputState: o.OutputState,
	}
}

// Apply rule to data from a specific cluster.
func (o PrometheusRuleGroupOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringPtrOutput { return v.ClusterName }).(pulumi.StringPtrOutput)
}

// Rule group description.
func (o PrometheusRuleGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Enable/disable rule group.
func (o PrometheusRuleGroupOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The interval in which to run the Prometheus rule group represented in ISO 8601 duration format. Should be between 1 and 15 minutes
func (o PrometheusRuleGroupOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringPtrOutput { return v.Interval }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o PrometheusRuleGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o PrometheusRuleGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defines the rules in the Prometheus rule group.
func (o PrometheusRuleGroupOutput) Rules() PrometheusRuleResponseArrayOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) PrometheusRuleResponseArrayOutput { return v.Rules }).(PrometheusRuleResponseArrayOutput)
}

// Target Azure Monitor workspaces resource ids. This api-version is currently limited to creating with one scope. This may change in future.
func (o PrometheusRuleGroupOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PrometheusRuleGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o PrometheusRuleGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrometheusRuleGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrometheusRuleGroupOutput{})
}
