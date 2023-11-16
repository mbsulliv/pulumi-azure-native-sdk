// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211120preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Response for Volume request.
type Volume struct {
	pulumi.CustomResourceState

	// State of the operation on the resource.
	CreationData SourceCreationDataResponsePtrOutput `pulumi:"creationData"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Volume size.
	SizeGiB pulumi.Float64PtrOutput `pulumi:"sizeGiB"`
	// Storage target information
	StorageTarget IscsiTargetInfoResponseOutput `pulumi:"storageTarget"`
	// Resource metadata required by ARM RPC
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique Id of the volume in GUID format
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ElasticSanName == nil {
		return nil, errors.New("invalid value for required argument 'ElasticSanName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VolumeGroupName == nil {
		return nil, errors.New("invalid value for required argument 'VolumeGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:elasticsan:Volume"),
		},
		{
			Type: pulumi.String("azure-native:elasticsan/v20221201preview:Volume"),
		},
		{
			Type: pulumi.String("azure-native:elasticsan/v20230101:Volume"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Volume
	err := ctx.RegisterResource("azure-native:elasticsan/v20211120preview:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure-native:elasticsan/v20211120preview:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
}

type VolumeState struct {
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// State of the operation on the resource.
	CreationData *SourceCreationData `pulumi:"creationData"`
	// The name of the ElasticSan.
	ElasticSanName string `pulumi:"elasticSanName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Volume size.
	SizeGiB *float64 `pulumi:"sizeGiB"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the VolumeGroup.
	VolumeGroupName string `pulumi:"volumeGroupName"`
	// The name of the Volume.
	VolumeName *string `pulumi:"volumeName"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// State of the operation on the resource.
	CreationData SourceCreationDataPtrInput
	// The name of the ElasticSan.
	ElasticSanName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Volume size.
	SizeGiB pulumi.Float64PtrInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// The name of the VolumeGroup.
	VolumeGroupName pulumi.StringInput
	// The name of the Volume.
	VolumeName pulumi.StringPtrInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(ctx context.Context) VolumeOutput
}

func (*Volume) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (i *Volume) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i *Volume) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

// State of the operation on the resource.
func (o VolumeOutput) CreationData() SourceCreationDataResponsePtrOutput {
	return o.ApplyT(func(v *Volume) SourceCreationDataResponsePtrOutput { return v.CreationData }).(SourceCreationDataResponsePtrOutput)
}

// Azure resource name.
func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Volume size.
func (o VolumeOutput) SizeGiB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.Float64PtrOutput { return v.SizeGiB }).(pulumi.Float64PtrOutput)
}

// Storage target information
func (o VolumeOutput) StorageTarget() IscsiTargetInfoResponseOutput {
	return o.ApplyT(func(v *Volume) IscsiTargetInfoResponseOutput { return v.StorageTarget }).(IscsiTargetInfoResponseOutput)
}

// Resource metadata required by ARM RPC
func (o VolumeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Volume) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o VolumeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o VolumeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Unique Id of the volume in GUID format
func (o VolumeOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.VolumeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeOutput{})
}
