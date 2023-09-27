// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the group link for the product.
func LookupWorkspaceProductGroupLink(ctx *pulumi.Context, args *LookupWorkspaceProductGroupLinkArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceProductGroupLinkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceProductGroupLinkResult
	err := ctx.Invoke("azure-native:apimanagement/v20220901preview:getWorkspaceProductGroupLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceProductGroupLinkArgs struct {
	// Product-Group link identifier. Must be unique in the current API Management service instance.
	GroupLinkId string `pulumi:"groupLinkId"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// Product-group link details.
type LookupWorkspaceProductGroupLinkResult struct {
	// Full resource Id of a group.
	GroupId string `pulumi:"groupId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspaceProductGroupLinkOutput(ctx *pulumi.Context, args LookupWorkspaceProductGroupLinkOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceProductGroupLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceProductGroupLinkResult, error) {
			args := v.(LookupWorkspaceProductGroupLinkArgs)
			r, err := LookupWorkspaceProductGroupLink(ctx, &args, opts...)
			var s LookupWorkspaceProductGroupLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceProductGroupLinkResultOutput)
}

type LookupWorkspaceProductGroupLinkOutputArgs struct {
	// Product-Group link identifier. Must be unique in the current API Management service instance.
	GroupLinkId pulumi.StringInput `pulumi:"groupLinkId"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceProductGroupLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceProductGroupLinkArgs)(nil)).Elem()
}

// Product-group link details.
type LookupWorkspaceProductGroupLinkResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceProductGroupLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceProductGroupLinkResult)(nil)).Elem()
}

func (o LookupWorkspaceProductGroupLinkResultOutput) ToLookupWorkspaceProductGroupLinkResultOutput() LookupWorkspaceProductGroupLinkResultOutput {
	return o
}

func (o LookupWorkspaceProductGroupLinkResultOutput) ToLookupWorkspaceProductGroupLinkResultOutputWithContext(ctx context.Context) LookupWorkspaceProductGroupLinkResultOutput {
	return o
}

func (o LookupWorkspaceProductGroupLinkResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkspaceProductGroupLinkResult] {
	return pulumix.Output[LookupWorkspaceProductGroupLinkResult]{
		OutputState: o.OutputState,
	}
}

// Full resource Id of a group.
func (o LookupWorkspaceProductGroupLinkResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceProductGroupLinkResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceProductGroupLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceProductGroupLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceProductGroupLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceProductGroupLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceProductGroupLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceProductGroupLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceProductGroupLinkResultOutput{})
}
