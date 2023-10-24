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
type WebAppMetadata struct {
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

// NewWebAppMetadata registers a new resource with the given unique name, arguments, and options.
func NewWebAppMetadata(ctx *pulumi.Context,
	name string, args *WebAppMetadataArgs, opts ...pulumi.ResourceOption) (*WebAppMetadata, error) {
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
			Type: pulumi.String("azure-native:web/v20150801:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppMetadata"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppMetadata"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppMetadata
	err := ctx.RegisterResource("azure-native:web:WebAppMetadata", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppMetadata gets an existing WebAppMetadata resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppMetadata(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppMetadataState, opts ...pulumi.ResourceOption) (*WebAppMetadata, error) {
	var resource WebAppMetadata
	err := ctx.ReadResource("azure-native:web:WebAppMetadata", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppMetadata resources.
type webAppMetadataState struct {
}

type WebAppMetadataState struct {
}

func (WebAppMetadataState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppMetadataState)(nil)).Elem()
}

type webAppMetadataArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Settings.
	Properties map[string]string `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppMetadata resource.
type WebAppMetadataArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Settings.
	Properties pulumi.StringMapInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppMetadataArgs)(nil)).Elem()
}

type WebAppMetadataInput interface {
	pulumi.Input

	ToWebAppMetadataOutput() WebAppMetadataOutput
	ToWebAppMetadataOutputWithContext(ctx context.Context) WebAppMetadataOutput
}

func (*WebAppMetadata) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppMetadata)(nil)).Elem()
}

func (i *WebAppMetadata) ToWebAppMetadataOutput() WebAppMetadataOutput {
	return i.ToWebAppMetadataOutputWithContext(context.Background())
}

func (i *WebAppMetadata) ToWebAppMetadataOutputWithContext(ctx context.Context) WebAppMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppMetadataOutput)
}

func (i *WebAppMetadata) ToOutput(ctx context.Context) pulumix.Output[*WebAppMetadata] {
	return pulumix.Output[*WebAppMetadata]{
		OutputState: i.ToWebAppMetadataOutputWithContext(ctx).OutputState,
	}
}

type WebAppMetadataOutput struct{ *pulumi.OutputState }

func (WebAppMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppMetadata)(nil)).Elem()
}

func (o WebAppMetadataOutput) ToWebAppMetadataOutput() WebAppMetadataOutput {
	return o
}

func (o WebAppMetadataOutput) ToWebAppMetadataOutputWithContext(ctx context.Context) WebAppMetadataOutput {
	return o
}

func (o WebAppMetadataOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppMetadata] {
	return pulumix.Output[*WebAppMetadata]{
		OutputState: o.OutputState,
	}
}

// Kind of resource.
func (o WebAppMetadataOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppMetadata) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppMetadataOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppMetadata) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Settings.
func (o WebAppMetadataOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebAppMetadata) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource type.
func (o WebAppMetadataOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppMetadata) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppMetadataOutput{})
}
