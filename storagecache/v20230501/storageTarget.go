// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Type of the Storage Target.
type StorageTarget struct {
	pulumi.CustomResourceState

	// The percentage of cache space allocated for this storage target
	AllocationPercentage pulumi.IntOutput `pulumi:"allocationPercentage"`
	// Properties when targetType is blobNfs.
	BlobNfs BlobNfsTargetResponsePtrOutput `pulumi:"blobNfs"`
	// Properties when targetType is clfs.
	Clfs ClfsTargetResponsePtrOutput `pulumi:"clfs"`
	// List of cache namespace junctions to target for namespace associations.
	Junctions NamespaceJunctionResponseArrayOutput `pulumi:"junctions"`
	// Region name string.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the Storage Target.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties when targetType is nfs3.
	Nfs3 Nfs3TargetResponsePtrOutput `pulumi:"nfs3"`
	// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Storage target operational state.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Type of the Storage Target.
	TargetType pulumi.StringOutput `pulumi:"targetType"`
	// Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
	Type pulumi.StringOutput `pulumi:"type"`
	// Properties when targetType is unknown.
	Unknown UnknownTargetResponsePtrOutput `pulumi:"unknown"`
}

// NewStorageTarget registers a new resource with the given unique name, arguments, and options.
func NewStorageTarget(ctx *pulumi.Context,
	name string, args *StorageTargetArgs, opts ...pulumi.ResourceOption) (*StorageTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CacheName == nil {
		return nil, errors.New("invalid value for required argument 'CacheName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TargetType == nil {
		return nil, errors.New("invalid value for required argument 'TargetType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagecache:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20190801preview:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20191101:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20200301:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20201001:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20210301:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20210501:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20210901:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20220101:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20220501:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20230101:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20230301preview:StorageTarget"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StorageTarget
	err := ctx.RegisterResource("azure-native:storagecache/v20230501:StorageTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageTarget gets an existing StorageTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageTargetState, opts ...pulumi.ResourceOption) (*StorageTarget, error) {
	var resource StorageTarget
	err := ctx.ReadResource("azure-native:storagecache/v20230501:StorageTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageTarget resources.
type storageTargetState struct {
}

type StorageTargetState struct {
}

func (StorageTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageTargetState)(nil)).Elem()
}

type storageTargetArgs struct {
	// Properties when targetType is blobNfs.
	BlobNfs *BlobNfsTarget `pulumi:"blobNfs"`
	// Name of cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
	CacheName string `pulumi:"cacheName"`
	// Properties when targetType is clfs.
	Clfs *ClfsTarget `pulumi:"clfs"`
	// List of cache namespace junctions to target for namespace associations.
	Junctions []NamespaceJunction `pulumi:"junctions"`
	// Properties when targetType is nfs3.
	Nfs3 *Nfs3Target `pulumi:"nfs3"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Storage target operational state.
	State *string `pulumi:"state"`
	// Name of Storage Target.
	StorageTargetName *string `pulumi:"storageTargetName"`
	// Type of the Storage Target.
	TargetType string `pulumi:"targetType"`
	// Properties when targetType is unknown.
	Unknown *UnknownTarget `pulumi:"unknown"`
}

// The set of arguments for constructing a StorageTarget resource.
type StorageTargetArgs struct {
	// Properties when targetType is blobNfs.
	BlobNfs BlobNfsTargetPtrInput
	// Name of cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
	CacheName pulumi.StringInput
	// Properties when targetType is clfs.
	Clfs ClfsTargetPtrInput
	// List of cache namespace junctions to target for namespace associations.
	Junctions NamespaceJunctionArrayInput
	// Properties when targetType is nfs3.
	Nfs3 Nfs3TargetPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Storage target operational state.
	State pulumi.StringPtrInput
	// Name of Storage Target.
	StorageTargetName pulumi.StringPtrInput
	// Type of the Storage Target.
	TargetType pulumi.StringInput
	// Properties when targetType is unknown.
	Unknown UnknownTargetPtrInput
}

func (StorageTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageTargetArgs)(nil)).Elem()
}

type StorageTargetInput interface {
	pulumi.Input

	ToStorageTargetOutput() StorageTargetOutput
	ToStorageTargetOutputWithContext(ctx context.Context) StorageTargetOutput
}

func (*StorageTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageTarget)(nil)).Elem()
}

func (i *StorageTarget) ToStorageTargetOutput() StorageTargetOutput {
	return i.ToStorageTargetOutputWithContext(context.Background())
}

func (i *StorageTarget) ToStorageTargetOutputWithContext(ctx context.Context) StorageTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageTargetOutput)
}

func (i *StorageTarget) ToOutput(ctx context.Context) pulumix.Output[*StorageTarget] {
	return pulumix.Output[*StorageTarget]{
		OutputState: i.ToStorageTargetOutputWithContext(ctx).OutputState,
	}
}

type StorageTargetOutput struct{ *pulumi.OutputState }

func (StorageTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageTarget)(nil)).Elem()
}

func (o StorageTargetOutput) ToStorageTargetOutput() StorageTargetOutput {
	return o
}

func (o StorageTargetOutput) ToStorageTargetOutputWithContext(ctx context.Context) StorageTargetOutput {
	return o
}

func (o StorageTargetOutput) ToOutput(ctx context.Context) pulumix.Output[*StorageTarget] {
	return pulumix.Output[*StorageTarget]{
		OutputState: o.OutputState,
	}
}

// The percentage of cache space allocated for this storage target
func (o StorageTargetOutput) AllocationPercentage() pulumi.IntOutput {
	return o.ApplyT(func(v *StorageTarget) pulumi.IntOutput { return v.AllocationPercentage }).(pulumi.IntOutput)
}

// Properties when targetType is blobNfs.
func (o StorageTargetOutput) BlobNfs() BlobNfsTargetResponsePtrOutput {
	return o.ApplyT(func(v *StorageTarget) BlobNfsTargetResponsePtrOutput { return v.BlobNfs }).(BlobNfsTargetResponsePtrOutput)
}

// Properties when targetType is clfs.
func (o StorageTargetOutput) Clfs() ClfsTargetResponsePtrOutput {
	return o.ApplyT(func(v *StorageTarget) ClfsTargetResponsePtrOutput { return v.Clfs }).(ClfsTargetResponsePtrOutput)
}

// List of cache namespace junctions to target for namespace associations.
func (o StorageTargetOutput) Junctions() NamespaceJunctionResponseArrayOutput {
	return o.ApplyT(func(v *StorageTarget) NamespaceJunctionResponseArrayOutput { return v.Junctions }).(NamespaceJunctionResponseArrayOutput)
}

// Region name string.
func (o StorageTargetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageTarget) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the Storage Target.
func (o StorageTargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageTarget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties when targetType is nfs3.
func (o StorageTargetOutput) Nfs3() Nfs3TargetResponsePtrOutput {
	return o.ApplyT(func(v *StorageTarget) Nfs3TargetResponsePtrOutput { return v.Nfs3 }).(Nfs3TargetResponsePtrOutput)
}

// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
func (o StorageTargetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageTarget) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Storage target operational state.
func (o StorageTargetOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageTarget) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// The system meta data relating to this resource.
func (o StorageTargetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StorageTarget) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the Storage Target.
func (o StorageTargetOutput) TargetType() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageTarget) pulumi.StringOutput { return v.TargetType }).(pulumi.StringOutput)
}

// Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
func (o StorageTargetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageTarget) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Properties when targetType is unknown.
func (o StorageTargetOutput) Unknown() UnknownTargetResponsePtrOutput {
	return o.ApplyT(func(v *StorageTarget) UnknownTargetResponsePtrOutput { return v.Unknown }).(UnknownTargetResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageTargetOutput{})
}
