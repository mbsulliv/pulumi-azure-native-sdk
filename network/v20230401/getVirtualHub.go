// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves the details of a VirtualHub.
func LookupVirtualHub(ctx *pulumi.Context, args *LookupVirtualHubArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualHubResult
	err := ctx.Invoke("azure-native:network/v20230401:getVirtualHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubArgs struct {
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// VirtualHub Resource.
type LookupVirtualHubResult struct {
	// Address-prefix for this VirtualHub.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// Flag to control transit for VirtualRouter hub.
	AllowBranchToBranchTraffic *bool `pulumi:"allowBranchToBranchTraffic"`
	// The azureFirewall associated with this VirtualHub.
	AzureFirewall *SubResourceResponse `pulumi:"azureFirewall"`
	// List of references to Bgp Connections.
	BgpConnections []SubResourceResponse `pulumi:"bgpConnections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The expressRouteGateway associated with this VirtualHub.
	ExpressRouteGateway *SubResourceResponse `pulumi:"expressRouteGateway"`
	// The hubRoutingPreference of this VirtualHub.
	HubRoutingPreference *string `pulumi:"hubRoutingPreference"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// List of references to IpConfigurations.
	IpConfigurations []SubResourceResponse `pulumi:"ipConfigurations"`
	// Kind of service virtual hub. This is metadata used for the Azure portal experience for Route Server.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The P2SVpnGateway associated with this VirtualHub.
	P2SVpnGateway *SubResourceResponse `pulumi:"p2SVpnGateway"`
	// The preferred gateway to route on-prem traffic
	PreferredRoutingGateway *string `pulumi:"preferredRoutingGateway"`
	// The provisioning state of the virtual hub resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// List of references to RouteMaps.
	RouteMaps []SubResourceResponse `pulumi:"routeMaps"`
	// The routeTable associated with this virtual hub.
	RouteTable *VirtualHubRouteTableResponse `pulumi:"routeTable"`
	// The routing state.
	RoutingState string `pulumi:"routingState"`
	// The securityPartnerProvider associated with this VirtualHub.
	SecurityPartnerProvider *SubResourceResponse `pulumi:"securityPartnerProvider"`
	// The Security Provider name.
	SecurityProviderName *string `pulumi:"securityProviderName"`
	// The sku of this VirtualHub.
	Sku *string `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// List of all virtual hub route table v2s associated with this VirtualHub.
	VirtualHubRouteTableV2s []VirtualHubRouteTableV2Response `pulumi:"virtualHubRouteTableV2s"`
	// VirtualRouter ASN.
	VirtualRouterAsn *float64 `pulumi:"virtualRouterAsn"`
	// The VirtualHub Router autoscale configuration.
	VirtualRouterAutoScaleConfiguration *VirtualRouterAutoScaleConfigurationResponse `pulumi:"virtualRouterAutoScaleConfiguration"`
	// VirtualRouter IPs.
	VirtualRouterIps []string `pulumi:"virtualRouterIps"`
	// The VirtualWAN to which the VirtualHub belongs.
	VirtualWan *SubResourceResponse `pulumi:"virtualWan"`
	// The VpnGateway associated with this VirtualHub.
	VpnGateway *SubResourceResponse `pulumi:"vpnGateway"`
}

func LookupVirtualHubOutput(ctx *pulumi.Context, args LookupVirtualHubOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualHubResult, error) {
			args := v.(LookupVirtualHubArgs)
			r, err := LookupVirtualHub(ctx, &args, opts...)
			var s LookupVirtualHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualHubResultOutput)
}

type LookupVirtualHubOutputArgs struct {
	// The resource group name of the VirtualHub.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupVirtualHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubArgs)(nil)).Elem()
}

// VirtualHub Resource.
type LookupVirtualHubResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubResult)(nil)).Elem()
}

func (o LookupVirtualHubResultOutput) ToLookupVirtualHubResultOutput() LookupVirtualHubResultOutput {
	return o
}

func (o LookupVirtualHubResultOutput) ToLookupVirtualHubResultOutputWithContext(ctx context.Context) LookupVirtualHubResultOutput {
	return o
}

func (o LookupVirtualHubResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualHubResult] {
	return pulumix.Output[LookupVirtualHubResult]{
		OutputState: o.OutputState,
	}
}

// Address-prefix for this VirtualHub.
func (o LookupVirtualHubResultOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

// Flag to control transit for VirtualRouter hub.
func (o LookupVirtualHubResultOutput) AllowBranchToBranchTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *bool { return v.AllowBranchToBranchTraffic }).(pulumi.BoolPtrOutput)
}

// The azureFirewall associated with this VirtualHub.
func (o LookupVirtualHubResultOutput) AzureFirewall() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.AzureFirewall }).(SubResourceResponsePtrOutput)
}

// List of references to Bgp Connections.
func (o LookupVirtualHubResultOutput) BgpConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) []SubResourceResponse { return v.BgpConnections }).(SubResourceResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVirtualHubResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The expressRouteGateway associated with this VirtualHub.
func (o LookupVirtualHubResultOutput) ExpressRouteGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.ExpressRouteGateway }).(SubResourceResponsePtrOutput)
}

// The hubRoutingPreference of this VirtualHub.
func (o LookupVirtualHubResultOutput) HubRoutingPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.HubRoutingPreference }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupVirtualHubResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// List of references to IpConfigurations.
func (o LookupVirtualHubResultOutput) IpConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) []SubResourceResponse { return v.IpConfigurations }).(SubResourceResponseArrayOutput)
}

// Kind of service virtual hub. This is metadata used for the Azure portal experience for Route Server.
func (o LookupVirtualHubResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupVirtualHubResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupVirtualHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Name }).(pulumi.StringOutput)
}

// The P2SVpnGateway associated with this VirtualHub.
func (o LookupVirtualHubResultOutput) P2SVpnGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.P2SVpnGateway }).(SubResourceResponsePtrOutput)
}

// The preferred gateway to route on-prem traffic
func (o LookupVirtualHubResultOutput) PreferredRoutingGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.PreferredRoutingGateway }).(pulumi.StringPtrOutput)
}

// The provisioning state of the virtual hub resource.
func (o LookupVirtualHubResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of references to RouteMaps.
func (o LookupVirtualHubResultOutput) RouteMaps() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) []SubResourceResponse { return v.RouteMaps }).(SubResourceResponseArrayOutput)
}

// The routeTable associated with this virtual hub.
func (o LookupVirtualHubResultOutput) RouteTable() VirtualHubRouteTableResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *VirtualHubRouteTableResponse { return v.RouteTable }).(VirtualHubRouteTableResponsePtrOutput)
}

// The routing state.
func (o LookupVirtualHubResultOutput) RoutingState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.RoutingState }).(pulumi.StringOutput)
}

// The securityPartnerProvider associated with this VirtualHub.
func (o LookupVirtualHubResultOutput) SecurityPartnerProvider() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.SecurityPartnerProvider }).(SubResourceResponsePtrOutput)
}

// The Security Provider name.
func (o LookupVirtualHubResultOutput) SecurityProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.SecurityProviderName }).(pulumi.StringPtrOutput)
}

// The sku of this VirtualHub.
func (o LookupVirtualHubResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupVirtualHubResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupVirtualHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Type }).(pulumi.StringOutput)
}

// List of all virtual hub route table v2s associated with this VirtualHub.
func (o LookupVirtualHubResultOutput) VirtualHubRouteTableV2s() VirtualHubRouteTableV2ResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) []VirtualHubRouteTableV2Response { return v.VirtualHubRouteTableV2s }).(VirtualHubRouteTableV2ResponseArrayOutput)
}

// VirtualRouter ASN.
func (o LookupVirtualHubResultOutput) VirtualRouterAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *float64 { return v.VirtualRouterAsn }).(pulumi.Float64PtrOutput)
}

// The VirtualHub Router autoscale configuration.
func (o LookupVirtualHubResultOutput) VirtualRouterAutoScaleConfiguration() VirtualRouterAutoScaleConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *VirtualRouterAutoScaleConfigurationResponse {
		return v.VirtualRouterAutoScaleConfiguration
	}).(VirtualRouterAutoScaleConfigurationResponsePtrOutput)
}

// VirtualRouter IPs.
func (o LookupVirtualHubResultOutput) VirtualRouterIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) []string { return v.VirtualRouterIps }).(pulumi.StringArrayOutput)
}

// The VirtualWAN to which the VirtualHub belongs.
func (o LookupVirtualHubResultOutput) VirtualWan() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.VirtualWan }).(SubResourceResponsePtrOutput)
}

// The VpnGateway associated with this VirtualHub.
func (o LookupVirtualHubResultOutput) VpnGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.VpnGateway }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualHubResultOutput{})
}
