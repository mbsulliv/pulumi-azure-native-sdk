// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the group link for the product.
// Azure REST API version: 2022-09-01-preview.
//
// Other available API versions: 2023-03-01-preview.
func LookupProductGroupLink(ctx *pulumi.Context, args *LookupProductGroupLinkArgs, opts ...pulumi.InvokeOption) (*LookupProductGroupLinkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupProductGroupLinkResult
	err := ctx.Invoke("azure-native:apimanagement:getProductGroupLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProductGroupLinkArgs struct {
	// Product-Group link identifier. Must be unique in the current API Management service instance.
	GroupLinkId string `pulumi:"groupLinkId"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Product-group link details.
type LookupProductGroupLinkResult struct {
	// Full resource Id of a group.
	GroupId string `pulumi:"groupId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupProductGroupLinkOutput(ctx *pulumi.Context, args LookupProductGroupLinkOutputArgs, opts ...pulumi.InvokeOption) LookupProductGroupLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProductGroupLinkResult, error) {
			args := v.(LookupProductGroupLinkArgs)
			r, err := LookupProductGroupLink(ctx, &args, opts...)
			var s LookupProductGroupLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProductGroupLinkResultOutput)
}

type LookupProductGroupLinkOutputArgs struct {
	// Product-Group link identifier. Must be unique in the current API Management service instance.
	GroupLinkId pulumi.StringInput `pulumi:"groupLinkId"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupProductGroupLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProductGroupLinkArgs)(nil)).Elem()
}

// Product-group link details.
type LookupProductGroupLinkResultOutput struct{ *pulumi.OutputState }

func (LookupProductGroupLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProductGroupLinkResult)(nil)).Elem()
}

func (o LookupProductGroupLinkResultOutput) ToLookupProductGroupLinkResultOutput() LookupProductGroupLinkResultOutput {
	return o
}

func (o LookupProductGroupLinkResultOutput) ToLookupProductGroupLinkResultOutputWithContext(ctx context.Context) LookupProductGroupLinkResultOutput {
	return o
}

// Full resource Id of a group.
func (o LookupProductGroupLinkResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductGroupLinkResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupProductGroupLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductGroupLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupProductGroupLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductGroupLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupProductGroupLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductGroupLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProductGroupLinkResultOutput{})
}
