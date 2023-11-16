// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the product link for the tag.
func LookupWorkspaceTagProductLink(ctx *pulumi.Context, args *LookupWorkspaceTagProductLinkArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceTagProductLinkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceTagProductLinkResult
	err := ctx.Invoke("azure-native:apimanagement/v20220901preview:getWorkspaceTagProductLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceTagProductLinkArgs struct {
	// Tag-product link identifier. Must be unique in the current API Management service instance.
	ProductLinkId string `pulumi:"productLinkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// Tag-product link details.
type LookupWorkspaceTagProductLinkResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Full resource Id of a product.
	ProductId string `pulumi:"productId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspaceTagProductLinkOutput(ctx *pulumi.Context, args LookupWorkspaceTagProductLinkOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceTagProductLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceTagProductLinkResult, error) {
			args := v.(LookupWorkspaceTagProductLinkArgs)
			r, err := LookupWorkspaceTagProductLink(ctx, &args, opts...)
			var s LookupWorkspaceTagProductLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceTagProductLinkResultOutput)
}

type LookupWorkspaceTagProductLinkOutputArgs struct {
	// Tag-product link identifier. Must be unique in the current API Management service instance.
	ProductLinkId pulumi.StringInput `pulumi:"productLinkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringInput `pulumi:"tagId"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceTagProductLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceTagProductLinkArgs)(nil)).Elem()
}

// Tag-product link details.
type LookupWorkspaceTagProductLinkResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceTagProductLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceTagProductLinkResult)(nil)).Elem()
}

func (o LookupWorkspaceTagProductLinkResultOutput) ToLookupWorkspaceTagProductLinkResultOutput() LookupWorkspaceTagProductLinkResultOutput {
	return o
}

func (o LookupWorkspaceTagProductLinkResultOutput) ToLookupWorkspaceTagProductLinkResultOutputWithContext(ctx context.Context) LookupWorkspaceTagProductLinkResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceTagProductLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceTagProductLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceTagProductLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceTagProductLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

// Full resource Id of a product.
func (o LookupWorkspaceTagProductLinkResultOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceTagProductLinkResult) string { return v.ProductId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceTagProductLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceTagProductLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceTagProductLinkResultOutput{})
}
