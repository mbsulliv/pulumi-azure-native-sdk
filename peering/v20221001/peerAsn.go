// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The essential information related to the peer's ASN.
type PeerAsn struct {
	pulumi.CustomResourceState

	// The error message for the validation state
	ErrorMessage pulumi.StringOutput `pulumi:"errorMessage"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Autonomous System Number (ASN) of the peer.
	PeerAsn pulumi.IntPtrOutput `pulumi:"peerAsn"`
	// The contact details of the peer.
	PeerContactDetail ContactDetailResponseArrayOutput `pulumi:"peerContactDetail"`
	// The name of the peer.
	PeerName pulumi.StringPtrOutput `pulumi:"peerName"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The validation state of the ASN associated with the peer.
	ValidationState pulumi.StringOutput `pulumi:"validationState"`
}

// NewPeerAsn registers a new resource with the given unique name, arguments, and options.
func NewPeerAsn(ctx *pulumi.Context,
	name string, args *PeerAsnArgs, opts ...pulumi.ResourceOption) (*PeerAsn, error) {
	if args == nil {
		args = &PeerAsnArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:peering:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190801preview:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190901preview:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200101preview:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200401:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20201001:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210101:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210601:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220101:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220601:PeerAsn"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PeerAsn
	err := ctx.RegisterResource("azure-native:peering/v20221001:PeerAsn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeerAsn gets an existing PeerAsn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeerAsn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeerAsnState, opts ...pulumi.ResourceOption) (*PeerAsn, error) {
	var resource PeerAsn
	err := ctx.ReadResource("azure-native:peering/v20221001:PeerAsn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PeerAsn resources.
type peerAsnState struct {
}

type PeerAsnState struct {
}

func (PeerAsnState) ElementType() reflect.Type {
	return reflect.TypeOf((*peerAsnState)(nil)).Elem()
}

type peerAsnArgs struct {
	// The Autonomous System Number (ASN) of the peer.
	PeerAsn *int `pulumi:"peerAsn"`
	// The peer ASN name.
	PeerAsnName *string `pulumi:"peerAsnName"`
	// The contact details of the peer.
	PeerContactDetail []ContactDetail `pulumi:"peerContactDetail"`
	// The name of the peer.
	PeerName *string `pulumi:"peerName"`
}

// The set of arguments for constructing a PeerAsn resource.
type PeerAsnArgs struct {
	// The Autonomous System Number (ASN) of the peer.
	PeerAsn pulumi.IntPtrInput
	// The peer ASN name.
	PeerAsnName pulumi.StringPtrInput
	// The contact details of the peer.
	PeerContactDetail ContactDetailArrayInput
	// The name of the peer.
	PeerName pulumi.StringPtrInput
}

func (PeerAsnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peerAsnArgs)(nil)).Elem()
}

type PeerAsnInput interface {
	pulumi.Input

	ToPeerAsnOutput() PeerAsnOutput
	ToPeerAsnOutputWithContext(ctx context.Context) PeerAsnOutput
}

func (*PeerAsn) ElementType() reflect.Type {
	return reflect.TypeOf((**PeerAsn)(nil)).Elem()
}

func (i *PeerAsn) ToPeerAsnOutput() PeerAsnOutput {
	return i.ToPeerAsnOutputWithContext(context.Background())
}

func (i *PeerAsn) ToPeerAsnOutputWithContext(ctx context.Context) PeerAsnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeerAsnOutput)
}

func (i *PeerAsn) ToOutput(ctx context.Context) pulumix.Output[*PeerAsn] {
	return pulumix.Output[*PeerAsn]{
		OutputState: i.ToPeerAsnOutputWithContext(ctx).OutputState,
	}
}

type PeerAsnOutput struct{ *pulumi.OutputState }

func (PeerAsnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeerAsn)(nil)).Elem()
}

func (o PeerAsnOutput) ToPeerAsnOutput() PeerAsnOutput {
	return o
}

func (o PeerAsnOutput) ToPeerAsnOutputWithContext(ctx context.Context) PeerAsnOutput {
	return o
}

func (o PeerAsnOutput) ToOutput(ctx context.Context) pulumix.Output[*PeerAsn] {
	return pulumix.Output[*PeerAsn]{
		OutputState: o.OutputState,
	}
}

// The error message for the validation state
func (o PeerAsnOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerAsn) pulumi.StringOutput { return v.ErrorMessage }).(pulumi.StringOutput)
}

// The name of the resource.
func (o PeerAsnOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerAsn) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Autonomous System Number (ASN) of the peer.
func (o PeerAsnOutput) PeerAsn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PeerAsn) pulumi.IntPtrOutput { return v.PeerAsn }).(pulumi.IntPtrOutput)
}

// The contact details of the peer.
func (o PeerAsnOutput) PeerContactDetail() ContactDetailResponseArrayOutput {
	return o.ApplyT(func(v *PeerAsn) ContactDetailResponseArrayOutput { return v.PeerContactDetail }).(ContactDetailResponseArrayOutput)
}

// The name of the peer.
func (o PeerAsnOutput) PeerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeerAsn) pulumi.StringPtrOutput { return v.PeerName }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o PeerAsnOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerAsn) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The validation state of the ASN associated with the peer.
func (o PeerAsnOutput) ValidationState() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerAsn) pulumi.StringOutput { return v.ValidationState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PeerAsnOutput{})
}
