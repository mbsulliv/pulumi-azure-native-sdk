// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220615preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Bot resource definition
type Bot struct {
	pulumi.CustomResourceState

	// Entity Tag.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Required. Gets or sets the Kind of the resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Specifies the location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The set of properties specific to bot resource
	Properties BotPropertiesResponseOutput `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Entity zones
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewBot registers a new resource with the given unique name, arguments, and options.
func NewBot(ctx *pulumi.Context,
	name string, args *BotArgs, opts ...pulumi.ResourceOption) (*Bot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToBotPropertiesPtrOutput().ApplyT(func(v *BotProperties) *BotProperties { return v.Defaults() }).(BotPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:botservice:Bot"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20171201:Bot"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20180712:Bot"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20200602:Bot"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20210301:Bot"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20210501preview:Bot"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20220915:Bot"),
		},
	})
	opts = append(opts, aliases)
	var resource Bot
	err := ctx.RegisterResource("azure-native:botservice/v20220615preview:Bot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBot gets an existing Bot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BotState, opts ...pulumi.ResourceOption) (*Bot, error) {
	var resource Bot
	err := ctx.ReadResource("azure-native:botservice/v20220615preview:Bot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bot resources.
type botState struct {
}

type BotState struct {
}

func (BotState) ElementType() reflect.Type {
	return reflect.TypeOf((*botState)(nil)).Elem()
}

type botArgs struct {
	// Required. Gets or sets the Kind of the resource.
	Kind *string `pulumi:"kind"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// The set of properties specific to bot resource
	Properties *BotProperties `pulumi:"properties"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName *string `pulumi:"resourceName"`
	// Gets or sets the SKU of the resource.
	Sku *Sku `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Bot resource.
type BotArgs struct {
	// Required. Gets or sets the Kind of the resource.
	Kind pulumi.StringPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// The set of properties specific to bot resource
	Properties BotPropertiesPtrInput
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName pulumi.StringInput
	// The name of the Bot resource.
	ResourceName pulumi.StringPtrInput
	// Gets or sets the SKU of the resource.
	Sku SkuPtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
}

func (BotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*botArgs)(nil)).Elem()
}

type BotInput interface {
	pulumi.Input

	ToBotOutput() BotOutput
	ToBotOutputWithContext(ctx context.Context) BotOutput
}

func (*Bot) ElementType() reflect.Type {
	return reflect.TypeOf((**Bot)(nil)).Elem()
}

func (i *Bot) ToBotOutput() BotOutput {
	return i.ToBotOutputWithContext(context.Background())
}

func (i *Bot) ToBotOutputWithContext(ctx context.Context) BotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotOutput)
}

type BotOutput struct{ *pulumi.OutputState }

func (BotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bot)(nil)).Elem()
}

func (o BotOutput) ToBotOutput() BotOutput {
	return o
}

func (o BotOutput) ToBotOutputWithContext(ctx context.Context) BotOutput {
	return o
}

// Entity Tag.
func (o BotOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Required. Gets or sets the Kind of the resource.
func (o BotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Specifies the location of the resource.
func (o BotOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Specifies the name of the resource.
func (o BotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The set of properties specific to bot resource
func (o BotOutput) Properties() BotPropertiesResponseOutput {
	return o.ApplyT(func(v *Bot) BotPropertiesResponseOutput { return v.Properties }).(BotPropertiesResponseOutput)
}

// Gets or sets the SKU of the resource.
func (o BotOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Bot) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Contains resource tags defined as key/value pairs.
func (o BotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the type of the resource.
func (o BotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Entity zones
func (o BotOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(BotOutput{})
}
