// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Virtual Appliance Site resource.
type VirtualHubBgpConnection struct {
	pulumi.CustomResourceState

	// The current state of the VirtualHub to Peer.
	ConnectionState pulumi.StringOutput `pulumi:"connectionState"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The reference to the HubVirtualNetworkConnection resource.
	HubVirtualNetworkConnection SubResourceResponsePtrOutput `pulumi:"hubVirtualNetworkConnection"`
	// Name of the connection.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Peer ASN.
	PeerAsn pulumi.Float64PtrOutput `pulumi:"peerAsn"`
	// Peer IP.
	PeerIp pulumi.StringPtrOutput `pulumi:"peerIp"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Connection type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualHubBgpConnection registers a new resource with the given unique name, arguments, and options.
func NewVirtualHubBgpConnection(ctx *pulumi.Context,
	name string, args *VirtualHubBgpConnectionArgs, opts ...pulumi.ResourceOption) (*VirtualHubBgpConnection, error) {
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
			Type: pulumi.String("azure-native:network:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:VirtualHubBgpConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualHubBgpConnection
	err := ctx.RegisterResource("azure-native:network/v20230401:VirtualHubBgpConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHubBgpConnection gets an existing VirtualHubBgpConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHubBgpConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubBgpConnectionState, opts ...pulumi.ResourceOption) (*VirtualHubBgpConnection, error) {
	var resource VirtualHubBgpConnection
	err := ctx.ReadResource("azure-native:network/v20230401:VirtualHubBgpConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHubBgpConnection resources.
type virtualHubBgpConnectionState struct {
}

type VirtualHubBgpConnectionState struct {
}

func (VirtualHubBgpConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubBgpConnectionState)(nil)).Elem()
}

type virtualHubBgpConnectionArgs struct {
	// The name of the connection.
	ConnectionName *string `pulumi:"connectionName"`
	// The reference to the HubVirtualNetworkConnection resource.
	HubVirtualNetworkConnection *SubResource `pulumi:"hubVirtualNetworkConnection"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Name of the connection.
	Name *string `pulumi:"name"`
	// Peer ASN.
	PeerAsn *float64 `pulumi:"peerAsn"`
	// Peer IP.
	PeerIp *string `pulumi:"peerIp"`
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// The set of arguments for constructing a VirtualHubBgpConnection resource.
type VirtualHubBgpConnectionArgs struct {
	// The name of the connection.
	ConnectionName pulumi.StringPtrInput
	// The reference to the HubVirtualNetworkConnection resource.
	HubVirtualNetworkConnection SubResourcePtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Name of the connection.
	Name pulumi.StringPtrInput
	// Peer ASN.
	PeerAsn pulumi.Float64PtrInput
	// Peer IP.
	PeerIp pulumi.StringPtrInput
	// The resource group name of the VirtualHub.
	ResourceGroupName pulumi.StringInput
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput
}

func (VirtualHubBgpConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubBgpConnectionArgs)(nil)).Elem()
}

type VirtualHubBgpConnectionInput interface {
	pulumi.Input

	ToVirtualHubBgpConnectionOutput() VirtualHubBgpConnectionOutput
	ToVirtualHubBgpConnectionOutputWithContext(ctx context.Context) VirtualHubBgpConnectionOutput
}

func (*VirtualHubBgpConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubBgpConnection)(nil)).Elem()
}

func (i *VirtualHubBgpConnection) ToVirtualHubBgpConnectionOutput() VirtualHubBgpConnectionOutput {
	return i.ToVirtualHubBgpConnectionOutputWithContext(context.Background())
}

func (i *VirtualHubBgpConnection) ToVirtualHubBgpConnectionOutputWithContext(ctx context.Context) VirtualHubBgpConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubBgpConnectionOutput)
}

func (i *VirtualHubBgpConnection) ToOutput(ctx context.Context) pulumix.Output[*VirtualHubBgpConnection] {
	return pulumix.Output[*VirtualHubBgpConnection]{
		OutputState: i.ToVirtualHubBgpConnectionOutputWithContext(ctx).OutputState,
	}
}

type VirtualHubBgpConnectionOutput struct{ *pulumi.OutputState }

func (VirtualHubBgpConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubBgpConnection)(nil)).Elem()
}

func (o VirtualHubBgpConnectionOutput) ToVirtualHubBgpConnectionOutput() VirtualHubBgpConnectionOutput {
	return o
}

func (o VirtualHubBgpConnectionOutput) ToVirtualHubBgpConnectionOutputWithContext(ctx context.Context) VirtualHubBgpConnectionOutput {
	return o
}

func (o VirtualHubBgpConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualHubBgpConnection] {
	return pulumix.Output[*VirtualHubBgpConnection]{
		OutputState: o.OutputState,
	}
}

// The current state of the VirtualHub to Peer.
func (o VirtualHubBgpConnectionOutput) ConnectionState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.StringOutput { return v.ConnectionState }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o VirtualHubBgpConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The reference to the HubVirtualNetworkConnection resource.
func (o VirtualHubBgpConnectionOutput) HubVirtualNetworkConnection() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) SubResourceResponsePtrOutput { return v.HubVirtualNetworkConnection }).(SubResourceResponsePtrOutput)
}

// Name of the connection.
func (o VirtualHubBgpConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Peer ASN.
func (o VirtualHubBgpConnectionOutput) PeerAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.Float64PtrOutput { return v.PeerAsn }).(pulumi.Float64PtrOutput)
}

// Peer IP.
func (o VirtualHubBgpConnectionOutput) PeerIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.StringPtrOutput { return v.PeerIp }).(pulumi.StringPtrOutput)
}

// The provisioning state of the resource.
func (o VirtualHubBgpConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Connection type.
func (o VirtualHubBgpConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualHubBgpConnectionOutput{})
}
