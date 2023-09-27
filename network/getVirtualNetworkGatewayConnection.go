// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified virtual network gateway connection by resource group.
// Azure REST API version: 2023-02-01.
func LookupVirtualNetworkGatewayConnection(ctx *pulumi.Context, args *LookupVirtualNetworkGatewayConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkGatewayConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualNetworkGatewayConnectionResult
	err := ctx.Invoke("azure-native:network:getVirtualNetworkGatewayConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkGatewayConnectionArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual network gateway connection.
	VirtualNetworkGatewayConnectionName string `pulumi:"virtualNetworkGatewayConnectionName"`
}

// A common class for general resource information.
type LookupVirtualNetworkGatewayConnectionResult struct {
	// The authorizationKey.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The connection mode for this connection.
	ConnectionMode *string `pulumi:"connectionMode"`
	// Connection protocol used for this connection.
	ConnectionProtocol *string `pulumi:"connectionProtocol"`
	// Virtual Network Gateway connection status.
	ConnectionStatus string `pulumi:"connectionStatus"`
	// Gateway connection type.
	ConnectionType string `pulumi:"connectionType"`
	// The dead peer detection timeout of this connection in seconds.
	DpdTimeoutSeconds *int `pulumi:"dpdTimeoutSeconds"`
	// The egress bytes transferred in this connection.
	EgressBytesTransferred float64 `pulumi:"egressBytesTransferred"`
	// List of egress NatRules.
	EgressNatRules []SubResourceResponse `pulumi:"egressNatRules"`
	// EnableBgp flag.
	EnableBgp *bool `pulumi:"enableBgp"`
	// Bypass the ExpressRoute gateway when accessing private-links. ExpressRoute FastPath (expressRouteGatewayBypass) must be enabled.
	EnablePrivateLinkFastPath *bool `pulumi:"enablePrivateLinkFastPath"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Bypass ExpressRoute Gateway for data forwarding.
	ExpressRouteGatewayBypass *bool `pulumi:"expressRouteGatewayBypass"`
	// GatewayCustomBgpIpAddresses to be used for virtual network gateway Connection.
	GatewayCustomBgpIpAddresses []GatewayCustomBgpIpAddressIpConfigurationResponse `pulumi:"gatewayCustomBgpIpAddresses"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The ingress bytes transferred in this connection.
	IngressBytesTransferred float64 `pulumi:"ingressBytesTransferred"`
	// List of ingress NatRules.
	IngressNatRules []SubResourceResponse `pulumi:"ingressNatRules"`
	// The IPSec Policies to be considered by this connection.
	IpsecPolicies []IpsecPolicyResponse `pulumi:"ipsecPolicies"`
	// The reference to local network gateway resource.
	LocalNetworkGateway2 *LocalNetworkGatewayResponse `pulumi:"localNetworkGateway2"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The reference to peerings resource.
	Peer *SubResourceResponse `pulumi:"peer"`
	// The provisioning state of the virtual network gateway connection resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource GUID property of the virtual network gateway connection resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The routing weight.
	RoutingWeight *int `pulumi:"routingWeight"`
	// The IPSec shared key.
	SharedKey *string `pulumi:"sharedKey"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The Traffic Selector Policies to be considered by this connection.
	TrafficSelectorPolicies []TrafficSelectorPolicyResponse `pulumi:"trafficSelectorPolicies"`
	// Collection of all tunnels' connection health status.
	TunnelConnectionStatus []TunnelConnectionHealthResponse `pulumi:"tunnelConnectionStatus"`
	// Resource type.
	Type string `pulumi:"type"`
	// Use private local Azure IP for the connection.
	UseLocalAzureIpAddress *bool `pulumi:"useLocalAzureIpAddress"`
	// Enable policy-based traffic selectors.
	UsePolicyBasedTrafficSelectors *bool `pulumi:"usePolicyBasedTrafficSelectors"`
	// The reference to virtual network gateway resource.
	VirtualNetworkGateway1 VirtualNetworkGatewayResponse `pulumi:"virtualNetworkGateway1"`
	// The reference to virtual network gateway resource.
	VirtualNetworkGateway2 *VirtualNetworkGatewayResponse `pulumi:"virtualNetworkGateway2"`
}

func LookupVirtualNetworkGatewayConnectionOutput(ctx *pulumi.Context, args LookupVirtualNetworkGatewayConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkGatewayConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkGatewayConnectionResult, error) {
			args := v.(LookupVirtualNetworkGatewayConnectionArgs)
			r, err := LookupVirtualNetworkGatewayConnection(ctx, &args, opts...)
			var s LookupVirtualNetworkGatewayConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkGatewayConnectionResultOutput)
}

type LookupVirtualNetworkGatewayConnectionOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the virtual network gateway connection.
	VirtualNetworkGatewayConnectionName pulumi.StringInput `pulumi:"virtualNetworkGatewayConnectionName"`
}

func (LookupVirtualNetworkGatewayConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkGatewayConnectionArgs)(nil)).Elem()
}

// A common class for general resource information.
type LookupVirtualNetworkGatewayConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkGatewayConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkGatewayConnectionResult)(nil)).Elem()
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) ToLookupVirtualNetworkGatewayConnectionResultOutput() LookupVirtualNetworkGatewayConnectionResultOutput {
	return o
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) ToLookupVirtualNetworkGatewayConnectionResultOutputWithContext(ctx context.Context) LookupVirtualNetworkGatewayConnectionResultOutput {
	return o
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualNetworkGatewayConnectionResult] {
	return pulumix.Output[LookupVirtualNetworkGatewayConnectionResult]{
		OutputState: o.OutputState,
	}
}

// The authorizationKey.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

// The connection mode for this connection.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) ConnectionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *string { return v.ConnectionMode }).(pulumi.StringPtrOutput)
}

// Connection protocol used for this connection.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) ConnectionProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *string { return v.ConnectionProtocol }).(pulumi.StringPtrOutput)
}

// Virtual Network Gateway connection status.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

// Gateway connection type.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) ConnectionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.ConnectionType }).(pulumi.StringOutput)
}

// The dead peer detection timeout of this connection in seconds.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) DpdTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *int { return v.DpdTimeoutSeconds }).(pulumi.IntPtrOutput)
}

// The egress bytes transferred in this connection.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

// List of egress NatRules.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) EgressNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) []SubResourceResponse { return v.EgressNatRules }).(SubResourceResponseArrayOutput)
}

// EnableBgp flag.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

// Bypass the ExpressRoute gateway when accessing private-links. ExpressRoute FastPath (expressRouteGatewayBypass) must be enabled.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) EnablePrivateLinkFastPath() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *bool { return v.EnablePrivateLinkFastPath }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Bypass ExpressRoute Gateway for data forwarding.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) ExpressRouteGatewayBypass() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *bool { return v.ExpressRouteGatewayBypass }).(pulumi.BoolPtrOutput)
}

// GatewayCustomBgpIpAddresses to be used for virtual network gateway Connection.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) GatewayCustomBgpIpAddresses() GatewayCustomBgpIpAddressIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) []GatewayCustomBgpIpAddressIpConfigurationResponse {
		return v.GatewayCustomBgpIpAddresses
	}).(GatewayCustomBgpIpAddressIpConfigurationResponseArrayOutput)
}

// Resource ID.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The ingress bytes transferred in this connection.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

// List of ingress NatRules.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) IngressNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) []SubResourceResponse { return v.IngressNatRules }).(SubResourceResponseArrayOutput)
}

// The IPSec Policies to be considered by this connection.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) IpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) []IpsecPolicyResponse { return v.IpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

// The reference to local network gateway resource.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) LocalNetworkGateway2() LocalNetworkGatewayResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *LocalNetworkGatewayResponse {
		return v.LocalNetworkGateway2
	}).(LocalNetworkGatewayResponsePtrOutput)
}

// Resource location.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The reference to peerings resource.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) Peer() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *SubResourceResponse { return v.Peer }).(SubResourceResponsePtrOutput)
}

// The provisioning state of the virtual network gateway connection resource.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the virtual network gateway connection resource.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The routing weight.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

// The IPSec shared key.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The Traffic Selector Policies to be considered by this connection.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) TrafficSelectorPolicies() TrafficSelectorPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) []TrafficSelectorPolicyResponse {
		return v.TrafficSelectorPolicies
	}).(TrafficSelectorPolicyResponseArrayOutput)
}

// Collection of all tunnels' connection health status.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) TunnelConnectionStatus() TunnelConnectionHealthResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) []TunnelConnectionHealthResponse {
		return v.TunnelConnectionStatus
	}).(TunnelConnectionHealthResponseArrayOutput)
}

// Resource type.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

// Use private local Azure IP for the connection.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) UseLocalAzureIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *bool { return v.UseLocalAzureIpAddress }).(pulumi.BoolPtrOutput)
}

// Enable policy-based traffic selectors.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *bool { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

// The reference to virtual network gateway resource.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) VirtualNetworkGateway1() VirtualNetworkGatewayResponseOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) VirtualNetworkGatewayResponse {
		return v.VirtualNetworkGateway1
	}).(VirtualNetworkGatewayResponseOutput)
}

// The reference to virtual network gateway resource.
func (o LookupVirtualNetworkGatewayConnectionResultOutput) VirtualNetworkGateway2() VirtualNetworkGatewayResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *VirtualNetworkGatewayResponse {
		return v.VirtualNetworkGateway2
	}).(VirtualNetworkGatewayResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkGatewayConnectionResultOutput{})
}
