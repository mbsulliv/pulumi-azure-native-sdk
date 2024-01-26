// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines binding between a resource and role
func LookupTrustedAccessRoleBinding(ctx *pulumi.Context, args *LookupTrustedAccessRoleBindingArgs, opts ...pulumi.InvokeOption) (*LookupTrustedAccessRoleBindingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTrustedAccessRoleBindingResult
	err := ctx.Invoke("azure-native:containerservice/v20231101:getTrustedAccessRoleBinding", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrustedAccessRoleBindingArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
	// The name of trusted access role binding.
	TrustedAccessRoleBindingName string `pulumi:"trustedAccessRoleBindingName"`
}

// Defines binding between a resource and role
type LookupTrustedAccessRoleBindingResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The current provisioning state of trusted access role binding.
	ProvisioningState string `pulumi:"provisioningState"`
	// A list of roles to bind, each item is a resource type qualified role name. For example: 'Microsoft.MachineLearningServices/workspaces/reader'.
	Roles []string `pulumi:"roles"`
	// The ARM resource ID of source resource that trusted access is configured for.
	SourceResourceId string `pulumi:"sourceResourceId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupTrustedAccessRoleBindingOutput(ctx *pulumi.Context, args LookupTrustedAccessRoleBindingOutputArgs, opts ...pulumi.InvokeOption) LookupTrustedAccessRoleBindingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTrustedAccessRoleBindingResult, error) {
			args := v.(LookupTrustedAccessRoleBindingArgs)
			r, err := LookupTrustedAccessRoleBinding(ctx, &args, opts...)
			var s LookupTrustedAccessRoleBindingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTrustedAccessRoleBindingResultOutput)
}

type LookupTrustedAccessRoleBindingOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// The name of trusted access role binding.
	TrustedAccessRoleBindingName pulumi.StringInput `pulumi:"trustedAccessRoleBindingName"`
}

func (LookupTrustedAccessRoleBindingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrustedAccessRoleBindingArgs)(nil)).Elem()
}

// Defines binding between a resource and role
type LookupTrustedAccessRoleBindingResultOutput struct{ *pulumi.OutputState }

func (LookupTrustedAccessRoleBindingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrustedAccessRoleBindingResult)(nil)).Elem()
}

func (o LookupTrustedAccessRoleBindingResultOutput) ToLookupTrustedAccessRoleBindingResultOutput() LookupTrustedAccessRoleBindingResultOutput {
	return o
}

func (o LookupTrustedAccessRoleBindingResultOutput) ToLookupTrustedAccessRoleBindingResultOutputWithContext(ctx context.Context) LookupTrustedAccessRoleBindingResultOutput {
	return o
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupTrustedAccessRoleBindingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupTrustedAccessRoleBindingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current provisioning state of trusted access role binding.
func (o LookupTrustedAccessRoleBindingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// A list of roles to bind, each item is a resource type qualified role name. For example: 'Microsoft.MachineLearningServices/workspaces/reader'.
func (o LookupTrustedAccessRoleBindingResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// The ARM resource ID of source resource that trusted access is configured for.
func (o LookupTrustedAccessRoleBindingResultOutput) SourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) string { return v.SourceResourceId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupTrustedAccessRoleBindingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupTrustedAccessRoleBindingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrustedAccessRoleBindingResultOutput{})
}
