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
type EnvironmentSpecificationVersion struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	Properties EnvironmentSpecificationVersionResponseOutput `pulumi:"properties"`
	// System data associated with resource provider
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEnvironmentSpecificationVersion registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentSpecificationVersion(ctx *pulumi.Context,
	name string, args *EnvironmentSpecificationVersionArgs, opts ...pulumi.ResourceOption) (*EnvironmentSpecificationVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
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
			Type: pulumi.String("azure-native:machinelearningservices:EnvironmentSpecificationVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:EnvironmentSpecificationVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:EnvironmentSpecificationVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:EnvironmentSpecificationVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:EnvironmentSpecificationVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:EnvironmentSpecificationVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:EnvironmentSpecificationVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:EnvironmentSpecificationVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:EnvironmentSpecificationVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:EnvironmentSpecificationVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:EnvironmentSpecificationVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:EnvironmentSpecificationVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:EnvironmentSpecificationVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EnvironmentSpecificationVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210301preview:EnvironmentSpecificationVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentSpecificationVersion gets an existing EnvironmentSpecificationVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentSpecificationVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentSpecificationVersionState, opts ...pulumi.ResourceOption) (*EnvironmentSpecificationVersion, error) {
	var resource EnvironmentSpecificationVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210301preview:EnvironmentSpecificationVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentSpecificationVersion resources.
type environmentSpecificationVersionState struct {
}

type EnvironmentSpecificationVersionState struct {
}

func (EnvironmentSpecificationVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentSpecificationVersionState)(nil)).Elem()
}

type environmentSpecificationVersionArgs struct {
	// Name of EnvironmentSpecificationVersion.
	Name string `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	Properties EnvironmentSpecificationVersionType `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version of EnvironmentSpecificationVersion.
	Version *string `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a EnvironmentSpecificationVersion resource.
type EnvironmentSpecificationVersionArgs struct {
	// Name of EnvironmentSpecificationVersion.
	Name pulumi.StringInput
	// [Required] Additional attributes of the entity.
	Properties EnvironmentSpecificationVersionTypeInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version of EnvironmentSpecificationVersion.
	Version pulumi.StringPtrInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (EnvironmentSpecificationVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentSpecificationVersionArgs)(nil)).Elem()
}

type EnvironmentSpecificationVersionInput interface {
	pulumi.Input

	ToEnvironmentSpecificationVersionOutput() EnvironmentSpecificationVersionOutput
	ToEnvironmentSpecificationVersionOutputWithContext(ctx context.Context) EnvironmentSpecificationVersionOutput
}

func (*EnvironmentSpecificationVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentSpecificationVersion)(nil)).Elem()
}

func (i *EnvironmentSpecificationVersion) ToEnvironmentSpecificationVersionOutput() EnvironmentSpecificationVersionOutput {
	return i.ToEnvironmentSpecificationVersionOutputWithContext(context.Background())
}

func (i *EnvironmentSpecificationVersion) ToEnvironmentSpecificationVersionOutputWithContext(ctx context.Context) EnvironmentSpecificationVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentSpecificationVersionOutput)
}

type EnvironmentSpecificationVersionOutput struct{ *pulumi.OutputState }

func (EnvironmentSpecificationVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentSpecificationVersion)(nil)).Elem()
}

func (o EnvironmentSpecificationVersionOutput) ToEnvironmentSpecificationVersionOutput() EnvironmentSpecificationVersionOutput {
	return o
}

func (o EnvironmentSpecificationVersionOutput) ToEnvironmentSpecificationVersionOutputWithContext(ctx context.Context) EnvironmentSpecificationVersionOutput {
	return o
}

// The name of the resource
func (o EnvironmentSpecificationVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentSpecificationVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o EnvironmentSpecificationVersionOutput) Properties() EnvironmentSpecificationVersionResponseOutput {
	return o.ApplyT(func(v *EnvironmentSpecificationVersion) EnvironmentSpecificationVersionResponseOutput {
		return v.Properties
	}).(EnvironmentSpecificationVersionResponseOutput)
}

// System data associated with resource provider
func (o EnvironmentSpecificationVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EnvironmentSpecificationVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EnvironmentSpecificationVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentSpecificationVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentSpecificationVersionOutput{})
}
