// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
type RegistryEnvironmentVersion struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	EnvironmentVersionProperties EnvironmentVersionResponseOutput `pulumi:"environmentVersionProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistryEnvironmentVersion registers a new resource with the given unique name, arguments, and options.
func NewRegistryEnvironmentVersion(ctx *pulumi.Context,
	name string, args *RegistryEnvironmentVersionArgs, opts ...pulumi.ResourceOption) (*RegistryEnvironmentVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.EnvironmentVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentVersionProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.EnvironmentVersionProperties = args.EnvironmentVersionProperties.ToEnvironmentVersionTypeOutput().ApplyT(func(v EnvironmentVersionType) EnvironmentVersionType { return *v.Defaults() }).(EnvironmentVersionTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:RegistryEnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:RegistryEnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:RegistryEnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:RegistryEnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:RegistryEnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:RegistryEnvironmentVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RegistryEnvironmentVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230601preview:RegistryEnvironmentVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryEnvironmentVersion gets an existing RegistryEnvironmentVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryEnvironmentVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryEnvironmentVersionState, opts ...pulumi.ResourceOption) (*RegistryEnvironmentVersion, error) {
	var resource RegistryEnvironmentVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230601preview:RegistryEnvironmentVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryEnvironmentVersion resources.
type registryEnvironmentVersionState struct {
}

type RegistryEnvironmentVersionState struct {
}

func (RegistryEnvironmentVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnvironmentVersionState)(nil)).Elem()
}

type registryEnvironmentVersionArgs struct {
	// Container name.
	EnvironmentName string `pulumi:"environmentName"`
	// [Required] Additional attributes of the entity.
	EnvironmentVersionProperties EnvironmentVersionType `pulumi:"environmentVersionProperties"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a RegistryEnvironmentVersion resource.
type RegistryEnvironmentVersionArgs struct {
	// Container name.
	EnvironmentName pulumi.StringInput
	// [Required] Additional attributes of the entity.
	EnvironmentVersionProperties EnvironmentVersionTypeInput
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version identifier.
	Version pulumi.StringPtrInput
}

func (RegistryEnvironmentVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnvironmentVersionArgs)(nil)).Elem()
}

type RegistryEnvironmentVersionInput interface {
	pulumi.Input

	ToRegistryEnvironmentVersionOutput() RegistryEnvironmentVersionOutput
	ToRegistryEnvironmentVersionOutputWithContext(ctx context.Context) RegistryEnvironmentVersionOutput
}

func (*RegistryEnvironmentVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnvironmentVersion)(nil)).Elem()
}

func (i *RegistryEnvironmentVersion) ToRegistryEnvironmentVersionOutput() RegistryEnvironmentVersionOutput {
	return i.ToRegistryEnvironmentVersionOutputWithContext(context.Background())
}

func (i *RegistryEnvironmentVersion) ToRegistryEnvironmentVersionOutputWithContext(ctx context.Context) RegistryEnvironmentVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnvironmentVersionOutput)
}

func (i *RegistryEnvironmentVersion) ToOutput(ctx context.Context) pulumix.Output[*RegistryEnvironmentVersion] {
	return pulumix.Output[*RegistryEnvironmentVersion]{
		OutputState: i.ToRegistryEnvironmentVersionOutputWithContext(ctx).OutputState,
	}
}

type RegistryEnvironmentVersionOutput struct{ *pulumi.OutputState }

func (RegistryEnvironmentVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnvironmentVersion)(nil)).Elem()
}

func (o RegistryEnvironmentVersionOutput) ToRegistryEnvironmentVersionOutput() RegistryEnvironmentVersionOutput {
	return o
}

func (o RegistryEnvironmentVersionOutput) ToRegistryEnvironmentVersionOutputWithContext(ctx context.Context) RegistryEnvironmentVersionOutput {
	return o
}

func (o RegistryEnvironmentVersionOutput) ToOutput(ctx context.Context) pulumix.Output[*RegistryEnvironmentVersion] {
	return pulumix.Output[*RegistryEnvironmentVersion]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o RegistryEnvironmentVersionOutput) EnvironmentVersionProperties() EnvironmentVersionResponseOutput {
	return o.ApplyT(func(v *RegistryEnvironmentVersion) EnvironmentVersionResponseOutput {
		return v.EnvironmentVersionProperties
	}).(EnvironmentVersionResponseOutput)
}

// The name of the resource
func (o RegistryEnvironmentVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnvironmentVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RegistryEnvironmentVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryEnvironmentVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RegistryEnvironmentVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnvironmentVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryEnvironmentVersionOutput{})
}
