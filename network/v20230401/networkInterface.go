// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A network interface in a resource group.
type NetworkInterface struct {
	pulumi.CustomResourceState

	// Auxiliary mode of Network Interface resource.
	AuxiliaryMode pulumi.StringPtrOutput `pulumi:"auxiliaryMode"`
	// Auxiliary sku of Network Interface resource.
	AuxiliarySku pulumi.StringPtrOutput `pulumi:"auxiliarySku"`
	// Indicates whether to disable tcp state tracking.
	DisableTcpStateTracking pulumi.BoolPtrOutput `pulumi:"disableTcpStateTracking"`
	// The DNS settings in network interface.
	DnsSettings NetworkInterfaceDnsSettingsResponsePtrOutput `pulumi:"dnsSettings"`
	// A reference to the dscp configuration to which the network interface is linked.
	DscpConfiguration SubResourceResponseOutput `pulumi:"dscpConfiguration"`
	// If the network interface is configured for accelerated networking. Not applicable to VM sizes which require accelerated networking.
	EnableAcceleratedNetworking pulumi.BoolPtrOutput `pulumi:"enableAcceleratedNetworking"`
	// Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding pulumi.BoolPtrOutput `pulumi:"enableIPForwarding"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The extended location of the network interface.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// A list of references to linked BareMetal resources.
	HostedWorkloads pulumi.StringArrayOutput `pulumi:"hostedWorkloads"`
	// A list of IPConfigurations of the network interface.
	IpConfigurations NetworkInterfaceIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The MAC address of the network interface.
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// Migration phase of Network Interface resource.
	MigrationPhase pulumi.StringPtrOutput `pulumi:"migrationPhase"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup NetworkSecurityGroupResponsePtrOutput `pulumi:"networkSecurityGroup"`
	// Type of Network Interface resource.
	NicType pulumi.StringPtrOutput `pulumi:"nicType"`
	// Whether this is a primary network interface on a virtual machine.
	Primary pulumi.BoolOutput `pulumi:"primary"`
	// A reference to the private endpoint to which the network interface is linked.
	PrivateEndpoint PrivateEndpointResponseOutput `pulumi:"privateEndpoint"`
	// Privatelinkservice of the network interface resource.
	PrivateLinkService PrivateLinkServiceResponsePtrOutput `pulumi:"privateLinkService"`
	// The provisioning state of the network interface resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource GUID property of the network interface resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A list of TapConfigurations of the network interface.
	TapConfigurations NetworkInterfaceTapConfigurationResponseArrayOutput `pulumi:"tapConfigurations"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The reference to a virtual machine.
	VirtualMachine SubResourceResponseOutput `pulumi:"virtualMachine"`
	// Whether the virtual machine this nic is attached to supports encryption.
	VnetEncryptionSupported pulumi.BoolOutput `pulumi:"vnetEncryptionSupported"`
	// WorkloadType of the NetworkInterface for BareMetal resources
	WorkloadType pulumi.StringPtrOutput `pulumi:"workloadType"`
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
			Type: pulumi.String("azure-native:network/v20180701:NetworkInterface"),
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
			Type: pulumi.String("azure-native:network/v20230501:NetworkInterface"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NetworkInterface
	err := ctx.RegisterResource("azure-native:network/v20230401:NetworkInterface", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:network/v20230401:NetworkInterface", name, id, state, &resource, opts...)
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
	// Auxiliary mode of Network Interface resource.
	AuxiliaryMode *string `pulumi:"auxiliaryMode"`
	// Auxiliary sku of Network Interface resource.
	AuxiliarySku *string `pulumi:"auxiliarySku"`
	// Indicates whether to disable tcp state tracking.
	DisableTcpStateTracking *bool `pulumi:"disableTcpStateTracking"`
	// The DNS settings in network interface.
	DnsSettings *NetworkInterfaceDnsSettings `pulumi:"dnsSettings"`
	// If the network interface is configured for accelerated networking. Not applicable to VM sizes which require accelerated networking.
	EnableAcceleratedNetworking *bool `pulumi:"enableAcceleratedNetworking"`
	// Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding *bool `pulumi:"enableIPForwarding"`
	// The extended location of the network interface.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// A list of IPConfigurations of the network interface.
	IpConfigurations []NetworkInterfaceIPConfiguration `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Migration phase of Network Interface resource.
	MigrationPhase *string `pulumi:"migrationPhase"`
	// The name of the network interface.
	NetworkInterfaceName *string `pulumi:"networkInterfaceName"`
	// The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroupType `pulumi:"networkSecurityGroup"`
	// Type of Network Interface resource.
	NicType *string `pulumi:"nicType"`
	// Privatelinkservice of the network interface resource.
	PrivateLinkService *PrivateLinkServiceType `pulumi:"privateLinkService"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// WorkloadType of the NetworkInterface for BareMetal resources
	WorkloadType *string `pulumi:"workloadType"`
}

// The set of arguments for constructing a NetworkInterface resource.
type NetworkInterfaceArgs struct {
	// Auxiliary mode of Network Interface resource.
	AuxiliaryMode pulumi.StringPtrInput
	// Auxiliary sku of Network Interface resource.
	AuxiliarySku pulumi.StringPtrInput
	// Indicates whether to disable tcp state tracking.
	DisableTcpStateTracking pulumi.BoolPtrInput
	// The DNS settings in network interface.
	DnsSettings NetworkInterfaceDnsSettingsPtrInput
	// If the network interface is configured for accelerated networking. Not applicable to VM sizes which require accelerated networking.
	EnableAcceleratedNetworking pulumi.BoolPtrInput
	// Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding pulumi.BoolPtrInput
	// The extended location of the network interface.
	ExtendedLocation ExtendedLocationPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// A list of IPConfigurations of the network interface.
	IpConfigurations NetworkInterfaceIPConfigurationArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Migration phase of Network Interface resource.
	MigrationPhase pulumi.StringPtrInput
	// The name of the network interface.
	NetworkInterfaceName pulumi.StringPtrInput
	// The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup NetworkSecurityGroupTypePtrInput
	// Type of Network Interface resource.
	NicType pulumi.StringPtrInput
	// Privatelinkservice of the network interface resource.
	PrivateLinkService PrivateLinkServiceTypePtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// WorkloadType of the NetworkInterface for BareMetal resources
	WorkloadType pulumi.StringPtrInput
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

// Auxiliary mode of Network Interface resource.
func (o NetworkInterfaceOutput) AuxiliaryMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.AuxiliaryMode }).(pulumi.StringPtrOutput)
}

// Auxiliary sku of Network Interface resource.
func (o NetworkInterfaceOutput) AuxiliarySku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.AuxiliarySku }).(pulumi.StringPtrOutput)
}

// Indicates whether to disable tcp state tracking.
func (o NetworkInterfaceOutput) DisableTcpStateTracking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolPtrOutput { return v.DisableTcpStateTracking }).(pulumi.BoolPtrOutput)
}

// The DNS settings in network interface.
func (o NetworkInterfaceOutput) DnsSettings() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkInterfaceDnsSettingsResponsePtrOutput { return v.DnsSettings }).(NetworkInterfaceDnsSettingsResponsePtrOutput)
}

// A reference to the dscp configuration to which the network interface is linked.
func (o NetworkInterfaceOutput) DscpConfiguration() SubResourceResponseOutput {
	return o.ApplyT(func(v *NetworkInterface) SubResourceResponseOutput { return v.DscpConfiguration }).(SubResourceResponseOutput)
}

// If the network interface is configured for accelerated networking. Not applicable to VM sizes which require accelerated networking.
func (o NetworkInterfaceOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolPtrOutput { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

// Indicates whether IP forwarding is enabled on this network interface.
func (o NetworkInterfaceOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolPtrOutput { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o NetworkInterfaceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the network interface.
func (o NetworkInterfaceOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// A list of references to linked BareMetal resources.
func (o NetworkInterfaceOutput) HostedWorkloads() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringArrayOutput { return v.HostedWorkloads }).(pulumi.StringArrayOutput)
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
func (o NetworkInterfaceOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

// Migration phase of Network Interface resource.
func (o NetworkInterfaceOutput) MigrationPhase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.MigrationPhase }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o NetworkInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The reference to the NetworkSecurityGroup resource.
func (o NetworkInterfaceOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkSecurityGroupResponsePtrOutput { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

// Type of Network Interface resource.
func (o NetworkInterfaceOutput) NicType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.NicType }).(pulumi.StringPtrOutput)
}

// Whether this is a primary network interface on a virtual machine.
func (o NetworkInterfaceOutput) Primary() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolOutput { return v.Primary }).(pulumi.BoolOutput)
}

// A reference to the private endpoint to which the network interface is linked.
func (o NetworkInterfaceOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *NetworkInterface) PrivateEndpointResponseOutput { return v.PrivateEndpoint }).(PrivateEndpointResponseOutput)
}

// Privatelinkservice of the network interface resource.
func (o NetworkInterfaceOutput) PrivateLinkService() PrivateLinkServiceResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) PrivateLinkServiceResponsePtrOutput { return v.PrivateLinkService }).(PrivateLinkServiceResponsePtrOutput)
}

// The provisioning state of the network interface resource.
func (o NetworkInterfaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the network interface resource.
func (o NetworkInterfaceOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Resource tags.
func (o NetworkInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A list of TapConfigurations of the network interface.
func (o NetworkInterfaceOutput) TapConfigurations() NetworkInterfaceTapConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkInterfaceTapConfigurationResponseArrayOutput {
		return v.TapConfigurations
	}).(NetworkInterfaceTapConfigurationResponseArrayOutput)
}

// Resource type.
func (o NetworkInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The reference to a virtual machine.
func (o NetworkInterfaceOutput) VirtualMachine() SubResourceResponseOutput {
	return o.ApplyT(func(v *NetworkInterface) SubResourceResponseOutput { return v.VirtualMachine }).(SubResourceResponseOutput)
}

// Whether the virtual machine this nic is attached to supports encryption.
func (o NetworkInterfaceOutput) VnetEncryptionSupported() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolOutput { return v.VnetEncryptionSupported }).(pulumi.BoolOutput)
}

// WorkloadType of the NetworkInterface for BareMetal resources
func (o NetworkInterfaceOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
}
