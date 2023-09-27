// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Virtual Router Peering resource.
type VirtualRouterPeering struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Name of the virtual router peering that is unique within a virtual router.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Peer ASN.
	PeerAsn pulumi.Float64PtrOutput `pulumi:"peerAsn"`
	// Peer IP.
	PeerIp pulumi.StringPtrOutput `pulumi:"peerIp"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Peering type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualRouterPeering registers a new resource with the given unique name, arguments, and options.
func NewVirtualRouterPeering(ctx *pulumi.Context,
	name string, args *VirtualRouterPeeringArgs, opts ...pulumi.ResourceOption) (*VirtualRouterPeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualRouterName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualRouterName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:VirtualRouterPeering"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualRouterPeering
	err := ctx.RegisterResource("azure-native:network/v20230501:VirtualRouterPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualRouterPeering gets an existing VirtualRouterPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualRouterPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualRouterPeeringState, opts ...pulumi.ResourceOption) (*VirtualRouterPeering, error) {
	var resource VirtualRouterPeering
	err := ctx.ReadResource("azure-native:network/v20230501:VirtualRouterPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualRouterPeering resources.
type virtualRouterPeeringState struct {
}

type VirtualRouterPeeringState struct {
}

func (VirtualRouterPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRouterPeeringState)(nil)).Elem()
}

type virtualRouterPeeringArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// Name of the virtual router peering that is unique within a virtual router.
	Name *string `pulumi:"name"`
	// Peer ASN.
	PeerAsn *float64 `pulumi:"peerAsn"`
	// Peer IP.
	PeerIp *string `pulumi:"peerIp"`
	// The name of the Virtual Router Peering.
	PeeringName *string `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Virtual Router.
	VirtualRouterName string `pulumi:"virtualRouterName"`
}

// The set of arguments for constructing a VirtualRouterPeering resource.
type VirtualRouterPeeringArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// Name of the virtual router peering that is unique within a virtual router.
	Name pulumi.StringPtrInput
	// Peer ASN.
	PeerAsn pulumi.Float64PtrInput
	// Peer IP.
	PeerIp pulumi.StringPtrInput
	// The name of the Virtual Router Peering.
	PeeringName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the Virtual Router.
	VirtualRouterName pulumi.StringInput
}

func (VirtualRouterPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRouterPeeringArgs)(nil)).Elem()
}

type VirtualRouterPeeringInput interface {
	pulumi.Input

	ToVirtualRouterPeeringOutput() VirtualRouterPeeringOutput
	ToVirtualRouterPeeringOutputWithContext(ctx context.Context) VirtualRouterPeeringOutput
}

func (*VirtualRouterPeering) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualRouterPeering)(nil)).Elem()
}

func (i *VirtualRouterPeering) ToVirtualRouterPeeringOutput() VirtualRouterPeeringOutput {
	return i.ToVirtualRouterPeeringOutputWithContext(context.Background())
}

func (i *VirtualRouterPeering) ToVirtualRouterPeeringOutputWithContext(ctx context.Context) VirtualRouterPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualRouterPeeringOutput)
}

func (i *VirtualRouterPeering) ToOutput(ctx context.Context) pulumix.Output[*VirtualRouterPeering] {
	return pulumix.Output[*VirtualRouterPeering]{
		OutputState: i.ToVirtualRouterPeeringOutputWithContext(ctx).OutputState,
	}
}

type VirtualRouterPeeringOutput struct{ *pulumi.OutputState }

func (VirtualRouterPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualRouterPeering)(nil)).Elem()
}

func (o VirtualRouterPeeringOutput) ToVirtualRouterPeeringOutput() VirtualRouterPeeringOutput {
	return o
}

func (o VirtualRouterPeeringOutput) ToVirtualRouterPeeringOutputWithContext(ctx context.Context) VirtualRouterPeeringOutput {
	return o
}

func (o VirtualRouterPeeringOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualRouterPeering] {
	return pulumix.Output[*VirtualRouterPeering]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o VirtualRouterPeeringOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRouterPeering) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Name of the virtual router peering that is unique within a virtual router.
func (o VirtualRouterPeeringOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRouterPeering) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Peer ASN.
func (o VirtualRouterPeeringOutput) PeerAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualRouterPeering) pulumi.Float64PtrOutput { return v.PeerAsn }).(pulumi.Float64PtrOutput)
}

// Peer IP.
func (o VirtualRouterPeeringOutput) PeerIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRouterPeering) pulumi.StringPtrOutput { return v.PeerIp }).(pulumi.StringPtrOutput)
}

// The provisioning state of the resource.
func (o VirtualRouterPeeringOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRouterPeering) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Peering type.
func (o VirtualRouterPeeringOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRouterPeering) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualRouterPeeringOutput{})
}
