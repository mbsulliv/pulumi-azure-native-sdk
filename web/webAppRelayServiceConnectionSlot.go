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

// Hybrid Connection for an App Service app.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2020-12-01
type WebAppRelayServiceConnectionSlot struct {
	pulumi.CustomResourceState

	BiztalkUri             pulumi.StringPtrOutput `pulumi:"biztalkUri"`
	EntityConnectionString pulumi.StringPtrOutput `pulumi:"entityConnectionString"`
	EntityName             pulumi.StringPtrOutput `pulumi:"entityName"`
	Hostname               pulumi.StringPtrOutput `pulumi:"hostname"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name                     pulumi.StringOutput    `pulumi:"name"`
	Port                     pulumi.IntPtrOutput    `pulumi:"port"`
	ResourceConnectionString pulumi.StringPtrOutput `pulumi:"resourceConnectionString"`
	ResourceType             pulumi.StringPtrOutput `pulumi:"resourceType"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppRelayServiceConnectionSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppRelayServiceConnectionSlot(ctx *pulumi.Context,
	name string, args *WebAppRelayServiceConnectionSlotArgs, opts ...pulumi.ResourceOption) (*WebAppRelayServiceConnectionSlot, error) {
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
			Type: pulumi.String("azure-native:web/v20150801:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppRelayServiceConnectionSlot"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppRelayServiceConnectionSlot
	err := ctx.RegisterResource("azure-native:web:WebAppRelayServiceConnectionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppRelayServiceConnectionSlot gets an existing WebAppRelayServiceConnectionSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppRelayServiceConnectionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppRelayServiceConnectionSlotState, opts ...pulumi.ResourceOption) (*WebAppRelayServiceConnectionSlot, error) {
	var resource WebAppRelayServiceConnectionSlot
	err := ctx.ReadResource("azure-native:web:WebAppRelayServiceConnectionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppRelayServiceConnectionSlot resources.
type webAppRelayServiceConnectionSlotState struct {
}

type WebAppRelayServiceConnectionSlotState struct {
}

func (WebAppRelayServiceConnectionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppRelayServiceConnectionSlotState)(nil)).Elem()
}

type webAppRelayServiceConnectionSlotArgs struct {
	BiztalkUri             *string `pulumi:"biztalkUri"`
	EntityConnectionString *string `pulumi:"entityConnectionString"`
	EntityName             *string `pulumi:"entityName"`
	Hostname               *string `pulumi:"hostname"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name                     string  `pulumi:"name"`
	Port                     *int    `pulumi:"port"`
	ResourceConnectionString *string `pulumi:"resourceConnectionString"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceType      *string `pulumi:"resourceType"`
	// Name of the deployment slot. If a slot is not specified, the API will create or update a hybrid connection for the production slot.
	Slot string `pulumi:"slot"`
}

// The set of arguments for constructing a WebAppRelayServiceConnectionSlot resource.
type WebAppRelayServiceConnectionSlotArgs struct {
	BiztalkUri             pulumi.StringPtrInput
	EntityConnectionString pulumi.StringPtrInput
	EntityName             pulumi.StringPtrInput
	Hostname               pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name                     pulumi.StringInput
	Port                     pulumi.IntPtrInput
	ResourceConnectionString pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	ResourceType      pulumi.StringPtrInput
	// Name of the deployment slot. If a slot is not specified, the API will create or update a hybrid connection for the production slot.
	Slot pulumi.StringInput
}

func (WebAppRelayServiceConnectionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppRelayServiceConnectionSlotArgs)(nil)).Elem()
}

type WebAppRelayServiceConnectionSlotInput interface {
	pulumi.Input

	ToWebAppRelayServiceConnectionSlotOutput() WebAppRelayServiceConnectionSlotOutput
	ToWebAppRelayServiceConnectionSlotOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionSlotOutput
}

func (*WebAppRelayServiceConnectionSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppRelayServiceConnectionSlot)(nil)).Elem()
}

func (i *WebAppRelayServiceConnectionSlot) ToWebAppRelayServiceConnectionSlotOutput() WebAppRelayServiceConnectionSlotOutput {
	return i.ToWebAppRelayServiceConnectionSlotOutputWithContext(context.Background())
}

func (i *WebAppRelayServiceConnectionSlot) ToWebAppRelayServiceConnectionSlotOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppRelayServiceConnectionSlotOutput)
}

func (i *WebAppRelayServiceConnectionSlot) ToOutput(ctx context.Context) pulumix.Output[*WebAppRelayServiceConnectionSlot] {
	return pulumix.Output[*WebAppRelayServiceConnectionSlot]{
		OutputState: i.ToWebAppRelayServiceConnectionSlotOutputWithContext(ctx).OutputState,
	}
}

type WebAppRelayServiceConnectionSlotOutput struct{ *pulumi.OutputState }

func (WebAppRelayServiceConnectionSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppRelayServiceConnectionSlot)(nil)).Elem()
}

func (o WebAppRelayServiceConnectionSlotOutput) ToWebAppRelayServiceConnectionSlotOutput() WebAppRelayServiceConnectionSlotOutput {
	return o
}

func (o WebAppRelayServiceConnectionSlotOutput) ToWebAppRelayServiceConnectionSlotOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionSlotOutput {
	return o
}

func (o WebAppRelayServiceConnectionSlotOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppRelayServiceConnectionSlot] {
	return pulumix.Output[*WebAppRelayServiceConnectionSlot]{
		OutputState: o.OutputState,
	}
}

func (o WebAppRelayServiceConnectionSlotOutput) BiztalkUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.BiztalkUri }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionSlotOutput) EntityConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.EntityConnectionString }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionSlotOutput) EntityName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.EntityName }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionSlotOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.Hostname }).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o WebAppRelayServiceConnectionSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppRelayServiceConnectionSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnectionSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppRelayServiceConnectionSlotOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnectionSlot) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o WebAppRelayServiceConnectionSlotOutput) ResourceConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.ResourceConnectionString }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionSlotOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnectionSlot) pulumi.StringPtrOutput { return v.ResourceType }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o WebAppRelayServiceConnectionSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnectionSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppRelayServiceConnectionSlotOutput{})
}
