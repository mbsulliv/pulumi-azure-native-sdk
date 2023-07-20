// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230615

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Network Rack resource definition.
type NetworkRack struct {
	pulumi.CustomResourceState

	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of network device ARM resource IDs.
	NetworkDevices pulumi.StringArrayOutput `pulumi:"networkDevices"`
	// ARM resource ID of the Network Fabric.
	NetworkFabricId pulumi.StringOutput `pulumi:"networkFabricId"`
	// Network Rack SKU name.
	NetworkRackType pulumi.StringPtrOutput `pulumi:"networkRackType"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkRack registers a new resource with the given unique name, arguments, and options.
func NewNetworkRack(ctx *pulumi.Context,
	name string, args *NetworkRackArgs, opts ...pulumi.ResourceOption) (*NetworkRack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkFabricId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkFabricId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric:NetworkRack"),
		},
		{
			Type: pulumi.String("azure-native:managednetworkfabric/v20230201preview:NetworkRack"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkRack
	err := ctx.RegisterResource("azure-native:managednetworkfabric/v20230615:NetworkRack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkRack gets an existing NetworkRack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkRack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkRackState, opts ...pulumi.ResourceOption) (*NetworkRack, error) {
	var resource NetworkRack
	err := ctx.ReadResource("azure-native:managednetworkfabric/v20230615:NetworkRack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkRack resources.
type networkRackState struct {
}

type NetworkRackState struct {
}

func (NetworkRackState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkRackState)(nil)).Elem()
}

type networkRackArgs struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// ARM resource ID of the Network Fabric.
	NetworkFabricId string `pulumi:"networkFabricId"`
	// Name of the Network Rack.
	NetworkRackName *string `pulumi:"networkRackName"`
	// Network Rack SKU name.
	NetworkRackType *string `pulumi:"networkRackType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkRack resource.
type NetworkRackArgs struct {
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// ARM resource ID of the Network Fabric.
	NetworkFabricId pulumi.StringInput
	// Name of the Network Rack.
	NetworkRackName pulumi.StringPtrInput
	// Network Rack SKU name.
	NetworkRackType pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkRackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkRackArgs)(nil)).Elem()
}

type NetworkRackInput interface {
	pulumi.Input

	ToNetworkRackOutput() NetworkRackOutput
	ToNetworkRackOutputWithContext(ctx context.Context) NetworkRackOutput
}

func (*NetworkRack) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRack)(nil)).Elem()
}

func (i *NetworkRack) ToNetworkRackOutput() NetworkRackOutput {
	return i.ToNetworkRackOutputWithContext(context.Background())
}

func (i *NetworkRack) ToNetworkRackOutputWithContext(ctx context.Context) NetworkRackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRackOutput)
}

type NetworkRackOutput struct{ *pulumi.OutputState }

func (NetworkRackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRack)(nil)).Elem()
}

func (o NetworkRackOutput) ToNetworkRackOutput() NetworkRackOutput {
	return o
}

func (o NetworkRackOutput) ToNetworkRackOutputWithContext(ctx context.Context) NetworkRackOutput {
	return o
}

// Switch configuration description.
func (o NetworkRackOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRack) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o NetworkRackOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkRack) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o NetworkRackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkRack) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of network device ARM resource IDs.
func (o NetworkRackOutput) NetworkDevices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkRack) pulumi.StringArrayOutput { return v.NetworkDevices }).(pulumi.StringArrayOutput)
}

// ARM resource ID of the Network Fabric.
func (o NetworkRackOutput) NetworkFabricId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkRack) pulumi.StringOutput { return v.NetworkFabricId }).(pulumi.StringOutput)
}

// Network Rack SKU name.
func (o NetworkRackOutput) NetworkRackType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRack) pulumi.StringPtrOutput { return v.NetworkRackType }).(pulumi.StringPtrOutput)
}

// Provisioning state of the resource.
func (o NetworkRackOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkRack) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o NetworkRackOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkRack) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o NetworkRackOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkRack) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NetworkRackOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkRack) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkRackOutput{})
}
