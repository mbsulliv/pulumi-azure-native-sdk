// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A custom image.
// Azure REST API version: 2018-09-15. Prior API version in Azure Native 1.x: 2018-09-15
type CustomImage struct {
	pulumi.CustomResourceState

	// The author of the custom image.
	Author pulumi.StringPtrOutput `pulumi:"author"`
	// The creation date of the custom image.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Storage information about the plan related to this custom image
	CustomImagePlan CustomImagePropertiesFromPlanResponsePtrOutput `pulumi:"customImagePlan"`
	// Storage information about the data disks present in the custom image
	DataDiskStorageInfo DataDiskStorageTypeInfoResponseArrayOutput `pulumi:"dataDiskStorageInfo"`
	// The description of the custom image.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether or not the custom images underlying offer/plan has been enabled for programmatic deployment
	IsPlanAuthorized pulumi.BoolPtrOutput `pulumi:"isPlanAuthorized"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The Managed Image Id backing the custom image.
	ManagedImageId pulumi.StringPtrOutput `pulumi:"managedImageId"`
	// The Managed Snapshot Id backing the custom image.
	ManagedSnapshotId pulumi.StringPtrOutput `pulumi:"managedSnapshotId"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringOutput `pulumi:"uniqueIdentifier"`
	// The VHD from which the image is to be created.
	Vhd CustomImagePropertiesCustomResponsePtrOutput `pulumi:"vhd"`
	// The virtual machine from which the image is to be created.
	Vm CustomImagePropertiesFromVmResponsePtrOutput `pulumi:"vm"`
}

// NewCustomImage registers a new resource with the given unique name, arguments, and options.
func NewCustomImage(ctx *pulumi.Context,
	name string, args *CustomImageArgs, opts ...pulumi.ResourceOption) (*CustomImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:CustomImage"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:CustomImage"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:CustomImage"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CustomImage
	err := ctx.RegisterResource("azure-native:devtestlab:CustomImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomImage gets an existing CustomImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomImageState, opts ...pulumi.ResourceOption) (*CustomImage, error) {
	var resource CustomImage
	err := ctx.ReadResource("azure-native:devtestlab:CustomImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomImage resources.
type customImageState struct {
}

type CustomImageState struct {
}

func (CustomImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*customImageState)(nil)).Elem()
}

type customImageArgs struct {
	// The author of the custom image.
	Author *string `pulumi:"author"`
	// Storage information about the plan related to this custom image
	CustomImagePlan *CustomImagePropertiesFromPlan `pulumi:"customImagePlan"`
	// Storage information about the data disks present in the custom image
	DataDiskStorageInfo []DataDiskStorageTypeInfo `pulumi:"dataDiskStorageInfo"`
	// The description of the custom image.
	Description *string `pulumi:"description"`
	// Whether or not the custom images underlying offer/plan has been enabled for programmatic deployment
	IsPlanAuthorized *bool `pulumi:"isPlanAuthorized"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The Managed Image Id backing the custom image.
	ManagedImageId *string `pulumi:"managedImageId"`
	// The Managed Snapshot Id backing the custom image.
	ManagedSnapshotId *string `pulumi:"managedSnapshotId"`
	// The name of the custom image.
	Name *string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VHD from which the image is to be created.
	Vhd *CustomImagePropertiesCustom `pulumi:"vhd"`
	// The virtual machine from which the image is to be created.
	Vm *CustomImagePropertiesFromVm `pulumi:"vm"`
}

// The set of arguments for constructing a CustomImage resource.
type CustomImageArgs struct {
	// The author of the custom image.
	Author pulumi.StringPtrInput
	// Storage information about the plan related to this custom image
	CustomImagePlan CustomImagePropertiesFromPlanPtrInput
	// Storage information about the data disks present in the custom image
	DataDiskStorageInfo DataDiskStorageTypeInfoArrayInput
	// The description of the custom image.
	Description pulumi.StringPtrInput
	// Whether or not the custom images underlying offer/plan has been enabled for programmatic deployment
	IsPlanAuthorized pulumi.BoolPtrInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The Managed Image Id backing the custom image.
	ManagedImageId pulumi.StringPtrInput
	// The Managed Snapshot Id backing the custom image.
	ManagedSnapshotId pulumi.StringPtrInput
	// The name of the custom image.
	Name pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The VHD from which the image is to be created.
	Vhd CustomImagePropertiesCustomPtrInput
	// The virtual machine from which the image is to be created.
	Vm CustomImagePropertiesFromVmPtrInput
}

func (CustomImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customImageArgs)(nil)).Elem()
}

type CustomImageInput interface {
	pulumi.Input

	ToCustomImageOutput() CustomImageOutput
	ToCustomImageOutputWithContext(ctx context.Context) CustomImageOutput
}

func (*CustomImage) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImage)(nil)).Elem()
}

func (i *CustomImage) ToCustomImageOutput() CustomImageOutput {
	return i.ToCustomImageOutputWithContext(context.Background())
}

func (i *CustomImage) ToCustomImageOutputWithContext(ctx context.Context) CustomImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImageOutput)
}

func (i *CustomImage) ToOutput(ctx context.Context) pulumix.Output[*CustomImage] {
	return pulumix.Output[*CustomImage]{
		OutputState: i.ToCustomImageOutputWithContext(ctx).OutputState,
	}
}

type CustomImageOutput struct{ *pulumi.OutputState }

func (CustomImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImage)(nil)).Elem()
}

func (o CustomImageOutput) ToCustomImageOutput() CustomImageOutput {
	return o
}

func (o CustomImageOutput) ToCustomImageOutputWithContext(ctx context.Context) CustomImageOutput {
	return o
}

func (o CustomImageOutput) ToOutput(ctx context.Context) pulumix.Output[*CustomImage] {
	return pulumix.Output[*CustomImage]{
		OutputState: o.OutputState,
	}
}

// The author of the custom image.
func (o CustomImageOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringPtrOutput { return v.Author }).(pulumi.StringPtrOutput)
}

// The creation date of the custom image.
func (o CustomImageOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// Storage information about the plan related to this custom image
func (o CustomImageOutput) CustomImagePlan() CustomImagePropertiesFromPlanResponsePtrOutput {
	return o.ApplyT(func(v *CustomImage) CustomImagePropertiesFromPlanResponsePtrOutput { return v.CustomImagePlan }).(CustomImagePropertiesFromPlanResponsePtrOutput)
}

// Storage information about the data disks present in the custom image
func (o CustomImageOutput) DataDiskStorageInfo() DataDiskStorageTypeInfoResponseArrayOutput {
	return o.ApplyT(func(v *CustomImage) DataDiskStorageTypeInfoResponseArrayOutput { return v.DataDiskStorageInfo }).(DataDiskStorageTypeInfoResponseArrayOutput)
}

// The description of the custom image.
func (o CustomImageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether or not the custom images underlying offer/plan has been enabled for programmatic deployment
func (o CustomImageOutput) IsPlanAuthorized() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.BoolPtrOutput { return v.IsPlanAuthorized }).(pulumi.BoolPtrOutput)
}

// The location of the resource.
func (o CustomImageOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The Managed Image Id backing the custom image.
func (o CustomImageOutput) ManagedImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringPtrOutput { return v.ManagedImageId }).(pulumi.StringPtrOutput)
}

// The Managed Snapshot Id backing the custom image.
func (o CustomImageOutput) ManagedSnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringPtrOutput { return v.ManagedSnapshotId }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o CustomImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning status of the resource.
func (o CustomImageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The tags of the resource.
func (o CustomImageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o CustomImageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o CustomImageOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

// The VHD from which the image is to be created.
func (o CustomImageOutput) Vhd() CustomImagePropertiesCustomResponsePtrOutput {
	return o.ApplyT(func(v *CustomImage) CustomImagePropertiesCustomResponsePtrOutput { return v.Vhd }).(CustomImagePropertiesCustomResponsePtrOutput)
}

// The virtual machine from which the image is to be created.
func (o CustomImageOutput) Vm() CustomImagePropertiesFromVmResponsePtrOutput {
	return o.ApplyT(func(v *CustomImage) CustomImagePropertiesFromVmResponsePtrOutput { return v.Vm }).(CustomImagePropertiesFromVmResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomImageOutput{})
}
