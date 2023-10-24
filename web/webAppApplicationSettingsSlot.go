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
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2020-10-01.
type WebAppApplicationSettingsSlot struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Settings.
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppApplicationSettingsSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppApplicationSettingsSlot(ctx *pulumi.Context,
	name string, args *WebAppApplicationSettingsSlotArgs, opts ...pulumi.ResourceOption) (*WebAppApplicationSettingsSlot, error) {
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
			Type: pulumi.String("azure-native:web/v20150801:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppApplicationSettingsSlot"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppApplicationSettingsSlot
	err := ctx.RegisterResource("azure-native:web:WebAppApplicationSettingsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppApplicationSettingsSlot gets an existing WebAppApplicationSettingsSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppApplicationSettingsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppApplicationSettingsSlotState, opts ...pulumi.ResourceOption) (*WebAppApplicationSettingsSlot, error) {
	var resource WebAppApplicationSettingsSlot
	err := ctx.ReadResource("azure-native:web:WebAppApplicationSettingsSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppApplicationSettingsSlot resources.
type webAppApplicationSettingsSlotState struct {
}

type WebAppApplicationSettingsSlotState struct {
}

func (WebAppApplicationSettingsSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppApplicationSettingsSlotState)(nil)).Elem()
}

type webAppApplicationSettingsSlotArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Settings.
	Properties map[string]string `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will update the application settings for the production slot.
	Slot string `pulumi:"slot"`
}

// The set of arguments for constructing a WebAppApplicationSettingsSlot resource.
type WebAppApplicationSettingsSlotArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Settings.
	Properties pulumi.StringMapInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API will update the application settings for the production slot.
	Slot pulumi.StringInput
}

func (WebAppApplicationSettingsSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppApplicationSettingsSlotArgs)(nil)).Elem()
}

type WebAppApplicationSettingsSlotInput interface {
	pulumi.Input

	ToWebAppApplicationSettingsSlotOutput() WebAppApplicationSettingsSlotOutput
	ToWebAppApplicationSettingsSlotOutputWithContext(ctx context.Context) WebAppApplicationSettingsSlotOutput
}

func (*WebAppApplicationSettingsSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppApplicationSettingsSlot)(nil)).Elem()
}

func (i *WebAppApplicationSettingsSlot) ToWebAppApplicationSettingsSlotOutput() WebAppApplicationSettingsSlotOutput {
	return i.ToWebAppApplicationSettingsSlotOutputWithContext(context.Background())
}

func (i *WebAppApplicationSettingsSlot) ToWebAppApplicationSettingsSlotOutputWithContext(ctx context.Context) WebAppApplicationSettingsSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppApplicationSettingsSlotOutput)
}

func (i *WebAppApplicationSettingsSlot) ToOutput(ctx context.Context) pulumix.Output[*WebAppApplicationSettingsSlot] {
	return pulumix.Output[*WebAppApplicationSettingsSlot]{
		OutputState: i.ToWebAppApplicationSettingsSlotOutputWithContext(ctx).OutputState,
	}
}

type WebAppApplicationSettingsSlotOutput struct{ *pulumi.OutputState }

func (WebAppApplicationSettingsSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppApplicationSettingsSlot)(nil)).Elem()
}

func (o WebAppApplicationSettingsSlotOutput) ToWebAppApplicationSettingsSlotOutput() WebAppApplicationSettingsSlotOutput {
	return o
}

func (o WebAppApplicationSettingsSlotOutput) ToWebAppApplicationSettingsSlotOutputWithContext(ctx context.Context) WebAppApplicationSettingsSlotOutput {
	return o
}

func (o WebAppApplicationSettingsSlotOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppApplicationSettingsSlot] {
	return pulumix.Output[*WebAppApplicationSettingsSlot]{
		OutputState: o.OutputState,
	}
}

// Kind of resource.
func (o WebAppApplicationSettingsSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppApplicationSettingsSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppApplicationSettingsSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppApplicationSettingsSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Settings.
func (o WebAppApplicationSettingsSlotOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebAppApplicationSettingsSlot) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource type.
func (o WebAppApplicationSettingsSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppApplicationSettingsSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppApplicationSettingsSlotOutput{})
}
