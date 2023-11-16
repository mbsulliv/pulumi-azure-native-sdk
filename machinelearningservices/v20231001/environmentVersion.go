// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
type EnvironmentVersion struct {
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

// NewEnvironmentVersion registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentVersion(ctx *pulumi.Context,
	name string, args *EnvironmentVersionArgs, opts ...pulumi.ResourceOption) (*EnvironmentVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentVersionProperties'")
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
	args.EnvironmentVersionProperties = args.EnvironmentVersionProperties.ToEnvironmentVersionTypeOutput().ApplyT(func(v EnvironmentVersionType) EnvironmentVersionType { return *v.Defaults() }).(EnvironmentVersionTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:EnvironmentVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:EnvironmentVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EnvironmentVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20231001:EnvironmentVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentVersion gets an existing EnvironmentVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentVersionState, opts ...pulumi.ResourceOption) (*EnvironmentVersion, error) {
	var resource EnvironmentVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20231001:EnvironmentVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentVersion resources.
type environmentVersionState struct {
}

type EnvironmentVersionState struct {
}

func (EnvironmentVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentVersionState)(nil)).Elem()
}

type environmentVersionArgs struct {
	// [Required] Additional attributes of the entity.
	EnvironmentVersionProperties EnvironmentVersionType `pulumi:"environmentVersionProperties"`
	// Name of EnvironmentVersion. This is case-sensitive.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version of EnvironmentVersion.
	Version *string `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a EnvironmentVersion resource.
type EnvironmentVersionArgs struct {
	// [Required] Additional attributes of the entity.
	EnvironmentVersionProperties EnvironmentVersionTypeInput
	// Name of EnvironmentVersion. This is case-sensitive.
	Name pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version of EnvironmentVersion.
	Version pulumi.StringPtrInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (EnvironmentVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentVersionArgs)(nil)).Elem()
}

type EnvironmentVersionInput interface {
	pulumi.Input

	ToEnvironmentVersionOutput() EnvironmentVersionOutput
	ToEnvironmentVersionOutputWithContext(ctx context.Context) EnvironmentVersionOutput
}

func (*EnvironmentVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentVersion)(nil)).Elem()
}

func (i *EnvironmentVersion) ToEnvironmentVersionOutput() EnvironmentVersionOutput {
	return i.ToEnvironmentVersionOutputWithContext(context.Background())
}

func (i *EnvironmentVersion) ToEnvironmentVersionOutputWithContext(ctx context.Context) EnvironmentVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVersionOutput)
}

type EnvironmentVersionOutput struct{ *pulumi.OutputState }

func (EnvironmentVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentVersion)(nil)).Elem()
}

func (o EnvironmentVersionOutput) ToEnvironmentVersionOutput() EnvironmentVersionOutput {
	return o
}

func (o EnvironmentVersionOutput) ToEnvironmentVersionOutputWithContext(ctx context.Context) EnvironmentVersionOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o EnvironmentVersionOutput) EnvironmentVersionProperties() EnvironmentVersionResponseOutput {
	return o.ApplyT(func(v *EnvironmentVersion) EnvironmentVersionResponseOutput { return v.EnvironmentVersionProperties }).(EnvironmentVersionResponseOutput)
}

// The name of the resource
func (o EnvironmentVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o EnvironmentVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EnvironmentVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EnvironmentVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentVersionOutput{})
}
