// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
// Azure REST API version: 2023-04-01.
//
// Other available API versions: 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01.
type RegistryDataContainer struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	DataContainerProperties DataContainerResponseOutput `pulumi:"dataContainerProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistryDataContainer registers a new resource with the given unique name, arguments, and options.
func NewRegistryDataContainer(ctx *pulumi.Context,
	name string, args *RegistryDataContainerArgs, opts ...pulumi.ResourceOption) (*RegistryDataContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'DataContainerProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.DataContainerProperties = args.DataContainerProperties.ToDataContainerTypeOutput().ApplyT(func(v DataContainerType) DataContainerType { return *v.Defaults() }).(DataContainerTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:RegistryDataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:RegistryDataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:RegistryDataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:RegistryDataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:RegistryDataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:RegistryDataContainer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RegistryDataContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices:RegistryDataContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryDataContainer gets an existing RegistryDataContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryDataContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryDataContainerState, opts ...pulumi.ResourceOption) (*RegistryDataContainer, error) {
	var resource RegistryDataContainer
	err := ctx.ReadResource("azure-native:machinelearningservices:RegistryDataContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryDataContainer resources.
type registryDataContainerState struct {
}

type RegistryDataContainerState struct {
}

func (RegistryDataContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryDataContainerState)(nil)).Elem()
}

type registryDataContainerArgs struct {
	// [Required] Additional attributes of the entity.
	DataContainerProperties DataContainerType `pulumi:"dataContainerProperties"`
	// Container name.
	Name *string `pulumi:"name"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a RegistryDataContainer resource.
type RegistryDataContainerArgs struct {
	// [Required] Additional attributes of the entity.
	DataContainerProperties DataContainerTypeInput
	// Container name.
	Name pulumi.StringPtrInput
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (RegistryDataContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryDataContainerArgs)(nil)).Elem()
}

type RegistryDataContainerInput interface {
	pulumi.Input

	ToRegistryDataContainerOutput() RegistryDataContainerOutput
	ToRegistryDataContainerOutputWithContext(ctx context.Context) RegistryDataContainerOutput
}

func (*RegistryDataContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryDataContainer)(nil)).Elem()
}

func (i *RegistryDataContainer) ToRegistryDataContainerOutput() RegistryDataContainerOutput {
	return i.ToRegistryDataContainerOutputWithContext(context.Background())
}

func (i *RegistryDataContainer) ToRegistryDataContainerOutputWithContext(ctx context.Context) RegistryDataContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryDataContainerOutput)
}

type RegistryDataContainerOutput struct{ *pulumi.OutputState }

func (RegistryDataContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryDataContainer)(nil)).Elem()
}

func (o RegistryDataContainerOutput) ToRegistryDataContainerOutput() RegistryDataContainerOutput {
	return o
}

func (o RegistryDataContainerOutput) ToRegistryDataContainerOutputWithContext(ctx context.Context) RegistryDataContainerOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o RegistryDataContainerOutput) DataContainerProperties() DataContainerResponseOutput {
	return o.ApplyT(func(v *RegistryDataContainer) DataContainerResponseOutput { return v.DataContainerProperties }).(DataContainerResponseOutput)
}

// The name of the resource
func (o RegistryDataContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryDataContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RegistryDataContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryDataContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RegistryDataContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryDataContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryDataContainerOutput{})
}
