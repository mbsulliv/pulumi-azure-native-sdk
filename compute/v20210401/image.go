// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.
type Image struct {
	pulumi.CustomResourceState

	// The extended location of the Image.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// Specifies the HyperVGenerationType of the VirtualMachine created from the image. From API Version 2019-03-01 if the image source is a blob, then we need the user to specify the value, if the source is managed resource like disk or snapshot, we may require the user to specify the property if we cannot deduce it from the source managed resource.
	HyperVGeneration pulumi.StringPtrOutput `pulumi:"hyperVGeneration"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The source virtual machine from which Image is created.
	SourceVirtualMachine SubResourceResponsePtrOutput `pulumi:"sourceVirtualMachine"`
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile ImageStorageProfileResponsePtrOutput `pulumi:"storageProfile"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230301:Image"),
		},
	})
	opts = append(opts, aliases)
	var resource Image
	err := ctx.RegisterResource("azure-native:compute/v20210401:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("azure-native:compute/v20210401:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
}

type ImageState struct {
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// The extended location of the Image.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Specifies the HyperVGenerationType of the VirtualMachine created from the image. From API Version 2019-03-01 if the image source is a blob, then we need the user to specify the value, if the source is managed resource like disk or snapshot, we may require the user to specify the property if we cannot deduce it from the source managed resource.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// The name of the image.
	ImageName *string `pulumi:"imageName"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The source virtual machine from which Image is created.
	SourceVirtualMachine *SubResource `pulumi:"sourceVirtualMachine"`
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile *ImageStorageProfile `pulumi:"storageProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// The extended location of the Image.
	ExtendedLocation ExtendedLocationPtrInput
	// Specifies the HyperVGenerationType of the VirtualMachine created from the image. From API Version 2019-03-01 if the image source is a blob, then we need the user to specify the value, if the source is managed resource like disk or snapshot, we may require the user to specify the property if we cannot deduce it from the source managed resource.
	HyperVGeneration pulumi.StringPtrInput
	// The name of the image.
	ImageName pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The source virtual machine from which Image is created.
	SourceVirtualMachine SubResourcePtrInput
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile ImageStorageProfilePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

// The extended location of the Image.
func (o ImageOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Image) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Specifies the HyperVGenerationType of the VirtualMachine created from the image. From API Version 2019-03-01 if the image source is a blob, then we need the user to specify the value, if the source is managed resource like disk or snapshot, we may require the user to specify the property if we cannot deduce it from the source managed resource.
func (o ImageOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// Resource location
func (o ImageOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o ImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state.
func (o ImageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The source virtual machine from which Image is created.
func (o ImageOutput) SourceVirtualMachine() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *Image) SubResourceResponsePtrOutput { return v.SourceVirtualMachine }).(SubResourceResponsePtrOutput)
}

// Specifies the storage settings for the virtual machine disks.
func (o ImageOutput) StorageProfile() ImageStorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *Image) ImageStorageProfileResponsePtrOutput { return v.StorageProfile }).(ImageStorageProfileResponsePtrOutput)
}

// Resource tags
func (o ImageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Image) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o ImageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ImageOutput{})
}
