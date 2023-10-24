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
//
// Other available API versions: 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01.
type RegistryDataVersion struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	DataVersionBaseProperties pulumi.AnyOutput `pulumi:"dataVersionBaseProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistryDataVersion registers a new resource with the given unique name, arguments, and options.
func NewRegistryDataVersion(ctx *pulumi.Context,
	name string, args *RegistryDataVersionArgs, opts ...pulumi.ResourceOption) (*RegistryDataVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataVersionBaseProperties == nil {
		return nil, errors.New("invalid value for required argument 'DataVersionBaseProperties'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:RegistryDataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:RegistryDataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:RegistryDataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:RegistryDataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:RegistryDataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:RegistryDataVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RegistryDataVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices:RegistryDataVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryDataVersion gets an existing RegistryDataVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryDataVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryDataVersionState, opts ...pulumi.ResourceOption) (*RegistryDataVersion, error) {
	var resource RegistryDataVersion
	err := ctx.ReadResource("azure-native:machinelearningservices:RegistryDataVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryDataVersion resources.
type registryDataVersionState struct {
}

type RegistryDataVersionState struct {
}

func (RegistryDataVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryDataVersionState)(nil)).Elem()
}

type registryDataVersionArgs struct {
	// [Required] Additional attributes of the entity.
	DataVersionBaseProperties interface{} `pulumi:"dataVersionBaseProperties"`
	// Container name.
	Name string `pulumi:"name"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a RegistryDataVersion resource.
type RegistryDataVersionArgs struct {
	// [Required] Additional attributes of the entity.
	DataVersionBaseProperties pulumi.Input
	// Container name.
	Name pulumi.StringInput
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version identifier.
	Version pulumi.StringPtrInput
}

func (RegistryDataVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryDataVersionArgs)(nil)).Elem()
}

type RegistryDataVersionInput interface {
	pulumi.Input

	ToRegistryDataVersionOutput() RegistryDataVersionOutput
	ToRegistryDataVersionOutputWithContext(ctx context.Context) RegistryDataVersionOutput
}

func (*RegistryDataVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryDataVersion)(nil)).Elem()
}

func (i *RegistryDataVersion) ToRegistryDataVersionOutput() RegistryDataVersionOutput {
	return i.ToRegistryDataVersionOutputWithContext(context.Background())
}

func (i *RegistryDataVersion) ToRegistryDataVersionOutputWithContext(ctx context.Context) RegistryDataVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryDataVersionOutput)
}

func (i *RegistryDataVersion) ToOutput(ctx context.Context) pulumix.Output[*RegistryDataVersion] {
	return pulumix.Output[*RegistryDataVersion]{
		OutputState: i.ToRegistryDataVersionOutputWithContext(ctx).OutputState,
	}
}

type RegistryDataVersionOutput struct{ *pulumi.OutputState }

func (RegistryDataVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryDataVersion)(nil)).Elem()
}

func (o RegistryDataVersionOutput) ToRegistryDataVersionOutput() RegistryDataVersionOutput {
	return o
}

func (o RegistryDataVersionOutput) ToRegistryDataVersionOutputWithContext(ctx context.Context) RegistryDataVersionOutput {
	return o
}

func (o RegistryDataVersionOutput) ToOutput(ctx context.Context) pulumix.Output[*RegistryDataVersion] {
	return pulumix.Output[*RegistryDataVersion]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o RegistryDataVersionOutput) DataVersionBaseProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v *RegistryDataVersion) pulumi.AnyOutput { return v.DataVersionBaseProperties }).(pulumi.AnyOutput)
}

// The name of the resource
func (o RegistryDataVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryDataVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RegistryDataVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryDataVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RegistryDataVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryDataVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryDataVersionOutput{})
}
