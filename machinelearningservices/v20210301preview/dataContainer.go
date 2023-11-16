// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
type DataContainer struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	Properties DataContainerResponseOutput `pulumi:"properties"`
	// System data associated with resource provider
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataContainer registers a new resource with the given unique name, arguments, and options.
func NewDataContainer(ctx *pulumi.Context,
	name string, args *DataContainerArgs, opts ...pulumi.ResourceOption) (*DataContainer, error) {
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
			Type: pulumi.String("azure-native:machinelearningservices:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:DataContainer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DataContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210301preview:DataContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataContainer gets an existing DataContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataContainerState, opts ...pulumi.ResourceOption) (*DataContainer, error) {
	var resource DataContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210301preview:DataContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataContainer resources.
type dataContainerState struct {
}

type DataContainerState struct {
}

func (DataContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataContainerState)(nil)).Elem()
}

type dataContainerArgs struct {
	// Container name.
	Name *string `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	Properties DataContainerType `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a DataContainer resource.
type DataContainerArgs struct {
	// Container name.
	Name pulumi.StringPtrInput
	// [Required] Additional attributes of the entity.
	Properties DataContainerTypeInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (DataContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataContainerArgs)(nil)).Elem()
}

type DataContainerInput interface {
	pulumi.Input

	ToDataContainerOutput() DataContainerOutput
	ToDataContainerOutputWithContext(ctx context.Context) DataContainerOutput
}

func (*DataContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**DataContainer)(nil)).Elem()
}

func (i *DataContainer) ToDataContainerOutput() DataContainerOutput {
	return i.ToDataContainerOutputWithContext(context.Background())
}

func (i *DataContainer) ToDataContainerOutputWithContext(ctx context.Context) DataContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataContainerOutput)
}

type DataContainerOutput struct{ *pulumi.OutputState }

func (DataContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataContainer)(nil)).Elem()
}

func (o DataContainerOutput) ToDataContainerOutput() DataContainerOutput {
	return o
}

func (o DataContainerOutput) ToDataContainerOutputWithContext(ctx context.Context) DataContainerOutput {
	return o
}

// The name of the resource
func (o DataContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o DataContainerOutput) Properties() DataContainerResponseOutput {
	return o.ApplyT(func(v *DataContainer) DataContainerResponseOutput { return v.Properties }).(DataContainerResponseOutput)
}

// System data associated with resource provider
func (o DataContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DataContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DataContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataContainerOutput{})
}
