// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The gallery images resource definition.
type GalleryImage struct {
	pulumi.CustomResourceState

	// Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
	CloudInitDataSource pulumi.StringPtrOutput `pulumi:"cloudInitDataSource"`
	// Storage ContainerID of the storage container to be used for gallery image
	ContainerId pulumi.StringPtrOutput `pulumi:"containerId"`
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine [V1, V2]
	HyperVGeneration pulumi.StringPtrOutput `pulumi:"hyperVGeneration"`
	// This is the gallery image definition identifier.
	Identifier GalleryImageIdentifierResponsePtrOutput `pulumi:"identifier"`
	// location of the image the gallery image should be created from
	ImagePath pulumi.StringPtrOutput `pulumi:"imagePath"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Operating system type that the gallery image uses [Windows, Linux]
	OsType pulumi.StringOutput `pulumi:"osType"`
	// Provisioning state of the gallery image.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The observed state of gallery images
	Status GalleryImageStatusResponseOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies information about the gallery image version that you want to create or update.
	Version GalleryImageVersionResponsePtrOutput `pulumi:"version"`
}

// NewGalleryImage registers a new resource with the given unique name, arguments, and options.
func NewGalleryImage(ctx *pulumi.Context,
	name string, args *GalleryImageArgs, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OsType == nil {
		return nil, errors.New("invalid value for required argument 'OsType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210701preview:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221215preview:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230901preview:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240101:GalleryImage"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GalleryImage
	err := ctx.RegisterResource("azure-native:azurestackhci/v20230701preview:GalleryImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGalleryImage gets an existing GalleryImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGalleryImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryImageState, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	var resource GalleryImage
	err := ctx.ReadResource("azure-native:azurestackhci/v20230701preview:GalleryImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GalleryImage resources.
type galleryImageState struct {
}

type GalleryImageState struct {
}

func (GalleryImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageState)(nil)).Elem()
}

type galleryImageArgs struct {
	// Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
	CloudInitDataSource *string `pulumi:"cloudInitDataSource"`
	// Storage ContainerID of the storage container to be used for gallery image
	ContainerId *string `pulumi:"containerId"`
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Name of the gallery image
	GalleryImageName *string `pulumi:"galleryImageName"`
	// The hypervisor generation of the Virtual Machine [V1, V2]
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// This is the gallery image definition identifier.
	Identifier *GalleryImageIdentifier `pulumi:"identifier"`
	// location of the image the gallery image should be created from
	ImagePath *string `pulumi:"imagePath"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Operating system type that the gallery image uses [Windows, Linux]
	OsType OperatingSystemTypes `pulumi:"osType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Specifies information about the gallery image version that you want to create or update.
	Version *GalleryImageVersion `pulumi:"version"`
}

// The set of arguments for constructing a GalleryImage resource.
type GalleryImageArgs struct {
	// Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
	CloudInitDataSource pulumi.StringPtrInput
	// Storage ContainerID of the storage container to be used for gallery image
	ContainerId pulumi.StringPtrInput
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationPtrInput
	// Name of the gallery image
	GalleryImageName pulumi.StringPtrInput
	// The hypervisor generation of the Virtual Machine [V1, V2]
	HyperVGeneration pulumi.StringPtrInput
	// This is the gallery image definition identifier.
	Identifier GalleryImageIdentifierPtrInput
	// location of the image the gallery image should be created from
	ImagePath pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Operating system type that the gallery image uses [Windows, Linux]
	OsType OperatingSystemTypesInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Specifies information about the gallery image version that you want to create or update.
	Version GalleryImageVersionPtrInput
}

func (GalleryImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageArgs)(nil)).Elem()
}

type GalleryImageInput interface {
	pulumi.Input

	ToGalleryImageOutput() GalleryImageOutput
	ToGalleryImageOutputWithContext(ctx context.Context) GalleryImageOutput
}

func (*GalleryImage) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImage)(nil)).Elem()
}

func (i *GalleryImage) ToGalleryImageOutput() GalleryImageOutput {
	return i.ToGalleryImageOutputWithContext(context.Background())
}

func (i *GalleryImage) ToGalleryImageOutputWithContext(ctx context.Context) GalleryImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageOutput)
}

type GalleryImageOutput struct{ *pulumi.OutputState }

func (GalleryImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImage)(nil)).Elem()
}

func (o GalleryImageOutput) ToGalleryImageOutput() GalleryImageOutput {
	return o
}

func (o GalleryImageOutput) ToGalleryImageOutputWithContext(ctx context.Context) GalleryImageOutput {
	return o
}

// Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
func (o GalleryImageOutput) CloudInitDataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.CloudInitDataSource }).(pulumi.StringPtrOutput)
}

// Storage ContainerID of the storage container to be used for gallery image
func (o GalleryImageOutput) ContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.ContainerId }).(pulumi.StringPtrOutput)
}

// The extendedLocation of the resource.
func (o GalleryImageOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *GalleryImage) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The hypervisor generation of the Virtual Machine [V1, V2]
func (o GalleryImageOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// This is the gallery image definition identifier.
func (o GalleryImageOutput) Identifier() GalleryImageIdentifierResponsePtrOutput {
	return o.ApplyT(func(v *GalleryImage) GalleryImageIdentifierResponsePtrOutput { return v.Identifier }).(GalleryImageIdentifierResponsePtrOutput)
}

// location of the image the gallery image should be created from
func (o GalleryImageOutput) ImagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.ImagePath }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o GalleryImageOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o GalleryImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Operating system type that the gallery image uses [Windows, Linux]
func (o GalleryImageOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.OsType }).(pulumi.StringOutput)
}

// Provisioning state of the gallery image.
func (o GalleryImageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The observed state of gallery images
func (o GalleryImageOutput) Status() GalleryImageStatusResponseOutput {
	return o.ApplyT(func(v *GalleryImage) GalleryImageStatusResponseOutput { return v.Status }).(GalleryImageStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o GalleryImageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GalleryImage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o GalleryImageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GalleryImageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies information about the gallery image version that you want to create or update.
func (o GalleryImageOutput) Version() GalleryImageVersionResponsePtrOutput {
	return o.ApplyT(func(v *GalleryImage) GalleryImageVersionResponsePtrOutput { return v.Version }).(GalleryImageVersionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GalleryImageOutput{})
}
