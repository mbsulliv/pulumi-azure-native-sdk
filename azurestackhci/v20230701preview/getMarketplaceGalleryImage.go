// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a marketplace gallery image
func LookupMarketplaceGalleryImage(ctx *pulumi.Context, args *LookupMarketplaceGalleryImageArgs, opts ...pulumi.InvokeOption) (*LookupMarketplaceGalleryImageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMarketplaceGalleryImageResult
	err := ctx.Invoke("azure-native:azurestackhci/v20230701preview:getMarketplaceGalleryImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMarketplaceGalleryImageArgs struct {
	// Name of the marketplace gallery image
	MarketplaceGalleryImageName string `pulumi:"marketplaceGalleryImageName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The marketplace gallery image resource definition.
type LookupMarketplaceGalleryImageResult struct {
	// Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
	CloudInitDataSource *string `pulumi:"cloudInitDataSource"`
	// Storage ContainerID of the storage container to be used for marketplace gallery image
	ContainerId *string `pulumi:"containerId"`
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine [V1, V2]
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// This is the gallery image definition identifier.
	Identifier *GalleryImageIdentifierResponse `pulumi:"identifier"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Operating system type that the gallery image uses [Windows, Linux]
	OsType string `pulumi:"osType"`
	// Provisioning state of the marketplace gallery image.
	ProvisioningState string `pulumi:"provisioningState"`
	// The observed state of marketplace gallery images
	Status MarketplaceGalleryImageStatusResponse `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Specifies information about the gallery image version that you want to create or update.
	Version *GalleryImageVersionResponse `pulumi:"version"`
}

func LookupMarketplaceGalleryImageOutput(ctx *pulumi.Context, args LookupMarketplaceGalleryImageOutputArgs, opts ...pulumi.InvokeOption) LookupMarketplaceGalleryImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMarketplaceGalleryImageResult, error) {
			args := v.(LookupMarketplaceGalleryImageArgs)
			r, err := LookupMarketplaceGalleryImage(ctx, &args, opts...)
			var s LookupMarketplaceGalleryImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMarketplaceGalleryImageResultOutput)
}

type LookupMarketplaceGalleryImageOutputArgs struct {
	// Name of the marketplace gallery image
	MarketplaceGalleryImageName pulumi.StringInput `pulumi:"marketplaceGalleryImageName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMarketplaceGalleryImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMarketplaceGalleryImageArgs)(nil)).Elem()
}

// The marketplace gallery image resource definition.
type LookupMarketplaceGalleryImageResultOutput struct{ *pulumi.OutputState }

func (LookupMarketplaceGalleryImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMarketplaceGalleryImageResult)(nil)).Elem()
}

func (o LookupMarketplaceGalleryImageResultOutput) ToLookupMarketplaceGalleryImageResultOutput() LookupMarketplaceGalleryImageResultOutput {
	return o
}

func (o LookupMarketplaceGalleryImageResultOutput) ToLookupMarketplaceGalleryImageResultOutputWithContext(ctx context.Context) LookupMarketplaceGalleryImageResultOutput {
	return o
}

func (o LookupMarketplaceGalleryImageResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMarketplaceGalleryImageResult] {
	return pulumix.Output[LookupMarketplaceGalleryImageResult]{
		OutputState: o.OutputState,
	}
}

// Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
func (o LookupMarketplaceGalleryImageResultOutput) CloudInitDataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) *string { return v.CloudInitDataSource }).(pulumi.StringPtrOutput)
}

// Storage ContainerID of the storage container to be used for marketplace gallery image
func (o LookupMarketplaceGalleryImageResultOutput) ContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) *string { return v.ContainerId }).(pulumi.StringPtrOutput)
}

// The extendedLocation of the resource.
func (o LookupMarketplaceGalleryImageResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The hypervisor generation of the Virtual Machine [V1, V2]
func (o LookupMarketplaceGalleryImageResultOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) *string { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupMarketplaceGalleryImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) string { return v.Id }).(pulumi.StringOutput)
}

// This is the gallery image definition identifier.
func (o LookupMarketplaceGalleryImageResultOutput) Identifier() GalleryImageIdentifierResponsePtrOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) *GalleryImageIdentifierResponse { return v.Identifier }).(GalleryImageIdentifierResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupMarketplaceGalleryImageResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMarketplaceGalleryImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) string { return v.Name }).(pulumi.StringOutput)
}

// Operating system type that the gallery image uses [Windows, Linux]
func (o LookupMarketplaceGalleryImageResultOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) string { return v.OsType }).(pulumi.StringOutput)
}

// Provisioning state of the marketplace gallery image.
func (o LookupMarketplaceGalleryImageResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The observed state of marketplace gallery images
func (o LookupMarketplaceGalleryImageResultOutput) Status() MarketplaceGalleryImageStatusResponseOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) MarketplaceGalleryImageStatusResponse { return v.Status }).(MarketplaceGalleryImageStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMarketplaceGalleryImageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupMarketplaceGalleryImageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMarketplaceGalleryImageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) string { return v.Type }).(pulumi.StringOutput)
}

// Specifies information about the gallery image version that you want to create or update.
func (o LookupMarketplaceGalleryImageResultOutput) Version() GalleryImageVersionResponsePtrOutput {
	return o.ApplyT(func(v LookupMarketplaceGalleryImageResult) *GalleryImageVersionResponse { return v.Version }).(GalleryImageVersionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMarketplaceGalleryImageResultOutput{})
}
