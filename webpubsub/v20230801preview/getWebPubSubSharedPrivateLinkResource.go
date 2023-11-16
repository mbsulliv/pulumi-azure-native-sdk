// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the specified shared private link resource
func LookupWebPubSubSharedPrivateLinkResource(ctx *pulumi.Context, args *LookupWebPubSubSharedPrivateLinkResourceArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubSharedPrivateLinkResourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebPubSubSharedPrivateLinkResourceResult
	err := ctx.Invoke("azure-native:webpubsub/v20230801preview:getWebPubSubSharedPrivateLinkResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebPubSubSharedPrivateLinkResourceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
	// The name of the shared private link resource
	SharedPrivateLinkResourceName string `pulumi:"sharedPrivateLinkResourceName"`
}

// Describes a Shared Private Link Resource
type LookupWebPubSubSharedPrivateLinkResourceResult struct {
	// The group id from the provider of resource the shared private link resource is for
	GroupId string `pulumi:"groupId"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource id of the resource the shared private link resource is for
	PrivateLinkResourceId string `pulumi:"privateLinkResourceId"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The request message for requesting approval of the shared private link resource
	RequestMessage *string `pulumi:"requestMessage"`
	// Status of the shared private link resource
	Status string `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWebPubSubSharedPrivateLinkResourceOutput(ctx *pulumi.Context, args LookupWebPubSubSharedPrivateLinkResourceOutputArgs, opts ...pulumi.InvokeOption) LookupWebPubSubSharedPrivateLinkResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebPubSubSharedPrivateLinkResourceResult, error) {
			args := v.(LookupWebPubSubSharedPrivateLinkResourceArgs)
			r, err := LookupWebPubSubSharedPrivateLinkResource(ctx, &args, opts...)
			var s LookupWebPubSubSharedPrivateLinkResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebPubSubSharedPrivateLinkResourceResultOutput)
}

type LookupWebPubSubSharedPrivateLinkResourceOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// The name of the shared private link resource
	SharedPrivateLinkResourceName pulumi.StringInput `pulumi:"sharedPrivateLinkResourceName"`
}

func (LookupWebPubSubSharedPrivateLinkResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubSharedPrivateLinkResourceArgs)(nil)).Elem()
}

// Describes a Shared Private Link Resource
type LookupWebPubSubSharedPrivateLinkResourceResultOutput struct{ *pulumi.OutputState }

func (LookupWebPubSubSharedPrivateLinkResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubSharedPrivateLinkResourceResult)(nil)).Elem()
}

func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) ToLookupWebPubSubSharedPrivateLinkResourceResultOutput() LookupWebPubSubSharedPrivateLinkResourceResultOutput {
	return o
}

func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) ToLookupWebPubSubSharedPrivateLinkResourceResultOutputWithContext(ctx context.Context) LookupWebPubSubSharedPrivateLinkResourceResultOutput {
	return o
}

// The group id from the provider of resource the shared private link resource is for
func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource id of the resource the shared private link resource is for
func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) PrivateLinkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.PrivateLinkResourceId }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The request message for requesting approval of the shared private link resource
func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

// Status of the shared private link resource
func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWebPubSubSharedPrivateLinkResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubSharedPrivateLinkResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebPubSubSharedPrivateLinkResourceResultOutput{})
}
