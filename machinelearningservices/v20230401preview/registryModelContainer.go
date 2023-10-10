// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
type RegistryModelContainer struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	ModelContainerProperties ModelContainerResponseOutput `pulumi:"modelContainerProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistryModelContainer registers a new resource with the given unique name, arguments, and options.
func NewRegistryModelContainer(ctx *pulumi.Context,
	name string, args *RegistryModelContainerArgs, opts ...pulumi.ResourceOption) (*RegistryModelContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ModelContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'ModelContainerProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.ModelContainerProperties = args.ModelContainerProperties.ToModelContainerTypeOutput().ApplyT(func(v ModelContainerType) ModelContainerType { return *v.Defaults() }).(ModelContainerTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:RegistryModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:RegistryModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:RegistryModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:RegistryModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:RegistryModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:RegistryModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:RegistryModelContainer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RegistryModelContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230401preview:RegistryModelContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryModelContainer gets an existing RegistryModelContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryModelContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryModelContainerState, opts ...pulumi.ResourceOption) (*RegistryModelContainer, error) {
	var resource RegistryModelContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230401preview:RegistryModelContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryModelContainer resources.
type registryModelContainerState struct {
}

type RegistryModelContainerState struct {
}

func (RegistryModelContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryModelContainerState)(nil)).Elem()
}

type registryModelContainerArgs struct {
	// [Required] Additional attributes of the entity.
	ModelContainerProperties ModelContainerType `pulumi:"modelContainerProperties"`
	// Container name.
	ModelName *string `pulumi:"modelName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a RegistryModelContainer resource.
type RegistryModelContainerArgs struct {
	// [Required] Additional attributes of the entity.
	ModelContainerProperties ModelContainerTypeInput
	// Container name.
	ModelName pulumi.StringPtrInput
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (RegistryModelContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryModelContainerArgs)(nil)).Elem()
}

type RegistryModelContainerInput interface {
	pulumi.Input

	ToRegistryModelContainerOutput() RegistryModelContainerOutput
	ToRegistryModelContainerOutputWithContext(ctx context.Context) RegistryModelContainerOutput
}

func (*RegistryModelContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryModelContainer)(nil)).Elem()
}

func (i *RegistryModelContainer) ToRegistryModelContainerOutput() RegistryModelContainerOutput {
	return i.ToRegistryModelContainerOutputWithContext(context.Background())
}

func (i *RegistryModelContainer) ToRegistryModelContainerOutputWithContext(ctx context.Context) RegistryModelContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryModelContainerOutput)
}

func (i *RegistryModelContainer) ToOutput(ctx context.Context) pulumix.Output[*RegistryModelContainer] {
	return pulumix.Output[*RegistryModelContainer]{
		OutputState: i.ToRegistryModelContainerOutputWithContext(ctx).OutputState,
	}
}

type RegistryModelContainerOutput struct{ *pulumi.OutputState }

func (RegistryModelContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryModelContainer)(nil)).Elem()
}

func (o RegistryModelContainerOutput) ToRegistryModelContainerOutput() RegistryModelContainerOutput {
	return o
}

func (o RegistryModelContainerOutput) ToRegistryModelContainerOutputWithContext(ctx context.Context) RegistryModelContainerOutput {
	return o
}

func (o RegistryModelContainerOutput) ToOutput(ctx context.Context) pulumix.Output[*RegistryModelContainer] {
	return pulumix.Output[*RegistryModelContainer]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o RegistryModelContainerOutput) ModelContainerProperties() ModelContainerResponseOutput {
	return o.ApplyT(func(v *RegistryModelContainer) ModelContainerResponseOutput { return v.ModelContainerProperties }).(ModelContainerResponseOutput)
}

// The name of the resource
func (o RegistryModelContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryModelContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RegistryModelContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryModelContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RegistryModelContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryModelContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryModelContainerOutput{})
}
