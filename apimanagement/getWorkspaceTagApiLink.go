// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the API link for the tag.
// Azure REST API version: 2022-09-01-preview.
func LookupWorkspaceTagApiLink(ctx *pulumi.Context, args *LookupWorkspaceTagApiLinkArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceTagApiLinkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceTagApiLinkResult
	err := ctx.Invoke("azure-native:apimanagement:getWorkspaceTagApiLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceTagApiLinkArgs struct {
	// Tag-API link identifier. Must be unique in the current API Management service instance.
	ApiLinkId string `pulumi:"apiLinkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// Tag-API link details.
type LookupWorkspaceTagApiLinkResult struct {
	// Full resource Id of an API.
	ApiId string `pulumi:"apiId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspaceTagApiLinkOutput(ctx *pulumi.Context, args LookupWorkspaceTagApiLinkOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceTagApiLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceTagApiLinkResult, error) {
			args := v.(LookupWorkspaceTagApiLinkArgs)
			r, err := LookupWorkspaceTagApiLink(ctx, &args, opts...)
			var s LookupWorkspaceTagApiLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceTagApiLinkResultOutput)
}

type LookupWorkspaceTagApiLinkOutputArgs struct {
	// Tag-API link identifier. Must be unique in the current API Management service instance.
	ApiLinkId pulumi.StringInput `pulumi:"apiLinkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringInput `pulumi:"tagId"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceTagApiLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceTagApiLinkArgs)(nil)).Elem()
}

// Tag-API link details.
type LookupWorkspaceTagApiLinkResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceTagApiLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceTagApiLinkResult)(nil)).Elem()
}

func (o LookupWorkspaceTagApiLinkResultOutput) ToLookupWorkspaceTagApiLinkResultOutput() LookupWorkspaceTagApiLinkResultOutput {
	return o
}

func (o LookupWorkspaceTagApiLinkResultOutput) ToLookupWorkspaceTagApiLinkResultOutputWithContext(ctx context.Context) LookupWorkspaceTagApiLinkResultOutput {
	return o
}

// Full resource Id of an API.
func (o LookupWorkspaceTagApiLinkResultOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceTagApiLinkResult) string { return v.ApiId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceTagApiLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceTagApiLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceTagApiLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceTagApiLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceTagApiLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceTagApiLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceTagApiLinkResultOutput{})
}