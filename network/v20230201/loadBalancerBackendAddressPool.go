// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Pool of backend IP addresses.
type LoadBalancerBackendAddressPool struct {
	pulumi.CustomResourceState

	// An array of references to IP addresses defined in network interfaces.
	BackendIPConfigurations NetworkInterfaceIPConfigurationResponseArrayOutput `pulumi:"backendIPConfigurations"`
	// Amount of seconds Load Balancer waits for before sending RESET to client and backend address.
	DrainPeriodInSeconds pulumi.IntPtrOutput `pulumi:"drainPeriodInSeconds"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// An array of references to inbound NAT rules that use this backend address pool.
	InboundNatRules SubResourceResponseArrayOutput `pulumi:"inboundNatRules"`
	// An array of backend addresses.
	LoadBalancerBackendAddresses LoadBalancerBackendAddressResponseArrayOutput `pulumi:"loadBalancerBackendAddresses"`
	// An array of references to load balancing rules that use this backend address pool.
	LoadBalancingRules SubResourceResponseArrayOutput `pulumi:"loadBalancingRules"`
	// The location of the backend address pool.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// A reference to an outbound rule that uses this backend address pool.
	OutboundRule SubResourceResponseOutput `pulumi:"outboundRule"`
	// An array of references to outbound rules that use this backend address pool.
	OutboundRules SubResourceResponseArrayOutput `pulumi:"outboundRules"`
	// The provisioning state of the backend address pool resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// An array of gateway load balancer tunnel interfaces.
	TunnelInterfaces GatewayLoadBalancerTunnelInterfaceResponseArrayOutput `pulumi:"tunnelInterfaces"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// A reference to a virtual network.
	VirtualNetwork SubResourceResponsePtrOutput `pulumi:"virtualNetwork"`
}

// NewLoadBalancerBackendAddressPool registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancerBackendAddressPool(ctx *pulumi.Context,
	name string, args *LoadBalancerBackendAddressPoolArgs, opts ...pulumi.ResourceOption) (*LoadBalancerBackendAddressPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerName == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:LoadBalancerBackendAddressPool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LoadBalancerBackendAddressPool
	err := ctx.RegisterResource("azure-native:network/v20230201:LoadBalancerBackendAddressPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancerBackendAddressPool gets an existing LoadBalancerBackendAddressPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancerBackendAddressPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerBackendAddressPoolState, opts ...pulumi.ResourceOption) (*LoadBalancerBackendAddressPool, error) {
	var resource LoadBalancerBackendAddressPool
	err := ctx.ReadResource("azure-native:network/v20230201:LoadBalancerBackendAddressPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancerBackendAddressPool resources.
type loadBalancerBackendAddressPoolState struct {
}

type LoadBalancerBackendAddressPoolState struct {
}

func (LoadBalancerBackendAddressPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerBackendAddressPoolState)(nil)).Elem()
}

type loadBalancerBackendAddressPoolArgs struct {
	// The name of the backend address pool.
	BackendAddressPoolName *string `pulumi:"backendAddressPoolName"`
	// Amount of seconds Load Balancer waits for before sending RESET to client and backend address.
	DrainPeriodInSeconds *int `pulumi:"drainPeriodInSeconds"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// An array of backend addresses.
	LoadBalancerBackendAddresses []LoadBalancerBackendAddress `pulumi:"loadBalancerBackendAddresses"`
	// The name of the load balancer.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// The location of the backend address pool.
	Location *string `pulumi:"location"`
	// The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// An array of gateway load balancer tunnel interfaces.
	TunnelInterfaces []GatewayLoadBalancerTunnelInterface `pulumi:"tunnelInterfaces"`
	// A reference to a virtual network.
	VirtualNetwork *SubResource `pulumi:"virtualNetwork"`
}

// The set of arguments for constructing a LoadBalancerBackendAddressPool resource.
type LoadBalancerBackendAddressPoolArgs struct {
	// The name of the backend address pool.
	BackendAddressPoolName pulumi.StringPtrInput
	// Amount of seconds Load Balancer waits for before sending RESET to client and backend address.
	DrainPeriodInSeconds pulumi.IntPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// An array of backend addresses.
	LoadBalancerBackendAddresses LoadBalancerBackendAddressArrayInput
	// The name of the load balancer.
	LoadBalancerName pulumi.StringInput
	// The location of the backend address pool.
	Location pulumi.StringPtrInput
	// The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// An array of gateway load balancer tunnel interfaces.
	TunnelInterfaces GatewayLoadBalancerTunnelInterfaceArrayInput
	// A reference to a virtual network.
	VirtualNetwork SubResourcePtrInput
}

func (LoadBalancerBackendAddressPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerBackendAddressPoolArgs)(nil)).Elem()
}

type LoadBalancerBackendAddressPoolInput interface {
	pulumi.Input

	ToLoadBalancerBackendAddressPoolOutput() LoadBalancerBackendAddressPoolOutput
	ToLoadBalancerBackendAddressPoolOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolOutput
}

func (*LoadBalancerBackendAddressPool) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerBackendAddressPool)(nil)).Elem()
}

func (i *LoadBalancerBackendAddressPool) ToLoadBalancerBackendAddressPoolOutput() LoadBalancerBackendAddressPoolOutput {
	return i.ToLoadBalancerBackendAddressPoolOutputWithContext(context.Background())
}

func (i *LoadBalancerBackendAddressPool) ToLoadBalancerBackendAddressPoolOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerBackendAddressPoolOutput)
}

type LoadBalancerBackendAddressPoolOutput struct{ *pulumi.OutputState }

func (LoadBalancerBackendAddressPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerBackendAddressPool)(nil)).Elem()
}

func (o LoadBalancerBackendAddressPoolOutput) ToLoadBalancerBackendAddressPoolOutput() LoadBalancerBackendAddressPoolOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolOutput) ToLoadBalancerBackendAddressPoolOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolOutput {
	return o
}

// An array of references to IP addresses defined in network interfaces.
func (o LoadBalancerBackendAddressPoolOutput) BackendIPConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) NetworkInterfaceIPConfigurationResponseArrayOutput {
		return v.BackendIPConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

// Amount of seconds Load Balancer waits for before sending RESET to client and backend address.
func (o LoadBalancerBackendAddressPoolOutput) DrainPeriodInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.IntPtrOutput { return v.DrainPeriodInSeconds }).(pulumi.IntPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LoadBalancerBackendAddressPoolOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// An array of references to inbound NAT rules that use this backend address pool.
func (o LoadBalancerBackendAddressPoolOutput) InboundNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) SubResourceResponseArrayOutput { return v.InboundNatRules }).(SubResourceResponseArrayOutput)
}

// An array of backend addresses.
func (o LoadBalancerBackendAddressPoolOutput) LoadBalancerBackendAddresses() LoadBalancerBackendAddressResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) LoadBalancerBackendAddressResponseArrayOutput {
		return v.LoadBalancerBackendAddresses
	}).(LoadBalancerBackendAddressResponseArrayOutput)
}

// An array of references to load balancing rules that use this backend address pool.
func (o LoadBalancerBackendAddressPoolOutput) LoadBalancingRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) SubResourceResponseArrayOutput { return v.LoadBalancingRules }).(SubResourceResponseArrayOutput)
}

// The location of the backend address pool.
func (o LoadBalancerBackendAddressPoolOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
func (o LoadBalancerBackendAddressPoolOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A reference to an outbound rule that uses this backend address pool.
func (o LoadBalancerBackendAddressPoolOutput) OutboundRule() SubResourceResponseOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) SubResourceResponseOutput { return v.OutboundRule }).(SubResourceResponseOutput)
}

// An array of references to outbound rules that use this backend address pool.
func (o LoadBalancerBackendAddressPoolOutput) OutboundRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) SubResourceResponseArrayOutput { return v.OutboundRules }).(SubResourceResponseArrayOutput)
}

// The provisioning state of the backend address pool resource.
func (o LoadBalancerBackendAddressPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// An array of gateway load balancer tunnel interfaces.
func (o LoadBalancerBackendAddressPoolOutput) TunnelInterfaces() GatewayLoadBalancerTunnelInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) GatewayLoadBalancerTunnelInterfaceResponseArrayOutput {
		return v.TunnelInterfaces
	}).(GatewayLoadBalancerTunnelInterfaceResponseArrayOutput)
}

// Type of the resource.
func (o LoadBalancerBackendAddressPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A reference to a virtual network.
func (o LoadBalancerBackendAddressPoolOutput) VirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) SubResourceResponsePtrOutput { return v.VirtualNetwork }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LoadBalancerBackendAddressPoolOutput{})
}
