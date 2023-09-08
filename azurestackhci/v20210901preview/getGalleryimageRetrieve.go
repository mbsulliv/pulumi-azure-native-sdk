// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets galleryimages by resource name
func LookupGalleryimageRetrieve(ctx *pulumi.Context, args *LookupGalleryimageRetrieveArgs, opts ...pulumi.InvokeOption) (*LookupGalleryimageRetrieveResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGalleryimageRetrieveResult
	err := ctx.Invoke("azure-native:azurestackhci/v20210901preview:getGalleryimageRetrieve", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryimageRetrieveArgs struct {
	// Name of the gallery image
	GalleryimagesName string `pulumi:"galleryimagesName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The gallery image resource definition.
type LookupGalleryimageRetrieveResult struct {
	// Datasource for the gallery image when provisioning with cloud-init (Azure or NoCloud)
	CloudInitDataSource *string `pulumi:"cloudInitDataSource"`
	// Container Name for storage container
	ContainerName *string `pulumi:"containerName"`
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine [V1, V2]
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// This is the gallery image definition identifier.
	Identifier *GalleryImageIdentifierResponse `pulumi:"identifier"`
	// location of the image the gallery image should be created from
	ImagePath *string `pulumi:"imagePath"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// operating system type that the gallery image uses. Expected to be linux or windows
	OsType *string `pulumi:"osType"`
	// Provisioning state of the gallery image.
	ProvisioningState string `pulumi:"provisioningState"`
	// name of the object to be used in moc
	ResourceName *string `pulumi:"resourceName"`
	// GalleryImageStatus defines the observed state of galleryimages
	Status GalleryImageStatusResponse `pulumi:"status"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Specifies information about the gallery image version that you want to create or update.
	Version *GalleryImageVersionResponse `pulumi:"version"`
}

func LookupGalleryimageRetrieveOutput(ctx *pulumi.Context, args LookupGalleryimageRetrieveOutputArgs, opts ...pulumi.InvokeOption) LookupGalleryimageRetrieveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGalleryimageRetrieveResult, error) {
			args := v.(LookupGalleryimageRetrieveArgs)
			r, err := LookupGalleryimageRetrieve(ctx, &args, opts...)
			var s LookupGalleryimageRetrieveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGalleryimageRetrieveResultOutput)
}

type LookupGalleryimageRetrieveOutputArgs struct {
	// Name of the gallery image
	GalleryimagesName pulumi.StringInput `pulumi:"galleryimagesName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGalleryimageRetrieveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryimageRetrieveArgs)(nil)).Elem()
}

// The gallery image resource definition.
type LookupGalleryimageRetrieveResultOutput struct{ *pulumi.OutputState }

func (LookupGalleryimageRetrieveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryimageRetrieveResult)(nil)).Elem()
}

func (o LookupGalleryimageRetrieveResultOutput) ToLookupGalleryimageRetrieveResultOutput() LookupGalleryimageRetrieveResultOutput {
	return o
}

func (o LookupGalleryimageRetrieveResultOutput) ToLookupGalleryimageRetrieveResultOutputWithContext(ctx context.Context) LookupGalleryimageRetrieveResultOutput {
	return o
}

// Datasource for the gallery image when provisioning with cloud-init (Azure or NoCloud)
func (o LookupGalleryimageRetrieveResultOutput) CloudInitDataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) *string { return v.CloudInitDataSource }).(pulumi.StringPtrOutput)
}

// Container Name for storage container
func (o LookupGalleryimageRetrieveResultOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

// The extendedLocation of the resource.
func (o LookupGalleryimageRetrieveResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The hypervisor generation of the Virtual Machine [V1, V2]
func (o LookupGalleryimageRetrieveResultOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) *string { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupGalleryimageRetrieveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) string { return v.Id }).(pulumi.StringOutput)
}

// This is the gallery image definition identifier.
func (o LookupGalleryimageRetrieveResultOutput) Identifier() GalleryImageIdentifierResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) *GalleryImageIdentifierResponse { return v.Identifier }).(GalleryImageIdentifierResponsePtrOutput)
}

// location of the image the gallery image should be created from
func (o LookupGalleryimageRetrieveResultOutput) ImagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) *string { return v.ImagePath }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupGalleryimageRetrieveResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupGalleryimageRetrieveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) string { return v.Name }).(pulumi.StringOutput)
}

// operating system type that the gallery image uses. Expected to be linux or windows
func (o LookupGalleryimageRetrieveResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

// Provisioning state of the gallery image.
func (o LookupGalleryimageRetrieveResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// name of the object to be used in moc
func (o LookupGalleryimageRetrieveResultOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

// GalleryImageStatus defines the observed state of galleryimages
func (o LookupGalleryimageRetrieveResultOutput) Status() GalleryImageStatusResponseOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) GalleryImageStatusResponse { return v.Status }).(GalleryImageStatusResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupGalleryimageRetrieveResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupGalleryimageRetrieveResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupGalleryimageRetrieveResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) string { return v.Type }).(pulumi.StringOutput)
}

// Specifies information about the gallery image version that you want to create or update.
func (o LookupGalleryimageRetrieveResultOutput) Version() GalleryImageVersionResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryimageRetrieveResult) *GalleryImageVersionResponse { return v.Version }).(GalleryImageVersionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGalleryimageRetrieveResultOutput{})
}