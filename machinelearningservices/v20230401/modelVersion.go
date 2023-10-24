// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
type ModelVersion struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	ModelVersionProperties ModelVersionResponseOutput `pulumi:"modelVersionProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewModelVersion registers a new resource with the given unique name, arguments, and options.
func NewModelVersion(ctx *pulumi.Context,
	name string, args *ModelVersionArgs, opts ...pulumi.ResourceOption) (*ModelVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ModelVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'ModelVersionProperties'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.ModelVersionProperties = args.ModelVersionProperties.ToModelVersionTypeOutput().ApplyT(func(v ModelVersionType) ModelVersionType { return *v.Defaults() }).(ModelVersionTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:ModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:ModelVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ModelVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230401:ModelVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModelVersion gets an existing ModelVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModelVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelVersionState, opts ...pulumi.ResourceOption) (*ModelVersion, error) {
	var resource ModelVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230401:ModelVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModelVersion resources.
type modelVersionState struct {
}

type ModelVersionState struct {
}

func (ModelVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelVersionState)(nil)).Elem()
}

type modelVersionArgs struct {
	// [Required] Additional attributes of the entity.
	ModelVersionProperties ModelVersionType `pulumi:"modelVersionProperties"`
	// Container name. This is case-sensitive.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier. This is case-sensitive.
	Version *string `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ModelVersion resource.
type ModelVersionArgs struct {
	// [Required] Additional attributes of the entity.
	ModelVersionProperties ModelVersionTypeInput
	// Container name. This is case-sensitive.
	Name pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version identifier. This is case-sensitive.
	Version pulumi.StringPtrInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (ModelVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelVersionArgs)(nil)).Elem()
}

type ModelVersionInput interface {
	pulumi.Input

	ToModelVersionOutput() ModelVersionOutput
	ToModelVersionOutputWithContext(ctx context.Context) ModelVersionOutput
}

func (*ModelVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelVersion)(nil)).Elem()
}

func (i *ModelVersion) ToModelVersionOutput() ModelVersionOutput {
	return i.ToModelVersionOutputWithContext(context.Background())
}

func (i *ModelVersion) ToModelVersionOutputWithContext(ctx context.Context) ModelVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelVersionOutput)
}

func (i *ModelVersion) ToOutput(ctx context.Context) pulumix.Output[*ModelVersion] {
	return pulumix.Output[*ModelVersion]{
		OutputState: i.ToModelVersionOutputWithContext(ctx).OutputState,
	}
}

type ModelVersionOutput struct{ *pulumi.OutputState }

func (ModelVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelVersion)(nil)).Elem()
}

func (o ModelVersionOutput) ToModelVersionOutput() ModelVersionOutput {
	return o
}

func (o ModelVersionOutput) ToModelVersionOutputWithContext(ctx context.Context) ModelVersionOutput {
	return o
}

func (o ModelVersionOutput) ToOutput(ctx context.Context) pulumix.Output[*ModelVersion] {
	return pulumix.Output[*ModelVersion]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o ModelVersionOutput) ModelVersionProperties() ModelVersionResponseOutput {
	return o.ApplyT(func(v *ModelVersion) ModelVersionResponseOutput { return v.ModelVersionProperties }).(ModelVersionResponseOutput)
}

// The name of the resource
func (o ModelVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ModelVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ModelVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ModelVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ModelVersionOutput{})
}
