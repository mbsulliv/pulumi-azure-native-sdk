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

// This operation retrieves the policy definition version in the given subscription with the given name.
func LookupPolicyDefinitionVersion(ctx *pulumi.Context, args *LookupPolicyDefinitionVersionArgs, opts ...pulumi.InvokeOption) (*LookupPolicyDefinitionVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyDefinitionVersionResult
	err := ctx.Invoke("azure-native:authorization/v20230401:getPolicyDefinitionVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPolicyDefinitionVersionArgs struct {
	// The name of the policy definition.
	PolicyDefinitionName string `pulumi:"policyDefinitionName"`
	// The policy definition version.  The format is x.y.z where x is the major version number, y is the minor version number, and z is the patch number
	PolicyDefinitionVersion string `pulumi:"policyDefinitionVersion"`
}

// The ID of the policy definition version.
type LookupPolicyDefinitionVersionResult struct {
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

// Defaults sets the appropriate defaults for LookupPolicyDefinitionVersionResult
func (val *LookupPolicyDefinitionVersionResult) Defaults() *LookupPolicyDefinitionVersionResult {
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

func LookupPolicyDefinitionVersionOutput(ctx *pulumi.Context, args LookupPolicyDefinitionVersionOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyDefinitionVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyDefinitionVersionResult, error) {
			args := v.(LookupPolicyDefinitionVersionArgs)
			r, err := LookupPolicyDefinitionVersion(ctx, &args, opts...)
			var s LookupPolicyDefinitionVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyDefinitionVersionResultOutput)
}

type LookupPolicyDefinitionVersionOutputArgs struct {
	// The name of the policy definition.
	PolicyDefinitionName pulumi.StringInput `pulumi:"policyDefinitionName"`
	// The policy definition version.  The format is x.y.z where x is the major version number, y is the minor version number, and z is the patch number
	PolicyDefinitionVersion pulumi.StringInput `pulumi:"policyDefinitionVersion"`
}

func (LookupPolicyDefinitionVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyDefinitionVersionArgs)(nil)).Elem()
}

// The ID of the policy definition version.
type LookupPolicyDefinitionVersionResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyDefinitionVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyDefinitionVersionResult)(nil)).Elem()
}

func (o LookupPolicyDefinitionVersionResultOutput) ToLookupPolicyDefinitionVersionResultOutput() LookupPolicyDefinitionVersionResultOutput {
	return o
}

func (o LookupPolicyDefinitionVersionResultOutput) ToLookupPolicyDefinitionVersionResultOutputWithContext(ctx context.Context) LookupPolicyDefinitionVersionResultOutput {
	return o
}

func (o LookupPolicyDefinitionVersionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPolicyDefinitionVersionResult] {
	return pulumix.Output[LookupPolicyDefinitionVersionResult]{
		OutputState: o.OutputState,
	}
}

// The policy definition description.
func (o LookupPolicyDefinitionVersionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the policy definition.
func (o LookupPolicyDefinitionVersionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ID of the policy definition version.
func (o LookupPolicyDefinitionVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
func (o LookupPolicyDefinitionVersionResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
func (o LookupPolicyDefinitionVersionResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the policy definition version.
func (o LookupPolicyDefinitionVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
func (o LookupPolicyDefinitionVersionResultOutput) Parameters() ParameterDefinitionsValueResponseMapOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionResult) map[string]ParameterDefinitionsValueResponse {
		return v.Parameters
	}).(ParameterDefinitionsValueResponseMapOutput)
}

// The policy rule.
func (o LookupPolicyDefinitionVersionResultOutput) PolicyRule() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionResult) interface{} { return v.PolicyRule }).(pulumi.AnyOutput)
}

// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
func (o LookupPolicyDefinitionVersionResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o LookupPolicyDefinitionVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource (Microsoft.Authorization/policyDefinitions/versions).
func (o LookupPolicyDefinitionVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

// The policy definition version in #.#.# format.
func (o LookupPolicyDefinitionVersionResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionVersionResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyDefinitionVersionResultOutput{})
}
