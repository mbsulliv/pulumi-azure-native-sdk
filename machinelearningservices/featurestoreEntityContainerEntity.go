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
// Azure REST API version: 2023-04-01-preview.
type FeaturestoreEntityContainerEntity struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	FeaturestoreEntityContainerProperties FeaturestoreEntityContainerResponseOutput `pulumi:"featurestoreEntityContainerProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFeaturestoreEntityContainerEntity registers a new resource with the given unique name, arguments, and options.
func NewFeaturestoreEntityContainerEntity(ctx *pulumi.Context,
	name string, args *FeaturestoreEntityContainerEntityArgs, opts ...pulumi.ResourceOption) (*FeaturestoreEntityContainerEntity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FeaturestoreEntityContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'FeaturestoreEntityContainerProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.FeaturestoreEntityContainerProperties = args.FeaturestoreEntityContainerProperties.ToFeaturestoreEntityContainerOutput().ApplyT(func(v FeaturestoreEntityContainer) FeaturestoreEntityContainer { return *v.Defaults() }).(FeaturestoreEntityContainerOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:FeaturestoreEntityContainerEntity"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:FeaturestoreEntityContainerEntity"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:FeaturestoreEntityContainerEntity"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FeaturestoreEntityContainerEntity
	err := ctx.RegisterResource("azure-native:machinelearningservices:FeaturestoreEntityContainerEntity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeaturestoreEntityContainerEntity gets an existing FeaturestoreEntityContainerEntity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeaturestoreEntityContainerEntity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeaturestoreEntityContainerEntityState, opts ...pulumi.ResourceOption) (*FeaturestoreEntityContainerEntity, error) {
	var resource FeaturestoreEntityContainerEntity
	err := ctx.ReadResource("azure-native:machinelearningservices:FeaturestoreEntityContainerEntity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeaturestoreEntityContainerEntity resources.
type featurestoreEntityContainerEntityState struct {
}

type FeaturestoreEntityContainerEntityState struct {
}

func (FeaturestoreEntityContainerEntityState) ElementType() reflect.Type {
	return reflect.TypeOf((*featurestoreEntityContainerEntityState)(nil)).Elem()
}

type featurestoreEntityContainerEntityArgs struct {
	// [Required] Additional attributes of the entity.
	FeaturestoreEntityContainerProperties FeaturestoreEntityContainer `pulumi:"featurestoreEntityContainerProperties"`
	// Container name. This is case-sensitive.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a FeaturestoreEntityContainerEntity resource.
type FeaturestoreEntityContainerEntityArgs struct {
	// [Required] Additional attributes of the entity.
	FeaturestoreEntityContainerProperties FeaturestoreEntityContainerInput
	// Container name. This is case-sensitive.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (FeaturestoreEntityContainerEntityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featurestoreEntityContainerEntityArgs)(nil)).Elem()
}

type FeaturestoreEntityContainerEntityInput interface {
	pulumi.Input

	ToFeaturestoreEntityContainerEntityOutput() FeaturestoreEntityContainerEntityOutput
	ToFeaturestoreEntityContainerEntityOutputWithContext(ctx context.Context) FeaturestoreEntityContainerEntityOutput
}

func (*FeaturestoreEntityContainerEntity) ElementType() reflect.Type {
	return reflect.TypeOf((**FeaturestoreEntityContainerEntity)(nil)).Elem()
}

func (i *FeaturestoreEntityContainerEntity) ToFeaturestoreEntityContainerEntityOutput() FeaturestoreEntityContainerEntityOutput {
	return i.ToFeaturestoreEntityContainerEntityOutputWithContext(context.Background())
}

func (i *FeaturestoreEntityContainerEntity) ToFeaturestoreEntityContainerEntityOutputWithContext(ctx context.Context) FeaturestoreEntityContainerEntityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeaturestoreEntityContainerEntityOutput)
}

func (i *FeaturestoreEntityContainerEntity) ToOutput(ctx context.Context) pulumix.Output[*FeaturestoreEntityContainerEntity] {
	return pulumix.Output[*FeaturestoreEntityContainerEntity]{
		OutputState: i.ToFeaturestoreEntityContainerEntityOutputWithContext(ctx).OutputState,
	}
}

type FeaturestoreEntityContainerEntityOutput struct{ *pulumi.OutputState }

func (FeaturestoreEntityContainerEntityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeaturestoreEntityContainerEntity)(nil)).Elem()
}

func (o FeaturestoreEntityContainerEntityOutput) ToFeaturestoreEntityContainerEntityOutput() FeaturestoreEntityContainerEntityOutput {
	return o
}

func (o FeaturestoreEntityContainerEntityOutput) ToFeaturestoreEntityContainerEntityOutputWithContext(ctx context.Context) FeaturestoreEntityContainerEntityOutput {
	return o
}

func (o FeaturestoreEntityContainerEntityOutput) ToOutput(ctx context.Context) pulumix.Output[*FeaturestoreEntityContainerEntity] {
	return pulumix.Output[*FeaturestoreEntityContainerEntity]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o FeaturestoreEntityContainerEntityOutput) FeaturestoreEntityContainerProperties() FeaturestoreEntityContainerResponseOutput {
	return o.ApplyT(func(v *FeaturestoreEntityContainerEntity) FeaturestoreEntityContainerResponseOutput {
		return v.FeaturestoreEntityContainerProperties
	}).(FeaturestoreEntityContainerResponseOutput)
}

// The name of the resource
func (o FeaturestoreEntityContainerEntityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FeaturestoreEntityContainerEntity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FeaturestoreEntityContainerEntityOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FeaturestoreEntityContainerEntity) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FeaturestoreEntityContainerEntityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FeaturestoreEntityContainerEntity) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FeaturestoreEntityContainerEntityOutput{})
}
