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
type RegistryEnvironmentContainer struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	EnvironmentContainerProperties EnvironmentContainerResponseOutput `pulumi:"environmentContainerProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistryEnvironmentContainer registers a new resource with the given unique name, arguments, and options.
func NewRegistryEnvironmentContainer(ctx *pulumi.Context,
	name string, args *RegistryEnvironmentContainerArgs, opts ...pulumi.ResourceOption) (*RegistryEnvironmentContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentContainerProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.EnvironmentContainerProperties = args.EnvironmentContainerProperties.ToEnvironmentContainerTypeOutput().ApplyT(func(v EnvironmentContainerType) EnvironmentContainerType { return *v.Defaults() }).(EnvironmentContainerTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:RegistryEnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:RegistryEnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:RegistryEnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:RegistryEnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:RegistryEnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:RegistryEnvironmentContainer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RegistryEnvironmentContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230401preview:RegistryEnvironmentContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryEnvironmentContainer gets an existing RegistryEnvironmentContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryEnvironmentContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryEnvironmentContainerState, opts ...pulumi.ResourceOption) (*RegistryEnvironmentContainer, error) {
	var resource RegistryEnvironmentContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230401preview:RegistryEnvironmentContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryEnvironmentContainer resources.
type registryEnvironmentContainerState struct {
}

type RegistryEnvironmentContainerState struct {
}

func (RegistryEnvironmentContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnvironmentContainerState)(nil)).Elem()
}

type registryEnvironmentContainerArgs struct {
	// [Required] Additional attributes of the entity.
	EnvironmentContainerProperties EnvironmentContainerType `pulumi:"environmentContainerProperties"`
	// Container name.
	EnvironmentName *string `pulumi:"environmentName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a RegistryEnvironmentContainer resource.
type RegistryEnvironmentContainerArgs struct {
	// [Required] Additional attributes of the entity.
	EnvironmentContainerProperties EnvironmentContainerTypeInput
	// Container name.
	EnvironmentName pulumi.StringPtrInput
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (RegistryEnvironmentContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnvironmentContainerArgs)(nil)).Elem()
}

type RegistryEnvironmentContainerInput interface {
	pulumi.Input

	ToRegistryEnvironmentContainerOutput() RegistryEnvironmentContainerOutput
	ToRegistryEnvironmentContainerOutputWithContext(ctx context.Context) RegistryEnvironmentContainerOutput
}

func (*RegistryEnvironmentContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnvironmentContainer)(nil)).Elem()
}

func (i *RegistryEnvironmentContainer) ToRegistryEnvironmentContainerOutput() RegistryEnvironmentContainerOutput {
	return i.ToRegistryEnvironmentContainerOutputWithContext(context.Background())
}

func (i *RegistryEnvironmentContainer) ToRegistryEnvironmentContainerOutputWithContext(ctx context.Context) RegistryEnvironmentContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnvironmentContainerOutput)
}

func (i *RegistryEnvironmentContainer) ToOutput(ctx context.Context) pulumix.Output[*RegistryEnvironmentContainer] {
	return pulumix.Output[*RegistryEnvironmentContainer]{
		OutputState: i.ToRegistryEnvironmentContainerOutputWithContext(ctx).OutputState,
	}
}

type RegistryEnvironmentContainerOutput struct{ *pulumi.OutputState }

func (RegistryEnvironmentContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnvironmentContainer)(nil)).Elem()
}

func (o RegistryEnvironmentContainerOutput) ToRegistryEnvironmentContainerOutput() RegistryEnvironmentContainerOutput {
	return o
}

func (o RegistryEnvironmentContainerOutput) ToRegistryEnvironmentContainerOutputWithContext(ctx context.Context) RegistryEnvironmentContainerOutput {
	return o
}

func (o RegistryEnvironmentContainerOutput) ToOutput(ctx context.Context) pulumix.Output[*RegistryEnvironmentContainer] {
	return pulumix.Output[*RegistryEnvironmentContainer]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o RegistryEnvironmentContainerOutput) EnvironmentContainerProperties() EnvironmentContainerResponseOutput {
	return o.ApplyT(func(v *RegistryEnvironmentContainer) EnvironmentContainerResponseOutput {
		return v.EnvironmentContainerProperties
	}).(EnvironmentContainerResponseOutput)
}

// The name of the resource
func (o RegistryEnvironmentContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnvironmentContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RegistryEnvironmentContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryEnvironmentContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RegistryEnvironmentContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnvironmentContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryEnvironmentContainerOutput{})
}
