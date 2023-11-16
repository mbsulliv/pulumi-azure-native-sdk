// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified network interface.
func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("azure-native:network/v20190201:getNetworkInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkInterfaceArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the network interface.
	NetworkInterfaceName string `pulumi:"networkInterfaceName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A network interface in a resource group.
type LookupNetworkInterfaceResult struct {
	// The DNS settings in network interface.
	DnsSettings *NetworkInterfaceDnsSettingsResponse `pulumi:"dnsSettings"`
	// If the network interface is accelerated networking enabled.
	EnableAcceleratedNetworking *bool `pulumi:"enableAcceleratedNetworking"`
	// Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding *bool `pulumi:"enableIPForwarding"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// A list of references to linked BareMetal resources
	HostedWorkloads []string `pulumi:"hostedWorkloads"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// A reference to the interface endpoint to which the network interface is linked.
	InterfaceEndpoint InterfaceEndpointResponse `pulumi:"interfaceEndpoint"`
	// A list of IPConfigurations of the network interface.
	IpConfigurations []NetworkInterfaceIPConfigurationResponse `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The MAC address of the network interface.
	MacAddress *string `pulumi:"macAddress"`
	// Resource name.
	Name string `pulumi:"name"`
	// The reference of the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroupResponse `pulumi:"networkSecurityGroup"`
	// Gets whether this is a primary network interface on a virtual machine.
	Primary *bool `pulumi:"primary"`
	// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource GUID property of the network interface resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A list of TapConfigurations of the network interface.
	TapConfigurations []NetworkInterfaceTapConfigurationResponse `pulumi:"tapConfigurations"`
	// Resource type.
	Type string `pulumi:"type"`
	// The reference of a virtual machine.
	VirtualMachine SubResourceResponse `pulumi:"virtualMachine"`
}

func LookupNetworkInterfaceOutput(ctx *pulumi.Context, args LookupNetworkInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkInterfaceResult, error) {
			args := v.(LookupNetworkInterfaceArgs)
			r, err := LookupNetworkInterface(ctx, &args, opts...)
			var s LookupNetworkInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkInterfaceResultOutput)
}

type LookupNetworkInterfaceOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the network interface.
	NetworkInterfaceName pulumi.StringInput `pulumi:"networkInterfaceName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceArgs)(nil)).Elem()
}

// A network interface in a resource group.
type LookupNetworkInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceResult)(nil)).Elem()
}

func (o LookupNetworkInterfaceResultOutput) ToLookupNetworkInterfaceResultOutput() LookupNetworkInterfaceResultOutput {
	return o
}

func (o LookupNetworkInterfaceResultOutput) ToLookupNetworkInterfaceResultOutputWithContext(ctx context.Context) LookupNetworkInterfaceResultOutput {
	return o
}

// The DNS settings in network interface.
func (o LookupNetworkInterfaceResultOutput) DnsSettings() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *NetworkInterfaceDnsSettingsResponse { return v.DnsSettings }).(NetworkInterfaceDnsSettingsResponsePtrOutput)
}

// If the network interface is accelerated networking enabled.
func (o LookupNetworkInterfaceResultOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

// Indicates whether IP forwarding is enabled on this network interface.
func (o LookupNetworkInterfaceResultOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *bool { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupNetworkInterfaceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// A list of references to linked BareMetal resources
func (o LookupNetworkInterfaceResultOutput) HostedWorkloads() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) []string { return v.HostedWorkloads }).(pulumi.StringArrayOutput)
}

// Resource ID.
func (o LookupNetworkInterfaceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// A reference to the interface endpoint to which the network interface is linked.
func (o LookupNetworkInterfaceResultOutput) InterfaceEndpoint() InterfaceEndpointResponseOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) InterfaceEndpointResponse { return v.InterfaceEndpoint }).(InterfaceEndpointResponseOutput)
}

// A list of IPConfigurations of the network interface.
func (o LookupNetworkInterfaceResultOutput) IpConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) []NetworkInterfaceIPConfigurationResponse {
		return v.IpConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

// Resource location.
func (o LookupNetworkInterfaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The MAC address of the network interface.
func (o LookupNetworkInterfaceResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupNetworkInterfaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The reference of the NetworkSecurityGroup resource.
func (o LookupNetworkInterfaceResultOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *NetworkSecurityGroupResponse { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

// Gets whether this is a primary network interface on a virtual machine.
func (o LookupNetworkInterfaceResultOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
func (o LookupNetworkInterfaceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The resource GUID property of the network interface resource.
func (o LookupNetworkInterfaceResultOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupNetworkInterfaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A list of TapConfigurations of the network interface.
func (o LookupNetworkInterfaceResultOutput) TapConfigurations() NetworkInterfaceTapConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) []NetworkInterfaceTapConfigurationResponse {
		return v.TapConfigurations
	}).(NetworkInterfaceTapConfigurationResponseArrayOutput)
}

// Resource type.
func (o LookupNetworkInterfaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The reference of a virtual machine.
func (o LookupNetworkInterfaceResultOutput) VirtualMachine() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) SubResourceResponse { return v.VirtualMachine }).(SubResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkInterfaceResultOutput{})
}
