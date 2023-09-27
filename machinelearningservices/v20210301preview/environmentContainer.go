// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
type EnvironmentContainer struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	Properties EnvironmentContainerResponseOutput `pulumi:"properties"`
	// System data associated with resource provider
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEnvironmentContainer registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentContainer(ctx *pulumi.Context,
	name string, args *EnvironmentContainerArgs, opts ...pulumi.ResourceOption) (*EnvironmentContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:EnvironmentContainer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EnvironmentContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210301preview:EnvironmentContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentContainer gets an existing EnvironmentContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentContainerState, opts ...pulumi.ResourceOption) (*EnvironmentContainer, error) {
	var resource EnvironmentContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210301preview:EnvironmentContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentContainer resources.
type environmentContainerState struct {
}

type EnvironmentContainerState struct {
}

func (EnvironmentContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentContainerState)(nil)).Elem()
}

type environmentContainerArgs struct {
	// Container name.
	Name *string `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	Properties EnvironmentContainerType `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a EnvironmentContainer resource.
type EnvironmentContainerArgs struct {
	// Container name.
	Name pulumi.StringPtrInput
	// [Required] Additional attributes of the entity.
	Properties EnvironmentContainerTypeInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (EnvironmentContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentContainerArgs)(nil)).Elem()
}

type EnvironmentContainerInput interface {
	pulumi.Input

	ToEnvironmentContainerOutput() EnvironmentContainerOutput
	ToEnvironmentContainerOutputWithContext(ctx context.Context) EnvironmentContainerOutput
}

func (*EnvironmentContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentContainer)(nil)).Elem()
}

func (i *EnvironmentContainer) ToEnvironmentContainerOutput() EnvironmentContainerOutput {
	return i.ToEnvironmentContainerOutputWithContext(context.Background())
}

func (i *EnvironmentContainer) ToEnvironmentContainerOutputWithContext(ctx context.Context) EnvironmentContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentContainerOutput)
}

func (i *EnvironmentContainer) ToOutput(ctx context.Context) pulumix.Output[*EnvironmentContainer] {
	return pulumix.Output[*EnvironmentContainer]{
		OutputState: i.ToEnvironmentContainerOutputWithContext(ctx).OutputState,
	}
}

type EnvironmentContainerOutput struct{ *pulumi.OutputState }

func (EnvironmentContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentContainer)(nil)).Elem()
}

func (o EnvironmentContainerOutput) ToEnvironmentContainerOutput() EnvironmentContainerOutput {
	return o
}

func (o EnvironmentContainerOutput) ToEnvironmentContainerOutputWithContext(ctx context.Context) EnvironmentContainerOutput {
	return o
}

func (o EnvironmentContainerOutput) ToOutput(ctx context.Context) pulumix.Output[*EnvironmentContainer] {
	return pulumix.Output[*EnvironmentContainer]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o EnvironmentContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o EnvironmentContainerOutput) Properties() EnvironmentContainerResponseOutput {
	return o.ApplyT(func(v *EnvironmentContainer) EnvironmentContainerResponseOutput { return v.Properties }).(EnvironmentContainerResponseOutput)
}

// System data associated with resource provider
func (o EnvironmentContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EnvironmentContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EnvironmentContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentContainerOutput{})
}
