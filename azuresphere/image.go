// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuresphere

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An image resource belonging to a catalog resource.
// Azure REST API version: 2022-09-01-preview. Prior API version in Azure Native 1.x: 2022-09-01-preview.
type Image struct {
	pulumi.CustomResourceState

	// The image component id.
	ComponentId pulumi.StringOutput `pulumi:"componentId"`
	// The image description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Image as a UTF-8 encoded base 64 string on image create. This field contains the image URI on image reads.
	Image pulumi.StringPtrOutput `pulumi:"image"`
	// Image ID
	ImageId pulumi.StringPtrOutput `pulumi:"imageId"`
	// Image name
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// The image type.
	ImageType pulumi.StringOutput `pulumi:"imageType"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Regional data boundary for an image
	RegionalDataBoundary pulumi.StringPtrOutput `pulumi:"regionalDataBoundary"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Location the image
	Uri pulumi.StringOutput `pulumi:"uri"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogName == nil {
		return nil, errors.New("invalid value for required argument 'CatalogName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azuresphere/v20220901preview:Image"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Image
	err := ctx.RegisterResource("azure-native:azuresphere:Image", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:azuresphere:Image", name, id, state, &resource, opts...)
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
	// Name of catalog
	CatalogName string `pulumi:"catalogName"`
	// Image as a UTF-8 encoded base 64 string on image create. This field contains the image URI on image reads.
	Image *string `pulumi:"image"`
	// Image ID
	ImageId *string `pulumi:"imageId"`
	// Image name. Use .default for image creation.
	ImageName *string `pulumi:"imageName"`
	// Regional data boundary for an image
	RegionalDataBoundary *string `pulumi:"regionalDataBoundary"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// Name of catalog
	CatalogName pulumi.StringInput
	// Image as a UTF-8 encoded base 64 string on image create. This field contains the image URI on image reads.
	Image pulumi.StringPtrInput
	// Image ID
	ImageId pulumi.StringPtrInput
	// Image name. Use .default for image creation.
	ImageName pulumi.StringPtrInput
	// Regional data boundary for an image
	RegionalDataBoundary pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
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

// The image component id.
func (o ImageOutput) ComponentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ComponentId }).(pulumi.StringOutput)
}

// The image description.
func (o ImageOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Image as a UTF-8 encoded base 64 string on image create. This field contains the image URI on image reads.
func (o ImageOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.Image }).(pulumi.StringPtrOutput)
}

// Image ID
func (o ImageOutput) ImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.ImageId }).(pulumi.StringPtrOutput)
}

// Image name
func (o ImageOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ImageName }).(pulumi.StringOutput)
}

// The image type.
func (o ImageOutput) ImageType() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ImageType }).(pulumi.StringOutput)
}

// The name of the resource
func (o ImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o ImageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Regional data boundary for an image
func (o ImageOutput) RegionalDataBoundary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.RegionalDataBoundary }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ImageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Image) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ImageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Location the image
func (o ImageOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ImageOutput{})
}
