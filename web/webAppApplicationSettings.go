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
type WebAppApplicationSettings struct {
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

// NewWebAppApplicationSettings registers a new resource with the given unique name, arguments, and options.
func NewWebAppApplicationSettings(ctx *pulumi.Context,
	name string, args *WebAppApplicationSettingsArgs, opts ...pulumi.ResourceOption) (*WebAppApplicationSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppApplicationSettings"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppApplicationSettings
	err := ctx.RegisterResource("azure-native:web:WebAppApplicationSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppApplicationSettings gets an existing WebAppApplicationSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppApplicationSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppApplicationSettingsState, opts ...pulumi.ResourceOption) (*WebAppApplicationSettings, error) {
	var resource WebAppApplicationSettings
	err := ctx.ReadResource("azure-native:web:WebAppApplicationSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppApplicationSettings resources.
type webAppApplicationSettingsState struct {
}

type WebAppApplicationSettingsState struct {
}

func (WebAppApplicationSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppApplicationSettingsState)(nil)).Elem()
}

type webAppApplicationSettingsArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Settings.
	Properties map[string]string `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppApplicationSettings resource.
type WebAppApplicationSettingsArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Settings.
	Properties pulumi.StringMapInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppApplicationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppApplicationSettingsArgs)(nil)).Elem()
}

type WebAppApplicationSettingsInput interface {
	pulumi.Input

	ToWebAppApplicationSettingsOutput() WebAppApplicationSettingsOutput
	ToWebAppApplicationSettingsOutputWithContext(ctx context.Context) WebAppApplicationSettingsOutput
}

func (*WebAppApplicationSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppApplicationSettings)(nil)).Elem()
}

func (i *WebAppApplicationSettings) ToWebAppApplicationSettingsOutput() WebAppApplicationSettingsOutput {
	return i.ToWebAppApplicationSettingsOutputWithContext(context.Background())
}

func (i *WebAppApplicationSettings) ToWebAppApplicationSettingsOutputWithContext(ctx context.Context) WebAppApplicationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppApplicationSettingsOutput)
}

func (i *WebAppApplicationSettings) ToOutput(ctx context.Context) pulumix.Output[*WebAppApplicationSettings] {
	return pulumix.Output[*WebAppApplicationSettings]{
		OutputState: i.ToWebAppApplicationSettingsOutputWithContext(ctx).OutputState,
	}
}

type WebAppApplicationSettingsOutput struct{ *pulumi.OutputState }

func (WebAppApplicationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppApplicationSettings)(nil)).Elem()
}

func (o WebAppApplicationSettingsOutput) ToWebAppApplicationSettingsOutput() WebAppApplicationSettingsOutput {
	return o
}

func (o WebAppApplicationSettingsOutput) ToWebAppApplicationSettingsOutputWithContext(ctx context.Context) WebAppApplicationSettingsOutput {
	return o
}

func (o WebAppApplicationSettingsOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppApplicationSettings] {
	return pulumix.Output[*WebAppApplicationSettings]{
		OutputState: o.OutputState,
	}
}

// Kind of resource.
func (o WebAppApplicationSettingsOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppApplicationSettings) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppApplicationSettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppApplicationSettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Settings.
func (o WebAppApplicationSettingsOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebAppApplicationSettings) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource type.
func (o WebAppApplicationSettingsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppApplicationSettings) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppApplicationSettingsOutput{})
}
