// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A network interface in a resource group.
type NetworkInterface struct {
	pulumi.CustomResourceState

	// The DNS settings in network interface.
	DnsSettings NetworkInterfaceDnsSettingsResponsePtrOutput `pulumi:"dnsSettings"`
	// If the network interface is accelerated networking enabled.
	EnableAcceleratedNetworking pulumi.BoolPtrOutput `pulumi:"enableAcceleratedNetworking"`
	// Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding pulumi.BoolPtrOutput `pulumi:"enableIPForwarding"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// A list of IPConfigurations of the network interface.
	IpConfigurations NetworkInterfaceIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The MAC address of the network interface.
	MacAddress pulumi.StringPtrOutput `pulumi:"macAddress"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The reference of the NetworkSecurityGroup resource.
	NetworkSecurityGroup NetworkSecurityGroupResponsePtrOutput `pulumi:"networkSecurityGroup"`
	// Gets whether this is a primary network interface on a virtual machine.
	Primary pulumi.BoolPtrOutput `pulumi:"primary"`
	// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The resource GUID property of the network interface resource.
	ResourceGuid pulumi.StringPtrOutput `pulumi:"resourceGuid"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The reference of a virtual machine.
	VirtualMachine SubResourceResponsePtrOutput `pulumi:"virtualMachine"`
}

// NewNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterface(ctx *pulumi.Context,
	name string, args *NetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:NetworkInterface"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NetworkInterface
	err := ctx.RegisterResource("azure-native:network/v20180701:NetworkInterface", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:network/v20180701:NetworkInterface", name, id, state, &resource, opts...)
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
	// The DNS settings in network interface.
	DnsSettings *NetworkInterfaceDnsSettings `pulumi:"dnsSettings"`
	// If the network interface is accelerated networking enabled.
	EnableAcceleratedNetworking *bool `pulumi:"enableAcceleratedNetworking"`
	// Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding *bool `pulumi:"enableIPForwarding"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// A list of IPConfigurations of the network interface.
	IpConfigurations []NetworkInterfaceIPConfiguration `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The MAC address of the network interface.
	MacAddress *string `pulumi:"macAddress"`
	// The name of the network interface.
	NetworkInterfaceName *string `pulumi:"networkInterfaceName"`
	// The reference of the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroup `pulumi:"networkSecurityGroup"`
	// Gets whether this is a primary network interface on a virtual machine.
	Primary *bool `pulumi:"primary"`
	// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource GUID property of the network interface resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The reference of a virtual machine.
	VirtualMachine *SubResource `pulumi:"virtualMachine"`
}

// The set of arguments for constructing a NetworkInterface resource.
type NetworkInterfaceArgs struct {
	// The DNS settings in network interface.
	DnsSettings NetworkInterfaceDnsSettingsPtrInput
	// If the network interface is accelerated networking enabled.
	EnableAcceleratedNetworking pulumi.BoolPtrInput
	// Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// A list of IPConfigurations of the network interface.
	IpConfigurations NetworkInterfaceIPConfigurationArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The MAC address of the network interface.
	MacAddress pulumi.StringPtrInput
	// The name of the network interface.
	NetworkInterfaceName pulumi.StringPtrInput
	// The reference of the NetworkSecurityGroup resource.
	NetworkSecurityGroup NetworkSecurityGroupPtrInput
	// Gets whether this is a primary network interface on a virtual machine.
	Primary pulumi.BoolPtrInput
	// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The resource GUID property of the network interface resource.
	ResourceGuid pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The reference of a virtual machine.
	VirtualMachine SubResourcePtrInput
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

// The DNS settings in network interface.
func (o NetworkInterfaceOutput) DnsSettings() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkInterfaceDnsSettingsResponsePtrOutput { return v.DnsSettings }).(NetworkInterfaceDnsSettingsResponsePtrOutput)
}

// If the network interface is accelerated networking enabled.
func (o NetworkInterfaceOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolPtrOutput { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

// Indicates whether IP forwarding is enabled on this network interface.
func (o NetworkInterfaceOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolPtrOutput { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o NetworkInterfaceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// A list of IPConfigurations of the network interface.
func (o NetworkInterfaceOutput) IpConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkInterfaceIPConfigurationResponseArrayOutput {
		return v.IpConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

// Resource location.
func (o NetworkInterfaceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The MAC address of the network interface.
func (o NetworkInterfaceOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.MacAddress }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o NetworkInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The reference of the NetworkSecurityGroup resource.
func (o NetworkInterfaceOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkSecurityGroupResponsePtrOutput { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

// Gets whether this is a primary network interface on a virtual machine.
func (o NetworkInterfaceOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolPtrOutput { return v.Primary }).(pulumi.BoolPtrOutput)
}

// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
func (o NetworkInterfaceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The resource GUID property of the network interface resource.
func (o NetworkInterfaceOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o NetworkInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o NetworkInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The reference of a virtual machine.
func (o NetworkInterfaceOutput) VirtualMachine() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) SubResourceResponsePtrOutput { return v.VirtualMachine }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
}
