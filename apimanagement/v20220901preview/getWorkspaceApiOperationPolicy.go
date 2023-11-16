// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the policy configuration at the API Operation level.
func LookupWorkspaceApiOperationPolicy(ctx *pulumi.Context, args *LookupWorkspaceApiOperationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceApiOperationPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceApiOperationPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20220901preview:getWorkspaceApiOperationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWorkspaceApiOperationPolicyArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Policy Export Format.
	Format *string `pulumi:"format"`
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId string `pulumi:"operationId"`
	// The identifier of the Policy.
	PolicyId string `pulumi:"policyId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// Policy Contract details.
type LookupWorkspaceApiOperationPolicyResult struct {
	// Format of the policyContent.
	Format *string `pulumi:"format"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Contents of the Policy as defined by the format.
	Value string `pulumi:"value"`
}

// Defaults sets the appropriate defaults for LookupWorkspaceApiOperationPolicyResult
func (val *LookupWorkspaceApiOperationPolicyResult) Defaults() *LookupWorkspaceApiOperationPolicyResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Format == nil {
		format_ := "xml"
		tmp.Format = &format_
	}
	return &tmp
}

func LookupWorkspaceApiOperationPolicyOutput(ctx *pulumi.Context, args LookupWorkspaceApiOperationPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceApiOperationPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceApiOperationPolicyResult, error) {
			args := v.(LookupWorkspaceApiOperationPolicyArgs)
			r, err := LookupWorkspaceApiOperationPolicy(ctx, &args, opts...)
			var s LookupWorkspaceApiOperationPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceApiOperationPolicyResultOutput)
}

type LookupWorkspaceApiOperationPolicyOutputArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// Policy Export Format.
	Format pulumi.StringPtrInput `pulumi:"format"`
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId pulumi.StringInput `pulumi:"operationId"`
	// The identifier of the Policy.
	PolicyId pulumi.StringInput `pulumi:"policyId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceApiOperationPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceApiOperationPolicyArgs)(nil)).Elem()
}

// Policy Contract details.
type LookupWorkspaceApiOperationPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceApiOperationPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceApiOperationPolicyResult)(nil)).Elem()
}

func (o LookupWorkspaceApiOperationPolicyResultOutput) ToLookupWorkspaceApiOperationPolicyResultOutput() LookupWorkspaceApiOperationPolicyResultOutput {
	return o
}

func (o LookupWorkspaceApiOperationPolicyResultOutput) ToLookupWorkspaceApiOperationPolicyResultOutputWithContext(ctx context.Context) LookupWorkspaceApiOperationPolicyResultOutput {
	return o
}

// Format of the policyContent.
func (o LookupWorkspaceApiOperationPolicyResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationPolicyResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceApiOperationPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceApiOperationPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceApiOperationPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

// Contents of the Policy as defined by the format.
func (o LookupWorkspaceApiOperationPolicyResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiOperationPolicyResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceApiOperationPolicyResultOutput{})
}
