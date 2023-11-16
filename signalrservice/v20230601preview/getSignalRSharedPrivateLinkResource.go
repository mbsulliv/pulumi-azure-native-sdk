// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the specified shared private link resource
func LookupSignalRSharedPrivateLinkResource(ctx *pulumi.Context, args *LookupSignalRSharedPrivateLinkResourceArgs, opts ...pulumi.InvokeOption) (*LookupSignalRSharedPrivateLinkResourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSignalRSharedPrivateLinkResourceResult
	err := ctx.Invoke("azure-native:signalrservice/v20230601preview:getSignalRSharedPrivateLinkResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalRSharedPrivateLinkResourceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
	// The name of the shared private link resource
	SharedPrivateLinkResourceName string `pulumi:"sharedPrivateLinkResourceName"`
}

// Describes a Shared Private Link Resource
type LookupSignalRSharedPrivateLinkResourceResult struct {
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

func LookupSignalRSharedPrivateLinkResourceOutput(ctx *pulumi.Context, args LookupSignalRSharedPrivateLinkResourceOutputArgs, opts ...pulumi.InvokeOption) LookupSignalRSharedPrivateLinkResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSignalRSharedPrivateLinkResourceResult, error) {
			args := v.(LookupSignalRSharedPrivateLinkResourceArgs)
			r, err := LookupSignalRSharedPrivateLinkResource(ctx, &args, opts...)
			var s LookupSignalRSharedPrivateLinkResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSignalRSharedPrivateLinkResourceResultOutput)
}

type LookupSignalRSharedPrivateLinkResourceOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// The name of the shared private link resource
	SharedPrivateLinkResourceName pulumi.StringInput `pulumi:"sharedPrivateLinkResourceName"`
}

func (LookupSignalRSharedPrivateLinkResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRSharedPrivateLinkResourceArgs)(nil)).Elem()
}

// Describes a Shared Private Link Resource
type LookupSignalRSharedPrivateLinkResourceResultOutput struct{ *pulumi.OutputState }

func (LookupSignalRSharedPrivateLinkResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRSharedPrivateLinkResourceResult)(nil)).Elem()
}

func (o LookupSignalRSharedPrivateLinkResourceResultOutput) ToLookupSignalRSharedPrivateLinkResourceResultOutput() LookupSignalRSharedPrivateLinkResourceResultOutput {
	return o
}

func (o LookupSignalRSharedPrivateLinkResourceResultOutput) ToLookupSignalRSharedPrivateLinkResourceResultOutputWithContext(ctx context.Context) LookupSignalRSharedPrivateLinkResourceResultOutput {
	return o
}

// The group id from the provider of resource the shared private link resource is for
func (o LookupSignalRSharedPrivateLinkResourceResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupSignalRSharedPrivateLinkResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSignalRSharedPrivateLinkResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource id of the resource the shared private link resource is for
func (o LookupSignalRSharedPrivateLinkResourceResultOutput) PrivateLinkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.PrivateLinkResourceId }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupSignalRSharedPrivateLinkResourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The request message for requesting approval of the shared private link resource
func (o LookupSignalRSharedPrivateLinkResourceResultOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

// Status of the shared private link resource
func (o LookupSignalRSharedPrivateLinkResourceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSignalRSharedPrivateLinkResourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSignalRSharedPrivateLinkResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRSharedPrivateLinkResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalRSharedPrivateLinkResourceResultOutput{})
}
