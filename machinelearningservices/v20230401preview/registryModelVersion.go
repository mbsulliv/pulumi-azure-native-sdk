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
type RegistryModelVersion struct {
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

// NewRegistryModelVersion registers a new resource with the given unique name, arguments, and options.
func NewRegistryModelVersion(ctx *pulumi.Context,
	name string, args *RegistryModelVersionArgs, opts ...pulumi.ResourceOption) (*RegistryModelVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ModelName == nil {
		return nil, errors.New("invalid value for required argument 'ModelName'")
	}
	if args.ModelVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'ModelVersionProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.ModelVersionProperties = args.ModelVersionProperties.ToModelVersionTypeOutput().ApplyT(func(v ModelVersionType) ModelVersionType { return *v.Defaults() }).(ModelVersionTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:RegistryModelVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RegistryModelVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230401preview:RegistryModelVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryModelVersion gets an existing RegistryModelVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryModelVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryModelVersionState, opts ...pulumi.ResourceOption) (*RegistryModelVersion, error) {
	var resource RegistryModelVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230401preview:RegistryModelVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryModelVersion resources.
type registryModelVersionState struct {
}

type RegistryModelVersionState struct {
}

func (RegistryModelVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryModelVersionState)(nil)).Elem()
}

type registryModelVersionArgs struct {
	// Container name.
	ModelName string `pulumi:"modelName"`
	// [Required] Additional attributes of the entity.
	ModelVersionProperties ModelVersionType `pulumi:"modelVersionProperties"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a RegistryModelVersion resource.
type RegistryModelVersionArgs struct {
	// Container name.
	ModelName pulumi.StringInput
	// [Required] Additional attributes of the entity.
	ModelVersionProperties ModelVersionTypeInput
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version identifier.
	Version pulumi.StringPtrInput
}

func (RegistryModelVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryModelVersionArgs)(nil)).Elem()
}

type RegistryModelVersionInput interface {
	pulumi.Input

	ToRegistryModelVersionOutput() RegistryModelVersionOutput
	ToRegistryModelVersionOutputWithContext(ctx context.Context) RegistryModelVersionOutput
}

func (*RegistryModelVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryModelVersion)(nil)).Elem()
}

func (i *RegistryModelVersion) ToRegistryModelVersionOutput() RegistryModelVersionOutput {
	return i.ToRegistryModelVersionOutputWithContext(context.Background())
}

func (i *RegistryModelVersion) ToRegistryModelVersionOutputWithContext(ctx context.Context) RegistryModelVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryModelVersionOutput)
}

func (i *RegistryModelVersion) ToOutput(ctx context.Context) pulumix.Output[*RegistryModelVersion] {
	return pulumix.Output[*RegistryModelVersion]{
		OutputState: i.ToRegistryModelVersionOutputWithContext(ctx).OutputState,
	}
}

type RegistryModelVersionOutput struct{ *pulumi.OutputState }

func (RegistryModelVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryModelVersion)(nil)).Elem()
}

func (o RegistryModelVersionOutput) ToRegistryModelVersionOutput() RegistryModelVersionOutput {
	return o
}

func (o RegistryModelVersionOutput) ToRegistryModelVersionOutputWithContext(ctx context.Context) RegistryModelVersionOutput {
	return o
}

func (o RegistryModelVersionOutput) ToOutput(ctx context.Context) pulumix.Output[*RegistryModelVersion] {
	return pulumix.Output[*RegistryModelVersion]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o RegistryModelVersionOutput) ModelVersionProperties() ModelVersionResponseOutput {
	return o.ApplyT(func(v *RegistryModelVersion) ModelVersionResponseOutput { return v.ModelVersionProperties }).(ModelVersionResponseOutput)
}

// The name of the resource
func (o RegistryModelVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryModelVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RegistryModelVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryModelVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RegistryModelVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryModelVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryModelVersionOutput{})
}
