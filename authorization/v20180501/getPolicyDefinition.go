// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This operation retrieves the policy definition in the given subscription with the given name.
func LookupPolicyDefinition(ctx *pulumi.Context, args *LookupPolicyDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupPolicyDefinitionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyDefinitionResult
	err := ctx.Invoke("azure-native:authorization/v20180501:getPolicyDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyDefinitionArgs struct {
	// The name of the policy definition to get.
	PolicyDefinitionName string `pulumi:"policyDefinitionName"`
}

// The policy definition.
type LookupPolicyDefinitionResult struct {
	// The policy definition description.
	Description *string `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the policy definition.
	Id string `pulumi:"id"`
	// The policy definition metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The policy definition mode. Possible values are NotSpecified, Indexed, and All.
	Mode *string `pulumi:"mode"`
	// The name of the policy definition.
	Name string `pulumi:"name"`
	// Required if a parameter is used in policy rule.
	Parameters interface{} `pulumi:"parameters"`
	// The policy rule.
	PolicyRule interface{} `pulumi:"policyRule"`
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
	PolicyType *string `pulumi:"policyType"`
	// The type of the resource (Microsoft.Authorization/policyDefinitions).
	Type string `pulumi:"type"`
}

func LookupPolicyDefinitionOutput(ctx *pulumi.Context, args LookupPolicyDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyDefinitionResult, error) {
			args := v.(LookupPolicyDefinitionArgs)
			r, err := LookupPolicyDefinition(ctx, &args, opts...)
			var s LookupPolicyDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyDefinitionResultOutput)
}

type LookupPolicyDefinitionOutputArgs struct {
	// The name of the policy definition to get.
	PolicyDefinitionName pulumi.StringInput `pulumi:"policyDefinitionName"`
}

func (LookupPolicyDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyDefinitionArgs)(nil)).Elem()
}

// The policy definition.
type LookupPolicyDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyDefinitionResult)(nil)).Elem()
}

func (o LookupPolicyDefinitionResultOutput) ToLookupPolicyDefinitionResultOutput() LookupPolicyDefinitionResultOutput {
	return o
}

func (o LookupPolicyDefinitionResultOutput) ToLookupPolicyDefinitionResultOutputWithContext(ctx context.Context) LookupPolicyDefinitionResultOutput {
	return o
}

func (o LookupPolicyDefinitionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPolicyDefinitionResult] {
	return pulumix.Output[LookupPolicyDefinitionResult]{
		OutputState: o.OutputState,
	}
}

// The policy definition description.
func (o LookupPolicyDefinitionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the policy definition.
func (o LookupPolicyDefinitionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ID of the policy definition.
func (o LookupPolicyDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The policy definition metadata.
func (o LookupPolicyDefinitionResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The policy definition mode. Possible values are NotSpecified, Indexed, and All.
func (o LookupPolicyDefinitionResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the policy definition.
func (o LookupPolicyDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Required if a parameter is used in policy rule.
func (o LookupPolicyDefinitionResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

// The policy rule.
func (o LookupPolicyDefinitionResultOutput) PolicyRule() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) interface{} { return v.PolicyRule }).(pulumi.AnyOutput)
}

// The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
func (o LookupPolicyDefinitionResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

// The type of the resource (Microsoft.Authorization/policyDefinitions).
func (o LookupPolicyDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyDefinitionResultOutput{})
}
