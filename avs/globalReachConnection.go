// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package avs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A global reach connection resource
// Azure REST API version: 2022-05-01. Prior API version in Azure Native 1.x: 2020-07-17-preview.
//
// Other available API versions: 2023-03-01.
type GlobalReachConnection struct {
	pulumi.CustomResourceState

	// The network used for global reach carved out from the original network block provided for the private cloud
	AddressPrefix pulumi.StringOutput `pulumi:"addressPrefix"`
	// Authorization key from the peer express route used for the global reach connection
	AuthorizationKey pulumi.StringPtrOutput `pulumi:"authorizationKey"`
	// The connection status of the global reach connection
	CircuitConnectionStatus pulumi.StringOutput `pulumi:"circuitConnectionStatus"`
	// The ID of the Private Cloud's ExpressRoute Circuit that is participating in the global reach connection
	ExpressRouteId pulumi.StringPtrOutput `pulumi:"expressRouteId"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Identifier of the ExpressRoute Circuit to peer with in the global reach connection
	PeerExpressRouteCircuit pulumi.StringPtrOutput `pulumi:"peerExpressRouteCircuit"`
	// The state of the  ExpressRoute Circuit Authorization provisioning
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGlobalReachConnection registers a new resource with the given unique name, arguments, and options.
func NewGlobalReachConnection(ctx *pulumi.Context,
	name string, args *GlobalReachConnectionArgs, opts ...pulumi.ResourceOption) (*GlobalReachConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:GlobalReachConnection"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:GlobalReachConnection"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:GlobalReachConnection"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:GlobalReachConnection"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:GlobalReachConnection"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230301:GlobalReachConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GlobalReachConnection
	err := ctx.RegisterResource("azure-native:avs:GlobalReachConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalReachConnection gets an existing GlobalReachConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalReachConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalReachConnectionState, opts ...pulumi.ResourceOption) (*GlobalReachConnection, error) {
	var resource GlobalReachConnection
	err := ctx.ReadResource("azure-native:avs:GlobalReachConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalReachConnection resources.
type globalReachConnectionState struct {
}

type GlobalReachConnectionState struct {
}

func (GlobalReachConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalReachConnectionState)(nil)).Elem()
}

type globalReachConnectionArgs struct {
	// Authorization key from the peer express route used for the global reach connection
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The ID of the Private Cloud's ExpressRoute Circuit that is participating in the global reach connection
	ExpressRouteId *string `pulumi:"expressRouteId"`
	// Name of the global reach connection in the private cloud
	GlobalReachConnectionName *string `pulumi:"globalReachConnectionName"`
	// Identifier of the ExpressRoute Circuit to peer with in the global reach connection
	PeerExpressRouteCircuit *string `pulumi:"peerExpressRouteCircuit"`
	// The name of the private cloud.
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a GlobalReachConnection resource.
type GlobalReachConnectionArgs struct {
	// Authorization key from the peer express route used for the global reach connection
	AuthorizationKey pulumi.StringPtrInput
	// The ID of the Private Cloud's ExpressRoute Circuit that is participating in the global reach connection
	ExpressRouteId pulumi.StringPtrInput
	// Name of the global reach connection in the private cloud
	GlobalReachConnectionName pulumi.StringPtrInput
	// Identifier of the ExpressRoute Circuit to peer with in the global reach connection
	PeerExpressRouteCircuit pulumi.StringPtrInput
	// The name of the private cloud.
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (GlobalReachConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalReachConnectionArgs)(nil)).Elem()
}

type GlobalReachConnectionInput interface {
	pulumi.Input

	ToGlobalReachConnectionOutput() GlobalReachConnectionOutput
	ToGlobalReachConnectionOutputWithContext(ctx context.Context) GlobalReachConnectionOutput
}

func (*GlobalReachConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalReachConnection)(nil)).Elem()
}

func (i *GlobalReachConnection) ToGlobalReachConnectionOutput() GlobalReachConnectionOutput {
	return i.ToGlobalReachConnectionOutputWithContext(context.Background())
}

func (i *GlobalReachConnection) ToGlobalReachConnectionOutputWithContext(ctx context.Context) GlobalReachConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalReachConnectionOutput)
}

type GlobalReachConnectionOutput struct{ *pulumi.OutputState }

func (GlobalReachConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalReachConnection)(nil)).Elem()
}

func (o GlobalReachConnectionOutput) ToGlobalReachConnectionOutput() GlobalReachConnectionOutput {
	return o
}

func (o GlobalReachConnectionOutput) ToGlobalReachConnectionOutputWithContext(ctx context.Context) GlobalReachConnectionOutput {
	return o
}

// The network used for global reach carved out from the original network block provided for the private cloud
func (o GlobalReachConnectionOutput) AddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalReachConnection) pulumi.StringOutput { return v.AddressPrefix }).(pulumi.StringOutput)
}

// Authorization key from the peer express route used for the global reach connection
func (o GlobalReachConnectionOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalReachConnection) pulumi.StringPtrOutput { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

// The connection status of the global reach connection
func (o GlobalReachConnectionOutput) CircuitConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalReachConnection) pulumi.StringOutput { return v.CircuitConnectionStatus }).(pulumi.StringOutput)
}

// The ID of the Private Cloud's ExpressRoute Circuit that is participating in the global reach connection
func (o GlobalReachConnectionOutput) ExpressRouteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalReachConnection) pulumi.StringPtrOutput { return v.ExpressRouteId }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o GlobalReachConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalReachConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Identifier of the ExpressRoute Circuit to peer with in the global reach connection
func (o GlobalReachConnectionOutput) PeerExpressRouteCircuit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalReachConnection) pulumi.StringPtrOutput { return v.PeerExpressRouteCircuit }).(pulumi.StringPtrOutput)
}

// The state of the  ExpressRoute Circuit Authorization provisioning
func (o GlobalReachConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalReachConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o GlobalReachConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalReachConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GlobalReachConnectionOutput{})
}
