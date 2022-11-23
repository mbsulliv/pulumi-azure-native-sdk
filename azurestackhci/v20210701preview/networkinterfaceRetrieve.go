// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The networkinterfaces resource definition.
type NetworkinterfaceRetrieve struct {
	pulumi.CustomResourceState

	// DNS Settings for the interface
	DnsSettings InterfaceDNSSettingsResponsePtrOutput `pulumi:"dnsSettings"`
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// IPConfigurations - A list of IPConfigurations of the network interface.
	IpConfigurations IpConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// MacAddress - The MAC address of the network interface.
	MacAddress pulumi.StringPtrOutput `pulumi:"macAddress"`
	// The name of the resource
	Name              pulumi.StringOutput `pulumi:"name"`
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// name of the object to be used in moc
	ResourceName pulumi.StringPtrOutput `pulumi:"resourceName"`
	// NetworkInterfaceStatus defines the observed state of NetworkInterface
	Status NetworkInterfaceStatusResponseOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkinterfaceRetrieve registers a new resource with the given unique name, arguments, and options.
func NewNetworkinterfaceRetrieve(ctx *pulumi.Context,
	name string, args *NetworkinterfaceRetrieveArgs, opts ...pulumi.ResourceOption) (*NetworkinterfaceRetrieve, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:networkinterfaceRetrieve"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkinterfaceRetrieve
	err := ctx.RegisterResource("azure-native:azurestackhci/v20210701preview:networkinterfaceRetrieve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkinterfaceRetrieve gets an existing NetworkinterfaceRetrieve resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkinterfaceRetrieve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkinterfaceRetrieveState, opts ...pulumi.ResourceOption) (*NetworkinterfaceRetrieve, error) {
	var resource NetworkinterfaceRetrieve
	err := ctx.ReadResource("azure-native:azurestackhci/v20210701preview:networkinterfaceRetrieve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkinterfaceRetrieve resources.
type networkinterfaceRetrieveState struct {
}

type NetworkinterfaceRetrieveState struct {
}

func (NetworkinterfaceRetrieveState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkinterfaceRetrieveState)(nil)).Elem()
}

type networkinterfaceRetrieveArgs struct {
	// DNS Settings for the interface
	DnsSettings *InterfaceDNSSettings `pulumi:"dnsSettings"`
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// IPConfigurations - A list of IPConfigurations of the network interface.
	IpConfigurations []IpConfiguration `pulumi:"ipConfigurations"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// MacAddress - The MAC address of the network interface.
	MacAddress            *string `pulumi:"macAddress"`
	NetworkinterfacesName *string `pulumi:"networkinterfacesName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// name of the object to be used in moc
	ResourceName *string `pulumi:"resourceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkinterfaceRetrieve resource.
type NetworkinterfaceRetrieveArgs struct {
	// DNS Settings for the interface
	DnsSettings InterfaceDNSSettingsPtrInput
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationPtrInput
	// IPConfigurations - A list of IPConfigurations of the network interface.
	IpConfigurations IpConfigurationArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// MacAddress - The MAC address of the network interface.
	MacAddress            pulumi.StringPtrInput
	NetworkinterfacesName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// name of the object to be used in moc
	ResourceName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkinterfaceRetrieveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkinterfaceRetrieveArgs)(nil)).Elem()
}

type NetworkinterfaceRetrieveInput interface {
	pulumi.Input

	ToNetworkinterfaceRetrieveOutput() NetworkinterfaceRetrieveOutput
	ToNetworkinterfaceRetrieveOutputWithContext(ctx context.Context) NetworkinterfaceRetrieveOutput
}

func (*NetworkinterfaceRetrieve) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkinterfaceRetrieve)(nil)).Elem()
}

func (i *NetworkinterfaceRetrieve) ToNetworkinterfaceRetrieveOutput() NetworkinterfaceRetrieveOutput {
	return i.ToNetworkinterfaceRetrieveOutputWithContext(context.Background())
}

func (i *NetworkinterfaceRetrieve) ToNetworkinterfaceRetrieveOutputWithContext(ctx context.Context) NetworkinterfaceRetrieveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkinterfaceRetrieveOutput)
}

type NetworkinterfaceRetrieveOutput struct{ *pulumi.OutputState }

func (NetworkinterfaceRetrieveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkinterfaceRetrieve)(nil)).Elem()
}

func (o NetworkinterfaceRetrieveOutput) ToNetworkinterfaceRetrieveOutput() NetworkinterfaceRetrieveOutput {
	return o
}

func (o NetworkinterfaceRetrieveOutput) ToNetworkinterfaceRetrieveOutputWithContext(ctx context.Context) NetworkinterfaceRetrieveOutput {
	return o
}

// DNS Settings for the interface
func (o NetworkinterfaceRetrieveOutput) DnsSettings() InterfaceDNSSettingsResponsePtrOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) InterfaceDNSSettingsResponsePtrOutput { return v.DnsSettings }).(InterfaceDNSSettingsResponsePtrOutput)
}

// The extendedLocation of the resource.
func (o NetworkinterfaceRetrieveOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// IPConfigurations - A list of IPConfigurations of the network interface.
func (o NetworkinterfaceRetrieveOutput) IpConfigurations() IpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) IpConfigurationResponseArrayOutput { return v.IpConfigurations }).(IpConfigurationResponseArrayOutput)
}

// The geo-location where the resource lives
func (o NetworkinterfaceRetrieveOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// MacAddress - The MAC address of the network interface.
func (o NetworkinterfaceRetrieveOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringPtrOutput { return v.MacAddress }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o NetworkinterfaceRetrieveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkinterfaceRetrieveOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// name of the object to be used in moc
func (o NetworkinterfaceRetrieveOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringPtrOutput { return v.ResourceName }).(pulumi.StringPtrOutput)
}

// NetworkInterfaceStatus defines the observed state of NetworkInterface
func (o NetworkinterfaceRetrieveOutput) Status() NetworkInterfaceStatusResponseOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) NetworkInterfaceStatusResponseOutput { return v.Status }).(NetworkInterfaceStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o NetworkinterfaceRetrieveOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o NetworkinterfaceRetrieveOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NetworkinterfaceRetrieveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkinterfaceRetrieveOutput{})
}
