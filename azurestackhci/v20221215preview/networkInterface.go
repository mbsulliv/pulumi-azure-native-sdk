// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221215preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The network interface resource definition.
type NetworkInterface struct {
	pulumi.CustomResourceState

	// DNS Settings for the interface
	DnsSettings InterfaceDNSSettingsResponsePtrOutput `pulumi:"dnsSettings"`
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// IPConfigurations - A list of IPConfigurations of the network interface.
	IpConfigurations IPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// MacAddress - The MAC address of the network interface.
	MacAddress pulumi.StringPtrOutput `pulumi:"macAddress"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the network interface.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The observed state of network interfaces
	Status NetworkInterfaceStatusResponseOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
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
			Type: pulumi.String("azure-native:azurestackhci:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210701preview:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230701preview:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230901preview:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240101:NetworkInterface"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NetworkInterface
	err := ctx.RegisterResource("azure-native:azurestackhci/v20221215preview:NetworkInterface", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:azurestackhci/v20221215preview:NetworkInterface", name, id, state, &resource, opts...)
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
	// DNS Settings for the interface
	DnsSettings *InterfaceDNSSettings `pulumi:"dnsSettings"`
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// IPConfigurations - A list of IPConfigurations of the network interface.
	IpConfigurations []IPConfiguration `pulumi:"ipConfigurations"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// MacAddress - The MAC address of the network interface.
	MacAddress *string `pulumi:"macAddress"`
	// Name of the network interface
	NetworkInterfaceName *string `pulumi:"networkInterfaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkInterface resource.
type NetworkInterfaceArgs struct {
	// DNS Settings for the interface
	DnsSettings InterfaceDNSSettingsPtrInput
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationPtrInput
	// IPConfigurations - A list of IPConfigurations of the network interface.
	IpConfigurations IPConfigurationArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// MacAddress - The MAC address of the network interface.
	MacAddress pulumi.StringPtrInput
	// Name of the network interface
	NetworkInterfaceName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
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

// DNS Settings for the interface
func (o NetworkInterfaceOutput) DnsSettings() InterfaceDNSSettingsResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) InterfaceDNSSettingsResponsePtrOutput { return v.DnsSettings }).(InterfaceDNSSettingsResponsePtrOutput)
}

// The extendedLocation of the resource.
func (o NetworkInterfaceOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// IPConfigurations - A list of IPConfigurations of the network interface.
func (o NetworkInterfaceOutput) IpConfigurations() IPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) IPConfigurationResponseArrayOutput { return v.IpConfigurations }).(IPConfigurationResponseArrayOutput)
}

// The geo-location where the resource lives
func (o NetworkInterfaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// MacAddress - The MAC address of the network interface.
func (o NetworkInterfaceOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.MacAddress }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o NetworkInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the network interface.
func (o NetworkInterfaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The observed state of network interfaces
func (o NetworkInterfaceOutput) Status() NetworkInterfaceStatusResponseOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkInterfaceStatusResponseOutput { return v.Status }).(NetworkInterfaceStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o NetworkInterfaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkInterface) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o NetworkInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NetworkInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
}
