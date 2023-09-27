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
type RegistryComponentContainer struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	ComponentContainerProperties ComponentContainerResponseOutput `pulumi:"componentContainerProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistryComponentContainer registers a new resource with the given unique name, arguments, and options.
func NewRegistryComponentContainer(ctx *pulumi.Context,
	name string, args *RegistryComponentContainerArgs, opts ...pulumi.ResourceOption) (*RegistryComponentContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComponentContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'ComponentContainerProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.ComponentContainerProperties = args.ComponentContainerProperties.ToComponentContainerTypeOutput().ApplyT(func(v ComponentContainerType) ComponentContainerType { return *v.Defaults() }).(ComponentContainerTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:RegistryComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:RegistryComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:RegistryComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:RegistryComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:RegistryComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:RegistryComponentContainer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RegistryComponentContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230401preview:RegistryComponentContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryComponentContainer gets an existing RegistryComponentContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryComponentContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryComponentContainerState, opts ...pulumi.ResourceOption) (*RegistryComponentContainer, error) {
	var resource RegistryComponentContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230401preview:RegistryComponentContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryComponentContainer resources.
type registryComponentContainerState struct {
}

type RegistryComponentContainerState struct {
}

func (RegistryComponentContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryComponentContainerState)(nil)).Elem()
}

type registryComponentContainerArgs struct {
	// [Required] Additional attributes of the entity.
	ComponentContainerProperties ComponentContainerType `pulumi:"componentContainerProperties"`
	// Container name.
	ComponentName *string `pulumi:"componentName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a RegistryComponentContainer resource.
type RegistryComponentContainerArgs struct {
	// [Required] Additional attributes of the entity.
	ComponentContainerProperties ComponentContainerTypeInput
	// Container name.
	ComponentName pulumi.StringPtrInput
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (RegistryComponentContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryComponentContainerArgs)(nil)).Elem()
}

type RegistryComponentContainerInput interface {
	pulumi.Input

	ToRegistryComponentContainerOutput() RegistryComponentContainerOutput
	ToRegistryComponentContainerOutputWithContext(ctx context.Context) RegistryComponentContainerOutput
}

func (*RegistryComponentContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryComponentContainer)(nil)).Elem()
}

func (i *RegistryComponentContainer) ToRegistryComponentContainerOutput() RegistryComponentContainerOutput {
	return i.ToRegistryComponentContainerOutputWithContext(context.Background())
}

func (i *RegistryComponentContainer) ToRegistryComponentContainerOutputWithContext(ctx context.Context) RegistryComponentContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryComponentContainerOutput)
}

func (i *RegistryComponentContainer) ToOutput(ctx context.Context) pulumix.Output[*RegistryComponentContainer] {
	return pulumix.Output[*RegistryComponentContainer]{
		OutputState: i.ToRegistryComponentContainerOutputWithContext(ctx).OutputState,
	}
}

type RegistryComponentContainerOutput struct{ *pulumi.OutputState }

func (RegistryComponentContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryComponentContainer)(nil)).Elem()
}

func (o RegistryComponentContainerOutput) ToRegistryComponentContainerOutput() RegistryComponentContainerOutput {
	return o
}

func (o RegistryComponentContainerOutput) ToRegistryComponentContainerOutputWithContext(ctx context.Context) RegistryComponentContainerOutput {
	return o
}

func (o RegistryComponentContainerOutput) ToOutput(ctx context.Context) pulumix.Output[*RegistryComponentContainer] {
	return pulumix.Output[*RegistryComponentContainer]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o RegistryComponentContainerOutput) ComponentContainerProperties() ComponentContainerResponseOutput {
	return o.ApplyT(func(v *RegistryComponentContainer) ComponentContainerResponseOutput {
		return v.ComponentContainerProperties
	}).(ComponentContainerResponseOutput)
}

// The name of the resource
func (o RegistryComponentContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryComponentContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RegistryComponentContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryComponentContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RegistryComponentContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryComponentContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryComponentContainerOutput{})
}
