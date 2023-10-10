// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
// Azure REST API version: 2023-04-01.
type RegistryComponentVersion struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	ComponentVersionProperties ComponentVersionResponseOutput `pulumi:"componentVersionProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistryComponentVersion registers a new resource with the given unique name, arguments, and options.
func NewRegistryComponentVersion(ctx *pulumi.Context,
	name string, args *RegistryComponentVersionArgs, opts ...pulumi.ResourceOption) (*RegistryComponentVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComponentName == nil {
		return nil, errors.New("invalid value for required argument 'ComponentName'")
	}
	if args.ComponentVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'ComponentVersionProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.ComponentVersionProperties = args.ComponentVersionProperties.ToComponentVersionTypeOutput().ApplyT(func(v ComponentVersionType) ComponentVersionType { return *v.Defaults() }).(ComponentVersionTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:RegistryComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:RegistryComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:RegistryComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:RegistryComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:RegistryComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:RegistryComponentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:RegistryComponentVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RegistryComponentVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices:RegistryComponentVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryComponentVersion gets an existing RegistryComponentVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryComponentVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryComponentVersionState, opts ...pulumi.ResourceOption) (*RegistryComponentVersion, error) {
	var resource RegistryComponentVersion
	err := ctx.ReadResource("azure-native:machinelearningservices:RegistryComponentVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryComponentVersion resources.
type registryComponentVersionState struct {
}

type RegistryComponentVersionState struct {
}

func (RegistryComponentVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryComponentVersionState)(nil)).Elem()
}

type registryComponentVersionArgs struct {
	// Container name.
	ComponentName string `pulumi:"componentName"`
	// [Required] Additional attributes of the entity.
	ComponentVersionProperties ComponentVersionType `pulumi:"componentVersionProperties"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a RegistryComponentVersion resource.
type RegistryComponentVersionArgs struct {
	// Container name.
	ComponentName pulumi.StringInput
	// [Required] Additional attributes of the entity.
	ComponentVersionProperties ComponentVersionTypeInput
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version identifier.
	Version pulumi.StringPtrInput
}

func (RegistryComponentVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryComponentVersionArgs)(nil)).Elem()
}

type RegistryComponentVersionInput interface {
	pulumi.Input

	ToRegistryComponentVersionOutput() RegistryComponentVersionOutput
	ToRegistryComponentVersionOutputWithContext(ctx context.Context) RegistryComponentVersionOutput
}

func (*RegistryComponentVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryComponentVersion)(nil)).Elem()
}

func (i *RegistryComponentVersion) ToRegistryComponentVersionOutput() RegistryComponentVersionOutput {
	return i.ToRegistryComponentVersionOutputWithContext(context.Background())
}

func (i *RegistryComponentVersion) ToRegistryComponentVersionOutputWithContext(ctx context.Context) RegistryComponentVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryComponentVersionOutput)
}

func (i *RegistryComponentVersion) ToOutput(ctx context.Context) pulumix.Output[*RegistryComponentVersion] {
	return pulumix.Output[*RegistryComponentVersion]{
		OutputState: i.ToRegistryComponentVersionOutputWithContext(ctx).OutputState,
	}
}

type RegistryComponentVersionOutput struct{ *pulumi.OutputState }

func (RegistryComponentVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryComponentVersion)(nil)).Elem()
}

func (o RegistryComponentVersionOutput) ToRegistryComponentVersionOutput() RegistryComponentVersionOutput {
	return o
}

func (o RegistryComponentVersionOutput) ToRegistryComponentVersionOutputWithContext(ctx context.Context) RegistryComponentVersionOutput {
	return o
}

func (o RegistryComponentVersionOutput) ToOutput(ctx context.Context) pulumix.Output[*RegistryComponentVersion] {
	return pulumix.Output[*RegistryComponentVersion]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o RegistryComponentVersionOutput) ComponentVersionProperties() ComponentVersionResponseOutput {
	return o.ApplyT(func(v *RegistryComponentVersion) ComponentVersionResponseOutput { return v.ComponentVersionProperties }).(ComponentVersionResponseOutput)
}

// The name of the resource
func (o RegistryComponentVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryComponentVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RegistryComponentVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryComponentVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RegistryComponentVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryComponentVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryComponentVersionOutput{})
}
