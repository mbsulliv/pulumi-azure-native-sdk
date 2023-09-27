// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Defines the NetworkInterface resource.
type NetworkInterface struct {
	pulumi.CustomResourceState

	// administrativeState of the network interface. Example: Enabled | Disabled.
	AdministrativeState pulumi.StringOutput `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// The arm resource id of the interface or compute server its connected to.
	ConnectedTo pulumi.StringOutput `pulumi:"connectedTo"`
	// The Interface Type. Example: Management/Data
	InterfaceType pulumi.StringOutput `pulumi:"interfaceType"`
	// ipv4Address.
	Ipv4Address pulumi.StringOutput `pulumi:"ipv4Address"`
	// ipv6Address.
	Ipv6Address pulumi.StringOutput `pulumi:"ipv6Address"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// physicalIdentifier of the network interface.
	PhysicalIdentifier pulumi.StringOutput `pulumi:"physicalIdentifier"`
	// Gets the provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterface(ctx *pulumi.Context,
	name string, args *NetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkDeviceName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkDeviceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:managednetworkfabric/v20230615:NetworkInterface"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NetworkInterface
	err := ctx.RegisterResource("azure-native:managednetworkfabric/v20230201preview:NetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterface gets an existing NetworkInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceState, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	var resource NetworkInterface
	err := ctx.ReadResource("azure-native:managednetworkfabric/v20230201preview:NetworkInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterface resources.
type networkInterfaceState struct {
}

type NetworkInterfaceState struct {
}

func (NetworkInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceState)(nil)).Elem()
}

type networkInterfaceArgs struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Name of the NetworkDevice
	NetworkDeviceName string `pulumi:"networkDeviceName"`
	// Name of the NetworkInterface
	NetworkInterfaceName *string `pulumi:"networkInterfaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a NetworkInterface resource.
type NetworkInterfaceArgs struct {
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// Name of the NetworkDevice
	NetworkDeviceName pulumi.StringInput
	// Name of the NetworkInterface
	NetworkInterfaceName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceArgs)(nil)).Elem()
}

type NetworkInterfaceInput interface {
	pulumi.Input

	ToNetworkInterfaceOutput() NetworkInterfaceOutput
	ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput
}

func (*NetworkInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil)).Elem()
}

func (i *NetworkInterface) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return i.ToNetworkInterfaceOutputWithContext(context.Background())
}

func (i *NetworkInterface) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceOutput)
}

func (i *NetworkInterface) ToOutput(ctx context.Context) pulumix.Output[*NetworkInterface] {
	return pulumix.Output[*NetworkInterface]{
		OutputState: i.ToNetworkInterfaceOutputWithContext(ctx).OutputState,
	}
}

type NetworkInterfaceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToOutput(ctx context.Context) pulumix.Output[*NetworkInterface] {
	return pulumix.Output[*NetworkInterface]{
		OutputState: o.OutputState,
	}
}

// administrativeState of the network interface. Example: Enabled | Disabled.
func (o NetworkInterfaceOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o NetworkInterfaceOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// The arm resource id of the interface or compute server its connected to.
func (o NetworkInterfaceOutput) ConnectedTo() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.ConnectedTo }).(pulumi.StringOutput)
}

// The Interface Type. Example: Management/Data
func (o NetworkInterfaceOutput) InterfaceType() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.InterfaceType }).(pulumi.StringOutput)
}

// ipv4Address.
func (o NetworkInterfaceOutput) Ipv4Address() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Ipv4Address }).(pulumi.StringOutput)
}

// ipv6Address.
func (o NetworkInterfaceOutput) Ipv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Ipv6Address }).(pulumi.StringOutput)
}

// The name of the resource
func (o NetworkInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// physicalIdentifier of the network interface.
func (o NetworkInterfaceOutput) PhysicalIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.PhysicalIdentifier }).(pulumi.StringOutput)
}

// Gets the provisioning state of the resource.
func (o NetworkInterfaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o NetworkInterfaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkInterface) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NetworkInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
}
