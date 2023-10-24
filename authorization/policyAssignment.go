// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The policy assignment.
// Azure REST API version: 2022-06-01. Prior API version in Azure Native 1.x: 2020-09-01.
//
// Other available API versions: 2016-04-01, 2016-12-01, 2019-06-01, 2020-03-01, 2023-04-01.
type PolicyAssignment struct {
	pulumi.CustomResourceState

	// This message will be part of response in case of policy violation.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the policy assignment.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
	EnforcementMode pulumi.StringPtrOutput `pulumi:"enforcementMode"`
	// The managed identity associated with the policy assignment.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The location of the policy assignment. Only required when utilizing managed identity.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The policy assignment metadata. Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The name of the policy assignment.
	Name pulumi.StringOutput `pulumi:"name"`
	// The messages that describe why a resource is non-compliant with the policy.
	NonComplianceMessages NonComplianceMessageResponseArrayOutput `pulumi:"nonComplianceMessages"`
	// The policy's excluded scopes.
	NotScopes pulumi.StringArrayOutput `pulumi:"notScopes"`
	// The policy property value override.
	Overrides OverrideResponseArrayOutput `pulumi:"overrides"`
	// The parameter values for the assigned policy rule. The keys are the parameter names.
	Parameters ParameterValuesValueResponseMapOutput `pulumi:"parameters"`
	// The ID of the policy definition or policy set definition being assigned.
	PolicyDefinitionId pulumi.StringPtrOutput `pulumi:"policyDefinitionId"`
	// The resource selector list to filter policies by resource properties.
	ResourceSelectors ResourceSelectorResponseArrayOutput `pulumi:"resourceSelectors"`
	// The scope for the policy assignment.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the policy assignment.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPolicyAssignment registers a new resource with the given unique name, arguments, and options.
func NewPolicyAssignment(ctx *pulumi.Context,
	name string, args *PolicyAssignmentArgs, opts ...pulumi.ResourceOption) (*PolicyAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.EnforcementMode == nil {
		args.EnforcementMode = pulumi.StringPtr("Default")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization/v20151001preview:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20160401:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20161201:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20170601preview:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180301:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180501:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190101:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190601:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190901:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200301:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200901:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20210601:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20220601:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20230401:PolicyAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PolicyAssignment
	err := ctx.RegisterResource("azure-native:authorization:PolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyAssignment gets an existing PolicyAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyAssignmentState, opts ...pulumi.ResourceOption) (*PolicyAssignment, error) {
	var resource PolicyAssignment
	err := ctx.ReadResource("azure-native:authorization:PolicyAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyAssignment resources.
type policyAssignmentState struct {
}

type PolicyAssignmentState struct {
}

func (PolicyAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAssignmentState)(nil)).Elem()
}

type policyAssignmentArgs struct {
	// This message will be part of response in case of policy violation.
	Description *string `pulumi:"description"`
	// The display name of the policy assignment.
	DisplayName *string `pulumi:"displayName"`
	// The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
	EnforcementMode *string `pulumi:"enforcementMode"`
	// The managed identity associated with the policy assignment.
	Identity *Identity `pulumi:"identity"`
	// The location of the policy assignment. Only required when utilizing managed identity.
	Location *string `pulumi:"location"`
	// The policy assignment metadata. Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// The messages that describe why a resource is non-compliant with the policy.
	NonComplianceMessages []NonComplianceMessage `pulumi:"nonComplianceMessages"`
	// The policy's excluded scopes.
	NotScopes []string `pulumi:"notScopes"`
	// The policy property value override.
	Overrides []Override `pulumi:"overrides"`
	// The parameter values for the assigned policy rule. The keys are the parameter names.
	Parameters map[string]ParameterValuesValue `pulumi:"parameters"`
	// The name of the policy assignment.
	PolicyAssignmentName *string `pulumi:"policyAssignmentName"`
	// The ID of the policy definition or policy set definition being assigned.
	PolicyDefinitionId *string `pulumi:"policyDefinitionId"`
	// The resource selector list to filter policies by resource properties.
	ResourceSelectors []ResourceSelector `pulumi:"resourceSelectors"`
	// The scope of the policy assignment. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}', or resource (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a PolicyAssignment resource.
type PolicyAssignmentArgs struct {
	// This message will be part of response in case of policy violation.
	Description pulumi.StringPtrInput
	// The display name of the policy assignment.
	DisplayName pulumi.StringPtrInput
	// The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
	EnforcementMode pulumi.StringPtrInput
	// The managed identity associated with the policy assignment.
	Identity IdentityPtrInput
	// The location of the policy assignment. Only required when utilizing managed identity.
	Location pulumi.StringPtrInput
	// The policy assignment metadata. Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata pulumi.Input
	// The messages that describe why a resource is non-compliant with the policy.
	NonComplianceMessages NonComplianceMessageArrayInput
	// The policy's excluded scopes.
	NotScopes pulumi.StringArrayInput
	// The policy property value override.
	Overrides OverrideArrayInput
	// The parameter values for the assigned policy rule. The keys are the parameter names.
	Parameters ParameterValuesValueMapInput
	// The name of the policy assignment.
	PolicyAssignmentName pulumi.StringPtrInput
	// The ID of the policy definition or policy set definition being assigned.
	PolicyDefinitionId pulumi.StringPtrInput
	// The resource selector list to filter policies by resource properties.
	ResourceSelectors ResourceSelectorArrayInput
	// The scope of the policy assignment. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}', or resource (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
	Scope pulumi.StringInput
}

func (PolicyAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAssignmentArgs)(nil)).Elem()
}

type PolicyAssignmentInput interface {
	pulumi.Input

	ToPolicyAssignmentOutput() PolicyAssignmentOutput
	ToPolicyAssignmentOutputWithContext(ctx context.Context) PolicyAssignmentOutput
}

func (*PolicyAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignment)(nil)).Elem()
}

func (i *PolicyAssignment) ToPolicyAssignmentOutput() PolicyAssignmentOutput {
	return i.ToPolicyAssignmentOutputWithContext(context.Background())
}

func (i *PolicyAssignment) ToPolicyAssignmentOutputWithContext(ctx context.Context) PolicyAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentOutput)
}

func (i *PolicyAssignment) ToOutput(ctx context.Context) pulumix.Output[*PolicyAssignment] {
	return pulumix.Output[*PolicyAssignment]{
		OutputState: i.ToPolicyAssignmentOutputWithContext(ctx).OutputState,
	}
}

type PolicyAssignmentOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignment)(nil)).Elem()
}

func (o PolicyAssignmentOutput) ToPolicyAssignmentOutput() PolicyAssignmentOutput {
	return o
}

func (o PolicyAssignmentOutput) ToPolicyAssignmentOutputWithContext(ctx context.Context) PolicyAssignmentOutput {
	return o
}

func (o PolicyAssignmentOutput) ToOutput(ctx context.Context) pulumix.Output[*PolicyAssignment] {
	return pulumix.Output[*PolicyAssignment]{
		OutputState: o.OutputState,
	}
}

// This message will be part of response in case of policy violation.
func (o PolicyAssignmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the policy assignment.
func (o PolicyAssignmentOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
func (o PolicyAssignmentOutput) EnforcementMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.EnforcementMode }).(pulumi.StringPtrOutput)
}

// The managed identity associated with the policy assignment.
func (o PolicyAssignmentOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The location of the policy assignment. Only required when utilizing managed identity.
func (o PolicyAssignmentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The policy assignment metadata. Metadata is an open ended object and is typically a collection of key value pairs.
func (o PolicyAssignmentOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

// The name of the policy assignment.
func (o PolicyAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The messages that describe why a resource is non-compliant with the policy.
func (o PolicyAssignmentOutput) NonComplianceMessages() NonComplianceMessageResponseArrayOutput {
	return o.ApplyT(func(v *PolicyAssignment) NonComplianceMessageResponseArrayOutput { return v.NonComplianceMessages }).(NonComplianceMessageResponseArrayOutput)
}

// The policy's excluded scopes.
func (o PolicyAssignmentOutput) NotScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringArrayOutput { return v.NotScopes }).(pulumi.StringArrayOutput)
}

// The policy property value override.
func (o PolicyAssignmentOutput) Overrides() OverrideResponseArrayOutput {
	return o.ApplyT(func(v *PolicyAssignment) OverrideResponseArrayOutput { return v.Overrides }).(OverrideResponseArrayOutput)
}

// The parameter values for the assigned policy rule. The keys are the parameter names.
func (o PolicyAssignmentOutput) Parameters() ParameterValuesValueResponseMapOutput {
	return o.ApplyT(func(v *PolicyAssignment) ParameterValuesValueResponseMapOutput { return v.Parameters }).(ParameterValuesValueResponseMapOutput)
}

// The ID of the policy definition or policy set definition being assigned.
func (o PolicyAssignmentOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringPtrOutput { return v.PolicyDefinitionId }).(pulumi.StringPtrOutput)
}

// The resource selector list to filter policies by resource properties.
func (o PolicyAssignmentOutput) ResourceSelectors() ResourceSelectorResponseArrayOutput {
	return o.ApplyT(func(v *PolicyAssignment) ResourceSelectorResponseArrayOutput { return v.ResourceSelectors }).(ResourceSelectorResponseArrayOutput)
}

// The scope for the policy assignment.
func (o PolicyAssignmentOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o PolicyAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PolicyAssignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the policy assignment.
func (o PolicyAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyAssignmentOutput{})
}
