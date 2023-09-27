// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The marketplace gallery image resource definition.
type Marketplacegalleryimage struct {
	pulumi.CustomResourceState

	// Datasource for the gallery image when provisioning with cloud-init [Azure, NoCloud]
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
	// operating system type that the gallery image uses. Expected to be linux or windows
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// Provisioning state of the gallery image.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// name of the object to be used in moc
	ResourceName pulumi.StringPtrOutput `pulumi:"resourceName"`
	// MarketplaceGalleryImageStatus defines the observed state of marketplacegalleryimages
	Status MarketplaceGalleryImageStatusResponseOutput `pulumi:"status"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies information about the gallery image version that you want to create or update.
	Version GalleryImageVersionResponsePtrOutput `pulumi:"version"`
}

// NewMarketplacegalleryimage registers a new resource with the given unique name, arguments, and options.
func NewMarketplacegalleryimage(ctx *pulumi.Context,
	name string, args *MarketplacegalleryimageArgs, opts ...pulumi.ResourceOption) (*Marketplacegalleryimage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:marketplacegalleryimage"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci:Marketplacegalleryimage"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci:marketplacegalleryimage"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221215preview:Marketplacegalleryimage"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221215preview:marketplacegalleryimage"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230701preview:Marketplacegalleryimage"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230701preview:marketplacegalleryimage"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Marketplacegalleryimage
	err := ctx.RegisterResource("azure-native:azurestackhci/v20210901preview:Marketplacegalleryimage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMarketplacegalleryimage gets an existing Marketplacegalleryimage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMarketplacegalleryimage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MarketplacegalleryimageState, opts ...pulumi.ResourceOption) (*Marketplacegalleryimage, error) {
	var resource Marketplacegalleryimage
	err := ctx.ReadResource("azure-native:azurestackhci/v20210901preview:Marketplacegalleryimage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Marketplacegalleryimage resources.
type marketplacegalleryimageState struct {
}

type MarketplacegalleryimageState struct {
}

func (MarketplacegalleryimageState) ElementType() reflect.Type {
	return reflect.TypeOf((*marketplacegalleryimageState)(nil)).Elem()
}

type marketplacegalleryimageArgs struct {
	// Datasource for the gallery image when provisioning with cloud-init [Azure, NoCloud]
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
	MarketplacegalleryimagesName *string `pulumi:"marketplacegalleryimagesName"`
	// operating system type that the gallery image uses. Expected to be linux or windows
	OsType *OperatingSystemTypes `pulumi:"osType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// name of the object to be used in moc
	ResourceName *string `pulumi:"resourceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Specifies information about the gallery image version that you want to create or update.
	Version *GalleryImageVersion `pulumi:"version"`
}

// The set of arguments for constructing a Marketplacegalleryimage resource.
type MarketplacegalleryimageArgs struct {
	// Datasource for the gallery image when provisioning with cloud-init [Azure, NoCloud]
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
	MarketplacegalleryimagesName pulumi.StringPtrInput
	// operating system type that the gallery image uses. Expected to be linux or windows
	OsType OperatingSystemTypesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// name of the object to be used in moc
	ResourceName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Specifies information about the gallery image version that you want to create or update.
	Version GalleryImageVersionPtrInput
}

func (MarketplacegalleryimageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*marketplacegalleryimageArgs)(nil)).Elem()
}

type MarketplacegalleryimageInput interface {
	pulumi.Input

	ToMarketplacegalleryimageOutput() MarketplacegalleryimageOutput
	ToMarketplacegalleryimageOutputWithContext(ctx context.Context) MarketplacegalleryimageOutput
}

func (*Marketplacegalleryimage) ElementType() reflect.Type {
	return reflect.TypeOf((**Marketplacegalleryimage)(nil)).Elem()
}

func (i *Marketplacegalleryimage) ToMarketplacegalleryimageOutput() MarketplacegalleryimageOutput {
	return i.ToMarketplacegalleryimageOutputWithContext(context.Background())
}

func (i *Marketplacegalleryimage) ToMarketplacegalleryimageOutputWithContext(ctx context.Context) MarketplacegalleryimageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarketplacegalleryimageOutput)
}

func (i *Marketplacegalleryimage) ToOutput(ctx context.Context) pulumix.Output[*Marketplacegalleryimage] {
	return pulumix.Output[*Marketplacegalleryimage]{
		OutputState: i.ToMarketplacegalleryimageOutputWithContext(ctx).OutputState,
	}
}

type MarketplacegalleryimageOutput struct{ *pulumi.OutputState }

func (MarketplacegalleryimageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Marketplacegalleryimage)(nil)).Elem()
}

func (o MarketplacegalleryimageOutput) ToMarketplacegalleryimageOutput() MarketplacegalleryimageOutput {
	return o
}

func (o MarketplacegalleryimageOutput) ToMarketplacegalleryimageOutputWithContext(ctx context.Context) MarketplacegalleryimageOutput {
	return o
}

func (o MarketplacegalleryimageOutput) ToOutput(ctx context.Context) pulumix.Output[*Marketplacegalleryimage] {
	return pulumix.Output[*Marketplacegalleryimage]{
		OutputState: o.OutputState,
	}
}

// Datasource for the gallery image when provisioning with cloud-init [Azure, NoCloud]
func (o MarketplacegalleryimageOutput) CloudInitDataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringPtrOutput { return v.CloudInitDataSource }).(pulumi.StringPtrOutput)
}

// Container Name for storage container
func (o MarketplacegalleryimageOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringPtrOutput { return v.ContainerName }).(pulumi.StringPtrOutput)
}

// The extendedLocation of the resource.
func (o MarketplacegalleryimageOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The hypervisor generation of the Virtual Machine [V1, V2]
func (o MarketplacegalleryimageOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringPtrOutput { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// This is the gallery image definition identifier.
func (o MarketplacegalleryimageOutput) Identifier() GalleryImageIdentifierResponsePtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) GalleryImageIdentifierResponsePtrOutput { return v.Identifier }).(GalleryImageIdentifierResponsePtrOutput)
}

// The geo-location where the resource lives
func (o MarketplacegalleryimageOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o MarketplacegalleryimageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// operating system type that the gallery image uses. Expected to be linux or windows
func (o MarketplacegalleryimageOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

// Provisioning state of the gallery image.
func (o MarketplacegalleryimageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// name of the object to be used in moc
func (o MarketplacegalleryimageOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringPtrOutput { return v.ResourceName }).(pulumi.StringPtrOutput)
}

// MarketplaceGalleryImageStatus defines the observed state of marketplacegalleryimages
func (o MarketplacegalleryimageOutput) Status() MarketplaceGalleryImageStatusResponseOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) MarketplaceGalleryImageStatusResponseOutput { return v.Status }).(MarketplaceGalleryImageStatusResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o MarketplacegalleryimageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o MarketplacegalleryimageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MarketplacegalleryimageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies information about the gallery image version that you want to create or update.
func (o MarketplacegalleryimageOutput) Version() GalleryImageVersionResponsePtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) GalleryImageVersionResponsePtrOutput { return v.Version }).(GalleryImageVersionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MarketplacegalleryimageOutput{})
}
