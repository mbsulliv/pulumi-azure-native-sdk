// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220303

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about a gallery image version.
func LookupGalleryImageVersion(ctx *pulumi.Context, args *LookupGalleryImageVersionArgs, opts ...pulumi.InvokeOption) (*LookupGalleryImageVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGalleryImageVersionResult
	err := ctx.Invoke("azure-native:compute/v20220303:getGalleryImageVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryImageVersionArgs struct {
	// The expand expression to apply on the operation.
	Expand *string `pulumi:"expand"`
	// The name of the gallery image definition in which the Image Version resides.
	GalleryImageName string `pulumi:"galleryImageName"`
	// The name of the gallery image version to be retrieved.
	GalleryImageVersionName string `pulumi:"galleryImageVersionName"`
	// The name of the Shared Image Gallery in which the Image Definition resides.
	GalleryName string `pulumi:"galleryName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Specifies information about the gallery image version that you want to create or update.
type LookupGalleryImageVersionResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// The provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// The publishing profile of a gallery image Version.
	PublishingProfile *GalleryImageVersionPublishingProfileResponse `pulumi:"publishingProfile"`
	// This is the replication status of the gallery image version.
	ReplicationStatus ReplicationStatusResponse `pulumi:"replicationStatus"`
	// This is the safety profile of the Gallery Image Version.
	SafetyProfile *GalleryImageVersionSafetyProfileResponse `pulumi:"safetyProfile"`
	// This is the storage profile of a Gallery Image Version.
	StorageProfile GalleryImageVersionStorageProfileResponse `pulumi:"storageProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupGalleryImageVersionOutput(ctx *pulumi.Context, args LookupGalleryImageVersionOutputArgs, opts ...pulumi.InvokeOption) LookupGalleryImageVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGalleryImageVersionResult, error) {
			args := v.(LookupGalleryImageVersionArgs)
			r, err := LookupGalleryImageVersion(ctx, &args, opts...)
			var s LookupGalleryImageVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGalleryImageVersionResultOutput)
}

type LookupGalleryImageVersionOutputArgs struct {
	// The expand expression to apply on the operation.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the gallery image definition in which the Image Version resides.
	GalleryImageName pulumi.StringInput `pulumi:"galleryImageName"`
	// The name of the gallery image version to be retrieved.
	GalleryImageVersionName pulumi.StringInput `pulumi:"galleryImageVersionName"`
	// The name of the Shared Image Gallery in which the Image Definition resides.
	GalleryName pulumi.StringInput `pulumi:"galleryName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGalleryImageVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryImageVersionArgs)(nil)).Elem()
}

// Specifies information about the gallery image version that you want to create or update.
type LookupGalleryImageVersionResultOutput struct{ *pulumi.OutputState }

func (LookupGalleryImageVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryImageVersionResult)(nil)).Elem()
}

func (o LookupGalleryImageVersionResultOutput) ToLookupGalleryImageVersionResultOutput() LookupGalleryImageVersionResultOutput {
	return o
}

func (o LookupGalleryImageVersionResultOutput) ToLookupGalleryImageVersionResultOutputWithContext(ctx context.Context) LookupGalleryImageVersionResultOutput {
	return o
}

// Resource Id
func (o LookupGalleryImageVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupGalleryImageVersionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupGalleryImageVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state, which only appears in the response.
func (o LookupGalleryImageVersionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The publishing profile of a gallery image Version.
func (o LookupGalleryImageVersionResultOutput) PublishingProfile() GalleryImageVersionPublishingProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) *GalleryImageVersionPublishingProfileResponse {
		return v.PublishingProfile
	}).(GalleryImageVersionPublishingProfileResponsePtrOutput)
}

// This is the replication status of the gallery image version.
func (o LookupGalleryImageVersionResultOutput) ReplicationStatus() ReplicationStatusResponseOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) ReplicationStatusResponse { return v.ReplicationStatus }).(ReplicationStatusResponseOutput)
}

// This is the safety profile of the Gallery Image Version.
func (o LookupGalleryImageVersionResultOutput) SafetyProfile() GalleryImageVersionSafetyProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) *GalleryImageVersionSafetyProfileResponse {
		return v.SafetyProfile
	}).(GalleryImageVersionSafetyProfileResponsePtrOutput)
}

// This is the storage profile of a Gallery Image Version.
func (o LookupGalleryImageVersionResultOutput) StorageProfile() GalleryImageVersionStorageProfileResponseOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) GalleryImageVersionStorageProfileResponse {
		return v.StorageProfile
	}).(GalleryImageVersionStorageProfileResponseOutput)
}

// Resource tags
func (o LookupGalleryImageVersionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupGalleryImageVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGalleryImageVersionResultOutput{})
}
