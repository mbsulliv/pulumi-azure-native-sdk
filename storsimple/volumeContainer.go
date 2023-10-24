// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The volume container.
// Azure REST API version: 2017-06-01. Prior API version in Azure Native 1.x: 2017-06-01.
type VolumeContainer struct {
	pulumi.CustomResourceState

	// The bandwidth-rate set on the volume container.
	BandWidthRateInMbps pulumi.IntPtrOutput `pulumi:"bandWidthRateInMbps"`
	// The ID of the bandwidth setting associated with the volume container.
	BandwidthSettingId pulumi.StringPtrOutput `pulumi:"bandwidthSettingId"`
	// The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
	EncryptionKey AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"encryptionKey"`
	// The flag to denote whether encryption is enabled or not.
	EncryptionStatus pulumi.StringOutput `pulumi:"encryptionStatus"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// The owner ship status of the volume container. Only when the status is "NotOwned", the delete operation on the volume container is permitted.
	OwnerShipStatus pulumi.StringOutput `pulumi:"ownerShipStatus"`
	// The path ID of storage account associated with the volume container.
	StorageAccountCredentialId pulumi.StringOutput `pulumi:"storageAccountCredentialId"`
	// The total cloud storage for the volume container.
	TotalCloudStorageUsageInBytes pulumi.Float64Output `pulumi:"totalCloudStorageUsageInBytes"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// The number of volumes in the volume Container.
	VolumeCount pulumi.IntOutput `pulumi:"volumeCount"`
}

// NewVolumeContainer registers a new resource with the given unique name, arguments, and options.
func NewVolumeContainer(ctx *pulumi.Context,
	name string, args *VolumeContainerArgs, opts ...pulumi.ResourceOption) (*VolumeContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountCredentialId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountCredentialId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storsimple/v20170601:VolumeContainer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VolumeContainer
	err := ctx.RegisterResource("azure-native:storsimple:VolumeContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeContainer gets an existing VolumeContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeContainerState, opts ...pulumi.ResourceOption) (*VolumeContainer, error) {
	var resource VolumeContainer
	err := ctx.ReadResource("azure-native:storsimple:VolumeContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeContainer resources.
type volumeContainerState struct {
}

type VolumeContainerState struct {
}

func (VolumeContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeContainerState)(nil)).Elem()
}

type volumeContainerArgs struct {
	// The bandwidth-rate set on the volume container.
	BandWidthRateInMbps *int `pulumi:"bandWidthRateInMbps"`
	// The ID of the bandwidth setting associated with the volume container.
	BandwidthSettingId *string `pulumi:"bandwidthSettingId"`
	// The device name
	DeviceName string `pulumi:"deviceName"`
	// The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
	EncryptionKey *AsymmetricEncryptedSecret `pulumi:"encryptionKey"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *Kind `pulumi:"kind"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The path ID of storage account associated with the volume container.
	StorageAccountCredentialId string `pulumi:"storageAccountCredentialId"`
	// The name of the volume container.
	VolumeContainerName *string `pulumi:"volumeContainerName"`
}

// The set of arguments for constructing a VolumeContainer resource.
type VolumeContainerArgs struct {
	// The bandwidth-rate set on the volume container.
	BandWidthRateInMbps pulumi.IntPtrInput
	// The ID of the bandwidth setting associated with the volume container.
	BandwidthSettingId pulumi.StringPtrInput
	// The device name
	DeviceName pulumi.StringInput
	// The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
	EncryptionKey AsymmetricEncryptedSecretPtrInput
	// The Kind of the object. Currently only Series8000 is supported
	Kind KindPtrInput
	// The manager name
	ManagerName pulumi.StringInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// The path ID of storage account associated with the volume container.
	StorageAccountCredentialId pulumi.StringInput
	// The name of the volume container.
	VolumeContainerName pulumi.StringPtrInput
}

func (VolumeContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeContainerArgs)(nil)).Elem()
}

type VolumeContainerInput interface {
	pulumi.Input

	ToVolumeContainerOutput() VolumeContainerOutput
	ToVolumeContainerOutputWithContext(ctx context.Context) VolumeContainerOutput
}

func (*VolumeContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeContainer)(nil)).Elem()
}

func (i *VolumeContainer) ToVolumeContainerOutput() VolumeContainerOutput {
	return i.ToVolumeContainerOutputWithContext(context.Background())
}

func (i *VolumeContainer) ToVolumeContainerOutputWithContext(ctx context.Context) VolumeContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeContainerOutput)
}

func (i *VolumeContainer) ToOutput(ctx context.Context) pulumix.Output[*VolumeContainer] {
	return pulumix.Output[*VolumeContainer]{
		OutputState: i.ToVolumeContainerOutputWithContext(ctx).OutputState,
	}
}

type VolumeContainerOutput struct{ *pulumi.OutputState }

func (VolumeContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeContainer)(nil)).Elem()
}

func (o VolumeContainerOutput) ToVolumeContainerOutput() VolumeContainerOutput {
	return o
}

func (o VolumeContainerOutput) ToVolumeContainerOutputWithContext(ctx context.Context) VolumeContainerOutput {
	return o
}

func (o VolumeContainerOutput) ToOutput(ctx context.Context) pulumix.Output[*VolumeContainer] {
	return pulumix.Output[*VolumeContainer]{
		OutputState: o.OutputState,
	}
}

// The bandwidth-rate set on the volume container.
func (o VolumeContainerOutput) BandWidthRateInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.IntPtrOutput { return v.BandWidthRateInMbps }).(pulumi.IntPtrOutput)
}

// The ID of the bandwidth setting associated with the volume container.
func (o VolumeContainerOutput) BandwidthSettingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringPtrOutput { return v.BandwidthSettingId }).(pulumi.StringPtrOutput)
}

// The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
func (o VolumeContainerOutput) EncryptionKey() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v *VolumeContainer) AsymmetricEncryptedSecretResponsePtrOutput { return v.EncryptionKey }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

// The flag to denote whether encryption is enabled or not.
func (o VolumeContainerOutput) EncryptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringOutput { return v.EncryptionStatus }).(pulumi.StringOutput)
}

// The Kind of the object. Currently only Series8000 is supported
func (o VolumeContainerOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The name of the object.
func (o VolumeContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The owner ship status of the volume container. Only when the status is "NotOwned", the delete operation on the volume container is permitted.
func (o VolumeContainerOutput) OwnerShipStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringOutput { return v.OwnerShipStatus }).(pulumi.StringOutput)
}

// The path ID of storage account associated with the volume container.
func (o VolumeContainerOutput) StorageAccountCredentialId() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringOutput { return v.StorageAccountCredentialId }).(pulumi.StringOutput)
}

// The total cloud storage for the volume container.
func (o VolumeContainerOutput) TotalCloudStorageUsageInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *VolumeContainer) pulumi.Float64Output { return v.TotalCloudStorageUsageInBytes }).(pulumi.Float64Output)
}

// The hierarchical type of the object.
func (o VolumeContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The number of volumes in the volume Container.
func (o VolumeContainerOutput) VolumeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.IntOutput { return v.VolumeCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeContainerOutput{})
}
