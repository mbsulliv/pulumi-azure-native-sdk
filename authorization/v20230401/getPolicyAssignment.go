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

// This operation retrieves a single policy assignment, given its name and the scope it was created at.
func LookupPolicyAssignment(ctx *pulumi.Context, args *LookupPolicyAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupPolicyAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20230401:getPolicyAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
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
	// The policy assignment metadata. Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// The name of the policy assignment.
	Name string `pulumi:"name"`
	// The messages that describe why a resource is non-compliant with the policy.
	NonComplianceMessages []NonComplianceMessageResponse `pulumi:"nonComplianceMessages"`
	// The policy's excluded scopes.
	NotScopes []string `pulumi:"notScopes"`
	// The policy property value override.
	Overrides []OverrideResponse `pulumi:"overrides"`
	// The parameter values for the assigned policy rule. The keys are the parameter names.
	Parameters map[string]ParameterValuesValueResponse `pulumi:"parameters"`
	// The ID of the policy definition or policy set definition being assigned.
	PolicyDefinitionId *string `pulumi:"policyDefinitionId"`
	// The resource selector list to filter policies by resource properties.
	ResourceSelectors []ResourceSelectorResponse `pulumi:"resourceSelectors"`
	// The scope for the policy assignment.
	Scope string `pulumi:"scope"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the policy assignment.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupPolicyAssignmentResult
func (val *LookupPolicyAssignmentResult) Defaults() *LookupPolicyAssignmentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EnforcementMode == nil {
		enforcementMode_ := "Default"
		tmp.EnforcementMode = &enforcementMode_
	}
	return &tmp
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

func (o LookupPolicyAssignmentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPolicyAssignmentResult] {
	return pulumix.Output[LookupPolicyAssignmentResult]{
		OutputState: o.OutputState,
	}
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

// The policy assignment metadata. Metadata is an open ended object and is typically a collection of key value pairs.
func (o LookupPolicyAssignmentResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The name of the policy assignment.
func (o LookupPolicyAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The messages that describe why a resource is non-compliant with the policy.
func (o LookupPolicyAssignmentResultOutput) NonComplianceMessages() NonComplianceMessageResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) []NonComplianceMessageResponse { return v.NonComplianceMessages }).(NonComplianceMessageResponseArrayOutput)
}

// The policy's excluded scopes.
func (o LookupPolicyAssignmentResultOutput) NotScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) []string { return v.NotScopes }).(pulumi.StringArrayOutput)
}

// The policy property value override.
func (o LookupPolicyAssignmentResultOutput) Overrides() OverrideResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) []OverrideResponse { return v.Overrides }).(OverrideResponseArrayOutput)
}

// The parameter values for the assigned policy rule. The keys are the parameter names.
func (o LookupPolicyAssignmentResultOutput) Parameters() ParameterValuesValueResponseMapOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) map[string]ParameterValuesValueResponse { return v.Parameters }).(ParameterValuesValueResponseMapOutput)
}

// The ID of the policy definition or policy set definition being assigned.
func (o LookupPolicyAssignmentResultOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.PolicyDefinitionId }).(pulumi.StringPtrOutput)
}

// The resource selector list to filter policies by resource properties.
func (o LookupPolicyAssignmentResultOutput) ResourceSelectors() ResourceSelectorResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) []ResourceSelectorResponse { return v.ResourceSelectors }).(ResourceSelectorResponseArrayOutput)
}

// The scope for the policy assignment.
func (o LookupPolicyAssignmentResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) string { return v.Scope }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o LookupPolicyAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the policy assignment.
func (o LookupPolicyAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyAssignmentResultOutput{})
}
