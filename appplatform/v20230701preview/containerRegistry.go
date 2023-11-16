// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Container registry resource payload.
type ContainerRegistry struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the container registry resource payload.
	Properties ContainerRegistryPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewContainerRegistry registers a new resource with the given unique name, arguments, and options.
func NewContainerRegistry(ctx *pulumi.Context,
	name string, args *ContainerRegistryArgs, opts ...pulumi.ResourceOption) (*ContainerRegistry, error) {
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
			Type: pulumi.String("azure-native:appplatform:ContainerRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230501preview:ContainerRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230901preview:ContainerRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20231101preview:ContainerRegistry"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ContainerRegistry
	err := ctx.RegisterResource("azure-native:appplatform/v20230701preview:ContainerRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRegistry gets an existing ContainerRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRegistryState, opts ...pulumi.ResourceOption) (*ContainerRegistry, error) {
	var resource ContainerRegistry
	err := ctx.ReadResource("azure-native:appplatform/v20230701preview:ContainerRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRegistry resources.
type containerRegistryState struct {
}

type ContainerRegistryState struct {
}

func (ContainerRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryState)(nil)).Elem()
}

type containerRegistryArgs struct {
	// The name of the container registry.
	ContainerRegistryName *string `pulumi:"containerRegistryName"`
	// Properties of the container registry resource payload.
	Properties *ContainerRegistryProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ContainerRegistry resource.
type ContainerRegistryArgs struct {
	// The name of the container registry.
	ContainerRegistryName pulumi.StringPtrInput
	// Properties of the container registry resource payload.
	Properties ContainerRegistryPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (ContainerRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryArgs)(nil)).Elem()
}

type ContainerRegistryInput interface {
	pulumi.Input

	ToContainerRegistryOutput() ContainerRegistryOutput
	ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput
}

func (*ContainerRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistry)(nil)).Elem()
}

func (i *ContainerRegistry) ToContainerRegistryOutput() ContainerRegistryOutput {
	return i.ToContainerRegistryOutputWithContext(context.Background())
}

func (i *ContainerRegistry) ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryOutput)
}

type ContainerRegistryOutput struct{ *pulumi.OutputState }

func (ContainerRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryOutput) ToContainerRegistryOutput() ContainerRegistryOutput {
	return o
}

func (o ContainerRegistryOutput) ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput {
	return o
}

// The name of the resource.
func (o ContainerRegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the container registry resource payload.
func (o ContainerRegistryOutput) Properties() ContainerRegistryPropertiesResponseOutput {
	return o.ApplyT(func(v *ContainerRegistry) ContainerRegistryPropertiesResponseOutput { return v.Properties }).(ContainerRegistryPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ContainerRegistryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ContainerRegistry) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o ContainerRegistryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerRegistryOutput{})
}
