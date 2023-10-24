// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package peering

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Peering is a logical representation of a set of connections to the Microsoft Cloud Edge at a location.
// Azure REST API version: 2022-10-01. Prior API version in Azure Native 1.x: 2021-01-01.
type Peering struct {
	pulumi.CustomResourceState

	// The properties that define a direct peering.
	Direct PeeringPropertiesDirectResponsePtrOutput `pulumi:"direct"`
	// The properties that define an exchange peering.
	Exchange PeeringPropertiesExchangeResponsePtrOutput `pulumi:"exchange"`
	// The kind of the peering.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The location of the peering.
	PeeringLocation pulumi.StringPtrOutput `pulumi:"peeringLocation"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The SKU that defines the tier and kind of the peering.
	Sku PeeringSkuResponseOutput `pulumi:"sku"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPeering registers a new resource with the given unique name, arguments, and options.
func NewPeering(ctx *pulumi.Context,
	name string, args *PeeringArgs, opts ...pulumi.ResourceOption) (*Peering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:peering/v20190801preview:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190901preview:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200101preview:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200401:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20201001:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210101:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210601:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220101:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220601:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20221001:Peering"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Peering
	err := ctx.RegisterResource("azure-native:peering:Peering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeering gets an existing Peering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringState, opts ...pulumi.ResourceOption) (*Peering, error) {
	var resource Peering
	err := ctx.ReadResource("azure-native:peering:Peering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Peering resources.
type peeringState struct {
}

type PeeringState struct {
}

func (PeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringState)(nil)).Elem()
}

type peeringArgs struct {
	// The properties that define a direct peering.
	Direct *PeeringPropertiesDirect `pulumi:"direct"`
	// The properties that define an exchange peering.
	Exchange *PeeringPropertiesExchange `pulumi:"exchange"`
	// The kind of the peering.
	Kind string `pulumi:"kind"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The location of the peering.
	PeeringLocation *string `pulumi:"peeringLocation"`
	// The name of the peering.
	PeeringName *string `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU that defines the tier and kind of the peering.
	Sku PeeringSku `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Peering resource.
type PeeringArgs struct {
	// The properties that define a direct peering.
	Direct PeeringPropertiesDirectPtrInput
	// The properties that define an exchange peering.
	Exchange PeeringPropertiesExchangePtrInput
	// The kind of the peering.
	Kind pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The location of the peering.
	PeeringLocation pulumi.StringPtrInput
	// The name of the peering.
	PeeringName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The SKU that defines the tier and kind of the peering.
	Sku PeeringSkuInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (PeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringArgs)(nil)).Elem()
}

type PeeringInput interface {
	pulumi.Input

	ToPeeringOutput() PeeringOutput
	ToPeeringOutputWithContext(ctx context.Context) PeeringOutput
}

func (*Peering) ElementType() reflect.Type {
	return reflect.TypeOf((**Peering)(nil)).Elem()
}

func (i *Peering) ToPeeringOutput() PeeringOutput {
	return i.ToPeeringOutputWithContext(context.Background())
}

func (i *Peering) ToPeeringOutputWithContext(ctx context.Context) PeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringOutput)
}

func (i *Peering) ToOutput(ctx context.Context) pulumix.Output[*Peering] {
	return pulumix.Output[*Peering]{
		OutputState: i.ToPeeringOutputWithContext(ctx).OutputState,
	}
}

type PeeringOutput struct{ *pulumi.OutputState }

func (PeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Peering)(nil)).Elem()
}

func (o PeeringOutput) ToPeeringOutput() PeeringOutput {
	return o
}

func (o PeeringOutput) ToPeeringOutputWithContext(ctx context.Context) PeeringOutput {
	return o
}

func (o PeeringOutput) ToOutput(ctx context.Context) pulumix.Output[*Peering] {
	return pulumix.Output[*Peering]{
		OutputState: o.OutputState,
	}
}

// The properties that define a direct peering.
func (o PeeringOutput) Direct() PeeringPropertiesDirectResponsePtrOutput {
	return o.ApplyT(func(v *Peering) PeeringPropertiesDirectResponsePtrOutput { return v.Direct }).(PeeringPropertiesDirectResponsePtrOutput)
}

// The properties that define an exchange peering.
func (o PeeringOutput) Exchange() PeeringPropertiesExchangeResponsePtrOutput {
	return o.ApplyT(func(v *Peering) PeeringPropertiesExchangeResponsePtrOutput { return v.Exchange }).(PeeringPropertiesExchangeResponsePtrOutput)
}

// The kind of the peering.
func (o PeeringOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The location of the resource.
func (o PeeringOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource.
func (o PeeringOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The location of the peering.
func (o PeeringOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringPtrOutput { return v.PeeringLocation }).(pulumi.StringPtrOutput)
}

// The provisioning state of the resource.
func (o PeeringOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU that defines the tier and kind of the peering.
func (o PeeringOutput) Sku() PeeringSkuResponseOutput {
	return o.ApplyT(func(v *Peering) PeeringSkuResponseOutput { return v.Sku }).(PeeringSkuResponseOutput)
}

// The resource tags.
func (o PeeringOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o PeeringOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PeeringOutput{})
}
