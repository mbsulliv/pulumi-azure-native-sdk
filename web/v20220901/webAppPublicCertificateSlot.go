// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Public certificate object
type WebAppPublicCertificateSlot struct {
	pulumi.CustomResourceState

	// Public Certificate byte array
	Blob pulumi.StringPtrOutput `pulumi:"blob"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Public Certificate Location
	PublicCertificateLocation pulumi.StringPtrOutput `pulumi:"publicCertificateLocation"`
	// Certificate Thumbprint
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppPublicCertificateSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppPublicCertificateSlot(ctx *pulumi.Context,
	name string, args *WebAppPublicCertificateSlotArgs, opts ...pulumi.ResourceOption) (*WebAppPublicCertificateSlot, error) {
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
			Type: pulumi.String("azure-native:web:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppPublicCertificateSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebAppPublicCertificateSlot"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppPublicCertificateSlot
	err := ctx.RegisterResource("azure-native:web/v20220901:WebAppPublicCertificateSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppPublicCertificateSlot gets an existing WebAppPublicCertificateSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppPublicCertificateSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppPublicCertificateSlotState, opts ...pulumi.ResourceOption) (*WebAppPublicCertificateSlot, error) {
	var resource WebAppPublicCertificateSlot
	err := ctx.ReadResource("azure-native:web/v20220901:WebAppPublicCertificateSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppPublicCertificateSlot resources.
type webAppPublicCertificateSlotState struct {
}

type WebAppPublicCertificateSlotState struct {
}

func (WebAppPublicCertificateSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPublicCertificateSlotState)(nil)).Elem()
}

type webAppPublicCertificateSlotArgs struct {
	// Public Certificate byte array
	Blob *string `pulumi:"blob"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Public Certificate Location
	PublicCertificateLocation *PublicCertificateLocation `pulumi:"publicCertificateLocation"`
	// Public certificate name.
	PublicCertificateName *string `pulumi:"publicCertificateName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will create a binding for the production slot.
	Slot string `pulumi:"slot"`
}

// The set of arguments for constructing a WebAppPublicCertificateSlot resource.
type WebAppPublicCertificateSlotArgs struct {
	// Public Certificate byte array
	Blob pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Public Certificate Location
	PublicCertificateLocation PublicCertificateLocationPtrInput
	// Public certificate name.
	PublicCertificateName pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API will create a binding for the production slot.
	Slot pulumi.StringInput
}

func (WebAppPublicCertificateSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPublicCertificateSlotArgs)(nil)).Elem()
}

type WebAppPublicCertificateSlotInput interface {
	pulumi.Input

	ToWebAppPublicCertificateSlotOutput() WebAppPublicCertificateSlotOutput
	ToWebAppPublicCertificateSlotOutputWithContext(ctx context.Context) WebAppPublicCertificateSlotOutput
}

func (*WebAppPublicCertificateSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppPublicCertificateSlot)(nil)).Elem()
}

func (i *WebAppPublicCertificateSlot) ToWebAppPublicCertificateSlotOutput() WebAppPublicCertificateSlotOutput {
	return i.ToWebAppPublicCertificateSlotOutputWithContext(context.Background())
}

func (i *WebAppPublicCertificateSlot) ToWebAppPublicCertificateSlotOutputWithContext(ctx context.Context) WebAppPublicCertificateSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPublicCertificateSlotOutput)
}

type WebAppPublicCertificateSlotOutput struct{ *pulumi.OutputState }

func (WebAppPublicCertificateSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppPublicCertificateSlot)(nil)).Elem()
}

func (o WebAppPublicCertificateSlotOutput) ToWebAppPublicCertificateSlotOutput() WebAppPublicCertificateSlotOutput {
	return o
}

func (o WebAppPublicCertificateSlotOutput) ToWebAppPublicCertificateSlotOutputWithContext(ctx context.Context) WebAppPublicCertificateSlotOutput {
	return o
}

// Public Certificate byte array
func (o WebAppPublicCertificateSlotOutput) Blob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPublicCertificateSlot) pulumi.StringPtrOutput { return v.Blob }).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o WebAppPublicCertificateSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPublicCertificateSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppPublicCertificateSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPublicCertificateSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Public Certificate Location
func (o WebAppPublicCertificateSlotOutput) PublicCertificateLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPublicCertificateSlot) pulumi.StringPtrOutput { return v.PublicCertificateLocation }).(pulumi.StringPtrOutput)
}

// Certificate Thumbprint
func (o WebAppPublicCertificateSlotOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPublicCertificateSlot) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

// Resource type.
func (o WebAppPublicCertificateSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPublicCertificateSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppPublicCertificateSlotOutput{})
}
