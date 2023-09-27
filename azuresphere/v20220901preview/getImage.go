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

// Get a Image
func LookupImage(ctx *pulumi.Context, args *LookupImageArgs, opts ...pulumi.InvokeOption) (*LookupImageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupImageResult
	err := ctx.Invoke("azure-native:azuresphere/v20220901preview:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupImageArgs struct {
	// Name of catalog
	CatalogName string `pulumi:"catalogName"`
	// Image name. Use .default for image creation.
	ImageName string `pulumi:"imageName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An image resource belonging to a catalog resource.
type LookupImageResult struct {
	// The image component id.
	ComponentId string `pulumi:"componentId"`
	// The image description.
	Description string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Image as a UTF-8 encoded base 64 string on image create. This field contains the image URI on image reads.
	Image *string `pulumi:"image"`
	// Image ID
	ImageId *string `pulumi:"imageId"`
	// Image name
	ImageName string `pulumi:"imageName"`
	// The image type.
	ImageType string `pulumi:"imageType"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// Regional data boundary for an image
	RegionalDataBoundary *string `pulumi:"regionalDataBoundary"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Location the image
	Uri string `pulumi:"uri"`
}

func LookupImageOutput(ctx *pulumi.Context, args LookupImageOutputArgs, opts ...pulumi.InvokeOption) LookupImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImageResult, error) {
			args := v.(LookupImageArgs)
			r, err := LookupImage(ctx, &args, opts...)
			var s LookupImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImageResultOutput)
}

type LookupImageOutputArgs struct {
	// Name of catalog
	CatalogName pulumi.StringInput `pulumi:"catalogName"`
	// Image name. Use .default for image creation.
	ImageName pulumi.StringInput `pulumi:"imageName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageArgs)(nil)).Elem()
}

// An image resource belonging to a catalog resource.
type LookupImageResultOutput struct{ *pulumi.OutputState }

func (LookupImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageResult)(nil)).Elem()
}

func (o LookupImageResultOutput) ToLookupImageResultOutput() LookupImageResultOutput {
	return o
}

func (o LookupImageResultOutput) ToLookupImageResultOutputWithContext(ctx context.Context) LookupImageResultOutput {
	return o
}

func (o LookupImageResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupImageResult] {
	return pulumix.Output[LookupImageResult]{
		OutputState: o.OutputState,
	}
}

// The image component id.
func (o LookupImageResultOutput) ComponentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ComponentId }).(pulumi.StringOutput)
}

// The image description.
func (o LookupImageResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Description }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Id }).(pulumi.StringOutput)
}

// Image as a UTF-8 encoded base 64 string on image create. This field contains the image URI on image reads.
func (o LookupImageResultOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.Image }).(pulumi.StringPtrOutput)
}

// Image ID
func (o LookupImageResultOutput) ImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.ImageId }).(pulumi.StringPtrOutput)
}

// Image name
func (o LookupImageResultOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ImageName }).(pulumi.StringOutput)
}

// The image type.
func (o LookupImageResultOutput) ImageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ImageType }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupImageResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Regional data boundary for an image
func (o LookupImageResultOutput) RegionalDataBoundary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.RegionalDataBoundary }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupImageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupImageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupImageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Type }).(pulumi.StringOutput)
}

// Location the image
func (o LookupImageResultOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Uri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImageResultOutput{})
}
