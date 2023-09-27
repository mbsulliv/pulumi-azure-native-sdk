// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
type ComponentContainer struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	ComponentContainerDetails ComponentContainerResponseOutput `pulumi:"componentContainerDetails"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewComponentContainer registers a new resource with the given unique name, arguments, and options.
func NewComponentContainer(ctx *pulumi.Context,
	name string, args *ComponentContainerArgs, opts ...pulumi.ResourceOption) (*ComponentContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComponentContainerDetails == nil {
		return nil, errors.New("invalid value for required argument 'ComponentContainerDetails'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.ComponentContainerDetails = args.ComponentContainerDetails.ToComponentContainerTypeOutput().ApplyT(func(v ComponentContainerType) ComponentContainerType { return *v.Defaults() }).(ComponentContainerTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:ComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:ComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:ComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:ComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:ComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:ComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:ComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:ComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:ComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:ComponentContainer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ComponentContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220201preview:ComponentContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComponentContainer gets an existing ComponentContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComponentContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentContainerState, opts ...pulumi.ResourceOption) (*ComponentContainer, error) {
	var resource ComponentContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220201preview:ComponentContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComponentContainer resources.
type componentContainerState struct {
}

type ComponentContainerState struct {
}

func (ComponentContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentContainerState)(nil)).Elem()
}

type componentContainerArgs struct {
	// [Required] Additional attributes of the entity.
	ComponentContainerDetails ComponentContainerType `pulumi:"componentContainerDetails"`
	// Container name.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ComponentContainer resource.
type ComponentContainerArgs struct {
	// [Required] Additional attributes of the entity.
	ComponentContainerDetails ComponentContainerTypeInput
	// Container name.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (ComponentContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentContainerArgs)(nil)).Elem()
}

type ComponentContainerInput interface {
	pulumi.Input

	ToComponentContainerOutput() ComponentContainerOutput
	ToComponentContainerOutputWithContext(ctx context.Context) ComponentContainerOutput
}

func (*ComponentContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**ComponentContainer)(nil)).Elem()
}

func (i *ComponentContainer) ToComponentContainerOutput() ComponentContainerOutput {
	return i.ToComponentContainerOutputWithContext(context.Background())
}

func (i *ComponentContainer) ToComponentContainerOutputWithContext(ctx context.Context) ComponentContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentContainerOutput)
}

func (i *ComponentContainer) ToOutput(ctx context.Context) pulumix.Output[*ComponentContainer] {
	return pulumix.Output[*ComponentContainer]{
		OutputState: i.ToComponentContainerOutputWithContext(ctx).OutputState,
	}
}

type ComponentContainerOutput struct{ *pulumi.OutputState }

func (ComponentContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComponentContainer)(nil)).Elem()
}

func (o ComponentContainerOutput) ToComponentContainerOutput() ComponentContainerOutput {
	return o
}

func (o ComponentContainerOutput) ToComponentContainerOutputWithContext(ctx context.Context) ComponentContainerOutput {
	return o
}

func (o ComponentContainerOutput) ToOutput(ctx context.Context) pulumix.Output[*ComponentContainer] {
	return pulumix.Output[*ComponentContainer]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o ComponentContainerOutput) ComponentContainerDetails() ComponentContainerResponseOutput {
	return o.ApplyT(func(v *ComponentContainer) ComponentContainerResponseOutput { return v.ComponentContainerDetails }).(ComponentContainerResponseOutput)
}

// The name of the resource
func (o ComponentContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComponentContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ComponentContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ComponentContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ComponentContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ComponentContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ComponentContainerOutput{})
}
