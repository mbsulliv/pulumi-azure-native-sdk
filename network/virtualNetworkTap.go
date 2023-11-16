// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Virtual Network Tap resource.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01.
type VirtualNetworkTap struct {
	pulumi.CustomResourceState

	// The reference to the private IP address on the internal Load Balancer that will receive the tap.
	DestinationLoadBalancerFrontEndIPConfiguration FrontendIPConfigurationResponsePtrOutput `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	// The reference to the private IP Address of the collector nic that will receive the tap.
	DestinationNetworkInterfaceIPConfiguration NetworkInterfaceIPConfigurationResponsePtrOutput `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	// The VXLAN destination port that will receive the tapped traffic.
	DestinationPort pulumi.IntPtrOutput `pulumi:"destinationPort"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the list of resource IDs for the network interface IP configuration that needs to be tapped.
	NetworkInterfaceTapConfigurations NetworkInterfaceTapConfigurationResponseArrayOutput `pulumi:"networkInterfaceTapConfigurations"`
	// The provisioning state of the virtual network tap resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource GUID property of the virtual network tap resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualNetworkTap registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkTap(ctx *pulumi.Context,
	name string, args *VirtualNetworkTapArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkTap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.DestinationLoadBalancerFrontEndIPConfiguration != nil {
		args.DestinationLoadBalancerFrontEndIPConfiguration = args.DestinationLoadBalancerFrontEndIPConfiguration.ToFrontendIPConfigurationPtrOutput().ApplyT(func(v *FrontendIPConfiguration) *FrontendIPConfiguration { return v.Defaults() }).(FrontendIPConfigurationPtrOutput)
	}
	if args.DestinationNetworkInterfaceIPConfiguration != nil {
		args.DestinationNetworkInterfaceIPConfiguration = args.DestinationNetworkInterfaceIPConfiguration.ToNetworkInterfaceIPConfigurationPtrOutput().ApplyT(func(v *NetworkInterfaceIPConfiguration) *NetworkInterfaceIPConfiguration { return v.Defaults() }).(NetworkInterfaceIPConfigurationPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:VirtualNetworkTap"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualNetworkTap
	err := ctx.RegisterResource("azure-native:network:VirtualNetworkTap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkTap gets an existing VirtualNetworkTap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkTap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkTapState, opts ...pulumi.ResourceOption) (*VirtualNetworkTap, error) {
	var resource VirtualNetworkTap
	err := ctx.ReadResource("azure-native:network:VirtualNetworkTap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkTap resources.
type virtualNetworkTapState struct {
}

type VirtualNetworkTapState struct {
}

func (VirtualNetworkTapState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkTapState)(nil)).Elem()
}

type virtualNetworkTapArgs struct {
	// The reference to the private IP address on the internal Load Balancer that will receive the tap.
	DestinationLoadBalancerFrontEndIPConfiguration *FrontendIPConfiguration `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	// The reference to the private IP Address of the collector nic that will receive the tap.
	DestinationNetworkInterfaceIPConfiguration *NetworkInterfaceIPConfiguration `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	// The VXLAN destination port that will receive the tapped traffic.
	DestinationPort *int `pulumi:"destinationPort"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the virtual network tap.
	TapName *string `pulumi:"tapName"`
}

// The set of arguments for constructing a VirtualNetworkTap resource.
type VirtualNetworkTapArgs struct {
	// The reference to the private IP address on the internal Load Balancer that will receive the tap.
	DestinationLoadBalancerFrontEndIPConfiguration FrontendIPConfigurationPtrInput
	// The reference to the private IP Address of the collector nic that will receive the tap.
	DestinationNetworkInterfaceIPConfiguration NetworkInterfaceIPConfigurationPtrInput
	// The VXLAN destination port that will receive the tapped traffic.
	DestinationPort pulumi.IntPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the virtual network tap.
	TapName pulumi.StringPtrInput
}

func (VirtualNetworkTapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkTapArgs)(nil)).Elem()
}

type VirtualNetworkTapInput interface {
	pulumi.Input

	ToVirtualNetworkTapOutput() VirtualNetworkTapOutput
	ToVirtualNetworkTapOutputWithContext(ctx context.Context) VirtualNetworkTapOutput
}

func (*VirtualNetworkTap) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkTap)(nil)).Elem()
}

func (i *VirtualNetworkTap) ToVirtualNetworkTapOutput() VirtualNetworkTapOutput {
	return i.ToVirtualNetworkTapOutputWithContext(context.Background())
}

func (i *VirtualNetworkTap) ToVirtualNetworkTapOutputWithContext(ctx context.Context) VirtualNetworkTapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkTapOutput)
}

type VirtualNetworkTapOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkTap)(nil)).Elem()
}

func (o VirtualNetworkTapOutput) ToVirtualNetworkTapOutput() VirtualNetworkTapOutput {
	return o
}

func (o VirtualNetworkTapOutput) ToVirtualNetworkTapOutputWithContext(ctx context.Context) VirtualNetworkTapOutput {
	return o
}

// The reference to the private IP address on the internal Load Balancer that will receive the tap.
func (o VirtualNetworkTapOutput) DestinationLoadBalancerFrontEndIPConfiguration() FrontendIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) FrontendIPConfigurationResponsePtrOutput {
		return v.DestinationLoadBalancerFrontEndIPConfiguration
	}).(FrontendIPConfigurationResponsePtrOutput)
}

// The reference to the private IP Address of the collector nic that will receive the tap.
func (o VirtualNetworkTapOutput) DestinationNetworkInterfaceIPConfiguration() NetworkInterfaceIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) NetworkInterfaceIPConfigurationResponsePtrOutput {
		return v.DestinationNetworkInterfaceIPConfiguration
	}).(NetworkInterfaceIPConfigurationResponsePtrOutput)
}

// The VXLAN destination port that will receive the tapped traffic.
func (o VirtualNetworkTapOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.IntPtrOutput { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o VirtualNetworkTapOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource location.
func (o VirtualNetworkTapOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o VirtualNetworkTapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the list of resource IDs for the network interface IP configuration that needs to be tapped.
func (o VirtualNetworkTapOutput) NetworkInterfaceTapConfigurations() NetworkInterfaceTapConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) NetworkInterfaceTapConfigurationResponseArrayOutput {
		return v.NetworkInterfaceTapConfigurations
	}).(NetworkInterfaceTapConfigurationResponseArrayOutput)
}

// The provisioning state of the virtual network tap resource.
func (o VirtualNetworkTapOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the virtual network tap resource.
func (o VirtualNetworkTapOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Resource tags.
func (o VirtualNetworkTapOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o VirtualNetworkTapOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkTapOutput{})
}
