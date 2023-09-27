// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The policy definition.
type PolicyDefinitionAtManagementGroup struct {
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

// NewPolicyDefinitionAtManagementGroup registers a new resource with the given unique name, arguments, and options.
func NewPolicyDefinitionAtManagementGroup(ctx *pulumi.Context,
	name string, args *PolicyDefinitionAtManagementGroupArgs, opts ...pulumi.ResourceOption) (*PolicyDefinitionAtManagementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	if args.Mode == nil {
		args.Mode = pulumi.StringPtr("Indexed")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20161201:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180301:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180501:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190101:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190601:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190901:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200301:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200901:PolicyDefinitionAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20230401:PolicyDefinitionAtManagementGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PolicyDefinitionAtManagementGroup
	err := ctx.RegisterResource("azure-native:authorization/v20210601:PolicyDefinitionAtManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyDefinitionAtManagementGroup gets an existing PolicyDefinitionAtManagementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyDefinitionAtManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyDefinitionAtManagementGroupState, opts ...pulumi.ResourceOption) (*PolicyDefinitionAtManagementGroup, error) {
	var resource PolicyDefinitionAtManagementGroup
	err := ctx.ReadResource("azure-native:authorization/v20210601:PolicyDefinitionAtManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyDefinitionAtManagementGroup resources.
type policyDefinitionAtManagementGroupState struct {
}

type PolicyDefinitionAtManagementGroupState struct {
}

func (PolicyDefinitionAtManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyDefinitionAtManagementGroupState)(nil)).Elem()
}

type policyDefinitionAtManagementGroupArgs struct {
	// The policy definition description.
	Description *string `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the management group.
	ManagementGroupId string `pulumi:"managementGroupId"`
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

// The set of arguments for constructing a PolicyDefinitionAtManagementGroup resource.
type PolicyDefinitionAtManagementGroupArgs struct {
	// The policy definition description.
	Description pulumi.StringPtrInput
	// The display name of the policy definition.
	DisplayName pulumi.StringPtrInput
	// The ID of the management group.
	ManagementGroupId pulumi.StringInput
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

func (PolicyDefinitionAtManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyDefinitionAtManagementGroupArgs)(nil)).Elem()
}

type PolicyDefinitionAtManagementGroupInput interface {
	pulumi.Input

	ToPolicyDefinitionAtManagementGroupOutput() PolicyDefinitionAtManagementGroupOutput
	ToPolicyDefinitionAtManagementGroupOutputWithContext(ctx context.Context) PolicyDefinitionAtManagementGroupOutput
}

func (*PolicyDefinitionAtManagementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyDefinitionAtManagementGroup)(nil)).Elem()
}

func (i *PolicyDefinitionAtManagementGroup) ToPolicyDefinitionAtManagementGroupOutput() PolicyDefinitionAtManagementGroupOutput {
	return i.ToPolicyDefinitionAtManagementGroupOutputWithContext(context.Background())
}

func (i *PolicyDefinitionAtManagementGroup) ToPolicyDefinitionAtManagementGroupOutputWithContext(ctx context.Context) PolicyDefinitionAtManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionAtManagementGroupOutput)
}

func (i *PolicyDefinitionAtManagementGroup) ToOutput(ctx context.Context) pulumix.Output[*PolicyDefinitionAtManagementGroup] {
	return pulumix.Output[*PolicyDefinitionAtManagementGroup]{
		OutputState: i.ToPolicyDefinitionAtManagementGroupOutputWithContext(ctx).OutputState,
	}
}

type PolicyDefinitionAtManagementGroupOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionAtManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyDefinitionAtManagementGroup)(nil)).Elem()
}

func (o PolicyDefinitionAtManagementGroupOutput) ToPolicyDefinitionAtManagementGroupOutput() PolicyDefinitionAtManagementGroupOutput {
	return o
}

func (o PolicyDefinitionAtManagementGroupOutput) ToPolicyDefinitionAtManagementGroupOutputWithContext(ctx context.Context) PolicyDefinitionAtManagementGroupOutput {
	return o
}

func (o PolicyDefinitionAtManagementGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*PolicyDefinitionAtManagementGroup] {
	return pulumix.Output[*PolicyDefinitionAtManagementGroup]{
		OutputState: o.OutputState,
	}
}

// The policy definition description.
func (o PolicyDefinitionAtManagementGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the policy definition.
func (o PolicyDefinitionAtManagementGroupOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
func (o PolicyDefinitionAtManagementGroupOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
func (o PolicyDefinitionAtManagementGroupOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the policy definition.
func (o PolicyDefinitionAtManagementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
func (o PolicyDefinitionAtManagementGroupOutput) Parameters() ParameterDefinitionsValueResponseMapOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) ParameterDefinitionsValueResponseMapOutput {
		return v.Parameters
	}).(ParameterDefinitionsValueResponseMapOutput)
}

// The policy rule.
func (o PolicyDefinitionAtManagementGroupOutput) PolicyRule() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.AnyOutput { return v.PolicyRule }).(pulumi.AnyOutput)
}

// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
func (o PolicyDefinitionAtManagementGroupOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.StringPtrOutput { return v.PolicyType }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o PolicyDefinitionAtManagementGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource (Microsoft.Authorization/policyDefinitions).
func (o PolicyDefinitionAtManagementGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyDefinitionAtManagementGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyDefinitionAtManagementGroupOutput{})
}
