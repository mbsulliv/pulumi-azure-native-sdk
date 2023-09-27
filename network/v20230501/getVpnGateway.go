// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves the details of a virtual wan vpn gateway.
func LookupVpnGateway(ctx *pulumi.Context, args *LookupVpnGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVpnGatewayResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVpnGatewayResult
	err := ctx.Invoke("azure-native:network/v20230501:getVpnGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnGatewayArgs struct {
	// The name of the gateway.
	GatewayName string `pulumi:"gatewayName"`
	// The resource group name of the VpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// VpnGateway Resource.
type LookupVpnGatewayResult struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings *BgpSettingsResponse `pulumi:"bgpSettings"`
	// List of all vpn connections to the gateway.
	Connections []VpnConnectionResponse `pulumi:"connections"`
	// Enable BGP routes translation for NAT on this VpnGateway.
	EnableBgpRouteTranslationForNat *bool `pulumi:"enableBgpRouteTranslationForNat"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// List of all IPs configured on the gateway.
	IpConfigurations []VpnGatewayIpConfigurationResponse `pulumi:"ipConfigurations"`
	// Enable Routing Preference property for the Public IP Interface of the VpnGateway.
	IsRoutingPreferenceInternet *bool `pulumi:"isRoutingPreferenceInternet"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// List of all the nat Rules associated with the gateway.
	NatRules []VpnGatewayNatRuleResponse `pulumi:"natRules"`
	// The provisioning state of the VPN gateway resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The VirtualHub to which the gateway belongs.
	VirtualHub *SubResourceResponse `pulumi:"virtualHub"`
	// The scale unit for this vpn gateway.
	VpnGatewayScaleUnit *int `pulumi:"vpnGatewayScaleUnit"`
}

func LookupVpnGatewayOutput(ctx *pulumi.Context, args LookupVpnGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupVpnGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpnGatewayResult, error) {
			args := v.(LookupVpnGatewayArgs)
			r, err := LookupVpnGateway(ctx, &args, opts...)
			var s LookupVpnGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpnGatewayResultOutput)
}

type LookupVpnGatewayOutputArgs struct {
	// The name of the gateway.
	GatewayName pulumi.StringInput `pulumi:"gatewayName"`
	// The resource group name of the VpnGateway.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupVpnGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnGatewayArgs)(nil)).Elem()
}

// VpnGateway Resource.
type LookupVpnGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupVpnGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnGatewayResult)(nil)).Elem()
}

func (o LookupVpnGatewayResultOutput) ToLookupVpnGatewayResultOutput() LookupVpnGatewayResultOutput {
	return o
}

func (o LookupVpnGatewayResultOutput) ToLookupVpnGatewayResultOutputWithContext(ctx context.Context) LookupVpnGatewayResultOutput {
	return o
}

func (o LookupVpnGatewayResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVpnGatewayResult] {
	return pulumix.Output[LookupVpnGatewayResult]{
		OutputState: o.OutputState,
	}
}

// Local network gateway's BGP speaker settings.
func (o LookupVpnGatewayResultOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) *BgpSettingsResponse { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

// List of all vpn connections to the gateway.
func (o LookupVpnGatewayResultOutput) Connections() VpnConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) []VpnConnectionResponse { return v.Connections }).(VpnConnectionResponseArrayOutput)
}

// Enable BGP routes translation for NAT on this VpnGateway.
func (o LookupVpnGatewayResultOutput) EnableBgpRouteTranslationForNat() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) *bool { return v.EnableBgpRouteTranslationForNat }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVpnGatewayResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupVpnGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// List of all IPs configured on the gateway.
func (o LookupVpnGatewayResultOutput) IpConfigurations() VpnGatewayIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) []VpnGatewayIpConfigurationResponse { return v.IpConfigurations }).(VpnGatewayIpConfigurationResponseArrayOutput)
}

// Enable Routing Preference property for the Public IP Interface of the VpnGateway.
func (o LookupVpnGatewayResultOutput) IsRoutingPreferenceInternet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) *bool { return v.IsRoutingPreferenceInternet }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o LookupVpnGatewayResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupVpnGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of all the nat Rules associated with the gateway.
func (o LookupVpnGatewayResultOutput) NatRules() VpnGatewayNatRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) []VpnGatewayNatRuleResponse { return v.NatRules }).(VpnGatewayNatRuleResponseArrayOutput)
}

// The provisioning state of the VPN gateway resource.
func (o LookupVpnGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupVpnGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupVpnGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

// The VirtualHub to which the gateway belongs.
func (o LookupVpnGatewayResultOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) *SubResourceResponse { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

// The scale unit for this vpn gateway.
func (o LookupVpnGatewayResultOutput) VpnGatewayScaleUnit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) *int { return v.VpnGatewayScaleUnit }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnGatewayResultOutput{})
}
