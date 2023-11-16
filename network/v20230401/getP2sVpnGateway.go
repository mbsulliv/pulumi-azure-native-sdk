// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the details of a virtual wan p2s vpn gateway.
func LookupP2sVpnGateway(ctx *pulumi.Context, args *LookupP2sVpnGatewayArgs, opts ...pulumi.InvokeOption) (*LookupP2sVpnGatewayResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupP2sVpnGatewayResult
	err := ctx.Invoke("azure-native:network/v20230401:getP2sVpnGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupP2sVpnGatewayArgs struct {
	// The name of the gateway.
	GatewayName string `pulumi:"gatewayName"`
	// The resource group name of the P2SVpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// P2SVpnGateway Resource.
type LookupP2sVpnGatewayResult struct {
	// List of all customer specified DNS servers IP addresses.
	CustomDnsServers []string `pulumi:"customDnsServers"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Enable Routing Preference property for the Public IP Interface of the P2SVpnGateway.
	IsRoutingPreferenceInternet *bool `pulumi:"isRoutingPreferenceInternet"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// List of all p2s connection configurations of the gateway.
	P2SConnectionConfigurations []P2SConnectionConfigurationResponse `pulumi:"p2SConnectionConfigurations"`
	// The provisioning state of the P2S VPN gateway resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The VirtualHub to which the gateway belongs.
	VirtualHub *SubResourceResponse `pulumi:"virtualHub"`
	// All P2S VPN clients' connection health status.
	VpnClientConnectionHealth VpnClientConnectionHealthResponse `pulumi:"vpnClientConnectionHealth"`
	// The scale unit for this p2s vpn gateway.
	VpnGatewayScaleUnit *int `pulumi:"vpnGatewayScaleUnit"`
	// The VpnServerConfiguration to which the p2sVpnGateway is attached to.
	VpnServerConfiguration *SubResourceResponse `pulumi:"vpnServerConfiguration"`
}

func LookupP2sVpnGatewayOutput(ctx *pulumi.Context, args LookupP2sVpnGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupP2sVpnGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupP2sVpnGatewayResult, error) {
			args := v.(LookupP2sVpnGatewayArgs)
			r, err := LookupP2sVpnGateway(ctx, &args, opts...)
			var s LookupP2sVpnGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupP2sVpnGatewayResultOutput)
}

type LookupP2sVpnGatewayOutputArgs struct {
	// The name of the gateway.
	GatewayName pulumi.StringInput `pulumi:"gatewayName"`
	// The resource group name of the P2SVpnGateway.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupP2sVpnGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupP2sVpnGatewayArgs)(nil)).Elem()
}

// P2SVpnGateway Resource.
type LookupP2sVpnGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupP2sVpnGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupP2sVpnGatewayResult)(nil)).Elem()
}

func (o LookupP2sVpnGatewayResultOutput) ToLookupP2sVpnGatewayResultOutput() LookupP2sVpnGatewayResultOutput {
	return o
}

func (o LookupP2sVpnGatewayResultOutput) ToLookupP2sVpnGatewayResultOutputWithContext(ctx context.Context) LookupP2sVpnGatewayResultOutput {
	return o
}

// List of all customer specified DNS servers IP addresses.
func (o LookupP2sVpnGatewayResultOutput) CustomDnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) []string { return v.CustomDnsServers }).(pulumi.StringArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupP2sVpnGatewayResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupP2sVpnGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Enable Routing Preference property for the Public IP Interface of the P2SVpnGateway.
func (o LookupP2sVpnGatewayResultOutput) IsRoutingPreferenceInternet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) *bool { return v.IsRoutingPreferenceInternet }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o LookupP2sVpnGatewayResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupP2sVpnGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of all p2s connection configurations of the gateway.
func (o LookupP2sVpnGatewayResultOutput) P2SConnectionConfigurations() P2SConnectionConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) []P2SConnectionConfigurationResponse {
		return v.P2SConnectionConfigurations
	}).(P2SConnectionConfigurationResponseArrayOutput)
}

// The provisioning state of the P2S VPN gateway resource.
func (o LookupP2sVpnGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupP2sVpnGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupP2sVpnGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

// The VirtualHub to which the gateway belongs.
func (o LookupP2sVpnGatewayResultOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) *SubResourceResponse { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

// All P2S VPN clients' connection health status.
func (o LookupP2sVpnGatewayResultOutput) VpnClientConnectionHealth() VpnClientConnectionHealthResponseOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) VpnClientConnectionHealthResponse {
		return v.VpnClientConnectionHealth
	}).(VpnClientConnectionHealthResponseOutput)
}

// The scale unit for this p2s vpn gateway.
func (o LookupP2sVpnGatewayResultOutput) VpnGatewayScaleUnit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) *int { return v.VpnGatewayScaleUnit }).(pulumi.IntPtrOutput)
}

// The VpnServerConfiguration to which the p2sVpnGateway is attached to.
func (o LookupP2sVpnGatewayResultOutput) VpnServerConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) *SubResourceResponse { return v.VpnServerConfiguration }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupP2sVpnGatewayResultOutput{})
}
