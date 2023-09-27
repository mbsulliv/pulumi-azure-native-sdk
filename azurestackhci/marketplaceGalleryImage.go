// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurestackhci

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The marketplace gallery image resource definition.
// Azure REST API version: 2022-12-15-preview.
type MarketplaceGalleryImage struct {
	pulumi.CustomResourceState

	// Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
	CloudInitDataSource pulumi.StringPtrOutput `pulumi:"cloudInitDataSource"`
	// Container Name for storage container
	ContainerName pulumi.StringPtrOutput `pulumi:"containerName"`
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine [V1, V2]
	HyperVGeneration pulumi.StringPtrOutput `pulumi:"hyperVGeneration"`
	// This is the gallery image definition identifier.
	Identifier GalleryImageIdentifierResponsePtrOutput `pulumi:"identifier"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Operating system type that the gallery image uses [Windows, Linux]
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// Provisioning state of the marketplace gallery image.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The observed state of marketplace gallery images
	Status MarketplaceGalleryImageStatusResponseOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies information about the gallery image version that you want to create or update.
	Version GalleryImageVersionResponsePtrOutput `pulumi:"version"`
}

// NewMarketplaceGalleryImage registers a new resource with the given unique name, arguments, and options.
func NewMarketplaceGalleryImage(ctx *pulumi.Context,
	name string, args *MarketplaceGalleryImageArgs, opts ...pulumi.ResourceOption) (*MarketplaceGalleryImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:MarketplaceGalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221215preview:MarketplaceGalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230701preview:MarketplaceGalleryImage"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MarketplaceGalleryImage
	err := ctx.RegisterResource("azure-native:azurestackhci:MarketplaceGalleryImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMarketplaceGalleryImage gets an existing MarketplaceGalleryImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMarketplaceGalleryImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MarketplaceGalleryImageState, opts ...pulumi.ResourceOption) (*MarketplaceGalleryImage, error) {
	var resource MarketplaceGalleryImage
	err := ctx.ReadResource("azure-native:azurestackhci:MarketplaceGalleryImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MarketplaceGalleryImage resources.
type marketplaceGalleryImageState struct {
}

type MarketplaceGalleryImageState struct {
}

func (MarketplaceGalleryImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*marketplaceGalleryImageState)(nil)).Elem()
}

type marketplaceGalleryImageArgs struct {
	// Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
	CloudInitDataSource *string `pulumi:"cloudInitDataSource"`
	// Container Name for storage container
	ContainerName *string `pulumi:"containerName"`
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine [V1, V2]
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// This is the gallery image definition identifier.
	Identifier *GalleryImageIdentifier `pulumi:"identifier"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of the marketplace gallery image
	MarketplaceGalleryImageName *string `pulumi:"marketplaceGalleryImageName"`
	// Operating system type that the gallery image uses [Windows, Linux]
	OsType *OperatingSystemTypes `pulumi:"osType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Specifies information about the gallery image version that you want to create or update.
	Version *GalleryImageVersion `pulumi:"version"`
}

// The set of arguments for constructing a MarketplaceGalleryImage resource.
type MarketplaceGalleryImageArgs struct {
	// Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
	CloudInitDataSource pulumi.StringPtrInput
	// Container Name for storage container
	ContainerName pulumi.StringPtrInput
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationPtrInput
	// The hypervisor generation of the Virtual Machine [V1, V2]
	HyperVGeneration pulumi.StringPtrInput
	// This is the gallery image definition identifier.
	Identifier GalleryImageIdentifierPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of the marketplace gallery image
	MarketplaceGalleryImageName pulumi.StringPtrInput
	// Operating system type that the gallery image uses [Windows, Linux]
	OsType OperatingSystemTypesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Specifies information about the gallery image version that you want to create or update.
	Version GalleryImageVersionPtrInput
}

func (MarketplaceGalleryImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*marketplaceGalleryImageArgs)(nil)).Elem()
}

type MarketplaceGalleryImageInput interface {
	pulumi.Input

	ToMarketplaceGalleryImageOutput() MarketplaceGalleryImageOutput
	ToMarketplaceGalleryImageOutputWithContext(ctx context.Context) MarketplaceGalleryImageOutput
}

func (*MarketplaceGalleryImage) ElementType() reflect.Type {
	return reflect.TypeOf((**MarketplaceGalleryImage)(nil)).Elem()
}

func (i *MarketplaceGalleryImage) ToMarketplaceGalleryImageOutput() MarketplaceGalleryImageOutput {
	return i.ToMarketplaceGalleryImageOutputWithContext(context.Background())
}

func (i *MarketplaceGalleryImage) ToMarketplaceGalleryImageOutputWithContext(ctx context.Context) MarketplaceGalleryImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarketplaceGalleryImageOutput)
}

func (i *MarketplaceGalleryImage) ToOutput(ctx context.Context) pulumix.Output[*MarketplaceGalleryImage] {
	return pulumix.Output[*MarketplaceGalleryImage]{
		OutputState: i.ToMarketplaceGalleryImageOutputWithContext(ctx).OutputState,
	}
}

type MarketplaceGalleryImageOutput struct{ *pulumi.OutputState }

func (MarketplaceGalleryImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MarketplaceGalleryImage)(nil)).Elem()
}

func (o MarketplaceGalleryImageOutput) ToMarketplaceGalleryImageOutput() MarketplaceGalleryImageOutput {
	return o
}

func (o MarketplaceGalleryImageOutput) ToMarketplaceGalleryImageOutputWithContext(ctx context.Context) MarketplaceGalleryImageOutput {
	return o
}

func (o MarketplaceGalleryImageOutput) ToOutput(ctx context.Context) pulumix.Output[*MarketplaceGalleryImage] {
	return pulumix.Output[*MarketplaceGalleryImage]{
		OutputState: o.OutputState,
	}
}

// Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
func (o MarketplaceGalleryImageOutput) CloudInitDataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) pulumi.StringPtrOutput { return v.CloudInitDataSource }).(pulumi.StringPtrOutput)
}

// Container Name for storage container
func (o MarketplaceGalleryImageOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) pulumi.StringPtrOutput { return v.ContainerName }).(pulumi.StringPtrOutput)
}

// The extendedLocation of the resource.
func (o MarketplaceGalleryImageOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The hypervisor generation of the Virtual Machine [V1, V2]
func (o MarketplaceGalleryImageOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) pulumi.StringPtrOutput { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// This is the gallery image definition identifier.
func (o MarketplaceGalleryImageOutput) Identifier() GalleryImageIdentifierResponsePtrOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) GalleryImageIdentifierResponsePtrOutput { return v.Identifier }).(GalleryImageIdentifierResponsePtrOutput)
}

// The geo-location where the resource lives
func (o MarketplaceGalleryImageOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o MarketplaceGalleryImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Operating system type that the gallery image uses [Windows, Linux]
func (o MarketplaceGalleryImageOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

// Provisioning state of the marketplace gallery image.
func (o MarketplaceGalleryImageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The observed state of marketplace gallery images
func (o MarketplaceGalleryImageOutput) Status() MarketplaceGalleryImageStatusResponseOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) MarketplaceGalleryImageStatusResponseOutput { return v.Status }).(MarketplaceGalleryImageStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MarketplaceGalleryImageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o MarketplaceGalleryImageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MarketplaceGalleryImageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies information about the gallery image version that you want to create or update.
func (o MarketplaceGalleryImageOutput) Version() GalleryImageVersionResponsePtrOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImage) GalleryImageVersionResponsePtrOutput { return v.Version }).(GalleryImageVersionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MarketplaceGalleryImageOutput{})
}
