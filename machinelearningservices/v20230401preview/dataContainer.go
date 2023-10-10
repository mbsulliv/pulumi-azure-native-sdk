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
type DataContainer struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	DataContainerProperties DataContainerResponseOutput `pulumi:"dataContainerProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
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

	if args.DataContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'DataContainerProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.DataContainerProperties = args.DataContainerProperties.ToDataContainerTypeOutput().ApplyT(func(v DataContainerType) DataContainerType { return *v.Defaults() }).(DataContainerTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:DataContainer"),
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
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:DataContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:DataContainer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DataContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230401preview:DataContainer", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230401preview:DataContainer", name, id, state, &resource, opts...)
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
	// [Required] Additional attributes of the entity.
	DataContainerProperties DataContainerType `pulumi:"dataContainerProperties"`
	// Container name.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a DataContainer resource.
type DataContainerArgs struct {
	// [Required] Additional attributes of the entity.
	DataContainerProperties DataContainerTypeInput
	// Container name.
	Name pulumi.StringPtrInput
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

func (i *DataContainer) ToOutput(ctx context.Context) pulumix.Output[*DataContainer] {
	return pulumix.Output[*DataContainer]{
		OutputState: i.ToDataContainerOutputWithContext(ctx).OutputState,
	}
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

func (o DataContainerOutput) ToOutput(ctx context.Context) pulumix.Output[*DataContainer] {
	return pulumix.Output[*DataContainer]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o DataContainerOutput) DataContainerProperties() DataContainerResponseOutput {
	return o.ApplyT(func(v *DataContainer) DataContainerResponseOutput { return v.DataContainerProperties }).(DataContainerResponseOutput)
}

// The name of the resource
func (o DataContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
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
