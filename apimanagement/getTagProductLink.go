// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the product link for the tag.
// Azure REST API version: 2022-09-01-preview.
//
// Other available API versions: 2023-03-01-preview.
func LookupTagProductLink(ctx *pulumi.Context, args *LookupTagProductLinkArgs, opts ...pulumi.InvokeOption) (*LookupTagProductLinkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTagProductLinkResult
	err := ctx.Invoke("azure-native:apimanagement:getTagProductLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagProductLinkArgs struct {
	// Tag-product link identifier. Must be unique in the current API Management service instance.
	ProductLinkId string `pulumi:"productLinkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
}

// Tag-product link details.
type LookupTagProductLinkResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Full resource Id of a product.
	ProductId string `pulumi:"productId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupTagProductLinkOutput(ctx *pulumi.Context, args LookupTagProductLinkOutputArgs, opts ...pulumi.InvokeOption) LookupTagProductLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagProductLinkResult, error) {
			args := v.(LookupTagProductLinkArgs)
			r, err := LookupTagProductLink(ctx, &args, opts...)
			var s LookupTagProductLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagProductLinkResultOutput)
}

type LookupTagProductLinkOutputArgs struct {
	// Tag-product link identifier. Must be unique in the current API Management service instance.
	ProductLinkId pulumi.StringInput `pulumi:"productLinkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringInput `pulumi:"tagId"`
}

func (LookupTagProductLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagProductLinkArgs)(nil)).Elem()
}

// Tag-product link details.
type LookupTagProductLinkResultOutput struct{ *pulumi.OutputState }

func (LookupTagProductLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagProductLinkResult)(nil)).Elem()
}

func (o LookupTagProductLinkResultOutput) ToLookupTagProductLinkResultOutput() LookupTagProductLinkResultOutput {
	return o
}

func (o LookupTagProductLinkResultOutput) ToLookupTagProductLinkResultOutputWithContext(ctx context.Context) LookupTagProductLinkResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupTagProductLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagProductLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupTagProductLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagProductLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

// Full resource Id of a product.
func (o LookupTagProductLinkResultOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagProductLinkResult) string { return v.ProductId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupTagProductLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagProductLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagProductLinkResultOutput{})
}
