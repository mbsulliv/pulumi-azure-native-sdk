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

type CapacityReservationGroup struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	CapacityReservationGroupProperties CapacityReservationGroupResponseOutput `pulumi:"capacityReservationGroupProperties"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
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

// NewCapacityReservationGroup registers a new resource with the given unique name, arguments, and options.
func NewCapacityReservationGroup(ctx *pulumi.Context,
	name string, args *CapacityReservationGroupArgs, opts ...pulumi.ResourceOption) (*CapacityReservationGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CapacityReservationGroupProperties == nil {
		return nil, errors.New("invalid value for required argument 'CapacityReservationGroupProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:CapacityReservationGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CapacityReservationGroup
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230801preview:CapacityReservationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapacityReservationGroup gets an existing CapacityReservationGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapacityReservationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityReservationGroupState, opts ...pulumi.ResourceOption) (*CapacityReservationGroup, error) {
	var resource CapacityReservationGroup
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230801preview:CapacityReservationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CapacityReservationGroup resources.
type capacityReservationGroupState struct {
}

type CapacityReservationGroupState struct {
}

func (CapacityReservationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationGroupState)(nil)).Elem()
}

type capacityReservationGroupArgs struct {
	// [Required] Additional attributes of the entity.
	CapacityReservationGroupProperties CapacityReservationGroupType `pulumi:"capacityReservationGroupProperties"`
	GroupId                            *string                      `pulumi:"groupId"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sku details required for ARM contract for Autoscaling.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CapacityReservationGroup resource.
type CapacityReservationGroupArgs struct {
	// [Required] Additional attributes of the entity.
	CapacityReservationGroupProperties CapacityReservationGroupTypeInput
	GroupId                            pulumi.StringPtrInput
	// Managed service identity (system assigned and/or user assigned identities)
	Identity ManagedServiceIdentityPtrInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Sku details required for ARM contract for Autoscaling.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (CapacityReservationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationGroupArgs)(nil)).Elem()
}

type CapacityReservationGroupInput interface {
	pulumi.Input

	ToCapacityReservationGroupOutput() CapacityReservationGroupOutput
	ToCapacityReservationGroupOutputWithContext(ctx context.Context) CapacityReservationGroupOutput
}

func (*CapacityReservationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservationGroup)(nil)).Elem()
}

func (i *CapacityReservationGroup) ToCapacityReservationGroupOutput() CapacityReservationGroupOutput {
	return i.ToCapacityReservationGroupOutputWithContext(context.Background())
}

func (i *CapacityReservationGroup) ToCapacityReservationGroupOutputWithContext(ctx context.Context) CapacityReservationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationGroupOutput)
}

func (i *CapacityReservationGroup) ToOutput(ctx context.Context) pulumix.Output[*CapacityReservationGroup] {
	return pulumix.Output[*CapacityReservationGroup]{
		OutputState: i.ToCapacityReservationGroupOutputWithContext(ctx).OutputState,
	}
}

type CapacityReservationGroupOutput struct{ *pulumi.OutputState }

func (CapacityReservationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservationGroup)(nil)).Elem()
}

func (o CapacityReservationGroupOutput) ToCapacityReservationGroupOutput() CapacityReservationGroupOutput {
	return o
}

func (o CapacityReservationGroupOutput) ToCapacityReservationGroupOutputWithContext(ctx context.Context) CapacityReservationGroupOutput {
	return o
}

func (o CapacityReservationGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*CapacityReservationGroup] {
	return pulumix.Output[*CapacityReservationGroup]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o CapacityReservationGroupOutput) CapacityReservationGroupProperties() CapacityReservationGroupResponseOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) CapacityReservationGroupResponseOutput {
		return v.CapacityReservationGroupProperties
	}).(CapacityReservationGroupResponseOutput)
}

// Managed service identity (system assigned and/or user assigned identities)
func (o CapacityReservationGroupOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
func (o CapacityReservationGroupOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o CapacityReservationGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o CapacityReservationGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Sku details required for ARM contract for Autoscaling.
func (o CapacityReservationGroupOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CapacityReservationGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o CapacityReservationGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CapacityReservationGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CapacityReservationGroupOutput{})
}
