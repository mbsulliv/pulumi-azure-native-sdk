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

// Push settings for the App.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2020-10-01.
type WebAppSitePushSettingsSlot struct {
	pulumi.CustomResourceState

	// Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in the push registration endpoint.
	DynamicTagsJson pulumi.StringPtrOutput `pulumi:"dynamicTagsJson"`
	// Gets or sets a flag indicating whether the Push endpoint is enabled.
	IsPushEnabled pulumi.BoolOutput `pulumi:"isPushEnabled"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets a JSON string containing a list of tags that are whitelisted for use by the push registration endpoint.
	TagWhitelistJson pulumi.StringPtrOutput `pulumi:"tagWhitelistJson"`
	// Gets or sets a JSON string containing a list of tags that require user authentication to be used in the push registration endpoint.
	// Tags can consist of alphanumeric characters and the following:
	// '_', '@', '#', '.', ':', '-'.
	// Validation should be performed at the PushRequestHandler.
	TagsRequiringAuth pulumi.StringPtrOutput `pulumi:"tagsRequiringAuth"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppSitePushSettingsSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppSitePushSettingsSlot(ctx *pulumi.Context,
	name string, args *WebAppSitePushSettingsSlotArgs, opts ...pulumi.ResourceOption) (*WebAppSitePushSettingsSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IsPushEnabled == nil {
		return nil, errors.New("invalid value for required argument 'IsPushEnabled'")
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
			Type: pulumi.String("azure-native:web/v20160801:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppSitePushSettingsSlot"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppSitePushSettingsSlot
	err := ctx.RegisterResource("azure-native:web:WebAppSitePushSettingsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppSitePushSettingsSlot gets an existing WebAppSitePushSettingsSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppSitePushSettingsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSitePushSettingsSlotState, opts ...pulumi.ResourceOption) (*WebAppSitePushSettingsSlot, error) {
	var resource WebAppSitePushSettingsSlot
	err := ctx.ReadResource("azure-native:web:WebAppSitePushSettingsSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppSitePushSettingsSlot resources.
type webAppSitePushSettingsSlotState struct {
}

type WebAppSitePushSettingsSlotState struct {
}

func (WebAppSitePushSettingsSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSitePushSettingsSlotState)(nil)).Elem()
}

type webAppSitePushSettingsSlotArgs struct {
	// Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in the push registration endpoint.
	DynamicTagsJson *string `pulumi:"dynamicTagsJson"`
	// Gets or sets a flag indicating whether the Push endpoint is enabled.
	IsPushEnabled bool `pulumi:"isPushEnabled"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of web app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot string `pulumi:"slot"`
	// Gets or sets a JSON string containing a list of tags that are whitelisted for use by the push registration endpoint.
	TagWhitelistJson *string `pulumi:"tagWhitelistJson"`
	// Gets or sets a JSON string containing a list of tags that require user authentication to be used in the push registration endpoint.
	// Tags can consist of alphanumeric characters and the following:
	// '_', '@', '#', '.', ':', '-'.
	// Validation should be performed at the PushRequestHandler.
	TagsRequiringAuth *string `pulumi:"tagsRequiringAuth"`
}

// The set of arguments for constructing a WebAppSitePushSettingsSlot resource.
type WebAppSitePushSettingsSlotArgs struct {
	// Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in the push registration endpoint.
	DynamicTagsJson pulumi.StringPtrInput
	// Gets or sets a flag indicating whether the Push endpoint is enabled.
	IsPushEnabled pulumi.BoolInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of web app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of web app slot. If not specified then will default to production slot.
	Slot pulumi.StringInput
	// Gets or sets a JSON string containing a list of tags that are whitelisted for use by the push registration endpoint.
	TagWhitelistJson pulumi.StringPtrInput
	// Gets or sets a JSON string containing a list of tags that require user authentication to be used in the push registration endpoint.
	// Tags can consist of alphanumeric characters and the following:
	// '_', '@', '#', '.', ':', '-'.
	// Validation should be performed at the PushRequestHandler.
	TagsRequiringAuth pulumi.StringPtrInput
}

func (WebAppSitePushSettingsSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSitePushSettingsSlotArgs)(nil)).Elem()
}

type WebAppSitePushSettingsSlotInput interface {
	pulumi.Input

	ToWebAppSitePushSettingsSlotOutput() WebAppSitePushSettingsSlotOutput
	ToWebAppSitePushSettingsSlotOutputWithContext(ctx context.Context) WebAppSitePushSettingsSlotOutput
}

func (*WebAppSitePushSettingsSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSitePushSettingsSlot)(nil)).Elem()
}

func (i *WebAppSitePushSettingsSlot) ToWebAppSitePushSettingsSlotOutput() WebAppSitePushSettingsSlotOutput {
	return i.ToWebAppSitePushSettingsSlotOutputWithContext(context.Background())
}

func (i *WebAppSitePushSettingsSlot) ToWebAppSitePushSettingsSlotOutputWithContext(ctx context.Context) WebAppSitePushSettingsSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSitePushSettingsSlotOutput)
}

func (i *WebAppSitePushSettingsSlot) ToOutput(ctx context.Context) pulumix.Output[*WebAppSitePushSettingsSlot] {
	return pulumix.Output[*WebAppSitePushSettingsSlot]{
		OutputState: i.ToWebAppSitePushSettingsSlotOutputWithContext(ctx).OutputState,
	}
}

type WebAppSitePushSettingsSlotOutput struct{ *pulumi.OutputState }

func (WebAppSitePushSettingsSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSitePushSettingsSlot)(nil)).Elem()
}

func (o WebAppSitePushSettingsSlotOutput) ToWebAppSitePushSettingsSlotOutput() WebAppSitePushSettingsSlotOutput {
	return o
}

func (o WebAppSitePushSettingsSlotOutput) ToWebAppSitePushSettingsSlotOutputWithContext(ctx context.Context) WebAppSitePushSettingsSlotOutput {
	return o
}

func (o WebAppSitePushSettingsSlotOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppSitePushSettingsSlot] {
	return pulumix.Output[*WebAppSitePushSettingsSlot]{
		OutputState: o.OutputState,
	}
}

// Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in the push registration endpoint.
func (o WebAppSitePushSettingsSlotOutput) DynamicTagsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSitePushSettingsSlot) pulumi.StringPtrOutput { return v.DynamicTagsJson }).(pulumi.StringPtrOutput)
}

// Gets or sets a flag indicating whether the Push endpoint is enabled.
func (o WebAppSitePushSettingsSlotOutput) IsPushEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *WebAppSitePushSettingsSlot) pulumi.BoolOutput { return v.IsPushEnabled }).(pulumi.BoolOutput)
}

// Kind of resource.
func (o WebAppSitePushSettingsSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSitePushSettingsSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppSitePushSettingsSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSitePushSettingsSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets a JSON string containing a list of tags that are whitelisted for use by the push registration endpoint.
func (o WebAppSitePushSettingsSlotOutput) TagWhitelistJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSitePushSettingsSlot) pulumi.StringPtrOutput { return v.TagWhitelistJson }).(pulumi.StringPtrOutput)
}

// Gets or sets a JSON string containing a list of tags that require user authentication to be used in the push registration endpoint.
// Tags can consist of alphanumeric characters and the following:
// '_', '@', '#', '.', ':', '-'.
// Validation should be performed at the PushRequestHandler.
func (o WebAppSitePushSettingsSlotOutput) TagsRequiringAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSitePushSettingsSlot) pulumi.StringPtrOutput { return v.TagsRequiringAuth }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o WebAppSitePushSettingsSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSitePushSettingsSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppSitePushSettingsSlotOutput{})
}
