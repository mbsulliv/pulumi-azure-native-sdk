// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurestackhci

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The storage container resource definition.
// Azure REST API version: 2022-12-15-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2024-01-01.
type StorageContainer struct {
	pulumi.CustomResourceState

	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Path of the storage container on the disk
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Provisioning state of the storage container.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The observed state of storage containers
	Status StorageContainerStatusResponseOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStorageContainer registers a new resource with the given unique name, arguments, and options.
func NewStorageContainer(ctx *pulumi.Context,
	name string, args *StorageContainerArgs, opts ...pulumi.ResourceOption) (*StorageContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:StorageContainer"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221215preview:StorageContainer"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230701preview:StorageContainer"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230901preview:StorageContainer"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240101:StorageContainer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StorageContainer
	err := ctx.RegisterResource("azure-native:azurestackhci:StorageContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageContainer gets an existing StorageContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageContainerState, opts ...pulumi.ResourceOption) (*StorageContainer, error) {
	var resource StorageContainer
	err := ctx.ReadResource("azure-native:azurestackhci:StorageContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageContainer resources.
type storageContainerState struct {
}

type StorageContainerState struct {
}

func (StorageContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageContainerState)(nil)).Elem()
}

type storageContainerArgs struct {
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Path of the storage container on the disk
	Path *string `pulumi:"path"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the storage container
	StorageContainerName *string `pulumi:"storageContainerName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StorageContainer resource.
type StorageContainerArgs struct {
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Path of the storage container on the disk
	Path pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of the storage container
	StorageContainerName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (StorageContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageContainerArgs)(nil)).Elem()
}

type StorageContainerInput interface {
	pulumi.Input

	ToStorageContainerOutput() StorageContainerOutput
	ToStorageContainerOutputWithContext(ctx context.Context) StorageContainerOutput
}

func (*StorageContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageContainer)(nil)).Elem()
}

func (i *StorageContainer) ToStorageContainerOutput() StorageContainerOutput {
	return i.ToStorageContainerOutputWithContext(context.Background())
}

func (i *StorageContainer) ToStorageContainerOutputWithContext(ctx context.Context) StorageContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageContainerOutput)
}

type StorageContainerOutput struct{ *pulumi.OutputState }

func (StorageContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageContainer)(nil)).Elem()
}

func (o StorageContainerOutput) ToStorageContainerOutput() StorageContainerOutput {
	return o
}

func (o StorageContainerOutput) ToStorageContainerOutputWithContext(ctx context.Context) StorageContainerOutput {
	return o
}

// The extendedLocation of the resource.
func (o StorageContainerOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *StorageContainer) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The geo-location where the resource lives
func (o StorageContainerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageContainer) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o StorageContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Path of the storage container on the disk
func (o StorageContainerOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageContainer) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// Provisioning state of the storage container.
func (o StorageContainerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageContainer) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The observed state of storage containers
func (o StorageContainerOutput) Status() StorageContainerStatusResponseOutput {
	return o.ApplyT(func(v *StorageContainer) StorageContainerStatusResponseOutput { return v.Status }).(StorageContainerStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o StorageContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StorageContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o StorageContainerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageContainer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o StorageContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageContainerOutput{})
}
