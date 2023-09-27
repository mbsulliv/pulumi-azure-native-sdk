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

// This operation retrieves the policy set definition in the given subscription with the given name.
func LookupPolicySetDefinition(ctx *pulumi.Context, args *LookupPolicySetDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupPolicySetDefinitionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicySetDefinitionResult
	err := ctx.Invoke("azure-native:authorization/v20230401:getPolicySetDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicySetDefinitionArgs struct {
	// The name of the policy set definition to get.
	PolicySetDefinitionName string `pulumi:"policySetDefinitionName"`
}

// The policy set definition.
type LookupPolicySetDefinitionResult struct {
	// The policy set definition description.
	Description *string `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the policy set definition.
	Id string `pulumi:"id"`
	// The policy set definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// The name of the policy set definition.
	Name string `pulumi:"name"`
	// The policy set definition parameters that can be used in policy definition references.
	Parameters map[string]ParameterDefinitionsValueResponse `pulumi:"parameters"`
	// The metadata describing groups of policy definition references within the policy set definition.
	PolicyDefinitionGroups []PolicyDefinitionGroupResponse `pulumi:"policyDefinitionGroups"`
	// An array of policy definition references.
	PolicyDefinitions []PolicyDefinitionReferenceResponse `pulumi:"policyDefinitions"`
	// The type of policy set definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType *string `pulumi:"policyType"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource (Microsoft.Authorization/policySetDefinitions).
	Type string `pulumi:"type"`
	// The policy set definition version in #.#.# format.
	Version *string `pulumi:"version"`
	// A list of available versions for this policy set definition.
	Versions []string `pulumi:"versions"`
}

func LookupPolicySetDefinitionOutput(ctx *pulumi.Context, args LookupPolicySetDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupPolicySetDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicySetDefinitionResult, error) {
			args := v.(LookupPolicySetDefinitionArgs)
			r, err := LookupPolicySetDefinition(ctx, &args, opts...)
			var s LookupPolicySetDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicySetDefinitionResultOutput)
}

type LookupPolicySetDefinitionOutputArgs struct {
	// The name of the policy set definition to get.
	PolicySetDefinitionName pulumi.StringInput `pulumi:"policySetDefinitionName"`
}

func (LookupPolicySetDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicySetDefinitionArgs)(nil)).Elem()
}

// The policy set definition.
type LookupPolicySetDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupPolicySetDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicySetDefinitionResult)(nil)).Elem()
}

func (o LookupPolicySetDefinitionResultOutput) ToLookupPolicySetDefinitionResultOutput() LookupPolicySetDefinitionResultOutput {
	return o
}

func (o LookupPolicySetDefinitionResultOutput) ToLookupPolicySetDefinitionResultOutputWithContext(ctx context.Context) LookupPolicySetDefinitionResultOutput {
	return o
}

func (o LookupPolicySetDefinitionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPolicySetDefinitionResult] {
	return pulumix.Output[LookupPolicySetDefinitionResult]{
		OutputState: o.OutputState,
	}
}

// The policy set definition description.
func (o LookupPolicySetDefinitionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the policy set definition.
func (o LookupPolicySetDefinitionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ID of the policy set definition.
func (o LookupPolicySetDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The policy set definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
func (o LookupPolicySetDefinitionResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The name of the policy set definition.
func (o LookupPolicySetDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The policy set definition parameters that can be used in policy definition references.
func (o LookupPolicySetDefinitionResultOutput) Parameters() ParameterDefinitionsValueResponseMapOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) map[string]ParameterDefinitionsValueResponse {
		return v.Parameters
	}).(ParameterDefinitionsValueResponseMapOutput)
}

// The metadata describing groups of policy definition references within the policy set definition.
func (o LookupPolicySetDefinitionResultOutput) PolicyDefinitionGroups() PolicyDefinitionGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) []PolicyDefinitionGroupResponse {
		return v.PolicyDefinitionGroups
	}).(PolicyDefinitionGroupResponseArrayOutput)
}

// An array of policy definition references.
func (o LookupPolicySetDefinitionResultOutput) PolicyDefinitions() PolicyDefinitionReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) []PolicyDefinitionReferenceResponse {
		return v.PolicyDefinitions
	}).(PolicyDefinitionReferenceResponseArrayOutput)
}

// The type of policy set definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
func (o LookupPolicySetDefinitionResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o LookupPolicySetDefinitionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource (Microsoft.Authorization/policySetDefinitions).
func (o LookupPolicySetDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

// The policy set definition version in #.#.# format.
func (o LookupPolicySetDefinitionResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

// A list of available versions for this policy set definition.
func (o LookupPolicySetDefinitionResultOutput) Versions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) []string { return v.Versions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicySetDefinitionResultOutput{})
}
