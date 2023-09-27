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

// Gets the details of the product specified by its identifier.
// Azure REST API version: 2022-09-01-preview.
func LookupWorkspaceProduct(ctx *pulumi.Context, args *LookupWorkspaceProductArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceProductResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceProductResult
	err := ctx.Invoke("azure-native:apimanagement:getWorkspaceProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceProductArgs struct {
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// Product details.
type LookupWorkspaceProductResult struct {
	// whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of false.
	ApprovalRequired *bool `pulumi:"approvalRequired"`
	// Product description. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// Product name.
	DisplayName string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
	State *string `pulumi:"state"`
	// Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
	SubscriptionRequired *bool `pulumi:"subscriptionRequired"`
	// Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of false.
	SubscriptionsLimit *int `pulumi:"subscriptionsLimit"`
	// Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
	Terms *string `pulumi:"terms"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspaceProductOutput(ctx *pulumi.Context, args LookupWorkspaceProductOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceProductResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceProductResult, error) {
			args := v.(LookupWorkspaceProductArgs)
			r, err := LookupWorkspaceProduct(ctx, &args, opts...)
			var s LookupWorkspaceProductResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceProductResultOutput)
}

type LookupWorkspaceProductOutputArgs struct {
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceProductOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceProductArgs)(nil)).Elem()
}

// Product details.
type LookupWorkspaceProductResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceProductResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceProductResult)(nil)).Elem()
}

func (o LookupWorkspaceProductResultOutput) ToLookupWorkspaceProductResultOutput() LookupWorkspaceProductResultOutput {
	return o
}

func (o LookupWorkspaceProductResultOutput) ToLookupWorkspaceProductResultOutputWithContext(ctx context.Context) LookupWorkspaceProductResultOutput {
	return o
}

func (o LookupWorkspaceProductResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkspaceProductResult] {
	return pulumix.Output[LookupWorkspaceProductResult]{
		OutputState: o.OutputState,
	}
}

// whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of false.
func (o LookupWorkspaceProductResultOutput) ApprovalRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceProductResult) *bool { return v.ApprovalRequired }).(pulumi.BoolPtrOutput)
}

// Product description. May include HTML formatting tags.
func (o LookupWorkspaceProductResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceProductResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Product name.
func (o LookupWorkspaceProductResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceProductResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceProductResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceProductResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceProductResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceProductResult) string { return v.Name }).(pulumi.StringOutput)
}

// whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
func (o LookupWorkspaceProductResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceProductResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
func (o LookupWorkspaceProductResultOutput) SubscriptionRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceProductResult) *bool { return v.SubscriptionRequired }).(pulumi.BoolPtrOutput)
}

// Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of false.
func (o LookupWorkspaceProductResultOutput) SubscriptionsLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceProductResult) *int { return v.SubscriptionsLimit }).(pulumi.IntPtrOutput)
}

// Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
func (o LookupWorkspaceProductResultOutput) Terms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceProductResult) *string { return v.Terms }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceProductResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceProductResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceProductResultOutput{})
}
