// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The policy definition.
// Azure REST API version: 2021-06-01. Prior API version in Azure Native 1.x: 2020-09-01.
//
// Other available API versions: 2016-04-01, 2018-05-01, 2019-06-01, 2023-04-01.
type PolicyDefinition struct {
	pulumi.CustomResourceState

	// The policy definition description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The name of the policy definition.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
	Parameters ParameterDefinitionsValueResponseMapOutput `pulumi:"parameters"`
	// The policy rule.
	PolicyRule pulumi.AnyOutput `pulumi:"policyRule"`
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType pulumi.StringPtrOutput `pulumi:"policyType"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource (Microsoft.Authorization/policyDefinitions).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPolicyDefinition registers a new resource with the given unique name, arguments, and options.
func NewPolicyDefinition(ctx *pulumi.Context,
	name string, args *PolicyDefinitionArgs, opts ...pulumi.ResourceOption) (*PolicyDefinition, error) {
	if args == nil {
		args = &PolicyDefinitionArgs{}
	}

	if args.Mode == nil {
		args.Mode = pulumi.StringPtr("Indexed")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization/v20151001preview:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20160401:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20161201:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180301:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180501:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190101:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190601:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190901:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200301:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200901:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20210601:PolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20230401:PolicyDefinition"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PolicyDefinition
	err := ctx.RegisterResource("azure-native:authorization:PolicyDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyDefinition gets an existing PolicyDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyDefinitionState, opts ...pulumi.ResourceOption) (*PolicyDefinition, error) {
	var resource PolicyDefinition
	err := ctx.ReadResource("azure-native:authorization:PolicyDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyDefinition resources.
type policyDefinitionState struct {
}

type PolicyDefinitionState struct {
}

func (PolicyDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyDefinitionState)(nil)).Elem()
}

type policyDefinitionArgs struct {
	// The policy definition description.
	Description *string `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName *string `pulumi:"displayName"`
	// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
	Mode *string `pulumi:"mode"`
	// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
	Parameters map[string]ParameterDefinitionsValue `pulumi:"parameters"`
	// The name of the policy definition to create.
	PolicyDefinitionName *string `pulumi:"policyDefinitionName"`
	// The policy rule.
	PolicyRule interface{} `pulumi:"policyRule"`
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType *string `pulumi:"policyType"`
}

// The set of arguments for constructing a PolicyDefinition resource.
type PolicyDefinitionArgs struct {
	// The policy definition description.
	Description pulumi.StringPtrInput
	// The display name of the policy definition.
	DisplayName pulumi.StringPtrInput
	// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata pulumi.Input
	// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
	Mode pulumi.StringPtrInput
	// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
	Parameters ParameterDefinitionsValueMapInput
	// The name of the policy definition to create.
	PolicyDefinitionName pulumi.StringPtrInput
	// The policy rule.
	PolicyRule pulumi.Input
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType pulumi.StringPtrInput
}

func (PolicyDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyDefinitionArgs)(nil)).Elem()
}

type PolicyDefinitionInput interface {
	pulumi.Input

	ToPolicyDefinitionOutput() PolicyDefinitionOutput
	ToPolicyDefinitionOutputWithContext(ctx context.Context) PolicyDefinitionOutput
}

func (*PolicyDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyDefinition)(nil)).Elem()
}

func (i *PolicyDefinition) ToPolicyDefinitionOutput() PolicyDefinitionOutput {
	return i.ToPolicyDefinitionOutputWithContext(context.Background())
}

func (i *PolicyDefinition) ToPolicyDefinitionOutputWithContext(ctx context.Context) PolicyDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionOutput)
}

type PolicyDefinitionOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyDefinition)(nil)).Elem()
}

func (o PolicyDefinitionOutput) ToPolicyDefinitionOutput() PolicyDefinitionOutput {
	return o
}

func (o PolicyDefinitionOutput) ToPolicyDefinitionOutputWithContext(ctx context.Context) PolicyDefinitionOutput {
	return o
}

// The policy definition description.
func (o PolicyDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinition) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the policy definition.
func (o PolicyDefinitionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinition) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
func (o PolicyDefinitionOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicyDefinition) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
func (o PolicyDefinitionOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinition) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the policy definition.
func (o PolicyDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
func (o PolicyDefinitionOutput) Parameters() ParameterDefinitionsValueResponseMapOutput {
	return o.ApplyT(func(v *PolicyDefinition) ParameterDefinitionsValueResponseMapOutput { return v.Parameters }).(ParameterDefinitionsValueResponseMapOutput)
}

// The policy rule.
func (o PolicyDefinitionOutput) PolicyRule() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicyDefinition) pulumi.AnyOutput { return v.PolicyRule }).(pulumi.AnyOutput)
}

// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
func (o PolicyDefinitionOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinition) pulumi.StringPtrOutput { return v.PolicyType }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o PolicyDefinitionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PolicyDefinition) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource (Microsoft.Authorization/policyDefinitions).
func (o PolicyDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyDefinitionOutput{})
}
