// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the operation link for the tag.
// Azure REST API version: 2022-09-01-preview.
func LookupWorkspaceTagOperationLink(ctx *pulumi.Context, args *LookupWorkspaceTagOperationLinkArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceTagOperationLinkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceTagOperationLinkResult
	err := ctx.Invoke("azure-native:apimanagement:getWorkspaceTagOperationLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceTagOperationLinkArgs struct {
	// Tag-operation link identifier. Must be unique in the current API Management service instance.
	OperationLinkId string `pulumi:"operationLinkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// Tag-operation link details.
type LookupWorkspaceTagOperationLinkResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Full resource Id of an API operation.
	OperationId string `pulumi:"operationId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspaceTagOperationLinkOutput(ctx *pulumi.Context, args LookupWorkspaceTagOperationLinkOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceTagOperationLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceTagOperationLinkResult, error) {
			args := v.(LookupWorkspaceTagOperationLinkArgs)
			r, err := LookupWorkspaceTagOperationLink(ctx, &args, opts...)
			var s LookupWorkspaceTagOperationLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceTagOperationLinkResultOutput)
}

type LookupWorkspaceTagOperationLinkOutputArgs struct {
	// Tag-operation link identifier. Must be unique in the current API Management service instance.
	OperationLinkId pulumi.StringInput `pulumi:"operationLinkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringInput `pulumi:"tagId"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceTagOperationLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceTagOperationLinkArgs)(nil)).Elem()
}

// Tag-operation link details.
type LookupWorkspaceTagOperationLinkResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceTagOperationLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceTagOperationLinkResult)(nil)).Elem()
}

func (o LookupWorkspaceTagOperationLinkResultOutput) ToLookupWorkspaceTagOperationLinkResultOutput() LookupWorkspaceTagOperationLinkResultOutput {
	return o
}

func (o LookupWorkspaceTagOperationLinkResultOutput) ToLookupWorkspaceTagOperationLinkResultOutputWithContext(ctx context.Context) LookupWorkspaceTagOperationLinkResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceTagOperationLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceTagOperationLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceTagOperationLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceTagOperationLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

// Full resource Id of an API operation.
func (o LookupWorkspaceTagOperationLinkResultOutput) OperationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceTagOperationLinkResult) string { return v.OperationId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceTagOperationLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceTagOperationLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceTagOperationLinkResultOutput{})
}