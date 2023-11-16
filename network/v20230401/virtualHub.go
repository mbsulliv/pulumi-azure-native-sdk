// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VirtualHub Resource.
type VirtualHub struct {
	pulumi.CustomResourceState

	// Address-prefix for this VirtualHub.
	AddressPrefix pulumi.StringPtrOutput `pulumi:"addressPrefix"`
	// Flag to control transit for VirtualRouter hub.
	AllowBranchToBranchTraffic pulumi.BoolPtrOutput `pulumi:"allowBranchToBranchTraffic"`
	// The azureFirewall associated with this VirtualHub.
	AzureFirewall SubResourceResponsePtrOutput `pulumi:"azureFirewall"`
	// List of references to Bgp Connections.
	BgpConnections SubResourceResponseArrayOutput `pulumi:"bgpConnections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The expressRouteGateway associated with this VirtualHub.
	ExpressRouteGateway SubResourceResponsePtrOutput `pulumi:"expressRouteGateway"`
	// The hubRoutingPreference of this VirtualHub.
	HubRoutingPreference pulumi.StringPtrOutput `pulumi:"hubRoutingPreference"`
	// List of references to IpConfigurations.
	IpConfigurations SubResourceResponseArrayOutput `pulumi:"ipConfigurations"`
	// Kind of service virtual hub. This is metadata used for the Azure portal experience for Route Server.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The P2SVpnGateway associated with this VirtualHub.
	P2SVpnGateway SubResourceResponsePtrOutput `pulumi:"p2SVpnGateway"`
	// The preferred gateway to route on-prem traffic
	PreferredRoutingGateway pulumi.StringPtrOutput `pulumi:"preferredRoutingGateway"`
	// The provisioning state of the virtual hub resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// List of references to RouteMaps.
	RouteMaps SubResourceResponseArrayOutput `pulumi:"routeMaps"`
	// The routeTable associated with this virtual hub.
	RouteTable VirtualHubRouteTableResponsePtrOutput `pulumi:"routeTable"`
	// The routing state.
	RoutingState pulumi.StringOutput `pulumi:"routingState"`
	// The securityPartnerProvider associated with this VirtualHub.
	SecurityPartnerProvider SubResourceResponsePtrOutput `pulumi:"securityPartnerProvider"`
	// The Security Provider name.
	SecurityProviderName pulumi.StringPtrOutput `pulumi:"securityProviderName"`
	// The sku of this VirtualHub.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// List of all virtual hub route table v2s associated with this VirtualHub.
	VirtualHubRouteTableV2s VirtualHubRouteTableV2ResponseArrayOutput `pulumi:"virtualHubRouteTableV2s"`
	// VirtualRouter ASN.
	VirtualRouterAsn pulumi.Float64PtrOutput `pulumi:"virtualRouterAsn"`
	// The VirtualHub Router autoscale configuration.
	VirtualRouterAutoScaleConfiguration VirtualRouterAutoScaleConfigurationResponsePtrOutput `pulumi:"virtualRouterAutoScaleConfiguration"`
	// VirtualRouter IPs.
	VirtualRouterIps pulumi.StringArrayOutput `pulumi:"virtualRouterIps"`
	// The VirtualWAN to which the VirtualHub belongs.
	VirtualWan SubResourceResponsePtrOutput `pulumi:"virtualWan"`
	// The VpnGateway associated with this VirtualHub.
	VpnGateway SubResourceResponsePtrOutput `pulumi:"vpnGateway"`
}

// NewVirtualHub registers a new resource with the given unique name, arguments, and options.
func NewVirtualHub(ctx *pulumi.Context,
	name string, args *VirtualHubArgs, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:VirtualHub"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualHub
	err := ctx.RegisterResource("azure-native:network/v20230401:VirtualHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHub gets an existing VirtualHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubState, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	var resource VirtualHub
	err := ctx.ReadResource("azure-native:network/v20230401:VirtualHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHub resources.
type virtualHubState struct {
}

type VirtualHubState struct {
}

func (VirtualHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubState)(nil)).Elem()
}

type virtualHubArgs struct {
	// Address-prefix for this VirtualHub.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// Flag to control transit for VirtualRouter hub.
	AllowBranchToBranchTraffic *bool `pulumi:"allowBranchToBranchTraffic"`
	// The azureFirewall associated with this VirtualHub.
	AzureFirewall *SubResource `pulumi:"azureFirewall"`
	// The expressRouteGateway associated with this VirtualHub.
	ExpressRouteGateway *SubResource `pulumi:"expressRouteGateway"`
	// The hubRoutingPreference of this VirtualHub.
	HubRoutingPreference *string `pulumi:"hubRoutingPreference"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The P2SVpnGateway associated with this VirtualHub.
	P2SVpnGateway *SubResource `pulumi:"p2SVpnGateway"`
	// The preferred gateway to route on-prem traffic
	PreferredRoutingGateway *string `pulumi:"preferredRoutingGateway"`
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The routeTable associated with this virtual hub.
	RouteTable *VirtualHubRouteTable `pulumi:"routeTable"`
	// The securityPartnerProvider associated with this VirtualHub.
	SecurityPartnerProvider *SubResource `pulumi:"securityPartnerProvider"`
	// The Security Provider name.
	SecurityProviderName *string `pulumi:"securityProviderName"`
	// The sku of this VirtualHub.
	Sku *string `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the VirtualHub.
	VirtualHubName *string `pulumi:"virtualHubName"`
	// List of all virtual hub route table v2s associated with this VirtualHub.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	VirtualHubRouteTableV2s []VirtualHubRouteTableV2Type `pulumi:"virtualHubRouteTableV2s"`
	// VirtualRouter ASN.
	VirtualRouterAsn *float64 `pulumi:"virtualRouterAsn"`
	// The VirtualHub Router autoscale configuration.
	VirtualRouterAutoScaleConfiguration *VirtualRouterAutoScaleConfiguration `pulumi:"virtualRouterAutoScaleConfiguration"`
	// VirtualRouter IPs.
	VirtualRouterIps []string `pulumi:"virtualRouterIps"`
	// The VirtualWAN to which the VirtualHub belongs.
	VirtualWan *SubResource `pulumi:"virtualWan"`
	// The VpnGateway associated with this VirtualHub.
	VpnGateway *SubResource `pulumi:"vpnGateway"`
}

// The set of arguments for constructing a VirtualHub resource.
type VirtualHubArgs struct {
	// Address-prefix for this VirtualHub.
	AddressPrefix pulumi.StringPtrInput
	// Flag to control transit for VirtualRouter hub.
	AllowBranchToBranchTraffic pulumi.BoolPtrInput
	// The azureFirewall associated with this VirtualHub.
	AzureFirewall SubResourcePtrInput
	// The expressRouteGateway associated with this VirtualHub.
	ExpressRouteGateway SubResourcePtrInput
	// The hubRoutingPreference of this VirtualHub.
	HubRoutingPreference pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The P2SVpnGateway associated with this VirtualHub.
	P2SVpnGateway SubResourcePtrInput
	// The preferred gateway to route on-prem traffic
	PreferredRoutingGateway pulumi.StringPtrInput
	// The resource group name of the VirtualHub.
	ResourceGroupName pulumi.StringInput
	// The routeTable associated with this virtual hub.
	RouteTable VirtualHubRouteTablePtrInput
	// The securityPartnerProvider associated with this VirtualHub.
	SecurityPartnerProvider SubResourcePtrInput
	// The Security Provider name.
	SecurityProviderName pulumi.StringPtrInput
	// The sku of this VirtualHub.
	Sku pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringPtrInput
	// List of all virtual hub route table v2s associated with this VirtualHub.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	VirtualHubRouteTableV2s VirtualHubRouteTableV2TypeArrayInput
	// VirtualRouter ASN.
	VirtualRouterAsn pulumi.Float64PtrInput
	// The VirtualHub Router autoscale configuration.
	VirtualRouterAutoScaleConfiguration VirtualRouterAutoScaleConfigurationPtrInput
	// VirtualRouter IPs.
	VirtualRouterIps pulumi.StringArrayInput
	// The VirtualWAN to which the VirtualHub belongs.
	VirtualWan SubResourcePtrInput
	// The VpnGateway associated with this VirtualHub.
	VpnGateway SubResourcePtrInput
}

func (VirtualHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubArgs)(nil)).Elem()
}

type VirtualHubInput interface {
	pulumi.Input

	ToVirtualHubOutput() VirtualHubOutput
	ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput
}

func (*VirtualHub) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHub)(nil)).Elem()
}

func (i *VirtualHub) ToVirtualHubOutput() VirtualHubOutput {
	return i.ToVirtualHubOutputWithContext(context.Background())
}

func (i *VirtualHub) ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubOutput)
}

type VirtualHubOutput struct{ *pulumi.OutputState }

func (VirtualHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHub)(nil)).Elem()
}

func (o VirtualHubOutput) ToVirtualHubOutput() VirtualHubOutput {
	return o
}

func (o VirtualHubOutput) ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput {
	return o
}

// Address-prefix for this VirtualHub.
func (o VirtualHubOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringPtrOutput { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

// Flag to control transit for VirtualRouter hub.
func (o VirtualHubOutput) AllowBranchToBranchTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.BoolPtrOutput { return v.AllowBranchToBranchTraffic }).(pulumi.BoolPtrOutput)
}

// The azureFirewall associated with this VirtualHub.
func (o VirtualHubOutput) AzureFirewall() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponsePtrOutput { return v.AzureFirewall }).(SubResourceResponsePtrOutput)
}

// List of references to Bgp Connections.
func (o VirtualHubOutput) BgpConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponseArrayOutput { return v.BgpConnections }).(SubResourceResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o VirtualHubOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The expressRouteGateway associated with this VirtualHub.
func (o VirtualHubOutput) ExpressRouteGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponsePtrOutput { return v.ExpressRouteGateway }).(SubResourceResponsePtrOutput)
}

// The hubRoutingPreference of this VirtualHub.
func (o VirtualHubOutput) HubRoutingPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringPtrOutput { return v.HubRoutingPreference }).(pulumi.StringPtrOutput)
}

// List of references to IpConfigurations.
func (o VirtualHubOutput) IpConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponseArrayOutput { return v.IpConfigurations }).(SubResourceResponseArrayOutput)
}

// Kind of service virtual hub. This is metadata used for the Azure portal experience for Route Server.
func (o VirtualHubOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o VirtualHubOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o VirtualHubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The P2SVpnGateway associated with this VirtualHub.
func (o VirtualHubOutput) P2SVpnGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponsePtrOutput { return v.P2SVpnGateway }).(SubResourceResponsePtrOutput)
}

// The preferred gateway to route on-prem traffic
func (o VirtualHubOutput) PreferredRoutingGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringPtrOutput { return v.PreferredRoutingGateway }).(pulumi.StringPtrOutput)
}

// The provisioning state of the virtual hub resource.
func (o VirtualHubOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of references to RouteMaps.
func (o VirtualHubOutput) RouteMaps() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponseArrayOutput { return v.RouteMaps }).(SubResourceResponseArrayOutput)
}

// The routeTable associated with this virtual hub.
func (o VirtualHubOutput) RouteTable() VirtualHubRouteTableResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) VirtualHubRouteTableResponsePtrOutput { return v.RouteTable }).(VirtualHubRouteTableResponsePtrOutput)
}

// The routing state.
func (o VirtualHubOutput) RoutingState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.RoutingState }).(pulumi.StringOutput)
}

// The securityPartnerProvider associated with this VirtualHub.
func (o VirtualHubOutput) SecurityPartnerProvider() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponsePtrOutput { return v.SecurityPartnerProvider }).(SubResourceResponsePtrOutput)
}

// The Security Provider name.
func (o VirtualHubOutput) SecurityProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringPtrOutput { return v.SecurityProviderName }).(pulumi.StringPtrOutput)
}

// The sku of this VirtualHub.
func (o VirtualHubOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringPtrOutput { return v.Sku }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o VirtualHubOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o VirtualHubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// List of all virtual hub route table v2s associated with this VirtualHub.
func (o VirtualHubOutput) VirtualHubRouteTableV2s() VirtualHubRouteTableV2ResponseArrayOutput {
	return o.ApplyT(func(v *VirtualHub) VirtualHubRouteTableV2ResponseArrayOutput { return v.VirtualHubRouteTableV2s }).(VirtualHubRouteTableV2ResponseArrayOutput)
}

// VirtualRouter ASN.
func (o VirtualHubOutput) VirtualRouterAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.Float64PtrOutput { return v.VirtualRouterAsn }).(pulumi.Float64PtrOutput)
}

// The VirtualHub Router autoscale configuration.
func (o VirtualHubOutput) VirtualRouterAutoScaleConfiguration() VirtualRouterAutoScaleConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) VirtualRouterAutoScaleConfigurationResponsePtrOutput {
		return v.VirtualRouterAutoScaleConfiguration
	}).(VirtualRouterAutoScaleConfigurationResponsePtrOutput)
}

// VirtualRouter IPs.
func (o VirtualHubOutput) VirtualRouterIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringArrayOutput { return v.VirtualRouterIps }).(pulumi.StringArrayOutput)
}

// The VirtualWAN to which the VirtualHub belongs.
func (o VirtualHubOutput) VirtualWan() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponsePtrOutput { return v.VirtualWan }).(SubResourceResponsePtrOutput)
}

// The VpnGateway associated with this VirtualHub.
func (o VirtualHubOutput) VpnGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponsePtrOutput { return v.VpnGateway }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualHubOutput{})
}
