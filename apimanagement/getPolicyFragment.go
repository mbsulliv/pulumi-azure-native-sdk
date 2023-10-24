// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a policy fragment.
// Azure REST API version: 2022-08-01.
//
// Other available API versions: 2022-09-01-preview, 2023-03-01-preview.
func LookupPolicyFragment(ctx *pulumi.Context, args *LookupPolicyFragmentArgs, opts ...pulumi.InvokeOption) (*LookupPolicyFragmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyFragmentResult
	err := ctx.Invoke("azure-native:apimanagement:getPolicyFragment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPolicyFragmentArgs struct {
	// Policy fragment content format.
	Format *string `pulumi:"format"`
	// A resource identifier.
	Id string `pulumi:"id"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Policy fragment contract details.
type LookupPolicyFragmentResult struct {
	// Policy fragment description.
	Description *string `pulumi:"description"`
	// Format of the policy fragment content.
	Format *string `pulumi:"format"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Contents of the policy fragment.
	Value string `pulumi:"value"`
}

// Defaults sets the appropriate defaults for LookupPolicyFragmentResult
func (val *LookupPolicyFragmentResult) Defaults() *LookupPolicyFragmentResult {
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

func LookupPolicyFragmentOutput(ctx *pulumi.Context, args LookupPolicyFragmentOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyFragmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyFragmentResult, error) {
			args := v.(LookupPolicyFragmentArgs)
			r, err := LookupPolicyFragment(ctx, &args, opts...)
			var s LookupPolicyFragmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyFragmentResultOutput)
}

type LookupPolicyFragmentOutputArgs struct {
	// Policy fragment content format.
	Format pulumi.StringPtrInput `pulumi:"format"`
	// A resource identifier.
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupPolicyFragmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyFragmentArgs)(nil)).Elem()
}

// Policy fragment contract details.
type LookupPolicyFragmentResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyFragmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyFragmentResult)(nil)).Elem()
}

func (o LookupPolicyFragmentResultOutput) ToLookupPolicyFragmentResultOutput() LookupPolicyFragmentResultOutput {
	return o
}

func (o LookupPolicyFragmentResultOutput) ToLookupPolicyFragmentResultOutputWithContext(ctx context.Context) LookupPolicyFragmentResultOutput {
	return o
}

func (o LookupPolicyFragmentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPolicyFragmentResult] {
	return pulumix.Output[LookupPolicyFragmentResult]{
		OutputState: o.OutputState,
	}
}

// Policy fragment description.
func (o LookupPolicyFragmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyFragmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Format of the policy fragment content.
func (o LookupPolicyFragmentResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyFragmentResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPolicyFragmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyFragmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPolicyFragmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyFragmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPolicyFragmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyFragmentResult) string { return v.Type }).(pulumi.StringOutput)
}

// Contents of the policy fragment.
func (o LookupPolicyFragmentResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyFragmentResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyFragmentResultOutput{})
}
