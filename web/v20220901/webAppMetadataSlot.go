// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// String dictionary resource.
type WebAppMetadataSlot struct {
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

// NewWebAppMetadataSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppMetadataSlot(ctx *pulumi.Context,
	name string, args *WebAppMetadataSlotArgs, opts ...pulumi.ResourceOption) (*WebAppMetadataSlot, error) {
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
			Type: pulumi.String("azure-native:web:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppMetadataSlot"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppMetadataSlot
	err := ctx.RegisterResource("azure-native:web/v20220901:WebAppMetadataSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppMetadataSlot gets an existing WebAppMetadataSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppMetadataSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppMetadataSlotState, opts ...pulumi.ResourceOption) (*WebAppMetadataSlot, error) {
	var resource WebAppMetadataSlot
	err := ctx.ReadResource("azure-native:web/v20220901:WebAppMetadataSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppMetadataSlot resources.
type webAppMetadataSlotState struct {
}

type WebAppMetadataSlotState struct {
}

func (WebAppMetadataSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppMetadataSlotState)(nil)).Elem()
}

type webAppMetadataSlotArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Settings.
	Properties map[string]string `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will update the metadata for the production slot.
	Slot string `pulumi:"slot"`
}

// The set of arguments for constructing a WebAppMetadataSlot resource.
type WebAppMetadataSlotArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Settings.
	Properties pulumi.StringMapInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API will update the metadata for the production slot.
	Slot pulumi.StringInput
}

func (WebAppMetadataSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppMetadataSlotArgs)(nil)).Elem()
}

type WebAppMetadataSlotInput interface {
	pulumi.Input

	ToWebAppMetadataSlotOutput() WebAppMetadataSlotOutput
	ToWebAppMetadataSlotOutputWithContext(ctx context.Context) WebAppMetadataSlotOutput
}

func (*WebAppMetadataSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppMetadataSlot)(nil)).Elem()
}

func (i *WebAppMetadataSlot) ToWebAppMetadataSlotOutput() WebAppMetadataSlotOutput {
	return i.ToWebAppMetadataSlotOutputWithContext(context.Background())
}

func (i *WebAppMetadataSlot) ToWebAppMetadataSlotOutputWithContext(ctx context.Context) WebAppMetadataSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppMetadataSlotOutput)
}

func (i *WebAppMetadataSlot) ToOutput(ctx context.Context) pulumix.Output[*WebAppMetadataSlot] {
	return pulumix.Output[*WebAppMetadataSlot]{
		OutputState: i.ToWebAppMetadataSlotOutputWithContext(ctx).OutputState,
	}
}

type WebAppMetadataSlotOutput struct{ *pulumi.OutputState }

func (WebAppMetadataSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppMetadataSlot)(nil)).Elem()
}

func (o WebAppMetadataSlotOutput) ToWebAppMetadataSlotOutput() WebAppMetadataSlotOutput {
	return o
}

func (o WebAppMetadataSlotOutput) ToWebAppMetadataSlotOutputWithContext(ctx context.Context) WebAppMetadataSlotOutput {
	return o
}

func (o WebAppMetadataSlotOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppMetadataSlot] {
	return pulumix.Output[*WebAppMetadataSlot]{
		OutputState: o.OutputState,
	}
}

// Kind of resource.
func (o WebAppMetadataSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppMetadataSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppMetadataSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppMetadataSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Settings.
func (o WebAppMetadataSlotOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebAppMetadataSlot) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource type.
func (o WebAppMetadataSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppMetadataSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppMetadataSlotOutput{})
}
