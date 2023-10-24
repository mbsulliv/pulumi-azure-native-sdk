// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// NetworkVirtualApplianceConnection resource.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01.
type NetworkVirtualApplianceConnection struct {
	pulumi.CustomResourceState

	// Network Virtual Appliance ASN.
	Asn pulumi.Float64PtrOutput `pulumi:"asn"`
	// List of bgpPeerAddresses for the NVA instances
	BgpPeerAddress pulumi.StringArrayOutput `pulumi:"bgpPeerAddress"`
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrOutput `pulumi:"enableInternetSecurity"`
	// The name of the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning state of the NetworkVirtualApplianceConnection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration RoutingConfigurationNfvResponsePtrOutput `pulumi:"routingConfiguration"`
	// Unique identifier for the connection.
	TunnelIdentifier pulumi.Float64PtrOutput `pulumi:"tunnelIdentifier"`
}

// NewNetworkVirtualApplianceConnection registers a new resource with the given unique name, arguments, and options.
func NewNetworkVirtualApplianceConnection(ctx *pulumi.Context,
	name string, args *NetworkVirtualApplianceConnectionArgs, opts ...pulumi.ResourceOption) (*NetworkVirtualApplianceConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkVirtualApplianceName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkVirtualApplianceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20221101:NetworkVirtualApplianceConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:NetworkVirtualApplianceConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:NetworkVirtualApplianceConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:NetworkVirtualApplianceConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NetworkVirtualApplianceConnection
	err := ctx.RegisterResource("azure-native:network:NetworkVirtualApplianceConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkVirtualApplianceConnection gets an existing NetworkVirtualApplianceConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkVirtualApplianceConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkVirtualApplianceConnectionState, opts ...pulumi.ResourceOption) (*NetworkVirtualApplianceConnection, error) {
	var resource NetworkVirtualApplianceConnection
	err := ctx.ReadResource("azure-native:network:NetworkVirtualApplianceConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkVirtualApplianceConnection resources.
type networkVirtualApplianceConnectionState struct {
}

type NetworkVirtualApplianceConnectionState struct {
}

func (NetworkVirtualApplianceConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceConnectionState)(nil)).Elem()
}

type networkVirtualApplianceConnectionArgs struct {
	// Network Virtual Appliance ASN.
	Asn *float64 `pulumi:"asn"`
	// List of bgpPeerAddresses for the NVA instances
	BgpPeerAddress []string `pulumi:"bgpPeerAddress"`
	// The name of the NVA connection.
	ConnectionName *string `pulumi:"connectionName"`
	// Enable internet security.
	EnableInternetSecurity *bool `pulumi:"enableInternetSecurity"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The name of the Network Virtual Appliance.
	NetworkVirtualApplianceName string `pulumi:"networkVirtualApplianceName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration *RoutingConfigurationNfv `pulumi:"routingConfiguration"`
	// Unique identifier for the connection.
	TunnelIdentifier *float64 `pulumi:"tunnelIdentifier"`
}

// The set of arguments for constructing a NetworkVirtualApplianceConnection resource.
type NetworkVirtualApplianceConnectionArgs struct {
	// Network Virtual Appliance ASN.
	Asn pulumi.Float64PtrInput
	// List of bgpPeerAddresses for the NVA instances
	BgpPeerAddress pulumi.StringArrayInput
	// The name of the NVA connection.
	ConnectionName pulumi.StringPtrInput
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The name of the Network Virtual Appliance.
	NetworkVirtualApplianceName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration RoutingConfigurationNfvPtrInput
	// Unique identifier for the connection.
	TunnelIdentifier pulumi.Float64PtrInput
}

func (NetworkVirtualApplianceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceConnectionArgs)(nil)).Elem()
}

type NetworkVirtualApplianceConnectionInput interface {
	pulumi.Input

	ToNetworkVirtualApplianceConnectionOutput() NetworkVirtualApplianceConnectionOutput
	ToNetworkVirtualApplianceConnectionOutputWithContext(ctx context.Context) NetworkVirtualApplianceConnectionOutput
}

func (*NetworkVirtualApplianceConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkVirtualApplianceConnection)(nil)).Elem()
}

func (i *NetworkVirtualApplianceConnection) ToNetworkVirtualApplianceConnectionOutput() NetworkVirtualApplianceConnectionOutput {
	return i.ToNetworkVirtualApplianceConnectionOutputWithContext(context.Background())
}

func (i *NetworkVirtualApplianceConnection) ToNetworkVirtualApplianceConnectionOutputWithContext(ctx context.Context) NetworkVirtualApplianceConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkVirtualApplianceConnectionOutput)
}

func (i *NetworkVirtualApplianceConnection) ToOutput(ctx context.Context) pulumix.Output[*NetworkVirtualApplianceConnection] {
	return pulumix.Output[*NetworkVirtualApplianceConnection]{
		OutputState: i.ToNetworkVirtualApplianceConnectionOutputWithContext(ctx).OutputState,
	}
}

type NetworkVirtualApplianceConnectionOutput struct{ *pulumi.OutputState }

func (NetworkVirtualApplianceConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkVirtualApplianceConnection)(nil)).Elem()
}

func (o NetworkVirtualApplianceConnectionOutput) ToNetworkVirtualApplianceConnectionOutput() NetworkVirtualApplianceConnectionOutput {
	return o
}

func (o NetworkVirtualApplianceConnectionOutput) ToNetworkVirtualApplianceConnectionOutputWithContext(ctx context.Context) NetworkVirtualApplianceConnectionOutput {
	return o
}

func (o NetworkVirtualApplianceConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*NetworkVirtualApplianceConnection] {
	return pulumix.Output[*NetworkVirtualApplianceConnection]{
		OutputState: o.OutputState,
	}
}

// Network Virtual Appliance ASN.
func (o NetworkVirtualApplianceConnectionOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *NetworkVirtualApplianceConnection) pulumi.Float64PtrOutput { return v.Asn }).(pulumi.Float64PtrOutput)
}

// List of bgpPeerAddresses for the NVA instances
func (o NetworkVirtualApplianceConnectionOutput) BgpPeerAddress() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkVirtualApplianceConnection) pulumi.StringArrayOutput { return v.BgpPeerAddress }).(pulumi.StringArrayOutput)
}

// Enable internet security.
func (o NetworkVirtualApplianceConnectionOutput) EnableInternetSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkVirtualApplianceConnection) pulumi.BoolPtrOutput { return v.EnableInternetSecurity }).(pulumi.BoolPtrOutput)
}

// The name of the resource.
func (o NetworkVirtualApplianceConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkVirtualApplianceConnection) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the NetworkVirtualApplianceConnection resource.
func (o NetworkVirtualApplianceConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVirtualApplianceConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The Routing Configuration indicating the associated and propagated route tables on this connection.
func (o NetworkVirtualApplianceConnectionOutput) RoutingConfiguration() RoutingConfigurationNfvResponsePtrOutput {
	return o.ApplyT(func(v *NetworkVirtualApplianceConnection) RoutingConfigurationNfvResponsePtrOutput {
		return v.RoutingConfiguration
	}).(RoutingConfigurationNfvResponsePtrOutput)
}

// Unique identifier for the connection.
func (o NetworkVirtualApplianceConnectionOutput) TunnelIdentifier() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *NetworkVirtualApplianceConnection) pulumi.Float64PtrOutput { return v.TunnelIdentifier }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkVirtualApplianceConnectionOutput{})
}
