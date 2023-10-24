// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Specifies information about the Shared Image Gallery that you want to create or update.
// Azure REST API version: 2022-03-03. Prior API version in Azure Native 1.x: 2020-09-30.
type Gallery struct {
	pulumi.CustomResourceState

	// The description of this Shared Image Gallery resource. This property is updatable.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Describes the gallery unique name.
	Identifier GalleryIdentifierResponsePtrOutput `pulumi:"identifier"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Profile for gallery sharing to subscription or tenant
	SharingProfile SharingProfileResponsePtrOutput `pulumi:"sharingProfile"`
	// Sharing status of current gallery.
	SharingStatus SharingStatusResponseOutput `pulumi:"sharingStatus"`
	// Contains information about the soft deletion policy of the gallery.
	SoftDeletePolicy SoftDeletePolicyResponsePtrOutput `pulumi:"softDeletePolicy"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGallery registers a new resource with the given unique name, arguments, and options.
func NewGallery(ctx *pulumi.Context,
	name string, args *GalleryArgs, opts ...pulumi.ResourceOption) (*Gallery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20180601:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211001:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220103:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220303:Gallery"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Gallery
	err := ctx.RegisterResource("azure-native:compute:Gallery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGallery gets an existing Gallery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGallery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryState, opts ...pulumi.ResourceOption) (*Gallery, error) {
	var resource Gallery
	err := ctx.ReadResource("azure-native:compute:Gallery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gallery resources.
type galleryState struct {
}

type GalleryState struct {
}

func (GalleryState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryState)(nil)).Elem()
}

type galleryArgs struct {
	// The description of this Shared Image Gallery resource. This property is updatable.
	Description *string `pulumi:"description"`
	// The name of the Shared Image Gallery. The allowed characters are alphabets and numbers with dots and periods allowed in the middle. The maximum length is 80 characters.
	GalleryName *string `pulumi:"galleryName"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Profile for gallery sharing to subscription or tenant
	SharingProfile *SharingProfile `pulumi:"sharingProfile"`
	// Contains information about the soft deletion policy of the gallery.
	SoftDeletePolicy *SoftDeletePolicy `pulumi:"softDeletePolicy"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Gallery resource.
type GalleryArgs struct {
	// The description of this Shared Image Gallery resource. This property is updatable.
	Description pulumi.StringPtrInput
	// The name of the Shared Image Gallery. The allowed characters are alphabets and numbers with dots and periods allowed in the middle. The maximum length is 80 characters.
	GalleryName pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Profile for gallery sharing to subscription or tenant
	SharingProfile SharingProfilePtrInput
	// Contains information about the soft deletion policy of the gallery.
	SoftDeletePolicy SoftDeletePolicyPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (GalleryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryArgs)(nil)).Elem()
}

type GalleryInput interface {
	pulumi.Input

	ToGalleryOutput() GalleryOutput
	ToGalleryOutputWithContext(ctx context.Context) GalleryOutput
}

func (*Gallery) ElementType() reflect.Type {
	return reflect.TypeOf((**Gallery)(nil)).Elem()
}

func (i *Gallery) ToGalleryOutput() GalleryOutput {
	return i.ToGalleryOutputWithContext(context.Background())
}

func (i *Gallery) ToGalleryOutputWithContext(ctx context.Context) GalleryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryOutput)
}

func (i *Gallery) ToOutput(ctx context.Context) pulumix.Output[*Gallery] {
	return pulumix.Output[*Gallery]{
		OutputState: i.ToGalleryOutputWithContext(ctx).OutputState,
	}
}

type GalleryOutput struct{ *pulumi.OutputState }

func (GalleryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gallery)(nil)).Elem()
}

func (o GalleryOutput) ToGalleryOutput() GalleryOutput {
	return o
}

func (o GalleryOutput) ToGalleryOutputWithContext(ctx context.Context) GalleryOutput {
	return o
}

func (o GalleryOutput) ToOutput(ctx context.Context) pulumix.Output[*Gallery] {
	return pulumix.Output[*Gallery]{
		OutputState: o.OutputState,
	}
}

// The description of this Shared Image Gallery resource. This property is updatable.
func (o GalleryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Describes the gallery unique name.
func (o GalleryOutput) Identifier() GalleryIdentifierResponsePtrOutput {
	return o.ApplyT(func(v *Gallery) GalleryIdentifierResponsePtrOutput { return v.Identifier }).(GalleryIdentifierResponsePtrOutput)
}

// Resource location
func (o GalleryOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o GalleryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state, which only appears in the response.
func (o GalleryOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Profile for gallery sharing to subscription or tenant
func (o GalleryOutput) SharingProfile() SharingProfileResponsePtrOutput {
	return o.ApplyT(func(v *Gallery) SharingProfileResponsePtrOutput { return v.SharingProfile }).(SharingProfileResponsePtrOutput)
}

// Sharing status of current gallery.
func (o GalleryOutput) SharingStatus() SharingStatusResponseOutput {
	return o.ApplyT(func(v *Gallery) SharingStatusResponseOutput { return v.SharingStatus }).(SharingStatusResponseOutput)
}

// Contains information about the soft deletion policy of the gallery.
func (o GalleryOutput) SoftDeletePolicy() SoftDeletePolicyResponsePtrOutput {
	return o.ApplyT(func(v *Gallery) SoftDeletePolicyResponsePtrOutput { return v.SoftDeletePolicy }).(SoftDeletePolicyResponsePtrOutput)
}

// Resource tags
func (o GalleryOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o GalleryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GalleryOutput{})
}
