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
// Azure REST API version: 2023-04-01. Prior API version in Azure Native 1.x: 2021-03-01-preview.
//
// Other available API versions: 2021-03-01-preview, 2022-02-01-preview, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01.
type ModelContainer struct {
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

// NewModelContainer registers a new resource with the given unique name, arguments, and options.
func NewModelContainer(ctx *pulumi.Context,
	name string, args *ModelContainerArgs, opts ...pulumi.ResourceOption) (*ModelContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ModelContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'ModelContainerProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.ModelContainerProperties = args.ModelContainerProperties.ToModelContainerTypeOutput().ApplyT(func(v ModelContainerType) ModelContainerType { return *v.Defaults() }).(ModelContainerTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:ModelContainer"),
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
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:ModelContainer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ModelContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices:ModelContainer", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:machinelearningservices:ModelContainer", name, id, state, &resource, opts...)
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
	// [Required] Additional attributes of the entity.
	ModelContainerProperties ModelContainerType `pulumi:"modelContainerProperties"`
	// Container name. This is case-sensitive.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ModelContainer resource.
type ModelContainerArgs struct {
	// [Required] Additional attributes of the entity.
	ModelContainerProperties ModelContainerTypeInput
	// Container name. This is case-sensitive.
	Name pulumi.StringPtrInput
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

// [Required] Additional attributes of the entity.
func (o ModelContainerOutput) ModelContainerProperties() ModelContainerResponseOutput {
	return o.ApplyT(func(v *ModelContainer) ModelContainerResponseOutput { return v.ModelContainerProperties }).(ModelContainerResponseOutput)
}

// The name of the resource
func (o ModelContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
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
