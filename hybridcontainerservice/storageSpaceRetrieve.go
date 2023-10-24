// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridcontainerservice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The storageSpaces resource definition.
// Azure REST API version: 2022-09-01-preview. Prior API version in Azure Native 1.x: 2022-05-01-preview.
type StorageSpaceRetrieve struct {
	pulumi.CustomResourceState

	ExtendedLocation StorageSpacesResponseExtendedLocationPtrOutput `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// HybridAKSStorageSpec defines the desired state of HybridAKSStorage
	Properties StorageSpacesPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStorageSpaceRetrieve registers a new resource with the given unique name, arguments, and options.
func NewStorageSpaceRetrieve(ctx *pulumi.Context,
	name string, args *StorageSpaceRetrieveArgs, opts ...pulumi.ResourceOption) (*StorageSpaceRetrieve, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcontainerservice:storageSpaceRetrieve"),
		},
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20220501preview:StorageSpaceRetrieve"),
		},
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20220501preview:storageSpaceRetrieve"),
		},
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20220901preview:StorageSpaceRetrieve"),
		},
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20220901preview:storageSpaceRetrieve"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StorageSpaceRetrieve
	err := ctx.RegisterResource("azure-native:hybridcontainerservice:StorageSpaceRetrieve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageSpaceRetrieve gets an existing StorageSpaceRetrieve resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageSpaceRetrieve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageSpaceRetrieveState, opts ...pulumi.ResourceOption) (*StorageSpaceRetrieve, error) {
	var resource StorageSpaceRetrieve
	err := ctx.ReadResource("azure-native:hybridcontainerservice:StorageSpaceRetrieve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageSpaceRetrieve resources.
type storageSpaceRetrieveState struct {
}

type StorageSpaceRetrieveState struct {
}

func (StorageSpaceRetrieveState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageSpaceRetrieveState)(nil)).Elem()
}

type storageSpaceRetrieveArgs struct {
	ExtendedLocation *StorageSpacesExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// HybridAKSStorageSpec defines the desired state of HybridAKSStorage
	Properties *StorageSpacesProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Parameter for the name of the storage object
	StorageSpacesName *string `pulumi:"storageSpacesName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StorageSpaceRetrieve resource.
type StorageSpaceRetrieveArgs struct {
	ExtendedLocation StorageSpacesExtendedLocationPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// HybridAKSStorageSpec defines the desired state of HybridAKSStorage
	Properties StorageSpacesPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Parameter for the name of the storage object
	StorageSpacesName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (StorageSpaceRetrieveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageSpaceRetrieveArgs)(nil)).Elem()
}

type StorageSpaceRetrieveInput interface {
	pulumi.Input

	ToStorageSpaceRetrieveOutput() StorageSpaceRetrieveOutput
	ToStorageSpaceRetrieveOutputWithContext(ctx context.Context) StorageSpaceRetrieveOutput
}

func (*StorageSpaceRetrieve) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpaceRetrieve)(nil)).Elem()
}

func (i *StorageSpaceRetrieve) ToStorageSpaceRetrieveOutput() StorageSpaceRetrieveOutput {
	return i.ToStorageSpaceRetrieveOutputWithContext(context.Background())
}

func (i *StorageSpaceRetrieve) ToStorageSpaceRetrieveOutputWithContext(ctx context.Context) StorageSpaceRetrieveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpaceRetrieveOutput)
}

func (i *StorageSpaceRetrieve) ToOutput(ctx context.Context) pulumix.Output[*StorageSpaceRetrieve] {
	return pulumix.Output[*StorageSpaceRetrieve]{
		OutputState: i.ToStorageSpaceRetrieveOutputWithContext(ctx).OutputState,
	}
}

type StorageSpaceRetrieveOutput struct{ *pulumi.OutputState }

func (StorageSpaceRetrieveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpaceRetrieve)(nil)).Elem()
}

func (o StorageSpaceRetrieveOutput) ToStorageSpaceRetrieveOutput() StorageSpaceRetrieveOutput {
	return o
}

func (o StorageSpaceRetrieveOutput) ToStorageSpaceRetrieveOutputWithContext(ctx context.Context) StorageSpaceRetrieveOutput {
	return o
}

func (o StorageSpaceRetrieveOutput) ToOutput(ctx context.Context) pulumix.Output[*StorageSpaceRetrieve] {
	return pulumix.Output[*StorageSpaceRetrieve]{
		OutputState: o.OutputState,
	}
}

func (o StorageSpaceRetrieveOutput) ExtendedLocation() StorageSpacesResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) StorageSpacesResponseExtendedLocationPtrOutput {
		return v.ExtendedLocation
	}).(StorageSpacesResponseExtendedLocationPtrOutput)
}

// The geo-location where the resource lives
func (o StorageSpaceRetrieveOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o StorageSpaceRetrieveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// HybridAKSStorageSpec defines the desired state of HybridAKSStorage
func (o StorageSpaceRetrieveOutput) Properties() StorageSpacesPropertiesResponseOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) StorageSpacesPropertiesResponseOutput { return v.Properties }).(StorageSpacesPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o StorageSpaceRetrieveOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o StorageSpaceRetrieveOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o StorageSpaceRetrieveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageSpaceRetrieveOutput{})
}
