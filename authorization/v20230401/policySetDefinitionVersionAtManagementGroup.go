// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The policy set definition version.
type PolicySetDefinitionVersionAtManagementGroup struct {
	pulumi.CustomResourceState

	// The policy set definition description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The policy set definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The name of the policy set definition version.
	Name pulumi.StringOutput `pulumi:"name"`
	// The policy set definition parameters that can be used in policy definition references.
	Parameters ParameterDefinitionsValueResponseMapOutput `pulumi:"parameters"`
	// The metadata describing groups of policy definition references within the policy set definition.
	PolicyDefinitionGroups PolicyDefinitionGroupResponseArrayOutput `pulumi:"policyDefinitionGroups"`
	// An array of policy definition references.
	PolicyDefinitions PolicyDefinitionReferenceResponseArrayOutput `pulumi:"policyDefinitions"`
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType pulumi.StringPtrOutput `pulumi:"policyType"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource (Microsoft.Authorization/policySetDefinitions/versions).
	Type pulumi.StringOutput `pulumi:"type"`
	// The policy set definition version in #.#.# format.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewPolicySetDefinitionVersionAtManagementGroup registers a new resource with the given unique name, arguments, and options.
func NewPolicySetDefinitionVersionAtManagementGroup(ctx *pulumi.Context,
	name string, args *PolicySetDefinitionVersionAtManagementGroupArgs, opts ...pulumi.ResourceOption) (*PolicySetDefinitionVersionAtManagementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupName'")
	}
	if args.PolicyDefinitions == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDefinitions'")
	}
	if args.PolicySetDefinitionName == nil {
		return nil, errors.New("invalid value for required argument 'PolicySetDefinitionName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:PolicySetDefinitionVersionAtManagementGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PolicySetDefinitionVersionAtManagementGroup
	err := ctx.RegisterResource("azure-native:authorization/v20230401:PolicySetDefinitionVersionAtManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicySetDefinitionVersionAtManagementGroup gets an existing PolicySetDefinitionVersionAtManagementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicySetDefinitionVersionAtManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicySetDefinitionVersionAtManagementGroupState, opts ...pulumi.ResourceOption) (*PolicySetDefinitionVersionAtManagementGroup, error) {
	var resource PolicySetDefinitionVersionAtManagementGroup
	err := ctx.ReadResource("azure-native:authorization/v20230401:PolicySetDefinitionVersionAtManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicySetDefinitionVersionAtManagementGroup resources.
type policySetDefinitionVersionAtManagementGroupState struct {
}

type PolicySetDefinitionVersionAtManagementGroupState struct {
}

func (PolicySetDefinitionVersionAtManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetDefinitionVersionAtManagementGroupState)(nil)).Elem()
}

type policySetDefinitionVersionAtManagementGroupArgs struct {
	// The policy set definition description.
	Description *string `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName *string `pulumi:"displayName"`
	// The name of the management group. The name is case insensitive.
	ManagementGroupName string `pulumi:"managementGroupName"`
	// The policy set definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// The policy set definition parameters that can be used in policy definition references.
	Parameters map[string]ParameterDefinitionsValue `pulumi:"parameters"`
	// The metadata describing groups of policy definition references within the policy set definition.
	PolicyDefinitionGroups []PolicyDefinitionGroup `pulumi:"policyDefinitionGroups"`
	// The policy set definition version.  The format is x.y.z where x is the major version number, y is the minor version number, and z is the patch number
	PolicyDefinitionVersion *string `pulumi:"policyDefinitionVersion"`
	// An array of policy definition references.
	PolicyDefinitions []PolicyDefinitionReference `pulumi:"policyDefinitions"`
	// The name of the policy set definition.
	PolicySetDefinitionName string `pulumi:"policySetDefinitionName"`
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType *string `pulumi:"policyType"`
	// The policy set definition version in #.#.# format.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a PolicySetDefinitionVersionAtManagementGroup resource.
type PolicySetDefinitionVersionAtManagementGroupArgs struct {
	// The policy set definition description.
	Description pulumi.StringPtrInput
	// The display name of the policy set definition.
	DisplayName pulumi.StringPtrInput
	// The name of the management group. The name is case insensitive.
	ManagementGroupName pulumi.StringInput
	// The policy set definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata pulumi.Input
	// The policy set definition parameters that can be used in policy definition references.
	Parameters ParameterDefinitionsValueMapInput
	// The metadata describing groups of policy definition references within the policy set definition.
	PolicyDefinitionGroups PolicyDefinitionGroupArrayInput
	// The policy set definition version.  The format is x.y.z where x is the major version number, y is the minor version number, and z is the patch number
	PolicyDefinitionVersion pulumi.StringPtrInput
	// An array of policy definition references.
	PolicyDefinitions PolicyDefinitionReferenceArrayInput
	// The name of the policy set definition.
	PolicySetDefinitionName pulumi.StringInput
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType pulumi.StringPtrInput
	// The policy set definition version in #.#.# format.
	Version pulumi.StringPtrInput
}

func (PolicySetDefinitionVersionAtManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetDefinitionVersionAtManagementGroupArgs)(nil)).Elem()
}

type PolicySetDefinitionVersionAtManagementGroupInput interface {
	pulumi.Input

	ToPolicySetDefinitionVersionAtManagementGroupOutput() PolicySetDefinitionVersionAtManagementGroupOutput
	ToPolicySetDefinitionVersionAtManagementGroupOutputWithContext(ctx context.Context) PolicySetDefinitionVersionAtManagementGroupOutput
}

func (*PolicySetDefinitionVersionAtManagementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySetDefinitionVersionAtManagementGroup)(nil)).Elem()
}

func (i *PolicySetDefinitionVersionAtManagementGroup) ToPolicySetDefinitionVersionAtManagementGroupOutput() PolicySetDefinitionVersionAtManagementGroupOutput {
	return i.ToPolicySetDefinitionVersionAtManagementGroupOutputWithContext(context.Background())
}

func (i *PolicySetDefinitionVersionAtManagementGroup) ToPolicySetDefinitionVersionAtManagementGroupOutputWithContext(ctx context.Context) PolicySetDefinitionVersionAtManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySetDefinitionVersionAtManagementGroupOutput)
}

func (i *PolicySetDefinitionVersionAtManagementGroup) ToOutput(ctx context.Context) pulumix.Output[*PolicySetDefinitionVersionAtManagementGroup] {
	return pulumix.Output[*PolicySetDefinitionVersionAtManagementGroup]{
		OutputState: i.ToPolicySetDefinitionVersionAtManagementGroupOutputWithContext(ctx).OutputState,
	}
}

type PolicySetDefinitionVersionAtManagementGroupOutput struct{ *pulumi.OutputState }

func (PolicySetDefinitionVersionAtManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySetDefinitionVersionAtManagementGroup)(nil)).Elem()
}

func (o PolicySetDefinitionVersionAtManagementGroupOutput) ToPolicySetDefinitionVersionAtManagementGroupOutput() PolicySetDefinitionVersionAtManagementGroupOutput {
	return o
}

func (o PolicySetDefinitionVersionAtManagementGroupOutput) ToPolicySetDefinitionVersionAtManagementGroupOutputWithContext(ctx context.Context) PolicySetDefinitionVersionAtManagementGroupOutput {
	return o
}

func (o PolicySetDefinitionVersionAtManagementGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*PolicySetDefinitionVersionAtManagementGroup] {
	return pulumix.Output[*PolicySetDefinitionVersionAtManagementGroup]{
		OutputState: o.OutputState,
	}
}

// The policy set definition description.
func (o PolicySetDefinitionVersionAtManagementGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySetDefinitionVersionAtManagementGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the policy set definition.
func (o PolicySetDefinitionVersionAtManagementGroupOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySetDefinitionVersionAtManagementGroup) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The policy set definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
func (o PolicySetDefinitionVersionAtManagementGroupOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicySetDefinitionVersionAtManagementGroup) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

// The name of the policy set definition version.
func (o PolicySetDefinitionVersionAtManagementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySetDefinitionVersionAtManagementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The policy set definition parameters that can be used in policy definition references.
func (o PolicySetDefinitionVersionAtManagementGroupOutput) Parameters() ParameterDefinitionsValueResponseMapOutput {
	return o.ApplyT(func(v *PolicySetDefinitionVersionAtManagementGroup) ParameterDefinitionsValueResponseMapOutput {
		return v.Parameters
	}).(ParameterDefinitionsValueResponseMapOutput)
}

// The metadata describing groups of policy definition references within the policy set definition.
func (o PolicySetDefinitionVersionAtManagementGroupOutput) PolicyDefinitionGroups() PolicyDefinitionGroupResponseArrayOutput {
	return o.ApplyT(func(v *PolicySetDefinitionVersionAtManagementGroup) PolicyDefinitionGroupResponseArrayOutput {
		return v.PolicyDefinitionGroups
	}).(PolicyDefinitionGroupResponseArrayOutput)
}

// An array of policy definition references.
func (o PolicySetDefinitionVersionAtManagementGroupOutput) PolicyDefinitions() PolicyDefinitionReferenceResponseArrayOutput {
	return o.ApplyT(func(v *PolicySetDefinitionVersionAtManagementGroup) PolicyDefinitionReferenceResponseArrayOutput {
		return v.PolicyDefinitions
	}).(PolicyDefinitionReferenceResponseArrayOutput)
}

// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
func (o PolicySetDefinitionVersionAtManagementGroupOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySetDefinitionVersionAtManagementGroup) pulumi.StringPtrOutput { return v.PolicyType }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o PolicySetDefinitionVersionAtManagementGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PolicySetDefinitionVersionAtManagementGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource (Microsoft.Authorization/policySetDefinitions/versions).
func (o PolicySetDefinitionVersionAtManagementGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySetDefinitionVersionAtManagementGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The policy set definition version in #.#.# format.
func (o PolicySetDefinitionVersionAtManagementGroupOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySetDefinitionVersionAtManagementGroup) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicySetDefinitionVersionAtManagementGroupOutput{})
}
