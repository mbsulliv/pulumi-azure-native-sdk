// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the shared private link resource managed by the search service in the given resource group.
func LookupSharedPrivateLinkResource(ctx *pulumi.Context, args *LookupSharedPrivateLinkResourceArgs, opts ...pulumi.InvokeOption) (*LookupSharedPrivateLinkResourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSharedPrivateLinkResourceResult
	err := ctx.Invoke("azure-native:search/v20220901:getSharedPrivateLinkResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSharedPrivateLinkResourceArgs struct {
	// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Azure Cognitive Search service associated with the specified resource group.
	SearchServiceName string `pulumi:"searchServiceName"`
	// The name of the shared private link resource managed by the Azure Cognitive Search service within the specified resource group.
	SharedPrivateLinkResourceName string `pulumi:"sharedPrivateLinkResourceName"`
}

// Describes a Shared Private Link Resource managed by the Azure Cognitive Search service.
type LookupSharedPrivateLinkResourceResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Describes the properties of a Shared Private Link Resource managed by the Azure Cognitive Search service.
	Properties SharedPrivateLinkResourcePropertiesResponse `pulumi:"properties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSharedPrivateLinkResourceOutput(ctx *pulumi.Context, args LookupSharedPrivateLinkResourceOutputArgs, opts ...pulumi.InvokeOption) LookupSharedPrivateLinkResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSharedPrivateLinkResourceResult, error) {
			args := v.(LookupSharedPrivateLinkResourceArgs)
			r, err := LookupSharedPrivateLinkResource(ctx, &args, opts...)
			var s LookupSharedPrivateLinkResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSharedPrivateLinkResourceResultOutput)
}

type LookupSharedPrivateLinkResourceOutputArgs struct {
	// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Azure Cognitive Search service associated with the specified resource group.
	SearchServiceName pulumi.StringInput `pulumi:"searchServiceName"`
	// The name of the shared private link resource managed by the Azure Cognitive Search service within the specified resource group.
	SharedPrivateLinkResourceName pulumi.StringInput `pulumi:"sharedPrivateLinkResourceName"`
}

func (LookupSharedPrivateLinkResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSharedPrivateLinkResourceArgs)(nil)).Elem()
}

// Describes a Shared Private Link Resource managed by the Azure Cognitive Search service.
type LookupSharedPrivateLinkResourceResultOutput struct{ *pulumi.OutputState }

func (LookupSharedPrivateLinkResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSharedPrivateLinkResourceResult)(nil)).Elem()
}

func (o LookupSharedPrivateLinkResourceResultOutput) ToLookupSharedPrivateLinkResourceResultOutput() LookupSharedPrivateLinkResourceResultOutput {
	return o
}

func (o LookupSharedPrivateLinkResourceResultOutput) ToLookupSharedPrivateLinkResourceResultOutputWithContext(ctx context.Context) LookupSharedPrivateLinkResourceResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSharedPrivateLinkResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSharedPrivateLinkResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSharedPrivateLinkResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSharedPrivateLinkResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Describes the properties of a Shared Private Link Resource managed by the Azure Cognitive Search service.
func (o LookupSharedPrivateLinkResourceResultOutput) Properties() SharedPrivateLinkResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupSharedPrivateLinkResourceResult) SharedPrivateLinkResourcePropertiesResponse {
		return v.Properties
	}).(SharedPrivateLinkResourcePropertiesResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSharedPrivateLinkResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSharedPrivateLinkResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSharedPrivateLinkResourceResultOutput{})
}
