// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the policy configuration at the API Operation level.
func LookupApiOperationPolicy(ctx *pulumi.Context, args *LookupApiOperationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupApiOperationPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApiOperationPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20220801:getApiOperationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApiOperationPolicyArgs struct {
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
}

// Policy Contract details.
type LookupApiOperationPolicyResult struct {
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

// Defaults sets the appropriate defaults for LookupApiOperationPolicyResult
func (val *LookupApiOperationPolicyResult) Defaults() *LookupApiOperationPolicyResult {
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

func LookupApiOperationPolicyOutput(ctx *pulumi.Context, args LookupApiOperationPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupApiOperationPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiOperationPolicyResult, error) {
			args := v.(LookupApiOperationPolicyArgs)
			r, err := LookupApiOperationPolicy(ctx, &args, opts...)
			var s LookupApiOperationPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiOperationPolicyResultOutput)
}

type LookupApiOperationPolicyOutputArgs struct {
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
}

func (LookupApiOperationPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiOperationPolicyArgs)(nil)).Elem()
}

// Policy Contract details.
type LookupApiOperationPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupApiOperationPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiOperationPolicyResult)(nil)).Elem()
}

func (o LookupApiOperationPolicyResultOutput) ToLookupApiOperationPolicyResultOutput() LookupApiOperationPolicyResultOutput {
	return o
}

func (o LookupApiOperationPolicyResultOutput) ToLookupApiOperationPolicyResultOutputWithContext(ctx context.Context) LookupApiOperationPolicyResultOutput {
	return o
}

func (o LookupApiOperationPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupApiOperationPolicyResult] {
	return pulumix.Output[LookupApiOperationPolicyResult]{
		OutputState: o.OutputState,
	}
}

// Format of the policyContent.
func (o LookupApiOperationPolicyResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiOperationPolicyResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApiOperationPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupApiOperationPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApiOperationPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

// Contents of the Policy as defined by the format.
func (o LookupApiOperationPolicyResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationPolicyResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiOperationPolicyResultOutput{})
}
