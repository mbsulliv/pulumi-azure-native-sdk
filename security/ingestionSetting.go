// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Configures how to correlate scan data and logs with resources associated with the subscription.
// Azure REST API version: 2021-01-15-preview. Prior API version in Azure Native 1.x: 2021-01-15-preview
type IngestionSetting struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIngestionSetting registers a new resource with the given unique name, arguments, and options.
func NewIngestionSetting(ctx *pulumi.Context,
	name string, args *IngestionSettingArgs, opts ...pulumi.ResourceOption) (*IngestionSetting, error) {
	if args == nil {
		args = &IngestionSettingArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security/v20210115preview:IngestionSetting"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IngestionSetting
	err := ctx.RegisterResource("azure-native:security:IngestionSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIngestionSetting gets an existing IngestionSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIngestionSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IngestionSettingState, opts ...pulumi.ResourceOption) (*IngestionSetting, error) {
	var resource IngestionSetting
	err := ctx.ReadResource("azure-native:security:IngestionSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IngestionSetting resources.
type ingestionSettingState struct {
}

type IngestionSettingState struct {
}

func (IngestionSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*ingestionSettingState)(nil)).Elem()
}

type ingestionSettingArgs struct {
	// Name of the ingestion setting
	IngestionSettingName *string `pulumi:"ingestionSettingName"`
}

// The set of arguments for constructing a IngestionSetting resource.
type IngestionSettingArgs struct {
	// Name of the ingestion setting
	IngestionSettingName pulumi.StringPtrInput
}

func (IngestionSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ingestionSettingArgs)(nil)).Elem()
}

type IngestionSettingInput interface {
	pulumi.Input

	ToIngestionSettingOutput() IngestionSettingOutput
	ToIngestionSettingOutputWithContext(ctx context.Context) IngestionSettingOutput
}

func (*IngestionSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**IngestionSetting)(nil)).Elem()
}

func (i *IngestionSetting) ToIngestionSettingOutput() IngestionSettingOutput {
	return i.ToIngestionSettingOutputWithContext(context.Background())
}

func (i *IngestionSetting) ToIngestionSettingOutputWithContext(ctx context.Context) IngestionSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionSettingOutput)
}

func (i *IngestionSetting) ToOutput(ctx context.Context) pulumix.Output[*IngestionSetting] {
	return pulumix.Output[*IngestionSetting]{
		OutputState: i.ToIngestionSettingOutputWithContext(ctx).OutputState,
	}
}

type IngestionSettingOutput struct{ *pulumi.OutputState }

func (IngestionSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngestionSetting)(nil)).Elem()
}

func (o IngestionSettingOutput) ToIngestionSettingOutput() IngestionSettingOutput {
	return o
}

func (o IngestionSettingOutput) ToIngestionSettingOutputWithContext(ctx context.Context) IngestionSettingOutput {
	return o
}

func (o IngestionSettingOutput) ToOutput(ctx context.Context) pulumix.Output[*IngestionSetting] {
	return pulumix.Output[*IngestionSetting]{
		OutputState: o.OutputState,
	}
}

// Resource name
func (o IngestionSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IngestionSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type
func (o IngestionSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IngestionSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IngestionSettingOutput{})
}
