// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220915

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Bot channel resource definition
type Channel struct {
	pulumi.CustomResourceState

	// Entity Tag.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Required. Gets or sets the Kind of the resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Specifies the location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The set of properties specific to bot channel resource
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Entity zones
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewChannel registers a new resource with the given unique name, arguments, and options.
func NewChannel(ctx *pulumi.Context,
	name string, args *ChannelArgs, opts ...pulumi.ResourceOption) (*Channel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:botservice:Channel"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20171201:Channel"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20180712:Channel"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20200602:Channel"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20210301:Channel"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20210501preview:Channel"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20220615preview:Channel"),
		},
	})
	opts = append(opts, aliases)
	var resource Channel
	err := ctx.RegisterResource("azure-native:botservice/v20220915:Channel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannel gets an existing Channel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelState, opts ...pulumi.ResourceOption) (*Channel, error) {
	var resource Channel
	err := ctx.ReadResource("azure-native:botservice/v20220915:Channel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Channel resources.
type channelState struct {
}

type ChannelState struct {
}

func (ChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelState)(nil)).Elem()
}

type channelArgs struct {
	// The name of the Channel resource.
	ChannelName *string `pulumi:"channelName"`
	// Required. Gets or sets the Kind of the resource.
	Kind *string `pulumi:"kind"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// The set of properties specific to bot channel resource
	Properties interface{} `pulumi:"properties"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName string `pulumi:"resourceName"`
	// Gets or sets the SKU of the resource.
	Sku *Sku `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Channel resource.
type ChannelArgs struct {
	// The name of the Channel resource.
	ChannelName pulumi.StringPtrInput
	// Required. Gets or sets the Kind of the resource.
	Kind pulumi.StringPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// The set of properties specific to bot channel resource
	Properties pulumi.Input
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName pulumi.StringInput
	// The name of the Bot resource.
	ResourceName pulumi.StringInput
	// Gets or sets the SKU of the resource.
	Sku SkuPtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
}

func (ChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelArgs)(nil)).Elem()
}

type ChannelInput interface {
	pulumi.Input

	ToChannelOutput() ChannelOutput
	ToChannelOutputWithContext(ctx context.Context) ChannelOutput
}

func (*Channel) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (i *Channel) ToChannelOutput() ChannelOutput {
	return i.ToChannelOutputWithContext(context.Background())
}

func (i *Channel) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelOutput)
}

type ChannelOutput struct{ *pulumi.OutputState }

func (ChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (o ChannelOutput) ToChannelOutput() ChannelOutput {
	return o
}

func (o ChannelOutput) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return o
}

// Entity Tag.
func (o ChannelOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Required. Gets or sets the Kind of the resource.
func (o ChannelOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Specifies the location of the resource.
func (o ChannelOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Specifies the name of the resource.
func (o ChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The set of properties specific to bot channel resource
func (o ChannelOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Channel) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Gets or sets the SKU of the resource.
func (o ChannelOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Channel) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Contains resource tags defined as key/value pairs.
func (o ChannelOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the type of the resource.
func (o ChannelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Entity zones
func (o ChannelOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ChannelOutput{})
}
