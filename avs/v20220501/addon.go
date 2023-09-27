// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An addon resource
type Addon struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of an addon resource
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAddon registers a new resource with the given unique name, arguments, and options.
func NewAddon(ctx *pulumi.Context,
	name string, args *AddonArgs, opts ...pulumi.ResourceOption) (*Addon, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230301:Addon"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Addon
	err := ctx.RegisterResource("azure-native:avs/v20220501:Addon", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddon gets an existing Addon resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddon(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddonState, opts ...pulumi.ResourceOption) (*Addon, error) {
	var resource Addon
	err := ctx.ReadResource("azure-native:avs/v20220501:Addon", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Addon resources.
type addonState struct {
}

type AddonState struct {
}

func (AddonState) ElementType() reflect.Type {
	return reflect.TypeOf((*addonState)(nil)).Elem()
}

type addonArgs struct {
	// Name of the addon for the private cloud
	AddonName *string `pulumi:"addonName"`
	// The name of the private cloud.
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The properties of an addon resource
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Addon resource.
type AddonArgs struct {
	// Name of the addon for the private cloud
	AddonName pulumi.StringPtrInput
	// The name of the private cloud.
	PrivateCloudName pulumi.StringInput
	// The properties of an addon resource
	Properties pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (AddonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addonArgs)(nil)).Elem()
}

type AddonInput interface {
	pulumi.Input

	ToAddonOutput() AddonOutput
	ToAddonOutputWithContext(ctx context.Context) AddonOutput
}

func (*Addon) ElementType() reflect.Type {
	return reflect.TypeOf((**Addon)(nil)).Elem()
}

func (i *Addon) ToAddonOutput() AddonOutput {
	return i.ToAddonOutputWithContext(context.Background())
}

func (i *Addon) ToAddonOutputWithContext(ctx context.Context) AddonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonOutput)
}

func (i *Addon) ToOutput(ctx context.Context) pulumix.Output[*Addon] {
	return pulumix.Output[*Addon]{
		OutputState: i.ToAddonOutputWithContext(ctx).OutputState,
	}
}

type AddonOutput struct{ *pulumi.OutputState }

func (AddonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Addon)(nil)).Elem()
}

func (o AddonOutput) ToAddonOutput() AddonOutput {
	return o
}

func (o AddonOutput) ToAddonOutputWithContext(ctx context.Context) AddonOutput {
	return o
}

func (o AddonOutput) ToOutput(ctx context.Context) pulumix.Output[*Addon] {
	return pulumix.Output[*Addon]{
		OutputState: o.OutputState,
	}
}

// Resource name.
func (o AddonOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The properties of an addon resource
func (o AddonOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Addon) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Resource type.
func (o AddonOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AddonOutput{})
}
