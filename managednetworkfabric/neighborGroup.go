// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managednetworkfabric

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Defines the Neighbor Group.
// Azure REST API version: 2023-06-15.
type NeighborGroup struct {
	pulumi.CustomResourceState

	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// An array of destination IPv4 Addresses or IPv6 Addresses.
	Destination NeighborGroupDestinationResponseOutput `pulumi:"destination"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of NetworkTap IDs where neighbor group is associated.
	NetworkTapIds pulumi.StringArrayOutput `pulumi:"networkTapIds"`
	// List of Network Tap Rule IDs where neighbor group is associated.
	NetworkTapRuleIds pulumi.StringArrayOutput `pulumi:"networkTapRuleIds"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNeighborGroup registers a new resource with the given unique name, arguments, and options.
func NewNeighborGroup(ctx *pulumi.Context,
	name string, args *NeighborGroupArgs, opts ...pulumi.ResourceOption) (*NeighborGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric/v20230615:NeighborGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NeighborGroup
	err := ctx.RegisterResource("azure-native:managednetworkfabric:NeighborGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNeighborGroup gets an existing NeighborGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNeighborGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NeighborGroupState, opts ...pulumi.ResourceOption) (*NeighborGroup, error) {
	var resource NeighborGroup
	err := ctx.ReadResource("azure-native:managednetworkfabric:NeighborGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NeighborGroup resources.
type neighborGroupState struct {
}

type NeighborGroupState struct {
}

func (NeighborGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*neighborGroupState)(nil)).Elem()
}

type neighborGroupArgs struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// An array of destination IPv4 Addresses or IPv6 Addresses.
	Destination NeighborGroupDestination `pulumi:"destination"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of the Neighbor Group.
	NeighborGroupName *string `pulumi:"neighborGroupName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NeighborGroup resource.
type NeighborGroupArgs struct {
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// An array of destination IPv4 Addresses or IPv6 Addresses.
	Destination NeighborGroupDestinationInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of the Neighbor Group.
	NeighborGroupName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NeighborGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*neighborGroupArgs)(nil)).Elem()
}

type NeighborGroupInput interface {
	pulumi.Input

	ToNeighborGroupOutput() NeighborGroupOutput
	ToNeighborGroupOutputWithContext(ctx context.Context) NeighborGroupOutput
}

func (*NeighborGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NeighborGroup)(nil)).Elem()
}

func (i *NeighborGroup) ToNeighborGroupOutput() NeighborGroupOutput {
	return i.ToNeighborGroupOutputWithContext(context.Background())
}

func (i *NeighborGroup) ToNeighborGroupOutputWithContext(ctx context.Context) NeighborGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NeighborGroupOutput)
}

func (i *NeighborGroup) ToOutput(ctx context.Context) pulumix.Output[*NeighborGroup] {
	return pulumix.Output[*NeighborGroup]{
		OutputState: i.ToNeighborGroupOutputWithContext(ctx).OutputState,
	}
}

type NeighborGroupOutput struct{ *pulumi.OutputState }

func (NeighborGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NeighborGroup)(nil)).Elem()
}

func (o NeighborGroupOutput) ToNeighborGroupOutput() NeighborGroupOutput {
	return o
}

func (o NeighborGroupOutput) ToNeighborGroupOutputWithContext(ctx context.Context) NeighborGroupOutput {
	return o
}

func (o NeighborGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*NeighborGroup] {
	return pulumix.Output[*NeighborGroup]{
		OutputState: o.OutputState,
	}
}

// Switch configuration description.
func (o NeighborGroupOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NeighborGroup) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// An array of destination IPv4 Addresses or IPv6 Addresses.
func (o NeighborGroupOutput) Destination() NeighborGroupDestinationResponseOutput {
	return o.ApplyT(func(v *NeighborGroup) NeighborGroupDestinationResponseOutput { return v.Destination }).(NeighborGroupDestinationResponseOutput)
}

// The geo-location where the resource lives
func (o NeighborGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NeighborGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o NeighborGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NeighborGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of NetworkTap IDs where neighbor group is associated.
func (o NeighborGroupOutput) NetworkTapIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NeighborGroup) pulumi.StringArrayOutput { return v.NetworkTapIds }).(pulumi.StringArrayOutput)
}

// List of Network Tap Rule IDs where neighbor group is associated.
func (o NeighborGroupOutput) NetworkTapRuleIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NeighborGroup) pulumi.StringArrayOutput { return v.NetworkTapRuleIds }).(pulumi.StringArrayOutput)
}

// The provisioning state of the resource.
func (o NeighborGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NeighborGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o NeighborGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NeighborGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o NeighborGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NeighborGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NeighborGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NeighborGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NeighborGroupOutput{})
}
