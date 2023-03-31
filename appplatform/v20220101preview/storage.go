// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Storage resource payload.
type Storage struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the storage resource payload.
	Properties StorageAccountResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStorage registers a new resource with the given unique name, arguments, and options.
func NewStorage(ctx *pulumi.Context,
	name string, args *StorageArgs, opts ...pulumi.ResourceOption) (*Storage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:Storage"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:Storage"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:Storage"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:Storage"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:Storage"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:Storage"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:Storage"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:Storage"),
		},
	})
	opts = append(opts, aliases)
	var resource Storage
	err := ctx.RegisterResource("azure-native:appplatform/v20220101preview:Storage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorage gets an existing Storage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageState, opts ...pulumi.ResourceOption) (*Storage, error) {
	var resource Storage
	err := ctx.ReadResource("azure-native:appplatform/v20220101preview:Storage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Storage resources.
type storageState struct {
}

type StorageState struct {
}

func (StorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageState)(nil)).Elem()
}

type storageArgs struct {
	// Properties of the storage resource payload.
	Properties *StorageAccount `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
	// The name of the storage resource.
	StorageName *string `pulumi:"storageName"`
}

// The set of arguments for constructing a Storage resource.
type StorageArgs struct {
	// Properties of the storage resource payload.
	Properties StorageAccountPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
	// The name of the storage resource.
	StorageName pulumi.StringPtrInput
}

func (StorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageArgs)(nil)).Elem()
}

type StorageInput interface {
	pulumi.Input

	ToStorageOutput() StorageOutput
	ToStorageOutputWithContext(ctx context.Context) StorageOutput
}

func (*Storage) ElementType() reflect.Type {
	return reflect.TypeOf((**Storage)(nil)).Elem()
}

func (i *Storage) ToStorageOutput() StorageOutput {
	return i.ToStorageOutputWithContext(context.Background())
}

func (i *Storage) ToStorageOutputWithContext(ctx context.Context) StorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageOutput)
}

type StorageOutput struct{ *pulumi.OutputState }

func (StorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Storage)(nil)).Elem()
}

func (o StorageOutput) ToStorageOutput() StorageOutput {
	return o
}

func (o StorageOutput) ToStorageOutputWithContext(ctx context.Context) StorageOutput {
	return o
}

// The name of the resource.
func (o StorageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the storage resource payload.
func (o StorageOutput) Properties() StorageAccountResponseOutput {
	return o.ApplyT(func(v *Storage) StorageAccountResponseOutput { return v.Properties }).(StorageAccountResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o StorageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Storage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o StorageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageOutput{})
}
