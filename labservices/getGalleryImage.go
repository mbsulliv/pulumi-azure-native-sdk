// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package labservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get gallery image
// Azure REST API version: 2018-10-15.
func LookupGalleryImage(ctx *pulumi.Context, args *LookupGalleryImageArgs, opts ...pulumi.InvokeOption) (*LookupGalleryImageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGalleryImageResult
	err := ctx.Invoke("azure-native:labservices:getGalleryImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryImageArgs struct {
	// Specify the $expand query. Example: 'properties($select=author)'
	Expand *string `pulumi:"expand"`
	// The name of the gallery Image.
	GalleryImageName string `pulumi:"galleryImageName"`
	// The name of the lab Account.
	LabAccountName string `pulumi:"labAccountName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents an image from the Azure Marketplace
type LookupGalleryImageResult struct {
	// The author of the gallery image.
	Author string `pulumi:"author"`
	// The creation date of the gallery image.
	CreatedDate string `pulumi:"createdDate"`
	// The description of the gallery image.
	Description string `pulumi:"description"`
	// The icon of the gallery image.
	Icon string `pulumi:"icon"`
	// The identifier of the resource.
	Id string `pulumi:"id"`
	// The image reference of the gallery image.
	ImageReference GalleryImageReferenceResponse `pulumi:"imageReference"`
	// Indicates whether this gallery image is enabled.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Indicates whether this gallery has been overridden for this lab account
	IsOverride *bool `pulumi:"isOverride"`
	// Indicates if the plan has been authorized for programmatic deployment.
	IsPlanAuthorized *bool `pulumi:"isPlanAuthorized"`
	// The details of the latest operation. ex: status, error
	LatestOperationResult LatestOperationResultResponse `pulumi:"latestOperationResult"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The third party plan that applies to this image
	PlanId string `pulumi:"planId"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

func LookupGalleryImageOutput(ctx *pulumi.Context, args LookupGalleryImageOutputArgs, opts ...pulumi.InvokeOption) LookupGalleryImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGalleryImageResult, error) {
			args := v.(LookupGalleryImageArgs)
			r, err := LookupGalleryImage(ctx, &args, opts...)
			var s LookupGalleryImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGalleryImageResultOutput)
}

type LookupGalleryImageOutputArgs struct {
	// Specify the $expand query. Example: 'properties($select=author)'
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the gallery Image.
	GalleryImageName pulumi.StringInput `pulumi:"galleryImageName"`
	// The name of the lab Account.
	LabAccountName pulumi.StringInput `pulumi:"labAccountName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGalleryImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryImageArgs)(nil)).Elem()
}

// Represents an image from the Azure Marketplace
type LookupGalleryImageResultOutput struct{ *pulumi.OutputState }

func (LookupGalleryImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryImageResult)(nil)).Elem()
}

func (o LookupGalleryImageResultOutput) ToLookupGalleryImageResultOutput() LookupGalleryImageResultOutput {
	return o
}

func (o LookupGalleryImageResultOutput) ToLookupGalleryImageResultOutputWithContext(ctx context.Context) LookupGalleryImageResultOutput {
	return o
}

func (o LookupGalleryImageResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGalleryImageResult] {
	return pulumix.Output[LookupGalleryImageResult]{
		OutputState: o.OutputState,
	}
}

// The author of the gallery image.
func (o LookupGalleryImageResultOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Author }).(pulumi.StringOutput)
}

// The creation date of the gallery image.
func (o LookupGalleryImageResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

// The description of the gallery image.
func (o LookupGalleryImageResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Description }).(pulumi.StringOutput)
}

// The icon of the gallery image.
func (o LookupGalleryImageResultOutput) Icon() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Icon }).(pulumi.StringOutput)
}

// The identifier of the resource.
func (o LookupGalleryImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Id }).(pulumi.StringOutput)
}

// The image reference of the gallery image.
func (o LookupGalleryImageResultOutput) ImageReference() GalleryImageReferenceResponseOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) GalleryImageReferenceResponse { return v.ImageReference }).(GalleryImageReferenceResponseOutput)
}

// Indicates whether this gallery image is enabled.
func (o LookupGalleryImageResultOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// Indicates whether this gallery has been overridden for this lab account
func (o LookupGalleryImageResultOutput) IsOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *bool { return v.IsOverride }).(pulumi.BoolPtrOutput)
}

// Indicates if the plan has been authorized for programmatic deployment.
func (o LookupGalleryImageResultOutput) IsPlanAuthorized() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *bool { return v.IsPlanAuthorized }).(pulumi.BoolPtrOutput)
}

// The details of the latest operation. ex: status, error
func (o LookupGalleryImageResultOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) LatestOperationResultResponse { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

// The location of the resource.
func (o LookupGalleryImageResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupGalleryImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Name }).(pulumi.StringOutput)
}

// The third party plan that applies to this image
func (o LookupGalleryImageResultOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.PlanId }).(pulumi.StringOutput)
}

// The provisioning status of the resource.
func (o LookupGalleryImageResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The tags of the resource.
func (o LookupGalleryImageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupGalleryImageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o LookupGalleryImageResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGalleryImageResultOutput{})
}
