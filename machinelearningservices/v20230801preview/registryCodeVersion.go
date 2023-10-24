// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
type RegistryCodeVersion struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	CodeVersionProperties CodeVersionResponseOutput `pulumi:"codeVersionProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistryCodeVersion registers a new resource with the given unique name, arguments, and options.
func NewRegistryCodeVersion(ctx *pulumi.Context,
	name string, args *RegistryCodeVersionArgs, opts ...pulumi.ResourceOption) (*RegistryCodeVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CodeName == nil {
		return nil, errors.New("invalid value for required argument 'CodeName'")
	}
	if args.CodeVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'CodeVersionProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.CodeVersionProperties = args.CodeVersionProperties.ToCodeVersionTypeOutput().ApplyT(func(v CodeVersionType) CodeVersionType { return *v.Defaults() }).(CodeVersionTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:RegistryCodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:RegistryCodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:RegistryCodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:RegistryCodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:RegistryCodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:RegistryCodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:RegistryCodeVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:RegistryCodeVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RegistryCodeVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230801preview:RegistryCodeVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryCodeVersion gets an existing RegistryCodeVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryCodeVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryCodeVersionState, opts ...pulumi.ResourceOption) (*RegistryCodeVersion, error) {
	var resource RegistryCodeVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230801preview:RegistryCodeVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryCodeVersion resources.
type registryCodeVersionState struct {
}

type RegistryCodeVersionState struct {
}

func (RegistryCodeVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryCodeVersionState)(nil)).Elem()
}

type registryCodeVersionArgs struct {
	// Container name.
	CodeName string `pulumi:"codeName"`
	// [Required] Additional attributes of the entity.
	CodeVersionProperties CodeVersionType `pulumi:"codeVersionProperties"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a RegistryCodeVersion resource.
type RegistryCodeVersionArgs struct {
	// Container name.
	CodeName pulumi.StringInput
	// [Required] Additional attributes of the entity.
	CodeVersionProperties CodeVersionTypeInput
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version identifier.
	Version pulumi.StringPtrInput
}

func (RegistryCodeVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryCodeVersionArgs)(nil)).Elem()
}

type RegistryCodeVersionInput interface {
	pulumi.Input

	ToRegistryCodeVersionOutput() RegistryCodeVersionOutput
	ToRegistryCodeVersionOutputWithContext(ctx context.Context) RegistryCodeVersionOutput
}

func (*RegistryCodeVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryCodeVersion)(nil)).Elem()
}

func (i *RegistryCodeVersion) ToRegistryCodeVersionOutput() RegistryCodeVersionOutput {
	return i.ToRegistryCodeVersionOutputWithContext(context.Background())
}

func (i *RegistryCodeVersion) ToRegistryCodeVersionOutputWithContext(ctx context.Context) RegistryCodeVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryCodeVersionOutput)
}

func (i *RegistryCodeVersion) ToOutput(ctx context.Context) pulumix.Output[*RegistryCodeVersion] {
	return pulumix.Output[*RegistryCodeVersion]{
		OutputState: i.ToRegistryCodeVersionOutputWithContext(ctx).OutputState,
	}
}

type RegistryCodeVersionOutput struct{ *pulumi.OutputState }

func (RegistryCodeVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryCodeVersion)(nil)).Elem()
}

func (o RegistryCodeVersionOutput) ToRegistryCodeVersionOutput() RegistryCodeVersionOutput {
	return o
}

func (o RegistryCodeVersionOutput) ToRegistryCodeVersionOutputWithContext(ctx context.Context) RegistryCodeVersionOutput {
	return o
}

func (o RegistryCodeVersionOutput) ToOutput(ctx context.Context) pulumix.Output[*RegistryCodeVersion] {
	return pulumix.Output[*RegistryCodeVersion]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o RegistryCodeVersionOutput) CodeVersionProperties() CodeVersionResponseOutput {
	return o.ApplyT(func(v *RegistryCodeVersion) CodeVersionResponseOutput { return v.CodeVersionProperties }).(CodeVersionResponseOutput)
}

// The name of the resource
func (o RegistryCodeVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryCodeVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RegistryCodeVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryCodeVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RegistryCodeVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryCodeVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryCodeVersionOutput{})
}
