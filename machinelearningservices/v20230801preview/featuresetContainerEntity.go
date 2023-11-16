// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
type FeaturesetContainerEntity struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	FeaturesetContainerProperties FeaturesetContainerResponseOutput `pulumi:"featuresetContainerProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFeaturesetContainerEntity registers a new resource with the given unique name, arguments, and options.
func NewFeaturesetContainerEntity(ctx *pulumi.Context,
	name string, args *FeaturesetContainerEntityArgs, opts ...pulumi.ResourceOption) (*FeaturesetContainerEntity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FeaturesetContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'FeaturesetContainerProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.FeaturesetContainerProperties = args.FeaturesetContainerProperties.ToFeaturesetContainerOutput().ApplyT(func(v FeaturesetContainer) FeaturesetContainer { return *v.Defaults() }).(FeaturesetContainerOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:FeaturesetContainerEntity"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:FeaturesetContainerEntity"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:FeaturesetContainerEntity"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:FeaturesetContainerEntity"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:FeaturesetContainerEntity"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FeaturesetContainerEntity
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230801preview:FeaturesetContainerEntity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeaturesetContainerEntity gets an existing FeaturesetContainerEntity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeaturesetContainerEntity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeaturesetContainerEntityState, opts ...pulumi.ResourceOption) (*FeaturesetContainerEntity, error) {
	var resource FeaturesetContainerEntity
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230801preview:FeaturesetContainerEntity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeaturesetContainerEntity resources.
type featuresetContainerEntityState struct {
}

type FeaturesetContainerEntityState struct {
}

func (FeaturesetContainerEntityState) ElementType() reflect.Type {
	return reflect.TypeOf((*featuresetContainerEntityState)(nil)).Elem()
}

type featuresetContainerEntityArgs struct {
	// [Required] Additional attributes of the entity.
	FeaturesetContainerProperties FeaturesetContainer `pulumi:"featuresetContainerProperties"`
	// Container name. This is case-sensitive.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a FeaturesetContainerEntity resource.
type FeaturesetContainerEntityArgs struct {
	// [Required] Additional attributes of the entity.
	FeaturesetContainerProperties FeaturesetContainerInput
	// Container name. This is case-sensitive.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (FeaturesetContainerEntityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featuresetContainerEntityArgs)(nil)).Elem()
}

type FeaturesetContainerEntityInput interface {
	pulumi.Input

	ToFeaturesetContainerEntityOutput() FeaturesetContainerEntityOutput
	ToFeaturesetContainerEntityOutputWithContext(ctx context.Context) FeaturesetContainerEntityOutput
}

func (*FeaturesetContainerEntity) ElementType() reflect.Type {
	return reflect.TypeOf((**FeaturesetContainerEntity)(nil)).Elem()
}

func (i *FeaturesetContainerEntity) ToFeaturesetContainerEntityOutput() FeaturesetContainerEntityOutput {
	return i.ToFeaturesetContainerEntityOutputWithContext(context.Background())
}

func (i *FeaturesetContainerEntity) ToFeaturesetContainerEntityOutputWithContext(ctx context.Context) FeaturesetContainerEntityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeaturesetContainerEntityOutput)
}

type FeaturesetContainerEntityOutput struct{ *pulumi.OutputState }

func (FeaturesetContainerEntityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeaturesetContainerEntity)(nil)).Elem()
}

func (o FeaturesetContainerEntityOutput) ToFeaturesetContainerEntityOutput() FeaturesetContainerEntityOutput {
	return o
}

func (o FeaturesetContainerEntityOutput) ToFeaturesetContainerEntityOutputWithContext(ctx context.Context) FeaturesetContainerEntityOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o FeaturesetContainerEntityOutput) FeaturesetContainerProperties() FeaturesetContainerResponseOutput {
	return o.ApplyT(func(v *FeaturesetContainerEntity) FeaturesetContainerResponseOutput {
		return v.FeaturesetContainerProperties
	}).(FeaturesetContainerResponseOutput)
}

// The name of the resource
func (o FeaturesetContainerEntityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FeaturesetContainerEntity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FeaturesetContainerEntityOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FeaturesetContainerEntity) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FeaturesetContainerEntityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FeaturesetContainerEntity) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FeaturesetContainerEntityOutput{})
}
