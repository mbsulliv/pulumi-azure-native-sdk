// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
type RegistryCodeContainer struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	CodeContainerProperties CodeContainerResponseOutput `pulumi:"codeContainerProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistryCodeContainer registers a new resource with the given unique name, arguments, and options.
func NewRegistryCodeContainer(ctx *pulumi.Context,
	name string, args *RegistryCodeContainerArgs, opts ...pulumi.ResourceOption) (*RegistryCodeContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CodeContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'CodeContainerProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.CodeContainerProperties = args.CodeContainerProperties.ToCodeContainerTypeOutput().ApplyT(func(v CodeContainerType) CodeContainerType { return *v.Defaults() }).(CodeContainerTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:RegistryCodeContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource RegistryCodeContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20221201preview:RegistryCodeContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryCodeContainer gets an existing RegistryCodeContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryCodeContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryCodeContainerState, opts ...pulumi.ResourceOption) (*RegistryCodeContainer, error) {
	var resource RegistryCodeContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20221201preview:RegistryCodeContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryCodeContainer resources.
type registryCodeContainerState struct {
}

type RegistryCodeContainerState struct {
}

func (RegistryCodeContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryCodeContainerState)(nil)).Elem()
}

type registryCodeContainerArgs struct {
	// [Required] Additional attributes of the entity.
	CodeContainerProperties CodeContainerType `pulumi:"codeContainerProperties"`
	// Container name.
	CodeName *string `pulumi:"codeName"`
	// Name of Azure Machine Learning registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a RegistryCodeContainer resource.
type RegistryCodeContainerArgs struct {
	// [Required] Additional attributes of the entity.
	CodeContainerProperties CodeContainerTypeInput
	// Container name.
	CodeName pulumi.StringPtrInput
	// Name of Azure Machine Learning registry.
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (RegistryCodeContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryCodeContainerArgs)(nil)).Elem()
}

type RegistryCodeContainerInput interface {
	pulumi.Input

	ToRegistryCodeContainerOutput() RegistryCodeContainerOutput
	ToRegistryCodeContainerOutputWithContext(ctx context.Context) RegistryCodeContainerOutput
}

func (*RegistryCodeContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryCodeContainer)(nil)).Elem()
}

func (i *RegistryCodeContainer) ToRegistryCodeContainerOutput() RegistryCodeContainerOutput {
	return i.ToRegistryCodeContainerOutputWithContext(context.Background())
}

func (i *RegistryCodeContainer) ToRegistryCodeContainerOutputWithContext(ctx context.Context) RegistryCodeContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryCodeContainerOutput)
}

type RegistryCodeContainerOutput struct{ *pulumi.OutputState }

func (RegistryCodeContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryCodeContainer)(nil)).Elem()
}

func (o RegistryCodeContainerOutput) ToRegistryCodeContainerOutput() RegistryCodeContainerOutput {
	return o
}

func (o RegistryCodeContainerOutput) ToRegistryCodeContainerOutputWithContext(ctx context.Context) RegistryCodeContainerOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o RegistryCodeContainerOutput) CodeContainerProperties() CodeContainerResponseOutput {
	return o.ApplyT(func(v *RegistryCodeContainer) CodeContainerResponseOutput { return v.CodeContainerProperties }).(CodeContainerResponseOutput)
}

// The name of the resource
func (o RegistryCodeContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryCodeContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RegistryCodeContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryCodeContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RegistryCodeContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryCodeContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryCodeContainerOutput{})
}