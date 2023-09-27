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
type ModelContainer struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	Properties ModelContainerResponseOutput `pulumi:"properties"`
	// System data associated with resource provider
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewModelContainer registers a new resource with the given unique name, arguments, and options.
func NewModelContainer(ctx *pulumi.Context,
	name string, args *ModelContainerArgs, opts ...pulumi.ResourceOption) (*ModelContainer, error) {
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
			Type: pulumi.String("azure-native:machinelearningservices:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:ModelContainer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ModelContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210301preview:ModelContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModelContainer gets an existing ModelContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModelContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelContainerState, opts ...pulumi.ResourceOption) (*ModelContainer, error) {
	var resource ModelContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210301preview:ModelContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModelContainer resources.
type modelContainerState struct {
}

type ModelContainerState struct {
}

func (ModelContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelContainerState)(nil)).Elem()
}

type modelContainerArgs struct {
	// Container name.
	Name *string `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	Properties ModelContainerType `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ModelContainer resource.
type ModelContainerArgs struct {
	// Container name.
	Name pulumi.StringPtrInput
	// [Required] Additional attributes of the entity.
	Properties ModelContainerTypeInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (ModelContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelContainerArgs)(nil)).Elem()
}

type ModelContainerInput interface {
	pulumi.Input

	ToModelContainerOutput() ModelContainerOutput
	ToModelContainerOutputWithContext(ctx context.Context) ModelContainerOutput
}

func (*ModelContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelContainer)(nil)).Elem()
}

func (i *ModelContainer) ToModelContainerOutput() ModelContainerOutput {
	return i.ToModelContainerOutputWithContext(context.Background())
}

func (i *ModelContainer) ToModelContainerOutputWithContext(ctx context.Context) ModelContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelContainerOutput)
}

func (i *ModelContainer) ToOutput(ctx context.Context) pulumix.Output[*ModelContainer] {
	return pulumix.Output[*ModelContainer]{
		OutputState: i.ToModelContainerOutputWithContext(ctx).OutputState,
	}
}

type ModelContainerOutput struct{ *pulumi.OutputState }

func (ModelContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelContainer)(nil)).Elem()
}

func (o ModelContainerOutput) ToModelContainerOutput() ModelContainerOutput {
	return o
}

func (o ModelContainerOutput) ToModelContainerOutputWithContext(ctx context.Context) ModelContainerOutput {
	return o
}

func (o ModelContainerOutput) ToOutput(ctx context.Context) pulumix.Output[*ModelContainer] {
	return pulumix.Output[*ModelContainer]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o ModelContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o ModelContainerOutput) Properties() ModelContainerResponseOutput {
	return o.ApplyT(func(v *ModelContainer) ModelContainerResponseOutput { return v.Properties }).(ModelContainerResponseOutput)
}

// System data associated with resource provider
func (o ModelContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ModelContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ModelContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ModelContainerOutput{})
}
