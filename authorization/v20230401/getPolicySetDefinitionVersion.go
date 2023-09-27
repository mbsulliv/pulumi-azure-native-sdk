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

// This operation retrieves the policy set definition version in the given subscription with the given name and version.
func LookupPolicySetDefinitionVersion(ctx *pulumi.Context, args *LookupPolicySetDefinitionVersionArgs, opts ...pulumi.InvokeOption) (*LookupPolicySetDefinitionVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicySetDefinitionVersionResult
	err := ctx.Invoke("azure-native:authorization/v20230401:getPolicySetDefinitionVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicySetDefinitionVersionArgs struct {
	// The policy set definition version.  The format is x.y.z where x is the major version number, y is the minor version number, and z is the patch number
	PolicyDefinitionVersion string `pulumi:"policyDefinitionVersion"`
	// The name of the policy set definition.
	PolicySetDefinitionName string `pulumi:"policySetDefinitionName"`
}

// The policy set definition version.
type LookupPolicySetDefinitionVersionResult struct {
	// The policy set definition description.
	Description *string `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the policy set definition version.
	Id string `pulumi:"id"`
	// The policy set definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// The name of the policy set definition version.
	Name string `pulumi:"name"`
	// The policy set definition parameters that can be used in policy definition references.
	Parameters map[string]ParameterDefinitionsValueResponse `pulumi:"parameters"`
	// The metadata describing groups of policy definition references within the policy set definition.
	PolicyDefinitionGroups []PolicyDefinitionGroupResponse `pulumi:"policyDefinitionGroups"`
	// An array of policy definition references.
	PolicyDefinitions []PolicyDefinitionReferenceResponse `pulumi:"policyDefinitions"`
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType *string `pulumi:"policyType"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource (Microsoft.Authorization/policySetDefinitions/versions).
	Type string `pulumi:"type"`
	// The policy set definition version in #.#.# format.
	Version *string `pulumi:"version"`
}

func LookupPolicySetDefinitionVersionOutput(ctx *pulumi.Context, args LookupPolicySetDefinitionVersionOutputArgs, opts ...pulumi.InvokeOption) LookupPolicySetDefinitionVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicySetDefinitionVersionResult, error) {
			args := v.(LookupPolicySetDefinitionVersionArgs)
			r, err := LookupPolicySetDefinitionVersion(ctx, &args, opts...)
			var s LookupPolicySetDefinitionVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicySetDefinitionVersionResultOutput)
}

type LookupPolicySetDefinitionVersionOutputArgs struct {
	// The policy set definition version.  The format is x.y.z where x is the major version number, y is the minor version number, and z is the patch number
	PolicyDefinitionVersion pulumi.StringInput `pulumi:"policyDefinitionVersion"`
	// The name of the policy set definition.
	PolicySetDefinitionName pulumi.StringInput `pulumi:"policySetDefinitionName"`
}

func (LookupPolicySetDefinitionVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicySetDefinitionVersionArgs)(nil)).Elem()
}

// The policy set definition version.
type LookupPolicySetDefinitionVersionResultOutput struct{ *pulumi.OutputState }

func (LookupPolicySetDefinitionVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicySetDefinitionVersionResult)(nil)).Elem()
}

func (o LookupPolicySetDefinitionVersionResultOutput) ToLookupPolicySetDefinitionVersionResultOutput() LookupPolicySetDefinitionVersionResultOutput {
	return o
}

func (o LookupPolicySetDefinitionVersionResultOutput) ToLookupPolicySetDefinitionVersionResultOutputWithContext(ctx context.Context) LookupPolicySetDefinitionVersionResultOutput {
	return o
}

func (o LookupPolicySetDefinitionVersionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPolicySetDefinitionVersionResult] {
	return pulumix.Output[LookupPolicySetDefinitionVersionResult]{
		OutputState: o.OutputState,
	}
}

// The policy set definition description.
func (o LookupPolicySetDefinitionVersionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionVersionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the policy set definition.
func (o LookupPolicySetDefinitionVersionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionVersionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ID of the policy set definition version.
func (o LookupPolicySetDefinitionVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The policy set definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
func (o LookupPolicySetDefinitionVersionResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionVersionResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The name of the policy set definition version.
func (o LookupPolicySetDefinitionVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The policy set definition parameters that can be used in policy definition references.
func (o LookupPolicySetDefinitionVersionResultOutput) Parameters() ParameterDefinitionsValueResponseMapOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionVersionResult) map[string]ParameterDefinitionsValueResponse {
		return v.Parameters
	}).(ParameterDefinitionsValueResponseMapOutput)
}

// The metadata describing groups of policy definition references within the policy set definition.
func (o LookupPolicySetDefinitionVersionResultOutput) PolicyDefinitionGroups() PolicyDefinitionGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionVersionResult) []PolicyDefinitionGroupResponse {
		return v.PolicyDefinitionGroups
	}).(PolicyDefinitionGroupResponseArrayOutput)
}

// An array of policy definition references.
func (o LookupPolicySetDefinitionVersionResultOutput) PolicyDefinitions() PolicyDefinitionReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionVersionResult) []PolicyDefinitionReferenceResponse {
		return v.PolicyDefinitions
	}).(PolicyDefinitionReferenceResponseArrayOutput)
}

// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
func (o LookupPolicySetDefinitionVersionResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionVersionResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o LookupPolicySetDefinitionVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource (Microsoft.Authorization/policySetDefinitions/versions).
func (o LookupPolicySetDefinitionVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

// The policy set definition version in #.#.# format.
func (o LookupPolicySetDefinitionVersionResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionVersionResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicySetDefinitionVersionResultOutput{})
}
