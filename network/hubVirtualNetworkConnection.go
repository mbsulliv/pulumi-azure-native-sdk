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

// HubVirtualNetworkConnection Resource.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01
type HubVirtualNetworkConnection struct {
	pulumi.CustomResourceState

	// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
	AllowHubToRemoteVnetTransit pulumi.BoolPtrOutput `pulumi:"allowHubToRemoteVnetTransit"`
	// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
	AllowRemoteVnetToUseHubVnetGateways pulumi.BoolPtrOutput `pulumi:"allowRemoteVnetToUseHubVnetGateways"`
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrOutput `pulumi:"enableInternetSecurity"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning state of the hub virtual network connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Reference to the remote virtual network.
	RemoteVirtualNetwork SubResourceResponsePtrOutput `pulumi:"remoteVirtualNetwork"`
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration RoutingConfigurationResponsePtrOutput `pulumi:"routingConfiguration"`
}

// NewHubVirtualNetworkConnection registers a new resource with the given unique name, arguments, and options.
func NewHubVirtualNetworkConnection(ctx *pulumi.Context,
	name string, args *HubVirtualNetworkConnectionArgs, opts ...pulumi.ResourceOption) (*HubVirtualNetworkConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualHubName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20200501:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:HubVirtualNetworkConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HubVirtualNetworkConnection
	err := ctx.RegisterResource("azure-native:network:HubVirtualNetworkConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHubVirtualNetworkConnection gets an existing HubVirtualNetworkConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHubVirtualNetworkConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HubVirtualNetworkConnectionState, opts ...pulumi.ResourceOption) (*HubVirtualNetworkConnection, error) {
	var resource HubVirtualNetworkConnection
	err := ctx.ReadResource("azure-native:network:HubVirtualNetworkConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HubVirtualNetworkConnection resources.
type hubVirtualNetworkConnectionState struct {
}

type HubVirtualNetworkConnectionState struct {
}

func (HubVirtualNetworkConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*hubVirtualNetworkConnectionState)(nil)).Elem()
}

type hubVirtualNetworkConnectionArgs struct {
	// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
	AllowHubToRemoteVnetTransit *bool `pulumi:"allowHubToRemoteVnetTransit"`
	// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
	AllowRemoteVnetToUseHubVnetGateways *bool `pulumi:"allowRemoteVnetToUseHubVnetGateways"`
	// The name of the HubVirtualNetworkConnection.
	ConnectionName *string `pulumi:"connectionName"`
	// Enable internet security.
	EnableInternetSecurity *bool `pulumi:"enableInternetSecurity"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Reference to the remote virtual network.
	RemoteVirtualNetwork *SubResource `pulumi:"remoteVirtualNetwork"`
	// The resource group name of the HubVirtualNetworkConnection.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration *RoutingConfiguration `pulumi:"routingConfiguration"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// The set of arguments for constructing a HubVirtualNetworkConnection resource.
type HubVirtualNetworkConnectionArgs struct {
	// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
	AllowHubToRemoteVnetTransit pulumi.BoolPtrInput
	// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
	AllowRemoteVnetToUseHubVnetGateways pulumi.BoolPtrInput
	// The name of the HubVirtualNetworkConnection.
	ConnectionName pulumi.StringPtrInput
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Reference to the remote virtual network.
	RemoteVirtualNetwork SubResourcePtrInput
	// The resource group name of the HubVirtualNetworkConnection.
	ResourceGroupName pulumi.StringInput
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration RoutingConfigurationPtrInput
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput
}

func (HubVirtualNetworkConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hubVirtualNetworkConnectionArgs)(nil)).Elem()
}

type HubVirtualNetworkConnectionInput interface {
	pulumi.Input

	ToHubVirtualNetworkConnectionOutput() HubVirtualNetworkConnectionOutput
	ToHubVirtualNetworkConnectionOutputWithContext(ctx context.Context) HubVirtualNetworkConnectionOutput
}

func (*HubVirtualNetworkConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**HubVirtualNetworkConnection)(nil)).Elem()
}

func (i *HubVirtualNetworkConnection) ToHubVirtualNetworkConnectionOutput() HubVirtualNetworkConnectionOutput {
	return i.ToHubVirtualNetworkConnectionOutputWithContext(context.Background())
}

func (i *HubVirtualNetworkConnection) ToHubVirtualNetworkConnectionOutputWithContext(ctx context.Context) HubVirtualNetworkConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubVirtualNetworkConnectionOutput)
}

func (i *HubVirtualNetworkConnection) ToOutput(ctx context.Context) pulumix.Output[*HubVirtualNetworkConnection] {
	return pulumix.Output[*HubVirtualNetworkConnection]{
		OutputState: i.ToHubVirtualNetworkConnectionOutputWithContext(ctx).OutputState,
	}
}

type HubVirtualNetworkConnectionOutput struct{ *pulumi.OutputState }

func (HubVirtualNetworkConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HubVirtualNetworkConnection)(nil)).Elem()
}

func (o HubVirtualNetworkConnectionOutput) ToHubVirtualNetworkConnectionOutput() HubVirtualNetworkConnectionOutput {
	return o
}

func (o HubVirtualNetworkConnectionOutput) ToHubVirtualNetworkConnectionOutputWithContext(ctx context.Context) HubVirtualNetworkConnectionOutput {
	return o
}

func (o HubVirtualNetworkConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*HubVirtualNetworkConnection] {
	return pulumix.Output[*HubVirtualNetworkConnection]{
		OutputState: o.OutputState,
	}
}

// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
func (o HubVirtualNetworkConnectionOutput) AllowHubToRemoteVnetTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) pulumi.BoolPtrOutput { return v.AllowHubToRemoteVnetTransit }).(pulumi.BoolPtrOutput)
}

// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
func (o HubVirtualNetworkConnectionOutput) AllowRemoteVnetToUseHubVnetGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) pulumi.BoolPtrOutput {
		return v.AllowRemoteVnetToUseHubVnetGateways
	}).(pulumi.BoolPtrOutput)
}

// Enable internet security.
func (o HubVirtualNetworkConnectionOutput) EnableInternetSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) pulumi.BoolPtrOutput { return v.EnableInternetSecurity }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o HubVirtualNetworkConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o HubVirtualNetworkConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the hub virtual network connection resource.
func (o HubVirtualNetworkConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Reference to the remote virtual network.
func (o HubVirtualNetworkConnectionOutput) RemoteVirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) SubResourceResponsePtrOutput { return v.RemoteVirtualNetwork }).(SubResourceResponsePtrOutput)
}

// The Routing Configuration indicating the associated and propagated route tables on this connection.
func (o HubVirtualNetworkConnectionOutput) RoutingConfiguration() RoutingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) RoutingConfigurationResponsePtrOutput {
		return v.RoutingConfiguration
	}).(RoutingConfigurationResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HubVirtualNetworkConnectionOutput{})
}
