// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a gallery
func LookupGallery(ctx *pulumi.Context, args *LookupGalleryArgs, opts ...pulumi.InvokeOption) (*LookupGalleryResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGalleryResult
	err := ctx.Invoke("azure-native:devcenter/v20230801preview:getGallery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryArgs struct {
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// The name of the gallery.
	GalleryName string `pulumi:"galleryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a gallery.
type LookupGalleryResult struct {
	// The resource ID of the backing Azure Compute Gallery.
	GalleryResourceId string `pulumi:"galleryResourceId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupGalleryOutput(ctx *pulumi.Context, args LookupGalleryOutputArgs, opts ...pulumi.InvokeOption) LookupGalleryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGalleryResult, error) {
			args := v.(LookupGalleryArgs)
			r, err := LookupGallery(ctx, &args, opts...)
			var s LookupGalleryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGalleryResultOutput)
}

type LookupGalleryOutputArgs struct {
	// The name of the devcenter.
	DevCenterName pulumi.StringInput `pulumi:"devCenterName"`
	// The name of the gallery.
	GalleryName pulumi.StringInput `pulumi:"galleryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGalleryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryArgs)(nil)).Elem()
}

// Represents a gallery.
type LookupGalleryResultOutput struct{ *pulumi.OutputState }

func (LookupGalleryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryResult)(nil)).Elem()
}

func (o LookupGalleryResultOutput) ToLookupGalleryResultOutput() LookupGalleryResultOutput {
	return o
}

func (o LookupGalleryResultOutput) ToLookupGalleryResultOutputWithContext(ctx context.Context) LookupGalleryResultOutput {
	return o
}

// The resource ID of the backing Azure Compute Gallery.
func (o LookupGalleryResultOutput) GalleryResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.GalleryResourceId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupGalleryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupGalleryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupGalleryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupGalleryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGalleryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupGalleryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGalleryResultOutput{})
}
