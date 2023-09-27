// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This operation retrieves the policy definition version in the given management group with the given name.
func LookupPolicyDefinitionVersionAtManagementGroup(ctx *pulumi.Context, args *LookupPolicyDefinitionVersionAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupPolicyDefinitionVersionAtManagementGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyDefinitionVersionAtManagementGroupResult
	err := ctx.Invoke("azure-native:authorization/v20230401:getPolicyDefinitionVersionAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPolicyDefinitionVersionAtManagementGroupArgs struct {
	// The name of the management group. The name is case insensitive.
	ManagementGroupName string `pulumi:"managementGroupName"`
	// The name of the policy definition.
	PolicyDefinitionName string `pulumi:"policyDefinitionName"`
	// The policy definition version.  The format is x.y.z where x is the major version number, y is the minor version number, and z is the patch number
	PolicyDefinitionVersion string `pulumi:"policyDefinitionVersion"`
}

// The ID of the policy definition version.
type LookupPolicyDefinitionVersionAtManagementGroupResult struct {
	// The policy definition description.
	Description *string `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the policy definition version.
	Id string `pulumi:"id"`
	// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
	Mode *string `pulumi:"mode"`
	// The name of the policy definition version.
	Name string `pulumi:"name"`
	// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
	Parameters map[string]ParameterDefinitionsValueResponse `pulumi:"parameters"`
	// The policy rule.
	PolicyRule interface{} `pulumi:"policyRule"`
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType *string `pulumi:"policyType"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource (Microsoft.Authorization/policyDefinitions/versions).
	Type string `pulumi:"type"`
	// The policy definition version in #.#.# format.
	Version *string `pulumi:"version"`
}

// Defaults sets the appropriate defaults for LookupPolicyDefinitionVersionAtManagementGroupResult
func (val *LookupPolicyDefinitionVersionAtManagementGroupResult) Defaults() *LookupPolicyDefinitionVersionAtManagementGroupResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Mode == nil {
		mode_ := "Indexed"
		tmp.Mode = &mode_
	}
	return &tmp
}

func LookupPolicyDefinitionVersionAtManagementGroupOutput(ctx *pulumi.Context, args LookupPolicyDefinitionVersionAtManagementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyDefinitionVersionAtManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyDefinitionVersionAtManagementGroupResult, error) {
			args := v.(LookupPolicyDefinitionVersionAtManagementGroupArgs)
			r, err := LookupPolicyDefinitionVersionAtManagementGroup(ctx, &args, opts...)
			var s LookupPolicyDefinitionVersionAtManagementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyDefinitionVersionAtManagementGroupResultOutput)
}

type LookupPolicyDefinitionVersionAtManagementGroupOutputArgs struct {
	// The name of the management group. The name is case insensitive.
	ManagementGroupName pulumi.StringInput `pulumi:"managementGroupName"`
	// The name of the policy definition.
	PolicyDefinitionName pulumi.StringInput `pulumi:"policyDefinitionName"`
	// The policy definition version.  The format is x.y.z where x is the major version number, y is the minor version number, and z is the patch number
	PolicyDefinitionVersion pulumi.StringInput `pulumi:"policyDefinitionVersion"`
}

func (LookupPolicyDefinitionVersionAtManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyDefinitionVersionAtManagementGroupArgs)(nil)).Elem()
}

// The ID of the policy definition version.
type LookupPolicyDefinitionVersionAtManagementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyDefinitionVersionAtManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyDefinitionVersionAtManagementGroupResult)(nil)).Elem()
}

func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) ToLookupPolicyDefinitionVersionAtManagementGroupResultOutput() LookupPolicyDefinitionVersionAtManagementGroupResultOutput {
	return o
}

func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) ToLookupPolicyDefinitionVersionAtManagementGroupResultOutputWithContext(ctx context.Context) LookupPolicyDefinitionVersionAtManagementGroupResultOutput {
	return o
}

func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPolicyDefinitionVersionAtManagementGroupResult] {
	return pulumix.Output[LookupPolicyDefinitionVersionAtManagementGroupResult]{
		OutputState: o.OutputState,
	}
}

// The policy definition description.
func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionAtManagementGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the policy definition.
func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionAtManagementGroupResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ID of the policy definition version.
func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionAtManagementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionAtManagementGroupResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionAtManagementGroupResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the policy definition version.
func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionAtManagementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) Parameters() ParameterDefinitionsValueResponseMapOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionAtManagementGroupResult) map[string]ParameterDefinitionsValueResponse {
		return v.Parameters
	}).(ParameterDefinitionsValueResponseMapOutput)
}

// The policy rule.
func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) PolicyRule() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionAtManagementGroupResult) interface{} { return v.PolicyRule }).(pulumi.AnyOutput)
}

// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionAtManagementGroupResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionAtManagementGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource (Microsoft.Authorization/policyDefinitions/versions).
func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionAtManagementGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// The policy definition version in #.#.# format.
func (o LookupPolicyDefinitionVersionAtManagementGroupResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionAtManagementGroupResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyDefinitionVersionAtManagementGroupResultOutput{})
}
