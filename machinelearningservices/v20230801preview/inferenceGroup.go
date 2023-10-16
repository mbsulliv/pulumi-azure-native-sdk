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

type InferenceGroup struct {
	pulumi.CustomResourceState

	// Managed service identity (system assigned and/or user assigned identities)
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// [Required] Additional attributes of the entity.
	InferenceGroupProperties InferenceGroupResponseOutput `pulumi:"inferenceGroupProperties"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Sku details required for ARM contract for Autoscaling.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewInferenceGroup registers a new resource with the given unique name, arguments, and options.
func NewInferenceGroup(ctx *pulumi.Context,
	name string, args *InferenceGroupArgs, opts ...pulumi.ResourceOption) (*InferenceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InferenceGroupProperties == nil {
		return nil, errors.New("invalid value for required argument 'InferenceGroupProperties'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.InferenceGroupProperties = args.InferenceGroupProperties.ToInferenceGroupTypeOutput().ApplyT(func(v InferenceGroupType) InferenceGroupType { return *v.Defaults() }).(InferenceGroupTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:InferenceGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource InferenceGroup
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230801preview:InferenceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInferenceGroup gets an existing InferenceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInferenceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InferenceGroupState, opts ...pulumi.ResourceOption) (*InferenceGroup, error) {
	var resource InferenceGroup
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230801preview:InferenceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InferenceGroup resources.
type inferenceGroupState struct {
}

type InferenceGroupState struct {
}

func (InferenceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*inferenceGroupState)(nil)).Elem()
}

type inferenceGroupArgs struct {
	// InferenceGroup name.
	GroupName *string `pulumi:"groupName"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// [Required] Additional attributes of the entity.
	InferenceGroupProperties InferenceGroupType `pulumi:"inferenceGroupProperties"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// InferencePool name.
	PoolName string `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sku details required for ARM contract for Autoscaling.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a InferenceGroup resource.
type InferenceGroupArgs struct {
	// InferenceGroup name.
	GroupName pulumi.StringPtrInput
	// Managed service identity (system assigned and/or user assigned identities)
	Identity ManagedServiceIdentityPtrInput
	// [Required] Additional attributes of the entity.
	InferenceGroupProperties InferenceGroupTypeInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// InferencePool name.
	PoolName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Sku details required for ARM contract for Autoscaling.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (InferenceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inferenceGroupArgs)(nil)).Elem()
}

type InferenceGroupInput interface {
	pulumi.Input

	ToInferenceGroupOutput() InferenceGroupOutput
	ToInferenceGroupOutputWithContext(ctx context.Context) InferenceGroupOutput
}

func (*InferenceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceGroup)(nil)).Elem()
}

func (i *InferenceGroup) ToInferenceGroupOutput() InferenceGroupOutput {
	return i.ToInferenceGroupOutputWithContext(context.Background())
}

func (i *InferenceGroup) ToInferenceGroupOutputWithContext(ctx context.Context) InferenceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceGroupOutput)
}

func (i *InferenceGroup) ToOutput(ctx context.Context) pulumix.Output[*InferenceGroup] {
	return pulumix.Output[*InferenceGroup]{
		OutputState: i.ToInferenceGroupOutputWithContext(ctx).OutputState,
	}
}

type InferenceGroupOutput struct{ *pulumi.OutputState }

func (InferenceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceGroup)(nil)).Elem()
}

func (o InferenceGroupOutput) ToInferenceGroupOutput() InferenceGroupOutput {
	return o
}

func (o InferenceGroupOutput) ToInferenceGroupOutputWithContext(ctx context.Context) InferenceGroupOutput {
	return o
}

func (o InferenceGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*InferenceGroup] {
	return pulumix.Output[*InferenceGroup]{
		OutputState: o.OutputState,
	}
}

// Managed service identity (system assigned and/or user assigned identities)
func (o InferenceGroupOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *InferenceGroup) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// [Required] Additional attributes of the entity.
func (o InferenceGroupOutput) InferenceGroupProperties() InferenceGroupResponseOutput {
	return o.ApplyT(func(v *InferenceGroup) InferenceGroupResponseOutput { return v.InferenceGroupProperties }).(InferenceGroupResponseOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
func (o InferenceGroupOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InferenceGroup) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o InferenceGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o InferenceGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Sku details required for ARM contract for Autoscaling.
func (o InferenceGroupOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *InferenceGroup) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o InferenceGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *InferenceGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o InferenceGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InferenceGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o InferenceGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InferenceGroupOutput{})
}