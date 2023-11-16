// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a policy fragment.
func LookupWorkspacePolicyFragment(ctx *pulumi.Context, args *LookupWorkspacePolicyFragmentArgs, opts ...pulumi.InvokeOption) (*LookupWorkspacePolicyFragmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspacePolicyFragmentResult
	err := ctx.Invoke("azure-native:apimanagement/v20220901preview:getWorkspacePolicyFragment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWorkspacePolicyFragmentArgs struct {
	// Policy fragment content format.
	Format *string `pulumi:"format"`
	// A resource identifier.
	Id string `pulumi:"id"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// Policy fragment contract details.
type LookupWorkspacePolicyFragmentResult struct {
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

// Defaults sets the appropriate defaults for LookupWorkspacePolicyFragmentResult
func (val *LookupWorkspacePolicyFragmentResult) Defaults() *LookupWorkspacePolicyFragmentResult {
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

func LookupWorkspacePolicyFragmentOutput(ctx *pulumi.Context, args LookupWorkspacePolicyFragmentOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspacePolicyFragmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspacePolicyFragmentResult, error) {
			args := v.(LookupWorkspacePolicyFragmentArgs)
			r, err := LookupWorkspacePolicyFragment(ctx, &args, opts...)
			var s LookupWorkspacePolicyFragmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspacePolicyFragmentResultOutput)
}

type LookupWorkspacePolicyFragmentOutputArgs struct {
	// Policy fragment content format.
	Format pulumi.StringPtrInput `pulumi:"format"`
	// A resource identifier.
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspacePolicyFragmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspacePolicyFragmentArgs)(nil)).Elem()
}

// Policy fragment contract details.
type LookupWorkspacePolicyFragmentResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspacePolicyFragmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspacePolicyFragmentResult)(nil)).Elem()
}

func (o LookupWorkspacePolicyFragmentResultOutput) ToLookupWorkspacePolicyFragmentResultOutput() LookupWorkspacePolicyFragmentResultOutput {
	return o
}

func (o LookupWorkspacePolicyFragmentResultOutput) ToLookupWorkspacePolicyFragmentResultOutputWithContext(ctx context.Context) LookupWorkspacePolicyFragmentResultOutput {
	return o
}

// Policy fragment description.
func (o LookupWorkspacePolicyFragmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspacePolicyFragmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Format of the policy fragment content.
func (o LookupWorkspacePolicyFragmentResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspacePolicyFragmentResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspacePolicyFragmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspacePolicyFragmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspacePolicyFragmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspacePolicyFragmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspacePolicyFragmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspacePolicyFragmentResult) string { return v.Type }).(pulumi.StringOutput)
}

// Contents of the policy fragment.
func (o LookupWorkspacePolicyFragmentResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspacePolicyFragmentResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspacePolicyFragmentResultOutput{})
}
