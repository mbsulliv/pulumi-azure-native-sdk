// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This operation retrieves a single policy assignment, given its name and the scope it was created at.
func LookupPolicyAssignment(ctx *pulumi.Context, args *LookupPolicyAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupPolicyAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20190601:getPolicyAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyAssignmentArgs struct {
	// The name of the policy assignment to get.
	PolicyAssignmentName string `pulumi:"policyAssignmentName"`
	// The scope of the policy assignment. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}', or resource (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
	Scope string `pulumi:"scope"`
}

// The policy assignment.
type LookupPolicyAssignmentResult struct {
	// This message will be part of response in case of policy violation.
	Description *string `pulumi:"description"`
	// The display name of the policy assignment.
	DisplayName *string `pulumi:"displayName"`
	// The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
	EnforcementMode *string `pulumi:"enforcementMode"`
	// The ID of the policy assignment.
	Id string `pulumi:"id"`
	// The managed identity associated with the policy assignment.
	Identity *IdentityResponse `pulumi:"identity"`
	// The location of the policy assignment. Only required when utilizing managed identity.
	Location *string `pulumi:"location"`
	// The policy assignment metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The name of the policy assignment.
	Name string `pulumi:"name"`
	// The policy's excluded scopes.
	NotScopes []string `pulumi:"notScopes"`
	// Required if a parameter is used in policy rule.
	Parameters interface{} `pulumi:"parameters"`
	// The ID of the policy definition or policy set definition being assigned.
	PolicyDefinitionId *string `pulumi:"policyDefinitionId"`
	// The scope for the policy assignment.
	Scope *string `pulumi:"scope"`
	// The policy sku. This property is optional, obsolete, and will be ignored.
	Sku *PolicySkuResponse `pulumi:"sku"`
	// The type of the policy assignment.
	Type string `pulumi:"type"`
}

func LookupPolicyAssignmentOutput(ctx *pulumi.Context, args LookupPolicyAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyAssignmentResult, error) {
			args := v.(LookupPolicyAssignmentArgs)
			r, err := LookupPolicyAssignment(ctx, &args, opts...)
			var s LookupPolicyAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyAssignmentResultOutput)
}

type LookupPolicyAssignmentOutputArgs struct {
	// The name of the policy assignment to get.
	PolicyAssignmentName pulumi.StringInput `pulumi:"policyAssignmentName"`
	// The scope of the policy assignment. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}', or resource (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupPolicyAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyAssignmentArgs)(nil)).Elem()
}

// The policy assignment.
type LookupPolicyAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyAssignmentResult)(nil)).Elem()
}

func (o LookupPolicyAssignmentResultOutput) ToLookupPolicyAssignmentResultOutput() LookupPolicyAssignmentResultOutput {
	return o
}

func (o LookupPolicyAssignmentResultOutput) ToLookupPolicyAssignmentResultOutputWithContext(ctx context.Context) LookupPolicyAssignmentResultOutput {
	return o
}

// This message will be part of response in case of policy violation.
func (o LookupPolicyAssignmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the policy assignment.
func (o LookupPolicyAssignmentResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
func (o LookupPolicyAssignmentResultOutput) EnforcementMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.EnforcementMode }).(pulumi.StringPtrOutput)
}

// The ID of the policy assignment.
func (o LookupPolicyAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The managed identity associated with the policy assignment.
func (o LookupPolicyAssignmentResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// The location of the policy assignment. Only required when utilizing managed identity.
func (o LookupPolicyAssignmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The policy assignment metadata.
func (o LookupPolicyAssignmentResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The name of the policy assignment.
func (o LookupPolicyAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The policy's excluded scopes.
func (o LookupPolicyAssignmentResultOutput) NotScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) []string { return v.NotScopes }).(pulumi.StringArrayOutput)
}

// Required if a parameter is used in policy rule.
func (o LookupPolicyAssignmentResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

// The ID of the policy definition or policy set definition being assigned.
func (o LookupPolicyAssignmentResultOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.PolicyDefinitionId }).(pulumi.StringPtrOutput)
}

// The scope for the policy assignment.
func (o LookupPolicyAssignmentResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// The policy sku. This property is optional, obsolete, and will be ignored.
func (o LookupPolicyAssignmentResultOutput) Sku() PolicySkuResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *PolicySkuResponse { return v.Sku }).(PolicySkuResponsePtrOutput)
}

// The type of the policy assignment.
func (o LookupPolicyAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyAssignmentResultOutput{})
}
