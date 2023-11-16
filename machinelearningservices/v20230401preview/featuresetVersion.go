// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
type FeaturesetVersion struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	FeaturesetVersionProperties FeaturesetVersionResponseOutput `pulumi:"featuresetVersionProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFeaturesetVersion registers a new resource with the given unique name, arguments, and options.
func NewFeaturesetVersion(ctx *pulumi.Context,
	name string, args *FeaturesetVersionArgs, opts ...pulumi.ResourceOption) (*FeaturesetVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FeaturesetVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'FeaturesetVersionProperties'")
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
	args.FeaturesetVersionProperties = args.FeaturesetVersionProperties.ToFeaturesetVersionTypeOutput().ApplyT(func(v FeaturesetVersionType) FeaturesetVersionType { return *v.Defaults() }).(FeaturesetVersionTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:FeaturesetVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:FeaturesetVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:FeaturesetVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:FeaturesetVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:FeaturesetVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FeaturesetVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230401preview:FeaturesetVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeaturesetVersion gets an existing FeaturesetVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeaturesetVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeaturesetVersionState, opts ...pulumi.ResourceOption) (*FeaturesetVersion, error) {
	var resource FeaturesetVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230401preview:FeaturesetVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeaturesetVersion resources.
type featuresetVersionState struct {
}

type FeaturesetVersionState struct {
}

func (FeaturesetVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*featuresetVersionState)(nil)).Elem()
}

type featuresetVersionArgs struct {
	// [Required] Additional attributes of the entity.
	FeaturesetVersionProperties FeaturesetVersionType `pulumi:"featuresetVersionProperties"`
	// Container name. This is case-sensitive.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier. This is case-sensitive.
	Version *string `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a FeaturesetVersion resource.
type FeaturesetVersionArgs struct {
	// [Required] Additional attributes of the entity.
	FeaturesetVersionProperties FeaturesetVersionTypeInput
	// Container name. This is case-sensitive.
	Name pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version identifier. This is case-sensitive.
	Version pulumi.StringPtrInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (FeaturesetVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featuresetVersionArgs)(nil)).Elem()
}

type FeaturesetVersionInput interface {
	pulumi.Input

	ToFeaturesetVersionOutput() FeaturesetVersionOutput
	ToFeaturesetVersionOutputWithContext(ctx context.Context) FeaturesetVersionOutput
}

func (*FeaturesetVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**FeaturesetVersion)(nil)).Elem()
}

func (i *FeaturesetVersion) ToFeaturesetVersionOutput() FeaturesetVersionOutput {
	return i.ToFeaturesetVersionOutputWithContext(context.Background())
}

func (i *FeaturesetVersion) ToFeaturesetVersionOutputWithContext(ctx context.Context) FeaturesetVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeaturesetVersionOutput)
}

type FeaturesetVersionOutput struct{ *pulumi.OutputState }

func (FeaturesetVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeaturesetVersion)(nil)).Elem()
}

func (o FeaturesetVersionOutput) ToFeaturesetVersionOutput() FeaturesetVersionOutput {
	return o
}

func (o FeaturesetVersionOutput) ToFeaturesetVersionOutputWithContext(ctx context.Context) FeaturesetVersionOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o FeaturesetVersionOutput) FeaturesetVersionProperties() FeaturesetVersionResponseOutput {
	return o.ApplyT(func(v *FeaturesetVersion) FeaturesetVersionResponseOutput { return v.FeaturesetVersionProperties }).(FeaturesetVersionResponseOutput)
}

// The name of the resource
func (o FeaturesetVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FeaturesetVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FeaturesetVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FeaturesetVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FeaturesetVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FeaturesetVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FeaturesetVersionOutput{})
}
