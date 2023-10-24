// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Friendly RuleSet name mapping to the any RuleSet or secret related information.
// Azure REST API version: 2023-05-01. Prior API version in Azure Native 1.x: 2020-09-01.
//
// Other available API versions: 2023-07-01-preview.
type RuleSet struct {
	pulumi.CustomResourceState

	DeploymentStatus pulumi.StringOutput `pulumi:"deploymentStatus"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the profile which holds the rule set.
	ProfileName pulumi.StringOutput `pulumi:"profileName"`
	// Provisioning status
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRuleSet registers a new resource with the given unique name, arguments, and options.
func NewRuleSet(ctx *pulumi.Context,
	name string, args *RuleSetArgs, opts ...pulumi.ResourceOption) (*RuleSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn/v20200901:RuleSet"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:RuleSet"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:RuleSet"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:RuleSet"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230501:RuleSet"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230701preview:RuleSet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RuleSet
	err := ctx.RegisterResource("azure-native:cdn:RuleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleSet gets an existing RuleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleSetState, opts ...pulumi.ResourceOption) (*RuleSet, error) {
	var resource RuleSet
	err := ctx.ReadResource("azure-native:cdn:RuleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleSet resources.
type ruleSetState struct {
}

type RuleSetState struct {
}

func (RuleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleSetState)(nil)).Elem()
}

type ruleSetArgs struct {
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the rule set under the profile which is unique globally
	RuleSetName *string `pulumi:"ruleSetName"`
}

// The set of arguments for constructing a RuleSet resource.
type RuleSetArgs struct {
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Name of the rule set under the profile which is unique globally
	RuleSetName pulumi.StringPtrInput
}

func (RuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleSetArgs)(nil)).Elem()
}

type RuleSetInput interface {
	pulumi.Input

	ToRuleSetOutput() RuleSetOutput
	ToRuleSetOutputWithContext(ctx context.Context) RuleSetOutput
}

func (*RuleSet) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleSet)(nil)).Elem()
}

func (i *RuleSet) ToRuleSetOutput() RuleSetOutput {
	return i.ToRuleSetOutputWithContext(context.Background())
}

func (i *RuleSet) ToRuleSetOutputWithContext(ctx context.Context) RuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleSetOutput)
}

func (i *RuleSet) ToOutput(ctx context.Context) pulumix.Output[*RuleSet] {
	return pulumix.Output[*RuleSet]{
		OutputState: i.ToRuleSetOutputWithContext(ctx).OutputState,
	}
}

type RuleSetOutput struct{ *pulumi.OutputState }

func (RuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleSet)(nil)).Elem()
}

func (o RuleSetOutput) ToRuleSetOutput() RuleSetOutput {
	return o
}

func (o RuleSetOutput) ToRuleSetOutputWithContext(ctx context.Context) RuleSetOutput {
	return o
}

func (o RuleSetOutput) ToOutput(ctx context.Context) pulumix.Output[*RuleSet] {
	return pulumix.Output[*RuleSet]{
		OutputState: o.OutputState,
	}
}

func (o RuleSetOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleSet) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Resource name.
func (o RuleSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the profile which holds the rule set.
func (o RuleSetOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleSet) pulumi.StringOutput { return v.ProfileName }).(pulumi.StringOutput)
}

// Provisioning status
func (o RuleSetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleSet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Read only system data
func (o RuleSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RuleSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o RuleSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RuleSetOutput{})
}
