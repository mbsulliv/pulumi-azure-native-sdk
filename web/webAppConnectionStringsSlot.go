// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// String dictionary resource.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2020-12-01
type WebAppConnectionStringsSlot struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Connection strings.
	Properties ConnStringValueTypePairResponseMapOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppConnectionStringsSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppConnectionStringsSlot(ctx *pulumi.Context,
	name string, args *WebAppConnectionStringsSlotArgs, opts ...pulumi.ResourceOption) (*WebAppConnectionStringsSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppConnectionStringsSlot"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppConnectionStringsSlot
	err := ctx.RegisterResource("azure-native:web:WebAppConnectionStringsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppConnectionStringsSlot gets an existing WebAppConnectionStringsSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppConnectionStringsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppConnectionStringsSlotState, opts ...pulumi.ResourceOption) (*WebAppConnectionStringsSlot, error) {
	var resource WebAppConnectionStringsSlot
	err := ctx.ReadResource("azure-native:web:WebAppConnectionStringsSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppConnectionStringsSlot resources.
type webAppConnectionStringsSlotState struct {
}

type WebAppConnectionStringsSlotState struct {
}

func (WebAppConnectionStringsSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppConnectionStringsSlotState)(nil)).Elem()
}

type webAppConnectionStringsSlotArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Connection strings.
	Properties map[string]ConnStringValueTypePair `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will update the connection settings for the production slot.
	Slot string `pulumi:"slot"`
}

// The set of arguments for constructing a WebAppConnectionStringsSlot resource.
type WebAppConnectionStringsSlotArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Connection strings.
	Properties ConnStringValueTypePairMapInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API will update the connection settings for the production slot.
	Slot pulumi.StringInput
}

func (WebAppConnectionStringsSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppConnectionStringsSlotArgs)(nil)).Elem()
}

type WebAppConnectionStringsSlotInput interface {
	pulumi.Input

	ToWebAppConnectionStringsSlotOutput() WebAppConnectionStringsSlotOutput
	ToWebAppConnectionStringsSlotOutputWithContext(ctx context.Context) WebAppConnectionStringsSlotOutput
}

func (*WebAppConnectionStringsSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppConnectionStringsSlot)(nil)).Elem()
}

func (i *WebAppConnectionStringsSlot) ToWebAppConnectionStringsSlotOutput() WebAppConnectionStringsSlotOutput {
	return i.ToWebAppConnectionStringsSlotOutputWithContext(context.Background())
}

func (i *WebAppConnectionStringsSlot) ToWebAppConnectionStringsSlotOutputWithContext(ctx context.Context) WebAppConnectionStringsSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppConnectionStringsSlotOutput)
}

func (i *WebAppConnectionStringsSlot) ToOutput(ctx context.Context) pulumix.Output[*WebAppConnectionStringsSlot] {
	return pulumix.Output[*WebAppConnectionStringsSlot]{
		OutputState: i.ToWebAppConnectionStringsSlotOutputWithContext(ctx).OutputState,
	}
}

type WebAppConnectionStringsSlotOutput struct{ *pulumi.OutputState }

func (WebAppConnectionStringsSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppConnectionStringsSlot)(nil)).Elem()
}

func (o WebAppConnectionStringsSlotOutput) ToWebAppConnectionStringsSlotOutput() WebAppConnectionStringsSlotOutput {
	return o
}

func (o WebAppConnectionStringsSlotOutput) ToWebAppConnectionStringsSlotOutputWithContext(ctx context.Context) WebAppConnectionStringsSlotOutput {
	return o
}

func (o WebAppConnectionStringsSlotOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppConnectionStringsSlot] {
	return pulumix.Output[*WebAppConnectionStringsSlot]{
		OutputState: o.OutputState,
	}
}

// Kind of resource.
func (o WebAppConnectionStringsSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppConnectionStringsSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppConnectionStringsSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppConnectionStringsSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Connection strings.
func (o WebAppConnectionStringsSlotOutput) Properties() ConnStringValueTypePairResponseMapOutput {
	return o.ApplyT(func(v *WebAppConnectionStringsSlot) ConnStringValueTypePairResponseMapOutput { return v.Properties }).(ConnStringValueTypePairResponseMapOutput)
}

// Resource type.
func (o WebAppConnectionStringsSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppConnectionStringsSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppConnectionStringsSlotOutput{})
}
